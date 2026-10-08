--- internal/filecache/filecache.go.orig	2026-09-30 18:55:09 UTC
+++ internal/filecache/filecache.go
@@ -38,7 +38,6 @@ import (
 	"sync/atomic"
 	"time"
 
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/gopls/internal/util/bug"
 	"golang.org/x/tools/gopls/internal/util/lru"
 )
@@ -63,7 +62,6 @@ func Start() {
 		// either re-create it or just fail the RPC with an
 		// informative error and terminate the process.
 		if _, err := Get("nonesuch", [32]byte{}, Bytes); err != nil && err != ErrNotFound {
-			counter.Inc("gopls/nocache")
 			log.Fatalf("gopls cannot access its persistent index (disk full?): %v", err)
 		}
 	}()
