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
+	Endpoint             string        // URL of your OTel Collector or backend
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
@@ -92,24 +81,14 @@ func RecordToolCallByName(ctx context.Context, startTi
 // shared by the mark3labs before/after-hook path and the go-sdk middleware
 // path (pkg/mcp-official/tools/middleware.Metrics).
 func RecordToolCallByName(ctx context.Context, startTime time.Time, toolErr bool, toolName string, config MetricsConfig, logger *log.Logger) {
-	if !config.Enabled || config.ToolCounter == nil {
-		logger.Debugf("Either metrics are not enabled or ToolCounter is NIL! Initialization failed.")
+	_ = ctx
+	if !config.Enabled {
 		return
 	}
-	// Calculate latency
 	elapsed := time.Since(startTime).Seconds()
-	attrs := metric.WithAttributes(
-		attribute.String("tool.name", toolName),
-		attribute.String("service.name", config.ServiceName),
-		attribute.String("service.version", config.ServiceVersion),
-	)
-	// Record tool call count
-	config.ToolCounter.Add(ctx, 1, attrs)
-	// Record Latency (Histogram)
-	config.ToolCallLatencyBucket.Record(ctx, elapsed, attrs)
+	logger.Debugf("Tool %s call completed in %.3fs", toolName, elapsed)
 	// Record errors if any
 	if toolErr {
-		config.ErrorCounter.Add(ctx, 1, attrs)
 		logger.Errorf("Recorded error for tool %s", toolName)
 	}
 }
@@ -121,18 +100,10 @@ func RecordClientType(ctx context.Context, ci ClientIn
 
 // RecordClientType records the type and version of the client making the tool call (e.g., CLI, VSCode, Web, etc.)
 func RecordClientType(ctx context.Context, ci ClientInfo, config MetricsConfig, logger *log.Logger) {
+	_ = ctx
 	logger.Infof("Recording client type for client: %s version: %s title: %s description: %s", ci.Name, ci.Version, ci.Title, ci.Description)
-	if !config.Enabled || config.ClientTypeCounter == nil {
-		logger.Debugf("Either metrics are not enabled or ClientTypeCounter is NIL! Initialization failed.")
+	if !config.Enabled {
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
+	logger.Debugf("Client type observed: %s %s", ci.Name, ci.Version)
 }
