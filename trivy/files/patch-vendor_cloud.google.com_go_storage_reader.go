--- vendor/cloud.google.com/go/storage/reader.go.orig	2026-10-22 14:31:43 UTC
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
@@ -119,9 +116,6 @@ func (o *ObjectHandle) NewRangeReader(ctx context.Cont
 func (o *ObjectHandle) NewRangeReader(ctx context.Context, offset, length int64, opts ...ReaderOption) (r *Reader, err error) {
 	// This span covers the life of the reader. It is closed via the context
 	// in Reader.Close.
-	ctx, _ = startSpanWithBucket(ctx, o.c, o.bucket, "Object.Reader")
-	defer func() { endSpan(ctx, err) }()
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
@@ -264,14 +258,6 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
 func (o *ObjectHandle) NewMultiRangeDownloader(ctx context.Context, opts ...MRDOption) (mrd *MultiRangeDownloader, err error) {
 	// This span covers the life of the MRD. It is closed via the context
 	// in MultiRangeDownloader.Close.
-	var spanCtx context.Context
-	spanCtx, _ = startSpanWithBucket(ctx, o.c, o.bucket, "Object.MultiRangeDownloader")
-	defer func() {
-		if err != nil {
-			endSpan(spanCtx, err)
-		}
-	}()
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
@@ -295,11 +281,8 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
 	for _, opt := range opts {
 		opt.apply(params)
 	}
-	if params.minConnections > 1 || params.maxConnections > 1 {
-		spanCtx = addFeatureAttributes(spanCtx, featureMultistreamInMRD)
-	}
 	// This call will return the *MultiRangeDownloader with the .impl field set.
-	return o.c.tc.NewMultiRangeDownloader(spanCtx, params, storageOpts...)
+	return o.c.tc.NewMultiRangeDownloader(ctx, params, storageOpts...)
 }
 
 // decompressiveTranscoding returns true if the request was served decompressed
@@ -408,16 +391,10 @@ func (r *Reader) Close() error {
 	r.mu.Unlock()
 
 	if r.metricsState != nil {
-		if r.metricsState.metrics != nil {
-			if total := atomic.SwapInt64(&r.bytesRead, 0); total > 0 {
-				r.metricsState.metrics.responseBodySize.Record(r.ctx, total, metric.WithAttributes(attribute.String("rpc.system.name", r.metricsState.getSystemName()), attribute.String("rpc.method", "ReadObject"), attribute.String("server.address", stripPort(r.metricsState.getTarget()))))
-			}
-		}
 		if r.metricsState.record != nil {
 			r.metricsState.record(err)
 		}
 	}
-	endSpan(r.ctx, err)
 	return err
 }
 
@@ -580,17 +557,6 @@ func (mrd *MultiRangeDownloader) Close() error {
 // it could lead to a deadlock.
 func (mrd *MultiRangeDownloader) Close() error {
 	err := mrd.impl.close(nil)
-	if state := metricsStateFromContext(mrd.impl.getSpanCtx()); state != nil {
-		if state.metrics != nil {
-			if total := mrd.impl.getBytesRead(); total > 0 {
-				state.metrics.responseBodySize.Record(mrd.impl.getSpanCtx(), total, metric.WithAttributes(attribute.String("rpc.system.name", state.getSystemName()), attribute.String("rpc.method", "ReadObject"), attribute.String("server.address", stripPort(state.getTarget()))))
-			}
-		}
-		if state.record != nil {
-			state.record(err)
-		}
-	}
-	endSpan(mrd.impl.getSpanCtx(), err)
 	return err
 }
 
