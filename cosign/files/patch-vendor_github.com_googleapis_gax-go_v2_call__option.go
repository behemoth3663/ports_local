--- vendor/github.com/googleapis/gax-go/v2/call_option.go.orig	2026-10-21 09:39:40 UTC
+++ vendor/github.com/googleapis/gax-go/v2/call_option.go
@@ -251,20 +251,15 @@ type clientMetricsOpt struct {
 }
 
 type clientMetricsOpt struct {
-	cm *ClientMetrics
 }
 
 // Resolve applies the ClientMetrics to the CallSettings.
 func (o clientMetricsOpt) Resolve(s *CallSettings) {
-	s.clientMetrics = o.cm
 }
 
 // WithClientMetrics applies metrics instrumentation to the CallSettings.
 //
 // This is for internal use only.
-func WithClientMetrics(cm *ClientMetrics) CallOption {
-	return clientMetricsOpt{cm: cm}
-}
 
 // CallSettings allow fine-grained control over how calls are made.
 type CallSettings struct {
@@ -281,48 +276,18 @@ type CallSettings struct {
 	// Timeout defines the amount of time that Invoke has to complete.
 	// Unexported so it cannot be changed by the code in an APICall.
 	timeout time.Duration
-
-	// clientMetrics holds the pre-allocated OpenTelemetry metrics instruments
-	// to use for this call.
-	clientMetrics *ClientMetrics
-
-	// clientTracing holds the pre-allocated OpenTelemetry tracer
-	// to use for this call.
-	clientTracing *ClientTracing
-
-	// clientLogging holds the pre-allocated OpenTelemetry/slog logger
-	// to use for this call.
-	clientLogging *ClientLogging
 }
 
 type clientTracingOpt struct {
-	ct *ClientTracing
 }
 
 // Resolve applies the ClientTracing to the CallSettings.
 func (o clientTracingOpt) Resolve(s *CallSettings) {
-	s.clientTracing = o.ct
 }
 
-// WithClientTracing applies tracing instrumentation to the CallSettings.
-//
-// This is for internal use only.
-func WithClientTracing(ct *ClientTracing) CallOption {
-	return clientTracingOpt{ct: ct}
-}
-
 type clientLoggingOpt struct {
-	cl *ClientLogging
 }
 
 // Resolve applies the ClientLogging to the CallSettings.
 func (o clientLoggingOpt) Resolve(s *CallSettings) {
-	s.clientLogging = o.cl
-}
-
-// WithClientLogging applies logging instrumentation to the CallSettings.
-//
-// This is for internal use only.
-func WithClientLogging(cl *ClientLogging) CallOption {
-	return clientLoggingOpt{cl: cl}
 }
