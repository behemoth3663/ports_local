--- vendor/cloud.google.com/go/spanner/grpc_client.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/grpc_client.go
@@ -20,14 +20,11 @@ import (
 	"context"
 	"strings"
 	"sync/atomic"
-	"time"
 
 	vkit "cloud.google.com/go/spanner/apiv1"
 	"cloud.google.com/go/spanner/apiv1/spannerpb"
 	"cloud.google.com/go/spanner/internal"
 	"github.com/googleapis/gax-go/v2"
-	"go.opentelemetry.io/otel/attribute"
-	oteltrace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/option"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/peer"
@@ -174,66 +171,7 @@ func (g *grpcSpannerClient) DeleteSession(ctx context.
 	return err
 }
 
-// setSpanAttributes dynamically sets span attributes based on the request type.
-func setSpanAttributes[T any](span oteltrace.Span, req T) {
-	if !span.IsRecording() {
-		return
-	}
-
-	var attrs []attribute.KeyValue
-
-	if r, ok := any(req).(interface {
-		GetRequestOptions() *spannerpb.RequestOptions
-	}); ok {
-		if tag := r.GetRequestOptions().GetTransactionTag(); tag != "" {
-			attrs = append(attrs, attribute.String("transaction.tag", tag))
-		}
-		if tag := r.GetRequestOptions().GetRequestTag(); tag != "" {
-			attrs = append(attrs, attribute.String("statement.tag", tag))
-		}
-	}
-
-	if r, ok := any(req).(interface{ GetSql() string }); ok {
-		if sql := r.GetSql(); sql != "" {
-			attrs = append(attrs, attribute.String("db.statement", sql))
-		}
-	} else if r, ok := any(req).(interface {
-		GetStatements() []*spannerpb.ExecuteBatchDmlRequest_Statement
-	}); ok {
-		if stmts := r.GetStatements(); len(stmts) > 0 {
-			sqls := make([]string, len(stmts))
-			for i, stmt := range stmts {
-				sqls[i] = stmt.GetSql()
-			}
-			attrs = append(attrs, attribute.StringSlice("db.statement", sqls))
-		}
-	}
-
-	if r, ok := any(req).(interface{ GetTable() string }); ok {
-		if table := r.GetTable(); table != "" {
-			attrs = append(attrs, attribute.String("db.table", table))
-		}
-	}
-
-	span.SetAttributes(attrs...)
-}
-
-func setGFEAndAFESpanAttributes(span oteltrace.Span, latencyMap map[string]time.Duration) {
-	if !span.IsRecording() {
-		return
-	}
-	for t, v := range latencyMap {
-		if t == gfeTimingHeader || t == afeTimingHeader {
-			span.SetAttributes(
-				attribute.Float64(t[:3]+".latency_ms", float64(v.Nanoseconds())/1e6),
-			)
-		}
-	}
-}
-
 func (g *grpcSpannerClient) ExecuteSql(ctx context.Context, req *spannerpb.ExecuteSqlRequest, opts ...gax.CallOption) (*spannerpb.ResultSet, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	mt := g.newBuiltinMetricsTracer(ctx)
 	defer recordOperationCompletion(mt)
 	ctx = context.WithValue(ctx, metricsTracerKey, mt)
@@ -244,8 +182,6 @@ func (g *grpcSpannerClient) ExecuteStreamingSql(ctx co
 }
 
 func (g *grpcSpannerClient) ExecuteStreamingSql(ctx context.Context, req *spannerpb.ExecuteSqlRequest, opts ...gax.CallOption) (spannerpb.Spanner_ExecuteStreamingSqlClient, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	// Note: This method does not add g.optsWithNextRequestID to inject x-goog-spanner-request-id
 	// as it is already manually added when creating Stream iterators for ExecuteStreamingSql.
 	client, err := g.raw.ExecuteStreamingSql(peer.NewContext(ctx, &peer.Peer{}), req, opts...)
@@ -256,7 +192,6 @@ func (g *grpcSpannerClient) ExecuteStreamingSql(ctx co
 	if mt != nil && client != nil && mt.currOp.currAttempt != nil {
 		md, _ := client.Header()
 		latencyMap := parseServerTimingHeader(md)
-		setGFEAndAFESpanAttributes(span, latencyMap)
 		mt.currOp.currAttempt.setServerTimingMetrics(latencyMap)
 		mt.currOp.currAttempt.setDirectPathUsed(client.Context())
 	}
@@ -264,8 +199,6 @@ func (g *grpcSpannerClient) ExecuteBatchDml(ctx contex
 }
 
 func (g *grpcSpannerClient) ExecuteBatchDml(ctx context.Context, req *spannerpb.ExecuteBatchDmlRequest, opts ...gax.CallOption) (*spannerpb.ExecuteBatchDmlResponse, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	mt := g.newBuiltinMetricsTracer(ctx)
 	defer recordOperationCompletion(mt)
 	ctx = context.WithValue(ctx, metricsTracerKey, mt)
@@ -276,8 +209,6 @@ func (g *grpcSpannerClient) Read(ctx context.Context, 
 }
 
 func (g *grpcSpannerClient) Read(ctx context.Context, req *spannerpb.ReadRequest, opts ...gax.CallOption) (*spannerpb.ResultSet, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	mt := g.newBuiltinMetricsTracer(ctx)
 	defer recordOperationCompletion(mt)
 	ctx = context.WithValue(ctx, metricsTracerKey, mt)
@@ -290,8 +221,6 @@ func (g *grpcSpannerClient) StreamingRead(ctx context.
 func (g *grpcSpannerClient) StreamingRead(ctx context.Context, req *spannerpb.ReadRequest, opts ...gax.CallOption) (spannerpb.Spanner_StreamingReadClient, error) {
 	// Note: This method does not add g.optsWithNextRequestID, as it is already
 	// manually added when creating Stream iterators for StreamingRead.
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	client, err := g.raw.StreamingRead(peer.NewContext(ctx, &peer.Peer{}), req, opts...)
 	mt, ok := ctx.Value(metricsTracerKey).(*builtinMetricsTracer)
 	if !ok {
@@ -300,7 +229,6 @@ func (g *grpcSpannerClient) StreamingRead(ctx context.
 	if mt != nil && client != nil && mt.currOp.currAttempt != nil {
 		md, _ := client.Header()
 		latencyMap := parseServerTimingHeader(md)
-		setGFEAndAFESpanAttributes(span, latencyMap)
 		mt.currOp.currAttempt.setServerTimingMetrics(latencyMap)
 		mt.currOp.currAttempt.setDirectPathUsed(client.Context())
 	}
@@ -318,8 +246,6 @@ func (g *grpcSpannerClient) Commit(ctx context.Context
 }
 
 func (g *grpcSpannerClient) Commit(ctx context.Context, req *spannerpb.CommitRequest, opts ...gax.CallOption) (*spannerpb.CommitResponse, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	mt := g.newBuiltinMetricsTracer(ctx)
 	defer recordOperationCompletion(mt)
 	ctx = context.WithValue(ctx, metricsTracerKey, mt)
@@ -340,8 +266,6 @@ func (g *grpcSpannerClient) PartitionQuery(ctx context
 }
 
 func (g *grpcSpannerClient) PartitionQuery(ctx context.Context, req *spannerpb.PartitionQueryRequest, opts ...gax.CallOption) (*spannerpb.PartitionResponse, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	mt := g.newBuiltinMetricsTracer(ctx)
 	defer recordOperationCompletion(mt)
 	ctx = context.WithValue(ctx, metricsTracerKey, mt)
@@ -352,8 +276,6 @@ func (g *grpcSpannerClient) PartitionRead(ctx context.
 }
 
 func (g *grpcSpannerClient) PartitionRead(ctx context.Context, req *spannerpb.PartitionReadRequest, opts ...gax.CallOption) (*spannerpb.PartitionResponse, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	mt := g.newBuiltinMetricsTracer(ctx)
 	defer recordOperationCompletion(mt)
 	ctx = context.WithValue(ctx, metricsTracerKey, mt)
@@ -364,8 +286,6 @@ func (g *grpcSpannerClient) BatchWrite(ctx context.Con
 }
 
 func (g *grpcSpannerClient) BatchWrite(ctx context.Context, req *spannerpb.BatchWriteRequest, opts ...gax.CallOption) (spannerpb.Spanner_BatchWriteClient, error) {
-	span := oteltrace.SpanFromContext(ctx)
-	setSpanAttributes(span, req)
 	client, err := g.raw.BatchWrite(peer.NewContext(ctx, &peer.Peer{}), req, g.optsWithNextRequestID(opts)...)
 	mt, ok := ctx.Value(metricsTracerKey).(*builtinMetricsTracer)
 	if !ok {
@@ -374,7 +294,6 @@ func (g *grpcSpannerClient) BatchWrite(ctx context.Con
 	if mt != nil && client != nil && mt.currOp.currAttempt != nil {
 		md, _ := client.Header()
 		latencyMap := parseServerTimingHeader(md)
-		setGFEAndAFESpanAttributes(span, latencyMap)
 		mt.currOp.currAttempt.setServerTimingMetrics(latencyMap)
 		mt.currOp.currAttempt.setDirectPathUsed(client.Context())
 	}
