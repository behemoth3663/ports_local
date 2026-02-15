--- internal/runner/configstack/builder.go.orig	2025-09-25 14:56:12 UTC
+++ internal/runner/configstack/builder.go
@@ -7,25 +7,13 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/common"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Build builds a new Runner.
 func Build(ctx context.Context, l log.Logger, terragruntOptions *options.TerragruntOptions, opts ...common.Option) (common.StackRunner, error) {
 	var terragruntConfigFiles []string
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "find_files_in_path", map[string]any{
-		"working_dir": terragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
-		result, err := config.FindConfigFilesInPath(terragruntOptions.WorkingDir, terragruntOptions)
-		if err != nil {
-			return err
-		}
-
-		terragruntConfigFiles = result
-
-		return nil
-	})
+	terragruntConfigFiles, err := config.FindConfigFilesInPath(terragruntOptions.WorkingDir, terragruntOptions)
 	if err != nil {
 		return nil, err
 	}
