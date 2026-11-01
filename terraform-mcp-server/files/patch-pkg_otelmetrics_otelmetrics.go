--- pkg/otelmetrics/otelmetrics.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/otelmetrics/otelmetrics.go
@@ -5,16 +5,9 @@ import (
 
 import (
 	"context"
-	"fmt"
 
 	"github.com/hashicorp/terraform-mcp-server/pkg/client"
 	log "github.com/sirupsen/logrus"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
-	"go.opentelemetry.io/otel/metric"
-	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
-	"go.opentelemetry.io/otel/sdk/resource"
 )
 
 // Setup reads metrics config from the environment and, if metrics are
@@ -22,80 +15,18 @@ func Setup(logger *log.Logger) (client.MetricsConfig, 
 // returns a shutdown func; the func does nothing when metrics are disabled.
 func Setup(logger *log.Logger) (client.MetricsConfig, func()) {
 	metricsConfig := client.LoadMetricsConfigFromEnv(logger)
-	logger.Infof("Metrics enabled: %t endpoint: %s exportInterval: %s", metricsConfig.Enabled, metricsConfig.Endpoint, metricsConfig.ExportInterval)
-	if !metricsConfig.Enabled {
-		return metricsConfig, func() {}
+	if metricsConfig.Enabled {
+		logger.Warn("OTEL_METRICS_ENABLED=true is set, but OpenTelemetry metrics are disabled in this build")
 	}
 
-	// Context for metrics is for tracking the lifecycle of the metrics setup and shutdown, not tied to individual tool calls.
-	ctxMetrics := context.Background()
-	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
-		logger.Errorf("OTel Internal Error: %v", err)
-	}))
-
-	shutdown, err := Init(ctxMetrics, &metricsConfig, logger)
-	if err != nil {
-		logger.Errorf("Failed to initialize metrics: %v", err)
-		return metricsConfig, func() {}
-	}
-
-	return metricsConfig, shutdown
+	return metricsConfig, func() {}
 }
 
 // Init builds the OTel exporter, meter provider, and the tool/error/latency/
 // client-type counters used to record metrics, storing them on config.
 func Init(ctx context.Context, config *client.MetricsConfig, logger *log.Logger) (func(), error) {
-	logger.Infof("Initializing exporter and meter provider for OTel metrics...")
-	// Create the Exporter (Sends data to the Collector)
-	// exporter, err := stdoutmetric.New() // For stdio debugging
-	exporter, err := otlpmetrichttp.New(ctx, otlpmetrichttp.WithEndpoint(config.Endpoint), otlpmetrichttp.WithInsecure())
-	if err != nil {
-		return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
-	}
-	// Create the MeterProvider with a PeriodicReader
-	// The reader flushes metrics to the exporter periodically
-	resourceAttrs := resource.NewSchemaless(
-		attribute.String("service.name", config.ServiceName),
-		attribute.String("service.version", config.ServiceVersion),
-	)
-	config.MeterProvider = sdkmetric.NewMeterProvider(
-		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(config.ExportInterval))),
-		sdkmetric.WithResource(resourceAttrs),
-	)
-
-	// Set it as the global provider
-	otel.SetMeterProvider(config.MeterProvider)
-
-	meter := config.MeterProvider.Meter(config.ServiceName)
-
-	config.ToolCounter, err = meter.Int64Counter("mcp_tool_calls_total")
-	if err != nil {
-		return nil, fmt.Errorf("failed to create tool counter: %w", err)
-	}
-
-	config.ErrorCounter, err = meter.Int64Counter("mcp_tool_errors_total",
-		metric.WithDescription("Total number of failed tool calls"))
-	if err != nil {
-		return nil, fmt.Errorf("failed to create error counter: %w", err)
-	}
-
-	config.ToolCallLatencyBucket, err = meter.Float64Histogram("mcp_tool_duration_seconds",
-		metric.WithDescription("Duration of tool calls in seconds"),
-		metric.WithUnit("s"))
-	if err != nil {
-		return nil, fmt.Errorf("failed to create latency histogram: %w", err)
-	}
-
-	config.ClientTypeCounter, err = meter.Int64Counter("mcp_client_type_total",
-		metric.WithDescription("Total number of connections by client type"))
-	if err != nil {
-		return nil, fmt.Errorf("failed to create client type counter: %w", err)
-	}
-
-	return func() {
-		logger.Infof("Shutting down metrics exporter..")
-		if err := config.MeterProvider.Shutdown(ctx); err != nil {
-			logger.Errorf("Error shutting down meter provider: %v", err)
-		}
-	}, nil
+	_ = ctx
+	_ = config
+	logger.Warn("OpenTelemetry metrics initialization is disabled in this build")
+	return func() {}, nil
 }
