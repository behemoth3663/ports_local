--- vendor/cloud.google.com/go/auth/internal/transport/transport.go.orig	2026-10-02 16:09:03 UTC
+++ vendor/cloud.google.com/go/auth/internal/transport/transport.go
@@ -24,7 +24,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/auth/credentials"
-	"go.opentelemetry.io/otel/attribute"
 )
 
 // knownKeys provides keys for reading telemetry attributes from Context.
@@ -43,14 +42,14 @@ var knownKeys = []string{
 
 // StaticTelemetryAttributes selectively converts known keys from a map of
 // strings to Open Telemetry attributes.
-func StaticTelemetryAttributes(m map[string]string) []attribute.KeyValue {
-	var staticAttrs []attribute.KeyValue
+func StaticTelemetryAttributes(m map[string]string) []string {
+	var staticAttrs []string
 	if m == nil {
 		return staticAttrs
 	}
 	for _, k := range knownKeys {
 		if v, ok := m[k]; ok {
-			staticAttrs = append(staticAttrs, attribute.String(k, v))
+			staticAttrs = append(staticAttrs, k+"="+v)
 		}
 	}
 	return staticAttrs
