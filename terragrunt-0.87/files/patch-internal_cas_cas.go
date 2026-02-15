--- internal/cas/cas.go.orig	2025-09-25 14:56:12 UTC
+++ internal/cas/cas.go
@@ -18,7 +18,6 @@ import (
 	"github.com/gofrs/flock"
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Options configures the behavior of CAS
@@ -103,11 +102,7 @@ func (c *CAS) Clone(ctx context.Context, l log.Logger,
 		}
 	}()
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "cas_clone", map[string]any{
-		"url":    url,
-		"branch": opts.Branch,
-	}, func(childCtx context.Context) error {
-		hash, err := c.resolveReference(childCtx, url, opts.Branch)
+		hash, err := c.resolveReference(ctx, url, opts.Branch)
 		if err != nil {
 			return err
 		}
@@ -130,7 +125,7 @@ func (c *CAS) Clone(ctx context.Context, l log.Logger,
 			// Set the working directory for git operations
 			c.git.SetWorkDir(tempDir)
 
-			if cloneAndStoreErr := c.cloneAndStoreContent(childCtx, l, opts, url, hash); cloneAndStoreErr != nil {
+			if cloneAndStoreErr := c.cloneAndStoreContent(ctx, l, opts, url, hash); cloneAndStoreErr != nil {
 				return cloneAndStoreErr
 			}
 		}
@@ -147,8 +142,7 @@ func (c *CAS) Clone(ctx context.Context, l log.Logger,
 			return err
 		}
 
-		return tree.LinkTree(childCtx, c.store, targetDir)
-	})
+		return tree.LinkTree(ctx, c.store, targetDir)
 }
 
 func (c *CAS) prepareTargetDirectory(dir, url string) string {
