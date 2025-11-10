--- internal/runner/run/run.go.orig	2025-11-10 13:46:52 UTC
+++ internal/runner/run/run.go
@@ -20,7 +20,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/externalcmd"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 
 	"github.com/gruntwork-io/terragrunt/tf"
 
@@ -182,12 +181,10 @@ func run(ctx context.Context, l log.Logger, opts *opti
 	}
 
 	if sourceURL != "" {
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "download_terraform_source", map[string]any{
-			"sourceUrl": sourceURL,
-		}, func(ctx context.Context) error {
+		err = func(ctx context.Context) error {
 			updatedTerragruntOptions, err = downloadTerraformSource(ctx, l, sourceURL, opts, terragruntConfig, r)
 			return err
-		})
+		}(ctx)
 		if err != nil {
 			return target.runErrorCallback(l, opts, terragruntConfig, err)
 		}
