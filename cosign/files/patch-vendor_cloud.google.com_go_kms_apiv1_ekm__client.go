--- vendor/cloud.google.com/go/kms/apiv1/ekm_client.go.orig	2026-11-26 19:15:42 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/ekm_client.go
@@ -30,7 +30,6 @@ import (
 	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -401,16 +400,7 @@ func NewEkmClient(ctx context.Context, opts ...option.
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
@@ -435,84 +425,7 @@ func NewEkmClient(ctx context.Context, opts ...option.
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
-	if gax.IsFeatureEnabled("TRACING") {
-		tracing := gax.NewClientTracing(
-			gax.WithTracingAttributes(map[string]string{
-				gax.ClientService:  "cloudkms",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/kms/apiv1",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "cloudkms.googleapis.com",
-			}),
-		)
-
-		client.CallOptions.ListEkmConnections = append(client.CallOptions.ListEkmConnections, gax.WithClientTracing(tracing))
-		client.CallOptions.GetEkmConnection = append(client.CallOptions.GetEkmConnection, gax.WithClientTracing(tracing))
-		client.CallOptions.CreateEkmConnection = append(client.CallOptions.CreateEkmConnection, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateEkmConnection = append(client.CallOptions.UpdateEkmConnection, gax.WithClientTracing(tracing))
-		client.CallOptions.GetEkmConfig = append(client.CallOptions.GetEkmConfig, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateEkmConfig = append(client.CallOptions.UpdateEkmConfig, gax.WithClientTracing(tracing))
-		client.CallOptions.VerifyConnectivity = append(client.CallOptions.VerifyConnectivity, gax.WithClientTracing(tracing))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientTracing(tracing))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientTracing(tracing))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientTracing(tracing))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientTracing(tracing))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientTracing(tracing))
-		client.CallOptions.GetOperation = append(client.CallOptions.GetOperation, gax.WithClientTracing(tracing))
-	}
-	if gax.IsFeatureEnabled("LOGGING") {
-		logging := gax.NewClientLogging(
-			gax.WithLoggerProvider(c.logger),
-			gax.WithLoggingAttributes(map[string]string{
-				gax.ClientService:  "cloudkms",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/kms/apiv1",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "cloudkms.googleapis.com",
-			}),
-		)
-
-		client.CallOptions.ListEkmConnections = append(client.CallOptions.ListEkmConnections, gax.WithClientLogging(logging))
-		client.CallOptions.GetEkmConnection = append(client.CallOptions.GetEkmConnection, gax.WithClientLogging(logging))
-		client.CallOptions.CreateEkmConnection = append(client.CallOptions.CreateEkmConnection, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateEkmConnection = append(client.CallOptions.UpdateEkmConnection, gax.WithClientLogging(logging))
-		client.CallOptions.GetEkmConfig = append(client.CallOptions.GetEkmConfig, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateEkmConfig = append(client.CallOptions.UpdateEkmConfig, gax.WithClientLogging(logging))
-		client.CallOptions.VerifyConnectivity = append(client.CallOptions.VerifyConnectivity, gax.WithClientLogging(logging))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientLogging(logging))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientLogging(logging))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientLogging(logging))
-		client.CallOptions.GetOperation = append(client.CallOptions.GetOperation, gax.WithClientLogging(logging))
-	}
-
 	client.internalClient = c
 
 	return &client, nil
@@ -570,16 +483,7 @@ func NewEkmRESTClient(ctx context.Context, opts ...opt
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
@@ -594,84 +498,6 @@ func NewEkmRESTClient(ctx context.Context, opts ...opt
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
-	if gax.IsFeatureEnabled("TRACING") {
-		tracing := gax.NewClientTracing(
-			gax.WithTracingAttributes(map[string]string{
-				gax.ClientService:  "cloudkms",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/kms/apiv1",
-				gax.RPCSystem:      "http",
-				gax.URLDomain:      "cloudkms.googleapis.com",
-			}),
-		)
-
-		callOpts.ListEkmConnections = append(callOpts.ListEkmConnections, gax.WithClientTracing(tracing))
-		callOpts.GetEkmConnection = append(callOpts.GetEkmConnection, gax.WithClientTracing(tracing))
-		callOpts.CreateEkmConnection = append(callOpts.CreateEkmConnection, gax.WithClientTracing(tracing))
-		callOpts.UpdateEkmConnection = append(callOpts.UpdateEkmConnection, gax.WithClientTracing(tracing))
-		callOpts.GetEkmConfig = append(callOpts.GetEkmConfig, gax.WithClientTracing(tracing))
-		callOpts.UpdateEkmConfig = append(callOpts.UpdateEkmConfig, gax.WithClientTracing(tracing))
-		callOpts.VerifyConnectivity = append(callOpts.VerifyConnectivity, gax.WithClientTracing(tracing))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientTracing(tracing))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientTracing(tracing))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientTracing(tracing))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientTracing(tracing))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientTracing(tracing))
-		callOpts.GetOperation = append(callOpts.GetOperation, gax.WithClientTracing(tracing))
-	}
-	if gax.IsFeatureEnabled("LOGGING") {
-		logging := gax.NewClientLogging(
-			gax.WithLoggerProvider(c.logger),
-			gax.WithLoggingAttributes(map[string]string{
-				gax.ClientService:  "cloudkms",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/kms/apiv1",
-				gax.RPCSystem:      "http",
-				gax.URLDomain:      "cloudkms.googleapis.com",
-			}),
-		)
-
-		callOpts.ListEkmConnections = append(callOpts.ListEkmConnections, gax.WithClientLogging(logging))
-		callOpts.GetEkmConnection = append(callOpts.GetEkmConnection, gax.WithClientLogging(logging))
-		callOpts.CreateEkmConnection = append(callOpts.CreateEkmConnection, gax.WithClientLogging(logging))
-		callOpts.UpdateEkmConnection = append(callOpts.UpdateEkmConnection, gax.WithClientLogging(logging))
-		callOpts.GetEkmConfig = append(callOpts.GetEkmConfig, gax.WithClientLogging(logging))
-		callOpts.UpdateEkmConfig = append(callOpts.UpdateEkmConfig, gax.WithClientLogging(logging))
-		callOpts.VerifyConnectivity = append(callOpts.VerifyConnectivity, gax.WithClientLogging(logging))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientLogging(logging))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientLogging(logging))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientLogging(logging))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientLogging(logging))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientLogging(logging))
-		callOpts.GetOperation = append(callOpts.GetOperation, gax.WithClientLogging(logging))
-	}
-
 	return &EkmClient{internalClient: c, CallOptions: callOpts}, nil
 }
 
@@ -717,9 +543,6 @@ func (c *ekmGRPCClient) ListEkmConnections(ctx context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/ListEkmConnections")
-	}
 	opts = append((*c.CallOptions).ListEkmConnections[0:len((*c.CallOptions).ListEkmConnections):len((*c.CallOptions).ListEkmConnections)], opts...)
 	it := &EkmConnectionIterator{}
 	req = proto.CloneOf(req)
@@ -766,9 +589,6 @@ func (c *ekmGRPCClient) GetEkmConnection(ctx context.C
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConnection")
-	}
 	opts = append((*c.CallOptions).GetEkmConnection[0:len((*c.CallOptions).GetEkmConnection):len((*c.CallOptions).GetEkmConnection)], opts...)
 	var resp *kmspb.EkmConnection
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -787,9 +607,6 @@ func (c *ekmGRPCClient) CreateEkmConnection(ctx contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/CreateEkmConnection")
-	}
 	opts = append((*c.CallOptions).CreateEkmConnection[0:len((*c.CallOptions).CreateEkmConnection):len((*c.CallOptions).CreateEkmConnection)], opts...)
 	var resp *kmspb.EkmConnection
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -808,9 +625,7 @@ func (c *ekmGRPCClient) UpdateEkmConnection(ctx contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/UpdateEkmConnection")
-	}
+
 	opts = append((*c.CallOptions).UpdateEkmConnection[0:len((*c.CallOptions).UpdateEkmConnection):len((*c.CallOptions).UpdateEkmConnection)], opts...)
 	var resp *kmspb.EkmConnection
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -829,9 +644,6 @@ func (c *ekmGRPCClient) GetEkmConfig(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConfig")
-	}
 	opts = append((*c.CallOptions).GetEkmConfig[0:len((*c.CallOptions).GetEkmConfig):len((*c.CallOptions).GetEkmConfig)], opts...)
 	var resp *kmspb.EkmConfig
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -850,9 +662,7 @@ func (c *ekmGRPCClient) UpdateEkmConfig(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/UpdateEkmConfig")
-	}
+
 	opts = append((*c.CallOptions).UpdateEkmConfig[0:len((*c.CallOptions).UpdateEkmConfig):len((*c.CallOptions).UpdateEkmConfig)], opts...)
 	var resp *kmspb.EkmConfig
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -871,9 +681,6 @@ func (c *ekmGRPCClient) VerifyConnectivity(ctx context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/VerifyConnectivity")
-	}
 	opts = append((*c.CallOptions).VerifyConnectivity[0:len((*c.CallOptions).VerifyConnectivity):len((*c.CallOptions).VerifyConnectivity)], opts...)
 	var resp *kmspb.VerifyConnectivityResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -892,9 +699,7 @@ func (c *ekmGRPCClient) GetLocation(ctx context.Contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-	}
+
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	var resp *locationpb.Location
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -913,9 +718,7 @@ func (c *ekmGRPCClient) ListLocations(ctx context.Cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/ListLocations")
-	}
+
 	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
 	it := &LocationIterator{}
 	req = proto.CloneOf(req)
@@ -962,9 +765,6 @@ func (c *ekmGRPCClient) GetIamPolicy(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/GetIamPolicy")
-	}
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -983,9 +783,6 @@ func (c *ekmGRPCClient) SetIamPolicy(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/SetIamPolicy")
-	}
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1004,9 +801,6 @@ func (c *ekmGRPCClient) TestIamPermissions(ctx context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/TestIamPermissions")
-	}
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	var resp *iampb.TestIamPermissionsResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1025,9 +819,7 @@ func (c *ekmGRPCClient) GetOperation(ctx context.Conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1145,10 +937,6 @@ func (c *ekmRESTClient) GetEkmConnection(ctx context.C
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConnection")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/ekmConnections/*}")
-	}
 	opts = append((*c.CallOptions).GetEkmConnection[0:len((*c.CallOptions).GetEkmConnection):len((*c.CallOptions).GetEkmConnection)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConnection{}
@@ -1208,10 +996,6 @@ func (c *ekmRESTClient) CreateEkmConnection(ctx contex
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/CreateEkmConnection")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*}/ekmConnections")
-	}
 	opts = append((*c.CallOptions).CreateEkmConnection[0:len((*c.CallOptions).CreateEkmConnection):len((*c.CallOptions).CreateEkmConnection)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConnection{}
@@ -1276,10 +1060,7 @@ func (c *ekmRESTClient) UpdateEkmConnection(ctx contex
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
@@ -1331,10 +1112,6 @@ func (c *ekmRESTClient) GetEkmConfig(ctx context.Conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/GetEkmConfig")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/ekmConfig}")
-	}
 	opts = append((*c.CallOptions).GetEkmConfig[0:len((*c.CallOptions).GetEkmConfig):len((*c.CallOptions).GetEkmConfig)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EkmConfig{}
@@ -1400,10 +1177,7 @@ func (c *ekmRESTClient) UpdateEkmConfig(ctx context.Co
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
@@ -1458,10 +1232,6 @@ func (c *ekmRESTClient) VerifyConnectivity(ctx context
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.EkmService/VerifyConnectivity")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/ekmConnections/*}:verifyConnectivity")
-	}
 	opts = append((*c.CallOptions).VerifyConnectivity[0:len((*c.CallOptions).VerifyConnectivity):len((*c.CallOptions).VerifyConnectivity)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.VerifyConnectivityResponse{}
@@ -1512,10 +1282,7 @@ func (c *ekmRESTClient) GetLocation(ctx context.Contex
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
@@ -1666,10 +1433,6 @@ func (c *ekmRESTClient) GetIamPolicy(ctx context.Conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/GetIamPolicy")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/locations/*/keyRings/*}:getIamPolicy")
-	}
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.Policy{}
@@ -1730,10 +1493,6 @@ func (c *ekmRESTClient) SetIamPolicy(ctx context.Conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/SetIamPolicy")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/locations/*/keyRings/*}:setIamPolicy")
-	}
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.Policy{}
@@ -1796,10 +1555,6 @@ func (c *ekmRESTClient) TestIamPermissions(ctx context
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/TestIamPermissions")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/locations/*/keyRings/*}:testIamPermissions")
-	}
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.TestIamPermissionsResponse{}
@@ -1850,10 +1605,7 @@ func (c *ekmRESTClient) GetOperation(ctx context.Conte
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
