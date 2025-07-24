--- vendor/github.com/transparency-dev/tessera/await.go.orig	2025-06-10 15:09:26 UTC
+++ vendor/github.com/transparency-dev/tessera/await.go
@@ -83,9 +83,6 @@ func (a *PublicationAwaiter) Await(ctx context.Context
 // or in the event that there is an error getting a valid checkpoint, an error
 // will be returned from this method.
 func (a *PublicationAwaiter) Await(ctx context.Context, future IndexFuture) (Index, []byte, error) {
-	ctx, span := tracer.Start(ctx, "tessera.Await")
-	defer span.End()
-
 	i, err := future()
 	if err != nil {
 		return i, nil, err
