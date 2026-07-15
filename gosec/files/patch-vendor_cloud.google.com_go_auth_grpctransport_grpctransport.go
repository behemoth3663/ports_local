--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2026-07-07 20:38:34 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -36,9 +36,6 @@ import (
 	"github.com/googleapis/gax-go/v2"
 	"github.com/googleapis/gax-go/v2/callctx"
 	"github.com/googleapis/gax-go/v2/internallog"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/grpc"
 	grpccreds "google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
@@ -371,7 +368,6 @@ func dial(ctx context.Context, secure bool, opts *Opti
 	// Add tracing, but before the other options, so that clients can override the
 	// gRPC stats handler.
 	// This assumes that gRPC options are processed in order, left to right.
-	grpcOpts = addOpenTelemetryStatsHandler(grpcOpts, opts, transportCreds.Endpoint)
 	grpcOpts = append(grpcOpts, opts.GRPCDialOpts...)
 
 	return grpc.DialContext(ctx, transportCreds.Endpoint, grpcOpts...)
@@ -461,41 +457,7 @@ func addOpenTelemetryStatsHandler(dialOpts []grpc.Dial
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
@@ -550,29 +512,6 @@ func (h *otelHandler) TagRPC(ctx context.Context, info
 // the current span.
 func (h *otelHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
 	ctx = h.Handler.TagRPC(ctx, info)
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
 
