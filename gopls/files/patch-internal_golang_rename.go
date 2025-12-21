--- internal/golang/rename.go.orig	2025-12-05 20:12:17 UTC
+++ internal/golang/rename.go
@@ -445,7 +445,6 @@ func Rename(ctx context.Context, snapshot *cache.Snaps
 		newPkgDir string // declared here so it can be used later to move the package's contents to the new directory
 	)
 	if inPackageName {
-		countRenamePackage.Inc()
 		var (
 			newPkgName PackageName
 			newPkgPath PackagePath
