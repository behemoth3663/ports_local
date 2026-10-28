--- vendor/github.com/aws/aws-sdk-go-v2/aws/middleware/request_id_retriever.go.orig	2026-10-28 15:34:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/middleware/request_id_retriever.go
@@ -4,7 +4,6 @@ import (
 	"context"
 
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -47,8 +46,6 @@ func (m *RequestIDRetriever) HandleDeserialize(ctx con
 			// set reqID on metadata for successful responses.
 			SetRequestIDMetadata(&metadata, v)
 
-			span, _ := tracing.GetSpan(ctx)
-			span.SetProperty("aws.request_id", v)
 			break
 		}
 	}
