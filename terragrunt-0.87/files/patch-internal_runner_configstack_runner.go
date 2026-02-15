--- internal/runner/configstack/runner.go.orig	2025-09-25 14:56:12 UTC
+++ internal/runner/configstack/runner.go
@@ -20,7 +20,6 @@ import (
 
 	"github.com/gruntwork-io/go-commons/collections"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tf"
 
 	"github.com/gruntwork-io/terragrunt/config"
@@ -260,9 +259,6 @@ func (runner *Runner) createStackForTerragruntConfigPa
 // assembles them into a stack, and checks for dependency cycles. Updates the Runner's stack with the resolved units.
 // Returns an error if discovery or validation fails.
 func (runner *Runner) createStackForTerragruntConfigPaths(ctx context.Context, l log.Logger, terragruntConfigPaths []string) error {
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "create_stack_for_terragrunt_config_paths", map[string]any{
-		"working_dir": runner.Stack.TerragruntOptions.WorkingDir,
-	}, func(ctx context.Context) error {
 		if len(terragruntConfigPaths) == 0 {
 			return errors.New(common.ErrNoUnitsFound)
 		}
@@ -274,24 +270,9 @@ func (runner *Runner) createStackForTerragruntConfigPa
 
 		runner.Stack.Units = units
 
-		return nil
-	})
-	if err != nil {
-		return errors.New(err)
-	}
-
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "check_for_cycles", map[string]any{
-		"working_dir": runner.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
 		if err := runner.Stack.Units.CheckForCycles(); err != nil {
 			return errors.New(err)
 		}
-
-		return nil
-	})
-	if err != nil {
-		return errors.New(err)
-	}
 
 	return nil
 }
