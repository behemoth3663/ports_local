--- internal/runner/run/run.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/run.go
@@ -6,6 +6,7 @@ import (
 import (
 	"context"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"io"
 	"maps"
@@ -16,8 +17,6 @@ import (
 	"strings"
 	"sync"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/cas"
 	"github.com/gruntwork-io/terragrunt/internal/codegen"
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
@@ -29,7 +28,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -172,13 +170,7 @@ func Run(
 
 	// Always download/copy source to cache directory for consistency.
 	// When no source is specified, sourceURL will be "." (current directory).
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "download_terraform_source", map[string]any{
-			"sourceUrl": sourceURL,
-		}, func(ctx context.Context, l log.Logger) error {
-			updatedOpts, err = DownloadTerraformSource(ctx, l, v, sourceURL, opts, cfg, r)
-			return err
-		})
+		updatedOpts, err = DownloadTerraformSource(ctx, l, v, sourceURL, opts, cfg, r)
 	if err != nil {
 		return err
 	}
@@ -946,7 +938,7 @@ func checkProtectedModuleRunCfg(opts *Options, cfg *ru
 
 // checkProtectedModuleRunCfg checks if module is protected using runcfg types.
 func checkProtectedModuleRunCfg(opts *Options, cfg *runcfg.RunConfig) error {
-	var destroyFlag = false
+	destroyFlag := false
 	if opts.TerraformCliArgs.First() == tf.CommandNameDestroy {
 		destroyFlag = true
 	}
@@ -990,7 +982,7 @@ func setTerragruntNullValuesRunCfg(v *venv.Venv, opts 
 
 	varFile := filepath.Join(opts.CacheDir, NullTFVarsFile)
 
-	const ownerReadWritePermissions = 0600
+	const ownerReadWritePermissions = 0o600
 
 	if err := vfs.WriteFile(
 		v.FS,
