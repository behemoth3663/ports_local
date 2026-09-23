diff --git a/vendor/cloud.google.com/go/monitoring/apiv3/v2/uptime_check_client.go b/vendor/cloud.google.com/go/monitoring/apiv3/v2/uptime_check_client.go
index 12811e790..cd09da955 100644
--- vendor/cloud.google.com/go/monitoring/apiv3/v2/uptime_check_client.go.orig
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/uptime_check_client.go
@@ -246,16 +246,7 @@ type uptimeCheckGRPCClient struct {
 // Monitoring, and then clicking on “Uptime”.
 func NewUptimeCheckClient(ctx context.Context, opts ...option.ClientOption) (*UptimeCheckClient, error) {
 	clientOpts := defaultUptimeCheckGRPCClientOptions()
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
 	if newUptimeCheckClientHook != nil {
 		hookOpts, err := newUptimeCheckClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -277,25 +268,6 @@ func NewUptimeCheckClient(ctx context.Context, opts ...option.ClientOption) (*Up
 		logger:            internaloption.GetLogger(opts),
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
-
-		client.CallOptions.ListUptimeCheckConfigs = append(client.CallOptions.ListUptimeCheckConfigs, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetUptimeCheckConfig = append(client.CallOptions.GetUptimeCheckConfig, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateUptimeCheckConfig = append(client.CallOptions.CreateUptimeCheckConfig, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateUptimeCheckConfig = append(client.CallOptions.UpdateUptimeCheckConfig, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteUptimeCheckConfig = append(client.CallOptions.DeleteUptimeCheckConfig, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListUptimeCheckIps = append(client.CallOptions.ListUptimeCheckIps, gax.WithClientMetrics(metrics))
-	}
 
 	client.internalClient = c
 
@@ -332,12 +304,7 @@ func (c *uptimeCheckGRPCClient) ListUptimeCheckConfigs(ctx context.Context, req
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.UptimeCheckService/ListUptimeCheckConfigs")
-	}
+
 	opts = append((*c.CallOptions).ListUptimeCheckConfigs[0:len((*c.CallOptions).ListUptimeCheckConfigs):len((*c.CallOptions).ListUptimeCheckConfigs)], opts...)
 	it := &UptimeCheckConfigIterator{}
 	req = proto.CloneOf(req)
@@ -384,12 +351,7 @@ func (c *uptimeCheckGRPCClient) GetUptimeCheckConfig(ctx context.Context, req *m
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.UptimeCheckService/GetUptimeCheckConfig")
-	}
+
 	opts = append((*c.CallOptions).GetUptimeCheckConfig[0:len((*c.CallOptions).GetUptimeCheckConfig):len((*c.CallOptions).GetUptimeCheckConfig)], opts...)
 	var resp *monitoringpb.UptimeCheckConfig
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -408,12 +370,7 @@ func (c *uptimeCheckGRPCClient) CreateUptimeCheckConfig(ctx context.Context, req
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.UptimeCheckService/CreateUptimeCheckConfig")
-	}
+
 	opts = append((*c.CallOptions).CreateUptimeCheckConfig[0:len((*c.CallOptions).CreateUptimeCheckConfig):len((*c.CallOptions).CreateUptimeCheckConfig)], opts...)
 	var resp *monitoringpb.UptimeCheckConfig
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -432,9 +389,7 @@ func (c *uptimeCheckGRPCClient) UpdateUptimeCheckConfig(ctx context.Context, req
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.UptimeCheckService/UpdateUptimeCheckConfig")
-	}
+
 	opts = append((*c.CallOptions).UpdateUptimeCheckConfig[0:len((*c.CallOptions).UpdateUptimeCheckConfig):len((*c.CallOptions).UpdateUptimeCheckConfig)], opts...)
 	var resp *monitoringpb.UptimeCheckConfig
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -453,12 +408,7 @@ func (c *uptimeCheckGRPCClient) DeleteUptimeCheckConfig(ctx context.Context, req
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.UptimeCheckService/DeleteUptimeCheckConfig")
-	}
+
 	opts = append((*c.CallOptions).DeleteUptimeCheckConfig[0:len((*c.CallOptions).DeleteUptimeCheckConfig):len((*c.CallOptions).DeleteUptimeCheckConfig)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -470,9 +420,7 @@ func (c *uptimeCheckGRPCClient) DeleteUptimeCheckConfig(ctx context.Context, req
 
 func (c *uptimeCheckGRPCClient) ListUptimeCheckIps(ctx context.Context, req *monitoringpb.ListUptimeCheckIpsRequest, opts ...gax.CallOption) *UptimeCheckIpIterator {
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.UptimeCheckService/ListUptimeCheckIps")
-	}
+
 	opts = append((*c.CallOptions).ListUptimeCheckIps[0:len((*c.CallOptions).ListUptimeCheckIps):len((*c.CallOptions).ListUptimeCheckIps)], opts...)
 	it := &UptimeCheckIpIterator{}
 	req = proto.CloneOf(req)
