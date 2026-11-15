--- vendor/github.com/aws/aws-sdk-go-v2/aws/signer/v4/middleware.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/signer/v4/middleware.go
@@ -15,7 +15,6 @@ import (
 	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
 	"github.com/aws/aws-sdk-go-v2/internal/sdk"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -161,10 +160,6 @@ func (m *ComputePayloadSHA256) HandleFinalize(
 	if GetPayloadHash(ctx) != "" {
 		return next.HandleFinalize(ctx, in)
 	}
-
-	_, span := tracing.StartSpan(ctx, "ComputePayloadSHA256")
-	defer span.End()
-
 	req, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &HashComputationError{
@@ -190,7 +185,6 @@ func (m *ComputePayloadSHA256) HandleFinalize(
 
 	ctx = SetPayloadHash(ctx, hex.EncodeToString(hash.Sum(nil)))
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
