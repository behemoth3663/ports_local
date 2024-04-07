--- vendor/cloud.google.com/go/internal/trace/trace.go.orig	2023-11-29 17:44:14 UTC
+++ vendor/cloud.google.com/go/internal/trace/trace.go
@@ -22,10 +22,6 @@ import (
 	"strings"
 
 	"go.opencensus.io/trace"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	ottrace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/genproto/googleapis/rpc/code"
 	"google.golang.org/grpc/status"
@@ -68,7 +64,6 @@ func StartSpan(ctx context.Context, name string) conte
 // "opencensus" will be required to continue using OpenCensus tracing.
 func StartSpan(ctx context.Context, name string) context.Context {
 	if IsOpenTelemetryTracingEnabled() {
-		ctx, _ = otel.GetTracerProvider().Tracer(openTelemetryTracerName).Start(ctx, name)
 	} else {
 		ctx, _ = trace.StartSpan(ctx, name)
 	}
@@ -85,12 +80,6 @@ func EndSpan(ctx context.Context, err error) {
 // "opencensus" will be required to continue using OpenCensus tracing.
 func EndSpan(ctx context.Context, err error) {
 	if IsOpenTelemetryTracingEnabled() {
-		span := ottrace.SpanFromContext(ctx)
-		if err != nil {
-			span.SetStatus(codes.Error, toOpenTelemetryStatusDescription(err))
-			span.RecordError(err)
-		}
-		span.End()
 	} else {
 		span := trace.FromContext(ctx)
 		if err != nil {
@@ -171,8 +160,6 @@ func TracePrintf(ctx context.Context, attrMap map[stri
 // "opencensus" will be required to continue using OpenCensus tracing.
 func TracePrintf(ctx context.Context, attrMap map[string]interface{}, format string, args ...interface{}) {
 	if IsOpenTelemetryTracingEnabled() {
-		attrs := otAttrs(attrMap)
-		ottrace.SpanFromContext(ctx).AddEvent(fmt.Sprintf(format, args...), ottrace.WithAttributes(attrs...))
 	} else {
 		attrs := ocAttrs(attrMap)
 		// TODO: (odeke-em): perhaps just pass around spans due to the cost
@@ -198,28 +185,6 @@ func ocAttrs(attrMap map[string]interface{}) []trace.A
 			a = trace.Int64Attribute(k, v)
 		default:
 			a = trace.StringAttribute(k, fmt.Sprintf("%#v", v))
-		}
-		attrs = append(attrs, a)
-	}
-	return attrs
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
 		}
 		attrs = append(attrs, a)
 	}
