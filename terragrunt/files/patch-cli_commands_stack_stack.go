--- cli/commands/stack/stack.go.orig	2025-11-05 17:03:04 UTC
+++ cli/commands/stack/stack.go
@@ -5,7 +5,6 @@ import (
 	"path/filepath"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/zclconf/go-cty/cty"
 
 	"github.com/gruntwork-io/terragrunt/cli/commands/common/runall"
@@ -29,36 +28,27 @@ func RunGenerate(ctx context.Context, l log.Logger, op
 
 	// Clean stack folders before calling `generate` when the `--source-update` flag is passed
 	if opts.SourceUpdate {
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_clean", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context) error {
+		err := func(ctx context.Context) error {
 			l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
 			return config.CleanStacks(ctx, l, opts)
-		})
+		}(ctx)
 		if err != nil {
 			return errors.Errorf("failed to clean stack directories under %q: %w", opts.WorkingDir, err)
 		}
 	}
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_generate", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context) error {
+	return func(ctx context.Context) error {
 		return config.GenerateStacks(ctx, l, opts)
-	})
+	}(ctx)
 }
 
 // Run execute stack command.
 func Run(ctx context.Context, l log.Logger, opts *options.TerragruntOptions) error {
 	opts.StackAction = "run"
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_run", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context) error {
+	err := func(ctx context.Context) error {
 		return RunGenerate(ctx, l, opts)
-	})
+	}(ctx)
 	if err != nil {
 		return err
 	}
@@ -73,15 +63,12 @@ func RunOutput(ctx context.Context, l log.Logger, opts
 	var outputs cty.Value
 
 	// collect outputs
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_output", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context) error {
+	err := func(ctx context.Context) error {
 		stackOutputs, err := config.StackOutput(ctx, l, opts)
 		outputs = stackOutputs
 
 		return err
-	})
+	}(ctx)
 	if err != nil {
 		return errors.New(err)
 	}
@@ -152,17 +139,5 @@ func RunClean(ctx context.Context, l log.Logger, opts 
 
 // RunClean recursively removes all stack directories under the specified WorkingDir.
 func RunClean(ctx context.Context, l log.Logger, opts *options.TerragruntOptions) error {
-	telemeter := telemetry.TelemeterFromContext(ctx)
-
-	err := telemeter.Collect(ctx, "stack_clean", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context) error {
 		return config.CleanStacks(ctx, l, opts)
-	})
-	if err != nil {
-		return errors.Errorf("failed to clean stack directories under %q: %w", opts.WorkingDir, err)
-	}
-
-	return nil
 }
