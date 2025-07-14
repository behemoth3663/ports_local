--- vendor/cloud.google.com/go/storage/http_client.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/http_client.go
@@ -31,7 +31,6 @@ import (
 
 	"cloud.google.com/go/iam/apiv1/iampb"
 	"cloud.google.com/go/internal/optional"
-	"cloud.google.com/go/internal/trace"
 	"golang.org/x/oauth2/google"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/api/iterator"
@@ -777,9 +776,6 @@ func (c *httpStorageClient) NewRangeReader(ctx context
 }
 
 func (c *httpStorageClient) NewRangeReader(ctx context.Context, params *newRangeReaderParams, opts ...storageOption) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.httpStorageClient.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 
 	if c.config.useJSONforReads {
@@ -1139,9 +1135,6 @@ func (c *httpStorageClient) ListNotifications(ctx cont
 // Note: This API does not support pagination. However, entity limits cap the number of notifications on a single bucket,
 // so all results will be returned in the first response. See https://cloud.google.com/storage/quotas#buckets.
 func (c *httpStorageClient) ListNotifications(ctx context.Context, bucket string, opts ...storageOption) (n map[string]*Notification, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.httpStorageClient.ListNotifications")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 	call := c.raw.Notifications.List(bucket)
 	if s.userProject != "" {
@@ -1159,9 +1152,6 @@ func (c *httpStorageClient) CreateNotification(ctx con
 }
 
 func (c *httpStorageClient) CreateNotification(ctx context.Context, bucket string, n *Notification, opts ...storageOption) (ret *Notification, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.httpStorageClient.CreateNotification")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 	call := c.raw.Notifications.Insert(bucket, toRawNotification(n))
 	if s.userProject != "" {
@@ -1179,9 +1169,6 @@ func (c *httpStorageClient) DeleteNotification(ctx con
 }
 
 func (c *httpStorageClient) DeleteNotification(ctx context.Context, bucket string, id string, opts ...storageOption) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.httpStorageClient.DeleteNotification")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	s := callSettings(c.settings, opts...)
 	call := c.raw.Notifications.Delete(bucket, id)
 	if s.userProject != "" {
