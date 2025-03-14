--- internal/server/command.go.orig	2025-02-24 17:39:34 UTC
+++ internal/server/command.go
@@ -23,7 +23,6 @@ import (
 	"sync"
 
 	"golang.org/x/mod/modfile"
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/go/ast/astutil"
 	"golang.org/x/tools/gopls/internal/cache"
 	"golang.org/x/tools/gopls/internal/cache/metadata"
@@ -284,7 +283,6 @@ func (*commandHandler) AddTelemetryCounters(_ context.
 		if n == "" || v < 0 {
 			continue
 		}
-		counter.Add("fwd/"+n, v)
 	}
 	return nil
 }
