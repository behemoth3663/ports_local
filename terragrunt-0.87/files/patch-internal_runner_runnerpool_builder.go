--- internal/runner/runnerpool/builder.go.orig	2025-09-25 14:56:12 UTC
+++ internal/runner/runnerpool/builder.go
@@ -8,7 +8,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/common"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Build stack runner using discovery and queueing mechanisms.
@@ -48,36 +47,12 @@ func Build(ctx context.Context, l log.Logger, terragru
 		d = d.WithIgnoreExternalDependencies()
 	}
 
-	// Wrap discovery with telemetry
-	var discovered discovery.DiscoveredConfigs
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "runner_pool_discovery", map[string]any{
-		"working_dir":       terragruntOptions.WorkingDir,
-		"terraform_command": terragruntOptions.TerraformCommand,
-	}, func(childCtx context.Context) error {
-		var discoveryErr error
-
-		discovered, discoveryErr = d.Discover(childCtx, l, terragruntOptions)
-
-		return discoveryErr
-	})
+	discovered, err := d.Discover(ctx, l, terragruntOptions)
 	if err != nil {
 		return nil, err
 	}
 
-	// Wrap runner pool creation with telemetry
-	var runner common.StackRunner
-
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "runner_pool_creation", map[string]any{
-		"discovered_configs": len(discovered),
-		"terraform_command":  terragruntOptions.TerraformCommand,
-	}, func(childCtx context.Context) error {
-		var runnerErr error
-
-		runner, runnerErr = NewRunnerPoolStack(childCtx, l, terragruntOptions, discovered, opts...)
-
-		return runnerErr
-	})
+	runner, err := NewRunnerPoolStack(ctx, l, terragruntOptions, discovered, opts...)
 	if err != nil {
 		return nil, err
 	}
