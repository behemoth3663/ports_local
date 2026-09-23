--- internal/runner/runall/runall.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/runall/runall.go
@@ -3,6 +3,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"path/filepath"
 
@@ -15,13 +16,10 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/worktrees"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/iacargs"
 	"github.com/gruntwork-io/terragrunt/internal/os/stdout"
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/options"
@@ -147,18 +145,7 @@ func Run(
 
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
-
-					return clean.CleanStacks(l, v.FS, opts)
-				})
+			errClean := clean.CleanStacks(l, v.FS, opts)
 			if errClean != nil {
 				return fmt.Errorf(
 					"failed to clean stack directories under %q: %w",
@@ -170,13 +157,7 @@ func Run(
 
 		// Generate the stack configuration with telemetry tracking
 		gen := generate.NewGenerator()
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_generate", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context, l log.Logger) error {
-			return gen.GenerateStacks(ctx, l, v, opts, wts)
-		})
-
+		err = gen.GenerateStacks(ctx, l, v, opts, wts)
 		// Handle any errors during stack generation
 		if err != nil {
 			return fmt.Errorf("failed to generate stack file: %w", err)
@@ -253,36 +234,17 @@ func RunAllOnStack(
 		}
 	}
 
-	var runErr error
+		err := rnr.Run(ctx, l, v, opts, r)
+		if err != nil {
+			// At this stage, we can't handle the error any further, so we just log it and return nil.
+			// After this point, we'll need to report on what happened, and we want that to happen
+			// after the error summary.
+			l.Errorf("Run failed: %v", err)
 
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
-				l.Errorf("Run failed: %v", err)
+			return err
+		}
 
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
-	}
-
-	return runErr
+		return nil
 }
 
 // shouldSkipSummary determines if summary output should be skipped for programmatic interactions.
