--- internal/cas/content.go.orig	2025-11-05 17:03:04 UTC
+++ internal/cas/content.go
@@ -9,7 +9,6 @@ import (
 	"runtime"
 
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Content manages git object storage and linking
@@ -37,10 +36,7 @@ func (c *Content) Link(ctx context.Context, hash, targ
 
 // Link creates a hard link from the store to the target path
 func (c *Content) Link(ctx context.Context, hash, targetPath string) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "cas_link", map[string]any{
-		"hash": hash,
-		"path": targetPath,
-	}, func(childCtx context.Context) error {
+	return func(childCtx context.Context) error {
 		sourcePath := c.getPath(hash)
 
 		// Try to create hard link directly (most efficient path)
@@ -82,7 +78,7 @@ func (c *Content) Link(ctx context.Context, hash, targ
 		}
 
 		return nil
-	})
+	}(ctx)
 }
 
 // Store stores a single content item. This is typically used for trees,
