--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2024-10-09 14:40:23 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -21,12 +21,10 @@ import (
 	"fmt"
 	"hash/crc32"
 	"io"
-	"log"
 	"net/url"
 	"os"
 
 	"cloud.google.com/go/iam/apiv1/iampb"
-	"cloud.google.com/go/internal/trace"
 	gapic "cloud.google.com/go/storage/internal/apiv2"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
@@ -127,15 +125,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
 		return nil, errors.New("storage: GRPC is incompatible with any option that specifies an API for reads")
 	}
 
-	if !config.disableClientMetrics {
-		// Do not fail client creation if enabling metrics fails.
-		if metricsContext, err := enableClientMetrics(ctx, s); err == nil {
-			s.metricsContext = metricsContext
-			s.clientOption = append(s.clientOption, metricsContext.clientOpts...)
-		} else {
-			log.Printf("Failed to enable client metrics: %v", err)
-		}
-	}
 	g, err := gapic.NewClient(ctx, s.clientOption...)
 	if err != nil {
 		return nil, err
@@ -148,9 +137,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	return c.raw.Close()
 }
 
@@ -1004,9 +990,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 }
 
 func (c *grpcStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	s.gax = append(s.gax, gax.WithGRPCOptions(
