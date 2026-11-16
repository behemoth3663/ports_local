--- vendor/github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpDecryptData) HandleSe
 func (m *awsRestjson1_serializeOpDecryptData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpDecryptData) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDecryptDataInput(v *DecryptDataInput, encoder *httpbinding.Encoder) error {
@@ -133,10 +129,8 @@ func (m *awsRestjson1_serializeOpEncryptData) HandleSe
 func (m *awsRestjson1_serializeOpEncryptData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -185,7 +179,6 @@ func (m *awsRestjson1_serializeOpEncryptData) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEncryptDataInput(v *EncryptDataInput, encoder *httpbinding.Encoder) error {
@@ -241,10 +234,8 @@ func (m *awsRestjson1_serializeOpGenerateCardValidatio
 func (m *awsRestjson1_serializeOpGenerateCardValidationData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -289,7 +280,6 @@ func (m *awsRestjson1_serializeOpGenerateCardValidatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateCardValidationDataInput(v *GenerateCardValidationDataInput, encoder *httpbinding.Encoder) error {
@@ -339,10 +329,8 @@ func (m *awsRestjson1_serializeOpGenerateMac) HandleSe
 func (m *awsRestjson1_serializeOpGenerateMac) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -387,7 +375,6 @@ func (m *awsRestjson1_serializeOpGenerateMac) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateMacInput(v *GenerateMacInput, encoder *httpbinding.Encoder) error {
@@ -437,10 +424,8 @@ func (m *awsRestjson1_serializeOpGenerateMacEmvPinChan
 func (m *awsRestjson1_serializeOpGenerateMacEmvPinChange) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -485,7 +470,6 @@ func (m *awsRestjson1_serializeOpGenerateMacEmvPinChan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateMacEmvPinChangeInput(v *GenerateMacEmvPinChangeInput, encoder *httpbinding.Encoder) error {
@@ -550,10 +534,8 @@ func (m *awsRestjson1_serializeOpGeneratePinData) Hand
 func (m *awsRestjson1_serializeOpGeneratePinData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -598,7 +580,6 @@ func (m *awsRestjson1_serializeOpGeneratePinData) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGeneratePinDataInput(v *GeneratePinDataInput, encoder *httpbinding.Encoder) error {
@@ -665,10 +646,8 @@ func (m *awsRestjson1_serializeOpReEncryptData) Handle
 func (m *awsRestjson1_serializeOpReEncryptData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -717,7 +696,6 @@ func (m *awsRestjson1_serializeOpReEncryptData) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsReEncryptDataInput(v *ReEncryptDataInput, encoder *httpbinding.Encoder) error {
@@ -792,10 +770,8 @@ func (m *awsRestjson1_serializeOpTranslatePinData) Han
 func (m *awsRestjson1_serializeOpTranslatePinData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -840,7 +816,6 @@ func (m *awsRestjson1_serializeOpTranslatePinData) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTranslatePinDataInput(v *TranslatePinDataInput, encoder *httpbinding.Encoder) error {
@@ -925,10 +900,8 @@ func (m *awsRestjson1_serializeOpVerifyAuthRequestCryp
 func (m *awsRestjson1_serializeOpVerifyAuthRequestCryptogram) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -973,7 +946,6 @@ func (m *awsRestjson1_serializeOpVerifyAuthRequestCryp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsVerifyAuthRequestCryptogramInput(v *VerifyAuthRequestCryptogramInput, encoder *httpbinding.Encoder) error {
@@ -1035,10 +1007,8 @@ func (m *awsRestjson1_serializeOpVerifyCardValidationD
 func (m *awsRestjson1_serializeOpVerifyCardValidationData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1083,7 +1053,6 @@ func (m *awsRestjson1_serializeOpVerifyCardValidationD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsVerifyCardValidationDataInput(v *VerifyCardValidationDataInput, encoder *httpbinding.Encoder) error {
@@ -1133,10 +1102,8 @@ func (m *awsRestjson1_serializeOpVerifyMac) HandleSeri
 func (m *awsRestjson1_serializeOpVerifyMac) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1181,7 +1148,6 @@ func (m *awsRestjson1_serializeOpVerifyMac) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsVerifyMacInput(v *VerifyMacInput, encoder *httpbinding.Encoder) error {
@@ -1236,10 +1202,8 @@ func (m *awsRestjson1_serializeOpVerifyPinData) Handle
 func (m *awsRestjson1_serializeOpVerifyPinData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1284,7 +1248,6 @@ func (m *awsRestjson1_serializeOpVerifyPinData) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsVerifyPinDataInput(v *VerifyPinDataInput, encoder *httpbinding.Encoder) error {
