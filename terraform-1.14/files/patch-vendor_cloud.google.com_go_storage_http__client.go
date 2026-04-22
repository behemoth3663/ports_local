--- vendor/cloud.google.com/go/storage/http_client.go.orig	2026-03-13 06:41:11 UTC
+++ vendor/cloud.google.com/go/storage/http_client.go
@@ -129,18 +129,6 @@ func newHTTPStorageClient(ctx context.Context, opts ..
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
@@ -346,8 +334,6 @@ func (c *httpStorageClient) ListObjects(ctx context.Co
 	fetch := func(pageSize int, pageToken string) (string, error) {
 		var err error
 		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "httpStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		req := c.raw.Objects.List(bucket)
 		if it.query.SoftDeleted {
 			req.SoftDeleted(it.query.SoftDeleted)
@@ -867,8 +853,6 @@ func (c *httpStorageClient) NewRangeReader(ctx context
 }
 
 func (c *httpStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx, _ = startSpan(ctx, "httpStorageClient.NewRangeReader")
-	defer func() { endSpan(ctx, err) }()
 
 	s := callSettings(c.settings, opts...)
 
