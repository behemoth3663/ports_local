--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-10-02 16:09:03 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -25,8 +25,6 @@ import (
 	"time"
 	"unicode/utf8"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 )
 
 // Interface internalWriter wraps low-level implementations which may vary
@@ -450,7 +448,7 @@ func (w *Writer) markClosed(err error) error {
 
 	if state := metricsStateFromContext(w.ctx); state != nil {
 		if state.metrics != nil && total > 0 {
-			state.metrics.requestBodySize.Record(w.ctx, total, metric.WithAttributes(attribute.String("rpc.method", "WriteObject")))
+			state.metrics.requestBodySize.Record(w.ctx, total)
 		}
 		if state.record != nil {
 			state.record(closingErr)
