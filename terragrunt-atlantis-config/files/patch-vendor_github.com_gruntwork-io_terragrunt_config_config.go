--- vendor/github.com/gruntwork-io/terragrunt/config/config.go.orig	2025-01-27 15:23:56 UTC
+++ vendor/github.com/gruntwork-io/terragrunt/config/config.go
@@ -18,7 +18,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/cache"
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/shell"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 
 	"github.com/mitchellh/mapstructure"
 
@@ -769,10 +768,7 @@ func ParseConfigFile(ctx *ParsingContext, configPath s
 
 	hclCache := cache.ContextCache[*hclparse.File](ctx, HclCacheContextKey)
 
-	err := telemetry.Telemetry(ctx, ctx.TerragruntOptions, "parse_config_file", map[string]interface{}{
-		"config_path": configPath,
-		"working_dir": ctx.TerragruntOptions.WorkingDir,
-	}, func(childCtx context.Context) error {
+	e := func(ctx *ParsingContext) error {
 		childKey := "nil"
 		if includeFromChild != nil {
 			childKey = includeFromChild.String()
@@ -818,7 +811,8 @@ func ParseConfigFile(ctx *ParsingContext, configPath s
 		}
 
 		return nil
-	})
+	}
+	err := e(ctx)
 	if err != nil {
 		return nil, err
 	}
