--- vendor/cloud.google.com/go/auth/internal/transport/transport.go.orig	2026-10-02 20:06:38 UTC
+++ vendor/cloud.google.com/go/auth/internal/transport/transport.go
@@ -24,7 +24,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/auth/credentials"
-	"go.opentelemetry.io/otel/attribute"
 )
 
 // knownKeys provides keys for reading telemetry attributes from Context.
@@ -42,15 +41,15 @@ var knownKeys = []string{
 }
 
 // StaticTelemetryAttributes selectively converts known keys from a map of
-// strings to Open Telemetry attributes.
-func StaticTelemetryAttributes(m map[string]string) []attribute.KeyValue {
-	var staticAttrs []attribute.KeyValue
+// strings and returns them for non-OpenTelemetry telemetry handling.
+func StaticTelemetryAttributes(m map[string]string) map[string]string {
+	staticAttrs := make(map[string]string)
 	if m == nil {
 		return staticAttrs
 	}
 	for _, k := range knownKeys {
 		if v, ok := m[k]; ok {
-			staticAttrs = append(staticAttrs, attribute.String(k, v))
+			staticAttrs[k] = v
 		}
 	}
 	return staticAttrs
