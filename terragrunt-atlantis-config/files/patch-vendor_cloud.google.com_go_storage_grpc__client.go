--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2025-01-02 18:25:02 UTC
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
 	"github.com/google/uuid"
@@ -35,7 +33,6 @@ import (
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
-	"google.golang.org/api/transport"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/encoding"
@@ -118,23 +115,6 @@ type grpcStorageClient struct {
 	settings *settings
 }
 
-func enableClientMetrics(ctx context.Context, s *settings, config storageConfig) (*metricsContext, error) {
-	var project string
-	// TODO: use new auth client
-	c, err := transport.Creds(ctx, s.clientOption...)
-	if err == nil {
-		project = c.ProjectID
-	}
-	metricsContext, err := newGRPCMetricContext(ctx, metricsConfig{
-		project:      project,
-		interval:     config.metricInterval,
-		manualReader: config.manualReader},
-	)
-	if err != nil {
-		return nil, fmt.Errorf("gRPC Metrics: %w", err)
-	}
-	return metricsContext, nil
-}
 
 // newGRPCStorageClient initializes a new storageClient that uses the gRPC
 // Storage API.
@@ -149,15 +129,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
 		return nil, errors.New("storage: GRPC is incompatible with any option that specifies an API for reads")
 	}
 
-	if !config.disableClientMetrics {
-		// Do not fail client creation if enabling metrics fails.
-		if metricsContext, err := enableClientMetrics(ctx, s, config); err == nil {
-			s.metricsContext = metricsContext
-			s.clientOption = append(s.clientOption, metricsContext.clientOpts...)
-		} else {
-			log.Printf("Failed to enable client metrics: %v", err)
-		}
-	}
 	g, err := gapic.NewClient(ctx, s.clientOption...)
 	if err != nil {
 		return nil, err
@@ -170,9 +141,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	return c.raw.Close()
 }
 
@@ -1056,9 +1024,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 }
 
 func (c *grpcStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	s.gax = append(s.gax, gax.WithGRPCOptions(
