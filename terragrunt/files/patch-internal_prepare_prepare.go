--- internal/prepare/prepare.go.orig	1979-11-29 21:00:00 UTC
+++ internal/prepare/prepare.go
@@ -23,7 +23,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/externalcmd"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -159,22 +158,16 @@ func PrepareSource(
 
 	// Always download/copy source to cache directory for consistency.
 	// When no source is specified, sourceURL will be "." (current directory).
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "download_terraform_source", map[string]any{
-			"sourceUrl": redact.NewURL(sourceURL),
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
-
-			return err
-		})
+	_ = redact.NewURL(sourceURL)
+	updatedRunOpts, err = run.DownloadTerraformSource(
+		ctx,
+		l,
+		v,
+		sourceURL,
+		runOpts,
+		runCfg,
+		r,
+	)
 	if err != nil {
 		return nil, err
 	}
