--- cli/commands/run/hook.go.orig	2025-09-25 14:56:12 UTC
+++ cli/commands/run/hook.go
@@ -11,7 +11,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/shell"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tflint"
 	"github.com/gruntwork-io/terragrunt/util"
 	"github.com/hashicorp/go-multierror"
@@ -116,12 +115,7 @@ func processHooks(
 
 		allPreviousErrors := previousExecErrors.Append(errorsOccured)
 		if shouldRunHook(curHook, opts, allPreviousErrors) {
-			err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "hook_"+curHook.Name, map[string]any{
-				"hook": curHook.Name,
-				"dir":  curHook.WorkingDir,
-			}, func(ctx context.Context) error {
-				return runHook(ctx, l, opts, cfg, curHook)
-			})
+			err := runHook(ctx, l, opts, cfg, curHook)
 			if err != nil {
 				errorsOccured = multierror.Append(errorsOccured, err)
 			}
