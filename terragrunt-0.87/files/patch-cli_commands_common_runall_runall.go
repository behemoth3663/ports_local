--- cli/commands/common/runall/runall.go.orig	2025-09-25 14:56:12 UTC
+++ cli/commands/common/runall/runall.go
@@ -15,7 +15,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/shell"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tf"
 )
 
@@ -120,10 +119,6 @@ func RunAllOnStack(ctx context.Context, l log.Logger, 
 		}
 	}
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "run_all_on_stack", map[string]any{
-		"terraform_command": opts.TerraformCommand,
-		"working_dir":       opts.WorkingDir,
-	}, func(ctx context.Context) error {
 		err := runner.Run(ctx, l, opts)
 		if err != nil {
 			// At this stage, we can't handle the error any further, so we just log it and return nil.
@@ -145,5 +140,4 @@ func RunAllOnStack(ctx context.Context, l log.Logger, 
 		}
 
 		return nil
-	})
 }
