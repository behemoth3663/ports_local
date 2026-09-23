--- internal/runner/run/hook.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/hook.go
@@ -14,7 +14,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/multierror"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tflint"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -91,13 +90,7 @@ func ProcessHooks(ctx context.Context, l log.Logger, v
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
@@ -132,8 +125,10 @@ func ProcessErrorHooks(
 	l.Debugf("Detected %d error Hooks", len(hooks))
 
 	errorMessages := make([]string, 0, len(previousExecErrors))
+
 	for _, e := range previousExecErrors {
 		errorMessage := e.Error()
+
 		// Process execution errors carry stdout that hook patterns need to match against.
 		// https://github.com/gruntwork-io/terragrunt/issues/2045
 		if processError, ok := errors.AsType[*util.ProcessExecutionError](e); ok {
@@ -152,42 +147,35 @@ func ProcessErrorHooks(
 	for _, curHook := range hooks {
 		if util.MatchesAny(curHook.OnErrors, errorMessage) &&
 			slices.Contains(curHook.Commands, opts.TerraformCommand) {
-			err := telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "error_hook_"+curHook.Name, map[string]any{
-					"hook": curHook.Name,
-					"dir":  curHook.WorkingDir,
-				}, func(ctx context.Context, l log.Logger) error {
-					l.Infof("Executing hook: %s", curHook.Name)
 
-					actionToExecute := curHook.Execute[0]
-					actionParams := curHook.Execute[1:]
+			l.Infof("Executing hook: %s", curHook.Name)
 
-					env, hookEnvErr := hookEnv(v.Env, opts, cfg, curHook.Name, HookTypeError)
-					if hookEnvErr != nil {
-						return hookEnvErr
-					}
+			actionToExecute := curHook.Execute[0]
+			actionParams := curHook.Execute[1:]
 
-					hookV := v.WithEnv(env)
+			env, hookEnvErr := hookEnv(v.Env, opts, cfg, curHook.Name, HookTypeError)
+			if hookEnvErr != nil {
+				errorsOccured = append(errorsOccured, hookEnvErr)
+				continue
+			}
 
-					_, possibleError := shell.RunCommandWithOutput(
-						ctx,
-						l,
-						hookV,
-						opts.shellRunOptions(env),
-						curHook.WorkingDir,
-						curHook.SuppressStdout,
-						false,
-						actionToExecute, actionParams...,
-					)
-					if possibleError != nil {
-						l.Errorf("%s", hookErrorMessage(curHook.Name, possibleError))
-						return possibleError
-					}
+			hookV := v.WithEnv(env)
 
-					return nil
-				})
-			if err != nil {
-				errorsOccured = append(errorsOccured, err)
+			_, possibleError := shell.RunCommandWithOutput(
+				ctx,
+				l,
+				hookV,
+				opts.shellRunOptions(env),
+				curHook.WorkingDir,
+				curHook.SuppressStdout,
+				false,
+				actionToExecute,
+				actionParams...,
+			)
+
+			if possibleError != nil {
+				l.Errorf("%s", hookErrorMessage(curHook.Name, possibleError))
+				errorsOccured = append(errorsOccured, possibleError)
 			}
 		}
 	}
