--- vendor/cloud.google.com/go/internal/trace/trace.go.orig	2025-02-20 19:18:58 UTC
+++ vendor/cloud.google.com/go/internal/trace/trace.go
@@ -17,12 +17,7 @@ import (
 import (
 	"context"
 	"errors"
-	"fmt"
 
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/grpc/status"
 )
@@ -36,7 +31,6 @@ func StartSpan(ctx context.Context, name string) conte
 // The default experimental tracing support for OpenCensus is now deprecated in
 // the Google Cloud client libraries for Go.
 func StartSpan(ctx context.Context, name string) context.Context {
-	ctx, _ = otel.GetTracerProvider().Tracer(OpenTelemetryTracerName).Start(ctx, name)
 	return ctx
 }
 
@@ -45,12 +39,6 @@ func EndSpan(ctx context.Context, err error) {
 // The default experimental tracing support for OpenCensus is now deprecated in
 // the Google Cloud client libraries for Go.
 func EndSpan(ctx context.Context, err error) {
-	span := trace.SpanFromContext(ctx)
-	if err != nil {
-		span.SetStatus(codes.Error, toOpenTelemetryStatusDescription(err))
-		span.RecordError(err)
-	}
-	span.End()
 }
 
 // toOpenTelemetryStatus converts an error to an equivalent OpenTelemetry status description.
@@ -70,28 +58,4 @@ func TracePrintf(ctx context.Context, attrMap map[stri
 // experimental tracing support for OpenCensus is now deprecated in the Google
 // Cloud client libraries for Go.
 func TracePrintf(ctx context.Context, attrMap map[string]interface{}, format string, args ...interface{}) {
-	attrs := otAttrs(attrMap)
-	trace.SpanFromContext(ctx).AddEvent(fmt.Sprintf(format, args...), trace.WithAttributes(attrs...))
-}
-
-// otAttrs converts a generic map to OpenTelemetry attributes.
-func otAttrs(attrMap map[string]interface{}) []attribute.KeyValue {
-	var attrs []attribute.KeyValue
-	for k, v := range attrMap {
-		var a attribute.KeyValue
-		switch v := v.(type) {
-		case string:
-			a = attribute.Key(k).String(v)
-		case bool:
-			a = attribute.Key(k).Bool(v)
-		case int:
-			a = attribute.Key(k).Int(v)
-		case int64:
-			a = attribute.Key(k).Int64(v)
-		default:
-			a = attribute.Key(k).String(fmt.Sprintf("%#v", v))
-		}
-		attrs = append(attrs, a)
-	}
-	return attrs
 }
