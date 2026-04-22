--- vendor/cloud.google.com/go/storage/grpc_reader.go.orig	2026-03-13 06:41:11 UTC
+++ vendor/cloud.google.com/go/storage/grpc_reader.go
@@ -82,8 +82,6 @@ func (c *grpcStorageClient) NewRangeReaderReadObject(c
 
 // NewRangeReaderReadObject is the legacy (non-bidi) implementation of reads.
 func (c *grpcStorageClient) NewRangeReaderReadObject(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx, _ = startSpan(ctx, "grpcStorageClient.NewRangeReaderReadObject")
-	defer func() { endSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
