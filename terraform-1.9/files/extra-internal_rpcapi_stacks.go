--- internal/rpcapi/stacks.go.orig	2024-10-16 12:28:59 UTC
+++ internal/rpcapi/stacks.go
@@ -9,9 +9,6 @@ import (
 
 	"github.com/hashicorp/go-slug/sourceaddrs"
 	"github.com/hashicorp/go-slug/sourcebundle"
-	"go.opentelemetry.io/otel/attribute"
-	otelCodes "go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/status"
 
@@ -629,11 +626,7 @@ func stackChangeHooks(send func(*terraform1.StackChang
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
@@ -641,13 +634,9 @@ func stackChangeHooks(send func(*terraform1.StackChang
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
 
@@ -656,13 +645,9 @@ func stackChangeHooks(send func(*terraform1.StackChang
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
 
@@ -692,21 +677,14 @@ func stackChangeHooks(send func(*terraform1.StackChang
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
 		PendingComponentInstanceApply: func(ctx context.Context, ci stackaddrs.AbsComponentInstance) {
@@ -714,21 +692,14 @@ func stackChangeHooks(send func(*terraform1.StackChang
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
 
@@ -763,11 +734,6 @@ func stackChangeHooks(send func(*terraform1.StackChang
 				// TODO: what do we do?
 				return span
 			}
-
-			span.(trace.Span).AddEvent("planned resource instance", trace.WithAttributes(
-				attribute.String("component_instance", ric.Addr.Component.String()),
-				attribute.String("resource_instance", ric.Addr.Item.String()),
-			))
 
 			var moved *terraform1.StackChangeProgress_ResourceInstancePlannedChange_Moved
 			if !ric.Change.PrevRunAddr.Equal(ric.Change.Addr) {
