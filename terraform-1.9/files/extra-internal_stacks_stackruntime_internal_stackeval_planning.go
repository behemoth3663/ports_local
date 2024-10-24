--- internal/stacks/stackruntime/internal/stackeval/planning.go.orig	2024-10-16 12:28:59 UTC
+++ internal/stacks/stackruntime/internal/stackeval/planning.go
@@ -52,8 +52,4 @@ type Plannable interface {
 	// so it's unnecessary and harmful to for one object to try to handle
 	// planning (or plan-time validation) on behalf of some other object.
 	PlanChanges(ctx context.Context) ([]stackplan.PlannedChange, tfdiags.Diagnostics)
-
-	// Our general async planning helper relies on this to name its
-	// tracing span.
-	tracingNamer
 }
