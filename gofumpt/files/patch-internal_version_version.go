--- internal/version/version.go.orig	2021-12-16 14:34:28 UTC
+++ internal/version/version.go
@@ -6,22 +6,13 @@ package version
 import (
 	"fmt"
 	"os"
-	"runtime/debug"
 )
 
-var version = "(devel)" // to match the default from runtime/debug
+var version = "v0.2.1"
 
 func String() string {
 	if testVersion := os.Getenv("GOFUMPT_VERSION_TEST"); testVersion != "" {
 		return testVersion
-	}
-	// don't overwrite the version if it was set by -ldflags=-X
-	if info, ok := debug.ReadBuildInfo(); ok && version == "(devel)" {
-		mod := &info.Main
-		if mod.Replace != nil {
-			mod = mod.Replace
-		}
-		version = mod.Version
 	}
 	return version
 }
