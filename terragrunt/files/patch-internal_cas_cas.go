--- internal/cas/cas.go.orig	2025-11-05 17:03:04 UTC
+++ internal/cas/cas.go
@@ -18,7 +18,6 @@ import (
 	"github.com/gofrs/flock"
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Options configures the behavior of CAS
@@ -103,10 +102,7 @@ func (c *CAS) Clone(ctx context.Context, l log.Logger,
 		}
 	}()
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "cas_clone", map[string]any{
-		"url":    url,
-		"branch": opts.Branch,
-	}, func(childCtx context.Context) error {
+	return func(childCtx context.Context) error {
 		hash, err := c.resolveReference(childCtx, url, opts.Branch)
 		if err != nil {
 			return err
@@ -148,7 +144,7 @@ func (c *CAS) Clone(ctx context.Context, l log.Logger,
 		}
 
 		return tree.LinkTree(childCtx, c.store, targetDir)
-	})
+	}(ctx)
 }
 
 func (c *CAS) prepareTargetDirectory(dir, url string) string {
