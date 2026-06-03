--- vendor/cloud.google.com/go/auth/internal/transport/transport.go.orig	2026-04-06 22:09:02 UTC
+++ vendor/cloud.google.com/go/auth/internal/transport/transport.go
@@ -24,7 +24,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/auth/credentials"
-	"go.opentelemetry.io/otel/attribute"
 )
 
 // knownKeys provides keys for reading telemetry attributes from Context.
@@ -43,18 +42,6 @@ var knownKeys = []string{
 
 // StaticTelemetryAttributes selectively converts known keys from a map of
 // strings to Open Telemetry attributes.
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
-}
 
 // CloneDetectOptions clones a user set detect option into some new memory that
 // we can internally manipulate before sending onto the detect package.
