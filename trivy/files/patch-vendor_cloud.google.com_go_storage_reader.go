--- vendor/cloud.google.com/go/storage/reader.go.orig	2026-10-02 12:18:10 UTC
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
@@ -413,11 +410,6 @@ func (r *Reader) Close() error {
 	r.mu.Unlock()
 
 	if r.metricsState != nil {
-		if r.metricsState.metrics != nil {
-			if total := atomic.SwapInt64(&r.bytesRead, 0); total > 0 {
-				r.metricsState.metrics.responseBodySize.Record(r.ctx, total, metric.WithAttributes(attribute.String("rpc.method", "ReadObject")))
-			}
-		}
 		if r.metricsState.record != nil {
 			r.metricsState.record(err)
 		}
