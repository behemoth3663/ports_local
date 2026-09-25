--- vendor/cloud.google.com/go/storage/bucket.go.orig	2026-08-26 17:25:54 UTC
+++ vendor/cloud.google.com/go/storage/bucket.go
@@ -82,9 +82,6 @@ func (b *BucketHandle) Create(ctx context.Context, pro
 // Create creates the Bucket in the project.
 // If attrs is nil the API defaults will be used.
 func (b *BucketHandle) Create(ctx context.Context, projectID string, attrs *BucketAttrs) (err error) {
-	ctx, _ = startSpanWithBucket(ctx, b.c, b.name, "Bucket.Create")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 
 	if _, err := b.c.tc.CreateBucket(ctx, projectID, b.name, attrs, b.enableObjectRetention, o...); err != nil {
@@ -95,9 +92,6 @@ func (b *BucketHandle) Delete(ctx context.Context) (er
 
 // Delete deletes the Bucket.
 func (b *BucketHandle) Delete(ctx context.Context) (err error) {
-	ctx, _ = startSpanWithBucket(ctx, b.c, b.name, "Bucket.Delete")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 	return b.c.tc.DeleteBucket(ctx, b.name, b.conds, o...)
 }
@@ -150,9 +144,6 @@ func (b *BucketHandle) Attrs(ctx context.Context) (att
 
 // Attrs returns the metadata for the bucket.
 func (b *BucketHandle) Attrs(ctx context.Context) (attrs *BucketAttrs, err error) {
-	ctx, _ = startSpanWithBucket(ctx, b.c, b.name, "Bucket.Attrs")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, b.retry, b.userProject)
 	attrs, err = b.c.tc.GetBucket(ctx, b.name, b.conds, o...)
 	if err == nil && b.c != nil && b.c.bucketMetadataCache != nil {
@@ -164,9 +155,6 @@ func (b *BucketHandle) Update(ctx context.Context, uat
 
 // Update updates a bucket's attributes.
 func (b *BucketHandle) Update(ctx context.Context, uattrs BucketAttrsToUpdate) (attrs *BucketAttrs, err error) {
-	ctx, _ = startSpanWithBucket(ctx, b.c, b.name, "Bucket.Update")
-	defer func() { endSpan(ctx, err) }()
-
 	isIdempotent := b.conds != nil && b.conds.MetagenerationMatch != 0
 	o := makeStorageOpts(isIdempotent, b.retry, b.userProject)
 	return b.c.tc.UpdateBucket(ctx, b.name, &uattrs, b.conds, o...)
