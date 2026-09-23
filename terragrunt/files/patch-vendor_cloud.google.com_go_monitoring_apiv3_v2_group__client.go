diff --git a/vendor/cloud.google.com/go/monitoring/apiv3/v2/group_client.go b/vendor/cloud.google.com/go/monitoring/apiv3/v2/group_client.go
index 17fe3d14c..ca16d5f78 100644
--- vendor/cloud.google.com/go/monitoring/apiv3/v2/group_client.go.orig
+++ vendor/cloud.google.com/go/monitoring/apiv3/v2/group_client.go
@@ -259,16 +259,7 @@ type groupGRPCClient struct {
 // from the infrastructure.
 func NewGroupClient(ctx context.Context, opts ...option.ClientOption) (*GroupClient, error) {
 	clientOpts := defaultGroupGRPCClientOptions()
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
 	if newGroupClientHook != nil {
 		hookOpts, err := newGroupClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -290,25 +281,6 @@ func NewGroupClient(ctx context.Context, opts ...option.ClientOption) (*GroupCli
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
-
-		client.CallOptions.ListGroups = append(client.CallOptions.ListGroups, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetGroup = append(client.CallOptions.GetGroup, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateGroup = append(client.CallOptions.CreateGroup, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateGroup = append(client.CallOptions.UpdateGroup, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteGroup = append(client.CallOptions.DeleteGroup, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListGroupMembers = append(client.CallOptions.ListGroupMembers, gax.WithClientMetrics(metrics))
-	}
 
 	client.internalClient = c
 
@@ -345,12 +317,7 @@ func (c *groupGRPCClient) ListGroups(ctx context.Context, req *monitoringpb.List
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.GroupService/ListGroups")
-	}
+
 	opts = append((*c.CallOptions).ListGroups[0:len((*c.CallOptions).ListGroups):len((*c.CallOptions).ListGroups)], opts...)
 	it := &GroupIterator{}
 	req = proto.CloneOf(req)
@@ -397,12 +364,7 @@ func (c *groupGRPCClient) GetGroup(ctx context.Context, req *monitoringpb.GetGro
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.GroupService/GetGroup")
-	}
+
 	opts = append((*c.CallOptions).GetGroup[0:len((*c.CallOptions).GetGroup):len((*c.CallOptions).GetGroup)], opts...)
 	var resp *monitoringpb.Group
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -421,12 +383,7 @@ func (c *groupGRPCClient) CreateGroup(ctx context.Context, req *monitoringpb.Cre
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.GroupService/CreateGroup")
-	}
+
 	opts = append((*c.CallOptions).CreateGroup[0:len((*c.CallOptions).CreateGroup):len((*c.CallOptions).CreateGroup)], opts...)
 	var resp *monitoringpb.Group
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -445,9 +402,7 @@ func (c *groupGRPCClient) UpdateGroup(ctx context.Context, req *monitoringpb.Upd
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.GroupService/UpdateGroup")
-	}
+
 	opts = append((*c.CallOptions).UpdateGroup[0:len((*c.CallOptions).UpdateGroup):len((*c.CallOptions).UpdateGroup)], opts...)
 	var resp *monitoringpb.Group
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -466,12 +421,7 @@ func (c *groupGRPCClient) DeleteGroup(ctx context.Context, req *monitoringpb.Del
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.GroupService/DeleteGroup")
-	}
+
 	opts = append((*c.CallOptions).DeleteGroup[0:len((*c.CallOptions).DeleteGroup):len((*c.CallOptions).DeleteGroup)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -486,12 +436,7 @@ func (c *groupGRPCClient) ListGroupMembers(ctx context.Context, req *monitoringp
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//monitoring.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.monitoring.v3.GroupService/ListGroupMembers")
-	}
+
 	opts = append((*c.CallOptions).ListGroupMembers[0:len((*c.CallOptions).ListGroupMembers):len((*c.CallOptions).ListGroupMembers)], opts...)
 	it := &MonitoredResourceIterator{}
 	req = proto.CloneOf(req)
