--- vendor/github.com/aws/aws-sdk-go-v2/service/pi/endpoints.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/pi/endpoints.go
@@ -16,7 +16,6 @@ import (
 	smithyendpoints "github.com/aws/smithy-go/endpoints"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 	"net/url"
@@ -484,8 +483,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 func (m *resolveEndpointV2Middleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveEndpoint")
-	defer span.End()
 
 	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
 		return next.HandleFinalize(ctx, in)
@@ -509,7 +506,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
 	}
 
-	span.SetProperty("client.call.resolved_endpoint", endpt.URI.String())
 
 	if endpt.URI.RawPath == "" && req.URL.RawPath != "" {
 		endpt.URI.RawPath = endpt.URI.Path
@@ -532,6 +528,5 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 		rscheme.SignerProperties.SetAll(&o.SignerProperties)
 	}
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
