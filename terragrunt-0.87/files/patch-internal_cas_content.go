--- internal/cas/content.go.orig	2025-09-25 14:56:12 UTC
+++ internal/cas/content.go
@@ -9,7 +9,6 @@ import (
 	"runtime"
 
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Content manages git object storage and linking
@@ -37,10 +36,6 @@ func (c *Content) Link(ctx context.Context, hash, targ
 
 // Link creates a hard link from the store to the target path
 func (c *Content) Link(ctx context.Context, hash, targetPath string) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "cas_link", map[string]any{
-		"hash": hash,
-		"path": targetPath,
-	}, func(childCtx context.Context) error {
 		sourcePath := c.getPath(hash)
 
 		// Try to create hard link directly (most efficient path)
@@ -82,7 +77,6 @@ func (c *Content) Link(ctx context.Context, hash, targ
 		}
 
 		return nil
-	})
 }
 
 // Store stores a single content item. This is typically used for trees,
