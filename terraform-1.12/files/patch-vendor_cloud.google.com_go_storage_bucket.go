--- vendor/cloud.google.com/go/storage/bucket.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/bucket.go
@@ -26,7 +26,6 @@ import (
 
 	"cloud.google.com/go/compute/metadata"
 	"cloud.google.com/go/internal/optional"
-	"cloud.google.com/go/internal/trace"
 	storagepb "cloud.google.com/go/storage/internal/apiv2/stubs"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/api/iamcredentials/v1"
@@ -81,9 +80,6 @@ func (b *BucketHandle) Create(ctx context.Context, pro
 // Create creates the Bucket in the project.
 // If attrs is nil the API defaults will be used.
 func (b *BucketHandle) Create(ctx context.Context, projectID string, attrs *BucketAttrs) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Bucket.Create")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 	if _, err := b.c.tc.CreateBucket(ctx, projectID, b.name, attrs, o...); err != nil {
 		return err
@@ -93,9 +89,6 @@ func (b *BucketHandle) Delete(ctx context.Context) (er
 
 // Delete deletes the Bucket.
 func (b *BucketHandle) Delete(ctx context.Context) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Bucket.Delete")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 	return b.c.tc.DeleteBucket(ctx, b.name, b.conds, o...)
 }
@@ -143,18 +136,12 @@ func (b *BucketHandle) Attrs(ctx context.Context) (att
 
 // Attrs returns the metadata for the bucket.
 func (b *BucketHandle) Attrs(ctx context.Context) (attrs *BucketAttrs, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Bucket.Attrs")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 	return b.c.tc.GetBucket(ctx, b.name, b.conds, o...)
 }
 
 // Update updates a bucket's attributes.
 func (b *BucketHandle) Update(ctx context.Context, uattrs BucketAttrsToUpdate) (attrs *BucketAttrs, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Bucket.Create")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	isIdempotent := b.conds != nil && b.conds.MetagenerationMatch != 0
 	o := makeStorageOpts(isIdempotent, b.retry, b.userProject)
 	return b.c.tc.UpdateBucket(ctx, b.name, &uattrs, b.conds, o...)
