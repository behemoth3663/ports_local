--- vendor/github.com/googleapis/gax-go/v2/invoke.go.orig	2026-10-22 19:52:07 UTC
+++ vendor/github.com/googleapis/gax-go/v2/invoke.go
@@ -37,7 +37,6 @@ import (
 
 	"github.com/googleapis/gax-go/v2/apierror"
 	"github.com/googleapis/gax-go/v2/callctx"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // APICall is a user defined call stub.
@@ -77,6 +76,56 @@ type sleeper func(ctx context.Context, d time.Duration
 
 type sleeper func(ctx context.Context, d time.Duration) error
 
+type ClientMetrics struct{}
+type ClientTracing struct{}
+type ClientLogging struct{}
+
+const (
+	ClientService  = "client.service"
+	ClientVersion  = "client.version"
+	ClientArtifact = "client.artifact"
+	RPCSystem      = "rpc.system"
+	URLDomain      = "url.domain"
+)
+
+func NewClientMetrics(...any) *ClientMetrics { return &ClientMetrics{} }
+
+func WithTelemetryLogger(any) any { return nil }
+
+func WithTelemetryAttributes(map[string]string) any { return nil }
+
+func NewClientTracing(...any) *ClientTracing { return &ClientTracing{} }
+
+func WithTracingAttributes(map[string]string) any { return nil }
+
+func NewClientLogging(...any) *ClientLogging { return &ClientLogging{} }
+
+func WithLoggerProvider(any) any { return nil }
+
+func WithLoggingAttributes(map[string]string) any { return nil }
+
+type TransportTelemetryData struct{}
+
+func (t *TransportTelemetryData) SetServerAddress(string) {}
+
+func (t *TransportTelemetryData) SetServerPort(int) {}
+
+type TelemetryErrorInfo struct{}
+
+func ExtractTransportTelemetry(context.Context) *TransportTelemetryData {
+	return nil
+}
+
+func InjectTransportTelemetry(ctx context.Context, _ *TransportTelemetryData) context.Context {
+	return ctx
+}
+
+func ExtractTelemetryErrorInfo(context.Context, error) TelemetryErrorInfo {
+	return TelemetryErrorInfo{}
+}
+
+func recordMetricWithInfo(context.Context, CallSettings, time.Duration, *TelemetryErrorInfo) {}
+
 // invoke implements Invoke, taking an additional sleeper argument for testing.
 func invoke(ctx context.Context, call APICall, settings CallSettings, sp sleeper) (err error) {
 	var retryer Retryer
@@ -91,11 +140,9 @@ func invoke(ctx context.Context, call APICall, setting
 	}
 
 	metricsEnabled := IsFeatureEnabled("METRICS")
-	tracingEnabled := IsFeatureEnabled("TRACING")
-	loggingEnabled := IsFeatureEnabled("LOGGING")
 
 	retryCount := 0
-	if metricsEnabled || tracingEnabled || loggingEnabled {
+	if metricsEnabled {
 		var start time.Time
 		if metricsEnabled {
 			start = time.Now()
@@ -104,33 +151,19 @@ func invoke(ctx context.Context, call APICall, setting
 			ctx = InjectTransportTelemetry(ctx, &TransportTelemetryData{})
 		}
 
-		var span trace.Span
-		if tracingEnabled {
-			ctx, span = startSpan(ctx, settings.clientTracing)
-		}
-
 		defer func() {
-			if !metricsEnabled && !tracingEnabled && err == nil {
+			if !metricsEnabled && err == nil {
 				return
 			}
 			errInfo := ExtractTelemetryErrorInfo(ctx, err)
 			if metricsEnabled {
 				recordMetricWithInfo(ctx, settings, time.Since(start), &errInfo)
 			}
-			if tracingEnabled {
-				endSpan(ctx, span, &errInfo, err)
-			}
-			if loggingEnabled && err != nil {
-				recordActionableLog(ctx, settings.clientLogging, &errInfo, retryCount, err)
-			}
 		}()
 	}
 
 	for {
 		ctxToUse := ctx
-		if tracingEnabled {
-			ctxToUse = withRetryCount(ctx, retryCount)
-		}
 		err = call(ctxToUse, settings)
 		if err == nil {
 			return nil
