--- vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestrestjsonlegacy/deserializers.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestrestjsonlegacy/deserializers.go
@@ -15,7 +15,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpGetResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetResource(response, &metadata)
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpGetResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -212,10 +208,8 @@ func (m *awsRestjson1_deserializeOpGetStreamingResourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetStreamingResource(response, &metadata)
@@ -228,7 +222,6 @@ func (m *awsRestjson1_deserializeOpGetStreamingResourc
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -322,10 +315,8 @@ func (m *awsRestjson1_deserializeOpPutCompressedData) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutCompressedData(response, &metadata)
@@ -333,7 +324,6 @@ func (m *awsRestjson1_deserializeOpPutCompressedData) 
 	output := &PutCompressedDataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -416,10 +406,8 @@ func (m *awsRestjson1_deserializeOpSubscribeEvents) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorSubscribeEvents(response, &metadata)
@@ -427,7 +415,6 @@ func (m *awsRestjson1_deserializeOpSubscribeEvents) Ha
 	output := &SubscribeEventsInitialReply{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
