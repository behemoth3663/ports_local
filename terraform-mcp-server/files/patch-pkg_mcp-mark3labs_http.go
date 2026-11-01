--- pkg/mcp-mark3labs/http.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/mcp-mark3labs/http.go
@@ -17,16 +17,13 @@ import (
 	"time"
 
 	"github.com/hashicorp/terraform-mcp-server/pkg/client"
-	"github.com/hashicorp/terraform-mcp-server/pkg/instana"
 	"github.com/hashicorp/terraform-mcp-server/pkg/logging"
 	"github.com/hashicorp/terraform-mcp-server/pkg/toolsets"
 	"github.com/hashicorp/terraform-mcp-server/version"
-	instanasdk "github.com/instana/go-sensor"
 	"github.com/mark3labs/mcp-go/mcp"
 	"github.com/mark3labs/mcp-go/server"
 	log "github.com/sirupsen/logrus"
 	"github.com/spf13/cobra"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // maxHeaderBytes limits the total size of request headers. MCP clients only send a
@@ -176,9 +173,6 @@ func streamableHTTPServerInit(ctx context.Context, hcS
 	endpointPath = path.Join("/", endpointPath)
 	var handler http.Handler
 
-	// Initialize the Instana collector if enabled (nil when disabled).
-	instanaCollector := instana.Setup(logger)
-
 	// Create StreamableHTTP server which implements the new streamable-http transport
 	// This is the modern MCP transport that supports both direct HTTP responses and SSE streams
 	opts := []server.StreamableHTTPOption{
@@ -300,14 +294,6 @@ func streamableHTTPServerInit(ctx context.Context, hcS
 
 	addr := fmt.Sprintf("%s:%s", host, port)
 	handler = mux
-	if enableOtelMetrics := os.Getenv("OTEL_METRICS_ENABLED"); enableOtelMetrics == "true" {
-		// Add http server instrumentation for standard server metrics
-		handler = otelhttp.NewHandler(handler, "terraform-mcp-server")
-	}
-	if instanaCollector != nil {
-		// Wrapping the handler so incoming HTTP requests will be able to be traced by Instana
-		handler = instanasdk.TracingHandlerFunc(instanaCollector, "", handler.ServeHTTP)
-	}
 
 	httpServer := &http.Server{
 		Addr:              addr,
