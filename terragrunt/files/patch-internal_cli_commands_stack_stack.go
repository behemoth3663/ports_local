--- internal/cli/commands/stack/stack.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/stack/stack.go
@@ -7,7 +7,6 @@ import (
 	"slices"
 
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/hashicorp/hcl/v2"
 	"github.com/hashicorp/hcl/v2/hclsyntax"
 	"github.com/zclconf/go-cty/cty"
@@ -46,13 +45,9 @@ func RunGenerate(
 
 	// Clean stack folders before calling `generate` when the `--source-update` flag is passed
 	if opts.SourceUpdate {
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_clean", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context, l log.Logger) error {
-			l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
-			return clean.CleanStacks(l, v.FS, opts)
-		})
+		_, _ = opts.TerragruntStackConfigPath, opts.WorkingDir
+		l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
+		err := clean.CleanStacks(l, v.FS, opts)
 		if err != nil {
 			return fmt.Errorf(
 				"failed to clean stack directories under %q: %w",
@@ -91,12 +86,8 @@ func RunGenerate(
 
 	gen := generate.NewGenerator()
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_generate", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return gen.GenerateStacks(ctx, l, v, opts, wts)
-	})
+	_, _ = opts.TerragruntStackConfigPath, opts.WorkingDir
+	err := gen.GenerateStacks(ctx, l, v, opts, wts)
 	if err != nil {
 		return err
 	}
@@ -120,12 +111,8 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 func Run(ctx context.Context, l log.Logger, v *venv.Venv, opts *options.TerragruntOptions) error {
 	opts.StackAction = "run"
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_run", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return RunGenerate(ctx, l, v, opts)
-	})
+	_, _ = opts.TerragruntStackConfigPath, opts.WorkingDir
+	err := RunGenerate(ctx, l, v, opts)
 	if err != nil {
 		return err
 	}
@@ -147,15 +134,9 @@ func RunOutput(
 	var outputs cty.Value
 
 	// collect outputs
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_output", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		stackOutputs, err := output.StackOutput(ctx, l, v, opts)
-		outputs = stackOutputs
-
-		return err
-	})
+	_, _ = opts.TerragruntStackConfigPath, opts.WorkingDir
+	stackOutputs, err := output.StackOutput(ctx, l, v, opts)
+	outputs = stackOutputs
 	if err != nil {
 		return err
 	}
@@ -280,14 +261,8 @@ func RunClean(
 	v *venv.Venv,
 	opts *options.TerragruntOptions,
 ) error {
-	telemeter := telemetry.TelemeterFromContext(ctx)
-
-	err := telemeter.Collect(ctx, l, "stack_clean", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return clean.CleanStacks(l, v.FS, opts)
-	})
+	_, _, _ = ctx, opts.TerragruntStackConfigPath, opts.WorkingDir
+	err := clean.CleanStacks(l, v.FS, opts)
 	if err != nil {
 		return fmt.Errorf("failed to clean stack directories under %q: %w", opts.WorkingDir, err)
 	}
