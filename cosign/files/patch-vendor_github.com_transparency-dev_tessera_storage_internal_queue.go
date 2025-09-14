--- vendor/github.com/transparency-dev/tessera/storage/internal/queue.go.orig	2025-09-22 08:57:07 UTC
+++ vendor/github.com/transparency-dev/tessera/storage/internal/queue.go
@@ -133,9 +133,6 @@ func (q *Queue) doFlush(ctx context.Context, f FlushFu
 
 // doFlush handles the queue flush, and sending notifications of assigned log indices.
 func (q *Queue) doFlush(ctx context.Context, f FlushFunc, entries []queueItem) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.queue.doFlush")
-	defer span.End()
-
 	entriesData := make([]*tessera.Entry, 0, len(entries))
 	for _, e := range entries {
 		entriesData = append(entriesData, e.entry)
