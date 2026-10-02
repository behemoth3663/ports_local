--- vendor/cloud.google.com/go/monitoring/apiv3/v2/query_client.go.orig	2026-10-02 16:09:03 UTC
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/query_client.go
@@ -145,16 +145,7 @@ func NewQueryClient(ctx context.Context, opts ...optio
 // the time-varying values of a metric.
 func NewQueryClient(ctx context.Context, opts ...option.ClientOption) (*QueryClient, error) {
 	clientOpts := defaultQueryGRPCClientOptions()
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "monitoring",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/monitoring/apiv3/v2",
-			"gcp.client.language": "go",
-			"url.domain":          "monitoring.googleapis.com",
-		}))
-	}
+
 	if newQueryClientHook != nil {
 		hookOpts, err := newQueryClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -176,21 +167,7 @@ func NewQueryClient(ctx context.Context, opts ...optio
 		logger:      internaloption.GetLogger(opts),
 	}
 	c.setGoogleClientInfo()
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "monitoring",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/monitoring/apiv3/v2",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "monitoring.googleapis.com",
-			}),
-		)
 
-		client.CallOptions.QueryTimeSeries = append(client.CallOptions.QueryTimeSeries, gax.WithClientMetrics(metrics))
-	}
-
 	client.internalClient = c
 
 	return &client, nil
@@ -226,9 +203,7 @@ func (c *queryGRPCClient) QueryTimeSeries(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.QueryService/QueryTimeSeries")
-	}
+
 	opts = append((*c.CallOptions).QueryTimeSeries[0:len((*c.CallOptions).QueryTimeSeries):len((*c.CallOptions).QueryTimeSeries)], opts...)
 	it := &TimeSeriesDataIterator{}
 	req = proto.CloneOf(req)
