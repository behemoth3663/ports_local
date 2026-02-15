--- config/config.go.orig	2025-09-25 14:56:12 UTC
+++ config/config.go
@@ -23,7 +23,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/remotestate"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 
 	"github.com/mitchellh/mapstructure"
 
@@ -1172,10 +1171,6 @@ func ParseConfigFile(ctx *ParsingContext, l log.Logger
 
 	hclCache := cache.ContextCache[*hclparse.File](ctx, HclCacheContextKey)
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "parse_config_file", map[string]any{
-		"config_path": configPath,
-		"working_dir": ctx.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
 		childKey := "nil"
 		if includeFromChild != nil {
 			childKey = includeFromChild.String()
@@ -1188,16 +1183,16 @@ func ParseConfigFile(ctx *ParsingContext, l log.Logger
 
 		dir, err := os.Getwd()
 		if err != nil {
-			return err
+			return config, err
 		}
 
 		fileInfo, err := os.Stat(configPath)
 		if err != nil {
 			if os.IsNotExist(err) {
-				return TerragruntConfigNotFoundError{Path: configPath}
+				return config, TerragruntConfigNotFoundError{Path: configPath}
 			}
 
-			return errors.Errorf("failed to get file info: %w", err)
+			return config, errors.Errorf("failed to get file info: %w", err)
 		}
 
 		var (
@@ -1212,7 +1207,7 @@ func ParseConfigFile(ctx *ParsingContext, l log.Logger
 			// Parse the HCL file into an AST body that can be decoded multiple times later without having to re-parse
 			file, err = hclparse.NewParser(ctx.ParserOptions...).ParseFromFile(configPath)
 			if err != nil {
-				return err
+				return config, err
 			}
 			// TODO: Remove lint ignore
 			hclCache.Put(ctx, cacheKey, file) //nolint:contextcheck
@@ -1221,14 +1216,8 @@ func ParseConfigFile(ctx *ParsingContext, l log.Logger
 		// TODO: Remove lint ignore
 		config, err = ParseConfig(ctx, l, file, includeFromChild) //nolint:contextcheck
 		if err != nil {
-			return err
+			return config, err
 		}
-
-		return nil
-	})
-	if err != nil {
-		return config, err
-	}
 
 	return config, nil
 }
