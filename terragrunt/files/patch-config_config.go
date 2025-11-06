--- config/config.go.orig	2025-11-05 17:03:04 UTC
+++ config/config.go
@@ -24,7 +24,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/remotestate"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 
 	"github.com/mitchellh/mapstructure"
 
@@ -1153,10 +1152,7 @@ func ParseConfigFile(ctx *ParsingContext, l log.Logger
 
 	hclCache := cache.ContextCache[*hclparse.File](ctx, HclCacheContextKey)
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "parse_config_file", map[string]any{
-		"config_path": configPath,
-		"working_dir": ctx.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		childKey := "nil"
 		if includeFromChild != nil {
 			childKey = includeFromChild.String()
@@ -1206,7 +1202,7 @@ func ParseConfigFile(ctx *ParsingContext, l log.Logger
 		}
 
 		return nil
-	})
+	}(ctx)
 	if err != nil {
 		return config, err
 	}
