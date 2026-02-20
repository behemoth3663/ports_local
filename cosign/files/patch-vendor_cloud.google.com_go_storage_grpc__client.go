--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2025-09-22 16:26:20 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -26,7 +26,6 @@ import (
 	"sync"
 
 	"cloud.google.com/go/iam/apiv1/iampb"
-	"cloud.google.com/go/internal/trace"
 	gapic "cloud.google.com/go/storage/internal/apiv2"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
@@ -127,7 +126,6 @@ func enableClientMetrics(ctx context.Context, s *setti
 	metricsContext, err := newGRPCMetricContext(ctx, metricsConfig{
 		project:       project,
 		interval:      config.metricInterval,
-		manualReader:  config.manualReader,
 		meterProvider: config.meterProvider,
 	},
 	)
@@ -459,9 +457,6 @@ func (c *grpcStorageClient) ListObjects(ctx context.Co
 		ctx = setUserProjectMetadata(ctx, s.userProject)
 	}
 	fetch := func(pageSize int, pageToken string) (token string, err error) {
-		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "grpcStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		var objects []*storagepb.Object
 		var gitr *gapic.ObjectIterator
 		err = run(it.ctx, func(ctx context.Context) error {
@@ -1069,8 +1064,6 @@ func (c *grpcStorageClient) NewMultiRangeDownloader(ct
 		return nil, errors.New("storage: MultiRangeDownloader requires the experimental.WithGRPCBidiReads option")
 	}
 
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewMultiRangeDownloader")
-	defer func() { trace.EndSpan(ctx, err) }()
 	s := callSettings(c.settings, opts...)
 	// Force the use of the custom codec to enable zero-copy reads.
 	s.gax = append(s.gax, gax.WithGRPCOptions(
@@ -1613,9 +1606,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 	if !c.config.grpcBidiReads {
 		return c.NewRangeReaderReadObject(ctx, params, opts...)
 	}
-
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
