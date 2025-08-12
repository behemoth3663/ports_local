--- internal/golang/extract.go.orig	2025-07-28 18:28:48 UTC
+++ internal/golang/extract.go
@@ -32,13 +32,11 @@ func extractVariableOne(pkg *cache.Package, pgf *parse
 
 // extractVariableOne implements the refactor.extract.{variable,constant} CodeAction command.
 func extractVariableOne(pkg *cache.Package, pgf *parsego.File, start, end token.Pos) (*token.FileSet, *analysis.SuggestedFix, error) {
-	countExtractVariable.Inc()
 	return extractVariable(pkg, pgf, start, end, false)
 }
 
 // extractVariableAll implements the refactor.extract.{variable,constant}-all CodeAction command.
 func extractVariableAll(pkg *cache.Package, pgf *parsego.File, start, end token.Pos) (*token.FileSet, *analysis.SuggestedFix, error) {
-	countExtractVariableAll.Inc()
 	return extractVariable(pkg, pgf, start, end, true)
 }
 
@@ -570,13 +568,11 @@ func extractMethod(pkg *cache.Package, pgf *parsego.Fi
 
 // extractMethod refactors the selected block of code into a new method.
 func extractMethod(pkg *cache.Package, pgf *parsego.File, start, end token.Pos) (*token.FileSet, *analysis.SuggestedFix, error) {
-	countExtractMethod.Inc()
 	return extractFunctionMethod(pkg, pgf, start, end, true)
 }
 
 // extractFunction refactors the selected block of code into a new function.
 func extractFunction(pkg *cache.Package, pgf *parsego.File, start, end token.Pos) (*token.FileSet, *analysis.SuggestedFix, error) {
-	countExtractFunction.Inc()
 	return extractFunctionMethod(pkg, pgf, start, end, false)
 }
 
