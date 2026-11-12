--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-11-15 13:51:30 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -25,8 +25,6 @@ import (
 	"time"
 	"unicode/utf8"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 )
 
 // Interface internalWriter wraps low-level implementations which may vary
@@ -467,7 +465,7 @@ func (w *Writer) markClosed(err error) error {
 
 	if state := metricsStateFromContext(w.ctx); state != nil {
 		if state.metrics != nil && total > 0 {
-			state.metrics.requestBodySize.Record(w.ctx, total, metric.WithAttributes(attribute.String("rpc.system.name", state.getSystemName()), attribute.String("rpc.method", "WriteObject"), attribute.String("server.address", stripPort(state.getTarget()))))
+			state.metrics.requestBodySize.Record(w.ctx, total)
 		}
 		if state.record != nil {
 			state.record(closingErr)
