--- vendor/github.com/transparency-dev/tessera/append_lifecycle.go.orig	2025-09-22 08:57:07 UTC
+++ vendor/github.com/transparency-dev/tessera/append_lifecycle.go
@@ -20,7 +20,6 @@ import (
 	"fmt"
 	"net/http"
 	"os"
-	"strings"
 	"sync"
 	"sync/atomic"
 	"time"
@@ -28,11 +27,8 @@ import (
 	f_log "github.com/transparency-dev/formats/log"
 	"github.com/transparency-dev/merkle/rfc6962"
 	"github.com/transparency-dev/tessera/api/layout"
-	"github.com/transparency-dev/tessera/internal/otel"
 	"github.com/transparency-dev/tessera/internal/parse"
 	"github.com/transparency-dev/tessera/internal/witness"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 	"golang.org/x/mod/sumdb/note"
 	"k8s.io/klog/v2"
 )
@@ -55,126 +51,10 @@ var (
 )
 
 var (
-	appenderAddsTotal         metric.Int64Counter
-	appenderAddHistogram      metric.Int64Histogram
-	appenderHighestIndex      metric.Int64Gauge
-	appenderIntegratedSize    metric.Int64Gauge
-	appenderIntegrateLatency  metric.Int64Histogram
-	appenderDeadlineRemaining metric.Int64Histogram
-	appenderNextIndex         metric.Int64Gauge
-	appenderSignedSize        metric.Int64Gauge
-	appenderWitnessedSize     metric.Int64Gauge
-	appenderWitnessRequests   metric.Int64Counter
-
-	followerEntriesProcessed metric.Int64Gauge
-	followerLag              metric.Int64Gauge
-
 	// Custom histogram buckets as we're still interested in details in the 1-2s area.
 	histogramBuckets = []float64{0, 10, 50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1200, 1400, 1600, 1800, 2000, 2500, 3000, 4000, 5000, 6000, 8000, 10000}
 )
 
-func init() {
-	var err error
-
-	appenderAddsTotal, err = meter.Int64Counter(
-		"tessera.appender.add.calls",
-		metric.WithDescription("Number of calls to the appender lifecycle Add function"),
-		metric.WithUnit("{call}"))
-	if err != nil {
-		klog.Exitf("Failed to create appenderAddsTotal metric: %v", err)
-	}
-
-	appenderAddHistogram, err = meter.Int64Histogram(
-		"tessera.appender.add.duration",
-		metric.WithDescription("Duration of calls to the appender lifecycle Add function"),
-		metric.WithUnit("ms"),
-		metric.WithExplicitBucketBoundaries(histogramBuckets...))
-	if err != nil {
-		klog.Exitf("Failed to create appenderAddDuration metric: %v", err)
-	}
-
-	appenderHighestIndex, err = meter.Int64Gauge(
-		"tessera.appender.index",
-		metric.WithDescription("Highest index assigned by appender lifecycle Add function"))
-	if err != nil {
-		klog.Exitf("Failed to create appenderHighestIndex metric: %v", err)
-	}
-
-	appenderIntegratedSize, err = meter.Int64Gauge(
-		"tessera.appender.integrated.size",
-		metric.WithDescription("Size of the integrated (but not necessarily published) tree"),
-		metric.WithUnit("{entry}"))
-	if err != nil {
-		klog.Exitf("Failed to create appenderIntegratedSize metric: %v", err)
-	}
-
-	appenderIntegrateLatency, err = meter.Int64Histogram(
-		"tessera.appender.integrate.latency",
-		metric.WithDescription("Duration between an index being assigned by Add, and that index being integrated in the tree"),
-		metric.WithUnit("ms"),
-		metric.WithExplicitBucketBoundaries(histogramBuckets...))
-	if err != nil {
-		klog.Exitf("Failed to create appenderIntegrateLatency metric: %v", err)
-	}
-
-	appenderDeadlineRemaining, err = meter.Int64Histogram(
-		"tessera.appender.deadline.remaining",
-		metric.WithDescription("Duration remaining before context cancellation when appender is invoked (only set for contexts with deadline)"),
-		metric.WithUnit("ms"),
-		metric.WithExplicitBucketBoundaries(histogramBuckets...))
-	if err != nil {
-		klog.Exitf("Failed to create appenderDeadlineRemaining metric: %v", err)
-	}
-
-	appenderNextIndex, err = meter.Int64Gauge(
-		"tessera.appender.next_index",
-		metric.WithDescription("The next available index to be assigned to entries"))
-	if err != nil {
-		klog.Exitf("Failed to create appenderNextIndex metric: %v", err)
-	}
-
-	appenderSignedSize, err = meter.Int64Gauge(
-		"tessera.appender.signed.size",
-		metric.WithDescription("Size of the latest signed checkpoint"),
-		metric.WithUnit("{entry}"))
-	if err != nil {
-		klog.Exitf("Failed to create appenderSignedSize metric: %v", err)
-	}
-
-	appenderWitnessedSize, err = meter.Int64Gauge(
-		"tessera.appender.witnessed.size",
-		metric.WithDescription("Size of the latest successfully witnessed checkpoint"),
-		metric.WithUnit("{entry}"))
-	if err != nil {
-		klog.Exitf("Failed to create appenderWitnessedSize metric: %v", err)
-	}
-
-	followerEntriesProcessed, err = meter.Int64Gauge(
-		"tessera.follower.processed",
-		metric.WithDescription("Number of entries processed"),
-		metric.WithUnit("{entry}"))
-	if err != nil {
-		klog.Exitf("Failed to create followerEntriesProcessed metric: %v", err)
-	}
-
-	followerLag, err = meter.Int64Gauge(
-		"tessera.follower.lag",
-		metric.WithDescription("Number of unprocessed entries in the current integrated tree"),
-		metric.WithUnit("{entry}"))
-	if err != nil {
-		klog.Exitf("Failed to create followerLag metric: %v", err)
-	}
-
-	appenderWitnessRequests, err = meter.Int64Counter(
-		"tessera.appender.witness.requests",
-		metric.WithDescription("Number of attempts to witness a log checkpoint"),
-		metric.WithUnit("{call}"))
-	if err != nil {
-		klog.Exitf("Failed to create appenderWitnessRequests metric: %v", err)
-	}
-
-}
-
 // AddFn adds a new entry to be sequenced by the storage implementation.
 //
 // This method should quickly return an IndexFuture, which can be called to resolve to the
@@ -269,12 +149,6 @@ func NewAppender(ctx context.Context, d Driver, opts *
 	}
 	// TODO(mhutchinson): move this into the decorators
 	a.Add = func(ctx context.Context, entry *Entry) IndexFuture {
-		if deadline, ok := ctx.Deadline(); ok {
-			appenderDeadlineRemaining.Record(ctx, time.Until(deadline).Milliseconds())
-		}
-		ctx, span := tracer.Start(ctx, "tessera.Appender.Add")
-		defer span.End()
-
 		// NOTE: We memoize the returned value here so that repeated calls to the returned
 		//		 future don't result in unexpected side-effects from inner AddFn functions
 		//		 being called multiple times.
@@ -305,18 +179,15 @@ func followerStats(ctx context.Context, f Follower, si
 		case <-t.C:
 		}
 
-		n, err := f.EntriesProcessed(ctx)
+		_, err := f.EntriesProcessed(ctx)
 		if err != nil {
 			klog.Errorf("followerStats: follower %q EntriesProcessed(): %v", name, err)
 			continue
 		}
-		s, err := size(ctx)
+		_, err = size(ctx)
 		if err != nil {
 			klog.Errorf("followerStats: follower %q size(): %v", name, err)
 		}
-		attrs := metric.WithAttributes(followerNameKey.String(name))
-		followerEntriesProcessed.Record(ctx, otel.Clamp64(n), attrs)
-		followerLag.Record(ctx, otel.Clamp64(s-n), attrs)
 	}
 }
 
@@ -385,15 +256,11 @@ func (i *integrationStats) updateStats(ctx context.Con
 			klog.Errorf("IntegratedSize: %v", err)
 			continue
 		}
-		appenderIntegratedSize.Record(ctx, otel.Clamp64(s))
-		if d, ok := i.latency(s); ok {
-			appenderIntegrateLatency.Record(ctx, d.Milliseconds())
-		}
-		i, err := r.NextIndex(ctx)
+		i.latency(s)
+		_, err = r.NextIndex(ctx)
 		if err != nil {
 			klog.Errorf("NextIndex: %v", err)
 		}
-		appenderNextIndex.Record(ctx, otel.Clamp64(i))
 	}
 }
 
@@ -401,32 +268,11 @@ func (i *integrationStats) statsDecorator(delegate Add
 // metric stats.
 func (i *integrationStats) statsDecorator(delegate AddFn) AddFn {
 	return func(ctx context.Context, entry *Entry) IndexFuture {
-		start := time.Now()
 		f := delegate(ctx, entry)
 
 		return func() (Index, error) {
 			idx, err := f()
-			attr := []attribute.KeyValue{}
-			pushbackType := "" // This will be used for the pushback attribute below, empty string means no pushback
 
-			if err != nil {
-				if errors.Is(err, ErrPushback) {
-					// record the the fact there was pushback, and use the error string as the type.
-					pushbackType = err.Error()
-				} else {
-					// Just flag that it's an errored request to avoid high cardinality of attribute values.
-					// TODO(al): We might want to bucket errors into OTel status codes in the future, though.
-					attr = append(attr, attribute.String("tessera.error.type", "_OTHER"))
-				}
-			}
-
-			attr = append(attr, attribute.String("tessera.pushback", strings.ReplaceAll(pushbackType, " ", "_")))
-			attr = append(attr, attribute.Bool("tessera.duplicate", idx.IsDup))
-
-			appenderAddsTotal.Add(ctx, 1, metric.WithAttributes(attr...))
-			d := time.Since(start)
-			appenderAddHistogram.Record(ctx, d.Milliseconds(), metric.WithAttributes(attr...))
-
 			if !idx.IsDup {
 				i.sample(idx.Index)
 			}
@@ -470,7 +316,6 @@ func (t *terminator) Add(ctx context.Context, entry *E
 		for old < i.Index && !t.largestIssued.CompareAndSwap(old, i.Index) {
 			old = t.largestIssued.Load()
 		}
-		appenderHighestIndex.Record(ctx, otel.Clamp64(t.largestIssued.Load()))
 
 		return i, err
 	}
@@ -592,29 +437,19 @@ func (o AppendOptions) CheckpointPublisher(lr LogReade
 func (o AppendOptions) CheckpointPublisher(lr LogReader, httpClient *http.Client) func(context.Context, uint64, []byte) ([]byte, error) {
 	wg := witness.NewWitnessGateway(o.witnesses, httpClient, lr.ReadTile)
 	return func(ctx context.Context, size uint64, root []byte) ([]byte, error) {
-		ctx, span := tracer.Start(ctx, "tessera.CheckpointPublisher")
-		defer span.End()
-
 		cp, err := o.newCP(ctx, size, root)
 		if err != nil {
 			return nil, fmt.Errorf("newCP: %v", err)
 		}
-		appenderSignedSize.Record(ctx, otel.Clamp64(size))
 
-		witAttr := []attribute.KeyValue{}
 		cp, err = wg.Witness(ctx, cp)
 		if err != nil {
 			if !o.witnessOpts.FailOpen {
-				appenderWitnessRequests.Add(ctx, 1, metric.WithAttributes(attribute.String("error.type", "failed")))
 				return nil, err
 			}
 			klog.Warningf("WitnessGateway: failing-open despite error: %v", err)
-			witAttr = append(witAttr, attribute.String("error.type", "failed_open"))
 		}
 
-		appenderWitnessRequests.Add(ctx, 1, metric.WithAttributes(witAttr...))
-		appenderWitnessedSize.Record(ctx, otel.Clamp64(size))
-
 		return cp, nil
 	}
 }
@@ -668,9 +503,6 @@ func (o *AppendOptions) WithCheckpointSigner(s note.Si
 		}
 	}
 	o.newCP = func(ctx context.Context, size uint64, hash []byte) ([]byte, error) {
-		_, span := tracer.Start(ctx, "tessera.SignCheckpoint")
-		defer span.End()
-
 		// If we're signing a zero-sized tree, the tlog-checkpoint spec says (via RFC6962) that
 		// the root must be SHA256 of the empty string, so we'll enforce that here:
 		if size == 0 {
