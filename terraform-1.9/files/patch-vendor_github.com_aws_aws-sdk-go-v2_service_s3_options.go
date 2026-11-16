--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/options.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3/options.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -138,7 +137,7 @@ type Options struct {
 	RuntimeEnvironment aws.RuntimeEnvironment
 
 	// The client tracer provider.
-	TracerProvider tracing.TracerProvider
+	TracerProvider interface{}
 
 	// Allows you to enable arn region support for the service.
 	UseARNRegion bool
