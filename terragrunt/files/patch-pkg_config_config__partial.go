--- pkg/config/config_partial.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/config_partial.go
@@ -551,27 +551,12 @@ func PartialParseConfigFile(
 
 	var config *TerragruntConfig
 
-	err = TraceParseConfigFile(
-		ctx,
-		l,
-		cfgPath,
-		pctx.WorkingDir,
-		true, // isPartial
-		pctx.PartialParseDecodeList,
-		include,
-		cached != nil,
-		func(ctx context.Context, l log.Logger) error {
-			file, parseErr := rebindOrParseHCLFile(ctx, l, v, pctx, hclCache, cached, cacheKey, cfgPath, fileInfo.ModTime())
-			if parseErr != nil {
-				return parseErr
+			file, err := rebindOrParseHCLFile(ctx, l, v, pctx, hclCache, cached, cacheKey, cfgPath, fileInfo.ModTime())
+			if err != nil {
+				return config, err
 			}
 
-			config, parseErr = TerragruntConfigFromPartialConfig(ctx, l, v, pctx, file, include)
-
-			return parseErr
-		})
-
-	return config, err
+	return TerragruntConfigFromPartialConfig(ctx, l, v, pctx, file, include)
 }
 
 // TerragruntConfigFromPartialConfig is a wrapper of PartialParseConfigString which checks for cached configs.
