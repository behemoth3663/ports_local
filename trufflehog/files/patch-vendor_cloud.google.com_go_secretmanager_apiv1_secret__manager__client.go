--- vendor/cloud.google.com/go/secretmanager/apiv1/secret_manager_client.go.orig	2026-11-25 07:43:23 UTC
+++ vendor/cloud.google.com/go/secretmanager/apiv1/secret_manager_client.go
@@ -29,7 +29,6 @@ import (
 	iampb "cloud.google.com/go/iam/apiv1/iampb"
 	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -460,16 +459,6 @@ func NewClient(ctx context.Context, opts ...option.Cli
 //	SecretVersion
 func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
 	clientOpts := defaultGRPCClientOptions()
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "secretmanager",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/secretmanager/apiv1",
-			"gcp.client.language": "go",
-			"url.domain":          "secretmanager.googleapis.com",
-		}))
-	}
 	if newClientHook != nil {
 		hookOpts, err := newClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -492,102 +481,7 @@ func NewClient(ctx context.Context, opts ...option.Cli
 		locationsClient: locationpb.NewLocationsClient(connPool),
 	}
 	c.setGoogleClientInfo()
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "secretmanager",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/secretmanager/apiv1",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "secretmanager.googleapis.com",
-			}),
-		)
 
-		client.CallOptions.ListSecrets = append(client.CallOptions.ListSecrets, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateSecret = append(client.CallOptions.CreateSecret, gax.WithClientMetrics(metrics))
-		client.CallOptions.AddSecretVersion = append(client.CallOptions.AddSecretVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetSecret = append(client.CallOptions.GetSecret, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateSecret = append(client.CallOptions.UpdateSecret, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteSecret = append(client.CallOptions.DeleteSecret, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListSecretVersions = append(client.CallOptions.ListSecretVersions, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetSecretVersion = append(client.CallOptions.GetSecretVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.AccessSecretVersion = append(client.CallOptions.AccessSecretVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.DisableSecretVersion = append(client.CallOptions.DisableSecretVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.EnableSecretVersion = append(client.CallOptions.EnableSecretVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.DestroySecretVersion = append(client.CallOptions.DestroySecretVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientMetrics(metrics))
-		client.CallOptions.EnableManagedRotation = append(client.CallOptions.EnableManagedRotation, gax.WithClientMetrics(metrics))
-		client.CallOptions.RotateSecret = append(client.CallOptions.RotateSecret, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientMetrics(metrics))
-	}
-	if gax.IsFeatureEnabled("TRACING") {
-		tracing := gax.NewClientTracing(
-			gax.WithTracingAttributes(map[string]string{
-				gax.ClientService:  "secretmanager",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/secretmanager/apiv1",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "secretmanager.googleapis.com",
-			}),
-		)
-
-		client.CallOptions.ListSecrets = append(client.CallOptions.ListSecrets, gax.WithClientTracing(tracing))
-		client.CallOptions.CreateSecret = append(client.CallOptions.CreateSecret, gax.WithClientTracing(tracing))
-		client.CallOptions.AddSecretVersion = append(client.CallOptions.AddSecretVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.GetSecret = append(client.CallOptions.GetSecret, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateSecret = append(client.CallOptions.UpdateSecret, gax.WithClientTracing(tracing))
-		client.CallOptions.DeleteSecret = append(client.CallOptions.DeleteSecret, gax.WithClientTracing(tracing))
-		client.CallOptions.ListSecretVersions = append(client.CallOptions.ListSecretVersions, gax.WithClientTracing(tracing))
-		client.CallOptions.GetSecretVersion = append(client.CallOptions.GetSecretVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.AccessSecretVersion = append(client.CallOptions.AccessSecretVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.DisableSecretVersion = append(client.CallOptions.DisableSecretVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.EnableSecretVersion = append(client.CallOptions.EnableSecretVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.DestroySecretVersion = append(client.CallOptions.DestroySecretVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientTracing(tracing))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientTracing(tracing))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientTracing(tracing))
-		client.CallOptions.EnableManagedRotation = append(client.CallOptions.EnableManagedRotation, gax.WithClientTracing(tracing))
-		client.CallOptions.RotateSecret = append(client.CallOptions.RotateSecret, gax.WithClientTracing(tracing))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientTracing(tracing))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientTracing(tracing))
-	}
-	if gax.IsFeatureEnabled("LOGGING") {
-		logging := gax.NewClientLogging(
-			gax.WithLoggerProvider(c.logger),
-			gax.WithLoggingAttributes(map[string]string{
-				gax.ClientService:  "secretmanager",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/secretmanager/apiv1",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "secretmanager.googleapis.com",
-			}),
-		)
-
-		client.CallOptions.ListSecrets = append(client.CallOptions.ListSecrets, gax.WithClientLogging(logging))
-		client.CallOptions.CreateSecret = append(client.CallOptions.CreateSecret, gax.WithClientLogging(logging))
-		client.CallOptions.AddSecretVersion = append(client.CallOptions.AddSecretVersion, gax.WithClientLogging(logging))
-		client.CallOptions.GetSecret = append(client.CallOptions.GetSecret, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateSecret = append(client.CallOptions.UpdateSecret, gax.WithClientLogging(logging))
-		client.CallOptions.DeleteSecret = append(client.CallOptions.DeleteSecret, gax.WithClientLogging(logging))
-		client.CallOptions.ListSecretVersions = append(client.CallOptions.ListSecretVersions, gax.WithClientLogging(logging))
-		client.CallOptions.GetSecretVersion = append(client.CallOptions.GetSecretVersion, gax.WithClientLogging(logging))
-		client.CallOptions.AccessSecretVersion = append(client.CallOptions.AccessSecretVersion, gax.WithClientLogging(logging))
-		client.CallOptions.DisableSecretVersion = append(client.CallOptions.DisableSecretVersion, gax.WithClientLogging(logging))
-		client.CallOptions.EnableSecretVersion = append(client.CallOptions.EnableSecretVersion, gax.WithClientLogging(logging))
-		client.CallOptions.DestroySecretVersion = append(client.CallOptions.DestroySecretVersion, gax.WithClientLogging(logging))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientLogging(logging))
-		client.CallOptions.EnableManagedRotation = append(client.CallOptions.EnableManagedRotation, gax.WithClientLogging(logging))
-		client.CallOptions.RotateSecret = append(client.CallOptions.RotateSecret, gax.WithClientLogging(logging))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientLogging(logging))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientLogging(logging))
-	}
-
 	client.internalClient = c
 
 	return &client, nil
@@ -647,16 +541,6 @@ func NewRESTClient(ctx context.Context, opts ...option
 //	SecretVersion
 func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
 	clientOpts := append(defaultRESTClientOptions(), opts...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "secretmanager",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/secretmanager/apiv1",
-			"gcp.client.language": "go",
-			"url.domain":          "secretmanager.googleapis.com",
-		}))
-	}
 	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
 	if err != nil {
 		return nil, err
@@ -671,102 +555,6 @@ func NewRESTClient(ctx context.Context, opts ...option
 	}
 	c.setGoogleClientInfo()
 
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "secretmanager",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/secretmanager/apiv1",
-				gax.RPCSystem:      "http",
-				gax.URLDomain:      "secretmanager.googleapis.com",
-			}),
-		)
-
-		callOpts.ListSecrets = append(callOpts.ListSecrets, gax.WithClientMetrics(metrics))
-		callOpts.CreateSecret = append(callOpts.CreateSecret, gax.WithClientMetrics(metrics))
-		callOpts.AddSecretVersion = append(callOpts.AddSecretVersion, gax.WithClientMetrics(metrics))
-		callOpts.GetSecret = append(callOpts.GetSecret, gax.WithClientMetrics(metrics))
-		callOpts.UpdateSecret = append(callOpts.UpdateSecret, gax.WithClientMetrics(metrics))
-		callOpts.DeleteSecret = append(callOpts.DeleteSecret, gax.WithClientMetrics(metrics))
-		callOpts.ListSecretVersions = append(callOpts.ListSecretVersions, gax.WithClientMetrics(metrics))
-		callOpts.GetSecretVersion = append(callOpts.GetSecretVersion, gax.WithClientMetrics(metrics))
-		callOpts.AccessSecretVersion = append(callOpts.AccessSecretVersion, gax.WithClientMetrics(metrics))
-		callOpts.DisableSecretVersion = append(callOpts.DisableSecretVersion, gax.WithClientMetrics(metrics))
-		callOpts.EnableSecretVersion = append(callOpts.EnableSecretVersion, gax.WithClientMetrics(metrics))
-		callOpts.DestroySecretVersion = append(callOpts.DestroySecretVersion, gax.WithClientMetrics(metrics))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientMetrics(metrics))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientMetrics(metrics))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientMetrics(metrics))
-		callOpts.EnableManagedRotation = append(callOpts.EnableManagedRotation, gax.WithClientMetrics(metrics))
-		callOpts.RotateSecret = append(callOpts.RotateSecret, gax.WithClientMetrics(metrics))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientMetrics(metrics))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientMetrics(metrics))
-	}
-	if gax.IsFeatureEnabled("TRACING") {
-		tracing := gax.NewClientTracing(
-			gax.WithTracingAttributes(map[string]string{
-				gax.ClientService:  "secretmanager",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/secretmanager/apiv1",
-				gax.RPCSystem:      "http",
-				gax.URLDomain:      "secretmanager.googleapis.com",
-			}),
-		)
-
-		callOpts.ListSecrets = append(callOpts.ListSecrets, gax.WithClientTracing(tracing))
-		callOpts.CreateSecret = append(callOpts.CreateSecret, gax.WithClientTracing(tracing))
-		callOpts.AddSecretVersion = append(callOpts.AddSecretVersion, gax.WithClientTracing(tracing))
-		callOpts.GetSecret = append(callOpts.GetSecret, gax.WithClientTracing(tracing))
-		callOpts.UpdateSecret = append(callOpts.UpdateSecret, gax.WithClientTracing(tracing))
-		callOpts.DeleteSecret = append(callOpts.DeleteSecret, gax.WithClientTracing(tracing))
-		callOpts.ListSecretVersions = append(callOpts.ListSecretVersions, gax.WithClientTracing(tracing))
-		callOpts.GetSecretVersion = append(callOpts.GetSecretVersion, gax.WithClientTracing(tracing))
-		callOpts.AccessSecretVersion = append(callOpts.AccessSecretVersion, gax.WithClientTracing(tracing))
-		callOpts.DisableSecretVersion = append(callOpts.DisableSecretVersion, gax.WithClientTracing(tracing))
-		callOpts.EnableSecretVersion = append(callOpts.EnableSecretVersion, gax.WithClientTracing(tracing))
-		callOpts.DestroySecretVersion = append(callOpts.DestroySecretVersion, gax.WithClientTracing(tracing))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientTracing(tracing))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientTracing(tracing))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientTracing(tracing))
-		callOpts.EnableManagedRotation = append(callOpts.EnableManagedRotation, gax.WithClientTracing(tracing))
-		callOpts.RotateSecret = append(callOpts.RotateSecret, gax.WithClientTracing(tracing))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientTracing(tracing))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientTracing(tracing))
-	}
-	if gax.IsFeatureEnabled("LOGGING") {
-		logging := gax.NewClientLogging(
-			gax.WithLoggerProvider(c.logger),
-			gax.WithLoggingAttributes(map[string]string{
-				gax.ClientService:  "secretmanager",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/secretmanager/apiv1",
-				gax.RPCSystem:      "http",
-				gax.URLDomain:      "secretmanager.googleapis.com",
-			}),
-		)
-
-		callOpts.ListSecrets = append(callOpts.ListSecrets, gax.WithClientLogging(logging))
-		callOpts.CreateSecret = append(callOpts.CreateSecret, gax.WithClientLogging(logging))
-		callOpts.AddSecretVersion = append(callOpts.AddSecretVersion, gax.WithClientLogging(logging))
-		callOpts.GetSecret = append(callOpts.GetSecret, gax.WithClientLogging(logging))
-		callOpts.UpdateSecret = append(callOpts.UpdateSecret, gax.WithClientLogging(logging))
-		callOpts.DeleteSecret = append(callOpts.DeleteSecret, gax.WithClientLogging(logging))
-		callOpts.ListSecretVersions = append(callOpts.ListSecretVersions, gax.WithClientLogging(logging))
-		callOpts.GetSecretVersion = append(callOpts.GetSecretVersion, gax.WithClientLogging(logging))
-		callOpts.AccessSecretVersion = append(callOpts.AccessSecretVersion, gax.WithClientLogging(logging))
-		callOpts.DisableSecretVersion = append(callOpts.DisableSecretVersion, gax.WithClientLogging(logging))
-		callOpts.EnableSecretVersion = append(callOpts.EnableSecretVersion, gax.WithClientLogging(logging))
-		callOpts.DestroySecretVersion = append(callOpts.DestroySecretVersion, gax.WithClientLogging(logging))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientLogging(logging))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientLogging(logging))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientLogging(logging))
-		callOpts.EnableManagedRotation = append(callOpts.EnableManagedRotation, gax.WithClientLogging(logging))
-		callOpts.RotateSecret = append(callOpts.RotateSecret, gax.WithClientLogging(logging))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientLogging(logging))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientLogging(logging))
-	}
-
 	return &Client{internalClient: c, CallOptions: callOpts}, nil
 }
 
@@ -812,9 +600,6 @@ func (c *gRPCClient) ListSecrets(ctx context.Context, 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/ListSecrets")
-	}
 	opts = append((*c.CallOptions).ListSecrets[0:len((*c.CallOptions).ListSecrets):len((*c.CallOptions).ListSecrets)], opts...)
 	it := &SecretIterator{}
 	req = proto.CloneOf(req)
@@ -861,9 +646,6 @@ func (c *gRPCClient) CreateSecret(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/CreateSecret")
-	}
 	opts = append((*c.CallOptions).CreateSecret[0:len((*c.CallOptions).CreateSecret):len((*c.CallOptions).CreateSecret)], opts...)
 	var resp *secretmanagerpb.Secret
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -882,9 +664,6 @@ func (c *gRPCClient) AddSecretVersion(ctx context.Cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/AddSecretVersion")
-	}
 	opts = append((*c.CallOptions).AddSecretVersion[0:len((*c.CallOptions).AddSecretVersion):len((*c.CallOptions).AddSecretVersion)], opts...)
 	var resp *secretmanagerpb.SecretVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -903,9 +682,6 @@ func (c *gRPCClient) GetSecret(ctx context.Context, re
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/GetSecret")
-	}
 	opts = append((*c.CallOptions).GetSecret[0:len((*c.CallOptions).GetSecret):len((*c.CallOptions).GetSecret)], opts...)
 	var resp *secretmanagerpb.Secret
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -924,9 +700,6 @@ func (c *gRPCClient) UpdateSecret(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/UpdateSecret")
-	}
 	opts = append((*c.CallOptions).UpdateSecret[0:len((*c.CallOptions).UpdateSecret):len((*c.CallOptions).UpdateSecret)], opts...)
 	var resp *secretmanagerpb.Secret
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -945,9 +718,6 @@ func (c *gRPCClient) DeleteSecret(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/DeleteSecret")
-	}
 	opts = append((*c.CallOptions).DeleteSecret[0:len((*c.CallOptions).DeleteSecret):len((*c.CallOptions).DeleteSecret)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -962,9 +732,6 @@ func (c *gRPCClient) ListSecretVersions(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/ListSecretVersions")
-	}
 	opts = append((*c.CallOptions).ListSecretVersions[0:len((*c.CallOptions).ListSecretVersions):len((*c.CallOptions).ListSecretVersions)], opts...)
 	it := &SecretVersionIterator{}
 	req = proto.CloneOf(req)
@@ -1011,9 +778,6 @@ func (c *gRPCClient) GetSecretVersion(ctx context.Cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/GetSecretVersion")
-	}
 	opts = append((*c.CallOptions).GetSecretVersion[0:len((*c.CallOptions).GetSecretVersion):len((*c.CallOptions).GetSecretVersion)], opts...)
 	var resp *secretmanagerpb.SecretVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1032,9 +796,6 @@ func (c *gRPCClient) AccessSecretVersion(ctx context.C
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/AccessSecretVersion")
-	}
 	opts = append((*c.CallOptions).AccessSecretVersion[0:len((*c.CallOptions).AccessSecretVersion):len((*c.CallOptions).AccessSecretVersion)], opts...)
 	var resp *secretmanagerpb.AccessSecretVersionResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1053,9 +814,6 @@ func (c *gRPCClient) DisableSecretVersion(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/DisableSecretVersion")
-	}
 	opts = append((*c.CallOptions).DisableSecretVersion[0:len((*c.CallOptions).DisableSecretVersion):len((*c.CallOptions).DisableSecretVersion)], opts...)
 	var resp *secretmanagerpb.SecretVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1074,9 +832,6 @@ func (c *gRPCClient) EnableSecretVersion(ctx context.C
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/EnableSecretVersion")
-	}
 	opts = append((*c.CallOptions).EnableSecretVersion[0:len((*c.CallOptions).EnableSecretVersion):len((*c.CallOptions).EnableSecretVersion)], opts...)
 	var resp *secretmanagerpb.SecretVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1095,9 +850,6 @@ func (c *gRPCClient) DestroySecretVersion(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/DestroySecretVersion")
-	}
 	opts = append((*c.CallOptions).DestroySecretVersion[0:len((*c.CallOptions).DestroySecretVersion):len((*c.CallOptions).DestroySecretVersion)], opts...)
 	var resp *secretmanagerpb.SecretVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1116,9 +868,6 @@ func (c *gRPCClient) SetIamPolicy(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/SetIamPolicy")
-	}
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1137,9 +886,6 @@ func (c *gRPCClient) GetIamPolicy(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/GetIamPolicy")
-	}
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1158,9 +904,6 @@ func (c *gRPCClient) TestIamPermissions(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/TestIamPermissions")
-	}
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	var resp *iampb.TestIamPermissionsResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1179,9 +922,6 @@ func (c *gRPCClient) EnableManagedRotation(ctx context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/EnableManagedRotation")
-	}
 	opts = append((*c.CallOptions).EnableManagedRotation[0:len((*c.CallOptions).EnableManagedRotation):len((*c.CallOptions).EnableManagedRotation)], opts...)
 	var resp *secretmanagerpb.SecretVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1200,9 +940,6 @@ func (c *gRPCClient) RotateSecret(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/RotateSecret")
-	}
 	opts = append((*c.CallOptions).RotateSecret[0:len((*c.CallOptions).RotateSecret):len((*c.CallOptions).RotateSecret)], opts...)
 	var resp *secretmanagerpb.SecretVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1221,9 +958,6 @@ func (c *gRPCClient) GetLocation(ctx context.Context, 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-	}
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	var resp *locationpb.Location
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1242,9 +976,6 @@ func (c *gRPCClient) ListLocations(ctx context.Context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/ListLocations")
-	}
 	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
 	it := &LocationIterator{}
 	req = proto.CloneOf(req)
@@ -1405,10 +1136,6 @@ func (c *restClient) CreateSecret(ctx context.Context,
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/CreateSecret")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*}/secrets")
-	}
 	opts = append((*c.CallOptions).CreateSecret[0:len((*c.CallOptions).CreateSecret):len((*c.CallOptions).CreateSecret)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.Secret{}
@@ -1472,10 +1199,6 @@ func (c *restClient) AddSecretVersion(ctx context.Cont
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/AddSecretVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/secrets/*}:addVersion")
-	}
 	opts = append((*c.CallOptions).AddSecretVersion[0:len((*c.CallOptions).AddSecretVersion):len((*c.CallOptions).AddSecretVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.SecretVersion{}
@@ -1531,10 +1254,6 @@ func (c *restClient) GetSecret(ctx context.Context, re
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/GetSecret")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/secrets/*}")
-	}
 	opts = append((*c.CallOptions).GetSecret[0:len((*c.CallOptions).GetSecret):len((*c.CallOptions).GetSecret)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.Secret{}
@@ -1605,10 +1324,6 @@ func (c *restClient) UpdateSecret(ctx context.Context,
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/UpdateSecret")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{secret.name=projects/*/secrets/*}")
-	}
 	opts = append((*c.CallOptions).UpdateSecret[0:len((*c.CallOptions).UpdateSecret):len((*c.CallOptions).UpdateSecret)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.Secret{}
@@ -1667,10 +1382,6 @@ func (c *restClient) DeleteSecret(ctx context.Context,
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/DeleteSecret")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/secrets/*}")
-	}
 	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		if settings.Path != "" {
 			baseUrl.Path = settings.Path
@@ -1802,10 +1513,6 @@ func (c *restClient) GetSecretVersion(ctx context.Cont
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/GetSecretVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/secrets/*/versions/*}")
-	}
 	opts = append((*c.CallOptions).GetSecretVersion[0:len((*c.CallOptions).GetSecretVersion):len((*c.CallOptions).GetSecretVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.SecretVersion{}
@@ -1865,10 +1572,6 @@ func (c *restClient) AccessSecretVersion(ctx context.C
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/AccessSecretVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/secrets/*/versions/*}:access")
-	}
 	opts = append((*c.CallOptions).AccessSecretVersion[0:len((*c.CallOptions).AccessSecretVersion):len((*c.CallOptions).AccessSecretVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.AccessSecretVersionResponse{}
@@ -1934,10 +1637,6 @@ func (c *restClient) DisableSecretVersion(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/DisableSecretVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/secrets/*/versions/*}:disable")
-	}
 	opts = append((*c.CallOptions).DisableSecretVersion[0:len((*c.CallOptions).DisableSecretVersion):len((*c.CallOptions).DisableSecretVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.SecretVersion{}
@@ -2003,10 +1702,6 @@ func (c *restClient) EnableSecretVersion(ctx context.C
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/EnableSecretVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/secrets/*/versions/*}:enable")
-	}
 	opts = append((*c.CallOptions).EnableSecretVersion[0:len((*c.CallOptions).EnableSecretVersion):len((*c.CallOptions).EnableSecretVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.SecretVersion{}
@@ -2073,10 +1768,6 @@ func (c *restClient) DestroySecretVersion(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/DestroySecretVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/secrets/*/versions/*}:destroy")
-	}
 	opts = append((*c.CallOptions).DestroySecretVersion[0:len((*c.CallOptions).DestroySecretVersion):len((*c.CallOptions).DestroySecretVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.SecretVersion{}
@@ -2144,10 +1835,6 @@ func (c *restClient) SetIamPolicy(ctx context.Context,
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/SetIamPolicy")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/secrets/*}:setIamPolicy")
-	}
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.Policy{}
@@ -2207,10 +1894,6 @@ func (c *restClient) GetIamPolicy(ctx context.Context,
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/GetIamPolicy")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/secrets/*}:getIamPolicy")
-	}
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.Policy{}
@@ -2278,10 +1961,6 @@ func (c *restClient) TestIamPermissions(ctx context.Co
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/TestIamPermissions")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{resource=projects/*/secrets/*}:testIamPermissions")
-	}
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &iampb.TestIamPermissionsResponse{}
@@ -2347,10 +2026,6 @@ func (c *restClient) EnableManagedRotation(ctx context
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/EnableManagedRotation")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/secrets/*}:enableManagedRotation")
-	}
 	opts = append((*c.CallOptions).EnableManagedRotation[0:len((*c.CallOptions).EnableManagedRotation):len((*c.CallOptions).EnableManagedRotation)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.SecretVersion{}
@@ -2414,10 +2089,6 @@ func (c *restClient) RotateSecret(ctx context.Context,
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.secretmanager.v1.SecretManagerService/RotateSecret")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/secrets/*}:rotateSecret")
-	}
 	opts = append((*c.CallOptions).RotateSecret[0:len((*c.CallOptions).RotateSecret):len((*c.CallOptions).RotateSecret)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &secretmanagerpb.SecretVersion{}
@@ -2473,10 +2144,6 @@ func (c *restClient) GetLocation(ctx context.Context, 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*}")
-	}
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &locationpb.Location{}
