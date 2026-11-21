--- vendor/github.com/aws/aws-sdk-go-v2/aws/signer/v4/middleware.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/signer/v4/middleware.go
@@ -15,7 +15,6 @@ import (
 	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
 	"github.com/aws/aws-sdk-go-v2/internal/sdk"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -162,8 +161,6 @@ func (m *ComputePayloadSHA256) HandleFinalize(
 		return next.HandleFinalize(ctx, in)
 	}
 
-	_, span := tracing.StartSpan(ctx, "ComputePayloadSHA256")
-	defer span.End()
 
 	req, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
@@ -190,7 +187,6 @@ func (m *ComputePayloadSHA256) HandleFinalize(
 
 	ctx = SetPayloadHash(ctx, hex.EncodeToString(hash.Sum(nil)))
 
-	span.End()
 	return next.HandleFinalize(ctx, in)
 }
 
