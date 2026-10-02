--- vendor/cloud.google.com/go/monitoring/apiv3/v2/notification_channel_client.go.orig	2026-10-02 16:09:03 UTC
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/notification_channel_client.go
@@ -343,16 +343,7 @@ func NewNotificationChannelClient(ctx context.Context,
 // controls how messages related to incidents are sent.
 func NewNotificationChannelClient(ctx context.Context, opts ...option.ClientOption) (*NotificationChannelClient, error) {
 	clientOpts := defaultNotificationChannelGRPCClientOptions()
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
 	if newNotificationChannelClientHook != nil {
 		hookOpts, err := newNotificationChannelClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -374,30 +365,7 @@ func NewNotificationChannelClient(ctx context.Context,
 		logger:                    internaloption.GetLogger(opts),
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
 
-		client.CallOptions.ListNotificationChannelDescriptors = append(client.CallOptions.ListNotificationChannelDescriptors, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetNotificationChannelDescriptor = append(client.CallOptions.GetNotificationChannelDescriptor, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListNotificationChannels = append(client.CallOptions.ListNotificationChannels, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetNotificationChannel = append(client.CallOptions.GetNotificationChannel, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateNotificationChannel = append(client.CallOptions.CreateNotificationChannel, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateNotificationChannel = append(client.CallOptions.UpdateNotificationChannel, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteNotificationChannel = append(client.CallOptions.DeleteNotificationChannel, gax.WithClientMetrics(metrics))
-		client.CallOptions.SendNotificationChannelVerificationCode = append(client.CallOptions.SendNotificationChannelVerificationCode, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetNotificationChannelVerificationCode = append(client.CallOptions.GetNotificationChannelVerificationCode, gax.WithClientMetrics(metrics))
-		client.CallOptions.VerifyNotificationChannel = append(client.CallOptions.VerifyNotificationChannel, gax.WithClientMetrics(metrics))
-	}
-
 	client.internalClient = c
 
 	return &client, nil
@@ -433,12 +401,7 @@ func (c *notificationChannelGRPCClient) ListNotificati
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/ListNotificationChannelDescriptors")
-	}
+
 	opts = append((*c.CallOptions).ListNotificationChannelDescriptors[0:len((*c.CallOptions).ListNotificationChannelDescriptors):len((*c.CallOptions).ListNotificationChannelDescriptors)], opts...)
 	it := &NotificationChannelDescriptorIterator{}
 	req = proto.CloneOf(req)
@@ -485,12 +448,7 @@ func (c *notificationChannelGRPCClient) GetNotificatio
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/GetNotificationChannelDescriptor")
-	}
+
 	opts = append((*c.CallOptions).GetNotificationChannelDescriptor[0:len((*c.CallOptions).GetNotificationChannelDescriptor):len((*c.CallOptions).GetNotificationChannelDescriptor)], opts...)
 	var resp *monitoringpb.NotificationChannelDescriptor
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -509,12 +467,7 @@ func (c *notificationChannelGRPCClient) ListNotificati
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/ListNotificationChannels")
-	}
+
 	opts = append((*c.CallOptions).ListNotificationChannels[0:len((*c.CallOptions).ListNotificationChannels):len((*c.CallOptions).ListNotificationChannels)], opts...)
 	it := &NotificationChannelIterator{}
 	req = proto.CloneOf(req)
@@ -561,12 +514,7 @@ func (c *notificationChannelGRPCClient) GetNotificatio
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/GetNotificationChannel")
-	}
+
 	opts = append((*c.CallOptions).GetNotificationChannel[0:len((*c.CallOptions).GetNotificationChannel):len((*c.CallOptions).GetNotificationChannel)], opts...)
 	var resp *monitoringpb.NotificationChannel
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -585,12 +533,7 @@ func (c *notificationChannelGRPCClient) CreateNotifica
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/CreateNotificationChannel")
-	}
+
 	opts = append((*c.CallOptions).CreateNotificationChannel[0:len((*c.CallOptions).CreateNotificationChannel):len((*c.CallOptions).CreateNotificationChannel)], opts...)
 	var resp *monitoringpb.NotificationChannel
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -609,9 +552,7 @@ func (c *notificationChannelGRPCClient) UpdateNotifica
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/UpdateNotificationChannel")
-	}
+
 	opts = append((*c.CallOptions).UpdateNotificationChannel[0:len((*c.CallOptions).UpdateNotificationChannel):len((*c.CallOptions).UpdateNotificationChannel)], opts...)
 	var resp *monitoringpb.NotificationChannel
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -630,12 +571,7 @@ func (c *notificationChannelGRPCClient) DeleteNotifica
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/DeleteNotificationChannel")
-	}
+
 	opts = append((*c.CallOptions).DeleteNotificationChannel[0:len((*c.CallOptions).DeleteNotificationChannel):len((*c.CallOptions).DeleteNotificationChannel)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -650,12 +586,7 @@ func (c *notificationChannelGRPCClient) SendNotificati
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/SendNotificationChannelVerificationCode")
-	}
+
 	opts = append((*c.CallOptions).SendNotificationChannelVerificationCode[0:len((*c.CallOptions).SendNotificationChannelVerificationCode):len((*c.CallOptions).SendNotificationChannelVerificationCode)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -670,12 +601,7 @@ func (c *notificationChannelGRPCClient) GetNotificatio
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/GetNotificationChannelVerificationCode")
-	}
+
 	opts = append((*c.CallOptions).GetNotificationChannelVerificationCode[0:len((*c.CallOptions).GetNotificationChannelVerificationCode):len((*c.CallOptions).GetNotificationChannelVerificationCode)], opts...)
 	var resp *monitoringpb.GetNotificationChannelVerificationCodeResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -694,12 +620,7 @@ func (c *notificationChannelGRPCClient) VerifyNotifica
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.NotificationChannelService/VerifyNotificationChannel")
-	}
+
 	opts = append((*c.CallOptions).VerifyNotificationChannel[0:len((*c.CallOptions).VerifyNotificationChannel):len((*c.CallOptions).VerifyNotificationChannel)], opts...)
 	var resp *monitoringpb.NotificationChannel
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
