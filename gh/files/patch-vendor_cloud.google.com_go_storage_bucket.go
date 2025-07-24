--- vendor/cloud.google.com/go/storage/bucket.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/bucket.go
@@ -81,9 +81,6 @@ func (b *BucketHandle) Create(ctx context.Context, pro
 // Create creates the Bucket in the project.
 // If attrs is nil the API defaults will be used.
 func (b *BucketHandle) Create(ctx context.Context, projectID string, attrs *BucketAttrs) (err error) {
-	ctx, _ = startSpan(ctx, "Bucket.Create")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 
 	if _, err := b.c.tc.CreateBucket(ctx, projectID, b.name, attrs, b.enableObjectRetention, o...); err != nil {
@@ -94,9 +91,6 @@ func (b *BucketHandle) Delete(ctx context.Context) (er
 
 // Delete deletes the Bucket.
 func (b *BucketHandle) Delete(ctx context.Context) (err error) {
-	ctx, _ = startSpan(ctx, "Bucket.Delete")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 	return b.c.tc.DeleteBucket(ctx, b.name, b.conds, o...)
 }
@@ -149,18 +143,12 @@ func (b *BucketHandle) Attrs(ctx context.Context) (att
 
 // Attrs returns the metadata for the bucket.
 func (b *BucketHandle) Attrs(ctx context.Context) (attrs *BucketAttrs, err error) {
-	ctx, _ = startSpan(ctx, "Bucket.Attrs")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 	return b.c.tc.GetBucket(ctx, b.name, b.conds, o...)
 }
 
 // Update updates a bucket's attributes.
 func (b *BucketHandle) Update(ctx context.Context, uattrs BucketAttrsToUpdate) (attrs *BucketAttrs, err error) {
-	ctx, _ = startSpan(ctx, "Bucket.Update")
-	defer func() { endSpan(ctx, err) }()
-
 	isIdempotent := b.conds != nil && b.conds.MetagenerationMatch != 0
 	o := makeStorageOpts(isIdempotent, b.retry, b.userProject)
 	return b.c.tc.UpdateBucket(ctx, b.name, &uattrs, b.conds, o...)
