--- vendor/cloud.google.com/go/spanner/read.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/read.go
@@ -28,8 +28,6 @@ import (
 	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/googleapis/gax-go/v2"
-	otcodes "go.opentelemetry.io/otel/codes"
-	otrace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/status"
@@ -95,7 +93,6 @@ func streamWithReplaceSessionFunc(
 	gsc *grpcSpannerClient,
 ) *RowIterator {
 	ctx, cancel := context.WithCancel(ctx)
-	ctx, _ = startSpan(ctx, "RowIterator")
 	return &RowIterator{
 		meterTracerFactory:   meterTracerFactory,
 		streamd:              newResumableStreamDecoder(ctx, cancel, logger, rpc, replaceSession, gsc),
@@ -680,10 +677,6 @@ func (d *resumableStreamDecoder) tryRecv(mt *builtinMe
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
