diff --git a/vendor/cloud.google.com/go/storage/option.go b/vendor/cloud.google.com/go/storage/option.go
index 758f3ee26..3d7dda429 100644
--- vendor/cloud.google.com/go/storage/option.go.orig
+++ vendor/cloud.google.com/go/storage/option.go
@@ -21,7 +21,6 @@ import (
 
 	"cloud.google.com/go/storage/experimental"
 	storageinternal "cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 )
@@ -85,10 +84,10 @@ type storageConfig struct {
 	disableClientMetrics   bool
 	enableOtelMetrics      bool
 	enableOtelDebugMetrics bool
-	metricExporter         *metric.Exporter
+	metricExporter         any
 	metricInterval         time.Duration
-	meterProvider          *metric.MeterProvider
-	manualReader           *metric.ManualReader
+	meterProvider          any
+	manualReader           any
 	readStallTimeoutConfig *experimental.ReadStallTimeoutConfig
 	grpcBidiReads          bool
 	grpcAppendableUploads  bool
@@ -206,10 +205,10 @@ func (w *withMeterOptions) ApplyStorageOpt(c *storageConfig) {
 type withMetricExporterConfig struct {
 	internaloption.EmbeddableAdapter
 	// exporter override
-	metricExporter *metric.Exporter
+	metricExporter any
 }
 
-func withMetricExporter(ex *metric.Exporter) option.ClientOption {
+func withMetricExporter(ex any) option.ClientOption {
 	return &withMetricExporterConfig{metricExporter: ex}
 }
 
@@ -220,16 +219,16 @@ func (w *withMetricExporterConfig) ApplyStorageOpt(c *storageConfig) {
 type withTestMetricReaderConfig struct {
 	internaloption.EmbeddableAdapter
 	// reader override
-	metricReader *metric.ManualReader
+	metricReader any
 }
 
 type withMeterProviderConfig struct {
 	internaloption.EmbeddableAdapter
 	// meter provider override
-	meterProvider *metric.MeterProvider
+	meterProvider any
 }
 
-func withMeterProvider(provider *metric.MeterProvider) option.ClientOption {
+func withMeterProvider(provider any) option.ClientOption {
 	return &withMeterProviderConfig{meterProvider: provider}
 }
 
@@ -237,7 +236,7 @@ func (w *withMeterProviderConfig) ApplyStorageOpt(c *storageConfig) {
 	c.meterProvider = w.meterProvider
 }
 
-func withTestMetricReader(ex *metric.ManualReader) option.ClientOption {
+func withTestMetricReader(ex any) option.ClientOption {
 	return &withTestMetricReaderConfig{metricReader: ex}
 }
 
