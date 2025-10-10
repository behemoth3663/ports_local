--- vendor/cloud.google.com/go/storage/option.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/option.go
@@ -21,7 +21,6 @@ import (
 
 	"cloud.google.com/go/storage/experimental"
 	storageinternal "cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 )
@@ -37,7 +36,6 @@ func init() {
 
 func init() {
 	// initialize experimental options
-	storageinternal.WithMetricExporter = withMetricExporter
 	storageinternal.WithMetricInterval = withMetricInterval
 	storageinternal.WithReadStallTimeout = withReadStallTimeout
 	storageinternal.WithGRPCBidiReads = withGRPCBidiReads
@@ -79,9 +77,7 @@ type storageConfig struct {
 	useJSONforReads        bool
 	readAPIWasSet          bool
 	disableClientMetrics   bool
-	metricExporter         *metric.Exporter
 	metricInterval         time.Duration
-	manualReader           *metric.ManualReader
 	readStallTimeoutConfig *experimental.ReadStallTimeoutConfig
 	grpcBidiReads          bool
 	grpcAppendableUploads  bool
@@ -185,30 +181,16 @@ type withMetricExporterConfig struct {
 
 type withMetricExporterConfig struct {
 	internaloption.EmbeddableAdapter
-	// exporter override
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
-	// reader override
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
