--- vendor/cloud.google.com/go/storage/grpc_metrics.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/grpc_metrics.go
@@ -16,22 +16,9 @@ import (
 
 import (
 	"context"
-	"errors"
-	"fmt"
-	"strings"
 	"time"
 
-	mexporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
-	"github.com/google/uuid"
-	"go.opentelemetry.io/contrib/detectors/gcp"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/sdk/metric"
-	"go.opentelemetry.io/otel/sdk/metric/metricdata"
-	"go.opentelemetry.io/otel/sdk/resource"
 	"google.golang.org/api/option"
-	"google.golang.org/grpc"
-	"google.golang.org/grpc/experimental/stats"
-	"google.golang.org/grpc/stats/opentelemetry"
 )
 
 const (
@@ -47,77 +34,11 @@ type storageMonitoredResource struct {
 	instance      string
 	cloudPlatform string
 	host          string
-	resource      *resource.Resource
 }
 
-func (smr *storageMonitoredResource) exporter() (metric.Exporter, error) {
-	exporter, err := mexporter.New(
-		mexporter.WithProjectID(smr.project),
-		mexporter.WithMetricDescriptorTypeFormatter(metricFormatter),
-		mexporter.WithCreateServiceTimeSeries(),
-		mexporter.WithMonitoredResourceDescription(monitoredResourceName, []string{"project_id", "location", "cloud_platform", "host_id", "instance_id", "api"}),
-	)
-	if err != nil {
-		return nil, fmt.Errorf("storage: creating metrics exporter: %w", err)
-	}
-	return exporter, nil
-}
-
-func newStorageMonitoredResource(ctx context.Context, project, api string, opts ...resource.Option) (*storageMonitoredResource, error) {
-	detectedAttrs, err := resource.New(ctx, opts...)
-	if err != nil {
-		return nil, err
-	}
-	smr := &storageMonitoredResource{
-		instance: uuid.New().String(),
-		api:      api,
-		project:  project,
-	}
-	s := detectedAttrs.Set()
-	// Attempt to use resource detector project id if project id wasn't
-	// identified using ADC as a last resort. Otherwise metrics cannot be started.
-	if p, present := s.Value("cloud.account.id"); present && smr.project == "" {
-		smr.project = p.AsString()
-	} else if !present && smr.project == "" {
-		return nil, errors.New("google cloud project is required to start client-side metrics")
-	}
-	if v, ok := s.Value("cloud.region"); ok {
-		smr.location = v.AsString()
-	} else {
-		smr.location = "global"
-	}
-	if v, ok := s.Value("cloud.platform"); ok {
-		smr.cloudPlatform = v.AsString()
-	} else {
-		smr.cloudPlatform = "unknown"
-	}
-	if v, ok := s.Value("host.id"); ok {
-		smr.host = v.AsString()
-	} else if v, ok := s.Value("faas.id"); ok {
-		smr.host = v.AsString()
-	} else {
-		smr.host = "unknown"
-	}
-	smr.resource, err = resource.New(ctx, resource.WithAttributes([]attribute.KeyValue{
-		{Key: "gcp.resource_type", Value: attribute.StringValue(monitoredResourceName)},
-		{Key: "project_id", Value: attribute.StringValue(smr.project)},
-		{Key: "api", Value: attribute.StringValue(smr.api)},
-		{Key: "instance_id", Value: attribute.StringValue(smr.instance)},
-		{Key: "location", Value: attribute.StringValue(smr.location)},
-		{Key: "cloud_platform", Value: attribute.StringValue(smr.cloudPlatform)},
-		{Key: "host_id", Value: attribute.StringValue(smr.host)},
-	}...))
-	if err != nil {
-		return nil, err
-	}
-	return smr, nil
-}
-
 type metricsContext struct {
 	// client options passed to gRPC channels
 	clientOpts []option.ClientOption
-	// instance of metric reader used by gRPC client-side metrics
-	provider *metric.MeterProvider
 	// clean func to call when closing gRPC client
 	close func()
 }
@@ -125,110 +46,20 @@ type metricsConfig struct {
 type metricsConfig struct {
 	project         string
 	interval        time.Duration
-	customExporter  *metric.Exporter
-	manualReader    *metric.ManualReader // used by tests
 	disableExporter bool                 // used by tests disables exports
-	resourceOpts    []resource.Option    // used by tests
 }
 
 func newGRPCMetricContext(ctx context.Context, cfg metricsConfig) (*metricsContext, error) {
-	var exporter metric.Exporter
-	meterOpts := []metric.Option{}
-	if cfg.customExporter == nil {
-		var ropts []resource.Option
-		if cfg.resourceOpts != nil {
-			ropts = cfg.resourceOpts
-		} else {
-			ropts = []resource.Option{resource.WithDetectors(gcp.NewDetector())}
-		}
-		smr, err := newStorageMonitoredResource(ctx, cfg.project, "grpc", ropts...)
-		if err != nil {
-			return nil, err
-		}
-		exporter, err = smr.exporter()
-		if err != nil {
-			return nil, err
-		}
-		meterOpts = append(meterOpts, metric.WithResource(smr.resource))
-	} else {
-		exporter = *cfg.customExporter
-	}
-	interval := time.Minute
-	if cfg.interval > 0 {
-		interval = cfg.interval
-	}
-	meterOpts = append(meterOpts,
-		// Metric views update histogram boundaries to be relevant to GCS
-		// otherwise default OTel histogram boundaries are used.
-		metric.WithView(
-			createHistogramView("grpc.client.attempt.duration", latencyHistogramBoundaries()),
-			createHistogramView("grpc.client.attempt.rcvd_total_compressed_message_size", sizeHistogramBoundaries()),
-			createHistogramView("grpc.client.attempt.sent_total_compressed_message_size", sizeHistogramBoundaries())),
-	)
-	if cfg.manualReader != nil {
-		meterOpts = append(meterOpts, metric.WithReader(cfg.manualReader))
-	}
-	if !cfg.disableExporter {
-		meterOpts = append(meterOpts, metric.WithReader(
-			metric.NewPeriodicReader(&exporterLogSuppressor{Exporter: exporter}, metric.WithInterval(interval))))
-	}
-	provider := metric.NewMeterProvider(meterOpts...)
-	mo := opentelemetry.MetricsOptions{
-		MeterProvider: provider,
-		Metrics: stats.NewMetrics(
-			"grpc.client.attempt.started",
-			"grpc.client.attempt.duration",
-			"grpc.client.attempt.sent_total_compressed_message_size",
-			"grpc.client.attempt.rcvd_total_compressed_message_size",
-			"grpc.client.call.duration",
-			"grpc.lb.wrr.rr_fallback",
-			"grpc.lb.wrr.endpoint_weight_not_yet_usable",
-			"grpc.lb.wrr.endpoint_weight_stale",
-			"grpc.lb.wrr.endpoint_weights",
-			"grpc.lb.rls.cache_entries",
-			"grpc.lb.rls.cache_size",
-			"grpc.lb.rls.default_target_picks",
-			"grpc.lb.rls.target_picks",
-			"grpc.lb.rls.failed_picks",
-		),
-		OptionalLabels: []string{"grpc.lb.locality"},
-	}
-	opts := []option.ClientOption{
-		option.WithGRPCDialOption(
-			opentelemetry.DialOption(opentelemetry.Options{MetricsOptions: mo})),
-		option.WithGRPCDialOption(
-			grpc.WithDefaultCallOptions(grpc.StaticMethodCallOption{})),
-	}
 	return &metricsContext{
-		clientOpts: opts,
-		provider:   provider,
-		close: func() {
-			provider.Shutdown(ctx)
-		},
 	}, nil
 }
 
 // Silences permission errors after initial error is emitted to prevent
 // chatty logs.
 type exporterLogSuppressor struct {
-	metric.Exporter
 	emittedFailure bool
 }
 
-// Implements OTel SDK metric.Exporter interface to prevent noisy logs from
-// lack of credentials after initial failure.
-// https://pkg.go.dev/go.opentelemetry.io/otel/sdk/metric@v1.28.0#Exporter
-func (e *exporterLogSuppressor) Export(ctx context.Context, rm *metricdata.ResourceMetrics) error {
-	if err := e.Exporter.Export(ctx, rm); err != nil && !e.emittedFailure {
-		if strings.Contains(err.Error(), "PermissionDenied") {
-			e.emittedFailure = true
-			return fmt.Errorf("gRPC metrics failed due permission issue: %w", err)
-		}
-		return err
-	}
-	return nil
-}
-
 func latencyHistogramBoundaries() []float64 {
 	boundaries := []float64{}
 	boundary := 0.0
@@ -266,18 +97,4 @@ func sizeHistogramBoundaries() []float64 {
 		}
 	}
 	return boundaries
-}
-
-func createHistogramView(name string, boundaries []float64) metric.View {
-	return metric.NewView(metric.Instrument{
-		Name: name,
-		Kind: metric.InstrumentKindHistogram,
-	}, metric.Stream{
-		Name:        name,
-		Aggregation: metric.AggregationExplicitBucketHistogram{Boundaries: boundaries},
-	})
-}
-
-func metricFormatter(m metricdata.Metrics) string {
-	return metricPrefix + strings.ReplaceAll(string(m.Name), ".", "/")
 }
