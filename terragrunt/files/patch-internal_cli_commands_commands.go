--- internal/cli/commands/commands.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/commands.go
@@ -3,6 +3,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"net"
 	"path/filepath"
@@ -15,7 +16,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/cache"
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
 	"github.com/gruntwork-io/terragrunt/internal/engine"
-	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/panicreport"
 	"github.com/gruntwork-io/terragrunt/internal/providercache"
@@ -27,8 +27,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/options"
 
-	"errors"
-
 	awsproviderpatch "github.com/gruntwork-io/terragrunt/internal/cli/commands/aws-provider-patch"
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/backend"
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/browse"
@@ -52,7 +50,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/os/exec"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run"
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
@@ -165,61 +162,28 @@ func WrapWithTelemetry(
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
+		_ = fmt.Sprintf("%s %s", cliCtx.Command.Name, opts.TerraformCommand)
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
-					_, _ = cliCtx.App.ErrWriter.Write([]byte(err.Error()))
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
 
@@ -511,7 +475,7 @@ func setupAutoProviderCacheDir(
 
 	providerCacheDir = filepath.Clean(providerCacheDir)
 
-	const cacheDirMode = 0755
+	const cacheDirMode = 0o755
 
 	// Create the cache directory if it doesn't exist
 	if err := v.FS.MkdirAll(providerCacheDir, cacheDirMode); err != nil {
