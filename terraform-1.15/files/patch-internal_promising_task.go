--- internal/promising/task.go.orig	2024-10-16 12:28:59 UTC
+++ internal/promising/task.go
@@ -6,9 +6,6 @@ import (
 import (
 	"context"
 	"sync/atomic"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // task represents one of a set of collaborating tasks that are communicating
@@ -68,21 +65,6 @@ func AsyncTask[P PromiseContainer](ctx context.Context
 		responsible: make(promiseSet),
 	}
 
-	// We treat async tasks as disconnected from their caller when tracing,
-	// because each task has an independent lifetime, but we do still track
-	// the causal relationship between the two using span links.
-	callerSpanContext := trace.SpanFromContext(ctx).SpanContext()
-
-	childCtx, childSpan := tracer.Start(
-		ctx, "async task",
-		trace.WithNewRoot(),
-		trace.WithLinks(
-			trace.Link{
-				SpanContext: callerSpanContext,
-			},
-		),
-	)
-
 	promises.AnnounceContainedPromises(func(apr AnyPromiseResolver) {
 		p := apr.promise()
 		if p.responsible.Load() != callerT {
@@ -93,18 +75,9 @@ func AsyncTask[P PromiseContainer](ctx context.Context
 		newT.responsible.Add(p)
 		callerT.responsible.Remove(p)
 		p.responsible.Store(newT)
-		p.traceSpan.AddEvent("delegated to new task", trace.WithAttributes(
-			attribute.String("promising.delegated_from", callerSpanContext.SpanID().String()),
-			attribute.String("promising.delegated_to", childSpan.SpanContext().SpanID().String()),
-		))
-		childSpan.AddEvent("inherited promise responsibility", trace.WithAttributes(
-			attribute.String("promising.responsible_for", p.traceSpan.SpanContext().SpanID().String()),
-		))
 	})
 
 	go func() {
-		ctx := childCtx
-		defer childSpan.End()
 		ctx = contextWithTask(ctx, newT)
 		impl(ctx, promises)
 
