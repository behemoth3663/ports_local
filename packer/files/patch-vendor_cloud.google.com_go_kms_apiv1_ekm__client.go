--- vendor/cloud.google.com/go/kms/apiv1/ekm_client.go.orig	2026-10-02 20:06:38 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/ekm_client.go
@@ -30,7 +30,6 @@ import (
 	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -402,16 +401,7 @@ func NewEkmClient(ctx context.Context, opts ...option.
 //	EkmConnection
 func NewEkmClient(ctx context.Context, opts ...option.ClientOption) (*EkmClient, error) {
 	clientOpts := defaultEkmGRPCClientOptions()
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "cloudkms",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/kms/apiv1",
-			"gcp.client.language": "go",
-			"url.domain":          "cloudkms.googleapis.com",
-		}))
-	}
+
 	if newEkmClientHook != nil {
 		hookOpts, err := newEkmClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -436,33 +426,7 @@ func NewEkmClient(ctx context.Context, opts ...option.
 		locationsClient:  locationpb.NewLocationsClient(connPool),
 	}
 	c.setGoogleClientInfo()
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "cloudkms",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/kms/apiv1",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "cloudkms.googleapis.com",
-			}),
-		)
 
-		client.CallOptions.ListEkmConnections = append(client.CallOptions.ListEkmConnections, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetEkmConnection = append(client.CallOptions.GetEkmConnection, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateEkmConnection = append(client.CallOptions.CreateEkmConnection, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateEkmConnection = append(client.CallOptions.UpdateEkmConnection, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetEkmConfig = append(client.CallOptions.GetEkmConfig, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateEkmConfig = append(client.CallOptions.UpdateEkmConfig, gax.WithClientMetrics(metrics))
-		client.CallOptions.VerifyConnectivity = append(client.CallOptions.VerifyConnectivity, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetOperation = append(client.CallOptions.GetOperation, gax.WithClientMetrics(metrics))
-	}
-
 	client.internalClient = c
 
 	return &client, nil
@@ -520,16 +484,7 @@ func NewEkmRESTClient(ctx context.Context, opts ...opt
 //	EkmConnection
 func NewEkmRESTClient(ctx context.Context, opts ...option.ClientOption) (*EkmClient, error) {
 	clientOpts := append(defaultEkmRESTClientOptions(), opts...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "cloudkms",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/kms/apiv1",
-			"gcp.client.language": "go",
-			"url.domain":          "cloudkms.googleapis.com",
-		}))
-	}
+
 	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
 	if err != nil {
 		return nil, err
@@ -544,33 +499,6 @@ func NewEkmRESTClient(ctx context.Context, opts ...opt
 	}
 	c.setGoogleClientInfo()
 
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "cloudkms",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/kms/apiv1",
-				gax.RPCSystem:      "http",
-				gax.URLDomain:      "cloudkms.googleapis.com",
-			}),
-		)
-
-		callOpts.ListEkmConnections = append(callOpts.ListEkmConnections, gax.WithClientMetrics(metrics))
-		callOpts.GetEkmConnection = append(callOpts.GetEkmConnection, gax.WithClientMetrics(metrics))
-		callOpts.CreateEkmConnection = append(callOpts.CreateEkmConnection, gax.WithClientMetrics(metrics))
-		callOpts.UpdateEkmConnection = append(callOpts.UpdateEkmConnection, gax.WithClientMetrics(metrics))
-		callOpts.GetEkmConfig = append(callOpts.GetEkmConfig, gax.WithClientMetrics(metrics))
-		callOpts.UpdateEkmConfig = append(callOpts.UpdateEkmConfig, gax.WithClientMetrics(metrics))
-		callOpts.VerifyConnectivity = append(callOpts.VerifyConnectivity, gax.WithClientMetrics(metrics))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientMetrics(metrics))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientMetrics(metrics))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientMetrics(metrics))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientMetrics(metrics))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientMetrics(metrics))
-		callOpts.GetOperation = append(callOpts.GetOperation, gax.WithClientMetrics(metrics))
-	}
-
 	return &EkmClient{internalClient: c, CallOptions: callOpts}, nil
 }
 
@@ -616,12 +544,7 @@ func (c *ekmGRPCClient) ListEkmConnections(ctx context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/ListEkmConnections")
-	}
+
 	opts = append((*c.CallOptions).ListEkmConnections[0:len((*c.CallOptions).ListEkmConnections):len((*c.CallOptions).ListEkmConnections)], opts...)
 	it := &EkmConnectionIterator{}
 	req = proto.CloneOf(req)
@@ -668,12 +591,7 @@ func (c *ekmGRPCClient) GetEkmConnection(ctx context.C
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConnection")
-	}
+
 	opts = append((*c.CallOptions).GetEkmConnection[0:len((*c.CallOptions).GetEkmConnection):len((*c.CallOptions).GetEkmConnection)], opts...)
 	var resp *kmspb.EkmConnection
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -692,12 +610,7 @@ func (c *ekmGRPCClient) CreateEkmConnection(ctx contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/CreateEkmConnection")
-	}
+
 	opts = append((*c.CallOptions).CreateEkmConnection[0:len((*c.CallOptions).CreateEkmConnection):len((*c.CallOptions).CreateEkmConnection)], opts...)
 	var resp *kmspb.EkmConnection
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -716,9 +629,7 @@ func (c *ekmGRPCClient) UpdateEkmConnection(ctx contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/UpdateEkmConnection")
-	}
+
 	opts = append((*c.CallOptions).UpdateEkmConnection[0:len((*c.CallOptions).UpdateEkmConnection):len((*c.CallOptions).UpdateEkmConnection)], opts...)
 	var resp *kmspb.EkmConnection
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -737,12 +648,7 @@ func (c *ekmGRPCClient) GetEkmConfig(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConfig")
-	}
+
 	opts = append((*c.CallOptions).GetEkmConfig[0:len((*c.CallOptions).GetEkmConfig):len((*c.CallOptions).GetEkmConfig)], opts...)
 	var resp *kmspb.EkmConfig
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -761,9 +667,7 @@ func (c *ekmGRPCClient) UpdateEkmConfig(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/UpdateEkmConfig")
-	}
+
 	opts = append((*c.CallOptions).UpdateEkmConfig[0:len((*c.CallOptions).UpdateEkmConfig):len((*c.CallOptions).UpdateEkmConfig)], opts...)
 	var resp *kmspb.EkmConfig
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -782,12 +686,7 @@ func (c *ekmGRPCClient) VerifyConnectivity(ctx context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/VerifyConnectivity")
-	}
+
 	opts = append((*c.CallOptions).VerifyConnectivity[0:len((*c.CallOptions).VerifyConnectivity):len((*c.CallOptions).VerifyConnectivity)], opts...)
 	var resp *kmspb.VerifyConnectivityResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -806,9 +705,7 @@ func (c *ekmGRPCClient) GetLocation(ctx context.Contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-	}
+
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	var resp *locationpb.Location
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -827,9 +724,7 @@ func (c *ekmGRPCClient) ListLocations(ctx context.Cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/ListLocations")
-	}
+
 	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
 	it := &LocationIterator{}
 	req = proto.CloneOf(req)
@@ -876,12 +771,7 @@ func (c *ekmGRPCClient) GetIamPolicy(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//iam-meta-api.googleapis.com/%v", req.GetResource()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/GetIamPolicy")
-	}
+
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -900,12 +790,7 @@ func (c *ekmGRPCClient) SetIamPolicy(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//iam-meta-api.googleapis.com/%v", req.GetResource()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/SetIamPolicy")
-	}
+
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -924,12 +809,7 @@ func (c *ekmGRPCClient) TestIamPermissions(ctx context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//iam-meta-api.googleapis.com/%v", req.GetResource()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/TestIamPermissions")
-	}
+
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	var resp *iampb.TestIamPermissionsResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -948,9 +828,7 @@ func (c *ekmGRPCClient) GetOperation(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1068,13 +946,7 @@ func (c *ekmRESTClient) GetEkmConnection(ctx context.C
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConnection")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/ekmConnections/*}")
-	}
+
 	opts = append((*c.CallOptions).GetEkmConnection[0:len((*c.CallOptions).GetEkmConnection):len((*c.CallOptions).GetEkmConnection)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConnection{}
@@ -1134,13 +1006,7 @@ func (c *ekmRESTClient) CreateEkmConnection(ctx contex
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/CreateEkmConnection")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*}/ekmConnections")
-	}
+
 	opts = append((*c.CallOptions).CreateEkmConnection[0:len((*c.CallOptions).CreateEkmConnection):len((*c.CallOptions).CreateEkmConnection)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConnection{}
@@ -1205,10 +1071,7 @@ func (c *ekmRESTClient) UpdateEkmConnection(ctx contex
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/UpdateEkmConnection")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{ekm_connection.name=projects/*/locations/*/ekmConnections/*}")
-	}
+
 	opts = append((*c.CallOptions).UpdateEkmConnection[0:len((*c.CallOptions).UpdateEkmConnection):len((*c.CallOptions).UpdateEkmConnection)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConnection{}
@@ -1260,13 +1123,7 @@ func (c *ekmRESTClient) GetEkmConfig(ctx context.Conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConfig")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/ekmConfig}")
-	}
+
 	opts = append((*c.CallOptions).GetEkmConfig[0:len((*c.CallOptions).GetEkmConfig):len((*c.CallOptions).GetEkmConfig)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConfig{}
@@ -1332,10 +1189,7 @@ func (c *ekmRESTClient) UpdateEkmConfig(ctx context.Co
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/UpdateEkmConfig")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{ekm_config.name=projects/*/locations/*/ekmConfig}")
-	}
+
 	opts = append((*c.CallOptions).UpdateEkmConfig[0:len((*c.CallOptions).UpdateEkmConfig):len((*c.CallOptions).UpdateEkmConfig)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConfig{}
@@ -1390,13 +1244,7 @@ func (c *ekmRESTClient) VerifyConnectivity(ctx context
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/VerifyConnectivity")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/ekmConnections/*}:verifyConnectivity")
-	}
+
 	opts = append((*c.CallOptions).VerifyConnectivity[0:len((*c.CallOptions).VerifyConnectivity):len((*c.CallOptions).VerifyConnectivity)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.VerifyConnectivityResponse{}
@@ -1447,10 +1295,7 @@ func (c *ekmRESTClient) GetLocation(ctx context.Contex
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*}")
-	}
+
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &locationpb.Location{}
@@ -1602,13 +1447,7 @@ func (c *ekmRESTClient) GetIamPolicy(ctx context.Conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//iam-meta-api.googleapis.com/%v", req.GetResource()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/GetIamPolicy")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/locations/*/keyRings/*}:getIamPolicy")
-	}
+
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.Policy{}
@@ -1669,13 +1508,7 @@ func (c *ekmRESTClient) SetIamPolicy(ctx context.Conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//iam-meta-api.googleapis.com/%v", req.GetResource()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/SetIamPolicy")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/locations/*/keyRings/*}:setIamPolicy")
-	}
+
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.Policy{}
@@ -1738,13 +1571,7 @@ func (c *ekmRESTClient) TestIamPermissions(ctx context
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//iam-meta-api.googleapis.com/%v", req.GetResource()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/TestIamPermissions")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/locations/*/keyRings/*}:testIamPermissions")
-	}
+
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.TestIamPermissionsResponse{}
@@ -1795,10 +1622,7 @@ func (c *ekmRESTClient) GetOperation(ctx context.Conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/operations/*}")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
