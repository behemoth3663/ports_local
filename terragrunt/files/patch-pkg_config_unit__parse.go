--- pkg/config/unit_parse.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/unit_parse.go
@@ -317,38 +317,27 @@ func (s *Store) resolveParsed(
 		cfg *TerragruntConfig
 	)
 
-	err := TraceParseConfigFile(
-		ctx,
-		l,
-		cfgPath,
-		pc.WorkingDir,
-		len(pc.PartialParseDecodeList) > 0,
-		pc.PartialParseDecodeList,
-		pc.include,
-		false,
-		func(childCtx context.Context, l log.Logger) error {
 			parser := hclparse.NewParser(pc.ParserOptions(l, v)...)
 
-			file, parseErr := parse(childCtx, parser)
+			file, parseErr := parse(ctx, parser)
 			if parseErr != nil {
-				return parseErr
+				return nil, nil, parseErr
 			}
 
-			base, parseErr := s.resolveBase(childCtx, l, v, pc, file)
+			base, parseErr := s.resolveBase(ctx, l, v, pc, file)
 			if base == nil {
-				return parseErr
+				return nil, nil, parseErr
 			}
 
-			fns := newParseFunctions(childCtx, l, v, s, file)
-			c = s.prepare(childCtx, l, v, pc, file, base, fns)
+			fns := newParseFunctions(ctx, l, v, s, file)
+			c = s.prepare(ctx, l, v, pc, file, base, fns)
 
-			cfg, parseErr = s.assembleFirst(childCtx, l, v, pc, c, fns)
+			cfg, parseErr = s.assembleFirst(ctx, l, v, pc, c, fns)
+			if parseErr != nil {
+				return nil, nil, parseErr
+			}
 
-			return parseErr
-		},
-	)
-
-	return c, cfg, err
+	return c, cfg, nil
 }
 
 // resolveBase returns the base of file, through s when pc is memoized. A base that holds errors is not retained, and
