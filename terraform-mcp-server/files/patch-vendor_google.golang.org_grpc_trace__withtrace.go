--- vendor/google.golang.org/grpc/trace_withtrace.go.orig	2026-11-01 19:57:46 UTC
+++ vendor/google.golang.org/grpc/trace_withtrace.go
@@ -22,18 +22,27 @@ import (
 
 import (
 	"context"
-
-	t "golang.org/x/net/trace"
+	"fmt"
 )
 
+type notrace struct{}
+
+func (notrace) LazyLog(x fmt.Stringer, sensitive bool) {}
+func (notrace) LazyPrintf(format string, a ...any)     {}
+func (notrace) SetError()                              {}
+func (notrace) SetRecycler(f func(any))                {}
+func (notrace) SetTraceInfo(traceID, spanID uint64)    {}
+func (notrace) SetMaxEvents(m int)                     {}
+func (notrace) Finish()                                {}
+
 func newTrace(family, title string) traceLog {
-	return t.New(family, title)
+	return notrace{}
 }
 
 func newTraceContext(ctx context.Context, tr traceLog) context.Context {
-	return t.NewContext(ctx, tr)
+	return ctx
 }
 
 func newTraceEventLog(family, title string) traceEventLog {
-	return t.NewEventLog(family, title)
+	return nil
 }
