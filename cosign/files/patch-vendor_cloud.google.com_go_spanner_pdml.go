--- vendor/cloud.google.com/go/spanner/pdml.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/pdml.go
@@ -17,10 +17,8 @@ import (
 import (
 	"context"
 
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/googleapis/gax-go/v2"
-	"go.opencensus.io/tag"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/metadata"
@@ -46,8 +44,6 @@ func (c *Client) partitionedUpdate(ctx context.Context
 }
 
 func (c *Client) partitionedUpdate(ctx context.Context, statement Statement, options QueryOptions) (count int64, err error) {
-	ctx, _ = startSpan(ctx, "PartitionedUpdate", c.otConfig.commonTraceStartOptions...)
-	defer func() { endSpan(ctx, err) }()
 	if err := checkNestedTxn(ctx); err != nil {
 		return 0, err
 	}
@@ -150,15 +146,6 @@ func executePdml(ctx context.Context, sh *sessionHandl
 
 	sh.updateLastUseTime()
 	resultSet, err := sh.getClient().ExecuteSql(ctx, req, gax.WithGRPCOptions(grpc.Header(&md)))
-	if getGFELatencyMetricsFlag() && md != nil && sh.session.pool != nil {
-		err := captureGFELatencyStats(tag.NewContext(ctx, sh.session.pool.tagMap), md, "executePdml_ExecuteSql")
-		if err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "executePdml_ExecuteSql", sh.session.pool.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if err != nil {
 		return 0, err
 	}
