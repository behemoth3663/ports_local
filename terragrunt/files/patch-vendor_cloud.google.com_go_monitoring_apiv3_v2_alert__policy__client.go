diff --git a/vendor/cloud.google.com/go/monitoring/apiv3/v2/alert_policy_client.go b/vendor/cloud.google.com/go/monitoring/apiv3/v2/alert_policy_client.go
index 72d871dc4..ce264fb17 100644
--- vendor/cloud.google.com/go/monitoring/apiv3/v2/alert_policy_client.go.orig
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/alert_policy_client.go
@@ -238,16 +238,7 @@ type alertPolicyGRPCClient struct {
 // Cloud console (at https://console.cloud.google.com/).
 func NewAlertPolicyClient(ctx context.Context, opts ...option.ClientOption) (*AlertPolicyClient, error) {
 	clientOpts := defaultAlertPolicyGRPCClientOptions()
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
 	if newAlertPolicyClientHook != nil {
 		hookOpts, err := newAlertPolicyClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -269,24 +260,6 @@ func NewAlertPolicyClient(ctx context.Context, opts ...option.ClientOption) (*Al
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
-		client.CallOptions.ListAlertPolicies = append(client.CallOptions.ListAlertPolicies, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetAlertPolicy = append(client.CallOptions.GetAlertPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateAlertPolicy = append(client.CallOptions.CreateAlertPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteAlertPolicy = append(client.CallOptions.DeleteAlertPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateAlertPolicy = append(client.CallOptions.UpdateAlertPolicy, gax.WithClientMetrics(metrics))
-	}
 
 	client.internalClient = c
 
@@ -323,12 +296,7 @@ func (c *alertPolicyGRPCClient) ListAlertPolicies(ctx context.Context, req *moni
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.AlertPolicyService/ListAlertPolicies")
-	}
+
 	opts = append((*c.CallOptions).ListAlertPolicies[0:len((*c.CallOptions).ListAlertPolicies):len((*c.CallOptions).ListAlertPolicies)], opts...)
 	it := &AlertPolicyIterator{}
 	req = proto.CloneOf(req)
@@ -375,12 +343,7 @@ func (c *alertPolicyGRPCClient) GetAlertPolicy(ctx context.Context, req *monitor
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.AlertPolicyService/GetAlertPolicy")
-	}
+
 	opts = append((*c.CallOptions).GetAlertPolicy[0:len((*c.CallOptions).GetAlertPolicy):len((*c.CallOptions).GetAlertPolicy)], opts...)
 	var resp *monitoringpb.AlertPolicy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -399,12 +362,7 @@ func (c *alertPolicyGRPCClient) CreateAlertPolicy(ctx context.Context, req *moni
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.AlertPolicyService/CreateAlertPolicy")
-	}
+
 	opts = append((*c.CallOptions).CreateAlertPolicy[0:len((*c.CallOptions).CreateAlertPolicy):len((*c.CallOptions).CreateAlertPolicy)], opts...)
 	var resp *monitoringpb.AlertPolicy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -423,12 +381,7 @@ func (c *alertPolicyGRPCClient) DeleteAlertPolicy(ctx context.Context, req *moni
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.AlertPolicyService/DeleteAlertPolicy")
-	}
+
 	opts = append((*c.CallOptions).DeleteAlertPolicy[0:len((*c.CallOptions).DeleteAlertPolicy):len((*c.CallOptions).DeleteAlertPolicy)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -443,9 +396,7 @@ func (c *alertPolicyGRPCClient) UpdateAlertPolicy(ctx context.Context, req *moni
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.AlertPolicyService/UpdateAlertPolicy")
-	}
+
 	opts = append((*c.CallOptions).UpdateAlertPolicy[0:len((*c.CallOptions).UpdateAlertPolicy):len((*c.CallOptions).UpdateAlertPolicy)], opts...)
 	var resp *monitoringpb.AlertPolicy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
