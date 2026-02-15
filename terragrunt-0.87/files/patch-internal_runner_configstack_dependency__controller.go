--- internal/runner/configstack/dependency_controller.go.orig	2025-09-25 14:56:12 UTC
+++ internal/runner/configstack/dependency_controller.go
@@ -10,7 +10,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/runner/common"
 	"github.com/gruntwork-io/terragrunt/options"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 const (
@@ -46,12 +45,7 @@ func (ctrl *DependencyController) runUnitWhenReady(ctx
 
 // runUnitWhenReady a unit once all of its dependencies have finished executing.
 func (ctrl *DependencyController) runUnitWhenReady(ctx context.Context, opts *options.TerragruntOptions, r *report.Report, semaphore chan struct{}) {
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "wait_for_unit_ready", map[string]any{
-		"path":             ctrl.Runner.Unit.Path,
-		"terraformCommand": ctrl.Runner.Unit.TerragruntOptions.TerraformCommand,
-	}, func(_ context.Context) error {
-		return ctrl.waitForDependencies(opts, r)
-	})
+	err := ctrl.waitForDependencies(opts, r)
 
 	semaphore <- struct{}{} // Add one to the buffered channel. Will block if parallelism limit is met
 
@@ -60,12 +54,7 @@ func (ctrl *DependencyController) runUnitWhenReady(ctx
 	}()
 
 	if err == nil {
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "run_unit", map[string]any{
-			"path":             ctrl.Runner.Unit.Path,
-			"terraformCommand": ctrl.Runner.Unit.TerragruntOptions.TerraformCommand,
-		}, func(ctx context.Context) error {
-			return ctrl.Runner.Run(ctx, opts, r)
-		})
+		err = ctrl.Runner.Run(ctx, opts, r)
 	}
 
 	ctrl.unitFinished(err, r, opts.Experiments.Evaluate(experiment.Report))
