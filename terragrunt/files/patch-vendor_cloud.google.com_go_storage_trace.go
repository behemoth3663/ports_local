--- vendor/cloud.google.com/go/storage/trace.go.orig	2026-11-04 11:39:43 UTC
+++ vendor/cloud.google.com/go/storage/trace.go
@@ -16,21 +16,11 @@ import (
 
 import (
 	"context"
-	"errors"
 	"fmt"
-	"net/http"
 	"os"
 	"strings"
 
 	internalTrace "cloud.google.com/go/internal/trace"
-	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	otelcodes "go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
-	"google.golang.org/api/googleapi"
-	"google.golang.org/grpc/codes"
-	"google.golang.org/grpc/status"
 )
 
 type traceContextKey string
@@ -39,18 +29,8 @@ const (
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
 const (
 	defaultTracerName = "cloud.google.com/go/storage"
 	gcpClientArtifact = "cloud.google.com/go/storage"
@@ -67,11 +47,7 @@ func isACOEnabled() bool {
 	return os.Getenv(storageBucketMetadataDisabledVar) != "true"
 }
 
-func tracer() trace.Tracer {
-	return otel.Tracer(defaultTracerName, trace.WithInstrumentationVersion(internal.Version))
-}
-
-func startSpanWithBucket(ctx context.Context, client *Client, bucket string, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
+func startSpanWithBucket(ctx context.Context, client *Client, bucket string, name string, opts ...any) (context.Context, any) {
 	if !isOTelTracingDevEnabled() {
 		return startSpan(ctx, name, opts...)
 	}
@@ -93,11 +69,8 @@ func startSpanWithBucket(ctx context.Context, client *
 			cache.fetchBackground(ctx, bucket)
 			meta = placeholder
 		}
-		attrs := []attribute.KeyValue{
-			attribute.String("gcp.resource.destination.id", destinationResourceName(meta.resource)),
-			attribute.String("gcp.resource.destination.location", meta.location),
-		}
-		ctx = contextWithTraceAttributes(ctx, attrs)
+		_ = destinationResourceName(meta.resource)
+		_ = meta.location
 	}
 	return startSpan(ctx, name, opts...)
 }
@@ -122,76 +95,19 @@ func destinationResourceName(resource string) string {
 // startSpan creates a span and a context.Context containing the newly-created span.
 // If the context.Context provided in `ctx` contains a span then the newly-created
 // span will be a child of that span, otherwise it will be a root span.
-func startSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
+func startSpan(ctx context.Context, name string, opts ...any) (context.Context, any) {
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
+	_ = opts
+	ctx = internalTrace.StartSpan(ctx, name)
+	return ctx, nil
 }
 
-func isNotFoundError(err error) bool {
-	if errors.Is(err, ErrBucketNotExist) {
-		return true
-	}
-	var e *googleapi.Error
-	if s, ok := status.FromError(err); (ok && s.Code() == codes.NotFound) ||
-		(errors.As(err, &e) && e.Code == http.StatusNotFound) {
-		return true
-	}
-	return false
-}
-
 // endSpan retrieves the current span from ctx and completes the span.
 // If an error occurs, the error is recorded as an exception span event for this span,
 // and the span status is set in the form of a code and a description.
 func endSpan(ctx context.Context, err error) {
-	if err != nil && isNotFoundError(err) {
-		isBucket, _ := ctx.Value(isBucketContextKey).(bool)
-		cache, _ := ctx.Value(cacheContextKey).(*bucketMetadataCache)
-		bucket, _ := ctx.Value(bucketContextKey).(string)
-
-		if isBucket && cache != nil && bucket != "" {
-			cache.evict(bucket)
-		}
-	}
-
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
-}
-
-// getCommonTraceOptions makes a SpanStartOption with common attributes.
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
-		attribute.String("gcp.client.service", gcpClientService),
-		attribute.String("gcp.client.version", internal.Version),
-		attribute.String("gcp.client.artifact", gcpClientArtifact),
-	}
+	internalTrace.EndSpan(ctx, err)
 }
 
 func appendPackageName(spanName string) string {
