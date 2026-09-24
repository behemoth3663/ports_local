--- vendor/github.com/googleapis/gax-go/v2/invoke.go.orig	2026-02-03 18:41:38 UTC
+++ vendor/github.com/googleapis/gax-go/v2/invoke.go
@@ -90,13 +90,8 @@ func invoke(ctx context.Context, call APICall, setting
 	}
 
 	retryCount := 0
-	// Feature gate: GOOGLE_SDK_GO_EXPERIMENTAL_TRACING=true
-	tracingEnabled := IsFeatureEnabled("TRACING")
 	for {
 		ctxToUse := ctx
-		if tracingEnabled {
-			ctxToUse = withRetryCount(ctx, retryCount)
-		}
 		err := call(ctxToUse, settings)
 		if err == nil {
 			return nil
