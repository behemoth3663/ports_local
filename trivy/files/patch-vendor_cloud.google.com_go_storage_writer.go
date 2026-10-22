--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-10-22 14:34:59 UTC
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
@@ -456,14 +453,10 @@ func (w *Writer) markClosed(err error) error {
 	w.mu.Unlock()
 
 	if state := metricsStateFromContext(w.ctx); state != nil {
-		if state.metrics != nil && total > 0 {
-			state.metrics.requestBodySize.Record(w.ctx, total, metric.WithAttributes(attribute.String("rpc.system.name", state.getSystemName()), attribute.String("rpc.method", "WriteObject"), attribute.String("server.address", stripPort(state.getTarget()))))
-		}
 		if state.record != nil {
 			state.record(closingErr)
 		}
 	}
-	endSpan(w.ctx, closingErr)
 	return closingErr
 }
 
