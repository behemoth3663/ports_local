--- vendor/k8s.io/component-base/metrics/histogram.go.orig	2025-12-17 22:45:12 UTC
+++ vendor/k8s.io/component-base/metrics/histogram.go
@@ -22,7 +22,6 @@ import (
 
 	"github.com/blang/semver/v4"
 	"github.com/prometheus/client_golang/prometheus"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Histogram is our internal representation for our wrapping struct around prometheus
@@ -47,18 +46,6 @@ func (e *exemplarHistogramMetric) Observe(v float64) {
 
 // Observe attaches an exemplar to the metric and then calls the delegate.
 func (e *exemplarHistogramMetric) Observe(v float64) {
-	if m, ok := e.delegate.(prometheus.ExemplarObserver); ok {
-		maybeSpanCtx := trace.SpanContextFromContext(e.ctx)
-		if maybeSpanCtx.IsValid() && maybeSpanCtx.IsSampled() {
-			exemplarLabels := prometheus.Labels{
-				"trace_id": maybeSpanCtx.TraceID().String(),
-				"span_id":  maybeSpanCtx.SpanID().String(),
-			}
-			m.ObserveWithExemplar(v, exemplarLabels)
-			return
-		}
-	}
-
 	e.delegate.Observe(v)
 }
 
@@ -263,17 +250,6 @@ func (h *exemplarHistogramVec) withExemplar(v float64)
 }
 
 func (h *exemplarHistogramVec) withExemplar(v float64) {
-	if m, ok := h.observer.(prometheus.ExemplarObserver); ok {
-		maybeSpanCtx := trace.SpanContextFromContext(h.HistogramVecWithContext.ctx)
-		if maybeSpanCtx.IsValid() && maybeSpanCtx.IsSampled() {
-			m.ObserveWithExemplar(v, prometheus.Labels{
-				"trace_id": maybeSpanCtx.TraceID().String(),
-				"span_id":  maybeSpanCtx.SpanID().String(),
-			})
-			return
-		}
-	}
-
 	h.observer.Observe(v)
 }
 
