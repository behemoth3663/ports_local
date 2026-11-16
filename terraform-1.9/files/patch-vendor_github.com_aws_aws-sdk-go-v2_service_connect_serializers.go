--- vendor/github.com/aws/aws-sdk-go-v2/service/connect/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/connect/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpActivateEvaluationFor
 func (m *awsRestjson1_serializeOpActivateEvaluationForm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpActivateEvaluationFor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsActivateEvaluationFormInput(v *ActivateEvaluationFormInput, encoder *httpbinding.Encoder) error {
@@ -130,10 +126,8 @@ func (m *awsRestjson1_serializeOpAssociateAnalyticsDat
 func (m *awsRestjson1_serializeOpAssociateAnalyticsDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -182,7 +176,6 @@ func (m *awsRestjson1_serializeOpAssociateAnalyticsDat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateAnalyticsDataSetInput(v *AssociateAnalyticsDataSetInput, encoder *httpbinding.Encoder) error {
@@ -229,10 +222,8 @@ func (m *awsRestjson1_serializeOpAssociateApprovedOrig
 func (m *awsRestjson1_serializeOpAssociateApprovedOrigin) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -281,7 +272,6 @@ func (m *awsRestjson1_serializeOpAssociateApprovedOrig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateApprovedOriginInput(v *AssociateApprovedOriginInput, encoder *httpbinding.Encoder) error {
@@ -328,10 +318,8 @@ func (m *awsRestjson1_serializeOpAssociateBot) HandleS
 func (m *awsRestjson1_serializeOpAssociateBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -380,7 +368,6 @@ func (m *awsRestjson1_serializeOpAssociateBot) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateBotInput(v *AssociateBotInput, encoder *httpbinding.Encoder) error {
@@ -436,10 +423,8 @@ func (m *awsRestjson1_serializeOpAssociateDefaultVocab
 func (m *awsRestjson1_serializeOpAssociateDefaultVocabulary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -488,7 +473,6 @@ func (m *awsRestjson1_serializeOpAssociateDefaultVocab
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateDefaultVocabularyInput(v *AssociateDefaultVocabularyInput, encoder *httpbinding.Encoder) error {
@@ -539,10 +523,8 @@ func (m *awsRestjson1_serializeOpAssociateFlow) Handle
 func (m *awsRestjson1_serializeOpAssociateFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -591,7 +573,6 @@ func (m *awsRestjson1_serializeOpAssociateFlow) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateFlowInput(v *AssociateFlowInput, encoder *httpbinding.Encoder) error {
@@ -643,10 +624,8 @@ func (m *awsRestjson1_serializeOpAssociateInstanceStor
 func (m *awsRestjson1_serializeOpAssociateInstanceStorageConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -695,7 +674,6 @@ func (m *awsRestjson1_serializeOpAssociateInstanceStor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateInstanceStorageConfigInput(v *AssociateInstanceStorageConfigInput, encoder *httpbinding.Encoder) error {
@@ -749,10 +727,8 @@ func (m *awsRestjson1_serializeOpAssociateLambdaFuncti
 func (m *awsRestjson1_serializeOpAssociateLambdaFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -801,7 +777,6 @@ func (m *awsRestjson1_serializeOpAssociateLambdaFuncti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateLambdaFunctionInput(v *AssociateLambdaFunctionInput, encoder *httpbinding.Encoder) error {
@@ -848,10 +823,8 @@ func (m *awsRestjson1_serializeOpAssociateLexBot) Hand
 func (m *awsRestjson1_serializeOpAssociateLexBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -900,7 +873,6 @@ func (m *awsRestjson1_serializeOpAssociateLexBot) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateLexBotInput(v *AssociateLexBotInput, encoder *httpbinding.Encoder) error {
@@ -949,10 +921,8 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumberC
 func (m *awsRestjson1_serializeOpAssociatePhoneNumberContactFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1001,7 +971,6 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumberC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociatePhoneNumberContactFlowInput(v *AssociatePhoneNumberContactFlowInput, encoder *httpbinding.Encoder) error {
@@ -1048,10 +1017,8 @@ func (m *awsRestjson1_serializeOpAssociateQueueQuickCo
 func (m *awsRestjson1_serializeOpAssociateQueueQuickConnects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1100,7 +1067,6 @@ func (m *awsRestjson1_serializeOpAssociateQueueQuickCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateQueueQuickConnectsInput(v *AssociateQueueQuickConnectsInput, encoder *httpbinding.Encoder) error {
@@ -1153,10 +1119,8 @@ func (m *awsRestjson1_serializeOpAssociateRoutingProfi
 func (m *awsRestjson1_serializeOpAssociateRoutingProfileQueues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1205,7 +1169,6 @@ func (m *awsRestjson1_serializeOpAssociateRoutingProfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateRoutingProfileQueuesInput(v *AssociateRoutingProfileQueuesInput, encoder *httpbinding.Encoder) error {
@@ -1258,10 +1221,8 @@ func (m *awsRestjson1_serializeOpAssociateSecurityKey)
 func (m *awsRestjson1_serializeOpAssociateSecurityKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1310,7 +1271,6 @@ func (m *awsRestjson1_serializeOpAssociateSecurityKey)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateSecurityKeyInput(v *AssociateSecurityKeyInput, encoder *httpbinding.Encoder) error {
@@ -1357,10 +1317,8 @@ func (m *awsRestjson1_serializeOpAssociateTrafficDistr
 func (m *awsRestjson1_serializeOpAssociateTrafficDistributionGroupUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1409,7 +1367,6 @@ func (m *awsRestjson1_serializeOpAssociateTrafficDistr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateTrafficDistributionGroupUserInput(v *AssociateTrafficDistributionGroupUserInput, encoder *httpbinding.Encoder) error {
@@ -1456,10 +1413,8 @@ func (m *awsRestjson1_serializeOpAssociateUserProficie
 func (m *awsRestjson1_serializeOpAssociateUserProficiencies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1508,7 +1463,6 @@ func (m *awsRestjson1_serializeOpAssociateUserProficie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateUserProficienciesInput(v *AssociateUserProficienciesInput, encoder *httpbinding.Encoder) error {
@@ -1561,10 +1515,8 @@ func (m *awsRestjson1_serializeOpBatchAssociateAnalyti
 func (m *awsRestjson1_serializeOpBatchAssociateAnalyticsDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1613,7 +1565,6 @@ func (m *awsRestjson1_serializeOpBatchAssociateAnalyti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchAssociateAnalyticsDataSetInput(v *BatchAssociateAnalyticsDataSetInput, encoder *httpbinding.Encoder) error {
@@ -1662,10 +1613,8 @@ func (m *awsRestjson1_serializeOpBatchDisassociateAnal
 func (m *awsRestjson1_serializeOpBatchDisassociateAnalyticsDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1714,7 +1663,6 @@ func (m *awsRestjson1_serializeOpBatchDisassociateAnal
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDisassociateAnalyticsDataSetInput(v *BatchDisassociateAnalyticsDataSetInput, encoder *httpbinding.Encoder) error {
@@ -1763,10 +1711,8 @@ func (m *awsRestjson1_serializeOpBatchGetAttachedFileM
 func (m *awsRestjson1_serializeOpBatchGetAttachedFileMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1815,7 +1761,6 @@ func (m *awsRestjson1_serializeOpBatchGetAttachedFileM
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetAttachedFileMetadataInput(v *BatchGetAttachedFileMetadataInput, encoder *httpbinding.Encoder) error {
@@ -1863,10 +1808,8 @@ func (m *awsRestjson1_serializeOpBatchGetFlowAssociati
 func (m *awsRestjson1_serializeOpBatchGetFlowAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1915,7 +1858,6 @@ func (m *awsRestjson1_serializeOpBatchGetFlowAssociati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetFlowAssociationInput(v *BatchGetFlowAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1964,10 +1906,8 @@ func (m *awsRestjson1_serializeOpBatchPutContact) Hand
 func (m *awsRestjson1_serializeOpBatchPutContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2016,7 +1956,6 @@ func (m *awsRestjson1_serializeOpBatchPutContact) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchPutContactInput(v *BatchPutContactInput, encoder *httpbinding.Encoder) error {
@@ -2065,10 +2004,8 @@ func (m *awsRestjson1_serializeOpClaimPhoneNumber) Han
 func (m *awsRestjson1_serializeOpClaimPhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2113,7 +2050,6 @@ func (m *awsRestjson1_serializeOpClaimPhoneNumber) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsClaimPhoneNumberInput(v *ClaimPhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -2173,10 +2109,8 @@ func (m *awsRestjson1_serializeOpCompleteAttachedFileU
 func (m *awsRestjson1_serializeOpCompleteAttachedFileUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2214,7 +2148,6 @@ func (m *awsRestjson1_serializeOpCompleteAttachedFileU
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCompleteAttachedFileUploadInput(v *CompleteAttachedFileUploadInput, encoder *httpbinding.Encoder) error {
@@ -2257,10 +2190,8 @@ func (m *awsRestjson1_serializeOpCreateAgentStatus) Ha
 func (m *awsRestjson1_serializeOpCreateAgentStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2309,7 +2240,6 @@ func (m *awsRestjson1_serializeOpCreateAgentStatus) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAgentStatusInput(v *CreateAgentStatusInput, encoder *httpbinding.Encoder) error {
@@ -2373,10 +2303,8 @@ func (m *awsRestjson1_serializeOpCreateContact) Handle
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
@@ -2421,7 +2349,6 @@ func (m *awsRestjson1_serializeOpCreateContact) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContactInput(v *CreateContactInput, encoder *httpbinding.Encoder) error {
@@ -2527,10 +2454,8 @@ func (m *awsRestjson1_serializeOpCreateContactFlow) Ha
 func (m *awsRestjson1_serializeOpCreateContactFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2579,7 +2504,6 @@ func (m *awsRestjson1_serializeOpCreateContactFlow) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContactFlowInput(v *CreateContactFlowInput, encoder *httpbinding.Encoder) error {
@@ -2648,10 +2572,8 @@ func (m *awsRestjson1_serializeOpCreateContactFlowModu
 func (m *awsRestjson1_serializeOpCreateContactFlowModule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2700,7 +2622,6 @@ func (m *awsRestjson1_serializeOpCreateContactFlowModu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContactFlowModuleInput(v *CreateContactFlowModuleInput, encoder *httpbinding.Encoder) error {
@@ -2764,10 +2685,8 @@ func (m *awsRestjson1_serializeOpCreateContactFlowVers
 func (m *awsRestjson1_serializeOpCreateContactFlowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2816,7 +2735,6 @@ func (m *awsRestjson1_serializeOpCreateContactFlowVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContactFlowVersionInput(v *CreateContactFlowVersionInput, encoder *httpbinding.Encoder) error {
@@ -2887,10 +2805,8 @@ func (m *awsRestjson1_serializeOpCreateEmailAddress) H
 func (m *awsRestjson1_serializeOpCreateEmailAddress) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2939,7 +2855,6 @@ func (m *awsRestjson1_serializeOpCreateEmailAddress) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEmailAddressInput(v *CreateEmailAddressInput, encoder *httpbinding.Encoder) error {
@@ -3003,10 +2918,8 @@ func (m *awsRestjson1_serializeOpCreateEvaluationForm)
 func (m *awsRestjson1_serializeOpCreateEvaluationForm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3055,7 +2968,6 @@ func (m *awsRestjson1_serializeOpCreateEvaluationForm)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEvaluationFormInput(v *CreateEvaluationFormInput, encoder *httpbinding.Encoder) error {
@@ -3121,10 +3033,8 @@ func (m *awsRestjson1_serializeOpCreateHoursOfOperatio
 func (m *awsRestjson1_serializeOpCreateHoursOfOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3173,7 +3083,6 @@ func (m *awsRestjson1_serializeOpCreateHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateHoursOfOperationInput(v *CreateHoursOfOperationInput, encoder *httpbinding.Encoder) error {
@@ -3239,10 +3148,8 @@ func (m *awsRestjson1_serializeOpCreateHoursOfOperatio
 func (m *awsRestjson1_serializeOpCreateHoursOfOperationOverride) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3291,7 +3198,6 @@ func (m *awsRestjson1_serializeOpCreateHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateHoursOfOperationOverrideInput(v *CreateHoursOfOperationOverrideInput, encoder *httpbinding.Encoder) error {
@@ -3364,10 +3270,8 @@ func (m *awsRestjson1_serializeOpCreateInstance) Handl
 func (m *awsRestjson1_serializeOpCreateInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3412,7 +3316,6 @@ func (m *awsRestjson1_serializeOpCreateInstance) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInstanceInput(v *CreateInstanceInput, encoder *httpbinding.Encoder) error {
@@ -3477,10 +3380,8 @@ func (m *awsRestjson1_serializeOpCreateIntegrationAsso
 func (m *awsRestjson1_serializeOpCreateIntegrationAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3529,7 +3430,6 @@ func (m *awsRestjson1_serializeOpCreateIntegrationAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIntegrationAssociationInput(v *CreateIntegrationAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3598,10 +3498,8 @@ func (m *awsRestjson1_serializeOpCreateParticipant) Ha
 func (m *awsRestjson1_serializeOpCreateParticipant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3646,7 +3544,6 @@ func (m *awsRestjson1_serializeOpCreateParticipant) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateParticipantInput(v *CreateParticipantInput, encoder *httpbinding.Encoder) error {
@@ -3696,10 +3593,8 @@ func (m *awsRestjson1_serializeOpCreatePersistentConta
 func (m *awsRestjson1_serializeOpCreatePersistentContactAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3748,7 +3643,6 @@ func (m *awsRestjson1_serializeOpCreatePersistentConta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePersistentContactAssociationInput(v *CreatePersistentContactAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3809,10 +3703,8 @@ func (m *awsRestjson1_serializeOpCreatePredefinedAttri
 func (m *awsRestjson1_serializeOpCreatePredefinedAttribute) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3861,7 +3753,6 @@ func (m *awsRestjson1_serializeOpCreatePredefinedAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePredefinedAttributeInput(v *CreatePredefinedAttributeInput, encoder *httpbinding.Encoder) error {
@@ -3910,10 +3801,8 @@ func (m *awsRestjson1_serializeOpCreatePrompt) HandleS
 func (m *awsRestjson1_serializeOpCreatePrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3962,7 +3851,6 @@ func (m *awsRestjson1_serializeOpCreatePrompt) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePromptInput(v *CreatePromptInput, encoder *httpbinding.Encoder) error {
@@ -4021,10 +3909,8 @@ func (m *awsRestjson1_serializeOpCreatePushNotificatio
 func (m *awsRestjson1_serializeOpCreatePushNotificationRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4073,7 +3959,6 @@ func (m *awsRestjson1_serializeOpCreatePushNotificatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePushNotificationRegistrationInput(v *CreatePushNotificationRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -4137,10 +4022,8 @@ func (m *awsRestjson1_serializeOpCreateQueue) HandleSe
 func (m *awsRestjson1_serializeOpCreateQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4189,7 +4072,6 @@ func (m *awsRestjson1_serializeOpCreateQueue) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateQueueInput(v *CreateQueueInput, encoder *httpbinding.Encoder) error {
@@ -4274,10 +4156,8 @@ func (m *awsRestjson1_serializeOpCreateQuickConnect) H
 func (m *awsRestjson1_serializeOpCreateQuickConnect) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4326,7 +4206,6 @@ func (m *awsRestjson1_serializeOpCreateQuickConnect) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateQuickConnectInput(v *CreateQuickConnectInput, encoder *httpbinding.Encoder) error {
@@ -4387,10 +4266,8 @@ func (m *awsRestjson1_serializeOpCreateRoutingProfile)
 func (m *awsRestjson1_serializeOpCreateRoutingProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4439,7 +4316,6 @@ func (m *awsRestjson1_serializeOpCreateRoutingProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRoutingProfileInput(v *CreateRoutingProfileInput, encoder *httpbinding.Encoder) error {
@@ -4517,10 +4393,8 @@ func (m *awsRestjson1_serializeOpCreateRule) HandleSer
 func (m *awsRestjson1_serializeOpCreateRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4569,7 +4443,6 @@ func (m *awsRestjson1_serializeOpCreateRule) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRuleInput(v *CreateRuleInput, encoder *httpbinding.Encoder) error {
@@ -4640,10 +4513,8 @@ func (m *awsRestjson1_serializeOpCreateSecurityProfile
 func (m *awsRestjson1_serializeOpCreateSecurityProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4692,7 +4563,6 @@ func (m *awsRestjson1_serializeOpCreateSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSecurityProfileInput(v *CreateSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -4786,10 +4656,8 @@ func (m *awsRestjson1_serializeOpCreateTaskTemplate) H
 func (m *awsRestjson1_serializeOpCreateTaskTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4838,7 +4706,6 @@ func (m *awsRestjson1_serializeOpCreateTaskTemplate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTaskTemplateInput(v *CreateTaskTemplateInput, encoder *httpbinding.Encoder) error {
@@ -4926,10 +4793,8 @@ func (m *awsRestjson1_serializeOpCreateTrafficDistribu
 func (m *awsRestjson1_serializeOpCreateTrafficDistributionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4974,7 +4839,6 @@ func (m *awsRestjson1_serializeOpCreateTrafficDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTrafficDistributionGroupInput(v *CreateTrafficDistributionGroupInput, encoder *httpbinding.Encoder) error {
@@ -5029,10 +4893,8 @@ func (m *awsRestjson1_serializeOpCreateUseCase) Handle
 func (m *awsRestjson1_serializeOpCreateUseCase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5081,7 +4943,6 @@ func (m *awsRestjson1_serializeOpCreateUseCase) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUseCaseInput(v *CreateUseCaseInput, encoder *httpbinding.Encoder) error {
@@ -5139,10 +5000,8 @@ func (m *awsRestjson1_serializeOpCreateUser) HandleSer
 func (m *awsRestjson1_serializeOpCreateUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5191,7 +5050,6 @@ func (m *awsRestjson1_serializeOpCreateUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUserInput(v *CreateUserInput, encoder *httpbinding.Encoder) error {
@@ -5281,10 +5139,8 @@ func (m *awsRestjson1_serializeOpCreateUserHierarchyGr
 func (m *awsRestjson1_serializeOpCreateUserHierarchyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5333,7 +5189,6 @@ func (m *awsRestjson1_serializeOpCreateUserHierarchyGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUserHierarchyGroupInput(v *CreateUserHierarchyGroupInput, encoder *httpbinding.Encoder) error {
@@ -5387,10 +5242,8 @@ func (m *awsRestjson1_serializeOpCreateView) HandleSer
 func (m *awsRestjson1_serializeOpCreateView) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5439,7 +5292,6 @@ func (m *awsRestjson1_serializeOpCreateView) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateViewInput(v *CreateViewInput, encoder *httpbinding.Encoder) error {
@@ -5510,10 +5362,8 @@ func (m *awsRestjson1_serializeOpCreateViewVersion) Ha
 func (m *awsRestjson1_serializeOpCreateViewVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5562,7 +5412,6 @@ func (m *awsRestjson1_serializeOpCreateViewVersion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateViewVersionInput(v *CreateViewVersionInput, encoder *httpbinding.Encoder) error {
@@ -5618,10 +5467,8 @@ func (m *awsRestjson1_serializeOpCreateVocabulary) Han
 func (m *awsRestjson1_serializeOpCreateVocabulary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5670,7 +5517,6 @@ func (m *awsRestjson1_serializeOpCreateVocabulary) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVocabularyInput(v *CreateVocabularyInput, encoder *httpbinding.Encoder) error {
@@ -5734,10 +5580,8 @@ func (m *awsRestjson1_serializeOpDeactivateEvaluationF
 func (m *awsRestjson1_serializeOpDeactivateEvaluationForm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5786,7 +5630,6 @@ func (m *awsRestjson1_serializeOpDeactivateEvaluationF
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeactivateEvaluationFormInput(v *DeactivateEvaluationFormInput, encoder *httpbinding.Encoder) error {
@@ -5837,10 +5680,8 @@ func (m *awsRestjson1_serializeOpDeleteAttachedFile) H
 func (m *awsRestjson1_serializeOpDeleteAttachedFile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5878,7 +5719,6 @@ func (m *awsRestjson1_serializeOpDeleteAttachedFile) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAttachedFileInput(v *DeleteAttachedFileInput, encoder *httpbinding.Encoder) error {
@@ -5921,10 +5761,8 @@ func (m *awsRestjson1_serializeOpDeleteContactEvaluati
 func (m *awsRestjson1_serializeOpDeleteContactEvaluation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5962,7 +5800,6 @@ func (m *awsRestjson1_serializeOpDeleteContactEvaluati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContactEvaluationInput(v *DeleteContactEvaluationInput, encoder *httpbinding.Encoder) error {
@@ -6001,10 +5838,8 @@ func (m *awsRestjson1_serializeOpDeleteContactFlow) Ha
 func (m *awsRestjson1_serializeOpDeleteContactFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6042,7 +5877,6 @@ func (m *awsRestjson1_serializeOpDeleteContactFlow) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContactFlowInput(v *DeleteContactFlowInput, encoder *httpbinding.Encoder) error {
@@ -6081,10 +5915,8 @@ func (m *awsRestjson1_serializeOpDeleteContactFlowModu
 func (m *awsRestjson1_serializeOpDeleteContactFlowModule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6122,7 +5954,6 @@ func (m *awsRestjson1_serializeOpDeleteContactFlowModu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContactFlowModuleInput(v *DeleteContactFlowModuleInput, encoder *httpbinding.Encoder) error {
@@ -6161,10 +5992,8 @@ func (m *awsRestjson1_serializeOpDeleteContactFlowVers
 func (m *awsRestjson1_serializeOpDeleteContactFlowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6202,7 +6031,6 @@ func (m *awsRestjson1_serializeOpDeleteContactFlowVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContactFlowVersionInput(v *DeleteContactFlowVersionInput, encoder *httpbinding.Encoder) error {
@@ -6250,10 +6078,8 @@ func (m *awsRestjson1_serializeOpDeleteEmailAddress) H
 func (m *awsRestjson1_serializeOpDeleteEmailAddress) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6291,7 +6117,6 @@ func (m *awsRestjson1_serializeOpDeleteEmailAddress) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEmailAddressInput(v *DeleteEmailAddressInput, encoder *httpbinding.Encoder) error {
@@ -6330,10 +6155,8 @@ func (m *awsRestjson1_serializeOpDeleteEvaluationForm)
 func (m *awsRestjson1_serializeOpDeleteEvaluationForm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6371,7 +6194,6 @@ func (m *awsRestjson1_serializeOpDeleteEvaluationForm)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEvaluationFormInput(v *DeleteEvaluationFormInput, encoder *httpbinding.Encoder) error {
@@ -6414,10 +6236,8 @@ func (m *awsRestjson1_serializeOpDeleteHoursOfOperatio
 func (m *awsRestjson1_serializeOpDeleteHoursOfOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6455,7 +6275,6 @@ func (m *awsRestjson1_serializeOpDeleteHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteHoursOfOperationInput(v *DeleteHoursOfOperationInput, encoder *httpbinding.Encoder) error {
@@ -6494,10 +6313,8 @@ func (m *awsRestjson1_serializeOpDeleteHoursOfOperatio
 func (m *awsRestjson1_serializeOpDeleteHoursOfOperationOverride) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6535,7 +6352,6 @@ func (m *awsRestjson1_serializeOpDeleteHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteHoursOfOperationOverrideInput(v *DeleteHoursOfOperationOverrideInput, encoder *httpbinding.Encoder) error {
@@ -6583,10 +6399,8 @@ func (m *awsRestjson1_serializeOpDeleteInstance) Handl
 func (m *awsRestjson1_serializeOpDeleteInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6624,7 +6438,6 @@ func (m *awsRestjson1_serializeOpDeleteInstance) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInstanceInput(v *DeleteInstanceInput, encoder *httpbinding.Encoder) error {
@@ -6658,10 +6471,8 @@ func (m *awsRestjson1_serializeOpDeleteIntegrationAsso
 func (m *awsRestjson1_serializeOpDeleteIntegrationAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6699,7 +6510,6 @@ func (m *awsRestjson1_serializeOpDeleteIntegrationAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIntegrationAssociationInput(v *DeleteIntegrationAssociationInput, encoder *httpbinding.Encoder) error {
@@ -6738,10 +6548,8 @@ func (m *awsRestjson1_serializeOpDeletePredefinedAttri
 func (m *awsRestjson1_serializeOpDeletePredefinedAttribute) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6779,7 +6587,6 @@ func (m *awsRestjson1_serializeOpDeletePredefinedAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePredefinedAttributeInput(v *DeletePredefinedAttributeInput, encoder *httpbinding.Encoder) error {
@@ -6818,10 +6625,8 @@ func (m *awsRestjson1_serializeOpDeletePrompt) HandleS
 func (m *awsRestjson1_serializeOpDeletePrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6859,7 +6664,6 @@ func (m *awsRestjson1_serializeOpDeletePrompt) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePromptInput(v *DeletePromptInput, encoder *httpbinding.Encoder) error {
@@ -6898,10 +6702,8 @@ func (m *awsRestjson1_serializeOpDeletePushNotificatio
 func (m *awsRestjson1_serializeOpDeletePushNotificationRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6939,7 +6741,6 @@ func (m *awsRestjson1_serializeOpDeletePushNotificatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePushNotificationRegistrationInput(v *DeletePushNotificationRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -6982,10 +6783,8 @@ func (m *awsRestjson1_serializeOpDeleteQueue) HandleSe
 func (m *awsRestjson1_serializeOpDeleteQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7023,7 +6822,6 @@ func (m *awsRestjson1_serializeOpDeleteQueue) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQueueInput(v *DeleteQueueInput, encoder *httpbinding.Encoder) error {
@@ -7062,10 +6860,8 @@ func (m *awsRestjson1_serializeOpDeleteQuickConnect) H
 func (m *awsRestjson1_serializeOpDeleteQuickConnect) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7103,7 +6899,6 @@ func (m *awsRestjson1_serializeOpDeleteQuickConnect) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQuickConnectInput(v *DeleteQuickConnectInput, encoder *httpbinding.Encoder) error {
@@ -7142,10 +6937,8 @@ func (m *awsRestjson1_serializeOpDeleteRoutingProfile)
 func (m *awsRestjson1_serializeOpDeleteRoutingProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7183,7 +6976,6 @@ func (m *awsRestjson1_serializeOpDeleteRoutingProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRoutingProfileInput(v *DeleteRoutingProfileInput, encoder *httpbinding.Encoder) error {
@@ -7222,10 +7014,8 @@ func (m *awsRestjson1_serializeOpDeleteRule) HandleSer
 func (m *awsRestjson1_serializeOpDeleteRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7263,7 +7053,6 @@ func (m *awsRestjson1_serializeOpDeleteRule) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRuleInput(v *DeleteRuleInput, encoder *httpbinding.Encoder) error {
@@ -7302,10 +7091,8 @@ func (m *awsRestjson1_serializeOpDeleteSecurityProfile
 func (m *awsRestjson1_serializeOpDeleteSecurityProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7343,7 +7130,6 @@ func (m *awsRestjson1_serializeOpDeleteSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSecurityProfileInput(v *DeleteSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -7382,10 +7168,8 @@ func (m *awsRestjson1_serializeOpDeleteTaskTemplate) H
 func (m *awsRestjson1_serializeOpDeleteTaskTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7423,7 +7207,6 @@ func (m *awsRestjson1_serializeOpDeleteTaskTemplate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTaskTemplateInput(v *DeleteTaskTemplateInput, encoder *httpbinding.Encoder) error {
@@ -7462,10 +7245,8 @@ func (m *awsRestjson1_serializeOpDeleteTrafficDistribu
 func (m *awsRestjson1_serializeOpDeleteTrafficDistributionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7503,7 +7284,6 @@ func (m *awsRestjson1_serializeOpDeleteTrafficDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTrafficDistributionGroupInput(v *DeleteTrafficDistributionGroupInput, encoder *httpbinding.Encoder) error {
@@ -7533,10 +7313,8 @@ func (m *awsRestjson1_serializeOpDeleteUseCase) Handle
 func (m *awsRestjson1_serializeOpDeleteUseCase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7574,7 +7352,6 @@ func (m *awsRestjson1_serializeOpDeleteUseCase) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUseCaseInput(v *DeleteUseCaseInput, encoder *httpbinding.Encoder) error {
@@ -7622,10 +7399,8 @@ func (m *awsRestjson1_serializeOpDeleteUser) HandleSer
 func (m *awsRestjson1_serializeOpDeleteUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7663,7 +7438,6 @@ func (m *awsRestjson1_serializeOpDeleteUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserInput(v *DeleteUserInput, encoder *httpbinding.Encoder) error {
@@ -7702,10 +7476,8 @@ func (m *awsRestjson1_serializeOpDeleteUserHierarchyGr
 func (m *awsRestjson1_serializeOpDeleteUserHierarchyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7743,7 +7515,6 @@ func (m *awsRestjson1_serializeOpDeleteUserHierarchyGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserHierarchyGroupInput(v *DeleteUserHierarchyGroupInput, encoder *httpbinding.Encoder) error {
@@ -7782,10 +7553,8 @@ func (m *awsRestjson1_serializeOpDeleteView) HandleSer
 func (m *awsRestjson1_serializeOpDeleteView) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7823,7 +7592,6 @@ func (m *awsRestjson1_serializeOpDeleteView) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteViewInput(v *DeleteViewInput, encoder *httpbinding.Encoder) error {
@@ -7862,10 +7630,8 @@ func (m *awsRestjson1_serializeOpDeleteViewVersion) Ha
 func (m *awsRestjson1_serializeOpDeleteViewVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7903,7 +7669,6 @@ func (m *awsRestjson1_serializeOpDeleteViewVersion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteViewVersionInput(v *DeleteViewVersionInput, encoder *httpbinding.Encoder) error {
@@ -7951,10 +7716,8 @@ func (m *awsRestjson1_serializeOpDeleteVocabulary) Han
 func (m *awsRestjson1_serializeOpDeleteVocabulary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7992,7 +7755,6 @@ func (m *awsRestjson1_serializeOpDeleteVocabulary) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVocabularyInput(v *DeleteVocabularyInput, encoder *httpbinding.Encoder) error {
@@ -8031,10 +7793,8 @@ func (m *awsRestjson1_serializeOpDescribeAgentStatus) 
 func (m *awsRestjson1_serializeOpDescribeAgentStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8072,7 +7832,6 @@ func (m *awsRestjson1_serializeOpDescribeAgentStatus) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAgentStatusInput(v *DescribeAgentStatusInput, encoder *httpbinding.Encoder) error {
@@ -8111,10 +7870,8 @@ func (m *awsRestjson1_serializeOpDescribeAuthenticatio
 func (m *awsRestjson1_serializeOpDescribeAuthenticationProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8152,7 +7909,6 @@ func (m *awsRestjson1_serializeOpDescribeAuthenticatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAuthenticationProfileInput(v *DescribeAuthenticationProfileInput, encoder *httpbinding.Encoder) error {
@@ -8191,10 +7947,8 @@ func (m *awsRestjson1_serializeOpDescribeContact) Hand
 func (m *awsRestjson1_serializeOpDescribeContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8232,7 +7986,6 @@ func (m *awsRestjson1_serializeOpDescribeContact) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeContactInput(v *DescribeContactInput, encoder *httpbinding.Encoder) error {
@@ -8271,10 +8024,8 @@ func (m *awsRestjson1_serializeOpDescribeContactEvalua
 func (m *awsRestjson1_serializeOpDescribeContactEvaluation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8312,7 +8063,6 @@ func (m *awsRestjson1_serializeOpDescribeContactEvalua
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeContactEvaluationInput(v *DescribeContactEvaluationInput, encoder *httpbinding.Encoder) error {
@@ -8351,10 +8101,8 @@ func (m *awsRestjson1_serializeOpDescribeContactFlow) 
 func (m *awsRestjson1_serializeOpDescribeContactFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8392,7 +8140,6 @@ func (m *awsRestjson1_serializeOpDescribeContactFlow) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeContactFlowInput(v *DescribeContactFlowInput, encoder *httpbinding.Encoder) error {
@@ -8431,10 +8178,8 @@ func (m *awsRestjson1_serializeOpDescribeContactFlowMo
 func (m *awsRestjson1_serializeOpDescribeContactFlowModule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8472,7 +8217,6 @@ func (m *awsRestjson1_serializeOpDescribeContactFlowMo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeContactFlowModuleInput(v *DescribeContactFlowModuleInput, encoder *httpbinding.Encoder) error {
@@ -8511,10 +8255,8 @@ func (m *awsRestjson1_serializeOpDescribeEmailAddress)
 func (m *awsRestjson1_serializeOpDescribeEmailAddress) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8552,7 +8294,6 @@ func (m *awsRestjson1_serializeOpDescribeEmailAddress)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeEmailAddressInput(v *DescribeEmailAddressInput, encoder *httpbinding.Encoder) error {
@@ -8591,10 +8332,8 @@ func (m *awsRestjson1_serializeOpDescribeEvaluationFor
 func (m *awsRestjson1_serializeOpDescribeEvaluationForm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8632,7 +8371,6 @@ func (m *awsRestjson1_serializeOpDescribeEvaluationFor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeEvaluationFormInput(v *DescribeEvaluationFormInput, encoder *httpbinding.Encoder) error {
@@ -8675,10 +8413,8 @@ func (m *awsRestjson1_serializeOpDescribeHoursOfOperat
 func (m *awsRestjson1_serializeOpDescribeHoursOfOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8716,7 +8452,6 @@ func (m *awsRestjson1_serializeOpDescribeHoursOfOperat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeHoursOfOperationInput(v *DescribeHoursOfOperationInput, encoder *httpbinding.Encoder) error {
@@ -8755,10 +8490,8 @@ func (m *awsRestjson1_serializeOpDescribeHoursOfOperat
 func (m *awsRestjson1_serializeOpDescribeHoursOfOperationOverride) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8796,7 +8529,6 @@ func (m *awsRestjson1_serializeOpDescribeHoursOfOperat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeHoursOfOperationOverrideInput(v *DescribeHoursOfOperationOverrideInput, encoder *httpbinding.Encoder) error {
@@ -8844,10 +8576,8 @@ func (m *awsRestjson1_serializeOpDescribeInstance) Han
 func (m *awsRestjson1_serializeOpDescribeInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8885,7 +8615,6 @@ func (m *awsRestjson1_serializeOpDescribeInstance) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInstanceInput(v *DescribeInstanceInput, encoder *httpbinding.Encoder) error {
@@ -8915,10 +8644,8 @@ func (m *awsRestjson1_serializeOpDescribeInstanceAttri
 func (m *awsRestjson1_serializeOpDescribeInstanceAttribute) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8956,7 +8683,6 @@ func (m *awsRestjson1_serializeOpDescribeInstanceAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInstanceAttributeInput(v *DescribeInstanceAttributeInput, encoder *httpbinding.Encoder) error {
@@ -8995,10 +8721,8 @@ func (m *awsRestjson1_serializeOpDescribeInstanceStora
 func (m *awsRestjson1_serializeOpDescribeInstanceStorageConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9036,7 +8760,6 @@ func (m *awsRestjson1_serializeOpDescribeInstanceStora
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInstanceStorageConfigInput(v *DescribeInstanceStorageConfigInput, encoder *httpbinding.Encoder) error {
@@ -9079,10 +8802,8 @@ func (m *awsRestjson1_serializeOpDescribePhoneNumber) 
 func (m *awsRestjson1_serializeOpDescribePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9120,7 +8841,6 @@ func (m *awsRestjson1_serializeOpDescribePhoneNumber) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePhoneNumberInput(v *DescribePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -9150,10 +8870,8 @@ func (m *awsRestjson1_serializeOpDescribePredefinedAtt
 func (m *awsRestjson1_serializeOpDescribePredefinedAttribute) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9191,7 +8909,6 @@ func (m *awsRestjson1_serializeOpDescribePredefinedAtt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePredefinedAttributeInput(v *DescribePredefinedAttributeInput, encoder *httpbinding.Encoder) error {
@@ -9230,10 +8947,8 @@ func (m *awsRestjson1_serializeOpDescribePrompt) Handl
 func (m *awsRestjson1_serializeOpDescribePrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9271,7 +8986,6 @@ func (m *awsRestjson1_serializeOpDescribePrompt) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePromptInput(v *DescribePromptInput, encoder *httpbinding.Encoder) error {
@@ -9310,10 +9024,8 @@ func (m *awsRestjson1_serializeOpDescribeQueue) Handle
 func (m *awsRestjson1_serializeOpDescribeQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9351,7 +9063,6 @@ func (m *awsRestjson1_serializeOpDescribeQueue) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeQueueInput(v *DescribeQueueInput, encoder *httpbinding.Encoder) error {
@@ -9390,10 +9101,8 @@ func (m *awsRestjson1_serializeOpDescribeQuickConnect)
 func (m *awsRestjson1_serializeOpDescribeQuickConnect) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9431,7 +9140,6 @@ func (m *awsRestjson1_serializeOpDescribeQuickConnect)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeQuickConnectInput(v *DescribeQuickConnectInput, encoder *httpbinding.Encoder) error {
@@ -9470,10 +9178,8 @@ func (m *awsRestjson1_serializeOpDescribeRoutingProfil
 func (m *awsRestjson1_serializeOpDescribeRoutingProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9511,7 +9217,6 @@ func (m *awsRestjson1_serializeOpDescribeRoutingProfil
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRoutingProfileInput(v *DescribeRoutingProfileInput, encoder *httpbinding.Encoder) error {
@@ -9550,10 +9255,8 @@ func (m *awsRestjson1_serializeOpDescribeRule) HandleS
 func (m *awsRestjson1_serializeOpDescribeRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9591,7 +9294,6 @@ func (m *awsRestjson1_serializeOpDescribeRule) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRuleInput(v *DescribeRuleInput, encoder *httpbinding.Encoder) error {
@@ -9630,10 +9332,8 @@ func (m *awsRestjson1_serializeOpDescribeSecurityProfi
 func (m *awsRestjson1_serializeOpDescribeSecurityProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9671,7 +9371,6 @@ func (m *awsRestjson1_serializeOpDescribeSecurityProfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSecurityProfileInput(v *DescribeSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -9710,10 +9409,8 @@ func (m *awsRestjson1_serializeOpDescribeTrafficDistri
 func (m *awsRestjson1_serializeOpDescribeTrafficDistributionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9751,7 +9448,6 @@ func (m *awsRestjson1_serializeOpDescribeTrafficDistri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTrafficDistributionGroupInput(v *DescribeTrafficDistributionGroupInput, encoder *httpbinding.Encoder) error {
@@ -9781,10 +9477,8 @@ func (m *awsRestjson1_serializeOpDescribeUser) HandleS
 func (m *awsRestjson1_serializeOpDescribeUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9822,7 +9516,6 @@ func (m *awsRestjson1_serializeOpDescribeUser) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeUserInput(v *DescribeUserInput, encoder *httpbinding.Encoder) error {
@@ -9861,10 +9554,8 @@ func (m *awsRestjson1_serializeOpDescribeUserHierarchy
 func (m *awsRestjson1_serializeOpDescribeUserHierarchyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9902,7 +9593,6 @@ func (m *awsRestjson1_serializeOpDescribeUserHierarchy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeUserHierarchyGroupInput(v *DescribeUserHierarchyGroupInput, encoder *httpbinding.Encoder) error {
@@ -9941,10 +9631,8 @@ func (m *awsRestjson1_serializeOpDescribeUserHierarchy
 func (m *awsRestjson1_serializeOpDescribeUserHierarchyStructure) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9982,7 +9670,6 @@ func (m *awsRestjson1_serializeOpDescribeUserHierarchy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeUserHierarchyStructureInput(v *DescribeUserHierarchyStructureInput, encoder *httpbinding.Encoder) error {
@@ -10012,10 +9699,8 @@ func (m *awsRestjson1_serializeOpDescribeView) HandleS
 func (m *awsRestjson1_serializeOpDescribeView) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10053,7 +9738,6 @@ func (m *awsRestjson1_serializeOpDescribeView) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeViewInput(v *DescribeViewInput, encoder *httpbinding.Encoder) error {
@@ -10092,10 +9776,8 @@ func (m *awsRestjson1_serializeOpDescribeVocabulary) H
 func (m *awsRestjson1_serializeOpDescribeVocabulary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10133,7 +9815,6 @@ func (m *awsRestjson1_serializeOpDescribeVocabulary) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeVocabularyInput(v *DescribeVocabularyInput, encoder *httpbinding.Encoder) error {
@@ -10172,10 +9853,8 @@ func (m *awsRestjson1_serializeOpDisassociateAnalytics
 func (m *awsRestjson1_serializeOpDisassociateAnalyticsDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10224,7 +9903,6 @@ func (m *awsRestjson1_serializeOpDisassociateAnalytics
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateAnalyticsDataSetInput(v *DisassociateAnalyticsDataSetInput, encoder *httpbinding.Encoder) error {
@@ -10271,10 +9949,8 @@ func (m *awsRestjson1_serializeOpDisassociateApprovedO
 func (m *awsRestjson1_serializeOpDisassociateApprovedOrigin) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10312,7 +9988,6 @@ func (m *awsRestjson1_serializeOpDisassociateApprovedO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateApprovedOriginInput(v *DisassociateApprovedOriginInput, encoder *httpbinding.Encoder) error {
@@ -10350,10 +10025,8 @@ func (m *awsRestjson1_serializeOpDisassociateBot) Hand
 func (m *awsRestjson1_serializeOpDisassociateBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10402,7 +10075,6 @@ func (m *awsRestjson1_serializeOpDisassociateBot) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateBotInput(v *DisassociateBotInput, encoder *httpbinding.Encoder) error {
@@ -10458,10 +10130,8 @@ func (m *awsRestjson1_serializeOpDisassociateFlow) Han
 func (m *awsRestjson1_serializeOpDisassociateFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10499,7 +10169,6 @@ func (m *awsRestjson1_serializeOpDisassociateFlow) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFlowInput(v *DisassociateFlowInput, encoder *httpbinding.Encoder) error {
@@ -10547,10 +10216,8 @@ func (m *awsRestjson1_serializeOpDisassociateInstanceS
 func (m *awsRestjson1_serializeOpDisassociateInstanceStorageConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10588,7 +10255,6 @@ func (m *awsRestjson1_serializeOpDisassociateInstanceS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateInstanceStorageConfigInput(v *DisassociateInstanceStorageConfigInput, encoder *httpbinding.Encoder) error {
@@ -10635,10 +10301,8 @@ func (m *awsRestjson1_serializeOpDisassociateLambdaFun
 func (m *awsRestjson1_serializeOpDisassociateLambdaFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10676,7 +10340,6 @@ func (m *awsRestjson1_serializeOpDisassociateLambdaFun
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateLambdaFunctionInput(v *DisassociateLambdaFunctionInput, encoder *httpbinding.Encoder) error {
@@ -10714,10 +10377,8 @@ func (m *awsRestjson1_serializeOpDisassociateLexBot) H
 func (m *awsRestjson1_serializeOpDisassociateLexBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10755,7 +10416,6 @@ func (m *awsRestjson1_serializeOpDisassociateLexBot) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateLexBotInput(v *DisassociateLexBotInput, encoder *httpbinding.Encoder) error {
@@ -10797,10 +10457,8 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 func (m *awsRestjson1_serializeOpDisassociatePhoneNumberContactFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10838,7 +10496,6 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociatePhoneNumberContactFlowInput(v *DisassociatePhoneNumberContactFlowInput, encoder *httpbinding.Encoder) error {
@@ -10872,10 +10529,8 @@ func (m *awsRestjson1_serializeOpDisassociateQueueQuic
 func (m *awsRestjson1_serializeOpDisassociateQueueQuickConnects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10924,7 +10579,6 @@ func (m *awsRestjson1_serializeOpDisassociateQueueQuic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateQueueQuickConnectsInput(v *DisassociateQueueQuickConnectsInput, encoder *httpbinding.Encoder) error {
@@ -10977,10 +10631,8 @@ func (m *awsRestjson1_serializeOpDisassociateRoutingPr
 func (m *awsRestjson1_serializeOpDisassociateRoutingProfileQueues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11029,7 +10681,6 @@ func (m *awsRestjson1_serializeOpDisassociateRoutingPr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateRoutingProfileQueuesInput(v *DisassociateRoutingProfileQueuesInput, encoder *httpbinding.Encoder) error {
@@ -11082,10 +10733,8 @@ func (m *awsRestjson1_serializeOpDisassociateSecurityK
 func (m *awsRestjson1_serializeOpDisassociateSecurityKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11123,7 +10772,6 @@ func (m *awsRestjson1_serializeOpDisassociateSecurityK
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateSecurityKeyInput(v *DisassociateSecurityKeyInput, encoder *httpbinding.Encoder) error {
@@ -11166,10 +10814,8 @@ func (m *awsRestjson1_serializeOpDisassociateTrafficDi
 func (m *awsRestjson1_serializeOpDisassociateTrafficDistributionGroupUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11207,7 +10853,6 @@ func (m *awsRestjson1_serializeOpDisassociateTrafficDi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateTrafficDistributionGroupUserInput(v *DisassociateTrafficDistributionGroupUserInput, encoder *httpbinding.Encoder) error {
@@ -11245,10 +10890,8 @@ func (m *awsRestjson1_serializeOpDisassociateUserProfi
 func (m *awsRestjson1_serializeOpDisassociateUserProficiencies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11297,7 +10940,6 @@ func (m *awsRestjson1_serializeOpDisassociateUserProfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateUserProficienciesInput(v *DisassociateUserProficienciesInput, encoder *httpbinding.Encoder) error {
@@ -11350,10 +10992,8 @@ func (m *awsRestjson1_serializeOpDismissUserContact) H
 func (m *awsRestjson1_serializeOpDismissUserContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11402,7 +11042,6 @@ func (m *awsRestjson1_serializeOpDismissUserContact) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDismissUserContactInput(v *DismissUserContactInput, encoder *httpbinding.Encoder) error {
@@ -11453,10 +11092,8 @@ func (m *awsRestjson1_serializeOpGetAttachedFile) Hand
 func (m *awsRestjson1_serializeOpGetAttachedFile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11494,7 +11131,6 @@ func (m *awsRestjson1_serializeOpGetAttachedFile) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAttachedFileInput(v *GetAttachedFileInput, encoder *httpbinding.Encoder) error {
@@ -11541,10 +11177,8 @@ func (m *awsRestjson1_serializeOpGetContactAttributes)
 func (m *awsRestjson1_serializeOpGetContactAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11582,7 +11216,6 @@ func (m *awsRestjson1_serializeOpGetContactAttributes)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContactAttributesInput(v *GetContactAttributesInput, encoder *httpbinding.Encoder) error {
@@ -11621,10 +11254,8 @@ func (m *awsRestjson1_serializeOpGetCurrentMetricData)
 func (m *awsRestjson1_serializeOpGetCurrentMetricData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11673,7 +11304,6 @@ func (m *awsRestjson1_serializeOpGetCurrentMetricData)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCurrentMetricDataInput(v *GetCurrentMetricDataInput, encoder *httpbinding.Encoder) error {
@@ -11748,10 +11378,8 @@ func (m *awsRestjson1_serializeOpGetCurrentUserData) H
 func (m *awsRestjson1_serializeOpGetCurrentUserData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11800,7 +11428,6 @@ func (m *awsRestjson1_serializeOpGetCurrentUserData) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCurrentUserDataInput(v *GetCurrentUserDataInput, encoder *httpbinding.Encoder) error {
@@ -11854,10 +11481,8 @@ func (m *awsRestjson1_serializeOpGetEffectiveHoursOfOp
 func (m *awsRestjson1_serializeOpGetEffectiveHoursOfOperations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11895,7 +11520,6 @@ func (m *awsRestjson1_serializeOpGetEffectiveHoursOfOp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEffectiveHoursOfOperationsInput(v *GetEffectiveHoursOfOperationsInput, encoder *httpbinding.Encoder) error {
@@ -11942,10 +11566,8 @@ func (m *awsRestjson1_serializeOpGetFederationToken) H
 func (m *awsRestjson1_serializeOpGetFederationToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11983,7 +11605,6 @@ func (m *awsRestjson1_serializeOpGetFederationToken) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFederationTokenInput(v *GetFederationTokenInput, encoder *httpbinding.Encoder) error {
@@ -12013,10 +11634,8 @@ func (m *awsRestjson1_serializeOpGetFlowAssociation) H
 func (m *awsRestjson1_serializeOpGetFlowAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12054,7 +11673,6 @@ func (m *awsRestjson1_serializeOpGetFlowAssociation) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFlowAssociationInput(v *GetFlowAssociationInput, encoder *httpbinding.Encoder) error {
@@ -12102,10 +11720,8 @@ func (m *awsRestjson1_serializeOpGetMetricData) Handle
 func (m *awsRestjson1_serializeOpGetMetricData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12154,7 +11770,6 @@ func (m *awsRestjson1_serializeOpGetMetricData) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMetricDataInput(v *GetMetricDataInput, encoder *httpbinding.Encoder) error {
@@ -12232,10 +11847,8 @@ func (m *awsRestjson1_serializeOpGetMetricDataV2) Hand
 func (m *awsRestjson1_serializeOpGetMetricDataV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12280,7 +11893,6 @@ func (m *awsRestjson1_serializeOpGetMetricDataV2) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMetricDataV2Input(v *GetMetricDataV2Input, encoder *httpbinding.Encoder) error {
@@ -12361,10 +11973,8 @@ func (m *awsRestjson1_serializeOpGetPromptFile) Handle
 func (m *awsRestjson1_serializeOpGetPromptFile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12402,7 +12012,6 @@ func (m *awsRestjson1_serializeOpGetPromptFile) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPromptFileInput(v *GetPromptFileInput, encoder *httpbinding.Encoder) error {
@@ -12441,10 +12050,8 @@ func (m *awsRestjson1_serializeOpGetTaskTemplate) Hand
 func (m *awsRestjson1_serializeOpGetTaskTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12482,7 +12089,6 @@ func (m *awsRestjson1_serializeOpGetTaskTemplate) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTaskTemplateInput(v *GetTaskTemplateInput, encoder *httpbinding.Encoder) error {
@@ -12525,10 +12131,8 @@ func (m *awsRestjson1_serializeOpGetTrafficDistributio
 func (m *awsRestjson1_serializeOpGetTrafficDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12566,7 +12170,6 @@ func (m *awsRestjson1_serializeOpGetTrafficDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTrafficDistributionInput(v *GetTrafficDistributionInput, encoder *httpbinding.Encoder) error {
@@ -12596,10 +12199,8 @@ func (m *awsRestjson1_serializeOpImportPhoneNumber) Ha
 func (m *awsRestjson1_serializeOpImportPhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12644,7 +12245,6 @@ func (m *awsRestjson1_serializeOpImportPhoneNumber) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportPhoneNumberInput(v *ImportPhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -12699,10 +12299,8 @@ func (m *awsRestjson1_serializeOpListAgentStatuses) Ha
 func (m *awsRestjson1_serializeOpListAgentStatuses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12740,7 +12338,6 @@ func (m *awsRestjson1_serializeOpListAgentStatuses) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentStatusesInput(v *ListAgentStatusesInput, encoder *httpbinding.Encoder) error {
@@ -12784,10 +12381,8 @@ func (m *awsRestjson1_serializeOpListAnalyticsDataAsso
 func (m *awsRestjson1_serializeOpListAnalyticsDataAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12825,7 +12420,6 @@ func (m *awsRestjson1_serializeOpListAnalyticsDataAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnalyticsDataAssociationsInput(v *ListAnalyticsDataAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -12867,10 +12461,8 @@ func (m *awsRestjson1_serializeOpListAnalyticsDataLake
 func (m *awsRestjson1_serializeOpListAnalyticsDataLakeDataSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12908,7 +12500,6 @@ func (m *awsRestjson1_serializeOpListAnalyticsDataLake
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnalyticsDataLakeDataSetsInput(v *ListAnalyticsDataLakeDataSetsInput, encoder *httpbinding.Encoder) error {
@@ -12946,10 +12537,8 @@ func (m *awsRestjson1_serializeOpListApprovedOrigins) 
 func (m *awsRestjson1_serializeOpListApprovedOrigins) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12987,7 +12576,6 @@ func (m *awsRestjson1_serializeOpListApprovedOrigins) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListApprovedOriginsInput(v *ListApprovedOriginsInput, encoder *httpbinding.Encoder) error {
@@ -13025,10 +12613,8 @@ func (m *awsRestjson1_serializeOpListAssociatedContact
 func (m *awsRestjson1_serializeOpListAssociatedContacts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13066,7 +12652,6 @@ func (m *awsRestjson1_serializeOpListAssociatedContact
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssociatedContactsInput(v *ListAssociatedContactsInput, encoder *httpbinding.Encoder) error {
@@ -13108,10 +12693,8 @@ func (m *awsRestjson1_serializeOpListAuthenticationPro
 func (m *awsRestjson1_serializeOpListAuthenticationProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13149,7 +12732,6 @@ func (m *awsRestjson1_serializeOpListAuthenticationPro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAuthenticationProfilesInput(v *ListAuthenticationProfilesInput, encoder *httpbinding.Encoder) error {
@@ -13187,10 +12769,8 @@ func (m *awsRestjson1_serializeOpListBots) HandleSeria
 func (m *awsRestjson1_serializeOpListBots) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13228,7 +12808,6 @@ func (m *awsRestjson1_serializeOpListBots) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotsInput(v *ListBotsInput, encoder *httpbinding.Encoder) error {
@@ -13270,10 +12849,8 @@ func (m *awsRestjson1_serializeOpListContactEvaluation
 func (m *awsRestjson1_serializeOpListContactEvaluations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13311,7 +12888,6 @@ func (m *awsRestjson1_serializeOpListContactEvaluation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContactEvaluationsInput(v *ListContactEvaluationsInput, encoder *httpbinding.Encoder) error {
@@ -13349,10 +12925,8 @@ func (m *awsRestjson1_serializeOpListContactFlowModule
 func (m *awsRestjson1_serializeOpListContactFlowModules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13390,7 +12964,6 @@ func (m *awsRestjson1_serializeOpListContactFlowModule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContactFlowModulesInput(v *ListContactFlowModulesInput, encoder *httpbinding.Encoder) error {
@@ -13432,10 +13005,8 @@ func (m *awsRestjson1_serializeOpListContactFlows) Han
 func (m *awsRestjson1_serializeOpListContactFlows) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13473,7 +13044,6 @@ func (m *awsRestjson1_serializeOpListContactFlows) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContactFlowsInput(v *ListContactFlowsInput, encoder *httpbinding.Encoder) error {
@@ -13517,10 +13087,8 @@ func (m *awsRestjson1_serializeOpListContactFlowVersio
 func (m *awsRestjson1_serializeOpListContactFlowVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13558,7 +13126,6 @@ func (m *awsRestjson1_serializeOpListContactFlowVersio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContactFlowVersionsInput(v *ListContactFlowVersionsInput, encoder *httpbinding.Encoder) error {
@@ -13605,10 +13172,8 @@ func (m *awsRestjson1_serializeOpListContactReferences
 func (m *awsRestjson1_serializeOpListContactReferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13646,7 +13211,6 @@ func (m *awsRestjson1_serializeOpListContactReferences
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContactReferencesInput(v *ListContactReferencesInput, encoder *httpbinding.Encoder) error {
@@ -13695,10 +13259,8 @@ func (m *awsRestjson1_serializeOpListDefaultVocabulari
 func (m *awsRestjson1_serializeOpListDefaultVocabularies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13747,7 +13309,6 @@ func (m *awsRestjson1_serializeOpListDefaultVocabulari
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDefaultVocabulariesInput(v *ListDefaultVocabulariesInput, encoder *httpbinding.Encoder) error {
@@ -13799,10 +13360,8 @@ func (m *awsRestjson1_serializeOpListEvaluationForms) 
 func (m *awsRestjson1_serializeOpListEvaluationForms) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13840,7 +13399,6 @@ func (m *awsRestjson1_serializeOpListEvaluationForms) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEvaluationFormsInput(v *ListEvaluationFormsInput, encoder *httpbinding.Encoder) error {
@@ -13878,10 +13436,8 @@ func (m *awsRestjson1_serializeOpListEvaluationFormVer
 func (m *awsRestjson1_serializeOpListEvaluationFormVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13919,7 +13475,6 @@ func (m *awsRestjson1_serializeOpListEvaluationFormVer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEvaluationFormVersionsInput(v *ListEvaluationFormVersionsInput, encoder *httpbinding.Encoder) error {
@@ -13966,10 +13521,8 @@ func (m *awsRestjson1_serializeOpListFlowAssociations)
 func (m *awsRestjson1_serializeOpListFlowAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14007,7 +13560,6 @@ func (m *awsRestjson1_serializeOpListFlowAssociations)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFlowAssociationsInput(v *ListFlowAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -14049,10 +13601,8 @@ func (m *awsRestjson1_serializeOpListHoursOfOperationO
 func (m *awsRestjson1_serializeOpListHoursOfOperationOverrides) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14090,7 +13640,6 @@ func (m *awsRestjson1_serializeOpListHoursOfOperationO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListHoursOfOperationOverridesInput(v *ListHoursOfOperationOverridesInput, encoder *httpbinding.Encoder) error {
@@ -14137,10 +13686,8 @@ func (m *awsRestjson1_serializeOpListHoursOfOperations
 func (m *awsRestjson1_serializeOpListHoursOfOperations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14178,7 +13725,6 @@ func (m *awsRestjson1_serializeOpListHoursOfOperations
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListHoursOfOperationsInput(v *ListHoursOfOperationsInput, encoder *httpbinding.Encoder) error {
@@ -14216,10 +13762,8 @@ func (m *awsRestjson1_serializeOpListInstanceAttribute
 func (m *awsRestjson1_serializeOpListInstanceAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14257,7 +13801,6 @@ func (m *awsRestjson1_serializeOpListInstanceAttribute
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInstanceAttributesInput(v *ListInstanceAttributesInput, encoder *httpbinding.Encoder) error {
@@ -14295,10 +13838,8 @@ func (m *awsRestjson1_serializeOpListInstances) Handle
 func (m *awsRestjson1_serializeOpListInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14336,7 +13877,6 @@ func (m *awsRestjson1_serializeOpListInstances) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInstancesInput(v *ListInstancesInput, encoder *httpbinding.Encoder) error {
@@ -14365,10 +13905,8 @@ func (m *awsRestjson1_serializeOpListInstanceStorageCo
 func (m *awsRestjson1_serializeOpListInstanceStorageConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14406,7 +13944,6 @@ func (m *awsRestjson1_serializeOpListInstanceStorageCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInstanceStorageConfigsInput(v *ListInstanceStorageConfigsInput, encoder *httpbinding.Encoder) error {
@@ -14448,10 +13985,8 @@ func (m *awsRestjson1_serializeOpListIntegrationAssoci
 func (m *awsRestjson1_serializeOpListIntegrationAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14489,7 +14024,6 @@ func (m *awsRestjson1_serializeOpListIntegrationAssoci
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIntegrationAssociationsInput(v *ListIntegrationAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -14535,10 +14069,8 @@ func (m *awsRestjson1_serializeOpListLambdaFunctions) 
 func (m *awsRestjson1_serializeOpListLambdaFunctions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14576,7 +14108,6 @@ func (m *awsRestjson1_serializeOpListLambdaFunctions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLambdaFunctionsInput(v *ListLambdaFunctionsInput, encoder *httpbinding.Encoder) error {
@@ -14614,10 +14145,8 @@ func (m *awsRestjson1_serializeOpListLexBots) HandleSe
 func (m *awsRestjson1_serializeOpListLexBots) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14655,7 +14184,6 @@ func (m *awsRestjson1_serializeOpListLexBots) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLexBotsInput(v *ListLexBotsInput, encoder *httpbinding.Encoder) error {
@@ -14693,10 +14221,8 @@ func (m *awsRestjson1_serializeOpListPhoneNumbers) Han
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
@@ -14734,7 +14260,6 @@ func (m *awsRestjson1_serializeOpListPhoneNumbers) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPhoneNumbersInput(v *ListPhoneNumbersInput, encoder *httpbinding.Encoder) error {
@@ -14784,10 +14309,8 @@ func (m *awsRestjson1_serializeOpListPhoneNumbersV2) H
 func (m *awsRestjson1_serializeOpListPhoneNumbersV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14832,7 +14355,6 @@ func (m *awsRestjson1_serializeOpListPhoneNumbersV2) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPhoneNumbersV2Input(v *ListPhoneNumbersV2Input, encoder *httpbinding.Encoder) error {
@@ -14899,10 +14421,8 @@ func (m *awsRestjson1_serializeOpListPredefinedAttribu
 func (m *awsRestjson1_serializeOpListPredefinedAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14940,7 +14460,6 @@ func (m *awsRestjson1_serializeOpListPredefinedAttribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPredefinedAttributesInput(v *ListPredefinedAttributesInput, encoder *httpbinding.Encoder) error {
@@ -14978,10 +14497,8 @@ func (m *awsRestjson1_serializeOpListPrompts) HandleSe
 func (m *awsRestjson1_serializeOpListPrompts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15019,7 +14536,6 @@ func (m *awsRestjson1_serializeOpListPrompts) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPromptsInput(v *ListPromptsInput, encoder *httpbinding.Encoder) error {
@@ -15057,10 +14573,8 @@ func (m *awsRestjson1_serializeOpListQueueQuickConnect
 func (m *awsRestjson1_serializeOpListQueueQuickConnects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15098,7 +14612,6 @@ func (m *awsRestjson1_serializeOpListQueueQuickConnect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueueQuickConnectsInput(v *ListQueueQuickConnectsInput, encoder *httpbinding.Encoder) error {
@@ -15145,10 +14658,8 @@ func (m *awsRestjson1_serializeOpListQueues) HandleSer
 func (m *awsRestjson1_serializeOpListQueues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15186,7 +14697,6 @@ func (m *awsRestjson1_serializeOpListQueues) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueuesInput(v *ListQueuesInput, encoder *httpbinding.Encoder) error {
@@ -15230,10 +14740,8 @@ func (m *awsRestjson1_serializeOpListQuickConnects) Ha
 func (m *awsRestjson1_serializeOpListQuickConnects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15271,7 +14779,6 @@ func (m *awsRestjson1_serializeOpListQuickConnects) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQuickConnectsInput(v *ListQuickConnectsInput, encoder *httpbinding.Encoder) error {
@@ -15315,10 +14822,8 @@ func (m *awsRestjson1_serializeOpListRealtimeContactAn
 func (m *awsRestjson1_serializeOpListRealtimeContactAnalysisSegmentsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15367,7 +14872,6 @@ func (m *awsRestjson1_serializeOpListRealtimeContactAn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRealtimeContactAnalysisSegmentsV2Input(v *ListRealtimeContactAnalysisSegmentsV2Input, encoder *httpbinding.Encoder) error {
@@ -15435,10 +14939,8 @@ func (m *awsRestjson1_serializeOpListRoutingProfileQue
 func (m *awsRestjson1_serializeOpListRoutingProfileQueues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15476,7 +14978,6 @@ func (m *awsRestjson1_serializeOpListRoutingProfileQue
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRoutingProfileQueuesInput(v *ListRoutingProfileQueuesInput, encoder *httpbinding.Encoder) error {
@@ -15523,10 +15024,8 @@ func (m *awsRestjson1_serializeOpListRoutingProfiles) 
 func (m *awsRestjson1_serializeOpListRoutingProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15564,7 +15063,6 @@ func (m *awsRestjson1_serializeOpListRoutingProfiles) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRoutingProfilesInput(v *ListRoutingProfilesInput, encoder *httpbinding.Encoder) error {
@@ -15602,10 +15100,8 @@ func (m *awsRestjson1_serializeOpListRules) HandleSeri
 func (m *awsRestjson1_serializeOpListRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15643,7 +15139,6 @@ func (m *awsRestjson1_serializeOpListRules) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRulesInput(v *ListRulesInput, encoder *httpbinding.Encoder) error {
@@ -15689,10 +15184,8 @@ func (m *awsRestjson1_serializeOpListSecurityKeys) Han
 func (m *awsRestjson1_serializeOpListSecurityKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15730,7 +15223,6 @@ func (m *awsRestjson1_serializeOpListSecurityKeys) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSecurityKeysInput(v *ListSecurityKeysInput, encoder *httpbinding.Encoder) error {
@@ -15768,10 +15260,8 @@ func (m *awsRestjson1_serializeOpListSecurityProfileAp
 func (m *awsRestjson1_serializeOpListSecurityProfileApplications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15809,7 +15299,6 @@ func (m *awsRestjson1_serializeOpListSecurityProfileAp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSecurityProfileApplicationsInput(v *ListSecurityProfileApplicationsInput, encoder *httpbinding.Encoder) error {
@@ -15856,10 +15345,8 @@ func (m *awsRestjson1_serializeOpListSecurityProfilePe
 func (m *awsRestjson1_serializeOpListSecurityProfilePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15897,7 +15384,6 @@ func (m *awsRestjson1_serializeOpListSecurityProfilePe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSecurityProfilePermissionsInput(v *ListSecurityProfilePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -15944,10 +15430,8 @@ func (m *awsRestjson1_serializeOpListSecurityProfiles)
 func (m *awsRestjson1_serializeOpListSecurityProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15985,7 +15469,6 @@ func (m *awsRestjson1_serializeOpListSecurityProfiles)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSecurityProfilesInput(v *ListSecurityProfilesInput, encoder *httpbinding.Encoder) error {
@@ -16023,10 +15506,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -16064,7 +15545,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -16094,10 +15574,8 @@ func (m *awsRestjson1_serializeOpListTaskTemplates) Ha
 func (m *awsRestjson1_serializeOpListTaskTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16135,7 +15613,6 @@ func (m *awsRestjson1_serializeOpListTaskTemplates) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTaskTemplatesInput(v *ListTaskTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -16181,10 +15658,8 @@ func (m *awsRestjson1_serializeOpListTrafficDistributi
 func (m *awsRestjson1_serializeOpListTrafficDistributionGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16222,7 +15697,6 @@ func (m *awsRestjson1_serializeOpListTrafficDistributi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrafficDistributionGroupsInput(v *ListTrafficDistributionGroupsInput, encoder *httpbinding.Encoder) error {
@@ -16255,10 +15729,8 @@ func (m *awsRestjson1_serializeOpListTrafficDistributi
 func (m *awsRestjson1_serializeOpListTrafficDistributionGroupUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16296,7 +15768,6 @@ func (m *awsRestjson1_serializeOpListTrafficDistributi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrafficDistributionGroupUsersInput(v *ListTrafficDistributionGroupUsersInput, encoder *httpbinding.Encoder) error {
@@ -16334,10 +15805,8 @@ func (m *awsRestjson1_serializeOpListUseCases) HandleS
 func (m *awsRestjson1_serializeOpListUseCases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16375,7 +15844,6 @@ func (m *awsRestjson1_serializeOpListUseCases) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUseCasesInput(v *ListUseCasesInput, encoder *httpbinding.Encoder) error {
@@ -16422,10 +15890,8 @@ func (m *awsRestjson1_serializeOpListUserHierarchyGrou
 func (m *awsRestjson1_serializeOpListUserHierarchyGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16463,7 +15929,6 @@ func (m *awsRestjson1_serializeOpListUserHierarchyGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUserHierarchyGroupsInput(v *ListUserHierarchyGroupsInput, encoder *httpbinding.Encoder) error {
@@ -16501,10 +15966,8 @@ func (m *awsRestjson1_serializeOpListUserProficiencies
 func (m *awsRestjson1_serializeOpListUserProficiencies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16542,7 +16005,6 @@ func (m *awsRestjson1_serializeOpListUserProficiencies
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUserProficienciesInput(v *ListUserProficienciesInput, encoder *httpbinding.Encoder) error {
@@ -16589,10 +16051,8 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
 func (m *awsRestjson1_serializeOpListUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16630,7 +16090,6 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUsersInput(v *ListUsersInput, encoder *httpbinding.Encoder) error {
@@ -16668,10 +16127,8 @@ func (m *awsRestjson1_serializeOpListViews) HandleSeri
 func (m *awsRestjson1_serializeOpListViews) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16709,7 +16166,6 @@ func (m *awsRestjson1_serializeOpListViews) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListViewsInput(v *ListViewsInput, encoder *httpbinding.Encoder) error {
@@ -16751,10 +16207,8 @@ func (m *awsRestjson1_serializeOpListViewVersions) Han
 func (m *awsRestjson1_serializeOpListViewVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16792,7 +16246,6 @@ func (m *awsRestjson1_serializeOpListViewVersions) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListViewVersionsInput(v *ListViewVersionsInput, encoder *httpbinding.Encoder) error {
@@ -16839,10 +16292,8 @@ func (m *awsRestjson1_serializeOpMonitorContact) Handl
 func (m *awsRestjson1_serializeOpMonitorContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16887,7 +16338,6 @@ func (m *awsRestjson1_serializeOpMonitorContact) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMonitorContactInput(v *MonitorContactInput, encoder *httpbinding.Encoder) error {
@@ -16942,10 +16392,8 @@ func (m *awsRestjson1_serializeOpPauseContact) HandleS
 func (m *awsRestjson1_serializeOpPauseContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16990,7 +16438,6 @@ func (m *awsRestjson1_serializeOpPauseContact) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPauseContactInput(v *PauseContactInput, encoder *httpbinding.Encoder) error {
@@ -17033,10 +16480,8 @@ func (m *awsRestjson1_serializeOpPutUserStatus) Handle
 func (m *awsRestjson1_serializeOpPutUserStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17085,7 +16530,6 @@ func (m *awsRestjson1_serializeOpPutUserStatus) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutUserStatusInput(v *PutUserStatusInput, encoder *httpbinding.Encoder) error {
@@ -17136,10 +16580,8 @@ func (m *awsRestjson1_serializeOpReleasePhoneNumber) H
 func (m *awsRestjson1_serializeOpReleasePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17177,7 +16619,6 @@ func (m *awsRestjson1_serializeOpReleasePhoneNumber) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsReleasePhoneNumberInput(v *ReleasePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -17211,10 +16652,8 @@ func (m *awsRestjson1_serializeOpReplicateInstance) Ha
 func (m *awsRestjson1_serializeOpReplicateInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17263,7 +16702,6 @@ func (m *awsRestjson1_serializeOpReplicateInstance) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsReplicateInstanceInput(v *ReplicateInstanceInput, encoder *httpbinding.Encoder) error {
@@ -17315,10 +16753,8 @@ func (m *awsRestjson1_serializeOpResumeContact) Handle
 func (m *awsRestjson1_serializeOpResumeContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17363,7 +16799,6 @@ func (m *awsRestjson1_serializeOpResumeContact) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResumeContactInput(v *ResumeContactInput, encoder *httpbinding.Encoder) error {
@@ -17406,10 +16841,8 @@ func (m *awsRestjson1_serializeOpResumeContactRecordin
 func (m *awsRestjson1_serializeOpResumeContactRecording) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17454,7 +16887,6 @@ func (m *awsRestjson1_serializeOpResumeContactRecordin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResumeContactRecordingInput(v *ResumeContactRecordingInput, encoder *httpbinding.Encoder) error {
@@ -17502,10 +16934,8 @@ func (m *awsRestjson1_serializeOpSearchAgentStatuses) 
 func (m *awsRestjson1_serializeOpSearchAgentStatuses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17550,7 +16980,6 @@ func (m *awsRestjson1_serializeOpSearchAgentStatuses) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchAgentStatusesInput(v *SearchAgentStatusesInput, encoder *httpbinding.Encoder) error {
@@ -17607,10 +17036,8 @@ func (m *awsRestjson1_serializeOpSearchAvailablePhoneN
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
@@ -17655,7 +17082,6 @@ func (m *awsRestjson1_serializeOpSearchAvailablePhoneN
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchAvailablePhoneNumbersInput(v *SearchAvailablePhoneNumbersInput, encoder *httpbinding.Encoder) error {
@@ -17718,10 +17144,8 @@ func (m *awsRestjson1_serializeOpSearchContactFlowModu
 func (m *awsRestjson1_serializeOpSearchContactFlowModules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17766,7 +17190,6 @@ func (m *awsRestjson1_serializeOpSearchContactFlowModu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchContactFlowModulesInput(v *SearchContactFlowModulesInput, encoder *httpbinding.Encoder) error {
@@ -17823,10 +17246,8 @@ func (m *awsRestjson1_serializeOpSearchContactFlows) H
 func (m *awsRestjson1_serializeOpSearchContactFlows) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17871,7 +17292,6 @@ func (m *awsRestjson1_serializeOpSearchContactFlows) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchContactFlowsInput(v *SearchContactFlowsInput, encoder *httpbinding.Encoder) error {
@@ -17928,10 +17348,8 @@ func (m *awsRestjson1_serializeOpSearchContacts) Handl
 func (m *awsRestjson1_serializeOpSearchContacts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17976,7 +17394,6 @@ func (m *awsRestjson1_serializeOpSearchContacts) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchContactsInput(v *SearchContactsInput, encoder *httpbinding.Encoder) error {
@@ -18040,10 +17457,8 @@ func (m *awsRestjson1_serializeOpSearchEmailAddresses)
 func (m *awsRestjson1_serializeOpSearchEmailAddresses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18088,7 +17503,6 @@ func (m *awsRestjson1_serializeOpSearchEmailAddresses)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchEmailAddressesInput(v *SearchEmailAddressesInput, encoder *httpbinding.Encoder) error {
@@ -18145,10 +17559,8 @@ func (m *awsRestjson1_serializeOpSearchHoursOfOperatio
 func (m *awsRestjson1_serializeOpSearchHoursOfOperationOverrides) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18193,7 +17605,6 @@ func (m *awsRestjson1_serializeOpSearchHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchHoursOfOperationOverridesInput(v *SearchHoursOfOperationOverridesInput, encoder *httpbinding.Encoder) error {
@@ -18250,10 +17661,8 @@ func (m *awsRestjson1_serializeOpSearchHoursOfOperatio
 func (m *awsRestjson1_serializeOpSearchHoursOfOperations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18298,7 +17707,6 @@ func (m *awsRestjson1_serializeOpSearchHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchHoursOfOperationsInput(v *SearchHoursOfOperationsInput, encoder *httpbinding.Encoder) error {
@@ -18355,10 +17763,8 @@ func (m *awsRestjson1_serializeOpSearchPredefinedAttri
 func (m *awsRestjson1_serializeOpSearchPredefinedAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18403,7 +17809,6 @@ func (m *awsRestjson1_serializeOpSearchPredefinedAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchPredefinedAttributesInput(v *SearchPredefinedAttributesInput, encoder *httpbinding.Encoder) error {
@@ -18453,10 +17858,8 @@ func (m *awsRestjson1_serializeOpSearchPrompts) Handle
 func (m *awsRestjson1_serializeOpSearchPrompts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18501,7 +17904,6 @@ func (m *awsRestjson1_serializeOpSearchPrompts) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchPromptsInput(v *SearchPromptsInput, encoder *httpbinding.Encoder) error {
@@ -18558,10 +17960,8 @@ func (m *awsRestjson1_serializeOpSearchQueues) HandleS
 func (m *awsRestjson1_serializeOpSearchQueues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18606,7 +18006,6 @@ func (m *awsRestjson1_serializeOpSearchQueues) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchQueuesInput(v *SearchQueuesInput, encoder *httpbinding.Encoder) error {
@@ -18663,10 +18062,8 @@ func (m *awsRestjson1_serializeOpSearchQuickConnects) 
 func (m *awsRestjson1_serializeOpSearchQuickConnects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18711,7 +18108,6 @@ func (m *awsRestjson1_serializeOpSearchQuickConnects) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchQuickConnectsInput(v *SearchQuickConnectsInput, encoder *httpbinding.Encoder) error {
@@ -18768,10 +18164,8 @@ func (m *awsRestjson1_serializeOpSearchResourceTags) H
 func (m *awsRestjson1_serializeOpSearchResourceTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18816,7 +18210,6 @@ func (m *awsRestjson1_serializeOpSearchResourceTags) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchResourceTagsInput(v *SearchResourceTagsInput, encoder *httpbinding.Encoder) error {
@@ -18873,10 +18266,8 @@ func (m *awsRestjson1_serializeOpSearchRoutingProfiles
 func (m *awsRestjson1_serializeOpSearchRoutingProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18921,7 +18312,6 @@ func (m *awsRestjson1_serializeOpSearchRoutingProfiles
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchRoutingProfilesInput(v *SearchRoutingProfilesInput, encoder *httpbinding.Encoder) error {
@@ -18978,10 +18368,8 @@ func (m *awsRestjson1_serializeOpSearchSecurityProfile
 func (m *awsRestjson1_serializeOpSearchSecurityProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19026,7 +18414,6 @@ func (m *awsRestjson1_serializeOpSearchSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchSecurityProfilesInput(v *SearchSecurityProfilesInput, encoder *httpbinding.Encoder) error {
@@ -19083,10 +18470,8 @@ func (m *awsRestjson1_serializeOpSearchUserHierarchyGr
 func (m *awsRestjson1_serializeOpSearchUserHierarchyGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19131,7 +18516,6 @@ func (m *awsRestjson1_serializeOpSearchUserHierarchyGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchUserHierarchyGroupsInput(v *SearchUserHierarchyGroupsInput, encoder *httpbinding.Encoder) error {
@@ -19188,10 +18572,8 @@ func (m *awsRestjson1_serializeOpSearchUsers) HandleSe
 func (m *awsRestjson1_serializeOpSearchUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19236,7 +18618,6 @@ func (m *awsRestjson1_serializeOpSearchUsers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchUsersInput(v *SearchUsersInput, encoder *httpbinding.Encoder) error {
@@ -19293,10 +18674,8 @@ func (m *awsRestjson1_serializeOpSearchVocabularies) H
 func (m *awsRestjson1_serializeOpSearchVocabularies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19345,7 +18724,6 @@ func (m *awsRestjson1_serializeOpSearchVocabularies) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchVocabulariesInput(v *SearchVocabulariesInput, encoder *httpbinding.Encoder) error {
@@ -19407,10 +18785,8 @@ func (m *awsRestjson1_serializeOpSendChatIntegrationEv
 func (m *awsRestjson1_serializeOpSendChatIntegrationEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19455,7 +18831,6 @@ func (m *awsRestjson1_serializeOpSendChatIntegrationEv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendChatIntegrationEventInput(v *SendChatIntegrationEventInput, encoder *httpbinding.Encoder) error {
@@ -19512,10 +18887,8 @@ func (m *awsRestjson1_serializeOpSendOutboundEmail) Ha
 func (m *awsRestjson1_serializeOpSendOutboundEmail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19564,7 +18937,6 @@ func (m *awsRestjson1_serializeOpSendOutboundEmail) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendOutboundEmailInput(v *SendOutboundEmailInput, encoder *httpbinding.Encoder) error {
@@ -19646,10 +19018,8 @@ func (m *awsRestjson1_serializeOpStartAttachedFileUplo
 func (m *awsRestjson1_serializeOpStartAttachedFileUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19698,7 +19068,6 @@ func (m *awsRestjson1_serializeOpStartAttachedFileUplo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAttachedFileUploadInput(v *StartAttachedFileUploadInput, encoder *httpbinding.Encoder) error {
@@ -19778,10 +19147,8 @@ func (m *awsRestjson1_serializeOpStartChatContact) Han
 func (m *awsRestjson1_serializeOpStartChatContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19826,7 +19193,6 @@ func (m *awsRestjson1_serializeOpStartChatContact) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartChatContactInput(v *StartChatContactInput, encoder *httpbinding.Encoder) error {
@@ -19926,10 +19292,8 @@ func (m *awsRestjson1_serializeOpStartContactEvaluatio
 func (m *awsRestjson1_serializeOpStartContactEvaluation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19978,7 +19342,6 @@ func (m *awsRestjson1_serializeOpStartContactEvaluatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartContactEvaluationInput(v *StartContactEvaluationInput, encoder *httpbinding.Encoder) error {
@@ -20030,10 +19393,8 @@ func (m *awsRestjson1_serializeOpStartContactRecording
 func (m *awsRestjson1_serializeOpStartContactRecording) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20078,7 +19439,6 @@ func (m *awsRestjson1_serializeOpStartContactRecording
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartContactRecordingInput(v *StartContactRecordingInput, encoder *httpbinding.Encoder) error {
@@ -20128,10 +19488,8 @@ func (m *awsRestjson1_serializeOpStartContactStreaming
 func (m *awsRestjson1_serializeOpStartContactStreaming) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20176,7 +19534,6 @@ func (m *awsRestjson1_serializeOpStartContactStreaming
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartContactStreamingInput(v *StartContactStreamingInput, encoder *httpbinding.Encoder) error {
@@ -20226,10 +19583,8 @@ func (m *awsRestjson1_serializeOpStartEmailContact) Ha
 func (m *awsRestjson1_serializeOpStartEmailContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20274,7 +19629,6 @@ func (m *awsRestjson1_serializeOpStartEmailContact) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartEmailContactInput(v *StartEmailContactInput, encoder *httpbinding.Encoder) error {
@@ -20386,10 +19740,8 @@ func (m *awsRestjson1_serializeOpStartOutboundChatCont
 func (m *awsRestjson1_serializeOpStartOutboundChatContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20434,7 +19786,6 @@ func (m *awsRestjson1_serializeOpStartOutboundChatCont
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartOutboundChatContactInput(v *StartOutboundChatContactInput, encoder *httpbinding.Encoder) error {
@@ -20536,10 +19887,8 @@ func (m *awsRestjson1_serializeOpStartOutboundEmailCon
 func (m *awsRestjson1_serializeOpStartOutboundEmailContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20584,7 +19933,6 @@ func (m *awsRestjson1_serializeOpStartOutboundEmailCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartOutboundEmailContactInput(v *StartOutboundEmailContactInput, encoder *httpbinding.Encoder) error {
@@ -20655,10 +20003,8 @@ func (m *awsRestjson1_serializeOpStartOutboundVoiceCon
 func (m *awsRestjson1_serializeOpStartOutboundVoiceContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20703,7 +20049,6 @@ func (m *awsRestjson1_serializeOpStartOutboundVoiceCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartOutboundVoiceContactInput(v *StartOutboundVoiceContactInput, encoder *httpbinding.Encoder) error {
@@ -20807,10 +20152,8 @@ func (m *awsRestjson1_serializeOpStartScreenSharing) H
 func (m *awsRestjson1_serializeOpStartScreenSharing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20855,7 +20198,6 @@ func (m *awsRestjson1_serializeOpStartScreenSharing) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartScreenSharingInput(v *StartScreenSharingInput, encoder *httpbinding.Encoder) error {
@@ -20898,10 +20240,8 @@ func (m *awsRestjson1_serializeOpStartTaskContact) Han
 func (m *awsRestjson1_serializeOpStartTaskContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20946,7 +20286,6 @@ func (m *awsRestjson1_serializeOpStartTaskContact) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTaskContactInput(v *StartTaskContactInput, encoder *httpbinding.Encoder) error {
@@ -21045,10 +20384,8 @@ func (m *awsRestjson1_serializeOpStartWebRTCContact) H
 func (m *awsRestjson1_serializeOpStartWebRTCContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21093,7 +20430,6 @@ func (m *awsRestjson1_serializeOpStartWebRTCContact) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartWebRTCContactInput(v *StartWebRTCContactInput, encoder *httpbinding.Encoder) error {
@@ -21174,10 +20510,8 @@ func (m *awsRestjson1_serializeOpStopContact) HandleSe
 func (m *awsRestjson1_serializeOpStopContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21222,7 +20556,6 @@ func (m *awsRestjson1_serializeOpStopContact) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopContactInput(v *StopContactInput, encoder *httpbinding.Encoder) error {
@@ -21267,10 +20600,8 @@ func (m *awsRestjson1_serializeOpStopContactRecording)
 func (m *awsRestjson1_serializeOpStopContactRecording) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21315,7 +20646,6 @@ func (m *awsRestjson1_serializeOpStopContactRecording)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopContactRecordingInput(v *StopContactRecordingInput, encoder *httpbinding.Encoder) error {
@@ -21363,10 +20693,8 @@ func (m *awsRestjson1_serializeOpStopContactStreaming)
 func (m *awsRestjson1_serializeOpStopContactStreaming) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21411,7 +20739,6 @@ func (m *awsRestjson1_serializeOpStopContactStreaming)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopContactStreamingInput(v *StopContactStreamingInput, encoder *httpbinding.Encoder) error {
@@ -21454,10 +20781,8 @@ func (m *awsRestjson1_serializeOpSubmitContactEvaluati
 func (m *awsRestjson1_serializeOpSubmitContactEvaluation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21506,7 +20831,6 @@ func (m *awsRestjson1_serializeOpSubmitContactEvaluati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSubmitContactEvaluationInput(v *SubmitContactEvaluationInput, encoder *httpbinding.Encoder) error {
@@ -21566,10 +20890,8 @@ func (m *awsRestjson1_serializeOpSuspendContactRecordi
 func (m *awsRestjson1_serializeOpSuspendContactRecording) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21614,7 +20936,6 @@ func (m *awsRestjson1_serializeOpSuspendContactRecordi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSuspendContactRecordingInput(v *SuspendContactRecordingInput, encoder *httpbinding.Encoder) error {
@@ -21662,10 +20983,8 @@ func (m *awsRestjson1_serializeOpTagContact) HandleSer
 func (m *awsRestjson1_serializeOpTagContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21710,7 +21029,6 @@ func (m *awsRestjson1_serializeOpTagContact) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagContactInput(v *TagContactInput, encoder *httpbinding.Encoder) error {
@@ -21755,10 +21073,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -21807,7 +21123,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -21851,10 +21166,8 @@ func (m *awsRestjson1_serializeOpTransferContact) Hand
 func (m *awsRestjson1_serializeOpTransferContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21899,7 +21212,6 @@ func (m *awsRestjson1_serializeOpTransferContact) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTransferContactInput(v *TransferContactInput, encoder *httpbinding.Encoder) error {
@@ -21957,10 +21269,8 @@ func (m *awsRestjson1_serializeOpUntagContact) HandleS
 func (m *awsRestjson1_serializeOpUntagContact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21998,7 +21308,6 @@ func (m *awsRestjson1_serializeOpUntagContact) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagContactInput(v *UntagContactInput, encoder *httpbinding.Encoder) error {
@@ -22043,10 +21352,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -22084,7 +21391,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -22120,10 +21426,8 @@ func (m *awsRestjson1_serializeOpUpdateAgentStatus) Ha
 func (m *awsRestjson1_serializeOpUpdateAgentStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22172,7 +21476,6 @@ func (m *awsRestjson1_serializeOpUpdateAgentStatus) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentStatusInput(v *UpdateAgentStatusInput, encoder *httpbinding.Encoder) error {
@@ -22243,10 +21546,8 @@ func (m *awsRestjson1_serializeOpUpdateAuthenticationP
 func (m *awsRestjson1_serializeOpUpdateAuthenticationProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22295,7 +21596,6 @@ func (m *awsRestjson1_serializeOpUpdateAuthenticationP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAuthenticationProfileInput(v *UpdateAuthenticationProfileInput, encoder *httpbinding.Encoder) error {
@@ -22370,10 +21670,8 @@ func (m *awsRestjson1_serializeOpUpdateContact) Handle
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
@@ -22422,7 +21720,6 @@ func (m *awsRestjson1_serializeOpUpdateContact) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactInput(v *UpdateContactInput, encoder *httpbinding.Encoder) error {
@@ -22520,10 +21817,8 @@ func (m *awsRestjson1_serializeOpUpdateContactAttribut
 func (m *awsRestjson1_serializeOpUpdateContactAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22568,7 +21863,6 @@ func (m *awsRestjson1_serializeOpUpdateContactAttribut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactAttributesInput(v *UpdateContactAttributesInput, encoder *httpbinding.Encoder) error {
@@ -22613,10 +21907,8 @@ func (m *awsRestjson1_serializeOpUpdateContactEvaluati
 func (m *awsRestjson1_serializeOpUpdateContactEvaluation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22665,7 +21957,6 @@ func (m *awsRestjson1_serializeOpUpdateContactEvaluati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactEvaluationInput(v *UpdateContactEvaluationInput, encoder *httpbinding.Encoder) error {
@@ -22725,10 +22016,8 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowCont
 func (m *awsRestjson1_serializeOpUpdateContactFlowContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22777,7 +22066,6 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowCont
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactFlowContentInput(v *UpdateContactFlowContentInput, encoder *httpbinding.Encoder) error {
@@ -22828,10 +22116,8 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowMeta
 func (m *awsRestjson1_serializeOpUpdateContactFlowMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22880,7 +22166,6 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowMeta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactFlowMetadataInput(v *UpdateContactFlowMetadataInput, encoder *httpbinding.Encoder) error {
@@ -22941,10 +22226,8 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowModu
 func (m *awsRestjson1_serializeOpUpdateContactFlowModuleContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22993,7 +22276,6 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowModu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactFlowModuleContentInput(v *UpdateContactFlowModuleContentInput, encoder *httpbinding.Encoder) error {
@@ -23044,10 +22326,8 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowModu
 func (m *awsRestjson1_serializeOpUpdateContactFlowModuleMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23096,7 +22376,6 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowModu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactFlowModuleMetadataInput(v *UpdateContactFlowModuleMetadataInput, encoder *httpbinding.Encoder) error {
@@ -23157,10 +22436,8 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowName
 func (m *awsRestjson1_serializeOpUpdateContactFlowName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23209,7 +22486,6 @@ func (m *awsRestjson1_serializeOpUpdateContactFlowName
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactFlowNameInput(v *UpdateContactFlowNameInput, encoder *httpbinding.Encoder) error {
@@ -23265,10 +22541,8 @@ func (m *awsRestjson1_serializeOpUpdateContactRoutingD
 func (m *awsRestjson1_serializeOpUpdateContactRoutingData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23317,7 +22591,6 @@ func (m *awsRestjson1_serializeOpUpdateContactRoutingD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactRoutingDataInput(v *UpdateContactRoutingDataInput, encoder *httpbinding.Encoder) error {
@@ -23380,10 +22653,8 @@ func (m *awsRestjson1_serializeOpUpdateContactSchedule
 func (m *awsRestjson1_serializeOpUpdateContactSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23428,7 +22699,6 @@ func (m *awsRestjson1_serializeOpUpdateContactSchedule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContactScheduleInput(v *UpdateContactScheduleInput, encoder *httpbinding.Encoder) error {
@@ -23471,10 +22741,8 @@ func (m *awsRestjson1_serializeOpUpdateEmailAddressMet
 func (m *awsRestjson1_serializeOpUpdateEmailAddressMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23523,7 +22791,6 @@ func (m *awsRestjson1_serializeOpUpdateEmailAddressMet
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEmailAddressMetadataInput(v *UpdateEmailAddressMetadataInput, encoder *httpbinding.Encoder) error {
@@ -23584,10 +22851,8 @@ func (m *awsRestjson1_serializeOpUpdateEvaluationForm)
 func (m *awsRestjson1_serializeOpUpdateEvaluationForm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23636,7 +22901,6 @@ func (m *awsRestjson1_serializeOpUpdateEvaluationForm)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEvaluationFormInput(v *UpdateEvaluationFormInput, encoder *httpbinding.Encoder) error {
@@ -23721,10 +22985,8 @@ func (m *awsRestjson1_serializeOpUpdateHoursOfOperatio
 func (m *awsRestjson1_serializeOpUpdateHoursOfOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23773,7 +23035,6 @@ func (m *awsRestjson1_serializeOpUpdateHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateHoursOfOperationInput(v *UpdateHoursOfOperationInput, encoder *httpbinding.Encoder) error {
@@ -23841,10 +23102,8 @@ func (m *awsRestjson1_serializeOpUpdateHoursOfOperatio
 func (m *awsRestjson1_serializeOpUpdateHoursOfOperationOverride) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23893,7 +23152,6 @@ func (m *awsRestjson1_serializeOpUpdateHoursOfOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateHoursOfOperationOverrideInput(v *UpdateHoursOfOperationOverrideInput, encoder *httpbinding.Encoder) error {
@@ -23975,10 +23233,8 @@ func (m *awsRestjson1_serializeOpUpdateInstanceAttribu
 func (m *awsRestjson1_serializeOpUpdateInstanceAttribute) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24027,7 +23283,6 @@ func (m *awsRestjson1_serializeOpUpdateInstanceAttribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInstanceAttributeInput(v *UpdateInstanceAttributeInput, encoder *httpbinding.Encoder) error {
@@ -24083,10 +23338,8 @@ func (m *awsRestjson1_serializeOpUpdateInstanceStorage
 func (m *awsRestjson1_serializeOpUpdateInstanceStorageConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24135,7 +23388,6 @@ func (m *awsRestjson1_serializeOpUpdateInstanceStorage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInstanceStorageConfigInput(v *UpdateInstanceStorageConfigInput, encoder *httpbinding.Encoder) error {
@@ -24197,10 +23449,8 @@ func (m *awsRestjson1_serializeOpUpdateParticipantAuth
 func (m *awsRestjson1_serializeOpUpdateParticipantAuthentication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24245,7 +23495,6 @@ func (m *awsRestjson1_serializeOpUpdateParticipantAuth
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateParticipantAuthenticationInput(v *UpdateParticipantAuthenticationInput, encoder *httpbinding.Encoder) error {
@@ -24298,10 +23547,8 @@ func (m *awsRestjson1_serializeOpUpdateParticipantRole
 func (m *awsRestjson1_serializeOpUpdateParticipantRoleConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24350,7 +23597,6 @@ func (m *awsRestjson1_serializeOpUpdateParticipantRole
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateParticipantRoleConfigInput(v *UpdateParticipantRoleConfigInput, encoder *httpbinding.Encoder) error {
@@ -24403,10 +23649,8 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumber) Ha
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
@@ -24455,7 +23699,6 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumber) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePhoneNumberInput(v *UpdatePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -24507,10 +23750,8 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumberMeta
 func (m *awsRestjson1_serializeOpUpdatePhoneNumberMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24559,7 +23800,6 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumberMeta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePhoneNumberMetadataInput(v *UpdatePhoneNumberMetadataInput, encoder *httpbinding.Encoder) error {
@@ -24606,10 +23846,8 @@ func (m *awsRestjson1_serializeOpUpdatePredefinedAttri
 func (m *awsRestjson1_serializeOpUpdatePredefinedAttribute) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24658,7 +23896,6 @@ func (m *awsRestjson1_serializeOpUpdatePredefinedAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePredefinedAttributeInput(v *UpdatePredefinedAttributeInput, encoder *httpbinding.Encoder) error {
@@ -24711,10 +23948,8 @@ func (m *awsRestjson1_serializeOpUpdatePrompt) HandleS
 func (m *awsRestjson1_serializeOpUpdatePrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24763,7 +23998,6 @@ func (m *awsRestjson1_serializeOpUpdatePrompt) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePromptInput(v *UpdatePromptInput, encoder *httpbinding.Encoder) error {
@@ -24824,10 +24058,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueHoursOfOpe
 func (m *awsRestjson1_serializeOpUpdateQueueHoursOfOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24876,7 +24108,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueHoursOfOpe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueHoursOfOperationInput(v *UpdateQueueHoursOfOperationInput, encoder *httpbinding.Encoder) error {
@@ -24927,10 +24158,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueMaxContact
 func (m *awsRestjson1_serializeOpUpdateQueueMaxContacts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -24979,7 +24208,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueMaxContact
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueMaxContactsInput(v *UpdateQueueMaxContactsInput, encoder *httpbinding.Encoder) error {
@@ -25030,10 +24258,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueName) Hand
 func (m *awsRestjson1_serializeOpUpdateQueueName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25082,7 +24308,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueName) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueNameInput(v *UpdateQueueNameInput, encoder *httpbinding.Encoder) error {
@@ -25138,10 +24363,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueOutboundCa
 func (m *awsRestjson1_serializeOpUpdateQueueOutboundCallerConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25190,7 +24413,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueOutboundCa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueOutboundCallerConfigInput(v *UpdateQueueOutboundCallerConfigInput, encoder *httpbinding.Encoder) error {
@@ -25243,10 +24465,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueOutboundEm
 func (m *awsRestjson1_serializeOpUpdateQueueOutboundEmailConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25295,7 +24515,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueOutboundEm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueOutboundEmailConfigInput(v *UpdateQueueOutboundEmailConfigInput, encoder *httpbinding.Encoder) error {
@@ -25348,10 +24567,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueStatus) Ha
 func (m *awsRestjson1_serializeOpUpdateQueueStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25400,7 +24617,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueStatus) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueStatusInput(v *UpdateQueueStatusInput, encoder *httpbinding.Encoder) error {
@@ -25451,10 +24667,8 @@ func (m *awsRestjson1_serializeOpUpdateQuickConnectCon
 func (m *awsRestjson1_serializeOpUpdateQuickConnectConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25503,7 +24717,6 @@ func (m *awsRestjson1_serializeOpUpdateQuickConnectCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQuickConnectConfigInput(v *UpdateQuickConnectConfigInput, encoder *httpbinding.Encoder) error {
@@ -25556,10 +24769,8 @@ func (m *awsRestjson1_serializeOpUpdateQuickConnectNam
 func (m *awsRestjson1_serializeOpUpdateQuickConnectName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25608,7 +24819,6 @@ func (m *awsRestjson1_serializeOpUpdateQuickConnectNam
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQuickConnectNameInput(v *UpdateQuickConnectNameInput, encoder *httpbinding.Encoder) error {
@@ -25664,10 +24874,8 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileA
 func (m *awsRestjson1_serializeOpUpdateRoutingProfileAgentAvailabilityTimer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25716,7 +24924,6 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoutingProfileAgentAvailabilityTimerInput(v *UpdateRoutingProfileAgentAvailabilityTimerInput, encoder *httpbinding.Encoder) error {
@@ -25767,10 +24974,8 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileC
 func (m *awsRestjson1_serializeOpUpdateRoutingProfileConcurrency) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25819,7 +25024,6 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoutingProfileConcurrencyInput(v *UpdateRoutingProfileConcurrencyInput, encoder *httpbinding.Encoder) error {
@@ -25872,10 +25076,8 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileD
 func (m *awsRestjson1_serializeOpUpdateRoutingProfileDefaultOutboundQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -25924,7 +25126,6 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoutingProfileDefaultOutboundQueueInput(v *UpdateRoutingProfileDefaultOutboundQueueInput, encoder *httpbinding.Encoder) error {
@@ -25975,10 +25176,8 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileN
 func (m *awsRestjson1_serializeOpUpdateRoutingProfileName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26027,7 +25226,6 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileN
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoutingProfileNameInput(v *UpdateRoutingProfileNameInput, encoder *httpbinding.Encoder) error {
@@ -26083,10 +25281,8 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileQ
 func (m *awsRestjson1_serializeOpUpdateRoutingProfileQueues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26135,7 +25331,6 @@ func (m *awsRestjson1_serializeOpUpdateRoutingProfileQ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoutingProfileQueuesInput(v *UpdateRoutingProfileQueuesInput, encoder *httpbinding.Encoder) error {
@@ -26188,10 +25383,8 @@ func (m *awsRestjson1_serializeOpUpdateRule) HandleSer
 func (m *awsRestjson1_serializeOpUpdateRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26240,7 +25433,6 @@ func (m *awsRestjson1_serializeOpUpdateRule) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRuleInput(v *UpdateRuleInput, encoder *httpbinding.Encoder) error {
@@ -26308,10 +25500,8 @@ func (m *awsRestjson1_serializeOpUpdateSecurityProfile
 func (m *awsRestjson1_serializeOpUpdateSecurityProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26360,7 +25550,6 @@ func (m *awsRestjson1_serializeOpUpdateSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSecurityProfileInput(v *UpdateSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -26451,10 +25640,8 @@ func (m *awsRestjson1_serializeOpUpdateTaskTemplate) H
 func (m *awsRestjson1_serializeOpUpdateTaskTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26503,7 +25690,6 @@ func (m *awsRestjson1_serializeOpUpdateTaskTemplate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTaskTemplateInput(v *UpdateTaskTemplateInput, encoder *httpbinding.Encoder) error {
@@ -26595,10 +25781,8 @@ func (m *awsRestjson1_serializeOpUpdateTrafficDistribu
 func (m *awsRestjson1_serializeOpUpdateTrafficDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26647,7 +25831,6 @@ func (m *awsRestjson1_serializeOpUpdateTrafficDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTrafficDistributionInput(v *UpdateTrafficDistributionInput, encoder *httpbinding.Encoder) error {
@@ -26705,10 +25888,8 @@ func (m *awsRestjson1_serializeOpUpdateUserHierarchy) 
 func (m *awsRestjson1_serializeOpUpdateUserHierarchy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26757,7 +25938,6 @@ func (m *awsRestjson1_serializeOpUpdateUserHierarchy) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserHierarchyInput(v *UpdateUserHierarchyInput, encoder *httpbinding.Encoder) error {
@@ -26808,10 +25988,8 @@ func (m *awsRestjson1_serializeOpUpdateUserHierarchyGr
 func (m *awsRestjson1_serializeOpUpdateUserHierarchyGroupName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26860,7 +26038,6 @@ func (m *awsRestjson1_serializeOpUpdateUserHierarchyGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserHierarchyGroupNameInput(v *UpdateUserHierarchyGroupNameInput, encoder *httpbinding.Encoder) error {
@@ -26911,10 +26088,8 @@ func (m *awsRestjson1_serializeOpUpdateUserHierarchySt
 func (m *awsRestjson1_serializeOpUpdateUserHierarchyStructure) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -26963,7 +26138,6 @@ func (m *awsRestjson1_serializeOpUpdateUserHierarchySt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserHierarchyStructureInput(v *UpdateUserHierarchyStructureInput, encoder *httpbinding.Encoder) error {
@@ -27007,10 +26181,8 @@ func (m *awsRestjson1_serializeOpUpdateUserIdentityInf
 func (m *awsRestjson1_serializeOpUpdateUserIdentityInfo) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -27059,7 +26231,6 @@ func (m *awsRestjson1_serializeOpUpdateUserIdentityInf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserIdentityInfoInput(v *UpdateUserIdentityInfoInput, encoder *httpbinding.Encoder) error {
@@ -27112,10 +26283,8 @@ func (m *awsRestjson1_serializeOpUpdateUserPhoneConfig
 func (m *awsRestjson1_serializeOpUpdateUserPhoneConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -27164,7 +26333,6 @@ func (m *awsRestjson1_serializeOpUpdateUserPhoneConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserPhoneConfigInput(v *UpdateUserPhoneConfigInput, encoder *httpbinding.Encoder) error {
@@ -27217,10 +26385,8 @@ func (m *awsRestjson1_serializeOpUpdateUserProficienci
 func (m *awsRestjson1_serializeOpUpdateUserProficiencies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -27269,7 +26435,6 @@ func (m *awsRestjson1_serializeOpUpdateUserProficienci
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserProficienciesInput(v *UpdateUserProficienciesInput, encoder *httpbinding.Encoder) error {
@@ -27322,10 +26487,8 @@ func (m *awsRestjson1_serializeOpUpdateUserRoutingProf
 func (m *awsRestjson1_serializeOpUpdateUserRoutingProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -27374,7 +26537,6 @@ func (m *awsRestjson1_serializeOpUpdateUserRoutingProf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserRoutingProfileInput(v *UpdateUserRoutingProfileInput, encoder *httpbinding.Encoder) error {
@@ -27425,10 +26587,8 @@ func (m *awsRestjson1_serializeOpUpdateUserSecurityPro
 func (m *awsRestjson1_serializeOpUpdateUserSecurityProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -27477,7 +26637,6 @@ func (m *awsRestjson1_serializeOpUpdateUserSecurityPro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserSecurityProfilesInput(v *UpdateUserSecurityProfilesInput, encoder *httpbinding.Encoder) error {
@@ -27530,10 +26689,8 @@ func (m *awsRestjson1_serializeOpUpdateViewContent) Ha
 func (m *awsRestjson1_serializeOpUpdateViewContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -27582,7 +26739,6 @@ func (m *awsRestjson1_serializeOpUpdateViewContent) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateViewContentInput(v *UpdateViewContentInput, encoder *httpbinding.Encoder) error {
@@ -27640,10 +26796,8 @@ func (m *awsRestjson1_serializeOpUpdateViewMetadata) H
 func (m *awsRestjson1_serializeOpUpdateViewMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -27692,7 +26846,6 @@ func (m *awsRestjson1_serializeOpUpdateViewMetadata) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateViewMetadataInput(v *UpdateViewMetadataInput, encoder *httpbinding.Encoder) error {
