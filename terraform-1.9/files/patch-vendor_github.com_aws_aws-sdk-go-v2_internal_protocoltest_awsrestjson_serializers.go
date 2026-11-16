--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/serializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 	"net/http"
@@ -35,10 +34,8 @@ func (m *awsRestjson1_serializeOpAllQueryStringTypes) 
 func (m *awsRestjson1_serializeOpAllQueryStringTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -76,7 +73,6 @@ func (m *awsRestjson1_serializeOpAllQueryStringTypes) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAllQueryStringTypesInput(v *AllQueryStringTypesInput, encoder *httpbinding.Encoder) error {
@@ -206,10 +202,8 @@ func (m *awsRestjson1_serializeOpConstantAndVariableQu
 func (m *awsRestjson1_serializeOpConstantAndVariableQueryString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -247,7 +241,6 @@ func (m *awsRestjson1_serializeOpConstantAndVariableQu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsConstantAndVariableQueryStringInput(v *ConstantAndVariableQueryStringInput, encoder *httpbinding.Encoder) error {
@@ -276,10 +269,8 @@ func (m *awsRestjson1_serializeOpConstantQueryString) 
 func (m *awsRestjson1_serializeOpConstantQueryString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -317,7 +308,6 @@ func (m *awsRestjson1_serializeOpConstantQueryString) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsConstantQueryStringInput(v *ConstantQueryStringInput, encoder *httpbinding.Encoder) error {
@@ -347,10 +337,8 @@ func (m *awsRestjson1_serializeOpContentTypeParameters
 func (m *awsRestjson1_serializeOpContentTypeParameters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -395,7 +383,6 @@ func (m *awsRestjson1_serializeOpContentTypeParameters
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsContentTypeParametersInput(v *ContentTypeParametersInput, encoder *httpbinding.Encoder) error {
@@ -428,10 +415,8 @@ func (m *awsRestjson1_serializeOpDatetimeOffsets) Hand
 func (m *awsRestjson1_serializeOpDatetimeOffsets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -465,7 +450,6 @@ func (m *awsRestjson1_serializeOpDatetimeOffsets) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDatetimeOffsetsInput(v *DatetimeOffsetsInput, encoder *httpbinding.Encoder) error {
@@ -486,10 +470,8 @@ func (m *awsRestjson1_serializeOpDocumentType) HandleS
 func (m *awsRestjson1_serializeOpDocumentType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -534,7 +516,6 @@ func (m *awsRestjson1_serializeOpDocumentType) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDocumentTypeInput(v *DocumentTypeInput, encoder *httpbinding.Encoder) error {
@@ -574,10 +555,8 @@ func (m *awsRestjson1_serializeOpDocumentTypeAsMapValu
 func (m *awsRestjson1_serializeOpDocumentTypeAsMapValue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -622,7 +601,6 @@ func (m *awsRestjson1_serializeOpDocumentTypeAsMapValu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDocumentTypeAsMapValueInput(v *DocumentTypeAsMapValueInput, encoder *httpbinding.Encoder) error {
@@ -657,10 +635,8 @@ func (m *awsRestjson1_serializeOpDocumentTypeAsPayload
 func (m *awsRestjson1_serializeOpDocumentTypeAsPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -710,7 +686,6 @@ func (m *awsRestjson1_serializeOpDocumentTypeAsPayload
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDocumentTypeAsPayloadInput(v *DocumentTypeAsPayloadInput, encoder *httpbinding.Encoder) error {
@@ -731,10 +706,8 @@ func (m *awsRestjson1_serializeOpEmptyInputAndEmptyOut
 func (m *awsRestjson1_serializeOpEmptyInputAndEmptyOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -768,7 +741,6 @@ func (m *awsRestjson1_serializeOpEmptyInputAndEmptyOut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEmptyInputAndEmptyOutputInput(v *EmptyInputAndEmptyOutputInput, encoder *httpbinding.Encoder) error {
@@ -789,10 +761,8 @@ func (m *awsRestjson1_serializeOpEndpointOperation) Ha
 func (m *awsRestjson1_serializeOpEndpointOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -826,7 +796,6 @@ func (m *awsRestjson1_serializeOpEndpointOperation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEndpointOperationInput(v *EndpointOperationInput, encoder *httpbinding.Encoder) error {
@@ -847,10 +816,8 @@ func (m *awsRestjson1_serializeOpEndpointWithHostLabel
 func (m *awsRestjson1_serializeOpEndpointWithHostLabelOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -895,7 +862,6 @@ func (m *awsRestjson1_serializeOpEndpointWithHostLabel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEndpointWithHostLabelOperationInput(v *EndpointWithHostLabelOperationInput, encoder *httpbinding.Encoder) error {
@@ -928,10 +894,8 @@ func (m *awsRestjson1_serializeOpFractionalSeconds) Ha
 func (m *awsRestjson1_serializeOpFractionalSeconds) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -965,7 +929,6 @@ func (m *awsRestjson1_serializeOpFractionalSeconds) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsFractionalSecondsInput(v *FractionalSecondsInput, encoder *httpbinding.Encoder) error {
@@ -986,10 +949,8 @@ func (m *awsRestjson1_serializeOpGreetingWithErrors) H
 func (m *awsRestjson1_serializeOpGreetingWithErrors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1023,7 +984,6 @@ func (m *awsRestjson1_serializeOpGreetingWithErrors) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGreetingWithErrorsInput(v *GreetingWithErrorsInput, encoder *httpbinding.Encoder) error {
@@ -1044,10 +1004,8 @@ func (m *awsRestjson1_serializeOpHostWithPathOperation
 func (m *awsRestjson1_serializeOpHostWithPathOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1081,7 +1039,6 @@ func (m *awsRestjson1_serializeOpHostWithPathOperation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHostWithPathOperationInput(v *HostWithPathOperationInput, encoder *httpbinding.Encoder) error {
@@ -1102,10 +1059,8 @@ func (m *awsRestjson1_serializeOpHttpChecksumRequired)
 func (m *awsRestjson1_serializeOpHttpChecksumRequired) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1150,7 +1105,6 @@ func (m *awsRestjson1_serializeOpHttpChecksumRequired)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpChecksumRequiredInput(v *HttpChecksumRequiredInput, encoder *httpbinding.Encoder) error {
@@ -1183,10 +1137,8 @@ func (m *awsRestjson1_serializeOpHttpEmptyPrefixHeader
 func (m *awsRestjson1_serializeOpHttpEmptyPrefixHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1224,7 +1176,6 @@ func (m *awsRestjson1_serializeOpHttpEmptyPrefixHeader
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpEmptyPrefixHeadersInput(v *HttpEmptyPrefixHeadersInput, encoder *httpbinding.Encoder) error {
@@ -1257,10 +1208,8 @@ func (m *awsRestjson1_serializeOpHttpEnumPayload) Hand
 func (m *awsRestjson1_serializeOpHttpEnumPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1306,7 +1255,6 @@ func (m *awsRestjson1_serializeOpHttpEnumPayload) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpEnumPayloadInput(v *HttpEnumPayloadInput, encoder *httpbinding.Encoder) error {
@@ -1327,10 +1275,8 @@ func (m *awsRestjson1_serializeOpHttpPayloadTraits) Ha
 func (m *awsRestjson1_serializeOpHttpPayloadTraits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1380,7 +1326,6 @@ func (m *awsRestjson1_serializeOpHttpPayloadTraits) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpPayloadTraitsInput(v *HttpPayloadTraitsInput, encoder *httpbinding.Encoder) error {
@@ -1406,10 +1351,8 @@ func (m *awsRestjson1_serializeOpHttpPayloadTraitsWith
 func (m *awsRestjson1_serializeOpHttpPayloadTraitsWithMediaType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1459,7 +1402,6 @@ func (m *awsRestjson1_serializeOpHttpPayloadTraitsWith
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpPayloadTraitsWithMediaTypeInput(v *HttpPayloadTraitsWithMediaTypeInput, encoder *httpbinding.Encoder) error {
@@ -1485,10 +1427,8 @@ func (m *awsRestjson1_serializeOpHttpPayloadWithStruct
 func (m *awsRestjson1_serializeOpHttpPayloadWithStructure) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1546,7 +1486,6 @@ func (m *awsRestjson1_serializeOpHttpPayloadWithStruct
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpPayloadWithStructureInput(v *HttpPayloadWithStructureInput, encoder *httpbinding.Encoder) error {
@@ -1567,10 +1506,8 @@ func (m *awsRestjson1_serializeOpHttpPayloadWithUnion)
 func (m *awsRestjson1_serializeOpHttpPayloadWithUnion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1620,7 +1557,6 @@ func (m *awsRestjson1_serializeOpHttpPayloadWithUnion)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpPayloadWithUnionInput(v *HttpPayloadWithUnionInput, encoder *httpbinding.Encoder) error {
@@ -1641,10 +1577,8 @@ func (m *awsRestjson1_serializeOpHttpPrefixHeaders) Ha
 func (m *awsRestjson1_serializeOpHttpPrefixHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1682,7 +1616,6 @@ func (m *awsRestjson1_serializeOpHttpPrefixHeaders) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpPrefixHeadersInput(v *HttpPrefixHeadersInput, encoder *httpbinding.Encoder) error {
@@ -1715,10 +1648,8 @@ func (m *awsRestjson1_serializeOpHttpPrefixHeadersInRe
 func (m *awsRestjson1_serializeOpHttpPrefixHeadersInResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1752,7 +1683,6 @@ func (m *awsRestjson1_serializeOpHttpPrefixHeadersInRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpPrefixHeadersInResponseInput(v *HttpPrefixHeadersInResponseInput, encoder *httpbinding.Encoder) error {
@@ -1773,10 +1703,8 @@ func (m *awsRestjson1_serializeOpHttpRequestWithFloatL
 func (m *awsRestjson1_serializeOpHttpRequestWithFloatLabels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1814,7 +1742,6 @@ func (m *awsRestjson1_serializeOpHttpRequestWithFloatL
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpRequestWithFloatLabelsInput(v *HttpRequestWithFloatLabelsInput, encoder *httpbinding.Encoder) error {
@@ -1853,10 +1780,8 @@ func (m *awsRestjson1_serializeOpHttpRequestWithGreedy
 func (m *awsRestjson1_serializeOpHttpRequestWithGreedyLabelInPath) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1894,7 +1819,6 @@ func (m *awsRestjson1_serializeOpHttpRequestWithGreedy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpRequestWithGreedyLabelInPathInput(v *HttpRequestWithGreedyLabelInPathInput, encoder *httpbinding.Encoder) error {
@@ -1933,10 +1857,8 @@ func (m *awsRestjson1_serializeOpHttpRequestWithLabels
 func (m *awsRestjson1_serializeOpHttpRequestWithLabels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1974,7 +1896,6 @@ func (m *awsRestjson1_serializeOpHttpRequestWithLabels
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpRequestWithLabelsInput(v *HttpRequestWithLabelsInput, encoder *httpbinding.Encoder) error {
@@ -2067,10 +1988,8 @@ func (m *awsRestjson1_serializeOpHttpRequestWithLabels
 func (m *awsRestjson1_serializeOpHttpRequestWithLabelsAndTimestampFormat) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2108,7 +2027,6 @@ func (m *awsRestjson1_serializeOpHttpRequestWithLabels
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpRequestWithLabelsAndTimestampFormatInput(v *HttpRequestWithLabelsAndTimestampFormatInput, encoder *httpbinding.Encoder) error {
@@ -2192,10 +2110,8 @@ func (m *awsRestjson1_serializeOpHttpRequestWithRegexL
 func (m *awsRestjson1_serializeOpHttpRequestWithRegexLiteral) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2233,7 +2149,6 @@ func (m *awsRestjson1_serializeOpHttpRequestWithRegexL
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpRequestWithRegexLiteralInput(v *HttpRequestWithRegexLiteralInput, encoder *httpbinding.Encoder) error {
@@ -2263,10 +2178,8 @@ func (m *awsRestjson1_serializeOpHttpResponseCode) Han
 func (m *awsRestjson1_serializeOpHttpResponseCode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2300,7 +2213,6 @@ func (m *awsRestjson1_serializeOpHttpResponseCode) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpResponseCodeInput(v *HttpResponseCodeInput, encoder *httpbinding.Encoder) error {
@@ -2321,10 +2233,8 @@ func (m *awsRestjson1_serializeOpHttpStringPayload) Ha
 func (m *awsRestjson1_serializeOpHttpStringPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2370,7 +2280,6 @@ func (m *awsRestjson1_serializeOpHttpStringPayload) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsHttpStringPayloadInput(v *HttpStringPayloadInput, encoder *httpbinding.Encoder) error {
@@ -2391,10 +2300,8 @@ func (m *awsRestjson1_serializeOpIgnoreQueryParamsInRe
 func (m *awsRestjson1_serializeOpIgnoreQueryParamsInResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2428,7 +2335,6 @@ func (m *awsRestjson1_serializeOpIgnoreQueryParamsInRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsIgnoreQueryParamsInResponseInput(v *IgnoreQueryParamsInResponseInput, encoder *httpbinding.Encoder) error {
@@ -2449,10 +2355,8 @@ func (m *awsRestjson1_serializeOpInputAndOutputWithHea
 func (m *awsRestjson1_serializeOpInputAndOutputWithHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2490,7 +2394,6 @@ func (m *awsRestjson1_serializeOpInputAndOutputWithHea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInputAndOutputWithHeadersInput(v *InputAndOutputWithHeadersInput, encoder *httpbinding.Encoder) error {
@@ -2665,10 +2568,8 @@ func (m *awsRestjson1_serializeOpJsonBlobs) HandleSeri
 func (m *awsRestjson1_serializeOpJsonBlobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2713,7 +2614,6 @@ func (m *awsRestjson1_serializeOpJsonBlobs) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsJsonBlobsInput(v *JsonBlobsInput, encoder *httpbinding.Encoder) error {
@@ -2746,10 +2646,8 @@ func (m *awsRestjson1_serializeOpJsonEnums) HandleSeri
 func (m *awsRestjson1_serializeOpJsonEnums) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2794,7 +2692,6 @@ func (m *awsRestjson1_serializeOpJsonEnums) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsJsonEnumsInput(v *JsonEnumsInput, encoder *httpbinding.Encoder) error {
@@ -2858,10 +2755,8 @@ func (m *awsRestjson1_serializeOpJsonIntEnums) HandleS
 func (m *awsRestjson1_serializeOpJsonIntEnums) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2906,7 +2801,6 @@ func (m *awsRestjson1_serializeOpJsonIntEnums) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsJsonIntEnumsInput(v *JsonIntEnumsInput, encoder *httpbinding.Encoder) error {
@@ -2970,10 +2864,8 @@ func (m *awsRestjson1_serializeOpJsonLists) HandleSeri
 func (m *awsRestjson1_serializeOpJsonLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3018,7 +2910,6 @@ func (m *awsRestjson1_serializeOpJsonLists) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsJsonListsInput(v *JsonListsInput, encoder *httpbinding.Encoder) error {
@@ -3109,10 +3000,8 @@ func (m *awsRestjson1_serializeOpJsonMaps) HandleSeria
 func (m *awsRestjson1_serializeOpJsonMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3157,7 +3046,6 @@ func (m *awsRestjson1_serializeOpJsonMaps) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsJsonMapsInput(v *JsonMapsInput, encoder *httpbinding.Encoder) error {
@@ -3220,10 +3108,8 @@ func (m *awsRestjson1_serializeOpJsonTimestamps) Handl
 func (m *awsRestjson1_serializeOpJsonTimestamps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3268,7 +3154,6 @@ func (m *awsRestjson1_serializeOpJsonTimestamps) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsJsonTimestampsInput(v *JsonTimestampsInput, encoder *httpbinding.Encoder) error {
@@ -3331,10 +3216,8 @@ func (m *awsRestjson1_serializeOpJsonUnions) HandleSer
 func (m *awsRestjson1_serializeOpJsonUnions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3379,7 +3262,6 @@ func (m *awsRestjson1_serializeOpJsonUnions) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsJsonUnionsInput(v *JsonUnionsInput, encoder *httpbinding.Encoder) error {
@@ -3414,10 +3296,8 @@ func (m *awsRestjson1_serializeOpMalformedAcceptWithBo
 func (m *awsRestjson1_serializeOpMalformedAcceptWithBody) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3451,7 +3331,6 @@ func (m *awsRestjson1_serializeOpMalformedAcceptWithBo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedAcceptWithBodyInput(v *MalformedAcceptWithBodyInput, encoder *httpbinding.Encoder) error {
@@ -3472,10 +3351,8 @@ func (m *awsRestjson1_serializeOpMalformedAcceptWithGe
 func (m *awsRestjson1_serializeOpMalformedAcceptWithGenericString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3509,7 +3386,6 @@ func (m *awsRestjson1_serializeOpMalformedAcceptWithGe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedAcceptWithGenericStringInput(v *MalformedAcceptWithGenericStringInput, encoder *httpbinding.Encoder) error {
@@ -3530,10 +3406,8 @@ func (m *awsRestjson1_serializeOpMalformedAcceptWithPa
 func (m *awsRestjson1_serializeOpMalformedAcceptWithPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3567,7 +3441,6 @@ func (m *awsRestjson1_serializeOpMalformedAcceptWithPa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedAcceptWithPayloadInput(v *MalformedAcceptWithPayloadInput, encoder *httpbinding.Encoder) error {
@@ -3588,10 +3461,8 @@ func (m *awsRestjson1_serializeOpMalformedBlob) Handle
 func (m *awsRestjson1_serializeOpMalformedBlob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3636,7 +3507,6 @@ func (m *awsRestjson1_serializeOpMalformedBlob) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedBlobInput(v *MalformedBlobInput, encoder *httpbinding.Encoder) error {
@@ -3669,10 +3539,8 @@ func (m *awsRestjson1_serializeOpMalformedBoolean) Han
 func (m *awsRestjson1_serializeOpMalformedBoolean) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3721,7 +3589,6 @@ func (m *awsRestjson1_serializeOpMalformedBoolean) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedBooleanInput(v *MalformedBooleanInput, encoder *httpbinding.Encoder) error {
@@ -3772,10 +3639,8 @@ func (m *awsRestjson1_serializeOpMalformedByte) Handle
 func (m *awsRestjson1_serializeOpMalformedByte) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3824,7 +3689,6 @@ func (m *awsRestjson1_serializeOpMalformedByte) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedByteInput(v *MalformedByteInput, encoder *httpbinding.Encoder) error {
@@ -3875,10 +3739,8 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 func (m *awsRestjson1_serializeOpMalformedContentTypeWithBody) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3923,7 +3785,6 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedContentTypeWithBodyInput(v *MalformedContentTypeWithBodyInput, encoder *httpbinding.Encoder) error {
@@ -3956,10 +3817,8 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 func (m *awsRestjson1_serializeOpMalformedContentTypeWithGenericString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4005,7 +3864,6 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedContentTypeWithGenericStringInput(v *MalformedContentTypeWithGenericStringInput, encoder *httpbinding.Encoder) error {
@@ -4026,10 +3884,8 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 func (m *awsRestjson1_serializeOpMalformedContentTypeWithoutBody) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4063,7 +3919,6 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedContentTypeWithoutBodyInput(v *MalformedContentTypeWithoutBodyInput, encoder *httpbinding.Encoder) error {
@@ -4084,10 +3939,8 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 func (m *awsRestjson1_serializeOpMalformedContentTypeWithoutBodyEmptyInput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4125,7 +3978,6 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedContentTypeWithoutBodyEmptyInputInput(v *MalformedContentTypeWithoutBodyEmptyInputInput, encoder *httpbinding.Encoder) error {
@@ -4151,10 +4003,8 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 func (m *awsRestjson1_serializeOpMalformedContentTypeWithPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4200,7 +4050,6 @@ func (m *awsRestjson1_serializeOpMalformedContentTypeW
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedContentTypeWithPayloadInput(v *MalformedContentTypeWithPayloadInput, encoder *httpbinding.Encoder) error {
@@ -4221,10 +4070,8 @@ func (m *awsRestjson1_serializeOpMalformedDouble) Hand
 func (m *awsRestjson1_serializeOpMalformedDouble) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4273,7 +4120,6 @@ func (m *awsRestjson1_serializeOpMalformedDouble) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedDoubleInput(v *MalformedDoubleInput, encoder *httpbinding.Encoder) error {
@@ -4337,10 +4183,8 @@ func (m *awsRestjson1_serializeOpMalformedFloat) Handl
 func (m *awsRestjson1_serializeOpMalformedFloat) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4389,7 +4233,6 @@ func (m *awsRestjson1_serializeOpMalformedFloat) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedFloatInput(v *MalformedFloatInput, encoder *httpbinding.Encoder) error {
@@ -4453,10 +4296,8 @@ func (m *awsRestjson1_serializeOpMalformedInteger) Han
 func (m *awsRestjson1_serializeOpMalformedInteger) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4505,7 +4346,6 @@ func (m *awsRestjson1_serializeOpMalformedInteger) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedIntegerInput(v *MalformedIntegerInput, encoder *httpbinding.Encoder) error {
@@ -4556,10 +4396,8 @@ func (m *awsRestjson1_serializeOpMalformedList) Handle
 func (m *awsRestjson1_serializeOpMalformedList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4604,7 +4442,6 @@ func (m *awsRestjson1_serializeOpMalformedList) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedListInput(v *MalformedListInput, encoder *httpbinding.Encoder) error {
@@ -4639,10 +4476,8 @@ func (m *awsRestjson1_serializeOpMalformedLong) Handle
 func (m *awsRestjson1_serializeOpMalformedLong) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4691,7 +4526,6 @@ func (m *awsRestjson1_serializeOpMalformedLong) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedLongInput(v *MalformedLongInput, encoder *httpbinding.Encoder) error {
@@ -4742,10 +4576,8 @@ func (m *awsRestjson1_serializeOpMalformedMap) HandleS
 func (m *awsRestjson1_serializeOpMalformedMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4790,7 +4622,6 @@ func (m *awsRestjson1_serializeOpMalformedMap) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedMapInput(v *MalformedMapInput, encoder *httpbinding.Encoder) error {
@@ -4825,10 +4656,8 @@ func (m *awsRestjson1_serializeOpMalformedRequestBody)
 func (m *awsRestjson1_serializeOpMalformedRequestBody) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4873,7 +4702,6 @@ func (m *awsRestjson1_serializeOpMalformedRequestBody)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedRequestBodyInput(v *MalformedRequestBodyInput, encoder *httpbinding.Encoder) error {
@@ -4924,10 +4752,8 @@ func (m *awsRestjson1_serializeOpMalformedShort) Handl
 func (m *awsRestjson1_serializeOpMalformedShort) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4976,7 +4802,6 @@ func (m *awsRestjson1_serializeOpMalformedShort) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedShortInput(v *MalformedShortInput, encoder *httpbinding.Encoder) error {
@@ -5027,10 +4852,8 @@ func (m *awsRestjson1_serializeOpMalformedString) Hand
 func (m *awsRestjson1_serializeOpMalformedString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5068,7 +4891,6 @@ func (m *awsRestjson1_serializeOpMalformedString) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedStringInput(v *MalformedStringInput, encoder *httpbinding.Encoder) error {
@@ -5096,10 +4918,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampBod
 func (m *awsRestjson1_serializeOpMalformedTimestampBodyDateTime) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5144,7 +4964,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampBod
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampBodyDateTimeInput(v *MalformedTimestampBodyDateTimeInput, encoder *httpbinding.Encoder) error {
@@ -5177,10 +4996,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampBod
 func (m *awsRestjson1_serializeOpMalformedTimestampBodyDefault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5225,7 +5042,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampBod
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampBodyDefaultInput(v *MalformedTimestampBodyDefaultInput, encoder *httpbinding.Encoder) error {
@@ -5258,10 +5074,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampBod
 func (m *awsRestjson1_serializeOpMalformedTimestampBodyHttpDate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5306,7 +5120,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampBod
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampBodyHttpDateInput(v *MalformedTimestampBodyHttpDateInput, encoder *httpbinding.Encoder) error {
@@ -5339,10 +5152,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampHea
 func (m *awsRestjson1_serializeOpMalformedTimestampHeaderDateTime) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5380,7 +5191,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampHea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampHeaderDateTimeInput(v *MalformedTimestampHeaderDateTimeInput, encoder *httpbinding.Encoder) error {
@@ -5406,10 +5216,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampHea
 func (m *awsRestjson1_serializeOpMalformedTimestampHeaderDefault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5447,7 +5255,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampHea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampHeaderDefaultInput(v *MalformedTimestampHeaderDefaultInput, encoder *httpbinding.Encoder) error {
@@ -5473,10 +5280,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampHea
 func (m *awsRestjson1_serializeOpMalformedTimestampHeaderEpoch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5514,7 +5319,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampHea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampHeaderEpochInput(v *MalformedTimestampHeaderEpochInput, encoder *httpbinding.Encoder) error {
@@ -5540,10 +5344,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampPat
 func (m *awsRestjson1_serializeOpMalformedTimestampPathDefault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5581,7 +5383,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampPat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampPathDefaultInput(v *MalformedTimestampPathDefaultInput, encoder *httpbinding.Encoder) error {
@@ -5611,10 +5412,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampPat
 func (m *awsRestjson1_serializeOpMalformedTimestampPathEpoch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5652,7 +5451,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampPat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampPathEpochInput(v *MalformedTimestampPathEpochInput, encoder *httpbinding.Encoder) error {
@@ -5682,10 +5480,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampPat
 func (m *awsRestjson1_serializeOpMalformedTimestampPathHttpDate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5723,7 +5519,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampPat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampPathHttpDateInput(v *MalformedTimestampPathHttpDateInput, encoder *httpbinding.Encoder) error {
@@ -5753,10 +5548,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampQue
 func (m *awsRestjson1_serializeOpMalformedTimestampQueryDefault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5794,7 +5587,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampQue
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampQueryDefaultInput(v *MalformedTimestampQueryDefaultInput, encoder *httpbinding.Encoder) error {
@@ -5819,10 +5611,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampQue
 func (m *awsRestjson1_serializeOpMalformedTimestampQueryEpoch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5860,7 +5650,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampQue
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampQueryEpochInput(v *MalformedTimestampQueryEpochInput, encoder *httpbinding.Encoder) error {
@@ -5885,10 +5674,8 @@ func (m *awsRestjson1_serializeOpMalformedTimestampQue
 func (m *awsRestjson1_serializeOpMalformedTimestampQueryHttpDate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5926,7 +5713,6 @@ func (m *awsRestjson1_serializeOpMalformedTimestampQue
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedTimestampQueryHttpDateInput(v *MalformedTimestampQueryHttpDateInput, encoder *httpbinding.Encoder) error {
@@ -5951,10 +5737,8 @@ func (m *awsRestjson1_serializeOpMalformedUnion) Handl
 func (m *awsRestjson1_serializeOpMalformedUnion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5999,7 +5783,6 @@ func (m *awsRestjson1_serializeOpMalformedUnion) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMalformedUnionInput(v *MalformedUnionInput, encoder *httpbinding.Encoder) error {
@@ -6034,10 +5817,8 @@ func (m *awsRestjson1_serializeOpMediaTypeHeader) Hand
 func (m *awsRestjson1_serializeOpMediaTypeHeader) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6075,7 +5856,6 @@ func (m *awsRestjson1_serializeOpMediaTypeHeader) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMediaTypeHeaderInput(v *MediaTypeHeaderInput, encoder *httpbinding.Encoder) error {
@@ -6103,10 +5883,8 @@ func (m *awsRestjson1_serializeOpNoInputAndNoOutput) H
 func (m *awsRestjson1_serializeOpNoInputAndNoOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6140,7 +5918,6 @@ func (m *awsRestjson1_serializeOpNoInputAndNoOutput) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsNoInputAndNoOutputInput(v *NoInputAndNoOutputInput, encoder *httpbinding.Encoder) error {
@@ -6161,10 +5938,8 @@ func (m *awsRestjson1_serializeOpNoInputAndOutput) Han
 func (m *awsRestjson1_serializeOpNoInputAndOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6198,7 +5973,6 @@ func (m *awsRestjson1_serializeOpNoInputAndOutput) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsNoInputAndOutputInput(v *NoInputAndOutputInput, encoder *httpbinding.Encoder) error {
@@ -6219,10 +5993,8 @@ func (m *awsRestjson1_serializeOpNullAndEmptyHeadersCl
 func (m *awsRestjson1_serializeOpNullAndEmptyHeadersClient) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6260,7 +6032,6 @@ func (m *awsRestjson1_serializeOpNullAndEmptyHeadersCl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsNullAndEmptyHeadersClientInput(v *NullAndEmptyHeadersClientInput, encoder *httpbinding.Encoder) error {
@@ -6308,10 +6079,8 @@ func (m *awsRestjson1_serializeOpNullAndEmptyHeadersSe
 func (m *awsRestjson1_serializeOpNullAndEmptyHeadersServer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6349,7 +6118,6 @@ func (m *awsRestjson1_serializeOpNullAndEmptyHeadersSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsNullAndEmptyHeadersServerInput(v *NullAndEmptyHeadersServerInput, encoder *httpbinding.Encoder) error {
@@ -6397,10 +6165,8 @@ func (m *awsRestjson1_serializeOpOmitsNullSerializesEm
 func (m *awsRestjson1_serializeOpOmitsNullSerializesEmptyString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6438,7 +6204,6 @@ func (m *awsRestjson1_serializeOpOmitsNullSerializesEm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsOmitsNullSerializesEmptyStringInput(v *OmitsNullSerializesEmptyStringInput, encoder *httpbinding.Encoder) error {
@@ -6467,10 +6232,8 @@ func (m *awsRestjson1_serializeOpOmitsSerializingEmpty
 func (m *awsRestjson1_serializeOpOmitsSerializingEmptyLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6508,7 +6271,6 @@ func (m *awsRestjson1_serializeOpOmitsSerializingEmpty
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsOmitsSerializingEmptyListsInput(v *OmitsSerializingEmptyListsInput, encoder *httpbinding.Encoder) error {
@@ -6571,10 +6333,8 @@ func (m *awsRestjson1_serializeOpOperationWithDefaults
 func (m *awsRestjson1_serializeOpOperationWithDefaults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6619,7 +6379,6 @@ func (m *awsRestjson1_serializeOpOperationWithDefaults
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsOperationWithDefaultsInput(v *OperationWithDefaultsInput, encoder *httpbinding.Encoder) error {
@@ -6671,10 +6430,8 @@ func (m *awsRestjson1_serializeOpOperationWithNestedSt
 func (m *awsRestjson1_serializeOpOperationWithNestedStructure) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6719,7 +6476,6 @@ func (m *awsRestjson1_serializeOpOperationWithNestedSt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsOperationWithNestedStructureInput(v *OperationWithNestedStructureInput, encoder *httpbinding.Encoder) error {
@@ -6754,10 +6510,8 @@ func (m *awsRestjson1_serializeOpPostPlayerAction) Han
 func (m *awsRestjson1_serializeOpPostPlayerAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6802,7 +6556,6 @@ func (m *awsRestjson1_serializeOpPostPlayerAction) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPostPlayerActionInput(v *PostPlayerActionInput, encoder *httpbinding.Encoder) error {
@@ -6837,10 +6590,8 @@ func (m *awsRestjson1_serializeOpPostUnionWithJsonName
 func (m *awsRestjson1_serializeOpPostUnionWithJsonName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6885,7 +6636,6 @@ func (m *awsRestjson1_serializeOpPostUnionWithJsonName
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPostUnionWithJsonNameInput(v *PostUnionWithJsonNameInput, encoder *httpbinding.Encoder) error {
@@ -6920,10 +6670,8 @@ func (m *awsRestjson1_serializeOpPutWithContentEncodin
 func (m *awsRestjson1_serializeOpPutWithContentEncoding) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6972,7 +6720,6 @@ func (m *awsRestjson1_serializeOpPutWithContentEncodin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutWithContentEncodingInput(v *PutWithContentEncodingInput, encoder *httpbinding.Encoder) error {
@@ -7010,10 +6757,8 @@ func (m *awsRestjson1_serializeOpQueryIdempotencyToken
 func (m *awsRestjson1_serializeOpQueryIdempotencyTokenAutoFill) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7051,7 +6796,6 @@ func (m *awsRestjson1_serializeOpQueryIdempotencyToken
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsQueryIdempotencyTokenAutoFillInput(v *QueryIdempotencyTokenAutoFillInput, encoder *httpbinding.Encoder) error {
@@ -7076,10 +6820,8 @@ func (m *awsRestjson1_serializeOpQueryParamsAsStringLi
 func (m *awsRestjson1_serializeOpQueryParamsAsStringListMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7117,7 +6859,6 @@ func (m *awsRestjson1_serializeOpQueryParamsAsStringLi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsQueryParamsAsStringListMapInput(v *QueryParamsAsStringListMapInput, encoder *httpbinding.Encoder) error {
@@ -7153,10 +6894,8 @@ func (m *awsRestjson1_serializeOpQueryPrecedence) Hand
 func (m *awsRestjson1_serializeOpQueryPrecedence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7194,7 +6933,6 @@ func (m *awsRestjson1_serializeOpQueryPrecedence) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsQueryPrecedenceInput(v *QueryPrecedenceInput, encoder *httpbinding.Encoder) error {
@@ -7228,10 +6966,8 @@ func (m *awsRestjson1_serializeOpRecursiveShapes) Hand
 func (m *awsRestjson1_serializeOpRecursiveShapes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7276,7 +7012,6 @@ func (m *awsRestjson1_serializeOpRecursiveShapes) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRecursiveShapesInput(v *RecursiveShapesInput, encoder *httpbinding.Encoder) error {
@@ -7311,10 +7046,8 @@ func (m *awsRestjson1_serializeOpResponseCodeHttpFallb
 func (m *awsRestjson1_serializeOpResponseCodeHttpFallback) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7348,7 +7081,6 @@ func (m *awsRestjson1_serializeOpResponseCodeHttpFallb
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResponseCodeHttpFallbackInput(v *ResponseCodeHttpFallbackInput, encoder *httpbinding.Encoder) error {
@@ -7369,10 +7101,8 @@ func (m *awsRestjson1_serializeOpResponseCodeRequired)
 func (m *awsRestjson1_serializeOpResponseCodeRequired) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7406,7 +7136,6 @@ func (m *awsRestjson1_serializeOpResponseCodeRequired)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResponseCodeRequiredInput(v *ResponseCodeRequiredInput, encoder *httpbinding.Encoder) error {
@@ -7427,10 +7156,8 @@ func (m *awsRestjson1_serializeOpSimpleScalarPropertie
 func (m *awsRestjson1_serializeOpSimpleScalarProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7479,7 +7206,6 @@ func (m *awsRestjson1_serializeOpSimpleScalarPropertie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSimpleScalarPropertiesInput(v *SimpleScalarPropertiesInput, encoder *httpbinding.Encoder) error {
@@ -7583,10 +7309,8 @@ func (m *awsRestjson1_serializeOpSparseJsonLists) Hand
 func (m *awsRestjson1_serializeOpSparseJsonLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7631,7 +7355,6 @@ func (m *awsRestjson1_serializeOpSparseJsonLists) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSparseJsonListsInput(v *SparseJsonListsInput, encoder *httpbinding.Encoder) error {
@@ -7673,10 +7396,8 @@ func (m *awsRestjson1_serializeOpSparseJsonMaps) Handl
 func (m *awsRestjson1_serializeOpSparseJsonMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7721,7 +7442,6 @@ func (m *awsRestjson1_serializeOpSparseJsonMaps) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSparseJsonMapsInput(v *SparseJsonMapsInput, encoder *httpbinding.Encoder) error {
@@ -7784,10 +7504,8 @@ func (m *awsRestjson1_serializeOpStreamingTraits) Hand
 func (m *awsRestjson1_serializeOpStreamingTraits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7837,7 +7555,6 @@ func (m *awsRestjson1_serializeOpStreamingTraits) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStreamingTraitsInput(v *StreamingTraitsInput, encoder *httpbinding.Encoder) error {
@@ -7863,10 +7580,8 @@ func (m *awsRestjson1_serializeOpStreamingTraitsRequir
 func (m *awsRestjson1_serializeOpStreamingTraitsRequireLength) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7916,7 +7631,6 @@ func (m *awsRestjson1_serializeOpStreamingTraitsRequir
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStreamingTraitsRequireLengthInput(v *StreamingTraitsRequireLengthInput, encoder *httpbinding.Encoder) error {
@@ -7942,10 +7656,8 @@ func (m *awsRestjson1_serializeOpStreamingTraitsWithMe
 func (m *awsRestjson1_serializeOpStreamingTraitsWithMediaType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7995,7 +7707,6 @@ func (m *awsRestjson1_serializeOpStreamingTraitsWithMe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStreamingTraitsWithMediaTypeInput(v *StreamingTraitsWithMediaTypeInput, encoder *httpbinding.Encoder) error {
@@ -8021,10 +7732,8 @@ func (m *awsRestjson1_serializeOpTestBodyStructure) Ha
 func (m *awsRestjson1_serializeOpTestBodyStructure) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8073,7 +7782,6 @@ func (m *awsRestjson1_serializeOpTestBodyStructure) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestBodyStructureInput(v *TestBodyStructureInput, encoder *httpbinding.Encoder) error {
@@ -8113,10 +7821,8 @@ func (m *awsRestjson1_serializeOpTestGetNoInputNoPaylo
 func (m *awsRestjson1_serializeOpTestGetNoInputNoPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8150,7 +7856,6 @@ func (m *awsRestjson1_serializeOpTestGetNoInputNoPaylo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestGetNoInputNoPayloadInput(v *TestGetNoInputNoPayloadInput, encoder *httpbinding.Encoder) error {
@@ -8171,10 +7876,8 @@ func (m *awsRestjson1_serializeOpTestGetNoPayload) Han
 func (m *awsRestjson1_serializeOpTestGetNoPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8212,7 +7915,6 @@ func (m *awsRestjson1_serializeOpTestGetNoPayload) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestGetNoPayloadInput(v *TestGetNoPayloadInput, encoder *httpbinding.Encoder) error {
@@ -8238,10 +7940,8 @@ func (m *awsRestjson1_serializeOpTestPayloadBlob) Hand
 func (m *awsRestjson1_serializeOpTestPayloadBlob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8291,7 +7991,6 @@ func (m *awsRestjson1_serializeOpTestPayloadBlob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestPayloadBlobInput(v *TestPayloadBlobInput, encoder *httpbinding.Encoder) error {
@@ -8317,10 +8016,8 @@ func (m *awsRestjson1_serializeOpTestPayloadStructure)
 func (m *awsRestjson1_serializeOpTestPayloadStructure) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8382,7 +8079,6 @@ func (m *awsRestjson1_serializeOpTestPayloadStructure)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestPayloadStructureInput(v *TestPayloadStructureInput, encoder *httpbinding.Encoder) error {
@@ -8408,10 +8104,8 @@ func (m *awsRestjson1_serializeOpTestPostNoInputNoPayl
 func (m *awsRestjson1_serializeOpTestPostNoInputNoPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8445,7 +8139,6 @@ func (m *awsRestjson1_serializeOpTestPostNoInputNoPayl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestPostNoInputNoPayloadInput(v *TestPostNoInputNoPayloadInput, encoder *httpbinding.Encoder) error {
@@ -8466,10 +8159,8 @@ func (m *awsRestjson1_serializeOpTestPostNoPayload) Ha
 func (m *awsRestjson1_serializeOpTestPostNoPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8507,7 +8198,6 @@ func (m *awsRestjson1_serializeOpTestPostNoPayload) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestPostNoPayloadInput(v *TestPostNoPayloadInput, encoder *httpbinding.Encoder) error {
@@ -8533,10 +8223,8 @@ func (m *awsRestjson1_serializeOpTimestampFormatHeader
 func (m *awsRestjson1_serializeOpTimestampFormatHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8574,7 +8262,6 @@ func (m *awsRestjson1_serializeOpTimestampFormatHeader
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTimestampFormatHeadersInput(v *TimestampFormatHeadersInput, encoder *httpbinding.Encoder) error {
@@ -8630,10 +8317,8 @@ func (m *awsRestjson1_serializeOpUnitInputAndOutput) H
 func (m *awsRestjson1_serializeOpUnitInputAndOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8667,7 +8352,6 @@ func (m *awsRestjson1_serializeOpUnitInputAndOutput) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUnitInputAndOutputInput(v *UnitInputAndOutputInput, encoder *httpbinding.Encoder) error {
