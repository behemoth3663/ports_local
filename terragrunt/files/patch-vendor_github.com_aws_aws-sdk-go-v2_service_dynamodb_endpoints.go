--- vendor/github.com/aws/aws-sdk-go-v2/service/dynamodb/endpoints.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/dynamodb/endpoints.go
@@ -19,7 +19,6 @@ import (
 	"github.com/aws/smithy-go/endpoints/private/rulesfn"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 	"net/url"
@@ -1049,9 +1048,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 func (m *resolveEndpointV2Middleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveEndpoint")
-	defer span.End()
-
 	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
 		return next.HandleFinalize(ctx, in)
 	}
@@ -1080,10 +1076,7 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 	if err != nil {
 		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
 	}
-
-	span.SetProperty("client.call.resolved_endpoint", endpt.URI.String())
-
-	if endpt.URI.RawPath == "" && req.URL.RawPath != "" {
+if endpt.URI.RawPath == "" && req.URL.RawPath != "" {
 		endpt.URI.RawPath = endpt.URI.Path
 	}
 	req.URL.Scheme = endpt.URI.Scheme
@@ -1104,6 +1097,5 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 		rscheme.SignerProperties.SetAll(&o.SignerProperties)
 	}
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
