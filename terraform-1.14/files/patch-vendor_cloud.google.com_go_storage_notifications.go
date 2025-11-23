--- vendor/cloud.google.com/go/storage/notifications.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/notifications.go
@@ -20,7 +20,6 @@ import (
 	"fmt"
 	"regexp"
 
-	"cloud.google.com/go/internal/trace"
 	storagepb "cloud.google.com/go/storage/internal/apiv2/stubs"
 	raw "google.golang.org/api/storage/v1"
 )
@@ -145,9 +144,6 @@ func (b *BucketHandle) AddNotification(ctx context.Con
 // and PayloadFormat, and must not set its ID. The other fields are all optional. The
 // returned Notification's ID can be used to refer to it.
 func (b *BucketHandle) AddNotification(ctx context.Context, n *Notification) (ret *Notification, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Bucket.AddNotification")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if n.ID != "" {
 		return nil, errors.New("storage: AddNotification: ID must not be set")
 	}
@@ -166,9 +162,6 @@ func (b *BucketHandle) Notifications(ctx context.Conte
 // Notifications returns all the Notifications configured for this bucket, as a map
 // indexed by notification ID.
 func (b *BucketHandle) Notifications(ctx context.Context) (n map[string]*Notification, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Bucket.Notifications")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	opts := makeStorageOpts(true, b.retry, b.userProject)
 	n, err = b.c.tc.ListNotifications(ctx, b.name, opts...)
 	return n, err
@@ -192,9 +185,6 @@ func (b *BucketHandle) DeleteNotification(ctx context.
 
 // DeleteNotification deletes the notification with the given ID.
 func (b *BucketHandle) DeleteNotification(ctx context.Context, id string) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Bucket.DeleteNotification")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	opts := makeStorageOpts(true, b.retry, b.userProject)
 	return b.c.tc.DeleteNotification(ctx, b.name, id, opts...)
 }
