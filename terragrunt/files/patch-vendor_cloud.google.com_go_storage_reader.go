--- vendor/cloud.google.com/go/storage/reader.go.orig	2026-11-15 13:51:30 UTC
+++ vendor/cloud.google.com/go/storage/reader.go
@@ -26,8 +26,6 @@ import (
 	"sync/atomic"
 	"time"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 )
 
 var crc32cTable = crc32.MakeTable(crc32.Castagnoli)
@@ -410,7 +408,7 @@ func (r *Reader) Close() error {
 	if r.metricsState != nil {
 		if r.metricsState.metrics != nil {
 			if total := atomic.SwapInt64(&r.bytesRead, 0); total > 0 {
-				r.metricsState.metrics.responseBodySize.Record(r.ctx, total, metric.WithAttributes(attribute.String("rpc.system.name", r.metricsState.getSystemName()), attribute.String("rpc.method", "ReadObject"), attribute.String("server.address", stripPort(r.metricsState.getTarget()))))
+				r.metricsState.metrics.responseBodySize.Record(r.ctx, total)
 			}
 		}
 		if r.metricsState.record != nil {
@@ -583,7 +581,7 @@ func (mrd *MultiRangeDownloader) Close() error {
 	if state := metricsStateFromContext(mrd.impl.getSpanCtx()); state != nil {
 		if state.metrics != nil {
 			if total := mrd.impl.getBytesRead(); total > 0 {
-				state.metrics.responseBodySize.Record(mrd.impl.getSpanCtx(), total, metric.WithAttributes(attribute.String("rpc.system.name", state.getSystemName()), attribute.String("rpc.method", "ReadObject"), attribute.String("server.address", stripPort(state.getTarget()))))
+				state.metrics.responseBodySize.Record(mrd.impl.getSpanCtx(), total)
 			}
 		}
 		if state.record != nil {
