--- vendor/cloud.google.com/go/storage/experimental/experimental.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/experimental/experimental.go
@@ -25,7 +25,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 )
 
@@ -35,13 +34,6 @@ func WithMetricInterval(metricInterval time.Duration) 
 // When using Cloud Monitoring interval must be at minimum 1 [time.Minute].
 func WithMetricInterval(metricInterval time.Duration) option.ClientOption {
 	return internal.WithMetricInterval.(func(time.Duration) option.ClientOption)(metricInterval)
-}
-
-// WithMetricExporter provides a [option.ClientOption] that may be passed to [storage.NewGRPCClient].
-// Set an alternate client-side metric Exporter to emit metrics through.
-// Must implement [metric.Exporter]
-func WithMetricExporter(ex *metric.Exporter) option.ClientOption {
-	return internal.WithMetricExporter.(func(*metric.Exporter) option.ClientOption)(ex)
 }
 
 // WithReadStallTimeout provides a [option.ClientOption] that may be passed to [storage.NewClient].
