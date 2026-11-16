--- vendor/github.com/aws/aws-sdk-go-v2/service/sesv2/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sesv2/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpBatchGetMetricData) H
 func (m *awsRestjson1_serializeOpBatchGetMetricData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpBatchGetMetricData) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetMetricDataInput(v *BatchGetMetricDataInput, encoder *httpbinding.Encoder) error {
@@ -109,10 +105,8 @@ func (m *awsRestjson1_serializeOpCancelExportJob) Hand
 func (m *awsRestjson1_serializeOpCancelExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -150,7 +144,6 @@ func (m *awsRestjson1_serializeOpCancelExportJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelExportJobInput(v *CancelExportJobInput, encoder *httpbinding.Encoder) error {
@@ -180,10 +173,8 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 func (m *awsRestjson1_serializeOpCreateConfigurationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -228,7 +219,6 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfigurationSetInput(v *CreateConfigurationSetInput, encoder *httpbinding.Encoder) error {
@@ -317,10 +307,8 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 func (m *awsRestjson1_serializeOpCreateConfigurationSetEventDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -369,7 +357,6 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfigurationSetEventDestinationInput(v *CreateConfigurationSetEventDestinationInput, encoder *httpbinding.Encoder) error {
@@ -418,10 +405,8 @@ func (m *awsRestjson1_serializeOpCreateContact) Handle
 func (m *awsRestjson1_serializeOpCreateContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -470,7 +455,6 @@ func (m *awsRestjson1_serializeOpCreateContact) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContactInput(v *CreateContactInput, encoder *httpbinding.Encoder) error {
@@ -529,10 +513,8 @@ func (m *awsRestjson1_serializeOpCreateContactList) Ha
 func (m *awsRestjson1_serializeOpCreateContactList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -577,7 +559,6 @@ func (m *awsRestjson1_serializeOpCreateContactList) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContactListInput(v *CreateContactListInput, encoder *httpbinding.Encoder) error {
@@ -629,10 +610,8 @@ func (m *awsRestjson1_serializeOpCreateCustomVerificat
 func (m *awsRestjson1_serializeOpCreateCustomVerificationEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -677,7 +656,6 @@ func (m *awsRestjson1_serializeOpCreateCustomVerificat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCustomVerificationEmailTemplateInput(v *CreateCustomVerificationEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -735,10 +713,8 @@ func (m *awsRestjson1_serializeOpCreateDedicatedIpPool
 func (m *awsRestjson1_serializeOpCreateDedicatedIpPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -783,7 +759,6 @@ func (m *awsRestjson1_serializeOpCreateDedicatedIpPool
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDedicatedIpPoolInput(v *CreateDedicatedIpPoolInput, encoder *httpbinding.Encoder) error {
@@ -828,10 +803,8 @@ func (m *awsRestjson1_serializeOpCreateDeliverabilityT
 func (m *awsRestjson1_serializeOpCreateDeliverabilityTestReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -876,7 +849,6 @@ func (m *awsRestjson1_serializeOpCreateDeliverabilityT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeliverabilityTestReportInput(v *CreateDeliverabilityTestReportInput, encoder *httpbinding.Encoder) error {
@@ -928,10 +900,8 @@ func (m *awsRestjson1_serializeOpCreateEmailIdentity) 
 func (m *awsRestjson1_serializeOpCreateEmailIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -976,7 +946,6 @@ func (m *awsRestjson1_serializeOpCreateEmailIdentity) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEmailIdentityInput(v *CreateEmailIdentityInput, encoder *httpbinding.Encoder) error {
@@ -1028,10 +997,8 @@ func (m *awsRestjson1_serializeOpCreateEmailIdentityPo
 func (m *awsRestjson1_serializeOpCreateEmailIdentityPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1080,7 +1047,6 @@ func (m *awsRestjson1_serializeOpCreateEmailIdentityPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEmailIdentityPolicyInput(v *CreateEmailIdentityPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1131,10 +1097,8 @@ func (m *awsRestjson1_serializeOpCreateEmailTemplate) 
 func (m *awsRestjson1_serializeOpCreateEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1179,7 +1143,6 @@ func (m *awsRestjson1_serializeOpCreateEmailTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEmailTemplateInput(v *CreateEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1219,10 +1182,8 @@ func (m *awsRestjson1_serializeOpCreateExportJob) Hand
 func (m *awsRestjson1_serializeOpCreateExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1267,7 +1228,6 @@ func (m *awsRestjson1_serializeOpCreateExportJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateExportJobInput(v *CreateExportJobInput, encoder *httpbinding.Encoder) error {
@@ -1309,10 +1269,8 @@ func (m *awsRestjson1_serializeOpCreateImportJob) Hand
 func (m *awsRestjson1_serializeOpCreateImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1357,7 +1315,6 @@ func (m *awsRestjson1_serializeOpCreateImportJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImportJobInput(v *CreateImportJobInput, encoder *httpbinding.Encoder) error {
@@ -1399,10 +1356,8 @@ func (m *awsRestjson1_serializeOpCreateMultiRegionEndp
 func (m *awsRestjson1_serializeOpCreateMultiRegionEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1447,7 +1402,6 @@ func (m *awsRestjson1_serializeOpCreateMultiRegionEndp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMultiRegionEndpointInput(v *CreateMultiRegionEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1494,10 +1448,8 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 func (m *awsRestjson1_serializeOpDeleteConfigurationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1535,7 +1487,6 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfigurationSetInput(v *DeleteConfigurationSetInput, encoder *httpbinding.Encoder) error {
@@ -1565,10 +1516,8 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 func (m *awsRestjson1_serializeOpDeleteConfigurationSetEventDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1606,7 +1555,6 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfigurationSetEventDestinationInput(v *DeleteConfigurationSetEventDestinationInput, encoder *httpbinding.Encoder) error {
@@ -1645,10 +1593,8 @@ func (m *awsRestjson1_serializeOpDeleteContact) Handle
 func (m *awsRestjson1_serializeOpDeleteContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1686,7 +1632,6 @@ func (m *awsRestjson1_serializeOpDeleteContact) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContactInput(v *DeleteContactInput, encoder *httpbinding.Encoder) error {
@@ -1725,10 +1670,8 @@ func (m *awsRestjson1_serializeOpDeleteContactList) Ha
 func (m *awsRestjson1_serializeOpDeleteContactList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1766,7 +1709,6 @@ func (m *awsRestjson1_serializeOpDeleteContactList) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContactListInput(v *DeleteContactListInput, encoder *httpbinding.Encoder) error {
@@ -1796,10 +1738,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomVerificat
 func (m *awsRestjson1_serializeOpDeleteCustomVerificationEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1837,7 +1777,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomVerificat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomVerificationEmailTemplateInput(v *DeleteCustomVerificationEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1867,10 +1806,8 @@ func (m *awsRestjson1_serializeOpDeleteDedicatedIpPool
 func (m *awsRestjson1_serializeOpDeleteDedicatedIpPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1908,7 +1845,6 @@ func (m *awsRestjson1_serializeOpDeleteDedicatedIpPool
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDedicatedIpPoolInput(v *DeleteDedicatedIpPoolInput, encoder *httpbinding.Encoder) error {
@@ -1938,10 +1874,8 @@ func (m *awsRestjson1_serializeOpDeleteEmailIdentity) 
 func (m *awsRestjson1_serializeOpDeleteEmailIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1979,7 +1913,6 @@ func (m *awsRestjson1_serializeOpDeleteEmailIdentity) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEmailIdentityInput(v *DeleteEmailIdentityInput, encoder *httpbinding.Encoder) error {
@@ -2009,10 +1942,8 @@ func (m *awsRestjson1_serializeOpDeleteEmailIdentityPo
 func (m *awsRestjson1_serializeOpDeleteEmailIdentityPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2050,7 +1981,6 @@ func (m *awsRestjson1_serializeOpDeleteEmailIdentityPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEmailIdentityPolicyInput(v *DeleteEmailIdentityPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2089,10 +2019,8 @@ func (m *awsRestjson1_serializeOpDeleteEmailTemplate) 
 func (m *awsRestjson1_serializeOpDeleteEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2130,7 +2058,6 @@ func (m *awsRestjson1_serializeOpDeleteEmailTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEmailTemplateInput(v *DeleteEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2160,10 +2087,8 @@ func (m *awsRestjson1_serializeOpDeleteMultiRegionEndp
 func (m *awsRestjson1_serializeOpDeleteMultiRegionEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2201,7 +2126,6 @@ func (m *awsRestjson1_serializeOpDeleteMultiRegionEndp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMultiRegionEndpointInput(v *DeleteMultiRegionEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2231,10 +2155,8 @@ func (m *awsRestjson1_serializeOpDeleteSuppressedDesti
 func (m *awsRestjson1_serializeOpDeleteSuppressedDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2272,7 +2194,6 @@ func (m *awsRestjson1_serializeOpDeleteSuppressedDesti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSuppressedDestinationInput(v *DeleteSuppressedDestinationInput, encoder *httpbinding.Encoder) error {
@@ -2302,10 +2223,8 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 func (m *awsRestjson1_serializeOpGetAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2339,7 +2258,6 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountInput(v *GetAccountInput, encoder *httpbinding.Encoder) error {
@@ -2360,10 +2278,8 @@ func (m *awsRestjson1_serializeOpGetBlacklistReports) 
 func (m *awsRestjson1_serializeOpGetBlacklistReports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2401,7 +2317,6 @@ func (m *awsRestjson1_serializeOpGetBlacklistReports) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBlacklistReportsInput(v *GetBlacklistReportsInput, encoder *httpbinding.Encoder) error {
@@ -2428,10 +2343,8 @@ func (m *awsRestjson1_serializeOpGetConfigurationSet) 
 func (m *awsRestjson1_serializeOpGetConfigurationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2469,7 +2382,6 @@ func (m *awsRestjson1_serializeOpGetConfigurationSet) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfigurationSetInput(v *GetConfigurationSetInput, encoder *httpbinding.Encoder) error {
@@ -2499,10 +2411,8 @@ func (m *awsRestjson1_serializeOpGetConfigurationSetEv
 func (m *awsRestjson1_serializeOpGetConfigurationSetEventDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2540,7 +2450,6 @@ func (m *awsRestjson1_serializeOpGetConfigurationSetEv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfigurationSetEventDestinationsInput(v *GetConfigurationSetEventDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -2570,10 +2479,8 @@ func (m *awsRestjson1_serializeOpGetContact) HandleSer
 func (m *awsRestjson1_serializeOpGetContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2611,7 +2518,6 @@ func (m *awsRestjson1_serializeOpGetContact) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContactInput(v *GetContactInput, encoder *httpbinding.Encoder) error {
@@ -2650,10 +2556,8 @@ func (m *awsRestjson1_serializeOpGetContactList) Handl
 func (m *awsRestjson1_serializeOpGetContactList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2691,7 +2595,6 @@ func (m *awsRestjson1_serializeOpGetContactList) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContactListInput(v *GetContactListInput, encoder *httpbinding.Encoder) error {
@@ -2721,10 +2624,8 @@ func (m *awsRestjson1_serializeOpGetCustomVerification
 func (m *awsRestjson1_serializeOpGetCustomVerificationEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2762,7 +2663,6 @@ func (m *awsRestjson1_serializeOpGetCustomVerification
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCustomVerificationEmailTemplateInput(v *GetCustomVerificationEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2792,10 +2692,8 @@ func (m *awsRestjson1_serializeOpGetDedicatedIp) Handl
 func (m *awsRestjson1_serializeOpGetDedicatedIp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2833,7 +2731,6 @@ func (m *awsRestjson1_serializeOpGetDedicatedIp) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDedicatedIpInput(v *GetDedicatedIpInput, encoder *httpbinding.Encoder) error {
@@ -2863,10 +2760,8 @@ func (m *awsRestjson1_serializeOpGetDedicatedIpPool) H
 func (m *awsRestjson1_serializeOpGetDedicatedIpPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2904,7 +2799,6 @@ func (m *awsRestjson1_serializeOpGetDedicatedIpPool) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDedicatedIpPoolInput(v *GetDedicatedIpPoolInput, encoder *httpbinding.Encoder) error {
@@ -2934,10 +2828,8 @@ func (m *awsRestjson1_serializeOpGetDedicatedIps) Hand
 func (m *awsRestjson1_serializeOpGetDedicatedIps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2975,7 +2867,6 @@ func (m *awsRestjson1_serializeOpGetDedicatedIps) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDedicatedIpsInput(v *GetDedicatedIpsInput, encoder *httpbinding.Encoder) error {
@@ -3008,10 +2899,8 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityDash
 func (m *awsRestjson1_serializeOpGetDeliverabilityDashboardOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3045,7 +2934,6 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityDash
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeliverabilityDashboardOptionsInput(v *GetDeliverabilityDashboardOptionsInput, encoder *httpbinding.Encoder) error {
@@ -3066,10 +2954,8 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityTest
 func (m *awsRestjson1_serializeOpGetDeliverabilityTestReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3107,7 +2993,6 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityTest
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeliverabilityTestReportInput(v *GetDeliverabilityTestReportInput, encoder *httpbinding.Encoder) error {
@@ -3137,10 +3022,8 @@ func (m *awsRestjson1_serializeOpGetDomainDeliverabili
 func (m *awsRestjson1_serializeOpGetDomainDeliverabilityCampaign) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3178,7 +3061,6 @@ func (m *awsRestjson1_serializeOpGetDomainDeliverabili
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainDeliverabilityCampaignInput(v *GetDomainDeliverabilityCampaignInput, encoder *httpbinding.Encoder) error {
@@ -3208,10 +3090,8 @@ func (m *awsRestjson1_serializeOpGetDomainStatisticsRe
 func (m *awsRestjson1_serializeOpGetDomainStatisticsReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3249,7 +3129,6 @@ func (m *awsRestjson1_serializeOpGetDomainStatisticsRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainStatisticsReportInput(v *GetDomainStatisticsReportInput, encoder *httpbinding.Encoder) error {
@@ -3287,10 +3166,8 @@ func (m *awsRestjson1_serializeOpGetEmailIdentity) Han
 func (m *awsRestjson1_serializeOpGetEmailIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3328,7 +3205,6 @@ func (m *awsRestjson1_serializeOpGetEmailIdentity) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEmailIdentityInput(v *GetEmailIdentityInput, encoder *httpbinding.Encoder) error {
@@ -3358,10 +3234,8 @@ func (m *awsRestjson1_serializeOpGetEmailIdentityPolic
 func (m *awsRestjson1_serializeOpGetEmailIdentityPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3399,7 +3273,6 @@ func (m *awsRestjson1_serializeOpGetEmailIdentityPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEmailIdentityPoliciesInput(v *GetEmailIdentityPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -3429,10 +3302,8 @@ func (m *awsRestjson1_serializeOpGetEmailTemplate) Han
 func (m *awsRestjson1_serializeOpGetEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3470,7 +3341,6 @@ func (m *awsRestjson1_serializeOpGetEmailTemplate) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEmailTemplateInput(v *GetEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3500,10 +3370,8 @@ func (m *awsRestjson1_serializeOpGetExportJob) HandleS
 func (m *awsRestjson1_serializeOpGetExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3541,7 +3409,6 @@ func (m *awsRestjson1_serializeOpGetExportJob) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetExportJobInput(v *GetExportJobInput, encoder *httpbinding.Encoder) error {
@@ -3571,10 +3438,8 @@ func (m *awsRestjson1_serializeOpGetImportJob) HandleS
 func (m *awsRestjson1_serializeOpGetImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3612,7 +3477,6 @@ func (m *awsRestjson1_serializeOpGetImportJob) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImportJobInput(v *GetImportJobInput, encoder *httpbinding.Encoder) error {
@@ -3642,10 +3506,8 @@ func (m *awsRestjson1_serializeOpGetMessageInsights) H
 func (m *awsRestjson1_serializeOpGetMessageInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3683,7 +3545,6 @@ func (m *awsRestjson1_serializeOpGetMessageInsights) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMessageInsightsInput(v *GetMessageInsightsInput, encoder *httpbinding.Encoder) error {
@@ -3713,10 +3574,8 @@ func (m *awsRestjson1_serializeOpGetMultiRegionEndpoin
 func (m *awsRestjson1_serializeOpGetMultiRegionEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3754,7 +3613,6 @@ func (m *awsRestjson1_serializeOpGetMultiRegionEndpoin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMultiRegionEndpointInput(v *GetMultiRegionEndpointInput, encoder *httpbinding.Encoder) error {
@@ -3784,10 +3642,8 @@ func (m *awsRestjson1_serializeOpGetSuppressedDestinat
 func (m *awsRestjson1_serializeOpGetSuppressedDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3825,7 +3681,6 @@ func (m *awsRestjson1_serializeOpGetSuppressedDestinat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSuppressedDestinationInput(v *GetSuppressedDestinationInput, encoder *httpbinding.Encoder) error {
@@ -3855,10 +3710,8 @@ func (m *awsRestjson1_serializeOpListConfigurationSets
 func (m *awsRestjson1_serializeOpListConfigurationSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3896,7 +3749,6 @@ func (m *awsRestjson1_serializeOpListConfigurationSets
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfigurationSetsInput(v *ListConfigurationSetsInput, encoder *httpbinding.Encoder) error {
@@ -3925,10 +3777,8 @@ func (m *awsRestjson1_serializeOpListContactLists) Han
 func (m *awsRestjson1_serializeOpListContactLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3966,7 +3816,6 @@ func (m *awsRestjson1_serializeOpListContactLists) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContactListsInput(v *ListContactListsInput, encoder *httpbinding.Encoder) error {
@@ -3995,10 +3844,8 @@ func (m *awsRestjson1_serializeOpListContacts) HandleS
 func (m *awsRestjson1_serializeOpListContacts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4047,7 +3894,6 @@ func (m *awsRestjson1_serializeOpListContacts) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContactsInput(v *ListContactsInput, encoder *httpbinding.Encoder) error {
@@ -4101,10 +3947,8 @@ func (m *awsRestjson1_serializeOpListCustomVerificatio
 func (m *awsRestjson1_serializeOpListCustomVerificationEmailTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4142,7 +3986,6 @@ func (m *awsRestjson1_serializeOpListCustomVerificatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomVerificationEmailTemplatesInput(v *ListCustomVerificationEmailTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -4171,10 +4014,8 @@ func (m *awsRestjson1_serializeOpListDedicatedIpPools)
 func (m *awsRestjson1_serializeOpListDedicatedIpPools) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4212,7 +4053,6 @@ func (m *awsRestjson1_serializeOpListDedicatedIpPools)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDedicatedIpPoolsInput(v *ListDedicatedIpPoolsInput, encoder *httpbinding.Encoder) error {
@@ -4241,10 +4081,8 @@ func (m *awsRestjson1_serializeOpListDeliverabilityTes
 func (m *awsRestjson1_serializeOpListDeliverabilityTestReports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4282,7 +4120,6 @@ func (m *awsRestjson1_serializeOpListDeliverabilityTes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDeliverabilityTestReportsInput(v *ListDeliverabilityTestReportsInput, encoder *httpbinding.Encoder) error {
@@ -4311,10 +4148,8 @@ func (m *awsRestjson1_serializeOpListDomainDeliverabil
 func (m *awsRestjson1_serializeOpListDomainDeliverabilityCampaigns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4352,7 +4187,6 @@ func (m *awsRestjson1_serializeOpListDomainDeliverabil
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainDeliverabilityCampaignsInput(v *ListDomainDeliverabilityCampaignsInput, encoder *httpbinding.Encoder) error {
@@ -4398,10 +4232,8 @@ func (m *awsRestjson1_serializeOpListEmailIdentities) 
 func (m *awsRestjson1_serializeOpListEmailIdentities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4439,7 +4271,6 @@ func (m *awsRestjson1_serializeOpListEmailIdentities) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEmailIdentitiesInput(v *ListEmailIdentitiesInput, encoder *httpbinding.Encoder) error {
@@ -4468,10 +4299,8 @@ func (m *awsRestjson1_serializeOpListEmailTemplates) H
 func (m *awsRestjson1_serializeOpListEmailTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4509,7 +4338,6 @@ func (m *awsRestjson1_serializeOpListEmailTemplates) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEmailTemplatesInput(v *ListEmailTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -4538,10 +4366,8 @@ func (m *awsRestjson1_serializeOpListExportJobs) Handl
 func (m *awsRestjson1_serializeOpListExportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4586,7 +4412,6 @@ func (m *awsRestjson1_serializeOpListExportJobs) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListExportJobsInput(v *ListExportJobsInput, encoder *httpbinding.Encoder) error {
@@ -4634,10 +4459,8 @@ func (m *awsRestjson1_serializeOpListImportJobs) Handl
 func (m *awsRestjson1_serializeOpListImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4682,7 +4505,6 @@ func (m *awsRestjson1_serializeOpListImportJobs) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImportJobsInput(v *ListImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -4725,10 +4547,8 @@ func (m *awsRestjson1_serializeOpListMultiRegionEndpoi
 func (m *awsRestjson1_serializeOpListMultiRegionEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4766,7 +4586,6 @@ func (m *awsRestjson1_serializeOpListMultiRegionEndpoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMultiRegionEndpointsInput(v *ListMultiRegionEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -4795,10 +4614,8 @@ func (m *awsRestjson1_serializeOpListRecommendations) 
 func (m *awsRestjson1_serializeOpListRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4843,7 +4660,6 @@ func (m *awsRestjson1_serializeOpListRecommendations) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecommendationsInput(v *ListRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -4888,10 +4704,8 @@ func (m *awsRestjson1_serializeOpListSuppressedDestina
 func (m *awsRestjson1_serializeOpListSuppressedDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4929,7 +4743,6 @@ func (m *awsRestjson1_serializeOpListSuppressedDestina
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSuppressedDestinationsInput(v *ListSuppressedDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -4972,10 +4785,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -5013,7 +4824,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -5038,10 +4848,8 @@ func (m *awsRestjson1_serializeOpPutAccountDedicatedIp
 func (m *awsRestjson1_serializeOpPutAccountDedicatedIpWarmupAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5086,7 +4894,6 @@ func (m *awsRestjson1_serializeOpPutAccountDedicatedIp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountDedicatedIpWarmupAttributesInput(v *PutAccountDedicatedIpWarmupAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5119,10 +4926,8 @@ func (m *awsRestjson1_serializeOpPutAccountDetails) Ha
 func (m *awsRestjson1_serializeOpPutAccountDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5167,7 +4972,6 @@ func (m *awsRestjson1_serializeOpPutAccountDetails) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountDetailsInput(v *PutAccountDetailsInput, encoder *httpbinding.Encoder) error {
@@ -5227,10 +5031,8 @@ func (m *awsRestjson1_serializeOpPutAccountSendingAttr
 func (m *awsRestjson1_serializeOpPutAccountSendingAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5275,7 +5077,6 @@ func (m *awsRestjson1_serializeOpPutAccountSendingAttr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountSendingAttributesInput(v *PutAccountSendingAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5308,10 +5109,8 @@ func (m *awsRestjson1_serializeOpPutAccountSuppression
 func (m *awsRestjson1_serializeOpPutAccountSuppressionAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5356,7 +5155,6 @@ func (m *awsRestjson1_serializeOpPutAccountSuppression
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountSuppressionAttributesInput(v *PutAccountSuppressionAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5391,10 +5189,8 @@ func (m *awsRestjson1_serializeOpPutAccountVdmAttribut
 func (m *awsRestjson1_serializeOpPutAccountVdmAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5439,7 +5235,6 @@ func (m *awsRestjson1_serializeOpPutAccountVdmAttribut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountVdmAttributesInput(v *PutAccountVdmAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5474,10 +5269,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetAr
 func (m *awsRestjson1_serializeOpPutConfigurationSetArchivingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5526,7 +5319,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetAr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetArchivingOptionsInput(v *PutConfigurationSetArchivingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -5568,10 +5360,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetDe
 func (m *awsRestjson1_serializeOpPutConfigurationSetDeliveryOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5620,7 +5410,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetDe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetDeliveryOptionsInput(v *PutConfigurationSetDeliveryOptionsInput, encoder *httpbinding.Encoder) error {
@@ -5672,10 +5461,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetRe
 func (m *awsRestjson1_serializeOpPutConfigurationSetReputationOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5724,7 +5511,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetReputationOptionsInput(v *PutConfigurationSetReputationOptionsInput, encoder *httpbinding.Encoder) error {
@@ -5766,10 +5552,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetSe
 func (m *awsRestjson1_serializeOpPutConfigurationSetSendingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5818,7 +5602,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetSendingOptionsInput(v *PutConfigurationSetSendingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -5860,10 +5643,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetSu
 func (m *awsRestjson1_serializeOpPutConfigurationSetSuppressionOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5912,7 +5693,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetSu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetSuppressionOptionsInput(v *PutConfigurationSetSuppressionOptionsInput, encoder *httpbinding.Encoder) error {
@@ -5956,10 +5736,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetTr
 func (m *awsRestjson1_serializeOpPutConfigurationSetTrackingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6008,7 +5786,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetTr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetTrackingOptionsInput(v *PutConfigurationSetTrackingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -6055,10 +5832,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetVd
 func (m *awsRestjson1_serializeOpPutConfigurationSetVdmOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6107,7 +5882,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetVd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetVdmOptionsInput(v *PutConfigurationSetVdmOptionsInput, encoder *httpbinding.Encoder) error {
@@ -6151,10 +5925,8 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpInPool)
 func (m *awsRestjson1_serializeOpPutDedicatedIpInPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6203,7 +5975,6 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpInPool)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDedicatedIpInPoolInput(v *PutDedicatedIpInPoolInput, encoder *httpbinding.Encoder) error {
@@ -6245,10 +6016,8 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpPoolSca
 func (m *awsRestjson1_serializeOpPutDedicatedIpPoolScalingAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6297,7 +6066,6 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpPoolSca
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDedicatedIpPoolScalingAttributesInput(v *PutDedicatedIpPoolScalingAttributesInput, encoder *httpbinding.Encoder) error {
@@ -6339,10 +6107,8 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpWarmupA
 func (m *awsRestjson1_serializeOpPutDedicatedIpWarmupAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6391,7 +6157,6 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpWarmupA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDedicatedIpWarmupAttributesInput(v *PutDedicatedIpWarmupAttributesInput, encoder *httpbinding.Encoder) error {
@@ -6433,10 +6198,8 @@ func (m *awsRestjson1_serializeOpPutDeliverabilityDash
 func (m *awsRestjson1_serializeOpPutDeliverabilityDashboardOption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6481,7 +6244,6 @@ func (m *awsRestjson1_serializeOpPutDeliverabilityDash
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDeliverabilityDashboardOptionInput(v *PutDeliverabilityDashboardOptionInput, encoder *httpbinding.Encoder) error {
@@ -6521,10 +6283,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityConfi
 func (m *awsRestjson1_serializeOpPutEmailIdentityConfigurationSetAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6573,7 +6333,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityConfigurationSetAttributesInput(v *PutEmailIdentityConfigurationSetAttributesInput, encoder *httpbinding.Encoder) error {
@@ -6615,10 +6374,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityDkimA
 func (m *awsRestjson1_serializeOpPutEmailIdentityDkimAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6667,7 +6424,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityDkimA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityDkimAttributesInput(v *PutEmailIdentityDkimAttributesInput, encoder *httpbinding.Encoder) error {
@@ -6709,10 +6465,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityDkimS
 func (m *awsRestjson1_serializeOpPutEmailIdentityDkimSigningAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6761,7 +6515,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityDkimS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityDkimSigningAttributesInput(v *PutEmailIdentityDkimSigningAttributesInput, encoder *httpbinding.Encoder) error {
@@ -6810,10 +6563,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityFeedb
 func (m *awsRestjson1_serializeOpPutEmailIdentityFeedbackAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6862,7 +6613,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityFeedb
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityFeedbackAttributesInput(v *PutEmailIdentityFeedbackAttributesInput, encoder *httpbinding.Encoder) error {
@@ -6904,10 +6654,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityMailF
 func (m *awsRestjson1_serializeOpPutEmailIdentityMailFromAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6956,7 +6704,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityMailF
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityMailFromAttributesInput(v *PutEmailIdentityMailFromAttributesInput, encoder *httpbinding.Encoder) error {
@@ -7003,10 +6750,8 @@ func (m *awsRestjson1_serializeOpPutSuppressedDestinat
 func (m *awsRestjson1_serializeOpPutSuppressedDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7051,7 +6796,6 @@ func (m *awsRestjson1_serializeOpPutSuppressedDestinat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutSuppressedDestinationInput(v *PutSuppressedDestinationInput, encoder *httpbinding.Encoder) error {
@@ -7089,10 +6833,8 @@ func (m *awsRestjson1_serializeOpSendBulkEmail) Handle
 func (m *awsRestjson1_serializeOpSendBulkEmail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7137,7 +6879,6 @@ func (m *awsRestjson1_serializeOpSendBulkEmail) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendBulkEmailInput(v *SendBulkEmailInput, encoder *httpbinding.Encoder) error {
@@ -7223,10 +6964,8 @@ func (m *awsRestjson1_serializeOpSendCustomVerificatio
 func (m *awsRestjson1_serializeOpSendCustomVerificationEmail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7271,7 +7010,6 @@ func (m *awsRestjson1_serializeOpSendCustomVerificatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendCustomVerificationEmailInput(v *SendCustomVerificationEmailInput, encoder *httpbinding.Encoder) error {
@@ -7314,10 +7052,8 @@ func (m *awsRestjson1_serializeOpSendEmail) HandleSeri
 func (m *awsRestjson1_serializeOpSendEmail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7362,7 +7098,6 @@ func (m *awsRestjson1_serializeOpSendEmail) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendEmailInput(v *SendEmailInput, encoder *httpbinding.Encoder) error {
@@ -7455,10 +7190,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7503,7 +7236,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7543,10 +7275,8 @@ func (m *awsRestjson1_serializeOpTestRenderEmailTempla
 func (m *awsRestjson1_serializeOpTestRenderEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7595,7 +7325,6 @@ func (m *awsRestjson1_serializeOpTestRenderEmailTempla
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestRenderEmailTemplateInput(v *TestRenderEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -7637,10 +7366,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7678,7 +7405,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7709,10 +7435,8 @@ func (m *awsRestjson1_serializeOpUpdateConfigurationSe
 func (m *awsRestjson1_serializeOpUpdateConfigurationSetEventDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7761,7 +7485,6 @@ func (m *awsRestjson1_serializeOpUpdateConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfigurationSetEventDestinationInput(v *UpdateConfigurationSetEventDestinationInput, encoder *httpbinding.Encoder) error {
@@ -7814,10 +7537,8 @@ func (m *awsRestjson1_serializeOpUpdateContact) Handle
 func (m *awsRestjson1_serializeOpUpdateContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7866,7 +7587,6 @@ func (m *awsRestjson1_serializeOpUpdateContact) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactInput(v *UpdateContactInput, encoder *httpbinding.Encoder) error {
@@ -7929,10 +7649,8 @@ func (m *awsRestjson1_serializeOpUpdateContactList) Ha
 func (m *awsRestjson1_serializeOpUpdateContactList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7981,7 +7699,6 @@ func (m *awsRestjson1_serializeOpUpdateContactList) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactListInput(v *UpdateContactListInput, encoder *httpbinding.Encoder) error {
@@ -8030,10 +7747,8 @@ func (m *awsRestjson1_serializeOpUpdateCustomVerificat
 func (m *awsRestjson1_serializeOpUpdateCustomVerificationEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8082,7 +7797,6 @@ func (m *awsRestjson1_serializeOpUpdateCustomVerificat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCustomVerificationEmailTemplateInput(v *UpdateCustomVerificationEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -8144,10 +7858,8 @@ func (m *awsRestjson1_serializeOpUpdateEmailIdentityPo
 func (m *awsRestjson1_serializeOpUpdateEmailIdentityPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8196,7 +7908,6 @@ func (m *awsRestjson1_serializeOpUpdateEmailIdentityPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEmailIdentityPolicyInput(v *UpdateEmailIdentityPolicyInput, encoder *httpbinding.Encoder) error {
@@ -8247,10 +7958,8 @@ func (m *awsRestjson1_serializeOpUpdateEmailTemplate) 
 func (m *awsRestjson1_serializeOpUpdateEmailTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8299,7 +8008,6 @@ func (m *awsRestjson1_serializeOpUpdateEmailTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEmailTemplateInput(v *UpdateEmailTemplateInput, encoder *httpbinding.Encoder) error {
