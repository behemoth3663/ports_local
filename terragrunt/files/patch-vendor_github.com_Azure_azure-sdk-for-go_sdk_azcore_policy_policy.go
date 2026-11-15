--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/policy/policy.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/policy/policy.go
@@ -11,7 +11,6 @@ import (
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 )
 
 // Policy represents an extensibility point for the Pipeline that can mutate the specified
@@ -51,8 +50,7 @@ type ClientOptions struct {
 	Telemetry TelemetryOptions
 
 	// TracingProvider configures the tracing provider.
-	// It defaults to a no-op tracer.
-	TracingProvider tracing.Provider
+	TracingProvider any
 
 	// Transport sets the transport for HTTP requests.
 	Transport Transporter
