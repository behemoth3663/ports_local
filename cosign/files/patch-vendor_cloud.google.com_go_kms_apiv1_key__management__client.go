--- vendor/cloud.google.com/go/kms/apiv1/key_management_client.go.orig	2026-10-02 15:35:42 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/key_management_client.go
@@ -32,7 +32,6 @@ import (
 	lroauto "cloud.google.com/go/longrunning/autogen"
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -1357,16 +1356,7 @@ func NewKeyManagementClient(ctx context.Context, opts 
 // Using gRPC with Cloud KMS (at https://cloud.google.com/kms/docs/grpc).
 func NewKeyManagementClient(ctx context.Context, opts ...option.ClientOption) (*KeyManagementClient, error) {
 	clientOpts := defaultKeyManagementGRPCClientOptions()
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
 	if newKeyManagementClientHook != nil {
 		hookOpts, err := newKeyManagementClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -1391,59 +1381,7 @@ func NewKeyManagementClient(ctx context.Context, opts 
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
 
-		client.CallOptions.ListKeyRings = append(client.CallOptions.ListKeyRings, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListCryptoKeys = append(client.CallOptions.ListCryptoKeys, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListCryptoKeyVersions = append(client.CallOptions.ListCryptoKeyVersions, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListImportJobs = append(client.CallOptions.ListImportJobs, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListRetiredResources = append(client.CallOptions.ListRetiredResources, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetKeyRing = append(client.CallOptions.GetKeyRing, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetCryptoKey = append(client.CallOptions.GetCryptoKey, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetCryptoKeyVersion = append(client.CallOptions.GetCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetPublicKey = append(client.CallOptions.GetPublicKey, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetImportJob = append(client.CallOptions.GetImportJob, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetRetiredResource = append(client.CallOptions.GetRetiredResource, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateKeyRing = append(client.CallOptions.CreateKeyRing, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateCryptoKey = append(client.CallOptions.CreateCryptoKey, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateCryptoKeyVersion = append(client.CallOptions.CreateCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteCryptoKey = append(client.CallOptions.DeleteCryptoKey, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteCryptoKeyVersion = append(client.CallOptions.DeleteCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.ImportCryptoKeyVersion = append(client.CallOptions.ImportCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateImportJob = append(client.CallOptions.CreateImportJob, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateCryptoKey = append(client.CallOptions.UpdateCryptoKey, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateCryptoKeyVersion = append(client.CallOptions.UpdateCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateCryptoKeyPrimaryVersion = append(client.CallOptions.UpdateCryptoKeyPrimaryVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.DestroyCryptoKeyVersion = append(client.CallOptions.DestroyCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.RestoreCryptoKeyVersion = append(client.CallOptions.RestoreCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.Encrypt = append(client.CallOptions.Encrypt, gax.WithClientMetrics(metrics))
-		client.CallOptions.Decrypt = append(client.CallOptions.Decrypt, gax.WithClientMetrics(metrics))
-		client.CallOptions.RawEncrypt = append(client.CallOptions.RawEncrypt, gax.WithClientMetrics(metrics))
-		client.CallOptions.RawDecrypt = append(client.CallOptions.RawDecrypt, gax.WithClientMetrics(metrics))
-		client.CallOptions.AsymmetricSign = append(client.CallOptions.AsymmetricSign, gax.WithClientMetrics(metrics))
-		client.CallOptions.AsymmetricDecrypt = append(client.CallOptions.AsymmetricDecrypt, gax.WithClientMetrics(metrics))
-		client.CallOptions.MacSign = append(client.CallOptions.MacSign, gax.WithClientMetrics(metrics))
-		client.CallOptions.MacVerify = append(client.CallOptions.MacVerify, gax.WithClientMetrics(metrics))
-		client.CallOptions.Decapsulate = append(client.CallOptions.Decapsulate, gax.WithClientMetrics(metrics))
-		client.CallOptions.GenerateRandomBytes = append(client.CallOptions.GenerateRandomBytes, gax.WithClientMetrics(metrics))
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
@@ -1526,16 +1464,7 @@ func NewKeyManagementRESTClient(ctx context.Context, o
 // Using gRPC with Cloud KMS (at https://cloud.google.com/kms/docs/grpc).
 func NewKeyManagementRESTClient(ctx context.Context, opts ...option.ClientOption) (*KeyManagementClient, error) {
 	clientOpts := append(defaultKeyManagementRESTClientOptions(), opts...)
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
@@ -1550,59 +1479,6 @@ func NewKeyManagementRESTClient(ctx context.Context, o
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
-		callOpts.ListKeyRings = append(callOpts.ListKeyRings, gax.WithClientMetrics(metrics))
-		callOpts.ListCryptoKeys = append(callOpts.ListCryptoKeys, gax.WithClientMetrics(metrics))
-		callOpts.ListCryptoKeyVersions = append(callOpts.ListCryptoKeyVersions, gax.WithClientMetrics(metrics))
-		callOpts.ListImportJobs = append(callOpts.ListImportJobs, gax.WithClientMetrics(metrics))
-		callOpts.ListRetiredResources = append(callOpts.ListRetiredResources, gax.WithClientMetrics(metrics))
-		callOpts.GetKeyRing = append(callOpts.GetKeyRing, gax.WithClientMetrics(metrics))
-		callOpts.GetCryptoKey = append(callOpts.GetCryptoKey, gax.WithClientMetrics(metrics))
-		callOpts.GetCryptoKeyVersion = append(callOpts.GetCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.GetPublicKey = append(callOpts.GetPublicKey, gax.WithClientMetrics(metrics))
-		callOpts.GetImportJob = append(callOpts.GetImportJob, gax.WithClientMetrics(metrics))
-		callOpts.GetRetiredResource = append(callOpts.GetRetiredResource, gax.WithClientMetrics(metrics))
-		callOpts.CreateKeyRing = append(callOpts.CreateKeyRing, gax.WithClientMetrics(metrics))
-		callOpts.CreateCryptoKey = append(callOpts.CreateCryptoKey, gax.WithClientMetrics(metrics))
-		callOpts.CreateCryptoKeyVersion = append(callOpts.CreateCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.DeleteCryptoKey = append(callOpts.DeleteCryptoKey, gax.WithClientMetrics(metrics))
-		callOpts.DeleteCryptoKeyVersion = append(callOpts.DeleteCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.ImportCryptoKeyVersion = append(callOpts.ImportCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.CreateImportJob = append(callOpts.CreateImportJob, gax.WithClientMetrics(metrics))
-		callOpts.UpdateCryptoKey = append(callOpts.UpdateCryptoKey, gax.WithClientMetrics(metrics))
-		callOpts.UpdateCryptoKeyVersion = append(callOpts.UpdateCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.UpdateCryptoKeyPrimaryVersion = append(callOpts.UpdateCryptoKeyPrimaryVersion, gax.WithClientMetrics(metrics))
-		callOpts.DestroyCryptoKeyVersion = append(callOpts.DestroyCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.RestoreCryptoKeyVersion = append(callOpts.RestoreCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.Encrypt = append(callOpts.Encrypt, gax.WithClientMetrics(metrics))
-		callOpts.Decrypt = append(callOpts.Decrypt, gax.WithClientMetrics(metrics))
-		callOpts.RawEncrypt = append(callOpts.RawEncrypt, gax.WithClientMetrics(metrics))
-		callOpts.RawDecrypt = append(callOpts.RawDecrypt, gax.WithClientMetrics(metrics))
-		callOpts.AsymmetricSign = append(callOpts.AsymmetricSign, gax.WithClientMetrics(metrics))
-		callOpts.AsymmetricDecrypt = append(callOpts.AsymmetricDecrypt, gax.WithClientMetrics(metrics))
-		callOpts.MacSign = append(callOpts.MacSign, gax.WithClientMetrics(metrics))
-		callOpts.MacVerify = append(callOpts.MacVerify, gax.WithClientMetrics(metrics))
-		callOpts.Decapsulate = append(callOpts.Decapsulate, gax.WithClientMetrics(metrics))
-		callOpts.GenerateRandomBytes = append(callOpts.GenerateRandomBytes, gax.WithClientMetrics(metrics))
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
@@ -1658,12 +1534,7 @@ func (c *keyManagementGRPCClient) ListKeyRings(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListKeyRings")
-	}
+
 	opts = append((*c.CallOptions).ListKeyRings[0:len((*c.CallOptions).ListKeyRings):len((*c.CallOptions).ListKeyRings)], opts...)
 	it := &KeyRingIterator{}
 	req = proto.CloneOf(req)
@@ -1710,12 +1581,7 @@ func (c *keyManagementGRPCClient) ListCryptoKeys(ctx c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListCryptoKeys")
-	}
+
 	opts = append((*c.CallOptions).ListCryptoKeys[0:len((*c.CallOptions).ListCryptoKeys):len((*c.CallOptions).ListCryptoKeys)], opts...)
 	it := &CryptoKeyIterator{}
 	req = proto.CloneOf(req)
@@ -1762,12 +1628,7 @@ func (c *keyManagementGRPCClient) ListCryptoKeyVersion
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListCryptoKeyVersions")
-	}
+
 	opts = append((*c.CallOptions).ListCryptoKeyVersions[0:len((*c.CallOptions).ListCryptoKeyVersions):len((*c.CallOptions).ListCryptoKeyVersions)], opts...)
 	it := &CryptoKeyVersionIterator{}
 	req = proto.CloneOf(req)
@@ -1814,12 +1675,7 @@ func (c *keyManagementGRPCClient) ListImportJobs(ctx c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListImportJobs")
-	}
+
 	opts = append((*c.CallOptions).ListImportJobs[0:len((*c.CallOptions).ListImportJobs):len((*c.CallOptions).ListImportJobs)], opts...)
 	it := &ImportJobIterator{}
 	req = proto.CloneOf(req)
@@ -1866,12 +1722,7 @@ func (c *keyManagementGRPCClient) ListRetiredResources
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListRetiredResources")
-	}
+
 	opts = append((*c.CallOptions).ListRetiredResources[0:len((*c.CallOptions).ListRetiredResources):len((*c.CallOptions).ListRetiredResources)], opts...)
 	it := &RetiredResourceIterator{}
 	req = proto.CloneOf(req)
@@ -1918,12 +1769,7 @@ func (c *keyManagementGRPCClient) GetKeyRing(ctx conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetKeyRing")
-	}
+
 	opts = append((*c.CallOptions).GetKeyRing[0:len((*c.CallOptions).GetKeyRing):len((*c.CallOptions).GetKeyRing)], opts...)
 	var resp *kmspb.KeyRing
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1942,12 +1788,7 @@ func (c *keyManagementGRPCClient) GetCryptoKey(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKey")
-	}
+
 	opts = append((*c.CallOptions).GetCryptoKey[0:len((*c.CallOptions).GetCryptoKey):len((*c.CallOptions).GetCryptoKey)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1966,12 +1807,7 @@ func (c *keyManagementGRPCClient) GetCryptoKeyVersion(
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).GetCryptoKeyVersion[0:len((*c.CallOptions).GetCryptoKeyVersion):len((*c.CallOptions).GetCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1990,12 +1826,7 @@ func (c *keyManagementGRPCClient) GetPublicKey(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetPublicKey")
-	}
+
 	opts = append((*c.CallOptions).GetPublicKey[0:len((*c.CallOptions).GetPublicKey):len((*c.CallOptions).GetPublicKey)], opts...)
 	var resp *kmspb.PublicKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2014,12 +1845,7 @@ func (c *keyManagementGRPCClient) GetImportJob(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetImportJob")
-	}
+
 	opts = append((*c.CallOptions).GetImportJob[0:len((*c.CallOptions).GetImportJob):len((*c.CallOptions).GetImportJob)], opts...)
 	var resp *kmspb.ImportJob
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2038,12 +1864,7 @@ func (c *keyManagementGRPCClient) GetRetiredResource(c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetRetiredResource")
-	}
+
 	opts = append((*c.CallOptions).GetRetiredResource[0:len((*c.CallOptions).GetRetiredResource):len((*c.CallOptions).GetRetiredResource)], opts...)
 	var resp *kmspb.RetiredResource
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2062,12 +1883,7 @@ func (c *keyManagementGRPCClient) CreateKeyRing(ctx co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateKeyRing")
-	}
+
 	opts = append((*c.CallOptions).CreateKeyRing[0:len((*c.CallOptions).CreateKeyRing):len((*c.CallOptions).CreateKeyRing)], opts...)
 	var resp *kmspb.KeyRing
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2086,12 +1902,7 @@ func (c *keyManagementGRPCClient) CreateCryptoKey(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKey")
-	}
+
 	opts = append((*c.CallOptions).CreateCryptoKey[0:len((*c.CallOptions).CreateCryptoKey):len((*c.CallOptions).CreateCryptoKey)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2110,12 +1921,7 @@ func (c *keyManagementGRPCClient) CreateCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).CreateCryptoKeyVersion[0:len((*c.CallOptions).CreateCryptoKeyVersion):len((*c.CallOptions).CreateCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2134,12 +1940,7 @@ func (c *keyManagementGRPCClient) DeleteCryptoKey(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKey")
-	}
+
 	opts = append((*c.CallOptions).DeleteCryptoKey[0:len((*c.CallOptions).DeleteCryptoKey):len((*c.CallOptions).DeleteCryptoKey)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2160,12 +1961,7 @@ func (c *keyManagementGRPCClient) DeleteCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).DeleteCryptoKeyVersion[0:len((*c.CallOptions).DeleteCryptoKeyVersion):len((*c.CallOptions).DeleteCryptoKeyVersion)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2186,12 +1982,7 @@ func (c *keyManagementGRPCClient) ImportCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ImportCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).ImportCryptoKeyVersion[0:len((*c.CallOptions).ImportCryptoKeyVersion):len((*c.CallOptions).ImportCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2210,12 +2001,7 @@ func (c *keyManagementGRPCClient) CreateImportJob(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateImportJob")
-	}
+
 	opts = append((*c.CallOptions).CreateImportJob[0:len((*c.CallOptions).CreateImportJob):len((*c.CallOptions).CreateImportJob)], opts...)
 	var resp *kmspb.ImportJob
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2234,9 +2020,7 @@ func (c *keyManagementGRPCClient) UpdateCryptoKey(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKey")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKey[0:len((*c.CallOptions).UpdateCryptoKey):len((*c.CallOptions).UpdateCryptoKey)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2255,9 +2039,7 @@ func (c *keyManagementGRPCClient) UpdateCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKeyVersion[0:len((*c.CallOptions).UpdateCryptoKeyVersion):len((*c.CallOptions).UpdateCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2276,12 +2058,7 @@ func (c *keyManagementGRPCClient) UpdateCryptoKeyPrima
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKeyPrimaryVersion")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKeyPrimaryVersion[0:len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion):len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2300,12 +2077,7 @@ func (c *keyManagementGRPCClient) DestroyCryptoKeyVers
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DestroyCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).DestroyCryptoKeyVersion[0:len((*c.CallOptions).DestroyCryptoKeyVersion):len((*c.CallOptions).DestroyCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2324,12 +2096,7 @@ func (c *keyManagementGRPCClient) RestoreCryptoKeyVers
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RestoreCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).RestoreCryptoKeyVersion[0:len((*c.CallOptions).RestoreCryptoKeyVersion):len((*c.CallOptions).RestoreCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2348,12 +2115,7 @@ func (c *keyManagementGRPCClient) Encrypt(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Encrypt")
-	}
+
 	opts = append((*c.CallOptions).Encrypt[0:len((*c.CallOptions).Encrypt):len((*c.CallOptions).Encrypt)], opts...)
 	var resp *kmspb.EncryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2372,12 +2134,7 @@ func (c *keyManagementGRPCClient) Decrypt(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decrypt")
-	}
+
 	opts = append((*c.CallOptions).Decrypt[0:len((*c.CallOptions).Decrypt):len((*c.CallOptions).Decrypt)], opts...)
 	var resp *kmspb.DecryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2396,9 +2153,7 @@ func (c *keyManagementGRPCClient) RawEncrypt(ctx conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RawEncrypt")
-	}
+
 	opts = append((*c.CallOptions).RawEncrypt[0:len((*c.CallOptions).RawEncrypt):len((*c.CallOptions).RawEncrypt)], opts...)
 	var resp *kmspb.RawEncryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2417,9 +2172,7 @@ func (c *keyManagementGRPCClient) RawDecrypt(ctx conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RawDecrypt")
-	}
+
 	opts = append((*c.CallOptions).RawDecrypt[0:len((*c.CallOptions).RawDecrypt):len((*c.CallOptions).RawDecrypt)], opts...)
 	var resp *kmspb.RawDecryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2438,12 +2191,7 @@ func (c *keyManagementGRPCClient) AsymmetricSign(ctx c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricSign")
-	}
+
 	opts = append((*c.CallOptions).AsymmetricSign[0:len((*c.CallOptions).AsymmetricSign):len((*c.CallOptions).AsymmetricSign)], opts...)
 	var resp *kmspb.AsymmetricSignResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2462,12 +2210,7 @@ func (c *keyManagementGRPCClient) AsymmetricDecrypt(ct
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricDecrypt")
-	}
+
 	opts = append((*c.CallOptions).AsymmetricDecrypt[0:len((*c.CallOptions).AsymmetricDecrypt):len((*c.CallOptions).AsymmetricDecrypt)], opts...)
 	var resp *kmspb.AsymmetricDecryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2486,12 +2229,7 @@ func (c *keyManagementGRPCClient) MacSign(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacSign")
-	}
+
 	opts = append((*c.CallOptions).MacSign[0:len((*c.CallOptions).MacSign):len((*c.CallOptions).MacSign)], opts...)
 	var resp *kmspb.MacSignResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2510,12 +2248,7 @@ func (c *keyManagementGRPCClient) MacVerify(ctx contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacVerify")
-	}
+
 	opts = append((*c.CallOptions).MacVerify[0:len((*c.CallOptions).MacVerify):len((*c.CallOptions).MacVerify)], opts...)
 	var resp *kmspb.MacVerifyResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2534,12 +2267,7 @@ func (c *keyManagementGRPCClient) Decapsulate(ctx cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decapsulate")
-	}
+
 	opts = append((*c.CallOptions).Decapsulate[0:len((*c.CallOptions).Decapsulate):len((*c.CallOptions).Decapsulate)], opts...)
 	var resp *kmspb.DecapsulateResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2558,9 +2286,7 @@ func (c *keyManagementGRPCClient) GenerateRandomBytes(
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GenerateRandomBytes")
-	}
+
 	opts = append((*c.CallOptions).GenerateRandomBytes[0:len((*c.CallOptions).GenerateRandomBytes):len((*c.CallOptions).GenerateRandomBytes)], opts...)
 	var resp *kmspb.GenerateRandomBytesResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2579,9 +2305,7 @@ func (c *keyManagementGRPCClient) GetLocation(ctx cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-	}
+
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	var resp *locationpb.Location
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2600,9 +2324,7 @@ func (c *keyManagementGRPCClient) ListLocations(ctx co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/ListLocations")
-	}
+
 	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
 	it := &LocationIterator{}
 	req = proto.CloneOf(req)
@@ -2649,12 +2371,7 @@ func (c *keyManagementGRPCClient) GetIamPolicy(ctx con
 
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
@@ -2673,12 +2390,7 @@ func (c *keyManagementGRPCClient) SetIamPolicy(ctx con
 
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
@@ -2697,12 +2409,7 @@ func (c *keyManagementGRPCClient) TestIamPermissions(c
 
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
@@ -2721,9 +2428,7 @@ func (c *keyManagementGRPCClient) GetOperation(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -3178,13 +2883,7 @@ func (c *keyManagementRESTClient) GetKeyRing(ctx conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetKeyRing")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*}")
-	}
+
 	opts = append((*c.CallOptions).GetKeyRing[0:len((*c.CallOptions).GetKeyRing):len((*c.CallOptions).GetKeyRing)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.KeyRing{}
@@ -3237,13 +2936,7 @@ func (c *keyManagementRESTClient) GetCryptoKey(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}")
-	}
+
 	opts = append((*c.CallOptions).GetCryptoKey[0:len((*c.CallOptions).GetCryptoKey):len((*c.CallOptions).GetCryptoKey)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKey{}
@@ -3295,13 +2988,7 @@ func (c *keyManagementRESTClient) GetCryptoKeyVersion(
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}")
-	}
+
 	opts = append((*c.CallOptions).GetCryptoKeyVersion[0:len((*c.CallOptions).GetCryptoKeyVersion):len((*c.CallOptions).GetCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -3360,13 +3047,7 @@ func (c *keyManagementRESTClient) GetPublicKey(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetPublicKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}/publicKey")
-	}
+
 	opts = append((*c.CallOptions).GetPublicKey[0:len((*c.CallOptions).GetPublicKey):len((*c.CallOptions).GetPublicKey)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.PublicKey{}
@@ -3417,13 +3098,7 @@ func (c *keyManagementRESTClient) GetImportJob(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetImportJob")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/importJobs/*}")
-	}
+
 	opts = append((*c.CallOptions).GetImportJob[0:len((*c.CallOptions).GetImportJob):len((*c.CallOptions).GetImportJob)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.ImportJob{}
@@ -3476,13 +3151,7 @@ func (c *keyManagementRESTClient) GetRetiredResource(c
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetRetiredResource")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/retiredResources/*}")
-	}
+
 	opts = append((*c.CallOptions).GetRetiredResource[0:len((*c.CallOptions).GetRetiredResource):len((*c.CallOptions).GetRetiredResource)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.RetiredResource{}
@@ -3542,13 +3211,7 @@ func (c *keyManagementRESTClient) CreateKeyRing(ctx co
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateKeyRing")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*}/keyRings")
-	}
+
 	opts = append((*c.CallOptions).CreateKeyRing[0:len((*c.CallOptions).CreateKeyRing):len((*c.CallOptions).CreateKeyRing)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.KeyRing{}
@@ -3615,13 +3278,7 @@ func (c *keyManagementRESTClient) CreateCryptoKey(ctx 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*}/cryptoKeys")
-	}
+
 	opts = append((*c.CallOptions).CreateCryptoKey[0:len((*c.CallOptions).CreateCryptoKey):len((*c.CallOptions).CreateCryptoKey)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKey{}
@@ -3684,13 +3341,7 @@ func (c *keyManagementRESTClient) CreateCryptoKeyVersi
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*/cryptoKeys/*}/cryptoKeyVersions")
-	}
+
 	opts = append((*c.CallOptions).CreateCryptoKeyVersion[0:len((*c.CallOptions).CreateCryptoKeyVersion):len((*c.CallOptions).CreateCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -3746,13 +3397,7 @@ func (c *keyManagementRESTClient) DeleteCryptoKey(ctx 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}")
-	}
+
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
 	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -3816,13 +3461,7 @@ func (c *keyManagementRESTClient) DeleteCryptoKeyVersi
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}")
-	}
+
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
 	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -3889,13 +3528,7 @@ func (c *keyManagementRESTClient) ImportCryptoKeyVersi
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ImportCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*/cryptoKeys/*}/cryptoKeyVersions:import")
-	}
+
 	opts = append((*c.CallOptions).ImportCryptoKeyVersion[0:len((*c.CallOptions).ImportCryptoKeyVersion):len((*c.CallOptions).ImportCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -3958,13 +3591,7 @@ func (c *keyManagementRESTClient) CreateImportJob(ctx 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetParent()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateImportJob")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*}/importJobs")
-	}
+
 	opts = append((*c.CallOptions).CreateImportJob[0:len((*c.CallOptions).CreateImportJob):len((*c.CallOptions).CreateImportJob)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.ImportJob{}
@@ -4029,10 +3656,7 @@ func (c *keyManagementRESTClient) UpdateCryptoKey(ctx 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{crypto_key.name=projects/*/locations/*/keyRings/*/cryptoKeys/*}")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKey[0:len((*c.CallOptions).UpdateCryptoKey):len((*c.CallOptions).UpdateCryptoKey)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKey{}
@@ -4108,10 +3732,7 @@ func (c *keyManagementRESTClient) UpdateCryptoKeyVersi
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{crypto_key_version.name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKeyVersion[0:len((*c.CallOptions).UpdateCryptoKeyVersion):len((*c.CallOptions).UpdateCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -4173,13 +3794,7 @@ func (c *keyManagementRESTClient) UpdateCryptoKeyPrima
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKeyPrimaryVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}:updatePrimaryVersion")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKeyPrimaryVersion[0:len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion):len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKey{}
@@ -4256,13 +3871,7 @@ func (c *keyManagementRESTClient) DestroyCryptoKeyVers
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DestroyCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:destroy")
-	}
+
 	opts = append((*c.CallOptions).DestroyCryptoKeyVersion[0:len((*c.CallOptions).DestroyCryptoKeyVersion):len((*c.CallOptions).DestroyCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -4327,13 +3936,7 @@ func (c *keyManagementRESTClient) RestoreCryptoKeyVers
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RestoreCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:restore")
-	}
+
 	opts = append((*c.CallOptions).RestoreCryptoKeyVersion[0:len((*c.CallOptions).RestoreCryptoKeyVersion):len((*c.CallOptions).RestoreCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -4393,13 +3996,7 @@ func (c *keyManagementRESTClient) Encrypt(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Encrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/**}:encrypt")
-	}
+
 	opts = append((*c.CallOptions).Encrypt[0:len((*c.CallOptions).Encrypt):len((*c.CallOptions).Encrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EncryptResponse{}
@@ -4459,13 +4056,7 @@ func (c *keyManagementRESTClient) Decrypt(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}:decrypt")
-	}
+
 	opts = append((*c.CallOptions).Decrypt[0:len((*c.CallOptions).Decrypt):len((*c.CallOptions).Decrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.DecryptResponse{}
@@ -4527,10 +4118,7 @@ func (c *keyManagementRESTClient) RawEncrypt(ctx conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RawEncrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:rawEncrypt")
-	}
+
 	opts = append((*c.CallOptions).RawEncrypt[0:len((*c.CallOptions).RawEncrypt):len((*c.CallOptions).RawEncrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.RawEncryptResponse{}
@@ -4590,10 +4178,7 @@ func (c *keyManagementRESTClient) RawDecrypt(ctx conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RawDecrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:rawDecrypt")
-	}
+
 	opts = append((*c.CallOptions).RawDecrypt[0:len((*c.CallOptions).RawDecrypt):len((*c.CallOptions).RawDecrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.RawDecryptResponse{}
@@ -4654,13 +4239,7 @@ func (c *keyManagementRESTClient) AsymmetricSign(ctx c
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricSign")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:asymmetricSign")
-	}
+
 	opts = append((*c.CallOptions).AsymmetricSign[0:len((*c.CallOptions).AsymmetricSign):len((*c.CallOptions).AsymmetricSign)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.AsymmetricSignResponse{}
@@ -4721,13 +4300,7 @@ func (c *keyManagementRESTClient) AsymmetricDecrypt(ct
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricDecrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:asymmetricDecrypt")
-	}
+
 	opts = append((*c.CallOptions).AsymmetricDecrypt[0:len((*c.CallOptions).AsymmetricDecrypt):len((*c.CallOptions).AsymmetricDecrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.AsymmetricDecryptResponse{}
@@ -4786,13 +4359,7 @@ func (c *keyManagementRESTClient) MacSign(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacSign")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:macSign")
-	}
+
 	opts = append((*c.CallOptions).MacSign[0:len((*c.CallOptions).MacSign):len((*c.CallOptions).MacSign)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.MacSignResponse{}
@@ -4852,13 +4419,7 @@ func (c *keyManagementRESTClient) MacVerify(ctx contex
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacVerify")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:macVerify")
-	}
+
 	opts = append((*c.CallOptions).MacVerify[0:len((*c.CallOptions).MacVerify):len((*c.CallOptions).MacVerify)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.MacVerifyResponse{}
@@ -4919,13 +4480,7 @@ func (c *keyManagementRESTClient) Decapsulate(ctx cont
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "resource_name", fmt.Sprintf("//cloudkms.googleapis.com/%v", req.GetName()))
-	}
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decapsulate")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:decapsulate")
-	}
+
 	opts = append((*c.CallOptions).Decapsulate[0:len((*c.CallOptions).Decapsulate):len((*c.CallOptions).Decapsulate)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.DecapsulateResponse{}
@@ -4983,10 +4538,7 @@ func (c *keyManagementRESTClient) GenerateRandomBytes(
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GenerateRandomBytes")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{location=projects/*/locations/*}:generateRandomBytes")
-	}
+
 	opts = append((*c.CallOptions).GenerateRandomBytes[0:len((*c.CallOptions).GenerateRandomBytes):len((*c.CallOptions).GenerateRandomBytes)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.GenerateRandomBytesResponse{}
@@ -5037,10 +4589,7 @@ func (c *keyManagementRESTClient) GetLocation(ctx cont
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
@@ -5192,13 +4741,7 @@ func (c *keyManagementRESTClient) GetIamPolicy(ctx con
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
@@ -5259,13 +4802,7 @@ func (c *keyManagementRESTClient) SetIamPolicy(ctx con
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
@@ -5328,13 +4865,7 @@ func (c *keyManagementRESTClient) TestIamPermissions(c
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
@@ -5385,10 +4916,7 @@ func (c *keyManagementRESTClient) GetOperation(ctx con
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
