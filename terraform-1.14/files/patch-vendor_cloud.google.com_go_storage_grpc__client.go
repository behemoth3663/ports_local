--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2026-03-13 06:41:11 UTC
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
@@ -126,26 +124,6 @@ type grpcStorageClient struct {
 	dpDiag   string
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
@@ -159,15 +137,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
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
@@ -244,9 +213,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	return c.raw.Close()
 }
 
@@ -546,8 +512,6 @@ func (c *grpcStorageClient) ListObjects(ctx context.Co
 	}
 	fetch := func(pageSize int, pageToken string) (token string, err error) {
 		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "grpcStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		var objects []*storagepb.Object
 		var gitr *gapic.ObjectIterator
 		err = run(it.ctx, func(ctx context.Context) error {
@@ -1177,9 +1141,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 	if !c.config.grpcBidiReads {
 		return c.NewRangeReaderReadObject(ctx, params, opts...)
 	}
-
-	ctx, _ = startSpan(ctx, "grpcStorageClient.NewRangeReader")
-	defer func() { endSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
