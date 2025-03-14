--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2025-02-06 22:11:02 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -21,13 +21,11 @@ import (
 	"fmt"
 	"hash/crc32"
 	"io"
-	"log"
 	"net/url"
 	"os"
 	"sync"
 
 	"cloud.google.com/go/iam/apiv1/iampb"
-	"cloud.google.com/go/internal/trace"
 	gapic "cloud.google.com/go/storage/internal/apiv2"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
@@ -35,7 +33,6 @@ import (
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
-	"google.golang.org/api/transport"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/encoding"
@@ -119,24 +116,6 @@ type grpcStorageClient struct {
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
@@ -150,15 +129,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
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
@@ -172,9 +142,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	return c.raw.Close()
 }
 
@@ -1071,8 +1038,6 @@ func (c *grpcStorageClient) NewMultiRangeDownloader(ct
 }
 
 func (c *grpcStorageClient) NewMultiRangeDownloader(ctx context.Context, params *newMultiRangeDownloaderParams, opts ...storageOption) (mr *MultiRangeDownloader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewMultiRangeDownloader")
-	defer func() { trace.EndSpan(ctx, err) }()
 	s := callSettings(c.settings, opts...)
 
 	if s.userProject != "" {
@@ -1491,9 +1456,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 	if !c.config.grpcBidiReads {
 		return c.NewRangeReaderReadObject(ctx, params, opts...)
 	}
-
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
