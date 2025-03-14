--- internal/util/bug/bug.go.orig	2025-02-24 17:39:34 UTC
+++ internal/util/bug/bug.go
@@ -18,8 +18,6 @@ import (
 	"sort"
 	"sync"
 	"time"
-
-	"golang.org/x/telemetry/counter"
 )
 
 // PanicOnBugs controls whether to panic when bugs are reported.
@@ -68,9 +66,6 @@ func Report(description string) {
 	report(description)
 }
 
-// BugReportCount is a telemetry counter that tracks # of bug reports.
-var BugReportCount = counter.NewStack("gopls/bug", 16)
-
 func report(description string) {
 	_, file, line, ok := runtime.Caller(2) // all exported reporting functions call report directly
 
@@ -92,22 +87,17 @@ func report(description string) {
 		AtTime:      time.Now(),
 	}
 
-	newBug := false
 	mu.Lock()
 	if _, ok := exemplars[key]; !ok {
 		if exemplars == nil {
 			exemplars = make(map[string]Bug)
 		}
 		exemplars[key] = bug // capture one exemplar per key
-		newBug = true
 	}
 	hh := handlers
 	handlers = nil
 	mu.Unlock()
 
-	if newBug {
-		BugReportCount.Inc()
-	}
 	// Call the handlers outside the critical section since a
 	// handler may itself fail and call bug.Report. Since handlers
 	// are one-shot, the inner call should be trivial.
