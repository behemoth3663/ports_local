--- vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestlegacy/deserializers.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestlegacy/deserializers.go
@@ -15,7 +15,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -48,10 +47,8 @@ func (m *awsAwsjson10_deserializeOpGetItem) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson10_deserializeOpErrorGetItem(response, &metadata)
@@ -167,10 +164,8 @@ func (m *awsAwsjson10_deserializeOpGetLists) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson10_deserializeOpErrorGetLists(response, &metadata)
@@ -283,10 +278,8 @@ func (m *awsAwsjson10_deserializeOpGetMaps) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson10_deserializeOpErrorGetMaps(response, &metadata)
@@ -399,10 +392,8 @@ func (m *awsAwsjson10_deserializeOpPutCompressedData) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson10_deserializeOpErrorPutCompressedData(response, &metadata)
@@ -516,10 +507,8 @@ func (m *awsAwsjson10_deserializeOpSubscribeEvents) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson10_deserializeOpErrorSubscribeEvents(response, &metadata)
