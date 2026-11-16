--- vendor/github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpBatchCreateCustomVoca
 func (m *awsRestjson1_serializeOpBatchCreateCustomVocabularyItem) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpBatchCreateCustomVoca
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchCreateCustomVocabularyItemInput(v *BatchCreateCustomVocabularyItemInput, encoder *httpbinding.Encoder) error {
@@ -141,10 +137,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteCustomVoca
 func (m *awsRestjson1_serializeOpBatchDeleteCustomVocabularyItem) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -193,7 +187,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteCustomVoca
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteCustomVocabularyItemInput(v *BatchDeleteCustomVocabularyItemInput, encoder *httpbinding.Encoder) error {
@@ -255,10 +248,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateCustomVoca
 func (m *awsRestjson1_serializeOpBatchUpdateCustomVocabularyItem) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -307,7 +298,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateCustomVoca
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateCustomVocabularyItemInput(v *BatchUpdateCustomVocabularyItemInput, encoder *httpbinding.Encoder) error {
@@ -369,10 +359,8 @@ func (m *awsRestjson1_serializeOpBuildBotLocale) Handl
 func (m *awsRestjson1_serializeOpBuildBotLocale) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -410,7 +398,6 @@ func (m *awsRestjson1_serializeOpBuildBotLocale) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBuildBotLocaleInput(v *BuildBotLocaleInput, encoder *httpbinding.Encoder) error {
@@ -458,10 +445,8 @@ func (m *awsRestjson1_serializeOpCreateBot) HandleSeri
 func (m *awsRestjson1_serializeOpCreateBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -506,7 +491,6 @@ func (m *awsRestjson1_serializeOpCreateBot) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBotInput(v *CreateBotInput, encoder *httpbinding.Encoder) error {
@@ -594,10 +578,8 @@ func (m *awsRestjson1_serializeOpCreateBotAlias) Handl
 func (m *awsRestjson1_serializeOpCreateBotAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -646,7 +628,6 @@ func (m *awsRestjson1_serializeOpCreateBotAlias) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBotAliasInput(v *CreateBotAliasInput, encoder *httpbinding.Encoder) error {
@@ -726,10 +707,8 @@ func (m *awsRestjson1_serializeOpCreateBotLocale) Hand
 func (m *awsRestjson1_serializeOpCreateBotLocale) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -778,7 +757,6 @@ func (m *awsRestjson1_serializeOpCreateBotLocale) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBotLocaleInput(v *CreateBotLocaleInput, encoder *httpbinding.Encoder) error {
@@ -866,10 +844,8 @@ func (m *awsRestjson1_serializeOpCreateBotReplica) Han
 func (m *awsRestjson1_serializeOpCreateBotReplica) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -918,7 +894,6 @@ func (m *awsRestjson1_serializeOpCreateBotReplica) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBotReplicaInput(v *CreateBotReplicaInput, encoder *httpbinding.Encoder) error {
@@ -960,10 +935,8 @@ func (m *awsRestjson1_serializeOpCreateBotVersion) Han
 func (m *awsRestjson1_serializeOpCreateBotVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1012,7 +985,6 @@ func (m *awsRestjson1_serializeOpCreateBotVersion) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBotVersionInput(v *CreateBotVersionInput, encoder *httpbinding.Encoder) error {
@@ -1061,10 +1033,8 @@ func (m *awsRestjson1_serializeOpCreateExport) HandleS
 func (m *awsRestjson1_serializeOpCreateExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1109,7 +1079,6 @@ func (m *awsRestjson1_serializeOpCreateExport) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateExportInput(v *CreateExportInput, encoder *httpbinding.Encoder) error {
@@ -1154,10 +1123,8 @@ func (m *awsRestjson1_serializeOpCreateIntent) HandleS
 func (m *awsRestjson1_serializeOpCreateIntent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1206,7 +1173,6 @@ func (m *awsRestjson1_serializeOpCreateIntent) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIntentInput(v *CreateIntentInput, encoder *httpbinding.Encoder) error {
@@ -1353,10 +1319,8 @@ func (m *awsRestjson1_serializeOpCreateResourcePolicy)
 func (m *awsRestjson1_serializeOpCreateResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1405,7 +1369,6 @@ func (m *awsRestjson1_serializeOpCreateResourcePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateResourcePolicyInput(v *CreateResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -1447,10 +1410,8 @@ func (m *awsRestjson1_serializeOpCreateResourcePolicyS
 func (m *awsRestjson1_serializeOpCreateResourcePolicyStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1499,7 +1460,6 @@ func (m *awsRestjson1_serializeOpCreateResourcePolicyS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateResourcePolicyStatementInput(v *CreateResourcePolicyStatementInput, encoder *httpbinding.Encoder) error {
@@ -1571,10 +1531,8 @@ func (m *awsRestjson1_serializeOpCreateSlot) HandleSer
 func (m *awsRestjson1_serializeOpCreateSlot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1623,7 +1581,6 @@ func (m *awsRestjson1_serializeOpCreateSlot) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSlotInput(v *CreateSlotInput, encoder *httpbinding.Encoder) error {
@@ -1730,10 +1687,8 @@ func (m *awsRestjson1_serializeOpCreateSlotType) Handl
 func (m *awsRestjson1_serializeOpCreateSlotType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1782,7 +1737,6 @@ func (m *awsRestjson1_serializeOpCreateSlotType) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSlotTypeInput(v *CreateSlotTypeInput, encoder *httpbinding.Encoder) error {
@@ -1880,10 +1834,8 @@ func (m *awsRestjson1_serializeOpCreateTestSetDiscrepa
 func (m *awsRestjson1_serializeOpCreateTestSetDiscrepancyReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1932,7 +1884,6 @@ func (m *awsRestjson1_serializeOpCreateTestSetDiscrepa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTestSetDiscrepancyReportInput(v *CreateTestSetDiscrepancyReportInput, encoder *httpbinding.Encoder) error {
@@ -1976,10 +1927,8 @@ func (m *awsRestjson1_serializeOpCreateUploadUrl) Hand
 func (m *awsRestjson1_serializeOpCreateUploadUrl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2013,7 +1962,6 @@ func (m *awsRestjson1_serializeOpCreateUploadUrl) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUploadUrlInput(v *CreateUploadUrlInput, encoder *httpbinding.Encoder) error {
@@ -2034,10 +1982,8 @@ func (m *awsRestjson1_serializeOpDeleteBot) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2075,7 +2021,6 @@ func (m *awsRestjson1_serializeOpDeleteBot) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBotInput(v *DeleteBotInput, encoder *httpbinding.Encoder) error {
@@ -2109,10 +2054,8 @@ func (m *awsRestjson1_serializeOpDeleteBotAlias) Handl
 func (m *awsRestjson1_serializeOpDeleteBotAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2150,7 +2093,6 @@ func (m *awsRestjson1_serializeOpDeleteBotAlias) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBotAliasInput(v *DeleteBotAliasInput, encoder *httpbinding.Encoder) error {
@@ -2193,10 +2135,8 @@ func (m *awsRestjson1_serializeOpDeleteBotLocale) Hand
 func (m *awsRestjson1_serializeOpDeleteBotLocale) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2234,7 +2174,6 @@ func (m *awsRestjson1_serializeOpDeleteBotLocale) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBotLocaleInput(v *DeleteBotLocaleInput, encoder *httpbinding.Encoder) error {
@@ -2282,10 +2221,8 @@ func (m *awsRestjson1_serializeOpDeleteBotReplica) Han
 func (m *awsRestjson1_serializeOpDeleteBotReplica) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2323,7 +2260,6 @@ func (m *awsRestjson1_serializeOpDeleteBotReplica) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBotReplicaInput(v *DeleteBotReplicaInput, encoder *httpbinding.Encoder) error {
@@ -2362,10 +2298,8 @@ func (m *awsRestjson1_serializeOpDeleteBotVersion) Han
 func (m *awsRestjson1_serializeOpDeleteBotVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2403,7 +2337,6 @@ func (m *awsRestjson1_serializeOpDeleteBotVersion) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBotVersionInput(v *DeleteBotVersionInput, encoder *httpbinding.Encoder) error {
@@ -2446,10 +2379,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomVocabular
 func (m *awsRestjson1_serializeOpDeleteCustomVocabulary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2487,7 +2418,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomVocabular
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomVocabularyInput(v *DeleteCustomVocabularyInput, encoder *httpbinding.Encoder) error {
@@ -2535,10 +2465,8 @@ func (m *awsRestjson1_serializeOpDeleteExport) HandleS
 func (m *awsRestjson1_serializeOpDeleteExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2576,7 +2504,6 @@ func (m *awsRestjson1_serializeOpDeleteExport) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteExportInput(v *DeleteExportInput, encoder *httpbinding.Encoder) error {
@@ -2606,10 +2533,8 @@ func (m *awsRestjson1_serializeOpDeleteImport) HandleS
 func (m *awsRestjson1_serializeOpDeleteImport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2647,7 +2572,6 @@ func (m *awsRestjson1_serializeOpDeleteImport) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImportInput(v *DeleteImportInput, encoder *httpbinding.Encoder) error {
@@ -2677,10 +2601,8 @@ func (m *awsRestjson1_serializeOpDeleteIntent) HandleS
 func (m *awsRestjson1_serializeOpDeleteIntent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2718,7 +2640,6 @@ func (m *awsRestjson1_serializeOpDeleteIntent) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIntentInput(v *DeleteIntentInput, encoder *httpbinding.Encoder) error {
@@ -2775,10 +2696,8 @@ func (m *awsRestjson1_serializeOpDeleteResourcePolicy)
 func (m *awsRestjson1_serializeOpDeleteResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2816,7 +2735,6 @@ func (m *awsRestjson1_serializeOpDeleteResourcePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResourcePolicyInput(v *DeleteResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2850,10 +2768,8 @@ func (m *awsRestjson1_serializeOpDeleteResourcePolicyS
 func (m *awsRestjson1_serializeOpDeleteResourcePolicyStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2891,7 +2807,6 @@ func (m *awsRestjson1_serializeOpDeleteResourcePolicyS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResourcePolicyStatementInput(v *DeleteResourcePolicyStatementInput, encoder *httpbinding.Encoder) error {
@@ -2934,10 +2849,8 @@ func (m *awsRestjson1_serializeOpDeleteSlot) HandleSer
 func (m *awsRestjson1_serializeOpDeleteSlot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2975,7 +2888,6 @@ func (m *awsRestjson1_serializeOpDeleteSlot) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSlotInput(v *DeleteSlotInput, encoder *httpbinding.Encoder) error {
@@ -3041,10 +2953,8 @@ func (m *awsRestjson1_serializeOpDeleteSlotType) Handl
 func (m *awsRestjson1_serializeOpDeleteSlotType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3082,7 +2992,6 @@ func (m *awsRestjson1_serializeOpDeleteSlotType) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSlotTypeInput(v *DeleteSlotTypeInput, encoder *httpbinding.Encoder) error {
@@ -3143,10 +3052,8 @@ func (m *awsRestjson1_serializeOpDeleteTestSet) Handle
 func (m *awsRestjson1_serializeOpDeleteTestSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3184,7 +3091,6 @@ func (m *awsRestjson1_serializeOpDeleteTestSet) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTestSetInput(v *DeleteTestSetInput, encoder *httpbinding.Encoder) error {
@@ -3214,10 +3120,8 @@ func (m *awsRestjson1_serializeOpDeleteUtterances) Han
 func (m *awsRestjson1_serializeOpDeleteUtterances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3255,7 +3159,6 @@ func (m *awsRestjson1_serializeOpDeleteUtterances) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUtterancesInput(v *DeleteUtterancesInput, encoder *httpbinding.Encoder) error {
@@ -3293,10 +3196,8 @@ func (m *awsRestjson1_serializeOpDescribeBot) HandleSe
 func (m *awsRestjson1_serializeOpDescribeBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3334,7 +3235,6 @@ func (m *awsRestjson1_serializeOpDescribeBot) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBotInput(v *DescribeBotInput, encoder *httpbinding.Encoder) error {
@@ -3364,10 +3264,8 @@ func (m *awsRestjson1_serializeOpDescribeBotAlias) Han
 func (m *awsRestjson1_serializeOpDescribeBotAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3405,7 +3303,6 @@ func (m *awsRestjson1_serializeOpDescribeBotAlias) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBotAliasInput(v *DescribeBotAliasInput, encoder *httpbinding.Encoder) error {
@@ -3444,10 +3341,8 @@ func (m *awsRestjson1_serializeOpDescribeBotLocale) Ha
 func (m *awsRestjson1_serializeOpDescribeBotLocale) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3485,7 +3380,6 @@ func (m *awsRestjson1_serializeOpDescribeBotLocale) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBotLocaleInput(v *DescribeBotLocaleInput, encoder *httpbinding.Encoder) error {
@@ -3533,10 +3427,8 @@ func (m *awsRestjson1_serializeOpDescribeBotRecommenda
 func (m *awsRestjson1_serializeOpDescribeBotRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3574,7 +3466,6 @@ func (m *awsRestjson1_serializeOpDescribeBotRecommenda
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBotRecommendationInput(v *DescribeBotRecommendationInput, encoder *httpbinding.Encoder) error {
@@ -3631,10 +3522,8 @@ func (m *awsRestjson1_serializeOpDescribeBotReplica) H
 func (m *awsRestjson1_serializeOpDescribeBotReplica) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3672,7 +3561,6 @@ func (m *awsRestjson1_serializeOpDescribeBotReplica) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBotReplicaInput(v *DescribeBotReplicaInput, encoder *httpbinding.Encoder) error {
@@ -3711,10 +3599,8 @@ func (m *awsRestjson1_serializeOpDescribeBotResourceGe
 func (m *awsRestjson1_serializeOpDescribeBotResourceGeneration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3752,7 +3638,6 @@ func (m *awsRestjson1_serializeOpDescribeBotResourceGe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBotResourceGenerationInput(v *DescribeBotResourceGenerationInput, encoder *httpbinding.Encoder) error {
@@ -3809,10 +3694,8 @@ func (m *awsRestjson1_serializeOpDescribeBotVersion) H
 func (m *awsRestjson1_serializeOpDescribeBotVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3850,7 +3733,6 @@ func (m *awsRestjson1_serializeOpDescribeBotVersion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBotVersionInput(v *DescribeBotVersionInput, encoder *httpbinding.Encoder) error {
@@ -3889,10 +3771,8 @@ func (m *awsRestjson1_serializeOpDescribeCustomVocabul
 func (m *awsRestjson1_serializeOpDescribeCustomVocabularyMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3930,7 +3810,6 @@ func (m *awsRestjson1_serializeOpDescribeCustomVocabul
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeCustomVocabularyMetadataInput(v *DescribeCustomVocabularyMetadataInput, encoder *httpbinding.Encoder) error {
@@ -3978,10 +3857,8 @@ func (m *awsRestjson1_serializeOpDescribeExport) Handl
 func (m *awsRestjson1_serializeOpDescribeExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4019,7 +3896,6 @@ func (m *awsRestjson1_serializeOpDescribeExport) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeExportInput(v *DescribeExportInput, encoder *httpbinding.Encoder) error {
@@ -4049,10 +3925,8 @@ func (m *awsRestjson1_serializeOpDescribeImport) Handl
 func (m *awsRestjson1_serializeOpDescribeImport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4090,7 +3964,6 @@ func (m *awsRestjson1_serializeOpDescribeImport) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeImportInput(v *DescribeImportInput, encoder *httpbinding.Encoder) error {
@@ -4120,10 +3993,8 @@ func (m *awsRestjson1_serializeOpDescribeIntent) Handl
 func (m *awsRestjson1_serializeOpDescribeIntent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4161,7 +4032,6 @@ func (m *awsRestjson1_serializeOpDescribeIntent) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeIntentInput(v *DescribeIntentInput, encoder *httpbinding.Encoder) error {
@@ -4218,10 +4088,8 @@ func (m *awsRestjson1_serializeOpDescribeResourcePolic
 func (m *awsRestjson1_serializeOpDescribeResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4259,7 +4127,6 @@ func (m *awsRestjson1_serializeOpDescribeResourcePolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeResourcePolicyInput(v *DescribeResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -4289,10 +4156,8 @@ func (m *awsRestjson1_serializeOpDescribeSlot) HandleS
 func (m *awsRestjson1_serializeOpDescribeSlot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4330,7 +4195,6 @@ func (m *awsRestjson1_serializeOpDescribeSlot) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSlotInput(v *DescribeSlotInput, encoder *httpbinding.Encoder) error {
@@ -4396,10 +4260,8 @@ func (m *awsRestjson1_serializeOpDescribeSlotType) Han
 func (m *awsRestjson1_serializeOpDescribeSlotType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4437,7 +4299,6 @@ func (m *awsRestjson1_serializeOpDescribeSlotType) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSlotTypeInput(v *DescribeSlotTypeInput, encoder *httpbinding.Encoder) error {
@@ -4494,10 +4355,8 @@ func (m *awsRestjson1_serializeOpDescribeTestExecution
 func (m *awsRestjson1_serializeOpDescribeTestExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4535,7 +4394,6 @@ func (m *awsRestjson1_serializeOpDescribeTestExecution
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTestExecutionInput(v *DescribeTestExecutionInput, encoder *httpbinding.Encoder) error {
@@ -4565,10 +4423,8 @@ func (m *awsRestjson1_serializeOpDescribeTestSet) Hand
 func (m *awsRestjson1_serializeOpDescribeTestSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4606,7 +4462,6 @@ func (m *awsRestjson1_serializeOpDescribeTestSet) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTestSetInput(v *DescribeTestSetInput, encoder *httpbinding.Encoder) error {
@@ -4636,10 +4491,8 @@ func (m *awsRestjson1_serializeOpDescribeTestSetDiscre
 func (m *awsRestjson1_serializeOpDescribeTestSetDiscrepancyReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4677,7 +4530,6 @@ func (m *awsRestjson1_serializeOpDescribeTestSetDiscre
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTestSetDiscrepancyReportInput(v *DescribeTestSetDiscrepancyReportInput, encoder *httpbinding.Encoder) error {
@@ -4707,10 +4559,8 @@ func (m *awsRestjson1_serializeOpDescribeTestSetGenera
 func (m *awsRestjson1_serializeOpDescribeTestSetGeneration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4748,7 +4598,6 @@ func (m *awsRestjson1_serializeOpDescribeTestSetGenera
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTestSetGenerationInput(v *DescribeTestSetGenerationInput, encoder *httpbinding.Encoder) error {
@@ -4778,10 +4627,8 @@ func (m *awsRestjson1_serializeOpGenerateBotElement) H
 func (m *awsRestjson1_serializeOpGenerateBotElement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4830,7 +4677,6 @@ func (m *awsRestjson1_serializeOpGenerateBotElement) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateBotElementInput(v *GenerateBotElementInput, encoder *httpbinding.Encoder) error {
@@ -4890,10 +4736,8 @@ func (m *awsRestjson1_serializeOpGetTestExecutionArtif
 func (m *awsRestjson1_serializeOpGetTestExecutionArtifactsUrl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4931,7 +4775,6 @@ func (m *awsRestjson1_serializeOpGetTestExecutionArtif
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTestExecutionArtifactsUrlInput(v *GetTestExecutionArtifactsUrlInput, encoder *httpbinding.Encoder) error {
@@ -4961,10 +4804,8 @@ func (m *awsRestjson1_serializeOpListAggregatedUtteran
 func (m *awsRestjson1_serializeOpListAggregatedUtterances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5013,7 +4854,6 @@ func (m *awsRestjson1_serializeOpListAggregatedUtteran
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAggregatedUtterancesInput(v *ListAggregatedUtterancesInput, encoder *httpbinding.Encoder) error {
@@ -5096,10 +4936,8 @@ func (m *awsRestjson1_serializeOpListBotAliases) Handl
 func (m *awsRestjson1_serializeOpListBotAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5148,7 +4986,6 @@ func (m *awsRestjson1_serializeOpListBotAliases) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotAliasesInput(v *ListBotAliasesInput, encoder *httpbinding.Encoder) error {
@@ -5195,10 +5032,8 @@ func (m *awsRestjson1_serializeOpListBotAliasReplicas)
 func (m *awsRestjson1_serializeOpListBotAliasReplicas) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5247,7 +5082,6 @@ func (m *awsRestjson1_serializeOpListBotAliasReplicas)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotAliasReplicasInput(v *ListBotAliasReplicasInput, encoder *httpbinding.Encoder) error {
@@ -5303,10 +5137,8 @@ func (m *awsRestjson1_serializeOpListBotLocales) Handl
 func (m *awsRestjson1_serializeOpListBotLocales) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5355,7 +5187,6 @@ func (m *awsRestjson1_serializeOpListBotLocales) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotLocalesInput(v *ListBotLocalesInput, encoder *httpbinding.Encoder) error {
@@ -5425,10 +5256,8 @@ func (m *awsRestjson1_serializeOpListBotRecommendation
 func (m *awsRestjson1_serializeOpListBotRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5477,7 +5306,6 @@ func (m *awsRestjson1_serializeOpListBotRecommendation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotRecommendationsInput(v *ListBotRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -5542,10 +5370,8 @@ func (m *awsRestjson1_serializeOpListBotReplicas) Hand
 func (m *awsRestjson1_serializeOpListBotReplicas) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5583,7 +5409,6 @@ func (m *awsRestjson1_serializeOpListBotReplicas) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotReplicasInput(v *ListBotReplicasInput, encoder *httpbinding.Encoder) error {
@@ -5613,10 +5438,8 @@ func (m *awsRestjson1_serializeOpListBotResourceGenera
 func (m *awsRestjson1_serializeOpListBotResourceGenerations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5665,7 +5488,6 @@ func (m *awsRestjson1_serializeOpListBotResourceGenera
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotResourceGenerationsInput(v *ListBotResourceGenerationsInput, encoder *httpbinding.Encoder) error {
@@ -5737,10 +5559,8 @@ func (m *awsRestjson1_serializeOpListBots) HandleSeria
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
@@ -5785,7 +5605,6 @@ func (m *awsRestjson1_serializeOpListBots) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotsInput(v *ListBotsInput, encoder *httpbinding.Encoder) error {
@@ -5837,10 +5656,8 @@ func (m *awsRestjson1_serializeOpListBotVersionReplica
 func (m *awsRestjson1_serializeOpListBotVersionReplicas) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5889,7 +5706,6 @@ func (m *awsRestjson1_serializeOpListBotVersionReplica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotVersionReplicasInput(v *ListBotVersionReplicasInput, encoder *httpbinding.Encoder) error {
@@ -5952,10 +5768,8 @@ func (m *awsRestjson1_serializeOpListBotVersions) Hand
 func (m *awsRestjson1_serializeOpListBotVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6004,7 +5818,6 @@ func (m *awsRestjson1_serializeOpListBotVersions) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotVersionsInput(v *ListBotVersionsInput, encoder *httpbinding.Encoder) error {
@@ -6058,10 +5871,8 @@ func (m *awsRestjson1_serializeOpListBuiltInIntents) H
 func (m *awsRestjson1_serializeOpListBuiltInIntents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6110,7 +5921,6 @@ func (m *awsRestjson1_serializeOpListBuiltInIntents) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBuiltInIntentsInput(v *ListBuiltInIntentsInput, encoder *httpbinding.Encoder) error {
@@ -6164,10 +5974,8 @@ func (m *awsRestjson1_serializeOpListBuiltInSlotTypes)
 func (m *awsRestjson1_serializeOpListBuiltInSlotTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6216,7 +6024,6 @@ func (m *awsRestjson1_serializeOpListBuiltInSlotTypes)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBuiltInSlotTypesInput(v *ListBuiltInSlotTypesInput, encoder *httpbinding.Encoder) error {
@@ -6270,10 +6077,8 @@ func (m *awsRestjson1_serializeOpListCustomVocabularyI
 func (m *awsRestjson1_serializeOpListCustomVocabularyItems) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6322,7 +6127,6 @@ func (m *awsRestjson1_serializeOpListCustomVocabularyI
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomVocabularyItemsInput(v *ListCustomVocabularyItemsInput, encoder *httpbinding.Encoder) error {
@@ -6387,10 +6191,8 @@ func (m *awsRestjson1_serializeOpListExports) HandleSe
 func (m *awsRestjson1_serializeOpListExports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6435,7 +6237,6 @@ func (m *awsRestjson1_serializeOpListExports) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListExportsInput(v *ListExportsInput, encoder *httpbinding.Encoder) error {
@@ -6502,10 +6303,8 @@ func (m *awsRestjson1_serializeOpListImports) HandleSe
 func (m *awsRestjson1_serializeOpListImports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6550,7 +6349,6 @@ func (m *awsRestjson1_serializeOpListImports) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImportsInput(v *ListImportsInput, encoder *httpbinding.Encoder) error {
@@ -6617,10 +6415,8 @@ func (m *awsRestjson1_serializeOpListIntentMetrics) Ha
 func (m *awsRestjson1_serializeOpListIntentMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6669,7 +6465,6 @@ func (m *awsRestjson1_serializeOpListIntentMetrics) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIntentMetricsInput(v *ListIntentMetricsInput, encoder *httpbinding.Encoder) error {
@@ -6754,10 +6549,8 @@ func (m *awsRestjson1_serializeOpListIntentPaths) Hand
 func (m *awsRestjson1_serializeOpListIntentPaths) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6806,7 +6599,6 @@ func (m *awsRestjson1_serializeOpListIntentPaths) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIntentPathsInput(v *ListIntentPathsInput, encoder *httpbinding.Encoder) error {
@@ -6865,10 +6657,8 @@ func (m *awsRestjson1_serializeOpListIntents) HandleSe
 func (m *awsRestjson1_serializeOpListIntents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6917,7 +6707,6 @@ func (m *awsRestjson1_serializeOpListIntents) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIntentsInput(v *ListIntentsInput, encoder *httpbinding.Encoder) error {
@@ -6996,10 +6785,8 @@ func (m *awsRestjson1_serializeOpListIntentStageMetric
 func (m *awsRestjson1_serializeOpListIntentStageMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7048,7 +6835,6 @@ func (m *awsRestjson1_serializeOpListIntentStageMetric
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIntentStageMetricsInput(v *ListIntentStageMetricsInput, encoder *httpbinding.Encoder) error {
@@ -7133,10 +6919,8 @@ func (m *awsRestjson1_serializeOpListRecommendedIntent
 func (m *awsRestjson1_serializeOpListRecommendedIntents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7185,7 +6969,6 @@ func (m *awsRestjson1_serializeOpListRecommendedIntent
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecommendedIntentsInput(v *ListRecommendedIntentsInput, encoder *httpbinding.Encoder) error {
@@ -7259,10 +7042,8 @@ func (m *awsRestjson1_serializeOpListSessionAnalyticsD
 func (m *awsRestjson1_serializeOpListSessionAnalyticsData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7311,7 +7092,6 @@ func (m *awsRestjson1_serializeOpListSessionAnalyticsD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSessionAnalyticsDataInput(v *ListSessionAnalyticsDataInput, encoder *httpbinding.Encoder) error {
@@ -7382,10 +7162,8 @@ func (m *awsRestjson1_serializeOpListSessionMetrics) H
 func (m *awsRestjson1_serializeOpListSessionMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7434,7 +7212,6 @@ func (m *awsRestjson1_serializeOpListSessionMetrics) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSessionMetricsInput(v *ListSessionMetricsInput, encoder *httpbinding.Encoder) error {
@@ -7519,10 +7296,8 @@ func (m *awsRestjson1_serializeOpListSlots) HandleSeri
 func (m *awsRestjson1_serializeOpListSlots) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7571,7 +7346,6 @@ func (m *awsRestjson1_serializeOpListSlots) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSlotsInput(v *ListSlotsInput, encoder *httpbinding.Encoder) error {
@@ -7659,10 +7433,8 @@ func (m *awsRestjson1_serializeOpListSlotTypes) Handle
 func (m *awsRestjson1_serializeOpListSlotTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7711,7 +7483,6 @@ func (m *awsRestjson1_serializeOpListSlotTypes) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSlotTypesInput(v *ListSlotTypesInput, encoder *httpbinding.Encoder) error {
@@ -7790,10 +7561,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -7831,7 +7600,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -7861,10 +7629,8 @@ func (m *awsRestjson1_serializeOpListTestExecutionResu
 func (m *awsRestjson1_serializeOpListTestExecutionResultItems) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7913,7 +7679,6 @@ func (m *awsRestjson1_serializeOpListTestExecutionResu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTestExecutionResultItemsInput(v *ListTestExecutionResultItemsInput, encoder *httpbinding.Encoder) error {
@@ -7967,10 +7732,8 @@ func (m *awsRestjson1_serializeOpListTestExecutions) H
 func (m *awsRestjson1_serializeOpListTestExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8015,7 +7778,6 @@ func (m *awsRestjson1_serializeOpListTestExecutions) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTestExecutionsInput(v *ListTestExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -8060,10 +7822,8 @@ func (m *awsRestjson1_serializeOpListTestSetRecords) H
 func (m *awsRestjson1_serializeOpListTestSetRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8112,7 +7872,6 @@ func (m *awsRestjson1_serializeOpListTestSetRecords) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTestSetRecordsInput(v *ListTestSetRecordsInput, encoder *httpbinding.Encoder) error {
@@ -8159,10 +7918,8 @@ func (m *awsRestjson1_serializeOpListTestSets) HandleS
 func (m *awsRestjson1_serializeOpListTestSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8207,7 +7964,6 @@ func (m *awsRestjson1_serializeOpListTestSets) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTestSetsInput(v *ListTestSetsInput, encoder *httpbinding.Encoder) error {
@@ -8252,10 +8008,8 @@ func (m *awsRestjson1_serializeOpListUtteranceAnalytic
 func (m *awsRestjson1_serializeOpListUtteranceAnalyticsData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8304,7 +8058,6 @@ func (m *awsRestjson1_serializeOpListUtteranceAnalytic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUtteranceAnalyticsDataInput(v *ListUtteranceAnalyticsDataInput, encoder *httpbinding.Encoder) error {
@@ -8375,10 +8128,8 @@ func (m *awsRestjson1_serializeOpListUtteranceMetrics)
 func (m *awsRestjson1_serializeOpListUtteranceMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8427,7 +8178,6 @@ func (m *awsRestjson1_serializeOpListUtteranceMetrics)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUtteranceMetricsInput(v *ListUtteranceMetricsInput, encoder *httpbinding.Encoder) error {
@@ -8519,10 +8269,8 @@ func (m *awsRestjson1_serializeOpSearchAssociatedTrans
 func (m *awsRestjson1_serializeOpSearchAssociatedTranscripts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8571,7 +8319,6 @@ func (m *awsRestjson1_serializeOpSearchAssociatedTrans
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchAssociatedTranscriptsInput(v *SearchAssociatedTranscriptsInput, encoder *httpbinding.Encoder) error {
@@ -8657,10 +8404,8 @@ func (m *awsRestjson1_serializeOpStartBotRecommendatio
 func (m *awsRestjson1_serializeOpStartBotRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8709,7 +8454,6 @@ func (m *awsRestjson1_serializeOpStartBotRecommendatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartBotRecommendationInput(v *StartBotRecommendationInput, encoder *httpbinding.Encoder) error {
@@ -8778,10 +8522,8 @@ func (m *awsRestjson1_serializeOpStartBotResourceGener
 func (m *awsRestjson1_serializeOpStartBotResourceGeneration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8830,7 +8572,6 @@ func (m *awsRestjson1_serializeOpStartBotResourceGener
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartBotResourceGenerationInput(v *StartBotResourceGenerationInput, encoder *httpbinding.Encoder) error {
@@ -8890,10 +8631,8 @@ func (m *awsRestjson1_serializeOpStartImport) HandleSe
 func (m *awsRestjson1_serializeOpStartImport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8938,7 +8677,6 @@ func (m *awsRestjson1_serializeOpStartImport) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartImportInput(v *StartImportInput, encoder *httpbinding.Encoder) error {
@@ -8988,10 +8726,8 @@ func (m *awsRestjson1_serializeOpStartTestExecution) H
 func (m *awsRestjson1_serializeOpStartTestExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9040,7 +8776,6 @@ func (m *awsRestjson1_serializeOpStartTestExecution) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTestExecutionInput(v *StartTestExecutionInput, encoder *httpbinding.Encoder) error {
@@ -9094,10 +8829,8 @@ func (m *awsRestjson1_serializeOpStartTestSetGeneratio
 func (m *awsRestjson1_serializeOpStartTestSetGeneration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9142,7 +8875,6 @@ func (m *awsRestjson1_serializeOpStartTestSetGeneratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTestSetGenerationInput(v *StartTestSetGenerationInput, encoder *httpbinding.Encoder) error {
@@ -9206,10 +8938,8 @@ func (m *awsRestjson1_serializeOpStopBotRecommendation
 func (m *awsRestjson1_serializeOpStopBotRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9247,7 +8977,6 @@ func (m *awsRestjson1_serializeOpStopBotRecommendation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopBotRecommendationInput(v *StopBotRecommendationInput, encoder *httpbinding.Encoder) error {
@@ -9304,10 +9033,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -9356,7 +9083,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -9400,10 +9126,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -9441,7 +9165,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -9477,10 +9200,8 @@ func (m *awsRestjson1_serializeOpUpdateBot) HandleSeri
 func (m *awsRestjson1_serializeOpUpdateBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9529,7 +9250,6 @@ func (m *awsRestjson1_serializeOpUpdateBot) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBotInput(v *UpdateBotInput, encoder *httpbinding.Encoder) error {
@@ -9612,10 +9332,8 @@ func (m *awsRestjson1_serializeOpUpdateBotAlias) Handl
 func (m *awsRestjson1_serializeOpUpdateBotAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9664,7 +9382,6 @@ func (m *awsRestjson1_serializeOpUpdateBotAlias) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBotAliasInput(v *UpdateBotAliasInput, encoder *httpbinding.Encoder) error {
@@ -9746,10 +9463,8 @@ func (m *awsRestjson1_serializeOpUpdateBotLocale) Hand
 func (m *awsRestjson1_serializeOpUpdateBotLocale) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9798,7 +9513,6 @@ func (m *awsRestjson1_serializeOpUpdateBotLocale) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBotLocaleInput(v *UpdateBotLocaleInput, encoder *httpbinding.Encoder) error {
@@ -9890,10 +9604,8 @@ func (m *awsRestjson1_serializeOpUpdateBotRecommendati
 func (m *awsRestjson1_serializeOpUpdateBotRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9942,7 +9654,6 @@ func (m *awsRestjson1_serializeOpUpdateBotRecommendati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBotRecommendationInput(v *UpdateBotRecommendationInput, encoder *httpbinding.Encoder) error {
@@ -10013,10 +9724,8 @@ func (m *awsRestjson1_serializeOpUpdateExport) HandleS
 func (m *awsRestjson1_serializeOpUpdateExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10065,7 +9774,6 @@ func (m *awsRestjson1_serializeOpUpdateExport) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateExportInput(v *UpdateExportInput, encoder *httpbinding.Encoder) error {
@@ -10107,10 +9815,8 @@ func (m *awsRestjson1_serializeOpUpdateIntent) HandleS
 func (m *awsRestjson1_serializeOpUpdateIntent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10159,7 +9865,6 @@ func (m *awsRestjson1_serializeOpUpdateIntent) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIntentInput(v *UpdateIntentInput, encoder *httpbinding.Encoder) error {
@@ -10322,10 +10027,8 @@ func (m *awsRestjson1_serializeOpUpdateResourcePolicy)
 func (m *awsRestjson1_serializeOpUpdateResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10374,7 +10077,6 @@ func (m *awsRestjson1_serializeOpUpdateResourcePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourcePolicyInput(v *UpdateResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -10420,10 +10122,8 @@ func (m *awsRestjson1_serializeOpUpdateSlot) HandleSer
 func (m *awsRestjson1_serializeOpUpdateSlot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10472,7 +10172,6 @@ func (m *awsRestjson1_serializeOpUpdateSlot) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSlotInput(v *UpdateSlotInput, encoder *httpbinding.Encoder) error {
@@ -10588,10 +10287,8 @@ func (m *awsRestjson1_serializeOpUpdateSlotType) Handl
 func (m *awsRestjson1_serializeOpUpdateSlotType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10640,7 +10337,6 @@ func (m *awsRestjson1_serializeOpUpdateSlotType) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSlotTypeInput(v *UpdateSlotTypeInput, encoder *httpbinding.Encoder) error {
@@ -10747,10 +10443,8 @@ func (m *awsRestjson1_serializeOpUpdateTestSet) Handle
 func (m *awsRestjson1_serializeOpUpdateTestSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10799,7 +10493,6 @@ func (m *awsRestjson1_serializeOpUpdateTestSet) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTestSetInput(v *UpdateTestSetInput, encoder *httpbinding.Encoder) error {
