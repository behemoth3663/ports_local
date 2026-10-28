--- vendor/github.com/aws/aws-sdk-go-v2/service/internal/s3shared/metadata_retriever.go.orig	2026-10-28 15:34:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/internal/s3shared/metadata_retriever.go
@@ -5,7 +5,6 @@ import (
 
 	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -31,8 +30,6 @@ func (m *metadataRetriever) HandleDeserialize(ctx cont
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	span, _ := tracing.GetSpan(ctx)
-
 	resp, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		// No raw response to wrap with.
@@ -43,14 +40,12 @@ func (m *metadataRetriever) HandleDeserialize(ctx cont
 	if v := resp.Header.Get("X-Amz-Request-Id"); len(v) != 0 {
 		// set reqID on metadata for successful responses.
 		awsmiddleware.SetRequestIDMetadata(&metadata, v)
-		span.SetProperty("aws.request_id", v)
 	}
 
 	// look up host-id
 	if v := resp.Header.Get("X-Amz-Id-2"); len(v) != 0 {
 		// set reqID on metadata for successful responses.
 		SetHostIDMetadata(&metadata, v)
-		span.SetProperty("aws.extended_request_id", v)
 	}
 
 	return out, metadata, err
