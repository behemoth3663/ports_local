--- vendor/cloud.google.com/go/spanner/retry.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/spanner/retry.go
@@ -22,7 +22,6 @@ import (
 	"strings"
 	"time"
 
-	"cloud.google.com/go/internal/trace"
 	"github.com/googleapis/gax-go/v2"
 	"google.golang.org/genproto/googleapis/rpc/errdetails"
 	"google.golang.org/grpc/codes"
@@ -113,18 +112,15 @@ func runWithRetryOnAbortedOrFailedInlineBeginOrSession
 				retryErr = err
 			}
 			if isSessionNotFoundError(retryErr) {
-				trace.TracePrintf(ctx, nil, "Retrying after Session not found")
 				continue
 			}
 			if isFailedInlineBeginTransaction(retryErr) {
-				trace.TracePrintf(ctx, nil, "Retrying after failed inline begin transaction")
 				continue
 			}
 			delay, shouldRetry := retryer.Retry(retryErr)
 			if !shouldRetry {
 				return err
 			}
-			trace.TracePrintf(ctx, nil, "Backing off after ABORTED for %s, then retrying", delay)
 			if err := gax.Sleep(ctx, delay); err != nil {
 				return err
 			}
