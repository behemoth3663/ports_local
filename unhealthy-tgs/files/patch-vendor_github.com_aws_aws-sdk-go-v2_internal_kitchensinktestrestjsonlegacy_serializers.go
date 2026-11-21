--- vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestrestjsonlegacy/serializers.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestrestjsonlegacy/serializers.go
@@ -10,7 +10,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -24,10 +23,8 @@ func (m *awsRestjson1_serializeOpGetResource) HandleSe
 func (m *awsRestjson1_serializeOpGetResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -65,7 +62,6 @@ func (m *awsRestjson1_serializeOpGetResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceInput(v *GetResourceInput, encoder *httpbinding.Encoder) error {
@@ -95,10 +91,8 @@ func (m *awsRestjson1_serializeOpGetStreamingResource)
 func (m *awsRestjson1_serializeOpGetStreamingResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -136,7 +130,6 @@ func (m *awsRestjson1_serializeOpGetStreamingResource)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStreamingResourceInput(v *GetStreamingResourceInput, encoder *httpbinding.Encoder) error {
@@ -166,10 +159,8 @@ func (m *awsRestjson1_serializeOpPutCompressedData) Ha
 func (m *awsRestjson1_serializeOpPutCompressedData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -214,7 +205,6 @@ func (m *awsRestjson1_serializeOpPutCompressedData) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutCompressedDataInput(v *PutCompressedDataInput, encoder *httpbinding.Encoder) error {
@@ -247,10 +237,8 @@ func (m *awsRestjson1_serializeOpSubscribeEvents) Hand
 func (m *awsRestjson1_serializeOpSubscribeEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -288,7 +276,6 @@ func (m *awsRestjson1_serializeOpSubscribeEvents) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSubscribeEventsInput(v *SubscribeEventsInput, encoder *httpbinding.Encoder) error {
