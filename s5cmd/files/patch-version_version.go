--- version/version.go.orig	2020-11-05 08:19:27 UTC
+++ version/version.go
@@ -1,7 +1,5 @@
 package version
 
-import "strings"
-
 var (
 	// Version represents the git tag of a particular release.
 	Version = "v0.0.0"
@@ -12,10 +10,5 @@ var (
 
 // GetHumanVersion returns human readable version information.
 func GetHumanVersion() string {
-	version := Version
-	if !strings.HasPrefix(version, "v") {
-		version = "v" + Version
-	}
-
-	return version + "-" + GitCommit
+	return Version
 }
