--- vendor/cloud.google.com/go/storage/trace.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/trace.go
@@ -17,14 +17,6 @@ import (
 import (
 	"context"
 	"fmt"
-	"os"
-
-	internalTrace "cloud.google.com/go/internal/trace"
-	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	otelcodes "go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 const (
@@ -37,60 +29,13 @@ func isOTelTracingDevEnabled() bool {
 // isOTelTracingDevEnabled checks the development flag until experimental feature is launched.
 // TODO: Remove development flag upon experimental launch.
 func isOTelTracingDevEnabled() bool {
-	return os.Getenv(storageOtelTracingDevVar) == "true"
+	return false
 }
 
-func tracer() trace.Tracer {
-	return otel.Tracer(defaultTracerName, trace.WithInstrumentationVersion(internal.Version))
-}
-
-// startSpan creates a span and a context.Context containing the newly-created span.
-// If the context.Context provided in `ctx` contains a span then the newly-created
-// span will be a child of that span, otherwise it will be a root span.
-func startSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
-	name = appendPackageName(name)
-	// TODO: Remove internalTrace upon experimental launch.
-	if !isOTelTracingDevEnabled() {
-		ctx = internalTrace.StartSpan(ctx, name)
-		return ctx, nil
-	}
-	opts = append(opts, getCommonTraceOptions()...)
-	ctx, span := tracer().Start(ctx, name, opts...)
-	return ctx, span
-}
-
 // endSpan retrieves the current span from ctx and completes the span.
 // If an error occurs, the error is recorded as an exception span event for this span,
 // and the span status is set in the form of a code and a description.
 func endSpan(ctx context.Context, err error) {
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
-		attribute.String("gcp.client.version", internal.Version),
-		attribute.String("gcp.client.repo", gcpClientRepo),
-		attribute.String("gcp.client.artifact", gcpClientArtifact),
-	}
 }
 
 func appendPackageName(spanName string) string {
