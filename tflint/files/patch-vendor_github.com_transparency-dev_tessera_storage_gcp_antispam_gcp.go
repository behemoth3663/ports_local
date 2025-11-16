--- vendor/github.com/transparency-dev/tessera/storage/gcp/antispam/gcp.go.orig	2025-09-22 08:57:07 UTC
+++ vendor/github.com/transparency-dev/tessera/storage/gcp/antispam/gcp.go
@@ -128,14 +128,10 @@ func (d *AntispamStorage) index(ctx context.Context, h
 
 // index returns the index (if any) previously associated with the provided hash
 func (d *AntispamStorage) index(ctx context.Context, h []byte) (*uint64, error) {
-	ctx, span := tracer.Start(ctx, "tessera.antispam.gcp.index")
-	defer span.End()
-
 	d.numLookups.Add(1)
 	var idx int64
 	if row, err := d.dbPool.Single().ReadRow(ctx, "IDSeq", spanner.Key{h}, []string{"idx"}); err != nil {
 		if c := spanner.ErrCode(err); c == codes.NotFound {
-			span.AddEvent("tessera.miss")
 			return nil, nil
 		}
 		return nil, err
@@ -144,7 +140,6 @@ func (d *AntispamStorage) index(ctx context.Context, h
 			return nil, fmt.Errorf("failed to read antispam index: %v", err)
 		}
 		idx := uint64(idx)
-		span.AddEvent("tessera.hit")
 		d.numHits.Add(1)
 		return &idx, nil
 	}
@@ -155,11 +150,7 @@ func (d *AntispamStorage) Decorator() func(f tessera.A
 func (d *AntispamStorage) Decorator() func(f tessera.AddFn) tessera.AddFn {
 	return func(delegate tessera.AddFn) tessera.AddFn {
 		return func(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
-			ctx, span := tracer.Start(ctx, "tessera.antispam.gcp.Add")
-			defer span.End()
-
 			if d.pushBack.Load() {
-				span.AddEvent("tessera.pushback")
 				// The follower is too far behind the currently integrated tree, so we're going to push back against
 				// the incoming requests.
 				// This should have two effects:
@@ -258,9 +249,6 @@ func (f *follower) Follow(ctx context.Context, lr tess
 		// Busy loop while there are entries to be consumed from the stream
 		for streamDone := false; !streamDone; {
 			_, err := f.as.dbPool.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
-				ctx, span := tracer.Start(ctx, "tessera.antispam.gcp.FollowTxn")
-				defer span.End()
-
 				// Figure out the last entry we used to populate our antispam storage.
 				row, err := txn.ReadRowWithOptions(ctx, "FollowCoord", spanner.Key{0}, []string{"nextIdx"}, &spanner.ReadOptions{LockHint: spannerpb.ReadRequest_LOCK_HINT_EXCLUSIVE})
 				if err != nil {
@@ -271,7 +259,6 @@ func (f *follower) Follow(ctx context.Context, lr tess
 				if err := row.Columns(&nextIdx); err != nil {
 					return fmt.Errorf("failed to read follow coordination info: %v", err)
 				}
-				span.SetAttributes(followFromKey.Int64(nextIdx))
 
 				followFrom := uint64(nextIdx)
 				if followFrom >= logSize {
@@ -296,13 +283,11 @@ func (f *follower) Follow(ctx context.Context, lr tess
 				}
 
 				pushback := logSize-followFrom > uint64(f.as.opts.PushbackThreshold)
-				span.SetAttributes(pushbackKey.Bool(pushback))
 				f.as.pushBack.Store(pushback)
 
 				// If this is the first time around the loop we need to start the stream of entries now that we know where we want to
 				// start reading from:
 				if next == nil {
-					span.AddEvent("Start streaming entries")
 					sizeFn := func(_ context.Context) (uint64, error) {
 						return logSize, nil
 					}
@@ -400,9 +385,6 @@ func (f *follower) batchUpdateIndex(ctx context.Contex
 //     This would work, but would also incur an extra round-trip of data which isn't really necessary but would
 //     slow the process down considerably and add extra load to Spanner for no benefit.
 func (f *follower) batchUpdateIndex(ctx context.Context, _ *spanner.ReadWriteTransaction, ms []*spanner.Mutation) error {
-	ctx, span := tracer.Start(ctx, "tessera.antispam.gcp.batchUpdateIndex")
-	defer span.End()
-
 	mgs := make([]*spanner.MutationGroup, 0, len(ms))
 	for _, m := range ms {
 		mgs = append(mgs, &spanner.MutationGroup{
