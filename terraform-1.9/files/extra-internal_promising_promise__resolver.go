--- internal/promising/promise_resolver.go.orig	2024-10-16 12:28:59 UTC
+++ internal/promising/promise_resolver.go
@@ -5,9 +5,6 @@ import (
 
 import (
 	"context"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // PromiseResolver is an object representing responsibility for a promise, which
@@ -27,23 +24,6 @@ func (pr PromiseResolver[T]) Resolve(ctx context.Conte
 		panic("promise resolved by incorrect task")
 	}
 	resolvePromise(pr.p, v, err)
-
-	resolvingTaskSpan := trace.SpanFromContext(ctx)
-	resolvingTaskSpanContext := resolvingTaskSpan.SpanContext()
-	promiseSpanContext := pr.p.traceSpan.SpanContext()
-	pr.p.traceSpan.AddEvent(
-		"resolved",
-		trace.WithAttributes(
-			attribute.String("promising.resolved_by", resolvingTaskSpanContext.SpanID().String()),
-		),
-	)
-	resolvingTaskSpan.AddEvent(
-		"resolved a promise",
-		trace.WithAttributes(
-			attribute.String("promising.resolved_id", promiseSpanContext.SpanID().String()),
-		),
-	)
-	pr.p.traceSpan.End()
 }
 
 func (pr PromiseResolver[T]) PromiseID() PromiseID {
