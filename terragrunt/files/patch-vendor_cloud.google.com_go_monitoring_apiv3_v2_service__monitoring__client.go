diff --git a/vendor/cloud.google.com/go/monitoring/apiv3/v2/service_monitoring_client.go b/vendor/cloud.google.com/go/monitoring/apiv3/v2/service_monitoring_client.go
index a0023d255..49a852a66 100644
--- vendor/cloud.google.com/go/monitoring/apiv3/v2/service_monitoring_client.go.orig
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/service_monitoring_client.go
@@ -290,16 +290,7 @@ type serviceMonitoringGRPCClient struct {
 // taxonomy of categorized Health Metrics.
 func NewServiceMonitoringClient(ctx context.Context, opts ...option.ClientOption) (*ServiceMonitoringClient, error) {
 	clientOpts := defaultServiceMonitoringGRPCClientOptions()
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
 	if newServiceMonitoringClientHook != nil {
 		hookOpts, err := newServiceMonitoringClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -321,29 +312,6 @@ func NewServiceMonitoringClient(ctx context.Context, opts ...option.ClientOption
 		logger:                  internaloption.GetLogger(opts),
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
-		client.CallOptions.CreateService = append(client.CallOptions.CreateService, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetService = append(client.CallOptions.GetService, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListServices = append(client.CallOptions.ListServices, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateService = append(client.CallOptions.UpdateService, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteService = append(client.CallOptions.DeleteService, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateServiceLevelObjective = append(client.CallOptions.CreateServiceLevelObjective, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetServiceLevelObjective = append(client.CallOptions.GetServiceLevelObjective, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListServiceLevelObjectives = append(client.CallOptions.ListServiceLevelObjectives, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateServiceLevelObjective = append(client.CallOptions.UpdateServiceLevelObjective, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteServiceLevelObjective = append(client.CallOptions.DeleteServiceLevelObjective, gax.WithClientMetrics(metrics))
-	}
 
 	client.internalClient = c
 
@@ -380,12 +348,7 @@ func (c *serviceMonitoringGRPCClient) CreateService(ctx context.Context, req *mo
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/CreateService")
-	}
+
 	opts = append((*c.CallOptions).CreateService[0:len((*c.CallOptions).CreateService):len((*c.CallOptions).CreateService)], opts...)
 	var resp *monitoringpb.Service
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -404,12 +367,7 @@ func (c *serviceMonitoringGRPCClient) GetService(ctx context.Context, req *monit
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/GetService")
-	}
+
 	opts = append((*c.CallOptions).GetService[0:len((*c.CallOptions).GetService):len((*c.CallOptions).GetService)], opts...)
 	var resp *monitoringpb.Service
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -428,12 +386,7 @@ func (c *serviceMonitoringGRPCClient) ListServices(ctx context.Context, req *mon
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/ListServices")
-	}
+
 	opts = append((*c.CallOptions).ListServices[0:len((*c.CallOptions).ListServices):len((*c.CallOptions).ListServices)], opts...)
 	it := &ServiceIterator{}
 	req = proto.CloneOf(req)
@@ -480,9 +433,7 @@ func (c *serviceMonitoringGRPCClient) UpdateService(ctx context.Context, req *mo
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/UpdateService")
-	}
+
 	opts = append((*c.CallOptions).UpdateService[0:len((*c.CallOptions).UpdateService):len((*c.CallOptions).UpdateService)], opts...)
 	var resp *monitoringpb.Service
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -501,12 +452,7 @@ func (c *serviceMonitoringGRPCClient) DeleteService(ctx context.Context, req *mo
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/DeleteService")
-	}
+
 	opts = append((*c.CallOptions).DeleteService[0:len((*c.CallOptions).DeleteService):len((*c.CallOptions).DeleteService)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -521,12 +467,7 @@ func (c *serviceMonitoringGRPCClient) CreateServiceLevelObjective(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/CreateServiceLevelObjective")
-	}
+
 	opts = append((*c.CallOptions).CreateServiceLevelObjective[0:len((*c.CallOptions).CreateServiceLevelObjective):len((*c.CallOptions).CreateServiceLevelObjective)], opts...)
 	var resp *monitoringpb.ServiceLevelObjective
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -545,12 +486,7 @@ func (c *serviceMonitoringGRPCClient) GetServiceLevelObjective(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/GetServiceLevelObjective")
-	}
+
 	opts = append((*c.CallOptions).GetServiceLevelObjective[0:len((*c.CallOptions).GetServiceLevelObjective):len((*c.CallOptions).GetServiceLevelObjective)], opts...)
 	var resp *monitoringpb.ServiceLevelObjective
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -569,12 +505,7 @@ func (c *serviceMonitoringGRPCClient) ListServiceLevelObjectives(ctx context.Con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/ListServiceLevelObjectives")
-	}
+
 	opts = append((*c.CallOptions).ListServiceLevelObjectives[0:len((*c.CallOptions).ListServiceLevelObjectives):len((*c.CallOptions).ListServiceLevelObjectives)], opts...)
 	it := &ServiceLevelObjectiveIterator{}
 	req = proto.CloneOf(req)
@@ -621,9 +552,7 @@ func (c *serviceMonitoringGRPCClient) UpdateServiceLevelObjective(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/UpdateServiceLevelObjective")
-	}
+
 	opts = append((*c.CallOptions).UpdateServiceLevelObjective[0:len((*c.CallOptions).UpdateServiceLevelObjective):len((*c.CallOptions).UpdateServiceLevelObjective)], opts...)
 	var resp *monitoringpb.ServiceLevelObjective
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -642,12 +571,7 @@ func (c *serviceMonitoringGRPCClient) DeleteServiceLevelObjective(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.ServiceMonitoringService/DeleteServiceLevelObjective")
-	}
+
 	opts = append((*c.CallOptions).DeleteServiceLevelObjective[0:len((*c.CallOptions).DeleteServiceLevelObjective):len((*c.CallOptions).DeleteServiceLevelObjective)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
