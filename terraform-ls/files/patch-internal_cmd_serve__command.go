--- internal/cmd/serve_command.go.orig	2023-10-12 21:03:18 UTC
+++ internal/cmd/serve_command.go
@@ -21,10 +21,6 @@ import (
 	"github.com/hashicorp/terraform-ls/internal/logging"
 	"github.com/hashicorp/terraform-ls/internal/pathtpl"
 	"github.com/mitchellh/cli"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type ServeCommand struct {
@@ -115,10 +111,6 @@ func (c *ServeCommand) Run(args []string) int {
 	var err error
 	shutdownFunc := func(context.Context) error { return nil }
 
-	// TODO: Currently unused until we decide how/where to export data
-	tp := trace.NewNoopTracerProvider()
-	otel.SetTracerProvider(tp)
-
 	if err != nil {
 		c.Ui.Error(fmt.Sprintf("Failed to init telemetry: %s", err))
 		return 1
@@ -151,17 +143,6 @@ func (c *ServeCommand) Run(args []string) int {
 	}
 
 	return 0
-}
-
-func (c *ServeCommand) otelResourceAttributes() []attribute.KeyValue {
-	return []attribute.KeyValue{
-		semconv.ServiceName("terraform-ls"),
-		semconv.ServiceVersion(c.Version),
-		attribute.Int("process.pid", os.Getpid()),
-		attribute.Int("runtime.NumCPU", runtime.NumCPU()),
-		attribute.Int("port", c.port),
-		attribute.Int("reqConcurrency", c.reqConcurrency),
-	}
 }
 
 type stopFunc func() error
