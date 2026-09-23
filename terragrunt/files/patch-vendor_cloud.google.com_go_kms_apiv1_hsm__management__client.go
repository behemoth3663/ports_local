--- vendor/cloud.google.com/go/kms/apiv1/hsm_management_client.go.orig	2026-09-23 21:56:03 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/hsm_management_client.go
@@ -32,8 +32,6 @@ import (
 	lroauto "cloud.google.com/go/longrunning/autogen"
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
-	trace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -590,16 +588,7 @@ func NewHsmManagementClient(ctx context.Context, opts 
 //	SingleTenantHsmInstanceProposal
 func NewHsmManagementClient(ctx context.Context, opts ...option.ClientOption) (*HsmManagementClient, error) {
 	clientOpts := defaultHsmManagementGRPCClientOptions()
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
 	if newHsmManagementClientHook != nil {
 		hookOpts, err := newHsmManagementClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -624,35 +613,7 @@ func NewHsmManagementClient(ctx context.Context, opts 
 		locationsClient:     locationpb.NewLocationsClient(connPool),
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
 
-		client.CallOptions.ListSingleTenantHsmInstances = append(client.CallOptions.ListSingleTenantHsmInstances, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetSingleTenantHsmInstance = append(client.CallOptions.GetSingleTenantHsmInstance, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateSingleTenantHsmInstance = append(client.CallOptions.CreateSingleTenantHsmInstance, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateSingleTenantHsmInstanceProposal = append(client.CallOptions.CreateSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		client.CallOptions.ApproveSingleTenantHsmInstanceProposal = append(client.CallOptions.ApproveSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		client.CallOptions.ExecuteSingleTenantHsmInstanceProposal = append(client.CallOptions.ExecuteSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetSingleTenantHsmInstanceProposal = append(client.CallOptions.GetSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListSingleTenantHsmInstanceProposals = append(client.CallOptions.ListSingleTenantHsmInstanceProposals, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteSingleTenantHsmInstanceProposal = append(client.CallOptions.DeleteSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetOperation = append(client.CallOptions.GetOperation, gax.WithClientMetrics(metrics))
-	}
-
 	client.internalClient = c
 
 	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
@@ -729,16 +690,7 @@ func NewHsmManagementRESTClient(ctx context.Context, o
 //	SingleTenantHsmInstanceProposal
 func NewHsmManagementRESTClient(ctx context.Context, opts ...option.ClientOption) (*HsmManagementClient, error) {
 	clientOpts := append(defaultHsmManagementRESTClientOptions(), opts...)
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
@@ -753,35 +705,6 @@ func NewHsmManagementRESTClient(ctx context.Context, o
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
-		callOpts.ListSingleTenantHsmInstances = append(callOpts.ListSingleTenantHsmInstances, gax.WithClientMetrics(metrics))
-		callOpts.GetSingleTenantHsmInstance = append(callOpts.GetSingleTenantHsmInstance, gax.WithClientMetrics(metrics))
-		callOpts.CreateSingleTenantHsmInstance = append(callOpts.CreateSingleTenantHsmInstance, gax.WithClientMetrics(metrics))
-		callOpts.CreateSingleTenantHsmInstanceProposal = append(callOpts.CreateSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		callOpts.ApproveSingleTenantHsmInstanceProposal = append(callOpts.ApproveSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		callOpts.ExecuteSingleTenantHsmInstanceProposal = append(callOpts.ExecuteSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		callOpts.GetSingleTenantHsmInstanceProposal = append(callOpts.GetSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		callOpts.ListSingleTenantHsmInstanceProposals = append(callOpts.ListSingleTenantHsmInstanceProposals, gax.WithClientMetrics(metrics))
-		callOpts.DeleteSingleTenantHsmInstanceProposal = append(callOpts.DeleteSingleTenantHsmInstanceProposal, gax.WithClientMetrics(metrics))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientMetrics(metrics))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientMetrics(metrics))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientMetrics(metrics))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientMetrics(metrics))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientMetrics(metrics))
-		callOpts.GetOperation = append(callOpts.GetOperation, gax.WithClientMetrics(metrics))
-	}
-
 	lroOpts := []option.ClientOption{
 		option.WithHTTPClient(httpClient),
 		option.WithEndpoint(endpoint),
@@ -837,12 +760,7 @@ func (c *hsmManagementGRPCClient) ListSingleTenantHsmI
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/ListSingleTenantHsmInstances")
-	}
+
 	opts = append((*c.CallOptions).ListSingleTenantHsmInstances[0:len((*c.CallOptions).ListSingleTenantHsmInstances):len((*c.CallOptions).ListSingleTenantHsmInstances)], opts...)
 	it := &SingleTenantHsmInstanceIterator{}
 	req = proto.CloneOf(req)
@@ -889,12 +807,7 @@ func (c *hsmManagementGRPCClient) GetSingleTenantHsmIn
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/GetSingleTenantHsmInstance")
-	}
+
 	opts = append((*c.CallOptions).GetSingleTenantHsmInstance[0:len((*c.CallOptions).GetSingleTenantHsmInstance):len((*c.CallOptions).GetSingleTenantHsmInstance)], opts...)
 	var resp *kmspb.SingleTenantHsmInstance
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -913,12 +826,7 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/CreateSingleTenantHsmInstance")
-	}
+
 	opts = append((*c.CallOptions).CreateSingleTenantHsmInstance[0:len((*c.CallOptions).CreateSingleTenantHsmInstance):len((*c.CallOptions).CreateSingleTenantHsmInstance)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -930,9 +838,6 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 		return nil, err
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &CreateSingleTenantHsmInstanceOperation{
 		lro: lro,
 	}, nil
@@ -943,12 +848,7 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/CreateSingleTenantHsmInstanceProposal")
-	}
+
 	opts = append((*c.CallOptions).CreateSingleTenantHsmInstanceProposal[0:len((*c.CallOptions).CreateSingleTenantHsmInstanceProposal):len((*c.CallOptions).CreateSingleTenantHsmInstanceProposal)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -960,9 +860,6 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 		return nil, err
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceProposalOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &CreateSingleTenantHsmInstanceProposalOperation{
 		lro: lro,
 	}, nil
@@ -973,12 +870,7 @@ func (c *hsmManagementGRPCClient) ApproveSingleTenantH
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/ApproveSingleTenantHsmInstanceProposal")
-	}
+
 	opts = append((*c.CallOptions).ApproveSingleTenantHsmInstanceProposal[0:len((*c.CallOptions).ApproveSingleTenantHsmInstanceProposal):len((*c.CallOptions).ApproveSingleTenantHsmInstanceProposal)], opts...)
 	var resp *kmspb.ApproveSingleTenantHsmInstanceProposalResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -997,12 +889,7 @@ func (c *hsmManagementGRPCClient) ExecuteSingleTenantH
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/ExecuteSingleTenantHsmInstanceProposal")
-	}
+
 	opts = append((*c.CallOptions).ExecuteSingleTenantHsmInstanceProposal[0:len((*c.CallOptions).ExecuteSingleTenantHsmInstanceProposal):len((*c.CallOptions).ExecuteSingleTenantHsmInstanceProposal)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1014,9 +901,6 @@ func (c *hsmManagementGRPCClient) ExecuteSingleTenantH
 		return nil, err
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.ExecuteSingleTenantHsmInstanceProposalOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &ExecuteSingleTenantHsmInstanceProposalOperation{
 		lro: lro,
 	}, nil
@@ -1027,12 +911,7 @@ func (c *hsmManagementGRPCClient) GetSingleTenantHsmIn
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/GetSingleTenantHsmInstanceProposal")
-	}
+
 	opts = append((*c.CallOptions).GetSingleTenantHsmInstanceProposal[0:len((*c.CallOptions).GetSingleTenantHsmInstanceProposal):len((*c.CallOptions).GetSingleTenantHsmInstanceProposal)], opts...)
 	var resp *kmspb.SingleTenantHsmInstanceProposal
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1051,12 +930,7 @@ func (c *hsmManagementGRPCClient) ListSingleTenantHsmI
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/ListSingleTenantHsmInstanceProposals")
-	}
+
 	opts = append((*c.CallOptions).ListSingleTenantHsmInstanceProposals[0:len((*c.CallOptions).ListSingleTenantHsmInstanceProposals):len((*c.CallOptions).ListSingleTenantHsmInstanceProposals)], opts...)
 	it := &SingleTenantHsmInstanceProposalIterator{}
 	req = proto.CloneOf(req)
@@ -1103,12 +977,7 @@ func (c *hsmManagementGRPCClient) DeleteSingleTenantHs
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/DeleteSingleTenantHsmInstanceProposal")
-	}
+
 	opts = append((*c.CallOptions).DeleteSingleTenantHsmInstanceProposal[0:len((*c.CallOptions).DeleteSingleTenantHsmInstanceProposal):len((*c.CallOptions).DeleteSingleTenantHsmInstanceProposal)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -1123,9 +992,7 @@ func (c *hsmManagementGRPCClient) GetLocation(ctx cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-	}
+
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	var resp *locationpb.Location
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1144,9 +1011,7 @@ func (c *hsmManagementGRPCClient) ListLocations(ctx co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/ListLocations")
-	}
+
 	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
 	it := &LocationIterator{}
 	req = proto.CloneOf(req)
@@ -1193,12 +1058,7 @@ func (c *hsmManagementGRPCClient) GetIamPolicy(ctx con
 
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
@@ -1217,12 +1077,7 @@ func (c *hsmManagementGRPCClient) SetIamPolicy(ctx con
 
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
@@ -1241,12 +1096,7 @@ func (c *hsmManagementGRPCClient) TestIamPermissions(c
 
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
@@ -1265,9 +1115,7 @@ func (c *hsmManagementGRPCClient) GetOperation(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1389,13 +1237,7 @@ func (c *hsmManagementRESTClient) GetSingleTenantHsmIn
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/GetSingleTenantHsmInstance")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/singleTenantHsmInstances/*}")
-	}
+
 	opts = append((*c.CallOptions).GetSingleTenantHsmInstance[0:len((*c.CallOptions).GetSingleTenantHsmInstance):len((*c.CallOptions).GetSingleTenantHsmInstance)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.SingleTenantHsmInstance{}
@@ -1460,13 +1302,7 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/CreateSingleTenantHsmInstance")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*}/singleTenantHsmInstances")
-	}
+
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
 	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1496,9 +1332,6 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
 
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &CreateSingleTenantHsmInstanceOperation{
 		lro:      lro,
 		pollPath: override,
@@ -1537,13 +1370,7 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/CreateSingleTenantHsmInstanceProposal")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/singleTenantHsmInstances/*}/proposals")
-	}
+
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
 	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1573,9 +1400,6 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
 
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceProposalOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &CreateSingleTenantHsmInstanceProposalOperation{
 		lro:      lro,
 		pollPath: override,
@@ -1613,13 +1437,7 @@ func (c *hsmManagementRESTClient) ApproveSingleTenantH
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/ApproveSingleTenantHsmInstanceProposal")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/singleTenantHsmInstances/*/proposals/*}:approve")
-	}
+
 	opts = append((*c.CallOptions).ApproveSingleTenantHsmInstanceProposal[0:len((*c.CallOptions).ApproveSingleTenantHsmInstanceProposal):len((*c.CallOptions).ApproveSingleTenantHsmInstanceProposal)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.ApproveSingleTenantHsmInstanceProposalResponse{}
@@ -1682,13 +1500,7 @@ func (c *hsmManagementRESTClient) ExecuteSingleTenantH
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/ExecuteSingleTenantHsmInstanceProposal")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/singleTenantHsmInstances/*/proposals/*}:execute")
-	}
+
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
 	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1718,9 +1530,6 @@ func (c *hsmManagementRESTClient) ExecuteSingleTenantH
 
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.ExecuteSingleTenantHsmInstanceProposalOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &ExecuteSingleTenantHsmInstanceProposalOperation{
 		lro:      lro,
 		pollPath: override,
@@ -1747,13 +1556,7 @@ func (c *hsmManagementRESTClient) GetSingleTenantHsmIn
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/GetSingleTenantHsmInstanceProposal")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/singleTenantHsmInstances/*/proposals/*}")
-	}
+
 	opts = append((*c.CallOptions).GetSingleTenantHsmInstanceProposal[0:len((*c.CallOptions).GetSingleTenantHsmInstanceProposal):len((*c.CallOptions).GetSingleTenantHsmInstanceProposal)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.SingleTenantHsmInstanceProposal{}
@@ -1893,13 +1696,7 @@ func (c *hsmManagementRESTClient) DeleteSingleTenantHs
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.HsmManagement/DeleteSingleTenantHsmInstanceProposal")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/singleTenantHsmInstances/*/proposals/*}")
-	}
+
 	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		if settings.Path != "" {
 			baseUrl.Path = settings.Path
@@ -1935,10 +1732,7 @@ func (c *hsmManagementRESTClient) GetLocation(ctx cont
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
@@ -2089,13 +1883,7 @@ func (c *hsmManagementRESTClient) GetIamPolicy(ctx con
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
@@ -2156,13 +1944,7 @@ func (c *hsmManagementRESTClient) SetIamPolicy(ctx con
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
@@ -2225,13 +2007,7 @@ func (c *hsmManagementRESTClient) TestIamPermissions(c
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
@@ -2282,10 +2058,7 @@ func (c *hsmManagementRESTClient) GetOperation(ctx con
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
