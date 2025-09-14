--- vendor/cloud.google.com/go/spanner/batch.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/batch.go
@@ -23,7 +23,6 @@ import (
 	"log"
 	"time"
 
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/googleapis/gax-go/v2"
 	"google.golang.org/grpc"
@@ -144,14 +143,6 @@ func (t *BatchReadOnlyTransaction) PartitionReadUsingI
 		PartitionOptions: opt.toProto(),
 	}, gax.WithGRPCOptions(grpc.Header(&md)))
 
-	if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-		if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "PartitionReadUsingIndexWithOptions"); err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "PartitionReadUsingIndexWithOptions", t.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if isUnimplementedErrorForMultiplexedPartitionReads(err) && t.sp.isMultiplexedSessionForPartitionedOpsEnabled() {
 		t.sp.disableMultiplexedSessionForPartitionedOps()
 	}
@@ -214,14 +205,6 @@ func (t *BatchReadOnlyTransaction) partitionQuery(ctx 
 	sh.updateLastUseTime()
 	resp, err := client.PartitionQuery(contextWithOutgoingMetadata(ctx, sh.getMetadata(), t.disableRouteToLeader), req, gax.WithGRPCOptions(grpc.Header(&md)))
 
-	if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-		if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "partitionQuery"); err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "partitionQuery", t.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if isUnimplementedErrorForMultiplexedPartitionReads(err) && t.sp.isMultiplexedSessionForPartitionedOpsEnabled() {
 		t.sp.disableMultiplexedSessionForPartitionedOps()
 	}
@@ -295,15 +278,6 @@ func (t *BatchReadOnlyTransaction) Cleanup(ctx context
 	var md metadata.MD
 	err := client.DeleteSession(contextWithOutgoingMetadata(ctx, sh.getMetadata(), true), &sppb.DeleteSessionRequest{Name: sid}, gax.WithGRPCOptions(grpc.Header(&md)))
 
-	if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-		if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "Cleanup"); err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "Cleanup", t.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
-
 	if err != nil {
 		var logger *log.Logger
 		if sh.session != nil {
@@ -349,15 +323,7 @@ func (t *BatchReadOnlyTransaction) Execute(ctx context
 			if err != nil {
 				return client, err
 			}
-			md, err := client.Header()
-			if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-				if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "Execute"); err != nil {
-					trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-				}
-			}
-			if metricErr := recordGFELatencyMetricsOT(ctx, md, "Execute", t.otConfig); metricErr != nil {
-				trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-			}
+			_, err = client.Header()
 			if isUnimplementedErrorForMultiplexedPartitionReads(err) && t.sp.isMultiplexedSessionForPartitionedOpsEnabled() {
 				t.sp.disableMultiplexedSessionForPartitionedOps()
 			}
@@ -381,16 +347,8 @@ func (t *BatchReadOnlyTransaction) Execute(ctx context
 			if err != nil {
 				return client, err
 			}
-			md, err := client.Header()
+			_, err = client.Header()
 
-			if getGFELatencyMetricsFlag() && md != nil && t.ct != nil {
-				if err := createContextAndCaptureGFELatencyMetrics(ctx, t.ct, md, "Execute"); err != nil {
-					trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-				}
-			}
-			if metricErr := recordGFELatencyMetricsOT(ctx, md, "Execute", t.otConfig); metricErr != nil {
-				trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-			}
 			if isUnimplementedErrorForMultiplexedPartitionReads(err) && t.sp.isMultiplexedSessionForPartitionedOpsEnabled() {
 				t.sp.disableMultiplexedSessionForPartitionedOps()
 			}
