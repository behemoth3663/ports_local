--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/api_op_InvokeAgentRuntime.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/api_op_InvokeAgentRuntime.go
@@ -57,7 +57,6 @@ type InvokeAgentRuntimeInput struct {
 	// application/json for JSON data.
 	Accept *string
 
-	// Additional context information for distributed tracing.
 	Baggage *string
 
 	// The MIME type of the input data in the payload. This tells the agent runtime
@@ -85,10 +84,8 @@ type InvokeAgentRuntimeInput struct {
 	// The trace identifier for request tracking.
 	TraceId *string
 
-	// The parent trace information for distributed tracing.
 	TraceParent *string
 
-	// The trace state information for distributed tracing.
 	TraceState *string
 
 	noSmithyDocumentSerde
@@ -102,7 +99,6 @@ type InvokeAgentRuntimeOutput struct {
 	// This member is required.
 	ContentType *string
 
-	// Additional context information for distributed tracing.
 	Baggage *string
 
 	// The version of the MCP protocol being used.
@@ -126,10 +122,8 @@ type InvokeAgentRuntimeOutput struct {
 	// The trace identifier for request tracking.
 	TraceId *string
 
-	// The parent trace information for distributed tracing.
 	TraceParent *string
 
-	// The trace state information for distributed tracing.
 	TraceState *string
 
 	// Metadata pertaining to the operation's result.
