--- vendor/cloud.google.com/go/spanner/metrics.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/spanner/metrics.go
@@ -18,27 +18,15 @@ import (
 
 import (
 	"context"
-	"errors"
 	"fmt"
 	"hash/fnv"
-	"log"
 	"os"
 	"strconv"
-	"strings"
 	"time"
 
 	"github.com/google/uuid"
-	"go.opentelemetry.io/contrib/detectors/gcp"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
-	"go.opentelemetry.io/otel/metric/noop"
-	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
-	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
 	"google.golang.org/api/option"
-	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
-	"google.golang.org/grpc/experimental/stats"
-	"google.golang.org/grpc/stats/opentelemetry"
 	"google.golang.org/grpc/status"
 
 	"cloud.google.com/go/spanner/internal"
@@ -187,15 +175,6 @@ var (
 	}
 
 	detectClientLocation = func(ctx context.Context) string {
-		resource, err := gcp.NewDetector().Detect(ctx)
-		if err != nil {
-			return "global"
-		}
-		for _, attr := range resource.Attributes() {
-			if attr.Key == semconv.CloudRegionKey {
-				return attr.Value.AsString()
-			}
-		}
 		// If region is not found, return global
 		return "global"
 	}
@@ -238,201 +217,8 @@ type builtinMetricsTracerFactory struct {
 
 	// client options passed to gRPC channels
 	clientOpts []option.ClientOption
-	// clientAttributes are attributes specific to a client instance that do not change across different function calls on the client.
-	clientAttributes []attribute.KeyValue
-
-	// Metrics instruments
-	operationLatencies metric.Float64Histogram // Histogram for operation latencies.
-	attemptLatencies   metric.Float64Histogram // Histogram for attempt latencies.
-	gfeLatencies       metric.Float64Histogram // Latency between Google's network receiving an RPC and reading back the first byte of the response
-	afeLatencies       metric.Float64Histogram // Latency between Spanner API Frontend receiving an RPC and starting to write back the response.
-	gfeErrorCount      metric.Int64Counter     // Counter for the number of requests that failed to reach the Google network.
-	afeErrorCount      metric.Int64Counter     // Counter for the number of requests that failed to reach the Spanner API Frontend.
-	operationCount     metric.Int64Counter     // Counter for the number of operations.
-	attemptCount       metric.Int64Counter     // Counter for the number of attempts.
 }
 
-func newBuiltinMetricsTracerFactory(ctx context.Context, dbpath, compression string, isAFEBuiltInMetricEnabled, isEnableGRPCBuiltInMetrics bool, metricsProvider metric.MeterProvider, opts ...option.ClientOption) (*builtinMetricsTracerFactory, error) {
-	clientUID, err := generateClientUID()
-	if err != nil {
-		log.Printf("built-in metrics: generateClientUID failed: %v. Using empty string in the %v metric atteribute", err, metricLabelKeyClientUID)
-	}
-	project, instance, database, err := parseDatabaseName(dbpath)
-	if err != nil {
-		return nil, err
-	}
-
-	tracerFactory := &builtinMetricsTracerFactory{
-		enabled: false,
-		clientAttributes: []attribute.KeyValue{
-			attribute.String(monitoredResLabelKeyProject, project),
-			attribute.String(monitoredResLabelKeyInstance, instance),
-			attribute.String(metricLabelKeyDatabase, database),
-			attribute.String(metricLabelKeyClientUID, clientUID),
-			attribute.String(metricLabelKeyClientName, clientName),
-			attribute.String(monitoredResLabelKeyClientHash, generateClientHash(clientUID)),
-			// Skipping instance config until we have a way to get it
-			attribute.String(monitoredResLabelKeyInstanceConfig, "unknown"),
-			attribute.String(monitoredResLabelKeyLocation, detectClientLocation(ctx)),
-		},
-		shutdown: func(ctx context.Context) {},
-	}
-	tracerFactory.isAFEBuiltInMetricEnabled = isAFEBuiltInMetricEnabled
-	tracerFactory.isDirectPathEnabled = false
-	tracerFactory.enabled = false
-	var meterProvider *sdkmetric.MeterProvider
-	if metricsProvider == nil {
-		// Create default meter provider
-		mpOptions, err := builtInMeterProviderOptions(project, compression, tracerFactory.clientAttributes, opts...)
-		if err != nil {
-			return tracerFactory, err
-		}
-		meterProvider = sdkmetric.NewMeterProvider(mpOptions...)
-
-		if isEnableGRPCBuiltInMetrics {
-			mo := opentelemetry.MetricsOptions{
-				MeterProvider: meterProvider,
-				Metrics:       stats.NewMetrics(grpcMetricsToEnable...),
-			}
-
-			// Configure gRPC dial options to enable gRPC metrics collection and static method call option.
-			// The static method call option ensures consistent method names in metrics by preventing gRPC from
-			// automatically adding service prefixes to method names. This helps maintain consistent metric
-			// naming across different gRPC calls.
-			tracerFactory.clientOpts = []option.ClientOption{
-				option.WithGRPCDialOption(
-					opentelemetry.DialOption(opentelemetry.Options{MetricsOptions: mo})),
-				option.WithGRPCDialOption(
-					grpc.WithDefaultCallOptions(grpc.StaticMethodCallOption{})),
-			}
-		}
-		tracerFactory.enabled = true
-		tracerFactory.shutdown = func(ctx context.Context) {
-			meterProvider.ForceFlush(ctx)
-			meterProvider.Shutdown(ctx)
-		}
-	} else {
-		switch metricsProvider.(type) {
-		case noop.MeterProvider:
-			return tracerFactory, nil
-		default:
-			return tracerFactory, errors.New("unknown MetricsProvider type")
-		}
-	}
-
-	// Create meter and instruments
-	meter := meterProvider.Meter(builtInMetricsMeterName, metric.WithInstrumentationVersion(internal.Version))
-	err = tracerFactory.createInstruments(meter)
-	return tracerFactory, err
-}
-
-func builtInMeterProviderOptions(project, compression string, clientAttributes []attribute.KeyValue, opts ...option.ClientOption) ([]sdkmetric.Option, error) {
-	allOpts := createExporterOptions(opts...)
-	defaultExporter, err := newMonitoringExporter(context.Background(), project, compression, clientAttributes, allOpts...)
-	if err != nil {
-		return nil, err
-	}
-	var views []sdkmetric.View
-	for _, m := range grpcMetricsToEnable {
-		views = append(views, sdkmetric.NewView(
-			sdkmetric.Instrument{
-				Name: m,
-			},
-			sdkmetric.Stream{
-				Aggregation: sdkmetric.AggregationSum{},
-				AttributeFilter: func(kv attribute.KeyValue) bool {
-					if _, ok := allowedMetricLabels[string(kv.Key)]; ok {
-						return true
-					}
-					return false
-				},
-			},
-		))
-	}
-	return []sdkmetric.Option{sdkmetric.WithReader(
-		sdkmetric.NewPeriodicReader(
-			defaultExporter,
-			sdkmetric.WithInterval(defaultSamplePeriod),
-		),
-	), sdkmetric.WithView(views...)}, nil
-}
-
-func (tf *builtinMetricsTracerFactory) createInstruments(meter metric.Meter) error {
-	var err error
-
-	// Create operation_latencies
-	tf.operationLatencies, err = meter.Float64Histogram(
-		nativeMetricsPrefix+metricNameOperationLatencies,
-		metric.WithDescription("Total time until final operation success or failure, including retries and backoff."),
-		metric.WithUnit(metricUnitMS),
-		metric.WithExplicitBucketBoundaries(bucketBounds...),
-	)
-	if err != nil {
-		return err
-	}
-
-	// Create attempt_latencies
-	tf.attemptLatencies, err = meter.Float64Histogram(
-		nativeMetricsPrefix+metricNameAttemptLatencies,
-		metric.WithDescription("Client observed latency per RPC attempt."),
-		metric.WithUnit(metricUnitMS),
-		metric.WithExplicitBucketBoundaries(bucketBounds...),
-	)
-	if err != nil {
-		return err
-	}
-
-	tf.gfeLatencies, err = meter.Float64Histogram(
-		nativeMetricsPrefix+metricNameGFELatencies,
-		metric.WithDescription("Latency between Google's network receiving an RPC and reading back the first byte of the response."),
-		metric.WithUnit(metricUnitMS),
-		metric.WithExplicitBucketBoundaries(bucketBounds...),
-	)
-	if err != nil {
-		return err
-	}
-
-	tf.afeLatencies, err = meter.Float64Histogram(
-		nativeMetricsPrefix+metricNameAFELatencies,
-		metric.WithDescription("Latency between Spanner API Frontend receiving an RPC and starting to write back the response."),
-		metric.WithUnit(metricUnitMS),
-		metric.WithExplicitBucketBoundaries(bucketBounds...),
-	)
-	if err != nil {
-		return err
-	}
-
-	// Create operation_count
-	tf.operationCount, err = meter.Int64Counter(
-		nativeMetricsPrefix+metricNameOperationCount,
-		metric.WithDescription("The count of database operations."),
-		metric.WithUnit(metricUnitCount),
-	)
-	if err != nil {
-		return err
-	}
-
-	// Create attempt_count
-	tf.attemptCount, err = meter.Int64Counter(
-		nativeMetricsPrefix+metricNameAttemptCount,
-		metric.WithDescription("The number of attempts made for the operation, including the initial attempt."),
-		metric.WithUnit(metricUnitCount),
-	)
-
-	tf.gfeErrorCount, err = meter.Int64Counter(
-		nativeMetricsPrefix+metricNameGFEConnectivityErrorCount,
-		metric.WithDescription("Number of requests that failed to reach the Google network."),
-		metric.WithUnit(metricUnitCount),
-	)
-
-	tf.afeErrorCount, err = meter.Int64Counter(
-		nativeMetricsPrefix+metricNameAFEConnectivityErrorCount,
-		metric.WithDescription("Number of requests that failed to reach the Spanner API Frontend."),
-		metric.WithUnit(metricUnitCount),
-	)
-	return err
-}
-
 // builtinMetricsTracer is created one per operation.
 // It is used to store metric instruments, attribute values, and other data required to obtain and record them.
 type builtinMetricsTracer struct {
@@ -440,19 +226,6 @@ type builtinMetricsTracer struct {
 	builtInEnabled            bool            // Indicates if built-in metrics are enabled.
 	isAFEBuiltInMetricEnabled bool
 
-	// clientAttributes are attributes specific to a client instance that do not change across different operations on the client.
-	clientAttributes []attribute.KeyValue
-
-	// Metrics instruments
-	instrumentOperationLatencies metric.Float64Histogram // Histogram for operation latencies.
-	instrumentAttemptLatencies   metric.Float64Histogram // Histogram for attempt latencies.
-	instrumentGFELatencies       metric.Float64Histogram // Histogram for GFE latencies.
-	instrumentAFELatencies       metric.Float64Histogram // Histogram for AFE latencies.
-	instrumentGFEErrorCount      metric.Int64Counter     // Counter for GFE connectivity errors.
-	instrumentAFEErrorCount      metric.Int64Counter     // Counter for AFE connectivity errors.
-	instrumentOperationCount     metric.Int64Counter     // Counter for the number of operations.
-	instrumentAttemptCount       metric.Int64Counter     // Counter for the number of attempts.
-
 	method string // The method being traced.
 
 	currOp *opTracer // The current operation tracer.
@@ -532,85 +305,20 @@ func (tf *builtinMetricsTracerFactory) createBuiltinMe
 		ctx:                       ctx,
 		builtInEnabled:            tf.enabled,
 		currOp:                    &currOpTracer,
-		clientAttributes:          tf.clientAttributes,
 		isAFEBuiltInMetricEnabled: tf.isAFEBuiltInMetricEnabled,
-
-		instrumentOperationLatencies: tf.operationLatencies,
-		instrumentAttemptLatencies:   tf.attemptLatencies,
-		instrumentOperationCount:     tf.operationCount,
-		instrumentAttemptCount:       tf.attemptCount,
-		instrumentGFELatencies:       tf.gfeLatencies,
-		instrumentAFELatencies:       tf.afeLatencies,
-		instrumentGFEErrorCount:      tf.gfeErrorCount,
-		instrumentAFEErrorCount:      tf.afeErrorCount,
 	}
 }
 
-// toOtelMetricAttrs:
-// - converts metric attributes values captured throughout the operation / attempt
-// to OpenTelemetry attributes format,
-// - combines these with common client attributes and returns
-func (mt *builtinMetricsTracer) toOtelMetricAttrs(metricName string) ([]attribute.KeyValue, error) {
-	if mt.currOp == nil || mt.currOp.currAttempt == nil {
-		return nil, fmt.Errorf("unable to create attributes list for unknown metric: %v", metricName)
-	}
-	// Get metric details
-	mDetails, found := metricsDetails[metricName]
-	if !found {
-		return nil, fmt.Errorf("unable to create attributes list for unknown metric: %v", metricName)
-	}
-
-	rpcStatus := mt.currOp.status
-	if mDetails.recordedPerAttempt {
-		rpcStatus = mt.currOp.currAttempt.status
-	}
-
-	return []attribute.KeyValue{
-		attribute.String(metricLabelKeyMethod, strings.ReplaceAll(strings.TrimPrefix(mt.method, "/google.spanner.v1."), "/", ".")),
-		attribute.String(metricLabelKeyDirectPathEnabled, strconv.FormatBool(mt.currOp.directPathEnabled)),
-		attribute.String(metricLabelKeyDirectPathUsed, strconv.FormatBool(mt.currOp.currAttempt.directPathUsed)),
-		attribute.String(metricLabelKeyStatus, rpcStatus),
-	}, nil
-}
-
 func (t *builtinMetricsTracer) recordGFELatency(latency time.Duration) {
-	if t.builtInEnabled {
-		attrs, err := t.toOtelMetricAttrs(metricNameGFELatencies)
-		if err != nil {
-			return
-		}
-		t.instrumentGFELatencies.Record(t.ctx, float64(latency.Milliseconds()), metric.WithAttributes(attrs...))
-	}
 }
 
 func (t *builtinMetricsTracer) recordAFELatency(latency time.Duration) {
-	if !t.isAFEBuiltInMetricEnabled {
-		return
-	}
-	attrs, err := t.toOtelMetricAttrs(metricNameAFELatencies)
-	if err != nil {
-		return
-	}
-	t.instrumentAFELatencies.Record(t.ctx, float64(latency.Milliseconds()), metric.WithAttributes(attrs...))
 }
 
 func (t *builtinMetricsTracer) recordGFEError() {
-	attrs, err := t.toOtelMetricAttrs(metricNameGFEConnectivityErrorCount)
-	if err != nil {
-		return
-	}
-	t.instrumentGFEErrorCount.Add(t.ctx, 1, metric.WithAttributes(attrs...))
 }
 
 func (t *builtinMetricsTracer) recordAFEError() {
-	if !t.isAFEBuiltInMetricEnabled {
-		return
-	}
-	attrs, err := t.toOtelMetricAttrs(metricNameAFEConnectivityErrorCount)
-	if err != nil {
-		return
-	}
-	t.instrumentAFEErrorCount.Add(t.ctx, 1, metric.WithAttributes(attrs...))
 }
 
 // Convert error to grpc status error
@@ -635,66 +343,12 @@ func recordAttemptCompletion(mt *builtinMetricsTracer)
 // Ignore errors seen while creating metric attributes since metric can still
 // be recorded with rest of the attributes
 func recordAttemptCompletion(mt *builtinMetricsTracer) {
-	if !mt.builtInEnabled {
-		return
-	}
-	// capture AFE metrics only if direct-path is enabled and used in current attempt
-	if mt.currOp.currAttempt.directPathUsed {
-		if dur, ok := mt.currOp.currAttempt.serverTimingMetrics[afeTimingHeader]; ok {
-			mt.recordAFELatency(dur)
-		} else {
-			mt.recordAFEError()
-		}
-	} else {
-		if dur, ok := mt.currOp.currAttempt.serverTimingMetrics[gfeTimingHeader]; ok {
-			mt.recordGFELatency(dur)
-		} else {
-			mt.recordGFEError()
-		}
-	}
-
-	// Calculate elapsed time
-	elapsedTime := convertToMs(time.Since(mt.currOp.currAttempt.startTime))
-
-	// Record attempt_latencies
-	attemptLatAttrs, err := mt.toOtelMetricAttrs(metricNameAttemptLatencies)
-	if err != nil {
-		return
-	}
-	mt.instrumentAttemptLatencies.Record(mt.ctx, elapsedTime, metric.WithAttributes(attemptLatAttrs...))
 }
 
 // recordOperationCompletion records as many operation specific metrics as it can
 // Ignores error seen while creating metric attributes since metric can still
 // be recorded with rest of the attributes
 func recordOperationCompletion(mt *builtinMetricsTracer) {
-	if !mt.builtInEnabled {
-		return
-	}
-
-	// Calculate elapsed time
-	elapsedTimeMs := convertToMs(time.Since(mt.currOp.startTime))
-
-	// Record operation_count
-	opCntAttrs, err := mt.toOtelMetricAttrs(metricNameOperationCount)
-	if err != nil {
-		return
-	}
-	mt.instrumentOperationCount.Add(mt.ctx, 1, metric.WithAttributes(opCntAttrs...))
-
-	// Record operation_latencies
-	opLatAttrs, err := mt.toOtelMetricAttrs(metricNameOperationLatencies)
-	if err != nil {
-		return
-	}
-	mt.instrumentOperationLatencies.Record(mt.ctx, elapsedTimeMs, metric.WithAttributes(opLatAttrs...))
-
-	// Record attempt_count
-	attemptCntAttrs, err := mt.toOtelMetricAttrs(metricNameAttemptCount)
-	if err != nil {
-		return
-	}
-	mt.instrumentAttemptCount.Add(mt.ctx, mt.currOp.attemptCount, metric.WithAttributes(attemptCntAttrs...))
 }
 
 func convertToMs(d time.Duration) float64 {
