diff --git a/vendor/cloud.google.com/go/storage/writer.go b/vendor/cloud.google.com/go/storage/writer.go
index d701ff22f..f6b6cea3e 100644
--- vendor/cloud.google.com/go/storage/writer.go.orig
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
