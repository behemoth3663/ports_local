--- internal/runner/runnerpool/builder.go.orig	2025-11-05 17:03:04 UTC
+++ internal/runner/runnerpool/builder.go
@@ -11,7 +11,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/common"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Build stack runner using discovery and queueing mechanisms.
@@ -95,16 +94,13 @@ func Build(
 	// Wrap discovery with telemetry
 	var discovered component.Components
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "runner_pool_discovery", map[string]any{
-		"working_dir":       terragruntOptions.WorkingDir,
-		"terraform_command": terragruntOptions.TerraformCommand,
-	}, func(childCtx context.Context) error {
+	err := func(childCtx context.Context) error {
 		var discoveryErr error
 
 		discovered, discoveryErr = d.Discover(childCtx, l, terragruntOptions)
 
 		return discoveryErr
-	})
+	}(ctx)
 	if err != nil {
 		return nil, err
 	}
@@ -112,16 +108,13 @@ func Build(
 	// Wrap runner pool creation with telemetry
 	var runner common.StackRunner
 
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "runner_pool_creation", map[string]any{
-		"discovered_configs": len(discovered),
-		"terraform_command":  terragruntOptions.TerraformCommand,
-	}, func(childCtx context.Context) error {
+	err = func(childCtx context.Context) error {
 		var runnerErr error
 
 		runner, runnerErr = NewRunnerPoolStack(childCtx, l, terragruntOptions, discovered, opts...)
 
 		return runnerErr
-	})
+	}(ctx)
 	if err != nil {
 		return nil, err
 	}
