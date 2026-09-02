--- vendor/cloud.google.com/go/storage/http_client.go.orig	2026-06-04 04:52:32 UTC
+++ vendor/cloud.google.com/go/storage/http_client.go
@@ -131,18 +131,6 @@ func newHTTPStorageClient(ctx context.Context, opts ..
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
@@ -349,8 +337,6 @@ func (c *httpStorageClient) ListObjects(ctx context.Co
 	fetch := func(pageSize int, pageToken string) (string, error) {
 		var err error
 		// Add trace span around List API call within the fetch.
-		ctx, _ = startSpan(ctx, "httpStorageClient.ObjectsListCall")
-		defer func() { endSpan(ctx, err) }()
 		req := c.raw.Objects.List(bucket)
 		if it.query.SoftDeleted {
 			req.SoftDeleted(it.query.SoftDeleted)
@@ -870,9 +856,6 @@ func (c *httpStorageClient) NewRangeReader(ctx context
 }
 
 func (c *httpStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx, _ = startSpan(ctx, "httpStorageClient.NewRangeReader")
-	defer func() { endSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	if c.config.useJSONforReads {
