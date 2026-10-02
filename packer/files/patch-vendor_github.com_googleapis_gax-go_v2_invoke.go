--- vendor/github.com/googleapis/gax-go/v2/invoke.go.orig	2026-10-02 20:06:39 UTC
+++ vendor/github.com/googleapis/gax-go/v2/invoke.go
@@ -31,12 +31,10 @@ import (
 
 import (
 	"context"
-	"strconv"
 	"strings"
 	"time"
 
 	"github.com/googleapis/gax-go/v2/apierror"
-	"github.com/googleapis/gax-go/v2/callctx"
 )
 
 // APICall is a user defined call stub.
@@ -47,8 +45,7 @@ func withRetryCount(ctx context.Context, retryCount in
 // attempted. On the initial request, retry count is 0.
 // On a second request (the first retry), retry count is 1.
 func withRetryCount(ctx context.Context, retryCount int) context.Context {
-	// Add to telemetry context so it's visible to observability wrappers
-	return callctx.WithTelemetryContext(ctx, "resend_count", strconv.Itoa(retryCount))
+	return ctx
 }
 
 // Invoke calls the given APICall, performing retries as specified by opts, if
@@ -89,22 +86,10 @@ func invoke(ctx context.Context, call APICall, setting
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
-	tracingEnabled := IsFeatureEnabled("TRACING")
 	for {
 		ctxToUse := ctx
-		if tracingEnabled {
-			ctxToUse = withRetryCount(ctx, retryCount)
-		}
 		err = call(ctxToUse, settings)
 		if err == nil {
 			return nil
