--- vendor/cloud.google.com/go/storage/trace.go.orig	2026-10-22 19:52:07 UTC
+++ vendor/cloud.google.com/go/storage/trace.go
@@ -23,57 +23,29 @@ import (
 	"strings"
 
 	internalTrace "cloud.google.com/go/internal/trace"
-	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	otelcodes "go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/status"
 )
 
 type traceContextKey string
+type apiMethodKey struct{}
 
 const (
 	cacheContextKey    traceContextKey = "bucketMetadataCache"
 	bucketContextKey   traceContextKey = "bucketName"
 	isBucketContextKey traceContextKey = "isBucketOp"
-	traceAttributesKey traceContextKey = "traceAttributes"
 )
 
-func contextWithTraceAttributes(ctx context.Context, attrs []attribute.KeyValue) context.Context {
-	return context.WithValue(ctx, traceAttributesKey, attrs)
-}
-
-func traceAttributesFromContext(ctx context.Context) ([]attribute.KeyValue, bool) {
-	attrs, ok := ctx.Value(traceAttributesKey).([]attribute.KeyValue)
-	return attrs, ok
-}
-
-const (
-	defaultTracerName = "cloud.google.com/go/storage"
-	gcpClientArtifact = "cloud.google.com/go/storage"
-)
-
 // isOTelTracingDevEnabled checks the development flag until experimental feature is launched.
 // TODO: Remove development flag upon experimental launch.
 func isOTelTracingDevEnabled() bool {
-	return os.Getenv(storageOtelTracingDevVar) == "true"
+	return false
 }
 
-func isACOEnabled() bool {
-	return os.Getenv(storageBucketMetadataDisabledVar) != "true"
-}
+func isACOEnabled() bool { return os.Getenv(storageBucketMetadataDisabledVar) != "true" }
 
-func tracer() trace.Tracer {
-	return otel.Tracer(defaultTracerName, trace.WithInstrumentationVersion(internal.Version))
-}
-
-func startSpanWithBucket(ctx context.Context, client *Client, bucket string, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
-	if !isOTelTracingDevEnabled() {
-		return startSpan(ctx, name, opts...)
-	}
+func startSpanWithBucket(ctx context.Context, client *Client, bucket string, name string, _ ...any) (context.Context, any) {
 	if client != nil && client.bucketMetadataCache != nil && bucket != "" {
 		ctx = context.WithValue(ctx, cacheContextKey, client.bucketMetadataCache)
 		ctx = context.WithValue(ctx, bucketContextKey, bucket)
@@ -90,34 +62,21 @@ func startSpanWithBucket(ctx context.Context, client *
 			}
 			cache.put(bucket, placeholder)
 			cache.fetchBackground(ctx, bucket)
-			meta = placeholder
+			_ = placeholder
 		}
-		attrs := []attribute.KeyValue{
-			attribute.String("gcp.resource.destination.id", meta.resource),
-			attribute.String("gcp.resource.destination.location", meta.location),
-		}
-		ctx = contextWithTraceAttributes(ctx, attrs)
+		_ = meta
 	}
-	return startSpan(ctx, name, opts...)
+	return startSpan(ctx, name)
 }
 
 // startSpan creates a span and a context.Context containing the newly-created span.
 // If the context.Context provided in `ctx` contains a span then the newly-created
 // span will be a child of that span, otherwise it will be a root span.
-func startSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
+func startSpan(ctx context.Context, name string, _ ...any) (context.Context, any) {
 	ctx = context.WithValue(ctx, apiMethodKey{}, name)
 	name = appendPackageName(name)
-	// TODO: Remove internalTrace upon experimental launch.
-	if !isOTelTracingDevEnabled() {
-		ctx = internalTrace.StartSpan(ctx, name)
-		return ctx, nil
-	}
-	opts = append(opts, getCommonTraceOptions()...)
-	if attrs, ok := traceAttributesFromContext(ctx); ok {
-		opts = append(opts, trace.WithAttributes(attrs...))
-	}
-	ctx, span := tracer().Start(ctx, name, opts...)
-	return ctx, span
+	ctx = internalTrace.StartSpan(ctx, name)
+	return ctx, nil
 }
 
 func isNotFoundError(err error) bool {
@@ -146,35 +105,10 @@ func endSpan(ctx context.Context, err error) {
 		}
 	}
 
-	// TODO: Remove internalTrace upon experimental launch.
-	if !isOTelTracingDevEnabled() {
-		internalTrace.EndSpan(ctx, err)
-	} else {
-		span := trace.SpanFromContext(ctx)
-		if err != nil {
-			span.SetStatus(otelcodes.Error, err.Error())
-			span.RecordError(err)
-		}
-		span.End()
-	}
+	internalTrace.EndSpan(ctx, err)
 }
 
 // getCommonTraceOptions makes a SpanStartOption with common attributes.
-func getCommonTraceOptions() []trace.SpanStartOption {
-	opts := []trace.SpanStartOption{
-		trace.WithAttributes(getCommonAttributes()...),
-	}
-	return opts
-}
-
-// getCommonAttributes includes the common attributes used for Cloud Trace adoption tracking.
-func getCommonAttributes() []attribute.KeyValue {
-	return []attribute.KeyValue{
-		attribute.String("gcp.client.version", internal.Version),
-		attribute.String("gcp.client.artifact", gcpClientArtifact),
-	}
-}
-
 func appendPackageName(spanName string) string {
-	return fmt.Sprintf("%s.%s", gcpClientArtifact, spanName)
+	return fmt.Sprintf("cloud.google.com/go/storage.%s", spanName)
 }
