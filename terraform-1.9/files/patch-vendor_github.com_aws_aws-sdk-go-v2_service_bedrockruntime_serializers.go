--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockruntime/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockruntime/serializers.go
@@ -14,7 +14,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -29,10 +28,8 @@ func (m *awsRestjson1_serializeOpApplyGuardrail) Handl
 func (m *awsRestjson1_serializeOpApplyGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -81,7 +78,6 @@ func (m *awsRestjson1_serializeOpApplyGuardrail) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsApplyGuardrailInput(v *ApplyGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -144,10 +140,8 @@ func (m *awsRestjson1_serializeOpConverse) HandleSeria
 func (m *awsRestjson1_serializeOpConverse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -196,7 +190,6 @@ func (m *awsRestjson1_serializeOpConverse) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsConverseInput(v *ConverseInput, encoder *httpbinding.Encoder) error {
@@ -303,10 +296,8 @@ func (m *awsRestjson1_serializeOpConverseStream) Handl
 func (m *awsRestjson1_serializeOpConverseStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -355,7 +346,6 @@ func (m *awsRestjson1_serializeOpConverseStream) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsConverseStreamInput(v *ConverseStreamInput, encoder *httpbinding.Encoder) error {
@@ -462,10 +452,8 @@ func (m *awsRestjson1_serializeOpGetAsyncInvoke) Handl
 func (m *awsRestjson1_serializeOpGetAsyncInvoke) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -503,7 +491,6 @@ func (m *awsRestjson1_serializeOpGetAsyncInvoke) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAsyncInvokeInput(v *GetAsyncInvokeInput, encoder *httpbinding.Encoder) error {
@@ -533,10 +520,8 @@ func (m *awsRestjson1_serializeOpInvokeModel) HandleSe
 func (m *awsRestjson1_serializeOpInvokeModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -586,7 +571,6 @@ func (m *awsRestjson1_serializeOpInvokeModel) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeModelInput(v *InvokeModelInput, encoder *httpbinding.Encoder) error {
@@ -646,10 +630,8 @@ func (m *awsRestjson1_serializeOpInvokeModelWithRespon
 func (m *awsRestjson1_serializeOpInvokeModelWithResponseStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -699,7 +681,6 @@ func (m *awsRestjson1_serializeOpInvokeModelWithRespon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeModelWithResponseStreamInput(v *InvokeModelWithResponseStreamInput, encoder *httpbinding.Encoder) error {
@@ -759,10 +740,8 @@ func (m *awsRestjson1_serializeOpListAsyncInvokes) Han
 func (m *awsRestjson1_serializeOpListAsyncInvokes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -800,7 +779,6 @@ func (m *awsRestjson1_serializeOpListAsyncInvokes) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAsyncInvokesInput(v *ListAsyncInvokesInput, encoder *httpbinding.Encoder) error {
@@ -849,10 +827,8 @@ func (m *awsRestjson1_serializeOpStartAsyncInvoke) Han
 func (m *awsRestjson1_serializeOpStartAsyncInvoke) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -897,7 +873,6 @@ func (m *awsRestjson1_serializeOpStartAsyncInvoke) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAsyncInvokeInput(v *StartAsyncInvokeInput, encoder *httpbinding.Encoder) error {
