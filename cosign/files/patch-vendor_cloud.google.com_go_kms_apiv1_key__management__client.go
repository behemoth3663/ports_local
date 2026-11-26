--- vendor/cloud.google.com/go/kms/apiv1/key_management_client.go.orig	2026-11-26 19:20:44 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/key_management_client.go
@@ -32,8 +32,6 @@ import (
 	lroauto "cloud.google.com/go/longrunning/autogen"
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
-	trace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -1396,16 +1394,7 @@ func NewKeyManagementClient(ctx context.Context, opts 
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
@@ -1430,168 +1419,7 @@ func NewKeyManagementClient(ctx context.Context, opts 
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
-		client.CallOptions.ImportTrustedKeyWrappedCryptoKeyVersion = append(client.CallOptions.ImportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		client.CallOptions.ExportTrustedKeyWrappedCryptoKeyVersion = append(client.CallOptions.ExportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientMetrics(metrics))
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
-		client.CallOptions.ListKeyRings = append(client.CallOptions.ListKeyRings, gax.WithClientTracing(tracing))
-		client.CallOptions.ListCryptoKeys = append(client.CallOptions.ListCryptoKeys, gax.WithClientTracing(tracing))
-		client.CallOptions.ListCryptoKeyVersions = append(client.CallOptions.ListCryptoKeyVersions, gax.WithClientTracing(tracing))
-		client.CallOptions.ListImportJobs = append(client.CallOptions.ListImportJobs, gax.WithClientTracing(tracing))
-		client.CallOptions.ListRetiredResources = append(client.CallOptions.ListRetiredResources, gax.WithClientTracing(tracing))
-		client.CallOptions.GetKeyRing = append(client.CallOptions.GetKeyRing, gax.WithClientTracing(tracing))
-		client.CallOptions.GetCryptoKey = append(client.CallOptions.GetCryptoKey, gax.WithClientTracing(tracing))
-		client.CallOptions.GetCryptoKeyVersion = append(client.CallOptions.GetCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.GetPublicKey = append(client.CallOptions.GetPublicKey, gax.WithClientTracing(tracing))
-		client.CallOptions.GetImportJob = append(client.CallOptions.GetImportJob, gax.WithClientTracing(tracing))
-		client.CallOptions.GetRetiredResource = append(client.CallOptions.GetRetiredResource, gax.WithClientTracing(tracing))
-		client.CallOptions.CreateKeyRing = append(client.CallOptions.CreateKeyRing, gax.WithClientTracing(tracing))
-		client.CallOptions.CreateCryptoKey = append(client.CallOptions.CreateCryptoKey, gax.WithClientTracing(tracing))
-		client.CallOptions.CreateCryptoKeyVersion = append(client.CallOptions.CreateCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.DeleteCryptoKey = append(client.CallOptions.DeleteCryptoKey, gax.WithClientTracing(tracing))
-		client.CallOptions.DeleteCryptoKeyVersion = append(client.CallOptions.DeleteCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.ImportCryptoKeyVersion = append(client.CallOptions.ImportCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.ImportTrustedKeyWrappedCryptoKeyVersion = append(client.CallOptions.ImportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.ExportTrustedKeyWrappedCryptoKeyVersion = append(client.CallOptions.ExportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.CreateImportJob = append(client.CallOptions.CreateImportJob, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateCryptoKey = append(client.CallOptions.UpdateCryptoKey, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateCryptoKeyVersion = append(client.CallOptions.UpdateCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateCryptoKeyPrimaryVersion = append(client.CallOptions.UpdateCryptoKeyPrimaryVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.DestroyCryptoKeyVersion = append(client.CallOptions.DestroyCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.RestoreCryptoKeyVersion = append(client.CallOptions.RestoreCryptoKeyVersion, gax.WithClientTracing(tracing))
-		client.CallOptions.Encrypt = append(client.CallOptions.Encrypt, gax.WithClientTracing(tracing))
-		client.CallOptions.Decrypt = append(client.CallOptions.Decrypt, gax.WithClientTracing(tracing))
-		client.CallOptions.RawEncrypt = append(client.CallOptions.RawEncrypt, gax.WithClientTracing(tracing))
-		client.CallOptions.RawDecrypt = append(client.CallOptions.RawDecrypt, gax.WithClientTracing(tracing))
-		client.CallOptions.AsymmetricSign = append(client.CallOptions.AsymmetricSign, gax.WithClientTracing(tracing))
-		client.CallOptions.AsymmetricDecrypt = append(client.CallOptions.AsymmetricDecrypt, gax.WithClientTracing(tracing))
-		client.CallOptions.MacSign = append(client.CallOptions.MacSign, gax.WithClientTracing(tracing))
-		client.CallOptions.MacVerify = append(client.CallOptions.MacVerify, gax.WithClientTracing(tracing))
-		client.CallOptions.Decapsulate = append(client.CallOptions.Decapsulate, gax.WithClientTracing(tracing))
-		client.CallOptions.GenerateRandomBytes = append(client.CallOptions.GenerateRandomBytes, gax.WithClientTracing(tracing))
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
-		client.CallOptions.ListKeyRings = append(client.CallOptions.ListKeyRings, gax.WithClientLogging(logging))
-		client.CallOptions.ListCryptoKeys = append(client.CallOptions.ListCryptoKeys, gax.WithClientLogging(logging))
-		client.CallOptions.ListCryptoKeyVersions = append(client.CallOptions.ListCryptoKeyVersions, gax.WithClientLogging(logging))
-		client.CallOptions.ListImportJobs = append(client.CallOptions.ListImportJobs, gax.WithClientLogging(logging))
-		client.CallOptions.ListRetiredResources = append(client.CallOptions.ListRetiredResources, gax.WithClientLogging(logging))
-		client.CallOptions.GetKeyRing = append(client.CallOptions.GetKeyRing, gax.WithClientLogging(logging))
-		client.CallOptions.GetCryptoKey = append(client.CallOptions.GetCryptoKey, gax.WithClientLogging(logging))
-		client.CallOptions.GetCryptoKeyVersion = append(client.CallOptions.GetCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.GetPublicKey = append(client.CallOptions.GetPublicKey, gax.WithClientLogging(logging))
-		client.CallOptions.GetImportJob = append(client.CallOptions.GetImportJob, gax.WithClientLogging(logging))
-		client.CallOptions.GetRetiredResource = append(client.CallOptions.GetRetiredResource, gax.WithClientLogging(logging))
-		client.CallOptions.CreateKeyRing = append(client.CallOptions.CreateKeyRing, gax.WithClientLogging(logging))
-		client.CallOptions.CreateCryptoKey = append(client.CallOptions.CreateCryptoKey, gax.WithClientLogging(logging))
-		client.CallOptions.CreateCryptoKeyVersion = append(client.CallOptions.CreateCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.DeleteCryptoKey = append(client.CallOptions.DeleteCryptoKey, gax.WithClientLogging(logging))
-		client.CallOptions.DeleteCryptoKeyVersion = append(client.CallOptions.DeleteCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.ImportCryptoKeyVersion = append(client.CallOptions.ImportCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.ImportTrustedKeyWrappedCryptoKeyVersion = append(client.CallOptions.ImportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.ExportTrustedKeyWrappedCryptoKeyVersion = append(client.CallOptions.ExportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.CreateImportJob = append(client.CallOptions.CreateImportJob, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateCryptoKey = append(client.CallOptions.UpdateCryptoKey, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateCryptoKeyVersion = append(client.CallOptions.UpdateCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateCryptoKeyPrimaryVersion = append(client.CallOptions.UpdateCryptoKeyPrimaryVersion, gax.WithClientLogging(logging))
-		client.CallOptions.DestroyCryptoKeyVersion = append(client.CallOptions.DestroyCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.RestoreCryptoKeyVersion = append(client.CallOptions.RestoreCryptoKeyVersion, gax.WithClientLogging(logging))
-		client.CallOptions.Encrypt = append(client.CallOptions.Encrypt, gax.WithClientLogging(logging))
-		client.CallOptions.Decrypt = append(client.CallOptions.Decrypt, gax.WithClientLogging(logging))
-		client.CallOptions.RawEncrypt = append(client.CallOptions.RawEncrypt, gax.WithClientLogging(logging))
-		client.CallOptions.RawDecrypt = append(client.CallOptions.RawDecrypt, gax.WithClientLogging(logging))
-		client.CallOptions.AsymmetricSign = append(client.CallOptions.AsymmetricSign, gax.WithClientLogging(logging))
-		client.CallOptions.AsymmetricDecrypt = append(client.CallOptions.AsymmetricDecrypt, gax.WithClientLogging(logging))
-		client.CallOptions.MacSign = append(client.CallOptions.MacSign, gax.WithClientLogging(logging))
-		client.CallOptions.MacVerify = append(client.CallOptions.MacVerify, gax.WithClientLogging(logging))
-		client.CallOptions.Decapsulate = append(client.CallOptions.Decapsulate, gax.WithClientLogging(logging))
-		client.CallOptions.GenerateRandomBytes = append(client.CallOptions.GenerateRandomBytes, gax.WithClientLogging(logging))
-		client.CallOptions.GetLocation = append(client.CallOptions.GetLocation, gax.WithClientLogging(logging))
-		client.CallOptions.ListLocations = append(client.CallOptions.ListLocations, gax.WithClientLogging(logging))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientLogging(logging))
-		client.CallOptions.GetOperation = append(client.CallOptions.GetOperation, gax.WithClientLogging(logging))
-	}
-
 	client.internalClient = c
 
 	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
@@ -1674,16 +1502,7 @@ func NewKeyManagementRESTClient(ctx context.Context, o
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
@@ -1698,168 +1517,6 @@ func NewKeyManagementRESTClient(ctx context.Context, o
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
-		callOpts.ImportTrustedKeyWrappedCryptoKeyVersion = append(callOpts.ImportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientMetrics(metrics))
-		callOpts.ExportTrustedKeyWrappedCryptoKeyVersion = append(callOpts.ExportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientMetrics(metrics))
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
-		callOpts.ListKeyRings = append(callOpts.ListKeyRings, gax.WithClientTracing(tracing))
-		callOpts.ListCryptoKeys = append(callOpts.ListCryptoKeys, gax.WithClientTracing(tracing))
-		callOpts.ListCryptoKeyVersions = append(callOpts.ListCryptoKeyVersions, gax.WithClientTracing(tracing))
-		callOpts.ListImportJobs = append(callOpts.ListImportJobs, gax.WithClientTracing(tracing))
-		callOpts.ListRetiredResources = append(callOpts.ListRetiredResources, gax.WithClientTracing(tracing))
-		callOpts.GetKeyRing = append(callOpts.GetKeyRing, gax.WithClientTracing(tracing))
-		callOpts.GetCryptoKey = append(callOpts.GetCryptoKey, gax.WithClientTracing(tracing))
-		callOpts.GetCryptoKeyVersion = append(callOpts.GetCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.GetPublicKey = append(callOpts.GetPublicKey, gax.WithClientTracing(tracing))
-		callOpts.GetImportJob = append(callOpts.GetImportJob, gax.WithClientTracing(tracing))
-		callOpts.GetRetiredResource = append(callOpts.GetRetiredResource, gax.WithClientTracing(tracing))
-		callOpts.CreateKeyRing = append(callOpts.CreateKeyRing, gax.WithClientTracing(tracing))
-		callOpts.CreateCryptoKey = append(callOpts.CreateCryptoKey, gax.WithClientTracing(tracing))
-		callOpts.CreateCryptoKeyVersion = append(callOpts.CreateCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.DeleteCryptoKey = append(callOpts.DeleteCryptoKey, gax.WithClientTracing(tracing))
-		callOpts.DeleteCryptoKeyVersion = append(callOpts.DeleteCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.ImportCryptoKeyVersion = append(callOpts.ImportCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.ImportTrustedKeyWrappedCryptoKeyVersion = append(callOpts.ImportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.ExportTrustedKeyWrappedCryptoKeyVersion = append(callOpts.ExportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.CreateImportJob = append(callOpts.CreateImportJob, gax.WithClientTracing(tracing))
-		callOpts.UpdateCryptoKey = append(callOpts.UpdateCryptoKey, gax.WithClientTracing(tracing))
-		callOpts.UpdateCryptoKeyVersion = append(callOpts.UpdateCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.UpdateCryptoKeyPrimaryVersion = append(callOpts.UpdateCryptoKeyPrimaryVersion, gax.WithClientTracing(tracing))
-		callOpts.DestroyCryptoKeyVersion = append(callOpts.DestroyCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.RestoreCryptoKeyVersion = append(callOpts.RestoreCryptoKeyVersion, gax.WithClientTracing(tracing))
-		callOpts.Encrypt = append(callOpts.Encrypt, gax.WithClientTracing(tracing))
-		callOpts.Decrypt = append(callOpts.Decrypt, gax.WithClientTracing(tracing))
-		callOpts.RawEncrypt = append(callOpts.RawEncrypt, gax.WithClientTracing(tracing))
-		callOpts.RawDecrypt = append(callOpts.RawDecrypt, gax.WithClientTracing(tracing))
-		callOpts.AsymmetricSign = append(callOpts.AsymmetricSign, gax.WithClientTracing(tracing))
-		callOpts.AsymmetricDecrypt = append(callOpts.AsymmetricDecrypt, gax.WithClientTracing(tracing))
-		callOpts.MacSign = append(callOpts.MacSign, gax.WithClientTracing(tracing))
-		callOpts.MacVerify = append(callOpts.MacVerify, gax.WithClientTracing(tracing))
-		callOpts.Decapsulate = append(callOpts.Decapsulate, gax.WithClientTracing(tracing))
-		callOpts.GenerateRandomBytes = append(callOpts.GenerateRandomBytes, gax.WithClientTracing(tracing))
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
-		callOpts.ListKeyRings = append(callOpts.ListKeyRings, gax.WithClientLogging(logging))
-		callOpts.ListCryptoKeys = append(callOpts.ListCryptoKeys, gax.WithClientLogging(logging))
-		callOpts.ListCryptoKeyVersions = append(callOpts.ListCryptoKeyVersions, gax.WithClientLogging(logging))
-		callOpts.ListImportJobs = append(callOpts.ListImportJobs, gax.WithClientLogging(logging))
-		callOpts.ListRetiredResources = append(callOpts.ListRetiredResources, gax.WithClientLogging(logging))
-		callOpts.GetKeyRing = append(callOpts.GetKeyRing, gax.WithClientLogging(logging))
-		callOpts.GetCryptoKey = append(callOpts.GetCryptoKey, gax.WithClientLogging(logging))
-		callOpts.GetCryptoKeyVersion = append(callOpts.GetCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.GetPublicKey = append(callOpts.GetPublicKey, gax.WithClientLogging(logging))
-		callOpts.GetImportJob = append(callOpts.GetImportJob, gax.WithClientLogging(logging))
-		callOpts.GetRetiredResource = append(callOpts.GetRetiredResource, gax.WithClientLogging(logging))
-		callOpts.CreateKeyRing = append(callOpts.CreateKeyRing, gax.WithClientLogging(logging))
-		callOpts.CreateCryptoKey = append(callOpts.CreateCryptoKey, gax.WithClientLogging(logging))
-		callOpts.CreateCryptoKeyVersion = append(callOpts.CreateCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.DeleteCryptoKey = append(callOpts.DeleteCryptoKey, gax.WithClientLogging(logging))
-		callOpts.DeleteCryptoKeyVersion = append(callOpts.DeleteCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.ImportCryptoKeyVersion = append(callOpts.ImportCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.ImportTrustedKeyWrappedCryptoKeyVersion = append(callOpts.ImportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.ExportTrustedKeyWrappedCryptoKeyVersion = append(callOpts.ExportTrustedKeyWrappedCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.CreateImportJob = append(callOpts.CreateImportJob, gax.WithClientLogging(logging))
-		callOpts.UpdateCryptoKey = append(callOpts.UpdateCryptoKey, gax.WithClientLogging(logging))
-		callOpts.UpdateCryptoKeyVersion = append(callOpts.UpdateCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.UpdateCryptoKeyPrimaryVersion = append(callOpts.UpdateCryptoKeyPrimaryVersion, gax.WithClientLogging(logging))
-		callOpts.DestroyCryptoKeyVersion = append(callOpts.DestroyCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.RestoreCryptoKeyVersion = append(callOpts.RestoreCryptoKeyVersion, gax.WithClientLogging(logging))
-		callOpts.Encrypt = append(callOpts.Encrypt, gax.WithClientLogging(logging))
-		callOpts.Decrypt = append(callOpts.Decrypt, gax.WithClientLogging(logging))
-		callOpts.RawEncrypt = append(callOpts.RawEncrypt, gax.WithClientLogging(logging))
-		callOpts.RawDecrypt = append(callOpts.RawDecrypt, gax.WithClientLogging(logging))
-		callOpts.AsymmetricSign = append(callOpts.AsymmetricSign, gax.WithClientLogging(logging))
-		callOpts.AsymmetricDecrypt = append(callOpts.AsymmetricDecrypt, gax.WithClientLogging(logging))
-		callOpts.MacSign = append(callOpts.MacSign, gax.WithClientLogging(logging))
-		callOpts.MacVerify = append(callOpts.MacVerify, gax.WithClientLogging(logging))
-		callOpts.Decapsulate = append(callOpts.Decapsulate, gax.WithClientLogging(logging))
-		callOpts.GenerateRandomBytes = append(callOpts.GenerateRandomBytes, gax.WithClientLogging(logging))
-		callOpts.GetLocation = append(callOpts.GetLocation, gax.WithClientLogging(logging))
-		callOpts.ListLocations = append(callOpts.ListLocations, gax.WithClientLogging(logging))
-		callOpts.GetIamPolicy = append(callOpts.GetIamPolicy, gax.WithClientLogging(logging))
-		callOpts.SetIamPolicy = append(callOpts.SetIamPolicy, gax.WithClientLogging(logging))
-		callOpts.TestIamPermissions = append(callOpts.TestIamPermissions, gax.WithClientLogging(logging))
-		callOpts.GetOperation = append(callOpts.GetOperation, gax.WithClientLogging(logging))
-	}
-
 	lroOpts := []option.ClientOption{
 		option.WithHTTPClient(httpClient),
 		option.WithEndpoint(endpoint),
@@ -1915,9 +1572,6 @@ func (c *keyManagementGRPCClient) ListKeyRings(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListKeyRings")
-	}
 	opts = append((*c.CallOptions).ListKeyRings[0:len((*c.CallOptions).ListKeyRings):len((*c.CallOptions).ListKeyRings)], opts...)
 	it := &KeyRingIterator{}
 	req = proto.CloneOf(req)
@@ -1964,9 +1618,6 @@ func (c *keyManagementGRPCClient) ListCryptoKeys(ctx c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListCryptoKeys")
-	}
 	opts = append((*c.CallOptions).ListCryptoKeys[0:len((*c.CallOptions).ListCryptoKeys):len((*c.CallOptions).ListCryptoKeys)], opts...)
 	it := &CryptoKeyIterator{}
 	req = proto.CloneOf(req)
@@ -2013,9 +1664,6 @@ func (c *keyManagementGRPCClient) ListCryptoKeyVersion
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListCryptoKeyVersions")
-	}
 	opts = append((*c.CallOptions).ListCryptoKeyVersions[0:len((*c.CallOptions).ListCryptoKeyVersions):len((*c.CallOptions).ListCryptoKeyVersions)], opts...)
 	it := &CryptoKeyVersionIterator{}
 	req = proto.CloneOf(req)
@@ -2062,9 +1710,6 @@ func (c *keyManagementGRPCClient) ListImportJobs(ctx c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListImportJobs")
-	}
 	opts = append((*c.CallOptions).ListImportJobs[0:len((*c.CallOptions).ListImportJobs):len((*c.CallOptions).ListImportJobs)], opts...)
 	it := &ImportJobIterator{}
 	req = proto.CloneOf(req)
@@ -2111,9 +1756,6 @@ func (c *keyManagementGRPCClient) ListRetiredResources
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ListRetiredResources")
-	}
 	opts = append((*c.CallOptions).ListRetiredResources[0:len((*c.CallOptions).ListRetiredResources):len((*c.CallOptions).ListRetiredResources)], opts...)
 	it := &RetiredResourceIterator{}
 	req = proto.CloneOf(req)
@@ -2160,9 +1802,6 @@ func (c *keyManagementGRPCClient) GetKeyRing(ctx conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetKeyRing")
-	}
 	opts = append((*c.CallOptions).GetKeyRing[0:len((*c.CallOptions).GetKeyRing):len((*c.CallOptions).GetKeyRing)], opts...)
 	var resp *kmspb.KeyRing
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2181,9 +1820,6 @@ func (c *keyManagementGRPCClient) GetCryptoKey(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKey")
-	}
 	opts = append((*c.CallOptions).GetCryptoKey[0:len((*c.CallOptions).GetCryptoKey):len((*c.CallOptions).GetCryptoKey)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2202,9 +1838,6 @@ func (c *keyManagementGRPCClient) GetCryptoKeyVersion(
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).GetCryptoKeyVersion[0:len((*c.CallOptions).GetCryptoKeyVersion):len((*c.CallOptions).GetCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2223,9 +1856,6 @@ func (c *keyManagementGRPCClient) GetPublicKey(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetPublicKey")
-	}
 	opts = append((*c.CallOptions).GetPublicKey[0:len((*c.CallOptions).GetPublicKey):len((*c.CallOptions).GetPublicKey)], opts...)
 	var resp *kmspb.PublicKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2244,9 +1874,6 @@ func (c *keyManagementGRPCClient) GetImportJob(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetImportJob")
-	}
 	opts = append((*c.CallOptions).GetImportJob[0:len((*c.CallOptions).GetImportJob):len((*c.CallOptions).GetImportJob)], opts...)
 	var resp *kmspb.ImportJob
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2265,9 +1892,6 @@ func (c *keyManagementGRPCClient) GetRetiredResource(c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetRetiredResource")
-	}
 	opts = append((*c.CallOptions).GetRetiredResource[0:len((*c.CallOptions).GetRetiredResource):len((*c.CallOptions).GetRetiredResource)], opts...)
 	var resp *kmspb.RetiredResource
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2286,9 +1910,6 @@ func (c *keyManagementGRPCClient) CreateKeyRing(ctx co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateKeyRing")
-	}
 	opts = append((*c.CallOptions).CreateKeyRing[0:len((*c.CallOptions).CreateKeyRing):len((*c.CallOptions).CreateKeyRing)], opts...)
 	var resp *kmspb.KeyRing
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2307,9 +1928,6 @@ func (c *keyManagementGRPCClient) CreateCryptoKey(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKey")
-	}
 	opts = append((*c.CallOptions).CreateCryptoKey[0:len((*c.CallOptions).CreateCryptoKey):len((*c.CallOptions).CreateCryptoKey)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2328,9 +1946,6 @@ func (c *keyManagementGRPCClient) CreateCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).CreateCryptoKeyVersion[0:len((*c.CallOptions).CreateCryptoKeyVersion):len((*c.CallOptions).CreateCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2349,9 +1964,6 @@ func (c *keyManagementGRPCClient) DeleteCryptoKey(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKey")
-	}
 	opts = append((*c.CallOptions).DeleteCryptoKey[0:len((*c.CallOptions).DeleteCryptoKey):len((*c.CallOptions).DeleteCryptoKey)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2363,9 +1975,6 @@ func (c *keyManagementGRPCClient) DeleteCryptoKey(ctx 
 		return nil, err
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &DeleteCryptoKeyOperation{
 		lro: lro,
 	}, nil
@@ -2376,9 +1985,6 @@ func (c *keyManagementGRPCClient) DeleteCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).DeleteCryptoKeyVersion[0:len((*c.CallOptions).DeleteCryptoKeyVersion):len((*c.CallOptions).DeleteCryptoKeyVersion)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2390,9 +1996,6 @@ func (c *keyManagementGRPCClient) DeleteCryptoKeyVersi
 		return nil, err
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyVersionOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &DeleteCryptoKeyVersionOperation{
 		lro: lro,
 	}, nil
@@ -2403,9 +2006,6 @@ func (c *keyManagementGRPCClient) ImportCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ImportCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).ImportCryptoKeyVersion[0:len((*c.CallOptions).ImportCryptoKeyVersion):len((*c.CallOptions).ImportCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2424,9 +2024,6 @@ func (c *keyManagementGRPCClient) ImportTrustedKeyWrap
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ImportTrustedKeyWrappedCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).ImportTrustedKeyWrappedCryptoKeyVersion[0:len((*c.CallOptions).ImportTrustedKeyWrappedCryptoKeyVersion):len((*c.CallOptions).ImportTrustedKeyWrappedCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2445,9 +2042,6 @@ func (c *keyManagementGRPCClient) ExportTrustedKeyWrap
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ExportTrustedKeyWrappedCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).ExportTrustedKeyWrappedCryptoKeyVersion[0:len((*c.CallOptions).ExportTrustedKeyWrappedCryptoKeyVersion):len((*c.CallOptions).ExportTrustedKeyWrappedCryptoKeyVersion)], opts...)
 	var resp *kmspb.ExportTrustedKeyWrappedCryptoKeyVersionResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2466,9 +2060,6 @@ func (c *keyManagementGRPCClient) CreateImportJob(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateImportJob")
-	}
 	opts = append((*c.CallOptions).CreateImportJob[0:len((*c.CallOptions).CreateImportJob):len((*c.CallOptions).CreateImportJob)], opts...)
 	var resp *kmspb.ImportJob
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2487,9 +2078,7 @@ func (c *keyManagementGRPCClient) UpdateCryptoKey(ctx 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKey")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKey[0:len((*c.CallOptions).UpdateCryptoKey):len((*c.CallOptions).UpdateCryptoKey)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2508,9 +2097,7 @@ func (c *keyManagementGRPCClient) UpdateCryptoKeyVersi
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKeyVersion")
-	}
+
 	opts = append((*c.CallOptions).UpdateCryptoKeyVersion[0:len((*c.CallOptions).UpdateCryptoKeyVersion):len((*c.CallOptions).UpdateCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2529,9 +2116,6 @@ func (c *keyManagementGRPCClient) UpdateCryptoKeyPrima
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKeyPrimaryVersion")
-	}
 	opts = append((*c.CallOptions).UpdateCryptoKeyPrimaryVersion[0:len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion):len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion)], opts...)
 	var resp *kmspb.CryptoKey
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2550,9 +2134,6 @@ func (c *keyManagementGRPCClient) DestroyCryptoKeyVers
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DestroyCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).DestroyCryptoKeyVersion[0:len((*c.CallOptions).DestroyCryptoKeyVersion):len((*c.CallOptions).DestroyCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2571,9 +2152,6 @@ func (c *keyManagementGRPCClient) RestoreCryptoKeyVers
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RestoreCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).RestoreCryptoKeyVersion[0:len((*c.CallOptions).RestoreCryptoKeyVersion):len((*c.CallOptions).RestoreCryptoKeyVersion)], opts...)
 	var resp *kmspb.CryptoKeyVersion
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2592,9 +2170,6 @@ func (c *keyManagementGRPCClient) Encrypt(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Encrypt")
-	}
 	opts = append((*c.CallOptions).Encrypt[0:len((*c.CallOptions).Encrypt):len((*c.CallOptions).Encrypt)], opts...)
 	var resp *kmspb.EncryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2613,9 +2188,6 @@ func (c *keyManagementGRPCClient) Decrypt(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decrypt")
-	}
 	opts = append((*c.CallOptions).Decrypt[0:len((*c.CallOptions).Decrypt):len((*c.CallOptions).Decrypt)], opts...)
 	var resp *kmspb.DecryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2634,9 +2206,7 @@ func (c *keyManagementGRPCClient) RawEncrypt(ctx conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RawEncrypt")
-	}
+
 	opts = append((*c.CallOptions).RawEncrypt[0:len((*c.CallOptions).RawEncrypt):len((*c.CallOptions).RawEncrypt)], opts...)
 	var resp *kmspb.RawEncryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2655,9 +2225,7 @@ func (c *keyManagementGRPCClient) RawDecrypt(ctx conte
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RawDecrypt")
-	}
+
 	opts = append((*c.CallOptions).RawDecrypt[0:len((*c.CallOptions).RawDecrypt):len((*c.CallOptions).RawDecrypt)], opts...)
 	var resp *kmspb.RawDecryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2676,9 +2244,6 @@ func (c *keyManagementGRPCClient) AsymmetricSign(ctx c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricSign")
-	}
 	opts = append((*c.CallOptions).AsymmetricSign[0:len((*c.CallOptions).AsymmetricSign):len((*c.CallOptions).AsymmetricSign)], opts...)
 	var resp *kmspb.AsymmetricSignResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2697,9 +2262,6 @@ func (c *keyManagementGRPCClient) AsymmetricDecrypt(ct
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricDecrypt")
-	}
 	opts = append((*c.CallOptions).AsymmetricDecrypt[0:len((*c.CallOptions).AsymmetricDecrypt):len((*c.CallOptions).AsymmetricDecrypt)], opts...)
 	var resp *kmspb.AsymmetricDecryptResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2718,9 +2280,6 @@ func (c *keyManagementGRPCClient) MacSign(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacSign")
-	}
 	opts = append((*c.CallOptions).MacSign[0:len((*c.CallOptions).MacSign):len((*c.CallOptions).MacSign)], opts...)
 	var resp *kmspb.MacSignResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2739,9 +2298,6 @@ func (c *keyManagementGRPCClient) MacVerify(ctx contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacVerify")
-	}
 	opts = append((*c.CallOptions).MacVerify[0:len((*c.CallOptions).MacVerify):len((*c.CallOptions).MacVerify)], opts...)
 	var resp *kmspb.MacVerifyResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2760,9 +2316,6 @@ func (c *keyManagementGRPCClient) Decapsulate(ctx cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decapsulate")
-	}
 	opts = append((*c.CallOptions).Decapsulate[0:len((*c.CallOptions).Decapsulate):len((*c.CallOptions).Decapsulate)], opts...)
 	var resp *kmspb.DecapsulateResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2781,9 +2334,7 @@ func (c *keyManagementGRPCClient) GenerateRandomBytes(
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GenerateRandomBytes")
-	}
+
 	opts = append((*c.CallOptions).GenerateRandomBytes[0:len((*c.CallOptions).GenerateRandomBytes):len((*c.CallOptions).GenerateRandomBytes)], opts...)
 	var resp *kmspb.GenerateRandomBytesResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2802,9 +2353,7 @@ func (c *keyManagementGRPCClient) GetLocation(ctx cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/GetLocation")
-	}
+
 	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
 	var resp *locationpb.Location
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2823,9 +2372,7 @@ func (c *keyManagementGRPCClient) ListLocations(ctx co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.location.Locations/ListLocations")
-	}
+
 	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
 	it := &LocationIterator{}
 	req = proto.CloneOf(req)
@@ -2872,9 +2419,6 @@ func (c *keyManagementGRPCClient) GetIamPolicy(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/GetIamPolicy")
-	}
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2893,9 +2437,6 @@ func (c *keyManagementGRPCClient) SetIamPolicy(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/SetIamPolicy")
-	}
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2914,9 +2455,6 @@ func (c *keyManagementGRPCClient) TestIamPermissions(c
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.iam.v1.IAMPolicy/TestIamPermissions")
-	}
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	var resp *iampb.TestIamPermissionsResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -2935,9 +2473,7 @@ func (c *keyManagementGRPCClient) GetOperation(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -3392,10 +2928,6 @@ func (c *keyManagementRESTClient) GetKeyRing(ctx conte
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetKeyRing")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*}")
-	}
 	opts = append((*c.CallOptions).GetKeyRing[0:len((*c.CallOptions).GetKeyRing):len((*c.CallOptions).GetKeyRing)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.KeyRing{}
@@ -3448,10 +2980,6 @@ func (c *keyManagementRESTClient) GetCryptoKey(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}")
-	}
 	opts = append((*c.CallOptions).GetCryptoKey[0:len((*c.CallOptions).GetCryptoKey):len((*c.CallOptions).GetCryptoKey)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKey{}
@@ -3503,10 +3031,6 @@ func (c *keyManagementRESTClient) GetCryptoKeyVersion(
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}")
-	}
 	opts = append((*c.CallOptions).GetCryptoKeyVersion[0:len((*c.CallOptions).GetCryptoKeyVersion):len((*c.CallOptions).GetCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -3565,10 +3089,6 @@ func (c *keyManagementRESTClient) GetPublicKey(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetPublicKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}/publicKey")
-	}
 	opts = append((*c.CallOptions).GetPublicKey[0:len((*c.CallOptions).GetPublicKey):len((*c.CallOptions).GetPublicKey)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.PublicKey{}
@@ -3622,10 +3142,6 @@ func (c *keyManagementRESTClient) GetImportJob(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetImportJob")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/importJobs/*}")
-	}
 	opts = append((*c.CallOptions).GetImportJob[0:len((*c.CallOptions).GetImportJob):len((*c.CallOptions).GetImportJob)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.ImportJob{}
@@ -3678,10 +3194,6 @@ func (c *keyManagementRESTClient) GetRetiredResource(c
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/GetRetiredResource")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/retiredResources/*}")
-	}
 	opts = append((*c.CallOptions).GetRetiredResource[0:len((*c.CallOptions).GetRetiredResource):len((*c.CallOptions).GetRetiredResource)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.RetiredResource{}
@@ -3741,10 +3253,6 @@ func (c *keyManagementRESTClient) CreateKeyRing(ctx co
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateKeyRing")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*}/keyRings")
-	}
 	opts = append((*c.CallOptions).CreateKeyRing[0:len((*c.CallOptions).CreateKeyRing):len((*c.CallOptions).CreateKeyRing)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.KeyRing{}
@@ -3814,10 +3322,6 @@ func (c *keyManagementRESTClient) CreateCryptoKey(ctx 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*}/cryptoKeys")
-	}
 	opts = append((*c.CallOptions).CreateCryptoKey[0:len((*c.CallOptions).CreateCryptoKey):len((*c.CallOptions).CreateCryptoKey)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKey{}
@@ -3880,10 +3384,6 @@ func (c *keyManagementRESTClient) CreateCryptoKeyVersi
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*/cryptoKeys/*}/cryptoKeyVersions")
-	}
 	opts = append((*c.CallOptions).CreateCryptoKeyVersion[0:len((*c.CallOptions).CreateCryptoKeyVersion):len((*c.CallOptions).CreateCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -3939,10 +3439,6 @@ func (c *keyManagementRESTClient) DeleteCryptoKey(ctx 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKey")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}")
-	}
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
 	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -3972,9 +3468,6 @@ func (c *keyManagementRESTClient) DeleteCryptoKey(ctx 
 
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &DeleteCryptoKeyOperation{
 		lro:      lro,
 		pollPath: override,
@@ -4010,10 +3503,6 @@ func (c *keyManagementRESTClient) DeleteCryptoKeyVersi
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DeleteCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}")
-	}
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
 	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -4043,9 +3532,6 @@ func (c *keyManagementRESTClient) DeleteCryptoKeyVersi
 
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyVersionOperation")
-	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
-	}
 	return &DeleteCryptoKeyVersionOperation{
 		lro:      lro,
 		pollPath: override,
@@ -4084,10 +3570,6 @@ func (c *keyManagementRESTClient) ImportCryptoKeyVersi
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ImportCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*/cryptoKeys/*}/cryptoKeyVersions:import")
-	}
 	opts = append((*c.CallOptions).ImportCryptoKeyVersion[0:len((*c.CallOptions).ImportCryptoKeyVersion):len((*c.CallOptions).ImportCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -4155,10 +3637,6 @@ func (c *keyManagementRESTClient) ImportTrustedKeyWrap
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ImportTrustedKeyWrappedCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*/cryptoKeys/*}/cryptoKeyVersions:importTrustedKeyWrappedCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).ImportTrustedKeyWrappedCryptoKeyVersion[0:len((*c.CallOptions).ImportTrustedKeyWrappedCryptoKeyVersion):len((*c.CallOptions).ImportTrustedKeyWrappedCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -4220,10 +3698,6 @@ func (c *keyManagementRESTClient) ExportTrustedKeyWrap
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/ExportTrustedKeyWrappedCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:exportTrustedKeyWrappedCryptoKeyVersion")
-	}
 	opts = append((*c.CallOptions).ExportTrustedKeyWrappedCryptoKeyVersion[0:len((*c.CallOptions).ExportTrustedKeyWrappedCryptoKeyVersion):len((*c.CallOptions).ExportTrustedKeyWrappedCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.ExportTrustedKeyWrappedCryptoKeyVersionResponse{}
@@ -4286,10 +3760,6 @@ func (c *keyManagementRESTClient) CreateImportJob(ctx 
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/CreateImportJob")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{parent=projects/*/locations/*/keyRings/*}/importJobs")
-	}
 	opts = append((*c.CallOptions).CreateImportJob[0:len((*c.CallOptions).CreateImportJob):len((*c.CallOptions).CreateImportJob)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.ImportJob{}
@@ -4354,10 +3824,7 @@ func (c *keyManagementRESTClient) UpdateCryptoKey(ctx 
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
@@ -4433,10 +3900,7 @@ func (c *keyManagementRESTClient) UpdateCryptoKeyVersi
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
@@ -4498,10 +3962,6 @@ func (c *keyManagementRESTClient) UpdateCryptoKeyPrima
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/UpdateCryptoKeyPrimaryVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}:updatePrimaryVersion")
-	}
 	opts = append((*c.CallOptions).UpdateCryptoKeyPrimaryVersion[0:len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion):len((*c.CallOptions).UpdateCryptoKeyPrimaryVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKey{}
@@ -4578,10 +4038,6 @@ func (c *keyManagementRESTClient) DestroyCryptoKeyVers
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/DestroyCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:destroy")
-	}
 	opts = append((*c.CallOptions).DestroyCryptoKeyVersion[0:len((*c.CallOptions).DestroyCryptoKeyVersion):len((*c.CallOptions).DestroyCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -4646,10 +4102,6 @@ func (c *keyManagementRESTClient) RestoreCryptoKeyVers
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/RestoreCryptoKeyVersion")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:restore")
-	}
 	opts = append((*c.CallOptions).RestoreCryptoKeyVersion[0:len((*c.CallOptions).RestoreCryptoKeyVersion):len((*c.CallOptions).RestoreCryptoKeyVersion)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.CryptoKeyVersion{}
@@ -4709,10 +4161,6 @@ func (c *keyManagementRESTClient) Encrypt(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Encrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/**}:encrypt")
-	}
 	opts = append((*c.CallOptions).Encrypt[0:len((*c.CallOptions).Encrypt):len((*c.CallOptions).Encrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.EncryptResponse{}
@@ -4772,10 +4220,6 @@ func (c *keyManagementRESTClient) Decrypt(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*}:decrypt")
-	}
 	opts = append((*c.CallOptions).Decrypt[0:len((*c.CallOptions).Decrypt):len((*c.CallOptions).Decrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.DecryptResponse{}
@@ -4837,10 +4281,7 @@ func (c *keyManagementRESTClient) RawEncrypt(ctx conte
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
@@ -4900,10 +4341,7 @@ func (c *keyManagementRESTClient) RawDecrypt(ctx conte
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
@@ -4964,10 +4402,6 @@ func (c *keyManagementRESTClient) AsymmetricSign(ctx c
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricSign")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:asymmetricSign")
-	}
 	opts = append((*c.CallOptions).AsymmetricSign[0:len((*c.CallOptions).AsymmetricSign):len((*c.CallOptions).AsymmetricSign)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.AsymmetricSignResponse{}
@@ -5028,10 +4462,6 @@ func (c *keyManagementRESTClient) AsymmetricDecrypt(ct
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/AsymmetricDecrypt")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:asymmetricDecrypt")
-	}
 	opts = append((*c.CallOptions).AsymmetricDecrypt[0:len((*c.CallOptions).AsymmetricDecrypt):len((*c.CallOptions).AsymmetricDecrypt)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.AsymmetricDecryptResponse{}
@@ -5090,10 +4520,6 @@ func (c *keyManagementRESTClient) MacSign(ctx context.
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacSign")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:macSign")
-	}
 	opts = append((*c.CallOptions).MacSign[0:len((*c.CallOptions).MacSign):len((*c.CallOptions).MacSign)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.MacSignResponse{}
@@ -5153,10 +4579,6 @@ func (c *keyManagementRESTClient) MacVerify(ctx contex
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/MacVerify")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:macVerify")
-	}
 	opts = append((*c.CallOptions).MacVerify[0:len((*c.CallOptions).MacVerify):len((*c.CallOptions).MacVerify)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.MacVerifyResponse{}
@@ -5217,10 +4639,6 @@ func (c *keyManagementRESTClient) Decapsulate(ctx cont
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.cloud.kms.v1.KeyManagementService/Decapsulate")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*}:decapsulate")
-	}
 	opts = append((*c.CallOptions).Decapsulate[0:len((*c.CallOptions).Decapsulate):len((*c.CallOptions).Decapsulate)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &kmspb.DecapsulateResponse{}
@@ -5278,10 +4696,7 @@ func (c *keyManagementRESTClient) GenerateRandomBytes(
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
@@ -5332,10 +4747,7 @@ func (c *keyManagementRESTClient) GetLocation(ctx cont
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
@@ -5486,10 +4898,6 @@ func (c *keyManagementRESTClient) GetIamPolicy(ctx con
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
@@ -5550,10 +4958,6 @@ func (c *keyManagementRESTClient) SetIamPolicy(ctx con
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
@@ -5616,10 +5020,6 @@ func (c *keyManagementRESTClient) TestIamPermissions(c
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
@@ -5670,10 +5070,7 @@ func (c *keyManagementRESTClient) GetOperation(ctx con
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
