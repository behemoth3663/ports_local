--- vendor/cloud.google.com/go/storage/writer.go.orig	2026-03-13 06:41:11 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -285,7 +285,6 @@ func (w *Writer) Close() error {
 	w.closed = true
 	w.mu.Lock()
 	defer w.mu.Unlock()
-	endSpan(w.ctx, w.err)
 	return w.err
 }
 
