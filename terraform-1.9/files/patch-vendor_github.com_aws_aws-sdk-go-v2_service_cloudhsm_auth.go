--- vendor/github.com/aws/aws-sdk-go-v2/service/cloudhsm/auth.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cloudhsm/auth.go
@@ -10,7 +10,6 @@ import (
 	smithyauth "github.com/aws/smithy-go/auth"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -147,8 +146,6 @@ func (m *resolveAuthSchemeMiddleware) HandleFinalize(c
 func (m *resolveAuthSchemeMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveAuthScheme")
-	defer span.End()
 
 	params := bindAuthResolverParams(ctx, m.operation, getOperationInput(ctx), m.options)
 	options, err := m.options.AuthSchemeResolver.ResolveAuthSchemes(ctx, params)
@@ -163,8 +160,6 @@ func (m *resolveAuthSchemeMiddleware) HandleFinalize(c
 
 	ctx = setResolvedAuthScheme(ctx, scheme)
 
-	span.SetProperty("auth.scheme_id", scheme.Scheme.SchemeID())
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
@@ -224,10 +219,8 @@ func (m *getIdentityMiddleware) HandleFinalize(ctx con
 func (m *getIdentityMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	innerCtx, span := tracing.StartSpan(ctx, "GetIdentity")
-	defer span.End()
 
-	rscheme := getResolvedAuthScheme(innerCtx)
+	rscheme := getResolvedAuthScheme(ctx)
 	if rscheme == nil {
 		return out, metadata, fmt.Errorf("no resolved auth scheme")
 	}
@@ -239,7 +232,7 @@ func (m *getIdentityMiddleware) HandleFinalize(ctx con
 
 	identity, err := timeOperationMetric(ctx, "client.call.resolve_identity_duration",
 		func() (smithyauth.Identity, error) {
-			return resolver.GetIdentity(innerCtx, rscheme.IdentityProperties)
+			return resolver.GetIdentity(ctx, rscheme.IdentityProperties)
 		},
 		func(o *metrics.RecordMetricOptions) {
 			o.Properties.Set("auth.scheme_id", rscheme.Scheme.SchemeID())
@@ -250,7 +243,6 @@ func (m *getIdentityMiddleware) HandleFinalize(ctx con
 
 	ctx = setIdentity(ctx, identity)
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
@@ -276,8 +268,6 @@ func (m *signRequestMiddleware) HandleFinalize(ctx con
 func (m *signRequestMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "SignRequest")
-	defer span.End()
 
 	req, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
@@ -308,6 +298,5 @@ func (m *signRequestMiddleware) HandleFinalize(ctx con
 		return out, metadata, fmt.Errorf("sign request: %w", err)
 	}
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
