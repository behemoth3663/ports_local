--- vendor/cloud.google.com/go/storage/option.go.orig	2026-11-02 21:24:35 UTC
+++ vendor/cloud.google.com/go/storage/option.go
@@ -21,7 +21,6 @@ import (
 
 	"cloud.google.com/go/storage/experimental"
 	storageinternal "cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 )
@@ -84,10 +83,10 @@ type storageConfig struct {
 	disableClientMetrics   bool
 	enableOtelMetrics      bool
 	enableOtelDebugMetrics bool
-	metricExporter         *metric.Exporter
+	metricExporter         interface{}
 	metricInterval         time.Duration
-	meterProvider          *metric.MeterProvider
-	manualReader           *metric.ManualReader
+	meterProvider          interface{}
+	manualReader           interface{}
 	readStallTimeoutConfig *experimental.ReadStallTimeoutConfig
 	grpcBidiReads          bool
 	grpcAppendableUploads  bool
@@ -206,10 +205,10 @@ type withMetricExporterConfig struct {
 type withMetricExporterConfig struct {
 	internaloption.EmbeddableAdapter
 	// exporter override
-	metricExporter *metric.Exporter
+	metricExporter interface{}
 }
 
-func withMetricExporter(ex *metric.Exporter) option.ClientOption {
+func withMetricExporter(ex interface{}) option.ClientOption {
 	return &withMetricExporterConfig{metricExporter: ex}
 }
 
@@ -220,16 +219,16 @@ type withTestMetricReaderConfig struct {
 type withTestMetricReaderConfig struct {
 	internaloption.EmbeddableAdapter
 	// reader override
-	metricReader *metric.ManualReader
+	metricReader interface{}
 }
 
 type withMeterProviderConfig struct {
 	internaloption.EmbeddableAdapter
 	// meter provider override
-	meterProvider *metric.MeterProvider
+	meterProvider interface{}
 }
 
-func withMeterProvider(provider *metric.MeterProvider) option.ClientOption {
+func withMeterProvider(provider interface{}) option.ClientOption {
 	return &withMeterProviderConfig{meterProvider: provider}
 }
 
@@ -237,7 +236,7 @@ func (w *withMeterProviderConfig) ApplyStorageOpt(c *s
 	c.meterProvider = w.meterProvider
 }
 
-func withTestMetricReader(ex *metric.ManualReader) option.ClientOption {
+func withTestMetricReader(ex interface{}) option.ClientOption {
 	return &withTestMetricReaderConfig{metricReader: ex}
 }
 
