--- internal/golang/inline.go.orig	2025-07-28 18:28:48 UTC
+++ internal/golang/inline.go
@@ -59,7 +59,6 @@ func inlineCall(ctx context.Context, snapshot *cache.S
 }
 
 func inlineCall(ctx context.Context, snapshot *cache.Snapshot, callerPkg *cache.Package, callerPGF *parsego.File, start, end token.Pos) (_ *token.FileSet, _ *analysis.SuggestedFix, err error) {
-	countInlineCall.Inc()
 	// Find enclosing static call.
 	call, fn, err := enclosingStaticCall(callerPkg, callerPGF, start, end)
 	if err != nil {
@@ -178,7 +177,6 @@ func inlineVariableOne(pkg *cache.Package, pgf *parseg
 // inlineVariableOne computes a fix to replace the selected variable by
 // its initialization expression.
 func inlineVariableOne(pkg *cache.Package, pgf *parsego.File, start, end token.Pos) (*token.FileSet, *analysis.SuggestedFix, error) {
-	countInlineVariable.Inc()
 	info := pkg.TypesInfo()
 	curUse, curRHS, ok := canInlineVariable(info, pgf.Cursor, start, end)
 	if !ok {
