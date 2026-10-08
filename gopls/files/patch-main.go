--- main.go.orig	2026-09-30 18:55:09 UTC
+++ main.go
@@ -11,7 +11,6 @@ import (
 package main
 
 import (
-	"golang.org/x/telemetry"
 	"golang.org/x/tools/gopls/internal/cmd"
 	versionpkg "golang.org/x/tools/gopls/internal/version"
 )
@@ -20,11 +19,6 @@ func main() {
 
 func main() {
 	versionpkg.VersionOverride = version
-
-	telemetry.Start(telemetry.Config{
-		ReportCrashes: true,
-		Upload:        true,
-	})
 
 	cmd.Main()
 }
