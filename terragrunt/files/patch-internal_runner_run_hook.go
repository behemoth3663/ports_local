--- internal/runner/run/hook.go.orig	2025-11-10 13:46:52 UTC
+++ internal/runner/run/hook.go
@@ -12,7 +12,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/shell"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tflint"
 	"github.com/gruntwork-io/terragrunt/util"
 	"github.com/hashicorp/go-multierror"
@@ -58,10 +57,7 @@ func processErrorHooks(ctx context.Context, l log.Logg
 
 	for _, curHook := range hooks {
 		if util.MatchesAny(curHook.OnErrors, errorMessage) && util.ListContainsElement(curHook.Commands, terragruntOptions.TerraformCommand) {
-			err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "error_hook_"+curHook.Name, map[string]any{
-				"hook": curHook.Name,
-				"dir":  curHook.WorkingDir,
-			}, func(ctx context.Context) error {
+			err := func(ctx context.Context) error {
 				l.Infof("Executing hook: %s", curHook.Name)
 
 				workingDir := ""
@@ -93,7 +89,7 @@ func processErrorHooks(ctx context.Context, l log.Logg
 				}
 
 				return nil
-			})
+			}(ctx)
 			if err != nil {
 				errorsOccured = multierror.Append(errorsOccured, err)
 			}
@@ -128,12 +124,9 @@ func processHooks(
 
 		allPreviousErrors := previousExecErrors.Append(errorsOccured)
 		if shouldRunHook(curHook, opts, allPreviousErrors) {
-			err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "hook_"+curHook.Name, map[string]any{
-				"hook": curHook.Name,
-				"dir":  curHook.WorkingDir,
-			}, func(ctx context.Context) error {
+			err := func(ctx context.Context) error {
 				return runHook(ctx, l, opts, cfg, curHook)
-			})
+			}(ctx)
 			if err != nil {
 				errorsOccured = multierror.Append(errorsOccured, err)
 			}
