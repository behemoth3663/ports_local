--- vendor/github.com/aws/aws-sdk-go-v2/aws/retry/middleware.go.orig	2026-10-28 11:23:21 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/retry/middleware.go
@@ -18,7 +18,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	smithymiddle "github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	"github.com/aws/smithy-go/transport/http"
 )
 
@@ -126,9 +125,6 @@ func (r *Attempt) HandleFinalize(ctx context.Context, 
 
 		var attemptResult AttemptResult
 
-		attemptCtx, span := tracing.StartSpan(attemptCtx, "Attempt", func(o *tracing.SpanOptions) {
-			o.Properties.Set("operation.attempt", attemptNum)
-		})
 		retryMetrics.Attempts.Add(ctx, 1, withOperationMetadata(ctx))
 
 		start := sdk.NowTime()
@@ -141,8 +137,6 @@ func (r *Attempt) HandleFinalize(ctx context.Context, 
 				o.Properties.Set("exception.type", errorType(err))
 			})
 		}
-
-		span.End()
 
 		attemptClockSkew, _ = awsmiddle.GetAttemptSkew(attemptResult.ResponseMetadata)
 
