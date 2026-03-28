--- vendor/cloud.google.com/go/storage/experimental/experimental.go.orig	2025-09-22 16:26:20 UTC
+++ vendor/cloud.google.com/go/storage/experimental/experimental.go
@@ -25,7 +25,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 )
 
@@ -41,9 +40,6 @@ func WithMetricInterval(metricInterval time.Duration) 
 // WithMetricExporter provides a [option.ClientOption] that may be passed to [storage.NewGRPCClient].
 // Set an alternate client-side metric Exporter to emit metrics through.
 // Must implement [metric.Exporter]. This option is ignored if WithMeterProvider is also set.
-func WithMetricExporter(ex *metric.Exporter) option.ClientOption {
-	return internal.WithMetricExporter.(func(*metric.Exporter) option.ClientOption)(ex)
-}
 
 // WithMeterProvider provides a [option.ClientOption] that may be passed to [storage.NewGRPCClient].
 // Set an alternate client-side meter provider to emit metrics through.
