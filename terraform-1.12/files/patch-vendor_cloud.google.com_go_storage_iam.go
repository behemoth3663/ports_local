--- vendor/cloud.google.com/go/storage/iam.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/iam.go
@@ -19,7 +19,6 @@ import (
 
 	"cloud.google.com/go/iam"
 	"cloud.google.com/go/iam/apiv1/iampb"
-	"cloud.google.com/go/internal/trace"
 	raw "google.golang.org/api/storage/v1"
 	"google.golang.org/genproto/googleapis/type/expr"
 )
@@ -45,26 +44,17 @@ func (c *iamClient) GetWithVersion(ctx context.Context
 }
 
 func (c *iamClient) GetWithVersion(ctx context.Context, resource string, requestedPolicyVersion int32) (p *iampb.Policy, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.IAM.Get")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, c.retry, c.userProject)
 	return c.client.tc.GetIamPolicy(ctx, resource, requestedPolicyVersion, o...)
 }
 
 func (c *iamClient) Set(ctx context.Context, resource string, p *iampb.Policy) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.IAM.Set")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	isIdempotent := len(p.Etag) > 0
 	o := makeStorageOpts(isIdempotent, c.retry, c.userProject)
 	return c.client.tc.SetIamPolicy(ctx, resource, p, o...)
 }
 
 func (c *iamClient) Test(ctx context.Context, resource string, perms []string) (permissions []string, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.IAM.Test")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	o := makeStorageOpts(true, c.retry, c.userProject)
 	return c.client.tc.TestIamPermissions(ctx, resource, perms, o...)
 }
