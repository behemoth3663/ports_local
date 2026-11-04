--- internal/runner/runall/runall.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/runall/runall.go
@@ -23,7 +23,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/os/stdout"
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/options"
@@ -149,18 +148,12 @@ func Run(
 
 		// Clean stack folders before calling `generate` when the `--source-update` flag is passed
 		if opts.SourceUpdate {
-			errClean := telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "stack_clean", map[string]any{
-					"stack_config_path": opts.TerragruntStackConfigPath,
-					"working_dir":       opts.WorkingDir,
-				}, func(ctx context.Context, l log.Logger) error {
-					l.Debugf(
-						"Running stack clean for %s, as part of generate command",
-						opts.WorkingDir,
-					)
+			l.Debugf(
+				"Running stack clean for %s, as part of generate command",
+				opts.WorkingDir,
+			)
 
-					return clean.CleanStacks(l, v.FS, opts)
-				})
+			errClean := clean.CleanStacks(l, v.FS, opts)
 			if errClean != nil {
 				return fmt.Errorf(
 					"failed to clean stack directories under %q: %w",
@@ -170,14 +163,8 @@ func Run(
 			}
 		}
 
-		// Generate the stack configuration with telemetry tracking
 		gen := generate.NewGenerator()
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_generate", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context, l log.Logger) error {
-			return gen.GenerateStacks(ctx, l, v, opts, wts)
-		})
+		err = gen.GenerateStacks(ctx, l, v, opts, wts)
 
 		// Handle any errors during stack generation
 		if err != nil {
@@ -263,33 +250,9 @@ func RunAllOnStack(
 		}
 	}
 
-	var runErr error
-
-	telemetryErr := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "run_all_on_stack", map[string]any{
-			"terraform_command": opts.TerraformCommand,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context, l log.Logger) error {
-			err := rnr.Run(ctx, l, v, opts, r)
-			if err != nil {
-				// At this stage, we can't handle the error any further, so we just log it and return nil.
-				// After this point, we'll need to report on what happened, and we want that to happen
-				// after the error summary.
-				l.Errorf("Run failed: %s", errfmt.Format(err))
-
-				// Save error to potentially return after telemetry completes
-				runErr = err
-
-				// Return nil to allow telemetry and reporting to complete
-				return nil
-			}
-
-			return nil
-		})
-
-	// log telemetry error and continue execution
-	if telemetryErr != nil {
-		l.Warnf("Telemetry collection failed: %v", telemetryErr)
+	runErr := rnr.Run(ctx, l, v, opts, r)
+	if runErr != nil {
+		l.Errorf("Run failed: %s", errfmt.Format(runErr))
 	}
 
 	return runErr
