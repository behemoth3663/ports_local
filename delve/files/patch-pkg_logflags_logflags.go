--- pkg/logflags/logflags.go.orig	2026-09-09 15:48:17 UTC
+++ pkg/logflags/logflags.go
@@ -15,7 +15,6 @@ import (
 	"strings"
 	"time"
 
-	"golang.org/x/telemetry/counter"
 )
 
 var any = false
@@ -31,7 +30,6 @@ var logOut io.WriteCloser
 
 var logOut io.WriteCloser
 
-var Bug = counter.NewStack("delve/bug", 16)
 
 func makeLogger(flag bool, attrs ...interface{}) Logger {
 	if lf := loggerFactory; lf != nil {
