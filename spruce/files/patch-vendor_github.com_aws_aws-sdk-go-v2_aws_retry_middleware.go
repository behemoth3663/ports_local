--- vendor/github.com/aws/aws-sdk-go-v2/aws/retry/middleware.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/retry/middleware.go
@@ -18,7 +18,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	smithymiddle "github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	"github.com/aws/smithy-go/transport/http"
 )
 
@@ -93,9 +92,6 @@ func (r *Attempt) HandleFinalize(ctx context.Context, 
 func (r *Attempt) HandleFinalize(ctx context.Context, in smithymiddle.FinalizeInput, next smithymiddle.FinalizeHandler) (
 	out smithymiddle.FinalizeOutput, metadata smithymiddle.Metadata, err error,
 ) {
-	ctx, span := tracing.StartSpan(ctx, "RetryLoop")
-	defer span.End()
-
 	var attemptClockSkew time.Duration
 	if !r.DisableClockSkewCorrection && r.ClientSkew != nil {
 		attemptClockSkew = time.Duration(r.ClientSkew.Load())
@@ -129,9 +125,6 @@ func (r *Attempt) HandleFinalize(ctx context.Context, 
 
 		var attemptResult AttemptResult
 
-		attemptCtx, span := tracing.StartSpan(attemptCtx, "Attempt", func(o *tracing.SpanOptions) {
-			o.Properties.Set("operation.attempt", attemptNum)
-		})
 		retryMetrics.Attempts.Add(ctx, 1, withOperationMetadata(ctx))
 
 		start := sdk.NowTime()
@@ -144,8 +137,6 @@ func (r *Attempt) HandleFinalize(ctx context.Context, 
 				o.Properties.Set("exception.type", errorType(err))
 			})
 		}
-
-		span.End()
 
 		attemptClockSkew, _ = awsmiddle.GetAttemptSkew(attemptResult.ResponseMetadata)
 
