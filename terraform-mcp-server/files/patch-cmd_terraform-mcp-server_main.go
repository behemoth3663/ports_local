--- cmd/terraform-mcp-server/main.go.orig	1979-11-29 21:00:00 UTC
+++ cmd/terraform-mcp-server/main.go
@@ -18,12 +18,6 @@ import (
 	"github.com/hashicorp/terraform-mcp-server/pkg/client"
 	"github.com/hashicorp/terraform-mcp-server/pkg/toolsets"
 	"github.com/hashicorp/terraform-mcp-server/version"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
-	"go.opentelemetry.io/otel/metric"
-	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
-	"go.opentelemetry.io/otel/sdk/resource"
 
 	"github.com/mark3labs/mcp-go/mcp"
 
@@ -417,77 +411,9 @@ func setupMetrics(logger *log.Logger) (client.MetricsC
 func setupMetrics(logger *log.Logger) (client.MetricsConfig, func()) {
 	metricsConfig := client.LoadMetricsConfigFromEnv(logger)
 	logger.Infof("Metrics enabled: %t endpoint: %s exportInterval: %s", metricsConfig.Enabled, metricsConfig.Endpoint, metricsConfig.ExportInterval)
-	if !metricsConfig.Enabled {
-		return metricsConfig, func() {}
+	if metricsConfig.Enabled {
+		logger.Warn("OTEL metrics are disabled in this build; continuing without telemetry exporters")
+		metricsConfig.Enabled = false
 	}
-
-	// Context for metrics is for tracking the lifecycle of the metrics setup and shutdown, not tied to individual tool calls.
-	ctxMetrics := context.Background()
-	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
-		logger.Errorf("OTel Internal Error: %v", err)
-	}))
-
-	shutdown, err := initMetrics(ctxMetrics, &metricsConfig, logger)
-	if err != nil {
-		logger.Errorf("Failed to initialize metrics: %v", err)
-		return metricsConfig, func() {}
-	}
-
-	return metricsConfig, shutdown
-}
-
-func initMetrics(ctx context.Context, config *client.MetricsConfig, logger *log.Logger) (func(), error) {
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
+	return metricsConfig, func() {}
 }
