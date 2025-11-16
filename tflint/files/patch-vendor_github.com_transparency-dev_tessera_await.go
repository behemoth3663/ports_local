--- vendor/github.com/transparency-dev/tessera/await.go.orig	2025-09-22 08:57:07 UTC
+++ vendor/github.com/transparency-dev/tessera/await.go
@@ -69,9 +69,6 @@ func (a *PublicationAwaiter) Await(ctx context.Context
 // or in the event that there is an error getting a valid checkpoint, an error
 // will be returned from this method.
 func (a *PublicationAwaiter) Await(ctx context.Context, future IndexFuture) (Index, []byte, error) {
-	_, span := tracer.Start(ctx, "tessera.Await")
-	defer span.End()
-
 	i, err := future()
 	if err != nil {
 		return i, nil, err
