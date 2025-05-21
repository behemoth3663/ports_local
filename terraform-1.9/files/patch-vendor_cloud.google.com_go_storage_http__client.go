--- vendor/cloud.google.com/go/storage/http_client.go.orig	2025-03-12 21:31:11 UTC
+++ vendor/cloud.google.com/go/storage/http_client.go
@@ -33,7 +33,6 @@ import (
 
 	"cloud.google.com/go/iam/apiv1/iampb"
 	"cloud.google.com/go/internal/optional"
-	"cloud.google.com/go/internal/trace"
 	"github.com/googleapis/gax-go/v2/callctx"
 	"golang.org/x/oauth2/google"
 	"google.golang.org/api/googleapi"
@@ -131,18 +130,6 @@ func newHTTPStorageClient(ctx context.Context, opts ..
 	}
 
 	var bd *bucketDelayManager
-	if config.readStallTimeoutConfig != nil {
-		drrstConfig := config.readStallTimeoutConfig
-		bd, err = newBucketDelayManager(
-			drrstConfig.TargetPercentile,
-			getDynamicReadReqIncreaseRateFromEnv(),
-			getDynamicReadReqInitialTimeoutSecFromEnv(drrstConfig.Min),
-			drrstConfig.Min,
-			defaultDynamicReqdReqMaxTimeout)
-		if err != nil {
-			return nil, fmt.Errorf("creating dynamic-delay: %w", err)
-		}
-	}
 
 	return &httpStorageClient{
 		creds:                      creds,
@@ -845,9 +832,6 @@ func (c *httpStorageClient) NewRangeReader(ctx context
 }
 
 func (c *httpStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.httpStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	if c.config.useJSONforReads {
