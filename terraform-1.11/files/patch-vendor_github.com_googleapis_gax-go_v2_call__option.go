--- vendor/github.com/googleapis/gax-go/v2/call_option.go.orig	2026-04-01 19:39:53 UTC
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
@@ -284,5 +279,4 @@ type CallSettings struct {
 
 	// clientMetrics holds the pre-allocated OpenTelemetry metrics instruments
 	// to use for this call.
-	clientMetrics *ClientMetrics
 }
