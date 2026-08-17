--- vendor/cloud.google.com/go/storage/experimental/experimental.go.orig	2025-03-20 16:11:04 UTC
+++ vendor/cloud.google.com/go/storage/experimental/experimental.go
@@ -25,7 +25,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel/sdk/metric"
 	"google.golang.org/api/option"
 )
 
@@ -40,9 +39,6 @@ func WithMetricInterval(metricInterval time.Duration) 
 // WithMetricExporter provides a [option.ClientOption] that may be passed to [storage.NewGRPCClient].
 // Set an alternate client-side metric Exporter to emit metrics through.
 // Must implement [metric.Exporter]
-func WithMetricExporter(ex *metric.Exporter) option.ClientOption {
-	return internal.WithMetricExporter.(func(*metric.Exporter) option.ClientOption)(ex)
-}
 
 // WithReadStallTimeout provides a [option.ClientOption] that may be passed to [storage.NewClient].
 // It enables the client to retry stalled requests when starting a download from
