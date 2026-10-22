--- vendor/cloud.google.com/go/auth/internal/transport/transport.go.orig	2026-10-22 19:52:07 UTC
+++ vendor/cloud.google.com/go/auth/internal/transport/transport.go
@@ -24,9 +24,13 @@ import (
 	"time"
 
 	"cloud.google.com/go/auth/credentials"
-	"go.opentelemetry.io/otel/attribute"
 )
 
+type TelemetryAttribute struct {
+	Key   string
+	Value string
+}
+
 // knownKeys provides keys for reading telemetry attributes from Context.
 // It provides an implicit contract with generated client library code
 // using the same keys. The keys in this collection should not be removed
@@ -42,15 +46,15 @@ var knownKeys = []string{
 }
 
 // StaticTelemetryAttributes selectively converts known keys from a map of
-// strings to Open Telemetry attributes.
-func StaticTelemetryAttributes(m map[string]string) []attribute.KeyValue {
-	var staticAttrs []attribute.KeyValue
+// strings to static telemetry attributes.
+func StaticTelemetryAttributes(m map[string]string) []TelemetryAttribute {
+	var staticAttrs []TelemetryAttribute
 	if m == nil {
 		return staticAttrs
 	}
 	for _, k := range knownKeys {
 		if v, ok := m[k]; ok {
-			staticAttrs = append(staticAttrs, attribute.String(k, v))
+			staticAttrs = append(staticAttrs, TelemetryAttribute{Key: k, Value: v})
 		}
 	}
 	return staticAttrs
