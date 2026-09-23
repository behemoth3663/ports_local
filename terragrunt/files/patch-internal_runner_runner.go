--- internal/runner/runner.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/runner.go
@@ -25,15 +25,11 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/queue"
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
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
@@ -471,139 +467,110 @@ func (rnr *Runner) Run(
 		}
 
 		unitPath := u.Path()
-		unitName := filepath.Base(unitPath)
+			l.Debugf(
+				"Runner Pool Task: starting unit=%s command=%s",
+				unitPath,
+				unitOpts.TerraformCommand,
+			)
 
-		return telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, unitLogger, "runner_pool_task", map[string]any{
-				"unit_path":              unitPath,
-				"unit_name":              unitName,
-				"terraform_command":      unitOpts.TerraformCommand,
-				"terraform_cli_args":     unitOpts.TerraformCliArgs,
-				"working_dir":            unitOpts.WorkingDir,
-				"terragrunt_config_path": unitOpts.TerragruntConfigPath,
-			}, func(childCtx context.Context, unitLogger log.Logger) error {
-				l.Debugf(
-					"Runner Pool Task: starting unit=%s command=%s",
-					unitPath,
-					unitOpts.TerraformCommand,
-				)
+			// Wrap the writer to buffer unit-scoped output. Build a per-unit
+			// venv so the wrapped writers flow through tf and shell calls.
+			unitWriter := NewUnitWriter(v.Writers.Writer)
 
-				// Wrap the writer to buffer unit-scoped output. Build a per-unit
-				// venv so the wrapped writers flow through tf and shell calls.
-				unitWriter := NewUnitWriter(v.Writers.Writer)
+			// Per-unit mutations (e.g. SetTerragruntInputsAsEnvVars writing
+			// TF_VAR_* in run.go) must not leak across concurrent units.
+			unitV := v.WithEnvCloned().WithWriter(unitWriter)
 
-				// Per-unit mutations (e.g. SetTerragruntInputsAsEnvVars writing
-				// TF_VAR_* in run.go) must not leak across concurrent units.
-				unitV := v.WithEnvCloned().WithWriter(unitWriter)
+			if unitErrWriterWrap != nil {
+				unitV = unitV.WithErrWriter(unitErrWriterWrap)
+			}
 
-				if unitErrWriterWrap != nil {
-					unitV = unitV.WithErrWriter(unitErrWriterWrap)
-				}
+			unitRunner := NewUnitRunner(u)
 
-				unitRunner := NewUnitRunner(u)
+			// Get credentials BEFORE config parsing: sops_decrypt_file() and
+			// get_aws_account_id() in locals need auth-provider credentials
+			// available in v.Env during HCL evaluation.
+			// See https://github.com/gruntwork-io/terragrunt/issues/5515
+			//
+			// The obtain_creds span is emitted by externalcmd.Provider.GetCredentials
+			// only when an auth provider is configured, so no conditional is needed here.
+			credsGetter, err := creds.ObtainCredsForParsing(
+				ctx,
+				unitLogger,
+				unitV,
+				unitOpts.AuthProviderCmd,
+				configbridge.ShellRunOptsFromOpts(v.Env, unitOpts),
+			)
+			if err != nil {
+				logTaskOutcome(ctx, l, unitPath, unitOpts.TerraformCommand, err)
 
-				// Get credentials BEFORE config parsing: sops_decrypt_file() and
-				// get_aws_account_id() in locals need auth-provider credentials
-				// available in v.Env during HCL evaluation.
-				// See https://github.com/gruntwork-io/terragrunt/issues/5515
-				//
-				// The obtain_creds span is emitted by externalcmd.Provider.GetCredentials
-				// only when an auth provider is configured, so no conditional is needed here.
-				credsGetter, err := creds.ObtainCredsForParsing(
-					childCtx,
-					unitLogger,
-					unitV,
-					unitOpts.AuthProviderCmd,
-					configbridge.ShellRunOptsFromOpts(v.Env, unitOpts),
-				)
-				if err != nil {
-					logTaskOutcome(childCtx, l, unitPath, unitOpts.TerraformCommand, err)
+				return err
+			}
 
-					return err
-				}
+			var cfg *config.TerragruntConfig
 
-				var cfg *config.TerragruntConfig
+			parseCtx, pctx := configbridge.NewParsingContext(
+				ctx,
+				unitLogger,
+				unitV,
+				unitOpts,
+			)
 
-				err = telemetry.TelemeterFromContext(childCtx).
-					Collect(childCtx, unitLogger, "unit_read_config", map[string]any{
-						"unit_path":              unitPath,
-						"unit_name":              unitName,
-						"terragrunt_config_path": unitOpts.TerragruntConfigPath,
-					}, func(readCtx context.Context, unitLogger log.Logger) error {
-						parseCtx, pctx := configbridge.NewParsingContext(
-							readCtx,
-							unitLogger,
-							unitV,
-							unitOpts,
-						)
+			cfg, err = config.ReadTerragruntConfig(
+				parseCtx,
+				unitLogger,
+				pctx,
+				pctx.ParserOptions,
+			)
 
-						var readErr error
+			if err != nil {
+				logTaskOutcome(ctx, l, unitPath, unitOpts.TerraformCommand, err)
 
-						cfg, readErr = config.ReadTerragruntConfig(
-							parseCtx,
-							unitLogger,
-							pctx,
-							pctx.ParserOptions,
-						)
+				return err
+			}
 
-						return readErr
-					})
-				if err != nil {
-					logTaskOutcome(childCtx, l, unitPath, unitOpts.TerraformCommand, err)
+			if !unitOpts.TFPathExplicitlySet && cfg.TerraformBinary != "" {
+				unitOpts.TFPath = cfg.TerraformBinary
+			}
 
-					return err
-				}
+			runCfg := cfg.ToRunConfig(unitLogger, unitV.FS)
 
-				if !unitOpts.TFPathExplicitlySet && cfg.TerraformBinary != "" {
-					unitOpts.TFPath = cfg.TerraformBinary
-				}
+			err = unitRunner.Run(
+					ctx,
+					unitLogger,
+					unitV,
+					unitOpts,
+					r,
+					runCfg,
+					credsGetter,
+				)
 
-				runCfg := cfg.ToRunConfig(unitLogger, unitV.FS)
-
-				err = telemetry.TelemeterFromContext(childCtx).
-					Collect(childCtx, unitLogger, "unit_run", map[string]any{
-						"unit_path":         unitPath,
-						"unit_name":         unitName,
-						"terraform_command": unitOpts.TerraformCommand,
-					}, func(runCtx context.Context, unitLogger log.Logger) error {
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
-
-				// This unit's terraform commands are all done, so release its engine now
-				// instead of holding it until the batch Shutdown. Skip units that another
-				// in-run unit depends on: that dependent re-reads this unit's outputs through
-				// engine.Run after this task ends, which would just re-spawn the engine we
-				// tore down.
-				if !withDependents[unitPath] {
-					noEngine := unitOpts.EngineOptions != nil && unitOpts.EngineOptions.NoEngine
-					if sErr := engine.ShutdownUnit(
-						childCtx, unitLogger, unitOpts.Experiments, noEngine,
-						unitOpts.WorkingDir,
-					); sErr != nil {
-						unitLogger.Errorf(
-							"Error shutting down engine for unit %s: %v",
-							unitPath,
-							sErr,
-						)
-					}
+			// This unit's terraform commands are all done, so release its engine now
+			// instead of holding it until the batch Shutdown. Skip units that another
+			// in-run unit depends on: that dependent re-reads this unit's outputs through
+			// engine.Run after this task ends, which would just re-spawn the engine we
+			// tore down.
+			if !withDependents[unitPath] {
+				noEngine := unitOpts.EngineOptions != nil && unitOpts.EngineOptions.NoEngine
+				if sErr := engine.ShutdownUnit(
+					ctx, unitLogger, unitOpts.Experiments, noEngine,
+					unitOpts.WorkingDir,
+				); sErr != nil {
+					unitLogger.Errorf(
+						"Error shutting down engine for unit %s: %v",
+						unitPath,
+						sErr,
+					)
 				}
+			}
 
-				if flushErr := unitWriter.Flush(); flushErr != nil && err == nil {
-					err = flushErr
-				}
+			if flushErr := unitWriter.Flush(); flushErr != nil && err == nil {
+				err = flushErr
+			}
 
-				logTaskOutcome(childCtx, l, unitPath, unitOpts.TerraformCommand, err)
+			logTaskOutcome(ctx, l, unitPath, unitOpts.TerraformCommand, err)
 
-				return err
-			})
+			return err
 	}
 
 	rnr.queue.FailFast = stackOpts.FailFast
@@ -1112,10 +1079,6 @@ func logTaskOutcome(ctx context.Context, l log.Logger,
 	outcome := "succeeded"
 	if err != nil {
 		outcome = "failed"
-	}
-
-	if span := trace.SpanFromContext(ctx); span.IsRecording() {
-		span.SetAttributes(attribute.String("outcome", outcome))
 	}
 
 	if err != nil {
