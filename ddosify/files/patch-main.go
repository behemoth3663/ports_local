--- main.go.orig	2021-10-11 12:25:31 UTC
+++ main.go
@@ -32,7 +32,6 @@ import (
 	"runtime"
 	"strings"
 	"text/tabwriter"
-	"time"
 
 	"go.ddosify.com/ddosify/config"
 	"go.ddosify.com/ddosify/core"
@@ -73,7 +72,7 @@ var (
 var (
 	GitVersion = "development"
 	GitCommit  = "unknown"
-	BuildDate  = time.Now().UTC().Format(time.RFC3339)
+	BuildDate  = ""
 )
 
 func main() {
