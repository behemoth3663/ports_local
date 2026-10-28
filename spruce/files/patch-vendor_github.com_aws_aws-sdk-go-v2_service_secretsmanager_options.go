--- vendor/github.com/aws/aws-sdk-go-v2/service/secretsmanager/options.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/secretsmanager/options.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -112,10 +111,6 @@ type Options struct {
 	// should not populate this structure programmatically, or rely on the values here
 	// within your applications.
 	RuntimeEnvironment aws.RuntimeEnvironment
-
-	// The client tracer provider.
-	TracerProvider tracing.TracerProvider
-
 	// The initial DefaultsMode used when the client options were constructed. If the
 	// DefaultsMode was set to aws.DefaultsModeAuto this will store what the resolved
 	// value was at that point in time.
