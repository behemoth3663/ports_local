--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2025-03-20 16:11:04 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -27,7 +27,6 @@ import (
 	"sync"
 
 	"cloud.google.com/go/iam/apiv1/iampb"
-	"cloud.google.com/go/internal/trace"
 	gapic "cloud.google.com/go/storage/internal/apiv2"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
@@ -128,8 +127,7 @@ func enableClientMetrics(ctx context.Context, s *setti
 	}
 	metricsContext, err := newGRPCMetricContext(ctx, metricsConfig{
 		project:      project,
-		interval:     config.metricInterval,
-		manualReader: config.manualReader},
+		interval:     config.metricInterval},
 	)
 	if err != nil {
 		return nil, fmt.Errorf("gRPC Metrics: %w", err)
@@ -1072,8 +1070,6 @@ func (c *grpcStorageClient) NewMultiRangeDownloader(ct
 }
 
 func (c *grpcStorageClient) NewMultiRangeDownloader(ctx context.Context, params *newMultiRangeDownloaderParams, opts ...storageOption) (mr *MultiRangeDownloader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewMultiRangeDownloader")
-	defer func() { trace.EndSpan(ctx, err) }()
 	s := callSettings(c.settings, opts...)
 
 	if s.userProject != "" {
@@ -1504,9 +1500,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 	if !c.config.grpcBidiReads {
 		return c.NewRangeReaderReadObject(ctx, params, opts...)
 	}
-
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
