--- vendor/github.com/aws/aws-sdk-go-v2/service/pinpoint/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/pinpoint/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpCreateApp) HandleSeri
 func (m *awsRestjson1_serializeOpCreateApp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_serializeOpCreateApp) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAppInput(v *CreateAppInput, encoder *httpbinding.Encoder) error {
@@ -109,10 +105,8 @@ func (m *awsRestjson1_serializeOpCreateCampaign) Handl
 func (m *awsRestjson1_serializeOpCreateCampaign) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -174,7 +168,6 @@ func (m *awsRestjson1_serializeOpCreateCampaign) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCampaignInput(v *CreateCampaignInput, encoder *httpbinding.Encoder) error {
@@ -204,10 +197,8 @@ func (m *awsRestjson1_serializeOpCreateEmailTemplate) 
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
@@ -269,7 +260,6 @@ func (m *awsRestjson1_serializeOpCreateEmailTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEmailTemplateInput(v *CreateEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -299,10 +289,8 @@ func (m *awsRestjson1_serializeOpCreateExportJob) Hand
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
@@ -364,7 +352,6 @@ func (m *awsRestjson1_serializeOpCreateExportJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateExportJobInput(v *CreateExportJobInput, encoder *httpbinding.Encoder) error {
@@ -394,10 +381,8 @@ func (m *awsRestjson1_serializeOpCreateImportJob) Hand
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
@@ -459,7 +444,6 @@ func (m *awsRestjson1_serializeOpCreateImportJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImportJobInput(v *CreateImportJobInput, encoder *httpbinding.Encoder) error {
@@ -489,10 +473,8 @@ func (m *awsRestjson1_serializeOpCreateInAppTemplate) 
 func (m *awsRestjson1_serializeOpCreateInAppTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -554,7 +536,6 @@ func (m *awsRestjson1_serializeOpCreateInAppTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInAppTemplateInput(v *CreateInAppTemplateInput, encoder *httpbinding.Encoder) error {
@@ -584,10 +565,8 @@ func (m *awsRestjson1_serializeOpCreateJourney) Handle
 func (m *awsRestjson1_serializeOpCreateJourney) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -649,7 +628,6 @@ func (m *awsRestjson1_serializeOpCreateJourney) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateJourneyInput(v *CreateJourneyInput, encoder *httpbinding.Encoder) error {
@@ -679,10 +657,8 @@ func (m *awsRestjson1_serializeOpCreatePushTemplate) H
 func (m *awsRestjson1_serializeOpCreatePushTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -744,7 +720,6 @@ func (m *awsRestjson1_serializeOpCreatePushTemplate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePushTemplateInput(v *CreatePushTemplateInput, encoder *httpbinding.Encoder) error {
@@ -774,10 +749,8 @@ func (m *awsRestjson1_serializeOpCreateRecommenderConf
 func (m *awsRestjson1_serializeOpCreateRecommenderConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -835,7 +808,6 @@ func (m *awsRestjson1_serializeOpCreateRecommenderConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRecommenderConfigurationInput(v *CreateRecommenderConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -856,10 +828,8 @@ func (m *awsRestjson1_serializeOpCreateSegment) Handle
 func (m *awsRestjson1_serializeOpCreateSegment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -921,7 +891,6 @@ func (m *awsRestjson1_serializeOpCreateSegment) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSegmentInput(v *CreateSegmentInput, encoder *httpbinding.Encoder) error {
@@ -951,10 +920,8 @@ func (m *awsRestjson1_serializeOpCreateSmsTemplate) Ha
 func (m *awsRestjson1_serializeOpCreateSmsTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1016,7 +983,6 @@ func (m *awsRestjson1_serializeOpCreateSmsTemplate) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSmsTemplateInput(v *CreateSmsTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1046,10 +1012,8 @@ func (m *awsRestjson1_serializeOpCreateVoiceTemplate) 
 func (m *awsRestjson1_serializeOpCreateVoiceTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1111,7 +1075,6 @@ func (m *awsRestjson1_serializeOpCreateVoiceTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVoiceTemplateInput(v *CreateVoiceTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1141,10 +1104,8 @@ func (m *awsRestjson1_serializeOpDeleteAdmChannel) Han
 func (m *awsRestjson1_serializeOpDeleteAdmChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1182,7 +1143,6 @@ func (m *awsRestjson1_serializeOpDeleteAdmChannel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAdmChannelInput(v *DeleteAdmChannelInput, encoder *httpbinding.Encoder) error {
@@ -1212,10 +1172,8 @@ func (m *awsRestjson1_serializeOpDeleteApnsChannel) Ha
 func (m *awsRestjson1_serializeOpDeleteApnsChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1253,7 +1211,6 @@ func (m *awsRestjson1_serializeOpDeleteApnsChannel) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteApnsChannelInput(v *DeleteApnsChannelInput, encoder *httpbinding.Encoder) error {
@@ -1283,10 +1240,8 @@ func (m *awsRestjson1_serializeOpDeleteApnsSandboxChan
 func (m *awsRestjson1_serializeOpDeleteApnsSandboxChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1324,7 +1279,6 @@ func (m *awsRestjson1_serializeOpDeleteApnsSandboxChan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteApnsSandboxChannelInput(v *DeleteApnsSandboxChannelInput, encoder *httpbinding.Encoder) error {
@@ -1354,10 +1308,8 @@ func (m *awsRestjson1_serializeOpDeleteApnsVoipChannel
 func (m *awsRestjson1_serializeOpDeleteApnsVoipChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1395,7 +1347,6 @@ func (m *awsRestjson1_serializeOpDeleteApnsVoipChannel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteApnsVoipChannelInput(v *DeleteApnsVoipChannelInput, encoder *httpbinding.Encoder) error {
@@ -1425,10 +1376,8 @@ func (m *awsRestjson1_serializeOpDeleteApnsVoipSandbox
 func (m *awsRestjson1_serializeOpDeleteApnsVoipSandboxChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1466,7 +1415,6 @@ func (m *awsRestjson1_serializeOpDeleteApnsVoipSandbox
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteApnsVoipSandboxChannelInput(v *DeleteApnsVoipSandboxChannelInput, encoder *httpbinding.Encoder) error {
@@ -1496,10 +1444,8 @@ func (m *awsRestjson1_serializeOpDeleteApp) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteApp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1537,7 +1483,6 @@ func (m *awsRestjson1_serializeOpDeleteApp) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAppInput(v *DeleteAppInput, encoder *httpbinding.Encoder) error {
@@ -1567,10 +1512,8 @@ func (m *awsRestjson1_serializeOpDeleteBaiduChannel) H
 func (m *awsRestjson1_serializeOpDeleteBaiduChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1608,7 +1551,6 @@ func (m *awsRestjson1_serializeOpDeleteBaiduChannel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBaiduChannelInput(v *DeleteBaiduChannelInput, encoder *httpbinding.Encoder) error {
@@ -1638,10 +1580,8 @@ func (m *awsRestjson1_serializeOpDeleteCampaign) Handl
 func (m *awsRestjson1_serializeOpDeleteCampaign) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1679,7 +1619,6 @@ func (m *awsRestjson1_serializeOpDeleteCampaign) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCampaignInput(v *DeleteCampaignInput, encoder *httpbinding.Encoder) error {
@@ -1718,10 +1657,8 @@ func (m *awsRestjson1_serializeOpDeleteEmailChannel) H
 func (m *awsRestjson1_serializeOpDeleteEmailChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1759,7 +1696,6 @@ func (m *awsRestjson1_serializeOpDeleteEmailChannel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEmailChannelInput(v *DeleteEmailChannelInput, encoder *httpbinding.Encoder) error {
@@ -1789,10 +1725,8 @@ func (m *awsRestjson1_serializeOpDeleteEmailTemplate) 
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
@@ -1830,7 +1764,6 @@ func (m *awsRestjson1_serializeOpDeleteEmailTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEmailTemplateInput(v *DeleteEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1864,10 +1797,8 @@ func (m *awsRestjson1_serializeOpDeleteEndpoint) Handl
 func (m *awsRestjson1_serializeOpDeleteEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1905,7 +1836,6 @@ func (m *awsRestjson1_serializeOpDeleteEndpoint) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEndpointInput(v *DeleteEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1944,10 +1874,8 @@ func (m *awsRestjson1_serializeOpDeleteEventStream) Ha
 func (m *awsRestjson1_serializeOpDeleteEventStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1985,7 +1913,6 @@ func (m *awsRestjson1_serializeOpDeleteEventStream) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventStreamInput(v *DeleteEventStreamInput, encoder *httpbinding.Encoder) error {
@@ -2015,10 +1942,8 @@ func (m *awsRestjson1_serializeOpDeleteGcmChannel) Han
 func (m *awsRestjson1_serializeOpDeleteGcmChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2056,7 +1981,6 @@ func (m *awsRestjson1_serializeOpDeleteGcmChannel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGcmChannelInput(v *DeleteGcmChannelInput, encoder *httpbinding.Encoder) error {
@@ -2086,10 +2010,8 @@ func (m *awsRestjson1_serializeOpDeleteInAppTemplate) 
 func (m *awsRestjson1_serializeOpDeleteInAppTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2127,7 +2049,6 @@ func (m *awsRestjson1_serializeOpDeleteInAppTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInAppTemplateInput(v *DeleteInAppTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2161,10 +2082,8 @@ func (m *awsRestjson1_serializeOpDeleteJourney) Handle
 func (m *awsRestjson1_serializeOpDeleteJourney) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2202,7 +2121,6 @@ func (m *awsRestjson1_serializeOpDeleteJourney) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteJourneyInput(v *DeleteJourneyInput, encoder *httpbinding.Encoder) error {
@@ -2241,10 +2159,8 @@ func (m *awsRestjson1_serializeOpDeletePushTemplate) H
 func (m *awsRestjson1_serializeOpDeletePushTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2282,7 +2198,6 @@ func (m *awsRestjson1_serializeOpDeletePushTemplate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePushTemplateInput(v *DeletePushTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2316,10 +2231,8 @@ func (m *awsRestjson1_serializeOpDeleteRecommenderConf
 func (m *awsRestjson1_serializeOpDeleteRecommenderConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2357,7 +2270,6 @@ func (m *awsRestjson1_serializeOpDeleteRecommenderConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRecommenderConfigurationInput(v *DeleteRecommenderConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2387,10 +2299,8 @@ func (m *awsRestjson1_serializeOpDeleteSegment) Handle
 func (m *awsRestjson1_serializeOpDeleteSegment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2428,7 +2338,6 @@ func (m *awsRestjson1_serializeOpDeleteSegment) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSegmentInput(v *DeleteSegmentInput, encoder *httpbinding.Encoder) error {
@@ -2467,10 +2376,8 @@ func (m *awsRestjson1_serializeOpDeleteSmsChannel) Han
 func (m *awsRestjson1_serializeOpDeleteSmsChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2508,7 +2415,6 @@ func (m *awsRestjson1_serializeOpDeleteSmsChannel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSmsChannelInput(v *DeleteSmsChannelInput, encoder *httpbinding.Encoder) error {
@@ -2538,10 +2444,8 @@ func (m *awsRestjson1_serializeOpDeleteSmsTemplate) Ha
 func (m *awsRestjson1_serializeOpDeleteSmsTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2579,7 +2483,6 @@ func (m *awsRestjson1_serializeOpDeleteSmsTemplate) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSmsTemplateInput(v *DeleteSmsTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2613,10 +2516,8 @@ func (m *awsRestjson1_serializeOpDeleteUserEndpoints) 
 func (m *awsRestjson1_serializeOpDeleteUserEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2654,7 +2555,6 @@ func (m *awsRestjson1_serializeOpDeleteUserEndpoints) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserEndpointsInput(v *DeleteUserEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -2693,10 +2593,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceChannel) H
 func (m *awsRestjson1_serializeOpDeleteVoiceChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2734,7 +2632,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceChannel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceChannelInput(v *DeleteVoiceChannelInput, encoder *httpbinding.Encoder) error {
@@ -2764,10 +2661,8 @@ func (m *awsRestjson1_serializeOpDeleteVoiceTemplate) 
 func (m *awsRestjson1_serializeOpDeleteVoiceTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2805,7 +2700,6 @@ func (m *awsRestjson1_serializeOpDeleteVoiceTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVoiceTemplateInput(v *DeleteVoiceTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2839,10 +2733,8 @@ func (m *awsRestjson1_serializeOpGetAdmChannel) Handle
 func (m *awsRestjson1_serializeOpGetAdmChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2880,7 +2772,6 @@ func (m *awsRestjson1_serializeOpGetAdmChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAdmChannelInput(v *GetAdmChannelInput, encoder *httpbinding.Encoder) error {
@@ -2910,10 +2801,8 @@ func (m *awsRestjson1_serializeOpGetApnsChannel) Handl
 func (m *awsRestjson1_serializeOpGetApnsChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2951,7 +2840,6 @@ func (m *awsRestjson1_serializeOpGetApnsChannel) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApnsChannelInput(v *GetApnsChannelInput, encoder *httpbinding.Encoder) error {
@@ -2981,10 +2869,8 @@ func (m *awsRestjson1_serializeOpGetApnsSandboxChannel
 func (m *awsRestjson1_serializeOpGetApnsSandboxChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3022,7 +2908,6 @@ func (m *awsRestjson1_serializeOpGetApnsSandboxChannel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApnsSandboxChannelInput(v *GetApnsSandboxChannelInput, encoder *httpbinding.Encoder) error {
@@ -3052,10 +2937,8 @@ func (m *awsRestjson1_serializeOpGetApnsVoipChannel) H
 func (m *awsRestjson1_serializeOpGetApnsVoipChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3093,7 +2976,6 @@ func (m *awsRestjson1_serializeOpGetApnsVoipChannel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApnsVoipChannelInput(v *GetApnsVoipChannelInput, encoder *httpbinding.Encoder) error {
@@ -3123,10 +3005,8 @@ func (m *awsRestjson1_serializeOpGetApnsVoipSandboxCha
 func (m *awsRestjson1_serializeOpGetApnsVoipSandboxChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3164,7 +3044,6 @@ func (m *awsRestjson1_serializeOpGetApnsVoipSandboxCha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApnsVoipSandboxChannelInput(v *GetApnsVoipSandboxChannelInput, encoder *httpbinding.Encoder) error {
@@ -3194,10 +3073,8 @@ func (m *awsRestjson1_serializeOpGetApp) HandleSeriali
 func (m *awsRestjson1_serializeOpGetApp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3235,7 +3112,6 @@ func (m *awsRestjson1_serializeOpGetApp) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAppInput(v *GetAppInput, encoder *httpbinding.Encoder) error {
@@ -3265,10 +3141,8 @@ func (m *awsRestjson1_serializeOpGetApplicationDateRan
 func (m *awsRestjson1_serializeOpGetApplicationDateRangeKpi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3306,7 +3180,6 @@ func (m *awsRestjson1_serializeOpGetApplicationDateRan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApplicationDateRangeKpiInput(v *GetApplicationDateRangeKpiInput, encoder *httpbinding.Encoder) error {
@@ -3361,10 +3234,8 @@ func (m *awsRestjson1_serializeOpGetApplicationSetting
 func (m *awsRestjson1_serializeOpGetApplicationSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3402,7 +3273,6 @@ func (m *awsRestjson1_serializeOpGetApplicationSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApplicationSettingsInput(v *GetApplicationSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3432,10 +3302,8 @@ func (m *awsRestjson1_serializeOpGetApps) HandleSerial
 func (m *awsRestjson1_serializeOpGetApps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3473,7 +3341,6 @@ func (m *awsRestjson1_serializeOpGetApps) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAppsInput(v *GetAppsInput, encoder *httpbinding.Encoder) error {
@@ -3502,10 +3369,8 @@ func (m *awsRestjson1_serializeOpGetBaiduChannel) Hand
 func (m *awsRestjson1_serializeOpGetBaiduChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3543,7 +3408,6 @@ func (m *awsRestjson1_serializeOpGetBaiduChannel) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBaiduChannelInput(v *GetBaiduChannelInput, encoder *httpbinding.Encoder) error {
@@ -3573,10 +3437,8 @@ func (m *awsRestjson1_serializeOpGetCampaign) HandleSe
 func (m *awsRestjson1_serializeOpGetCampaign) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3614,7 +3476,6 @@ func (m *awsRestjson1_serializeOpGetCampaign) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCampaignInput(v *GetCampaignInput, encoder *httpbinding.Encoder) error {
@@ -3653,10 +3514,8 @@ func (m *awsRestjson1_serializeOpGetCampaignActivities
 func (m *awsRestjson1_serializeOpGetCampaignActivities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3694,7 +3553,6 @@ func (m *awsRestjson1_serializeOpGetCampaignActivities
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCampaignActivitiesInput(v *GetCampaignActivitiesInput, encoder *httpbinding.Encoder) error {
@@ -3741,10 +3599,8 @@ func (m *awsRestjson1_serializeOpGetCampaignDateRangeK
 func (m *awsRestjson1_serializeOpGetCampaignDateRangeKpi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3782,7 +3638,6 @@ func (m *awsRestjson1_serializeOpGetCampaignDateRangeK
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCampaignDateRangeKpiInput(v *GetCampaignDateRangeKpiInput, encoder *httpbinding.Encoder) error {
@@ -3846,10 +3701,8 @@ func (m *awsRestjson1_serializeOpGetCampaigns) HandleS
 func (m *awsRestjson1_serializeOpGetCampaigns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3887,7 +3740,6 @@ func (m *awsRestjson1_serializeOpGetCampaigns) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCampaignsInput(v *GetCampaignsInput, encoder *httpbinding.Encoder) error {
@@ -3925,10 +3777,8 @@ func (m *awsRestjson1_serializeOpGetCampaignVersion) H
 func (m *awsRestjson1_serializeOpGetCampaignVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3966,7 +3816,6 @@ func (m *awsRestjson1_serializeOpGetCampaignVersion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCampaignVersionInput(v *GetCampaignVersionInput, encoder *httpbinding.Encoder) error {
@@ -4014,10 +3863,8 @@ func (m *awsRestjson1_serializeOpGetCampaignVersions) 
 func (m *awsRestjson1_serializeOpGetCampaignVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4055,7 +3902,6 @@ func (m *awsRestjson1_serializeOpGetCampaignVersions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCampaignVersionsInput(v *GetCampaignVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4102,10 +3948,8 @@ func (m *awsRestjson1_serializeOpGetChannels) HandleSe
 func (m *awsRestjson1_serializeOpGetChannels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4143,7 +3987,6 @@ func (m *awsRestjson1_serializeOpGetChannels) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetChannelsInput(v *GetChannelsInput, encoder *httpbinding.Encoder) error {
@@ -4173,10 +4016,8 @@ func (m *awsRestjson1_serializeOpGetEmailChannel) Hand
 func (m *awsRestjson1_serializeOpGetEmailChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4214,7 +4055,6 @@ func (m *awsRestjson1_serializeOpGetEmailChannel) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEmailChannelInput(v *GetEmailChannelInput, encoder *httpbinding.Encoder) error {
@@ -4244,10 +4084,8 @@ func (m *awsRestjson1_serializeOpGetEmailTemplate) Han
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
@@ -4285,7 +4123,6 @@ func (m *awsRestjson1_serializeOpGetEmailTemplate) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEmailTemplateInput(v *GetEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -4319,10 +4156,8 @@ func (m *awsRestjson1_serializeOpGetEndpoint) HandleSe
 func (m *awsRestjson1_serializeOpGetEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4360,7 +4195,6 @@ func (m *awsRestjson1_serializeOpGetEndpoint) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEndpointInput(v *GetEndpointInput, encoder *httpbinding.Encoder) error {
@@ -4399,10 +4233,8 @@ func (m *awsRestjson1_serializeOpGetEventStream) Handl
 func (m *awsRestjson1_serializeOpGetEventStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4440,7 +4272,6 @@ func (m *awsRestjson1_serializeOpGetEventStream) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventStreamInput(v *GetEventStreamInput, encoder *httpbinding.Encoder) error {
@@ -4470,10 +4301,8 @@ func (m *awsRestjson1_serializeOpGetExportJob) HandleS
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
@@ -4511,7 +4340,6 @@ func (m *awsRestjson1_serializeOpGetExportJob) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetExportJobInput(v *GetExportJobInput, encoder *httpbinding.Encoder) error {
@@ -4550,10 +4378,8 @@ func (m *awsRestjson1_serializeOpGetExportJobs) Handle
 func (m *awsRestjson1_serializeOpGetExportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4591,7 +4417,6 @@ func (m *awsRestjson1_serializeOpGetExportJobs) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetExportJobsInput(v *GetExportJobsInput, encoder *httpbinding.Encoder) error {
@@ -4629,10 +4454,8 @@ func (m *awsRestjson1_serializeOpGetGcmChannel) Handle
 func (m *awsRestjson1_serializeOpGetGcmChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4670,7 +4493,6 @@ func (m *awsRestjson1_serializeOpGetGcmChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGcmChannelInput(v *GetGcmChannelInput, encoder *httpbinding.Encoder) error {
@@ -4700,10 +4522,8 @@ func (m *awsRestjson1_serializeOpGetImportJob) HandleS
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
@@ -4741,7 +4561,6 @@ func (m *awsRestjson1_serializeOpGetImportJob) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImportJobInput(v *GetImportJobInput, encoder *httpbinding.Encoder) error {
@@ -4780,10 +4599,8 @@ func (m *awsRestjson1_serializeOpGetImportJobs) Handle
 func (m *awsRestjson1_serializeOpGetImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4821,7 +4638,6 @@ func (m *awsRestjson1_serializeOpGetImportJobs) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImportJobsInput(v *GetImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -4859,10 +4675,8 @@ func (m *awsRestjson1_serializeOpGetInAppMessages) Han
 func (m *awsRestjson1_serializeOpGetInAppMessages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4900,7 +4714,6 @@ func (m *awsRestjson1_serializeOpGetInAppMessages) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInAppMessagesInput(v *GetInAppMessagesInput, encoder *httpbinding.Encoder) error {
@@ -4939,10 +4752,8 @@ func (m *awsRestjson1_serializeOpGetInAppTemplate) Han
 func (m *awsRestjson1_serializeOpGetInAppTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4980,7 +4791,6 @@ func (m *awsRestjson1_serializeOpGetInAppTemplate) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInAppTemplateInput(v *GetInAppTemplateInput, encoder *httpbinding.Encoder) error {
@@ -5014,10 +4824,8 @@ func (m *awsRestjson1_serializeOpGetJourney) HandleSer
 func (m *awsRestjson1_serializeOpGetJourney) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5055,7 +4863,6 @@ func (m *awsRestjson1_serializeOpGetJourney) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJourneyInput(v *GetJourneyInput, encoder *httpbinding.Encoder) error {
@@ -5094,10 +4901,8 @@ func (m *awsRestjson1_serializeOpGetJourneyDateRangeKp
 func (m *awsRestjson1_serializeOpGetJourneyDateRangeKpi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5135,7 +4940,6 @@ func (m *awsRestjson1_serializeOpGetJourneyDateRangeKp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJourneyDateRangeKpiInput(v *GetJourneyDateRangeKpiInput, encoder *httpbinding.Encoder) error {
@@ -5199,10 +5003,8 @@ func (m *awsRestjson1_serializeOpGetJourneyExecutionAc
 func (m *awsRestjson1_serializeOpGetJourneyExecutionActivityMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5240,7 +5042,6 @@ func (m *awsRestjson1_serializeOpGetJourneyExecutionAc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJourneyExecutionActivityMetricsInput(v *GetJourneyExecutionActivityMetricsInput, encoder *httpbinding.Encoder) error {
@@ -5296,10 +5097,8 @@ func (m *awsRestjson1_serializeOpGetJourneyExecutionMe
 func (m *awsRestjson1_serializeOpGetJourneyExecutionMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5337,7 +5136,6 @@ func (m *awsRestjson1_serializeOpGetJourneyExecutionMe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJourneyExecutionMetricsInput(v *GetJourneyExecutionMetricsInput, encoder *httpbinding.Encoder) error {
@@ -5384,10 +5182,8 @@ func (m *awsRestjson1_serializeOpGetJourneyRunExecutio
 func (m *awsRestjson1_serializeOpGetJourneyRunExecutionActivityMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5425,7 +5221,6 @@ func (m *awsRestjson1_serializeOpGetJourneyRunExecutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJourneyRunExecutionActivityMetricsInput(v *GetJourneyRunExecutionActivityMetricsInput, encoder *httpbinding.Encoder) error {
@@ -5490,10 +5285,8 @@ func (m *awsRestjson1_serializeOpGetJourneyRunExecutio
 func (m *awsRestjson1_serializeOpGetJourneyRunExecutionMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5531,7 +5324,6 @@ func (m *awsRestjson1_serializeOpGetJourneyRunExecutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJourneyRunExecutionMetricsInput(v *GetJourneyRunExecutionMetricsInput, encoder *httpbinding.Encoder) error {
@@ -5587,10 +5379,8 @@ func (m *awsRestjson1_serializeOpGetJourneyRuns) Handl
 func (m *awsRestjson1_serializeOpGetJourneyRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5628,7 +5418,6 @@ func (m *awsRestjson1_serializeOpGetJourneyRuns) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJourneyRunsInput(v *GetJourneyRunsInput, encoder *httpbinding.Encoder) error {
@@ -5675,10 +5464,8 @@ func (m *awsRestjson1_serializeOpGetPushTemplate) Hand
 func (m *awsRestjson1_serializeOpGetPushTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5716,7 +5503,6 @@ func (m *awsRestjson1_serializeOpGetPushTemplate) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPushTemplateInput(v *GetPushTemplateInput, encoder *httpbinding.Encoder) error {
@@ -5750,10 +5536,8 @@ func (m *awsRestjson1_serializeOpGetRecommenderConfigu
 func (m *awsRestjson1_serializeOpGetRecommenderConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5791,7 +5575,6 @@ func (m *awsRestjson1_serializeOpGetRecommenderConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRecommenderConfigurationInput(v *GetRecommenderConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5821,10 +5604,8 @@ func (m *awsRestjson1_serializeOpGetRecommenderConfigu
 func (m *awsRestjson1_serializeOpGetRecommenderConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5862,7 +5643,6 @@ func (m *awsRestjson1_serializeOpGetRecommenderConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRecommenderConfigurationsInput(v *GetRecommenderConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5891,10 +5671,8 @@ func (m *awsRestjson1_serializeOpGetSegment) HandleSer
 func (m *awsRestjson1_serializeOpGetSegment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5932,7 +5710,6 @@ func (m *awsRestjson1_serializeOpGetSegment) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentInput(v *GetSegmentInput, encoder *httpbinding.Encoder) error {
@@ -5971,10 +5748,8 @@ func (m *awsRestjson1_serializeOpGetSegmentExportJobs)
 func (m *awsRestjson1_serializeOpGetSegmentExportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6012,7 +5787,6 @@ func (m *awsRestjson1_serializeOpGetSegmentExportJobs)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentExportJobsInput(v *GetSegmentExportJobsInput, encoder *httpbinding.Encoder) error {
@@ -6059,10 +5833,8 @@ func (m *awsRestjson1_serializeOpGetSegmentImportJobs)
 func (m *awsRestjson1_serializeOpGetSegmentImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6100,7 +5872,6 @@ func (m *awsRestjson1_serializeOpGetSegmentImportJobs)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentImportJobsInput(v *GetSegmentImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -6147,10 +5918,8 @@ func (m *awsRestjson1_serializeOpGetSegments) HandleSe
 func (m *awsRestjson1_serializeOpGetSegments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6188,7 +5957,6 @@ func (m *awsRestjson1_serializeOpGetSegments) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentsInput(v *GetSegmentsInput, encoder *httpbinding.Encoder) error {
@@ -6226,10 +5994,8 @@ func (m *awsRestjson1_serializeOpGetSegmentVersion) Ha
 func (m *awsRestjson1_serializeOpGetSegmentVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6267,7 +6033,6 @@ func (m *awsRestjson1_serializeOpGetSegmentVersion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentVersionInput(v *GetSegmentVersionInput, encoder *httpbinding.Encoder) error {
@@ -6315,10 +6080,8 @@ func (m *awsRestjson1_serializeOpGetSegmentVersions) H
 func (m *awsRestjson1_serializeOpGetSegmentVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6356,7 +6119,6 @@ func (m *awsRestjson1_serializeOpGetSegmentVersions) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentVersionsInput(v *GetSegmentVersionsInput, encoder *httpbinding.Encoder) error {
@@ -6403,10 +6165,8 @@ func (m *awsRestjson1_serializeOpGetSmsChannel) Handle
 func (m *awsRestjson1_serializeOpGetSmsChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6444,7 +6204,6 @@ func (m *awsRestjson1_serializeOpGetSmsChannel) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSmsChannelInput(v *GetSmsChannelInput, encoder *httpbinding.Encoder) error {
@@ -6474,10 +6233,8 @@ func (m *awsRestjson1_serializeOpGetSmsTemplate) Handl
 func (m *awsRestjson1_serializeOpGetSmsTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6515,7 +6272,6 @@ func (m *awsRestjson1_serializeOpGetSmsTemplate) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSmsTemplateInput(v *GetSmsTemplateInput, encoder *httpbinding.Encoder) error {
@@ -6549,10 +6305,8 @@ func (m *awsRestjson1_serializeOpGetUserEndpoints) Han
 func (m *awsRestjson1_serializeOpGetUserEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6590,7 +6344,6 @@ func (m *awsRestjson1_serializeOpGetUserEndpoints) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserEndpointsInput(v *GetUserEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -6629,10 +6382,8 @@ func (m *awsRestjson1_serializeOpGetVoiceChannel) Hand
 func (m *awsRestjson1_serializeOpGetVoiceChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6670,7 +6421,6 @@ func (m *awsRestjson1_serializeOpGetVoiceChannel) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceChannelInput(v *GetVoiceChannelInput, encoder *httpbinding.Encoder) error {
@@ -6700,10 +6450,8 @@ func (m *awsRestjson1_serializeOpGetVoiceTemplate) Han
 func (m *awsRestjson1_serializeOpGetVoiceTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6741,7 +6489,6 @@ func (m *awsRestjson1_serializeOpGetVoiceTemplate) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceTemplateInput(v *GetVoiceTemplateInput, encoder *httpbinding.Encoder) error {
@@ -6775,10 +6522,8 @@ func (m *awsRestjson1_serializeOpListJourneys) HandleS
 func (m *awsRestjson1_serializeOpListJourneys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6816,7 +6561,6 @@ func (m *awsRestjson1_serializeOpListJourneys) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJourneysInput(v *ListJourneysInput, encoder *httpbinding.Encoder) error {
@@ -6854,10 +6598,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6895,7 +6637,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6925,10 +6666,8 @@ func (m *awsRestjson1_serializeOpListTemplates) Handle
 func (m *awsRestjson1_serializeOpListTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6966,7 +6705,6 @@ func (m *awsRestjson1_serializeOpListTemplates) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTemplatesInput(v *ListTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -7003,10 +6741,8 @@ func (m *awsRestjson1_serializeOpListTemplateVersions)
 func (m *awsRestjson1_serializeOpListTemplateVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7044,7 +6780,6 @@ func (m *awsRestjson1_serializeOpListTemplateVersions)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTemplateVersionsInput(v *ListTemplateVersionsInput, encoder *httpbinding.Encoder) error {
@@ -7091,10 +6826,8 @@ func (m *awsRestjson1_serializeOpPhoneNumberValidate) 
 func (m *awsRestjson1_serializeOpPhoneNumberValidate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7152,7 +6885,6 @@ func (m *awsRestjson1_serializeOpPhoneNumberValidate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPhoneNumberValidateInput(v *PhoneNumberValidateInput, encoder *httpbinding.Encoder) error {
@@ -7173,10 +6905,8 @@ func (m *awsRestjson1_serializeOpPutEvents) HandleSeri
 func (m *awsRestjson1_serializeOpPutEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7238,7 +6968,6 @@ func (m *awsRestjson1_serializeOpPutEvents) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEventsInput(v *PutEventsInput, encoder *httpbinding.Encoder) error {
@@ -7268,10 +6997,8 @@ func (m *awsRestjson1_serializeOpPutEventStream) Handl
 func (m *awsRestjson1_serializeOpPutEventStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7333,7 +7060,6 @@ func (m *awsRestjson1_serializeOpPutEventStream) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEventStreamInput(v *PutEventStreamInput, encoder *httpbinding.Encoder) error {
@@ -7363,10 +7089,8 @@ func (m *awsRestjson1_serializeOpRemoveAttributes) Han
 func (m *awsRestjson1_serializeOpRemoveAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7428,7 +7152,6 @@ func (m *awsRestjson1_serializeOpRemoveAttributes) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveAttributesInput(v *RemoveAttributesInput, encoder *httpbinding.Encoder) error {
@@ -7467,10 +7190,8 @@ func (m *awsRestjson1_serializeOpSendMessages) HandleS
 func (m *awsRestjson1_serializeOpSendMessages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7532,7 +7253,6 @@ func (m *awsRestjson1_serializeOpSendMessages) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendMessagesInput(v *SendMessagesInput, encoder *httpbinding.Encoder) error {
@@ -7562,10 +7282,8 @@ func (m *awsRestjson1_serializeOpSendOTPMessage) Handl
 func (m *awsRestjson1_serializeOpSendOTPMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7627,7 +7345,6 @@ func (m *awsRestjson1_serializeOpSendOTPMessage) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendOTPMessageInput(v *SendOTPMessageInput, encoder *httpbinding.Encoder) error {
@@ -7657,10 +7374,8 @@ func (m *awsRestjson1_serializeOpSendUsersMessages) Ha
 func (m *awsRestjson1_serializeOpSendUsersMessages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7722,7 +7437,6 @@ func (m *awsRestjson1_serializeOpSendUsersMessages) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendUsersMessagesInput(v *SendUsersMessagesInput, encoder *httpbinding.Encoder) error {
@@ -7752,10 +7466,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7847,10 +7558,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7888,7 +7597,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7924,10 +7632,8 @@ func (m *awsRestjson1_serializeOpUpdateAdmChannel) Han
 func (m *awsRestjson1_serializeOpUpdateAdmChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7989,7 +7695,6 @@ func (m *awsRestjson1_serializeOpUpdateAdmChannel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAdmChannelInput(v *UpdateAdmChannelInput, encoder *httpbinding.Encoder) error {
@@ -8019,10 +7724,8 @@ func (m *awsRestjson1_serializeOpUpdateApnsChannel) Ha
 func (m *awsRestjson1_serializeOpUpdateApnsChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8084,7 +7787,6 @@ func (m *awsRestjson1_serializeOpUpdateApnsChannel) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApnsChannelInput(v *UpdateApnsChannelInput, encoder *httpbinding.Encoder) error {
@@ -8114,10 +7816,8 @@ func (m *awsRestjson1_serializeOpUpdateApnsSandboxChan
 func (m *awsRestjson1_serializeOpUpdateApnsSandboxChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8179,7 +7879,6 @@ func (m *awsRestjson1_serializeOpUpdateApnsSandboxChan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApnsSandboxChannelInput(v *UpdateApnsSandboxChannelInput, encoder *httpbinding.Encoder) error {
@@ -8209,10 +7908,8 @@ func (m *awsRestjson1_serializeOpUpdateApnsVoipChannel
 func (m *awsRestjson1_serializeOpUpdateApnsVoipChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8274,7 +7971,6 @@ func (m *awsRestjson1_serializeOpUpdateApnsVoipChannel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApnsVoipChannelInput(v *UpdateApnsVoipChannelInput, encoder *httpbinding.Encoder) error {
@@ -8304,10 +8000,8 @@ func (m *awsRestjson1_serializeOpUpdateApnsVoipSandbox
 func (m *awsRestjson1_serializeOpUpdateApnsVoipSandboxChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8369,7 +8063,6 @@ func (m *awsRestjson1_serializeOpUpdateApnsVoipSandbox
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApnsVoipSandboxChannelInput(v *UpdateApnsVoipSandboxChannelInput, encoder *httpbinding.Encoder) error {
@@ -8399,10 +8092,8 @@ func (m *awsRestjson1_serializeOpUpdateApplicationSett
 func (m *awsRestjson1_serializeOpUpdateApplicationSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8464,7 +8155,6 @@ func (m *awsRestjson1_serializeOpUpdateApplicationSett
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApplicationSettingsInput(v *UpdateApplicationSettingsInput, encoder *httpbinding.Encoder) error {
@@ -8494,10 +8184,8 @@ func (m *awsRestjson1_serializeOpUpdateBaiduChannel) H
 func (m *awsRestjson1_serializeOpUpdateBaiduChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8559,7 +8247,6 @@ func (m *awsRestjson1_serializeOpUpdateBaiduChannel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBaiduChannelInput(v *UpdateBaiduChannelInput, encoder *httpbinding.Encoder) error {
@@ -8589,10 +8276,8 @@ func (m *awsRestjson1_serializeOpUpdateCampaign) Handl
 func (m *awsRestjson1_serializeOpUpdateCampaign) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8654,7 +8339,6 @@ func (m *awsRestjson1_serializeOpUpdateCampaign) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCampaignInput(v *UpdateCampaignInput, encoder *httpbinding.Encoder) error {
@@ -8693,10 +8377,8 @@ func (m *awsRestjson1_serializeOpUpdateEmailChannel) H
 func (m *awsRestjson1_serializeOpUpdateEmailChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8758,7 +8440,6 @@ func (m *awsRestjson1_serializeOpUpdateEmailChannel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEmailChannelInput(v *UpdateEmailChannelInput, encoder *httpbinding.Encoder) error {
@@ -8788,10 +8469,8 @@ func (m *awsRestjson1_serializeOpUpdateEmailTemplate) 
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
@@ -8853,7 +8532,6 @@ func (m *awsRestjson1_serializeOpUpdateEmailTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEmailTemplateInput(v *UpdateEmailTemplateInput, encoder *httpbinding.Encoder) error {
@@ -8891,10 +8569,8 @@ func (m *awsRestjson1_serializeOpUpdateEndpoint) Handl
 func (m *awsRestjson1_serializeOpUpdateEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8956,7 +8632,6 @@ func (m *awsRestjson1_serializeOpUpdateEndpoint) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEndpointInput(v *UpdateEndpointInput, encoder *httpbinding.Encoder) error {
@@ -8995,10 +8670,8 @@ func (m *awsRestjson1_serializeOpUpdateEndpointsBatch)
 func (m *awsRestjson1_serializeOpUpdateEndpointsBatch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9060,7 +8733,6 @@ func (m *awsRestjson1_serializeOpUpdateEndpointsBatch)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEndpointsBatchInput(v *UpdateEndpointsBatchInput, encoder *httpbinding.Encoder) error {
@@ -9090,10 +8762,8 @@ func (m *awsRestjson1_serializeOpUpdateGcmChannel) Han
 func (m *awsRestjson1_serializeOpUpdateGcmChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9155,7 +8825,6 @@ func (m *awsRestjson1_serializeOpUpdateGcmChannel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGcmChannelInput(v *UpdateGcmChannelInput, encoder *httpbinding.Encoder) error {
@@ -9185,10 +8854,8 @@ func (m *awsRestjson1_serializeOpUpdateInAppTemplate) 
 func (m *awsRestjson1_serializeOpUpdateInAppTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9250,7 +8917,6 @@ func (m *awsRestjson1_serializeOpUpdateInAppTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInAppTemplateInput(v *UpdateInAppTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9288,10 +8954,8 @@ func (m *awsRestjson1_serializeOpUpdateJourney) Handle
 func (m *awsRestjson1_serializeOpUpdateJourney) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9353,7 +9017,6 @@ func (m *awsRestjson1_serializeOpUpdateJourney) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateJourneyInput(v *UpdateJourneyInput, encoder *httpbinding.Encoder) error {
@@ -9392,10 +9055,8 @@ func (m *awsRestjson1_serializeOpUpdateJourneyState) H
 func (m *awsRestjson1_serializeOpUpdateJourneyState) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9457,7 +9118,6 @@ func (m *awsRestjson1_serializeOpUpdateJourneyState) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateJourneyStateInput(v *UpdateJourneyStateInput, encoder *httpbinding.Encoder) error {
@@ -9496,10 +9156,8 @@ func (m *awsRestjson1_serializeOpUpdatePushTemplate) H
 func (m *awsRestjson1_serializeOpUpdatePushTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9561,7 +9219,6 @@ func (m *awsRestjson1_serializeOpUpdatePushTemplate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePushTemplateInput(v *UpdatePushTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9599,10 +9256,8 @@ func (m *awsRestjson1_serializeOpUpdateRecommenderConf
 func (m *awsRestjson1_serializeOpUpdateRecommenderConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9664,7 +9319,6 @@ func (m *awsRestjson1_serializeOpUpdateRecommenderConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRecommenderConfigurationInput(v *UpdateRecommenderConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -9694,10 +9348,8 @@ func (m *awsRestjson1_serializeOpUpdateSegment) Handle
 func (m *awsRestjson1_serializeOpUpdateSegment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9759,7 +9411,6 @@ func (m *awsRestjson1_serializeOpUpdateSegment) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSegmentInput(v *UpdateSegmentInput, encoder *httpbinding.Encoder) error {
@@ -9798,10 +9449,8 @@ func (m *awsRestjson1_serializeOpUpdateSmsChannel) Han
 func (m *awsRestjson1_serializeOpUpdateSmsChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9863,7 +9512,6 @@ func (m *awsRestjson1_serializeOpUpdateSmsChannel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSmsChannelInput(v *UpdateSmsChannelInput, encoder *httpbinding.Encoder) error {
@@ -9893,10 +9541,8 @@ func (m *awsRestjson1_serializeOpUpdateSmsTemplate) Ha
 func (m *awsRestjson1_serializeOpUpdateSmsTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9958,7 +9604,6 @@ func (m *awsRestjson1_serializeOpUpdateSmsTemplate) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSmsTemplateInput(v *UpdateSmsTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9996,10 +9641,8 @@ func (m *awsRestjson1_serializeOpUpdateTemplateActiveV
 func (m *awsRestjson1_serializeOpUpdateTemplateActiveVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10061,7 +9704,6 @@ func (m *awsRestjson1_serializeOpUpdateTemplateActiveV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTemplateActiveVersionInput(v *UpdateTemplateActiveVersionInput, encoder *httpbinding.Encoder) error {
@@ -10100,10 +9742,8 @@ func (m *awsRestjson1_serializeOpUpdateVoiceChannel) H
 func (m *awsRestjson1_serializeOpUpdateVoiceChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10165,7 +9805,6 @@ func (m *awsRestjson1_serializeOpUpdateVoiceChannel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVoiceChannelInput(v *UpdateVoiceChannelInput, encoder *httpbinding.Encoder) error {
@@ -10195,10 +9834,8 @@ func (m *awsRestjson1_serializeOpUpdateVoiceTemplate) 
 func (m *awsRestjson1_serializeOpUpdateVoiceTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10260,7 +9897,6 @@ func (m *awsRestjson1_serializeOpUpdateVoiceTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVoiceTemplateInput(v *UpdateVoiceTemplateInput, encoder *httpbinding.Encoder) error {
@@ -10298,10 +9934,8 @@ func (m *awsRestjson1_serializeOpVerifyOTPMessage) Han
 func (m *awsRestjson1_serializeOpVerifyOTPMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10363,7 +9997,6 @@ func (m *awsRestjson1_serializeOpVerifyOTPMessage) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsVerifyOTPMessageInput(v *VerifyOTPMessageInput, encoder *httpbinding.Encoder) error {
