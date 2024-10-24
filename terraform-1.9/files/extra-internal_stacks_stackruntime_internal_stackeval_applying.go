--- internal/stacks/stackruntime/internal/stackeval/applying.go.orig	2024-10-16 12:28:59 UTC
+++ internal/stacks/stackruntime/internal/stackeval/applying.go
@@ -60,8 +60,4 @@ type Applyable interface {
 	// evaluating other objects. Those will be collected separately by calling
 	// this same method on those other objects.
 	CheckApply(ctx context.Context) ([]stackstate.AppliedChange, tfdiags.Diagnostics)
-
-	// Our general async planning helper relies on this to name its
-	// tracing span.
-	tracingNamer
 }
