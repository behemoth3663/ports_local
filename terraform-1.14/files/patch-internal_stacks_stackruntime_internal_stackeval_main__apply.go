--- internal/stacks/stackruntime/internal/stackeval/main_apply.go.orig	2025-06-11 10:22:08 UTC
+++ internal/stacks/stackruntime/internal/stackeval/main_apply.go
@@ -9,9 +9,6 @@ import (
 	"log"
 	"sync/atomic"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/protobuf/types/known/anypb"
 
 	"github.com/hashicorp/terraform/internal/addrs"
@@ -87,8 +84,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 				reg.RegisterComponentInstanceChange(
 					ctx, addr,
 					func(ctx context.Context, main *Main) (*ComponentInstanceApplyResult, tfdiags.Diagnostics) {
-						ctx, span := tracer.Start(ctx, addr.String()+" apply")
-						defer span.End()
 						log.Printf("[TRACE] stackeval: %s preparing to apply", addr)
 
 						stack := main.Stack(ctx, addr.Stack, ApplyPhase)
@@ -129,7 +124,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 											fmt.Sprintf("There was a plan for %s, which matched both removed and component blocks.", addr),
 										))
 										log.Printf("[ERROR] stackeval: %s has both a component and a removed block that point to the same address", addr)
-										span.SetStatus(codes.Error, "both component and removed block present")
 										return nil, diags
 									}
 									inst = i
@@ -145,7 +139,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 								fmt.Sprintf("There was a plan for %s, which matched no known component instances.", addr),
 							))
 							log.Printf("[ERROR] stackeval: %s has planned changes, but no instance to apply them to", addr)
-							span.SetStatus(codes.Error, "no instance to apply changes to")
 							return nil, diags
 						}
 
@@ -164,13 +157,11 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 								fmt.Sprintf("The plan for %s is inconsistent with its prior state: %s.", addr, err),
 							))
 							log.Printf("[ERROR] stackeval: %s has a plan inconsistent with its prior state: %s", addr, err)
-							span.SetStatus(codes.Error, "plan is inconsistent with prior state")
 							return nil, diags
 						}
 
 						if matchedUnknownBlock && inst == nil {
 							log.Printf("[TRACE] stackeval: %s matched an unknown block so returning prior state unchanged.", addr)
-							span.SetStatus(codes.Ok, "apply unknown")
 
 							for _, step := range addr.Stack {
 								if step.Key == addrs.WildcardKey {
@@ -225,18 +216,11 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 						for waitComponentAddr := range waitForComponents.All() {
 							if stack := main.Stack(ctx, waitComponentAddr.Stack, ApplyPhase); stack != nil {
 								if component := stack.Component(waitComponentAddr.Item); component != nil {
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
@@ -251,9 +235,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 						for waitComponentAddr := range waitForRemoveds.All() {
 							if stack := main.Stack(ctx, waitComponentAddr.Stack, ApplyPhase); stack != nil {
 								if removed := stack.RemovedComponent(waitComponentAddr.Item); removed != nil {
-									span.AddEvent("awaiting predecessor", trace.WithAttributes(
-										attribute.String("component_addr", waitComponentAddr.String()),
-									))
 									success := true
 									for _, block := range removed {
 										if !block.ApplySuccessful(ctx) {
@@ -264,10 +245,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
 										// If anything we're waiting on does not succeed then we can't proceed without
 										// violating the dependency invariants.
 										log.Printf("[TRACE] stackeval: %s cannot start because %s changes did not apply completely", addr, waitComponentAddr)
-										span.AddEvent("predecessor is incomplete", trace.WithAttributes(
-											attribute.String("component_addr", waitComponentAddr.String()),
-										))
-										span.SetStatus(codes.Error, "predecessors did not completely apply")
 
 										// We'll return a stub result that reports that nothing was changed, since
 										// we're not going to run our apply phase at all.
@@ -282,11 +259,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
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
@@ -297,12 +269,6 @@ func ApplyPlan(ctx context.Context, config *stackconfi
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
@@ -407,9 +373,6 @@ func (m *Main) walkApplyCheckObjectChanges(ctx context
 // deals with changes.)
 func (m *Main) walkApplyCheckObjectChanges(ctx context.Context, walk *applyWalk, obj Applyable) {
 	walk.AsyncTask(ctx, func(ctx context.Context) {
-		ctx, span := tracer.Start(ctx, obj.tracingName()+" apply-time checks")
-		defer span.End()
-
 		changes, diags := obj.CheckApply(ctx)
 		for _, change := range changes {
 			walk.out.AnnounceAppliedChange(ctx, change)
