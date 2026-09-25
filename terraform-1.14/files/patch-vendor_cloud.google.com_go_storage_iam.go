--- vendor/cloud.google.com/go/storage/iam.go.orig	2026-08-26 17:25:54 UTC
+++ vendor/cloud.google.com/go/storage/iam.go
@@ -46,26 +46,17 @@ func (c *iamClient) GetWithVersion(ctx context.Context
 }
 
 func (c *iamClient) GetWithVersion(ctx context.Context, resource string, requestedPolicyVersion int32) (p *iampb.Policy, err error) {
-	ctx, _ = startSpanWithBucket(ctx, c.client, c.bucket, "storage.IAM.Get")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, c.retry, c.userProject)
 	return c.client.tc.GetIamPolicy(ctx, resource, requestedPolicyVersion, o...)
 }
 
 func (c *iamClient) Set(ctx context.Context, resource string, p *iampb.Policy) (err error) {
-	ctx, _ = startSpanWithBucket(ctx, c.client, c.bucket, "storage.IAM.Set")
-	defer func() { endSpan(ctx, err) }()
-
 	isIdempotent := len(p.Etag) > 0
 	o := makeStorageOpts(isIdempotent, c.retry, c.userProject)
 	return c.client.tc.SetIamPolicy(ctx, resource, p, o...)
 }
 
 func (c *iamClient) Test(ctx context.Context, resource string, perms []string) (permissions []string, err error) {
-	ctx, _ = startSpanWithBucket(ctx, c.client, c.bucket, "storage.IAM.Test")
-	defer func() { endSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, c.retry, c.userProject)
 	return c.client.tc.TestIamPermissions(ctx, resource, perms, o...)
 }
