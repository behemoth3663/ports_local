--- main.go.orig	2026-05-15 14:53:57 UTC
+++ main.go
@@ -15,8 +15,6 @@ import (
 	"log"
 	"os"
 
-	"golang.org/x/telemetry"
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/gopls/internal/cmd"
 	"golang.org/x/tools/gopls/internal/filecache"
 	versionpkg "golang.org/x/tools/gopls/internal/version"
@@ -28,11 +26,6 @@ func main() {
 func main() {
 	versionpkg.VersionOverride = version
 
-	telemetry.Start(telemetry.Config{
-		ReportCrashes: true,
-		Upload:        true,
-	})
-
 	// Force early creation of the filecache and refuse to start
 	// if there were unexpected errors such as ENOSPC. This
 	// minimizes the window of exposure to deletion of the
@@ -47,7 +40,6 @@ func main() {
 	// either re-create it or just fail the RPC with an
 	// informative error and terminate the process.
 	if _, err := filecache.Get("nonesuch", [32]byte{}, filecache.Bytes); err != nil && err != filecache.ErrNotFound {
-		counter.Inc("gopls/nocache")
 		log.Fatalf("gopls cannot access its persistent index (disk full?): %v", err)
 	}
 
