--- vendor/cloud.google.com/go/storage/acl.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/acl.go
@@ -19,7 +19,6 @@ import (
 	"net/http"
 	"reflect"
 
-	"cloud.google.com/go/internal/trace"
 	storagepb "cloud.google.com/go/storage/internal/apiv2/stubs"
 	raw "google.golang.org/api/storage/v1"
 )
@@ -79,9 +78,6 @@ func (a *ACLHandle) Delete(ctx context.Context, entity
 
 // Delete permanently deletes the ACL entry for the given entity.
 func (a *ACLHandle) Delete(ctx context.Context, entity ACLEntity) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.ACL.Delete")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if a.object != "" {
 		return a.objectDelete(ctx, entity)
 	}
@@ -93,9 +89,6 @@ func (a *ACLHandle) Set(ctx context.Context, entity AC
 
 // Set sets the role for the given entity.
 func (a *ACLHandle) Set(ctx context.Context, entity ACLEntity, role ACLRole) (err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.ACL.Set")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if a.object != "" {
 		return a.objectSet(ctx, entity, role, false)
 	}
@@ -107,9 +100,6 @@ func (a *ACLHandle) List(ctx context.Context) (rules [
 
 // List retrieves ACL entries.
 func (a *ACLHandle) List(ctx context.Context) (rules []ACLRule, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.ACL.List")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if a.object != "" {
 		return a.objectList(ctx)
 	}
