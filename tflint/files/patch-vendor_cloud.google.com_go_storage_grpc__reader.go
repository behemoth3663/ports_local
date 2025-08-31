--- vendor/cloud.google.com/go/storage/grpc_reader.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/grpc_reader.go
@@ -22,7 +22,6 @@ import (
 	"hash/crc32"
 	"io"
 
-	"cloud.google.com/go/internal/trace"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
 	"google.golang.org/grpc"
@@ -82,9 +81,6 @@ func (c *grpcStorageClient) NewRangeReaderReadObject(c
 }
 
 func (c *grpcStorageClient) NewRangeReaderReadObject(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReaderReadObject")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	s.gax = append(s.gax, gax.WithGRPCOptions(
