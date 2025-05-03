--- vendor/cloud.google.com/go/storage/notifications.go.orig	2025-02-20 19:18:58 UTC
+++ vendor/cloud.google.com/go/storage/notifications.go
@@ -120,9 +120,6 @@ func (b *BucketHandle) AddNotification(ctx context.Con
 // returned Notification's ID can be used to refer to it.
 // Note: gRPC is not supported.
 func (b *BucketHandle) AddNotification(ctx context.Context, n *Notification) (ret *Notification, err error) {
-	ctx, _ = startSpan(ctx, "Bucket.AddNotification")
-	defer func() { endSpan(ctx, err) }()
-
 	if n.ID != "" {
 		return nil, errors.New("storage: AddNotification: ID must not be set")
 	}
@@ -142,9 +139,6 @@ func (b *BucketHandle) Notifications(ctx context.Conte
 // indexed by notification ID.
 // Note: gRPC is not supported.
 func (b *BucketHandle) Notifications(ctx context.Context) (n map[string]*Notification, err error) {
-	ctx, _ = startSpan(ctx, "Bucket.Notifications")
-	defer func() { endSpan(ctx, err) }()
-
 	opts := makeStorageOpts(true, b.retry, b.userProject)
 	n, err = b.c.tc.ListNotifications(ctx, b.name, opts...)
 	return n, err
@@ -161,9 +155,6 @@ func (b *BucketHandle) DeleteNotification(ctx context.
 // DeleteNotification deletes the notification with the given ID.
 // Note: gRPC is not supported.
 func (b *BucketHandle) DeleteNotification(ctx context.Context, id string) (err error) {
-	ctx, _ = startSpan(ctx, "Bucket.DeleteNotification")
-	defer func() { endSpan(ctx, err) }()
-
 	opts := makeStorageOpts(true, b.retry, b.userProject)
 	return b.c.tc.DeleteNotification(ctx, b.name, id, opts...)
 }
