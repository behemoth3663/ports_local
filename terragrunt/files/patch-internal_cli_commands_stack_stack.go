--- internal/cli/commands/stack/stack.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/stack/stack.go
@@ -7,7 +7,6 @@ import (
 	"slices"
 
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/hashicorp/hcl/v2"
 	"github.com/hashicorp/hcl/v2/hclsyntax"
 	"github.com/zclconf/go-cty/cty"
@@ -46,13 +45,7 @@ func RunGenerate(
 
 	// Clean stack folders before calling `generate` when the `--source-update` flag is passed
 	if opts.SourceUpdate {
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_clean", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context, l log.Logger) error {
-			l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
-			return clean.CleanStacks(l, v.FS, opts)
-		})
+		err := clean.CleanStacks(l, v.FS, opts)
 		if err != nil {
 			return fmt.Errorf(
 				"failed to clean stack directories under %q: %w",
@@ -91,12 +84,7 @@ func RunGenerate(
 
 	gen := generate.NewGenerator()
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_generate", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return gen.GenerateStacks(ctx, l, v, opts, wts)
-	})
+	err := gen.GenerateStacks(ctx, l, v, opts, wts)
 	if err != nil {
 		return err
 	}
@@ -119,12 +107,7 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 func Run(ctx context.Context, l log.Logger, v *venv.Venv, opts *options.TerragruntOptions) error {
 	opts.StackAction = "run"
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_run", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return RunGenerate(ctx, l, v, opts)
-	})
+	err := RunGenerate(ctx, l, v, opts)
 	if err != nil {
 		return err
 	}
@@ -143,18 +126,7 @@ func RunOutput(
 	opts.StackAction = "output"
 	opts.TerraformCommand = "output" // required for discovery exclude action matching in StackOutput
 
-	var outputs cty.Value
-
-	// collect outputs
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_output", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		stackOutputs, err := output.StackOutput(ctx, l, v, opts)
-		outputs = stackOutputs
-
-		return err
-	})
+	outputs, err := output.StackOutput(ctx, l, v, opts)
 	if err != nil {
 		return err
 	}
@@ -274,14 +246,7 @@ func RunClean(ctx context.Context, l log.Logger, v *ve
 
 // RunClean recursively removes all stack directories under the specified WorkingDir.
 func RunClean(ctx context.Context, l log.Logger, v *venv.Venv, opts *options.TerragruntOptions) error {
-	telemeter := telemetry.TelemeterFromContext(ctx)
-
-	err := telemeter.Collect(ctx, l, "stack_clean", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return clean.CleanStacks(l, v.FS, opts)
-	})
+	err := clean.CleanStacks(l, v.FS, opts)
 	if err != nil {
 		return fmt.Errorf("failed to clean stack directories under %q: %w", opts.WorkingDir, err)
 	}
