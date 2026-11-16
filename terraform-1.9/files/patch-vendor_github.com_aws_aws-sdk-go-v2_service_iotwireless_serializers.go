--- vendor/github.com/aws/aws-sdk-go-v2/service/iotwireless/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotwireless/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAssociateAwsAccountWi
 func (m *awsRestjson1_serializeOpAssociateAwsAccountWithPartnerAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -75,7 +72,6 @@ func (m *awsRestjson1_serializeOpAssociateAwsAccountWi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateAwsAccountWithPartnerAccountInput(v *AssociateAwsAccountWithPartnerAccountInput, encoder *httpbinding.Encoder) error {
@@ -122,10 +118,8 @@ func (m *awsRestjson1_serializeOpAssociateMulticastGro
 func (m *awsRestjson1_serializeOpAssociateMulticastGroupWithFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -174,7 +168,6 @@ func (m *awsRestjson1_serializeOpAssociateMulticastGro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateMulticastGroupWithFuotaTaskInput(v *AssociateMulticastGroupWithFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -216,10 +209,8 @@ func (m *awsRestjson1_serializeOpAssociateWirelessDevi
 func (m *awsRestjson1_serializeOpAssociateWirelessDeviceWithFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -268,7 +259,6 @@ func (m *awsRestjson1_serializeOpAssociateWirelessDevi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateWirelessDeviceWithFuotaTaskInput(v *AssociateWirelessDeviceWithFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -310,10 +300,8 @@ func (m *awsRestjson1_serializeOpAssociateWirelessDevi
 func (m *awsRestjson1_serializeOpAssociateWirelessDeviceWithMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -362,7 +350,6 @@ func (m *awsRestjson1_serializeOpAssociateWirelessDevi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateWirelessDeviceWithMulticastGroupInput(v *AssociateWirelessDeviceWithMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -404,10 +391,8 @@ func (m *awsRestjson1_serializeOpAssociateWirelessDevi
 func (m *awsRestjson1_serializeOpAssociateWirelessDeviceWithThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -456,7 +441,6 @@ func (m *awsRestjson1_serializeOpAssociateWirelessDevi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateWirelessDeviceWithThingInput(v *AssociateWirelessDeviceWithThingInput, encoder *httpbinding.Encoder) error {
@@ -498,10 +482,8 @@ func (m *awsRestjson1_serializeOpAssociateWirelessGate
 func (m *awsRestjson1_serializeOpAssociateWirelessGatewayWithCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -550,7 +532,6 @@ func (m *awsRestjson1_serializeOpAssociateWirelessGate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateWirelessGatewayWithCertificateInput(v *AssociateWirelessGatewayWithCertificateInput, encoder *httpbinding.Encoder) error {
@@ -592,10 +573,8 @@ func (m *awsRestjson1_serializeOpAssociateWirelessGate
 func (m *awsRestjson1_serializeOpAssociateWirelessGatewayWithThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -644,7 +623,6 @@ func (m *awsRestjson1_serializeOpAssociateWirelessGate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateWirelessGatewayWithThingInput(v *AssociateWirelessGatewayWithThingInput, encoder *httpbinding.Encoder) error {
@@ -686,10 +664,8 @@ func (m *awsRestjson1_serializeOpCancelMulticastGroupS
 func (m *awsRestjson1_serializeOpCancelMulticastGroupSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -727,7 +703,6 @@ func (m *awsRestjson1_serializeOpCancelMulticastGroupS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelMulticastGroupSessionInput(v *CancelMulticastGroupSessionInput, encoder *httpbinding.Encoder) error {
@@ -757,10 +732,8 @@ func (m *awsRestjson1_serializeOpCreateDestination) Ha
 func (m *awsRestjson1_serializeOpCreateDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -805,7 +778,6 @@ func (m *awsRestjson1_serializeOpCreateDestination) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDestinationInput(v *CreateDestinationInput, encoder *httpbinding.Encoder) error {
@@ -870,10 +842,8 @@ func (m *awsRestjson1_serializeOpCreateDeviceProfile) 
 func (m *awsRestjson1_serializeOpCreateDeviceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -918,7 +888,6 @@ func (m *awsRestjson1_serializeOpCreateDeviceProfile) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeviceProfileInput(v *CreateDeviceProfileInput, encoder *httpbinding.Encoder) error {
@@ -977,10 +946,8 @@ func (m *awsRestjson1_serializeOpCreateFuotaTask) Hand
 func (m *awsRestjson1_serializeOpCreateFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1025,7 +992,6 @@ func (m *awsRestjson1_serializeOpCreateFuotaTask) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFuotaTaskInput(v *CreateFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -1112,10 +1078,8 @@ func (m *awsRestjson1_serializeOpCreateMulticastGroup)
 func (m *awsRestjson1_serializeOpCreateMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1160,7 +1124,6 @@ func (m *awsRestjson1_serializeOpCreateMulticastGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMulticastGroupInput(v *CreateMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -1217,10 +1180,8 @@ func (m *awsRestjson1_serializeOpCreateNetworkAnalyzer
 func (m *awsRestjson1_serializeOpCreateNetworkAnalyzerConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1265,7 +1226,6 @@ func (m *awsRestjson1_serializeOpCreateNetworkAnalyzer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNetworkAnalyzerConfigurationInput(v *CreateNetworkAnalyzerConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1343,10 +1303,8 @@ func (m *awsRestjson1_serializeOpCreateServiceProfile)
 func (m *awsRestjson1_serializeOpCreateServiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1391,7 +1349,6 @@ func (m *awsRestjson1_serializeOpCreateServiceProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateServiceProfileInput(v *CreateServiceProfileInput, encoder *httpbinding.Encoder) error {
@@ -1443,10 +1400,8 @@ func (m *awsRestjson1_serializeOpCreateWirelessDevice)
 func (m *awsRestjson1_serializeOpCreateWirelessDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1491,7 +1446,6 @@ func (m *awsRestjson1_serializeOpCreateWirelessDevice)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWirelessDeviceInput(v *CreateWirelessDeviceInput, encoder *httpbinding.Encoder) error {
@@ -1570,10 +1524,8 @@ func (m *awsRestjson1_serializeOpCreateWirelessGateway
 func (m *awsRestjson1_serializeOpCreateWirelessGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1618,7 +1570,6 @@ func (m *awsRestjson1_serializeOpCreateWirelessGateway
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWirelessGatewayInput(v *CreateWirelessGatewayInput, encoder *httpbinding.Encoder) error {
@@ -1675,10 +1626,8 @@ func (m *awsRestjson1_serializeOpCreateWirelessGateway
 func (m *awsRestjson1_serializeOpCreateWirelessGatewayTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1727,7 +1676,6 @@ func (m *awsRestjson1_serializeOpCreateWirelessGateway
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWirelessGatewayTaskInput(v *CreateWirelessGatewayTaskInput, encoder *httpbinding.Encoder) error {
@@ -1769,10 +1717,8 @@ func (m *awsRestjson1_serializeOpCreateWirelessGateway
 func (m *awsRestjson1_serializeOpCreateWirelessGatewayTaskDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1817,7 +1763,6 @@ func (m *awsRestjson1_serializeOpCreateWirelessGateway
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWirelessGatewayTaskDefinitionInput(v *CreateWirelessGatewayTaskDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -1874,10 +1819,8 @@ func (m *awsRestjson1_serializeOpDeleteDestination) Ha
 func (m *awsRestjson1_serializeOpDeleteDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1915,7 +1858,6 @@ func (m *awsRestjson1_serializeOpDeleteDestination) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDestinationInput(v *DeleteDestinationInput, encoder *httpbinding.Encoder) error {
@@ -1945,10 +1887,8 @@ func (m *awsRestjson1_serializeOpDeleteDeviceProfile) 
 func (m *awsRestjson1_serializeOpDeleteDeviceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1986,7 +1926,6 @@ func (m *awsRestjson1_serializeOpDeleteDeviceProfile) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDeviceProfileInput(v *DeleteDeviceProfileInput, encoder *httpbinding.Encoder) error {
@@ -2016,10 +1955,8 @@ func (m *awsRestjson1_serializeOpDeleteFuotaTask) Hand
 func (m *awsRestjson1_serializeOpDeleteFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2057,7 +1994,6 @@ func (m *awsRestjson1_serializeOpDeleteFuotaTask) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFuotaTaskInput(v *DeleteFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -2087,10 +2023,8 @@ func (m *awsRestjson1_serializeOpDeleteMulticastGroup)
 func (m *awsRestjson1_serializeOpDeleteMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2128,7 +2062,6 @@ func (m *awsRestjson1_serializeOpDeleteMulticastGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMulticastGroupInput(v *DeleteMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -2158,10 +2091,8 @@ func (m *awsRestjson1_serializeOpDeleteNetworkAnalyzer
 func (m *awsRestjson1_serializeOpDeleteNetworkAnalyzerConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2199,7 +2130,6 @@ func (m *awsRestjson1_serializeOpDeleteNetworkAnalyzer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteNetworkAnalyzerConfigurationInput(v *DeleteNetworkAnalyzerConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2229,10 +2159,8 @@ func (m *awsRestjson1_serializeOpDeleteQueuedMessages)
 func (m *awsRestjson1_serializeOpDeleteQueuedMessages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2270,7 +2198,6 @@ func (m *awsRestjson1_serializeOpDeleteQueuedMessages)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQueuedMessagesInput(v *DeleteQueuedMessagesInput, encoder *httpbinding.Encoder) error {
@@ -2308,10 +2235,8 @@ func (m *awsRestjson1_serializeOpDeleteServiceProfile)
 func (m *awsRestjson1_serializeOpDeleteServiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2349,7 +2274,6 @@ func (m *awsRestjson1_serializeOpDeleteServiceProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteServiceProfileInput(v *DeleteServiceProfileInput, encoder *httpbinding.Encoder) error {
@@ -2379,10 +2303,8 @@ func (m *awsRestjson1_serializeOpDeleteWirelessDevice)
 func (m *awsRestjson1_serializeOpDeleteWirelessDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2420,7 +2342,6 @@ func (m *awsRestjson1_serializeOpDeleteWirelessDevice)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWirelessDeviceInput(v *DeleteWirelessDeviceInput, encoder *httpbinding.Encoder) error {
@@ -2450,10 +2371,8 @@ func (m *awsRestjson1_serializeOpDeleteWirelessDeviceI
 func (m *awsRestjson1_serializeOpDeleteWirelessDeviceImportTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2491,7 +2410,6 @@ func (m *awsRestjson1_serializeOpDeleteWirelessDeviceI
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWirelessDeviceImportTaskInput(v *DeleteWirelessDeviceImportTaskInput, encoder *httpbinding.Encoder) error {
@@ -2521,10 +2439,8 @@ func (m *awsRestjson1_serializeOpDeleteWirelessGateway
 func (m *awsRestjson1_serializeOpDeleteWirelessGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2562,7 +2478,6 @@ func (m *awsRestjson1_serializeOpDeleteWirelessGateway
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWirelessGatewayInput(v *DeleteWirelessGatewayInput, encoder *httpbinding.Encoder) error {
@@ -2592,10 +2507,8 @@ func (m *awsRestjson1_serializeOpDeleteWirelessGateway
 func (m *awsRestjson1_serializeOpDeleteWirelessGatewayTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2633,7 +2546,6 @@ func (m *awsRestjson1_serializeOpDeleteWirelessGateway
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWirelessGatewayTaskInput(v *DeleteWirelessGatewayTaskInput, encoder *httpbinding.Encoder) error {
@@ -2663,10 +2575,8 @@ func (m *awsRestjson1_serializeOpDeleteWirelessGateway
 func (m *awsRestjson1_serializeOpDeleteWirelessGatewayTaskDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2704,7 +2614,6 @@ func (m *awsRestjson1_serializeOpDeleteWirelessGateway
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWirelessGatewayTaskDefinitionInput(v *DeleteWirelessGatewayTaskDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2734,10 +2643,8 @@ func (m *awsRestjson1_serializeOpDeregisterWirelessDev
 func (m *awsRestjson1_serializeOpDeregisterWirelessDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2775,7 +2682,6 @@ func (m *awsRestjson1_serializeOpDeregisterWirelessDev
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeregisterWirelessDeviceInput(v *DeregisterWirelessDeviceInput, encoder *httpbinding.Encoder) error {
@@ -2809,10 +2715,8 @@ func (m *awsRestjson1_serializeOpDisassociateAwsAccoun
 func (m *awsRestjson1_serializeOpDisassociateAwsAccountFromPartnerAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2850,7 +2754,6 @@ func (m *awsRestjson1_serializeOpDisassociateAwsAccoun
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateAwsAccountFromPartnerAccountInput(v *DisassociateAwsAccountFromPartnerAccountInput, encoder *httpbinding.Encoder) error {
@@ -2884,10 +2787,8 @@ func (m *awsRestjson1_serializeOpDisassociateMulticast
 func (m *awsRestjson1_serializeOpDisassociateMulticastGroupFromFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2925,7 +2826,6 @@ func (m *awsRestjson1_serializeOpDisassociateMulticast
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMulticastGroupFromFuotaTaskInput(v *DisassociateMulticastGroupFromFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -2964,10 +2864,8 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessD
 func (m *awsRestjson1_serializeOpDisassociateWirelessDeviceFromFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3005,7 +2903,6 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateWirelessDeviceFromFuotaTaskInput(v *DisassociateWirelessDeviceFromFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -3044,10 +2941,8 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessD
 func (m *awsRestjson1_serializeOpDisassociateWirelessDeviceFromMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3085,7 +2980,6 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateWirelessDeviceFromMulticastGroupInput(v *DisassociateWirelessDeviceFromMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -3124,10 +3018,8 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessD
 func (m *awsRestjson1_serializeOpDisassociateWirelessDeviceFromThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3165,7 +3057,6 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateWirelessDeviceFromThingInput(v *DisassociateWirelessDeviceFromThingInput, encoder *httpbinding.Encoder) error {
@@ -3195,10 +3086,8 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessG
 func (m *awsRestjson1_serializeOpDisassociateWirelessGatewayFromCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3236,7 +3125,6 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateWirelessGatewayFromCertificateInput(v *DisassociateWirelessGatewayFromCertificateInput, encoder *httpbinding.Encoder) error {
@@ -3266,10 +3154,8 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessG
 func (m *awsRestjson1_serializeOpDisassociateWirelessGatewayFromThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3307,7 +3193,6 @@ func (m *awsRestjson1_serializeOpDisassociateWirelessG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateWirelessGatewayFromThingInput(v *DisassociateWirelessGatewayFromThingInput, encoder *httpbinding.Encoder) error {
@@ -3337,10 +3222,8 @@ func (m *awsRestjson1_serializeOpGetDestination) Handl
 func (m *awsRestjson1_serializeOpGetDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3378,7 +3261,6 @@ func (m *awsRestjson1_serializeOpGetDestination) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDestinationInput(v *GetDestinationInput, encoder *httpbinding.Encoder) error {
@@ -3408,10 +3290,8 @@ func (m *awsRestjson1_serializeOpGetDeviceProfile) Han
 func (m *awsRestjson1_serializeOpGetDeviceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3449,7 +3329,6 @@ func (m *awsRestjson1_serializeOpGetDeviceProfile) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeviceProfileInput(v *GetDeviceProfileInput, encoder *httpbinding.Encoder) error {
@@ -3479,10 +3358,8 @@ func (m *awsRestjson1_serializeOpGetEventConfiguration
 func (m *awsRestjson1_serializeOpGetEventConfigurationByResourceTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3516,7 +3393,6 @@ func (m *awsRestjson1_serializeOpGetEventConfiguration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventConfigurationByResourceTypesInput(v *GetEventConfigurationByResourceTypesInput, encoder *httpbinding.Encoder) error {
@@ -3537,10 +3413,8 @@ func (m *awsRestjson1_serializeOpGetFuotaTask) HandleS
 func (m *awsRestjson1_serializeOpGetFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3578,7 +3452,6 @@ func (m *awsRestjson1_serializeOpGetFuotaTask) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFuotaTaskInput(v *GetFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -3608,10 +3481,8 @@ func (m *awsRestjson1_serializeOpGetLogLevelsByResourc
 func (m *awsRestjson1_serializeOpGetLogLevelsByResourceTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3645,7 +3516,6 @@ func (m *awsRestjson1_serializeOpGetLogLevelsByResourc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLogLevelsByResourceTypesInput(v *GetLogLevelsByResourceTypesInput, encoder *httpbinding.Encoder) error {
@@ -3666,10 +3536,8 @@ func (m *awsRestjson1_serializeOpGetMetricConfiguratio
 func (m *awsRestjson1_serializeOpGetMetricConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3703,7 +3571,6 @@ func (m *awsRestjson1_serializeOpGetMetricConfiguratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMetricConfigurationInput(v *GetMetricConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3724,10 +3591,8 @@ func (m *awsRestjson1_serializeOpGetMetrics) HandleSer
 func (m *awsRestjson1_serializeOpGetMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3772,7 +3637,6 @@ func (m *awsRestjson1_serializeOpGetMetrics) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMetricsInput(v *GetMetricsInput, encoder *httpbinding.Encoder) error {
@@ -3807,10 +3671,8 @@ func (m *awsRestjson1_serializeOpGetMulticastGroup) Ha
 func (m *awsRestjson1_serializeOpGetMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3848,7 +3710,6 @@ func (m *awsRestjson1_serializeOpGetMulticastGroup) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMulticastGroupInput(v *GetMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -3878,10 +3739,8 @@ func (m *awsRestjson1_serializeOpGetMulticastGroupSess
 func (m *awsRestjson1_serializeOpGetMulticastGroupSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3919,7 +3778,6 @@ func (m *awsRestjson1_serializeOpGetMulticastGroupSess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMulticastGroupSessionInput(v *GetMulticastGroupSessionInput, encoder *httpbinding.Encoder) error {
@@ -3949,10 +3807,8 @@ func (m *awsRestjson1_serializeOpGetNetworkAnalyzerCon
 func (m *awsRestjson1_serializeOpGetNetworkAnalyzerConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3990,7 +3846,6 @@ func (m *awsRestjson1_serializeOpGetNetworkAnalyzerCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNetworkAnalyzerConfigurationInput(v *GetNetworkAnalyzerConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4020,10 +3875,8 @@ func (m *awsRestjson1_serializeOpGetPartnerAccount) Ha
 func (m *awsRestjson1_serializeOpGetPartnerAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4061,7 +3914,6 @@ func (m *awsRestjson1_serializeOpGetPartnerAccount) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPartnerAccountInput(v *GetPartnerAccountInput, encoder *httpbinding.Encoder) error {
@@ -4095,10 +3947,8 @@ func (m *awsRestjson1_serializeOpGetPosition) HandleSe
 func (m *awsRestjson1_serializeOpGetPosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4136,7 +3986,6 @@ func (m *awsRestjson1_serializeOpGetPosition) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPositionInput(v *GetPositionInput, encoder *httpbinding.Encoder) error {
@@ -4170,10 +4019,8 @@ func (m *awsRestjson1_serializeOpGetPositionConfigurat
 func (m *awsRestjson1_serializeOpGetPositionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4211,7 +4058,6 @@ func (m *awsRestjson1_serializeOpGetPositionConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPositionConfigurationInput(v *GetPositionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4245,10 +4091,8 @@ func (m *awsRestjson1_serializeOpGetPositionEstimate) 
 func (m *awsRestjson1_serializeOpGetPositionEstimate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4293,7 +4137,6 @@ func (m *awsRestjson1_serializeOpGetPositionEstimate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPositionEstimateInput(v *GetPositionEstimateInput, encoder *httpbinding.Encoder) error {
@@ -4354,10 +4197,8 @@ func (m *awsRestjson1_serializeOpGetResourceEventConfi
 func (m *awsRestjson1_serializeOpGetResourceEventConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4395,7 +4236,6 @@ func (m *awsRestjson1_serializeOpGetResourceEventConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceEventConfigurationInput(v *GetResourceEventConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4433,10 +4273,8 @@ func (m *awsRestjson1_serializeOpGetResourceLogLevel) 
 func (m *awsRestjson1_serializeOpGetResourceLogLevel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4474,7 +4312,6 @@ func (m *awsRestjson1_serializeOpGetResourceLogLevel) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceLogLevelInput(v *GetResourceLogLevelInput, encoder *httpbinding.Encoder) error {
@@ -4508,10 +4345,8 @@ func (m *awsRestjson1_serializeOpGetResourcePosition) 
 func (m *awsRestjson1_serializeOpGetResourcePosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4549,7 +4384,6 @@ func (m *awsRestjson1_serializeOpGetResourcePosition) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourcePositionInput(v *GetResourcePositionInput, encoder *httpbinding.Encoder) error {
@@ -4583,10 +4417,8 @@ func (m *awsRestjson1_serializeOpGetServiceEndpoint) H
 func (m *awsRestjson1_serializeOpGetServiceEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4624,7 +4456,6 @@ func (m *awsRestjson1_serializeOpGetServiceEndpoint) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetServiceEndpointInput(v *GetServiceEndpointInput, encoder *httpbinding.Encoder) error {
@@ -4649,10 +4480,8 @@ func (m *awsRestjson1_serializeOpGetServiceProfile) Ha
 func (m *awsRestjson1_serializeOpGetServiceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4690,7 +4519,6 @@ func (m *awsRestjson1_serializeOpGetServiceProfile) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetServiceProfileInput(v *GetServiceProfileInput, encoder *httpbinding.Encoder) error {
@@ -4720,10 +4548,8 @@ func (m *awsRestjson1_serializeOpGetWirelessDevice) Ha
 func (m *awsRestjson1_serializeOpGetWirelessDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4761,7 +4587,6 @@ func (m *awsRestjson1_serializeOpGetWirelessDevice) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessDeviceInput(v *GetWirelessDeviceInput, encoder *httpbinding.Encoder) error {
@@ -4795,10 +4620,8 @@ func (m *awsRestjson1_serializeOpGetWirelessDeviceImpo
 func (m *awsRestjson1_serializeOpGetWirelessDeviceImportTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4836,7 +4659,6 @@ func (m *awsRestjson1_serializeOpGetWirelessDeviceImpo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessDeviceImportTaskInput(v *GetWirelessDeviceImportTaskInput, encoder *httpbinding.Encoder) error {
@@ -4866,10 +4688,8 @@ func (m *awsRestjson1_serializeOpGetWirelessDeviceStat
 func (m *awsRestjson1_serializeOpGetWirelessDeviceStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4907,7 +4727,6 @@ func (m *awsRestjson1_serializeOpGetWirelessDeviceStat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessDeviceStatisticsInput(v *GetWirelessDeviceStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -4937,10 +4756,8 @@ func (m *awsRestjson1_serializeOpGetWirelessGateway) H
 func (m *awsRestjson1_serializeOpGetWirelessGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4978,7 +4795,6 @@ func (m *awsRestjson1_serializeOpGetWirelessGateway) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessGatewayInput(v *GetWirelessGatewayInput, encoder *httpbinding.Encoder) error {
@@ -5012,10 +4828,8 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayCer
 func (m *awsRestjson1_serializeOpGetWirelessGatewayCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5053,7 +4867,6 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayCer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessGatewayCertificateInput(v *GetWirelessGatewayCertificateInput, encoder *httpbinding.Encoder) error {
@@ -5083,10 +4896,8 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayFir
 func (m *awsRestjson1_serializeOpGetWirelessGatewayFirmwareInformation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5124,7 +4935,6 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayFir
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessGatewayFirmwareInformationInput(v *GetWirelessGatewayFirmwareInformationInput, encoder *httpbinding.Encoder) error {
@@ -5154,10 +4964,8 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewaySta
 func (m *awsRestjson1_serializeOpGetWirelessGatewayStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5195,7 +5003,6 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewaySta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessGatewayStatisticsInput(v *GetWirelessGatewayStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -5225,10 +5032,8 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayTas
 func (m *awsRestjson1_serializeOpGetWirelessGatewayTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5266,7 +5071,6 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayTas
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessGatewayTaskInput(v *GetWirelessGatewayTaskInput, encoder *httpbinding.Encoder) error {
@@ -5296,10 +5100,8 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayTas
 func (m *awsRestjson1_serializeOpGetWirelessGatewayTaskDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5337,7 +5139,6 @@ func (m *awsRestjson1_serializeOpGetWirelessGatewayTas
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWirelessGatewayTaskDefinitionInput(v *GetWirelessGatewayTaskDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -5367,10 +5168,8 @@ func (m *awsRestjson1_serializeOpListDestinations) Han
 func (m *awsRestjson1_serializeOpListDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5408,7 +5207,6 @@ func (m *awsRestjson1_serializeOpListDestinations) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDestinationsInput(v *ListDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -5437,10 +5235,8 @@ func (m *awsRestjson1_serializeOpListDeviceProfiles) H
 func (m *awsRestjson1_serializeOpListDeviceProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5478,7 +5274,6 @@ func (m *awsRestjson1_serializeOpListDeviceProfiles) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDeviceProfilesInput(v *ListDeviceProfilesInput, encoder *httpbinding.Encoder) error {
@@ -5511,10 +5306,8 @@ func (m *awsRestjson1_serializeOpListDevicesForWireles
 func (m *awsRestjson1_serializeOpListDevicesForWirelessDeviceImportTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5552,7 +5345,6 @@ func (m *awsRestjson1_serializeOpListDevicesForWireles
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDevicesForWirelessDeviceImportTaskInput(v *ListDevicesForWirelessDeviceImportTaskInput, encoder *httpbinding.Encoder) error {
@@ -5589,10 +5381,8 @@ func (m *awsRestjson1_serializeOpListEventConfiguratio
 func (m *awsRestjson1_serializeOpListEventConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5630,7 +5420,6 @@ func (m *awsRestjson1_serializeOpListEventConfiguratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventConfigurationsInput(v *ListEventConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5663,10 +5452,8 @@ func (m *awsRestjson1_serializeOpListFuotaTasks) Handl
 func (m *awsRestjson1_serializeOpListFuotaTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5704,7 +5491,6 @@ func (m *awsRestjson1_serializeOpListFuotaTasks) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFuotaTasksInput(v *ListFuotaTasksInput, encoder *httpbinding.Encoder) error {
@@ -5733,10 +5519,8 @@ func (m *awsRestjson1_serializeOpListMulticastGroups) 
 func (m *awsRestjson1_serializeOpListMulticastGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5774,7 +5558,6 @@ func (m *awsRestjson1_serializeOpListMulticastGroups) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMulticastGroupsInput(v *ListMulticastGroupsInput, encoder *httpbinding.Encoder) error {
@@ -5803,10 +5586,8 @@ func (m *awsRestjson1_serializeOpListMulticastGroupsBy
 func (m *awsRestjson1_serializeOpListMulticastGroupsByFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5844,7 +5625,6 @@ func (m *awsRestjson1_serializeOpListMulticastGroupsBy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMulticastGroupsByFuotaTaskInput(v *ListMulticastGroupsByFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -5882,10 +5662,8 @@ func (m *awsRestjson1_serializeOpListNetworkAnalyzerCo
 func (m *awsRestjson1_serializeOpListNetworkAnalyzerConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5923,7 +5701,6 @@ func (m *awsRestjson1_serializeOpListNetworkAnalyzerCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNetworkAnalyzerConfigurationsInput(v *ListNetworkAnalyzerConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5952,10 +5729,8 @@ func (m *awsRestjson1_serializeOpListPartnerAccounts) 
 func (m *awsRestjson1_serializeOpListPartnerAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5993,7 +5768,6 @@ func (m *awsRestjson1_serializeOpListPartnerAccounts) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPartnerAccountsInput(v *ListPartnerAccountsInput, encoder *httpbinding.Encoder) error {
@@ -6022,10 +5796,8 @@ func (m *awsRestjson1_serializeOpListPositionConfigura
 func (m *awsRestjson1_serializeOpListPositionConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6063,7 +5835,6 @@ func (m *awsRestjson1_serializeOpListPositionConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPositionConfigurationsInput(v *ListPositionConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -6096,10 +5867,8 @@ func (m *awsRestjson1_serializeOpListQueuedMessages) H
 func (m *awsRestjson1_serializeOpListQueuedMessages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6137,7 +5906,6 @@ func (m *awsRestjson1_serializeOpListQueuedMessages) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueuedMessagesInput(v *ListQueuedMessagesInput, encoder *httpbinding.Encoder) error {
@@ -6179,10 +5947,8 @@ func (m *awsRestjson1_serializeOpListServiceProfiles) 
 func (m *awsRestjson1_serializeOpListServiceProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6220,7 +5986,6 @@ func (m *awsRestjson1_serializeOpListServiceProfiles) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListServiceProfilesInput(v *ListServiceProfilesInput, encoder *httpbinding.Encoder) error {
@@ -6249,10 +6014,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6290,7 +6053,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6315,10 +6077,8 @@ func (m *awsRestjson1_serializeOpListWirelessDeviceImp
 func (m *awsRestjson1_serializeOpListWirelessDeviceImportTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6356,7 +6116,6 @@ func (m *awsRestjson1_serializeOpListWirelessDeviceImp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWirelessDeviceImportTasksInput(v *ListWirelessDeviceImportTasksInput, encoder *httpbinding.Encoder) error {
@@ -6385,10 +6144,8 @@ func (m *awsRestjson1_serializeOpListWirelessDevices) 
 func (m *awsRestjson1_serializeOpListWirelessDevices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6426,7 +6183,6 @@ func (m *awsRestjson1_serializeOpListWirelessDevices) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWirelessDevicesInput(v *ListWirelessDevicesInput, encoder *httpbinding.Encoder) error {
@@ -6479,10 +6235,8 @@ func (m *awsRestjson1_serializeOpListWirelessGateways)
 func (m *awsRestjson1_serializeOpListWirelessGateways) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6520,7 +6274,6 @@ func (m *awsRestjson1_serializeOpListWirelessGateways)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWirelessGatewaysInput(v *ListWirelessGatewaysInput, encoder *httpbinding.Encoder) error {
@@ -6549,10 +6302,8 @@ func (m *awsRestjson1_serializeOpListWirelessGatewayTa
 func (m *awsRestjson1_serializeOpListWirelessGatewayTaskDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6590,7 +6341,6 @@ func (m *awsRestjson1_serializeOpListWirelessGatewayTa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWirelessGatewayTaskDefinitionsInput(v *ListWirelessGatewayTaskDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -6623,10 +6373,8 @@ func (m *awsRestjson1_serializeOpPutPositionConfigurat
 func (m *awsRestjson1_serializeOpPutPositionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6675,7 +6423,6 @@ func (m *awsRestjson1_serializeOpPutPositionConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutPositionConfigurationInput(v *PutPositionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6728,10 +6475,8 @@ func (m *awsRestjson1_serializeOpPutResourceLogLevel) 
 func (m *awsRestjson1_serializeOpPutResourceLogLevel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6780,7 +6525,6 @@ func (m *awsRestjson1_serializeOpPutResourceLogLevel) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutResourceLogLevelInput(v *PutResourceLogLevelInput, encoder *httpbinding.Encoder) error {
@@ -6826,10 +6570,8 @@ func (m *awsRestjson1_serializeOpResetAllResourceLogLe
 func (m *awsRestjson1_serializeOpResetAllResourceLogLevels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6863,7 +6605,6 @@ func (m *awsRestjson1_serializeOpResetAllResourceLogLe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResetAllResourceLogLevelsInput(v *ResetAllResourceLogLevelsInput, encoder *httpbinding.Encoder) error {
@@ -6884,10 +6625,8 @@ func (m *awsRestjson1_serializeOpResetResourceLogLevel
 func (m *awsRestjson1_serializeOpResetResourceLogLevel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6925,7 +6664,6 @@ func (m *awsRestjson1_serializeOpResetResourceLogLevel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResetResourceLogLevelInput(v *ResetResourceLogLevelInput, encoder *httpbinding.Encoder) error {
@@ -6959,10 +6697,8 @@ func (m *awsRestjson1_serializeOpSendDataToMulticastGr
 func (m *awsRestjson1_serializeOpSendDataToMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7011,7 +6747,6 @@ func (m *awsRestjson1_serializeOpSendDataToMulticastGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendDataToMulticastGroupInput(v *SendDataToMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -7060,10 +6795,8 @@ func (m *awsRestjson1_serializeOpSendDataToWirelessDev
 func (m *awsRestjson1_serializeOpSendDataToWirelessDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7112,7 +6845,6 @@ func (m *awsRestjson1_serializeOpSendDataToWirelessDev
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendDataToWirelessDeviceInput(v *SendDataToWirelessDeviceInput, encoder *httpbinding.Encoder) error {
@@ -7166,10 +6898,8 @@ func (m *awsRestjson1_serializeOpStartBulkAssociateWir
 func (m *awsRestjson1_serializeOpStartBulkAssociateWirelessDeviceWithMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7218,7 +6948,6 @@ func (m *awsRestjson1_serializeOpStartBulkAssociateWir
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartBulkAssociateWirelessDeviceWithMulticastGroupInput(v *StartBulkAssociateWirelessDeviceWithMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -7267,10 +6996,8 @@ func (m *awsRestjson1_serializeOpStartBulkDisassociate
 func (m *awsRestjson1_serializeOpStartBulkDisassociateWirelessDeviceFromMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7319,7 +7046,6 @@ func (m *awsRestjson1_serializeOpStartBulkDisassociate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartBulkDisassociateWirelessDeviceFromMulticastGroupInput(v *StartBulkDisassociateWirelessDeviceFromMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -7368,10 +7094,8 @@ func (m *awsRestjson1_serializeOpStartFuotaTask) Handl
 func (m *awsRestjson1_serializeOpStartFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7420,7 +7144,6 @@ func (m *awsRestjson1_serializeOpStartFuotaTask) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartFuotaTaskInput(v *StartFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -7464,10 +7187,8 @@ func (m *awsRestjson1_serializeOpStartMulticastGroupSe
 func (m *awsRestjson1_serializeOpStartMulticastGroupSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7516,7 +7237,6 @@ func (m *awsRestjson1_serializeOpStartMulticastGroupSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMulticastGroupSessionInput(v *StartMulticastGroupSessionInput, encoder *httpbinding.Encoder) error {
@@ -7560,10 +7280,8 @@ func (m *awsRestjson1_serializeOpStartSingleWirelessDe
 func (m *awsRestjson1_serializeOpStartSingleWirelessDeviceImportTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7608,7 +7326,6 @@ func (m *awsRestjson1_serializeOpStartSingleWirelessDe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartSingleWirelessDeviceImportTaskInput(v *StartSingleWirelessDeviceImportTaskInput, encoder *httpbinding.Encoder) error {
@@ -7665,10 +7382,8 @@ func (m *awsRestjson1_serializeOpStartWirelessDeviceIm
 func (m *awsRestjson1_serializeOpStartWirelessDeviceImportTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7713,7 +7428,6 @@ func (m *awsRestjson1_serializeOpStartWirelessDeviceIm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartWirelessDeviceImportTaskInput(v *StartWirelessDeviceImportTaskInput, encoder *httpbinding.Encoder) error {
@@ -7765,10 +7479,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7817,7 +7529,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7856,10 +7567,8 @@ func (m *awsRestjson1_serializeOpTestWirelessDevice) H
 func (m *awsRestjson1_serializeOpTestWirelessDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7897,7 +7606,6 @@ func (m *awsRestjson1_serializeOpTestWirelessDevice) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestWirelessDeviceInput(v *TestWirelessDeviceInput, encoder *httpbinding.Encoder) error {
@@ -7927,10 +7635,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7968,7 +7674,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7999,10 +7704,8 @@ func (m *awsRestjson1_serializeOpUpdateDestination) Ha
 func (m *awsRestjson1_serializeOpUpdateDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8051,7 +7754,6 @@ func (m *awsRestjson1_serializeOpUpdateDestination) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDestinationInput(v *UpdateDestinationInput, encoder *httpbinding.Encoder) error {
@@ -8108,10 +7810,8 @@ func (m *awsRestjson1_serializeOpUpdateEventConfigurat
 func (m *awsRestjson1_serializeOpUpdateEventConfigurationByResourceTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8156,7 +7856,6 @@ func (m *awsRestjson1_serializeOpUpdateEventConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventConfigurationByResourceTypesInput(v *UpdateEventConfigurationByResourceTypesInput, encoder *httpbinding.Encoder) error {
@@ -8219,10 +7918,8 @@ func (m *awsRestjson1_serializeOpUpdateFuotaTask) Hand
 func (m *awsRestjson1_serializeOpUpdateFuotaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8271,7 +7968,6 @@ func (m *awsRestjson1_serializeOpUpdateFuotaTask) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFuotaTaskInput(v *UpdateFuotaTaskInput, encoder *httpbinding.Encoder) error {
@@ -8355,10 +8051,8 @@ func (m *awsRestjson1_serializeOpUpdateLogLevelsByReso
 func (m *awsRestjson1_serializeOpUpdateLogLevelsByResourceTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8403,7 +8097,6 @@ func (m *awsRestjson1_serializeOpUpdateLogLevelsByReso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLogLevelsByResourceTypesInput(v *UpdateLogLevelsByResourceTypesInput, encoder *httpbinding.Encoder) error {
@@ -8457,10 +8150,8 @@ func (m *awsRestjson1_serializeOpUpdateMetricConfigura
 func (m *awsRestjson1_serializeOpUpdateMetricConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8505,7 +8196,6 @@ func (m *awsRestjson1_serializeOpUpdateMetricConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMetricConfigurationInput(v *UpdateMetricConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8540,10 +8230,8 @@ func (m *awsRestjson1_serializeOpUpdateMulticastGroup)
 func (m *awsRestjson1_serializeOpUpdateMulticastGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8592,7 +8280,6 @@ func (m *awsRestjson1_serializeOpUpdateMulticastGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMulticastGroupInput(v *UpdateMulticastGroupInput, encoder *httpbinding.Encoder) error {
@@ -8646,10 +8333,8 @@ func (m *awsRestjson1_serializeOpUpdateNetworkAnalyzer
 func (m *awsRestjson1_serializeOpUpdateNetworkAnalyzerConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8698,7 +8383,6 @@ func (m *awsRestjson1_serializeOpUpdateNetworkAnalyzer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateNetworkAnalyzerConfigurationInput(v *UpdateNetworkAnalyzerConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8789,10 +8473,8 @@ func (m *awsRestjson1_serializeOpUpdatePartnerAccount)
 func (m *awsRestjson1_serializeOpUpdatePartnerAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8841,7 +8523,6 @@ func (m *awsRestjson1_serializeOpUpdatePartnerAccount)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePartnerAccountInput(v *UpdatePartnerAccountInput, encoder *httpbinding.Encoder) error {
@@ -8889,10 +8570,8 @@ func (m *awsRestjson1_serializeOpUpdatePosition) Handl
 func (m *awsRestjson1_serializeOpUpdatePosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8941,7 +8620,6 @@ func (m *awsRestjson1_serializeOpUpdatePosition) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePositionInput(v *UpdatePositionInput, encoder *httpbinding.Encoder) error {
@@ -8989,10 +8667,8 @@ func (m *awsRestjson1_serializeOpUpdateResourceEventCo
 func (m *awsRestjson1_serializeOpUpdateResourceEventConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9041,7 +8717,6 @@ func (m *awsRestjson1_serializeOpUpdateResourceEventCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceEventConfigurationInput(v *UpdateResourceEventConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -9121,10 +8796,8 @@ func (m *awsRestjson1_serializeOpUpdateResourcePositio
 func (m *awsRestjson1_serializeOpUpdateResourcePosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9174,7 +8847,6 @@ func (m *awsRestjson1_serializeOpUpdateResourcePositio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourcePositionInput(v *UpdateResourcePositionInput, encoder *httpbinding.Encoder) error {
@@ -9208,10 +8880,8 @@ func (m *awsRestjson1_serializeOpUpdateWirelessDevice)
 func (m *awsRestjson1_serializeOpUpdateWirelessDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9260,7 +8930,6 @@ func (m *awsRestjson1_serializeOpUpdateWirelessDevice)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWirelessDeviceInput(v *UpdateWirelessDeviceInput, encoder *httpbinding.Encoder) error {
@@ -9324,10 +8993,8 @@ func (m *awsRestjson1_serializeOpUpdateWirelessDeviceI
 func (m *awsRestjson1_serializeOpUpdateWirelessDeviceImportTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9376,7 +9043,6 @@ func (m *awsRestjson1_serializeOpUpdateWirelessDeviceI
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWirelessDeviceImportTaskInput(v *UpdateWirelessDeviceImportTaskInput, encoder *httpbinding.Encoder) error {
@@ -9420,10 +9086,8 @@ func (m *awsRestjson1_serializeOpUpdateWirelessGateway
 func (m *awsRestjson1_serializeOpUpdateWirelessGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9472,7 +9136,6 @@ func (m *awsRestjson1_serializeOpUpdateWirelessGateway
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWirelessGatewayInput(v *UpdateWirelessGatewayInput, encoder *httpbinding.Encoder) error {
