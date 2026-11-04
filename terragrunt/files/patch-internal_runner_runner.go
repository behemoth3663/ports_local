--- internal/runner/runner.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/runner.go
@@ -27,15 +27,12 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/queue"
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/view/dag"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
 	"github.com/gruntwork-io/terragrunt/pkg/options"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Seen together in a unit's plan stderr, these mean the unit failed because a
@@ -442,17 +439,8 @@ func (rnr *Runner) Run(
 		}
 
 		unitPath := u.Path()
-		unitName := filepath.Base(unitPath)
 
-		return telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, unitLogger, "runner_pool_task", map[string]any{
-				"unit_path":              unitPath,
-				"unit_name":              unitName,
-				"terraform_command":      unitOpts.TerraformCommand,
-				"terraform_cli_args":     unitOpts.TerraformCliArgs,
-				"working_dir":            unitOpts.WorkingDir,
-				"terragrunt_config_path": unitOpts.TerragruntConfigPath,
-			}, func(childCtx context.Context, unitLogger log.Logger) error {
+		return func(childCtx context.Context, unitLogger log.Logger) error {
 				l.Debugf(
 					"Runner Pool Task: starting unit=%s command=%s",
 					unitPath,
@@ -502,25 +490,13 @@ func (rnr *Runner) Run(
 
 				var cfg *config.TerragruntConfig
 
-				err = telemetry.TelemeterFromContext(childCtx).
-					Collect(childCtx, unitLogger, "unit_read_config", map[string]any{
-						"unit_path":              unitPath,
-						"unit_name":              unitName,
-						"terragrunt_config_path": unitOpts.TerragruntConfigPath,
-					}, func(readCtx context.Context, unitLogger log.Logger) error {
-						pctx := configbridge.NewParsingContext(unitOpts)
-
-						var readErr error
-
-						cfg, readErr = config.ReadTerragruntConfig(
-							readCtx,
-							unitLogger,
-							unitV,
-							pctx,
-						)
-
-						return readErr
-					})
+				pctx := configbridge.NewParsingContext(unitOpts)
+				cfg, err = config.ReadTerragruntConfig(
+					childCtx,
+					unitLogger,
+					unitV,
+					pctx,
+				)
 				if err != nil {
 					logTaskOutcome(childCtx, l, unitPath, unitOpts.TerraformCommand, err)
 
@@ -538,25 +514,18 @@ func (rnr *Runner) Run(
 
 				runCfg := cfg.ToRunConfig(unitLogger, unitV.FS)
 
-				err = telemetry.TelemeterFromContext(childCtx).
-					Collect(childCtx, unitLogger, "unit_run", map[string]any{
-						"unit_path":         unitPath,
-						"unit_name":         unitName,
-						"terraform_command": unitOpts.TerraformCommand,
-					}, func(runCtx context.Context, unitLogger log.Logger) error {
-						unlockOutputs := config.LockDependencyOutputs(runCtx, unitOpts.TerragruntConfigPath)
-						defer unlockOutputs()
+				unlockOutputs := config.LockDependencyOutputs(childCtx, unitOpts.TerragruntConfigPath)
+				defer unlockOutputs()
 
-						return unitRunner.Run(
-							runCtx,
-							unitLogger,
-							unitV,
-							unitOpts,
-							r,
-							runCfg,
-							credsGetter,
-						)
-					})
+				err = unitRunner.Run(
+					childCtx,
+					unitLogger,
+					unitV,
+					unitOpts,
+					r,
+					runCfg,
+					credsGetter,
+				)
 
 				// This unit's terraform commands are all done, so release its engine now
 				// instead of holding it until the batch Shutdown. Skip units that another
@@ -584,7 +553,7 @@ func (rnr *Runner) Run(
 				logTaskOutcome(childCtx, l, unitPath, unitOpts.TerraformCommand, err)
 
 				return err
-			})
+			}(ctx, unitLogger)
 	}
 
 	rnr.queue.FailFast = stackOpts.FailFast
@@ -969,18 +938,12 @@ func collectDependenciesBounded(unit *component.Unit, 
 	}
 }
 
-// logTaskOutcome stamps the task outcome on the active span and emits a debug log,
-// so the runner_pool_task span carries succeeded/failed status without callers
-// needing to call SpanFromContext directly.
 func logTaskOutcome(ctx context.Context, l log.Logger, unitPath, command string, err error) {
 	outcome := "succeeded"
 	if err != nil {
 		outcome = "failed"
 	}
-
-	if span := trace.SpanFromContext(ctx); span.IsRecording() {
-		span.SetAttributes(attribute.String("outcome", outcome))
-	}
+	_ = ctx
 
 	if err != nil {
 		l.Debugf(
