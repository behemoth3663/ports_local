--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrock/auth.go.orig	2026-11-21 14:46:33 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrock/auth.go
@@ -10,7 +10,6 @@ import (
 	smithyauth "github.com/aws/smithy-go/auth"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"slices"
 	"strings"
@@ -154,8 +153,6 @@ func (m *resolveAuthSchemeMiddleware) HandleFinalize(c
 func (m *resolveAuthSchemeMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveAuthScheme")
-	defer span.End()
 
 	params, err := bindAuthResolverParams(ctx, m.operation, getOperationInput(ctx), m.options)
 	if err != nil {
@@ -173,8 +170,6 @@ func (m *resolveAuthSchemeMiddleware) HandleFinalize(c
 
 	ctx = setResolvedAuthScheme(ctx, scheme)
 
-	span.SetProperty("auth.scheme_id", scheme.Scheme.SchemeID())
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
@@ -268,8 +263,7 @@ func (m *getIdentityMiddleware) HandleFinalize(ctx con
 func (m *getIdentityMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	innerCtx, span := tracing.StartSpan(ctx, "GetIdentity")
-	defer span.End()
+	innerCtx := ctx
 
 	rscheme := getResolvedAuthScheme(innerCtx)
 	if rscheme == nil {
@@ -294,7 +288,6 @@ func (m *getIdentityMiddleware) HandleFinalize(ctx con
 
 	ctx = setIdentity(ctx, identity)
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
@@ -320,8 +313,6 @@ func (m *signRequestMiddleware) HandleFinalize(ctx con
 func (m *signRequestMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "SignRequest")
-	defer span.End()
 
 	req, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
@@ -352,6 +343,5 @@ func (m *signRequestMiddleware) HandleFinalize(ctx con
 		return out, metadata, fmt.Errorf("sign request: %w", err)
 	}
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
