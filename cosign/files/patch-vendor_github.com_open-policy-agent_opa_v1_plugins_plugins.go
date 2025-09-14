--- vendor/github.com/open-policy-agent/opa/v1/plugins/plugins.go.orig	2025-09-26 08:53:51 UTC
+++ vendor/github.com/open-policy-agent/opa/v1/plugins/plugins.go
@@ -17,7 +17,6 @@ import (
 
 	"github.com/open-policy-agent/opa/internal/report"
 	"github.com/prometheus/client_golang/prometheus"
-	"go.opentelemetry.io/otel/sdk/trace"
 
 	bundleUtils "github.com/open-policy-agent/opa/internal/bundle"
 	cfg "github.com/open-policy-agent/opa/internal/config"
@@ -209,7 +208,6 @@ type Manager struct {
 	enablePrintStatements        bool
 	router                       *http.ServeMux
 	prometheusRegister           prometheus.Registerer
-	tracerProvider               *trace.TracerProvider
 	distributedTacingOpts        tracing.Options
 	registeredNDCacheTriggers    []func(bool)
 	registeredTelemetryGatherers map[string]report.Gatherer
@@ -390,13 +388,6 @@ func WithPrometheusRegister(prometheusRegister prometh
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
@@ -1143,11 +1134,6 @@ func (m *Manager) PrometheusRegister() prometheus.Regi
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
