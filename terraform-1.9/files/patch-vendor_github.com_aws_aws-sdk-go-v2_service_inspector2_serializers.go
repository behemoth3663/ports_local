--- vendor/github.com/aws/aws-sdk-go-v2/service/inspector2/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/inspector2/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAssociateMember) Hand
 func (m *awsRestjson1_serializeOpAssociateMember) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -75,7 +72,6 @@ func (m *awsRestjson1_serializeOpAssociateMember) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateMemberInput(v *AssociateMemberInput, encoder *httpbinding.Encoder) error {
@@ -108,10 +104,8 @@ func (m *awsRestjson1_serializeOpBatchAssociateCodeSec
 func (m *awsRestjson1_serializeOpBatchAssociateCodeSecurityScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -156,7 +150,6 @@ func (m *awsRestjson1_serializeOpBatchAssociateCodeSec
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchAssociateCodeSecurityScanConfigurationInput(v *BatchAssociateCodeSecurityScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -191,10 +184,8 @@ func (m *awsRestjson1_serializeOpBatchDisassociateCode
 func (m *awsRestjson1_serializeOpBatchDisassociateCodeSecurityScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -239,7 +230,6 @@ func (m *awsRestjson1_serializeOpBatchDisassociateCode
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDisassociateCodeSecurityScanConfigurationInput(v *BatchDisassociateCodeSecurityScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -274,10 +264,8 @@ func (m *awsRestjson1_serializeOpBatchGetAccountStatus
 func (m *awsRestjson1_serializeOpBatchGetAccountStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -322,7 +310,6 @@ func (m *awsRestjson1_serializeOpBatchGetAccountStatus
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetAccountStatusInput(v *BatchGetAccountStatusInput, encoder *httpbinding.Encoder) error {
@@ -357,10 +344,8 @@ func (m *awsRestjson1_serializeOpBatchGetCodeSnippet) 
 func (m *awsRestjson1_serializeOpBatchGetCodeSnippet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -405,7 +390,6 @@ func (m *awsRestjson1_serializeOpBatchGetCodeSnippet) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetCodeSnippetInput(v *BatchGetCodeSnippetInput, encoder *httpbinding.Encoder) error {
@@ -440,10 +424,8 @@ func (m *awsRestjson1_serializeOpBatchGetFindingDetail
 func (m *awsRestjson1_serializeOpBatchGetFindingDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -488,7 +470,6 @@ func (m *awsRestjson1_serializeOpBatchGetFindingDetail
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetFindingDetailsInput(v *BatchGetFindingDetailsInput, encoder *httpbinding.Encoder) error {
@@ -523,10 +504,8 @@ func (m *awsRestjson1_serializeOpBatchGetFreeTrialInfo
 func (m *awsRestjson1_serializeOpBatchGetFreeTrialInfo) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -571,7 +550,6 @@ func (m *awsRestjson1_serializeOpBatchGetFreeTrialInfo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetFreeTrialInfoInput(v *BatchGetFreeTrialInfoInput, encoder *httpbinding.Encoder) error {
@@ -606,10 +584,8 @@ func (m *awsRestjson1_serializeOpBatchGetMemberEc2Deep
 func (m *awsRestjson1_serializeOpBatchGetMemberEc2DeepInspectionStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -654,7 +630,6 @@ func (m *awsRestjson1_serializeOpBatchGetMemberEc2Deep
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetMemberEc2DeepInspectionStatusInput(v *BatchGetMemberEc2DeepInspectionStatusInput, encoder *httpbinding.Encoder) error {
@@ -689,10 +664,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateMemberEc2D
 func (m *awsRestjson1_serializeOpBatchUpdateMemberEc2DeepInspectionStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -737,7 +710,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateMemberEc2D
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateMemberEc2DeepInspectionStatusInput(v *BatchUpdateMemberEc2DeepInspectionStatusInput, encoder *httpbinding.Encoder) error {
@@ -772,10 +744,8 @@ func (m *awsRestjson1_serializeOpCancelFindingsReport)
 func (m *awsRestjson1_serializeOpCancelFindingsReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -820,7 +790,6 @@ func (m *awsRestjson1_serializeOpCancelFindingsReport)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelFindingsReportInput(v *CancelFindingsReportInput, encoder *httpbinding.Encoder) error {
@@ -853,10 +822,8 @@ func (m *awsRestjson1_serializeOpCancelSbomExport) Han
 func (m *awsRestjson1_serializeOpCancelSbomExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -901,7 +868,6 @@ func (m *awsRestjson1_serializeOpCancelSbomExport) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelSbomExportInput(v *CancelSbomExportInput, encoder *httpbinding.Encoder) error {
@@ -934,10 +900,8 @@ func (m *awsRestjson1_serializeOpCreateCisScanConfigur
 func (m *awsRestjson1_serializeOpCreateCisScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -982,7 +946,6 @@ func (m *awsRestjson1_serializeOpCreateCisScanConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCisScanConfigurationInput(v *CreateCisScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1041,10 +1004,8 @@ func (m *awsRestjson1_serializeOpCreateCodeSecurityInt
 func (m *awsRestjson1_serializeOpCreateCodeSecurityIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1089,7 +1050,6 @@ func (m *awsRestjson1_serializeOpCreateCodeSecurityInt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCodeSecurityIntegrationInput(v *CreateCodeSecurityIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -1141,10 +1101,8 @@ func (m *awsRestjson1_serializeOpCreateCodeSecuritySca
 func (m *awsRestjson1_serializeOpCreateCodeSecurityScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1189,7 +1147,6 @@ func (m *awsRestjson1_serializeOpCreateCodeSecuritySca
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCodeSecurityScanConfigurationInput(v *CreateCodeSecurityScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1248,10 +1205,8 @@ func (m *awsRestjson1_serializeOpCreateFilter) HandleS
 func (m *awsRestjson1_serializeOpCreateFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1296,7 +1251,6 @@ func (m *awsRestjson1_serializeOpCreateFilter) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFilterInput(v *CreateFilterInput, encoder *httpbinding.Encoder) error {
@@ -1358,10 +1312,8 @@ func (m *awsRestjson1_serializeOpCreateFindingsReport)
 func (m *awsRestjson1_serializeOpCreateFindingsReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1406,7 +1358,6 @@ func (m *awsRestjson1_serializeOpCreateFindingsReport)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFindingsReportInput(v *CreateFindingsReportInput, encoder *httpbinding.Encoder) error {
@@ -1453,10 +1404,8 @@ func (m *awsRestjson1_serializeOpCreateSbomExport) Han
 func (m *awsRestjson1_serializeOpCreateSbomExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1501,7 +1450,6 @@ func (m *awsRestjson1_serializeOpCreateSbomExport) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSbomExportInput(v *CreateSbomExportInput, encoder *httpbinding.Encoder) error {
@@ -1548,10 +1496,8 @@ func (m *awsRestjson1_serializeOpDeleteCisScanConfigur
 func (m *awsRestjson1_serializeOpDeleteCisScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1596,7 +1542,6 @@ func (m *awsRestjson1_serializeOpDeleteCisScanConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCisScanConfigurationInput(v *DeleteCisScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1629,10 +1574,8 @@ func (m *awsRestjson1_serializeOpDeleteCodeSecurityInt
 func (m *awsRestjson1_serializeOpDeleteCodeSecurityIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1677,7 +1620,6 @@ func (m *awsRestjson1_serializeOpDeleteCodeSecurityInt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCodeSecurityIntegrationInput(v *DeleteCodeSecurityIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -1710,10 +1652,8 @@ func (m *awsRestjson1_serializeOpDeleteCodeSecuritySca
 func (m *awsRestjson1_serializeOpDeleteCodeSecurityScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1758,7 +1698,6 @@ func (m *awsRestjson1_serializeOpDeleteCodeSecuritySca
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCodeSecurityScanConfigurationInput(v *DeleteCodeSecurityScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1791,10 +1730,8 @@ func (m *awsRestjson1_serializeOpDeleteFilter) HandleS
 func (m *awsRestjson1_serializeOpDeleteFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1839,7 +1776,6 @@ func (m *awsRestjson1_serializeOpDeleteFilter) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFilterInput(v *DeleteFilterInput, encoder *httpbinding.Encoder) error {
@@ -1872,10 +1808,8 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
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
@@ -1909,7 +1843,6 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOrganizationConfigurationInput(v *DescribeOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1930,10 +1863,8 @@ func (m *awsRestjson1_serializeOpDisable) HandleSerial
 func (m *awsRestjson1_serializeOpDisable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1978,7 +1909,6 @@ func (m *awsRestjson1_serializeOpDisable) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableInput(v *DisableInput, encoder *httpbinding.Encoder) error {
@@ -2020,10 +1950,8 @@ func (m *awsRestjson1_serializeOpDisableDelegatedAdmin
 func (m *awsRestjson1_serializeOpDisableDelegatedAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2068,7 +1996,6 @@ func (m *awsRestjson1_serializeOpDisableDelegatedAdmin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableDelegatedAdminAccountInput(v *DisableDelegatedAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -2101,10 +2028,8 @@ func (m *awsRestjson1_serializeOpDisassociateMember) H
 func (m *awsRestjson1_serializeOpDisassociateMember) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2149,7 +2074,6 @@ func (m *awsRestjson1_serializeOpDisassociateMember) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMemberInput(v *DisassociateMemberInput, encoder *httpbinding.Encoder) error {
@@ -2182,10 +2106,8 @@ func (m *awsRestjson1_serializeOpEnable) HandleSeriali
 func (m *awsRestjson1_serializeOpEnable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2230,7 +2152,6 @@ func (m *awsRestjson1_serializeOpEnable) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableInput(v *EnableInput, encoder *httpbinding.Encoder) error {
@@ -2277,10 +2198,8 @@ func (m *awsRestjson1_serializeOpEnableDelegatedAdminA
 func (m *awsRestjson1_serializeOpEnableDelegatedAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2325,7 +2244,6 @@ func (m *awsRestjson1_serializeOpEnableDelegatedAdminA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableDelegatedAdminAccountInput(v *EnableDelegatedAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -2363,10 +2281,8 @@ func (m *awsRestjson1_serializeOpGetCisScanReport) Han
 func (m *awsRestjson1_serializeOpGetCisScanReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2411,7 +2327,6 @@ func (m *awsRestjson1_serializeOpGetCisScanReport) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCisScanReportInput(v *GetCisScanReportInput, encoder *httpbinding.Encoder) error {
@@ -2456,10 +2371,8 @@ func (m *awsRestjson1_serializeOpGetCisScanResultDetai
 func (m *awsRestjson1_serializeOpGetCisScanResultDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2504,7 +2417,6 @@ func (m *awsRestjson1_serializeOpGetCisScanResultDetai
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCisScanResultDetailsInput(v *GetCisScanResultDetailsInput, encoder *httpbinding.Encoder) error {
@@ -2574,10 +2486,8 @@ func (m *awsRestjson1_serializeOpGetClustersForImage) 
 func (m *awsRestjson1_serializeOpGetClustersForImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2622,7 +2532,6 @@ func (m *awsRestjson1_serializeOpGetClustersForImage) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetClustersForImageInput(v *GetClustersForImageInput, encoder *httpbinding.Encoder) error {
@@ -2667,10 +2576,8 @@ func (m *awsRestjson1_serializeOpGetCodeSecurityIntegr
 func (m *awsRestjson1_serializeOpGetCodeSecurityIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2715,7 +2622,6 @@ func (m *awsRestjson1_serializeOpGetCodeSecurityIntegr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCodeSecurityIntegrationInput(v *GetCodeSecurityIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -2755,10 +2661,8 @@ func (m *awsRestjson1_serializeOpGetCodeSecurityScan) 
 func (m *awsRestjson1_serializeOpGetCodeSecurityScan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2803,7 +2707,6 @@ func (m *awsRestjson1_serializeOpGetCodeSecurityScan) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCodeSecurityScanInput(v *GetCodeSecurityScanInput, encoder *httpbinding.Encoder) error {
@@ -2843,10 +2746,8 @@ func (m *awsRestjson1_serializeOpGetCodeSecurityScanCo
 func (m *awsRestjson1_serializeOpGetCodeSecurityScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2891,7 +2792,6 @@ func (m *awsRestjson1_serializeOpGetCodeSecurityScanCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCodeSecurityScanConfigurationInput(v *GetCodeSecurityScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2924,10 +2824,8 @@ func (m *awsRestjson1_serializeOpGetConfiguration) Han
 func (m *awsRestjson1_serializeOpGetConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2961,7 +2859,6 @@ func (m *awsRestjson1_serializeOpGetConfiguration) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfigurationInput(v *GetConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2982,10 +2879,8 @@ func (m *awsRestjson1_serializeOpGetDelegatedAdminAcco
 func (m *awsRestjson1_serializeOpGetDelegatedAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3019,7 +2914,6 @@ func (m *awsRestjson1_serializeOpGetDelegatedAdminAcco
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDelegatedAdminAccountInput(v *GetDelegatedAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -3040,10 +2934,8 @@ func (m *awsRestjson1_serializeOpGetEc2DeepInspectionC
 func (m *awsRestjson1_serializeOpGetEc2DeepInspectionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3077,7 +2969,6 @@ func (m *awsRestjson1_serializeOpGetEc2DeepInspectionC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEc2DeepInspectionConfigurationInput(v *GetEc2DeepInspectionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3098,10 +2989,8 @@ func (m *awsRestjson1_serializeOpGetEncryptionKey) Han
 func (m *awsRestjson1_serializeOpGetEncryptionKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3139,7 +3028,6 @@ func (m *awsRestjson1_serializeOpGetEncryptionKey) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEncryptionKeyInput(v *GetEncryptionKeyInput, encoder *httpbinding.Encoder) error {
@@ -3168,10 +3056,8 @@ func (m *awsRestjson1_serializeOpGetFindingsReportStat
 func (m *awsRestjson1_serializeOpGetFindingsReportStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3216,7 +3102,6 @@ func (m *awsRestjson1_serializeOpGetFindingsReportStat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsReportStatusInput(v *GetFindingsReportStatusInput, encoder *httpbinding.Encoder) error {
@@ -3249,10 +3134,8 @@ func (m *awsRestjson1_serializeOpGetMember) HandleSeri
 func (m *awsRestjson1_serializeOpGetMember) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3297,7 +3180,6 @@ func (m *awsRestjson1_serializeOpGetMember) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMemberInput(v *GetMemberInput, encoder *httpbinding.Encoder) error {
@@ -3330,10 +3212,8 @@ func (m *awsRestjson1_serializeOpGetSbomExport) Handle
 func (m *awsRestjson1_serializeOpGetSbomExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3378,7 +3258,6 @@ func (m *awsRestjson1_serializeOpGetSbomExport) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSbomExportInput(v *GetSbomExportInput, encoder *httpbinding.Encoder) error {
@@ -3411,10 +3290,8 @@ func (m *awsRestjson1_serializeOpListAccountPermission
 func (m *awsRestjson1_serializeOpListAccountPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3459,7 +3336,6 @@ func (m *awsRestjson1_serializeOpListAccountPermission
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAccountPermissionsInput(v *ListAccountPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -3502,10 +3378,8 @@ func (m *awsRestjson1_serializeOpListCisScanConfigurat
 func (m *awsRestjson1_serializeOpListCisScanConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3550,7 +3424,6 @@ func (m *awsRestjson1_serializeOpListCisScanConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCisScanConfigurationsInput(v *ListCisScanConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -3605,10 +3478,8 @@ func (m *awsRestjson1_serializeOpListCisScanResultsAgg
 func (m *awsRestjson1_serializeOpListCisScanResultsAggregatedByChecks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3653,7 +3524,6 @@ func (m *awsRestjson1_serializeOpListCisScanResultsAgg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCisScanResultsAggregatedByChecksInput(v *ListCisScanResultsAggregatedByChecksInput, encoder *httpbinding.Encoder) error {
@@ -3713,10 +3583,8 @@ func (m *awsRestjson1_serializeOpListCisScanResultsAgg
 func (m *awsRestjson1_serializeOpListCisScanResultsAggregatedByTargetResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3761,7 +3629,6 @@ func (m *awsRestjson1_serializeOpListCisScanResultsAgg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCisScanResultsAggregatedByTargetResourceInput(v *ListCisScanResultsAggregatedByTargetResourceInput, encoder *httpbinding.Encoder) error {
@@ -3821,10 +3688,8 @@ func (m *awsRestjson1_serializeOpListCisScans) HandleS
 func (m *awsRestjson1_serializeOpListCisScans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3869,7 +3734,6 @@ func (m *awsRestjson1_serializeOpListCisScans) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCisScansInput(v *ListCisScansInput, encoder *httpbinding.Encoder) error {
@@ -3929,10 +3793,8 @@ func (m *awsRestjson1_serializeOpListCodeSecurityInteg
 func (m *awsRestjson1_serializeOpListCodeSecurityIntegrations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3970,7 +3832,6 @@ func (m *awsRestjson1_serializeOpListCodeSecurityInteg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCodeSecurityIntegrationsInput(v *ListCodeSecurityIntegrationsInput, encoder *httpbinding.Encoder) error {
@@ -3999,10 +3860,8 @@ func (m *awsRestjson1_serializeOpListCodeSecurityScanC
 func (m *awsRestjson1_serializeOpListCodeSecurityScanConfigurationAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4051,7 +3910,6 @@ func (m *awsRestjson1_serializeOpListCodeSecurityScanC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCodeSecurityScanConfigurationAssociationsInput(v *ListCodeSecurityScanConfigurationAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4092,10 +3950,8 @@ func (m *awsRestjson1_serializeOpListCodeSecurityScanC
 func (m *awsRestjson1_serializeOpListCodeSecurityScanConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4133,7 +3989,6 @@ func (m *awsRestjson1_serializeOpListCodeSecurityScanC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCodeSecurityScanConfigurationsInput(v *ListCodeSecurityScanConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -4162,10 +4017,8 @@ func (m *awsRestjson1_serializeOpListCoverage) HandleS
 func (m *awsRestjson1_serializeOpListCoverage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4210,7 +4063,6 @@ func (m *awsRestjson1_serializeOpListCoverage) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCoverageInput(v *ListCoverageInput, encoder *httpbinding.Encoder) error {
@@ -4255,10 +4107,8 @@ func (m *awsRestjson1_serializeOpListCoverageStatistic
 func (m *awsRestjson1_serializeOpListCoverageStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4303,7 +4153,6 @@ func (m *awsRestjson1_serializeOpListCoverageStatistic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCoverageStatisticsInput(v *ListCoverageStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -4348,10 +4197,8 @@ func (m *awsRestjson1_serializeOpListDelegatedAdminAcc
 func (m *awsRestjson1_serializeOpListDelegatedAdminAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4396,7 +4243,6 @@ func (m *awsRestjson1_serializeOpListDelegatedAdminAcc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDelegatedAdminAccountsInput(v *ListDelegatedAdminAccountsInput, encoder *httpbinding.Encoder) error {
@@ -4434,10 +4280,8 @@ func (m *awsRestjson1_serializeOpListFilters) HandleSe
 func (m *awsRestjson1_serializeOpListFilters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4482,7 +4326,6 @@ func (m *awsRestjson1_serializeOpListFilters) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFiltersInput(v *ListFiltersInput, encoder *httpbinding.Encoder) error {
@@ -4532,10 +4375,8 @@ func (m *awsRestjson1_serializeOpListFindingAggregatio
 func (m *awsRestjson1_serializeOpListFindingAggregations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4580,7 +4421,6 @@ func (m *awsRestjson1_serializeOpListFindingAggregatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFindingAggregationsInput(v *ListFindingAggregationsInput, encoder *httpbinding.Encoder) error {
@@ -4637,10 +4477,8 @@ func (m *awsRestjson1_serializeOpListFindings) HandleS
 func (m *awsRestjson1_serializeOpListFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4685,7 +4523,6 @@ func (m *awsRestjson1_serializeOpListFindings) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFindingsInput(v *ListFindingsInput, encoder *httpbinding.Encoder) error {
@@ -4737,10 +4574,8 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
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
@@ -4785,7 +4620,6 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMembersInput(v *ListMembersInput, encoder *httpbinding.Encoder) error {
@@ -4828,10 +4662,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -4869,7 +4701,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4899,10 +4730,8 @@ func (m *awsRestjson1_serializeOpListUsageTotals) Hand
 func (m *awsRestjson1_serializeOpListUsageTotals) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4947,7 +4776,6 @@ func (m *awsRestjson1_serializeOpListUsageTotals) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUsageTotalsInput(v *ListUsageTotalsInput, encoder *httpbinding.Encoder) error {
@@ -4992,10 +4820,8 @@ func (m *awsRestjson1_serializeOpResetEncryptionKey) H
 func (m *awsRestjson1_serializeOpResetEncryptionKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5040,7 +4866,6 @@ func (m *awsRestjson1_serializeOpResetEncryptionKey) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResetEncryptionKeyInput(v *ResetEncryptionKeyInput, encoder *httpbinding.Encoder) error {
@@ -5078,10 +4903,8 @@ func (m *awsRestjson1_serializeOpSearchVulnerabilities
 func (m *awsRestjson1_serializeOpSearchVulnerabilities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5126,7 +4949,6 @@ func (m *awsRestjson1_serializeOpSearchVulnerabilities
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchVulnerabilitiesInput(v *SearchVulnerabilitiesInput, encoder *httpbinding.Encoder) error {
@@ -5166,10 +4988,8 @@ func (m *awsRestjson1_serializeOpSendCisSessionHealth)
 func (m *awsRestjson1_serializeOpSendCisSessionHealth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5214,7 +5034,6 @@ func (m *awsRestjson1_serializeOpSendCisSessionHealth)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendCisSessionHealthInput(v *SendCisSessionHealthInput, encoder *httpbinding.Encoder) error {
@@ -5252,10 +5071,8 @@ func (m *awsRestjson1_serializeOpSendCisSessionTelemet
 func (m *awsRestjson1_serializeOpSendCisSessionTelemetry) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5300,7 +5117,6 @@ func (m *awsRestjson1_serializeOpSendCisSessionTelemet
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendCisSessionTelemetryInput(v *SendCisSessionTelemetryInput, encoder *httpbinding.Encoder) error {
@@ -5345,10 +5161,8 @@ func (m *awsRestjson1_serializeOpStartCisSession) Hand
 func (m *awsRestjson1_serializeOpStartCisSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5393,7 +5207,6 @@ func (m *awsRestjson1_serializeOpStartCisSession) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartCisSessionInput(v *StartCisSessionInput, encoder *httpbinding.Encoder) error {
@@ -5433,10 +5246,8 @@ func (m *awsRestjson1_serializeOpStartCodeSecurityScan
 func (m *awsRestjson1_serializeOpStartCodeSecurityScan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5481,7 +5292,6 @@ func (m *awsRestjson1_serializeOpStartCodeSecurityScan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartCodeSecurityScanInput(v *StartCodeSecurityScanInput, encoder *httpbinding.Encoder) error {
@@ -5521,10 +5331,8 @@ func (m *awsRestjson1_serializeOpStopCisSession) Handl
 func (m *awsRestjson1_serializeOpStopCisSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5569,7 +5377,6 @@ func (m *awsRestjson1_serializeOpStopCisSession) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopCisSessionInput(v *StopCisSessionInput, encoder *httpbinding.Encoder) error {
@@ -5614,10 +5421,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5666,7 +5471,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5710,10 +5514,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5751,7 +5553,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5787,10 +5588,8 @@ func (m *awsRestjson1_serializeOpUpdateCisScanConfigur
 func (m *awsRestjson1_serializeOpUpdateCisScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5835,7 +5634,6 @@ func (m *awsRestjson1_serializeOpUpdateCisScanConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCisScanConfigurationInput(v *UpdateCisScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5892,10 +5690,8 @@ func (m *awsRestjson1_serializeOpUpdateCodeSecurityInt
 func (m *awsRestjson1_serializeOpUpdateCodeSecurityIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5940,7 +5736,6 @@ func (m *awsRestjson1_serializeOpUpdateCodeSecurityInt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCodeSecurityIntegrationInput(v *UpdateCodeSecurityIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -5980,10 +5775,8 @@ func (m *awsRestjson1_serializeOpUpdateCodeSecuritySca
 func (m *awsRestjson1_serializeOpUpdateCodeSecurityScanConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6028,7 +5821,6 @@ func (m *awsRestjson1_serializeOpUpdateCodeSecuritySca
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCodeSecurityScanConfigurationInput(v *UpdateCodeSecurityScanConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6068,10 +5860,8 @@ func (m *awsRestjson1_serializeOpUpdateConfiguration) 
 func (m *awsRestjson1_serializeOpUpdateConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6116,7 +5906,6 @@ func (m *awsRestjson1_serializeOpUpdateConfiguration) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfigurationInput(v *UpdateConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6158,10 +5947,8 @@ func (m *awsRestjson1_serializeOpUpdateEc2DeepInspecti
 func (m *awsRestjson1_serializeOpUpdateEc2DeepInspectionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6206,7 +5993,6 @@ func (m *awsRestjson1_serializeOpUpdateEc2DeepInspecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEc2DeepInspectionConfigurationInput(v *UpdateEc2DeepInspectionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6246,10 +6032,8 @@ func (m *awsRestjson1_serializeOpUpdateEncryptionKey) 
 func (m *awsRestjson1_serializeOpUpdateEncryptionKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6294,7 +6078,6 @@ func (m *awsRestjson1_serializeOpUpdateEncryptionKey) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEncryptionKeyInput(v *UpdateEncryptionKeyInput, encoder *httpbinding.Encoder) error {
@@ -6337,10 +6120,8 @@ func (m *awsRestjson1_serializeOpUpdateFilter) HandleS
 func (m *awsRestjson1_serializeOpUpdateFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6385,7 +6166,6 @@ func (m *awsRestjson1_serializeOpUpdateFilter) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFilterInput(v *UpdateFilterInput, encoder *httpbinding.Encoder) error {
@@ -6445,10 +6225,8 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
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
@@ -6493,7 +6271,6 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOrganizationConfigurationInput(v *UpdateOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6528,10 +6305,8 @@ func (m *awsRestjson1_serializeOpUpdateOrgEc2DeepInspe
 func (m *awsRestjson1_serializeOpUpdateOrgEc2DeepInspectionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6576,7 +6351,6 @@ func (m *awsRestjson1_serializeOpUpdateOrgEc2DeepInspe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOrgEc2DeepInspectionConfigurationInput(v *UpdateOrgEc2DeepInspectionConfigurationInput, encoder *httpbinding.Encoder) error {
