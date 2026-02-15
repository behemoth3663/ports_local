--- internal/runner/runnerpool/runner.go.orig	2025-09-25 14:56:12 UTC
+++ internal/runner/runnerpool/runner.go
@@ -26,7 +26,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Runner implements the Stack interface for runner pool execution.
@@ -162,15 +161,8 @@ func (r *Runner) Run(ctx context.Context, l log.Logger
 	}
 
 	task := func(ctx context.Context, u *common.Unit) error {
-		return telemetry.TelemeterFromContext(ctx).Collect(ctx, "runner_pool_task", map[string]any{
-			"terraform_command":      u.TerragruntOptions.TerraformCommand,
-			"terraform_cli_args":     u.TerragruntOptions.TerraformCliArgs,
-			"working_dir":            u.TerragruntOptions.WorkingDir,
-			"terragrunt_config_path": u.TerragruntOptions.TerragruntConfigPath,
-		}, func(childCtx context.Context) error {
 			unitRunner := common.NewUnitRunner(u)
-			return unitRunner.Run(childCtx, u.TerragruntOptions, r.Stack.Report)
-		})
+			return unitRunner.Run(ctx, u.TerragruntOptions, r.Stack.Report)
 	}
 
 	r.queue.FailFast = opts.FailFast
