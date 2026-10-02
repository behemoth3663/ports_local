--- vendor/cloud.google.com/go/auth/internal/transport/transport.go.orig	2026-10-02 15:35:42 UTC
+++ vendor/cloud.google.com/go/auth/internal/transport/transport.go
@@ -24,7 +24,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/auth/credentials"
-	"go.opentelemetry.io/otel/attribute"
 )
 
 // knownKeys provides keys for reading telemetry attributes from Context.
@@ -39,21 +38,6 @@ var knownKeys = []string{
 	"gcp.client.artifact",
 	"gcp.client.language",
 	"url.domain",
-}
-
-// StaticTelemetryAttributes selectively converts known keys from a map of
-// strings to Open Telemetry attributes.
-func StaticTelemetryAttributes(m map[string]string) []attribute.KeyValue {
-	var staticAttrs []attribute.KeyValue
-	if m == nil {
-		return staticAttrs
-	}
-	for _, k := range knownKeys {
-		if v, ok := m[k]; ok {
-			staticAttrs = append(staticAttrs, attribute.String(k, v))
-		}
-	}
-	return staticAttrs
 }
 
 // CloneDetectOptions clones a user set detect option into some new memory that
