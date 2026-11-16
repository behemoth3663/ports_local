--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyxml "github.com/aws/smithy-go/encoding/xml"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 	"net/http"
@@ -31,10 +30,8 @@ func (m *awsRestxml_serializeOpAllQueryStringTypes) Ha
 func (m *awsRestxml_serializeOpAllQueryStringTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -72,7 +69,6 @@ func (m *awsRestxml_serializeOpAllQueryStringTypes) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsAllQueryStringTypesInput(v *AllQueryStringTypesInput, encoder *httpbinding.Encoder) error {
@@ -200,10 +196,8 @@ func (m *awsRestxml_serializeOpBodyWithXmlName) Handle
 func (m *awsRestxml_serializeOpBodyWithXmlName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -254,7 +248,6 @@ func (m *awsRestxml_serializeOpBodyWithXmlName) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsBodyWithXmlNameInput(v *BodyWithXmlNameInput, encoder *httpbinding.Encoder) error {
@@ -293,10 +286,8 @@ func (m *awsRestxml_serializeOpConstantAndVariableQuer
 func (m *awsRestxml_serializeOpConstantAndVariableQueryString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -334,7 +325,6 @@ func (m *awsRestxml_serializeOpConstantAndVariableQuer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsConstantAndVariableQueryStringInput(v *ConstantAndVariableQueryStringInput, encoder *httpbinding.Encoder) error {
@@ -363,10 +353,8 @@ func (m *awsRestxml_serializeOpConstantQueryString) Ha
 func (m *awsRestxml_serializeOpConstantQueryString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -404,7 +392,6 @@ func (m *awsRestxml_serializeOpConstantQueryString) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsConstantQueryStringInput(v *ConstantQueryStringInput, encoder *httpbinding.Encoder) error {
@@ -434,10 +421,8 @@ func (m *awsRestxml_serializeOpContentTypeParameters) 
 func (m *awsRestxml_serializeOpContentTypeParameters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -488,7 +473,6 @@ func (m *awsRestxml_serializeOpContentTypeParameters) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsContentTypeParametersInput(v *ContentTypeParametersInput, encoder *httpbinding.Encoder) error {
@@ -525,10 +509,8 @@ func (m *awsRestxml_serializeOpDatetimeOffsets) Handle
 func (m *awsRestxml_serializeOpDatetimeOffsets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -562,7 +544,6 @@ func (m *awsRestxml_serializeOpDatetimeOffsets) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDatetimeOffsetsInput(v *DatetimeOffsetsInput, encoder *httpbinding.Encoder) error {
@@ -583,10 +564,8 @@ func (m *awsRestxml_serializeOpEmptyInputAndEmptyOutpu
 func (m *awsRestxml_serializeOpEmptyInputAndEmptyOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -620,7 +599,6 @@ func (m *awsRestxml_serializeOpEmptyInputAndEmptyOutpu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsEmptyInputAndEmptyOutputInput(v *EmptyInputAndEmptyOutputInput, encoder *httpbinding.Encoder) error {
@@ -641,10 +619,8 @@ func (m *awsRestxml_serializeOpEndpointOperation) Hand
 func (m *awsRestxml_serializeOpEndpointOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -678,7 +654,6 @@ func (m *awsRestxml_serializeOpEndpointOperation) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsEndpointOperationInput(v *EndpointOperationInput, encoder *httpbinding.Encoder) error {
@@ -699,10 +674,8 @@ func (m *awsRestxml_serializeOpEndpointWithHostLabelHe
 func (m *awsRestxml_serializeOpEndpointWithHostLabelHeaderOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -740,7 +713,6 @@ func (m *awsRestxml_serializeOpEndpointWithHostLabelHe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsEndpointWithHostLabelHeaderOperationInput(v *EndpointWithHostLabelHeaderOperationInput, encoder *httpbinding.Encoder) error {
@@ -766,10 +738,8 @@ func (m *awsRestxml_serializeOpEndpointWithHostLabelOp
 func (m *awsRestxml_serializeOpEndpointWithHostLabelOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -820,7 +790,6 @@ func (m *awsRestxml_serializeOpEndpointWithHostLabelOp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsEndpointWithHostLabelOperationInput(v *EndpointWithHostLabelOperationInput, encoder *httpbinding.Encoder) error {
@@ -857,10 +826,8 @@ func (m *awsRestxml_serializeOpFlattenedXmlMap) Handle
 func (m *awsRestxml_serializeOpFlattenedXmlMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -911,7 +878,6 @@ func (m *awsRestxml_serializeOpFlattenedXmlMap) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsFlattenedXmlMapInput(v *FlattenedXmlMapInput, encoder *httpbinding.Encoder) error {
@@ -950,10 +916,8 @@ func (m *awsRestxml_serializeOpFlattenedXmlMapWithXmlN
 func (m *awsRestxml_serializeOpFlattenedXmlMapWithXmlName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1004,7 +968,6 @@ func (m *awsRestxml_serializeOpFlattenedXmlMapWithXmlN
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsFlattenedXmlMapWithXmlNameInput(v *FlattenedXmlMapWithXmlNameInput, encoder *httpbinding.Encoder) error {
@@ -1043,10 +1006,8 @@ func (m *awsRestxml_serializeOpFlattenedXmlMapWithXmlN
 func (m *awsRestxml_serializeOpFlattenedXmlMapWithXmlNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1080,7 +1041,6 @@ func (m *awsRestxml_serializeOpFlattenedXmlMapWithXmlN
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsFlattenedXmlMapWithXmlNamespaceInput(v *FlattenedXmlMapWithXmlNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -1101,10 +1061,8 @@ func (m *awsRestxml_serializeOpFractionalSeconds) Hand
 func (m *awsRestxml_serializeOpFractionalSeconds) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1138,7 +1096,6 @@ func (m *awsRestxml_serializeOpFractionalSeconds) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsFractionalSecondsInput(v *FractionalSecondsInput, encoder *httpbinding.Encoder) error {
@@ -1159,10 +1116,8 @@ func (m *awsRestxml_serializeOpGreetingWithErrors) Han
 func (m *awsRestxml_serializeOpGreetingWithErrors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1196,7 +1151,6 @@ func (m *awsRestxml_serializeOpGreetingWithErrors) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGreetingWithErrorsInput(v *GreetingWithErrorsInput, encoder *httpbinding.Encoder) error {
@@ -1217,10 +1171,8 @@ func (m *awsRestxml_serializeOpHttpEmptyPrefixHeaders)
 func (m *awsRestxml_serializeOpHttpEmptyPrefixHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1258,7 +1210,6 @@ func (m *awsRestxml_serializeOpHttpEmptyPrefixHeaders)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpEmptyPrefixHeadersInput(v *HttpEmptyPrefixHeadersInput, encoder *httpbinding.Encoder) error {
@@ -1291,10 +1242,8 @@ func (m *awsRestxml_serializeOpHttpEnumPayload) Handle
 func (m *awsRestxml_serializeOpHttpEnumPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1340,7 +1289,6 @@ func (m *awsRestxml_serializeOpHttpEnumPayload) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpEnumPayloadInput(v *HttpEnumPayloadInput, encoder *httpbinding.Encoder) error {
@@ -1361,10 +1309,8 @@ func (m *awsRestxml_serializeOpHttpPayloadTraits) Hand
 func (m *awsRestxml_serializeOpHttpPayloadTraits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1414,7 +1360,6 @@ func (m *awsRestxml_serializeOpHttpPayloadTraits) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadTraitsInput(v *HttpPayloadTraitsInput, encoder *httpbinding.Encoder) error {
@@ -1440,10 +1385,8 @@ func (m *awsRestxml_serializeOpHttpPayloadTraitsWithMe
 func (m *awsRestxml_serializeOpHttpPayloadTraitsWithMediaType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1493,7 +1436,6 @@ func (m *awsRestxml_serializeOpHttpPayloadTraitsWithMe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadTraitsWithMediaTypeInput(v *HttpPayloadTraitsWithMediaTypeInput, encoder *httpbinding.Encoder) error {
@@ -1519,10 +1461,8 @@ func (m *awsRestxml_serializeOpHttpPayloadWithMemberXm
 func (m *awsRestxml_serializeOpHttpPayloadWithMemberXmlName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1579,7 +1519,6 @@ func (m *awsRestxml_serializeOpHttpPayloadWithMemberXm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadWithMemberXmlNameInput(v *HttpPayloadWithMemberXmlNameInput, encoder *httpbinding.Encoder) error {
@@ -1600,10 +1539,8 @@ func (m *awsRestxml_serializeOpHttpPayloadWithStructur
 func (m *awsRestxml_serializeOpHttpPayloadWithStructure) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1660,7 +1597,6 @@ func (m *awsRestxml_serializeOpHttpPayloadWithStructur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadWithStructureInput(v *HttpPayloadWithStructureInput, encoder *httpbinding.Encoder) error {
@@ -1681,10 +1617,8 @@ func (m *awsRestxml_serializeOpHttpPayloadWithUnion) H
 func (m *awsRestxml_serializeOpHttpPayloadWithUnion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1741,7 +1675,6 @@ func (m *awsRestxml_serializeOpHttpPayloadWithUnion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadWithUnionInput(v *HttpPayloadWithUnionInput, encoder *httpbinding.Encoder) error {
@@ -1762,10 +1695,8 @@ func (m *awsRestxml_serializeOpHttpPayloadWithXmlName)
 func (m *awsRestxml_serializeOpHttpPayloadWithXmlName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1822,7 +1753,6 @@ func (m *awsRestxml_serializeOpHttpPayloadWithXmlName)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadWithXmlNameInput(v *HttpPayloadWithXmlNameInput, encoder *httpbinding.Encoder) error {
@@ -1843,10 +1773,8 @@ func (m *awsRestxml_serializeOpHttpPayloadWithXmlNames
 func (m *awsRestxml_serializeOpHttpPayloadWithXmlNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1904,7 +1832,6 @@ func (m *awsRestxml_serializeOpHttpPayloadWithXmlNames
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadWithXmlNamespaceInput(v *HttpPayloadWithXmlNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -1925,10 +1852,8 @@ func (m *awsRestxml_serializeOpHttpPayloadWithXmlNames
 func (m *awsRestxml_serializeOpHttpPayloadWithXmlNamespaceAndPrefix) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1986,7 +1911,6 @@ func (m *awsRestxml_serializeOpHttpPayloadWithXmlNames
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPayloadWithXmlNamespaceAndPrefixInput(v *HttpPayloadWithXmlNamespaceAndPrefixInput, encoder *httpbinding.Encoder) error {
@@ -2007,10 +1931,8 @@ func (m *awsRestxml_serializeOpHttpPrefixHeaders) Hand
 func (m *awsRestxml_serializeOpHttpPrefixHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2048,7 +1970,6 @@ func (m *awsRestxml_serializeOpHttpPrefixHeaders) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpPrefixHeadersInput(v *HttpPrefixHeadersInput, encoder *httpbinding.Encoder) error {
@@ -2081,10 +2002,8 @@ func (m *awsRestxml_serializeOpHttpRequestWithFloatLab
 func (m *awsRestxml_serializeOpHttpRequestWithFloatLabels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2122,7 +2041,6 @@ func (m *awsRestxml_serializeOpHttpRequestWithFloatLab
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpRequestWithFloatLabelsInput(v *HttpRequestWithFloatLabelsInput, encoder *httpbinding.Encoder) error {
@@ -2161,10 +2079,8 @@ func (m *awsRestxml_serializeOpHttpRequestWithGreedyLa
 func (m *awsRestxml_serializeOpHttpRequestWithGreedyLabelInPath) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2202,7 +2118,6 @@ func (m *awsRestxml_serializeOpHttpRequestWithGreedyLa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpRequestWithGreedyLabelInPathInput(v *HttpRequestWithGreedyLabelInPathInput, encoder *httpbinding.Encoder) error {
@@ -2241,10 +2156,8 @@ func (m *awsRestxml_serializeOpHttpRequestWithLabels) 
 func (m *awsRestxml_serializeOpHttpRequestWithLabels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2282,7 +2195,6 @@ func (m *awsRestxml_serializeOpHttpRequestWithLabels) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpRequestWithLabelsInput(v *HttpRequestWithLabelsInput, encoder *httpbinding.Encoder) error {
@@ -2375,10 +2287,8 @@ func (m *awsRestxml_serializeOpHttpRequestWithLabelsAn
 func (m *awsRestxml_serializeOpHttpRequestWithLabelsAndTimestampFormat) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2416,7 +2326,6 @@ func (m *awsRestxml_serializeOpHttpRequestWithLabelsAn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpRequestWithLabelsAndTimestampFormatInput(v *HttpRequestWithLabelsAndTimestampFormatInput, encoder *httpbinding.Encoder) error {
@@ -2500,10 +2409,8 @@ func (m *awsRestxml_serializeOpHttpResponseCode) Handl
 func (m *awsRestxml_serializeOpHttpResponseCode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2537,7 +2444,6 @@ func (m *awsRestxml_serializeOpHttpResponseCode) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpResponseCodeInput(v *HttpResponseCodeInput, encoder *httpbinding.Encoder) error {
@@ -2558,10 +2464,8 @@ func (m *awsRestxml_serializeOpHttpStringPayload) Hand
 func (m *awsRestxml_serializeOpHttpStringPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2607,7 +2511,6 @@ func (m *awsRestxml_serializeOpHttpStringPayload) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHttpStringPayloadInput(v *HttpStringPayloadInput, encoder *httpbinding.Encoder) error {
@@ -2628,10 +2531,8 @@ func (m *awsRestxml_serializeOpIgnoreQueryParamsInResp
 func (m *awsRestxml_serializeOpIgnoreQueryParamsInResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2665,7 +2566,6 @@ func (m *awsRestxml_serializeOpIgnoreQueryParamsInResp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsIgnoreQueryParamsInResponseInput(v *IgnoreQueryParamsInResponseInput, encoder *httpbinding.Encoder) error {
@@ -2686,10 +2586,8 @@ func (m *awsRestxml_serializeOpInputAndOutputWithHeade
 func (m *awsRestxml_serializeOpInputAndOutputWithHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2727,7 +2625,6 @@ func (m *awsRestxml_serializeOpInputAndOutputWithHeade
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsInputAndOutputWithHeadersInput(v *InputAndOutputWithHeadersInput, encoder *httpbinding.Encoder) error {
@@ -2885,10 +2782,8 @@ func (m *awsRestxml_serializeOpNestedXmlMaps) HandleSe
 func (m *awsRestxml_serializeOpNestedXmlMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2939,7 +2834,6 @@ func (m *awsRestxml_serializeOpNestedXmlMaps) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsNestedXmlMapsInput(v *NestedXmlMapsInput, encoder *httpbinding.Encoder) error {
@@ -2991,10 +2885,8 @@ func (m *awsRestxml_serializeOpNestedXmlMapWithXmlName
 func (m *awsRestxml_serializeOpNestedXmlMapWithXmlName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3045,7 +2937,6 @@ func (m *awsRestxml_serializeOpNestedXmlMapWithXmlName
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsNestedXmlMapWithXmlNameInput(v *NestedXmlMapWithXmlNameInput, encoder *httpbinding.Encoder) error {
@@ -3084,10 +2975,8 @@ func (m *awsRestxml_serializeOpNoInputAndNoOutput) Han
 func (m *awsRestxml_serializeOpNoInputAndNoOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3121,7 +3010,6 @@ func (m *awsRestxml_serializeOpNoInputAndNoOutput) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsNoInputAndNoOutputInput(v *NoInputAndNoOutputInput, encoder *httpbinding.Encoder) error {
@@ -3142,10 +3030,8 @@ func (m *awsRestxml_serializeOpNoInputAndOutput) Handl
 func (m *awsRestxml_serializeOpNoInputAndOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3179,7 +3065,6 @@ func (m *awsRestxml_serializeOpNoInputAndOutput) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsNoInputAndOutputInput(v *NoInputAndOutputInput, encoder *httpbinding.Encoder) error {
@@ -3200,10 +3085,8 @@ func (m *awsRestxml_serializeOpNullAndEmptyHeadersClie
 func (m *awsRestxml_serializeOpNullAndEmptyHeadersClient) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3241,7 +3124,6 @@ func (m *awsRestxml_serializeOpNullAndEmptyHeadersClie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsNullAndEmptyHeadersClientInput(v *NullAndEmptyHeadersClientInput, encoder *httpbinding.Encoder) error {
@@ -3289,10 +3171,8 @@ func (m *awsRestxml_serializeOpNullAndEmptyHeadersServ
 func (m *awsRestxml_serializeOpNullAndEmptyHeadersServer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3330,7 +3210,6 @@ func (m *awsRestxml_serializeOpNullAndEmptyHeadersServ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsNullAndEmptyHeadersServerInput(v *NullAndEmptyHeadersServerInput, encoder *httpbinding.Encoder) error {
@@ -3378,10 +3257,8 @@ func (m *awsRestxml_serializeOpOmitsNullSerializesEmpt
 func (m *awsRestxml_serializeOpOmitsNullSerializesEmptyString) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3419,7 +3296,6 @@ func (m *awsRestxml_serializeOpOmitsNullSerializesEmpt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsOmitsNullSerializesEmptyStringInput(v *OmitsNullSerializesEmptyStringInput, encoder *httpbinding.Encoder) error {
@@ -3448,10 +3324,8 @@ func (m *awsRestxml_serializeOpPutWithContentEncoding)
 func (m *awsRestxml_serializeOpPutWithContentEncoding) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3506,7 +3380,6 @@ func (m *awsRestxml_serializeOpPutWithContentEncoding)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutWithContentEncodingInput(v *PutWithContentEncodingInput, encoder *httpbinding.Encoder) error {
@@ -3548,10 +3421,8 @@ func (m *awsRestxml_serializeOpQueryIdempotencyTokenAu
 func (m *awsRestxml_serializeOpQueryIdempotencyTokenAutoFill) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3589,7 +3460,6 @@ func (m *awsRestxml_serializeOpQueryIdempotencyTokenAu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsQueryIdempotencyTokenAutoFillInput(v *QueryIdempotencyTokenAutoFillInput, encoder *httpbinding.Encoder) error {
@@ -3614,10 +3484,8 @@ func (m *awsRestxml_serializeOpQueryParamsAsStringList
 func (m *awsRestxml_serializeOpQueryParamsAsStringListMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3655,7 +3523,6 @@ func (m *awsRestxml_serializeOpQueryParamsAsStringList
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsQueryParamsAsStringListMapInput(v *QueryParamsAsStringListMapInput, encoder *httpbinding.Encoder) error {
@@ -3691,10 +3558,8 @@ func (m *awsRestxml_serializeOpQueryPrecedence) Handle
 func (m *awsRestxml_serializeOpQueryPrecedence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3732,7 +3597,6 @@ func (m *awsRestxml_serializeOpQueryPrecedence) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsQueryPrecedenceInput(v *QueryPrecedenceInput, encoder *httpbinding.Encoder) error {
@@ -3766,10 +3630,8 @@ func (m *awsRestxml_serializeOpRecursiveShapes) Handle
 func (m *awsRestxml_serializeOpRecursiveShapes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3820,7 +3682,6 @@ func (m *awsRestxml_serializeOpRecursiveShapes) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsRecursiveShapesInput(v *RecursiveShapesInput, encoder *httpbinding.Encoder) error {
@@ -3859,10 +3720,8 @@ func (m *awsRestxml_serializeOpSimpleScalarProperties)
 func (m *awsRestxml_serializeOpSimpleScalarProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3917,7 +3776,6 @@ func (m *awsRestxml_serializeOpSimpleScalarProperties)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsSimpleScalarPropertiesInput(v *SimpleScalarPropertiesInput, encoder *httpbinding.Encoder) error {
@@ -4073,10 +3931,8 @@ func (m *awsRestxml_serializeOpTimestampFormatHeaders)
 func (m *awsRestxml_serializeOpTimestampFormatHeaders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4114,7 +3970,6 @@ func (m *awsRestxml_serializeOpTimestampFormatHeaders)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTimestampFormatHeadersInput(v *TimestampFormatHeadersInput, encoder *httpbinding.Encoder) error {
@@ -4170,10 +4025,8 @@ func (m *awsRestxml_serializeOpXmlAttributes) HandleSe
 func (m *awsRestxml_serializeOpXmlAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4229,7 +4082,6 @@ func (m *awsRestxml_serializeOpXmlAttributes) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlAttributesInput(v *XmlAttributesInput, encoder *httpbinding.Encoder) error {
@@ -4266,10 +4118,8 @@ func (m *awsRestxml_serializeOpXmlAttributesOnPayload)
 func (m *awsRestxml_serializeOpXmlAttributesOnPayload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4331,7 +4181,6 @@ func (m *awsRestxml_serializeOpXmlAttributesOnPayload)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlAttributesOnPayloadInput(v *XmlAttributesOnPayloadInput, encoder *httpbinding.Encoder) error {
@@ -4352,10 +4201,8 @@ func (m *awsRestxml_serializeOpXmlBlobs) HandleSeriali
 func (m *awsRestxml_serializeOpXmlBlobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4406,7 +4253,6 @@ func (m *awsRestxml_serializeOpXmlBlobs) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlBlobsInput(v *XmlBlobsInput, encoder *httpbinding.Encoder) error {
@@ -4443,10 +4289,8 @@ func (m *awsRestxml_serializeOpXmlEmptyBlobs) HandleSe
 func (m *awsRestxml_serializeOpXmlEmptyBlobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4497,7 +4341,6 @@ func (m *awsRestxml_serializeOpXmlEmptyBlobs) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlEmptyBlobsInput(v *XmlEmptyBlobsInput, encoder *httpbinding.Encoder) error {
@@ -4534,10 +4377,8 @@ func (m *awsRestxml_serializeOpXmlEmptyLists) HandleSe
 func (m *awsRestxml_serializeOpXmlEmptyLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4588,7 +4429,6 @@ func (m *awsRestxml_serializeOpXmlEmptyLists) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlEmptyListsInput(v *XmlEmptyListsInput, encoder *httpbinding.Encoder) error {
@@ -4809,10 +4649,8 @@ func (m *awsRestxml_serializeOpXmlEmptyMaps) HandleSer
 func (m *awsRestxml_serializeOpXmlEmptyMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4863,7 +4701,6 @@ func (m *awsRestxml_serializeOpXmlEmptyMaps) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlEmptyMapsInput(v *XmlEmptyMapsInput, encoder *httpbinding.Encoder) error {
@@ -4902,10 +4739,8 @@ func (m *awsRestxml_serializeOpXmlEmptyStrings) Handle
 func (m *awsRestxml_serializeOpXmlEmptyStrings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4956,7 +4791,6 @@ func (m *awsRestxml_serializeOpXmlEmptyStrings) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlEmptyStringsInput(v *XmlEmptyStringsInput, encoder *httpbinding.Encoder) error {
@@ -4993,10 +4827,8 @@ func (m *awsRestxml_serializeOpXmlEnums) HandleSeriali
 func (m *awsRestxml_serializeOpXmlEnums) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5047,7 +4879,6 @@ func (m *awsRestxml_serializeOpXmlEnums) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlEnumsInput(v *XmlEnumsInput, encoder *httpbinding.Encoder) error {
@@ -5145,10 +4976,8 @@ func (m *awsRestxml_serializeOpXmlIntEnums) HandleSeri
 func (m *awsRestxml_serializeOpXmlIntEnums) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5199,7 +5028,6 @@ func (m *awsRestxml_serializeOpXmlIntEnums) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlIntEnumsInput(v *XmlIntEnumsInput, encoder *httpbinding.Encoder) error {
@@ -5297,10 +5125,8 @@ func (m *awsRestxml_serializeOpXmlLists) HandleSeriali
 func (m *awsRestxml_serializeOpXmlLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5351,7 +5177,6 @@ func (m *awsRestxml_serializeOpXmlLists) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlListsInput(v *XmlListsInput, encoder *httpbinding.Encoder) error {
@@ -5572,10 +5397,8 @@ func (m *awsRestxml_serializeOpXmlMaps) HandleSerializ
 func (m *awsRestxml_serializeOpXmlMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5626,7 +5449,6 @@ func (m *awsRestxml_serializeOpXmlMaps) HandleSerializ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlMapsInput(v *XmlMapsInput, encoder *httpbinding.Encoder) error {
@@ -5665,10 +5487,8 @@ func (m *awsRestxml_serializeOpXmlMapsXmlName) HandleS
 func (m *awsRestxml_serializeOpXmlMapsXmlName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5719,7 +5539,6 @@ func (m *awsRestxml_serializeOpXmlMapsXmlName) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlMapsXmlNameInput(v *XmlMapsXmlNameInput, encoder *httpbinding.Encoder) error {
@@ -5758,10 +5577,8 @@ func (m *awsRestxml_serializeOpXmlMapWithXmlNamespace)
 func (m *awsRestxml_serializeOpXmlMapWithXmlNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5812,7 +5629,6 @@ func (m *awsRestxml_serializeOpXmlMapWithXmlNamespace)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlMapWithXmlNamespaceInput(v *XmlMapWithXmlNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -5852,10 +5668,8 @@ func (m *awsRestxml_serializeOpXmlNamespaces) HandleSe
 func (m *awsRestxml_serializeOpXmlNamespaces) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5907,7 +5721,6 @@ func (m *awsRestxml_serializeOpXmlNamespaces) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlNamespacesInput(v *XmlNamespacesInput, encoder *httpbinding.Encoder) error {
@@ -5946,10 +5759,8 @@ func (m *awsRestxml_serializeOpXmlTimestamps) HandleSe
 func (m *awsRestxml_serializeOpXmlTimestamps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6000,7 +5811,6 @@ func (m *awsRestxml_serializeOpXmlTimestamps) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlTimestampsInput(v *XmlTimestampsInput, encoder *httpbinding.Encoder) error {
@@ -6103,10 +5913,8 @@ func (m *awsRestxml_serializeOpXmlUnions) HandleSerial
 func (m *awsRestxml_serializeOpXmlUnions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6157,7 +5965,6 @@ func (m *awsRestxml_serializeOpXmlUnions) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsXmlUnionsInput(v *XmlUnionsInput, encoder *httpbinding.Encoder) error {
