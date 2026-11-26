--- pkg/config/unit_dependency.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/unit_dependency.go
@@ -117,21 +117,7 @@ func (set *dependencySet) resolve(
 
 	var result *cty.Value
 
-	err = TraceParseDependencies(
-		ctx,
-		l,
-		set.configPath,
-		pctx.SkipOutputsResolution,
-		len(blocks),
-		names,
-		func(ctx context.Context, l log.Logger) error {
-			var encodeErr error
-
-			result, encodeErr = DependencyBlocksToCtyValue(ctx, l, v, s, pctx, blocks)
-
-			return encodeErr
-		},
-	)
+			result, err = DependencyBlocksToCtyValue(ctx, l, v, s, pctx, blocks)
 	if err != nil {
 		return &dependencyResolution{value: failedDependencyValue(result), blocks: blocks}, err
 	}
