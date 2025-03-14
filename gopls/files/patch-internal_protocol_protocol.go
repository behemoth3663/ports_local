--- internal/protocol/protocol.go.orig	2025-02-24 17:39:34 UTC
+++ internal/protocol/protocol.go
@@ -11,7 +11,6 @@ import (
 	"fmt"
 	"io"
 
-	"golang.org/x/telemetry/crashmonitor"
 	"golang.org/x/tools/gopls/internal/util/bug"
 	"golang.org/x/tools/internal/event"
 	"golang.org/x/tools/internal/jsonrpc2"
@@ -299,15 +298,10 @@ func recoverHandlerPanic(method string) {
 }
 
 func recoverHandlerPanic(method string) {
-	// Report panics in the handler goroutine,
-	// unless we have enabled the monitor,
-	// which reports all crashes.
-	if !crashmonitor.Supported() {
 		defer func() {
 			if x := recover(); x != nil {
 				bug.Reportf("panic in %s request", method)
 				panic(x)
 			}
 		}()
-	}
 }
