--- vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAssociateChannelFlow)
 func (m *awsRestjson1_serializeOpAssociateChannelFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -78,7 +75,6 @@ func (m *awsRestjson1_serializeOpAssociateChannelFlow)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateChannelFlowInput(v *AssociateChannelFlowInput, encoder *httpbinding.Encoder) error {
@@ -125,10 +121,8 @@ func (m *awsRestjson1_serializeOpBatchCreateChannelMem
 func (m *awsRestjson1_serializeOpBatchCreateChannelMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -177,7 +171,6 @@ func (m *awsRestjson1_serializeOpBatchCreateChannelMem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchCreateChannelMembershipInput(v *BatchCreateChannelMembershipInput, encoder *httpbinding.Encoder) error {
@@ -236,10 +229,8 @@ func (m *awsRestjson1_serializeOpChannelFlowCallback) 
 func (m *awsRestjson1_serializeOpChannelFlowCallback) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -288,7 +279,6 @@ func (m *awsRestjson1_serializeOpChannelFlowCallback) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsChannelFlowCallbackInput(v *ChannelFlowCallbackInput, encoder *httpbinding.Encoder) error {
@@ -342,10 +332,8 @@ func (m *awsRestjson1_serializeOpCreateChannel) Handle
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
@@ -394,7 +382,6 @@ func (m *awsRestjson1_serializeOpCreateChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChannelInput(v *CreateChannelInput, encoder *httpbinding.Encoder) error {
@@ -497,10 +484,8 @@ func (m *awsRestjson1_serializeOpCreateChannelBan) Han
 func (m *awsRestjson1_serializeOpCreateChannelBan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -549,7 +534,6 @@ func (m *awsRestjson1_serializeOpCreateChannelBan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChannelBanInput(v *CreateChannelBanInput, encoder *httpbinding.Encoder) error {
@@ -596,10 +580,8 @@ func (m *awsRestjson1_serializeOpCreateChannelFlow) Ha
 func (m *awsRestjson1_serializeOpCreateChannelFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -644,7 +626,6 @@ func (m *awsRestjson1_serializeOpCreateChannelFlow) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChannelFlowInput(v *CreateChannelFlowInput, encoder *httpbinding.Encoder) error {
@@ -701,10 +682,8 @@ func (m *awsRestjson1_serializeOpCreateChannelMembersh
 func (m *awsRestjson1_serializeOpCreateChannelMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -753,7 +732,6 @@ func (m *awsRestjson1_serializeOpCreateChannelMembersh
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChannelMembershipInput(v *CreateChannelMembershipInput, encoder *httpbinding.Encoder) error {
@@ -810,10 +788,8 @@ func (m *awsRestjson1_serializeOpCreateChannelModerato
 func (m *awsRestjson1_serializeOpCreateChannelModerator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -862,7 +838,6 @@ func (m *awsRestjson1_serializeOpCreateChannelModerato
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChannelModeratorInput(v *CreateChannelModeratorInput, encoder *httpbinding.Encoder) error {
@@ -909,10 +884,8 @@ func (m *awsRestjson1_serializeOpDeleteChannel) Handle
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
@@ -950,7 +923,6 @@ func (m *awsRestjson1_serializeOpDeleteChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelInput(v *DeleteChannelInput, encoder *httpbinding.Encoder) error {
@@ -985,10 +957,8 @@ func (m *awsRestjson1_serializeOpDeleteChannelBan) Han
 func (m *awsRestjson1_serializeOpDeleteChannelBan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1026,7 +996,6 @@ func (m *awsRestjson1_serializeOpDeleteChannelBan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelBanInput(v *DeleteChannelBanInput, encoder *httpbinding.Encoder) error {
@@ -1070,10 +1039,8 @@ func (m *awsRestjson1_serializeOpDeleteChannelFlow) Ha
 func (m *awsRestjson1_serializeOpDeleteChannelFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1111,7 +1078,6 @@ func (m *awsRestjson1_serializeOpDeleteChannelFlow) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelFlowInput(v *DeleteChannelFlowInput, encoder *httpbinding.Encoder) error {
@@ -1141,10 +1107,8 @@ func (m *awsRestjson1_serializeOpDeleteChannelMembersh
 func (m *awsRestjson1_serializeOpDeleteChannelMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1182,7 +1146,6 @@ func (m *awsRestjson1_serializeOpDeleteChannelMembersh
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelMembershipInput(v *DeleteChannelMembershipInput, encoder *httpbinding.Encoder) error {
@@ -1230,10 +1193,8 @@ func (m *awsRestjson1_serializeOpDeleteChannelMessage)
 func (m *awsRestjson1_serializeOpDeleteChannelMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1271,7 +1232,6 @@ func (m *awsRestjson1_serializeOpDeleteChannelMessage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelMessageInput(v *DeleteChannelMessageInput, encoder *httpbinding.Encoder) error {
@@ -1319,10 +1279,8 @@ func (m *awsRestjson1_serializeOpDeleteChannelModerato
 func (m *awsRestjson1_serializeOpDeleteChannelModerator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1360,7 +1318,6 @@ func (m *awsRestjson1_serializeOpDeleteChannelModerato
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChannelModeratorInput(v *DeleteChannelModeratorInput, encoder *httpbinding.Encoder) error {
@@ -1404,10 +1361,8 @@ func (m *awsRestjson1_serializeOpDeleteMessagingStream
 func (m *awsRestjson1_serializeOpDeleteMessagingStreamingConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1445,7 +1400,6 @@ func (m *awsRestjson1_serializeOpDeleteMessagingStream
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMessagingStreamingConfigurationsInput(v *DeleteMessagingStreamingConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -1475,10 +1429,8 @@ func (m *awsRestjson1_serializeOpDescribeChannel) Hand
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
@@ -1516,7 +1468,6 @@ func (m *awsRestjson1_serializeOpDescribeChannel) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelInput(v *DescribeChannelInput, encoder *httpbinding.Encoder) error {
@@ -1551,10 +1502,8 @@ func (m *awsRestjson1_serializeOpDescribeChannelBan) H
 func (m *awsRestjson1_serializeOpDescribeChannelBan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1592,7 +1541,6 @@ func (m *awsRestjson1_serializeOpDescribeChannelBan) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelBanInput(v *DescribeChannelBanInput, encoder *httpbinding.Encoder) error {
@@ -1636,10 +1584,8 @@ func (m *awsRestjson1_serializeOpDescribeChannelFlow) 
 func (m *awsRestjson1_serializeOpDescribeChannelFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1677,7 +1623,6 @@ func (m *awsRestjson1_serializeOpDescribeChannelFlow) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelFlowInput(v *DescribeChannelFlowInput, encoder *httpbinding.Encoder) error {
@@ -1707,10 +1652,8 @@ func (m *awsRestjson1_serializeOpDescribeChannelMember
 func (m *awsRestjson1_serializeOpDescribeChannelMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1748,7 +1691,6 @@ func (m *awsRestjson1_serializeOpDescribeChannelMember
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelMembershipInput(v *DescribeChannelMembershipInput, encoder *httpbinding.Encoder) error {
@@ -1796,10 +1738,8 @@ func (m *awsRestjson1_serializeOpDescribeChannelMember
 func (m *awsRestjson1_serializeOpDescribeChannelMembershipForAppInstanceUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1837,7 +1777,6 @@ func (m *awsRestjson1_serializeOpDescribeChannelMember
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelMembershipForAppInstanceUserInput(v *DescribeChannelMembershipForAppInstanceUserInput, encoder *httpbinding.Encoder) error {
@@ -1876,10 +1815,8 @@ func (m *awsRestjson1_serializeOpDescribeChannelModera
 func (m *awsRestjson1_serializeOpDescribeChannelModeratedByAppInstanceUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1917,7 +1854,6 @@ func (m *awsRestjson1_serializeOpDescribeChannelModera
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelModeratedByAppInstanceUserInput(v *DescribeChannelModeratedByAppInstanceUserInput, encoder *httpbinding.Encoder) error {
@@ -1956,10 +1892,8 @@ func (m *awsRestjson1_serializeOpDescribeChannelModera
 func (m *awsRestjson1_serializeOpDescribeChannelModerator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1997,7 +1931,6 @@ func (m *awsRestjson1_serializeOpDescribeChannelModera
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChannelModeratorInput(v *DescribeChannelModeratorInput, encoder *httpbinding.Encoder) error {
@@ -2041,10 +1974,8 @@ func (m *awsRestjson1_serializeOpDisassociateChannelFl
 func (m *awsRestjson1_serializeOpDisassociateChannelFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2082,7 +2013,6 @@ func (m *awsRestjson1_serializeOpDisassociateChannelFl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateChannelFlowInput(v *DisassociateChannelFlowInput, encoder *httpbinding.Encoder) error {
@@ -2126,10 +2056,8 @@ func (m *awsRestjson1_serializeOpGetChannelMembershipP
 func (m *awsRestjson1_serializeOpGetChannelMembershipPreferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2167,7 +2095,6 @@ func (m *awsRestjson1_serializeOpGetChannelMembershipP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetChannelMembershipPreferencesInput(v *GetChannelMembershipPreferencesInput, encoder *httpbinding.Encoder) error {
@@ -2211,10 +2138,8 @@ func (m *awsRestjson1_serializeOpGetChannelMessage) Ha
 func (m *awsRestjson1_serializeOpGetChannelMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2252,7 +2177,6 @@ func (m *awsRestjson1_serializeOpGetChannelMessage) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetChannelMessageInput(v *GetChannelMessageInput, encoder *httpbinding.Encoder) error {
@@ -2300,10 +2224,8 @@ func (m *awsRestjson1_serializeOpGetChannelMessageStat
 func (m *awsRestjson1_serializeOpGetChannelMessageStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2341,7 +2263,6 @@ func (m *awsRestjson1_serializeOpGetChannelMessageStat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetChannelMessageStatusInput(v *GetChannelMessageStatusInput, encoder *httpbinding.Encoder) error {
@@ -2389,10 +2310,8 @@ func (m *awsRestjson1_serializeOpGetMessagingSessionEn
 func (m *awsRestjson1_serializeOpGetMessagingSessionEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2426,7 +2345,6 @@ func (m *awsRestjson1_serializeOpGetMessagingSessionEn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMessagingSessionEndpointInput(v *GetMessagingSessionEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2447,10 +2365,8 @@ func (m *awsRestjson1_serializeOpGetMessagingStreaming
 func (m *awsRestjson1_serializeOpGetMessagingStreamingConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2488,7 +2404,6 @@ func (m *awsRestjson1_serializeOpGetMessagingStreaming
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMessagingStreamingConfigurationsInput(v *GetMessagingStreamingConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -2518,10 +2433,8 @@ func (m *awsRestjson1_serializeOpListChannelBans) Hand
 func (m *awsRestjson1_serializeOpListChannelBans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2559,7 +2472,6 @@ func (m *awsRestjson1_serializeOpListChannelBans) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelBansInput(v *ListChannelBansInput, encoder *httpbinding.Encoder) error {
@@ -2602,10 +2514,8 @@ func (m *awsRestjson1_serializeOpListChannelFlows) Han
 func (m *awsRestjson1_serializeOpListChannelFlows) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2643,7 +2553,6 @@ func (m *awsRestjson1_serializeOpListChannelFlows) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelFlowsInput(v *ListChannelFlowsInput, encoder *httpbinding.Encoder) error {
@@ -2676,10 +2585,8 @@ func (m *awsRestjson1_serializeOpListChannelMembership
 func (m *awsRestjson1_serializeOpListChannelMemberships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2717,7 +2624,6 @@ func (m *awsRestjson1_serializeOpListChannelMembership
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelMembershipsInput(v *ListChannelMembershipsInput, encoder *httpbinding.Encoder) error {
@@ -2768,10 +2674,8 @@ func (m *awsRestjson1_serializeOpListChannelMembership
 func (m *awsRestjson1_serializeOpListChannelMembershipsForAppInstanceUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2809,7 +2713,6 @@ func (m *awsRestjson1_serializeOpListChannelMembership
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelMembershipsForAppInstanceUserInput(v *ListChannelMembershipsForAppInstanceUserInput, encoder *httpbinding.Encoder) error {
@@ -2847,10 +2750,8 @@ func (m *awsRestjson1_serializeOpListChannelMessages) 
 func (m *awsRestjson1_serializeOpListChannelMessages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2888,7 +2789,6 @@ func (m *awsRestjson1_serializeOpListChannelMessages) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelMessagesInput(v *ListChannelMessagesInput, encoder *httpbinding.Encoder) error {
@@ -2947,10 +2847,8 @@ func (m *awsRestjson1_serializeOpListChannelModerators
 func (m *awsRestjson1_serializeOpListChannelModerators) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2988,7 +2886,6 @@ func (m *awsRestjson1_serializeOpListChannelModerators
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelModeratorsInput(v *ListChannelModeratorsInput, encoder *httpbinding.Encoder) error {
@@ -3031,10 +2928,8 @@ func (m *awsRestjson1_serializeOpListChannels) HandleS
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
@@ -3072,7 +2967,6 @@ func (m *awsRestjson1_serializeOpListChannels) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelsInput(v *ListChannelsInput, encoder *httpbinding.Encoder) error {
@@ -3114,10 +3008,8 @@ func (m *awsRestjson1_serializeOpListChannelsAssociate
 func (m *awsRestjson1_serializeOpListChannelsAssociatedWithChannelFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3155,7 +3047,6 @@ func (m *awsRestjson1_serializeOpListChannelsAssociate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelsAssociatedWithChannelFlowInput(v *ListChannelsAssociatedWithChannelFlowInput, encoder *httpbinding.Encoder) error {
@@ -3188,10 +3079,8 @@ func (m *awsRestjson1_serializeOpListChannelsModerated
 func (m *awsRestjson1_serializeOpListChannelsModeratedByAppInstanceUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3229,7 +3118,6 @@ func (m *awsRestjson1_serializeOpListChannelsModerated
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChannelsModeratedByAppInstanceUserInput(v *ListChannelsModeratedByAppInstanceUserInput, encoder *httpbinding.Encoder) error {
@@ -3267,10 +3155,8 @@ func (m *awsRestjson1_serializeOpListSubChannels) Hand
 func (m *awsRestjson1_serializeOpListSubChannels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3308,7 +3194,6 @@ func (m *awsRestjson1_serializeOpListSubChannels) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubChannelsInput(v *ListSubChannelsInput, encoder *httpbinding.Encoder) error {
@@ -3351,10 +3236,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -3392,7 +3275,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -3417,10 +3299,8 @@ func (m *awsRestjson1_serializeOpPutChannelExpirationS
 func (m *awsRestjson1_serializeOpPutChannelExpirationSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3469,7 +3349,6 @@ func (m *awsRestjson1_serializeOpPutChannelExpirationS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutChannelExpirationSettingsInput(v *PutChannelExpirationSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3518,10 +3397,8 @@ func (m *awsRestjson1_serializeOpPutChannelMembershipP
 func (m *awsRestjson1_serializeOpPutChannelMembershipPreferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3570,7 +3447,6 @@ func (m *awsRestjson1_serializeOpPutChannelMembershipP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutChannelMembershipPreferencesInput(v *PutChannelMembershipPreferencesInput, encoder *httpbinding.Encoder) error {
@@ -3628,10 +3504,8 @@ func (m *awsRestjson1_serializeOpPutMessagingStreaming
 func (m *awsRestjson1_serializeOpPutMessagingStreamingConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3680,7 +3554,6 @@ func (m *awsRestjson1_serializeOpPutMessagingStreaming
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutMessagingStreamingConfigurationsInput(v *PutMessagingStreamingConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -3724,10 +3597,8 @@ func (m *awsRestjson1_serializeOpRedactChannelMessage)
 func (m *awsRestjson1_serializeOpRedactChannelMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3776,7 +3647,6 @@ func (m *awsRestjson1_serializeOpRedactChannelMessage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRedactChannelMessageInput(v *RedactChannelMessageInput, encoder *httpbinding.Encoder) error {
@@ -3832,10 +3702,8 @@ func (m *awsRestjson1_serializeOpSearchChannels) Handl
 func (m *awsRestjson1_serializeOpSearchChannels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3884,7 +3752,6 @@ func (m *awsRestjson1_serializeOpSearchChannels) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchChannelsInput(v *SearchChannelsInput, encoder *httpbinding.Encoder) error {
@@ -3932,10 +3799,8 @@ func (m *awsRestjson1_serializeOpSendChannelMessage) H
 func (m *awsRestjson1_serializeOpSendChannelMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3984,7 +3849,6 @@ func (m *awsRestjson1_serializeOpSendChannelMessage) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendChannelMessageInput(v *SendChannelMessageInput, encoder *httpbinding.Encoder) error {
@@ -4082,10 +3946,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -4130,7 +3992,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4170,10 +4031,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -4218,7 +4077,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4258,10 +4116,8 @@ func (m *awsRestjson1_serializeOpUpdateChannel) Handle
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
@@ -4310,7 +4166,6 @@ func (m *awsRestjson1_serializeOpUpdateChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChannelInput(v *UpdateChannelInput, encoder *httpbinding.Encoder) error {
@@ -4367,10 +4222,8 @@ func (m *awsRestjson1_serializeOpUpdateChannelFlow) Ha
 func (m *awsRestjson1_serializeOpUpdateChannelFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4419,7 +4272,6 @@ func (m *awsRestjson1_serializeOpUpdateChannelFlow) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChannelFlowInput(v *UpdateChannelFlowInput, encoder *httpbinding.Encoder) error {
@@ -4468,10 +4320,8 @@ func (m *awsRestjson1_serializeOpUpdateChannelMessage)
 func (m *awsRestjson1_serializeOpUpdateChannelMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4520,7 +4370,6 @@ func (m *awsRestjson1_serializeOpUpdateChannelMessage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChannelMessageInput(v *UpdateChannelMessageInput, encoder *httpbinding.Encoder) error {
@@ -4591,10 +4440,8 @@ func (m *awsRestjson1_serializeOpUpdateChannelReadMark
 func (m *awsRestjson1_serializeOpUpdateChannelReadMarker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4632,7 +4479,6 @@ func (m *awsRestjson1_serializeOpUpdateChannelReadMark
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChannelReadMarkerInput(v *UpdateChannelReadMarkerInput, encoder *httpbinding.Encoder) error {
