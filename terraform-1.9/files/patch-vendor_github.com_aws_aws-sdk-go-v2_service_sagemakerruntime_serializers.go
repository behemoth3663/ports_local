--- vendor/github.com/aws/aws-sdk-go-v2/service/sagemakerruntime/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sagemakerruntime/serializers.go
@@ -9,7 +9,6 @@ import (
 	smithy "github.com/aws/smithy-go"
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -23,10 +22,8 @@ func (m *awsRestjson1_serializeOpInvokeEndpoint) Handl
 func (m *awsRestjson1_serializeOpInvokeEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -76,7 +73,6 @@ func (m *awsRestjson1_serializeOpInvokeEndpoint) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeEndpointInput(v *InvokeEndpointInput, encoder *httpbinding.Encoder) error {
@@ -156,10 +152,8 @@ func (m *awsRestjson1_serializeOpInvokeEndpointAsync) 
 func (m *awsRestjson1_serializeOpInvokeEndpointAsync) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -197,7 +191,6 @@ func (m *awsRestjson1_serializeOpInvokeEndpointAsync) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeEndpointAsyncInput(v *InvokeEndpointAsyncInput, encoder *httpbinding.Encoder) error {
@@ -262,10 +255,8 @@ func (m *awsRestjson1_serializeOpInvokeEndpointWithRes
 func (m *awsRestjson1_serializeOpInvokeEndpointWithResponseStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -315,7 +306,6 @@ func (m *awsRestjson1_serializeOpInvokeEndpointWithRes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeEndpointWithResponseStreamInput(v *InvokeEndpointWithResponseStreamInput, encoder *httpbinding.Encoder) error {
