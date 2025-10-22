--- vendor/cloud.google.com/go/spanner/trace.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/trace.go
@@ -19,10 +19,6 @@ import (
 	"errors"
 	"fmt"
 
-	"cloud.google.com/go/spanner/internal"
-	"go.opentelemetry.io/otel"
-	otelcodes "go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/grpc/status"
 )
@@ -33,29 +29,10 @@ const (
 	gcpClientArtifact = "cloud.google.com/go/spanner"
 )
 
-func tracer() trace.Tracer {
-	return otel.Tracer(defaultTracerName, trace.WithInstrumentationVersion(internal.Version))
-}
-
-// startSpan creates a span and a context.Context containing the newly-created span.
-// If the context.Context provided in `ctx` contains a span then the newly-created
-// span will be a child of that span, otherwise it will be a root span.
-func startSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
-	name = prependPackageName(name)
-	ctx, span := tracer().Start(ctx, name, opts...)
-	return ctx, span
-}
-
 // endSpan retrieves the current span from ctx and completes the span.
 // If an error occurs, the error is recorded as an exception span event for this span,
 // and the span status is set in the form of a code and a description.
 func endSpan(ctx context.Context, err error) {
-	span := trace.SpanFromContext(ctx)
-	if err != nil {
-		span.SetStatus(otelcodes.Error, toOpenTelemetryStatusDescription(err))
-		span.RecordError(err)
-	}
-	span.End()
 }
 
 // toOpenTelemetryStatus converts an error to an equivalent OpenTelemetry status description.
