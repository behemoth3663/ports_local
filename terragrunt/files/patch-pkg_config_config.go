--- pkg/config/config.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/config.go
@@ -1643,20 +1643,8 @@ func ParseConfigFile(
 
 	cached := cachedHCLFile(ctx, v.FS, hclCache, cacheKey, cfgPath)
 
-	isPartial := len(pctx.PartialParseDecodeList) > 0
-
-	err = TraceParseConfigFile(
-		ctx,
-		l,
-		cfgPath,
-		pctx.WorkingDir,
-		isPartial,
-		pctx.PartialParseDecodeList,
-		includeFromChild,
-		cached != nil,
-		func(childCtx context.Context, l log.Logger) error {
-			file, parseErr := rebindOrParseHCLFile(
-				childCtx,
+			file, err := rebindOrParseHCLFile(
+				ctx,
 				l,
 				v,
 				pctx,
@@ -1666,20 +1654,14 @@ func ParseConfigFile(
 				cfgPath,
 				fileInfo.ModTime(),
 			)
-			if parseErr != nil {
-				return parseErr
+			if err != nil {
+				return config, err
 			}
 
-			config, parseErr = ParseConfig(childCtx, l, v, pctx, file, includeFromChild)
-			if parseErr != nil {
-				return parseErr
+			config, err = ParseConfig(ctx, l, v, pctx, file, includeFromChild)
+			if err != nil {
+				return config, err
 			}
-
-			return nil
-		})
-	if err != nil {
-		return config, err
-	}
 
 	return config, nil
 }
