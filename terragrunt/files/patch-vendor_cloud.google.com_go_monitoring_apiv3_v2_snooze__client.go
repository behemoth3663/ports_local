diff --git a/vendor/cloud.google.com/go/monitoring/apiv3/v2/snooze_client.go b/vendor/cloud.google.com/go/monitoring/apiv3/v2/snooze_client.go
index 3a45e85ff..5cb4cde59 100644
--- vendor/cloud.google.com/go/monitoring/apiv3/v2/snooze_client.go.orig
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/snooze_client.go
@@ -196,16 +196,7 @@ type snoozeGRPCClient struct {
 // or more alert policies should not fire alerts for the specified duration.
 func NewSnoozeClient(ctx context.Context, opts ...option.ClientOption) (*SnoozeClient, error) {
 	clientOpts := defaultSnoozeGRPCClientOptions()
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
 	if newSnoozeClientHook != nil {
 		hookOpts, err := newSnoozeClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -227,23 +218,6 @@ func NewSnoozeClient(ctx context.Context, opts ...option.ClientOption) (*SnoozeC
 		logger:       internaloption.GetLogger(opts),
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
-		client.CallOptions.CreateSnooze = append(client.CallOptions.CreateSnooze, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListSnoozes = append(client.CallOptions.ListSnoozes, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetSnooze = append(client.CallOptions.GetSnooze, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateSnooze = append(client.CallOptions.UpdateSnooze, gax.WithClientMetrics(metrics))
-	}
 
 	client.internalClient = c
 
@@ -280,12 +254,7 @@ func (c *snoozeGRPCClient) CreateSnooze(ctx context.Context, req *monitoringpb.C
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.SnoozeService/CreateSnooze")
-	}
+
 	opts = append((*c.CallOptions).CreateSnooze[0:len((*c.CallOptions).CreateSnooze):len((*c.CallOptions).CreateSnooze)], opts...)
 	var resp *monitoringpb.Snooze
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -304,12 +273,7 @@ func (c *snoozeGRPCClient) ListSnoozes(ctx context.Context, req *monitoringpb.Li
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.SnoozeService/ListSnoozes")
-	}
+
 	opts = append((*c.CallOptions).ListSnoozes[0:len((*c.CallOptions).ListSnoozes):len((*c.CallOptions).ListSnoozes)], opts...)
 	it := &SnoozeIterator{}
 	req = proto.CloneOf(req)
@@ -356,12 +320,7 @@ func (c *snoozeGRPCClient) GetSnooze(ctx context.Context, req *monitoringpb.GetS
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.SnoozeService/GetSnooze")
-	}
+
 	opts = append((*c.CallOptions).GetSnooze[0:len((*c.CallOptions).GetSnooze):len((*c.CallOptions).GetSnooze)], opts...)
 	var resp *monitoringpb.Snooze
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -380,9 +339,7 @@ func (c *snoozeGRPCClient) UpdateSnooze(ctx context.Context, req *monitoringpb.U
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.SnoozeService/UpdateSnooze")
-	}
+
 	opts = append((*c.CallOptions).UpdateSnooze[0:len((*c.CallOptions).UpdateSnooze):len((*c.CallOptions).UpdateSnooze)], opts...)
 	var resp *monitoringpb.Snooze
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
