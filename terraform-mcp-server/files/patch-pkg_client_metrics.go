--- pkg/client/metrics.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/client/metrics.go
@@ -8,24 +8,15 @@ import (
 	"github.com/hashicorp/terraform-mcp-server/version"
 	"github.com/mark3labs/mcp-go/mcp"
 	log "github.com/sirupsen/logrus"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
-	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
 )
 
 type MetricsConfig struct {
-	Enabled               bool
-	Endpoint              string                   // URL of your OTel Collector or backend
-	ExportInterval        time.Duration            // Controls the frequency of metric flushes
-	ServiceName           string                   // ServiceName identifies the source of the metrics (e.g., "terraform-mcp-server")
-	ServiceVersion        string                   // ServiceVersion helps track metrics across different deployments
-	MeterProvider         *sdkmetric.MeterProvider // MeterProvider is the OTel provider used to create instruments
-	Attributes            []attribute.KeyValue     // Attributes are global labels applied to every metric emitted
-	EnableRuntimeMetrics  bool                     // EnableRuntimeMetrics toggles the collection of Go runtime stats (GC, Memory)
-	ToolCounter           metric.Int64Counter      // ToolCounter tracks the total number of tool calls initiated
-	ErrorCounter          metric.Int64Counter      // Error count
-	ToolCallLatencyBucket metric.Float64Histogram  // Latency distribution
-	ClientTypeCounter     metric.Int64Counter      // Client type count (e.g. cli, cpi, vscode, web etc.)
+	Enabled              bool
+	Endpoint             string        // URL of your metrics endpoint
+	ExportInterval       time.Duration // Controls the frequency of metric flushes
+	ServiceName          string        // ServiceName identifies the source of the metrics (e.g., "terraform-mcp-server")
+	ServiceVersion       string        // ServiceVersion helps track metrics across different deployments
+	EnableRuntimeMetrics bool          // EnableRuntimeMetrics toggles the collection of Go runtime stats (GC, Memory)
 }
 
 type ClientInfo struct {
@@ -42,8 +33,6 @@ func DefaultMetricsConfig() MetricsConfig {
 		ExportInterval:       2 * time.Second,
 		ServiceName:          "terraform-mcp-server",
 		ServiceVersion:       version.GetHumanVersion(),
-		MeterProvider:        nil,
-		Attributes:           []attribute.KeyValue{},
 		EnableRuntimeMetrics: true,
 	}
 }
@@ -89,44 +78,25 @@ func RecordToolCall(ctx context.Context, startTime tim
 }
 
 func RecordToolCall(ctx context.Context, startTime time.Time, toolErr bool, id any, message *mcp.CallToolRequest, config MetricsConfig, logger *log.Logger) {
+	_ = ctx
+	_ = startTime
+	_ = toolErr
+	_ = id
+	_ = message
 	logger.Infof("Recording tool call for tool: %s id: %v", message.Params.Name, id)
-	if !config.Enabled || config.ToolCounter == nil {
-		logger.Debugf("Either metrics are not enabled or ToolCounter is NIL! Initialization failed.")
+	if !config.Enabled {
+		logger.Debugf("Metrics are not enabled.")
 		return
 	}
-	// Calculate latency
-	elapsed := time.Since(startTime).Seconds()
-
-	attrs := metric.WithAttributes(
-		attribute.String("tool.name", message.Params.Name),
-		attribute.String("service.name", config.ServiceName),
-		attribute.String("service.version", config.ServiceVersion),
-	)
-	// Record tool call count
-	config.ToolCounter.Add(ctx, 1, attrs)
-	// Record Latency (Histogram)
-	config.ToolCallLatencyBucket.Record(ctx, elapsed, attrs)
-	// Record errors if any
-	if toolErr == true {
-		config.ErrorCounter.Add(ctx, 1, attrs)
-		logger.Errorf("Recorded error for tool %s", message.Params.Name)
-	}
 }
 
 // RecordClientType records the type and version of the client making the tool call (e.g., CLI, VSCode, Web, etc.)
 func RecordClientType(ctx context.Context, ci ClientInfo, config MetricsConfig, logger *log.Logger) {
+	_ = ctx
+	_ = ci
 	logger.Infof("Recording client type for client: %s version: %s title: %s description: %s", ci.Name, ci.Version, ci.Title, ci.Description)
-	if !config.Enabled || config.ClientTypeCounter == nil {
-		logger.Debugf("Either metrics are not enabled or ClientTypeCounter is NIL! Initialization failed.")
+	if !config.Enabled {
+		logger.Debugf("Metrics are not enabled.")
 		return
 	}
-	attrs := metric.WithAttributes(
-		attribute.String("client.name", ci.Name),
-		attribute.String("client.version", ci.Version),
-		attribute.String("client.title", ci.Title),
-		attribute.String("client.description", ci.Description),
-		attribute.String("service.name", config.ServiceName),
-		attribute.String("service.version", config.ServiceVersion),
-	)
-	config.ClientTypeCounter.Add(ctx, 1, attrs)
 }
