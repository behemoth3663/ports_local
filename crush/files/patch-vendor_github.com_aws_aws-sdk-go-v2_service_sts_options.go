--- vendor/github.com/aws/aws-sdk-go-v2/service/sts/options.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sts/options.go
@@ -13,7 +13,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -110,10 +109,6 @@ type Options struct {
 	// should not populate this structure programmatically, or rely on the values here
 	// within your applications.
 	RuntimeEnvironment aws.RuntimeEnvironment
-
-	// The client tracer provider.
-	TracerProvider tracing.TracerProvider
-
 	// Signature Version 4a (SigV4a) Signer
 	httpSignerV4a httpSignerV4a
 
