--- internal/cli/commands/stack/stack.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/stack/stack.go
@@ -8,7 +8,6 @@ import (
 
 	rewrite "github.com/gruntwork-io/terragrunt/internal/config/rewrite"
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/hashicorp/hcl/v2"
 	"github.com/hashicorp/hcl/v2/hclsyntax"
 	"github.com/zclconf/go-cty/cty"
@@ -48,13 +47,8 @@ func RunGenerate(
 
 	// Clean stack folders before calling `generate` when the `--source-update` flag is passed
 	if opts.SourceUpdate {
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_clean", map[string]any{
-			"stack_config_path": opts.TerragruntStackConfigPath,
-			"working_dir":       opts.WorkingDir,
-		}, func(ctx context.Context, l log.Logger) error {
-			l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
-			return clean.CleanStacks(l, v.FS, opts)
-		})
+		l.Debugf("Running stack clean for %s, as part of generate command", opts.WorkingDir)
+		err := clean.CleanStacks(l, v.FS, opts)
 		if err != nil {
 			return fmt.Errorf(
 				"failed to clean stack directories under %q: %w",
@@ -93,12 +87,7 @@ func RunGenerate(
 
 	gen := generate.NewGenerator()
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_generate", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return gen.GenerateStacks(ctx, l, v, s, opts, wts)
-	})
+	err := gen.GenerateStacks(ctx, l, v, s, opts, wts)
 	if err != nil {
 		return err
 	}
@@ -122,12 +111,7 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 func Run(ctx context.Context, l log.Logger, v *venv.Venv, s *rewrite.Store, opts *options.TerragruntOptions) error {
 	opts.StackAction = "run"
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_run", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		return RunGenerate(ctx, l, v, s, opts)
-	})
+	err := RunGenerate(ctx, l, v, s, opts)
 	if err != nil {
 		return err
 	}
@@ -150,15 +134,8 @@ func RunOutput(
 	var outputs cty.Value
 
 	// collect outputs
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "stack_output", map[string]any{
-		"stack_config_path": opts.TerragruntStackConfigPath,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		stackOutputs, err := output.StackOutput(ctx, l, v, s, opts)
-		outputs = stackOutputs
-
-		return err
-	})
+	stackOutputs, err := output.StackOutput(ctx, l, v, s, opts)
+	outputs = stackOutputs
 	if err != nil {
 		return err
 	}
@@ -283,14 +260,7 @@ func RunClean(
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
+	err := clean.CleanStacks(l, v.FS, opts)
 	if err != nil {
 		return fmt.Errorf("failed to clean stack directories under %q: %w", opts.WorkingDir, err)
 	}
