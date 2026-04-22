--- vendor/cloud.google.com/go/storage/option.go.orig	2026-03-13 06:41:11 UTC
+++ vendor/cloud.google.com/go/storage/option.go
@@ -19,9 +19,7 @@ import (
 	"strconv"
 	"time"
 
-	"cloud.google.com/go/storage/experimental"
 	storageinternal "cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 )
@@ -37,10 +35,7 @@ func init() {
 
 func init() {
 	// initialize experimental options
-	storageinternal.WithMetricExporter = withMetricExporter
 	storageinternal.WithMetricInterval = withMetricInterval
-	storageinternal.WithMeterProvider = withMeterProvider
-	storageinternal.WithReadStallTimeout = withReadStallTimeout
 	storageinternal.WithGRPCBidiReads = withGRPCBidiReads
 	storageinternal.WithZonalBucketAPIs = withZonalBucketAPIs
 	storageinternal.WithDirectConnectivityEnforced = withDirectConnectivityEnforced
@@ -81,11 +76,7 @@ type storageConfig struct {
 	useJSONforReads        bool
 	readAPIWasSet          bool
 	disableClientMetrics   bool
-	metricExporter         *metric.Exporter
 	metricInterval         time.Duration
-	meterProvider          *metric.MeterProvider
-	manualReader           *metric.ManualReader
-	readStallTimeoutConfig *experimental.ReadStallTimeoutConfig
 	grpcBidiReads          bool
 	grpcAppendableUploads  bool
 	grpcDirectPathEnforced bool
@@ -202,43 +193,25 @@ type withMetricExporterConfig struct {
 type withMetricExporterConfig struct {
 	internaloption.EmbeddableAdapter
 	// exporter override
-	metricExporter *metric.Exporter
 }
 
-func withMetricExporter(ex *metric.Exporter) option.ClientOption {
-	return &withMetricExporterConfig{metricExporter: ex}
-}
-
 func (w *withMetricExporterConfig) ApplyStorageOpt(c *storageConfig) {
-	c.metricExporter = w.metricExporter
 }
 
 type withTestMetricReaderConfig struct {
 	internaloption.EmbeddableAdapter
 	// reader override
-	metricReader *metric.ManualReader
 }
 
 type withMeterProviderConfig struct {
 	internaloption.EmbeddableAdapter
 	// meter provider override
-	meterProvider *metric.MeterProvider
 }
 
-func withMeterProvider(provider *metric.MeterProvider) option.ClientOption {
-	return &withMeterProviderConfig{meterProvider: provider}
-}
-
 func (w *withMeterProviderConfig) ApplyStorageOpt(c *storageConfig) {
-	c.meterProvider = w.meterProvider
 }
 
-func withTestMetricReader(ex *metric.ManualReader) option.ClientOption {
-	return &withTestMetricReaderConfig{metricReader: ex}
-}
-
 func (w *withTestMetricReaderConfig) ApplyStorageOpt(c *storageConfig) {
-	c.manualReader = w.metricReader
 }
 
 // WithReadStallTimeout is an option that may be passed to [NewClient].
@@ -248,31 +221,12 @@ func (w *withTestMetricReaderConfig) ApplyStorageOpt(c
 //
 // This is only supported for the read operation and that too for http(XML) client.
 // Grpc read-operation will be supported soon.
-func withReadStallTimeout(rstc *experimental.ReadStallTimeoutConfig) option.ClientOption {
-	// TODO (raj-prince): To keep separate dynamicDelay instance for different BucketHandle.
-	// Currently, dynamicTimeout is kept at the client and hence shared across all the
-	// BucketHandle, which is not the ideal state. As latency depends on location of VM
-	// and Bucket, and read latency of different buckets may lie in different range.
-	// Hence having a separate dynamicTimeout instance at BucketHandle level will
-	// be better
-	if rstc.Min == time.Duration(0) {
-		rstc.Min = defaultDynamicReadReqMinTimeout
-	}
-	if rstc.TargetPercentile == 0 {
-		rstc.TargetPercentile = defaultTargetPercentile
-	}
-	return &withReadStallTimeoutConfig{
-		readStallTimeoutConfig: rstc,
-	}
-}
 
 type withReadStallTimeoutConfig struct {
 	internaloption.EmbeddableAdapter
-	readStallTimeoutConfig *experimental.ReadStallTimeoutConfig
 }
 
 func (wrstc *withReadStallTimeoutConfig) ApplyStorageOpt(config *storageConfig) {
-	config.readStallTimeoutConfig = wrstc.readStallTimeoutConfig
 }
 
 func withGRPCBidiReads() option.ClientOption {
