--- vendor/cloud.google.com/go/storage/trace.go.orig	2026-10-20 09:33:13 UTC
+++ vendor/cloud.google.com/go/storage/trace.go
@@ -21,10 +21,6 @@ import (
 
 	internalTrace "cloud.google.com/go/internal/trace"
 	"cloud.google.com/go/storage/internal"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	otelcodes "go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 const (
