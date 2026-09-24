--- vendor/cloud.google.com/go/storage/iam.go.orig	2025-03-12 21:31:11 UTC
+++ vendor/cloud.google.com/go/storage/iam.go
@@ -44,16 +44,12 @@ func (c *iamClient) GetWithVersion(ctx context.Context
 }
 
 func (c *iamClient) GetWithVersion(ctx context.Context, resource string, requestedPolicyVersion int32) (p *iampb.Policy, err error) {
-	ctx, _ = startSpan(ctx, "storage.IAM.Get")
-	defer func() { endSpan(ctx, err) }()
 
 	o := makeStorageOpts(true, c.retry, c.userProject)
 	return c.client.tc.GetIamPolicy(ctx, resource, requestedPolicyVersion, o...)
 }
 
 func (c *iamClient) Set(ctx context.Context, resource string, p *iampb.Policy) (err error) {
-	ctx, _ = startSpan(ctx, "storage.IAM.Set")
-	defer func() { endSpan(ctx, err) }()
 
 	isIdempotent := len(p.Etag) > 0
 	o := makeStorageOpts(isIdempotent, c.retry, c.userProject)
@@ -61,8 +57,6 @@ func (c *iamClient) Test(ctx context.Context, resource
 }
 
 func (c *iamClient) Test(ctx context.Context, resource string, perms []string) (permissions []string, err error) {
-	ctx, _ = startSpan(ctx, "storage.IAM.Test")
-	defer func() { endSpan(ctx, err) }()
 
 	o := makeStorageOpts(true, c.retry, c.userProject)
 	return c.client.tc.TestIamPermissions(ctx, resource, perms, o...)
