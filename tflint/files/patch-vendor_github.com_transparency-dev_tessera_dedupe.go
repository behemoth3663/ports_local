--- vendor/github.com/transparency-dev/tessera/dedupe.go.orig	2025-06-10 15:09:26 UTC
+++ vendor/github.com/transparency-dev/tessera/dedupe.go
@@ -57,9 +57,6 @@ func (d *inMemoryDedupe) add(ctx context.Context, e *E
 // Add adds the entry to the underlying delegate only if e hasn't been recently seen. In either case,
 // an IndexFuture will be returned that the client can use to get the sequence number of this entry.
 func (d *inMemoryDedupe) add(ctx context.Context, e *Entry) IndexFuture {
-	ctx, span := tracer.Start(ctx, "tessera.Appender.inmemoryDedup.Add")
-	defer span.End()
-
 	id := string(e.Identity())
 
 	f := sync.OnceValue(func() IndexFuture {
