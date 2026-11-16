--- vendor/github.com/aws/aws-sdk-go-v2/service/ivsrealtime/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ivsrealtime/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCreateEncoderConfigur
 func (m *awsRestjson1_serializeOpCreateEncoderConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpCreateEncoderConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEncoderConfigurationInput(v *CreateEncoderConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -121,10 +117,8 @@ func (m *awsRestjson1_serializeOpCreateIngestConfigura
 func (m *awsRestjson1_serializeOpCreateIngestConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -169,7 +163,6 @@ func (m *awsRestjson1_serializeOpCreateIngestConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIngestConfigurationInput(v *CreateIngestConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -236,10 +229,8 @@ func (m *awsRestjson1_serializeOpCreateParticipantToke
 func (m *awsRestjson1_serializeOpCreateParticipantToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -284,7 +275,6 @@ func (m *awsRestjson1_serializeOpCreateParticipantToke
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateParticipantTokenInput(v *CreateParticipantTokenInput, encoder *httpbinding.Encoder) error {
@@ -341,10 +331,8 @@ func (m *awsRestjson1_serializeOpCreateStage) HandleSe
 func (m *awsRestjson1_serializeOpCreateStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -389,7 +377,6 @@ func (m *awsRestjson1_serializeOpCreateStage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateStageInput(v *CreateStageInput, encoder *httpbinding.Encoder) error {
@@ -443,10 +430,8 @@ func (m *awsRestjson1_serializeOpCreateStorageConfigur
 func (m *awsRestjson1_serializeOpCreateStorageConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -491,7 +476,6 @@ func (m *awsRestjson1_serializeOpCreateStorageConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateStorageConfigurationInput(v *CreateStorageConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -538,10 +522,8 @@ func (m *awsRestjson1_serializeOpDeleteEncoderConfigur
 func (m *awsRestjson1_serializeOpDeleteEncoderConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -586,7 +568,6 @@ func (m *awsRestjson1_serializeOpDeleteEncoderConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEncoderConfigurationInput(v *DeleteEncoderConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -619,10 +600,8 @@ func (m *awsRestjson1_serializeOpDeleteIngestConfigura
 func (m *awsRestjson1_serializeOpDeleteIngestConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -667,7 +646,6 @@ func (m *awsRestjson1_serializeOpDeleteIngestConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIngestConfigurationInput(v *DeleteIngestConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -705,10 +683,8 @@ func (m *awsRestjson1_serializeOpDeletePublicKey) Hand
 func (m *awsRestjson1_serializeOpDeletePublicKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -753,7 +729,6 @@ func (m *awsRestjson1_serializeOpDeletePublicKey) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePublicKeyInput(v *DeletePublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -786,10 +761,8 @@ func (m *awsRestjson1_serializeOpDeleteStage) HandleSe
 func (m *awsRestjson1_serializeOpDeleteStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -834,7 +807,6 @@ func (m *awsRestjson1_serializeOpDeleteStage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteStageInput(v *DeleteStageInput, encoder *httpbinding.Encoder) error {
@@ -867,10 +839,8 @@ func (m *awsRestjson1_serializeOpDeleteStorageConfigur
 func (m *awsRestjson1_serializeOpDeleteStorageConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -915,7 +885,6 @@ func (m *awsRestjson1_serializeOpDeleteStorageConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteStorageConfigurationInput(v *DeleteStorageConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -948,10 +917,8 @@ func (m *awsRestjson1_serializeOpDisconnectParticipant
 func (m *awsRestjson1_serializeOpDisconnectParticipant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -996,7 +963,6 @@ func (m *awsRestjson1_serializeOpDisconnectParticipant
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisconnectParticipantInput(v *DisconnectParticipantInput, encoder *httpbinding.Encoder) error {
@@ -1039,10 +1005,8 @@ func (m *awsRestjson1_serializeOpGetComposition) Handl
 func (m *awsRestjson1_serializeOpGetComposition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1087,7 +1051,6 @@ func (m *awsRestjson1_serializeOpGetComposition) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCompositionInput(v *GetCompositionInput, encoder *httpbinding.Encoder) error {
@@ -1120,10 +1083,8 @@ func (m *awsRestjson1_serializeOpGetEncoderConfigurati
 func (m *awsRestjson1_serializeOpGetEncoderConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1168,7 +1129,6 @@ func (m *awsRestjson1_serializeOpGetEncoderConfigurati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEncoderConfigurationInput(v *GetEncoderConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1201,10 +1161,8 @@ func (m *awsRestjson1_serializeOpGetIngestConfiguratio
 func (m *awsRestjson1_serializeOpGetIngestConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1249,7 +1207,6 @@ func (m *awsRestjson1_serializeOpGetIngestConfiguratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIngestConfigurationInput(v *GetIngestConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1282,10 +1239,8 @@ func (m *awsRestjson1_serializeOpGetParticipant) Handl
 func (m *awsRestjson1_serializeOpGetParticipant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1330,7 +1285,6 @@ func (m *awsRestjson1_serializeOpGetParticipant) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetParticipantInput(v *GetParticipantInput, encoder *httpbinding.Encoder) error {
@@ -1373,10 +1327,8 @@ func (m *awsRestjson1_serializeOpGetPublicKey) HandleS
 func (m *awsRestjson1_serializeOpGetPublicKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1421,7 +1373,6 @@ func (m *awsRestjson1_serializeOpGetPublicKey) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPublicKeyInput(v *GetPublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -1454,10 +1405,8 @@ func (m *awsRestjson1_serializeOpGetStage) HandleSeria
 func (m *awsRestjson1_serializeOpGetStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1502,7 +1451,6 @@ func (m *awsRestjson1_serializeOpGetStage) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStageInput(v *GetStageInput, encoder *httpbinding.Encoder) error {
@@ -1535,10 +1483,8 @@ func (m *awsRestjson1_serializeOpGetStageSession) Hand
 func (m *awsRestjson1_serializeOpGetStageSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1583,7 +1529,6 @@ func (m *awsRestjson1_serializeOpGetStageSession) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStageSessionInput(v *GetStageSessionInput, encoder *httpbinding.Encoder) error {
@@ -1621,10 +1566,8 @@ func (m *awsRestjson1_serializeOpGetStorageConfigurati
 func (m *awsRestjson1_serializeOpGetStorageConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1669,7 +1612,6 @@ func (m *awsRestjson1_serializeOpGetStorageConfigurati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStorageConfigurationInput(v *GetStorageConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1702,10 +1644,8 @@ func (m *awsRestjson1_serializeOpImportPublicKey) Hand
 func (m *awsRestjson1_serializeOpImportPublicKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1750,7 +1690,6 @@ func (m *awsRestjson1_serializeOpImportPublicKey) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportPublicKeyInput(v *ImportPublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -1795,10 +1734,8 @@ func (m *awsRestjson1_serializeOpListCompositions) Han
 func (m *awsRestjson1_serializeOpListCompositions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1843,7 +1780,6 @@ func (m *awsRestjson1_serializeOpListCompositions) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCompositionsInput(v *ListCompositionsInput, encoder *httpbinding.Encoder) error {
@@ -1891,10 +1827,8 @@ func (m *awsRestjson1_serializeOpListEncoderConfigurat
 func (m *awsRestjson1_serializeOpListEncoderConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1939,7 +1873,6 @@ func (m *awsRestjson1_serializeOpListEncoderConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEncoderConfigurationsInput(v *ListEncoderConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -1977,10 +1910,8 @@ func (m *awsRestjson1_serializeOpListIngestConfigurati
 func (m *awsRestjson1_serializeOpListIngestConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2025,7 +1956,6 @@ func (m *awsRestjson1_serializeOpListIngestConfigurati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIngestConfigurationsInput(v *ListIngestConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -2073,10 +2003,8 @@ func (m *awsRestjson1_serializeOpListParticipantEvents
 func (m *awsRestjson1_serializeOpListParticipantEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2121,7 +2049,6 @@ func (m *awsRestjson1_serializeOpListParticipantEvents
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListParticipantEventsInput(v *ListParticipantEventsInput, encoder *httpbinding.Encoder) error {
@@ -2174,10 +2101,8 @@ func (m *awsRestjson1_serializeOpListParticipantReplic
 func (m *awsRestjson1_serializeOpListParticipantReplicas) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2222,7 +2147,6 @@ func (m *awsRestjson1_serializeOpListParticipantReplic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListParticipantReplicasInput(v *ListParticipantReplicasInput, encoder *httpbinding.Encoder) error {
@@ -2270,10 +2194,8 @@ func (m *awsRestjson1_serializeOpListParticipants) Han
 func (m *awsRestjson1_serializeOpListParticipants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2318,7 +2240,6 @@ func (m *awsRestjson1_serializeOpListParticipants) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListParticipantsInput(v *ListParticipantsInput, encoder *httpbinding.Encoder) error {
@@ -2386,10 +2307,8 @@ func (m *awsRestjson1_serializeOpListPublicKeys) Handl
 func (m *awsRestjson1_serializeOpListPublicKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2434,7 +2353,6 @@ func (m *awsRestjson1_serializeOpListPublicKeys) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPublicKeysInput(v *ListPublicKeysInput, encoder *httpbinding.Encoder) error {
@@ -2472,10 +2390,8 @@ func (m *awsRestjson1_serializeOpListStages) HandleSer
 func (m *awsRestjson1_serializeOpListStages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2520,7 +2436,6 @@ func (m *awsRestjson1_serializeOpListStages) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStagesInput(v *ListStagesInput, encoder *httpbinding.Encoder) error {
@@ -2558,10 +2473,8 @@ func (m *awsRestjson1_serializeOpListStageSessions) Ha
 func (m *awsRestjson1_serializeOpListStageSessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2606,7 +2519,6 @@ func (m *awsRestjson1_serializeOpListStageSessions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStageSessionsInput(v *ListStageSessionsInput, encoder *httpbinding.Encoder) error {
@@ -2649,10 +2561,8 @@ func (m *awsRestjson1_serializeOpListStorageConfigurat
 func (m *awsRestjson1_serializeOpListStorageConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2697,7 +2607,6 @@ func (m *awsRestjson1_serializeOpListStorageConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStorageConfigurationsInput(v *ListStorageConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -2735,10 +2644,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -2776,7 +2683,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -2806,10 +2712,8 @@ func (m *awsRestjson1_serializeOpStartComposition) Han
 func (m *awsRestjson1_serializeOpStartComposition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2854,7 +2758,6 @@ func (m *awsRestjson1_serializeOpStartComposition) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartCompositionInput(v *StartCompositionInput, encoder *httpbinding.Encoder) error {
@@ -2913,10 +2816,8 @@ func (m *awsRestjson1_serializeOpStartParticipantRepli
 func (m *awsRestjson1_serializeOpStartParticipantReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2961,7 +2862,6 @@ func (m *awsRestjson1_serializeOpStartParticipantRepli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartParticipantReplicationInput(v *StartParticipantReplicationInput, encoder *httpbinding.Encoder) error {
@@ -3016,10 +2916,8 @@ func (m *awsRestjson1_serializeOpStopComposition) Hand
 func (m *awsRestjson1_serializeOpStopComposition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3064,7 +2962,6 @@ func (m *awsRestjson1_serializeOpStopComposition) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopCompositionInput(v *StopCompositionInput, encoder *httpbinding.Encoder) error {
@@ -3097,10 +2994,8 @@ func (m *awsRestjson1_serializeOpStopParticipantReplic
 func (m *awsRestjson1_serializeOpStopParticipantReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3145,7 +3040,6 @@ func (m *awsRestjson1_serializeOpStopParticipantReplic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopParticipantReplicationInput(v *StopParticipantReplicationInput, encoder *httpbinding.Encoder) error {
@@ -3188,10 +3082,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -3240,7 +3132,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3284,10 +3175,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -3325,7 +3214,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3361,10 +3249,8 @@ func (m *awsRestjson1_serializeOpUpdateIngestConfigura
 func (m *awsRestjson1_serializeOpUpdateIngestConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3409,7 +3295,6 @@ func (m *awsRestjson1_serializeOpUpdateIngestConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIngestConfigurationInput(v *UpdateIngestConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3447,10 +3332,8 @@ func (m *awsRestjson1_serializeOpUpdateStage) HandleSe
 func (m *awsRestjson1_serializeOpUpdateStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3495,7 +3378,6 @@ func (m *awsRestjson1_serializeOpUpdateStage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateStageInput(v *UpdateStageInput, encoder *httpbinding.Encoder) error {
