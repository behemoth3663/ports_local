--- vendor/github.com/googleapis/gax-go/v2/invoke.go.orig	2026-09-26 01:03:15 UTC
+++ vendor/github.com/googleapis/gax-go/v2/invoke.go
@@ -31,13 +31,10 @@ import (
 
 import (
 	"context"
-	"strconv"
 	"strings"
 	"time"
 
 	"github.com/googleapis/gax-go/v2/apierror"
-	"github.com/googleapis/gax-go/v2/callctx"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // APICall is a user defined call stub.
@@ -49,7 +46,7 @@ func withRetryCount(ctx context.Context, retryCount in
 // On a second request (the first retry), retry count is 1.
 func withRetryCount(ctx context.Context, retryCount int) context.Context {
 	// Add to telemetry context so it's visible to observability wrappers
-	return callctx.WithTelemetryContext(ctx, "resend_count", strconv.Itoa(retryCount))
+	return ctx
 }
 
 // Invoke calls the given APICall, performing retries as specified by opts, if
@@ -90,47 +87,10 @@ func invoke(ctx context.Context, call APICall, setting
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
