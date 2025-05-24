--- vendor/cloud.google.com/go/storage/option.go.orig	2025-03-12 21:31:11 UTC
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
@@ -37,9 +35,7 @@ func init() {
 
 func init() {
 	// initialize experimental options
-	storageinternal.WithMetricExporter = withMetricExporter
 	storageinternal.WithMetricInterval = withMetricInterval
-	storageinternal.WithReadStallTimeout = withReadStallTimeout
 	storageinternal.WithGRPCBidiReads = withGRPCBidiReads
 }
 
@@ -78,10 +74,7 @@ type storageConfig struct {
 	useJSONforReads        bool
 	readAPIWasSet          bool
 	disableClientMetrics   bool
-	metricExporter         *metric.Exporter
 	metricInterval         time.Duration
-	manualReader           *metric.ManualReader
-	readStallTimeoutConfig *experimental.ReadStallTimeoutConfig
 	grpcBidiReads          bool
 }
 
@@ -184,29 +177,17 @@ type withMetricExporterConfig struct {
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
 
-func withTestMetricReader(ex *metric.ManualReader) option.ClientOption {
-	return &withTestMetricReaderConfig{metricReader: ex}
-}
-
 func (w *withTestMetricReaderConfig) ApplyStorageOpt(c *storageConfig) {
-	c.manualReader = w.metricReader
 }
 
 // WithReadStallTimeout is an option that may be passed to [NewClient].
@@ -216,31 +197,18 @@ func (w *withTestMetricReaderConfig) ApplyStorageOpt(c
 //
 // This is only supported for the read operation and that too for http(XML) client.
 // Grpc read-operation will be supported soon.
-func withReadStallTimeout(rstc *experimental.ReadStallTimeoutConfig) option.ClientOption {
 	// TODO (raj-prince): To keep separate dynamicDelay instance for different BucketHandle.
 	// Currently, dynamicTimeout is kept at the client and hence shared across all the
 	// BucketHandle, which is not the ideal state. As latency depends on location of VM
 	// and Bucket, and read latency of different buckets may lie in different range.
 	// Hence having a separate dynamicTimeout instance at BucketHandle level will
 	// be better
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
