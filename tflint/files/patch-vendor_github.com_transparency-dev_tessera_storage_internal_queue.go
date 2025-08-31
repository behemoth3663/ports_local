--- vendor/github.com/transparency-dev/tessera/storage/internal/queue.go.orig	2025-06-10 15:09:26 UTC
+++ vendor/github.com/transparency-dev/tessera/storage/internal/queue.go
@@ -95,9 +95,6 @@ func (q *Queue) Add(ctx context.Context, e *tessera.En
 
 // Add places e into the queue, and returns a func which may be called to retrieve the assigned index.
 func (q *Queue) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
-	_, span := tracer.Start(ctx, "tessera.storage.queue.Add")
-	defer span.End()
-
 	qi := newEntry(ctx, e)
 
 	if err := q.buf.Push(qi); err != nil {
@@ -108,9 +105,6 @@ func (q *Queue) doFlush(ctx context.Context, entries [
 
 // doFlush handles the queue flush, and sending notifications of assigned log indices.
 func (q *Queue) doFlush(ctx context.Context, entries []*queueItem) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.queue.doFlush")
-	defer span.End()
-
 	entriesData := make([]*tessera.Entry, 0, len(entries))
 	for _, e := range entries {
 		entriesData = append(entriesData, e.entry)
@@ -136,14 +130,11 @@ func newEntry(ctx context.Context, data *tessera.Entry
 
 // newEntry creates a new entry for the provided data.
 func newEntry(ctx context.Context, data *tessera.Entry) *queueItem {
-	_, span := tracer.Start(ctx, "tessera.storage.queue.future")
-
 	e := &queueItem{
 		entry: data,
 		c:     make(chan tessera.IndexFuture, 1),
 	}
 	e.f = sync.OnceValues(func() (tessera.Index, error) {
-		defer span.End()
 		return (<-e.c)()
 	})
 	return e
