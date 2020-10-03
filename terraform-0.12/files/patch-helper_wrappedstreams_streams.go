--- helper/wrappedstreams/streams.go.orig	2020-07-25 03:39:12 UTC
+++ helper/wrappedstreams/streams.go
@@ -10,32 +10,31 @@ import (
 
 // Stdin returns the true stdin of the process.
 func Stdin() *os.File {
-	stdin := os.Stdin
-	if panicwrap.Wrapped(nil) {
-		stdin = wrappedStdin
-	}
-
+	stdin, _, _ := fds()
 	return stdin
 }
 
 // Stdout returns the true stdout of the process.
 func Stdout() *os.File {
-	stdout := os.Stdout
-	if panicwrap.Wrapped(nil) {
-		stdout = wrappedStdout
-	}
-
+	_, stdout, _ := fds()
 	return stdout
 }
 
 // Stderr returns the true stderr of the process.
 func Stderr() *os.File {
-	stderr := os.Stderr
+	_, _, stderr := fds()
+	return stderr
+}
+
+func fds() (stdin, stdout, stderr *os.File) {
+	stdin, stdout, stderr = os.Stdin, os.Stdout, os.Stderr
+
 	if panicwrap.Wrapped(nil) {
-		stderr = wrappedStderr
+		initPlatform()
+		stdin, stdout, stderr = wrappedStdin, wrappedStdout, wrappedStderr
 	}
 
-	return stderr
+	return
 }
 
 // These are the wrapped standard streams. These are setup by the
@@ -45,8 +44,3 @@ var (
 	wrappedStdout *os.File
 	wrappedStderr *os.File
 )
-
-func init() {
-	// Initialize the platform-specific code
-	initPlatform()
-}
