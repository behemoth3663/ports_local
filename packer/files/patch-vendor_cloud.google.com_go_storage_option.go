--- vendor/cloud.google.com/go/storage/option.go.orig	2026-10-02 20:06:38 UTC
+++ vendor/cloud.google.com/go/storage/option.go
@@ -21,7 +21,7 @@ import (
 
 	"cloud.google.com/go/storage/experimental"
 	storageinternal "cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
+
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 )
@@ -81,10 +81,10 @@ type storageConfig struct {
 	useJSONforReads        bool
 	readAPIWasSet          bool
 	disableClientMetrics   bool
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
@@ -202,10 +202,10 @@ type withMetricExporterConfig struct {
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
 
@@ -216,16 +216,16 @@ type withTestMetricReaderConfig struct {
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
 
@@ -233,7 +233,7 @@ func (w *withMeterProviderConfig) ApplyStorageOpt(c *s
 	c.meterProvider = w.meterProvider
 }
 
-func withTestMetricReader(ex *metric.ManualReader) option.ClientOption {
+func withTestMetricReader(ex any) option.ClientOption {
 	return &withTestMetricReaderConfig{metricReader: ex}
 }
 
