--- vendor/github.com/open-policy-agent/opa/plugins/plugins.go.orig	2024-08-29 15:23:19 UTC
+++ vendor/github.com/open-policy-agent/opa/plugins/plugins.go
@@ -15,7 +15,6 @@ import (
 
 	"github.com/open-policy-agent/opa/internal/report"
 	"github.com/prometheus/client_golang/prometheus"
-	"go.opentelemetry.io/otel/sdk/trace"
 
 	"github.com/gorilla/mux"
 
@@ -200,7 +199,6 @@ type Manager struct {
 	enablePrintStatements        bool
 	router                       *mux.Router
 	prometheusRegister           prometheus.Registerer
-	tracerProvider               *trace.TracerProvider
 	distributedTacingOpts        tracing.Options
 	registeredNDCacheTriggers    []func(bool)
 	registeredTelemetryGatherers map[string]report.Gatherer
@@ -374,13 +372,6 @@ func WithPrometheusRegister(prometheusRegister prometh
 	}
 }
 
-// WithTracerProvider sets the passed *trace.TracerProvider to be used by plugins
-func WithTracerProvider(tracerProvider *trace.TracerProvider) func(*Manager) {
-	return func(m *Manager) {
-		m.tracerProvider = tracerProvider
-	}
-}
-
 // WithDistributedTracingOpts sets the options to be used by distributed tracing.
 func WithDistributedTracingOpts(tr tracing.Options) func(*Manager) {
 	return func(m *Manager) {
@@ -1046,11 +1037,6 @@ func (m *Manager) PrometheusRegister() prometheus.Regi
 // PrometheusRegister gets the prometheus.Registerer for this plugin manager.
 func (m *Manager) PrometheusRegister() prometheus.Registerer {
 	return m.prometheusRegister
-}
-
-// TracerProvider gets the *trace.TracerProvider for this plugin manager.
-func (m *Manager) TracerProvider() *trace.TracerProvider {
-	return m.tracerProvider
 }
 
 func (m *Manager) RegisterNDCacheTrigger(trigger func(bool)) {
