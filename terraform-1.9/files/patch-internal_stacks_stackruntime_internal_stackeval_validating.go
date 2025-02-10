--- internal/stacks/stackruntime/internal/stackeval/validating.go.orig	2024-10-16 12:28:59 UTC
+++ internal/stacks/stackruntime/internal/stackeval/validating.go
@@ -30,8 +30,4 @@ type Validatable interface {
 	// so it's unnecessary and harmful to try to handle validation on behalf of
 	// some other related object.
 	Validate(ctx context.Context) tfdiags.Diagnostics
-
-	// Our general async validation helper relies on this to name its
-	// tracing span.
-	tracingNamer
 }
