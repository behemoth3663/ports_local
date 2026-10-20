--- vendor/github.com/open-policy-agent/opa/v1/plugins/plugins.go.orig	2026-10-20 09:33:13 UTC
+++ vendor/github.com/open-policy-agent/opa/v1/plugins/plugins.go
@@ -17,7 +17,6 @@ import (
 
 	"github.com/open-policy-agent/opa/internal/report"
 	"github.com/prometheus/client_golang/prometheus"
-	"go.opentelemetry.io/otel/sdk/trace"
 
 	bundleUtils "github.com/open-policy-agent/opa/internal/bundle"
 	cfg "github.com/open-policy-agent/opa/internal/config"
@@ -208,7 +207,7 @@ type Manager struct {
 	enablePrintStatements        bool
 	router                       *http.ServeMux
 	prometheusRegister           prometheus.Registerer
-	tracerProvider               *trace.TracerProvider
+	tracerProvider               any
 	distributedTacingOpts        tracing.Options
 	registeredNDCacheTriggers    []func(bool)
 	registeredTelemetryGatherers map[string]report.Gatherer
@@ -382,8 +381,8 @@ func WithPrometheusRegister(prometheusRegister prometh
 	}
 }
 
-// WithTracerProvider sets the passed *trace.TracerProvider to be used by plugins
-func WithTracerProvider(tracerProvider *trace.TracerProvider) func(*Manager) {
+// WithTracerProvider sets the passed tracer provider to be used by plugins.
+func WithTracerProvider(tracerProvider any) func(*Manager) {
 	return func(m *Manager) {
 		m.tracerProvider = tracerProvider
 	}
@@ -1059,8 +1058,8 @@ func (m *Manager) PrometheusRegister() prometheus.Regi
 	return m.prometheusRegister
 }
 
-// TracerProvider gets the *trace.TracerProvider for this plugin manager.
-func (m *Manager) TracerProvider() *trace.TracerProvider {
+// TracerProvider gets the tracer provider for this plugin manager.
+func (m *Manager) TracerProvider() any {
 	return m.tracerProvider
 }
 
