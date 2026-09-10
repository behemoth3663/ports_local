--- cmd/dlv/main.go.orig	2026-09-09 15:48:17 UTC
+++ cmd/dlv/main.go
@@ -6,17 +6,12 @@ import (
 	"github.com/go-delve/delve/cmd/dlv/cmds"
 	"github.com/go-delve/delve/pkg/logflags"
 	"github.com/go-delve/delve/pkg/version"
-	"golang.org/x/telemetry"
 )
 
 // Build is the git sha of this binaries build.
 var Build string
 
 func main() {
-	telemetry.Start(telemetry.Config{
-		ReportCrashes: true,
-	})
-
 	if Build != "" {
 		version.DelveVersion.Build = Build
 	}
