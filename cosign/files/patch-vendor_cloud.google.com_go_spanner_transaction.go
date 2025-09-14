--- vendor/cloud.google.com/go/spanner/transaction.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/transaction.go
@@ -24,7 +24,6 @@ import (
 	"sync/atomic"
 	"time"
 
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/googleapis/gax-go/v2"
 	"github.com/googleapis/gax-go/v2/apierror"
@@ -105,8 +104,6 @@ type txReadOnly struct {
 	// disableRouteToLeader specifies if all the requests of type read-write and PDML
 	// need to be routed to the leader region.
 	disableRouteToLeader bool
-
-	otConfig *openTelemetryConfig
 }
 
 func (t *txReadOnly) isDefaultInlinedBegin() bool {
@@ -297,8 +294,6 @@ func (t *txReadOnly) ReadWithOptions(ctx context.Conte
 // ReadWithOptions returns a RowIterator for reading multiple rows from the
 // database. Pass a ReadOptions to modify the read operation.
 func (t *txReadOnly) ReadWithOptions(ctx context.Context, table string, keys KeySet, columns []string, opts *ReadOptions) (ri *RowIterator) {
-	ctx, _ = startSpan(ctx, "Read", t.otConfig.commonTraceStartOptions...)
-	defer func() { endSpan(ctx, ri.err) }()
 	var (
 		sh  *sessionHandle
 		ts  *sppb.TransactionSelector
@@ -389,15 +384,7 @@ func (t *txReadOnly) ReadWithOptions(ctx context.Conte
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
@@ -680,8 +667,6 @@ func (t *txReadOnly) query(ctx context.Context, statem
 }
 
 func (t *txReadOnly) query(ctx context.Context, statement Statement, options QueryOptions) (ri *RowIterator) {
-	ctx, _ = startSpan(ctx, "Query", t.otConfig.commonTraceStartOptions...)
-	defer func() { endSpan(ctx, ri.err) }()
 	req, sh, err := t.prepareExecuteSQL(ctx, statement, options)
 	if err != nil {
 		return &RowIterator{
@@ -714,15 +699,7 @@ func (t *txReadOnly) query(ctx context.Context, statem
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
@@ -918,15 +895,6 @@ func (t *ReadOnlyTransaction) begin(ctx context.Contex
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
@@ -1375,8 +1343,6 @@ func (t *ReadWriteTransaction) update(ctx context.Cont
 }
 
 func (t *ReadWriteTransaction) update(ctx context.Context, stmt Statement, opts QueryOptions) (rowCount int64, err error) {
-	ctx, _ = startSpan(ctx, "Update", t.otConfig.commonTraceStartOptions...)
-	defer func() { endSpan(ctx, err) }()
 	req, sh, err := t.prepareExecuteSQL(ctx, stmt, opts)
 	if err != nil {
 		return 0, err
@@ -1390,14 +1356,6 @@ func (t *ReadWriteTransaction) update(ctx context.Cont
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
@@ -1451,9 +1409,6 @@ func (t *ReadWriteTransaction) batchUpdateWithOptions(
 }
 
 func (t *ReadWriteTransaction) batchUpdateWithOptions(ctx context.Context, stmts []Statement, opts QueryOptions) (_ []int64, err error) {
-	ctx, _ = startSpan(ctx, "BatchUpdate", t.otConfig.commonTraceStartOptions...)
-	defer func() { endSpan(ctx, err) }()
-
 	sh, ts, err := t.acquire(ctx)
 	if err != nil {
 		return nil, err
@@ -1501,14 +1456,6 @@ func (t *ReadWriteTransaction) batchUpdateWithOptions(
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
@@ -1917,14 +1864,6 @@ func (t *ReadWriteTransaction) commit(ctx context.Cont
 		t.updatePrecommitToken(res.GetPrecommitToken())
 		res, err = performCommit(false)
 	}
-	if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-		if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "commit"); err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "commit", t.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if err != nil {
 		return resp, t.txReadOnly.updateTxState(toSpannerErrorWithCommitInfo(err, true))
 	}
@@ -2123,7 +2062,6 @@ func newReadWriteStmtBasedTransactionWithSessionHandle
 	t.options = options
 	t.txOpts = c.txo.merge(options)
 	t.ct = c.ct
-	t.otConfig = c.otConfig
 
 	if t.shouldExplicitBegin(0, t.txOpts) {
 		// Explicitly begin the transactions
