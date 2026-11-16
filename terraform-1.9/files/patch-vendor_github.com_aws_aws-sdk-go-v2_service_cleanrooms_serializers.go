--- vendor/github.com/aws/aws-sdk-go-v2/service/cleanrooms/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cleanrooms/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpBatchGetCollaboration
 func (m *awsRestjson1_serializeOpBatchGetCollaborationAnalysisTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpBatchGetCollaboration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetCollaborationAnalysisTemplateInput(v *BatchGetCollaborationAnalysisTemplateInput, encoder *httpbinding.Encoder) error {
@@ -121,10 +117,8 @@ func (m *awsRestjson1_serializeOpBatchGetSchema) Handl
 func (m *awsRestjson1_serializeOpBatchGetSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -173,7 +167,6 @@ func (m *awsRestjson1_serializeOpBatchGetSchema) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetSchemaInput(v *BatchGetSchemaInput, encoder *httpbinding.Encoder) error {
@@ -217,10 +210,8 @@ func (m *awsRestjson1_serializeOpBatchGetSchemaAnalysi
 func (m *awsRestjson1_serializeOpBatchGetSchemaAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -269,7 +260,6 @@ func (m *awsRestjson1_serializeOpBatchGetSchemaAnalysi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetSchemaAnalysisRuleInput(v *BatchGetSchemaAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -313,10 +303,8 @@ func (m *awsRestjson1_serializeOpCreateAnalysisTemplat
 func (m *awsRestjson1_serializeOpCreateAnalysisTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -365,7 +353,6 @@ func (m *awsRestjson1_serializeOpCreateAnalysisTemplat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAnalysisTemplateInput(v *CreateAnalysisTemplateInput, encoder *httpbinding.Encoder) error {
@@ -445,10 +432,8 @@ func (m *awsRestjson1_serializeOpCreateCollaboration) 
 func (m *awsRestjson1_serializeOpCreateCollaboration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -493,7 +478,6 @@ func (m *awsRestjson1_serializeOpCreateCollaboration) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCollaborationInput(v *CreateCollaborationInput, encoder *httpbinding.Encoder) error {
@@ -593,10 +577,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredAudie
 func (m *awsRestjson1_serializeOpCreateConfiguredAudienceModelAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -645,7 +627,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredAudie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredAudienceModelAssociationInput(v *CreateConfiguredAudienceModelAssociationInput, encoder *httpbinding.Encoder) error {
@@ -709,10 +690,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 func (m *awsRestjson1_serializeOpCreateConfiguredTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -757,7 +736,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredTableInput(v *CreateConfiguredTableInput, encoder *httpbinding.Encoder) error {
@@ -828,10 +806,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 func (m *awsRestjson1_serializeOpCreateConfiguredTableAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -880,7 +856,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredTableAnalysisRuleInput(v *CreateConfiguredTableAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -929,10 +904,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 func (m *awsRestjson1_serializeOpCreateConfiguredTableAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -981,7 +954,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredTableAssociationInput(v *CreateConfiguredTableAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1045,10 +1017,8 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 func (m *awsRestjson1_serializeOpCreateConfiguredTableAssociationAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1097,7 +1067,6 @@ func (m *awsRestjson1_serializeOpCreateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfiguredTableAssociationAnalysisRuleInput(v *CreateConfiguredTableAssociationAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -1155,10 +1124,8 @@ func (m *awsRestjson1_serializeOpCreateIdMappingTable)
 func (m *awsRestjson1_serializeOpCreateIdMappingTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1207,7 +1174,6 @@ func (m *awsRestjson1_serializeOpCreateIdMappingTable)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIdMappingTableInput(v *CreateIdMappingTableInput, encoder *httpbinding.Encoder) error {
@@ -1273,10 +1239,8 @@ func (m *awsRestjson1_serializeOpCreateIdNamespaceAsso
 func (m *awsRestjson1_serializeOpCreateIdNamespaceAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1325,7 +1289,6 @@ func (m *awsRestjson1_serializeOpCreateIdNamespaceAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIdNamespaceAssociationInput(v *CreateIdNamespaceAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1393,10 +1356,8 @@ func (m *awsRestjson1_serializeOpCreateMembership) Han
 func (m *awsRestjson1_serializeOpCreateMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1441,7 +1402,6 @@ func (m *awsRestjson1_serializeOpCreateMembership) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMembershipInput(v *CreateMembershipInput, encoder *httpbinding.Encoder) error {
@@ -1512,10 +1472,8 @@ func (m *awsRestjson1_serializeOpCreatePrivacyBudgetTe
 func (m *awsRestjson1_serializeOpCreatePrivacyBudgetTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1564,7 +1522,6 @@ func (m *awsRestjson1_serializeOpCreatePrivacyBudgetTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePrivacyBudgetTemplateInput(v *CreatePrivacyBudgetTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1625,10 +1582,8 @@ func (m *awsRestjson1_serializeOpDeleteAnalysisTemplat
 func (m *awsRestjson1_serializeOpDeleteAnalysisTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1666,7 +1621,6 @@ func (m *awsRestjson1_serializeOpDeleteAnalysisTemplat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAnalysisTemplateInput(v *DeleteAnalysisTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1705,10 +1659,8 @@ func (m *awsRestjson1_serializeOpDeleteCollaboration) 
 func (m *awsRestjson1_serializeOpDeleteCollaboration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1746,7 +1698,6 @@ func (m *awsRestjson1_serializeOpDeleteCollaboration) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCollaborationInput(v *DeleteCollaborationInput, encoder *httpbinding.Encoder) error {
@@ -1776,10 +1727,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredAudie
 func (m *awsRestjson1_serializeOpDeleteConfiguredAudienceModelAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1817,7 +1766,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredAudie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredAudienceModelAssociationInput(v *DeleteConfiguredAudienceModelAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1856,10 +1804,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 func (m *awsRestjson1_serializeOpDeleteConfiguredTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1897,7 +1843,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredTableInput(v *DeleteConfiguredTableInput, encoder *httpbinding.Encoder) error {
@@ -1927,10 +1872,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 func (m *awsRestjson1_serializeOpDeleteConfiguredTableAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1968,7 +1911,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredTableAnalysisRuleInput(v *DeleteConfiguredTableAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -2007,10 +1949,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 func (m *awsRestjson1_serializeOpDeleteConfiguredTableAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2048,7 +1988,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredTableAssociationInput(v *DeleteConfiguredTableAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2087,10 +2026,8 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 func (m *awsRestjson1_serializeOpDeleteConfiguredTableAssociationAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2128,7 +2065,6 @@ func (m *awsRestjson1_serializeOpDeleteConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfiguredTableAssociationAnalysisRuleInput(v *DeleteConfiguredTableAssociationAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -2176,10 +2112,8 @@ func (m *awsRestjson1_serializeOpDeleteIdMappingTable)
 func (m *awsRestjson1_serializeOpDeleteIdMappingTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2217,7 +2151,6 @@ func (m *awsRestjson1_serializeOpDeleteIdMappingTable)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIdMappingTableInput(v *DeleteIdMappingTableInput, encoder *httpbinding.Encoder) error {
@@ -2256,10 +2189,8 @@ func (m *awsRestjson1_serializeOpDeleteIdNamespaceAsso
 func (m *awsRestjson1_serializeOpDeleteIdNamespaceAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2297,7 +2228,6 @@ func (m *awsRestjson1_serializeOpDeleteIdNamespaceAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIdNamespaceAssociationInput(v *DeleteIdNamespaceAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2336,10 +2266,8 @@ func (m *awsRestjson1_serializeOpDeleteMember) HandleS
 func (m *awsRestjson1_serializeOpDeleteMember) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2377,7 +2305,6 @@ func (m *awsRestjson1_serializeOpDeleteMember) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMemberInput(v *DeleteMemberInput, encoder *httpbinding.Encoder) error {
@@ -2416,10 +2343,8 @@ func (m *awsRestjson1_serializeOpDeleteMembership) Han
 func (m *awsRestjson1_serializeOpDeleteMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2457,7 +2382,6 @@ func (m *awsRestjson1_serializeOpDeleteMembership) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMembershipInput(v *DeleteMembershipInput, encoder *httpbinding.Encoder) error {
@@ -2487,10 +2411,8 @@ func (m *awsRestjson1_serializeOpDeletePrivacyBudgetTe
 func (m *awsRestjson1_serializeOpDeletePrivacyBudgetTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2528,7 +2450,6 @@ func (m *awsRestjson1_serializeOpDeletePrivacyBudgetTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePrivacyBudgetTemplateInput(v *DeletePrivacyBudgetTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2567,10 +2488,8 @@ func (m *awsRestjson1_serializeOpGetAnalysisTemplate) 
 func (m *awsRestjson1_serializeOpGetAnalysisTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2608,7 +2527,6 @@ func (m *awsRestjson1_serializeOpGetAnalysisTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAnalysisTemplateInput(v *GetAnalysisTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2647,10 +2565,8 @@ func (m *awsRestjson1_serializeOpGetCollaboration) Han
 func (m *awsRestjson1_serializeOpGetCollaboration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2688,7 +2604,6 @@ func (m *awsRestjson1_serializeOpGetCollaboration) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationInput(v *GetCollaborationInput, encoder *httpbinding.Encoder) error {
@@ -2718,10 +2633,8 @@ func (m *awsRestjson1_serializeOpGetCollaborationAnaly
 func (m *awsRestjson1_serializeOpGetCollaborationAnalysisTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2759,7 +2672,6 @@ func (m *awsRestjson1_serializeOpGetCollaborationAnaly
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationAnalysisTemplateInput(v *GetCollaborationAnalysisTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2798,10 +2710,8 @@ func (m *awsRestjson1_serializeOpGetCollaborationConfi
 func (m *awsRestjson1_serializeOpGetCollaborationConfiguredAudienceModelAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2839,7 +2749,6 @@ func (m *awsRestjson1_serializeOpGetCollaborationConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationConfiguredAudienceModelAssociationInput(v *GetCollaborationConfiguredAudienceModelAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2878,10 +2787,8 @@ func (m *awsRestjson1_serializeOpGetCollaborationIdNam
 func (m *awsRestjson1_serializeOpGetCollaborationIdNamespaceAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2919,7 +2826,6 @@ func (m *awsRestjson1_serializeOpGetCollaborationIdNam
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationIdNamespaceAssociationInput(v *GetCollaborationIdNamespaceAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2958,10 +2864,8 @@ func (m *awsRestjson1_serializeOpGetCollaborationPriva
 func (m *awsRestjson1_serializeOpGetCollaborationPrivacyBudgetTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2999,7 +2903,6 @@ func (m *awsRestjson1_serializeOpGetCollaborationPriva
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCollaborationPrivacyBudgetTemplateInput(v *GetCollaborationPrivacyBudgetTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3038,10 +2941,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredAudience
 func (m *awsRestjson1_serializeOpGetConfiguredAudienceModelAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3079,7 +2980,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredAudience
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredAudienceModelAssociationInput(v *GetConfiguredAudienceModelAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3118,10 +3018,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredTable) H
 func (m *awsRestjson1_serializeOpGetConfiguredTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3159,7 +3057,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredTable) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredTableInput(v *GetConfiguredTableInput, encoder *httpbinding.Encoder) error {
@@ -3189,10 +3086,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredTableAna
 func (m *awsRestjson1_serializeOpGetConfiguredTableAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3230,7 +3125,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredTableAna
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredTableAnalysisRuleInput(v *GetConfiguredTableAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -3269,10 +3163,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredTableAss
 func (m *awsRestjson1_serializeOpGetConfiguredTableAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3310,7 +3202,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredTableAss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredTableAssociationInput(v *GetConfiguredTableAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3349,10 +3240,8 @@ func (m *awsRestjson1_serializeOpGetConfiguredTableAss
 func (m *awsRestjson1_serializeOpGetConfiguredTableAssociationAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3390,7 +3279,6 @@ func (m *awsRestjson1_serializeOpGetConfiguredTableAss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfiguredTableAssociationAnalysisRuleInput(v *GetConfiguredTableAssociationAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -3438,10 +3326,8 @@ func (m *awsRestjson1_serializeOpGetIdMappingTable) Ha
 func (m *awsRestjson1_serializeOpGetIdMappingTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3479,7 +3365,6 @@ func (m *awsRestjson1_serializeOpGetIdMappingTable) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIdMappingTableInput(v *GetIdMappingTableInput, encoder *httpbinding.Encoder) error {
@@ -3518,10 +3403,8 @@ func (m *awsRestjson1_serializeOpGetIdNamespaceAssocia
 func (m *awsRestjson1_serializeOpGetIdNamespaceAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3559,7 +3442,6 @@ func (m *awsRestjson1_serializeOpGetIdNamespaceAssocia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIdNamespaceAssociationInput(v *GetIdNamespaceAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3598,10 +3480,8 @@ func (m *awsRestjson1_serializeOpGetMembership) Handle
 func (m *awsRestjson1_serializeOpGetMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3639,7 +3519,6 @@ func (m *awsRestjson1_serializeOpGetMembership) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMembershipInput(v *GetMembershipInput, encoder *httpbinding.Encoder) error {
@@ -3669,10 +3548,8 @@ func (m *awsRestjson1_serializeOpGetPrivacyBudgetTempl
 func (m *awsRestjson1_serializeOpGetPrivacyBudgetTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3710,7 +3587,6 @@ func (m *awsRestjson1_serializeOpGetPrivacyBudgetTempl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPrivacyBudgetTemplateInput(v *GetPrivacyBudgetTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3749,10 +3625,8 @@ func (m *awsRestjson1_serializeOpGetProtectedJob) Hand
 func (m *awsRestjson1_serializeOpGetProtectedJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3790,7 +3664,6 @@ func (m *awsRestjson1_serializeOpGetProtectedJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProtectedJobInput(v *GetProtectedJobInput, encoder *httpbinding.Encoder) error {
@@ -3829,10 +3702,8 @@ func (m *awsRestjson1_serializeOpGetProtectedQuery) Ha
 func (m *awsRestjson1_serializeOpGetProtectedQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3870,7 +3741,6 @@ func (m *awsRestjson1_serializeOpGetProtectedQuery) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProtectedQueryInput(v *GetProtectedQueryInput, encoder *httpbinding.Encoder) error {
@@ -3909,10 +3779,8 @@ func (m *awsRestjson1_serializeOpGetSchema) HandleSeri
 func (m *awsRestjson1_serializeOpGetSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3950,7 +3818,6 @@ func (m *awsRestjson1_serializeOpGetSchema) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSchemaInput(v *GetSchemaInput, encoder *httpbinding.Encoder) error {
@@ -3989,10 +3856,8 @@ func (m *awsRestjson1_serializeOpGetSchemaAnalysisRule
 func (m *awsRestjson1_serializeOpGetSchemaAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4030,7 +3895,6 @@ func (m *awsRestjson1_serializeOpGetSchemaAnalysisRule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSchemaAnalysisRuleInput(v *GetSchemaAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -4078,10 +3942,8 @@ func (m *awsRestjson1_serializeOpListAnalysisTemplates
 func (m *awsRestjson1_serializeOpListAnalysisTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4119,7 +3981,6 @@ func (m *awsRestjson1_serializeOpListAnalysisTemplates
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnalysisTemplatesInput(v *ListAnalysisTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -4157,10 +4018,8 @@ func (m *awsRestjson1_serializeOpListCollaborationAnal
 func (m *awsRestjson1_serializeOpListCollaborationAnalysisTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4198,7 +4057,6 @@ func (m *awsRestjson1_serializeOpListCollaborationAnal
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationAnalysisTemplatesInput(v *ListCollaborationAnalysisTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -4236,10 +4094,8 @@ func (m *awsRestjson1_serializeOpListCollaborationConf
 func (m *awsRestjson1_serializeOpListCollaborationConfiguredAudienceModelAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4277,7 +4133,6 @@ func (m *awsRestjson1_serializeOpListCollaborationConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationConfiguredAudienceModelAssociationsInput(v *ListCollaborationConfiguredAudienceModelAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4315,10 +4170,8 @@ func (m *awsRestjson1_serializeOpListCollaborationIdNa
 func (m *awsRestjson1_serializeOpListCollaborationIdNamespaceAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4356,7 +4209,6 @@ func (m *awsRestjson1_serializeOpListCollaborationIdNa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationIdNamespaceAssociationsInput(v *ListCollaborationIdNamespaceAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4394,10 +4246,8 @@ func (m *awsRestjson1_serializeOpListCollaborationPriv
 func (m *awsRestjson1_serializeOpListCollaborationPrivacyBudgets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4435,7 +4285,6 @@ func (m *awsRestjson1_serializeOpListCollaborationPriv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationPrivacyBudgetsInput(v *ListCollaborationPrivacyBudgetsInput, encoder *httpbinding.Encoder) error {
@@ -4477,10 +4326,8 @@ func (m *awsRestjson1_serializeOpListCollaborationPriv
 func (m *awsRestjson1_serializeOpListCollaborationPrivacyBudgetTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4518,7 +4365,6 @@ func (m *awsRestjson1_serializeOpListCollaborationPriv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationPrivacyBudgetTemplatesInput(v *ListCollaborationPrivacyBudgetTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -4556,10 +4402,8 @@ func (m *awsRestjson1_serializeOpListCollaborations) H
 func (m *awsRestjson1_serializeOpListCollaborations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4597,7 +4441,6 @@ func (m *awsRestjson1_serializeOpListCollaborations) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollaborationsInput(v *ListCollaborationsInput, encoder *httpbinding.Encoder) error {
@@ -4630,10 +4473,8 @@ func (m *awsRestjson1_serializeOpListConfiguredAudienc
 func (m *awsRestjson1_serializeOpListConfiguredAudienceModelAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4671,7 +4512,6 @@ func (m *awsRestjson1_serializeOpListConfiguredAudienc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfiguredAudienceModelAssociationsInput(v *ListConfiguredAudienceModelAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4709,10 +4549,8 @@ func (m *awsRestjson1_serializeOpListConfiguredTableAs
 func (m *awsRestjson1_serializeOpListConfiguredTableAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4750,7 +4588,6 @@ func (m *awsRestjson1_serializeOpListConfiguredTableAs
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfiguredTableAssociationsInput(v *ListConfiguredTableAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4788,10 +4625,8 @@ func (m *awsRestjson1_serializeOpListConfiguredTables)
 func (m *awsRestjson1_serializeOpListConfiguredTables) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4829,7 +4664,6 @@ func (m *awsRestjson1_serializeOpListConfiguredTables)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfiguredTablesInput(v *ListConfiguredTablesInput, encoder *httpbinding.Encoder) error {
@@ -4858,10 +4692,8 @@ func (m *awsRestjson1_serializeOpListIdMappingTables) 
 func (m *awsRestjson1_serializeOpListIdMappingTables) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4899,7 +4731,6 @@ func (m *awsRestjson1_serializeOpListIdMappingTables) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIdMappingTablesInput(v *ListIdMappingTablesInput, encoder *httpbinding.Encoder) error {
@@ -4937,10 +4768,8 @@ func (m *awsRestjson1_serializeOpListIdNamespaceAssoci
 func (m *awsRestjson1_serializeOpListIdNamespaceAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4978,7 +4807,6 @@ func (m *awsRestjson1_serializeOpListIdNamespaceAssoci
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIdNamespaceAssociationsInput(v *ListIdNamespaceAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -5016,10 +4844,8 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
 func (m *awsRestjson1_serializeOpListMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5057,7 +4883,6 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMembersInput(v *ListMembersInput, encoder *httpbinding.Encoder) error {
@@ -5095,10 +4920,8 @@ func (m *awsRestjson1_serializeOpListMemberships) Hand
 func (m *awsRestjson1_serializeOpListMemberships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5136,7 +4959,6 @@ func (m *awsRestjson1_serializeOpListMemberships) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMembershipsInput(v *ListMembershipsInput, encoder *httpbinding.Encoder) error {
@@ -5169,10 +4991,8 @@ func (m *awsRestjson1_serializeOpListPrivacyBudgets) H
 func (m *awsRestjson1_serializeOpListPrivacyBudgets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5210,7 +5030,6 @@ func (m *awsRestjson1_serializeOpListPrivacyBudgets) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPrivacyBudgetsInput(v *ListPrivacyBudgetsInput, encoder *httpbinding.Encoder) error {
@@ -5252,10 +5071,8 @@ func (m *awsRestjson1_serializeOpListPrivacyBudgetTemp
 func (m *awsRestjson1_serializeOpListPrivacyBudgetTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5293,7 +5110,6 @@ func (m *awsRestjson1_serializeOpListPrivacyBudgetTemp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPrivacyBudgetTemplatesInput(v *ListPrivacyBudgetTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -5331,10 +5147,8 @@ func (m *awsRestjson1_serializeOpListProtectedJobs) Ha
 func (m *awsRestjson1_serializeOpListProtectedJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5372,7 +5186,6 @@ func (m *awsRestjson1_serializeOpListProtectedJobs) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProtectedJobsInput(v *ListProtectedJobsInput, encoder *httpbinding.Encoder) error {
@@ -5414,10 +5227,8 @@ func (m *awsRestjson1_serializeOpListProtectedQueries)
 func (m *awsRestjson1_serializeOpListProtectedQueries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5455,7 +5266,6 @@ func (m *awsRestjson1_serializeOpListProtectedQueries)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProtectedQueriesInput(v *ListProtectedQueriesInput, encoder *httpbinding.Encoder) error {
@@ -5497,10 +5307,8 @@ func (m *awsRestjson1_serializeOpListSchemas) HandleSe
 func (m *awsRestjson1_serializeOpListSchemas) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5538,7 +5346,6 @@ func (m *awsRestjson1_serializeOpListSchemas) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSchemasInput(v *ListSchemasInput, encoder *httpbinding.Encoder) error {
@@ -5580,10 +5387,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -5621,7 +5426,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -5651,10 +5455,8 @@ func (m *awsRestjson1_serializeOpPopulateIdMappingTabl
 func (m *awsRestjson1_serializeOpPopulateIdMappingTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5692,7 +5494,6 @@ func (m *awsRestjson1_serializeOpPopulateIdMappingTabl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPopulateIdMappingTableInput(v *PopulateIdMappingTableInput, encoder *httpbinding.Encoder) error {
@@ -5731,10 +5532,8 @@ func (m *awsRestjson1_serializeOpPreviewPrivacyImpact)
 func (m *awsRestjson1_serializeOpPreviewPrivacyImpact) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5783,7 +5582,6 @@ func (m *awsRestjson1_serializeOpPreviewPrivacyImpact)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPreviewPrivacyImpactInput(v *PreviewPrivacyImpactInput, encoder *httpbinding.Encoder) error {
@@ -5827,10 +5625,8 @@ func (m *awsRestjson1_serializeOpStartProtectedJob) Ha
 func (m *awsRestjson1_serializeOpStartProtectedJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5879,7 +5675,6 @@ func (m *awsRestjson1_serializeOpStartProtectedJob) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartProtectedJobInput(v *StartProtectedJobInput, encoder *httpbinding.Encoder) error {
@@ -5935,10 +5730,8 @@ func (m *awsRestjson1_serializeOpStartProtectedQuery) 
 func (m *awsRestjson1_serializeOpStartProtectedQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5987,7 +5780,6 @@ func (m *awsRestjson1_serializeOpStartProtectedQuery) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartProtectedQueryInput(v *StartProtectedQueryInput, encoder *httpbinding.Encoder) error {
@@ -6050,10 +5842,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -6102,7 +5892,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6146,10 +5935,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -6187,7 +5974,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6223,10 +6009,8 @@ func (m *awsRestjson1_serializeOpUpdateAnalysisTemplat
 func (m *awsRestjson1_serializeOpUpdateAnalysisTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6275,7 +6059,6 @@ func (m *awsRestjson1_serializeOpUpdateAnalysisTemplat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAnalysisTemplateInput(v *UpdateAnalysisTemplateInput, encoder *httpbinding.Encoder) error {
@@ -6326,10 +6109,8 @@ func (m *awsRestjson1_serializeOpUpdateCollaboration) 
 func (m *awsRestjson1_serializeOpUpdateCollaboration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6378,7 +6159,6 @@ func (m *awsRestjson1_serializeOpUpdateCollaboration) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCollaborationInput(v *UpdateCollaborationInput, encoder *httpbinding.Encoder) error {
@@ -6430,10 +6210,8 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredAudie
 func (m *awsRestjson1_serializeOpUpdateConfiguredAudienceModelAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6482,7 +6260,6 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredAudie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfiguredAudienceModelAssociationInput(v *UpdateConfiguredAudienceModelAssociationInput, encoder *httpbinding.Encoder) error {
@@ -6538,10 +6315,8 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 func (m *awsRestjson1_serializeOpUpdateConfiguredTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6590,7 +6365,6 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfiguredTableInput(v *UpdateConfiguredTableInput, encoder *httpbinding.Encoder) error {
@@ -6649,10 +6423,8 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 func (m *awsRestjson1_serializeOpUpdateConfiguredTableAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6701,7 +6473,6 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfiguredTableAnalysisRuleInput(v *UpdateConfiguredTableAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -6754,10 +6525,8 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 func (m *awsRestjson1_serializeOpUpdateConfiguredTableAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6806,7 +6575,6 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfiguredTableAssociationInput(v *UpdateConfiguredTableAssociationInput, encoder *httpbinding.Encoder) error {
@@ -6862,10 +6630,8 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 func (m *awsRestjson1_serializeOpUpdateConfiguredTableAssociationAnalysisRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6914,7 +6680,6 @@ func (m *awsRestjson1_serializeOpUpdateConfiguredTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfiguredTableAssociationAnalysisRuleInput(v *UpdateConfiguredTableAssociationAnalysisRuleInput, encoder *httpbinding.Encoder) error {
@@ -6976,10 +6741,8 @@ func (m *awsRestjson1_serializeOpUpdateIdMappingTable)
 func (m *awsRestjson1_serializeOpUpdateIdMappingTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7028,7 +6791,6 @@ func (m *awsRestjson1_serializeOpUpdateIdMappingTable)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIdMappingTableInput(v *UpdateIdMappingTableInput, encoder *httpbinding.Encoder) error {
@@ -7084,10 +6846,8 @@ func (m *awsRestjson1_serializeOpUpdateIdNamespaceAsso
 func (m *awsRestjson1_serializeOpUpdateIdNamespaceAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7136,7 +6896,6 @@ func (m *awsRestjson1_serializeOpUpdateIdNamespaceAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIdNamespaceAssociationInput(v *UpdateIdNamespaceAssociationInput, encoder *httpbinding.Encoder) error {
@@ -7199,10 +6958,8 @@ func (m *awsRestjson1_serializeOpUpdateMembership) Han
 func (m *awsRestjson1_serializeOpUpdateMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7251,7 +7008,6 @@ func (m *awsRestjson1_serializeOpUpdateMembership) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMembershipInput(v *UpdateMembershipInput, encoder *httpbinding.Encoder) error {
@@ -7312,10 +7068,8 @@ func (m *awsRestjson1_serializeOpUpdatePrivacyBudgetTe
 func (m *awsRestjson1_serializeOpUpdatePrivacyBudgetTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7364,7 +7118,6 @@ func (m *awsRestjson1_serializeOpUpdatePrivacyBudgetTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePrivacyBudgetTemplateInput(v *UpdatePrivacyBudgetTemplateInput, encoder *httpbinding.Encoder) error {
@@ -7422,10 +7175,8 @@ func (m *awsRestjson1_serializeOpUpdateProtectedJob) H
 func (m *awsRestjson1_serializeOpUpdateProtectedJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7474,7 +7225,6 @@ func (m *awsRestjson1_serializeOpUpdateProtectedJob) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProtectedJobInput(v *UpdateProtectedJobInput, encoder *httpbinding.Encoder) error {
@@ -7525,10 +7275,8 @@ func (m *awsRestjson1_serializeOpUpdateProtectedQuery)
 func (m *awsRestjson1_serializeOpUpdateProtectedQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7577,7 +7325,6 @@ func (m *awsRestjson1_serializeOpUpdateProtectedQuery)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProtectedQueryInput(v *UpdateProtectedQueryInput, encoder *httpbinding.Encoder) error {
