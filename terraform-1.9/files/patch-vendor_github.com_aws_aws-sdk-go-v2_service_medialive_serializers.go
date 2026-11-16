--- vendor/github.com/aws/aws-sdk-go-v2/service/medialive/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/medialive/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAcceptInputDeviceTran
 func (m *awsRestjson1_serializeOpAcceptInputDeviceTransfer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_serializeOpAcceptInputDeviceTran
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptInputDeviceTransferInput(v *AcceptInputDeviceTransferInput, encoder *httpbinding.Encoder) error {
@@ -97,10 +93,8 @@ func (m *awsRestjson1_serializeOpBatchDelete) HandleSe
 func (m *awsRestjson1_serializeOpBatchDelete) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -145,7 +139,6 @@ func (m *awsRestjson1_serializeOpBatchDelete) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteInput(v *BatchDeleteInput, encoder *httpbinding.Encoder) error {
@@ -201,10 +194,8 @@ func (m *awsRestjson1_serializeOpBatchStart) HandleSer
 func (m *awsRestjson1_serializeOpBatchStart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -249,7 +240,6 @@ func (m *awsRestjson1_serializeOpBatchStart) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchStartInput(v *BatchStartInput, encoder *httpbinding.Encoder) error {
@@ -291,10 +281,8 @@ func (m *awsRestjson1_serializeOpBatchStop) HandleSeri
 func (m *awsRestjson1_serializeOpBatchStop) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -339,7 +327,6 @@ func (m *awsRestjson1_serializeOpBatchStop) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchStopInput(v *BatchStopInput, encoder *httpbinding.Encoder) error {
@@ -381,10 +368,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateSchedule) 
 func (m *awsRestjson1_serializeOpBatchUpdateSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -433,7 +418,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateSchedule) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateScheduleInput(v *BatchUpdateScheduleInput, encoder *httpbinding.Encoder) error {
@@ -484,10 +468,8 @@ func (m *awsRestjson1_serializeOpCancelInputDeviceTran
 func (m *awsRestjson1_serializeOpCancelInputDeviceTransfer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -525,7 +507,6 @@ func (m *awsRestjson1_serializeOpCancelInputDeviceTran
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelInputDeviceTransferInput(v *CancelInputDeviceTransferInput, encoder *httpbinding.Encoder) error {
@@ -555,10 +536,8 @@ func (m *awsRestjson1_serializeOpClaimDevice) HandleSe
 func (m *awsRestjson1_serializeOpClaimDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -603,7 +582,6 @@ func (m *awsRestjson1_serializeOpClaimDevice) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsClaimDeviceInput(v *ClaimDeviceInput, encoder *httpbinding.Encoder) error {
@@ -636,10 +614,8 @@ func (m *awsRestjson1_serializeOpCreateChannel) Handle
 func (m *awsRestjson1_serializeOpCreateChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -684,7 +660,6 @@ func (m *awsRestjson1_serializeOpCreateChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChannelInput(v *CreateChannelInput, encoder *httpbinding.Encoder) error {
@@ -817,10 +792,8 @@ func (m *awsRestjson1_serializeOpCreateChannelPlacemen
 func (m *awsRestjson1_serializeOpCreateChannelPlacementGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -869,7 +842,6 @@ func (m *awsRestjson1_serializeOpCreateChannelPlacemen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChannelPlacementGroupInput(v *CreateChannelPlacementGroupInput, encoder *httpbinding.Encoder) error {
@@ -930,10 +902,8 @@ func (m *awsRestjson1_serializeOpCreateCloudWatchAlarm
 func (m *awsRestjson1_serializeOpCreateCloudWatchAlarmTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -978,7 +948,6 @@ func (m *awsRestjson1_serializeOpCreateCloudWatchAlarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCloudWatchAlarmTemplateInput(v *CreateCloudWatchAlarmTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1091,10 +1060,8 @@ func (m *awsRestjson1_serializeOpCreateCloudWatchAlarm
 func (m *awsRestjson1_serializeOpCreateCloudWatchAlarmTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1139,7 +1106,6 @@ func (m *awsRestjson1_serializeOpCreateCloudWatchAlarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCloudWatchAlarmTemplateGroupInput(v *CreateCloudWatchAlarmTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -1189,10 +1155,8 @@ func (m *awsRestjson1_serializeOpCreateCluster) Handle
 func (m *awsRestjson1_serializeOpCreateCluster) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1237,7 +1201,6 @@ func (m *awsRestjson1_serializeOpCreateCluster) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateClusterInput(v *CreateClusterInput, encoder *httpbinding.Encoder) error {
@@ -1299,10 +1262,8 @@ func (m *awsRestjson1_serializeOpCreateEventBridgeRule
 func (m *awsRestjson1_serializeOpCreateEventBridgeRuleTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1347,7 +1308,6 @@ func (m *awsRestjson1_serializeOpCreateEventBridgeRule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEventBridgeRuleTemplateInput(v *CreateEventBridgeRuleTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1414,10 +1374,8 @@ func (m *awsRestjson1_serializeOpCreateEventBridgeRule
 func (m *awsRestjson1_serializeOpCreateEventBridgeRuleTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1462,7 +1420,6 @@ func (m *awsRestjson1_serializeOpCreateEventBridgeRule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEventBridgeRuleTemplateGroupInput(v *CreateEventBridgeRuleTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -1512,10 +1469,8 @@ func (m *awsRestjson1_serializeOpCreateInput) HandleSe
 func (m *awsRestjson1_serializeOpCreateInput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1560,7 +1515,6 @@ func (m *awsRestjson1_serializeOpCreateInput) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInputInput(v *CreateInputInput, encoder *httpbinding.Encoder) error {
@@ -1690,10 +1644,8 @@ func (m *awsRestjson1_serializeOpCreateInputSecurityGr
 func (m *awsRestjson1_serializeOpCreateInputSecurityGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1738,7 +1690,6 @@ func (m *awsRestjson1_serializeOpCreateInputSecurityGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInputSecurityGroupInput(v *CreateInputSecurityGroupInput, encoder *httpbinding.Encoder) error {
@@ -1780,10 +1731,8 @@ func (m *awsRestjson1_serializeOpCreateMultiplex) Hand
 func (m *awsRestjson1_serializeOpCreateMultiplex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1828,7 +1777,6 @@ func (m *awsRestjson1_serializeOpCreateMultiplex) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMultiplexInput(v *CreateMultiplexInput, encoder *httpbinding.Encoder) error {
@@ -1887,10 +1835,8 @@ func (m *awsRestjson1_serializeOpCreateMultiplexProgra
 func (m *awsRestjson1_serializeOpCreateMultiplexProgram) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1939,7 +1885,6 @@ func (m *awsRestjson1_serializeOpCreateMultiplexProgra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMultiplexProgramInput(v *CreateMultiplexProgramInput, encoder *httpbinding.Encoder) error {
@@ -1993,10 +1938,8 @@ func (m *awsRestjson1_serializeOpCreateNetwork) Handle
 func (m *awsRestjson1_serializeOpCreateNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2041,7 +1984,6 @@ func (m *awsRestjson1_serializeOpCreateNetwork) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNetworkInput(v *CreateNetworkInput, encoder *httpbinding.Encoder) error {
@@ -2100,10 +2042,8 @@ func (m *awsRestjson1_serializeOpCreateNode) HandleSer
 func (m *awsRestjson1_serializeOpCreateNode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2152,7 +2092,6 @@ func (m *awsRestjson1_serializeOpCreateNode) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNodeInput(v *CreateNodeInput, encoder *httpbinding.Encoder) error {
@@ -2218,10 +2157,8 @@ func (m *awsRestjson1_serializeOpCreateNodeRegistratio
 func (m *awsRestjson1_serializeOpCreateNodeRegistrationScript) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2270,7 +2207,6 @@ func (m *awsRestjson1_serializeOpCreateNodeRegistratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNodeRegistrationScriptInput(v *CreateNodeRegistrationScriptInput, encoder *httpbinding.Encoder) error {
@@ -2334,10 +2270,8 @@ func (m *awsRestjson1_serializeOpCreatePartnerInput) H
 func (m *awsRestjson1_serializeOpCreatePartnerInput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2386,7 +2320,6 @@ func (m *awsRestjson1_serializeOpCreatePartnerInput) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePartnerInputInput(v *CreatePartnerInputInput, encoder *httpbinding.Encoder) error {
@@ -2435,10 +2368,8 @@ func (m *awsRestjson1_serializeOpCreateSdiSource) Hand
 func (m *awsRestjson1_serializeOpCreateSdiSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2483,7 +2414,6 @@ func (m *awsRestjson1_serializeOpCreateSdiSource) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSdiSourceInput(v *CreateSdiSourceInput, encoder *httpbinding.Encoder) error {
@@ -2538,10 +2468,8 @@ func (m *awsRestjson1_serializeOpCreateSignalMap) Hand
 func (m *awsRestjson1_serializeOpCreateSignalMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2586,7 +2514,6 @@ func (m *awsRestjson1_serializeOpCreateSignalMap) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSignalMapInput(v *CreateSignalMapInput, encoder *httpbinding.Encoder) error {
@@ -2655,10 +2582,8 @@ func (m *awsRestjson1_serializeOpCreateTags) HandleSer
 func (m *awsRestjson1_serializeOpCreateTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2707,7 +2632,6 @@ func (m *awsRestjson1_serializeOpCreateTags) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTagsInput(v *CreateTagsInput, encoder *httpbinding.Encoder) error {
@@ -2751,10 +2675,8 @@ func (m *awsRestjson1_serializeOpDeleteChannel) Handle
 func (m *awsRestjson1_serializeOpDeleteChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2792,7 +2714,6 @@ func (m *awsRestjson1_serializeOpDeleteChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelInput(v *DeleteChannelInput, encoder *httpbinding.Encoder) error {
@@ -2822,10 +2743,8 @@ func (m *awsRestjson1_serializeOpDeleteChannelPlacemen
 func (m *awsRestjson1_serializeOpDeleteChannelPlacementGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2863,7 +2782,6 @@ func (m *awsRestjson1_serializeOpDeleteChannelPlacemen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelPlacementGroupInput(v *DeleteChannelPlacementGroupInput, encoder *httpbinding.Encoder) error {
@@ -2902,10 +2820,8 @@ func (m *awsRestjson1_serializeOpDeleteCloudWatchAlarm
 func (m *awsRestjson1_serializeOpDeleteCloudWatchAlarmTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2943,7 +2859,6 @@ func (m *awsRestjson1_serializeOpDeleteCloudWatchAlarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCloudWatchAlarmTemplateInput(v *DeleteCloudWatchAlarmTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2973,10 +2888,8 @@ func (m *awsRestjson1_serializeOpDeleteCloudWatchAlarm
 func (m *awsRestjson1_serializeOpDeleteCloudWatchAlarmTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3014,7 +2927,6 @@ func (m *awsRestjson1_serializeOpDeleteCloudWatchAlarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCloudWatchAlarmTemplateGroupInput(v *DeleteCloudWatchAlarmTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -3044,10 +2956,8 @@ func (m *awsRestjson1_serializeOpDeleteCluster) Handle
 func (m *awsRestjson1_serializeOpDeleteCluster) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3085,7 +2995,6 @@ func (m *awsRestjson1_serializeOpDeleteCluster) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteClusterInput(v *DeleteClusterInput, encoder *httpbinding.Encoder) error {
@@ -3115,10 +3024,8 @@ func (m *awsRestjson1_serializeOpDeleteEventBridgeRule
 func (m *awsRestjson1_serializeOpDeleteEventBridgeRuleTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3156,7 +3063,6 @@ func (m *awsRestjson1_serializeOpDeleteEventBridgeRule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventBridgeRuleTemplateInput(v *DeleteEventBridgeRuleTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3186,10 +3092,8 @@ func (m *awsRestjson1_serializeOpDeleteEventBridgeRule
 func (m *awsRestjson1_serializeOpDeleteEventBridgeRuleTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3227,7 +3131,6 @@ func (m *awsRestjson1_serializeOpDeleteEventBridgeRule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventBridgeRuleTemplateGroupInput(v *DeleteEventBridgeRuleTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -3257,10 +3160,8 @@ func (m *awsRestjson1_serializeOpDeleteInput) HandleSe
 func (m *awsRestjson1_serializeOpDeleteInput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3298,7 +3199,6 @@ func (m *awsRestjson1_serializeOpDeleteInput) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInputInput(v *DeleteInputInput, encoder *httpbinding.Encoder) error {
@@ -3328,10 +3228,8 @@ func (m *awsRestjson1_serializeOpDeleteInputSecurityGr
 func (m *awsRestjson1_serializeOpDeleteInputSecurityGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3369,7 +3267,6 @@ func (m *awsRestjson1_serializeOpDeleteInputSecurityGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInputSecurityGroupInput(v *DeleteInputSecurityGroupInput, encoder *httpbinding.Encoder) error {
@@ -3399,10 +3296,8 @@ func (m *awsRestjson1_serializeOpDeleteMultiplex) Hand
 func (m *awsRestjson1_serializeOpDeleteMultiplex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3440,7 +3335,6 @@ func (m *awsRestjson1_serializeOpDeleteMultiplex) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMultiplexInput(v *DeleteMultiplexInput, encoder *httpbinding.Encoder) error {
@@ -3470,10 +3364,8 @@ func (m *awsRestjson1_serializeOpDeleteMultiplexProgra
 func (m *awsRestjson1_serializeOpDeleteMultiplexProgram) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3511,7 +3403,6 @@ func (m *awsRestjson1_serializeOpDeleteMultiplexProgra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMultiplexProgramInput(v *DeleteMultiplexProgramInput, encoder *httpbinding.Encoder) error {
@@ -3550,10 +3441,8 @@ func (m *awsRestjson1_serializeOpDeleteNetwork) Handle
 func (m *awsRestjson1_serializeOpDeleteNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3591,7 +3480,6 @@ func (m *awsRestjson1_serializeOpDeleteNetwork) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteNetworkInput(v *DeleteNetworkInput, encoder *httpbinding.Encoder) error {
@@ -3621,10 +3509,8 @@ func (m *awsRestjson1_serializeOpDeleteNode) HandleSer
 func (m *awsRestjson1_serializeOpDeleteNode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3662,7 +3548,6 @@ func (m *awsRestjson1_serializeOpDeleteNode) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteNodeInput(v *DeleteNodeInput, encoder *httpbinding.Encoder) error {
@@ -3701,10 +3586,8 @@ func (m *awsRestjson1_serializeOpDeleteReservation) Ha
 func (m *awsRestjson1_serializeOpDeleteReservation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3742,7 +3625,6 @@ func (m *awsRestjson1_serializeOpDeleteReservation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteReservationInput(v *DeleteReservationInput, encoder *httpbinding.Encoder) error {
@@ -3772,10 +3654,8 @@ func (m *awsRestjson1_serializeOpDeleteSchedule) Handl
 func (m *awsRestjson1_serializeOpDeleteSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3813,7 +3693,6 @@ func (m *awsRestjson1_serializeOpDeleteSchedule) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteScheduleInput(v *DeleteScheduleInput, encoder *httpbinding.Encoder) error {
@@ -3843,10 +3722,8 @@ func (m *awsRestjson1_serializeOpDeleteSdiSource) Hand
 func (m *awsRestjson1_serializeOpDeleteSdiSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3884,7 +3761,6 @@ func (m *awsRestjson1_serializeOpDeleteSdiSource) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSdiSourceInput(v *DeleteSdiSourceInput, encoder *httpbinding.Encoder) error {
@@ -3914,10 +3790,8 @@ func (m *awsRestjson1_serializeOpDeleteSignalMap) Hand
 func (m *awsRestjson1_serializeOpDeleteSignalMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3955,7 +3829,6 @@ func (m *awsRestjson1_serializeOpDeleteSignalMap) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSignalMapInput(v *DeleteSignalMapInput, encoder *httpbinding.Encoder) error {
@@ -3985,10 +3858,8 @@ func (m *awsRestjson1_serializeOpDeleteTags) HandleSer
 func (m *awsRestjson1_serializeOpDeleteTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4026,7 +3897,6 @@ func (m *awsRestjson1_serializeOpDeleteTags) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTagsInput(v *DeleteTagsInput, encoder *httpbinding.Encoder) error {
@@ -4062,10 +3932,8 @@ func (m *awsRestjson1_serializeOpDescribeAccountConfig
 func (m *awsRestjson1_serializeOpDescribeAccountConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4099,7 +3967,6 @@ func (m *awsRestjson1_serializeOpDescribeAccountConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccountConfigurationInput(v *DescribeAccountConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4120,10 +3987,8 @@ func (m *awsRestjson1_serializeOpDescribeChannel) Hand
 func (m *awsRestjson1_serializeOpDescribeChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4161,7 +4026,6 @@ func (m *awsRestjson1_serializeOpDescribeChannel) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelInput(v *DescribeChannelInput, encoder *httpbinding.Encoder) error {
@@ -4191,10 +4055,8 @@ func (m *awsRestjson1_serializeOpDescribeChannelPlacem
 func (m *awsRestjson1_serializeOpDescribeChannelPlacementGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4232,7 +4094,6 @@ func (m *awsRestjson1_serializeOpDescribeChannelPlacem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelPlacementGroupInput(v *DescribeChannelPlacementGroupInput, encoder *httpbinding.Encoder) error {
@@ -4271,10 +4132,8 @@ func (m *awsRestjson1_serializeOpDescribeCluster) Hand
 func (m *awsRestjson1_serializeOpDescribeCluster) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4312,7 +4171,6 @@ func (m *awsRestjson1_serializeOpDescribeCluster) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeClusterInput(v *DescribeClusterInput, encoder *httpbinding.Encoder) error {
@@ -4342,10 +4200,8 @@ func (m *awsRestjson1_serializeOpDescribeInput) Handle
 func (m *awsRestjson1_serializeOpDescribeInput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4383,7 +4239,6 @@ func (m *awsRestjson1_serializeOpDescribeInput) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInputInput(v *DescribeInputInput, encoder *httpbinding.Encoder) error {
@@ -4413,10 +4268,8 @@ func (m *awsRestjson1_serializeOpDescribeInputDevice) 
 func (m *awsRestjson1_serializeOpDescribeInputDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4454,7 +4307,6 @@ func (m *awsRestjson1_serializeOpDescribeInputDevice) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInputDeviceInput(v *DescribeInputDeviceInput, encoder *httpbinding.Encoder) error {
@@ -4484,10 +4336,8 @@ func (m *awsRestjson1_serializeOpDescribeInputDeviceTh
 func (m *awsRestjson1_serializeOpDescribeInputDeviceThumbnail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4525,7 +4375,6 @@ func (m *awsRestjson1_serializeOpDescribeInputDeviceTh
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInputDeviceThumbnailInput(v *DescribeInputDeviceThumbnailInput, encoder *httpbinding.Encoder) error {
@@ -4560,10 +4409,8 @@ func (m *awsRestjson1_serializeOpDescribeInputSecurity
 func (m *awsRestjson1_serializeOpDescribeInputSecurityGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4601,7 +4448,6 @@ func (m *awsRestjson1_serializeOpDescribeInputSecurity
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInputSecurityGroupInput(v *DescribeInputSecurityGroupInput, encoder *httpbinding.Encoder) error {
@@ -4631,10 +4477,8 @@ func (m *awsRestjson1_serializeOpDescribeMultiplex) Ha
 func (m *awsRestjson1_serializeOpDescribeMultiplex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4672,7 +4516,6 @@ func (m *awsRestjson1_serializeOpDescribeMultiplex) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeMultiplexInput(v *DescribeMultiplexInput, encoder *httpbinding.Encoder) error {
@@ -4702,10 +4545,8 @@ func (m *awsRestjson1_serializeOpDescribeMultiplexProg
 func (m *awsRestjson1_serializeOpDescribeMultiplexProgram) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4743,7 +4584,6 @@ func (m *awsRestjson1_serializeOpDescribeMultiplexProg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeMultiplexProgramInput(v *DescribeMultiplexProgramInput, encoder *httpbinding.Encoder) error {
@@ -4782,10 +4622,8 @@ func (m *awsRestjson1_serializeOpDescribeNetwork) Hand
 func (m *awsRestjson1_serializeOpDescribeNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4823,7 +4661,6 @@ func (m *awsRestjson1_serializeOpDescribeNetwork) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeNetworkInput(v *DescribeNetworkInput, encoder *httpbinding.Encoder) error {
@@ -4853,10 +4690,8 @@ func (m *awsRestjson1_serializeOpDescribeNode) HandleS
 func (m *awsRestjson1_serializeOpDescribeNode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4894,7 +4729,6 @@ func (m *awsRestjson1_serializeOpDescribeNode) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeNodeInput(v *DescribeNodeInput, encoder *httpbinding.Encoder) error {
@@ -4933,10 +4767,8 @@ func (m *awsRestjson1_serializeOpDescribeOffering) Han
 func (m *awsRestjson1_serializeOpDescribeOffering) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4974,7 +4806,6 @@ func (m *awsRestjson1_serializeOpDescribeOffering) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOfferingInput(v *DescribeOfferingInput, encoder *httpbinding.Encoder) error {
@@ -5004,10 +4835,8 @@ func (m *awsRestjson1_serializeOpDescribeReservation) 
 func (m *awsRestjson1_serializeOpDescribeReservation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5045,7 +4874,6 @@ func (m *awsRestjson1_serializeOpDescribeReservation) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReservationInput(v *DescribeReservationInput, encoder *httpbinding.Encoder) error {
@@ -5075,10 +4903,8 @@ func (m *awsRestjson1_serializeOpDescribeSchedule) Han
 func (m *awsRestjson1_serializeOpDescribeSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5116,7 +4942,6 @@ func (m *awsRestjson1_serializeOpDescribeSchedule) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeScheduleInput(v *DescribeScheduleInput, encoder *httpbinding.Encoder) error {
@@ -5154,10 +4979,8 @@ func (m *awsRestjson1_serializeOpDescribeSdiSource) Ha
 func (m *awsRestjson1_serializeOpDescribeSdiSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5195,7 +5018,6 @@ func (m *awsRestjson1_serializeOpDescribeSdiSource) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSdiSourceInput(v *DescribeSdiSourceInput, encoder *httpbinding.Encoder) error {
@@ -5225,10 +5047,8 @@ func (m *awsRestjson1_serializeOpDescribeThumbnails) H
 func (m *awsRestjson1_serializeOpDescribeThumbnails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5266,7 +5086,6 @@ func (m *awsRestjson1_serializeOpDescribeThumbnails) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThumbnailsInput(v *DescribeThumbnailsInput, encoder *httpbinding.Encoder) error {
@@ -5304,10 +5123,8 @@ func (m *awsRestjson1_serializeOpGetCloudWatchAlarmTem
 func (m *awsRestjson1_serializeOpGetCloudWatchAlarmTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5345,7 +5162,6 @@ func (m *awsRestjson1_serializeOpGetCloudWatchAlarmTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCloudWatchAlarmTemplateInput(v *GetCloudWatchAlarmTemplateInput, encoder *httpbinding.Encoder) error {
@@ -5375,10 +5191,8 @@ func (m *awsRestjson1_serializeOpGetCloudWatchAlarmTem
 func (m *awsRestjson1_serializeOpGetCloudWatchAlarmTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5416,7 +5230,6 @@ func (m *awsRestjson1_serializeOpGetCloudWatchAlarmTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCloudWatchAlarmTemplateGroupInput(v *GetCloudWatchAlarmTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -5446,10 +5259,8 @@ func (m *awsRestjson1_serializeOpGetEventBridgeRuleTem
 func (m *awsRestjson1_serializeOpGetEventBridgeRuleTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5487,7 +5298,6 @@ func (m *awsRestjson1_serializeOpGetEventBridgeRuleTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventBridgeRuleTemplateInput(v *GetEventBridgeRuleTemplateInput, encoder *httpbinding.Encoder) error {
@@ -5517,10 +5327,8 @@ func (m *awsRestjson1_serializeOpGetEventBridgeRuleTem
 func (m *awsRestjson1_serializeOpGetEventBridgeRuleTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5558,7 +5366,6 @@ func (m *awsRestjson1_serializeOpGetEventBridgeRuleTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventBridgeRuleTemplateGroupInput(v *GetEventBridgeRuleTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -5588,10 +5395,8 @@ func (m *awsRestjson1_serializeOpGetSignalMap) HandleS
 func (m *awsRestjson1_serializeOpGetSignalMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5629,7 +5434,6 @@ func (m *awsRestjson1_serializeOpGetSignalMap) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSignalMapInput(v *GetSignalMapInput, encoder *httpbinding.Encoder) error {
@@ -5659,10 +5463,8 @@ func (m *awsRestjson1_serializeOpListChannelPlacementG
 func (m *awsRestjson1_serializeOpListChannelPlacementGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5700,7 +5502,6 @@ func (m *awsRestjson1_serializeOpListChannelPlacementG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelPlacementGroupsInput(v *ListChannelPlacementGroupsInput, encoder *httpbinding.Encoder) error {
@@ -5738,10 +5539,8 @@ func (m *awsRestjson1_serializeOpListChannels) HandleS
 func (m *awsRestjson1_serializeOpListChannels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5779,7 +5578,6 @@ func (m *awsRestjson1_serializeOpListChannels) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelsInput(v *ListChannelsInput, encoder *httpbinding.Encoder) error {
@@ -5808,10 +5606,8 @@ func (m *awsRestjson1_serializeOpListCloudWatchAlarmTe
 func (m *awsRestjson1_serializeOpListCloudWatchAlarmTemplateGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5849,7 +5645,6 @@ func (m *awsRestjson1_serializeOpListCloudWatchAlarmTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCloudWatchAlarmTemplateGroupsInput(v *ListCloudWatchAlarmTemplateGroupsInput, encoder *httpbinding.Encoder) error {
@@ -5886,10 +5681,8 @@ func (m *awsRestjson1_serializeOpListCloudWatchAlarmTe
 func (m *awsRestjson1_serializeOpListCloudWatchAlarmTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5927,7 +5720,6 @@ func (m *awsRestjson1_serializeOpListCloudWatchAlarmTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCloudWatchAlarmTemplatesInput(v *ListCloudWatchAlarmTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -5968,10 +5760,8 @@ func (m *awsRestjson1_serializeOpListClusters) HandleS
 func (m *awsRestjson1_serializeOpListClusters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6009,7 +5799,6 @@ func (m *awsRestjson1_serializeOpListClusters) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListClustersInput(v *ListClustersInput, encoder *httpbinding.Encoder) error {
@@ -6038,10 +5827,8 @@ func (m *awsRestjson1_serializeOpListEventBridgeRuleTe
 func (m *awsRestjson1_serializeOpListEventBridgeRuleTemplateGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6079,7 +5866,6 @@ func (m *awsRestjson1_serializeOpListEventBridgeRuleTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventBridgeRuleTemplateGroupsInput(v *ListEventBridgeRuleTemplateGroupsInput, encoder *httpbinding.Encoder) error {
@@ -6112,10 +5898,8 @@ func (m *awsRestjson1_serializeOpListEventBridgeRuleTe
 func (m *awsRestjson1_serializeOpListEventBridgeRuleTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6153,7 +5937,6 @@ func (m *awsRestjson1_serializeOpListEventBridgeRuleTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventBridgeRuleTemplatesInput(v *ListEventBridgeRuleTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -6190,10 +5973,8 @@ func (m *awsRestjson1_serializeOpListInputDevices) Han
 func (m *awsRestjson1_serializeOpListInputDevices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6231,7 +6012,6 @@ func (m *awsRestjson1_serializeOpListInputDevices) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInputDevicesInput(v *ListInputDevicesInput, encoder *httpbinding.Encoder) error {
@@ -6260,10 +6040,8 @@ func (m *awsRestjson1_serializeOpListInputDeviceTransf
 func (m *awsRestjson1_serializeOpListInputDeviceTransfers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6301,7 +6079,6 @@ func (m *awsRestjson1_serializeOpListInputDeviceTransf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInputDeviceTransfersInput(v *ListInputDeviceTransfersInput, encoder *httpbinding.Encoder) error {
@@ -6334,10 +6111,8 @@ func (m *awsRestjson1_serializeOpListInputs) HandleSer
 func (m *awsRestjson1_serializeOpListInputs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6375,7 +6150,6 @@ func (m *awsRestjson1_serializeOpListInputs) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInputsInput(v *ListInputsInput, encoder *httpbinding.Encoder) error {
@@ -6404,10 +6178,8 @@ func (m *awsRestjson1_serializeOpListInputSecurityGrou
 func (m *awsRestjson1_serializeOpListInputSecurityGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6445,7 +6217,6 @@ func (m *awsRestjson1_serializeOpListInputSecurityGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInputSecurityGroupsInput(v *ListInputSecurityGroupsInput, encoder *httpbinding.Encoder) error {
@@ -6474,10 +6245,8 @@ func (m *awsRestjson1_serializeOpListMultiplexes) Hand
 func (m *awsRestjson1_serializeOpListMultiplexes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6515,7 +6284,6 @@ func (m *awsRestjson1_serializeOpListMultiplexes) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMultiplexesInput(v *ListMultiplexesInput, encoder *httpbinding.Encoder) error {
@@ -6544,10 +6312,8 @@ func (m *awsRestjson1_serializeOpListMultiplexPrograms
 func (m *awsRestjson1_serializeOpListMultiplexPrograms) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6585,7 +6351,6 @@ func (m *awsRestjson1_serializeOpListMultiplexPrograms
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMultiplexProgramsInput(v *ListMultiplexProgramsInput, encoder *httpbinding.Encoder) error {
@@ -6623,10 +6388,8 @@ func (m *awsRestjson1_serializeOpListNetworks) HandleS
 func (m *awsRestjson1_serializeOpListNetworks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6664,7 +6427,6 @@ func (m *awsRestjson1_serializeOpListNetworks) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNetworksInput(v *ListNetworksInput, encoder *httpbinding.Encoder) error {
@@ -6693,10 +6455,8 @@ func (m *awsRestjson1_serializeOpListNodes) HandleSeri
 func (m *awsRestjson1_serializeOpListNodes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6734,7 +6494,6 @@ func (m *awsRestjson1_serializeOpListNodes) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNodesInput(v *ListNodesInput, encoder *httpbinding.Encoder) error {
@@ -6772,10 +6531,8 @@ func (m *awsRestjson1_serializeOpListOfferings) Handle
 func (m *awsRestjson1_serializeOpListOfferings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6813,7 +6570,6 @@ func (m *awsRestjson1_serializeOpListOfferings) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOfferingsInput(v *ListOfferingsInput, encoder *httpbinding.Encoder) error {
@@ -6882,10 +6638,8 @@ func (m *awsRestjson1_serializeOpListReservations) Han
 func (m *awsRestjson1_serializeOpListReservations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6923,7 +6677,6 @@ func (m *awsRestjson1_serializeOpListReservations) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReservationsInput(v *ListReservationsInput, encoder *httpbinding.Encoder) error {
@@ -6984,10 +6737,8 @@ func (m *awsRestjson1_serializeOpListSdiSources) Handl
 func (m *awsRestjson1_serializeOpListSdiSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7025,7 +6776,6 @@ func (m *awsRestjson1_serializeOpListSdiSources) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSdiSourcesInput(v *ListSdiSourcesInput, encoder *httpbinding.Encoder) error {
@@ -7054,10 +6804,8 @@ func (m *awsRestjson1_serializeOpListSignalMaps) Handl
 func (m *awsRestjson1_serializeOpListSignalMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7095,7 +6843,6 @@ func (m *awsRestjson1_serializeOpListSignalMaps) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSignalMapsInput(v *ListSignalMapsInput, encoder *httpbinding.Encoder) error {
@@ -7132,10 +6879,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -7173,7 +6918,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -7203,10 +6947,8 @@ func (m *awsRestjson1_serializeOpListVersions) HandleS
 func (m *awsRestjson1_serializeOpListVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7240,7 +6982,6 @@ func (m *awsRestjson1_serializeOpListVersions) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVersionsInput(v *ListVersionsInput, encoder *httpbinding.Encoder) error {
@@ -7261,10 +7002,8 @@ func (m *awsRestjson1_serializeOpPurchaseOffering) Han
 func (m *awsRestjson1_serializeOpPurchaseOffering) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7313,7 +7052,6 @@ func (m *awsRestjson1_serializeOpPurchaseOffering) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPurchaseOfferingInput(v *PurchaseOfferingInput, encoder *httpbinding.Encoder) error {
@@ -7384,10 +7122,8 @@ func (m *awsRestjson1_serializeOpRebootInputDevice) Ha
 func (m *awsRestjson1_serializeOpRebootInputDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7436,7 +7172,6 @@ func (m *awsRestjson1_serializeOpRebootInputDevice) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRebootInputDeviceInput(v *RebootInputDeviceInput, encoder *httpbinding.Encoder) error {
@@ -7478,10 +7213,8 @@ func (m *awsRestjson1_serializeOpRejectInputDeviceTran
 func (m *awsRestjson1_serializeOpRejectInputDeviceTransfer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7519,7 +7252,6 @@ func (m *awsRestjson1_serializeOpRejectInputDeviceTran
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectInputDeviceTransferInput(v *RejectInputDeviceTransferInput, encoder *httpbinding.Encoder) error {
@@ -7549,10 +7281,8 @@ func (m *awsRestjson1_serializeOpRestartChannelPipelin
 func (m *awsRestjson1_serializeOpRestartChannelPipelines) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7601,7 +7331,6 @@ func (m *awsRestjson1_serializeOpRestartChannelPipelin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRestartChannelPipelinesInput(v *RestartChannelPipelinesInput, encoder *httpbinding.Encoder) error {
@@ -7645,10 +7374,8 @@ func (m *awsRestjson1_serializeOpStartChannel) HandleS
 func (m *awsRestjson1_serializeOpStartChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7686,7 +7413,6 @@ func (m *awsRestjson1_serializeOpStartChannel) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartChannelInput(v *StartChannelInput, encoder *httpbinding.Encoder) error {
@@ -7716,10 +7442,8 @@ func (m *awsRestjson1_serializeOpStartDeleteMonitorDep
 func (m *awsRestjson1_serializeOpStartDeleteMonitorDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7757,7 +7481,6 @@ func (m *awsRestjson1_serializeOpStartDeleteMonitorDep
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDeleteMonitorDeploymentInput(v *StartDeleteMonitorDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -7787,10 +7510,8 @@ func (m *awsRestjson1_serializeOpStartInputDevice) Han
 func (m *awsRestjson1_serializeOpStartInputDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7828,7 +7549,6 @@ func (m *awsRestjson1_serializeOpStartInputDevice) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartInputDeviceInput(v *StartInputDeviceInput, encoder *httpbinding.Encoder) error {
@@ -7858,10 +7578,8 @@ func (m *awsRestjson1_serializeOpStartInputDeviceMaint
 func (m *awsRestjson1_serializeOpStartInputDeviceMaintenanceWindow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7899,7 +7617,6 @@ func (m *awsRestjson1_serializeOpStartInputDeviceMaint
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartInputDeviceMaintenanceWindowInput(v *StartInputDeviceMaintenanceWindowInput, encoder *httpbinding.Encoder) error {
@@ -7929,10 +7646,8 @@ func (m *awsRestjson1_serializeOpStartMonitorDeploymen
 func (m *awsRestjson1_serializeOpStartMonitorDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7981,7 +7696,6 @@ func (m *awsRestjson1_serializeOpStartMonitorDeploymen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMonitorDeploymentInput(v *StartMonitorDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -8023,10 +7737,8 @@ func (m *awsRestjson1_serializeOpStartMultiplex) Handl
 func (m *awsRestjson1_serializeOpStartMultiplex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8064,7 +7776,6 @@ func (m *awsRestjson1_serializeOpStartMultiplex) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMultiplexInput(v *StartMultiplexInput, encoder *httpbinding.Encoder) error {
@@ -8094,10 +7805,8 @@ func (m *awsRestjson1_serializeOpStartUpdateSignalMap)
 func (m *awsRestjson1_serializeOpStartUpdateSignalMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8146,7 +7855,6 @@ func (m *awsRestjson1_serializeOpStartUpdateSignalMap)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartUpdateSignalMapInput(v *StartUpdateSignalMapInput, encoder *httpbinding.Encoder) error {
@@ -8217,10 +7925,8 @@ func (m *awsRestjson1_serializeOpStopChannel) HandleSe
 func (m *awsRestjson1_serializeOpStopChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8258,7 +7964,6 @@ func (m *awsRestjson1_serializeOpStopChannel) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopChannelInput(v *StopChannelInput, encoder *httpbinding.Encoder) error {
@@ -8288,10 +7993,8 @@ func (m *awsRestjson1_serializeOpStopInputDevice) Hand
 func (m *awsRestjson1_serializeOpStopInputDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8329,7 +8032,6 @@ func (m *awsRestjson1_serializeOpStopInputDevice) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopInputDeviceInput(v *StopInputDeviceInput, encoder *httpbinding.Encoder) error {
@@ -8359,10 +8061,8 @@ func (m *awsRestjson1_serializeOpStopMultiplex) Handle
 func (m *awsRestjson1_serializeOpStopMultiplex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8400,7 +8100,6 @@ func (m *awsRestjson1_serializeOpStopMultiplex) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopMultiplexInput(v *StopMultiplexInput, encoder *httpbinding.Encoder) error {
@@ -8430,10 +8129,8 @@ func (m *awsRestjson1_serializeOpTransferInputDevice) 
 func (m *awsRestjson1_serializeOpTransferInputDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8482,7 +8179,6 @@ func (m *awsRestjson1_serializeOpTransferInputDevice) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTransferInputDeviceInput(v *TransferInputDeviceInput, encoder *httpbinding.Encoder) error {
@@ -8534,10 +8230,8 @@ func (m *awsRestjson1_serializeOpUpdateAccountConfigur
 func (m *awsRestjson1_serializeOpUpdateAccountConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8582,7 +8276,6 @@ func (m *awsRestjson1_serializeOpUpdateAccountConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountConfigurationInput(v *UpdateAccountConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8617,10 +8310,8 @@ func (m *awsRestjson1_serializeOpUpdateChannel) Handle
 func (m *awsRestjson1_serializeOpUpdateChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8669,7 +8360,6 @@ func (m *awsRestjson1_serializeOpUpdateChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChannelInput(v *UpdateChannelInput, encoder *httpbinding.Encoder) error {
@@ -8782,10 +8472,8 @@ func (m *awsRestjson1_serializeOpUpdateChannelClass) H
 func (m *awsRestjson1_serializeOpUpdateChannelClass) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8834,7 +8522,6 @@ func (m *awsRestjson1_serializeOpUpdateChannelClass) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChannelClassInput(v *UpdateChannelClassInput, encoder *httpbinding.Encoder) error {
@@ -8883,10 +8570,8 @@ func (m *awsRestjson1_serializeOpUpdateChannelPlacemen
 func (m *awsRestjson1_serializeOpUpdateChannelPlacementGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8935,7 +8620,6 @@ func (m *awsRestjson1_serializeOpUpdateChannelPlacemen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChannelPlacementGroupInput(v *UpdateChannelPlacementGroupInput, encoder *httpbinding.Encoder) error {
@@ -8993,10 +8677,8 @@ func (m *awsRestjson1_serializeOpUpdateCloudWatchAlarm
 func (m *awsRestjson1_serializeOpUpdateCloudWatchAlarmTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9045,7 +8727,6 @@ func (m *awsRestjson1_serializeOpUpdateCloudWatchAlarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCloudWatchAlarmTemplateInput(v *UpdateCloudWatchAlarmTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9155,10 +8836,8 @@ func (m *awsRestjson1_serializeOpUpdateCloudWatchAlarm
 func (m *awsRestjson1_serializeOpUpdateCloudWatchAlarmTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9207,7 +8886,6 @@ func (m *awsRestjson1_serializeOpUpdateCloudWatchAlarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCloudWatchAlarmTemplateGroupInput(v *UpdateCloudWatchAlarmTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -9249,10 +8927,8 @@ func (m *awsRestjson1_serializeOpUpdateCluster) Handle
 func (m *awsRestjson1_serializeOpUpdateCluster) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9301,7 +8977,6 @@ func (m *awsRestjson1_serializeOpUpdateCluster) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateClusterInput(v *UpdateClusterInput, encoder *httpbinding.Encoder) error {
@@ -9350,10 +9025,8 @@ func (m *awsRestjson1_serializeOpUpdateEventBridgeRule
 func (m *awsRestjson1_serializeOpUpdateEventBridgeRuleTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9402,7 +9075,6 @@ func (m *awsRestjson1_serializeOpUpdateEventBridgeRule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventBridgeRuleTemplateInput(v *UpdateEventBridgeRuleTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9466,10 +9138,8 @@ func (m *awsRestjson1_serializeOpUpdateEventBridgeRule
 func (m *awsRestjson1_serializeOpUpdateEventBridgeRuleTemplateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9518,7 +9188,6 @@ func (m *awsRestjson1_serializeOpUpdateEventBridgeRule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventBridgeRuleTemplateGroupInput(v *UpdateEventBridgeRuleTemplateGroupInput, encoder *httpbinding.Encoder) error {
@@ -9560,10 +9229,8 @@ func (m *awsRestjson1_serializeOpUpdateInput) HandleSe
 func (m *awsRestjson1_serializeOpUpdateInput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9612,7 +9279,6 @@ func (m *awsRestjson1_serializeOpUpdateInput) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInputInput(v *UpdateInputInput, encoder *httpbinding.Encoder) error {
@@ -9722,10 +9388,8 @@ func (m *awsRestjson1_serializeOpUpdateInputDevice) Ha
 func (m *awsRestjson1_serializeOpUpdateInputDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9774,7 +9438,6 @@ func (m *awsRestjson1_serializeOpUpdateInputDevice) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInputDeviceInput(v *UpdateInputDeviceInput, encoder *httpbinding.Encoder) error {
@@ -9835,10 +9498,8 @@ func (m *awsRestjson1_serializeOpUpdateInputSecurityGr
 func (m *awsRestjson1_serializeOpUpdateInputSecurityGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9887,7 +9548,6 @@ func (m *awsRestjson1_serializeOpUpdateInputSecurityGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInputSecurityGroupInput(v *UpdateInputSecurityGroupInput, encoder *httpbinding.Encoder) error {
@@ -9938,10 +9598,8 @@ func (m *awsRestjson1_serializeOpUpdateMultiplex) Hand
 func (m *awsRestjson1_serializeOpUpdateMultiplex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9990,7 +9648,6 @@ func (m *awsRestjson1_serializeOpUpdateMultiplex) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMultiplexInput(v *UpdateMultiplexInput, encoder *httpbinding.Encoder) error {
@@ -10046,10 +9703,8 @@ func (m *awsRestjson1_serializeOpUpdateMultiplexProgra
 func (m *awsRestjson1_serializeOpUpdateMultiplexProgram) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10098,7 +9753,6 @@ func (m *awsRestjson1_serializeOpUpdateMultiplexProgra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMultiplexProgramInput(v *UpdateMultiplexProgramInput, encoder *httpbinding.Encoder) error {
@@ -10151,10 +9805,8 @@ func (m *awsRestjson1_serializeOpUpdateNetwork) Handle
 func (m *awsRestjson1_serializeOpUpdateNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10203,7 +9855,6 @@ func (m *awsRestjson1_serializeOpUpdateNetwork) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateNetworkInput(v *UpdateNetworkInput, encoder *httpbinding.Encoder) error {
@@ -10259,10 +9910,8 @@ func (m *awsRestjson1_serializeOpUpdateNode) HandleSer
 func (m *awsRestjson1_serializeOpUpdateNode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10311,7 +9960,6 @@ func (m *awsRestjson1_serializeOpUpdateNode) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateNodeInput(v *UpdateNodeInput, encoder *httpbinding.Encoder) error {
@@ -10374,10 +10022,8 @@ func (m *awsRestjson1_serializeOpUpdateNodeState) Hand
 func (m *awsRestjson1_serializeOpUpdateNodeState) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10426,7 +10072,6 @@ func (m *awsRestjson1_serializeOpUpdateNodeState) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateNodeStateInput(v *UpdateNodeStateInput, encoder *httpbinding.Encoder) error {
@@ -10477,10 +10122,8 @@ func (m *awsRestjson1_serializeOpUpdateReservation) Ha
 func (m *awsRestjson1_serializeOpUpdateReservation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10529,7 +10172,6 @@ func (m *awsRestjson1_serializeOpUpdateReservation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateReservationInput(v *UpdateReservationInput, encoder *httpbinding.Encoder) error {
@@ -10578,10 +10220,8 @@ func (m *awsRestjson1_serializeOpUpdateSdiSource) Hand
 func (m *awsRestjson1_serializeOpUpdateSdiSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10630,7 +10270,6 @@ func (m *awsRestjson1_serializeOpUpdateSdiSource) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSdiSourceInput(v *UpdateSdiSourceInput, encoder *httpbinding.Encoder) error {
