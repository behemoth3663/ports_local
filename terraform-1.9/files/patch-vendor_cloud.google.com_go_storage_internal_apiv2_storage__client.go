--- vendor/cloud.google.com/go/storage/internal/apiv2/storage_client.go.orig	2026-06-04 04:52:32 UTC
+++ vendor/cloud.google.com/go/storage/internal/apiv2/storage_client.go
@@ -1017,43 +1017,6 @@ func NewClient(ctx context.Context, opts ...option.Cli
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
-
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
 
 	client.internalClient = c
 
