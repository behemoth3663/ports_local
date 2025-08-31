--- vendor/github.com/transparency-dev/tessera/client/client.go.orig	2025-06-10 15:09:26 UTC
+++ vendor/github.com/transparency-dev/tessera/client/client.go
@@ -31,7 +31,6 @@ import (
 	"github.com/transparency-dev/merkle/rfc6962"
 	"github.com/transparency-dev/tessera/api"
 	"github.com/transparency-dev/tessera/api/layout"
-	"github.com/transparency-dev/tessera/internal/otel"
 	"golang.org/x/mod/sumdb/note"
 )
 
@@ -100,9 +99,6 @@ func FetchCheckpoint(ctx context.Context, f Checkpoint
 // FetchCheckpoint retrieves and opens a checkpoint from the log.
 // Returns both the parsed structure and the raw serialised checkpoint.
 func FetchCheckpoint(ctx context.Context, f CheckpointFetcherFunc, v note.Verifier, origin string) (*log.Checkpoint, []byte, *note.Note, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.FetchCheckpoint")
-	defer span.End()
-
 	cpRaw, err := f(ctx)
 	if err != nil {
 		return nil, nil, nil, err
@@ -117,10 +113,6 @@ func FetchRangeNodes(ctx context.Context, s uint64, f 
 // FetchRangeNodes returns the set of nodes representing the compact range covering
 // a log of size s.
 func FetchRangeNodes(ctx context.Context, s uint64, f TileFetcherFunc) ([][]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.FetchRangeNodes")
-	defer span.End()
-	span.SetAttributes(logSizeKey.Int64(otel.Clamp64(s)))
-
 	nc := newNodeCache(f, s)
 	nIDs := make([]compact.NodeID, 0, compact.RangeSize(0, s))
 	nIDs = compact.RangeNodes(0, s, nIDs)
@@ -137,11 +129,6 @@ func FetchLeafHashes(ctx context.Context, f TileFetche
 
 // FetchLeafHashes fetches N consecutive leaf hashes starting with the leaf at index first.
 func FetchLeafHashes(ctx context.Context, f TileFetcherFunc, first, N, logSize uint64) ([][]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.FetchLeafHashes")
-	defer span.End()
-
-	span.SetAttributes(firstKey.Int64(otel.Clamp64(first)), NKey.Int64(otel.Clamp64(N)), logSizeKey.Int64(otel.Clamp64(logSize)))
-
 	nc := newNodeCache(f, logSize)
 	hashes := make([][]byte, 0, N)
 	for i, end := first, first+N; i < end; i++ {
@@ -157,11 +144,6 @@ func GetEntryBundle(ctx context.Context, f EntryBundle
 
 // GetEntryBundle fetches the entry bundle at the given _tile index_.
 func GetEntryBundle(ctx context.Context, f EntryBundleFetcherFunc, i, logSize uint64) (api.EntryBundle, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.GetEntryBundle")
-	defer span.End()
-
-	span.SetAttributes(indexKey.Int64(otel.Clamp64(i)), logSizeKey.Int64(otel.Clamp64(logSize)))
-
 	bundle := api.EntryBundle{}
 	p := layout.PartialTileSize(0, i, logSize)
 	sRaw, err := f(ctx, i, p)
@@ -210,11 +192,6 @@ func (pb *ProofBuilder) InclusionProof(ctx context.Con
 // This function uses the passed-in function to retrieve tiles containing any log tree
 // nodes necessary to build the proof.
 func (pb *ProofBuilder) InclusionProof(ctx context.Context, index uint64) ([][]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.InclusionProof")
-	defer span.End()
-
-	span.SetAttributes(indexKey.Int64(otel.Clamp64(index)))
-
 	nodes, err := proof.Inclusion(index, pb.treeSize)
 	if err != nil {
 		return nil, fmt.Errorf("failed to calculate inclusion proof node list: %v", err)
@@ -226,10 +203,6 @@ func (pb *ProofBuilder) ConsistencyProof(ctx context.C
 // This function uses the passed-in function to retrieve tiles containing any log tree
 // nodes necessary to build the proof.
 func (pb *ProofBuilder) ConsistencyProof(ctx context.Context, smaller, larger uint64) ([][]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.ConsistencyProof")
-	defer span.End()
-	span.SetAttributes(smallerKey.Int64(otel.Clamp64(smaller)), largerKey.Int64(otel.Clamp64(larger)))
-
 	if m := max(smaller, larger); m > pb.treeSize {
 		return nil, fmt.Errorf("requested consistency proof to %d which is larger than tree size %d", m, pb.treeSize)
 	}
@@ -316,9 +289,6 @@ func (lst *LogStateTracker) Update(ctx context.Context
 // If the LatestConsistent checkpoint is 0 sized, no consistency proof will be returned
 // since it would be meaningless to do so.
 func (lst *LogStateTracker) Update(ctx context.Context) ([]byte, [][]byte, []byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.logstatetracker.Update")
-	defer span.End()
-
 	c, cRaw, _, err := lst.consensusCheckpoint(ctx, lst.cpSigVerifier, lst.origin)
 	if err != nil {
 		return nil, nil, nil, err
@@ -398,11 +368,6 @@ func (n *nodeCache) GetNode(ctx context.Context, id co
 // the tile containing the requested node will be fetched and cached, and the
 // node hash returned.
 func (n *nodeCache) GetNode(ctx context.Context, id compact.NodeID) ([]byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.client.nodecache.GetNode")
-	defer span.End()
-
-	span.SetAttributes(indexKey.Int64(otel.Clamp64(id.Index)), levelKey.Int64(int64(id.Level)))
-
 	// First check for ephemeral nodes:
 	if e := n.ephemeral[id]; len(e) != 0 {
 		return e, nil
@@ -412,7 +377,6 @@ func (n *nodeCache) GetNode(ctx context.Context, id co
 	tKey := tileKey{tileLevel, tileIndex}
 	t, ok := n.tiles[tKey]
 	if !ok {
-		span.AddEvent("cache miss")
 		p := layout.PartialTileSize(tileLevel, tileIndex, n.logSize)
 		tileRaw, err := fetchPartialOrFullTile(ctx, n.getTile, tileLevel, tileIndex, p)
 		if err != nil {
