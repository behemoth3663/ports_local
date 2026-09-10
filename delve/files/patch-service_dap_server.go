--- service/dap/server.go.orig	2026-09-09 15:48:17 UTC
+++ service/dap/server.go
@@ -691,7 +691,6 @@ func (s *Session) recoverPanic(request dap.Message) {
 // in case it's a dup and ignored by the client, we also log the error.
 func (s *Session) recoverPanic(request dap.Message) {
 	if ierr := recover(); ierr != nil {
-		logflags.Bug.Inc()
 		s.config.log.Errorf("recovered panic: %s\n%s\n", ierr, debug.Stack())
 		s.sendInternalErrorResponse(request.GetSeq(), fmt.Sprintf("%v", ierr))
 	}
