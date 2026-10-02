--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-10-02 12:18:10 UTC
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
@@ -449,9 +446,7 @@ func (w *Writer) markClosed(err error) error {
 	w.mu.Unlock()
 
 	if state := metricsStateFromContext(w.ctx); state != nil {
-		if state.metrics != nil && total > 0 {
-			state.metrics.requestBodySize.Record(w.ctx, total, metric.WithAttributes(attribute.String("rpc.method", "WriteObject")))
-		}
+		_ = total
 		if state.record != nil {
 			state.record(closingErr)
 		}
