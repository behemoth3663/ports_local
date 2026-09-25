--- vendor/cloud.google.com/go/storage/reader.go.orig	2026-08-26 17:25:54 UTC
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
@@ -117,11 +114,6 @@ func (o *ObjectHandle) NewRangeReader(ctx context.Cont
 // operations, which all use JSON. JSON will become the default in a future
 // release.
 func (o *ObjectHandle) NewRangeReader(ctx context.Context, offset, length int64, opts ...ReaderOption) (r *Reader, err error) {
-	// This span covers the life of the reader. It is closed via the context
-	// in Reader.Close.
-	ctx, _ = startSpanWithBucket(ctx, o.c, o.bucket, "Object.Reader")
-	defer func() { endSpan(ctx, err) }()
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
@@ -152,9 +144,6 @@ func (o *ObjectHandle) NewRangeReader(ctx context.Cont
 	}
 
 	r, err = o.c.tc.NewRangeReader(ctx, params, storageOpts...)
-
-	// Pass the context so that the span can be closed in Reader.Close, or close the
-	// span now if there is an error.
 	if err == nil {
 		r.ctx = ctx
 	}
@@ -267,16 +256,6 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
 // NewMultiRangeDownloader creates a multi-range reader for an object.
 // Must be called on a gRPC client created using [NewGRPCClient].
 func (o *ObjectHandle) NewMultiRangeDownloader(ctx context.Context, opts ...MRDOption) (mrd *MultiRangeDownloader, err error) {
-	// This span covers the life of the MRD. It is closed via the context
-	// in MultiRangeDownloader.Close.
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
@@ -300,11 +279,8 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
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
@@ -415,14 +391,13 @@ func (r *Reader) Close() error {
 	if r.metricsState != nil {
 		if r.metricsState.metrics != nil {
 			if total := atomic.SwapInt64(&r.bytesRead, 0); total > 0 {
-				r.metricsState.metrics.responseBodySize.Record(r.ctx, total, metric.WithAttributes(attribute.String("rpc.method", "ReadObject")))
+				r.metricsState.metrics.responseBodySize.Record(r.ctx, total)
 			}
 		}
 		if r.metricsState.record != nil {
 			r.metricsState.record(err)
 		}
 	}
-	endSpan(r.ctx, err)
 	return err
 }
 
@@ -589,10 +564,6 @@ func (mrd *MultiRangeDownloader) Close() error {
 // it could lead to a deadlock.
 func (mrd *MultiRangeDownloader) Close() error {
 	err := mrd.impl.close(nil)
-	if state := metricsStateFromContext(mrd.impl.getSpanCtx()); state != nil && state.record != nil {
-		state.record(err)
-	}
-	endSpan(mrd.impl.getSpanCtx(), err)
 	return err
 }
 
