--- vendor/cloud.google.com/go/spanner/read.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/spanner/read.go
@@ -25,7 +25,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/internal/protostruct"
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/googleapis/gax-go/v2"
 	"google.golang.org/api/iterator"
@@ -92,7 +91,6 @@ func streamWithReplaceSessionFunc(
 	gsc *grpcSpannerClient,
 ) *RowIterator {
 	ctx, cancel := context.WithCancel(ctx)
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.RowIterator")
 	return &RowIterator{
 		meterTracerFactory:   meterTracerFactory,
 		streamd:              newResumableStreamDecoder(ctx, logger, rpc, replaceSession, gsc),
@@ -281,13 +279,6 @@ func (r *RowIterator) Stop() {
 // Stop terminates the iteration. It should be called after you finish using the
 // iterator.
 func (r *RowIterator) Stop() {
-	if r.streamd != nil {
-		if r.err != nil && r.err != iterator.Done {
-			defer trace.EndSpan(r.streamd.ctx, r.err)
-		} else {
-			defer trace.EndSpan(r.streamd.ctx, nil)
-		}
-	}
 	if r.cancel != nil {
 		r.cancel()
 	}
@@ -581,7 +572,6 @@ func (d *resumableStreamDecoder) next(mt *builtinMetri
 				d.changeState(aborted)
 				continue
 			}
-			trace.TracePrintf(d.ctx, nil, "Backing off stream read for %s", delay)
 			if err := gax.Sleep(d.ctx, delay); err == nil {
 				// record the attempt completion
 				mt.currOp.currAttempt.setStatus(status.Code(d.err).String())
