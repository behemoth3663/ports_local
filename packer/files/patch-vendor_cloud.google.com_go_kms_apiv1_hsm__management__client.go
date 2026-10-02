--- vendor/cloud.google.com/go/kms/apiv1/hsm_management_client.go.orig	2026-10-02 20:06:38 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/hsm_management_client.go
@@ -32,7 +32,6 @@ import (
 	lroauto "cloud.google.com/go/longrunning/autogen"
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -590,16 +589,7 @@ func NewHsmManagementClient(ctx context.Context, opts 
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
@@ -624,35 +614,7 @@ func NewHsmManagementClient(ctx context.Context, opts 
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
@@ -729,16 +691,7 @@ func NewHsmManagementRESTClient(ctx context.Context, o
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
@@ -753,35 +706,6 @@ func NewHsmManagementRESTClient(ctx context.Context, o
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
@@ -837,12 +761,7 @@ func (c *hsmManagementGRPCClient) ListSingleTenantHsmI
 
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
@@ -889,12 +808,7 @@ func (c *hsmManagementGRPCClient) GetSingleTenantHsmIn
 
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
@@ -913,12 +827,7 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 
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
@@ -939,12 +848,7 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 
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
@@ -965,12 +869,7 @@ func (c *hsmManagementGRPCClient) ApproveSingleTenantH
 
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
@@ -989,12 +888,7 @@ func (c *hsmManagementGRPCClient) ExecuteSingleTenantH
 
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
@@ -1015,12 +909,7 @@ func (c *hsmManagementGRPCClient) GetSingleTenantHsmIn
 
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
@@ -1039,12 +928,7 @@ func (c *hsmManagementGRPCClient) ListSingleTenantHsmI
 
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
@@ -1091,12 +975,7 @@ func (c *hsmManagementGRPCClient) DeleteSingleTenantHs
 
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
@@ -1111,9 +990,7 @@ func (c *hsmManagementGRPCClient) GetLocation(ctx cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-	}
+
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	var resp *locationpb.Location
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1132,9 +1009,7 @@ func (c *hsmManagementGRPCClient) ListLocations(ctx co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/ListLocations")
-	}
+
 	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
 	it := &LocationIterator{}
 	req = proto.CloneOf(req)
@@ -1181,12 +1056,7 @@ func (c *hsmManagementGRPCClient) GetIamPolicy(ctx con
 
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
@@ -1205,12 +1075,7 @@ func (c *hsmManagementGRPCClient) SetIamPolicy(ctx con
 
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
@@ -1229,12 +1094,7 @@ func (c *hsmManagementGRPCClient) TestIamPermissions(c
 
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
@@ -1253,9 +1113,7 @@ func (c *hsmManagementGRPCClient) GetOperation(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1377,13 +1235,7 @@ func (c *hsmManagementRESTClient) GetSingleTenantHsmIn
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
@@ -1448,13 +1300,7 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
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
@@ -1521,13 +1367,7 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
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
@@ -1593,13 +1433,7 @@ func (c *hsmManagementRESTClient) ApproveSingleTenantH
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
@@ -1662,13 +1496,7 @@ func (c *hsmManagementRESTClient) ExecuteSingleTenantH
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
@@ -1723,13 +1551,7 @@ func (c *hsmManagementRESTClient) GetSingleTenantHsmIn
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
@@ -1869,13 +1691,7 @@ func (c *hsmManagementRESTClient) DeleteSingleTenantHs
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
@@ -1911,10 +1727,7 @@ func (c *hsmManagementRESTClient) GetLocation(ctx cont
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
@@ -2066,13 +1879,7 @@ func (c *hsmManagementRESTClient) GetIamPolicy(ctx con
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
@@ -2133,13 +1940,7 @@ func (c *hsmManagementRESTClient) SetIamPolicy(ctx con
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
@@ -2202,13 +2003,7 @@ func (c *hsmManagementRESTClient) TestIamPermissions(c
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
@@ -2259,10 +2054,7 @@ func (c *hsmManagementRESTClient) GetOperation(ctx con
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
