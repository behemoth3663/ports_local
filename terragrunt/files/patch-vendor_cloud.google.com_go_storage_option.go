--- vendor/cloud.google.com/go/storage/option.go.orig	2026-11-04 11:39:43 UTC
+++ vendor/cloud.google.com/go/storage/option.go
@@ -21,7 +21,6 @@ import (
 
 	"cloud.google.com/go/storage/experimental"
 	storageinternal "cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 )
@@ -86,10 +85,10 @@ type storageConfig struct {
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
@@ -213,10 +212,10 @@ type withMetricExporterConfig struct {
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
 
@@ -227,16 +226,16 @@ type withTestMetricReaderConfig struct {
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
 
@@ -244,7 +243,7 @@ func (w *withMeterProviderConfig) ApplyStorageOpt(c *s
 	c.meterProvider = w.meterProvider
 }
 
-func withTestMetricReader(ex *metric.ManualReader) option.ClientOption {
+func withTestMetricReader(ex any) option.ClientOption {
 	return &withTestMetricReaderConfig{metricReader: ex}
 }
 
