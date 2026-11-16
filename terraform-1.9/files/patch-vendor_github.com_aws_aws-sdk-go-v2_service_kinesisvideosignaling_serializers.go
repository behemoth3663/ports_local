--- vendor/github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling/serializers.go
@@ -10,7 +10,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -24,10 +23,8 @@ func (m *awsRestjson1_serializeOpGetIceServerConfig) H
 func (m *awsRestjson1_serializeOpGetIceServerConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -72,7 +69,6 @@ func (m *awsRestjson1_serializeOpGetIceServerConfig) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIceServerConfigInput(v *GetIceServerConfigInput, encoder *httpbinding.Encoder) error {
@@ -120,10 +116,8 @@ func (m *awsRestjson1_serializeOpSendAlexaOfferToMaste
 func (m *awsRestjson1_serializeOpSendAlexaOfferToMaster) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -168,7 +162,6 @@ func (m *awsRestjson1_serializeOpSendAlexaOfferToMaste
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendAlexaOfferToMasterInput(v *SendAlexaOfferToMasterInput, encoder *httpbinding.Encoder) error {
