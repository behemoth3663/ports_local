--- cmd/deadcode/deadcode.go.orig	2025-03-05 20:18:18 UTC
+++ cmd/deadcode/deadcode.go
@@ -26,7 +26,6 @@ import (
 	"strings"
 	"text/template"
 
-	"golang.org/x/telemetry"
 	"golang.org/x/tools/go/callgraph"
 	"golang.org/x/tools/go/callgraph/rta"
 	"golang.org/x/tools/go/packages"
@@ -64,8 +63,6 @@ func main() {
 }
 
 func main() {
-	telemetry.Start(telemetry.Config{ReportCrashes: true})
-
 	log.SetPrefix("deadcode: ")
 	log.SetFlags(0) // no time prefix
 
