--- vendor/cloud.google.com/go/spanner/transaction.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/spanner/transaction.go
@@ -23,7 +23,6 @@ import (
 	"sync/atomic"
 	"time"
 
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/googleapis/gax-go/v2"
 	"google.golang.org/api/iterator"
@@ -256,8 +255,6 @@ func (t *txReadOnly) ReadWithOptions(ctx context.Conte
 // ReadWithOptions returns a RowIterator for reading multiple rows from the
 // database. Pass a ReadOptions to modify the read operation.
 func (t *txReadOnly) ReadWithOptions(ctx context.Context, table string, keys KeySet, columns []string, opts *ReadOptions) (ri *RowIterator) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.Read")
-	defer func() { trace.EndSpan(ctx, ri.err) }()
 	var (
 		sh  *sessionHandle
 		ts  *sppb.TransactionSelector
@@ -348,15 +345,7 @@ func (t *txReadOnly) ReadWithOptions(ctx context.Conte
 				}
 				return client, t.updateTxState(err)
 			}
-			md, err := client.Header()
-			if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-				if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "ReadWithOptions"); err != nil {
-					trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-				}
-			}
-			if metricErr := recordGFELatencyMetricsOT(ctx, md, "ReadWithOptions", t.otConfig); metricErr != nil {
-				trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-			}
+			_, err = client.Header()
 			return client, err
 		},
 		t.replaceSessionFunc,
@@ -615,8 +604,6 @@ func (t *txReadOnly) query(ctx context.Context, statem
 }
 
 func (t *txReadOnly) query(ctx context.Context, statement Statement, options QueryOptions) (ri *RowIterator) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.Query")
-	defer func() { trace.EndSpan(ctx, ri.err) }()
 	req, sh, err := t.prepareExecuteSQL(ctx, statement, options)
 	if err != nil {
 		return &RowIterator{
@@ -649,15 +636,7 @@ func (t *txReadOnly) query(ctx context.Context, statem
 				}
 				return client, t.updateTxState(err)
 			}
-			md, err := client.Header()
-			if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-				if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "query"); err != nil {
-					trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-				}
-			}
-			if metricErr := recordGFELatencyMetricsOT(ctx, md, "query", t.otConfig); metricErr != nil {
-				trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-			}
+			_, err = client.Header()
 			return client, err
 		},
 		t.replaceSessionFunc,
@@ -850,15 +829,6 @@ func (t *ReadOnlyTransaction) begin(ctx context.Contex
 			},
 		}, gax.WithGRPCOptions(grpc.Header(&md)))
 
-		if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-			if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "begin_BeginTransaction"); err != nil {
-				trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-			}
-		}
-		if metricErr := recordGFELatencyMetricsOT(ctx, md, "begin_BeginTransaction", t.otConfig); metricErr != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-		}
-
 		if isSessionNotFoundError(err) {
 			sh.destroy()
 			continue
@@ -1244,8 +1214,6 @@ func (t *ReadWriteTransaction) update(ctx context.Cont
 }
 
 func (t *ReadWriteTransaction) update(ctx context.Context, stmt Statement, opts QueryOptions) (rowCount int64, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.Update")
-	defer func() { trace.EndSpan(ctx, err) }()
 	req, sh, err := t.prepareExecuteSQL(ctx, stmt, opts)
 	if err != nil {
 		return 0, err
@@ -1259,14 +1227,6 @@ func (t *ReadWriteTransaction) update(ctx context.Cont
 	var md metadata.MD
 	resultSet, err := sh.getClient().ExecuteSql(contextWithOutgoingMetadata(ctx, sh.getMetadata(), t.disableRouteToLeader), req, gax.WithGRPCOptions(grpc.Header(&md)))
 
-	if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-		if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "update"); err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "update", t.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if err != nil {
 		if hasInlineBeginTransaction {
 			t.setTransactionID(nil)
@@ -1320,9 +1280,6 @@ func (t *ReadWriteTransaction) batchUpdateWithOptions(
 }
 
 func (t *ReadWriteTransaction) batchUpdateWithOptions(ctx context.Context, stmts []Statement, opts QueryOptions) (_ []int64, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.BatchUpdate")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	sh, ts, err := t.acquire(ctx)
 	if err != nil {
 		return nil, err
@@ -1370,14 +1327,6 @@ func (t *ReadWriteTransaction) batchUpdateWithOptions(
 		LastStatements: opts.LastStatement,
 	}, gax.WithGRPCOptions(grpc.Header(&md)))
 
-	if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-		if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "batchUpdateWithOptions"); err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", ToSpannerError(err))
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "batchUpdateWithOptions", t.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if err != nil {
 		if hasInlineBeginTransaction {
 			t.setTransactionID(nil)
@@ -1770,14 +1719,6 @@ func (t *ReadWriteTransaction) commit(ctx context.Cont
 	if res.GetMultiplexedSessionRetry() != nil {
 		t.updatePrecommitToken(res.GetPrecommitToken())
 		res, err = performCommit(false)
-	}
-	if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-		if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "commit"); err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "commit", t.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
 	}
 	if err != nil {
 		return resp, t.txReadOnly.updateTxState(toSpannerErrorWithCommitInfo(err, true))
