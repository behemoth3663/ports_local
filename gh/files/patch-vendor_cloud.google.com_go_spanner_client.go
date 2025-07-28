--- vendor/cloud.google.com/go/spanner/client.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/spanner/client.go
@@ -27,14 +27,10 @@ import (
 	"strings"
 	"time"
 
-	"cloud.google.com/go/internal/trace"
 	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
 	"github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp"
 	grpcgcppb "github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp/grpc_gcp"
 	"github.com/googleapis/gax-go/v2"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
-	"go.opentelemetry.io/otel/metric/noop"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -350,8 +346,6 @@ type ClientConfig struct {
 	// should be used for non-transactional reads or queries.
 	DirectedReadOptions *sppb.DirectedReadOptions
 
-	OpenTelemetryMeterProvider metric.MeterProvider
-
 	// EnableEndToEndTracing indicates whether end to end tracing is enabled or not. If
 	// it is enabled, trace spans will be created at Spanner layer. Enabling end to end
 	// tracing requires OpenTelemetry to be set up. Simply enabling this option won't
@@ -369,20 +363,6 @@ type openTelemetryConfig struct {
 
 type openTelemetryConfig struct {
 	enabled                        bool
-	meterProvider                  metric.MeterProvider
-	attributeMap                   []attribute.KeyValue
-	attributeMapWithMultiplexed    []attribute.KeyValue
-	attributeMapWithoutMultiplexed []attribute.KeyValue
-	otMetricRegistration           metric.Registration
-	openSessionCount               metric.Int64ObservableGauge
-	maxAllowedSessionsCount        metric.Int64ObservableGauge
-	sessionsCount                  metric.Int64ObservableGauge
-	maxInUseSessionsCount          metric.Int64ObservableGauge
-	getSessionTimeoutsCount        metric.Int64Counter
-	acquiredSessionsCount          metric.Int64Counter
-	releasedSessionsCount          metric.Int64Counter
-	gfeLatency                     metric.Int64Histogram
-	gfeHeaderMissingCount          metric.Int64Counter
 }
 
 func contextWithOutgoingMetadata(ctx context.Context, md metadata.MD, disableRouteToLeader bool) context.Context {
@@ -415,9 +395,6 @@ func newClientWithConfig(ctx context.Context, database
 		return nil, err
 	}
 
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.NewClient")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	// Explicitly disable some gRPC experiments as they are not stable yet.
 	gRPCPickFirstEnvVarName := "GRPC_EXPERIMENTAL_ENABLE_NEW_PICK_FIRST"
 	if os.Getenv(gRPCPickFirstEnvVarName) == "" {
@@ -445,43 +422,21 @@ func newClientWithConfig(ctx context.Context, database
 		config.NumChannels = numChannels
 	}
 
-	var metricsProvider metric.MeterProvider
-	if emulatorAddr := os.Getenv("SPANNER_EMULATOR_HOST"); emulatorAddr != "" {
-		// Do not emit native metrics when emulator is being used
-		metricsProvider = noop.NewMeterProvider()
-	}
 	// Check if native metrics are disabled via env.
 	if disableNativeMetrics, _ := strconv.ParseBool(os.Getenv("SPANNER_DISABLE_BUILTIN_METRICS")); disableNativeMetrics {
 		config.DisableNativeMetrics = true
 	}
-	if config.DisableNativeMetrics {
-		// Do not emit native metrics when DisableNativeMetrics is set
-		metricsProvider = noop.NewMeterProvider()
-	}
 	isAFEBuiltInMetricEnabled := strings.EqualFold("false", os.Getenv("SPANNER_DISABLE_AFE_SERVER_TIMING"))
-	isGRPCBuiltInMetricsEnabled := strings.EqualFold("false", os.Getenv("SPANNER_DISABLE_DIRECT_ACCESS_GRPC_BUILTIN_METRICS"))
 	// enable the AFE/GRPC built-in metrics if direct-path is enabled
 	isDirectPathEnabled, _ := strconv.ParseBool(os.Getenv("GOOGLE_SPANNER_ENABLE_DIRECT_ACCESS"))
 	if isDirectPathEnabled {
 		isAFEBuiltInMetricEnabled = true
-		isGRPCBuiltInMetricsEnabled = true
 	}
 	// disable the AFE/GRPC built-in metrics if the env var is explicitly set
 	if ok, _ := strconv.ParseBool(os.Getenv("SPANNER_DISABLE_AFE_SERVER_TIMING")); ok {
 		isAFEBuiltInMetricEnabled = false
 	}
-	if ok, _ := strconv.ParseBool(os.Getenv("SPANNER_DISABLE_DIRECT_ACCESS_GRPC_BUILTIN_METRICS")); ok {
-		isGRPCBuiltInMetricsEnabled = false
-	}
 
-	metricsTracerFactory, err := newBuiltinMetricsTracerFactory(ctx, database, config.Compression, isAFEBuiltInMetricEnabled, isGRPCBuiltInMetricsEnabled, metricsProvider, opts...)
-	if err != nil {
-		return nil, err
-	}
-	if len(metricsTracerFactory.clientOpts) > 0 {
-		opts = append(opts, metricsTracerFactory.clientOpts...)
-	}
-
 	var pool gtransport.ConnPool
 
 	if gme != nil {
@@ -572,18 +527,6 @@ func newClientWithConfig(ctx context.Context, database
 	// Create a session client.
 	sc := newSessionClient(pool, database, config.UserAgent, sessionLabels, config.DatabaseRole, config.DisableRouteToLeader, md, config.BatchTimeout, config.Logger, config.CallOptions)
 
-	// Create a OpenTelemetry configuration
-	otConfig, err := createOpenTelemetryConfig(config.OpenTelemetryMeterProvider, config.Logger, sc.id, database)
-	if err != nil {
-		// The error returned here will be due to database name parsing
-		return nil, err
-	}
-	// To prevent data race in unit tests (ex: TestClient_SessionNotFound)
-	sc.mu.Lock()
-	sc.otConfig = otConfig
-	sc.metricsTracerFactory = metricsTracerFactory
-	sc.mu.Unlock()
-
 	// Create a session pool.
 	config.SessionPoolConfig.sessionLabels = sessionLabels
 	sp, err := newSessionPool(sc, config.SessionPoolConfig)
@@ -604,8 +547,6 @@ func newClientWithConfig(ctx context.Context, database
 		ct:                   getCommonTags(sc),
 		disableRouteToLeader: config.DisableRouteToLeader,
 		dro:                  config.DirectedReadOptions,
-		otConfig:             otConfig,
-		metricsTracerFactory: metricsTracerFactory,
 	}
 	return c, nil
 }
@@ -1043,8 +984,6 @@ func (c *Client) ReadWriteTransaction(ctx context.Cont
 // See https://godoc.org/cloud.google.com/go/spanner#ReadWriteTransaction for
 // more details.
 func (c *Client) ReadWriteTransaction(ctx context.Context, f func(context.Context, *ReadWriteTransaction) error) (commitTimestamp time.Time, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.ReadWriteTransaction")
-	defer func() { trace.EndSpan(ctx, err) }()
 	resp, err := c.rwTransaction(ctx, f, TransactionOptions{})
 	return resp.CommitTs, err
 }
@@ -1057,8 +996,6 @@ func (c *Client) ReadWriteTransactionWithOptions(ctx c
 // See https://godoc.org/cloud.google.com/go/spanner#ReadWriteTransaction for
 // more details.
 func (c *Client) ReadWriteTransactionWithOptions(ctx context.Context, f func(context.Context, *ReadWriteTransaction) error, options TransactionOptions) (resp CommitResponse, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.ReadWriteTransactionWithOptions")
-	defer func() { trace.EndSpan(ctx, err) }()
 	resp, err = c.rwTransaction(ctx, f, options)
 	return resp, err
 }
@@ -1102,7 +1039,6 @@ func (c *Client) rwTransaction(ctx context.Context, f 
 			// BeginTransaction RPC invocation will be retried on a new session if it returns SessionNotFound.
 			t.txReadOnly.sh = sh
 			if err = t.begin(ctx, nil); err != nil {
-				trace.TracePrintf(ctx, nil, "Error while BeginTransaction during retrying a ReadWrite transaction: %v", ToSpannerError(err))
 				return ToSpannerError(err)
 			}
 		} else {
@@ -1127,9 +1063,6 @@ func (c *Client) rwTransaction(ctx context.Context, f 
 		t.ct = c.ct
 		t.otConfig = c.otConfig
 
-		trace.TracePrintf(ctx, map[string]interface{}{"transactionSelector": t.getTransactionSelector().String()},
-			"Starting transaction attempt")
-
 		resp, err = t.runInTransaction(ctx, f)
 		return err
 	})
@@ -1227,9 +1160,6 @@ func (c *Client) Apply(ctx context.Context, ms []*Muta
 		opt(ao)
 	}
 
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.Apply")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if !ao.atLeastOnce {
 		resp, err := c.ReadWriteTransactionWithOptions(ctx, func(ctx context.Context, t *ReadWriteTransaction) error {
 			return t.BufferWrite(ms)
@@ -1340,7 +1270,6 @@ func (r *BatchWriteResponseIterator) Stop() {
 		if err == iterator.Done {
 			err = nil
 		}
-		defer trace.EndSpan(r.ctx, err)
 	}
 	if r.cancel != nil {
 		r.cancel()
@@ -1400,12 +1329,7 @@ func (c *Client) BatchWriteWithOptions(ctx context.Con
 
 // BatchWriteWithOptions is same as BatchWrite. It accepts additional options to customize the request.
 func (c *Client) BatchWriteWithOptions(ctx context.Context, mgs []*MutationGroup, opts BatchWriteOptions) *BatchWriteResponseIterator {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.BatchWrite")
-
 	var err error
-	defer func() {
-		trace.EndSpan(ctx, err)
-	}()
 
 	opts = c.bwo.merge(opts)
 
@@ -1430,14 +1354,6 @@ func (c *Client) BatchWriteWithOptions(ctx context.Con
 			ExcludeTxnFromChangeStreams: opts.ExcludeTxnFromChangeStreams,
 		}, gax.WithGRPCOptions(grpc.Header(&md)))
 
-		if getGFELatencyMetricsFlag() && md != nil && c.ct != nil {
-			if metricErr := createContextAndCaptureGFELatencyMetrics(ct, c.ct, md, "BatchWrite"); metricErr != nil {
-				trace.TracePrintf(ct, nil, "Error in recording GFE Latency. Try disabling and rerunning. Error: %v", err)
-			}
-		}
-		if metricErr := recordGFELatencyMetricsOT(ct, md, "BatchWrite", c.otConfig); metricErr != nil {
-			trace.TracePrintf(ct, nil, "Error in recording GFE Latency through OpenTelemetry. Error: %v", err)
-		}
 		return stream, rpcErr
 	}
 
@@ -1461,7 +1377,6 @@ func (c *Client) BatchWriteWithOptions(ctx context.Con
 	}
 
 	ctx, cancel := context.WithCancel(ctx)
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/spanner.BatchWriteResponseIterator")
 	return &BatchWriteResponseIterator{
 		ctx:                ctx,
 		meterTracerFactory: c.metricsTracerFactory,
