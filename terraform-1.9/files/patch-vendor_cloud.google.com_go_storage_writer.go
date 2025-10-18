--- vendor/cloud.google.com/go/storage/writer.go.orig	2025-10-15 02:58:16 UTC
+++ vendor/cloud.google.com/go/storage/writer.go
@@ -22,8 +22,6 @@ import (
 	"sync"
 	"time"
 	"unicode/utf8"
-
-	"cloud.google.com/go/internal/trace"
 )
 
 // Interface internalWriter wraps low-level implementations which may vary
@@ -267,7 +265,6 @@ func (w *Writer) Close() error {
 	if w.obj == nil && w.err == nil {
 		w.err = errors.New("storage: write succeeded but no object attributes returned from the server")
 	}
-	trace.EndSpan(w.ctx, w.err)
 	return w.err
 }
 
