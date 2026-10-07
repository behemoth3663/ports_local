--- cmd/terraform-mcp-server/init.go.orig	1979-11-29 21:00:00 UTC
+++ cmd/terraform-mcp-server/init.go
@@ -21,13 +21,11 @@ import (
 	"github.com/hashicorp/terraform-mcp-server/pkg/tools"
 	"github.com/hashicorp/terraform-mcp-server/pkg/toolsets"
 	"github.com/hashicorp/terraform-mcp-server/version"
-	instana "github.com/instana/go-sensor"
 	"github.com/mark3labs/mcp-go/server"
 	"github.com/modelcontextprotocol/go-sdk/mcp"
 	log "github.com/sirupsen/logrus"
 	"github.com/spf13/cobra"
 	"github.com/spf13/viper"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 type healthResponse struct {
@@ -282,22 +280,11 @@ func serverInit(ctx context.Context, hcServer *server.
 	return nil
 }
 
-// setupInstana initializes the Instana collector when INSTANA_ENABLED is set,
-// Once it is initialized, the application metrics such as (CPU,
-// memory, goroutines) will be collected automatically;
-func setupInstana(logger *log.Logger) instana.TracerLogger {
-	if os.Getenv("INSTANA_ENABLED") != "true" {
-		return nil
+func setupInstana(logger *log.Logger) bool {
+	if os.Getenv("INSTANA_ENABLED") == "true" {
+		logger.Warn("Instana tracing is disabled in this build")
 	}
-	serviceName := "terraform-mcp-server"
-	if n := os.Getenv("INSTANA_SERVICE_NAME"); n != "" {
-		serviceName = n
-	}
-	logger.Info("Instana instrumentation enabled")
-	return instana.InitCollector(&instana.Options{
-		Service: serviceName,
-		Tracer:  instana.DefaultTracerOptions(),
-	})
+	return false
 }
 
 func streamableHTTPServerInit(ctx context.Context, hcServer *server.MCPServer, logger *log.Logger, host string, port string, endpointPath string, heartbeatInterval time.Duration, organizationAllowlist []string, enabledToolsets []string) error {
@@ -306,7 +293,7 @@ func streamableHTTPServerInit(ctx context.Context, hcS
 	var handler http.Handler
 
 	// Initialize the Instana collector if enabled (nil when disabled).
-	instanaCollector := setupInstana(logger)
+	instanaEnabled := setupInstana(logger)
 
 	// Create StreamableHTTP server which implements the new streamable-http transport
 	// This is the modern MCP transport that supports both direct HTTP responses and SSE streams
@@ -403,13 +390,11 @@ func streamableHTTPServerInit(ctx context.Context, hcS
 
 	addr := fmt.Sprintf("%s:%s", host, port)
 	handler = mux
-	if enableOtelMetrics := os.Getenv("OTEL_METRICS_ENABLED"); enableOtelMetrics == "true" {
-		// Add http server instrumentation for standard server metrics
-		handler = otelhttp.NewHandler(handler, "terraform-mcp-server")
+	if os.Getenv("OTEL_METRICS_ENABLED") == "true" {
+		logger.Warn("OTEL HTTP instrumentation is disabled in this build")
 	}
-	if instanaCollector != nil {
-		// Wrapping the handler so incoming HTTP requests will be able to be traced by Instana
-		handler = instana.TracingHandlerFunc(instanaCollector, "", handler.ServeHTTP)
+	if instanaEnabled {
+		logger.Warn("Instana HTTP instrumentation is disabled in this build")
 	}
 
 	httpServer := &http.Server{
