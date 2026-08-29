--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2026-06-04 04:52:32 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -21,7 +21,6 @@ import (
 	"fmt"
 	"hash/crc32"
 	"io"
-	"log"
 	"os"
 	"strconv"
 	"strings"
@@ -33,7 +32,6 @@ import (
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
-	"google.golang.org/api/transport"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/encoding"
@@ -130,26 +128,6 @@ type grpcStorageClient struct {
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
 func newGRPCStorageClient(ctx context.Context, opts ...storageOption) (*grpcStorageClient, error) {
@@ -163,15 +141,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
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
 	c := &grpcStorageClient{
 		settings: s,
 		config:   &config,
@@ -272,9 +241,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	return c.raw.Close()
 }
 
@@ -574,8 +540,6 @@ func (c *grpcStorageClient) ListObjects(ctx context.Co
 	}
 	fetch := func(pageSize int, pageToken string) (token string, err error) {
 		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "grpcStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		var objects []*storagepb.Object
 		var gitr *gapic.ObjectIterator
 		err = run(it.ctx, func(ctx context.Context) error {
@@ -1208,8 +1172,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 		return c.NewRangeReaderReadObject(ctx, params, opts...)
 	}
 
-	ctx, _ = startSpan(ctx, "grpcStorageClient.NewRangeReader")
-	defer func() { endSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
