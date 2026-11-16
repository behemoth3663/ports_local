--- vendor/github.com/aws/aws-sdk-go-v2/service/cleanroomsml/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cleanroomsml/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpCancelTrainedModel) H
 func (m *awsRestjson1_serializeOpCancelTrainedModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -68,7 +65,6 @@ func (m *awsRestjson1_serializeOpCancelTrainedModel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelTrainedModelInput(v *CancelTrainedModelInput, encoder *httpbinding.Encoder) error {
@@ -111,10 +107,8 @@ func (m *awsRestjson1_serializeOpCancelTrainedModelInf
 func (m *awsRestjson1_serializeOpCancelTrainedModelInferenceJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -152,7 +146,6 @@ func (m *awsRestjson1_serializeOpCancelTrainedModelInf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelTrainedModelInferenceJobInput(v *CancelTrainedModelInferenceJobInput, encoder *httpbinding.Encoder) error {
@@ -191,10 +184,8 @@ func (m *awsRestjson1_serializeOpCreateAudienceModel) 
 func (m *awsRestjson1_serializeOpCreateAudienceModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -239,7 +230,6 @@ func (m *awsRestjson1_serializeOpCreateAudienceModel) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAudienceModelInput(v *CreateAudienceModelInput, encoder *httpbinding.Encoder) error {
@@ -304,10 +294,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredAudie
 func (m *awsRestjson1_serializeOpCreateConfiguredAudienceModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -352,7 +340,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredAudie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredAudienceModelInput(v *CreateConfiguredAudienceModelInput, encoder *httpbinding.Encoder) error {
@@ -433,10 +420,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredModel
 func (m *awsRestjson1_serializeOpCreateConfiguredModelAlgorithm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -481,7 +466,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredModel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredModelAlgorithmInput(v *CreateConfiguredModelAlgorithmInput, encoder *httpbinding.Encoder) error {
@@ -550,10 +534,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredModel
 func (m *awsRestjson1_serializeOpCreateConfiguredModelAlgorithmAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -602,7 +584,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredModel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredModelAlgorithmAssociationInput(v *CreateConfiguredModelAlgorithmAssociationInput, encoder *httpbinding.Encoder) error {
@@ -668,10 +649,8 @@ func (m *awsRestjson1_serializeOpCreateMLInputChannel)
 func (m *awsRestjson1_serializeOpCreateMLInputChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -720,7 +699,6 @@ func (m *awsRestjson1_serializeOpCreateMLInputChannel)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMLInputChannelInput(v *CreateMLInputChannelInput, encoder *httpbinding.Encoder) error {
@@ -798,10 +776,8 @@ func (m *awsRestjson1_serializeOpCreateTrainedModel) H
 func (m *awsRestjson1_serializeOpCreateTrainedModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -850,7 +826,6 @@ func (m *awsRestjson1_serializeOpCreateTrainedModel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTrainedModelInput(v *CreateTrainedModelInput, encoder *httpbinding.Encoder) error {
@@ -961,10 +936,8 @@ func (m *awsRestjson1_serializeOpCreateTrainingDataset
 func (m *awsRestjson1_serializeOpCreateTrainingDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1009,7 +982,6 @@ func (m *awsRestjson1_serializeOpCreateTrainingDataset
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTrainingDatasetInput(v *CreateTrainingDatasetInput, encoder *httpbinding.Encoder) error {
@@ -1066,10 +1038,8 @@ func (m *awsRestjson1_serializeOpDeleteAudienceGenerat
 func (m *awsRestjson1_serializeOpDeleteAudienceGenerationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1107,7 +1077,6 @@ func (m *awsRestjson1_serializeOpDeleteAudienceGenerat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAudienceGenerationJobInput(v *DeleteAudienceGenerationJobInput, encoder *httpbinding.Encoder) error {
@@ -1137,10 +1106,8 @@ func (m *awsRestjson1_serializeOpDeleteAudienceModel) 
 func (m *awsRestjson1_serializeOpDeleteAudienceModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1178,7 +1145,6 @@ func (m *awsRestjson1_serializeOpDeleteAudienceModel) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAudienceModelInput(v *DeleteAudienceModelInput, encoder *httpbinding.Encoder) error {
@@ -1208,10 +1174,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredAudie
 func (m *awsRestjson1_serializeOpDeleteConfiguredAudienceModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1249,7 +1213,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredAudie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredAudienceModelInput(v *DeleteConfiguredAudienceModelInput, encoder *httpbinding.Encoder) error {
@@ -1279,10 +1242,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredAudie
 func (m *awsRestjson1_serializeOpDeleteConfiguredAudienceModelPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1320,7 +1281,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredAudie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredAudienceModelPolicyInput(v *DeleteConfiguredAudienceModelPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1350,10 +1310,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredModel
 func (m *awsRestjson1_serializeOpDeleteConfiguredModelAlgorithm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1391,7 +1349,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredModel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredModelAlgorithmInput(v *DeleteConfiguredModelAlgorithmInput, encoder *httpbinding.Encoder) error {
@@ -1421,10 +1378,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredModel
 func (m *awsRestjson1_serializeOpDeleteConfiguredModelAlgorithmAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1462,7 +1417,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredModel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredModelAlgorithmAssociationInput(v *DeleteConfiguredModelAlgorithmAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1501,10 +1455,8 @@ func (m *awsRestjson1_serializeOpDeleteMLConfiguration
 func (m *awsRestjson1_serializeOpDeleteMLConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1542,7 +1494,6 @@ func (m *awsRestjson1_serializeOpDeleteMLConfiguration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMLConfigurationInput(v *DeleteMLConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1572,10 +1523,8 @@ func (m *awsRestjson1_serializeOpDeleteMLInputChannelD
 func (m *awsRestjson1_serializeOpDeleteMLInputChannelData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1613,7 +1562,6 @@ func (m *awsRestjson1_serializeOpDeleteMLInputChannelD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMLInputChannelDataInput(v *DeleteMLInputChannelDataInput, encoder *httpbinding.Encoder) error {
@@ -1652,10 +1600,8 @@ func (m *awsRestjson1_serializeOpDeleteTrainedModelOut
 func (m *awsRestjson1_serializeOpDeleteTrainedModelOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1693,7 +1639,6 @@ func (m *awsRestjson1_serializeOpDeleteTrainedModelOut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTrainedModelOutputInput(v *DeleteTrainedModelOutputInput, encoder *httpbinding.Encoder) error {
@@ -1736,10 +1681,8 @@ func (m *awsRestjson1_serializeOpDeleteTrainingDataset
 func (m *awsRestjson1_serializeOpDeleteTrainingDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1777,7 +1720,6 @@ func (m *awsRestjson1_serializeOpDeleteTrainingDataset
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTrainingDatasetInput(v *DeleteTrainingDatasetInput, encoder *httpbinding.Encoder) error {
@@ -1807,10 +1749,8 @@ func (m *awsRestjson1_serializeOpGetAudienceGeneration
 func (m *awsRestjson1_serializeOpGetAudienceGenerationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1848,7 +1788,6 @@ func (m *awsRestjson1_serializeOpGetAudienceGeneration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAudienceGenerationJobInput(v *GetAudienceGenerationJobInput, encoder *httpbinding.Encoder) error {
@@ -1878,10 +1817,8 @@ func (m *awsRestjson1_serializeOpGetAudienceModel) Han
 func (m *awsRestjson1_serializeOpGetAudienceModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1919,7 +1856,6 @@ func (m *awsRestjson1_serializeOpGetAudienceModel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAudienceModelInput(v *GetAudienceModelInput, encoder *httpbinding.Encoder) error {
@@ -1949,10 +1885,8 @@ func (m *awsRestjson1_serializeOpGetCollaborationConfi
 func (m *awsRestjson1_serializeOpGetCollaborationConfiguredModelAlgorithmAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1990,7 +1924,6 @@ func (m *awsRestjson1_serializeOpGetCollaborationConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationConfiguredModelAlgorithmAssociationInput(v *GetCollaborationConfiguredModelAlgorithmAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2029,10 +1962,8 @@ func (m *awsRestjson1_serializeOpGetCollaborationMLInp
 func (m *awsRestjson1_serializeOpGetCollaborationMLInputChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2070,7 +2001,6 @@ func (m *awsRestjson1_serializeOpGetCollaborationMLInp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationMLInputChannelInput(v *GetCollaborationMLInputChannelInput, encoder *httpbinding.Encoder) error {
@@ -2109,10 +2039,8 @@ func (m *awsRestjson1_serializeOpGetCollaborationTrain
 func (m *awsRestjson1_serializeOpGetCollaborationTrainedModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2150,7 +2078,6 @@ func (m *awsRestjson1_serializeOpGetCollaborationTrain
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationTrainedModelInput(v *GetCollaborationTrainedModelInput, encoder *httpbinding.Encoder) error {
@@ -2193,10 +2120,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredAudience
 func (m *awsRestjson1_serializeOpGetConfiguredAudienceModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2234,7 +2159,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredAudience
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredAudienceModelInput(v *GetConfiguredAudienceModelInput, encoder *httpbinding.Encoder) error {
@@ -2264,10 +2188,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredAudience
 func (m *awsRestjson1_serializeOpGetConfiguredAudienceModelPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2305,7 +2227,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredAudience
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredAudienceModelPolicyInput(v *GetConfiguredAudienceModelPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2335,10 +2256,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredModelAlg
 func (m *awsRestjson1_serializeOpGetConfiguredModelAlgorithm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2376,7 +2295,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredModelAlg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredModelAlgorithmInput(v *GetConfiguredModelAlgorithmInput, encoder *httpbinding.Encoder) error {
@@ -2406,10 +2324,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredModelAlg
 func (m *awsRestjson1_serializeOpGetConfiguredModelAlgorithmAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2447,7 +2363,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredModelAlg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredModelAlgorithmAssociationInput(v *GetConfiguredModelAlgorithmAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2486,10 +2401,8 @@ func (m *awsRestjson1_serializeOpGetMLConfiguration) H
 func (m *awsRestjson1_serializeOpGetMLConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2527,7 +2440,6 @@ func (m *awsRestjson1_serializeOpGetMLConfiguration) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMLConfigurationInput(v *GetMLConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2557,10 +2469,8 @@ func (m *awsRestjson1_serializeOpGetMLInputChannel) Ha
 func (m *awsRestjson1_serializeOpGetMLInputChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2598,7 +2508,6 @@ func (m *awsRestjson1_serializeOpGetMLInputChannel) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMLInputChannelInput(v *GetMLInputChannelInput, encoder *httpbinding.Encoder) error {
@@ -2637,10 +2546,8 @@ func (m *awsRestjson1_serializeOpGetTrainedModel) Hand
 func (m *awsRestjson1_serializeOpGetTrainedModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2678,7 +2585,6 @@ func (m *awsRestjson1_serializeOpGetTrainedModel) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTrainedModelInput(v *GetTrainedModelInput, encoder *httpbinding.Encoder) error {
@@ -2721,10 +2627,8 @@ func (m *awsRestjson1_serializeOpGetTrainedModelInfere
 func (m *awsRestjson1_serializeOpGetTrainedModelInferenceJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2762,7 +2666,6 @@ func (m *awsRestjson1_serializeOpGetTrainedModelInfere
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTrainedModelInferenceJobInput(v *GetTrainedModelInferenceJobInput, encoder *httpbinding.Encoder) error {
@@ -2801,10 +2704,8 @@ func (m *awsRestjson1_serializeOpGetTrainingDataset) H
 func (m *awsRestjson1_serializeOpGetTrainingDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2842,7 +2743,6 @@ func (m *awsRestjson1_serializeOpGetTrainingDataset) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTrainingDatasetInput(v *GetTrainingDatasetInput, encoder *httpbinding.Encoder) error {
@@ -2872,10 +2772,8 @@ func (m *awsRestjson1_serializeOpListAudienceExportJob
 func (m *awsRestjson1_serializeOpListAudienceExportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2913,7 +2811,6 @@ func (m *awsRestjson1_serializeOpListAudienceExportJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAudienceExportJobsInput(v *ListAudienceExportJobsInput, encoder *httpbinding.Encoder) error {
@@ -2946,10 +2843,8 @@ func (m *awsRestjson1_serializeOpListAudienceGeneratio
 func (m *awsRestjson1_serializeOpListAudienceGenerationJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2987,7 +2882,6 @@ func (m *awsRestjson1_serializeOpListAudienceGeneratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAudienceGenerationJobsInput(v *ListAudienceGenerationJobsInput, encoder *httpbinding.Encoder) error {
@@ -3024,10 +2918,8 @@ func (m *awsRestjson1_serializeOpListAudienceModels) H
 func (m *awsRestjson1_serializeOpListAudienceModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3065,7 +2957,6 @@ func (m *awsRestjson1_serializeOpListAudienceModels) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAudienceModelsInput(v *ListAudienceModelsInput, encoder *httpbinding.Encoder) error {
@@ -3094,10 +2985,8 @@ func (m *awsRestjson1_serializeOpListCollaborationConf
 func (m *awsRestjson1_serializeOpListCollaborationConfiguredModelAlgorithmAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3135,7 +3024,6 @@ func (m *awsRestjson1_serializeOpListCollaborationConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationConfiguredModelAlgorithmAssociationsInput(v *ListCollaborationConfiguredModelAlgorithmAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -3173,10 +3061,8 @@ func (m *awsRestjson1_serializeOpListCollaborationMLIn
 func (m *awsRestjson1_serializeOpListCollaborationMLInputChannels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3214,7 +3100,6 @@ func (m *awsRestjson1_serializeOpListCollaborationMLIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationMLInputChannelsInput(v *ListCollaborationMLInputChannelsInput, encoder *httpbinding.Encoder) error {
@@ -3252,10 +3137,8 @@ func (m *awsRestjson1_serializeOpListCollaborationTrai
 func (m *awsRestjson1_serializeOpListCollaborationTrainedModelExportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3293,7 +3176,6 @@ func (m *awsRestjson1_serializeOpListCollaborationTrai
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationTrainedModelExportJobsInput(v *ListCollaborationTrainedModelExportJobsInput, encoder *httpbinding.Encoder) error {
@@ -3344,10 +3226,8 @@ func (m *awsRestjson1_serializeOpListCollaborationTrai
 func (m *awsRestjson1_serializeOpListCollaborationTrainedModelInferenceJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3385,7 +3265,6 @@ func (m *awsRestjson1_serializeOpListCollaborationTrai
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationTrainedModelInferenceJobsInput(v *ListCollaborationTrainedModelInferenceJobsInput, encoder *httpbinding.Encoder) error {
@@ -3431,10 +3310,8 @@ func (m *awsRestjson1_serializeOpListCollaborationTrai
 func (m *awsRestjson1_serializeOpListCollaborationTrainedModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3472,7 +3349,6 @@ func (m *awsRestjson1_serializeOpListCollaborationTrai
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationTrainedModelsInput(v *ListCollaborationTrainedModelsInput, encoder *httpbinding.Encoder) error {
@@ -3510,10 +3386,8 @@ func (m *awsRestjson1_serializeOpListConfiguredAudienc
 func (m *awsRestjson1_serializeOpListConfiguredAudienceModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3551,7 +3425,6 @@ func (m *awsRestjson1_serializeOpListConfiguredAudienc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfiguredAudienceModelsInput(v *ListConfiguredAudienceModelsInput, encoder *httpbinding.Encoder) error {
@@ -3580,10 +3453,8 @@ func (m *awsRestjson1_serializeOpListConfiguredModelAl
 func (m *awsRestjson1_serializeOpListConfiguredModelAlgorithmAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3621,7 +3492,6 @@ func (m *awsRestjson1_serializeOpListConfiguredModelAl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfiguredModelAlgorithmAssociationsInput(v *ListConfiguredModelAlgorithmAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -3659,10 +3529,8 @@ func (m *awsRestjson1_serializeOpListConfiguredModelAl
 func (m *awsRestjson1_serializeOpListConfiguredModelAlgorithms) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3700,7 +3568,6 @@ func (m *awsRestjson1_serializeOpListConfiguredModelAl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfiguredModelAlgorithmsInput(v *ListConfiguredModelAlgorithmsInput, encoder *httpbinding.Encoder) error {
@@ -3729,10 +3596,8 @@ func (m *awsRestjson1_serializeOpListMLInputChannels) 
 func (m *awsRestjson1_serializeOpListMLInputChannels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3770,7 +3635,6 @@ func (m *awsRestjson1_serializeOpListMLInputChannels) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMLInputChannelsInput(v *ListMLInputChannelsInput, encoder *httpbinding.Encoder) error {
@@ -3808,10 +3672,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -3849,7 +3711,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -3879,10 +3740,8 @@ func (m *awsRestjson1_serializeOpListTrainedModelInfer
 func (m *awsRestjson1_serializeOpListTrainedModelInferenceJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3920,7 +3779,6 @@ func (m *awsRestjson1_serializeOpListTrainedModelInfer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrainedModelInferenceJobsInput(v *ListTrainedModelInferenceJobsInput, encoder *httpbinding.Encoder) error {
@@ -3966,10 +3824,8 @@ func (m *awsRestjson1_serializeOpListTrainedModels) Ha
 func (m *awsRestjson1_serializeOpListTrainedModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4007,7 +3863,6 @@ func (m *awsRestjson1_serializeOpListTrainedModels) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrainedModelsInput(v *ListTrainedModelsInput, encoder *httpbinding.Encoder) error {
@@ -4045,10 +3900,8 @@ func (m *awsRestjson1_serializeOpListTrainedModelVersi
 func (m *awsRestjson1_serializeOpListTrainedModelVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4086,7 +3939,6 @@ func (m *awsRestjson1_serializeOpListTrainedModelVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrainedModelVersionsInput(v *ListTrainedModelVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4137,10 +3989,8 @@ func (m *awsRestjson1_serializeOpListTrainingDatasets)
 func (m *awsRestjson1_serializeOpListTrainingDatasets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4178,7 +4028,6 @@ func (m *awsRestjson1_serializeOpListTrainingDatasets)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrainingDatasetsInput(v *ListTrainingDatasetsInput, encoder *httpbinding.Encoder) error {
@@ -4207,10 +4056,8 @@ func (m *awsRestjson1_serializeOpPutConfiguredAudience
 func (m *awsRestjson1_serializeOpPutConfiguredAudienceModelPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4259,7 +4106,6 @@ func (m *awsRestjson1_serializeOpPutConfiguredAudience
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfiguredAudienceModelPolicyInput(v *PutConfiguredAudienceModelPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4311,10 +4157,8 @@ func (m *awsRestjson1_serializeOpPutMLConfiguration) H
 func (m *awsRestjson1_serializeOpPutMLConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4363,7 +4207,6 @@ func (m *awsRestjson1_serializeOpPutMLConfiguration) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutMLConfigurationInput(v *PutMLConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4407,10 +4250,8 @@ func (m *awsRestjson1_serializeOpStartAudienceExportJo
 func (m *awsRestjson1_serializeOpStartAudienceExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4455,7 +4296,6 @@ func (m *awsRestjson1_serializeOpStartAudienceExportJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAudienceExportJobInput(v *StartAudienceExportJobInput, encoder *httpbinding.Encoder) error {
@@ -4505,10 +4345,8 @@ func (m *awsRestjson1_serializeOpStartAudienceGenerati
 func (m *awsRestjson1_serializeOpStartAudienceGenerationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4553,7 +4391,6 @@ func (m *awsRestjson1_serializeOpStartAudienceGenerati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAudienceGenerationJobInput(v *StartAudienceGenerationJobInput, encoder *httpbinding.Encoder) error {
@@ -4620,10 +4457,8 @@ func (m *awsRestjson1_serializeOpStartTrainedModelExpo
 func (m *awsRestjson1_serializeOpStartTrainedModelExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4672,7 +4507,6 @@ func (m *awsRestjson1_serializeOpStartTrainedModelExpo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTrainedModelExportJobInput(v *StartTrainedModelExportJobInput, encoder *httpbinding.Encoder) error {
@@ -4740,10 +4574,8 @@ func (m *awsRestjson1_serializeOpStartTrainedModelInfe
 func (m *awsRestjson1_serializeOpStartTrainedModelInferenceJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4792,7 +4624,6 @@ func (m *awsRestjson1_serializeOpStartTrainedModelInfe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTrainedModelInferenceJobInput(v *StartTrainedModelInferenceJobInput, encoder *httpbinding.Encoder) error {
@@ -4901,10 +4732,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -4953,7 +4782,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4997,10 +4825,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5038,7 +4864,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5074,10 +4899,8 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredAudie
 func (m *awsRestjson1_serializeOpUpdateConfiguredAudienceModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5126,7 +4949,6 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredAudie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfiguredAudienceModelInput(v *UpdateConfiguredAudienceModelInput, encoder *httpbinding.Encoder) error {
