--- vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumbers
 func (m *awsRestjson1_serializeOpAssociatePhoneNumbersWithVoiceConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumbers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociatePhoneNumbersWithVoiceConnectorInput(v *AssociatePhoneNumbersWithVoiceConnectorInput, encoder *httpbinding.Encoder) error {
@@ -126,10 +122,8 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumbers
 func (m *awsRestjson1_serializeOpAssociatePhoneNumbersWithVoiceConnectorGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -178,7 +172,6 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumbers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociatePhoneNumbersWithVoiceConnectorGroupInput(v *AssociatePhoneNumbersWithVoiceConnectorGroupInput, encoder *httpbinding.Encoder) error {
@@ -227,10 +220,8 @@ func (m *awsRestjson1_serializeOpBatchDeletePhoneNumbe
 func (m *awsRestjson1_serializeOpBatchDeletePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -275,7 +266,6 @@ func (m *awsRestjson1_serializeOpBatchDeletePhoneNumbe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeletePhoneNumberInput(v *BatchDeletePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -310,10 +300,8 @@ func (m *awsRestjson1_serializeOpBatchUpdatePhoneNumbe
 func (m *awsRestjson1_serializeOpBatchUpdatePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -358,7 +346,6 @@ func (m *awsRestjson1_serializeOpBatchUpdatePhoneNumbe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdatePhoneNumberInput(v *BatchUpdatePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -393,10 +380,8 @@ func (m *awsRestjson1_serializeOpCreatePhoneNumberOrde
 func (m *awsRestjson1_serializeOpCreatePhoneNumberOrder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -441,7 +426,6 @@ func (m *awsRestjson1_serializeOpCreatePhoneNumberOrde
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePhoneNumberOrderInput(v *CreatePhoneNumberOrderInput, encoder *httpbinding.Encoder) error {
@@ -486,10 +470,8 @@ func (m *awsRestjson1_serializeOpCreateProxySession) H
 func (m *awsRestjson1_serializeOpCreateProxySession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -538,7 +520,6 @@ func (m *awsRestjson1_serializeOpCreateProxySession) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProxySessionInput(v *CreateProxySessionInput, encoder *httpbinding.Encoder) error {
@@ -616,10 +597,8 @@ func (m *awsRestjson1_serializeOpCreateSipMediaApplica
 func (m *awsRestjson1_serializeOpCreateSipMediaApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -664,7 +643,6 @@ func (m *awsRestjson1_serializeOpCreateSipMediaApplica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSipMediaApplicationInput(v *CreateSipMediaApplicationInput, encoder *httpbinding.Encoder) error {
@@ -716,10 +694,8 @@ func (m *awsRestjson1_serializeOpCreateSipMediaApplica
 func (m *awsRestjson1_serializeOpCreateSipMediaApplicationCall) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -768,7 +744,6 @@ func (m *awsRestjson1_serializeOpCreateSipMediaApplica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSipMediaApplicationCallInput(v *CreateSipMediaApplicationCallInput, encoder *httpbinding.Encoder) error {
@@ -829,10 +804,8 @@ func (m *awsRestjson1_serializeOpCreateSipRule) Handle
 func (m *awsRestjson1_serializeOpCreateSipRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -877,7 +850,6 @@ func (m *awsRestjson1_serializeOpCreateSipRule) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSipRuleInput(v *CreateSipRuleInput, encoder *httpbinding.Encoder) error {
@@ -932,10 +904,8 @@ func (m *awsRestjson1_serializeOpCreateVoiceConnector)
 func (m *awsRestjson1_serializeOpCreateVoiceConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -980,7 +950,6 @@ func (m *awsRestjson1_serializeOpCreateVoiceConnector)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVoiceConnectorInput(v *CreateVoiceConnectorInput, encoder *httpbinding.Encoder) error {
@@ -1035,10 +1004,8 @@ func (m *awsRestjson1_serializeOpCreateVoiceConnectorG
 func (m *awsRestjson1_serializeOpCreateVoiceConnectorGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1083,7 +1050,6 @@ func (m *awsRestjson1_serializeOpCreateVoiceConnectorG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVoiceConnectorGroupInput(v *CreateVoiceConnectorGroupInput, encoder *httpbinding.Encoder) error {
@@ -1123,10 +1089,8 @@ func (m *awsRestjson1_serializeOpCreateVoiceProfile) H
 func (m *awsRestjson1_serializeOpCreateVoiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1171,7 +1135,6 @@ func (m *awsRestjson1_serializeOpCreateVoiceProfile) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVoiceProfileInput(v *CreateVoiceProfileInput, encoder *httpbinding.Encoder) error {
@@ -1204,10 +1167,8 @@ func (m *awsRestjson1_serializeOpCreateVoiceProfileDom
 func (m *awsRestjson1_serializeOpCreateVoiceProfileDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1252,7 +1213,6 @@ func (m *awsRestjson1_serializeOpCreateVoiceProfileDom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVoiceProfileDomainInput(v *CreateVoiceProfileDomainInput, encoder *httpbinding.Encoder) error {
@@ -1309,10 +1269,8 @@ func (m *awsRestjson1_serializeOpDeletePhoneNumber) Ha
 func (m *awsRestjson1_serializeOpDeletePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1350,7 +1308,6 @@ func (m *awsRestjson1_serializeOpDeletePhoneNumber) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePhoneNumberInput(v *DeletePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -1380,10 +1337,8 @@ func (m *awsRestjson1_serializeOpDeleteProxySession) H
 func (m *awsRestjson1_serializeOpDeleteProxySession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1421,7 +1376,6 @@ func (m *awsRestjson1_serializeOpDeleteProxySession) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProxySessionInput(v *DeleteProxySessionInput, encoder *httpbinding.Encoder) error {
@@ -1460,10 +1414,8 @@ func (m *awsRestjson1_serializeOpDeleteSipMediaApplica
 func (m *awsRestjson1_serializeOpDeleteSipMediaApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1501,7 +1453,6 @@ func (m *awsRestjson1_serializeOpDeleteSipMediaApplica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSipMediaApplicationInput(v *DeleteSipMediaApplicationInput, encoder *httpbinding.Encoder) error {
@@ -1531,10 +1482,8 @@ func (m *awsRestjson1_serializeOpDeleteSipRule) Handle
 func (m *awsRestjson1_serializeOpDeleteSipRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1572,7 +1521,6 @@ func (m *awsRestjson1_serializeOpDeleteSipRule) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSipRuleInput(v *DeleteSipRuleInput, encoder *httpbinding.Encoder) error {
@@ -1602,10 +1550,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnector)
 func (m *awsRestjson1_serializeOpDeleteVoiceConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1643,7 +1589,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnector)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorInput(v *DeleteVoiceConnectorInput, encoder *httpbinding.Encoder) error {
@@ -1673,10 +1618,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorE
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorEmergencyCallingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1714,7 +1657,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorE
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorEmergencyCallingConfigurationInput(v *DeleteVoiceConnectorEmergencyCallingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1744,10 +1686,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorE
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorExternalSystemsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1785,7 +1725,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorE
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorExternalSystemsConfigurationInput(v *DeleteVoiceConnectorExternalSystemsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1815,10 +1754,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorG
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1856,7 +1793,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorGroupInput(v *DeleteVoiceConnectorGroupInput, encoder *httpbinding.Encoder) error {
@@ -1886,10 +1822,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorO
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorOrigination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1927,7 +1861,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorOriginationInput(v *DeleteVoiceConnectorOriginationInput, encoder *httpbinding.Encoder) error {
@@ -1957,10 +1890,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorP
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorProxy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1998,7 +1929,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorProxyInput(v *DeleteVoiceConnectorProxyInput, encoder *httpbinding.Encoder) error {
@@ -2028,10 +1958,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorS
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorStreamingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2069,7 +1997,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorStreamingConfigurationInput(v *DeleteVoiceConnectorStreamingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2099,10 +2026,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorT
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorTermination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2140,7 +2065,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorTerminationInput(v *DeleteVoiceConnectorTerminationInput, encoder *httpbinding.Encoder) error {
@@ -2170,10 +2094,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorT
 func (m *awsRestjson1_serializeOpDeleteVoiceConnectorTerminationCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2222,7 +2144,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceConnectorT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceConnectorTerminationCredentialsInput(v *DeleteVoiceConnectorTerminationCredentialsInput, encoder *httpbinding.Encoder) error {
@@ -2266,10 +2187,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceProfile) H
 func (m *awsRestjson1_serializeOpDeleteVoiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2307,7 +2226,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceProfile) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceProfileInput(v *DeleteVoiceProfileInput, encoder *httpbinding.Encoder) error {
@@ -2337,10 +2255,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceProfileDom
 func (m *awsRestjson1_serializeOpDeleteVoiceProfileDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2378,7 +2294,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceProfileDom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceProfileDomainInput(v *DeleteVoiceProfileDomainInput, encoder *httpbinding.Encoder) error {
@@ -2408,10 +2323,8 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 func (m *awsRestjson1_serializeOpDisassociatePhoneNumbersFromVoiceConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2460,7 +2373,6 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociatePhoneNumbersFromVoiceConnectorInput(v *DisassociatePhoneNumbersFromVoiceConnectorInput, encoder *httpbinding.Encoder) error {
@@ -2504,10 +2416,8 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 func (m *awsRestjson1_serializeOpDisassociatePhoneNumbersFromVoiceConnectorGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2556,7 +2466,6 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociatePhoneNumbersFromVoiceConnectorGroupInput(v *DisassociatePhoneNumbersFromVoiceConnectorGroupInput, encoder *httpbinding.Encoder) error {
@@ -2600,10 +2509,8 @@ func (m *awsRestjson1_serializeOpGetGlobalSettings) Ha
 func (m *awsRestjson1_serializeOpGetGlobalSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2637,7 +2544,6 @@ func (m *awsRestjson1_serializeOpGetGlobalSettings) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGlobalSettingsInput(v *GetGlobalSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2658,10 +2564,8 @@ func (m *awsRestjson1_serializeOpGetPhoneNumber) Handl
 func (m *awsRestjson1_serializeOpGetPhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2699,7 +2603,6 @@ func (m *awsRestjson1_serializeOpGetPhoneNumber) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPhoneNumberInput(v *GetPhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -2729,10 +2632,8 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberOrder) 
 func (m *awsRestjson1_serializeOpGetPhoneNumberOrder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2770,7 +2671,6 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberOrder) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPhoneNumberOrderInput(v *GetPhoneNumberOrderInput, encoder *httpbinding.Encoder) error {
@@ -2800,10 +2700,8 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberSetting
 func (m *awsRestjson1_serializeOpGetPhoneNumberSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2837,7 +2735,6 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPhoneNumberSettingsInput(v *GetPhoneNumberSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2858,10 +2755,8 @@ func (m *awsRestjson1_serializeOpGetProxySession) Hand
 func (m *awsRestjson1_serializeOpGetProxySession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2899,7 +2794,6 @@ func (m *awsRestjson1_serializeOpGetProxySession) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProxySessionInput(v *GetProxySessionInput, encoder *httpbinding.Encoder) error {
@@ -2938,10 +2832,8 @@ func (m *awsRestjson1_serializeOpGetSipMediaApplicatio
 func (m *awsRestjson1_serializeOpGetSipMediaApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2979,7 +2871,6 @@ func (m *awsRestjson1_serializeOpGetSipMediaApplicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSipMediaApplicationInput(v *GetSipMediaApplicationInput, encoder *httpbinding.Encoder) error {
@@ -3009,10 +2900,8 @@ func (m *awsRestjson1_serializeOpGetSipMediaApplicatio
 func (m *awsRestjson1_serializeOpGetSipMediaApplicationAlexaSkillConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3050,7 +2939,6 @@ func (m *awsRestjson1_serializeOpGetSipMediaApplicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSipMediaApplicationAlexaSkillConfigurationInput(v *GetSipMediaApplicationAlexaSkillConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3080,10 +2968,8 @@ func (m *awsRestjson1_serializeOpGetSipMediaApplicatio
 func (m *awsRestjson1_serializeOpGetSipMediaApplicationLoggingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3121,7 +3007,6 @@ func (m *awsRestjson1_serializeOpGetSipMediaApplicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSipMediaApplicationLoggingConfigurationInput(v *GetSipMediaApplicationLoggingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3151,10 +3036,8 @@ func (m *awsRestjson1_serializeOpGetSipRule) HandleSer
 func (m *awsRestjson1_serializeOpGetSipRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3192,7 +3075,6 @@ func (m *awsRestjson1_serializeOpGetSipRule) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSipRuleInput(v *GetSipRuleInput, encoder *httpbinding.Encoder) error {
@@ -3222,10 +3104,8 @@ func (m *awsRestjson1_serializeOpGetSpeakerSearchTask)
 func (m *awsRestjson1_serializeOpGetSpeakerSearchTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3263,7 +3143,6 @@ func (m *awsRestjson1_serializeOpGetSpeakerSearchTask)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSpeakerSearchTaskInput(v *GetSpeakerSearchTaskInput, encoder *httpbinding.Encoder) error {
@@ -3302,10 +3181,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnector) Ha
 func (m *awsRestjson1_serializeOpGetVoiceConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3343,7 +3220,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnector) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorInput(v *GetVoiceConnectorInput, encoder *httpbinding.Encoder) error {
@@ -3373,10 +3249,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorEmer
 func (m *awsRestjson1_serializeOpGetVoiceConnectorEmergencyCallingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3414,7 +3288,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorEmer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorEmergencyCallingConfigurationInput(v *GetVoiceConnectorEmergencyCallingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3444,10 +3317,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorExte
 func (m *awsRestjson1_serializeOpGetVoiceConnectorExternalSystemsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3485,7 +3356,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorExte
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorExternalSystemsConfigurationInput(v *GetVoiceConnectorExternalSystemsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3515,10 +3385,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorGrou
 func (m *awsRestjson1_serializeOpGetVoiceConnectorGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3556,7 +3424,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorGroupInput(v *GetVoiceConnectorGroupInput, encoder *httpbinding.Encoder) error {
@@ -3586,10 +3453,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorLogg
 func (m *awsRestjson1_serializeOpGetVoiceConnectorLoggingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3627,7 +3492,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorLogg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorLoggingConfigurationInput(v *GetVoiceConnectorLoggingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3657,10 +3521,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorOrig
 func (m *awsRestjson1_serializeOpGetVoiceConnectorOrigination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3698,7 +3560,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorOrig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorOriginationInput(v *GetVoiceConnectorOriginationInput, encoder *httpbinding.Encoder) error {
@@ -3728,10 +3589,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorProx
 func (m *awsRestjson1_serializeOpGetVoiceConnectorProxy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3769,7 +3628,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorProx
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorProxyInput(v *GetVoiceConnectorProxyInput, encoder *httpbinding.Encoder) error {
@@ -3799,10 +3657,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorStre
 func (m *awsRestjson1_serializeOpGetVoiceConnectorStreamingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3840,7 +3696,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorStre
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorStreamingConfigurationInput(v *GetVoiceConnectorStreamingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3870,10 +3725,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorTerm
 func (m *awsRestjson1_serializeOpGetVoiceConnectorTermination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3911,7 +3764,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorTerm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorTerminationInput(v *GetVoiceConnectorTerminationInput, encoder *httpbinding.Encoder) error {
@@ -3941,10 +3793,8 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorTerm
 func (m *awsRestjson1_serializeOpGetVoiceConnectorTerminationHealth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3982,7 +3832,6 @@ func (m *awsRestjson1_serializeOpGetVoiceConnectorTerm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceConnectorTerminationHealthInput(v *GetVoiceConnectorTerminationHealthInput, encoder *httpbinding.Encoder) error {
@@ -4012,10 +3861,8 @@ func (m *awsRestjson1_serializeOpGetVoiceProfile) Hand
 func (m *awsRestjson1_serializeOpGetVoiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4053,7 +3900,6 @@ func (m *awsRestjson1_serializeOpGetVoiceProfile) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceProfileInput(v *GetVoiceProfileInput, encoder *httpbinding.Encoder) error {
@@ -4083,10 +3929,8 @@ func (m *awsRestjson1_serializeOpGetVoiceProfileDomain
 func (m *awsRestjson1_serializeOpGetVoiceProfileDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4124,7 +3968,6 @@ func (m *awsRestjson1_serializeOpGetVoiceProfileDomain
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceProfileDomainInput(v *GetVoiceProfileDomainInput, encoder *httpbinding.Encoder) error {
@@ -4154,10 +3997,8 @@ func (m *awsRestjson1_serializeOpGetVoiceToneAnalysisT
 func (m *awsRestjson1_serializeOpGetVoiceToneAnalysisTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4195,7 +4036,6 @@ func (m *awsRestjson1_serializeOpGetVoiceToneAnalysisT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceToneAnalysisTaskInput(v *GetVoiceToneAnalysisTaskInput, encoder *httpbinding.Encoder) error {
@@ -4238,10 +4078,8 @@ func (m *awsRestjson1_serializeOpListAvailableVoiceCon
 func (m *awsRestjson1_serializeOpListAvailableVoiceConnectorRegions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4275,7 +4113,6 @@ func (m *awsRestjson1_serializeOpListAvailableVoiceCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAvailableVoiceConnectorRegionsInput(v *ListAvailableVoiceConnectorRegionsInput, encoder *httpbinding.Encoder) error {
@@ -4296,10 +4133,8 @@ func (m *awsRestjson1_serializeOpListPhoneNumberOrders
 func (m *awsRestjson1_serializeOpListPhoneNumberOrders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4337,7 +4172,6 @@ func (m *awsRestjson1_serializeOpListPhoneNumberOrders
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPhoneNumberOrdersInput(v *ListPhoneNumberOrdersInput, encoder *httpbinding.Encoder) error {
@@ -4366,10 +4200,8 @@ func (m *awsRestjson1_serializeOpListPhoneNumbers) Han
 func (m *awsRestjson1_serializeOpListPhoneNumbers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4407,7 +4239,6 @@ func (m *awsRestjson1_serializeOpListPhoneNumbers) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPhoneNumbersInput(v *ListPhoneNumbersInput, encoder *httpbinding.Encoder) error {
@@ -4452,10 +4283,8 @@ func (m *awsRestjson1_serializeOpListProxySessions) Ha
 func (m *awsRestjson1_serializeOpListProxySessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4493,7 +4322,6 @@ func (m *awsRestjson1_serializeOpListProxySessions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProxySessionsInput(v *ListProxySessionsInput, encoder *httpbinding.Encoder) error {
@@ -4535,10 +4363,8 @@ func (m *awsRestjson1_serializeOpListSipMediaApplicati
 func (m *awsRestjson1_serializeOpListSipMediaApplications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4576,7 +4402,6 @@ func (m *awsRestjson1_serializeOpListSipMediaApplicati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSipMediaApplicationsInput(v *ListSipMediaApplicationsInput, encoder *httpbinding.Encoder) error {
@@ -4605,10 +4430,8 @@ func (m *awsRestjson1_serializeOpListSipRules) HandleS
 func (m *awsRestjson1_serializeOpListSipRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4646,7 +4469,6 @@ func (m *awsRestjson1_serializeOpListSipRules) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSipRulesInput(v *ListSipRulesInput, encoder *httpbinding.Encoder) error {
@@ -4679,10 +4501,8 @@ func (m *awsRestjson1_serializeOpListSupportedPhoneNum
 func (m *awsRestjson1_serializeOpListSupportedPhoneNumberCountries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4720,7 +4540,6 @@ func (m *awsRestjson1_serializeOpListSupportedPhoneNum
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSupportedPhoneNumberCountriesInput(v *ListSupportedPhoneNumberCountriesInput, encoder *httpbinding.Encoder) error {
@@ -4745,10 +4564,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4786,7 +4603,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4811,10 +4627,8 @@ func (m *awsRestjson1_serializeOpListVoiceConnectorGro
 func (m *awsRestjson1_serializeOpListVoiceConnectorGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4852,7 +4666,6 @@ func (m *awsRestjson1_serializeOpListVoiceConnectorGro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVoiceConnectorGroupsInput(v *ListVoiceConnectorGroupsInput, encoder *httpbinding.Encoder) error {
@@ -4881,10 +4694,8 @@ func (m *awsRestjson1_serializeOpListVoiceConnectors) 
 func (m *awsRestjson1_serializeOpListVoiceConnectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4922,7 +4733,6 @@ func (m *awsRestjson1_serializeOpListVoiceConnectors) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVoiceConnectorsInput(v *ListVoiceConnectorsInput, encoder *httpbinding.Encoder) error {
@@ -4951,10 +4761,8 @@ func (m *awsRestjson1_serializeOpListVoiceConnectorTer
 func (m *awsRestjson1_serializeOpListVoiceConnectorTerminationCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4992,7 +4800,6 @@ func (m *awsRestjson1_serializeOpListVoiceConnectorTer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVoiceConnectorTerminationCredentialsInput(v *ListVoiceConnectorTerminationCredentialsInput, encoder *httpbinding.Encoder) error {
@@ -5022,10 +4829,8 @@ func (m *awsRestjson1_serializeOpListVoiceProfileDomai
 func (m *awsRestjson1_serializeOpListVoiceProfileDomains) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5063,7 +4868,6 @@ func (m *awsRestjson1_serializeOpListVoiceProfileDomai
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVoiceProfileDomainsInput(v *ListVoiceProfileDomainsInput, encoder *httpbinding.Encoder) error {
@@ -5092,10 +4896,8 @@ func (m *awsRestjson1_serializeOpListVoiceProfiles) Ha
 func (m *awsRestjson1_serializeOpListVoiceProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5133,7 +4935,6 @@ func (m *awsRestjson1_serializeOpListVoiceProfiles) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVoiceProfilesInput(v *ListVoiceProfilesInput, encoder *httpbinding.Encoder) error {
@@ -5166,10 +4967,8 @@ func (m *awsRestjson1_serializeOpPutSipMediaApplicatio
 func (m *awsRestjson1_serializeOpPutSipMediaApplicationAlexaSkillConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5218,7 +5017,6 @@ func (m *awsRestjson1_serializeOpPutSipMediaApplicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutSipMediaApplicationAlexaSkillConfigurationInput(v *PutSipMediaApplicationAlexaSkillConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5262,10 +5060,8 @@ func (m *awsRestjson1_serializeOpPutSipMediaApplicatio
 func (m *awsRestjson1_serializeOpPutSipMediaApplicationLoggingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5314,7 +5110,6 @@ func (m *awsRestjson1_serializeOpPutSipMediaApplicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutSipMediaApplicationLoggingConfigurationInput(v *PutSipMediaApplicationLoggingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5358,10 +5153,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorEmer
 func (m *awsRestjson1_serializeOpPutVoiceConnectorEmergencyCallingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5410,7 +5203,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorEmer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorEmergencyCallingConfigurationInput(v *PutVoiceConnectorEmergencyCallingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5454,10 +5246,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorExte
 func (m *awsRestjson1_serializeOpPutVoiceConnectorExternalSystemsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5506,7 +5296,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorExte
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorExternalSystemsConfigurationInput(v *PutVoiceConnectorExternalSystemsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5557,10 +5346,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorLogg
 func (m *awsRestjson1_serializeOpPutVoiceConnectorLoggingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5609,7 +5396,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorLogg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorLoggingConfigurationInput(v *PutVoiceConnectorLoggingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5653,10 +5439,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorOrig
 func (m *awsRestjson1_serializeOpPutVoiceConnectorOrigination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5705,7 +5489,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorOrig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorOriginationInput(v *PutVoiceConnectorOriginationInput, encoder *httpbinding.Encoder) error {
@@ -5749,10 +5532,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorProx
 func (m *awsRestjson1_serializeOpPutVoiceConnectorProxy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5801,7 +5582,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorProx
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorProxyInput(v *PutVoiceConnectorProxyInput, encoder *httpbinding.Encoder) error {
@@ -5860,10 +5640,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorStre
 func (m *awsRestjson1_serializeOpPutVoiceConnectorStreamingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5912,7 +5690,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorStre
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorStreamingConfigurationInput(v *PutVoiceConnectorStreamingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5956,10 +5733,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorTerm
 func (m *awsRestjson1_serializeOpPutVoiceConnectorTermination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6008,7 +5783,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorTerm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorTerminationInput(v *PutVoiceConnectorTerminationInput, encoder *httpbinding.Encoder) error {
@@ -6052,10 +5826,8 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorTerm
 func (m *awsRestjson1_serializeOpPutVoiceConnectorTerminationCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6104,7 +5876,6 @@ func (m *awsRestjson1_serializeOpPutVoiceConnectorTerm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVoiceConnectorTerminationCredentialsInput(v *PutVoiceConnectorTerminationCredentialsInput, encoder *httpbinding.Encoder) error {
@@ -6148,10 +5919,8 @@ func (m *awsRestjson1_serializeOpRestorePhoneNumber) H
 func (m *awsRestjson1_serializeOpRestorePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6189,7 +5958,6 @@ func (m *awsRestjson1_serializeOpRestorePhoneNumber) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRestorePhoneNumberInput(v *RestorePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -6219,10 +5987,8 @@ func (m *awsRestjson1_serializeOpSearchAvailablePhoneN
 func (m *awsRestjson1_serializeOpSearchAvailablePhoneNumbers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6260,7 +6026,6 @@ func (m *awsRestjson1_serializeOpSearchAvailablePhoneN
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchAvailablePhoneNumbersInput(v *SearchAvailablePhoneNumbersInput, encoder *httpbinding.Encoder) error {
@@ -6313,10 +6078,8 @@ func (m *awsRestjson1_serializeOpStartSpeakerSearchTas
 func (m *awsRestjson1_serializeOpStartSpeakerSearchTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6365,7 +6128,6 @@ func (m *awsRestjson1_serializeOpStartSpeakerSearchTas
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartSpeakerSearchTaskInput(v *StartSpeakerSearchTaskInput, encoder *httpbinding.Encoder) error {
@@ -6422,10 +6184,8 @@ func (m *awsRestjson1_serializeOpStartVoiceToneAnalysi
 func (m *awsRestjson1_serializeOpStartVoiceToneAnalysisTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6474,7 +6234,6 @@ func (m *awsRestjson1_serializeOpStartVoiceToneAnalysi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartVoiceToneAnalysisTaskInput(v *StartVoiceToneAnalysisTaskInput, encoder *httpbinding.Encoder) error {
@@ -6526,10 +6285,8 @@ func (m *awsRestjson1_serializeOpStopSpeakerSearchTask
 func (m *awsRestjson1_serializeOpStopSpeakerSearchTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6567,7 +6324,6 @@ func (m *awsRestjson1_serializeOpStopSpeakerSearchTask
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopSpeakerSearchTaskInput(v *StopSpeakerSearchTaskInput, encoder *httpbinding.Encoder) error {
@@ -6606,10 +6362,8 @@ func (m *awsRestjson1_serializeOpStopVoiceToneAnalysis
 func (m *awsRestjson1_serializeOpStopVoiceToneAnalysisTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6647,7 +6401,6 @@ func (m *awsRestjson1_serializeOpStopVoiceToneAnalysis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopVoiceToneAnalysisTaskInput(v *StopVoiceToneAnalysisTaskInput, encoder *httpbinding.Encoder) error {
@@ -6686,10 +6439,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6734,7 +6485,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6774,10 +6524,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6822,7 +6570,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6862,10 +6609,8 @@ func (m *awsRestjson1_serializeOpUpdateGlobalSettings)
 func (m *awsRestjson1_serializeOpUpdateGlobalSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6910,7 +6655,6 @@ func (m *awsRestjson1_serializeOpUpdateGlobalSettings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGlobalSettingsInput(v *UpdateGlobalSettingsInput, encoder *httpbinding.Encoder) error {
@@ -6945,10 +6689,8 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumber) Ha
 func (m *awsRestjson1_serializeOpUpdatePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6997,7 +6739,6 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumber) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePhoneNumberInput(v *UpdatePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -7049,10 +6790,8 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumberSett
 func (m *awsRestjson1_serializeOpUpdatePhoneNumberSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7097,7 +6836,6 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumberSett
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePhoneNumberSettingsInput(v *UpdatePhoneNumberSettingsInput, encoder *httpbinding.Encoder) error {
@@ -7130,10 +6868,8 @@ func (m *awsRestjson1_serializeOpUpdateProxySession) H
 func (m *awsRestjson1_serializeOpUpdateProxySession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7182,7 +6918,6 @@ func (m *awsRestjson1_serializeOpUpdateProxySession) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProxySessionInput(v *UpdateProxySessionInput, encoder *httpbinding.Encoder) error {
@@ -7240,10 +6975,8 @@ func (m *awsRestjson1_serializeOpUpdateSipMediaApplica
 func (m *awsRestjson1_serializeOpUpdateSipMediaApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7292,7 +7025,6 @@ func (m *awsRestjson1_serializeOpUpdateSipMediaApplica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSipMediaApplicationInput(v *UpdateSipMediaApplicationInput, encoder *httpbinding.Encoder) error {
@@ -7341,10 +7073,8 @@ func (m *awsRestjson1_serializeOpUpdateSipMediaApplica
 func (m *awsRestjson1_serializeOpUpdateSipMediaApplicationCall) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7393,7 +7123,6 @@ func (m *awsRestjson1_serializeOpUpdateSipMediaApplica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSipMediaApplicationCallInput(v *UpdateSipMediaApplicationCallInput, encoder *httpbinding.Encoder) error {
@@ -7446,10 +7175,8 @@ func (m *awsRestjson1_serializeOpUpdateSipRule) Handle
 func (m *awsRestjson1_serializeOpUpdateSipRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7498,7 +7225,6 @@ func (m *awsRestjson1_serializeOpUpdateSipRule) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSipRuleInput(v *UpdateSipRuleInput, encoder *httpbinding.Encoder) error {
@@ -7552,10 +7278,8 @@ func (m *awsRestjson1_serializeOpUpdateVoiceConnector)
 func (m *awsRestjson1_serializeOpUpdateVoiceConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7604,7 +7328,6 @@ func (m *awsRestjson1_serializeOpUpdateVoiceConnector)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVoiceConnectorInput(v *UpdateVoiceConnectorInput, encoder *httpbinding.Encoder) error {
@@ -7651,10 +7374,8 @@ func (m *awsRestjson1_serializeOpUpdateVoiceConnectorG
 func (m *awsRestjson1_serializeOpUpdateVoiceConnectorGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7703,7 +7424,6 @@ func (m *awsRestjson1_serializeOpUpdateVoiceConnectorG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVoiceConnectorGroupInput(v *UpdateVoiceConnectorGroupInput, encoder *httpbinding.Encoder) error {
@@ -7752,10 +7472,8 @@ func (m *awsRestjson1_serializeOpUpdateVoiceProfile) H
 func (m *awsRestjson1_serializeOpUpdateVoiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7804,7 +7522,6 @@ func (m *awsRestjson1_serializeOpUpdateVoiceProfile) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVoiceProfileInput(v *UpdateVoiceProfileInput, encoder *httpbinding.Encoder) error {
@@ -7846,10 +7563,8 @@ func (m *awsRestjson1_serializeOpUpdateVoiceProfileDom
 func (m *awsRestjson1_serializeOpUpdateVoiceProfileDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7898,7 +7613,6 @@ func (m *awsRestjson1_serializeOpUpdateVoiceProfileDom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVoiceProfileDomainInput(v *UpdateVoiceProfileDomainInput, encoder *httpbinding.Encoder) error {
@@ -7945,10 +7659,8 @@ func (m *awsRestjson1_serializeOpValidateE911Address) 
 func (m *awsRestjson1_serializeOpValidateE911Address) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7993,7 +7705,6 @@ func (m *awsRestjson1_serializeOpValidateE911Address) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsValidateE911AddressInput(v *ValidateE911AddressInput, encoder *httpbinding.Encoder) error {
