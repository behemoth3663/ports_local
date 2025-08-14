--- vendor/k8s.io/component-base/metrics/counter.go.orig	2025-07-16 01:44:15 UTC
+++ vendor/k8s.io/component-base/metrics/counter.go
@@ -22,7 +22,6 @@ import (
 
 	"github.com/blang/semver/v4"
 	"github.com/prometheus/client_golang/prometheus"
-	"go.opentelemetry.io/otel/trace"
 
 	dto "github.com/prometheus/client_model/go"
 )
@@ -126,18 +125,6 @@ func (e *exemplarCounterMetric) withExemplar(v float64
 
 // withExemplar attaches an exemplar to the metric.
 func (e *exemplarCounterMetric) withExemplar(v float64) {
-	if m, ok := e.CounterMetric.(prometheus.ExemplarAdder); ok {
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
 	e.CounterMetric.Add(v)
 }
 
