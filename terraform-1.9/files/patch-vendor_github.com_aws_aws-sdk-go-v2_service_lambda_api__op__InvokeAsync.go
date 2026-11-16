--- vendor/github.com/aws/aws-sdk-go-v2/service/lambda/api_op_InvokeAsync.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lambda/api_op_InvokeAsync.go
@@ -16,7 +16,6 @@ import (
 // Invokes a function asynchronously.
 //
 // If you do use the InvokeAsync action, note that it doesn't support the use of
-// X-Ray active tracing. Trace ID is not propagated to the function, even if X-Ray
 // active tracing is turned on.
 //
 // Deprecated: This operation has been deprecated.
