diff --git a/vendor/cloud.google.com/go/storage/reader.go b/vendor/cloud.google.com/go/storage/reader.go
index de56f9f2a..c6a4b3e33 100644
--- vendor/cloud.google.com/go/storage/reader.go.orig
+++ vendor/cloud.google.com/go/storage/reader.go
@@ -26,8 +26,6 @@ import (
 	"sync/atomic"
 	"time"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 )
 
 var crc32cTable = crc32.MakeTable(crc32.Castagnoli)
@@ -415,7 +413,7 @@ func (r *Reader) Close() error {
 	if r.metricsState != nil {
 		if r.metricsState.metrics != nil {
 			if total := atomic.SwapInt64(&r.bytesRead, 0); total > 0 {
-				r.metricsState.metrics.responseBodySize.Record(r.ctx, total, metric.WithAttributes(attribute.String("rpc.method", "ReadObject")))
+				r.metricsState.metrics.responseBodySize.Record(r.ctx, total)
 			}
 		}
 		if r.metricsState.record != nil {
