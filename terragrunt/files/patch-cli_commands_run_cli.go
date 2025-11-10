--- cli/commands/run/cli.go.orig	2025-11-10 20:41:38 UTC
+++ cli/commands/run/cli.go
@@ -17,7 +17,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/stacks/generate"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tf"
 )
 
@@ -130,25 +129,19 @@ func wrapWithStackGenerate(l log.Logger, opts *options
 
 		// Clean stack folders before calling `generate` when the `--source-update` flag is passed
 		if opts.SourceUpdate {
-			errClean := telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_clean", map[string]any{
-				"stack_config_path": opts.TerragruntStackConfigPath,
-				"working_dir":       opts.WorkingDir,
-			}, func(ctx context.Context) error {
+			errClean := func(ctx context.Context) error {
 				l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
 				return config.CleanStacks(ctx, l, opts)
-			})
+			}(ctx)
 			if errClean != nil {
 				return errors.Errorf("failed to clean stack directories under %q: %w", opts.WorkingDir, errClean)
 			}
 		}
 
 		// Generate the stack configuration with telemetry tracking
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_generate", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context) error {
+		err := func(ctx context.Context) error {
 			return generate.GenerateStacks(ctx, l, opts)
-		})
+		}(ctx)
 
 		// Handle any errors during stack generation
 		if err != nil {
