--- internal/runner/runner.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/runner.go
@@ -27,16 +27,12 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/queue"
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/internal/view/dag"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
 	"github.com/gruntwork-io/terragrunt/pkg/options"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Seen together in a unit's plan stderr, these mean the unit failed because a
@@ -444,17 +440,8 @@ func (rnr *Runner) Run(
 		}
 
 		unitPath := u.Path()
-		unitName := filepath.Base(unitPath)
 
-		return telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, unitLogger, "runner_pool_task", map[string]any{
-				"unit_path":              unitPath,
-				"unit_name":              unitName,
-				"tofu_command":           unitOpts.TofuCommand,
-				"tofu_cli_args":          unitOpts.TerraformCliArgs.Slice(),
-				"working_dir":            unitOpts.WorkingDir,
-				"terragrunt_config_path": unitOpts.TerragruntConfigPath,
-			}, func(childCtx context.Context, unitLogger log.Logger) error {
+		return func(childCtx context.Context, unitLogger log.Logger) error {
 				l.Debugf(
 					"Runner Pool Task: starting unit=%s command=%s",
 					unitPath,
@@ -504,28 +491,17 @@ func (rnr *Runner) Run(
 
 				var cfg *config.TerragruntConfig
 
-				err = telemetry.TelemeterFromContext(childCtx).
-					Collect(childCtx, unitLogger, "unit_read_config", map[string]any{
-						"unit_path":              unitPath,
-						"unit_name":              unitName,
-						"terragrunt_config_path": unitOpts.TerragruntConfigPath,
-					}, func(readCtx context.Context, unitLogger log.Logger) error {
 						pctx := configbridge.NewParsingContext(unitOpts)
-
-						var readErr error
-
-						cfg, readErr = s.ReadTerragruntConfig(
-							readCtx,
+						cfg, err = s.ReadTerragruntConfig(
+							childCtx,
 							unitLogger,
 							unitV,
 							pctx,
 						)
-						if readErr != nil {
-							return readErr
+						if err == nil {
+							err = tips.GiveOutputsBlockTips(childCtx, unitLogger, unitV, s, pctx, cfg, unitOpts.Tips)
 						}
 
-						return tips.GiveOutputsBlockTips(readCtx, unitLogger, unitV, s, pctx, cfg, unitOpts.Tips)
-					})
 				if err != nil {
 					logTaskOutcome(childCtx, l, unitPath, unitOpts.TofuCommand, err)
 
@@ -557,17 +533,10 @@ func (rnr *Runner) Run(
 
 				runCfg := cfg.ToRunConfig(unitLogger, unitV.FS)
 
-				err = telemetry.TelemeterFromContext(childCtx).
-					Collect(childCtx, unitLogger, "unit_run", map[string]any{
-						"unit_path":    unitPath,
-						"unit_name":    unitName,
-						"tofu_command": unitOpts.TofuCommand,
-					}, func(runCtx context.Context, unitLogger log.Logger) error {
-						unlockOutputs := config.LockDependencyOutputs(runCtx, unitOpts.TerragruntConfigPath)
-						defer unlockOutputs()
+						unlockOutputs := config.LockDependencyOutputs(childCtx, unitOpts.TerragruntConfigPath)
 
-						return unitRunner.Run(
-							runCtx,
+						err = unitRunner.Run(
+							childCtx,
 							unitLogger,
 							unitV,
 							unitOpts,
@@ -576,8 +545,9 @@ func (rnr *Runner) Run(
 							runCfg,
 							credsGetter,
 						)
-					})
 
+						unlockOutputs()
+
 				// This unit's terraform commands are all done, so release its engine now
 				// instead of holding it until the batch Shutdown. Skip units that another
 				// in-run unit depends on: that dependent re-reads this unit's outputs through
@@ -604,7 +574,7 @@ func (rnr *Runner) Run(
 				logTaskOutcome(childCtx, l, unitPath, unitOpts.TofuCommand, err)
 
 				return err
-			})
+			}(ctx, unitLogger)
 	}
 
 	rnr.queue.FailFast = stackOpts.FailFast
@@ -976,18 +946,12 @@ func reachableDependencies(unit *component.Unit) []str
 	return paths
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
