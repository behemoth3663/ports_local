--- internal/rpcapi/stacks.go.orig	2025-06-11 10:22:08 UTC
+++ internal/rpcapi/stacks.go
@@ -13,9 +13,6 @@ import (
 	"github.com/hashicorp/go-slug/sourceaddrs"
 	"github.com/hashicorp/go-slug/sourcebundle"
 	"github.com/hashicorp/terraform-svchost/disco"
-	"go.opentelemetry.io/otel/attribute"
-	otelCodes "go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/status"
 
@@ -1041,11 +1038,7 @@ func stackChangeHooks(send func(*stacks.StackChangePro
 		// span, we'll wrap it in a context so that the runtime's downstream
 		// operations will appear as children of it.
 		ContextAttach: func(parent context.Context, tracking any) context.Context {
-			span, ok := tracking.(trace.Span)
-			if !ok {
-				return parent
-			}
-			return trace.ContextWithSpan(parent, span)
+			return parent
 		},
 
 		// For the overall plan operation we don't emit any events to the client,
@@ -1053,13 +1046,9 @@ func stackChangeHooks(send func(*stacks.StackChangePro
 		// a root tracing span for all of the downstream planning operations to
 		// attach themselves to.
 		BeginPlan: func(ctx context.Context, s struct{}) any {
-			_, span := tracer.Start(ctx, "planning", trace.WithAttributes(
-				attribute.String("main_stack_source", mainStackSource.String()),
-			))
-			return span
+			return nil
 		},
 		EndPlan: func(ctx context.Context, span any, s struct{}) any {
-			span.(trace.Span).End()
 			return nil
 		},
 
@@ -1068,13 +1057,9 @@ func stackChangeHooks(send func(*stacks.StackChangePro
 		// a root tracing span for all of the downstream planning operations to
 		// attach themselves to.
 		BeginApply: func(ctx context.Context, s struct{}) any {
-			_, span := tracer.Start(ctx, "applying", trace.WithAttributes(
-				attribute.String("main_stack_source", mainStackSource.String()),
-			))
-			return span
+			return nil
 		},
 		EndApply: func(ctx context.Context, span any, s struct{}) any {
-			span.(trace.Span).End()
 			return nil
 		},
 
@@ -1104,27 +1089,18 @@ func stackChangeHooks(send func(*stacks.StackChangePro
 		},
 		BeginComponentInstancePlan: func(ctx context.Context, ci stackaddrs.AbsComponentInstance) any {
 			send(evtComponentInstanceStatus(ci, hooks.ComponentInstancePlanning))
-			_, span := tracer.Start(ctx, "planning", trace.WithAttributes(
-				attribute.String("component_instance", ci.String()),
-			))
-			return span
+			return nil
 		},
 		EndComponentInstancePlan: func(ctx context.Context, span any, ci stackaddrs.AbsComponentInstance) any {
 			send(evtComponentInstanceStatus(ci, hooks.ComponentInstancePlanned))
-			span.(trace.Span).SetStatus(otelCodes.Ok, "planning succeeded")
-			span.(trace.Span).End()
 			return nil
 		},
 		ErrorComponentInstancePlan: func(ctx context.Context, span any, ci stackaddrs.AbsComponentInstance) any {
 			send(evtComponentInstanceStatus(ci, hooks.ComponentInstanceErrored))
-			span.(trace.Span).SetStatus(otelCodes.Error, "planning failed")
-			span.(trace.Span).End()
 			return nil
 		},
 		DeferComponentInstancePlan: func(ctx context.Context, span any, ci stackaddrs.AbsComponentInstance) any {
 			send(evtComponentInstanceStatus(ci, hooks.ComponentInstanceDeferred))
-			span.(trace.Span).SetStatus(otelCodes.Error, "planning succeeded, but deferred")
-			span.(trace.Span).End()
 			return nil
 		},
 		PendingComponentInstanceApply: func(ctx context.Context, ci stackaddrs.AbsComponentInstance) {
@@ -1132,21 +1108,14 @@ func stackChangeHooks(send func(*stacks.StackChangePro
 		},
 		BeginComponentInstanceApply: func(ctx context.Context, ci stackaddrs.AbsComponentInstance) any {
 			send(evtComponentInstanceStatus(ci, hooks.ComponentInstanceApplying))
-			_, span := tracer.Start(ctx, "applying", trace.WithAttributes(
-				attribute.String("component_instance", ci.String()),
-			))
-			return span
+			return nil
 		},
 		EndComponentInstanceApply: func(ctx context.Context, span any, ci stackaddrs.AbsComponentInstance) any {
 			send(evtComponentInstanceStatus(ci, hooks.ComponentInstanceApplied))
-			span.(trace.Span).SetStatus(otelCodes.Ok, "applying succeeded")
-			span.(trace.Span).End()
 			return nil
 		},
 		ErrorComponentInstanceApply: func(ctx context.Context, span any, ci stackaddrs.AbsComponentInstance) any {
 			send(evtComponentInstanceStatus(ci, hooks.ComponentInstanceErrored))
-			span.(trace.Span).SetStatus(otelCodes.Error, "applying failed")
-			span.(trace.Span).End()
 			return nil
 		},
 
@@ -1176,11 +1145,6 @@ func stackChangeHooks(send func(*stacks.StackChangePro
 		// Upon completion of a component instance plan, we emit a planned
 		// change sumary event to the client for each resource instance.
 		ReportResourceInstancePlanned: func(ctx context.Context, span any, ric *hooks.ResourceInstanceChange) any {
-			span.(trace.Span).AddEvent("planned resource instance", trace.WithAttributes(
-				attribute.String("component_instance", ric.Addr.Component.String()),
-				attribute.String("resource_instance", ric.Addr.Item.String()),
-			))
-
 			ripc, err := resourceInstancePlanned(ric)
 			if err != nil {
 				return span
@@ -1195,11 +1159,6 @@ func stackChangeHooks(send func(*stacks.StackChangePro
 		},
 
 		ReportResourceInstanceDeferred: func(ctx context.Context, span any, change *hooks.DeferredResourceInstanceChange) any {
-			span.(trace.Span).AddEvent("deferred resource instance", trace.WithAttributes(
-				attribute.String("component_instance", change.Change.Addr.Component.String()),
-				attribute.String("resource_instance", change.Change.Addr.Item.String()),
-			))
-
 			ripc, err := resourceInstancePlanned(change.Change)
 			if err != nil {
 				return span
