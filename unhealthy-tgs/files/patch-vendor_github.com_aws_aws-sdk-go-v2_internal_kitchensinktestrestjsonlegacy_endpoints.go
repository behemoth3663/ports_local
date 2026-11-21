--- vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestrestjsonlegacy/endpoints.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestrestjsonlegacy/endpoints.go
@@ -15,7 +15,6 @@ import (
 	smithyendpoints "github.com/aws/smithy-go/endpoints"
 	"github.com/aws/smithy-go/endpoints/private/rulesfn"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 	"net/url"
@@ -335,8 +334,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 func (m *resolveEndpointV2Middleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveEndpoint")
-	defer span.End()
 
 	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
 		return next.HandleFinalize(ctx, in)
@@ -363,7 +360,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
 	}
 
-	span.SetProperty("client.call.resolved_endpoint", endpt.URI.String())
 
 	if endpt.URI.RawPath == "" && req.URL.RawPath != "" {
 		endpt.URI.RawPath = endpt.URI.Path
@@ -386,6 +382,5 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 		rscheme.SignerProperties.SetAll(&o.SignerProperties)
 	}
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
