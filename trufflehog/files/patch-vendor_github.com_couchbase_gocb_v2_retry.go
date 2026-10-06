--- vendor/github.com/couchbase/gocb/v2/retry.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/couchbase/gocb/v2/retry.go
@@ -3,7 +3,6 @@ import (
 import (
 	"context"
 	"errors"
-	"go.opentelemetry.io/otel/trace"
 	"time"
 
 	"google.golang.org/grpc"
@@ -102,8 +101,7 @@ type RetryAction interface {
 }
 
 // NoRetryRetryAction represents an action that indicates to not retry.
-type NoRetryRetryAction struct {
-}
+type NoRetryRetryAction struct{}
 
 // Duration is the length of time to wait before retrying an operation.
 func (ra *NoRetryRetryAction) Duration() time.Duration {
@@ -219,7 +217,8 @@ func newRetriableRequestPS(operation string, idempoten
 }
 
 func newRetriableRequestPS(operation string, idempotent bool, parentSpan RequestSpan, traceIdentifier string,
-	strategy RetryStrategy) *retriableRequestPs {
+	strategy RetryStrategy,
+) *retriableRequestPs {
 	loggerIdentifier := traceIdentifier
 	if loggerIdentifier == "" {
 		loggerIdentifier = uuid.NewString()[:6]
@@ -280,13 +279,10 @@ func handleRetriableRequest[ReqT any, RespT any](
 	retryReq *retriableRequestPs,
 	sendFn func(context.Context, ReqT, ...grpc.CallOption) (RespT, error),
 	retryReasonFn func(err error) RetryReason,
-	peekResult func(RespT) error) (RespT, error) {
+	peekResult func(RespT) error,
+) (RespT, error) {
 	for {
 		logSchedf("Writing request ID=%s, OP=%s", retryReq.loggerIdentifier, retryReq.operation)
-
-		if s, ok := retryReq.parentSpan.(OtelAwareRequestSpan); ok {
-			ctx = trace.ContextWithSpan(ctx, s.Wrapped())
-		}
 
 		res, err := sendFn(ctx, req)
 		logSchedf("Handling response ID=%s, OP=%s", retryReq.loggerIdentifier, retryReq.operation)
