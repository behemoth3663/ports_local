--- vendor/cloud.google.com/go/spanner/read.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/read.go
@@ -25,11 +25,8 @@ import (
 	"time"
 
 	"cloud.google.com/go/internal/protostruct"
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/googleapis/gax-go/v2"
-	otcodes "go.opentelemetry.io/otel/codes"
-	otrace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/status"
@@ -95,7 +92,6 @@ func streamWithReplaceSessionFunc(
 	gsc *grpcSpannerClient,
 ) *RowIterator {
 	ctx, cancel := context.WithCancel(ctx)
-	ctx, _ = startSpan(ctx, "RowIterator")
 	return &RowIterator{
 		meterTracerFactory:   meterTracerFactory,
 		streamd:              newResumableStreamDecoder(ctx, cancel, logger, rpc, replaceSession, gsc),
@@ -285,13 +281,6 @@ func (r *RowIterator) Stop() {
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
@@ -589,7 +578,6 @@ func (d *resumableStreamDecoder) next(mt *builtinMetri
 				d.changeState(aborted)
 				continue
 			}
-			trace.TracePrintf(d.ctx, nil, "Backing off stream read for %s", delay)
 			if err := gax.Sleep(d.ctx, delay); err == nil {
 				// record the attempt completion
 				mt.currOp.currAttempt.setStatus(status.Code(d.err).String())
@@ -680,10 +668,6 @@ func (d *resumableStreamDecoder) tryRecv(mt *builtinMe
 	if d.err == nil {
 		d.q.push(res)
 		if res.GetLast() {
-			if span := otrace.SpanFromContext(d.stream.Context()); span != nil && span.IsRecording() {
-				span.SetStatus(otcodes.Ok, "Stream finished successfully")
-				span.End()
-			}
 			if d.cancel != nil {
 				// Remove the cancel function to prevent iter.Stop from also calling it.
 				cancel := d.cancel
