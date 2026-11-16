--- vendor/github.com/aws/aws-sdk-go-v2/service/taxsettings/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/taxsettings/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteTaxRegistr
 func (m *awsRestjson1_serializeOpBatchDeleteTaxRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -73,7 +70,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteTaxRegistr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteTaxRegistrationInput(v *BatchDeleteTaxRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -108,10 +104,8 @@ func (m *awsRestjson1_serializeOpBatchGetTaxExemptions
 func (m *awsRestjson1_serializeOpBatchGetTaxExemptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -156,7 +150,6 @@ func (m *awsRestjson1_serializeOpBatchGetTaxExemptions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetTaxExemptionsInput(v *BatchGetTaxExemptionsInput, encoder *httpbinding.Encoder) error {
@@ -191,10 +184,8 @@ func (m *awsRestjson1_serializeOpBatchPutTaxRegistrati
 func (m *awsRestjson1_serializeOpBatchPutTaxRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -239,7 +230,6 @@ func (m *awsRestjson1_serializeOpBatchPutTaxRegistrati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchPutTaxRegistrationInput(v *BatchPutTaxRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -281,10 +271,8 @@ func (m *awsRestjson1_serializeOpDeleteSupplementalTax
 func (m *awsRestjson1_serializeOpDeleteSupplementalTaxRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -329,7 +317,6 @@ func (m *awsRestjson1_serializeOpDeleteSupplementalTax
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSupplementalTaxRegistrationInput(v *DeleteSupplementalTaxRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -362,10 +349,8 @@ func (m *awsRestjson1_serializeOpDeleteTaxRegistration
 func (m *awsRestjson1_serializeOpDeleteTaxRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -410,7 +395,6 @@ func (m *awsRestjson1_serializeOpDeleteTaxRegistration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTaxRegistrationInput(v *DeleteTaxRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -443,10 +427,8 @@ func (m *awsRestjson1_serializeOpGetTaxExemptionTypes)
 func (m *awsRestjson1_serializeOpGetTaxExemptionTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -480,7 +462,6 @@ func (m *awsRestjson1_serializeOpGetTaxExemptionTypes)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTaxExemptionTypesInput(v *GetTaxExemptionTypesInput, encoder *httpbinding.Encoder) error {
@@ -501,10 +482,8 @@ func (m *awsRestjson1_serializeOpGetTaxInheritance) Ha
 func (m *awsRestjson1_serializeOpGetTaxInheritance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -538,7 +517,6 @@ func (m *awsRestjson1_serializeOpGetTaxInheritance) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTaxInheritanceInput(v *GetTaxInheritanceInput, encoder *httpbinding.Encoder) error {
@@ -559,10 +537,8 @@ func (m *awsRestjson1_serializeOpGetTaxRegistration) H
 func (m *awsRestjson1_serializeOpGetTaxRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -607,7 +583,6 @@ func (m *awsRestjson1_serializeOpGetTaxRegistration) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTaxRegistrationInput(v *GetTaxRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -640,10 +615,8 @@ func (m *awsRestjson1_serializeOpGetTaxRegistrationDoc
 func (m *awsRestjson1_serializeOpGetTaxRegistrationDocument) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -688,7 +661,6 @@ func (m *awsRestjson1_serializeOpGetTaxRegistrationDoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTaxRegistrationDocumentInput(v *GetTaxRegistrationDocumentInput, encoder *httpbinding.Encoder) error {
@@ -730,10 +702,8 @@ func (m *awsRestjson1_serializeOpListSupplementalTaxRe
 func (m *awsRestjson1_serializeOpListSupplementalTaxRegistrations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -778,7 +748,6 @@ func (m *awsRestjson1_serializeOpListSupplementalTaxRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSupplementalTaxRegistrationsInput(v *ListSupplementalTaxRegistrationsInput, encoder *httpbinding.Encoder) error {
@@ -816,10 +785,8 @@ func (m *awsRestjson1_serializeOpListTaxExemptions) Ha
 func (m *awsRestjson1_serializeOpListTaxExemptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -864,7 +831,6 @@ func (m *awsRestjson1_serializeOpListTaxExemptions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTaxExemptionsInput(v *ListTaxExemptionsInput, encoder *httpbinding.Encoder) error {
@@ -902,10 +868,8 @@ func (m *awsRestjson1_serializeOpListTaxRegistrations)
 func (m *awsRestjson1_serializeOpListTaxRegistrations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -950,7 +914,6 @@ func (m *awsRestjson1_serializeOpListTaxRegistrations)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTaxRegistrationsInput(v *ListTaxRegistrationsInput, encoder *httpbinding.Encoder) error {
@@ -988,10 +951,8 @@ func (m *awsRestjson1_serializeOpPutSupplementalTaxReg
 func (m *awsRestjson1_serializeOpPutSupplementalTaxRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1036,7 +997,6 @@ func (m *awsRestjson1_serializeOpPutSupplementalTaxReg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutSupplementalTaxRegistrationInput(v *PutSupplementalTaxRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -1071,10 +1031,8 @@ func (m *awsRestjson1_serializeOpPutTaxExemption) Hand
 func (m *awsRestjson1_serializeOpPutTaxExemption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1119,7 +1077,6 @@ func (m *awsRestjson1_serializeOpPutTaxExemption) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTaxExemptionInput(v *PutTaxExemptionInput, encoder *httpbinding.Encoder) error {
@@ -1173,10 +1130,8 @@ func (m *awsRestjson1_serializeOpPutTaxInheritance) Ha
 func (m *awsRestjson1_serializeOpPutTaxInheritance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1221,7 +1176,6 @@ func (m *awsRestjson1_serializeOpPutTaxInheritance) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTaxInheritanceInput(v *PutTaxInheritanceInput, encoder *httpbinding.Encoder) error {
@@ -1254,10 +1208,8 @@ func (m *awsRestjson1_serializeOpPutTaxRegistration) H
 func (m *awsRestjson1_serializeOpPutTaxRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1302,7 +1254,6 @@ func (m *awsRestjson1_serializeOpPutTaxRegistration) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTaxRegistrationInput(v *PutTaxRegistrationInput, encoder *httpbinding.Encoder) error {
