--- vendor/github.com/aws/aws-sdk-go-v2/service/transcribestreaming/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/transcribestreaming/serializers.go
@@ -12,7 +12,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpGetMedicalScribeStrea
 func (m *awsRestjson1_serializeOpGetMedicalScribeStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_serializeOpGetMedicalScribeStrea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMedicalScribeStreamInput(v *GetMedicalScribeStreamInput, encoder *httpbinding.Encoder) error {
@@ -97,10 +93,8 @@ func (m *awsRestjson1_serializeOpStartCallAnalyticsStr
 func (m *awsRestjson1_serializeOpStartCallAnalyticsStreamTranscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -140,7 +134,6 @@ func (m *awsRestjson1_serializeOpStartCallAnalyticsStr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartCallAnalyticsStreamTranscriptionInput(v *StartCallAnalyticsStreamTranscriptionInput, encoder *httpbinding.Encoder) error {
@@ -226,10 +219,8 @@ func (m *awsRestjson1_serializeOpStartMedicalScribeStr
 func (m *awsRestjson1_serializeOpStartMedicalScribeStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -269,7 +260,6 @@ func (m *awsRestjson1_serializeOpStartMedicalScribeStr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMedicalScribeStreamInput(v *StartMedicalScribeStreamInput, encoder *httpbinding.Encoder) error {
@@ -310,10 +300,8 @@ func (m *awsRestjson1_serializeOpStartMedicalStreamTra
 func (m *awsRestjson1_serializeOpStartMedicalStreamTranscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -353,7 +341,6 @@ func (m *awsRestjson1_serializeOpStartMedicalStreamTra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMedicalStreamTranscriptionInput(v *StartMedicalStreamTranscriptionInput, encoder *httpbinding.Encoder) error {
@@ -429,10 +416,8 @@ func (m *awsRestjson1_serializeOpStartStreamTranscript
 func (m *awsRestjson1_serializeOpStartStreamTranscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -472,7 +457,6 @@ func (m *awsRestjson1_serializeOpStartStreamTranscript
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartStreamTranscriptionInput(v *StartStreamTranscriptionInput, encoder *httpbinding.Encoder) error {
