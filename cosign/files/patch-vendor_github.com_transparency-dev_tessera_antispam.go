--- vendor/github.com/transparency-dev/tessera/antispam.go.orig	2025-09-22 08:57:07 UTC
+++ vendor/github.com/transparency-dev/tessera/antispam.go
@@ -57,9 +57,6 @@ func (d *inMemoryDedup) add(ctx context.Context, e *En
 // Add adds the entry to the underlying delegate only if e hasn't been recently seen. In either case,
 // an IndexFuture will be returned that the client can use to get the sequence number of this entry.
 func (d *inMemoryDedup) add(ctx context.Context, e *Entry) IndexFuture {
-	ctx, span := tracer.Start(ctx, "tessera.Appender.inmemoryDedup.Add")
-	defer span.End()
-
 	id := string(e.Identity())
 
 	f := sync.OnceValue(func() IndexFuture {
