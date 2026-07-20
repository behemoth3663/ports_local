--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2026-04-06 22:09:02 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -33,12 +33,7 @@ import (
 	"cloud.google.com/go/auth/internal"
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/headers"
-	"github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"github.com/googleapis/gax-go/v2/internallog"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/grpc"
 	grpccreds "google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
@@ -459,41 +454,7 @@ func addOpenTelemetryStatsHandler(dialOpts []grpc.Dial
 }
 
 func addOpenTelemetryStatsHandler(dialOpts []grpc.DialOption, opts *Options, endpoint string) []grpc.DialOption {
-	if opts.DisableTelemetry {
 		return dialOpts
-	}
-	if gax.IsFeatureEnabled("METRICS") {
-		host, port := extractHostPort(endpoint)
-		dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(openTelemetryUnaryClientInterceptor(host, port)))
-	}
-	if !gax.IsFeatureEnabled("TRACING") && !gax.IsFeatureEnabled("LOGGING") {
-		return append(dialOpts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
-	}
-	var staticAttrs []attribute.KeyValue
-	var scopedLogger *slog.Logger
-
-	if gax.IsFeatureEnabled("LOGGING") && opts.Logger != nil {
-		scopedLogger = opts.Logger
-	}
-
-	if opts.InternalOptions != nil {
-		staticAttrs = transport.StaticTelemetryAttributes(opts.InternalOptions.TelemetryAttributes)
-		if scopedLogger != nil {
-			var staticLogAttrs []any
-			for _, attr := range staticAttrs {
-				staticLogAttrs = append(staticLogAttrs, slog.String(string(attr.Key), attr.Value.AsString()))
-			}
-			scopedLogger = scopedLogger.With(staticLogAttrs...)
-		}
-	}
-	var otelOpts []otelgrpc.Option
-	if gax.IsFeatureEnabled("TRACING") {
-		otelOpts = append(otelOpts, otelgrpc.WithSpanAttributes(staticAttrs...))
-	}
-	return append(dialOpts, grpc.WithStatsHandler(&otelHandler{
-		Handler: otelgrpc.NewClientHandler(otelOpts...),
-		logger:  scopedLogger,
-	}))
 }
 
 // Extract the host and port from a target address
@@ -520,16 +481,6 @@ func openTelemetryUnaryClientInterceptor(host string, 
 // TransportTelemetryData with the server peer address.
 func openTelemetryUnaryClientInterceptor(host string, port int) grpc.UnaryClientInterceptor {
 	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
-		transportData := gax.ExtractTransportTelemetry(ctx)
-		if transportData != nil {
-			if host != "" {
-				transportData.SetServerAddress(host)
-			}
-			if port != 0 {
-				transportData.SetServerPort(port)
-			}
-		}
-
 		err := invoker(ctx, method, req, reply, cc, opts...)
 
 		return err
@@ -547,129 +498,12 @@ func (h *otelHandler) TagRPC(ctx context.Context, info
 // name and retry count from the outgoing context metadata and attach them to
 // the current span.
 func (h *otelHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
-	ctx = h.Handler.TagRPC(ctx, info)
-
-	var span trace.Span
-	if gax.IsFeatureEnabled("TRACING") {
-		if s := trace.SpanFromContext(ctx); s != nil && s.IsRecording() {
-			span = s
-		}
-	}
-
-	if span == nil {
-		return ctx
-	}
-	var attrs []attribute.KeyValue
-	if resName, ok := callctx.TelemetryFromContext(ctx, "resource_name"); ok {
-		attrs = append(attrs, attribute.String("gcp.resource.destination.id", resName))
-	}
-	if resendCountStr, ok := callctx.TelemetryFromContext(ctx, "resend_count"); ok {
-		if count, err := strconv.Atoi(resendCountStr); err == nil {
-			attrs = append(attrs, attribute.Int("gcp.grpc.resend_count", count))
-		}
-	}
-	if len(attrs) > 0 {
-		span.SetAttributes(attrs...)
-	}
 	return ctx
 }
 
 // HandleRPC intercepts the RPC completion to capture and format error-related
 // attributes ensuring they conform to Google Cloud observability standards.
 func (h *otelHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {
-	end, ok := s.(*stats.End)
-	if !ok {
-		h.Handler.HandleRPC(ctx, s)
 		return
-	}
-
-	var span trace.Span
-	if gax.IsFeatureEnabled("TRACING") {
-		if s := trace.SpanFromContext(ctx); s != nil && s.IsRecording() {
-			span = s
-		}
-	}
-
-	var logger *slog.Logger
-	if gax.IsFeatureEnabled("LOGGING") {
-		if l := h.logger; l != nil && l.Enabled(ctx, slog.LevelDebug) {
-			logger = l
-		}
-	}
-
-	if span == nil && logger == nil {
-		h.Handler.HandleRPC(ctx, s)
-		return
-	}
-
-	if end.Error != nil {
-		h.handleRPCError(ctx, span, logger, end)
-	} else {
-		h.handleRPCSuccess(span)
-	}
-	h.Handler.HandleRPC(ctx, s)
 }
 
-func (h *otelHandler) handleRPCError(ctx context.Context, span trace.Span, logger *slog.Logger, end *stats.End) {
-	info := gax.ExtractTelemetryErrorInfo(ctx, end.Error)
-
-	if logger != nil {
-		logActionableError(ctx, logger, info)
-	}
-
-	if span != nil {
-		attrs := []attribute.KeyValue{
-			attribute.String("error.type", info.ErrorType),
-			attribute.String("status.message", info.StatusMessage),
-			attribute.String("rpc.response.status_code", info.StatusCode),
-			attribute.String("exception.type", fmt.Sprintf("%T", end.Error)),
-		}
-		span.SetAttributes(attrs...)
-	}
-}
-
-func (h *otelHandler) handleRPCSuccess(span trace.Span) {
-	if span != nil {
-		attrs := []attribute.KeyValue{
-			attribute.String("rpc.response.status_code", "OK"),
-		}
-		span.SetAttributes(attrs...)
-	}
-}
-
-func logActionableError(ctx context.Context, logger *slog.Logger, info gax.TelemetryErrorInfo) {
-	baseLogAttrs := []slog.Attr{
-		slog.String("rpc.system.name", "grpc"),
-		slog.String("rpc.response.status_code", info.StatusCode),
-	}
-
-	if resendCountStr, ok := callctx.TelemetryFromContext(ctx, "resend_count"); ok {
-		if count, e := strconv.Atoi(resendCountStr); e == nil {
-			baseLogAttrs = append(baseLogAttrs, slog.Int64("gcp.grpc.resend_count", int64(count)))
-		}
-	}
-
-	msg := info.StatusMessage
-	if msg == "" {
-		msg = "API call failed"
-	}
-
-	if rpcMethod, ok := callctx.TelemetryFromContext(ctx, "rpc_method"); ok {
-		baseLogAttrs = append(baseLogAttrs, slog.String("rpc.method", rpcMethod))
-	}
-
-	if resName, ok := callctx.TelemetryFromContext(ctx, "resource_name"); ok {
-		baseLogAttrs = append(baseLogAttrs, slog.String("gcp.resource.destination.id", resName))
-	}
-
-	if info.Domain != "" {
-		baseLogAttrs = append(baseLogAttrs, slog.String("gcp.errors.domain", info.Domain))
-	}
-	for k, v := range info.Metadata {
-		baseLogAttrs = append(baseLogAttrs, slog.String("gcp.errors.metadata."+k, v))
-	}
-
-	baseLogAttrs = append(baseLogAttrs, slog.String("error.type", info.ErrorType))
-
-	logger.LogAttrs(ctx, slog.LevelDebug, msg, baseLogAttrs...)
-}
