--- cmd/govulncheck/main.go.orig	2025-01-06 19:26:26 UTC
+++ cmd/govulncheck/main.go
@@ -9,13 +9,10 @@ import (
 	"fmt"
 	"os"
 
-	"golang.org/x/telemetry"
 	"golang.org/x/vuln/scan"
 )
 
 func main() {
-	telemetry.Start(telemetry.Config{ReportCrashes: true})
-
 	ctx := context.Background()
 
 	cmd := scan.Command(ctx, os.Args[1:]...)
