--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-06-04 04:52:32 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -372,7 +372,6 @@ func (w *Writer) Close() error {
 			if w.err == nil && err != nil {
 				w.err = err
 			}
-			endSpan(w.ctx, w.err)
 			return w.err
 		}
 	}
@@ -391,7 +390,6 @@ func (w *Writer) Close() error {
 	w.mu.Lock()
 	defer w.mu.Unlock()
 	w.closed = true
-	endSpan(w.ctx, w.err)
 	return w.err
 }
 
