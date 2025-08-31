--- vendor/cloud.google.com/go/spanner/session.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/spanner/session.go
@@ -29,14 +29,7 @@ import (
 	"sync"
 	"time"
 
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
-	"cloud.google.com/go/spanner/internal"
-	"go.opencensus.io/stats"
-	"go.opencensus.io/tag"
-	octrace "go.opencensus.io/trace"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 	"google.golang.org/grpc/codes"
 	"google.golang.org/grpc/metadata"
 )
@@ -295,10 +288,6 @@ func (s *session) ping() error {
 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 	defer cancel()
 
-	// Start parent span that doesn't record.
-	_, span := octrace.StartSpan(ctx, "cloud.google.com/go/spanner.ping", octrace.WithSampler(octrace.NeverSample()))
-	defer span.End()
-
 	// s.getID is safe even when s is invalid.
 	_, err := s.client.ExecuteSql(contextWithOutgoingMetadata(ctx, s.md, true), &sppb.ExecuteSqlRequest{
 		Session: s.getID(),
@@ -658,9 +647,6 @@ type sessionPool struct {
 	// sessions checked out of the pool during the last 10 minutes.
 	mw *maintenanceWindow
 
-	// tagMap is a map of all tags that are associated with the emitted metrics.
-	tagMap *tag.Map
-
 	// indicates the number of leaked sessions removed from the session pool.
 	// This is valid only when ActionOnInactiveTransaction is WarnAndClose or ActionOnInactiveTransaction is Close in InactiveTransactionRemovalOptions.
 	numOfLeakedSessionsRemoved uint64
@@ -720,21 +706,12 @@ func newSessionPool(sc *sessionClient, config SessionP
 		enableMultiplexSession:   config.enableMultiplexSession,
 	}
 
-	_, instance, database, err := parseDatabaseName(sc.database)
+	_, _, _, err := parseDatabaseName(sc.database)
 	if err != nil {
 		return nil, err
 	}
 	// Errors should not prevent initializing the session pool.
-	ctx, err := tag.New(context.Background(),
-		tag.Upsert(tagKeyClientID, sc.id),
-		tag.Upsert(tagKeyDatabase, database),
-		tag.Upsert(tagKeyInstance, instance),
-		tag.Upsert(tagKeyLibVersion, internal.Version),
-	)
-	if err != nil {
-		logf(pool.sc.logger, "Failed to create tag map, error: %v", err)
-	}
-	pool.tagMap = tag.FromContext(ctx)
+	ctx := context.Background()
 
 	// On GCE VM, within the same region an healthcheck ping takes on average
 	// 10ms to finish, given a 5 minutes interval and 10 healthcheck workers, a
@@ -767,7 +744,6 @@ func newSessionPool(sc *sessionClient, config SessionP
 			}
 		}()
 	}
-	pool.recordStat(context.Background(), MaxAllowedSessionsCount, int64(config.MaxOpened))
 
 	err = registerSessionPoolOTMetrics(pool)
 	if err != nil {
@@ -778,33 +754,6 @@ func newSessionPool(sc *sessionClient, config SessionP
 	return pool, nil
 }
 
-func (p *sessionPool) recordStat(ctx context.Context, m *stats.Int64Measure, n int64, tags ...tag.Tag) {
-	ctx = tag.NewContext(ctx, p.tagMap)
-	mutators := make([]tag.Mutator, len(tags))
-	for i, t := range tags {
-		mutators[i] = tag.Upsert(t.Key, t.Value)
-	}
-	ctx, err := tag.New(ctx, mutators...)
-	if err != nil {
-		logf(p.sc.logger, "Failed to tag metrics, error: %v", err)
-	}
-	recordStat(ctx, m, n)
-}
-
-type recordOTStatOption struct {
-	attr []attribute.KeyValue
-}
-
-func (p *sessionPool) recordOTStat(ctx context.Context, m metric.Int64Counter, val int64, option recordOTStatOption) {
-	if m != nil {
-		attrs := p.otConfig.attributeMap
-		if len(option.attr) > 0 {
-			attrs = option.attr
-		}
-		m.Add(ctx, val, metric.WithAttributes(attrs...))
-	}
-}
-
 func (p *sessionPool) getRatioOfSessionsInUseLocked() float64 {
 	maxSessions := p.MaxOpened
 	if maxSessions == 0 {
@@ -912,7 +861,6 @@ func (p *sessionPool) growPoolLocked(numSessions uint6
 	// Take budget before the actual session creation.
 	numSessions = minUint64(numSessions, math.MaxInt32)
 	p.numOpened += uint64(numSessions)
-	p.recordStat(context.Background(), OpenSessionCount, int64(p.numOpened))
 	p.createReqs += uint64(numSessions)
 	// Asynchronously create a batch of sessions for the pool.
 	return p.sc.batchCreateSessions(int32(numSessions), distributeOverChannels, p)
@@ -962,8 +910,6 @@ func (p *sessionPool) sessionReady(ctx context.Context
 		s.pool = p
 		p.multiplexedSession = s
 		p.multiplexedSessionCreationError = nil
-		p.recordStat(context.Background(), OpenSessionCount, int64(1), tag.Tag{Key: tagKeyIsMultiplexed, Value: "true"})
-		p.recordStat(context.Background(), SessionsCount, 1, tagNumSessions, tag.Tag{Key: tagKeyIsMultiplexed, Value: "true"})
 		// either notify the waiting goroutine or skip if no one is waiting
 		select {
 		case p.mayGetMultiplexedSession <- true:
@@ -1015,7 +961,6 @@ func (p *sessionPool) sessionCreationFailed(ctx contex
 			}
 			return
 		}
-		p.recordStat(context.Background(), OpenSessionCount, int64(0), tag.Tag{Key: tagKeyIsMultiplexed, Value: "true"})
 		p.multiplexedSessionCreationError = err
 		select {
 		case p.mayGetMultiplexedSession <- true:
@@ -1026,7 +971,6 @@ func (p *sessionPool) sessionCreationFailed(ctx contex
 	}
 	p.createReqs -= uint64(numSessions)
 	p.numOpened -= uint64(numSessions)
-	p.recordStat(context.Background(), OpenSessionCount, int64(p.numOpened), tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
 	// Notify other waiters blocking on session creation.
 	p.sessionCreationError = err
 	close(p.mayGetSession)
@@ -1057,12 +1001,6 @@ func (p *sessionPool) close(ctx context.Context) {
 		return
 	}
 	p.valid = false
-	if p.otConfig != nil && p.otConfig.otMetricRegistration != nil {
-		err := p.otConfig.otMetricRegistration.Unregister()
-		if err != nil {
-			logf(p.sc.logger, "Failed to unregister callback from the OpenTelemetry meter, error : %v", err)
-		}
-	}
 	p.mu.Unlock()
 	p.hc.close()
 	// destroy all the sessions
@@ -1202,7 +1140,6 @@ func (p *sessionPool) take(ctx context.Context) (*sess
 // take returns a cached session if there are available ones; if there isn't
 // any, it tries to allocate a new one.
 func (p *sessionPool) take(ctx context.Context) (*sessionHandle, error) {
-	trace.TracePrintf(ctx, nil, "Acquiring a session")
 	for {
 		var s *session
 
@@ -1215,8 +1152,6 @@ func (p *sessionPool) take(ctx context.Context) (*sess
 			// Idle sessions are available, get one from the top of the idle
 			// list.
 			s = p.idleList.Remove(p.idleList.Front()).(*session)
-			trace.TracePrintf(ctx, map[string]interface{}{"sessionID": s.getID()},
-				"Acquired session")
 			p.decNumSessionsLocked(ctx)
 		}
 		if s != nil {
@@ -1249,14 +1184,8 @@ func (p *sessionPool) take(ctx context.Context) (*sess
 		p.numWaiters++
 		mayGetSession := p.mayGetSession
 		p.mu.Unlock()
-		trace.TracePrintf(ctx, nil, "Waiting for read-only session to become available")
 		select {
 		case <-ctx.Done():
-			trace.TracePrintf(ctx, nil, "Context done waiting for session")
-			p.recordStat(ctx, GetSessionTimeoutsCount, 1, tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
-			if p.otConfig != nil {
-				p.recordOTStat(ctx, p.otConfig.getSessionTimeoutsCount, 1, recordOTStatOption{attr: p.otConfig.attributeMapWithoutMultiplexed})
-			}
 			p.mu.Lock()
 			p.numWaiters--
 			p.mu.Unlock()
@@ -1265,7 +1194,6 @@ func (p *sessionPool) take(ctx context.Context) (*sess
 			p.mu.Lock()
 			p.numWaiters--
 			if p.sessionCreationError != nil {
-				trace.TracePrintf(ctx, nil, "Error creating session: %v", p.sessionCreationError)
 				err := p.sessionCreationError
 				p.mu.Unlock()
 				return nil, err
@@ -1278,7 +1206,6 @@ func (p *sessionPool) takeMultiplexed(ctx context.Cont
 // takeMultiplexed returns a cached session if there is available one; if there isn't
 // any, it tries to allocate a new one.
 func (p *sessionPool) takeMultiplexed(ctx context.Context) (*sessionHandle, error) {
-	trace.TracePrintf(ctx, nil, "Acquiring a multiplexed session")
 	for {
 		var s *session
 		p.mu.Lock()
@@ -1294,8 +1221,6 @@ func (p *sessionPool) takeMultiplexed(ctx context.Cont
 		if p.multiplexedSession != nil {
 			// Multiplexed session is available, get it.
 			s = p.multiplexedSession
-			trace.TracePrintf(ctx, map[string]interface{}{"sessionID": s.getID()},
-				"Acquired multiplexed session")
 			p.mu.Unlock()
 			p.incNumMultiplexedInUse(ctx)
 			return p.newSessionHandle(s), nil
@@ -1305,16 +1230,10 @@ func (p *sessionPool) takeMultiplexed(ctx context.Cont
 		p.multiplexedSessionReq <- muxSessionCreateRequest{force: false, ctx: ctx}
 		select {
 		case <-ctx.Done():
-			trace.TracePrintf(ctx, nil, "Context done waiting for multiplexed session")
-			p.recordStat(ctx, GetSessionTimeoutsCount, 1, tag.Tag{Key: tagKeyIsMultiplexed, Value: "true"})
-			if p.otConfig != nil {
-				p.recordOTStat(ctx, p.otConfig.getSessionTimeoutsCount, 1, recordOTStatOption{attr: p.otConfig.attributeMapWithMultiplexed})
-			}
 			return nil, p.errGetSessionTimeout(ctx)
 		case <-mayGetSession: // Block until multiplexed session is created.
 			p.mu.Lock()
 			if p.multiplexedSessionCreationError != nil {
-				trace.TracePrintf(ctx, nil, "Error creating multiplexed session: %v", p.multiplexedSessionCreationError)
 				err := p.multiplexedSessionCreationError
 				if isUnimplementedError(err) {
 					logf(p.sc.logger, "Multiplexed session is not enabled on this project, continuing with regular sessions")
@@ -1386,7 +1305,6 @@ func (p *sessionPool) remove(s *session, isExpire bool
 		if wasInUse {
 			p.decNumInUseLocked(ctx)
 		}
-		p.recordStat(ctx, OpenSessionCount, int64(p.numOpened))
 		close(p.mayGetSession)
 		p.mayGetSession = make(chan struct{})
 		return true
@@ -1406,22 +1324,12 @@ func (p *sessionPool) incNumInUseLocked(ctx context.Co
 
 func (p *sessionPool) incNumInUseLocked(ctx context.Context) {
 	p.numInUse++
-	p.recordStat(ctx, SessionsCount, int64(p.numInUse), tagNumInUseSessions, tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
-	p.recordStat(ctx, AcquiredSessionsCount, 1, tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
-	if p.otConfig != nil {
-		p.recordOTStat(ctx, p.otConfig.acquiredSessionsCount, 1, recordOTStatOption{attr: p.otConfig.attributeMapWithoutMultiplexed})
-	}
 	if p.numInUse > p.maxNumInUse {
 		p.maxNumInUse = p.numInUse
-		p.recordStat(ctx, MaxInUseSessionsCount, int64(p.maxNumInUse), tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
 	}
 }
 
 func (p *sessionPool) incNumMultiplexedInUse(ctx context.Context) {
-	p.recordStat(ctx, AcquiredSessionsCount, 1, tag.Tag{Key: tagKeyIsMultiplexed, Value: "true"})
-	if p.otConfig != nil {
-		p.recordOTStat(ctx, p.otConfig.acquiredSessionsCount, 1, recordOTStatOption{attr: p.otConfig.attributeMapWithMultiplexed})
-	}
 }
 
 func (p *sessionPool) decNumInUseLocked(ctx context.Context) {
@@ -1431,28 +1339,17 @@ func (p *sessionPool) decNumInUseLocked(ctx context.Co
 		logf(p.sc.logger, "Number of sessions in use is negative, resetting it to currSessionsCheckedOutLocked. Stack trace: %s", string(debug.Stack()))
 		p.numInUse = p.currSessionsCheckedOutLocked()
 	}
-	p.recordStat(ctx, SessionsCount, int64(p.numInUse), tagNumInUseSessions, tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
-	p.recordStat(ctx, ReleasedSessionsCount, 1, tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
-	if p.otConfig != nil {
-		p.recordOTStat(ctx, p.otConfig.releasedSessionsCount, 1, recordOTStatOption{attr: p.otConfig.attributeMapWithoutMultiplexed})
-	}
 }
 
 func (p *sessionPool) decNumMultiplexedInUseLocked(ctx context.Context) {
-	p.recordStat(ctx, ReleasedSessionsCount, 1, tag.Tag{Key: tagKeyIsMultiplexed, Value: "true"})
-	if p.otConfig != nil {
-		p.recordOTStat(ctx, p.otConfig.releasedSessionsCount, 1, recordOTStatOption{attr: p.otConfig.attributeMapWithMultiplexed})
-	}
 }
 
 func (p *sessionPool) incNumSessionsLocked(ctx context.Context) {
 	p.numSessions++
-	p.recordStat(ctx, SessionsCount, int64(p.numSessions), tagNumSessions)
 }
 
 func (p *sessionPool) decNumSessionsLocked(ctx context.Context) {
 	p.numSessions--
-	p.recordStat(ctx, SessionsCount, int64(p.numSessions), tagNumSessions)
 }
 
 // hcHeap implements heap.Interface. It is used to create the priority queue for
@@ -1802,7 +1699,6 @@ func (hc *healthChecker) maintainer() {
 		now := time.Now()
 		if now.After(hc.pool.lastResetTime.Add(10 * time.Minute)) {
 			hc.pool.maxNumInUse = hc.pool.numInUse
-			hc.pool.recordStat(context.Background(), MaxInUseSessionsCount, int64(hc.pool.maxNumInUse), tag.Tag{Key: tagKeyIsMultiplexed, Value: "false"})
 			hc.pool.lastResetTime = now
 		}
 		hc.pool.mu.Unlock()
