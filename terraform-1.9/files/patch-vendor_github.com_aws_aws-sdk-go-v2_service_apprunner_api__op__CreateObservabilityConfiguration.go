--- vendor/github.com/aws/aws-sdk-go-v2/service/apprunner/api_op_CreateObservabilityConfiguration.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/apprunner/api_op_CreateObservabilityConfiguration.go
@@ -65,7 +65,6 @@ type CreateObservabilityConfigurationInput struct {
 	Tags []types.Tag
 
 	// The configuration of the tracing feature within this observability
-	// configuration. If you don't specify it, App Runner doesn't enable tracing.
 	TraceConfiguration *types.TraceConfiguration
 
 	noSmithyDocumentSerde
