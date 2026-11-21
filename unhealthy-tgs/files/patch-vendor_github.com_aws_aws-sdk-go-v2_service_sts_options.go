--- vendor/github.com/aws/aws-sdk-go-v2/service/sts/options.go.orig	2026-11-21 14:46:46 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sts/options.go
@@ -13,7 +13,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -112,7 +111,7 @@ type Options struct {
 	RuntimeEnvironment aws.RuntimeEnvironment
 
 	// The client tracer provider.
-	TracerProvider tracing.TracerProvider
+	TracerProvider interface{}
 
 	// Signature Version 4a (SigV4a) Signer
 	httpSignerV4a httpSignerV4a
