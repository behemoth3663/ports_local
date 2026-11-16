--- vendor/github.com/aws/aws-sdk-go-v2/service/securityhub/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/securityhub/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAcceptAdministratorIn
 func (m *awsRestjson1_serializeOpAcceptAdministratorInvitation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -75,7 +72,6 @@ func (m *awsRestjson1_serializeOpAcceptAdministratorIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptAdministratorInvitationInput(v *AcceptAdministratorInvitationInput, encoder *httpbinding.Encoder) error {
@@ -113,10 +109,8 @@ func (m *awsRestjson1_serializeOpAcceptInvitation) Han
 func (m *awsRestjson1_serializeOpAcceptInvitation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -161,7 +155,6 @@ func (m *awsRestjson1_serializeOpAcceptInvitation) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptInvitationInput(v *AcceptInvitationInput, encoder *httpbinding.Encoder) error {
@@ -199,10 +192,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteAutomation
 func (m *awsRestjson1_serializeOpBatchDeleteAutomationRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -247,7 +238,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteAutomation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteAutomationRulesInput(v *BatchDeleteAutomationRulesInput, encoder *httpbinding.Encoder) error {
@@ -282,10 +272,8 @@ func (m *awsRestjson1_serializeOpBatchDisableStandards
 func (m *awsRestjson1_serializeOpBatchDisableStandards) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -330,7 +318,6 @@ func (m *awsRestjson1_serializeOpBatchDisableStandards
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDisableStandardsInput(v *BatchDisableStandardsInput, encoder *httpbinding.Encoder) error {
@@ -365,10 +352,8 @@ func (m *awsRestjson1_serializeOpBatchEnableStandards)
 func (m *awsRestjson1_serializeOpBatchEnableStandards) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -413,7 +398,6 @@ func (m *awsRestjson1_serializeOpBatchEnableStandards)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchEnableStandardsInput(v *BatchEnableStandardsInput, encoder *httpbinding.Encoder) error {
@@ -448,10 +432,8 @@ func (m *awsRestjson1_serializeOpBatchGetAutomationRul
 func (m *awsRestjson1_serializeOpBatchGetAutomationRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -496,7 +478,6 @@ func (m *awsRestjson1_serializeOpBatchGetAutomationRul
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetAutomationRulesInput(v *BatchGetAutomationRulesInput, encoder *httpbinding.Encoder) error {
@@ -531,10 +512,8 @@ func (m *awsRestjson1_serializeOpBatchGetConfiguration
 func (m *awsRestjson1_serializeOpBatchGetConfigurationPolicyAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -579,7 +558,6 @@ func (m *awsRestjson1_serializeOpBatchGetConfiguration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetConfigurationPolicyAssociationsInput(v *BatchGetConfigurationPolicyAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -614,10 +592,8 @@ func (m *awsRestjson1_serializeOpBatchGetSecurityContr
 func (m *awsRestjson1_serializeOpBatchGetSecurityControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -662,7 +638,6 @@ func (m *awsRestjson1_serializeOpBatchGetSecurityContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetSecurityControlsInput(v *BatchGetSecurityControlsInput, encoder *httpbinding.Encoder) error {
@@ -697,10 +672,8 @@ func (m *awsRestjson1_serializeOpBatchGetStandardsCont
 func (m *awsRestjson1_serializeOpBatchGetStandardsControlAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -745,7 +718,6 @@ func (m *awsRestjson1_serializeOpBatchGetStandardsCont
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetStandardsControlAssociationsInput(v *BatchGetStandardsControlAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -780,10 +752,8 @@ func (m *awsRestjson1_serializeOpBatchImportFindings) 
 func (m *awsRestjson1_serializeOpBatchImportFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -828,7 +798,6 @@ func (m *awsRestjson1_serializeOpBatchImportFindings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchImportFindingsInput(v *BatchImportFindingsInput, encoder *httpbinding.Encoder) error {
@@ -863,10 +832,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateAutomation
 func (m *awsRestjson1_serializeOpBatchUpdateAutomationRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -911,7 +878,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateAutomation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateAutomationRulesInput(v *BatchUpdateAutomationRulesInput, encoder *httpbinding.Encoder) error {
@@ -946,10 +912,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateFindings) 
 func (m *awsRestjson1_serializeOpBatchUpdateFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -994,7 +958,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateFindings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateFindingsInput(v *BatchUpdateFindingsInput, encoder *httpbinding.Encoder) error {
@@ -1086,10 +1049,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateFindingsV2
 func (m *awsRestjson1_serializeOpBatchUpdateFindingsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1134,7 +1095,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateFindingsV2
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateFindingsV2Input(v *BatchUpdateFindingsV2Input, encoder *httpbinding.Encoder) error {
@@ -1191,10 +1151,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateStandardsC
 func (m *awsRestjson1_serializeOpBatchUpdateStandardsControlAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1239,7 +1197,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateStandardsC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateStandardsControlAssociationsInput(v *BatchUpdateStandardsControlAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -1274,10 +1231,8 @@ func (m *awsRestjson1_serializeOpConnectorRegistration
 func (m *awsRestjson1_serializeOpConnectorRegistrationsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1322,7 +1277,6 @@ func (m *awsRestjson1_serializeOpConnectorRegistration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsConnectorRegistrationsV2Input(v *ConnectorRegistrationsV2Input, encoder *httpbinding.Encoder) error {
@@ -1360,10 +1314,8 @@ func (m *awsRestjson1_serializeOpCreateActionTarget) H
 func (m *awsRestjson1_serializeOpCreateActionTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1408,7 +1360,6 @@ func (m *awsRestjson1_serializeOpCreateActionTarget) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateActionTargetInput(v *CreateActionTargetInput, encoder *httpbinding.Encoder) error {
@@ -1451,10 +1402,8 @@ func (m *awsRestjson1_serializeOpCreateAggregatorV2) H
 func (m *awsRestjson1_serializeOpCreateAggregatorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1499,7 +1448,6 @@ func (m *awsRestjson1_serializeOpCreateAggregatorV2) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAggregatorV2Input(v *CreateAggregatorV2Input, encoder *httpbinding.Encoder) error {
@@ -1551,10 +1499,8 @@ func (m *awsRestjson1_serializeOpCreateAutomationRule)
 func (m *awsRestjson1_serializeOpCreateAutomationRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1599,7 +1545,6 @@ func (m *awsRestjson1_serializeOpCreateAutomationRule)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAutomationRuleInput(v *CreateAutomationRuleInput, encoder *httpbinding.Encoder) error {
@@ -1673,10 +1618,8 @@ func (m *awsRestjson1_serializeOpCreateAutomationRuleV
 func (m *awsRestjson1_serializeOpCreateAutomationRuleV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1721,7 +1664,6 @@ func (m *awsRestjson1_serializeOpCreateAutomationRuleV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAutomationRuleV2Input(v *CreateAutomationRuleV2Input, encoder *httpbinding.Encoder) error {
@@ -1808,10 +1750,8 @@ func (m *awsRestjson1_serializeOpCreateConfigurationPo
 func (m *awsRestjson1_serializeOpCreateConfigurationPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1856,7 +1796,6 @@ func (m *awsRestjson1_serializeOpCreateConfigurationPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfigurationPolicyInput(v *CreateConfigurationPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1908,10 +1847,8 @@ func (m *awsRestjson1_serializeOpCreateConnectorV2) Ha
 func (m *awsRestjson1_serializeOpCreateConnectorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1956,7 +1893,6 @@ func (m *awsRestjson1_serializeOpCreateConnectorV2) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectorV2Input(v *CreateConnectorV2Input, encoder *httpbinding.Encoder) error {
@@ -2018,10 +1954,8 @@ func (m *awsRestjson1_serializeOpCreateFindingAggregat
 func (m *awsRestjson1_serializeOpCreateFindingAggregator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2066,7 +2000,6 @@ func (m *awsRestjson1_serializeOpCreateFindingAggregat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFindingAggregatorInput(v *CreateFindingAggregatorInput, encoder *httpbinding.Encoder) error {
@@ -2106,10 +2039,8 @@ func (m *awsRestjson1_serializeOpCreateInsight) Handle
 func (m *awsRestjson1_serializeOpCreateInsight) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2154,7 +2085,6 @@ func (m *awsRestjson1_serializeOpCreateInsight) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInsightInput(v *CreateInsightInput, encoder *httpbinding.Encoder) error {
@@ -2199,10 +2129,8 @@ func (m *awsRestjson1_serializeOpCreateMembers) Handle
 func (m *awsRestjson1_serializeOpCreateMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2247,7 +2175,6 @@ func (m *awsRestjson1_serializeOpCreateMembers) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMembersInput(v *CreateMembersInput, encoder *httpbinding.Encoder) error {
@@ -2282,10 +2209,8 @@ func (m *awsRestjson1_serializeOpCreateTicketV2) Handl
 func (m *awsRestjson1_serializeOpCreateTicketV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2330,7 +2255,6 @@ func (m *awsRestjson1_serializeOpCreateTicketV2) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTicketV2Input(v *CreateTicketV2Input, encoder *httpbinding.Encoder) error {
@@ -2373,10 +2297,8 @@ func (m *awsRestjson1_serializeOpDeclineInvitations) H
 func (m *awsRestjson1_serializeOpDeclineInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2421,7 +2343,6 @@ func (m *awsRestjson1_serializeOpDeclineInvitations) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeclineInvitationsInput(v *DeclineInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -2456,10 +2377,8 @@ func (m *awsRestjson1_serializeOpDeleteActionTarget) H
 func (m *awsRestjson1_serializeOpDeleteActionTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2497,7 +2416,6 @@ func (m *awsRestjson1_serializeOpDeleteActionTarget) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteActionTargetInput(v *DeleteActionTargetInput, encoder *httpbinding.Encoder) error {
@@ -2527,10 +2445,8 @@ func (m *awsRestjson1_serializeOpDeleteAggregatorV2) H
 func (m *awsRestjson1_serializeOpDeleteAggregatorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2568,7 +2484,6 @@ func (m *awsRestjson1_serializeOpDeleteAggregatorV2) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAggregatorV2Input(v *DeleteAggregatorV2Input, encoder *httpbinding.Encoder) error {
@@ -2598,10 +2513,8 @@ func (m *awsRestjson1_serializeOpDeleteAutomationRuleV
 func (m *awsRestjson1_serializeOpDeleteAutomationRuleV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2639,7 +2552,6 @@ func (m *awsRestjson1_serializeOpDeleteAutomationRuleV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAutomationRuleV2Input(v *DeleteAutomationRuleV2Input, encoder *httpbinding.Encoder) error {
@@ -2669,10 +2581,8 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationPo
 func (m *awsRestjson1_serializeOpDeleteConfigurationPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2710,7 +2620,6 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfigurationPolicyInput(v *DeleteConfigurationPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2740,10 +2649,8 @@ func (m *awsRestjson1_serializeOpDeleteConnectorV2) Ha
 func (m *awsRestjson1_serializeOpDeleteConnectorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2781,7 +2688,6 @@ func (m *awsRestjson1_serializeOpDeleteConnectorV2) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConnectorV2Input(v *DeleteConnectorV2Input, encoder *httpbinding.Encoder) error {
@@ -2811,10 +2717,8 @@ func (m *awsRestjson1_serializeOpDeleteFindingAggregat
 func (m *awsRestjson1_serializeOpDeleteFindingAggregator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2852,7 +2756,6 @@ func (m *awsRestjson1_serializeOpDeleteFindingAggregat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFindingAggregatorInput(v *DeleteFindingAggregatorInput, encoder *httpbinding.Encoder) error {
@@ -2882,10 +2785,8 @@ func (m *awsRestjson1_serializeOpDeleteInsight) Handle
 func (m *awsRestjson1_serializeOpDeleteInsight) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2923,7 +2824,6 @@ func (m *awsRestjson1_serializeOpDeleteInsight) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInsightInput(v *DeleteInsightInput, encoder *httpbinding.Encoder) error {
@@ -2953,10 +2853,8 @@ func (m *awsRestjson1_serializeOpDeleteInvitations) Ha
 func (m *awsRestjson1_serializeOpDeleteInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3001,7 +2899,6 @@ func (m *awsRestjson1_serializeOpDeleteInvitations) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInvitationsInput(v *DeleteInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -3036,10 +2933,8 @@ func (m *awsRestjson1_serializeOpDeleteMembers) Handle
 func (m *awsRestjson1_serializeOpDeleteMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3084,7 +2979,6 @@ func (m *awsRestjson1_serializeOpDeleteMembers) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMembersInput(v *DeleteMembersInput, encoder *httpbinding.Encoder) error {
@@ -3119,10 +3013,8 @@ func (m *awsRestjson1_serializeOpDescribeActionTargets
 func (m *awsRestjson1_serializeOpDescribeActionTargets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3167,7 +3059,6 @@ func (m *awsRestjson1_serializeOpDescribeActionTargets
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeActionTargetsInput(v *DescribeActionTargetsInput, encoder *httpbinding.Encoder) error {
@@ -3212,10 +3103,8 @@ func (m *awsRestjson1_serializeOpDescribeHub) HandleSe
 func (m *awsRestjson1_serializeOpDescribeHub) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3253,7 +3142,6 @@ func (m *awsRestjson1_serializeOpDescribeHub) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeHubInput(v *DescribeHubInput, encoder *httpbinding.Encoder) error {
@@ -3278,10 +3166,8 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
 func (m *awsRestjson1_serializeOpDescribeOrganizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3315,7 +3201,6 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOrganizationConfigurationInput(v *DescribeOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3336,10 +3221,8 @@ func (m *awsRestjson1_serializeOpDescribeProducts) Han
 func (m *awsRestjson1_serializeOpDescribeProducts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3377,7 +3260,6 @@ func (m *awsRestjson1_serializeOpDescribeProducts) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeProductsInput(v *DescribeProductsInput, encoder *httpbinding.Encoder) error {
@@ -3410,10 +3292,8 @@ func (m *awsRestjson1_serializeOpDescribeProductsV2) H
 func (m *awsRestjson1_serializeOpDescribeProductsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3451,7 +3331,6 @@ func (m *awsRestjson1_serializeOpDescribeProductsV2) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeProductsV2Input(v *DescribeProductsV2Input, encoder *httpbinding.Encoder) error {
@@ -3480,10 +3359,8 @@ func (m *awsRestjson1_serializeOpDescribeSecurityHubV2
 func (m *awsRestjson1_serializeOpDescribeSecurityHubV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3517,7 +3394,6 @@ func (m *awsRestjson1_serializeOpDescribeSecurityHubV2
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSecurityHubV2Input(v *DescribeSecurityHubV2Input, encoder *httpbinding.Encoder) error {
@@ -3538,10 +3414,8 @@ func (m *awsRestjson1_serializeOpDescribeStandards) Ha
 func (m *awsRestjson1_serializeOpDescribeStandards) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3579,7 +3453,6 @@ func (m *awsRestjson1_serializeOpDescribeStandards) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeStandardsInput(v *DescribeStandardsInput, encoder *httpbinding.Encoder) error {
@@ -3608,10 +3481,8 @@ func (m *awsRestjson1_serializeOpDescribeStandardsCont
 func (m *awsRestjson1_serializeOpDescribeStandardsControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3649,7 +3520,6 @@ func (m *awsRestjson1_serializeOpDescribeStandardsCont
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeStandardsControlsInput(v *DescribeStandardsControlsInput, encoder *httpbinding.Encoder) error {
@@ -3687,10 +3557,8 @@ func (m *awsRestjson1_serializeOpDisableImportFindings
 func (m *awsRestjson1_serializeOpDisableImportFindingsForProduct) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3728,7 +3596,6 @@ func (m *awsRestjson1_serializeOpDisableImportFindings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableImportFindingsForProductInput(v *DisableImportFindingsForProductInput, encoder *httpbinding.Encoder) error {
@@ -3758,10 +3625,8 @@ func (m *awsRestjson1_serializeOpDisableOrganizationAd
 func (m *awsRestjson1_serializeOpDisableOrganizationAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3806,7 +3671,6 @@ func (m *awsRestjson1_serializeOpDisableOrganizationAd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableOrganizationAdminAccountInput(v *DisableOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -3844,10 +3708,8 @@ func (m *awsRestjson1_serializeOpDisableSecurityHub) H
 func (m *awsRestjson1_serializeOpDisableSecurityHub) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3881,7 +3743,6 @@ func (m *awsRestjson1_serializeOpDisableSecurityHub) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableSecurityHubInput(v *DisableSecurityHubInput, encoder *httpbinding.Encoder) error {
@@ -3902,10 +3763,8 @@ func (m *awsRestjson1_serializeOpDisableSecurityHubV2)
 func (m *awsRestjson1_serializeOpDisableSecurityHubV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3939,7 +3798,6 @@ func (m *awsRestjson1_serializeOpDisableSecurityHubV2)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableSecurityHubV2Input(v *DisableSecurityHubV2Input, encoder *httpbinding.Encoder) error {
@@ -3960,10 +3818,8 @@ func (m *awsRestjson1_serializeOpDisassociateFromAdmin
 func (m *awsRestjson1_serializeOpDisassociateFromAdministratorAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3997,7 +3853,6 @@ func (m *awsRestjson1_serializeOpDisassociateFromAdmin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFromAdministratorAccountInput(v *DisassociateFromAdministratorAccountInput, encoder *httpbinding.Encoder) error {
@@ -4018,10 +3873,8 @@ func (m *awsRestjson1_serializeOpDisassociateFromMaste
 func (m *awsRestjson1_serializeOpDisassociateFromMasterAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4055,7 +3908,6 @@ func (m *awsRestjson1_serializeOpDisassociateFromMaste
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFromMasterAccountInput(v *DisassociateFromMasterAccountInput, encoder *httpbinding.Encoder) error {
@@ -4076,10 +3928,8 @@ func (m *awsRestjson1_serializeOpDisassociateMembers) 
 func (m *awsRestjson1_serializeOpDisassociateMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4124,7 +3974,6 @@ func (m *awsRestjson1_serializeOpDisassociateMembers) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMembersInput(v *DisassociateMembersInput, encoder *httpbinding.Encoder) error {
@@ -4159,10 +4008,8 @@ func (m *awsRestjson1_serializeOpEnableImportFindingsF
 func (m *awsRestjson1_serializeOpEnableImportFindingsForProduct) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4207,7 +4054,6 @@ func (m *awsRestjson1_serializeOpEnableImportFindingsF
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableImportFindingsForProductInput(v *EnableImportFindingsForProductInput, encoder *httpbinding.Encoder) error {
@@ -4240,10 +4086,8 @@ func (m *awsRestjson1_serializeOpEnableOrganizationAdm
 func (m *awsRestjson1_serializeOpEnableOrganizationAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4288,7 +4132,6 @@ func (m *awsRestjson1_serializeOpEnableOrganizationAdm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableOrganizationAdminAccountInput(v *EnableOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -4326,10 +4169,8 @@ func (m *awsRestjson1_serializeOpEnableSecurityHub) Ha
 func (m *awsRestjson1_serializeOpEnableSecurityHub) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4374,7 +4215,6 @@ func (m *awsRestjson1_serializeOpEnableSecurityHub) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableSecurityHubInput(v *EnableSecurityHubInput, encoder *httpbinding.Encoder) error {
@@ -4419,10 +4259,8 @@ func (m *awsRestjson1_serializeOpEnableSecurityHubV2) 
 func (m *awsRestjson1_serializeOpEnableSecurityHubV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4467,7 +4305,6 @@ func (m *awsRestjson1_serializeOpEnableSecurityHubV2) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableSecurityHubV2Input(v *EnableSecurityHubV2Input, encoder *httpbinding.Encoder) error {
@@ -4502,10 +4339,8 @@ func (m *awsRestjson1_serializeOpGetAdministratorAccou
 func (m *awsRestjson1_serializeOpGetAdministratorAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4539,7 +4374,6 @@ func (m *awsRestjson1_serializeOpGetAdministratorAccou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAdministratorAccountInput(v *GetAdministratorAccountInput, encoder *httpbinding.Encoder) error {
@@ -4560,10 +4394,8 @@ func (m *awsRestjson1_serializeOpGetAggregatorV2) Hand
 func (m *awsRestjson1_serializeOpGetAggregatorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4601,7 +4433,6 @@ func (m *awsRestjson1_serializeOpGetAggregatorV2) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAggregatorV2Input(v *GetAggregatorV2Input, encoder *httpbinding.Encoder) error {
@@ -4631,10 +4462,8 @@ func (m *awsRestjson1_serializeOpGetAutomationRuleV2) 
 func (m *awsRestjson1_serializeOpGetAutomationRuleV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4672,7 +4501,6 @@ func (m *awsRestjson1_serializeOpGetAutomationRuleV2) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAutomationRuleV2Input(v *GetAutomationRuleV2Input, encoder *httpbinding.Encoder) error {
@@ -4702,10 +4530,8 @@ func (m *awsRestjson1_serializeOpGetConfigurationPolic
 func (m *awsRestjson1_serializeOpGetConfigurationPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4743,7 +4569,6 @@ func (m *awsRestjson1_serializeOpGetConfigurationPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfigurationPolicyInput(v *GetConfigurationPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4773,10 +4598,8 @@ func (m *awsRestjson1_serializeOpGetConfigurationPolic
 func (m *awsRestjson1_serializeOpGetConfigurationPolicyAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4821,7 +4644,6 @@ func (m *awsRestjson1_serializeOpGetConfigurationPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfigurationPolicyAssociationInput(v *GetConfigurationPolicyAssociationInput, encoder *httpbinding.Encoder) error {
@@ -4856,10 +4678,8 @@ func (m *awsRestjson1_serializeOpGetConnectorV2) Handl
 func (m *awsRestjson1_serializeOpGetConnectorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4897,7 +4717,6 @@ func (m *awsRestjson1_serializeOpGetConnectorV2) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectorV2Input(v *GetConnectorV2Input, encoder *httpbinding.Encoder) error {
@@ -4927,10 +4746,8 @@ func (m *awsRestjson1_serializeOpGetEnabledStandards) 
 func (m *awsRestjson1_serializeOpGetEnabledStandards) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4975,7 +4792,6 @@ func (m *awsRestjson1_serializeOpGetEnabledStandards) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEnabledStandardsInput(v *GetEnabledStandardsInput, encoder *httpbinding.Encoder) error {
@@ -5020,10 +4836,8 @@ func (m *awsRestjson1_serializeOpGetFindingAggregator)
 func (m *awsRestjson1_serializeOpGetFindingAggregator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5061,7 +4875,6 @@ func (m *awsRestjson1_serializeOpGetFindingAggregator)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingAggregatorInput(v *GetFindingAggregatorInput, encoder *httpbinding.Encoder) error {
@@ -5091,10 +4904,8 @@ func (m *awsRestjson1_serializeOpGetFindingHistory) Ha
 func (m *awsRestjson1_serializeOpGetFindingHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5139,7 +4950,6 @@ func (m *awsRestjson1_serializeOpGetFindingHistory) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingHistoryInput(v *GetFindingHistoryInput, encoder *httpbinding.Encoder) error {
@@ -5194,10 +5004,8 @@ func (m *awsRestjson1_serializeOpGetFindings) HandleSe
 func (m *awsRestjson1_serializeOpGetFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5242,7 +5050,6 @@ func (m *awsRestjson1_serializeOpGetFindings) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsInput(v *GetFindingsInput, encoder *httpbinding.Encoder) error {
@@ -5294,10 +5101,8 @@ func (m *awsRestjson1_serializeOpGetFindingStatisticsV
 func (m *awsRestjson1_serializeOpGetFindingStatisticsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5342,7 +5147,6 @@ func (m *awsRestjson1_serializeOpGetFindingStatisticsV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingStatisticsV2Input(v *GetFindingStatisticsV2Input, encoder *httpbinding.Encoder) error {
@@ -5387,10 +5191,8 @@ func (m *awsRestjson1_serializeOpGetFindingsV2) Handle
 func (m *awsRestjson1_serializeOpGetFindingsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5435,7 +5237,6 @@ func (m *awsRestjson1_serializeOpGetFindingsV2) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsV2Input(v *GetFindingsV2Input, encoder *httpbinding.Encoder) error {
@@ -5487,10 +5288,8 @@ func (m *awsRestjson1_serializeOpGetInsightResults) Ha
 func (m *awsRestjson1_serializeOpGetInsightResults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5528,7 +5327,6 @@ func (m *awsRestjson1_serializeOpGetInsightResults) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInsightResultsInput(v *GetInsightResultsInput, encoder *httpbinding.Encoder) error {
@@ -5558,10 +5356,8 @@ func (m *awsRestjson1_serializeOpGetInsights) HandleSe
 func (m *awsRestjson1_serializeOpGetInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5606,7 +5402,6 @@ func (m *awsRestjson1_serializeOpGetInsights) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInsightsInput(v *GetInsightsInput, encoder *httpbinding.Encoder) error {
@@ -5651,10 +5446,8 @@ func (m *awsRestjson1_serializeOpGetInvitationsCount) 
 func (m *awsRestjson1_serializeOpGetInvitationsCount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5688,7 +5481,6 @@ func (m *awsRestjson1_serializeOpGetInvitationsCount) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInvitationsCountInput(v *GetInvitationsCountInput, encoder *httpbinding.Encoder) error {
@@ -5709,10 +5501,8 @@ func (m *awsRestjson1_serializeOpGetMasterAccount) Han
 func (m *awsRestjson1_serializeOpGetMasterAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5746,7 +5536,6 @@ func (m *awsRestjson1_serializeOpGetMasterAccount) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMasterAccountInput(v *GetMasterAccountInput, encoder *httpbinding.Encoder) error {
@@ -5767,10 +5556,8 @@ func (m *awsRestjson1_serializeOpGetMembers) HandleSer
 func (m *awsRestjson1_serializeOpGetMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5815,7 +5602,6 @@ func (m *awsRestjson1_serializeOpGetMembers) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMembersInput(v *GetMembersInput, encoder *httpbinding.Encoder) error {
@@ -5850,10 +5636,8 @@ func (m *awsRestjson1_serializeOpGetResourcesStatistic
 func (m *awsRestjson1_serializeOpGetResourcesStatisticsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5898,7 +5682,6 @@ func (m *awsRestjson1_serializeOpGetResourcesStatistic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourcesStatisticsV2Input(v *GetResourcesStatisticsV2Input, encoder *httpbinding.Encoder) error {
@@ -5943,10 +5726,8 @@ func (m *awsRestjson1_serializeOpGetResourcesV2) Handl
 func (m *awsRestjson1_serializeOpGetResourcesV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5991,7 +5772,6 @@ func (m *awsRestjson1_serializeOpGetResourcesV2) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourcesV2Input(v *GetResourcesV2Input, encoder *httpbinding.Encoder) error {
@@ -6043,10 +5823,8 @@ func (m *awsRestjson1_serializeOpGetSecurityControlDef
 func (m *awsRestjson1_serializeOpGetSecurityControlDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6084,7 +5862,6 @@ func (m *awsRestjson1_serializeOpGetSecurityControlDef
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSecurityControlDefinitionInput(v *GetSecurityControlDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -6109,10 +5886,8 @@ func (m *awsRestjson1_serializeOpInviteMembers) Handle
 func (m *awsRestjson1_serializeOpInviteMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6157,7 +5932,6 @@ func (m *awsRestjson1_serializeOpInviteMembers) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInviteMembersInput(v *InviteMembersInput, encoder *httpbinding.Encoder) error {
@@ -6192,10 +5966,8 @@ func (m *awsRestjson1_serializeOpListAggregatorsV2) Ha
 func (m *awsRestjson1_serializeOpListAggregatorsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6233,7 +6005,6 @@ func (m *awsRestjson1_serializeOpListAggregatorsV2) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAggregatorsV2Input(v *ListAggregatorsV2Input, encoder *httpbinding.Encoder) error {
@@ -6262,10 +6033,8 @@ func (m *awsRestjson1_serializeOpListAutomationRules) 
 func (m *awsRestjson1_serializeOpListAutomationRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6303,7 +6072,6 @@ func (m *awsRestjson1_serializeOpListAutomationRules) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAutomationRulesInput(v *ListAutomationRulesInput, encoder *httpbinding.Encoder) error {
@@ -6332,10 +6100,8 @@ func (m *awsRestjson1_serializeOpListAutomationRulesV2
 func (m *awsRestjson1_serializeOpListAutomationRulesV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6373,7 +6139,6 @@ func (m *awsRestjson1_serializeOpListAutomationRulesV2
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAutomationRulesV2Input(v *ListAutomationRulesV2Input, encoder *httpbinding.Encoder) error {
@@ -6402,10 +6167,8 @@ func (m *awsRestjson1_serializeOpListConfigurationPoli
 func (m *awsRestjson1_serializeOpListConfigurationPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6443,7 +6206,6 @@ func (m *awsRestjson1_serializeOpListConfigurationPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfigurationPoliciesInput(v *ListConfigurationPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -6472,10 +6234,8 @@ func (m *awsRestjson1_serializeOpListConfigurationPoli
 func (m *awsRestjson1_serializeOpListConfigurationPolicyAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6520,7 +6280,6 @@ func (m *awsRestjson1_serializeOpListConfigurationPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfigurationPolicyAssociationsInput(v *ListConfigurationPolicyAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -6565,10 +6324,8 @@ func (m *awsRestjson1_serializeOpListConnectorsV2) Han
 func (m *awsRestjson1_serializeOpListConnectorsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6606,7 +6363,6 @@ func (m *awsRestjson1_serializeOpListConnectorsV2) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConnectorsV2Input(v *ListConnectorsV2Input, encoder *httpbinding.Encoder) error {
@@ -6643,10 +6399,8 @@ func (m *awsRestjson1_serializeOpListEnabledProductsFo
 func (m *awsRestjson1_serializeOpListEnabledProductsForImport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6684,7 +6438,6 @@ func (m *awsRestjson1_serializeOpListEnabledProductsFo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEnabledProductsForImportInput(v *ListEnabledProductsForImportInput, encoder *httpbinding.Encoder) error {
@@ -6713,10 +6466,8 @@ func (m *awsRestjson1_serializeOpListFindingAggregator
 func (m *awsRestjson1_serializeOpListFindingAggregators) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6754,7 +6505,6 @@ func (m *awsRestjson1_serializeOpListFindingAggregator
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFindingAggregatorsInput(v *ListFindingAggregatorsInput, encoder *httpbinding.Encoder) error {
@@ -6783,10 +6533,8 @@ func (m *awsRestjson1_serializeOpListInvitations) Hand
 func (m *awsRestjson1_serializeOpListInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6824,7 +6572,6 @@ func (m *awsRestjson1_serializeOpListInvitations) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInvitationsInput(v *ListInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -6853,10 +6600,8 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
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
@@ -6894,7 +6639,6 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMembersInput(v *ListMembersInput, encoder *httpbinding.Encoder) error {
@@ -6927,10 +6671,8 @@ func (m *awsRestjson1_serializeOpListOrganizationAdmin
 func (m *awsRestjson1_serializeOpListOrganizationAdminAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6968,7 +6710,6 @@ func (m *awsRestjson1_serializeOpListOrganizationAdmin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationAdminAccountsInput(v *ListOrganizationAdminAccountsInput, encoder *httpbinding.Encoder) error {
@@ -7001,10 +6742,8 @@ func (m *awsRestjson1_serializeOpListSecurityControlDe
 func (m *awsRestjson1_serializeOpListSecurityControlDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7042,7 +6781,6 @@ func (m *awsRestjson1_serializeOpListSecurityControlDe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSecurityControlDefinitionsInput(v *ListSecurityControlDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -7075,10 +6813,8 @@ func (m *awsRestjson1_serializeOpListStandardsControlA
 func (m *awsRestjson1_serializeOpListStandardsControlAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7116,7 +6852,6 @@ func (m *awsRestjson1_serializeOpListStandardsControlA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStandardsControlAssociationsInput(v *ListStandardsControlAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -7149,10 +6884,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -7190,7 +6923,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -7220,10 +6952,8 @@ func (m *awsRestjson1_serializeOpStartConfigurationPol
 func (m *awsRestjson1_serializeOpStartConfigurationPolicyAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7268,7 +6998,6 @@ func (m *awsRestjson1_serializeOpStartConfigurationPol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartConfigurationPolicyAssociationInput(v *StartConfigurationPolicyAssociationInput, encoder *httpbinding.Encoder) error {
@@ -7308,10 +7037,8 @@ func (m *awsRestjson1_serializeOpStartConfigurationPol
 func (m *awsRestjson1_serializeOpStartConfigurationPolicyDisassociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7356,7 +7083,6 @@ func (m *awsRestjson1_serializeOpStartConfigurationPol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartConfigurationPolicyDisassociationInput(v *StartConfigurationPolicyDisassociationInput, encoder *httpbinding.Encoder) error {
@@ -7396,10 +7122,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7448,7 +7172,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7492,10 +7215,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7533,7 +7254,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7569,10 +7289,8 @@ func (m *awsRestjson1_serializeOpUpdateActionTarget) H
 func (m *awsRestjson1_serializeOpUpdateActionTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7621,7 +7339,6 @@ func (m *awsRestjson1_serializeOpUpdateActionTarget) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateActionTargetInput(v *UpdateActionTargetInput, encoder *httpbinding.Encoder) error {
@@ -7668,10 +7385,8 @@ func (m *awsRestjson1_serializeOpUpdateAggregatorV2) H
 func (m *awsRestjson1_serializeOpUpdateAggregatorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7720,7 +7435,6 @@ func (m *awsRestjson1_serializeOpUpdateAggregatorV2) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAggregatorV2Input(v *UpdateAggregatorV2Input, encoder *httpbinding.Encoder) error {
@@ -7769,10 +7483,8 @@ func (m *awsRestjson1_serializeOpUpdateAutomationRuleV
 func (m *awsRestjson1_serializeOpUpdateAutomationRuleV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7821,7 +7533,6 @@ func (m *awsRestjson1_serializeOpUpdateAutomationRuleV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAutomationRuleV2Input(v *UpdateAutomationRuleV2Input, encoder *httpbinding.Encoder) error {
@@ -7905,10 +7616,8 @@ func (m *awsRestjson1_serializeOpUpdateConfigurationPo
 func (m *awsRestjson1_serializeOpUpdateConfigurationPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7957,7 +7666,6 @@ func (m *awsRestjson1_serializeOpUpdateConfigurationPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfigurationPolicyInput(v *UpdateConfigurationPolicyInput, encoder *httpbinding.Encoder) error {
@@ -8016,10 +7724,8 @@ func (m *awsRestjson1_serializeOpUpdateConnectorV2) Ha
 func (m *awsRestjson1_serializeOpUpdateConnectorV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8068,7 +7774,6 @@ func (m *awsRestjson1_serializeOpUpdateConnectorV2) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConnectorV2Input(v *UpdateConnectorV2Input, encoder *httpbinding.Encoder) error {
@@ -8122,10 +7827,8 @@ func (m *awsRestjson1_serializeOpUpdateFindingAggregat
 func (m *awsRestjson1_serializeOpUpdateFindingAggregator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8170,7 +7873,6 @@ func (m *awsRestjson1_serializeOpUpdateFindingAggregat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFindingAggregatorInput(v *UpdateFindingAggregatorInput, encoder *httpbinding.Encoder) error {
@@ -8215,10 +7917,8 @@ func (m *awsRestjson1_serializeOpUpdateFindings) Handl
 func (m *awsRestjson1_serializeOpUpdateFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8263,7 +7963,6 @@ func (m *awsRestjson1_serializeOpUpdateFindings) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFindingsInput(v *UpdateFindingsInput, encoder *httpbinding.Encoder) error {
@@ -8310,10 +8009,8 @@ func (m *awsRestjson1_serializeOpUpdateInsight) Handle
 func (m *awsRestjson1_serializeOpUpdateInsight) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8362,7 +8059,6 @@ func (m *awsRestjson1_serializeOpUpdateInsight) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInsightInput(v *UpdateInsightInput, encoder *httpbinding.Encoder) error {
@@ -8416,10 +8112,8 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
 func (m *awsRestjson1_serializeOpUpdateOrganizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8464,7 +8158,6 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOrganizationConfigurationInput(v *UpdateOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8509,10 +8202,8 @@ func (m *awsRestjson1_serializeOpUpdateSecurityControl
 func (m *awsRestjson1_serializeOpUpdateSecurityControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8557,7 +8248,6 @@ func (m *awsRestjson1_serializeOpUpdateSecurityControl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSecurityControlInput(v *UpdateSecurityControlInput, encoder *httpbinding.Encoder) error {
@@ -8602,10 +8292,8 @@ func (m *awsRestjson1_serializeOpUpdateSecurityHubConf
 func (m *awsRestjson1_serializeOpUpdateSecurityHubConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8650,7 +8338,6 @@ func (m *awsRestjson1_serializeOpUpdateSecurityHubConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSecurityHubConfigurationInput(v *UpdateSecurityHubConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8688,10 +8375,8 @@ func (m *awsRestjson1_serializeOpUpdateStandardsContro
 func (m *awsRestjson1_serializeOpUpdateStandardsControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8740,7 +8425,6 @@ func (m *awsRestjson1_serializeOpUpdateStandardsContro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateStandardsControlInput(v *UpdateStandardsControlInput, encoder *httpbinding.Encoder) error {
