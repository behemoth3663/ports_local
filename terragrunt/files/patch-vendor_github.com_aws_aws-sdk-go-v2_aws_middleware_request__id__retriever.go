--- vendor/github.com/aws/aws-sdk-go-v2/aws/middleware/request_id_retriever.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/middleware/request_id_retriever.go
@@ -4,7 +4,6 @@ import (
 	"context"
 
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -47,9 +46,8 @@ func (m *RequestIDRetriever) HandleDeserialize(ctx con
 			// set reqID on metadata for successful responses.
 			SetRequestIDMetadata(&metadata, v)
 
-			span, _ := tracing.GetSpan(ctx)
-			span.SetProperty("aws.request_id", v)
-			break
+			_ , _ = 0, 0
+break
 		}
 	}
 
