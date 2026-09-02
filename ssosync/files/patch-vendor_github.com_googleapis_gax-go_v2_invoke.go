--- vendor/github.com/googleapis/gax-go/v2/invoke.go.orig	2026-04-01 19:39:53 UTC
+++ vendor/github.com/googleapis/gax-go/v2/invoke.go
@@ -89,14 +89,6 @@ func invoke(ctx context.Context, call APICall, setting
 		ctx = c
 	}
 
-	if IsFeatureEnabled("METRICS") {
-		start := time.Now()
-		ctx = InjectTransportTelemetry(ctx, &TransportTelemetryData{})
-		defer func() {
-			recordMetric(ctx, settings, time.Since(start), err)
-		}()
-	}
-
 	retryCount := 0
 	// Feature gate: GOOGLE_SDK_GO_EXPERIMENTAL_TRACING=true
 	tracingEnabled := IsFeatureEnabled("TRACING")
