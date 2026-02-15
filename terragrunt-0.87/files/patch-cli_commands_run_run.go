--- cli/commands/run/run.go.orig	2025-09-25 14:56:12 UTC
+++ cli/commands/run/run.go
@@ -19,7 +19,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/cli/commands/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/cli/commands/run/creds/providers/externalcmd"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 
 	"github.com/gruntwork-io/terragrunt/tf"
 
@@ -210,12 +209,7 @@ func run(ctx context.Context, l log.Logger, opts *opti
 	}
 
 	if sourceURL != "" {
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "download_terraform_source", map[string]any{
-			"sourceUrl": sourceURL,
-		}, func(ctx context.Context) error {
 			updatedTerragruntOptions, err = downloadTerraformSource(ctx, l, sourceURL, opts, terragruntConfig, r)
-			return err
-		})
 		if err != nil {
 			return target.runErrorCallback(l, opts, terragruntConfig, err)
 		}
