--- vendor/cloud.google.com/go/auth/internal/transport/transport.go.orig	2026-10-12 21:23:13 UTC
+++ vendor/cloud.google.com/go/auth/internal/transport/transport.go
@@ -24,7 +24,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/auth/credentials"
-	"go.opentelemetry.io/otel/attribute"
 )
 
 // knownKeys provides keys for reading telemetry attributes from Context.
@@ -41,16 +40,15 @@ var knownKeys = []string{
 	"url.domain",
 }
 
-// StaticTelemetryAttributes selectively converts known keys from a map of
-// strings to Open Telemetry attributes.
-func StaticTelemetryAttributes(m map[string]string) []attribute.KeyValue {
-	var staticAttrs []attribute.KeyValue
+// StaticTelemetryAttributes selectively returns known telemetry attributes.
+func StaticTelemetryAttributes(m map[string]string) map[string]string {
+	staticAttrs := map[string]string{}
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
