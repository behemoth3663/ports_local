--- helper/wrappedstreams/streams_other.go.orig	2020-07-25 04:03:45 UTC
+++ helper/wrappedstreams/streams_other.go
@@ -4,11 +4,18 @@ package wrappedstreams
 
 import (
 	"os"
+	"sync"
 )
 
+var initOnce sync.Once
+
 func initPlatform() {
-	// The standard streams are passed in via extra file descriptors.
-	wrappedStdin = os.NewFile(uintptr(3), "stdin")
-	wrappedStdout = os.NewFile(uintptr(4), "stdout")
-	wrappedStderr = os.NewFile(uintptr(5), "stderr")
+	// These must be initialized lazily, once it's been determined that this is
+	// a wrapped process.
+	initOnce.Do(func() {
+		// The standard streams are passed in via extra file descriptors.
+		wrappedStdin = os.NewFile(uintptr(3), "stdin")
+		wrappedStdout = os.NewFile(uintptr(4), "stdout")
+		wrappedStderr = os.NewFile(uintptr(5), "stderr")
+	})
 }
