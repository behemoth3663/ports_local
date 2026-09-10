--- service/rpccommon/server.go.orig	2026-09-09 15:48:17 UTC
+++ service/rpccommon/server.go
@@ -392,7 +392,6 @@ func newInternalError(ierr any, skip int) *internalErr
 }
 
 func newInternalError(ierr any, skip int) *internalError {
-	logflags.Bug.Inc()
 	r := &internalError{ierr, nil}
 	for i := skip; ; i++ {
 		pc, file, line, ok := runtime.Caller(i)
