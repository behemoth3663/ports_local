--- vendor/cloud.google.com/go/storage/internal/apiv2/storage_client.go.orig	2026-11-04 11:39:43 UTC
+++ vendor/cloud.google.com/go/storage/internal/apiv2/storage_client.go
@@ -29,7 +29,6 @@ import (
 	iampb "cloud.google.com/go/iam/apiv1/iampb"
 	storagepb "cloud.google.com/go/storage/internal/apiv2/storagepb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -986,16 +985,7 @@ func NewClient(ctx context.Context, opts ...option.Cli
 // any other character (no special directory semantics).
 func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
 	clientOpts := defaultGRPCClientOptions()
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "storage",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/storage/internal/apiv2",
-			"gcp.client.language": "go",
-			"url.domain":          "storage.googleapis.com",
-		}))
-	}
+
 	if newClientHook != nil {
 		hookOpts, err := newClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -1017,117 +1007,7 @@ func NewClient(ctx context.Context, opts ...option.Cli
 		logger:      internaloption.GetLogger(opts),
 	}
 	c.setGoogleClientInfo()
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "storage",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/storage/internal/apiv2",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "storage.googleapis.com",
-			}),
-		)
 
-		client.CallOptions.DeleteBucket = append(client.CallOptions.DeleteBucket, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetBucket = append(client.CallOptions.GetBucket, gax.WithClientMetrics(metrics))
-		client.CallOptions.CreateBucket = append(client.CallOptions.CreateBucket, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListBuckets = append(client.CallOptions.ListBuckets, gax.WithClientMetrics(metrics))
-		client.CallOptions.LockBucketRetentionPolicy = append(client.CallOptions.LockBucketRetentionPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientMetrics(metrics))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateBucket = append(client.CallOptions.UpdateBucket, gax.WithClientMetrics(metrics))
-		client.CallOptions.ComposeObject = append(client.CallOptions.ComposeObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteObject = append(client.CallOptions.DeleteObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.RestoreObject = append(client.CallOptions.RestoreObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.CancelResumableWrite = append(client.CallOptions.CancelResumableWrite, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetObject = append(client.CallOptions.GetObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.ReadObject = append(client.CallOptions.ReadObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.BidiReadObject = append(client.CallOptions.BidiReadObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.UpdateObject = append(client.CallOptions.UpdateObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.WriteObject = append(client.CallOptions.WriteObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.BidiWriteObject = append(client.CallOptions.BidiWriteObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.ListObjects = append(client.CallOptions.ListObjects, gax.WithClientMetrics(metrics))
-		client.CallOptions.RewriteObject = append(client.CallOptions.RewriteObject, gax.WithClientMetrics(metrics))
-		client.CallOptions.StartResumableWrite = append(client.CallOptions.StartResumableWrite, gax.WithClientMetrics(metrics))
-		client.CallOptions.QueryWriteStatus = append(client.CallOptions.QueryWriteStatus, gax.WithClientMetrics(metrics))
-		client.CallOptions.MoveObject = append(client.CallOptions.MoveObject, gax.WithClientMetrics(metrics))
-	}
-	if gax.IsFeatureEnabled("TRACING") {
-		tracing := gax.NewClientTracing(
-			gax.WithTracingAttributes(map[string]string{
-				gax.ClientService:  "storage",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/storage/internal/apiv2",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "storage.googleapis.com",
-			}),
-		)
-
-		client.CallOptions.DeleteBucket = append(client.CallOptions.DeleteBucket, gax.WithClientTracing(tracing))
-		client.CallOptions.GetBucket = append(client.CallOptions.GetBucket, gax.WithClientTracing(tracing))
-		client.CallOptions.CreateBucket = append(client.CallOptions.CreateBucket, gax.WithClientTracing(tracing))
-		client.CallOptions.ListBuckets = append(client.CallOptions.ListBuckets, gax.WithClientTracing(tracing))
-		client.CallOptions.LockBucketRetentionPolicy = append(client.CallOptions.LockBucketRetentionPolicy, gax.WithClientTracing(tracing))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientTracing(tracing))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientTracing(tracing))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateBucket = append(client.CallOptions.UpdateBucket, gax.WithClientTracing(tracing))
-		client.CallOptions.ComposeObject = append(client.CallOptions.ComposeObject, gax.WithClientTracing(tracing))
-		client.CallOptions.DeleteObject = append(client.CallOptions.DeleteObject, gax.WithClientTracing(tracing))
-		client.CallOptions.RestoreObject = append(client.CallOptions.RestoreObject, gax.WithClientTracing(tracing))
-		client.CallOptions.CancelResumableWrite = append(client.CallOptions.CancelResumableWrite, gax.WithClientTracing(tracing))
-		client.CallOptions.GetObject = append(client.CallOptions.GetObject, gax.WithClientTracing(tracing))
-		client.CallOptions.ReadObject = append(client.CallOptions.ReadObject, gax.WithClientTracing(tracing))
-		client.CallOptions.BidiReadObject = append(client.CallOptions.BidiReadObject, gax.WithClientTracing(tracing))
-		client.CallOptions.UpdateObject = append(client.CallOptions.UpdateObject, gax.WithClientTracing(tracing))
-		client.CallOptions.WriteObject = append(client.CallOptions.WriteObject, gax.WithClientTracing(tracing))
-		client.CallOptions.BidiWriteObject = append(client.CallOptions.BidiWriteObject, gax.WithClientTracing(tracing))
-		client.CallOptions.ListObjects = append(client.CallOptions.ListObjects, gax.WithClientTracing(tracing))
-		client.CallOptions.RewriteObject = append(client.CallOptions.RewriteObject, gax.WithClientTracing(tracing))
-		client.CallOptions.StartResumableWrite = append(client.CallOptions.StartResumableWrite, gax.WithClientTracing(tracing))
-		client.CallOptions.QueryWriteStatus = append(client.CallOptions.QueryWriteStatus, gax.WithClientTracing(tracing))
-		client.CallOptions.MoveObject = append(client.CallOptions.MoveObject, gax.WithClientTracing(tracing))
-	}
-	if gax.IsFeatureEnabled("LOGGING") {
-		logging := gax.NewClientLogging(
-			gax.WithLoggerProvider(c.logger),
-			gax.WithLoggingAttributes(map[string]string{
-				gax.ClientService:  "storage",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/storage/internal/apiv2",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "storage.googleapis.com",
-			}),
-		)
-
-		client.CallOptions.DeleteBucket = append(client.CallOptions.DeleteBucket, gax.WithClientLogging(logging))
-		client.CallOptions.GetBucket = append(client.CallOptions.GetBucket, gax.WithClientLogging(logging))
-		client.CallOptions.CreateBucket = append(client.CallOptions.CreateBucket, gax.WithClientLogging(logging))
-		client.CallOptions.ListBuckets = append(client.CallOptions.ListBuckets, gax.WithClientLogging(logging))
-		client.CallOptions.LockBucketRetentionPolicy = append(client.CallOptions.LockBucketRetentionPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.GetIamPolicy = append(client.CallOptions.GetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.SetIamPolicy = append(client.CallOptions.SetIamPolicy, gax.WithClientLogging(logging))
-		client.CallOptions.TestIamPermissions = append(client.CallOptions.TestIamPermissions, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateBucket = append(client.CallOptions.UpdateBucket, gax.WithClientLogging(logging))
-		client.CallOptions.ComposeObject = append(client.CallOptions.ComposeObject, gax.WithClientLogging(logging))
-		client.CallOptions.DeleteObject = append(client.CallOptions.DeleteObject, gax.WithClientLogging(logging))
-		client.CallOptions.RestoreObject = append(client.CallOptions.RestoreObject, gax.WithClientLogging(logging))
-		client.CallOptions.CancelResumableWrite = append(client.CallOptions.CancelResumableWrite, gax.WithClientLogging(logging))
-		client.CallOptions.GetObject = append(client.CallOptions.GetObject, gax.WithClientLogging(logging))
-		client.CallOptions.ReadObject = append(client.CallOptions.ReadObject, gax.WithClientLogging(logging))
-		client.CallOptions.BidiReadObject = append(client.CallOptions.BidiReadObject, gax.WithClientLogging(logging))
-		client.CallOptions.UpdateObject = append(client.CallOptions.UpdateObject, gax.WithClientLogging(logging))
-		client.CallOptions.WriteObject = append(client.CallOptions.WriteObject, gax.WithClientLogging(logging))
-		client.CallOptions.BidiWriteObject = append(client.CallOptions.BidiWriteObject, gax.WithClientLogging(logging))
-		client.CallOptions.ListObjects = append(client.CallOptions.ListObjects, gax.WithClientLogging(logging))
-		client.CallOptions.RewriteObject = append(client.CallOptions.RewriteObject, gax.WithClientLogging(logging))
-		client.CallOptions.StartResumableWrite = append(client.CallOptions.StartResumableWrite, gax.WithClientLogging(logging))
-		client.CallOptions.QueryWriteStatus = append(client.CallOptions.QueryWriteStatus, gax.WithClientLogging(logging))
-		client.CallOptions.MoveObject = append(client.CallOptions.MoveObject, gax.WithClientLogging(logging))
-	}
-
 	client.internalClient = c
 
 	return &client, nil
@@ -1172,9 +1052,6 @@ func (c *gRPCClient) DeleteBucket(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/DeleteBucket")
-	}
 	opts = append((*c.CallOptions).DeleteBucket[0:len((*c.CallOptions).DeleteBucket):len((*c.CallOptions).DeleteBucket)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -1198,9 +1075,6 @@ func (c *gRPCClient) GetBucket(ctx context.Context, re
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/GetBucket")
-	}
 	opts = append((*c.CallOptions).GetBucket[0:len((*c.CallOptions).GetBucket):len((*c.CallOptions).GetBucket)], opts...)
 	var resp *storagepb.Bucket
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1231,9 +1105,6 @@ func (c *gRPCClient) CreateBucket(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/CreateBucket")
-	}
 	opts = append((*c.CallOptions).CreateBucket[0:len((*c.CallOptions).CreateBucket):len((*c.CallOptions).CreateBucket)], opts...)
 	var resp *storagepb.Bucket
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1261,9 +1132,6 @@ func (c *gRPCClient) ListBuckets(ctx context.Context, 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/ListBuckets")
-	}
 	opts = append((*c.CallOptions).ListBuckets[0:len((*c.CallOptions).ListBuckets):len((*c.CallOptions).ListBuckets)], opts...)
 	it := &BucketIterator{}
 	req = proto.CloneOf(req)
@@ -1319,9 +1187,6 @@ func (c *gRPCClient) LockBucketRetentionPolicy(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/LockBucketRetentionPolicy")
-	}
 	opts = append((*c.CallOptions).LockBucketRetentionPolicy[0:len((*c.CallOptions).LockBucketRetentionPolicy):len((*c.CallOptions).LockBucketRetentionPolicy)], opts...)
 	var resp *storagepb.Bucket
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1352,9 +1217,6 @@ func (c *gRPCClient) GetIamPolicy(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/GetIamPolicy")
-	}
 	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1385,9 +1247,6 @@ func (c *gRPCClient) SetIamPolicy(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/SetIamPolicy")
-	}
 	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
 	var resp *iampb.Policy
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1421,9 +1280,6 @@ func (c *gRPCClient) TestIamPermissions(ctx context.Co
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/TestIamPermissions")
-	}
 	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
 	var resp *iampb.TestIamPermissionsResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1451,9 +1307,7 @@ func (c *gRPCClient) UpdateBucket(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/UpdateBucket")
-	}
+
 	opts = append((*c.CallOptions).UpdateBucket[0:len((*c.CallOptions).UpdateBucket):len((*c.CallOptions).UpdateBucket)], opts...)
 	var resp *storagepb.Bucket
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1481,9 +1335,6 @@ func (c *gRPCClient) ComposeObject(ctx context.Context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/ComposeObject")
-	}
 	opts = append((*c.CallOptions).ComposeObject[0:len((*c.CallOptions).ComposeObject):len((*c.CallOptions).ComposeObject)], opts...)
 	var resp *storagepb.Object
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1511,9 +1362,6 @@ func (c *gRPCClient) DeleteObject(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/DeleteObject")
-	}
 	opts = append((*c.CallOptions).DeleteObject[0:len((*c.CallOptions).DeleteObject):len((*c.CallOptions).DeleteObject)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -1537,9 +1385,6 @@ func (c *gRPCClient) RestoreObject(ctx context.Context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/RestoreObject")
-	}
 	opts = append((*c.CallOptions).RestoreObject[0:len((*c.CallOptions).RestoreObject):len((*c.CallOptions).RestoreObject)], opts...)
 	var resp *storagepb.Object
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1567,9 +1412,7 @@ func (c *gRPCClient) CancelResumableWrite(ctx context.
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/CancelResumableWrite")
-	}
+
 	opts = append((*c.CallOptions).CancelResumableWrite[0:len((*c.CallOptions).CancelResumableWrite):len((*c.CallOptions).CancelResumableWrite)], opts...)
 	var resp *storagepb.CancelResumableWriteResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1597,9 +1440,6 @@ func (c *gRPCClient) GetObject(ctx context.Context, re
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/GetObject")
-	}
 	opts = append((*c.CallOptions).GetObject[0:len((*c.CallOptions).GetObject):len((*c.CallOptions).GetObject)], opts...)
 	var resp *storagepb.Object
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1627,9 +1467,6 @@ func (c *gRPCClient) ReadObject(ctx context.Context, r
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/ReadObject")
-	}
 	opts = append((*c.CallOptions).ReadObject[0:len((*c.CallOptions).ReadObject):len((*c.CallOptions).ReadObject)], opts...)
 	var resp storagepb.Storage_ReadObjectClient
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1647,9 +1484,7 @@ func (c *gRPCClient) BidiReadObject(ctx context.Contex
 
 func (c *gRPCClient) BidiReadObject(ctx context.Context, opts ...gax.CallOption) (storagepb.Storage_BidiReadObjectClient, error) {
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/BidiReadObject")
-	}
+
 	var resp storagepb.Storage_BidiReadObjectClient
 	opts = append((*c.CallOptions).BidiReadObject[0:len((*c.CallOptions).BidiReadObject):len((*c.CallOptions).BidiReadObject)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1679,9 +1514,7 @@ func (c *gRPCClient) UpdateObject(ctx context.Context,
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/UpdateObject")
-	}
+
 	opts = append((*c.CallOptions).UpdateObject[0:len((*c.CallOptions).UpdateObject):len((*c.CallOptions).UpdateObject)], opts...)
 	var resp *storagepb.Object
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1697,9 +1530,7 @@ func (c *gRPCClient) WriteObject(ctx context.Context, 
 
 func (c *gRPCClient) WriteObject(ctx context.Context, opts ...gax.CallOption) (storagepb.Storage_WriteObjectClient, error) {
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/WriteObject")
-	}
+
 	var resp storagepb.Storage_WriteObjectClient
 	opts = append((*c.CallOptions).WriteObject[0:len((*c.CallOptions).WriteObject):len((*c.CallOptions).WriteObject)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1717,9 +1548,7 @@ func (c *gRPCClient) BidiWriteObject(ctx context.Conte
 
 func (c *gRPCClient) BidiWriteObject(ctx context.Context, opts ...gax.CallOption) (storagepb.Storage_BidiWriteObjectClient, error) {
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/BidiWriteObject")
-	}
+
 	var resp storagepb.Storage_BidiWriteObjectClient
 	opts = append((*c.CallOptions).BidiWriteObject[0:len((*c.CallOptions).BidiWriteObject):len((*c.CallOptions).BidiWriteObject)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1749,9 +1578,6 @@ func (c *gRPCClient) ListObjects(ctx context.Context, 
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/ListObjects")
-	}
 	opts = append((*c.CallOptions).ListObjects[0:len((*c.CallOptions).ListObjects):len((*c.CallOptions).ListObjects)], opts...)
 	it := &ObjectIterator{}
 	req = proto.CloneOf(req)
@@ -1810,9 +1636,6 @@ func (c *gRPCClient) RewriteObject(ctx context.Context
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/RewriteObject")
-	}
 	opts = append((*c.CallOptions).RewriteObject[0:len((*c.CallOptions).RewriteObject):len((*c.CallOptions).RewriteObject)], opts...)
 	var resp *storagepb.RewriteResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1840,9 +1663,7 @@ func (c *gRPCClient) StartResumableWrite(ctx context.C
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/StartResumableWrite")
-	}
+
 	opts = append((*c.CallOptions).StartResumableWrite[0:len((*c.CallOptions).StartResumableWrite):len((*c.CallOptions).StartResumableWrite)], opts...)
 	var resp *storagepb.StartResumableWriteResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1870,9 +1691,7 @@ func (c *gRPCClient) QueryWriteStatus(ctx context.Cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/QueryWriteStatus")
-	}
+
 	opts = append((*c.CallOptions).QueryWriteStatus[0:len((*c.CallOptions).QueryWriteStatus):len((*c.CallOptions).QueryWriteStatus)], opts...)
 	var resp *storagepb.QueryWriteStatusResponse
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -1900,9 +1719,6 @@ func (c *gRPCClient) MoveObject(ctx context.Context, r
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.storage.v2.Storage/MoveObject")
-	}
 	opts = append((*c.CallOptions).MoveObject[0:len((*c.CallOptions).MoveObject):len((*c.CallOptions).MoveObject)], opts...)
 	var resp *storagepb.Object
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
