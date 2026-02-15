--- vendor/cloud.google.com/go/storage/writer.go.orig	2025-02-06 22:11:02 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -22,8 +22,6 @@ import (
 	"sync"
 	"time"
 	"unicode/utf8"
-
-	"cloud.google.com/go/internal/trace"
 )
 
 // A Writer writes a Cloud Storage object.
@@ -190,7 +188,6 @@ func (w *Writer) Close() error {
 	<-w.donec
 	w.mu.Lock()
 	defer w.mu.Unlock()
-	trace.EndSpan(w.ctx, w.err)
 	return w.err
 }
 
