--- internal/runner/run/run.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/run.go
@@ -28,7 +28,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -171,13 +170,8 @@ func Run(
 
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
