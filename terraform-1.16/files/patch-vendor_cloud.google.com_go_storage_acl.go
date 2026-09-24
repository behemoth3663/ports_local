--- vendor/cloud.google.com/go/storage/acl.go.orig	2025-03-12 21:31:11 UTC
+++ vendor/cloud.google.com/go/storage/acl.go
@@ -76,8 +76,6 @@ func (a *ACLHandle) Delete(ctx context.Context, entity
 
 // Delete permanently deletes the ACL entry for the given entity.
 func (a *ACLHandle) Delete(ctx context.Context, entity ACLEntity) (err error) {
-	ctx, _ = startSpan(ctx, "ACL.Delete")
-	defer func() { endSpan(ctx, err) }()
 
 	if a.object != "" {
 		return a.objectDelete(ctx, entity)
@@ -90,8 +88,6 @@ func (a *ACLHandle) Set(ctx context.Context, entity AC
 
 // Set sets the role for the given entity.
 func (a *ACLHandle) Set(ctx context.Context, entity ACLEntity, role ACLRole) (err error) {
-	ctx, _ = startSpan(ctx, "ACL.Set")
-	defer func() { endSpan(ctx, err) }()
 
 	if a.object != "" {
 		return a.objectSet(ctx, entity, role, false)
@@ -104,8 +100,6 @@ func (a *ACLHandle) List(ctx context.Context) (rules [
 
 // List retrieves ACL entries.
 func (a *ACLHandle) List(ctx context.Context) (rules []ACLRule, err error) {
-	ctx, _ = startSpan(ctx, "ACL.List")
-	defer func() { endSpan(ctx, err) }()
 
 	if a.object != "" {
 		return a.objectList(ctx)
