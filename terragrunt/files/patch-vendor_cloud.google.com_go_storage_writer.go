--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-10-22 19:52:07 UTC
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
@@ -271,9 +268,6 @@ func (w *Writer) isGRPCClient() bool {
 
 func (w *Writer) isGRPCClient() bool {
 	tc := w.o.c.tc
-	if mc, ok := tc.(*metricsStorageClient); ok {
-		tc = mc.storageClient
-	}
 	_, ok := tc.(*grpcStorageClient)
 	return ok
 }
@@ -452,13 +446,9 @@ func (w *Writer) markClosed(err error) error {
 		w.err = err
 	}
 	closingErr := w.err
-	total := atomic.LoadInt64(&w.bytesWritten)
 	w.mu.Unlock()
 
 	if state := metricsStateFromContext(w.ctx); state != nil {
-		if state.metrics != nil && total > 0 {
-			state.metrics.requestBodySize.Record(w.ctx, total, metric.WithAttributes(attribute.String("rpc.system.name", state.getSystemName()), attribute.String("rpc.method", "WriteObject"), attribute.String("server.address", stripPort(state.getTarget()))))
-		}
 		if state.record != nil {
 			state.record(closingErr)
 		}
