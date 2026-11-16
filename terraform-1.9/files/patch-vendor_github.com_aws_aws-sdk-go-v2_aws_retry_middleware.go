--- vendor/github.com/aws/aws-sdk-go-v2/aws/retry/middleware.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/retry/middleware.go
@@ -17,7 +17,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	smithymiddle "github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	"github.com/aws/smithy-go/transport/http"
 )
 
@@ -112,9 +111,6 @@ func (r *Attempt) HandleFinalize(ctx context.Context, 
 
 		var attemptResult AttemptResult
 
-		attemptCtx, span := tracing.StartSpan(attemptCtx, "Attempt", func(o *tracing.SpanOptions) {
-			o.Properties.Set("operation.attempt", attemptNum)
-		})
 		retryMetrics.Attempts.Add(ctx, 1, withOperationMetadata(ctx))
 
 		start := sdk.NowTime()
@@ -127,8 +123,6 @@ func (r *Attempt) HandleFinalize(ctx context.Context, 
 				o.Properties.Set("exception.type", errorType(err))
 			})
 		}
-
-		span.End()
 
 		attemptClockSkew, _ = awsmiddle.GetAttemptSkew(attemptResult.ResponseMetadata)
 
