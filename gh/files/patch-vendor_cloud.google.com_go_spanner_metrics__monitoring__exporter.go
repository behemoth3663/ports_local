--- vendor/cloud.google.com/go/spanner/metrics_monitoring_exporter.go.orig	2025-07-17 21:08:43 UTC
+++ vendor/cloud.google.com/go/spanner/metrics_monitoring_exporter.go
@@ -22,26 +22,11 @@ import (
 	"context"
 	"errors"
 	"fmt"
-	"math"
-	"reflect"
-	"strings"
 	"sync"
 	"time"
 
 	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
 	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
-	"cloud.google.com/go/spanner/internal"
-	"github.com/googleapis/gax-go/v2/callctx"
-	"go.opentelemetry.io/otel/attribute"
-	otelmetric "go.opentelemetry.io/otel/sdk/metric"
-	otelmetricdata "go.opentelemetry.io/otel/sdk/metric/metricdata"
-	"google.golang.org/api/option"
-	"google.golang.org/genproto/googleapis/api/distribution"
-	googlemetricpb "google.golang.org/genproto/googleapis/api/metric"
-	monitoredrespb "google.golang.org/genproto/googleapis/api/monitoredres"
-	"google.golang.org/grpc/codes"
-	"google.golang.org/grpc/encoding/gzip"
-	"google.golang.org/grpc/status"
 	"google.golang.org/protobuf/types/known/timestamppb"
 )
 
@@ -92,7 +77,6 @@ type monitoringExporter struct {
 type monitoringExporter struct {
 	projectID        string
 	compression      string
-	clientAttributes []attribute.KeyValue
 	shutdown         chan struct{}
 	client           *monitoring.MetricClient
 	shutdownOnce     sync.Once
@@ -102,21 +86,6 @@ type monitoringExporter struct {
 	lastExportedAt time.Time
 }
 
-func newMonitoringExporter(ctx context.Context, project, compression string, clientAttributes []attribute.KeyValue, opts ...option.ClientOption) (*monitoringExporter, error) {
-	client, err := monitoring.NewMetricClient(ctx, opts...)
-	if err != nil {
-		return nil, err
-	}
-	return &monitoringExporter{
-		projectID:        project,
-		compression:      compression,
-		clientAttributes: clientAttributes,
-		lastExportedAt:   time.Now().Add(-time.Minute),
-		client:           client,
-		shutdown:         make(chan struct{}),
-	}, nil
-}
-
 func (me *monitoringExporter) stop() {
 	// stop the exporter if last export happens within half-time of default sample period
 	me.mu.Lock()
@@ -139,210 +108,6 @@ func (e *monitoringExporter) Shutdown(ctx context.Cont
 	return err
 }
 
-// Export exports OpenTelemetry Metrics to Google Cloud Monitoring.
-func (me *monitoringExporter) Export(ctx context.Context, rm *otelmetricdata.ResourceMetrics) error {
-	select {
-	case <-me.shutdown:
-		return errShutdown
-	default:
-	}
-
-	me.mu.Lock()
-	if me.stopExport {
-		me.mu.Unlock()
-		return nil
-	}
-	me.mu.Unlock()
-	return me.exportTimeSeries(ctx, rm)
-}
-
-// Temporality returns the Temporality to use for an instrument kind.
-func (me *monitoringExporter) Temporality(ik otelmetric.InstrumentKind) otelmetricdata.Temporality {
-	return otelmetricdata.CumulativeTemporality
-}
-
-// Aggregation returns the Aggregation to use for an instrument kind.
-func (me *monitoringExporter) Aggregation(ik otelmetric.InstrumentKind) otelmetric.Aggregation {
-	return otelmetric.DefaultAggregationSelector(ik)
-}
-
-// exportTimeSeries create TimeSeries from the records in cps.
-// res should be the common resource among all TimeSeries, such as instance id, application name and so on.
-func (me *monitoringExporter) exportTimeSeries(ctx context.Context, rm *otelmetricdata.ResourceMetrics) error {
-	tss, err := me.recordsToTimeSeriesPbs(rm)
-	if len(tss) == 0 {
-		return err
-	}
-	name := fmt.Sprintf("projects/%s", me.projectID)
-	ctx = callctx.SetHeaders(ctx, "x-goog-api-client", "gccl/"+internal.Version)
-	if me.compression == gzip.Name {
-		ctx = callctx.SetHeaders(ctx, requestsCompressionHeader, gzip.Name)
-	}
-	errs := []error{err}
-	for i := 0; i < len(tss); i += sendBatchSize {
-		j := i + sendBatchSize
-		if j >= len(tss) {
-			j = len(tss)
-		}
-		req := &monitoringpb.CreateTimeSeriesRequest{
-			Name:       name,
-			TimeSeries: tss[i:j],
-		}
-		err = me.client.CreateServiceTimeSeries(ctx, req)
-		if err != nil {
-			if status.Code(err) == codes.PermissionDenied {
-				err = fmt.Errorf("%w Need monitoring metric writer permission on project=%s. Follow https://cloud.google.com/spanner/docs/view-manage-client-side-metrics#access-client-side-metrics to set up permissions",
-					err, me.projectID)
-			}
-		}
-		errs = append(errs, err)
-	}
-
-	me.mu.Lock()
-	me.lastExportedAt = time.Now()
-	me.mu.Unlock()
-	return errors.Join(errs...)
-}
-
-// recordToMetricAndMonitoredResourcePbs converts data from records to Metric and Monitored resource proto type for Cloud Monitoring.
-func (me *monitoringExporter) recordToMetricAndMonitoredResourcePbs(metrics otelmetricdata.Metrics, attributes attribute.Set) (*googlemetricpb.Metric, *monitoredrespb.MonitoredResource) {
-	mr := &monitoredrespb.MonitoredResource{
-		Type:   spannerResourceType,
-		Labels: map[string]string{},
-	}
-	labels := make(map[string]string)
-	addAttributes := func(attr *attribute.Set) {
-		iter := attr.Iter()
-		for iter.Next() {
-			kv := iter.Attribute()
-			// Replace metric label names by converting "." to "_" since Cloud Monitoring does not
-			// support labels containing "."
-			labelKey := strings.Replace(string(kv.Key), ".", "_", -1)
-			if _, isResLabel := monitoredResLabelsSet[labelKey]; isResLabel {
-				mr.Labels[labelKey] = kv.Value.Emit()
-			} else {
-				if _, ok := allowedMetricLabels[string(kv.Key)]; ok {
-					labels[labelKey] = kv.Value.Emit()
-				}
-			}
-		}
-		for _, label := range me.clientAttributes {
-			if _, isResLabel := monitoredResLabelsSet[string(label.Key)]; isResLabel {
-				mr.Labels[string(label.Key)] = label.Value.Emit()
-			} else {
-				labels[string(label.Key)] = label.Value.Emit()
-			}
-		}
-	}
-	metricName := metrics.Name
-	if !strings.HasPrefix(metricName, nativeMetricsPrefix) {
-		metricName = nativeMetricsPrefix + strings.Replace(metricName, ".", "/", -1)
-	}
-	addAttributes(&attributes)
-	return &googlemetricpb.Metric{
-		Type:   metricName,
-		Labels: labels,
-	}, mr
-}
-
-func (me *monitoringExporter) recordsToTimeSeriesPbs(rm *otelmetricdata.ResourceMetrics) ([]*monitoringpb.TimeSeries, error) {
-	var (
-		tss  []*monitoringpb.TimeSeries
-		errs []error
-	)
-	for _, scope := range rm.ScopeMetrics {
-		if !(scope.Scope.Name == builtInMetricsMeterName || scope.Scope.Name == grpcMetricMeterName) {
-			continue
-		}
-		for _, metrics := range scope.Metrics {
-			ts, err := me.recordToTimeSeriesPb(metrics)
-			errs = append(errs, err)
-			tss = append(tss, ts...)
-		}
-	}
-
-	return tss, errors.Join(errs...)
-}
-
-// recordToTimeSeriesPb converts record to TimeSeries proto type with common resource.
-// ref. https://cloud.google.com/monitoring/api/ref_v3/rest/v3/TimeSeries
-func (me *monitoringExporter) recordToTimeSeriesPb(m otelmetricdata.Metrics) ([]*monitoringpb.TimeSeries, error) {
-	var tss []*monitoringpb.TimeSeries
-	var errs []error
-	if m.Data == nil {
-		return nil, nil
-	}
-	switch a := m.Data.(type) {
-	case otelmetricdata.Histogram[float64]:
-		for _, point := range a.DataPoints {
-			metric, mr := me.recordToMetricAndMonitoredResourcePbs(m, point.Attributes)
-			ts, err := histogramToTimeSeries(point, m, mr)
-			if err != nil {
-				errs = append(errs, err)
-				continue
-			}
-			ts.Metric = metric
-			tss = append(tss, ts)
-		}
-	case otelmetricdata.Sum[int64]:
-		for _, point := range a.DataPoints {
-			metric, mr := me.recordToMetricAndMonitoredResourcePbs(m, point.Attributes)
-			var ts *monitoringpb.TimeSeries
-			var err error
-			ts, err = sumToTimeSeries[int64](point, m, mr)
-			if err != nil {
-				errs = append(errs, err)
-				continue
-			}
-			ts.Metric = metric
-			tss = append(tss, ts)
-		}
-	default:
-		errs = append(errs, errUnexpectedAggregationKind{kind: reflect.TypeOf(m.Data).String()})
-	}
-	return tss, errors.Join(errs...)
-}
-
-func sumToTimeSeries[N int64 | float64](point otelmetricdata.DataPoint[N], metrics otelmetricdata.Metrics, mr *monitoredrespb.MonitoredResource) (*monitoringpb.TimeSeries, error) {
-	interval, err := toNonemptyTimeIntervalpb(point.StartTime, point.Time)
-	if err != nil {
-		return nil, err
-	}
-	value, valueType := numberDataPointToValue[N](point)
-	return &monitoringpb.TimeSeries{
-		Resource:   mr,
-		Unit:       string(metrics.Unit),
-		MetricKind: googlemetricpb.MetricDescriptor_CUMULATIVE,
-		ValueType:  valueType,
-		Points: []*monitoringpb.Point{{
-			Interval: interval,
-			Value:    value,
-		}},
-	}, nil
-}
-
-func histogramToTimeSeries[N int64 | float64](point otelmetricdata.HistogramDataPoint[N], metrics otelmetricdata.Metrics, mr *monitoredrespb.MonitoredResource) (*monitoringpb.TimeSeries, error) {
-	interval, err := toNonemptyTimeIntervalpb(point.StartTime, point.Time)
-	if err != nil {
-		return nil, err
-	}
-	distributionValue := histToDistribution(point)
-	return &monitoringpb.TimeSeries{
-		Resource:   mr,
-		Unit:       string(metrics.Unit),
-		MetricKind: googlemetricpb.MetricDescriptor_CUMULATIVE,
-		ValueType:  googlemetricpb.MetricDescriptor_DISTRIBUTION,
-		Points: []*monitoringpb.Point{{
-			Interval: interval,
-			Value: &monitoringpb.TypedValue{
-				Value: &monitoringpb.TypedValue_DistributionValue{
-					DistributionValue: distributionValue,
-				},
-			},
-		}},
-	}, nil
-}
-
 func toNonemptyTimeIntervalpb(start, end time.Time) (*monitoringpb.TimeInterval, error) {
 	// The end time of a new interval must be at least a millisecond after the end time of the
 	// previous interval, for all non-gauge types.
@@ -364,46 +129,4 @@ func toNonemptyTimeIntervalpb(start, end time.Time) (*
 		StartTime: startpb,
 		EndTime:   endpb,
 	}, nil
-}
-
-func histToDistribution[N int64 | float64](hist otelmetricdata.HistogramDataPoint[N]) *distribution.Distribution {
-	counts := make([]int64, len(hist.BucketCounts))
-	for i, v := range hist.BucketCounts {
-		counts[i] = int64(v)
-	}
-	var mean float64
-	if !math.IsNaN(float64(hist.Sum)) && hist.Count > 0 { // Avoid divide-by-zero
-		mean = float64(hist.Sum) / float64(hist.Count)
-	}
-	return &distribution.Distribution{
-		Count:        int64(hist.Count),
-		Mean:         mean,
-		BucketCounts: counts,
-		BucketOptions: &distribution.Distribution_BucketOptions{
-			Options: &distribution.Distribution_BucketOptions_ExplicitBuckets{
-				ExplicitBuckets: &distribution.Distribution_BucketOptions_Explicit{
-					Bounds: hist.Bounds,
-				},
-			},
-		},
-	}
-}
-
-func numberDataPointToValue[N int64 | float64](
-	point otelmetricdata.DataPoint[N],
-) (*monitoringpb.TypedValue, googlemetricpb.MetricDescriptor_ValueType) {
-	switch v := any(point.Value).(type) {
-	case int64:
-		return &monitoringpb.TypedValue{Value: &monitoringpb.TypedValue_Int64Value{
-				Int64Value: v,
-			}},
-			googlemetricpb.MetricDescriptor_INT64
-	case float64:
-		return &monitoringpb.TypedValue{Value: &monitoringpb.TypedValue_DoubleValue{
-				DoubleValue: v,
-			}},
-			googlemetricpb.MetricDescriptor_DOUBLE
-	}
-	// It is impossible to reach this statement
-	return nil, googlemetricpb.MetricDescriptor_INT64
 }
