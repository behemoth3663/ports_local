--- vendor/cloud.google.com/go/storage/writer.go.orig	2025-03-12 21:31:11 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -22,8 +22,6 @@ import (
 	"sync"
 	"time"
 	"unicode/utf8"
-
-	"cloud.google.com/go/internal/trace"
 )
 
 // A Writer writes a Cloud Storage object.
@@ -233,7 +231,6 @@ func (w *Writer) Close() error {
 	w.closed = true
 	w.mu.Lock()
 	defer w.mu.Unlock()
-	trace.EndSpan(w.ctx, w.err)
 	return w.err
 }
 
