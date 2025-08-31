--- vendor/github.com/transparency-dev/tessera/storage/gcp/gcp.go.orig	2025-06-10 15:09:26 UTC
+++ vendor/github.com/transparency-dev/tessera/storage/gcp/gcp.go
@@ -54,7 +54,6 @@ import (
 	"github.com/transparency-dev/tessera/api"
 	"github.com/transparency-dev/tessera/api/layout"
 	"github.com/transparency-dev/tessera/internal/migrate"
-	"github.com/transparency-dev/tessera/internal/otel"
 	"github.com/transparency-dev/tessera/internal/parse"
 	"github.com/transparency-dev/tessera/internal/stream"
 	storage "github.com/transparency-dev/tessera/storage/internal"
@@ -147,9 +146,6 @@ func (lr *LogReader) ReadCheckpoint(ctx context.Contex
 }
 
 func (lr *LogReader) ReadCheckpoint(ctx context.Context) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.ReadCheckpoint")
-	defer span.End()
-
 	r, err := lr.lrs.getCheckpoint(ctx)
 	if err != nil {
 		if errors.Is(err, gcs.ErrObjectNotExist) {
@@ -160,37 +156,22 @@ func (lr *LogReader) ReadTile(ctx context.Context, l, 
 }
 
 func (lr *LogReader) ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.ReadTile")
-	defer span.End()
-
 	return lr.lrs.getTile(ctx, l, i, p)
 }
 
 func (lr *LogReader) ReadEntryBundle(ctx context.Context, i uint64, p uint8) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.ReadEntryBundle")
-	defer span.End()
-
 	return lr.lrs.getEntryBundle(ctx, i, p)
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
 
 func (lr *LogReader) StreamEntries(ctx context.Context, startEntry, N uint64) iter.Seq2[stream.Bundle, error] {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.StreamEntries")
-	defer span.End()
-
 	klog.Infof("StreamEntries from %d", startEntry)
 
 	// TODO(al): Consider making this configurable.
@@ -280,9 +261,6 @@ func (a *Appender) Add(ctx context.Context, e *tessera
 
 // Add is the entrypoint for adding entries to a sequencing log.
 func (a *Appender) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.Add")
-	defer span.End()
-
 	return a.queue.Add(ctx, e)
 }
 
@@ -300,9 +278,6 @@ func (a *Appender) integrateEntriesJob(ctx context.Con
 		}
 
 		func() {
-			ctx, span := tracer.Start(ctx, "tessera.storage.gcp.integrateEntriesJob")
-			defer span.End()
-
 			// Don't quickloop for now, it causes issues updating checkpoint too frequently.
 			cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
 			defer cancel()
@@ -334,8 +309,6 @@ func (a *Appender) publishCheckpointJob(ctx context.Co
 		case <-t.C:
 		}
 		func() {
-			ctx, span := tracer.Start(ctx, "tessera.storage.gcp.publishCheckpointJob")
-			defer span.End()
 			if err := a.sequencer.publishCheckpoint(ctx, i, a.publishCheckpoint); err != nil {
 				klog.Warningf("publishCheckpoint failed: %v", err)
 			}
@@ -360,9 +333,6 @@ func (a *Appender) garbageCollectorJob(ctx context.Con
 		case <-t.C:
 		}
 		func() {
-			ctx, span := tracer.Start(ctx, "tessera.storage.gcp.garbageCollectTask")
-			defer span.End()
-
 			// Figure out the size of the latest published checkpoint - we can't be removing partial tiles implied by
 			// that checkpoint just because we've done an integration and know about a larger (but as yet unpublished)
 			// checkpoint!
@@ -408,10 +378,6 @@ func (a *Appender) publishCheckpoint(ctx context.Conte
 }
 
 func (a *Appender) publishCheckpoint(ctx context.Context, size uint64, root []byte) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.publishCheckpoint")
-	defer span.End()
-	span.SetAttributes(treeSizeKey.Int64(otel.Clamp64(size)))
-
 	cpRaw, err := a.newCP(ctx, size, root)
 	if err != nil {
 		return fmt.Errorf("newCP: %v", err)
@@ -470,9 +436,6 @@ func (s *logResourceStore) getTiles(ctx context.Contex
 //
 // Tiles are returned in the same order as they're requested, nils represent tiles which were not found.
 func (s *logResourceStore) getTiles(ctx context.Context, tileIDs []storage.TileID, logSize uint64) ([]*api.HashTile, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.getTiles")
-	defer span.End()
-
 	r := make([]*api.HashTile, len(tileIDs))
 	errG := errgroup.Group{}
 	for i, id := range tileIDs {
@@ -539,9 +502,6 @@ func (a *Appender) integrateEntries(ctx context.Contex
 //
 // Returns the new root hash of the log with the entries added.
 func (a *Appender) integrateEntries(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.integrateEntries")
-	defer span.End()
-
 	var newRoot []byte
 
 	errG := errgroup.Group{}
@@ -573,11 +533,6 @@ func integrate(ctx context.Context, fromSeq uint64, lh
 
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
@@ -614,9 +569,6 @@ func (a *Appender) updateEntryBundles(ctx context.Cont
 //
 // The right-most bundle will be grown, if it's partial, and/or new bundles will be created as required.
 func (a *Appender) updateEntryBundles(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.updateEntryBundles")
-	defer span.End()
-
 	if len(entries) == 0 {
 		return nil
 	}
@@ -763,11 +715,6 @@ func (s *spannerCoordinator) assignEntries(ctx context
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
@@ -778,7 +725,6 @@ func (s *spannerCoordinator) assignEntries(ctx context
 			return fmt.Errorf("failed to read integration coordination info: %v", err)
 		}
 	}
-	span.SetAttributes(treeSizeKey.Int64(treeSize))
 
 	var next int64 // Unfortunately, Spanner doesn't support uint64 so we'll have to cast around a bit.
 
@@ -847,9 +793,6 @@ func (s *spannerCoordinator) consumeEntries(ctx contex
 //
 // Returns true if some entries were consumed as a weak signal that there may be further entries waiting to be consumed.
 func (s *spannerCoordinator) consumeEntries(ctx context.Context, limit uint64, f consumeFunc, forceUpdate bool) (bool, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.gcp.consumeEntries")
-	defer span.End()
-
 	didWork := false
 	_, err := s.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
 		// Figure out which is the starting index of sequenced entries to start consuming from.
@@ -1108,15 +1051,10 @@ func (s *gcsStorage) getObject(ctx context.Context, ob
 
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
@@ -1138,15 +1076,10 @@ func (s *gcsStorage) setObject(ctx context.Context, ob
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
 
@@ -1173,12 +1106,10 @@ func (s *gcsStorage) setObject(ctx context.Context, ob
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
@@ -1190,13 +1121,9 @@ func (s *gcsStorage) deleteObjectsWithPrefix(ctx conte
 
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
 
