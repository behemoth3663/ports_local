--- cli/commands/common/runall/runall.go.orig	2025-11-05 17:03:04 UTC
+++ cli/commands/common/runall/runall.go
@@ -14,7 +14,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/shell"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tf"
 )
 
@@ -122,10 +121,7 @@ func RunAllOnStack(ctx context.Context, l log.Logger, 
 
 	var runErr error
 
-	telemetryErr := telemetry.TelemeterFromContext(ctx).Collect(ctx, "run_all_on_stack", map[string]any{
-		"terraform_command": opts.TerraformCommand,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context) error {
+	telemetryErr := func(ctx context.Context) error {
 		err := runner.Run(ctx, l, opts)
 		if err != nil {
 			// At this stage, we can't handle the error any further, so we just log it and return nil.
@@ -151,7 +147,7 @@ func RunAllOnStack(ctx context.Context, l log.Logger, 
 		}
 
 		return nil
-	})
+	}(ctx)
 
 	// log telemetry error and continue execution
 	if telemetryErr != nil {
