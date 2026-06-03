--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2026-03-23 17:47:48 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -34,9 +34,6 @@ import (
 	"github.com/googleapis/gax-go/v2"
 	"github.com/googleapis/gax-go/v2/callctx"
 	"github.com/googleapis/gax-go/v2/internallog"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
 	grpccreds "google.golang.org/grpc/credentials"
@@ -459,22 +456,7 @@ func addOpenTelemetryStatsHandler(dialOpts []grpc.Dial
 }
 
 func addOpenTelemetryStatsHandler(dialOpts []grpc.DialOption, opts *Options) []grpc.DialOption {
-	if opts.DisableTelemetry {
 		return dialOpts
-	}
-	if !gax.IsFeatureEnabled("TRACING") {
-		return append(dialOpts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
-	}
-	var staticAttrs []attribute.KeyValue
-	if opts.InternalOptions != nil {
-		staticAttrs = transport.StaticTelemetryAttributes(opts.InternalOptions.TelemetryAttributes)
-	}
-	otelOpts := []otelgrpc.Option{
-		otelgrpc.WithSpanAttributes(staticAttrs...),
-	}
-	return append(dialOpts, grpc.WithStatsHandler(&otelHandler{
-		Handler: otelgrpc.NewClientHandler(otelOpts...),
-	}))
 }
 
 // otelHandler is a wrapper around the OpenTelemetry gRPC client handler that
@@ -487,80 +469,12 @@ func (h *otelHandler) TagRPC(ctx context.Context, info
 // name and retry count from the outgoing context metadata and attach them to
 // the current span.
 func (h *otelHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
-	ctx = h.Handler.TagRPC(ctx, info)
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
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
-		return
-	}
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
-		h.Handler.HandleRPC(ctx, s)
-		return
-	}
-
-	var attrs []attribute.KeyValue
-	if end.Error != nil {
-		st, ok := status.FromError(end.Error)
-		rpcStatusCode := codeToCanonicalStr(st.Code())
-
-		var errorType string
-		// 1. Check if the local context expired or was cancelled. This is the only
-		// reliable way to distinguish a local client timeout from a server timeout
-		// because gRPC does not wrap context errors in its status.Error types.
-		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
-			errorType = "CLIENT_TIMEOUT"
-		} else if errors.Is(ctx.Err(), context.Canceled) {
-			errorType = "CLIENT_CANCELLED"
-		} else if !ok || st.Code() == codes.Unknown || st.Code() == codes.Internal {
-			// 2. If the error isn't a context breakdown and the gRPC framework
-			// doesn't "understand" it (returning ok=false or a generic catch-all
-			// bucket like Unknown/Internal), we "pack" the actual Go error type
-			// name into error.type (e.g., "*net.OpError"). This is per the error.type
-			// [spec](https://opentelemetry.io/docs/specs/semconv/registry/attributes/error/#error-type).
-			// "When error.type is set to a type (e.g., an exception type), its canonical
-			// class name identifying the type within the artifact SHOULD be used."
-			errorType = fmt.Sprintf("%T", end.Error)
-		} else {
-			// 3. Otherwise, it is a well-understood gRPC protocol error (e.g.,
-			// PERMISSION_DENIED) likely returned by the server.
-			errorType = rpcStatusCode
-		}
-
-		attrs = []attribute.KeyValue{
-			attribute.String("error.type", errorType),
-			attribute.String("status.message", st.Message()),
-			attribute.String("rpc.response.status_code", rpcStatusCode),
-			attribute.String("exception.type", fmt.Sprintf("%T", end.Error)),
-		}
-	} else {
-		attrs = []attribute.KeyValue{
-			attribute.String("rpc.response.status_code", "OK"),
-		}
-	}
-	span.SetAttributes(attrs...)
 	h.Handler.HandleRPC(ctx, s)
 }
 
