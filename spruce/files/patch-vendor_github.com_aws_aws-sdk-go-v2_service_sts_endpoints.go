--- vendor/github.com/aws/aws-sdk-go-v2/service/sts/endpoints.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sts/endpoints.go
@@ -19,7 +19,6 @@ import (
 	"github.com/aws/smithy-go/endpoints/private/rulesfn"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 	"net/url"
@@ -627,8 +626,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 func (m *resolveEndpointV2Middleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveEndpoint")
-	defer span.End()
 
 	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
 		return next.HandleFinalize(ctx, in)
@@ -655,8 +652,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
 	}
 
-	span.SetProperty("client.call.resolved_endpoint", endpt.URI.String())
-
 	if endpt.URI.RawPath == "" && req.URL.RawPath != "" {
 		endpt.URI.RawPath = endpt.URI.Path
 	}
@@ -677,7 +672,5 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 	for _, o := range opts {
 		rscheme.SignerProperties.SetAll(&o.SignerProperties)
 	}
-
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
