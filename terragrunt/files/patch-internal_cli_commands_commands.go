diff --git a/internal/cli/commands/commands.go b/internal/cli/commands/commands.go
index 3a189da2c..0eab0cc16 100644
--- internal/cli/commands/commands.go.orig
+++ internal/cli/commands/commands.go
@@ -3,6 +3,7 @@ package commands
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"net"
 	"path/filepath"
@@ -27,8 +28,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/options"
 
-	"errors"
-
 	awsproviderpatch "github.com/gruntwork-io/terragrunt/internal/cli/commands/aws-provider-patch"
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/backend"
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/browse"
@@ -50,7 +49,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/os/exec"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run"
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
@@ -145,12 +143,7 @@ func New(l log.Logger, opts *options.TerragruntOptions, v *venv.Venv) clihelper.
 	return allCommands
 }
 
-// WrapWithTelemetry wraps CLI command execution with telemetry initialization,
-// context setting and labels. If telemetry is disabled, just runs the command.
-//
-// The telemeter is created here rather than at app startup because command
-// actions run after the full CLI parse, so telemetry options and experiments
-// (such as otel-logs) are honored whether they were set via flags or env vars.
+// WrapWithTelemetry keeps the command lifecycle hook point but executes without telemetry.
 func WrapWithTelemetry(
 	l log.Logger,
 	opts *options.TerragruntOptions,
@@ -161,61 +154,27 @@ func WrapWithTelemetry(
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
 				_, _ = cliCtx.App.ErrWriter.Write([]byte(err.Error()))
 			}
 		}()
 
-		ctx = telemetry.ContextWithTelemeter(ctx, telemeter)
-
-		cmdName := fmt.Sprintf(
-			"%s %s", cliCtx.Command.Name, opts.TerraformCommand,
-		)
-
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
-					_, _ = cliCtx.App.ErrWriter.Write([]byte(err.Error()))
-				}
-			}()
-
-			if err := initialSetup(cliCtx, l, v, opts); err != nil {
-				return err
-			}
+		if err := initialSetup(cliCtx, l, v, opts); err != nil {
+			return err
+		}
 
-			if err := RunAction(childCtx, cliCtx, l, opts, v, action); err != nil {
-				opts.Tips.Find(tips.DebuggingDocs).Evaluate(l)
-				return err
-			}
+		if err := RunAction(ctx, cliCtx, l, opts, v, action); err != nil {
+			opts.Tips.Find(tips.DebuggingDocs).Evaluate(l)
+			return err
+		}
 
-			return nil
-		})
+		return nil
 	}
 }
 
@@ -507,7 +466,7 @@ func setupAutoProviderCacheDir(
 
 	providerCacheDir = filepath.Clean(providerCacheDir)
 
-	const cacheDirMode = 0755
+	const cacheDirMode = 0o755
 
 	// Create the cache directory if it doesn't exist
 	if err := v.FS.MkdirAll(providerCacheDir, cacheDirMode); err != nil {
