--- vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCreateMediaCapturePip
 func (m *awsRestjson1_serializeOpCreateMediaCapturePipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpCreateMediaCapturePip
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMediaCapturePipelineInput(v *CreateMediaCapturePipelineInput, encoder *httpbinding.Encoder) error {
@@ -153,10 +149,8 @@ func (m *awsRestjson1_serializeOpCreateMediaConcatenat
 func (m *awsRestjson1_serializeOpCreateMediaConcatenationPipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -201,7 +195,6 @@ func (m *awsRestjson1_serializeOpCreateMediaConcatenat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMediaConcatenationPipelineInput(v *CreateMediaConcatenationPipelineInput, encoder *httpbinding.Encoder) error {
@@ -255,10 +248,8 @@ func (m *awsRestjson1_serializeOpCreateMediaInsightsPi
 func (m *awsRestjson1_serializeOpCreateMediaInsightsPipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -303,7 +294,6 @@ func (m *awsRestjson1_serializeOpCreateMediaInsightsPi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMediaInsightsPipelineInput(v *CreateMediaInsightsPipelineInput, encoder *httpbinding.Encoder) error {
@@ -376,10 +366,8 @@ func (m *awsRestjson1_serializeOpCreateMediaInsightsPi
 func (m *awsRestjson1_serializeOpCreateMediaInsightsPipelineConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -424,7 +412,6 @@ func (m *awsRestjson1_serializeOpCreateMediaInsightsPi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMediaInsightsPipelineConfigurationInput(v *CreateMediaInsightsPipelineConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -488,10 +475,8 @@ func (m *awsRestjson1_serializeOpCreateMediaLiveConnec
 func (m *awsRestjson1_serializeOpCreateMediaLiveConnectorPipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -536,7 +521,6 @@ func (m *awsRestjson1_serializeOpCreateMediaLiveConnec
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMediaLiveConnectorPipelineInput(v *CreateMediaLiveConnectorPipelineInput, encoder *httpbinding.Encoder) error {
@@ -590,10 +574,8 @@ func (m *awsRestjson1_serializeOpCreateMediaPipelineKi
 func (m *awsRestjson1_serializeOpCreateMediaPipelineKinesisVideoStreamPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -638,7 +620,6 @@ func (m *awsRestjson1_serializeOpCreateMediaPipelineKi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMediaPipelineKinesisVideoStreamPoolInput(v *CreateMediaPipelineKinesisVideoStreamPoolInput, encoder *httpbinding.Encoder) error {
@@ -690,10 +671,8 @@ func (m *awsRestjson1_serializeOpCreateMediaStreamPipe
 func (m *awsRestjson1_serializeOpCreateMediaStreamPipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -738,7 +717,6 @@ func (m *awsRestjson1_serializeOpCreateMediaStreamPipe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMediaStreamPipelineInput(v *CreateMediaStreamPipelineInput, encoder *httpbinding.Encoder) error {
@@ -792,10 +770,8 @@ func (m *awsRestjson1_serializeOpDeleteMediaCapturePip
 func (m *awsRestjson1_serializeOpDeleteMediaCapturePipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -833,7 +809,6 @@ func (m *awsRestjson1_serializeOpDeleteMediaCapturePip
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMediaCapturePipelineInput(v *DeleteMediaCapturePipelineInput, encoder *httpbinding.Encoder) error {
@@ -863,10 +838,8 @@ func (m *awsRestjson1_serializeOpDeleteMediaInsightsPi
 func (m *awsRestjson1_serializeOpDeleteMediaInsightsPipelineConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -904,7 +877,6 @@ func (m *awsRestjson1_serializeOpDeleteMediaInsightsPi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMediaInsightsPipelineConfigurationInput(v *DeleteMediaInsightsPipelineConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -934,10 +906,8 @@ func (m *awsRestjson1_serializeOpDeleteMediaPipeline) 
 func (m *awsRestjson1_serializeOpDeleteMediaPipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -975,7 +945,6 @@ func (m *awsRestjson1_serializeOpDeleteMediaPipeline) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMediaPipelineInput(v *DeleteMediaPipelineInput, encoder *httpbinding.Encoder) error {
@@ -1005,10 +974,8 @@ func (m *awsRestjson1_serializeOpDeleteMediaPipelineKi
 func (m *awsRestjson1_serializeOpDeleteMediaPipelineKinesisVideoStreamPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1046,7 +1013,6 @@ func (m *awsRestjson1_serializeOpDeleteMediaPipelineKi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMediaPipelineKinesisVideoStreamPoolInput(v *DeleteMediaPipelineKinesisVideoStreamPoolInput, encoder *httpbinding.Encoder) error {
@@ -1076,10 +1042,8 @@ func (m *awsRestjson1_serializeOpGetMediaCapturePipeli
 func (m *awsRestjson1_serializeOpGetMediaCapturePipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1117,7 +1081,6 @@ func (m *awsRestjson1_serializeOpGetMediaCapturePipeli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMediaCapturePipelineInput(v *GetMediaCapturePipelineInput, encoder *httpbinding.Encoder) error {
@@ -1147,10 +1110,8 @@ func (m *awsRestjson1_serializeOpGetMediaInsightsPipel
 func (m *awsRestjson1_serializeOpGetMediaInsightsPipelineConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1188,7 +1149,6 @@ func (m *awsRestjson1_serializeOpGetMediaInsightsPipel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMediaInsightsPipelineConfigurationInput(v *GetMediaInsightsPipelineConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1218,10 +1178,8 @@ func (m *awsRestjson1_serializeOpGetMediaPipeline) Han
 func (m *awsRestjson1_serializeOpGetMediaPipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1259,7 +1217,6 @@ func (m *awsRestjson1_serializeOpGetMediaPipeline) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMediaPipelineInput(v *GetMediaPipelineInput, encoder *httpbinding.Encoder) error {
@@ -1289,10 +1246,8 @@ func (m *awsRestjson1_serializeOpGetMediaPipelineKines
 func (m *awsRestjson1_serializeOpGetMediaPipelineKinesisVideoStreamPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1330,7 +1285,6 @@ func (m *awsRestjson1_serializeOpGetMediaPipelineKines
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMediaPipelineKinesisVideoStreamPoolInput(v *GetMediaPipelineKinesisVideoStreamPoolInput, encoder *httpbinding.Encoder) error {
@@ -1360,10 +1314,8 @@ func (m *awsRestjson1_serializeOpGetSpeakerSearchTask)
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
@@ -1401,7 +1353,6 @@ func (m *awsRestjson1_serializeOpGetSpeakerSearchTask)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSpeakerSearchTaskInput(v *GetSpeakerSearchTaskInput, encoder *httpbinding.Encoder) error {
@@ -1440,10 +1391,8 @@ func (m *awsRestjson1_serializeOpGetVoiceToneAnalysisT
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
@@ -1481,7 +1430,6 @@ func (m *awsRestjson1_serializeOpGetVoiceToneAnalysisT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVoiceToneAnalysisTaskInput(v *GetVoiceToneAnalysisTaskInput, encoder *httpbinding.Encoder) error {
@@ -1520,10 +1468,8 @@ func (m *awsRestjson1_serializeOpListMediaCapturePipel
 func (m *awsRestjson1_serializeOpListMediaCapturePipelines) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1561,7 +1507,6 @@ func (m *awsRestjson1_serializeOpListMediaCapturePipel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMediaCapturePipelinesInput(v *ListMediaCapturePipelinesInput, encoder *httpbinding.Encoder) error {
@@ -1590,10 +1535,8 @@ func (m *awsRestjson1_serializeOpListMediaInsightsPipe
 func (m *awsRestjson1_serializeOpListMediaInsightsPipelineConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1631,7 +1574,6 @@ func (m *awsRestjson1_serializeOpListMediaInsightsPipe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMediaInsightsPipelineConfigurationsInput(v *ListMediaInsightsPipelineConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -1660,10 +1602,8 @@ func (m *awsRestjson1_serializeOpListMediaPipelineKine
 func (m *awsRestjson1_serializeOpListMediaPipelineKinesisVideoStreamPools) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1701,7 +1641,6 @@ func (m *awsRestjson1_serializeOpListMediaPipelineKine
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMediaPipelineKinesisVideoStreamPoolsInput(v *ListMediaPipelineKinesisVideoStreamPoolsInput, encoder *httpbinding.Encoder) error {
@@ -1730,10 +1669,8 @@ func (m *awsRestjson1_serializeOpListMediaPipelines) H
 func (m *awsRestjson1_serializeOpListMediaPipelines) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1771,7 +1708,6 @@ func (m *awsRestjson1_serializeOpListMediaPipelines) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMediaPipelinesInput(v *ListMediaPipelinesInput, encoder *httpbinding.Encoder) error {
@@ -1800,10 +1736,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -1841,7 +1775,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -1866,10 +1799,8 @@ func (m *awsRestjson1_serializeOpStartSpeakerSearchTas
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
@@ -1918,7 +1849,6 @@ func (m *awsRestjson1_serializeOpStartSpeakerSearchTas
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartSpeakerSearchTaskInput(v *StartSpeakerSearchTaskInput, encoder *httpbinding.Encoder) error {
@@ -1972,10 +1902,8 @@ func (m *awsRestjson1_serializeOpStartVoiceToneAnalysi
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
@@ -2024,7 +1952,6 @@ func (m *awsRestjson1_serializeOpStartVoiceToneAnalysi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartVoiceToneAnalysisTaskInput(v *StartVoiceToneAnalysisTaskInput, encoder *httpbinding.Encoder) error {
@@ -2078,10 +2005,8 @@ func (m *awsRestjson1_serializeOpStopSpeakerSearchTask
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
@@ -2119,7 +2044,6 @@ func (m *awsRestjson1_serializeOpStopSpeakerSearchTask
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopSpeakerSearchTaskInput(v *StopSpeakerSearchTaskInput, encoder *httpbinding.Encoder) error {
@@ -2158,10 +2082,8 @@ func (m *awsRestjson1_serializeOpStopVoiceToneAnalysis
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
@@ -2199,7 +2121,6 @@ func (m *awsRestjson1_serializeOpStopVoiceToneAnalysis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopVoiceToneAnalysisTaskInput(v *StopVoiceToneAnalysisTaskInput, encoder *httpbinding.Encoder) error {
@@ -2238,10 +2159,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -2286,7 +2205,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2326,10 +2244,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -2374,7 +2290,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2414,10 +2329,8 @@ func (m *awsRestjson1_serializeOpUpdateMediaInsightsPi
 func (m *awsRestjson1_serializeOpUpdateMediaInsightsPipelineConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2466,7 +2379,6 @@ func (m *awsRestjson1_serializeOpUpdateMediaInsightsPi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMediaInsightsPipelineConfigurationInput(v *UpdateMediaInsightsPipelineConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2522,10 +2434,8 @@ func (m *awsRestjson1_serializeOpUpdateMediaInsightsPi
 func (m *awsRestjson1_serializeOpUpdateMediaInsightsPipelineStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2574,7 +2484,6 @@ func (m *awsRestjson1_serializeOpUpdateMediaInsightsPi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMediaInsightsPipelineStatusInput(v *UpdateMediaInsightsPipelineStatusInput, encoder *httpbinding.Encoder) error {
@@ -2616,10 +2525,8 @@ func (m *awsRestjson1_serializeOpUpdateMediaPipelineKi
 func (m *awsRestjson1_serializeOpUpdateMediaPipelineKinesisVideoStreamPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2668,7 +2575,6 @@ func (m *awsRestjson1_serializeOpUpdateMediaPipelineKi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMediaPipelineKinesisVideoStreamPoolInput(v *UpdateMediaPipelineKinesisVideoStreamPoolInput, encoder *httpbinding.Encoder) error {
