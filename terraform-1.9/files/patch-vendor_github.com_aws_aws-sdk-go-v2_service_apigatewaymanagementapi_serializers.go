--- vendor/github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/serializers.go
@@ -9,7 +9,6 @@ import (
 	smithy "github.com/aws/smithy-go"
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -23,10 +22,8 @@ func (m *awsRestjson1_serializeOpDeleteConnection) Han
 func (m *awsRestjson1_serializeOpDeleteConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -64,7 +61,6 @@ func (m *awsRestjson1_serializeOpDeleteConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConnectionInput(v *DeleteConnectionInput, encoder *httpbinding.Encoder) error {
@@ -94,10 +90,8 @@ func (m *awsRestjson1_serializeOpGetConnection) Handle
 func (m *awsRestjson1_serializeOpGetConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -135,7 +129,6 @@ func (m *awsRestjson1_serializeOpGetConnection) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectionInput(v *GetConnectionInput, encoder *httpbinding.Encoder) error {
@@ -165,10 +158,8 @@ func (m *awsRestjson1_serializeOpPostToConnection) Han
 func (m *awsRestjson1_serializeOpPostToConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -218,7 +209,6 @@ func (m *awsRestjson1_serializeOpPostToConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPostToConnectionInput(v *PostToConnectionInput, encoder *httpbinding.Encoder) error {
