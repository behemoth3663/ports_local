--- vendor/github.com/aws/aws-sdk-go-v2/service/synthetics/types/types.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/synthetics/types/types.go
@@ -297,7 +297,6 @@ type CanaryRunConfigInput struct {
 	// You can enable active tracing only for canaries that use version syn-nodejs-2.0
 	// or later for their canary runtime.
 	//
-	// [Canaries and X-Ray tracing]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_tracing.html
 	ActiveTracing *bool
 
 	// Specifies the keys and values to use for any environment variables used in the
@@ -341,7 +340,6 @@ type CanaryRunConfigOutput struct {
 // A structure that contains information about a canary run.
 type CanaryRunConfigOutput struct {
 
-	// Displays whether this canary run used active X-Ray tracing.
 	ActiveTracing *bool
 
 	// Specifies the amount of ephemeral storage (in MB) to allocate for the canary
