--- vendor/github.com/aws/aws-sdk-go-v2/service/lambda/api_op_CreateFunction.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lambda/api_op_CreateFunction.go
@@ -15,7 +15,6 @@ import (
 // deployment package is a .zip file archive or container image that contains your
 // function code. The execution role grants the function permission to use Amazon
 // Web Services services, such as Amazon CloudWatch Logs for log streaming and
-// X-Ray for request tracing.
 //
 // If the deployment package is a [container image], then you set the package type to Image . For a
 // container image, the code property must include the URI of a container image in
