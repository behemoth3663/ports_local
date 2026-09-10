--- pkg/proc/eval.go.orig	2026-09-09 15:48:17 UTC
+++ pkg/proc/eval.go
@@ -1100,7 +1100,6 @@ func (stack *evalStack) executeOp() {
 	defer func() {
 		err := recover()
 		if err != nil {
-			logflags.Bug.Inc()
 			stack.err = fmt.Errorf("internal debugger error: %v (recovered)\n%s", err, string(debug.Stack()))
 		}
 	}()
