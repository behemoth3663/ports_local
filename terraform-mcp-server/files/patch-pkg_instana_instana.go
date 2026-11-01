--- pkg/instana/instana.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/instana/instana.go
@@ -6,24 +6,13 @@ import (
 import (
 	"os"
 
-	instanasdk "github.com/instana/go-sensor"
 	log "github.com/sirupsen/logrus"
 )
 
-// Setup initializes the Instana collector when INSTANA_ENABLED is set. Once
-// initialized, application metrics (CPU, memory, goroutines) are collected
-// automatically.
-func Setup(logger *log.Logger) instanasdk.TracerLogger {
-	if os.Getenv("INSTANA_ENABLED") != "true" {
-		return nil
+// Setup validates whether Instana is requested and logs that this build does
+// not include Instana instrumentation.
+func Setup(logger *log.Logger) {
+	if os.Getenv("INSTANA_ENABLED") == "true" {
+		logger.Warn("INSTANA_ENABLED=true is set, but Instana instrumentation is disabled in this build")
 	}
-	serviceName := "terraform-mcp-server"
-	if n := os.Getenv("INSTANA_SERVICE_NAME"); n != "" {
-		serviceName = n
-	}
-	logger.Info("Instana instrumentation enabled")
-	return instanasdk.InitCollector(&instanasdk.Options{
-		Service: serviceName,
-		Tracer:  instanasdk.DefaultTracerOptions(),
-	})
 }
