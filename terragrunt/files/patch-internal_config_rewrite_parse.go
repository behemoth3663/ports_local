--- internal/config/rewrite/parse.go.orig	1979-11-29 21:00:00 UTC
+++ internal/config/rewrite/parse.go
@@ -222,35 +222,21 @@ func (s *Store) resolveFileOnce(
 		cfg *pkgconfig.TerragruntConfig
 	)
 
-	err = pkgconfig.TraceParseConfigFile(
-		ctx,
-		l,
-		cfgPath,
-		pc.settings.WorkingDir,
-		len(pc.settings.PartialParseDecodeList) > 0,
-		pc.settings.PartialParseDecodeList,
-		pc.file.include,
-		false,
-		func(childCtx context.Context, l log.Logger) error {
 			parser := pkghclparse.NewParser(pc.parserOptions(l, v)...)
 
-			file, parseErr := s.ParseUnitAST(childCtx, v.FS, cfgPath, info, parser)
-			if parseErr != nil {
-				return parseErr
+			file, err := s.ParseUnitAST(ctx, v.FS, cfgPath, info, parser)
+			if err != nil {
+				return c, cfg, err
 			}
 
-			base, parseErr := s.resolveBase(childCtx, l, v, pc, file)
+			base, err := s.resolveBase(ctx, l, v, pc, file)
 			if base == nil {
-				return parseErr
+				return c, cfg, err
 			}
 
-			c = prepare(childCtx, l, v, pc, file, base)
+			c = prepare(ctx, l, v, pc, file, base)
 
-			cfg, parseErr = s.assembleFirst(childCtx, l, v, pc, c)
-
-			return parseErr
-		},
-	)
+			cfg, err = s.assembleFirst(ctx, l, v, pc, c)
 
 	return c, cfg, err
 }
