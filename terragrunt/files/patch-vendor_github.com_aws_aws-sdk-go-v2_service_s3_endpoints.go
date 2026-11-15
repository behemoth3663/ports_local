--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/endpoints.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3/endpoints.go
@@ -21,7 +21,6 @@ import (
 	"github.com/aws/smithy-go/endpoints/private/rulesfn"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 	"net/url"
@@ -3490,9 +3489,6 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 func (m *resolveEndpointV2Middleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveEndpoint")
-	defer span.End()
-
 	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
 		return next.HandleFinalize(ctx, in)
 	}
@@ -3517,10 +3513,7 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
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
@@ -3546,6 +3539,5 @@ func (m *resolveEndpointV2Middleware) HandleFinalize(c
 	backend := s3cust.GetPropertiesBackend(&endpt.Properties)
 	ctx = internalcontext.SetS3Backend(ctx, backend)
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
