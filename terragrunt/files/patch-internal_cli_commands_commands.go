--- internal/cli/commands/commands.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/commands.go
@@ -15,7 +15,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/cache"
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
 	"github.com/gruntwork-io/terragrunt/internal/engine"
-	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/panicreport"
 	"github.com/gruntwork-io/terragrunt/internal/providercache"
@@ -52,7 +51,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/os/exec"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run"
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
@@ -149,12 +147,7 @@ func New(l log.Logger, opts *options.TerragruntOptions
 	return allCommands
 }
 
-// WrapWithTelemetry wraps CLI command execution with telemetry initialization,
-// context setting and labels. If telemetry is disabled, just runs the command.
-//
-// The telemeter is created here rather than at app startup because command
-// actions run after the full CLI parse, so telemetry options and experiments
-// (such as otel-logs) are honored whether they were set via flags or env vars.
+// WrapWithTelemetry wraps CLI command execution.
 func WrapWithTelemetry(
 	l log.Logger,
 	opts *options.TerragruntOptions,
@@ -165,65 +158,29 @@ func WrapWithTelemetry(
 		cliCtx *clihelper.Context,
 		action clihelper.ActionFunc,
 	) error {
-		telemeter, err := telemetry.NewTelemeter(
-			ctx,
-			l,
-			cliCtx.App.Name,
-			cliCtx.App.Version,
-			cliCtx.App.Writer,
-			opts.Telemetry,
-			opts.Experiments.Evaluate(experiment.OtelLogs),
-		)
-		if err != nil {
-			return err
-		}
-
 		defer func() {
-			if err := telemeter.Shutdown(ctx); err != nil {
+			if err := engine.Shutdown(
+				ctx,
+				l,
+				opts.Experiments,
+				opts.EngineOptions.NoEngine,
+			); err != nil {
 				if _, writeErr := cliCtx.App.ErrWriter.Write([]byte(err.Error())); writeErr != nil {
-					l.Debugf("Failed to report telemetry shutdown error %v: %v", err, writeErr)
+					l.Debugf("Failed to report engine shutdown error %v: %v", err, writeErr)
 				}
 			}
 		}()
 
-		ctx = telemetry.ContextWithTelemeter(ctx, telemeter)
+		if err := initialSetup(cliCtx, l, v, opts); err != nil {
+			return err
+		}
 
-		cmdName := fmt.Sprintf(
-			"%s %s", cliCtx.Command.Name, opts.TerraformCommand,
-		)
+		if err := RunAction(ctx, cliCtx, l, opts, v, action); err != nil {
+			opts.Tips.Find(tips.DebuggingDocs).Evaluate(l)
+			return err
+		}
 
-		return telemeter.Collect(ctx, l, cmdName, map[string]any{
-			"terraformCommand": opts.TerraformCommand,
-			"args":             opts.TerraformCliArgs,
-			"dir":              opts.WorkingDir,
-		}, func(childCtx context.Context, l log.Logger) error {
-			// Engines emit telemetry during shutdown, so stop them while the
-			// command span is still open (and before the telemeter's deferred
-			// Shutdown flushes) to keep their records correlated with it.
-			defer func() {
-				if err := engine.Shutdown(
-					childCtx,
-					l,
-					opts.Experiments,
-					opts.EngineOptions.NoEngine,
-				); err != nil {
-					if _, writeErr := cliCtx.App.ErrWriter.Write([]byte(err.Error())); writeErr != nil {
-						l.Debugf("Failed to report engine shutdown error %v: %v", err, writeErr)
-					}
-				}
-			}()
-
-			if err := initialSetup(cliCtx, l, v, opts); err != nil {
-				return err
-			}
-
-			if err := RunAction(childCtx, cliCtx, l, opts, v, action); err != nil {
-				opts.Tips.Find(tips.DebuggingDocs).Evaluate(l)
-				return err
-			}
-
-			return nil
-		})
+		return nil
 	}
 }
 
