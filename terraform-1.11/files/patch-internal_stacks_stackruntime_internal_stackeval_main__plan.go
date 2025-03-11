--- internal/stacks/stackruntime/internal/stackeval/main_plan.go.orig	2024-10-16 12:28:59 UTC
+++ internal/stacks/stackruntime/internal/stackeval/main_plan.go
@@ -176,9 +176,6 @@ func (m *Main) walkPlanObjectChanges(ctx context.Conte
 // contribute changes to the plan, which should each implement [Plannable].
 func (m *Main) walkPlanObjectChanges(ctx context.Context, walk *planWalk, obj Plannable) {
 	walk.AsyncTask(ctx, func(ctx context.Context) {
-		ctx, span := tracer.Start(ctx, obj.tracingName()+" planning")
-		defer span.End()
-
 		changes, diags := obj.PlanChanges(ctx)
 		for _, change := range changes {
 			walk.out.AnnouncePlannedChange(ctx, change)
