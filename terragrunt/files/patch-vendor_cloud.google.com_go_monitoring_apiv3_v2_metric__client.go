diff --git a/vendor/cloud.google.com/go/monitoring/apiv3/v2/metric_client.go b/vendor/cloud.google.com/go/monitoring/apiv3/v2/metric_client.go
index 02544a1d8..4995f1faf 100644
--- vendor/cloud.google.com/go/monitoring/apiv3/v2/metric_client.go.orig
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/metric_client.go
@@ -297,16 +297,7 @@ type metricGRPCClient struct {
 // time series data.
 func NewMetricClient(ctx context.Context, opts ...option.ClientOption) (*MetricClient, error) {
 	clientOpts := defaultMetricGRPCClientOptions()
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
 	if newMetricClientHook != nil {
 		hookOpts, err := newMetricClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -328,28 +319,6 @@ func NewMetricClient(ctx context.Context, opts ...option.ClientOption) (*MetricC
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
-		client.CallOptions.ListMonitoredResourceDescriptors = append(client.CallOptions.ListMonitoredResourceDescriptors, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetMonitoredResourceDescriptor = append(client.CallOptions.GetMonitoredResourceDescriptor, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListMetricDescriptors = append(client.CallOptions.ListMetricDescriptors, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetMetricDescriptor = append(client.CallOptions.GetMetricDescriptor, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateMetricDescriptor = append(client.CallOptions.CreateMetricDescriptor, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteMetricDescriptor = append(client.CallOptions.DeleteMetricDescriptor, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListTimeSeries = append(client.CallOptions.ListTimeSeries, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateTimeSeries = append(client.CallOptions.CreateTimeSeries, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateServiceTimeSeries = append(client.CallOptions.CreateServiceTimeSeries, gax.WithClientMetrics(metrics))
-	}
 
 	client.internalClient = c
 
@@ -386,12 +355,7 @@ func (c *metricGRPCClient) ListMonitoredResourceDescriptors(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/ListMonitoredResourceDescriptors")
-	}
+
 	opts = append((*c.CallOptions).ListMonitoredResourceDescriptors[0:len((*c.CallOptions).ListMonitoredResourceDescriptors):len((*c.CallOptions).ListMonitoredResourceDescriptors)], opts...)
 	it := &MonitoredResourceDescriptorIterator{}
 	req = proto.CloneOf(req)
@@ -438,12 +402,7 @@ func (c *metricGRPCClient) GetMonitoredResourceDescriptor(ctx context.Context, r
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/GetMonitoredResourceDescriptor")
-	}
+
 	opts = append((*c.CallOptions).GetMonitoredResourceDescriptor[0:len((*c.CallOptions).GetMonitoredResourceDescriptor):len((*c.CallOptions).GetMonitoredResourceDescriptor)], opts...)
 	var resp *monitoredrespb.MonitoredResourceDescriptor
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -462,12 +421,7 @@ func (c *metricGRPCClient) ListMetricDescriptors(ctx context.Context, req *monit
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/ListMetricDescriptors")
-	}
+
 	opts = append((*c.CallOptions).ListMetricDescriptors[0:len((*c.CallOptions).ListMetricDescriptors):len((*c.CallOptions).ListMetricDescriptors)], opts...)
 	it := &MetricDescriptorIterator{}
 	req = proto.CloneOf(req)
@@ -514,12 +468,7 @@ func (c *metricGRPCClient) GetMetricDescriptor(ctx context.Context, req *monitor
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/GetMetricDescriptor")
-	}
+
 	opts = append((*c.CallOptions).GetMetricDescriptor[0:len((*c.CallOptions).GetMetricDescriptor):len((*c.CallOptions).GetMetricDescriptor)], opts...)
 	var resp *metricpb.MetricDescriptor
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -538,12 +487,7 @@ func (c *metricGRPCClient) CreateMetricDescriptor(ctx context.Context, req *moni
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/CreateMetricDescriptor")
-	}
+
 	opts = append((*c.CallOptions).CreateMetricDescriptor[0:len((*c.CallOptions).CreateMetricDescriptor):len((*c.CallOptions).CreateMetricDescriptor)], opts...)
 	var resp *metricpb.MetricDescriptor
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -562,12 +506,7 @@ func (c *metricGRPCClient) DeleteMetricDescriptor(ctx context.Context, req *moni
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/DeleteMetricDescriptor")
-	}
+
 	opts = append((*c.CallOptions).DeleteMetricDescriptor[0:len((*c.CallOptions).DeleteMetricDescriptor):len((*c.CallOptions).DeleteMetricDescriptor)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -582,12 +521,7 @@ func (c *metricGRPCClient) ListTimeSeries(ctx context.Context, req *monitoringpb
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/ListTimeSeries")
-	}
+
 	opts = append((*c.CallOptions).ListTimeSeries[0:len((*c.CallOptions).ListTimeSeries):len((*c.CallOptions).ListTimeSeries)], opts...)
 	it := &TimeSeriesIterator{}
 	req = proto.CloneOf(req)
@@ -634,12 +568,7 @@ func (c *metricGRPCClient) CreateTimeSeries(ctx context.Context, req *monitoring
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/CreateTimeSeries")
-	}
+
 	opts = append((*c.CallOptions).CreateTimeSeries[0:len((*c.CallOptions).CreateTimeSeries):len((*c.CallOptions).CreateTimeSeries)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -654,12 +583,7 @@ func (c *metricGRPCClient) CreateServiceTimeSeries(ctx context.Context, req *mon
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.MetricService/CreateServiceTimeSeries")
-	}
+
 	opts = append((*c.CallOptions).CreateServiceTimeSeries[0:len((*c.CallOptions).CreateServiceTimeSeries):len((*c.CallOptions).CreateServiceTimeSeries)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
