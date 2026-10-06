--- vendor/github.com/couchbase/gocb/v2/client_ps.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/couchbase/gocb/v2/client_ps.go
@@ -4,13 +4,12 @@ import (
 	"context"
 	"errors"
 	"fmt"
-	"github.com/couchbaselabs/gocbconnstr/v2"
-	"go.opentelemetry.io/otel/metric"
-	"go.opentelemetry.io/otel/trace"
 	"sync"
 	"sync/atomic"
 	"time"
 
+	"github.com/couchbaselabs/gocbconnstr/v2"
+
 	"google.golang.org/grpc"
 
 	"github.com/couchbase/gocbcore/v10"
@@ -65,13 +64,13 @@ func (c *psConnectionMgr) buildConfig(cluster *Cluster
 
 	logger := newZapLogger()
 
-	var tp trace.TracerProvider
+	var tp interface{}
 	if c.tracer != nil {
 		if tracer, ok := c.tracer.tracer.(OtelAwareRequestTracer); ok {
 			tp = tracer.Provider()
 		}
 	}
-	var mp metric.MeterProvider
+	var mp interface{}
 	if c.meter != nil {
 		if meter, ok := c.meter.meter.(OtelAwareMeter); ok {
 			mp = meter.Provider()
@@ -320,7 +319,8 @@ func newPsOpManagerProvider(retry RetryStrategy, trace
 }
 
 func newPsOpManagerProvider(retry RetryStrategy, tracer *tracerWrapper, timeout time.Duration, meter *meterWrapper,
-	service string) *psOpManagerProvider {
+	service string,
+) *psOpManagerProvider {
 	return &psOpManagerProvider{
 		defaultRetryStrategy: retry,
 		tracer:               tracer,
@@ -495,12 +495,14 @@ func wrapPSOpCtx[ReqT any, RespT any](ctx context.Cont
 
 func wrapPSOpCtx[ReqT any, RespT any](ctx context.Context, m psOpManager,
 	req ReqT,
-	fn func(context.Context, ReqT, ...grpc.CallOption) (RespT, error)) (RespT, error) {
+	fn func(context.Context, ReqT, ...grpc.CallOption) (RespT, error),
+) (RespT, error) {
 	return wrapPSOpCtxWithPeek(ctx, m, req, m.TraceSpan(), fn, nil)
 }
 
 func wrapPSOp[ReqT any, RespT any](m psOpManager, req ReqT,
-	fn func(context.Context, ReqT, ...grpc.CallOption) (RespT, error)) (RespT, error) {
+	fn func(context.Context, ReqT, ...grpc.CallOption) (RespT, error),
+) (RespT, error) {
 	ctx, cancel := context.WithTimeout(m.Context(), m.Timeout())
 	defer cancel()
 
@@ -512,7 +514,8 @@ func wrapPSOpCtxWithPeek[ReqT any, RespT any](ctx cont
 	req ReqT,
 	parentSpan RequestSpan,
 	fn func(context.Context, ReqT, ...grpc.CallOption) (RespT, error),
-	peekResult func(RespT) error) (RespT, error) {
+	peekResult func(RespT) error,
+) (RespT, error) {
 	retryReq := newRetriableRequestPS(m.OpName(), m.IsIdempotent(), parentSpan, m.OperationID(), m.RetryStrategy())
 	m.SetRetryRequest(retryReq)
 
