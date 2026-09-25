--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2026-08-26 17:25:54 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -21,7 +21,6 @@ import (
 	"fmt"
 	"hash/crc32"
 	"io"
-	"log"
 	"os"
 	"strconv"
 	"strings"
@@ -134,26 +133,6 @@ type grpcStorageClient struct {
 	configFeatureAttributes uint32
 }
 
-func enableClientMetrics(ctx context.Context, s *settings, config storageConfig) (*metricsContext, error) {
-	var project string
-	// TODO: use new auth client
-	c, err := transport.Creds(ctx, s.clientOption...)
-	if err == nil {
-		project = c.ProjectID
-	}
-	metricsContext, err := newGRPCMetricContext(ctx, metricsConfig{
-		project:       project,
-		interval:      config.metricInterval,
-		manualReader:  config.manualReader,
-		meterProvider: config.meterProvider,
-	},
-	)
-	if err != nil {
-		return nil, fmt.Errorf("gRPC Metrics: %w", err)
-	}
-	return metricsContext, nil
-}
-
 // newGRPCStorageClient initializes a new storageClient that uses the gRPC
 // Storage API.
 func newGRPCStorageClient(ctx context.Context, opts ...storageOption) (client *grpcStorageClient, err error) {
@@ -167,16 +146,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
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
-
 	var clientMetrics *clientMetrics
 	var metricsCleanup func()
 	if isOtelMetricsEnabled(&config) {
@@ -323,9 +292,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	if c.metricsCleanup != nil {
 		c.metricsCleanup()
 	}
@@ -632,8 +598,6 @@ func (c *grpcStorageClient) ListObjects(ctx context.Co
 		ctx, record := startMetricsOp(it.ctx, "ListObjects", false)
 		defer func() { record(err) }()
 		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "grpcStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		var objects []*storagepb.Object
 		var gitr *gapic.ObjectIterator
 		err = run(ctx, func(ctx context.Context) error {
@@ -1268,9 +1232,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 	if !c.config.grpcBidiReads {
 		return c.NewRangeReaderReadObject(ctx, params, opts...)
 	}
-
-	ctx, _ = startSpan(ctx, "grpcStorageClient.NewRangeReader")
-	defer func() { endSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
