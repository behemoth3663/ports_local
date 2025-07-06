--- vendor/cloud.google.com/go/storage/http_client.go.orig	2025-05-30 06:48:19 UTC
+++ vendor/cloud.google.com/go/storage/http_client.go
@@ -33,7 +33,6 @@ import (
 	"cloud.google.com/go/auth"
 	"cloud.google.com/go/iam/apiv1/iampb"
 	"cloud.google.com/go/internal/optional"
-	"cloud.google.com/go/internal/trace"
 	"github.com/google/uuid"
 	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/googleapi"
@@ -130,18 +129,6 @@ func newHTTPStorageClient(ctx context.Context, opts ..
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
@@ -345,8 +332,6 @@ func (c *httpStorageClient) ListObjects(ctx context.Co
 	fetch := func(pageSize int, pageToken string) (string, error) {
 		var err error
 		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "httpStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		req := c.raw.Objects.List(bucket)
 		if it.query.SoftDeleted {
 			req.SoftDeleted(it.query.SoftDeleted)
@@ -847,9 +832,6 @@ func (c *httpStorageClient) NewRangeReader(ctx context
 }
 
 func (c *httpStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.httpStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	if c.config.useJSONforReads {
