--- vendor/github.com/aws/aws-sdk-go-v2/service/internal/s3shared/metadata_retriever.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/internal/s3shared/metadata_retriever.go
@@ -5,7 +5,6 @@ import (
 
 	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -31,7 +30,7 @@ func (m *metadataRetriever) HandleDeserialize(ctx cont
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	span, _ := tracing.GetSpan(ctx)
+	_ , _ = 0, 0
 
 	resp, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
@@ -43,15 +42,13 @@ func (m *metadataRetriever) HandleDeserialize(ctx cont
 	if v := resp.Header.Get("X-Amz-Request-Id"); len(v) != 0 {
 		// set reqID on metadata for successful responses.
 		awsmiddleware.SetRequestIDMetadata(&metadata, v)
-		span.SetProperty("aws.request_id", v)
-	}
+}
 
 	// look up host-id
 	if v := resp.Header.Get("X-Amz-Id-2"); len(v) != 0 {
 		// set reqID on metadata for successful responses.
 		SetHostIDMetadata(&metadata, v)
-		span.SetProperty("aws.extended_request_id", v)
-	}
+}
 
 	return out, metadata, err
 }
