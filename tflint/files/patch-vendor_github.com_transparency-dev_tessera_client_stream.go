--- vendor/github.com/transparency-dev/tessera/client/stream.go.orig	2025-09-22 08:57:07 UTC
+++ vendor/github.com/transparency-dev/tessera/client/stream.go
@@ -48,9 +48,6 @@ func EntryBundles(ctx context.Context, numWorkers uint
 // to balance throughput against consumption of resources, but such balancing needs to be mindful of the nature of the
 // source infrastructure, and how concurrent requests affect performance (e.g. GCS buckets vs. files on a single disk).
 func EntryBundles(ctx context.Context, numWorkers uint, getSize TreeSizeFunc, getBundle EntryBundleFetcherFunc, fromEntry uint64, N uint64) iter.Seq2[Bundle, error] {
-	ctx, span := tracer.Start(ctx, "tessera.storage.StreamAdaptor")
-	defer span.End()
-
 	// bundleOrErr represents a fetched entry bundle and its params, or an error if we couldn't fetch it for
 	// some reason.
 	type bundleOrErr struct {
@@ -69,9 +66,6 @@ func EntryBundles(ctx context.Context, numWorkers uint
 	// We use a limited number of tokens here to prevent this from
 	// consuming an unbounded amount of resources.
 	go func() {
-		ctx, span := tracer.Start(ctx, "tessera.storage.StreamAdaptorWorker")
-		defer span.End()
-
 		defer close(bundles)
 
 		treeSize, err := getSize(ctx)
