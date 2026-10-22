--- vendor/cloud.google.com/go/storage/reader.go.orig	2026-10-22 19:52:07 UTC
+++ vendor/cloud.google.com/go/storage/reader.go
@@ -25,9 +25,6 @@ import (
 	"sync"
 	"sync/atomic"
 	"time"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 )
 
 var crc32cTable = crc32.MakeTable(crc32.Castagnoli)
@@ -408,11 +405,7 @@ func (r *Reader) Close() error {
 	r.mu.Unlock()
 
 	if r.metricsState != nil {
-		if r.metricsState.metrics != nil {
-			if total := atomic.SwapInt64(&r.bytesRead, 0); total > 0 {
-				r.metricsState.metrics.responseBodySize.Record(r.ctx, total, metric.WithAttributes(attribute.String("rpc.system.name", r.metricsState.getSystemName()), attribute.String("rpc.method", "ReadObject"), attribute.String("server.address", stripPort(r.metricsState.getTarget()))))
-			}
-		}
+		_ = atomic.SwapInt64(&r.bytesRead, 0)
 		if r.metricsState.record != nil {
 			r.metricsState.record(err)
 		}
@@ -581,11 +574,7 @@ func (mrd *MultiRangeDownloader) Close() error {
 func (mrd *MultiRangeDownloader) Close() error {
 	err := mrd.impl.close(nil)
 	if state := metricsStateFromContext(mrd.impl.getSpanCtx()); state != nil {
-		if state.metrics != nil {
-			if total := mrd.impl.getBytesRead(); total > 0 {
-				state.metrics.responseBodySize.Record(mrd.impl.getSpanCtx(), total, metric.WithAttributes(attribute.String("rpc.system.name", state.getSystemName()), attribute.String("rpc.method", "ReadObject"), attribute.String("server.address", stripPort(state.getTarget()))))
-			}
-		}
+		_ = mrd.impl.getBytesRead()
 		if state.record != nil {
 			state.record(err)
 		}
