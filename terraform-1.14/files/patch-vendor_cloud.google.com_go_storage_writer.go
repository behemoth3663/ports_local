--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-08-26 17:25:54 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -24,9 +24,6 @@ import (
 	"sync/atomic"
 	"time"
 	"unicode/utf8"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 )
 
 // Interface internalWriter wraps low-level implementations which may vary
@@ -450,13 +447,12 @@ func (w *Writer) markClosed(err error) error {
 
 	if state := metricsStateFromContext(w.ctx); state != nil {
 		if state.metrics != nil && total > 0 {
-			state.metrics.requestBodySize.Record(w.ctx, total, metric.WithAttributes(attribute.String("rpc.method", "WriteObject")))
+			state.metrics.requestBodySize.Record(w.ctx, total)
 		}
 		if state.record != nil {
 			state.record(closingErr)
 		}
 	}
-	endSpan(w.ctx, closingErr)
 	return closingErr
 }
 
