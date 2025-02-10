--- internal/stacks/stackruntime/internal/stackeval/main_apply.go.orig	2025-02-10 14:18:24 UTC
+++ internal/stacks/stackruntime/internal/stackeval/main_apply.go
@@ -9,9 +9,6 @@ import (
 	"log"
 	"sync/atomic"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/protobuf/types/known/anypb"
 
 	"github.com/hashicorp/terraform/internal/collections"
@@ -94,8 +91,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 				reg.RegisterComponentInstanceChange(
 					ctx, addr,
 					func(ctx context.Context, main *Main) (*ComponentInstanceApplyResult, tfdiags.Diagnostics) {
-						ctx, span := tracer.Start(ctx, addr.String()+" apply")
-						defer span.End()
 						log.Printf("[TRACE] stackeval: %s preparing to apply", addr)
 
 						stack := main.Stack(ctx, addr.Stack, ApplyPhase)
@@ -114,7 +109,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 							// to be reported by some object when we do the apply
 							// checking walk below.
 							log.Printf("[ERROR] stackeval: %s has planned changes, but does not seem to be declared", addr)
-							span.SetStatus(codes.Error, "missing component instance declaration")
 							return nil, nil
 						}
 
@@ -133,7 +127,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 								fmt.Sprintf("The plan for %s is inconsistent with its prior state: %s.", addr, err),
 							))
 							log.Printf("[ERROR] stackeval: %s has a plan inconsistent with its prior state: %s", addr, err)
-							span.SetStatus(codes.Error, "plan is inconsistent with prior state")
 							return nil, diags
 						}
 
@@ -161,18 +154,11 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 						for _, waitComponentAddr := range waitForComponents.Elems() {
 							if stack := main.Stack(ctx, waitComponentAddr.Stack, ApplyPhase); stack != nil {
 								if component := stack.Component(ctx, waitComponentAddr.Item); component != nil {
-									span.AddEvent("awaiting predecessor", trace.WithAttributes(
-										attribute.String("component_addr", waitComponentAddr.String()),
-									))
 									success := component.ApplySuccessful(ctx)
 									if !success {
 										// If anything we're waiting on does not succeed then we can't proceed without
 										// violating the dependency invariants.
 										log.Printf("[TRACE] stackeval: %s cannot start because %s changes did not apply completely", addr, waitComponentAddr)
-										span.AddEvent("predecessor is incomplete", trace.WithAttributes(
-											attribute.String("component_addr", waitComponentAddr.String()),
-										))
-										span.SetStatus(codes.Error, "predecessors did not completely apply")
 
 										// We'll return a stub result that reports that nothing was changed, since
 										// we're not going to run our apply phase at all.
@@ -187,11 +173,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 						log.Printf("[TRACE] stackeval: %s now applying", addr)
 
 						ret, diags := inst.ApplyModuleTreePlan(ctx, modulesRuntimePlan)
-						if !ret.Complete {
-							span.SetStatus(codes.Error, "apply did not complete successfully")
-						} else {
-							span.SetStatus(codes.Ok, "apply complete")
-						}
 						return ret, diags
 					},
 				)
@@ -202,12 +183,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 		main.AllowLanguageExperiments(opts.ExperimentsAllowed)
 		begin(ctx, main) // the change tasks registered above become runnable
 
-		// With the planned changes now in progress, we'll visit everything and
-		// each object to check itself (producing diagnostics) and announce any
-		// changes that were applied to it.
-		ctx, span := tracer.Start(ctx, "apply-time checks")
-		defer span.End()
-
 		var seenSelfDepDiag atomic.Bool
 		ws, complete := newWalkStateCustomDiags(
 			func(diags tfdiags.Diagnostics) {
@@ -312,9 +287,6 @@ func (m *Main) walkApplyCheckObjectChanges(ctx context
 // deals with changes.)
 func (m *Main) walkApplyCheckObjectChanges(ctx context.Context, walk *applyWalk, obj Applyable) {
 	walk.AsyncTask(ctx, func(ctx context.Context) {
-		ctx, span := tracer.Start(ctx, obj.tracingName()+" apply-time checks")
-		defer span.End()
-
 		changes, diags := obj.CheckApply(ctx)
 		for _, change := range changes {
 			walk.out.AnnounceAppliedChange(ctx, change)
