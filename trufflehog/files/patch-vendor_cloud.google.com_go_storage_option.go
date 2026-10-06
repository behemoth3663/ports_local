--- vendor/cloud.google.com/go/storage/option.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/cloud.google.com/go/storage/option.go
@@ -21,7 +21,6 @@ import (
 
 	"cloud.google.com/go/storage/experimental"
 	storageinternal "cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 )
@@ -79,9 +78,9 @@ type storageConfig struct {
 	useJSONforReads        bool
 	readAPIWasSet          bool
 	disableClientMetrics   bool
-	metricExporter         *metric.Exporter
+	metricExporter         interface{}
 	metricInterval         time.Duration
-	manualReader           *metric.ManualReader
+	manualReader           interface{}
 	readStallTimeoutConfig *experimental.ReadStallTimeoutConfig
 	grpcBidiReads          bool
 	grpcAppendableUploads  bool
@@ -186,10 +185,10 @@ type withMetricExporterConfig struct {
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
 
@@ -200,10 +199,10 @@ type withTestMetricReaderConfig struct {
 type withTestMetricReaderConfig struct {
 	internaloption.EmbeddableAdapter
 	// reader override
-	metricReader *metric.ManualReader
+	metricReader interface{}
 }
 
-func withTestMetricReader(ex *metric.ManualReader) option.ClientOption {
+func withTestMetricReader(ex interface{}) option.ClientOption {
 	return &withTestMetricReaderConfig{metricReader: ex}
 }
 
