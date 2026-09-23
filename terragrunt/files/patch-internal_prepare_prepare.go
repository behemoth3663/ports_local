--- internal/prepare/prepare.go.orig	1979-11-29 21:00:00 UTC
+++ internal/prepare/prepare.go
@@ -22,7 +22,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/externalcmd"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -158,22 +157,16 @@ func PrepareSource(
 
 	// Always download/copy source to cache directory for consistency.
 	// When no source is specified, sourceURL will be "." (current directory).
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "download_terraform_source", map[string]any{
-			"sourceUrl": sourceURL,
-		}, func(ctx context.Context, l log.Logger) error {
-			updatedRunOpts, err = run.DownloadTerraformSource(
-				ctx,
-				l,
-				v,
-				sourceURL,
-				runOpts,
-				runCfg,
-				r,
-			)
+		updatedRunOpts, err = run.DownloadTerraformSource(
+			ctx,
+			l,
+			v,
+			sourceURL,
+			runOpts,
+			runCfg,
+			r,
+		)
 
-			return err
-		})
 	if err != nil {
 		return nil, err
 	}
