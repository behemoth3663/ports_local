--- cmd/goimports/goimports.go.orig	2025-09-10 14:57:49 UTC
+++ cmd/goimports/goimports.go
@@ -20,7 +20,6 @@ import (
 	"runtime/pprof"
 	"strings"
 
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/internal/gocommand"
 	"golang.org/x/tools/internal/imports"
 )
@@ -200,8 +199,6 @@ func main() {
 
 func main() {
 	// is anyone using this command?
-	counter.Open()
-	counter.Inc("tools/cmd:goimports")
 	runtime.GOMAXPROCS(runtime.NumCPU())
 
 	// call gofmtMain in a separate function
