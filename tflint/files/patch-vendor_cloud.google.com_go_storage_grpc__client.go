--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2025-02-20 19:18:58 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -21,7 +21,6 @@ import (
 	"fmt"
 	"hash/crc32"
 	"io"
-	"log"
 	"net/url"
 	"os"
 	"sync"
@@ -35,7 +34,6 @@ import (
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
-	"google.golang.org/api/transport"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/encoding"
@@ -119,24 +117,6 @@ type grpcStorageClient struct {
 	config   *storageConfig
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
-
 // newGRPCStorageClient initializes a new storageClient that uses the gRPC
 // Storage API.
 func newGRPCStorageClient(ctx context.Context, opts ...storageOption) (storageClient, error) {
@@ -150,15 +130,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
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
@@ -172,9 +143,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	return c.raw.Close()
 }
 
