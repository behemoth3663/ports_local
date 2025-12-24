--- vendor/k8s.io/component-base/metrics/counter.go.orig	2025-12-17 22:45:12 UTC
+++ vendor/k8s.io/component-base/metrics/counter.go
@@ -22,7 +22,6 @@ import (
 
 	"github.com/blang/semver/v4"
 	"github.com/prometheus/client_golang/prometheus"
-	"go.opentelemetry.io/otel/trace"
 
 	dto "github.com/prometheus/client_model/go"
 )
@@ -109,18 +108,6 @@ func (e *exemplarCounterMetric) Add(v float64) {
 
 // Add attaches an exemplar to the metric and then calls the delegate.
 func (e *exemplarCounterMetric) Add(v float64) {
-	if m, ok := e.delegate.(prometheus.ExemplarAdder); ok {
-		maybeSpanCtx := trace.SpanContextFromContext(e.ctx)
-		if maybeSpanCtx.IsValid() && maybeSpanCtx.IsSampled() {
-			exemplarLabels := prometheus.Labels{
-				"trace_id": maybeSpanCtx.TraceID().String(),
-				"span_id":  maybeSpanCtx.SpanID().String(),
-			}
-			m.AddWithExemplar(v, exemplarLabels)
-			return
-		}
-	}
-
 	e.delegate.Add(v)
 }
 
