--- vendor/github.com/aws/aws-sdk-go-v2/service/sso/auth.go.orig	2026-10-28 15:34:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sso/auth.go
@@ -10,7 +10,6 @@ import (
 	smithyauth "github.com/aws/smithy-go/auth"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"slices"
 	"strings"
@@ -176,8 +175,6 @@ func (m *resolveAuthSchemeMiddleware) HandleFinalize(c
 func (m *resolveAuthSchemeMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "ResolveAuthScheme")
-	defer span.End()
 
 	params, err := bindAuthResolverParams(ctx, m.operation, getOperationInput(ctx), m.options)
 	if err != nil {
@@ -194,9 +191,6 @@ func (m *resolveAuthSchemeMiddleware) HandleFinalize(c
 	}
 
 	ctx = setResolvedAuthScheme(ctx, scheme)
-
-	span.SetProperty("auth.scheme_id", scheme.Scheme.SchemeID())
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
@@ -290,10 +284,8 @@ func (m *getIdentityMiddleware) HandleFinalize(ctx con
 func (m *getIdentityMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	innerCtx, span := tracing.StartSpan(ctx, "GetIdentity")
-	defer span.End()
-
-	rscheme := getResolvedAuthScheme(innerCtx)
+	innerCtx := ctx
+rscheme := getResolvedAuthScheme(innerCtx)
 	if rscheme == nil {
 		return out, metadata, fmt.Errorf("no resolved auth scheme")
 	}
@@ -315,8 +307,6 @@ func (m *getIdentityMiddleware) HandleFinalize(ctx con
 	}
 
 	ctx = setIdentity(ctx, identity)
-
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
@@ -342,8 +332,6 @@ func (m *signRequestMiddleware) HandleFinalize(ctx con
 func (m *signRequestMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
 	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "SignRequest")
-	defer span.End()
 
 	req, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
@@ -373,7 +361,5 @@ func (m *signRequestMiddleware) HandleFinalize(ctx con
 	if err != nil {
 		return out, metadata, fmt.Errorf("sign request: %w", err)
 	}
-
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
