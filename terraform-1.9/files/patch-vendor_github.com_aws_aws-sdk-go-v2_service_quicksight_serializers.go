--- vendor/github.com/aws/aws-sdk-go-v2/service/quicksight/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/quicksight/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 	"time"
@@ -28,10 +27,8 @@ func (m *awsRestjson1_serializeOpBatchCreateTopicRevie
 func (m *awsRestjson1_serializeOpBatchCreateTopicReviewedAnswer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -80,7 +77,6 @@ func (m *awsRestjson1_serializeOpBatchCreateTopicRevie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchCreateTopicReviewedAnswerInput(v *BatchCreateTopicReviewedAnswerInput, encoder *httpbinding.Encoder) error {
@@ -133,10 +129,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteTopicRevie
 func (m *awsRestjson1_serializeOpBatchDeleteTopicReviewedAnswer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -185,7 +179,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteTopicRevie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteTopicReviewedAnswerInput(v *BatchDeleteTopicReviewedAnswerInput, encoder *httpbinding.Encoder) error {
@@ -238,10 +231,8 @@ func (m *awsRestjson1_serializeOpCancelIngestion) Hand
 func (m *awsRestjson1_serializeOpCancelIngestion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -279,7 +270,6 @@ func (m *awsRestjson1_serializeOpCancelIngestion) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelIngestionInput(v *CancelIngestionInput, encoder *httpbinding.Encoder) error {
@@ -327,10 +317,8 @@ func (m *awsRestjson1_serializeOpCreateAccountCustomiz
 func (m *awsRestjson1_serializeOpCreateAccountCustomization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -379,7 +367,6 @@ func (m *awsRestjson1_serializeOpCreateAccountCustomiz
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAccountCustomizationInput(v *CreateAccountCustomizationInput, encoder *httpbinding.Encoder) error {
@@ -434,10 +421,8 @@ func (m *awsRestjson1_serializeOpCreateAccountSubscrip
 func (m *awsRestjson1_serializeOpCreateAccountSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -486,7 +471,6 @@ func (m *awsRestjson1_serializeOpCreateAccountSubscrip
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAccountSubscriptionInput(v *CreateAccountSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -625,10 +609,8 @@ func (m *awsRestjson1_serializeOpCreateAnalysis) Handl
 func (m *awsRestjson1_serializeOpCreateAnalysis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -677,7 +659,6 @@ func (m *awsRestjson1_serializeOpCreateAnalysis) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAnalysisInput(v *CreateAnalysisInput, encoder *httpbinding.Encoder) error {
@@ -782,10 +763,8 @@ func (m *awsRestjson1_serializeOpCreateBrand) HandleSe
 func (m *awsRestjson1_serializeOpCreateBrand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -834,7 +813,6 @@ func (m *awsRestjson1_serializeOpCreateBrand) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBrandInput(v *CreateBrandInput, encoder *httpbinding.Encoder) error {
@@ -894,10 +872,8 @@ func (m *awsRestjson1_serializeOpCreateCustomPermissio
 func (m *awsRestjson1_serializeOpCreateCustomPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -946,7 +922,6 @@ func (m *awsRestjson1_serializeOpCreateCustomPermissio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCustomPermissionsInput(v *CreateCustomPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -1002,10 +977,8 @@ func (m *awsRestjson1_serializeOpCreateDashboard) Hand
 func (m *awsRestjson1_serializeOpCreateDashboard) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1054,7 +1027,6 @@ func (m *awsRestjson1_serializeOpCreateDashboard) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDashboardInput(v *CreateDashboardInput, encoder *httpbinding.Encoder) error {
@@ -1185,10 +1157,8 @@ func (m *awsRestjson1_serializeOpCreateDataSet) Handle
 func (m *awsRestjson1_serializeOpCreateDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1237,7 +1207,6 @@ func (m *awsRestjson1_serializeOpCreateDataSet) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataSetInput(v *CreateDataSetInput, encoder *httpbinding.Encoder) error {
@@ -1385,10 +1354,8 @@ func (m *awsRestjson1_serializeOpCreateDataSource) Han
 func (m *awsRestjson1_serializeOpCreateDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1437,7 +1404,6 @@ func (m *awsRestjson1_serializeOpCreateDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataSourceInput(v *CreateDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -1538,10 +1504,8 @@ func (m *awsRestjson1_serializeOpCreateFolder) HandleS
 func (m *awsRestjson1_serializeOpCreateFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1590,7 +1554,6 @@ func (m *awsRestjson1_serializeOpCreateFolder) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFolderInput(v *CreateFolderInput, encoder *httpbinding.Encoder) error {
@@ -1670,10 +1633,8 @@ func (m *awsRestjson1_serializeOpCreateFolderMembershi
 func (m *awsRestjson1_serializeOpCreateFolderMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1711,7 +1672,6 @@ func (m *awsRestjson1_serializeOpCreateFolderMembershi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFolderMembershipInput(v *CreateFolderMembershipInput, encoder *httpbinding.Encoder) error {
@@ -1768,10 +1728,8 @@ func (m *awsRestjson1_serializeOpCreateGroup) HandleSe
 func (m *awsRestjson1_serializeOpCreateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1820,7 +1778,6 @@ func (m *awsRestjson1_serializeOpCreateGroup) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGroupInput(v *CreateGroupInput, encoder *httpbinding.Encoder) error {
@@ -1876,10 +1833,8 @@ func (m *awsRestjson1_serializeOpCreateGroupMembership
 func (m *awsRestjson1_serializeOpCreateGroupMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1917,7 +1872,6 @@ func (m *awsRestjson1_serializeOpCreateGroupMembership
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGroupMembershipInput(v *CreateGroupMembershipInput, encoder *httpbinding.Encoder) error {
@@ -1974,10 +1928,8 @@ func (m *awsRestjson1_serializeOpCreateIAMPolicyAssign
 func (m *awsRestjson1_serializeOpCreateIAMPolicyAssignment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2026,7 +1978,6 @@ func (m *awsRestjson1_serializeOpCreateIAMPolicyAssign
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIAMPolicyAssignmentInput(v *CreateIAMPolicyAssignmentInput, encoder *httpbinding.Encoder) error {
@@ -2094,10 +2045,8 @@ func (m *awsRestjson1_serializeOpCreateIngestion) Hand
 func (m *awsRestjson1_serializeOpCreateIngestion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2146,7 +2095,6 @@ func (m *awsRestjson1_serializeOpCreateIngestion) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIngestionInput(v *CreateIngestionInput, encoder *httpbinding.Encoder) error {
@@ -2206,10 +2154,8 @@ func (m *awsRestjson1_serializeOpCreateNamespace) Hand
 func (m *awsRestjson1_serializeOpCreateNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2258,7 +2204,6 @@ func (m *awsRestjson1_serializeOpCreateNamespace) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNamespaceInput(v *CreateNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -2312,10 +2257,8 @@ func (m *awsRestjson1_serializeOpCreateRefreshSchedule
 func (m *awsRestjson1_serializeOpCreateRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2364,7 +2307,6 @@ func (m *awsRestjson1_serializeOpCreateRefreshSchedule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRefreshScheduleInput(v *CreateRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -2417,10 +2359,8 @@ func (m *awsRestjson1_serializeOpCreateRoleMembership)
 func (m *awsRestjson1_serializeOpCreateRoleMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2458,7 +2398,6 @@ func (m *awsRestjson1_serializeOpCreateRoleMembership)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRoleMembershipInput(v *CreateRoleMembershipInput, encoder *httpbinding.Encoder) error {
@@ -2515,10 +2454,8 @@ func (m *awsRestjson1_serializeOpCreateTemplate) Handl
 func (m *awsRestjson1_serializeOpCreateTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2567,7 +2504,6 @@ func (m *awsRestjson1_serializeOpCreateTemplate) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTemplateInput(v *CreateTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2658,10 +2594,8 @@ func (m *awsRestjson1_serializeOpCreateTemplateAlias) 
 func (m *awsRestjson1_serializeOpCreateTemplateAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2710,7 +2644,6 @@ func (m *awsRestjson1_serializeOpCreateTemplateAlias) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTemplateAliasInput(v *CreateTemplateAliasInput, encoder *httpbinding.Encoder) error {
@@ -2770,10 +2703,8 @@ func (m *awsRestjson1_serializeOpCreateTheme) HandleSe
 func (m *awsRestjson1_serializeOpCreateTheme) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2822,7 +2753,6 @@ func (m *awsRestjson1_serializeOpCreateTheme) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateThemeInput(v *CreateThemeInput, encoder *httpbinding.Encoder) error {
@@ -2904,10 +2834,8 @@ func (m *awsRestjson1_serializeOpCreateThemeAlias) Han
 func (m *awsRestjson1_serializeOpCreateThemeAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2956,7 +2884,6 @@ func (m *awsRestjson1_serializeOpCreateThemeAlias) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateThemeAliasInput(v *CreateThemeAliasInput, encoder *httpbinding.Encoder) error {
@@ -3016,10 +2943,8 @@ func (m *awsRestjson1_serializeOpCreateTopic) HandleSe
 func (m *awsRestjson1_serializeOpCreateTopic) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3068,7 +2993,6 @@ func (m *awsRestjson1_serializeOpCreateTopic) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTopicInput(v *CreateTopicInput, encoder *httpbinding.Encoder) error {
@@ -3138,10 +3062,8 @@ func (m *awsRestjson1_serializeOpCreateTopicRefreshSch
 func (m *awsRestjson1_serializeOpCreateTopicRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3190,7 +3112,6 @@ func (m *awsRestjson1_serializeOpCreateTopicRefreshSch
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTopicRefreshScheduleInput(v *CreateTopicRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -3253,10 +3174,8 @@ func (m *awsRestjson1_serializeOpCreateVPCConnection) 
 func (m *awsRestjson1_serializeOpCreateVPCConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3305,7 +3224,6 @@ func (m *awsRestjson1_serializeOpCreateVPCConnection) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVPCConnectionInput(v *CreateVPCConnectionInput, encoder *httpbinding.Encoder) error {
@@ -3385,10 +3303,8 @@ func (m *awsRestjson1_serializeOpDeleteAccountCustomiz
 func (m *awsRestjson1_serializeOpDeleteAccountCustomization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3426,7 +3342,6 @@ func (m *awsRestjson1_serializeOpDeleteAccountCustomiz
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccountCustomizationInput(v *DeleteAccountCustomizationInput, encoder *httpbinding.Encoder) error {
@@ -3460,10 +3375,8 @@ func (m *awsRestjson1_serializeOpDeleteAccountSubscrip
 func (m *awsRestjson1_serializeOpDeleteAccountSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3501,7 +3414,6 @@ func (m *awsRestjson1_serializeOpDeleteAccountSubscrip
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccountSubscriptionInput(v *DeleteAccountSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -3531,10 +3443,8 @@ func (m *awsRestjson1_serializeOpDeleteAnalysis) Handl
 func (m *awsRestjson1_serializeOpDeleteAnalysis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3572,7 +3482,6 @@ func (m *awsRestjson1_serializeOpDeleteAnalysis) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAnalysisInput(v *DeleteAnalysisInput, encoder *httpbinding.Encoder) error {
@@ -3619,10 +3528,8 @@ func (m *awsRestjson1_serializeOpDeleteBrand) HandleSe
 func (m *awsRestjson1_serializeOpDeleteBrand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3660,7 +3567,6 @@ func (m *awsRestjson1_serializeOpDeleteBrand) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBrandInput(v *DeleteBrandInput, encoder *httpbinding.Encoder) error {
@@ -3699,10 +3605,8 @@ func (m *awsRestjson1_serializeOpDeleteBrandAssignment
 func (m *awsRestjson1_serializeOpDeleteBrandAssignment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3740,7 +3644,6 @@ func (m *awsRestjson1_serializeOpDeleteBrandAssignment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBrandAssignmentInput(v *DeleteBrandAssignmentInput, encoder *httpbinding.Encoder) error {
@@ -3770,10 +3673,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomPermissio
 func (m *awsRestjson1_serializeOpDeleteCustomPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3811,7 +3712,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomPermissio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomPermissionsInput(v *DeleteCustomPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -3850,10 +3750,8 @@ func (m *awsRestjson1_serializeOpDeleteDashboard) Hand
 func (m *awsRestjson1_serializeOpDeleteDashboard) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3891,7 +3789,6 @@ func (m *awsRestjson1_serializeOpDeleteDashboard) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDashboardInput(v *DeleteDashboardInput, encoder *httpbinding.Encoder) error {
@@ -3934,10 +3831,8 @@ func (m *awsRestjson1_serializeOpDeleteDataSet) Handle
 func (m *awsRestjson1_serializeOpDeleteDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3975,7 +3870,6 @@ func (m *awsRestjson1_serializeOpDeleteDataSet) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataSetInput(v *DeleteDataSetInput, encoder *httpbinding.Encoder) error {
@@ -4014,10 +3908,8 @@ func (m *awsRestjson1_serializeOpDeleteDataSetRefreshP
 func (m *awsRestjson1_serializeOpDeleteDataSetRefreshProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4055,7 +3947,6 @@ func (m *awsRestjson1_serializeOpDeleteDataSetRefreshP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataSetRefreshPropertiesInput(v *DeleteDataSetRefreshPropertiesInput, encoder *httpbinding.Encoder) error {
@@ -4094,10 +3985,8 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 func (m *awsRestjson1_serializeOpDeleteDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4135,7 +4024,6 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataSourceInput(v *DeleteDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -4174,10 +4062,8 @@ func (m *awsRestjson1_serializeOpDeleteDefaultQBusines
 func (m *awsRestjson1_serializeOpDeleteDefaultQBusinessApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4215,7 +4101,6 @@ func (m *awsRestjson1_serializeOpDeleteDefaultQBusines
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDefaultQBusinessApplicationInput(v *DeleteDefaultQBusinessApplicationInput, encoder *httpbinding.Encoder) error {
@@ -4249,10 +4134,8 @@ func (m *awsRestjson1_serializeOpDeleteFolder) HandleS
 func (m *awsRestjson1_serializeOpDeleteFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4290,7 +4173,6 @@ func (m *awsRestjson1_serializeOpDeleteFolder) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFolderInput(v *DeleteFolderInput, encoder *httpbinding.Encoder) error {
@@ -4329,10 +4211,8 @@ func (m *awsRestjson1_serializeOpDeleteFolderMembershi
 func (m *awsRestjson1_serializeOpDeleteFolderMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4370,7 +4250,6 @@ func (m *awsRestjson1_serializeOpDeleteFolderMembershi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFolderMembershipInput(v *DeleteFolderMembershipInput, encoder *httpbinding.Encoder) error {
@@ -4427,10 +4306,8 @@ func (m *awsRestjson1_serializeOpDeleteGroup) HandleSe
 func (m *awsRestjson1_serializeOpDeleteGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4468,7 +4345,6 @@ func (m *awsRestjson1_serializeOpDeleteGroup) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGroupInput(v *DeleteGroupInput, encoder *httpbinding.Encoder) error {
@@ -4516,10 +4392,8 @@ func (m *awsRestjson1_serializeOpDeleteGroupMembership
 func (m *awsRestjson1_serializeOpDeleteGroupMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4557,7 +4431,6 @@ func (m *awsRestjson1_serializeOpDeleteGroupMembership
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGroupMembershipInput(v *DeleteGroupMembershipInput, encoder *httpbinding.Encoder) error {
@@ -4614,10 +4487,8 @@ func (m *awsRestjson1_serializeOpDeleteIAMPolicyAssign
 func (m *awsRestjson1_serializeOpDeleteIAMPolicyAssignment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4655,7 +4526,6 @@ func (m *awsRestjson1_serializeOpDeleteIAMPolicyAssign
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIAMPolicyAssignmentInput(v *DeleteIAMPolicyAssignmentInput, encoder *httpbinding.Encoder) error {
@@ -4703,10 +4573,8 @@ func (m *awsRestjson1_serializeOpDeleteIdentityPropaga
 func (m *awsRestjson1_serializeOpDeleteIdentityPropagationConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4744,7 +4612,6 @@ func (m *awsRestjson1_serializeOpDeleteIdentityPropaga
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIdentityPropagationConfigInput(v *DeleteIdentityPropagationConfigInput, encoder *httpbinding.Encoder) error {
@@ -4783,10 +4650,8 @@ func (m *awsRestjson1_serializeOpDeleteNamespace) Hand
 func (m *awsRestjson1_serializeOpDeleteNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4824,7 +4689,6 @@ func (m *awsRestjson1_serializeOpDeleteNamespace) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteNamespaceInput(v *DeleteNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -4863,10 +4727,8 @@ func (m *awsRestjson1_serializeOpDeleteRefreshSchedule
 func (m *awsRestjson1_serializeOpDeleteRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4904,7 +4766,6 @@ func (m *awsRestjson1_serializeOpDeleteRefreshSchedule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRefreshScheduleInput(v *DeleteRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -4952,10 +4813,8 @@ func (m *awsRestjson1_serializeOpDeleteRoleCustomPermi
 func (m *awsRestjson1_serializeOpDeleteRoleCustomPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4993,7 +4852,6 @@ func (m *awsRestjson1_serializeOpDeleteRoleCustomPermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRoleCustomPermissionInput(v *DeleteRoleCustomPermissionInput, encoder *httpbinding.Encoder) error {
@@ -5041,10 +4899,8 @@ func (m *awsRestjson1_serializeOpDeleteRoleMembership)
 func (m *awsRestjson1_serializeOpDeleteRoleMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5082,7 +4938,6 @@ func (m *awsRestjson1_serializeOpDeleteRoleMembership)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRoleMembershipInput(v *DeleteRoleMembershipInput, encoder *httpbinding.Encoder) error {
@@ -5139,10 +4994,8 @@ func (m *awsRestjson1_serializeOpDeleteTemplate) Handl
 func (m *awsRestjson1_serializeOpDeleteTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5180,7 +5033,6 @@ func (m *awsRestjson1_serializeOpDeleteTemplate) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTemplateInput(v *DeleteTemplateInput, encoder *httpbinding.Encoder) error {
@@ -5223,10 +5075,8 @@ func (m *awsRestjson1_serializeOpDeleteTemplateAlias) 
 func (m *awsRestjson1_serializeOpDeleteTemplateAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5264,7 +5114,6 @@ func (m *awsRestjson1_serializeOpDeleteTemplateAlias) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTemplateAliasInput(v *DeleteTemplateAliasInput, encoder *httpbinding.Encoder) error {
@@ -5312,10 +5161,8 @@ func (m *awsRestjson1_serializeOpDeleteTheme) HandleSe
 func (m *awsRestjson1_serializeOpDeleteTheme) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5353,7 +5200,6 @@ func (m *awsRestjson1_serializeOpDeleteTheme) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteThemeInput(v *DeleteThemeInput, encoder *httpbinding.Encoder) error {
@@ -5396,10 +5242,8 @@ func (m *awsRestjson1_serializeOpDeleteThemeAlias) Han
 func (m *awsRestjson1_serializeOpDeleteThemeAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5437,7 +5281,6 @@ func (m *awsRestjson1_serializeOpDeleteThemeAlias) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteThemeAliasInput(v *DeleteThemeAliasInput, encoder *httpbinding.Encoder) error {
@@ -5485,10 +5328,8 @@ func (m *awsRestjson1_serializeOpDeleteTopic) HandleSe
 func (m *awsRestjson1_serializeOpDeleteTopic) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5526,7 +5367,6 @@ func (m *awsRestjson1_serializeOpDeleteTopic) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTopicInput(v *DeleteTopicInput, encoder *httpbinding.Encoder) error {
@@ -5565,10 +5405,8 @@ func (m *awsRestjson1_serializeOpDeleteTopicRefreshSch
 func (m *awsRestjson1_serializeOpDeleteTopicRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5606,7 +5444,6 @@ func (m *awsRestjson1_serializeOpDeleteTopicRefreshSch
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTopicRefreshScheduleInput(v *DeleteTopicRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -5654,10 +5491,8 @@ func (m *awsRestjson1_serializeOpDeleteUser) HandleSer
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
@@ -5695,7 +5530,6 @@ func (m *awsRestjson1_serializeOpDeleteUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserInput(v *DeleteUserInput, encoder *httpbinding.Encoder) error {
@@ -5743,10 +5577,8 @@ func (m *awsRestjson1_serializeOpDeleteUserByPrincipal
 func (m *awsRestjson1_serializeOpDeleteUserByPrincipalId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5784,7 +5616,6 @@ func (m *awsRestjson1_serializeOpDeleteUserByPrincipal
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserByPrincipalIdInput(v *DeleteUserByPrincipalIdInput, encoder *httpbinding.Encoder) error {
@@ -5832,10 +5663,8 @@ func (m *awsRestjson1_serializeOpDeleteUserCustomPermi
 func (m *awsRestjson1_serializeOpDeleteUserCustomPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5873,7 +5702,6 @@ func (m *awsRestjson1_serializeOpDeleteUserCustomPermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserCustomPermissionInput(v *DeleteUserCustomPermissionInput, encoder *httpbinding.Encoder) error {
@@ -5921,10 +5749,8 @@ func (m *awsRestjson1_serializeOpDeleteVPCConnection) 
 func (m *awsRestjson1_serializeOpDeleteVPCConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5962,7 +5788,6 @@ func (m *awsRestjson1_serializeOpDeleteVPCConnection) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVPCConnectionInput(v *DeleteVPCConnectionInput, encoder *httpbinding.Encoder) error {
@@ -6001,10 +5826,8 @@ func (m *awsRestjson1_serializeOpDescribeAccountCustom
 func (m *awsRestjson1_serializeOpDescribeAccountCustomization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6042,7 +5865,6 @@ func (m *awsRestjson1_serializeOpDescribeAccountCustom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccountCustomizationInput(v *DescribeAccountCustomizationInput, encoder *httpbinding.Encoder) error {
@@ -6080,10 +5902,8 @@ func (m *awsRestjson1_serializeOpDescribeAccountSettin
 func (m *awsRestjson1_serializeOpDescribeAccountSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6121,7 +5941,6 @@ func (m *awsRestjson1_serializeOpDescribeAccountSettin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccountSettingsInput(v *DescribeAccountSettingsInput, encoder *httpbinding.Encoder) error {
@@ -6151,10 +5970,8 @@ func (m *awsRestjson1_serializeOpDescribeAccountSubscr
 func (m *awsRestjson1_serializeOpDescribeAccountSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6192,7 +6009,6 @@ func (m *awsRestjson1_serializeOpDescribeAccountSubscr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccountSubscriptionInput(v *DescribeAccountSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -6222,10 +6038,8 @@ func (m *awsRestjson1_serializeOpDescribeAnalysis) Han
 func (m *awsRestjson1_serializeOpDescribeAnalysis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6263,7 +6077,6 @@ func (m *awsRestjson1_serializeOpDescribeAnalysis) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAnalysisInput(v *DescribeAnalysisInput, encoder *httpbinding.Encoder) error {
@@ -6302,10 +6115,8 @@ func (m *awsRestjson1_serializeOpDescribeAnalysisDefin
 func (m *awsRestjson1_serializeOpDescribeAnalysisDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6343,7 +6154,6 @@ func (m *awsRestjson1_serializeOpDescribeAnalysisDefin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAnalysisDefinitionInput(v *DescribeAnalysisDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -6382,10 +6192,8 @@ func (m *awsRestjson1_serializeOpDescribeAnalysisPermi
 func (m *awsRestjson1_serializeOpDescribeAnalysisPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6423,7 +6231,6 @@ func (m *awsRestjson1_serializeOpDescribeAnalysisPermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAnalysisPermissionsInput(v *DescribeAnalysisPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -6462,10 +6269,8 @@ func (m *awsRestjson1_serializeOpDescribeAssetBundleEx
 func (m *awsRestjson1_serializeOpDescribeAssetBundleExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6503,7 +6308,6 @@ func (m *awsRestjson1_serializeOpDescribeAssetBundleEx
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAssetBundleExportJobInput(v *DescribeAssetBundleExportJobInput, encoder *httpbinding.Encoder) error {
@@ -6542,10 +6346,8 @@ func (m *awsRestjson1_serializeOpDescribeAssetBundleIm
 func (m *awsRestjson1_serializeOpDescribeAssetBundleImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6583,7 +6385,6 @@ func (m *awsRestjson1_serializeOpDescribeAssetBundleIm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAssetBundleImportJobInput(v *DescribeAssetBundleImportJobInput, encoder *httpbinding.Encoder) error {
@@ -6622,10 +6423,8 @@ func (m *awsRestjson1_serializeOpDescribeBrand) Handle
 func (m *awsRestjson1_serializeOpDescribeBrand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6663,7 +6462,6 @@ func (m *awsRestjson1_serializeOpDescribeBrand) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBrandInput(v *DescribeBrandInput, encoder *httpbinding.Encoder) error {
@@ -6706,10 +6504,8 @@ func (m *awsRestjson1_serializeOpDescribeBrandAssignme
 func (m *awsRestjson1_serializeOpDescribeBrandAssignment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6747,7 +6543,6 @@ func (m *awsRestjson1_serializeOpDescribeBrandAssignme
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBrandAssignmentInput(v *DescribeBrandAssignmentInput, encoder *httpbinding.Encoder) error {
@@ -6777,10 +6572,8 @@ func (m *awsRestjson1_serializeOpDescribeBrandPublishe
 func (m *awsRestjson1_serializeOpDescribeBrandPublishedVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6818,7 +6611,6 @@ func (m *awsRestjson1_serializeOpDescribeBrandPublishe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBrandPublishedVersionInput(v *DescribeBrandPublishedVersionInput, encoder *httpbinding.Encoder) error {
@@ -6857,10 +6649,8 @@ func (m *awsRestjson1_serializeOpDescribeCustomPermiss
 func (m *awsRestjson1_serializeOpDescribeCustomPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6898,7 +6688,6 @@ func (m *awsRestjson1_serializeOpDescribeCustomPermiss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeCustomPermissionsInput(v *DescribeCustomPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -6937,10 +6726,8 @@ func (m *awsRestjson1_serializeOpDescribeDashboard) Ha
 func (m *awsRestjson1_serializeOpDescribeDashboard) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6978,7 +6765,6 @@ func (m *awsRestjson1_serializeOpDescribeDashboard) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDashboardInput(v *DescribeDashboardInput, encoder *httpbinding.Encoder) error {
@@ -7025,10 +6811,8 @@ func (m *awsRestjson1_serializeOpDescribeDashboardDefi
 func (m *awsRestjson1_serializeOpDescribeDashboardDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7066,7 +6850,6 @@ func (m *awsRestjson1_serializeOpDescribeDashboardDefi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDashboardDefinitionInput(v *DescribeDashboardDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -7113,10 +6896,8 @@ func (m *awsRestjson1_serializeOpDescribeDashboardPerm
 func (m *awsRestjson1_serializeOpDescribeDashboardPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7154,7 +6935,6 @@ func (m *awsRestjson1_serializeOpDescribeDashboardPerm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDashboardPermissionsInput(v *DescribeDashboardPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -7193,10 +6973,8 @@ func (m *awsRestjson1_serializeOpDescribeDashboardSnap
 func (m *awsRestjson1_serializeOpDescribeDashboardSnapshotJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7234,7 +7012,6 @@ func (m *awsRestjson1_serializeOpDescribeDashboardSnap
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDashboardSnapshotJobInput(v *DescribeDashboardSnapshotJobInput, encoder *httpbinding.Encoder) error {
@@ -7282,10 +7059,8 @@ func (m *awsRestjson1_serializeOpDescribeDashboardSnap
 func (m *awsRestjson1_serializeOpDescribeDashboardSnapshotJobResult) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7323,7 +7098,6 @@ func (m *awsRestjson1_serializeOpDescribeDashboardSnap
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDashboardSnapshotJobResultInput(v *DescribeDashboardSnapshotJobResultInput, encoder *httpbinding.Encoder) error {
@@ -7371,10 +7145,8 @@ func (m *awsRestjson1_serializeOpDescribeDashboardsQAC
 func (m *awsRestjson1_serializeOpDescribeDashboardsQAConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7412,7 +7184,6 @@ func (m *awsRestjson1_serializeOpDescribeDashboardsQAC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDashboardsQAConfigurationInput(v *DescribeDashboardsQAConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7442,10 +7213,8 @@ func (m *awsRestjson1_serializeOpDescribeDataSet) Hand
 func (m *awsRestjson1_serializeOpDescribeDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7483,7 +7252,6 @@ func (m *awsRestjson1_serializeOpDescribeDataSet) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDataSetInput(v *DescribeDataSetInput, encoder *httpbinding.Encoder) error {
@@ -7522,10 +7290,8 @@ func (m *awsRestjson1_serializeOpDescribeDataSetPermis
 func (m *awsRestjson1_serializeOpDescribeDataSetPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7563,7 +7329,6 @@ func (m *awsRestjson1_serializeOpDescribeDataSetPermis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDataSetPermissionsInput(v *DescribeDataSetPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -7602,10 +7367,8 @@ func (m *awsRestjson1_serializeOpDescribeDataSetRefres
 func (m *awsRestjson1_serializeOpDescribeDataSetRefreshProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7643,7 +7406,6 @@ func (m *awsRestjson1_serializeOpDescribeDataSetRefres
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDataSetRefreshPropertiesInput(v *DescribeDataSetRefreshPropertiesInput, encoder *httpbinding.Encoder) error {
@@ -7682,10 +7444,8 @@ func (m *awsRestjson1_serializeOpDescribeDataSource) H
 func (m *awsRestjson1_serializeOpDescribeDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7723,7 +7483,6 @@ func (m *awsRestjson1_serializeOpDescribeDataSource) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDataSourceInput(v *DescribeDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -7762,10 +7521,8 @@ func (m *awsRestjson1_serializeOpDescribeDataSourcePer
 func (m *awsRestjson1_serializeOpDescribeDataSourcePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7803,7 +7560,6 @@ func (m *awsRestjson1_serializeOpDescribeDataSourcePer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDataSourcePermissionsInput(v *DescribeDataSourcePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -7842,10 +7598,8 @@ func (m *awsRestjson1_serializeOpDescribeDefaultQBusin
 func (m *awsRestjson1_serializeOpDescribeDefaultQBusinessApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7883,7 +7637,6 @@ func (m *awsRestjson1_serializeOpDescribeDefaultQBusin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDefaultQBusinessApplicationInput(v *DescribeDefaultQBusinessApplicationInput, encoder *httpbinding.Encoder) error {
@@ -7917,10 +7670,8 @@ func (m *awsRestjson1_serializeOpDescribeFolder) Handl
 func (m *awsRestjson1_serializeOpDescribeFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7958,7 +7709,6 @@ func (m *awsRestjson1_serializeOpDescribeFolder) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeFolderInput(v *DescribeFolderInput, encoder *httpbinding.Encoder) error {
@@ -7997,10 +7747,8 @@ func (m *awsRestjson1_serializeOpDescribeFolderPermiss
 func (m *awsRestjson1_serializeOpDescribeFolderPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8038,7 +7786,6 @@ func (m *awsRestjson1_serializeOpDescribeFolderPermiss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeFolderPermissionsInput(v *DescribeFolderPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -8089,10 +7836,8 @@ func (m *awsRestjson1_serializeOpDescribeFolderResolve
 func (m *awsRestjson1_serializeOpDescribeFolderResolvedPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8130,7 +7875,6 @@ func (m *awsRestjson1_serializeOpDescribeFolderResolve
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeFolderResolvedPermissionsInput(v *DescribeFolderResolvedPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -8181,10 +7925,8 @@ func (m *awsRestjson1_serializeOpDescribeGroup) Handle
 func (m *awsRestjson1_serializeOpDescribeGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8222,7 +7964,6 @@ func (m *awsRestjson1_serializeOpDescribeGroup) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeGroupInput(v *DescribeGroupInput, encoder *httpbinding.Encoder) error {
@@ -8270,10 +8011,8 @@ func (m *awsRestjson1_serializeOpDescribeGroupMembersh
 func (m *awsRestjson1_serializeOpDescribeGroupMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8311,7 +8050,6 @@ func (m *awsRestjson1_serializeOpDescribeGroupMembersh
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeGroupMembershipInput(v *DescribeGroupMembershipInput, encoder *httpbinding.Encoder) error {
@@ -8368,10 +8106,8 @@ func (m *awsRestjson1_serializeOpDescribeIAMPolicyAssi
 func (m *awsRestjson1_serializeOpDescribeIAMPolicyAssignment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8409,7 +8145,6 @@ func (m *awsRestjson1_serializeOpDescribeIAMPolicyAssi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeIAMPolicyAssignmentInput(v *DescribeIAMPolicyAssignmentInput, encoder *httpbinding.Encoder) error {
@@ -8457,10 +8192,8 @@ func (m *awsRestjson1_serializeOpDescribeIngestion) Ha
 func (m *awsRestjson1_serializeOpDescribeIngestion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8498,7 +8231,6 @@ func (m *awsRestjson1_serializeOpDescribeIngestion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeIngestionInput(v *DescribeIngestionInput, encoder *httpbinding.Encoder) error {
@@ -8546,10 +8278,8 @@ func (m *awsRestjson1_serializeOpDescribeIpRestriction
 func (m *awsRestjson1_serializeOpDescribeIpRestriction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8587,7 +8317,6 @@ func (m *awsRestjson1_serializeOpDescribeIpRestriction
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeIpRestrictionInput(v *DescribeIpRestrictionInput, encoder *httpbinding.Encoder) error {
@@ -8617,10 +8346,8 @@ func (m *awsRestjson1_serializeOpDescribeKeyRegistrati
 func (m *awsRestjson1_serializeOpDescribeKeyRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8658,7 +8385,6 @@ func (m *awsRestjson1_serializeOpDescribeKeyRegistrati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeKeyRegistrationInput(v *DescribeKeyRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -8692,10 +8418,8 @@ func (m *awsRestjson1_serializeOpDescribeNamespace) Ha
 func (m *awsRestjson1_serializeOpDescribeNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8733,7 +8457,6 @@ func (m *awsRestjson1_serializeOpDescribeNamespace) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeNamespaceInput(v *DescribeNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -8772,10 +8495,8 @@ func (m *awsRestjson1_serializeOpDescribeQPersonalizat
 func (m *awsRestjson1_serializeOpDescribeQPersonalizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8813,7 +8534,6 @@ func (m *awsRestjson1_serializeOpDescribeQPersonalizat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeQPersonalizationConfigurationInput(v *DescribeQPersonalizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8843,10 +8563,8 @@ func (m *awsRestjson1_serializeOpDescribeQuickSightQSe
 func (m *awsRestjson1_serializeOpDescribeQuickSightQSearchConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8884,7 +8602,6 @@ func (m *awsRestjson1_serializeOpDescribeQuickSightQSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeQuickSightQSearchConfigurationInput(v *DescribeQuickSightQSearchConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8914,10 +8631,8 @@ func (m *awsRestjson1_serializeOpDescribeRefreshSchedu
 func (m *awsRestjson1_serializeOpDescribeRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8955,7 +8670,6 @@ func (m *awsRestjson1_serializeOpDescribeRefreshSchedu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRefreshScheduleInput(v *DescribeRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -9003,10 +8717,8 @@ func (m *awsRestjson1_serializeOpDescribeRoleCustomPer
 func (m *awsRestjson1_serializeOpDescribeRoleCustomPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9044,7 +8756,6 @@ func (m *awsRestjson1_serializeOpDescribeRoleCustomPer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRoleCustomPermissionInput(v *DescribeRoleCustomPermissionInput, encoder *httpbinding.Encoder) error {
@@ -9092,10 +8803,8 @@ func (m *awsRestjson1_serializeOpDescribeTemplate) Han
 func (m *awsRestjson1_serializeOpDescribeTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9133,7 +8842,6 @@ func (m *awsRestjson1_serializeOpDescribeTemplate) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTemplateInput(v *DescribeTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9180,10 +8888,8 @@ func (m *awsRestjson1_serializeOpDescribeTemplateAlias
 func (m *awsRestjson1_serializeOpDescribeTemplateAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9221,7 +8927,6 @@ func (m *awsRestjson1_serializeOpDescribeTemplateAlias
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTemplateAliasInput(v *DescribeTemplateAliasInput, encoder *httpbinding.Encoder) error {
@@ -9269,10 +8974,8 @@ func (m *awsRestjson1_serializeOpDescribeTemplateDefin
 func (m *awsRestjson1_serializeOpDescribeTemplateDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9310,7 +9013,6 @@ func (m *awsRestjson1_serializeOpDescribeTemplateDefin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTemplateDefinitionInput(v *DescribeTemplateDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -9357,10 +9059,8 @@ func (m *awsRestjson1_serializeOpDescribeTemplatePermi
 func (m *awsRestjson1_serializeOpDescribeTemplatePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9398,7 +9098,6 @@ func (m *awsRestjson1_serializeOpDescribeTemplatePermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTemplatePermissionsInput(v *DescribeTemplatePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -9437,10 +9136,8 @@ func (m *awsRestjson1_serializeOpDescribeTheme) Handle
 func (m *awsRestjson1_serializeOpDescribeTheme) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9478,7 +9175,6 @@ func (m *awsRestjson1_serializeOpDescribeTheme) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThemeInput(v *DescribeThemeInput, encoder *httpbinding.Encoder) error {
@@ -9525,10 +9221,8 @@ func (m *awsRestjson1_serializeOpDescribeThemeAlias) H
 func (m *awsRestjson1_serializeOpDescribeThemeAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9566,7 +9260,6 @@ func (m *awsRestjson1_serializeOpDescribeThemeAlias) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThemeAliasInput(v *DescribeThemeAliasInput, encoder *httpbinding.Encoder) error {
@@ -9614,10 +9307,8 @@ func (m *awsRestjson1_serializeOpDescribeThemePermissi
 func (m *awsRestjson1_serializeOpDescribeThemePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9655,7 +9346,6 @@ func (m *awsRestjson1_serializeOpDescribeThemePermissi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThemePermissionsInput(v *DescribeThemePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -9694,10 +9384,8 @@ func (m *awsRestjson1_serializeOpDescribeTopic) Handle
 func (m *awsRestjson1_serializeOpDescribeTopic) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9735,7 +9423,6 @@ func (m *awsRestjson1_serializeOpDescribeTopic) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTopicInput(v *DescribeTopicInput, encoder *httpbinding.Encoder) error {
@@ -9774,10 +9461,8 @@ func (m *awsRestjson1_serializeOpDescribeTopicPermissi
 func (m *awsRestjson1_serializeOpDescribeTopicPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9815,7 +9500,6 @@ func (m *awsRestjson1_serializeOpDescribeTopicPermissi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTopicPermissionsInput(v *DescribeTopicPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -9854,10 +9538,8 @@ func (m *awsRestjson1_serializeOpDescribeTopicRefresh)
 func (m *awsRestjson1_serializeOpDescribeTopicRefresh) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9895,7 +9577,6 @@ func (m *awsRestjson1_serializeOpDescribeTopicRefresh)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTopicRefreshInput(v *DescribeTopicRefreshInput, encoder *httpbinding.Encoder) error {
@@ -9943,10 +9624,8 @@ func (m *awsRestjson1_serializeOpDescribeTopicRefreshS
 func (m *awsRestjson1_serializeOpDescribeTopicRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9984,7 +9663,6 @@ func (m *awsRestjson1_serializeOpDescribeTopicRefreshS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTopicRefreshScheduleInput(v *DescribeTopicRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -10032,10 +9710,8 @@ func (m *awsRestjson1_serializeOpDescribeUser) HandleS
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
@@ -10073,7 +9749,6 @@ func (m *awsRestjson1_serializeOpDescribeUser) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeUserInput(v *DescribeUserInput, encoder *httpbinding.Encoder) error {
@@ -10121,10 +9796,8 @@ func (m *awsRestjson1_serializeOpDescribeVPCConnection
 func (m *awsRestjson1_serializeOpDescribeVPCConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10162,7 +9835,6 @@ func (m *awsRestjson1_serializeOpDescribeVPCConnection
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeVPCConnectionInput(v *DescribeVPCConnectionInput, encoder *httpbinding.Encoder) error {
@@ -10201,10 +9873,8 @@ func (m *awsRestjson1_serializeOpGenerateEmbedUrlForAn
 func (m *awsRestjson1_serializeOpGenerateEmbedUrlForAnonymousUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10253,7 +9923,6 @@ func (m *awsRestjson1_serializeOpGenerateEmbedUrlForAn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateEmbedUrlForAnonymousUserInput(v *GenerateEmbedUrlForAnonymousUserInput, encoder *httpbinding.Encoder) error {
@@ -10328,10 +9997,8 @@ func (m *awsRestjson1_serializeOpGenerateEmbedUrlForRe
 func (m *awsRestjson1_serializeOpGenerateEmbedUrlForRegisteredUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10380,7 +10047,6 @@ func (m *awsRestjson1_serializeOpGenerateEmbedUrlForRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateEmbedUrlForRegisteredUserInput(v *GenerateEmbedUrlForRegisteredUserInput, encoder *httpbinding.Encoder) error {
@@ -10441,10 +10107,8 @@ func (m *awsRestjson1_serializeOpGenerateEmbedUrlForRe
 func (m *awsRestjson1_serializeOpGenerateEmbedUrlForRegisteredUserWithIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10493,7 +10157,6 @@ func (m *awsRestjson1_serializeOpGenerateEmbedUrlForRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateEmbedUrlForRegisteredUserWithIdentityInput(v *GenerateEmbedUrlForRegisteredUserWithIdentityInput, encoder *httpbinding.Encoder) error {
@@ -10549,10 +10212,8 @@ func (m *awsRestjson1_serializeOpGetDashboardEmbedUrl)
 func (m *awsRestjson1_serializeOpGetDashboardEmbedUrl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10590,7 +10251,6 @@ func (m *awsRestjson1_serializeOpGetDashboardEmbedUrl)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDashboardEmbedUrlInput(v *GetDashboardEmbedUrlInput, encoder *httpbinding.Encoder) error {
@@ -10663,10 +10323,8 @@ func (m *awsRestjson1_serializeOpGetSessionEmbedUrl) H
 func (m *awsRestjson1_serializeOpGetSessionEmbedUrl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10704,7 +10362,6 @@ func (m *awsRestjson1_serializeOpGetSessionEmbedUrl) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSessionEmbedUrlInput(v *GetSessionEmbedUrlInput, encoder *httpbinding.Encoder) error {
@@ -10746,10 +10403,8 @@ func (m *awsRestjson1_serializeOpListAnalyses) HandleS
 func (m *awsRestjson1_serializeOpListAnalyses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10787,7 +10442,6 @@ func (m *awsRestjson1_serializeOpListAnalyses) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnalysesInput(v *ListAnalysesInput, encoder *httpbinding.Encoder) error {
@@ -10825,10 +10479,8 @@ func (m *awsRestjson1_serializeOpListAssetBundleExport
 func (m *awsRestjson1_serializeOpListAssetBundleExportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10866,7 +10518,6 @@ func (m *awsRestjson1_serializeOpListAssetBundleExport
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetBundleExportJobsInput(v *ListAssetBundleExportJobsInput, encoder *httpbinding.Encoder) error {
@@ -10904,10 +10555,8 @@ func (m *awsRestjson1_serializeOpListAssetBundleImport
 func (m *awsRestjson1_serializeOpListAssetBundleImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10945,7 +10594,6 @@ func (m *awsRestjson1_serializeOpListAssetBundleImport
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetBundleImportJobsInput(v *ListAssetBundleImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -10983,10 +10631,8 @@ func (m *awsRestjson1_serializeOpListBrands) HandleSer
 func (m *awsRestjson1_serializeOpListBrands) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11024,7 +10670,6 @@ func (m *awsRestjson1_serializeOpListBrands) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBrandsInput(v *ListBrandsInput, encoder *httpbinding.Encoder) error {
@@ -11062,10 +10707,8 @@ func (m *awsRestjson1_serializeOpListCustomPermissions
 func (m *awsRestjson1_serializeOpListCustomPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11103,7 +10746,6 @@ func (m *awsRestjson1_serializeOpListCustomPermissions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomPermissionsInput(v *ListCustomPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -11141,10 +10783,8 @@ func (m *awsRestjson1_serializeOpListDashboards) Handl
 func (m *awsRestjson1_serializeOpListDashboards) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11182,7 +10822,6 @@ func (m *awsRestjson1_serializeOpListDashboards) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDashboardsInput(v *ListDashboardsInput, encoder *httpbinding.Encoder) error {
@@ -11220,10 +10859,8 @@ func (m *awsRestjson1_serializeOpListDashboardVersions
 func (m *awsRestjson1_serializeOpListDashboardVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11261,7 +10898,6 @@ func (m *awsRestjson1_serializeOpListDashboardVersions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDashboardVersionsInput(v *ListDashboardVersionsInput, encoder *httpbinding.Encoder) error {
@@ -11308,10 +10944,8 @@ func (m *awsRestjson1_serializeOpListDataSets) HandleS
 func (m *awsRestjson1_serializeOpListDataSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11349,7 +10983,6 @@ func (m *awsRestjson1_serializeOpListDataSets) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataSetsInput(v *ListDataSetsInput, encoder *httpbinding.Encoder) error {
@@ -11387,10 +11020,8 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 func (m *awsRestjson1_serializeOpListDataSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11428,7 +11059,6 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataSourcesInput(v *ListDataSourcesInput, encoder *httpbinding.Encoder) error {
@@ -11466,10 +11096,8 @@ func (m *awsRestjson1_serializeOpListFolderMembers) Ha
 func (m *awsRestjson1_serializeOpListFolderMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11507,7 +11135,6 @@ func (m *awsRestjson1_serializeOpListFolderMembers) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFolderMembersInput(v *ListFolderMembersInput, encoder *httpbinding.Encoder) error {
@@ -11554,10 +11181,8 @@ func (m *awsRestjson1_serializeOpListFolders) HandleSe
 func (m *awsRestjson1_serializeOpListFolders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11595,7 +11220,6 @@ func (m *awsRestjson1_serializeOpListFolders) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFoldersInput(v *ListFoldersInput, encoder *httpbinding.Encoder) error {
@@ -11633,10 +11257,8 @@ func (m *awsRestjson1_serializeOpListFoldersForResourc
 func (m *awsRestjson1_serializeOpListFoldersForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11674,7 +11296,6 @@ func (m *awsRestjson1_serializeOpListFoldersForResourc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFoldersForResourceInput(v *ListFoldersForResourceInput, encoder *httpbinding.Encoder) error {
@@ -11721,10 +11342,8 @@ func (m *awsRestjson1_serializeOpListGroupMemberships)
 func (m *awsRestjson1_serializeOpListGroupMemberships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11762,7 +11381,6 @@ func (m *awsRestjson1_serializeOpListGroupMemberships)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGroupMembershipsInput(v *ListGroupMembershipsInput, encoder *httpbinding.Encoder) error {
@@ -11818,10 +11436,8 @@ func (m *awsRestjson1_serializeOpListGroups) HandleSer
 func (m *awsRestjson1_serializeOpListGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11859,7 +11475,6 @@ func (m *awsRestjson1_serializeOpListGroups) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGroupsInput(v *ListGroupsInput, encoder *httpbinding.Encoder) error {
@@ -11906,10 +11521,8 @@ func (m *awsRestjson1_serializeOpListIAMPolicyAssignme
 func (m *awsRestjson1_serializeOpListIAMPolicyAssignments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11947,7 +11560,6 @@ func (m *awsRestjson1_serializeOpListIAMPolicyAssignme
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIAMPolicyAssignmentsInput(v *ListIAMPolicyAssignmentsInput, encoder *httpbinding.Encoder) error {
@@ -11998,10 +11610,8 @@ func (m *awsRestjson1_serializeOpListIAMPolicyAssignme
 func (m *awsRestjson1_serializeOpListIAMPolicyAssignmentsForUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12039,7 +11649,6 @@ func (m *awsRestjson1_serializeOpListIAMPolicyAssignme
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIAMPolicyAssignmentsForUserInput(v *ListIAMPolicyAssignmentsForUserInput, encoder *httpbinding.Encoder) error {
@@ -12095,10 +11704,8 @@ func (m *awsRestjson1_serializeOpListIdentityPropagati
 func (m *awsRestjson1_serializeOpListIdentityPropagationConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12136,7 +11743,6 @@ func (m *awsRestjson1_serializeOpListIdentityPropagati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIdentityPropagationConfigsInput(v *ListIdentityPropagationConfigsInput, encoder *httpbinding.Encoder) error {
@@ -12174,10 +11780,8 @@ func (m *awsRestjson1_serializeOpListIngestions) Handl
 func (m *awsRestjson1_serializeOpListIngestions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12215,7 +11819,6 @@ func (m *awsRestjson1_serializeOpListIngestions) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIngestionsInput(v *ListIngestionsInput, encoder *httpbinding.Encoder) error {
@@ -12262,10 +11865,8 @@ func (m *awsRestjson1_serializeOpListNamespaces) Handl
 func (m *awsRestjson1_serializeOpListNamespaces) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12303,7 +11904,6 @@ func (m *awsRestjson1_serializeOpListNamespaces) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNamespacesInput(v *ListNamespacesInput, encoder *httpbinding.Encoder) error {
@@ -12341,10 +11941,8 @@ func (m *awsRestjson1_serializeOpListRefreshSchedules)
 func (m *awsRestjson1_serializeOpListRefreshSchedules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12382,7 +11980,6 @@ func (m *awsRestjson1_serializeOpListRefreshSchedules)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRefreshSchedulesInput(v *ListRefreshSchedulesInput, encoder *httpbinding.Encoder) error {
@@ -12421,10 +12018,8 @@ func (m *awsRestjson1_serializeOpListRoleMemberships) 
 func (m *awsRestjson1_serializeOpListRoleMemberships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12462,7 +12057,6 @@ func (m *awsRestjson1_serializeOpListRoleMemberships) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRoleMembershipsInput(v *ListRoleMembershipsInput, encoder *httpbinding.Encoder) error {
@@ -12518,10 +12112,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -12559,7 +12151,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -12589,10 +12180,8 @@ func (m *awsRestjson1_serializeOpListTemplateAliases) 
 func (m *awsRestjson1_serializeOpListTemplateAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12630,7 +12219,6 @@ func (m *awsRestjson1_serializeOpListTemplateAliases) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTemplateAliasesInput(v *ListTemplateAliasesInput, encoder *httpbinding.Encoder) error {
@@ -12677,10 +12265,8 @@ func (m *awsRestjson1_serializeOpListTemplates) Handle
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
@@ -12718,7 +12304,6 @@ func (m *awsRestjson1_serializeOpListTemplates) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTemplatesInput(v *ListTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -12756,10 +12341,8 @@ func (m *awsRestjson1_serializeOpListTemplateVersions)
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
@@ -12797,7 +12380,6 @@ func (m *awsRestjson1_serializeOpListTemplateVersions)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTemplateVersionsInput(v *ListTemplateVersionsInput, encoder *httpbinding.Encoder) error {
@@ -12844,10 +12426,8 @@ func (m *awsRestjson1_serializeOpListThemeAliases) Han
 func (m *awsRestjson1_serializeOpListThemeAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12885,7 +12465,6 @@ func (m *awsRestjson1_serializeOpListThemeAliases) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThemeAliasesInput(v *ListThemeAliasesInput, encoder *httpbinding.Encoder) error {
@@ -12932,10 +12511,8 @@ func (m *awsRestjson1_serializeOpListThemes) HandleSer
 func (m *awsRestjson1_serializeOpListThemes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12973,7 +12550,6 @@ func (m *awsRestjson1_serializeOpListThemes) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThemesInput(v *ListThemesInput, encoder *httpbinding.Encoder) error {
@@ -13015,10 +12591,8 @@ func (m *awsRestjson1_serializeOpListThemeVersions) Ha
 func (m *awsRestjson1_serializeOpListThemeVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13056,7 +12630,6 @@ func (m *awsRestjson1_serializeOpListThemeVersions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThemeVersionsInput(v *ListThemeVersionsInput, encoder *httpbinding.Encoder) error {
@@ -13103,10 +12676,8 @@ func (m *awsRestjson1_serializeOpListTopicRefreshSched
 func (m *awsRestjson1_serializeOpListTopicRefreshSchedules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13144,7 +12715,6 @@ func (m *awsRestjson1_serializeOpListTopicRefreshSched
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTopicRefreshSchedulesInput(v *ListTopicRefreshSchedulesInput, encoder *httpbinding.Encoder) error {
@@ -13183,10 +12753,8 @@ func (m *awsRestjson1_serializeOpListTopicReviewedAnsw
 func (m *awsRestjson1_serializeOpListTopicReviewedAnswers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13224,7 +12792,6 @@ func (m *awsRestjson1_serializeOpListTopicReviewedAnsw
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTopicReviewedAnswersInput(v *ListTopicReviewedAnswersInput, encoder *httpbinding.Encoder) error {
@@ -13263,10 +12830,8 @@ func (m *awsRestjson1_serializeOpListTopics) HandleSer
 func (m *awsRestjson1_serializeOpListTopics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13304,7 +12869,6 @@ func (m *awsRestjson1_serializeOpListTopics) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTopicsInput(v *ListTopicsInput, encoder *httpbinding.Encoder) error {
@@ -13342,10 +12906,8 @@ func (m *awsRestjson1_serializeOpListUserGroups) Handl
 func (m *awsRestjson1_serializeOpListUserGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13383,7 +12945,6 @@ func (m *awsRestjson1_serializeOpListUserGroups) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUserGroupsInput(v *ListUserGroupsInput, encoder *httpbinding.Encoder) error {
@@ -13439,10 +13000,8 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
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
@@ -13480,7 +13039,6 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUsersInput(v *ListUsersInput, encoder *httpbinding.Encoder) error {
@@ -13527,10 +13085,8 @@ func (m *awsRestjson1_serializeOpListVPCConnections) H
 func (m *awsRestjson1_serializeOpListVPCConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13568,7 +13124,6 @@ func (m *awsRestjson1_serializeOpListVPCConnections) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVPCConnectionsInput(v *ListVPCConnectionsInput, encoder *httpbinding.Encoder) error {
@@ -13606,10 +13161,8 @@ func (m *awsRestjson1_serializeOpPredictQAResults) Han
 func (m *awsRestjson1_serializeOpPredictQAResults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13658,7 +13211,6 @@ func (m *awsRestjson1_serializeOpPredictQAResults) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPredictQAResultsInput(v *PredictQAResultsInput, encoder *httpbinding.Encoder) error {
@@ -13715,10 +13267,8 @@ func (m *awsRestjson1_serializeOpPutDataSetRefreshProp
 func (m *awsRestjson1_serializeOpPutDataSetRefreshProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13767,7 +13317,6 @@ func (m *awsRestjson1_serializeOpPutDataSetRefreshProp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDataSetRefreshPropertiesInput(v *PutDataSetRefreshPropertiesInput, encoder *httpbinding.Encoder) error {
@@ -13820,10 +13369,8 @@ func (m *awsRestjson1_serializeOpRegisterUser) HandleS
 func (m *awsRestjson1_serializeOpRegisterUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13872,7 +13419,6 @@ func (m *awsRestjson1_serializeOpRegisterUser) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterUserInput(v *RegisterUserInput, encoder *httpbinding.Encoder) error {
@@ -13975,10 +13521,8 @@ func (m *awsRestjson1_serializeOpRestoreAnalysis) Hand
 func (m *awsRestjson1_serializeOpRestoreAnalysis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14016,7 +13560,6 @@ func (m *awsRestjson1_serializeOpRestoreAnalysis) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRestoreAnalysisInput(v *RestoreAnalysisInput, encoder *httpbinding.Encoder) error {
@@ -14059,10 +13602,8 @@ func (m *awsRestjson1_serializeOpSearchAnalyses) Handl
 func (m *awsRestjson1_serializeOpSearchAnalyses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14111,7 +13652,6 @@ func (m *awsRestjson1_serializeOpSearchAnalyses) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchAnalysesInput(v *SearchAnalysesInput, encoder *httpbinding.Encoder) error {
@@ -14165,10 +13705,8 @@ func (m *awsRestjson1_serializeOpSearchDashboards) Han
 func (m *awsRestjson1_serializeOpSearchDashboards) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14217,7 +13755,6 @@ func (m *awsRestjson1_serializeOpSearchDashboards) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchDashboardsInput(v *SearchDashboardsInput, encoder *httpbinding.Encoder) error {
@@ -14271,10 +13808,8 @@ func (m *awsRestjson1_serializeOpSearchDataSets) Handl
 func (m *awsRestjson1_serializeOpSearchDataSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14323,7 +13858,6 @@ func (m *awsRestjson1_serializeOpSearchDataSets) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchDataSetsInput(v *SearchDataSetsInput, encoder *httpbinding.Encoder) error {
@@ -14377,10 +13911,8 @@ func (m *awsRestjson1_serializeOpSearchDataSources) Ha
 func (m *awsRestjson1_serializeOpSearchDataSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14429,7 +13961,6 @@ func (m *awsRestjson1_serializeOpSearchDataSources) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchDataSourcesInput(v *SearchDataSourcesInput, encoder *httpbinding.Encoder) error {
@@ -14483,10 +14014,8 @@ func (m *awsRestjson1_serializeOpSearchFolders) Handle
 func (m *awsRestjson1_serializeOpSearchFolders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14535,7 +14064,6 @@ func (m *awsRestjson1_serializeOpSearchFolders) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchFoldersInput(v *SearchFoldersInput, encoder *httpbinding.Encoder) error {
@@ -14589,10 +14117,8 @@ func (m *awsRestjson1_serializeOpSearchGroups) HandleS
 func (m *awsRestjson1_serializeOpSearchGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14641,7 +14167,6 @@ func (m *awsRestjson1_serializeOpSearchGroups) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchGroupsInput(v *SearchGroupsInput, encoder *httpbinding.Encoder) error {
@@ -14702,10 +14227,8 @@ func (m *awsRestjson1_serializeOpSearchTopics) HandleS
 func (m *awsRestjson1_serializeOpSearchTopics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14754,7 +14277,6 @@ func (m *awsRestjson1_serializeOpSearchTopics) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchTopicsInput(v *SearchTopicsInput, encoder *httpbinding.Encoder) error {
@@ -14808,10 +14330,8 @@ func (m *awsRestjson1_serializeOpStartAssetBundleExpor
 func (m *awsRestjson1_serializeOpStartAssetBundleExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14860,7 +14380,6 @@ func (m *awsRestjson1_serializeOpStartAssetBundleExpor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAssetBundleExportJobInput(v *StartAssetBundleExportJobInput, encoder *httpbinding.Encoder) error {
@@ -14953,10 +14472,8 @@ func (m *awsRestjson1_serializeOpStartAssetBundleImpor
 func (m *awsRestjson1_serializeOpStartAssetBundleImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15005,7 +14522,6 @@ func (m *awsRestjson1_serializeOpStartAssetBundleImpor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAssetBundleImportJobInput(v *StartAssetBundleImportJobInput, encoder *httpbinding.Encoder) error {
@@ -15087,10 +14603,8 @@ func (m *awsRestjson1_serializeOpStartDashboardSnapsho
 func (m *awsRestjson1_serializeOpStartDashboardSnapshotJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15139,7 +14653,6 @@ func (m *awsRestjson1_serializeOpStartDashboardSnapsho
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDashboardSnapshotJobInput(v *StartDashboardSnapshotJobInput, encoder *httpbinding.Encoder) error {
@@ -15204,10 +14717,8 @@ func (m *awsRestjson1_serializeOpStartDashboardSnapsho
 func (m *awsRestjson1_serializeOpStartDashboardSnapshotJobSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15245,7 +14756,6 @@ func (m *awsRestjson1_serializeOpStartDashboardSnapsho
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDashboardSnapshotJobScheduleInput(v *StartDashboardSnapshotJobScheduleInput, encoder *httpbinding.Encoder) error {
@@ -15293,10 +14803,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -15345,7 +14853,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -15389,10 +14896,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -15430,7 +14935,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -15466,10 +14970,8 @@ func (m *awsRestjson1_serializeOpUpdateAccountCustomiz
 func (m *awsRestjson1_serializeOpUpdateAccountCustomization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15518,7 +15020,6 @@ func (m *awsRestjson1_serializeOpUpdateAccountCustomiz
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountCustomizationInput(v *UpdateAccountCustomizationInput, encoder *httpbinding.Encoder) error {
@@ -15566,10 +15067,8 @@ func (m *awsRestjson1_serializeOpUpdateAccountSettings
 func (m *awsRestjson1_serializeOpUpdateAccountSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15618,7 +15117,6 @@ func (m *awsRestjson1_serializeOpUpdateAccountSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountSettingsInput(v *UpdateAccountSettingsInput, encoder *httpbinding.Encoder) error {
@@ -15670,10 +15168,8 @@ func (m *awsRestjson1_serializeOpUpdateAnalysis) Handl
 func (m *awsRestjson1_serializeOpUpdateAnalysis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15722,7 +15218,6 @@ func (m *awsRestjson1_serializeOpUpdateAnalysis) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAnalysisInput(v *UpdateAnalysisInput, encoder *httpbinding.Encoder) error {
@@ -15806,10 +15301,8 @@ func (m *awsRestjson1_serializeOpUpdateAnalysisPermiss
 func (m *awsRestjson1_serializeOpUpdateAnalysisPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15858,7 +15351,6 @@ func (m *awsRestjson1_serializeOpUpdateAnalysisPermiss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAnalysisPermissionsInput(v *UpdateAnalysisPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -15918,10 +15410,8 @@ func (m *awsRestjson1_serializeOpUpdateApplicationWith
 func (m *awsRestjson1_serializeOpUpdateApplicationWithTokenExchangeGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15959,7 +15449,6 @@ func (m *awsRestjson1_serializeOpUpdateApplicationWith
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApplicationWithTokenExchangeGrantInput(v *UpdateApplicationWithTokenExchangeGrantInput, encoder *httpbinding.Encoder) error {
@@ -15993,10 +15482,8 @@ func (m *awsRestjson1_serializeOpUpdateBrand) HandleSe
 func (m *awsRestjson1_serializeOpUpdateBrand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16045,7 +15532,6 @@ func (m *awsRestjson1_serializeOpUpdateBrand) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBrandInput(v *UpdateBrandInput, encoder *httpbinding.Encoder) error {
@@ -16098,10 +15584,8 @@ func (m *awsRestjson1_serializeOpUpdateBrandAssignment
 func (m *awsRestjson1_serializeOpUpdateBrandAssignment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16150,7 +15634,6 @@ func (m *awsRestjson1_serializeOpUpdateBrandAssignment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBrandAssignmentInput(v *UpdateBrandAssignmentInput, encoder *httpbinding.Encoder) error {
@@ -16192,10 +15675,8 @@ func (m *awsRestjson1_serializeOpUpdateBrandPublishedV
 func (m *awsRestjson1_serializeOpUpdateBrandPublishedVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16244,7 +15725,6 @@ func (m *awsRestjson1_serializeOpUpdateBrandPublishedV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBrandPublishedVersionInput(v *UpdateBrandPublishedVersionInput, encoder *httpbinding.Encoder) error {
@@ -16295,10 +15775,8 @@ func (m *awsRestjson1_serializeOpUpdateCustomPermissio
 func (m *awsRestjson1_serializeOpUpdateCustomPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16347,7 +15825,6 @@ func (m *awsRestjson1_serializeOpUpdateCustomPermissio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCustomPermissionsInput(v *UpdateCustomPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -16400,10 +15877,8 @@ func (m *awsRestjson1_serializeOpUpdateDashboard) Hand
 func (m *awsRestjson1_serializeOpUpdateDashboard) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16452,7 +15927,6 @@ func (m *awsRestjson1_serializeOpUpdateDashboard) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDashboardInput(v *UpdateDashboardInput, encoder *httpbinding.Encoder) error {
@@ -16548,10 +16022,8 @@ func (m *awsRestjson1_serializeOpUpdateDashboardLinks)
 func (m *awsRestjson1_serializeOpUpdateDashboardLinks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16600,7 +16072,6 @@ func (m *awsRestjson1_serializeOpUpdateDashboardLinks)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDashboardLinksInput(v *UpdateDashboardLinksInput, encoder *httpbinding.Encoder) error {
@@ -16653,10 +16124,8 @@ func (m *awsRestjson1_serializeOpUpdateDashboardPermis
 func (m *awsRestjson1_serializeOpUpdateDashboardPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16705,7 +16174,6 @@ func (m *awsRestjson1_serializeOpUpdateDashboardPermis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDashboardPermissionsInput(v *UpdateDashboardPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -16779,10 +16247,8 @@ func (m *awsRestjson1_serializeOpUpdateDashboardPublis
 func (m *awsRestjson1_serializeOpUpdateDashboardPublishedVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16820,7 +16286,6 @@ func (m *awsRestjson1_serializeOpUpdateDashboardPublis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDashboardPublishedVersionInput(v *UpdateDashboardPublishedVersionInput, encoder *httpbinding.Encoder) error {
@@ -16868,10 +16333,8 @@ func (m *awsRestjson1_serializeOpUpdateDashboardsQACon
 func (m *awsRestjson1_serializeOpUpdateDashboardsQAConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16920,7 +16383,6 @@ func (m *awsRestjson1_serializeOpUpdateDashboardsQACon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDashboardsQAConfigurationInput(v *UpdateDashboardsQAConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -16962,10 +16424,8 @@ func (m *awsRestjson1_serializeOpUpdateDataSet) Handle
 func (m *awsRestjson1_serializeOpUpdateDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17014,7 +16474,6 @@ func (m *awsRestjson1_serializeOpUpdateDataSet) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataSetInput(v *UpdateDataSetInput, encoder *httpbinding.Encoder) error {
@@ -17140,10 +16599,8 @@ func (m *awsRestjson1_serializeOpUpdateDataSetPermissi
 func (m *awsRestjson1_serializeOpUpdateDataSetPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17192,7 +16649,6 @@ func (m *awsRestjson1_serializeOpUpdateDataSetPermissi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataSetPermissionsInput(v *UpdateDataSetPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -17252,10 +16708,8 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 func (m *awsRestjson1_serializeOpUpdateDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17304,7 +16758,6 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataSourceInput(v *UpdateDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -17383,10 +16836,8 @@ func (m *awsRestjson1_serializeOpUpdateDataSourcePermi
 func (m *awsRestjson1_serializeOpUpdateDataSourcePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17435,7 +16886,6 @@ func (m *awsRestjson1_serializeOpUpdateDataSourcePermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataSourcePermissionsInput(v *UpdateDataSourcePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -17495,10 +16945,8 @@ func (m *awsRestjson1_serializeOpUpdateDefaultQBusines
 func (m *awsRestjson1_serializeOpUpdateDefaultQBusinessApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17547,7 +16995,6 @@ func (m *awsRestjson1_serializeOpUpdateDefaultQBusines
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDefaultQBusinessApplicationInput(v *UpdateDefaultQBusinessApplicationInput, encoder *httpbinding.Encoder) error {
@@ -17593,10 +17040,8 @@ func (m *awsRestjson1_serializeOpUpdateFolder) HandleS
 func (m *awsRestjson1_serializeOpUpdateFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17645,7 +17090,6 @@ func (m *awsRestjson1_serializeOpUpdateFolder) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFolderInput(v *UpdateFolderInput, encoder *httpbinding.Encoder) error {
@@ -17696,10 +17140,8 @@ func (m *awsRestjson1_serializeOpUpdateFolderPermissio
 func (m *awsRestjson1_serializeOpUpdateFolderPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17748,7 +17190,6 @@ func (m *awsRestjson1_serializeOpUpdateFolderPermissio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFolderPermissionsInput(v *UpdateFolderPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -17808,10 +17249,8 @@ func (m *awsRestjson1_serializeOpUpdateGroup) HandleSe
 func (m *awsRestjson1_serializeOpUpdateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17860,7 +17299,6 @@ func (m *awsRestjson1_serializeOpUpdateGroup) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGroupInput(v *UpdateGroupInput, encoder *httpbinding.Encoder) error {
@@ -17920,10 +17358,8 @@ func (m *awsRestjson1_serializeOpUpdateIAMPolicyAssign
 func (m *awsRestjson1_serializeOpUpdateIAMPolicyAssignment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17972,7 +17408,6 @@ func (m *awsRestjson1_serializeOpUpdateIAMPolicyAssign
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIAMPolicyAssignmentInput(v *UpdateIAMPolicyAssignmentInput, encoder *httpbinding.Encoder) error {
@@ -18044,10 +17479,8 @@ func (m *awsRestjson1_serializeOpUpdateIdentityPropaga
 func (m *awsRestjson1_serializeOpUpdateIdentityPropagationConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18096,7 +17529,6 @@ func (m *awsRestjson1_serializeOpUpdateIdentityPropaga
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIdentityPropagationConfigInput(v *UpdateIdentityPropagationConfigInput, encoder *httpbinding.Encoder) error {
@@ -18149,10 +17581,8 @@ func (m *awsRestjson1_serializeOpUpdateIpRestriction) 
 func (m *awsRestjson1_serializeOpUpdateIpRestriction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18201,7 +17631,6 @@ func (m *awsRestjson1_serializeOpUpdateIpRestriction) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIpRestrictionInput(v *UpdateIpRestrictionInput, encoder *httpbinding.Encoder) error {
@@ -18264,10 +17693,8 @@ func (m *awsRestjson1_serializeOpUpdateKeyRegistration
 func (m *awsRestjson1_serializeOpUpdateKeyRegistration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18316,7 +17743,6 @@ func (m *awsRestjson1_serializeOpUpdateKeyRegistration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateKeyRegistrationInput(v *UpdateKeyRegistrationInput, encoder *httpbinding.Encoder) error {
@@ -18360,10 +17786,8 @@ func (m *awsRestjson1_serializeOpUpdatePublicSharingSe
 func (m *awsRestjson1_serializeOpUpdatePublicSharingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18412,7 +17836,6 @@ func (m *awsRestjson1_serializeOpUpdatePublicSharingSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePublicSharingSettingsInput(v *UpdatePublicSharingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -18454,10 +17877,8 @@ func (m *awsRestjson1_serializeOpUpdateQPersonalizatio
 func (m *awsRestjson1_serializeOpUpdateQPersonalizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18506,7 +17927,6 @@ func (m *awsRestjson1_serializeOpUpdateQPersonalizatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQPersonalizationConfigurationInput(v *UpdateQPersonalizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -18548,10 +17968,8 @@ func (m *awsRestjson1_serializeOpUpdateQuickSightQSear
 func (m *awsRestjson1_serializeOpUpdateQuickSightQSearchConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18600,7 +18018,6 @@ func (m *awsRestjson1_serializeOpUpdateQuickSightQSear
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQuickSightQSearchConfigurationInput(v *UpdateQuickSightQSearchConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -18642,10 +18059,8 @@ func (m *awsRestjson1_serializeOpUpdateRefreshSchedule
 func (m *awsRestjson1_serializeOpUpdateRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18694,7 +18109,6 @@ func (m *awsRestjson1_serializeOpUpdateRefreshSchedule
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRefreshScheduleInput(v *UpdateRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -18747,10 +18161,8 @@ func (m *awsRestjson1_serializeOpUpdateRoleCustomPermi
 func (m *awsRestjson1_serializeOpUpdateRoleCustomPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18799,7 +18211,6 @@ func (m *awsRestjson1_serializeOpUpdateRoleCustomPermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoleCustomPermissionInput(v *UpdateRoleCustomPermissionInput, encoder *httpbinding.Encoder) error {
@@ -18859,10 +18270,8 @@ func (m *awsRestjson1_serializeOpUpdateSPICECapacityCo
 func (m *awsRestjson1_serializeOpUpdateSPICECapacityConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18911,7 +18320,6 @@ func (m *awsRestjson1_serializeOpUpdateSPICECapacityCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSPICECapacityConfigurationInput(v *UpdateSPICECapacityConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -18953,10 +18361,8 @@ func (m *awsRestjson1_serializeOpUpdateTemplate) Handl
 func (m *awsRestjson1_serializeOpUpdateTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19005,7 +18411,6 @@ func (m *awsRestjson1_serializeOpUpdateTemplate) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTemplateInput(v *UpdateTemplateInput, encoder *httpbinding.Encoder) error {
@@ -19082,10 +18487,8 @@ func (m *awsRestjson1_serializeOpUpdateTemplateAlias) 
 func (m *awsRestjson1_serializeOpUpdateTemplateAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19134,7 +18537,6 @@ func (m *awsRestjson1_serializeOpUpdateTemplateAlias) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTemplateAliasInput(v *UpdateTemplateAliasInput, encoder *httpbinding.Encoder) error {
@@ -19194,10 +18596,8 @@ func (m *awsRestjson1_serializeOpUpdateTemplatePermiss
 func (m *awsRestjson1_serializeOpUpdateTemplatePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19246,7 +18646,6 @@ func (m *awsRestjson1_serializeOpUpdateTemplatePermiss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTemplatePermissionsInput(v *UpdateTemplatePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -19306,10 +18705,8 @@ func (m *awsRestjson1_serializeOpUpdateTheme) HandleSe
 func (m *awsRestjson1_serializeOpUpdateTheme) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19358,7 +18755,6 @@ func (m *awsRestjson1_serializeOpUpdateTheme) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThemeInput(v *UpdateThemeInput, encoder *httpbinding.Encoder) error {
@@ -19426,10 +18822,8 @@ func (m *awsRestjson1_serializeOpUpdateThemeAlias) Han
 func (m *awsRestjson1_serializeOpUpdateThemeAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19478,7 +18872,6 @@ func (m *awsRestjson1_serializeOpUpdateThemeAlias) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThemeAliasInput(v *UpdateThemeAliasInput, encoder *httpbinding.Encoder) error {
@@ -19538,10 +18931,8 @@ func (m *awsRestjson1_serializeOpUpdateThemePermission
 func (m *awsRestjson1_serializeOpUpdateThemePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19590,7 +18981,6 @@ func (m *awsRestjson1_serializeOpUpdateThemePermission
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThemePermissionsInput(v *UpdateThemePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -19650,10 +19040,8 @@ func (m *awsRestjson1_serializeOpUpdateTopic) HandleSe
 func (m *awsRestjson1_serializeOpUpdateTopic) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19702,7 +19090,6 @@ func (m *awsRestjson1_serializeOpUpdateTopic) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTopicInput(v *UpdateTopicInput, encoder *httpbinding.Encoder) error {
@@ -19762,10 +19149,8 @@ func (m *awsRestjson1_serializeOpUpdateTopicPermission
 func (m *awsRestjson1_serializeOpUpdateTopicPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19814,7 +19199,6 @@ func (m *awsRestjson1_serializeOpUpdateTopicPermission
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTopicPermissionsInput(v *UpdateTopicPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -19874,10 +19258,8 @@ func (m *awsRestjson1_serializeOpUpdateTopicRefreshSch
 func (m *awsRestjson1_serializeOpUpdateTopicRefreshSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19926,7 +19308,6 @@ func (m *awsRestjson1_serializeOpUpdateTopicRefreshSch
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTopicRefreshScheduleInput(v *UpdateTopicRefreshScheduleInput, encoder *httpbinding.Encoder) error {
@@ -19988,10 +19369,8 @@ func (m *awsRestjson1_serializeOpUpdateUser) HandleSer
 func (m *awsRestjson1_serializeOpUpdateUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20040,7 +19419,6 @@ func (m *awsRestjson1_serializeOpUpdateUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserInput(v *UpdateUserInput, encoder *httpbinding.Encoder) error {
@@ -20130,10 +19508,8 @@ func (m *awsRestjson1_serializeOpUpdateUserCustomPermi
 func (m *awsRestjson1_serializeOpUpdateUserCustomPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20182,7 +19558,6 @@ func (m *awsRestjson1_serializeOpUpdateUserCustomPermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserCustomPermissionInput(v *UpdateUserCustomPermissionInput, encoder *httpbinding.Encoder) error {
@@ -20242,10 +19617,8 @@ func (m *awsRestjson1_serializeOpUpdateVPCConnection) 
 func (m *awsRestjson1_serializeOpUpdateVPCConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20294,7 +19667,6 @@ func (m *awsRestjson1_serializeOpUpdateVPCConnection) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVPCConnectionInput(v *UpdateVPCConnectionInput, encoder *httpbinding.Encoder) error {
