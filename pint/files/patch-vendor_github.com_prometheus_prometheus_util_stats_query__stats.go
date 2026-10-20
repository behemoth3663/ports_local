--- vendor/github.com/prometheus/prometheus/util/stats/query_stats.go.orig	2026-09-25 07:25:38 UTC
+++ vendor/github.com/prometheus/prometheus/util/stats/query_stats.go
@@ -19,8 +19,6 @@ import (
 	"fmt"
 
 	"github.com/prometheus/client_golang/prometheus"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // QueryTiming identifies the code area or functionality in which time is spent
@@ -221,25 +219,20 @@ type SpanTimer struct {
 type SpanTimer struct {
 	timer     *Timer
 	observers []prometheus.Observer
-
-	span trace.Span
 }
 
 func NewSpanTimer(ctx context.Context, operation string, timer *Timer, observers ...prometheus.Observer) (*SpanTimer, context.Context) {
-	ctx, span := otel.Tracer("").Start(ctx, operation)
 	timer.Start()
+	_ = operation
 
 	return &SpanTimer{
 		timer:     timer,
 		observers: observers,
-
-		span: span,
 	}, ctx
 }
 
 func (s *SpanTimer) Finish() {
 	s.timer.Stop()
-	s.span.End()
 
 	for _, obs := range s.observers {
 		obs.Observe(s.timer.ElapsedTime().Seconds())
