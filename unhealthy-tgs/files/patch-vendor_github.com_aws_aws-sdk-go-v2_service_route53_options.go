--- vendor/github.com/aws/aws-sdk-go-v2/service/route53/options.go.orig	2026-11-21 14:46:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/route53/options.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -110,7 +109,7 @@ type Options struct {
 	RuntimeEnvironment aws.RuntimeEnvironment
 
 	// The client tracer provider.
-	TracerProvider tracing.TracerProvider
+	TracerProvider interface{}
 
 	// The initial DefaultsMode used when the client options were constructed. If the
 	// DefaultsMode was set to aws.DefaultsModeAuto this will store what the resolved
