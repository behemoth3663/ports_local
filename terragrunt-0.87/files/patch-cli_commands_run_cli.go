--- cli/commands/run/cli.go.orig	2025-09-25 14:56:12 UTC
+++ cli/commands/run/cli.go
@@ -2,7 +2,6 @@ import (
 package run
 
 import (
-	"context"
 	"path/filepath"
 	"strings"
 
@@ -15,7 +14,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tf"
 )
 
@@ -128,25 +126,15 @@ func wrapWithStackGenerate(l log.Logger, opts *options
 
 		// Clean stack folders before calling `generate` when the `--source-update` flag is passed
 		if opts.SourceUpdate {
-			errClean := telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_clean", map[string]any{
-				"stack_config_path": opts.TerragruntStackConfigPath,
-				"working_dir":       opts.WorkingDir,
-			}, func(ctx context.Context) error {
 				l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
-				return config.CleanStacks(ctx, l, opts)
-			})
+			errClean := config.CleanStacks(ctx, l, opts)
 			if errClean != nil {
 				return errors.Errorf("failed to clean stack directories under %q: %w", opts.WorkingDir, errClean)
 			}
 		}
 
 		// Generate the stack configuration with telemetry tracking
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_generate", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context) error {
-			return config.GenerateStacks(ctx, l, opts)
-		})
+		err := config.GenerateStacks(ctx, l, opts)
 
 		// Handle any errors during stack generation
 		if err != nil {
