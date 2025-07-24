--- vendor/github.com/transparency-dev/tessera/storage/internal/integrate.go.orig	2025-06-10 15:09:26 UTC
+++ vendor/github.com/transparency-dev/tessera/storage/internal/integrate.go
@@ -24,7 +24,6 @@ import (
 	"github.com/transparency-dev/merkle/rfc6962"
 	"github.com/transparency-dev/tessera/api"
 	"github.com/transparency-dev/tessera/api/layout"
-	"github.com/transparency-dev/tessera/internal/otel"
 	"golang.org/x/exp/maps"
 	"k8s.io/klog/v2"
 )
@@ -100,11 +99,6 @@ func (t *treeBuilder) integrate(ctx context.Context, f
 }
 
 func (t *treeBuilder) integrate(ctx context.Context, fromSize uint64, leafHashes [][]byte) (newSize uint64, rootHash []byte, tiles map[TileID]*api.HashTile, err error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.integrate")
-	defer span.End()
-
-	span.SetAttributes(fromSizeKey.Int64(otel.Clamp64(fromSize)), numEntriesKey.Int(len(leafHashes)))
-
 	baseRange, err := t.newRange(ctx, fromSize)
 	if err != nil {
 		return 0, nil, nil, fmt.Errorf("failed to create range covering existing log: %w", err)
@@ -126,7 +120,6 @@ func (t *treeBuilder) integrate(ctx context.Context, f
 		return fromSize, r, nil, nil
 	}
 
-	span.AddEvent("Loaded state")
 	klog.V(1).Infof("Loaded state with roothash %x", r)
 	// Create a new compact range which represents the update to the tree
 	newRange := t.rf.NewEmptyRange(fromSize)
@@ -143,7 +136,6 @@ func (t *treeBuilder) integrate(ctx context.Context, f
 	if err := tc.Err(); err != nil {
 		return 0, nil, nil, err
 	}
-	span.AddEvent("Updated tile cache")
 
 	// Merge the update range into the old tree
 	if err := baseRange.AppendRange(newRange, visitor); err != nil {
@@ -162,8 +154,6 @@ func (t *treeBuilder) integrate(ctx context.Context, f
 		return 0, nil, nil, fmt.Errorf("failed to calculate new root hash: %w", err)
 	}
 
-	span.AddEvent("Calculated new root")
-
 	// All calculation is now complete, all that remains is to store the new
 	// tiles and updated log state.
 	klog.V(1).Infof("New log state: size 0x%x hash: %x", baseRange.End(), newRoot)
@@ -187,16 +177,10 @@ func (r *tileReadCache) Get(ctx context.Context, tileI
 
 // Get returns a previously set tile and true, or, if no such tile is in the cache, attempt to fetch it.
 func (r *tileReadCache) Get(ctx context.Context, tileID TileID, treeSize uint64) (*populatedTile, error) {
-	ctx, span := tracer.Start(ctx, "tessera.storage.readCache.Get")
-	defer span.End()
-
-	span.SetAttributes(indexKey.Int64(otel.Clamp64(tileID.Index)), levelKey.Int64(otel.Clamp64(tileID.Level)), treeSizeKey.Int64(otel.Clamp64(treeSize)))
-
 	k := layout.TilePath(uint64(tileID.Level), tileID.Index, layout.PartialTileSize(tileID.Level, tileID.Index, treeSize))
 	e, ok := r.entries[k]
 	if !ok {
 		klog.V(1).Infof("Readcache miss: %q", k)
-		span.AddEvent(fmt.Sprintf("Cache miss %q", k))
 		t, err := r.getTiles(ctx, []TileID{tileID}, treeSize)
 		if err != nil {
 			return nil, err
@@ -214,9 +198,6 @@ func (r *tileReadCache) Prewarm(ctx context.Context, t
 //
 // Returns an error if any of the tiles couldn't be fetched.
 func (r *tileReadCache) Prewarm(ctx context.Context, tileIDs []TileID, treeSize uint64) error {
-	ctx, span := tracer.Start(ctx, "tessera.storage.readCache.Prewarm")
-	defer span.End()
-
 	t, err := r.getTiles(ctx, tileIDs, treeSize)
 	if err != nil {
 		return err
