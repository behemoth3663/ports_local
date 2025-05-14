--- telemetry/context.go.orig	2025-05-13 18:23:29 UTC
+++ telemetry/context.go
@@ -2,9 +2,6 @@ import (
 
 import (
 	"context"
-	"fmt"
-
-	"go.opentelemetry.io/otel/trace"
 )
 
 type contextKey byte
@@ -32,24 +29,5 @@ func TraceParentFromContext(ctx context.Context, telem
 
 // TraceParentFromContext returns the W3C traceparent header value from the context's span, or an error if not available.
 func TraceParentFromContext(ctx context.Context, telemetry *Options) string {
-	span := trace.SpanFromContext(ctx)
-	spanContext := span.SpanContext()
-
-	if !spanContext.IsValid() {
-		return ""
-	}
-
-	if len(telemetry.TraceParent) > 0 {
 		return telemetry.TraceParent
-	}
-
-	traceID := spanContext.TraceID().String()
-	spanID := spanContext.SpanID().String()
-	flags := "00"
-
-	if spanContext.TraceFlags().IsSampled() {
-		flags = "01"
-	}
-
-	return fmt.Sprintf("00-%s-%s-%s", traceID, spanID, flags)
 }
