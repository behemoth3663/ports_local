--- internal/cli/commands/commands.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/commands.go
@@ -11,15 +11,12 @@ import (
 	"strings"
 	"time"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 	"golang.org/x/sync/errgroup"
 
 	"github.com/gruntwork-io/terragrunt/internal/cache"
 	rewrite "github.com/gruntwork-io/terragrunt/internal/config/rewrite"
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
 	"github.com/gruntwork-io/terragrunt/internal/engine"
-	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/panicreport"
 	"github.com/gruntwork-io/terragrunt/internal/providercache"
@@ -57,7 +54,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/os/exec"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run"
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
@@ -169,12 +165,7 @@ const telemetryShutdownTimeout = 30 * time.Second
 // telemetryShutdownTimeout limits how long the telemetry flush can delay exit.
 const telemetryShutdownTimeout = 30 * time.Second
 
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
@@ -185,44 +176,14 @@ func WrapWithTelemetry(
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
-		// Flush even after an interrupt cancels ctx, but do not wait forever.
-		defer func() {
-			shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), telemetryShutdownTimeout)
-			defer cancel()
-
-			if err := telemeter.Shutdown(shutdownCtx); err != nil {
-				if _, writeErr := cliCtx.App.ErrWriter.Write([]byte(err.Error())); writeErr != nil {
-					l.Debugf("Failed to report telemetry shutdown error %v: %v", err, writeErr)
-				}
-			}
-		}()
-
-		ctx = telemetry.ContextWithTelemeter(ctx, telemeter)
-
 		resolveErr := resolveWorkingDir(v, opts)
 
-		return telemeter.Collect(ctx, l, cliCtx.Command.Name, map[string]any{
-			"dir": opts.WorkingDir,
-		}, func(childCtx context.Context, l log.Logger) error {
 			// Engines emit telemetry during shutdown, so stop them while the
 			// command span is still open (and before the telemeter's deferred
 			// Shutdown flushes) to keep their records correlated with it.
 			defer func() {
 				if err := engine.Shutdown(
-					childCtx,
+					ctx,
 					l,
 					opts.Experiments,
 					opts.EngineOptions.NoEngine,
@@ -241,15 +202,12 @@ func WrapWithTelemetry(
 				return err
 			}
 
-			setCommandSpanAttributes(childCtx, cliCtx, opts)
-
-			if err := RunAction(childCtx, cliCtx, l, opts, v, action); err != nil {
+			if err := RunAction(ctx, cliCtx, l, opts, v, action); err != nil {
 				opts.Tips.Find(tips.DebuggingDocs).Evaluate(l)
 				return err
 			}
 
 			return nil
-		})
 	}
 }
 
@@ -771,21 +729,6 @@ func guardTerminal(l log.Logger, v *venv.Venv, reporte
 
 // setCommandSpanAttributes records on the command span the tofu command and
 // arguments that initialSetup resolves, for a command that runs OpenTofu/Terraform.
-func setCommandSpanAttributes(ctx context.Context, cliCtx *clihelper.Context, opts *options.TerragruntOptions) {
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
-		return
-	}
-
-	if !runsTofu(cliCtx) || opts.TofuCommand == "" {
-		return
-	}
-
-	span.SetAttributes(
-		attribute.String("tofu_command", opts.TofuCommand),
-		attribute.StringSlice("args", opts.TerraformCliArgs.Slice()),
-	)
-}
 
 // runsTofu reports whether cliCtx is a command that passes its arguments to
 // OpenTofu/Terraform: run, stack run, a subcommand of run, or a shortcut.
