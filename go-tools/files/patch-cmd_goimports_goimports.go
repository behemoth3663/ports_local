--- cmd/goimports/goimports.go.orig	2026-07-09 02:42:41 UTC
+++ cmd/goimports/goimports.go
@@ -21,7 +21,6 @@ import (
 	"strings"
 	"testing"
 
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/internal/gocommand"
 	"golang.org/x/tools/internal/imports"
 )
@@ -202,8 +201,6 @@ func main() {
 func main() {
 	// Measure how many people still use goimports.
 	// (See https://go.dev/issue/78671 for one.)
-	counter.Open()
-	counter.Inc("tools/cmd:goimports")
 	runtime.GOMAXPROCS(runtime.NumCPU())
 
 	// call gofmtMain in a separate function
