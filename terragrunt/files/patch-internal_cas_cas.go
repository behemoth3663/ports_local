--- internal/cas/cas.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/cas.go
@@ -752,7 +752,6 @@ func (c *CAS) populateTreeFromSymbolicRef(
 		ref.URL,
 		err,
 	)
-	RecordFallback(ctx, l, FallbackReasonGitStoreUnavailable, map[string]any{"url": ref.URL})
 
 	tempDir, cleanup, err := c.makeFallbackCloneDir(l, v)
 	if err != nil {
@@ -817,7 +816,6 @@ func (c *CAS) populateTreeFromCommitRef(
 		ref.URL,
 		err,
 	)
-	RecordFallback(ctx, l, FallbackReasonGitStoreUnavailable, map[string]any{"url": ref.URL})
 
 	tempDir, cleanup, err := c.makeFallbackCloneDir(l, v)
 	if err != nil {
