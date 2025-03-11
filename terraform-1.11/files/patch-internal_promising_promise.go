--- internal/promising/promise.go.orig	2024-10-16 12:28:59 UTC
+++ internal/promising/promise.go
@@ -7,9 +7,6 @@ import (
 	"context"
 	"sync"
 	"sync/atomic"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // promise represents a result that will become available at some point
@@ -17,7 +14,6 @@ type promise struct {
 type promise struct {
 	responsible atomic.Pointer[task]
 	result      atomic.Pointer[promiseResult]
-	traceSpan   trace.Span
 
 	waiting   []chan<- struct{}
 	waitingMu sync.Mutex
@@ -75,33 +71,15 @@ func NewPromise[T any](ctx context.Context) (PromiseRe
 // The caller should retain the resolver for its own use and pass the getter
 // to any other tasks that will consume the result of the promise.
 func NewPromise[T any](ctx context.Context) (PromiseResolver[T], PromiseGet[T]) {
-	callerSpan := trace.SpanFromContext(ctx)
 	initialResponsible := mustTaskFromContext(ctx)
 	p := &promise{}
 	p.responsible.Store(initialResponsible)
 	initialResponsible.responsible[p] = struct{}{}
 
-	ctx, span := tracer.Start(
-		ctx, "promise",
-		trace.WithNewRoot(),
-		trace.WithLinks(trace.Link{
-			SpanContext: trace.SpanContextFromContext(ctx),
-		}),
-	)
-	_ = ctx // prevent staticcheck from complaining until we have something actually using this
-	p.traceSpan = span
-	promiseSpanContext := span.SpanContext()
-
-	callerSpan.AddEvent("new promise", trace.WithAttributes(
-		attribute.String("promising.responsible_for", promiseSpanContext.SpanID().String()),
-	))
-
 	resolver := PromiseResolver[T]{p}
 	getter := PromiseGet[T](func(ctx context.Context) (T, error) {
 		reqT := mustTaskFromContext(ctx)
 
-		waiterSpan := trace.SpanFromContext(ctx)
-
 		ok := reqT.awaiting.CompareAndSwap(nil, p)
 		if !ok {
 			// If we get here then the task seems to have forked into two
@@ -164,12 +142,6 @@ func NewPromise[T any](ctx context.Context) (PromiseRe
 				err = append(err, checkP.promiseID())
 				affectedPromises = append(affectedPromises, checkP)
 			}
-			waiterSpan.AddEvent(
-				"task is self-dependent",
-				trace.WithAttributes(
-					attribute.String("promise.waiting_for_id", promiseSpanContext.SpanID().String()),
-				),
-			)
 
 			// All waiters for this promise need to see this error, because
 			// otherwise the other waiters might stall forever waiting for
@@ -188,41 +160,14 @@ func NewPromise[T any](ctx context.Context) (PromiseRe
 		if result := p.result.Load(); result != nil {
 			// No need to wait because the result is already available.
 			p.waitingMu.Unlock()
-			waiterSpan.AddEvent(
-				"promise is already resolved",
-				trace.WithAttributes(
-					attribute.String("promise.waiting_for_id", promiseSpanContext.SpanID().String()),
-				),
-			)
 			return getResolvedPromiseResult[T](result)
 		}
 
 		ch := make(chan struct{})
 		p.waiting = append(p.waiting, ch)
-		waiterCount := len(p.waiting)
 		p.waitingMu.Unlock()
 
-		waiterSpan.AddEvent(
-			"waiting for promise result",
-			trace.WithAttributes(
-				attribute.String("promise.waiting_for_id", promiseSpanContext.SpanID().String()),
-				attribute.Int("promise.waiter_count", waiterCount),
-			),
-		)
-		p.traceSpan.AddEvent(
-			"new task waiting",
-			trace.WithAttributes(
-				attribute.String("promise.waiter_id", waiterSpan.SpanContext().SpanID().String()),
-				attribute.Int("promise.waiter_count", waiterCount),
-			),
-		)
 		<-ch // channel will be closed once promise is resolved
-		waiterSpan.AddEvent(
-			"promise resolved",
-			trace.WithAttributes(
-				attribute.String("promise.waiting_for_id", promiseSpanContext.SpanID().String()),
-			),
-		)
 		if result := p.result.Load(); result != nil {
 			return getResolvedPromiseResult[T](result)
 		} else {
@@ -276,31 +221,16 @@ func resolvePromiseInternalFailure(p *promise, err err
 	p.waitingMu.Lock()
 	defer p.waitingMu.Unlock()
 
-	p.traceSpan.AddEvent("internal promise failure", trace.WithAttributes(
-		attribute.String("error", err.Error()),
-	))
-
 	// For internal failures we leave the responsibility data in place so
 	// that the responsible task can still try to resolve the promise and
 	// have it be a no-op, since the task that's responsible for resolving
 	// will not typically also call the promise getter, and so it won't
 	// know about the failure.
 
-	ok := p.result.CompareAndSwap(nil, &promiseResult{
+	p.result.CompareAndSwap(nil, &promiseResult{
 		err:    err,
 		forced: true,
 	})
-	if !ok {
-		// This suggests either that the responsible task beat us to the punch
-		// and resolved first, or that this promise was involved in two
-		// different self-dependence situations simultaneously and a different
-		// one got recorded already.
-		//
-		// Both situations are no big deal -- the promise got resolved one
-		// way or another -- but we'll record a tracing event for it just
-		// in case it's helpful while debugging something.
-		p.traceSpan.AddEvent("internal promise failure conflict")
-	}
 
 	for _, waitingCh := range p.waiting {
 		close(waitingCh)
