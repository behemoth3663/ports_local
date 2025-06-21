--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2025-05-21 15:38:31 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -21,19 +21,16 @@ import (
 	"fmt"
 	"hash/crc32"
 	"io"
-	"log"
 	"os"
 	"sync"
 
 	"cloud.google.com/go/iam/apiv1/iampb"
-	"cloud.google.com/go/internal/trace"
 	gapic "cloud.google.com/go/storage/internal/apiv2"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
-	"google.golang.org/api/transport"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/encoding"
@@ -112,24 +109,6 @@ type grpcStorageClient struct {
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
 func newGRPCStorageClient(ctx context.Context, opts ...storageOption) (*grpcStorageClient, error) {
@@ -143,15 +122,6 @@ func newGRPCStorageClient(ctx context.Context, opts ..
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
@@ -165,9 +135,6 @@ func (c *grpcStorageClient) Close() error {
 }
 
 func (c *grpcStorageClient) Close() error {
-	if c.settings.metricsContext != nil {
-		c.settings.metricsContext.close()
-	}
 	return c.raw.Close()
 }
 
@@ -453,8 +420,6 @@ func (c *grpcStorageClient) ListObjects(ctx context.Co
 	}
 	fetch := func(pageSize int, pageToken string) (token string, err error) {
 		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "grpcStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		var objects []*storagepb.Object
 		var gitr *gapic.ObjectIterator
 		err = run(it.ctx, func(ctx context.Context) error {
@@ -1062,8 +1027,6 @@ func (c *grpcStorageClient) NewMultiRangeDownloader(ct
 		return nil, errors.New("storage: MultiRangeDownloader requires the experimental.WithGRPCBidiReads option")
 	}
 
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewMultiRangeDownloader")
-	defer func() { trace.EndSpan(ctx, err) }()
 	s := callSettings(c.settings, opts...)
 
 	if s.userProject != "" {
@@ -1536,9 +1499,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 	if !c.config.grpcBidiReads {
 		return c.NewRangeReaderReadObject(ctx, params, opts...)
 	}
-
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
