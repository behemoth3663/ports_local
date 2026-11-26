--- internal/runner/run/hook.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/hook.go
@@ -14,7 +14,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/tflint"
 	"github.com/gruntwork-io/terragrunt/internal/util"
@@ -106,13 +105,7 @@ func ProcessHooks(ctx context.Context, l log.Logger, v
 		}
 
 		if shouldRunHook(curHook, p.Opts, allPreviousErrors) {
-			err := telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "hook_"+curHook.Name, map[string]any{
-					"hook": curHook.Name,
-					"dir":  curHook.WorkingDir,
-				}, func(ctx context.Context, l log.Logger) error {
-					return runHook(ctx, l, v, p.Opts, p.Cfg, curHook, p.HookType)
-				})
+			err := runHook(ctx, l, v, p.Opts, p.Cfg, curHook, p.HookType)
 			if err != nil {
 				allPreviousErrors = append(allPreviousErrors, err)
 			}
@@ -174,11 +167,6 @@ func ProcessErrorHooks(
 
 		if matched &&
 			slices.Contains(curHook.Commands, opts.TerraformCliArgs.CommandName(opts.TerraformCommand)) {
-			err := telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "error_hook_"+curHook.Name, map[string]any{
-					"hook": curHook.Name,
-					"dir":  curHook.WorkingDir,
-				}, func(ctx context.Context, l log.Logger) error {
 					l.Infof("Executing hook: %s", curHook.Name)
 
 					actionToExecute := curHook.Execute[0]
@@ -186,8 +174,8 @@ func ProcessErrorHooks(
 
 					env, hookEnvErr := hookEnv(v, opts, cfg, curHook.Name, HookTypeError)
 					if hookEnvErr != nil {
-						return hookEnvErr
-					}
+						errorsOccured = append(errorsOccured, hookEnvErr)
+					} else {
 
 					hookV := v.WithEnv(env)
 
@@ -203,14 +191,9 @@ func ProcessErrorHooks(
 					)
 					if possibleError != nil {
 						l.Errorf("%s", hookErrorMessage(curHook.Name, possibleError))
-						return possibleError
+						errorsOccured = append(errorsOccured, possibleError)
 					}
-
-					return nil
-				})
-			if err != nil {
-				errorsOccured = append(errorsOccured, err)
-			}
+					}
 		}
 	}
 
