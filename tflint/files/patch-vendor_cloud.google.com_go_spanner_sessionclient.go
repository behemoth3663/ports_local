--- vendor/cloud.google.com/go/spanner/sessionclient.go.orig	2025-08-14 14:42:16 UTC
+++ vendor/cloud.google.com/go/spanner/sessionclient.go
@@ -25,12 +25,9 @@ import (
 	"sync/atomic"
 	"time"
 
-	"cloud.google.com/go/internal/trace"
 	vkit "cloud.google.com/go/spanner/apiv1"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
-	"cloud.google.com/go/spanner/internal"
 	"github.com/googleapis/gax-go/v2"
-	"go.opencensus.io/tag"
 	"google.golang.org/api/option"
 
 	gtransport "google.golang.org/api/transport/grpc"
@@ -106,7 +103,6 @@ type sessionClient struct {
 	batchTimeout         time.Duration
 	logger               *log.Logger
 	callOptions          *vkit.CallOptions
-	otConfig             *openTelemetryConfig
 	metricsTracerFactory *builtinMetricsTracerFactory
 	channelIDMap         map[*grpc.ClientConn]uint64
 
@@ -171,28 +167,6 @@ func (sc *sessionClient) createSession(ctx context.Con
 		Session:  &sppb.Session{Labels: sc.sessionLabels, CreatorRole: sc.databaseRole},
 	}, gax.WithGRPCOptions(grpc.Header(&md)))
 
-	if getGFELatencyMetricsFlag() && md != nil {
-		_, instance, database, err := parseDatabaseName(sc.database)
-		if err != nil {
-			return nil, ToSpannerError(err)
-		}
-		ctxGFE, err := tag.New(ctx,
-			tag.Upsert(tagKeyClientID, sc.id),
-			tag.Upsert(tagKeyDatabase, database),
-			tag.Upsert(tagKeyInstance, instance),
-			tag.Upsert(tagKeyLibVersion, internal.Version),
-		)
-		if err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", ToSpannerError(err))
-		}
-		err = captureGFELatencyStats(ctxGFE, md, "createSession")
-		if err != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", ToSpannerError(err))
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, md, "createSession", sc.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if err != nil {
 		return nil, ToSpannerError(err)
 	}
@@ -270,9 +244,6 @@ func (sc *sessionClient) executeBatchCreateSessions(cl
 	defer sc.waitWorkers.Done()
 	ctx, cancel := context.WithTimeout(context.Background(), sc.batchTimeout)
 	defer cancel()
-	ctx, _ = startSpan(ctx, "BatchCreateSessions", sc.otConfig.commonTraceStartOptions...)
-	defer func() { endSpan(ctx, nil) }()
-	trace.TracePrintf(ctx, nil, "Creating a batch of %d sessions", createCount)
 
 	remainingCreateCount := createCount
 	for {
@@ -281,12 +252,10 @@ func (sc *sessionClient) executeBatchCreateSessions(cl
 		sc.mu.Unlock()
 		if closed {
 			err := spannerErrorf(codes.Canceled, "Session client closed")
-			trace.TracePrintf(ctx, nil, "Session client closed while creating a batch of %d sessions: %v", createCount, err)
 			consumer.sessionCreationFailed(ctx, err, remainingCreateCount, false)
 			break
 		}
 		if ctx.Err() != nil {
-			trace.TracePrintf(ctx, nil, "Context error while creating a batch of %d sessions: %v", createCount, ctx.Err())
 			consumer.sessionCreationFailed(ctx, ToSpannerError(ctx.Err()), remainingCreateCount, false)
 			break
 		}
@@ -297,36 +266,11 @@ func (sc *sessionClient) executeBatchCreateSessions(cl
 			SessionTemplate: &sppb.Session{Labels: labels, CreatorRole: sc.databaseRole},
 		}, gax.WithGRPCOptions(grpc.Header(&mdForGFELatency)))
 
-		if getGFELatencyMetricsFlag() && mdForGFELatency != nil {
-			_, instance, database, err := parseDatabaseName(sc.database)
-			if err != nil {
-				trace.TracePrintf(ctx, nil, "Error getting instance and database name: %v", err)
-			}
-			// Errors should not prevent initializing the session pool.
-			ctxGFE, err := tag.New(ctx,
-				tag.Upsert(tagKeyClientID, sc.id),
-				tag.Upsert(tagKeyDatabase, database),
-				tag.Upsert(tagKeyInstance, instance),
-				tag.Upsert(tagKeyLibVersion, internal.Version),
-			)
-			if err != nil {
-				trace.TracePrintf(ctx, nil, "Error in adding tags in BatchCreateSessions for GFE Latency: %v", err)
-			}
-			err = captureGFELatencyStats(ctxGFE, mdForGFELatency, "executeBatchCreateSessions")
-			if err != nil {
-				trace.TracePrintf(ctx, nil, "Error in Capturing GFE Latency and Header Missing count. Try disabling and rerunning. Error: %v", err)
-			}
-		}
-		if metricErr := recordGFELatencyMetricsOT(ctx, mdForGFELatency, "executeBatchCreateSessions", sc.otConfig); metricErr != nil {
-			trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-		}
 		if err != nil {
-			trace.TracePrintf(ctx, nil, "Error creating a batch of %d sessions: %v", remainingCreateCount, err)
 			consumer.sessionCreationFailed(ctx, ToSpannerError(err), remainingCreateCount, false)
 			break
 		}
 		actuallyCreated := int32(len(response.Session))
-		trace.TracePrintf(ctx, nil, "Received a batch of %d sessions", actuallyCreated)
 		for _, s := range response.Session {
 			consumer.sessionReady(ctx, &session{valid: true, client: client, id: s.Name, createTime: time.Now(), md: md, logger: sc.logger})
 		}
@@ -335,26 +279,20 @@ func (sc *sessionClient) executeBatchCreateSessions(cl
 			// should do another call using the same gRPC channel.
 			remainingCreateCount -= actuallyCreated
 		} else {
-			trace.TracePrintf(ctx, nil, "Finished creating %d sessions", createCount)
 			break
 		}
 	}
 }
 
 func (sc *sessionClient) executeCreateMultiplexedSession(ctx context.Context, client spannerClient, md metadata.MD, consumer sessionConsumer) {
-	ctx, _ = startSpan(ctx, "CreateSession", sc.otConfig.commonTraceStartOptions...)
-	defer func() { endSpan(ctx, nil) }()
-	trace.TracePrintf(ctx, nil, "Creating a multiplexed session")
 	sc.mu.Lock()
 	closed := sc.closed
 	sc.mu.Unlock()
 	if closed {
-		err := spannerErrorf(codes.Canceled, "Session client closed")
-		trace.TracePrintf(ctx, nil, "Session client closed while creating a multiplexed session: %v", err)
+		_ = spannerErrorf(codes.Canceled, "Session client closed")
 		return
 	}
 	if ctx.Err() != nil {
-		trace.TracePrintf(ctx, nil, "Context error while creating a multiplexed session: %v", ctx.Err())
 		consumer.sessionCreationFailed(ctx, ToSpannerError(ctx.Err()), 1, true)
 		return
 	}
@@ -365,36 +303,11 @@ func (sc *sessionClient) executeCreateMultiplexedSessi
 		Session: &sppb.Session{CreatorRole: sc.databaseRole, Multiplexed: true},
 	}, gax.WithGRPCOptions(grpc.Header(&mdForGFELatency)))
 
-	if getGFELatencyMetricsFlag() && mdForGFELatency != nil {
-		_, instance, database, err := parseDatabaseName(sc.database)
-		if err != nil {
-			trace.TracePrintf(ctx, nil, "Error getting instance and database name: %v", err)
-		}
-		// Errors should not prevent initializing the session pool.
-		ctxGFE, err := tag.New(ctx,
-			tag.Upsert(tagKeyClientID, sc.id),
-			tag.Upsert(tagKeyDatabase, database),
-			tag.Upsert(tagKeyInstance, instance),
-			tag.Upsert(tagKeyLibVersion, internal.Version),
-		)
-		if err != nil {
-			trace.TracePrintf(ctx, nil, "Error in adding tags in CreateSession for GFE Latency: %v", err)
-		}
-		err = captureGFELatencyStats(ctxGFE, mdForGFELatency, "executeCreateSession")
-		if err != nil {
-			trace.TracePrintf(ctx, nil, "Error in Capturing GFE Latency and Header Missing count. Try disabling and rerunning. Error: %v", err)
-		}
-	}
-	if metricErr := recordGFELatencyMetricsOT(ctx, mdForGFELatency, "executeCreateSession", sc.otConfig); metricErr != nil {
-		trace.TracePrintf(ctx, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", metricErr)
-	}
 	if err != nil {
-		trace.TracePrintf(ctx, nil, "Error creating a multiplexed sessions: %v", err)
 		consumer.sessionCreationFailed(ctx, ToSpannerError(err), 1, true)
 		return
 	}
 	consumer.sessionReady(ctx, &session{valid: true, client: client, id: response.Name, createTime: time.Now(), md: md, logger: sc.logger, isMultiplexed: response.Multiplexed})
-	trace.TracePrintf(ctx, nil, "Finished creating multiplexed sessions")
 }
 
 func (sc *sessionClient) sessionWithID(id string) (*session, error) {
