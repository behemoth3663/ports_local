--- vendor/github.com/googleapis/gax-go/v2/invoke.go.orig	2026-09-26 01:03:15 UTC
+++ vendor/github.com/googleapis/gax-go/v2/invoke.go
@@ -37,7 +37,6 @@ import (
 
 	"github.com/googleapis/gax-go/v2/apierror"
 	"github.com/googleapis/gax-go/v2/callctx"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // APICall is a user defined call stub.
@@ -90,47 +89,9 @@ func invoke(ctx context.Context, call APICall, setting
 		ctx = c
 	}
 
-	metricsEnabled := IsFeatureEnabled("METRICS")
-	tracingEnabled := IsFeatureEnabled("TRACING")
-	loggingEnabled := IsFeatureEnabled("LOGGING")
-
 	retryCount := 0
-	if metricsEnabled || tracingEnabled || loggingEnabled {
-		var start time.Time
-		if metricsEnabled {
-			start = time.Now()
-		}
-		if ExtractTransportTelemetry(ctx) == nil {
-			ctx = InjectTransportTelemetry(ctx, &TransportTelemetryData{})
-		}
-
-		var span trace.Span
-		if tracingEnabled {
-			ctx, span = startSpan(ctx, settings.clientTracing)
-		}
-
-		defer func() {
-			if !metricsEnabled && !tracingEnabled && err == nil {
-				return
-			}
-			errInfo := ExtractTelemetryErrorInfo(ctx, err)
-			if metricsEnabled {
-				recordMetricWithInfo(ctx, settings, time.Since(start), &errInfo)
-			}
-			if tracingEnabled {
-				endSpan(ctx, span, &errInfo, err)
-			}
-			if loggingEnabled && err != nil {
-				recordActionableLog(ctx, settings.clientLogging, &errInfo, retryCount, err)
-			}
-		}()
-	}
-
 	for {
 		ctxToUse := ctx
-		if tracingEnabled {
-			ctxToUse = withRetryCount(ctx, retryCount)
-		}
 		err = call(ctxToUse, settings)
 		if err == nil {
 			return nil
