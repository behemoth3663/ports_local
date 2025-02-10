--- vendor/cloud.google.com/go/storage/grpc_client.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/grpc_client.go
@@ -24,7 +24,6 @@ import (
 	"os"
 
 	"cloud.google.com/go/iam/apiv1/iampb"
-	"cloud.google.com/go/internal/trace"
 	gapic "cloud.google.com/go/storage/internal/apiv2"
 	storagepb "cloud.google.com/go/storage/internal/apiv2/stubs"
 	"github.com/googleapis/gax-go/v2"
@@ -852,9 +851,6 @@ func (c *grpcStorageClient) NewRangeReader(ctx context
 }
 
 func (c *grpcStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	if s.userProject != "" {
@@ -1253,9 +1249,6 @@ func (c *grpcStorageClient) ListNotifications(ctx cont
 // Notification methods.
 
 func (c *grpcStorageClient) ListNotifications(ctx context.Context, bucket string, opts ...storageOption) (n map[string]*Notification, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.ListNotifications")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 	if s.userProject != "" {
 		ctx = setUserProjectMetadata(ctx, s.userProject)
@@ -1288,9 +1281,6 @@ func (c *grpcStorageClient) CreateNotification(ctx con
 }
 
 func (c *grpcStorageClient) CreateNotification(ctx context.Context, bucket string, n *Notification, opts ...storageOption) (ret *Notification, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.CreateNotification")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 	req := &storagepb.CreateNotificationConfigRequest{
 		Parent:             bucketResourceName(globalProjectAlias, bucket),
@@ -1309,9 +1299,6 @@ func (c *grpcStorageClient) DeleteNotification(ctx con
 }
 
 func (c *grpcStorageClient) DeleteNotification(ctx context.Context, bucket string, id string, opts ...storageOption) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.grpcStorageClient.DeleteNotification")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 	req := &storagepb.DeleteNotificationConfigRequest{Name: id}
 	return run(ctx, func() error {
