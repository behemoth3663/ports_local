--- vendor/cloud.google.com/go/storage/reader.go.orig	2026-03-13 06:41:11 UTC
+++ vendor/cloud.google.com/go/storage/reader.go
@@ -114,8 +114,6 @@ func (o *ObjectHandle) NewRangeReader(ctx context.Cont
 func (o *ObjectHandle) NewRangeReader(ctx context.Context, offset, length int64) (r *Reader, err error) {
 	// This span covers the life of the reader. It is closed via the context
 	// in Reader.Close.
-	ctx, _ = startSpan(ctx, "Object.Reader")
-	defer func() { endSpan(ctx, err) }()
 
 	if err := o.validate(); err != nil {
 		return nil, err
@@ -234,13 +232,6 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
 func (o *ObjectHandle) NewMultiRangeDownloader(ctx context.Context, opts ...MRDOption) (mrd *MultiRangeDownloader, err error) {
 	// This span covers the life of the MRD. It is closed via the context
 	// in MultiRangeDownloader.Close.
-	var spanCtx context.Context
-	spanCtx, _ = startSpan(ctx, "Object.MultiRangeDownloader")
-	defer func() {
-		if err != nil {
-			endSpan(spanCtx, err)
-		}
-	}()
 
 	if err := o.validate(); err != nil {
 		return nil, err
@@ -267,7 +258,7 @@ func (o *ObjectHandle) NewMultiRangeDownloader(ctx con
 	}
 
 	// This call will return the *MultiRangeDownloader with the .impl field set.
-	return o.c.tc.NewMultiRangeDownloader(spanCtx, params, storageOpts...)
+	return o.c.tc.NewMultiRangeDownloader(ctx, params, storageOpts...)
 }
 
 // decompressiveTranscoding returns true if the request was served decompressed
@@ -353,7 +344,6 @@ func (r *Reader) Close() error {
 // Close closes the Reader. It must be called when done reading.
 func (r *Reader) Close() error {
 	err := r.reader.Close()
-	endSpan(r.ctx, err)
 	return err
 }
 
@@ -501,7 +491,6 @@ func (mrd *MultiRangeDownloader) Close() error {
 // If the downloader is in a permanent error state, this will return an error.
 func (mrd *MultiRangeDownloader) Close() error {
 	err := mrd.impl.close(nil)
-	endSpan(mrd.impl.getSpanCtx(), err)
 	return err
 }
 
