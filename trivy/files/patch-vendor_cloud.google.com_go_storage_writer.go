--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-11-02 21:24:35 UTC
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
@@ -456,9 +453,7 @@ func (w *Writer) markClosed(err error) error {
 	w.mu.Unlock()
 
 	if state := metricsStateFromContext(w.ctx); state != nil {
-		if state.metrics != nil && total > 0 {
-			state.metrics.requestBodySize.Record(w.ctx, total, metric.WithAttributes(attribute.String("rpc.system.name", state.getSystemName()), attribute.String("rpc.method", "WriteObject"), attribute.String("server.address", stripPort(state.getTarget()))))
-		}
+		_ = total
 		if state.record != nil {
 			state.record(closingErr)
 		}
