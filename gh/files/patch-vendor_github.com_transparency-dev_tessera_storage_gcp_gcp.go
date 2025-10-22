--- vendor/github.com/transparency-dev/tessera/storage/gcp/gcp.go.orig	2025-09-22 08:57:07 UTC
+++ vendor/github.com/transparency-dev/tessera/storage/gcp/gcp.go
@@ -54,7 +54,6 @@ import (
 	"github.com/transparency-dev/tessera/api/layout"
 	"github.com/transparency-dev/tessera/internal/fetcher"
 	"github.com/transparency-dev/tessera/internal/migrate"
-	"github.com/transparency-dev/tessera/internal/otel"
 	"github.com/transparency-dev/tessera/internal/parse"
 	storage "github.com/transparency-dev/tessera/storage/internal"
 	"golang.org/x/sync/errgroup"
@@ -157,9 +156,6 @@ func (lr *LogReader) ReadCheckpoint(ctx context.Contex
 }
 
 func (lr *LogReader) ReadCheckpoint(ctx context.Context) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.ReadCheckpoint")
-	defer span.End()
-
 	r, err := lr.lrs.getCheckpoint(ctx)
 	if err != nil {
 		if errors.Is(err, gcs.ErrObjectNotExist) {
@@ -170,34 +166,22 @@ func (lr *LogReader) ReadTile(ctx context.Context, l, 
 }
 
 func (lr *LogReader) ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.ReadTile")
-	defer span.End()
-
 	return fetcher.PartialOrFullResource(ctx, p, func(ctx context.Context, p uint8) ([]byte, error) {
 		return lr.lrs.getTile(ctx, l, i, p)
 	})
 }
 
 func (lr *LogReader) ReadEntryBundle(ctx context.Context, i uint64, p uint8) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.ReadEntryBundle")
-	defer span.End()
-
 	return fetcher.PartialOrFullResource(ctx, p, func(ctx context.Context, p uint8) ([]byte, error) {
 		return lr.lrs.getEntryBundle(ctx, i, p)
 	})
 }
 
 func (lr *LogReader) IntegratedSize(ctx context.Context) (uint64, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.IntegratedSize")
-	defer span.End()
-
 	return lr.integratedSize(ctx)
 }
 
 func (lr *LogReader) NextIndex(ctx context.Context) (uint64, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.NextIndex")
-	defer span.End()
-
 	return lr.nextIndex(ctx)
 }
 
@@ -296,9 +280,6 @@ func (a *Appender) Add(ctx context.Context, e *tessera
 
 // Add is the entrypoint for adding entries to a sequencing log.
 func (a *Appender) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.Add")
-	defer span.End()
-
 	return a.queue.Add(ctx, e)
 }
 
@@ -316,9 +297,6 @@ func (a *Appender) integrateEntriesJob(ctx context.Con
 		}
 
 		func() {
-			ctx, span := tracer.Start(ctx, "tessera.storage.gcp.integrateEntriesJob")
-			defer span.End()
-
 			// Don't quickloop for now, it causes issues updating checkpoint too frequently.
 			cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
 			defer cancel()
@@ -350,8 +328,6 @@ func (a *Appender) publishCheckpointJob(ctx context.Co
 		case <-t.C:
 		}
 		func() {
-			ctx, span := tracer.Start(ctx, "tessera.storage.gcp.publishCheckpointJob")
-			defer span.End()
 			if err := a.sequencer.publishCheckpoint(ctx, i, a.publishCheckpoint); err != nil {
 				klog.Warningf("publishCheckpoint failed: %v", err)
 			}
@@ -376,9 +352,6 @@ func (a *Appender) garbageCollectorJob(ctx context.Con
 		case <-t.C:
 		}
 		func() {
-			ctx, span := tracer.Start(ctx, "tessera.storage.gcp.garbageCollectTask")
-			defer span.End()
-
 			// Figure out the size of the latest published checkpoint - we can't be removing partial tiles implied by
 			// that checkpoint just because we've done an integration and know about a larger (but as yet unpublished)
 			// checkpoint!
@@ -427,10 +400,6 @@ func (a *Appender) publishCheckpoint(ctx context.Conte
 }
 
 func (a *Appender) publishCheckpoint(ctx context.Context, size uint64, root []byte) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.publishCheckpoint")
-	defer span.End()
-	span.SetAttributes(treeSizeKey.Int64(otel.Clamp64(size)))
-
 	cpRaw, err := a.newCP(ctx, size, root)
 	if err != nil {
 		return fmt.Errorf("newCP: %v", err)
@@ -489,9 +458,6 @@ func (s *logResourceStore) getTiles(ctx context.Contex
 //
 // Tiles are returned in the same order as they're requested, nils represent tiles which were not found.
 func (s *logResourceStore) getTiles(ctx context.Context, tileIDs []storage.TileID, logSize uint64) ([]*api.HashTile, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.getTiles")
-	defer span.End()
-
 	r := make([]*api.HashTile, len(tileIDs))
 	errG := errgroup.Group{}
 	for i, id := range tileIDs {
@@ -558,9 +524,6 @@ func (a *Appender) integrateEntries(ctx context.Contex
 //
 // Returns the new root hash of the log with the entries added.
 func (a *Appender) integrateEntries(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.integrateEntries")
-	defer span.End()
-
 	var newRoot []byte
 
 	errG := errgroup.Group{}
@@ -592,11 +555,6 @@ func integrate(ctx context.Context, fromSeq uint64, lh
 
 // integrate adds the provided leaf hashes to the merkle tree, starting at the provided location.
 func integrate(ctx context.Context, fromSeq uint64, lh [][]byte, logStore *logResourceStore) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.integrate")
-	defer span.End()
-
-	span.SetAttributes(fromSizeKey.Int64(otel.Clamp64(fromSeq)), numEntriesKey.Int(len(lh)))
-
 	errG := errgroup.Group{}
 	getTiles := func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
 		n, err := logStore.getTiles(ctx, tileIDs, treeSize)
@@ -633,9 +591,6 @@ func (a *Appender) updateEntryBundles(ctx context.Cont
 //
 // The right-most bundle will be grown, if it's partial, and/or new bundles will be created as required.
 func (a *Appender) updateEntryBundles(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.updateEntryBundles")
-	defer span.End()
-
 	if len(entries) == 0 {
 		return nil
 	}
@@ -775,11 +730,6 @@ func (s *spannerCoordinator) assignEntries(ctx context
 // This is achieved by storing the passed-in entries in the Seq table in Spanner, keyed by the
 // index assigned to the first entry in the batch.
 func (s *spannerCoordinator) assignEntries(ctx context.Context, entries []*tessera.Entry) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.assignEntries")
-	defer span.End()
-
-	span.SetAttributes(numEntriesKey.Int(len(entries)))
-
 	// First grab the treeSize in a non-locking read-only fashion (we don't want to block/collide with integration).
 	// We'll use this value to determine whether we need to apply back-pressure.
 	var treeSize int64
@@ -790,7 +740,6 @@ func (s *spannerCoordinator) assignEntries(ctx context
 			return fmt.Errorf("failed to read integration coordination info: %v", err)
 		}
 	}
-	span.SetAttributes(treeSizeKey.Int64(treeSize))
 
 	var next int64 // Unfortunately, Spanner doesn't support uint64 so we'll have to cast around a bit.
 
@@ -859,9 +808,6 @@ func (s *spannerCoordinator) consumeEntries(ctx contex
 //
 // Returns true if some entries were consumed as a weak signal that there may be further entries waiting to be consumed.
 func (s *spannerCoordinator) consumeEntries(ctx context.Context, limit uint64, f consumeFunc, forceUpdate bool) (bool, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.consumeEntries")
-	defer span.End()
-
 	didWork := false
 	_, err := s.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
 		// Figure out which is the starting index of sequenced entries to start consuming from.
@@ -1120,15 +1066,10 @@ func (s *gcsStorage) getObject(ctx context.Context, ob
 
 // getObject returns the data and generation of the specified object, or an error.
 func (s *gcsStorage) getObject(ctx context.Context, obj string) ([]byte, int64, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.getObject")
-	defer span.End()
-
 	if s.bucketPrefix != "" {
 		obj = filepath.Join(s.bucketPrefix, obj)
 	}
 
-	span.SetAttributes(objectPathKey.String(obj))
-
 	r, err := s.gcsClient.Bucket(s.bucket).Object(obj).NewReader(ctx)
 	if err != nil {
 		return nil, -1, fmt.Errorf("getObject: failed to create reader for object %q in bucket %q: %w", obj, s.bucket, err)
@@ -1150,15 +1091,10 @@ func (s *gcsStorage) setObject(ctx context.Context, ob
 // the currently stored data is bit-for-bit identical to the data to-be-written.
 // This is intended to provide idempotentency for writes.
 func (s *gcsStorage) setObject(ctx context.Context, objName string, data []byte, cond *gcs.Conditions, contType string, cacheCtl string) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.setObject")
-	defer span.End()
-
 	if s.bucketPrefix != "" {
 		objName = filepath.Join(s.bucketPrefix, objName)
 	}
 
-	span.SetAttributes(objectPathKey.String(objName))
-
 	bkt := s.gcsClient.Bucket(s.bucket)
 	obj := bkt.Object(objName)
 
@@ -1196,12 +1132,10 @@ func (s *gcsStorage) setObject(ctx context.Context, ob
 				return fmt.Errorf("failed to fetch existing content for %q (@%d): %v", objName, existingGen, err)
 			}
 			if !bytes.Equal(existing, data) {
-				span.AddEvent("Non-idempotent write")
 				klog.Errorf("Resource %q non-idempotent write:\n%s", objName, cmp.Diff(existing, data))
 				return fmt.Errorf("precondition failed: resource content for %q differs from data to-be-written", objName)
 			}
 
-			span.AddEvent("Idempotent write")
 			klog.V(2).Infof("setObject: identical resource already exists for %q, continuing", objName)
 			return nil
 		}
@@ -1213,13 +1147,9 @@ func (s *gcsStorage) deleteObjectsWithPrefix(ctx conte
 
 // deleteObjectsWithPrefix removes any objects with the provided prefix from GCS.
 func (s *gcsStorage) deleteObjectsWithPrefix(ctx context.Context, objPrefix string) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.deleteObject")
-	defer span.End()
-
 	if s.bucketPrefix != "" {
 		objPrefix = filepath.Join(s.bucketPrefix, objPrefix)
 	}
-	span.SetAttributes(objectPathKey.String(objPrefix))
 
 	bkt := s.gcsClient.Bucket(s.bucket)
 
