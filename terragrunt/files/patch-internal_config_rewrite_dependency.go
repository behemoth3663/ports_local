--- internal/config/rewrite/dependency.go.orig	1979-11-29 21:00:00 UTC
+++ internal/config/rewrite/dependency.go
@@ -52,23 +52,7 @@ func (s *dependencySet) resolve(
 		names = append(names, dep.Name)
 	}
 
-	var result *cty.Value
-
-	err := pkgconfig.TraceParseDependencies(
-		ctx,
-		l,
-		s.configPath,
-		pctx.SkipOutputsResolution,
-		len(s.blocks),
-		names,
-		func(ctx context.Context, l log.Logger) error {
-			var encodeErr error
-
-			result, encodeErr = pkgconfig.DependencyBlocksToCtyValue(ctx, l, v, pctx, s.blocks)
-
-			return encodeErr
-		},
-	)
+			result, err := pkgconfig.DependencyBlocksToCtyValue(ctx, l, v, pctx, s.blocks)
 	if err != nil {
 		return failedDependencyValue(result), err
 	}
