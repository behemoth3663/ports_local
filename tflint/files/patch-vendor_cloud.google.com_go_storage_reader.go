--- vendor/cloud.google.com/go/storage/reader.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/reader.go
@@ -24,8 +24,6 @@ import (
 	"strings"
 	"sync"
 	"time"
-
-	"cloud.google.com/go/internal/trace"
 )
 
 var crc32cTable = crc32.MakeTable(crc32.Castagnoli)
@@ -114,10 +112,6 @@ func (o *ObjectHandle) NewRangeReader(ctx context.Cont
 // operations, which all use JSON. JSON will become the default in a future
 // release.
 func (o *ObjectHandle) NewRangeReader(ctx context.Context, offset, length int64) (r *Reader, err error) {
-	// This span covers the life of the reader. It is closed via the context
-	// in Reader.Close.
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Object.Reader")
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
@@ -150,8 +144,6 @@ func (o *ObjectHandle) NewRangeReader(ctx context.Cont
 	// span now if there is an error.
 	if err == nil {
 		r.ctx = ctx
-	} else {
-		trace.EndSpan(ctx, err)
 	}
 
 	return r, err
@@ -165,10 +157,6 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
 // [experimental.WithGRPCBidiReads] or [experimental.WithZonalBucketAPIs]
 // must be selected in order to use this API.
 func (o *ObjectHandle) NewMultiRangeDownloader(ctx context.Context) (mrd *MultiRangeDownloader, err error) {
-	// This span covers the life of the reader. It is closed via the context
-	// in Reader.Close.
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Object.MultiRangeDownloader")
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
@@ -195,8 +183,6 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
 	// span now if there is an error.
 	if err == nil {
 		r.ctx = ctx
-	} else {
-		trace.EndSpan(ctx, err)
 	}
 
 	return r, err
@@ -284,7 +270,6 @@ func (r *Reader) Close() error {
 // Close closes the Reader. It must be called when done reading.
 func (r *Reader) Close() error {
 	err := r.reader.Close()
-	trace.EndSpan(r.ctx, err)
 	return err
 }
 
@@ -430,7 +415,6 @@ func (mrd *MultiRangeDownloader) Close() error {
 // Call [MultiRangeDownloader.Wait] to avoid this error.
 func (mrd *MultiRangeDownloader) Close() error {
 	err := mrd.reader.close()
-	trace.EndSpan(mrd.ctx, err)
 	return err
 }
 
