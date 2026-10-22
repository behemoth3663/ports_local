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
 	"github.com/gruntwork-io/terragrunt/internal/iacargs"
@@ -28,7 +27,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -171,13 +169,8 @@ func Run(
 
 	// Always download/copy source to cache directory for consistency.
 	// When no source is specified, sourceURL will be "." (current directory).
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "download_terraform_source", map[string]any{
-			"sourceUrl": redact.NewURL(sourceURL),
-		}, func(ctx context.Context, l log.Logger) error {
-			updatedOpts, err = DownloadTerraformSource(ctx, l, v, sourceURL, opts, cfg, r)
-			return err
-		})
+	_ = redact.NewURL(sourceURL)
+	updatedOpts, err = DownloadTerraformSource(ctx, l, v, sourceURL, opts, cfg, r)
 	if err != nil {
 		return err
 	}
@@ -984,7 +977,7 @@ func setTerragruntNullValuesRunCfg(v *venv.Venv, opts 
 
 	varFile := filepath.Join(opts.CacheDir, NullTFVarsFile)
 
-	const ownerReadWritePermissions = 0600
+	const ownerReadWritePermissions = 0o600
 
 	if err := vfs.WriteFile(
 		v.FS,
