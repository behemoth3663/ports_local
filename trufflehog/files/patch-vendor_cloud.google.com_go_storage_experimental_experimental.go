--- vendor/cloud.google.com/go/storage/experimental/experimental.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/cloud.google.com/go/storage/experimental/experimental.go
@@ -25,7 +25,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 )
 
@@ -40,8 +39,8 @@ func WithMetricInterval(metricInterval time.Duration) 
 // WithMetricExporter provides a [option.ClientOption] that may be passed to [storage.NewGRPCClient].
 // Set an alternate client-side metric Exporter to emit metrics through.
 // Must implement [metric.Exporter]
-func WithMetricExporter(ex *metric.Exporter) option.ClientOption {
-	return internal.WithMetricExporter.(func(*metric.Exporter) option.ClientOption)(ex)
+func WithMetricExporter(ex interface{}) option.ClientOption {
+	return internal.WithMetricExporter.(func(interface{}) option.ClientOption)(ex)
 }
 
 // WithReadStallTimeout provides a [option.ClientOption] that may be passed to [storage.NewClient].
