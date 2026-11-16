--- vendor/github.com/aws/aws-sdk-go-v2/service/macie2/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/macie2/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAcceptInvitation) Han
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
@@ -73,7 +70,6 @@ func (m *awsRestjson1_serializeOpAcceptInvitation) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptInvitationInput(v *AcceptInvitationInput, encoder *httpbinding.Encoder) error {
@@ -116,10 +112,8 @@ func (m *awsRestjson1_serializeOpBatchGetCustomDataIde
 func (m *awsRestjson1_serializeOpBatchGetCustomDataIdentifiers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -164,7 +158,6 @@ func (m *awsRestjson1_serializeOpBatchGetCustomDataIde
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetCustomDataIdentifiersInput(v *BatchGetCustomDataIdentifiersInput, encoder *httpbinding.Encoder) error {
@@ -199,10 +192,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateAutomatedD
 func (m *awsRestjson1_serializeOpBatchUpdateAutomatedDiscoveryAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -247,7 +238,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateAutomatedD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateAutomatedDiscoveryAccountsInput(v *BatchUpdateAutomatedDiscoveryAccountsInput, encoder *httpbinding.Encoder) error {
@@ -282,10 +272,8 @@ func (m *awsRestjson1_serializeOpCreateAllowList) Hand
 func (m *awsRestjson1_serializeOpCreateAllowList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -330,7 +318,6 @@ func (m *awsRestjson1_serializeOpCreateAllowList) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAllowListInput(v *CreateAllowListInput, encoder *httpbinding.Encoder) error {
@@ -387,10 +374,8 @@ func (m *awsRestjson1_serializeOpCreateClassificationJ
 func (m *awsRestjson1_serializeOpCreateClassificationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -435,7 +420,6 @@ func (m *awsRestjson1_serializeOpCreateClassificationJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateClassificationJobInput(v *CreateClassificationJobInput, encoder *httpbinding.Encoder) error {
@@ -540,10 +524,8 @@ func (m *awsRestjson1_serializeOpCreateCustomDataIdent
 func (m *awsRestjson1_serializeOpCreateCustomDataIdentifier) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -588,7 +570,6 @@ func (m *awsRestjson1_serializeOpCreateCustomDataIdent
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCustomDataIdentifierInput(v *CreateCustomDataIdentifierInput, encoder *httpbinding.Encoder) error {
@@ -669,10 +650,8 @@ func (m *awsRestjson1_serializeOpCreateFindingsFilter)
 func (m *awsRestjson1_serializeOpCreateFindingsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -717,7 +696,6 @@ func (m *awsRestjson1_serializeOpCreateFindingsFilter)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFindingsFilterInput(v *CreateFindingsFilterInput, encoder *httpbinding.Encoder) error {
@@ -784,10 +762,8 @@ func (m *awsRestjson1_serializeOpCreateInvitations) Ha
 func (m *awsRestjson1_serializeOpCreateInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -832,7 +808,6 @@ func (m *awsRestjson1_serializeOpCreateInvitations) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInvitationsInput(v *CreateInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -877,10 +852,8 @@ func (m *awsRestjson1_serializeOpCreateMember) HandleS
 func (m *awsRestjson1_serializeOpCreateMember) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -925,7 +898,6 @@ func (m *awsRestjson1_serializeOpCreateMember) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMemberInput(v *CreateMemberInput, encoder *httpbinding.Encoder) error {
@@ -967,10 +939,8 @@ func (m *awsRestjson1_serializeOpCreateSampleFindings)
 func (m *awsRestjson1_serializeOpCreateSampleFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1015,7 +985,6 @@ func (m *awsRestjson1_serializeOpCreateSampleFindings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSampleFindingsInput(v *CreateSampleFindingsInput, encoder *httpbinding.Encoder) error {
@@ -1050,10 +1019,8 @@ func (m *awsRestjson1_serializeOpDeclineInvitations) H
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
@@ -1098,7 +1065,6 @@ func (m *awsRestjson1_serializeOpDeclineInvitations) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeclineInvitationsInput(v *DeclineInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -1133,10 +1099,8 @@ func (m *awsRestjson1_serializeOpDeleteAllowList) Hand
 func (m *awsRestjson1_serializeOpDeleteAllowList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1174,7 +1138,6 @@ func (m *awsRestjson1_serializeOpDeleteAllowList) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAllowListInput(v *DeleteAllowListInput, encoder *httpbinding.Encoder) error {
@@ -1208,10 +1171,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomDataIdent
 func (m *awsRestjson1_serializeOpDeleteCustomDataIdentifier) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1249,7 +1210,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomDataIdent
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomDataIdentifierInput(v *DeleteCustomDataIdentifierInput, encoder *httpbinding.Encoder) error {
@@ -1279,10 +1239,8 @@ func (m *awsRestjson1_serializeOpDeleteFindingsFilter)
 func (m *awsRestjson1_serializeOpDeleteFindingsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1320,7 +1278,6 @@ func (m *awsRestjson1_serializeOpDeleteFindingsFilter)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFindingsFilterInput(v *DeleteFindingsFilterInput, encoder *httpbinding.Encoder) error {
@@ -1350,10 +1307,8 @@ func (m *awsRestjson1_serializeOpDeleteInvitations) Ha
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
@@ -1398,7 +1353,6 @@ func (m *awsRestjson1_serializeOpDeleteInvitations) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInvitationsInput(v *DeleteInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -1433,10 +1387,8 @@ func (m *awsRestjson1_serializeOpDeleteMember) HandleS
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
@@ -1474,7 +1426,6 @@ func (m *awsRestjson1_serializeOpDeleteMember) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMemberInput(v *DeleteMemberInput, encoder *httpbinding.Encoder) error {
@@ -1504,10 +1455,8 @@ func (m *awsRestjson1_serializeOpDescribeBuckets) Hand
 func (m *awsRestjson1_serializeOpDescribeBuckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1552,7 +1501,6 @@ func (m *awsRestjson1_serializeOpDescribeBuckets) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBucketsInput(v *DescribeBucketsInput, encoder *httpbinding.Encoder) error {
@@ -1604,10 +1552,8 @@ func (m *awsRestjson1_serializeOpDescribeClassificatio
 func (m *awsRestjson1_serializeOpDescribeClassificationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1645,7 +1591,6 @@ func (m *awsRestjson1_serializeOpDescribeClassificatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeClassificationJobInput(v *DescribeClassificationJobInput, encoder *httpbinding.Encoder) error {
@@ -1675,10 +1620,8 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
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
@@ -1712,7 +1655,6 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOrganizationConfigurationInput(v *DescribeOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1733,10 +1675,8 @@ func (m *awsRestjson1_serializeOpDisableMacie) HandleS
 func (m *awsRestjson1_serializeOpDisableMacie) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1770,7 +1710,6 @@ func (m *awsRestjson1_serializeOpDisableMacie) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableMacieInput(v *DisableMacieInput, encoder *httpbinding.Encoder) error {
@@ -1791,10 +1730,8 @@ func (m *awsRestjson1_serializeOpDisableOrganizationAd
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
@@ -1832,7 +1769,6 @@ func (m *awsRestjson1_serializeOpDisableOrganizationAd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableOrganizationAdminAccountInput(v *DisableOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -1857,10 +1793,8 @@ func (m *awsRestjson1_serializeOpDisassociateFromAdmin
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
@@ -1894,7 +1828,6 @@ func (m *awsRestjson1_serializeOpDisassociateFromAdmin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFromAdministratorAccountInput(v *DisassociateFromAdministratorAccountInput, encoder *httpbinding.Encoder) error {
@@ -1915,10 +1848,8 @@ func (m *awsRestjson1_serializeOpDisassociateFromMaste
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
@@ -1952,7 +1883,6 @@ func (m *awsRestjson1_serializeOpDisassociateFromMaste
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFromMasterAccountInput(v *DisassociateFromMasterAccountInput, encoder *httpbinding.Encoder) error {
@@ -1973,10 +1903,8 @@ func (m *awsRestjson1_serializeOpDisassociateMember) H
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
@@ -2014,7 +1942,6 @@ func (m *awsRestjson1_serializeOpDisassociateMember) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMemberInput(v *DisassociateMemberInput, encoder *httpbinding.Encoder) error {
@@ -2044,10 +1971,8 @@ func (m *awsRestjson1_serializeOpEnableMacie) HandleSe
 func (m *awsRestjson1_serializeOpEnableMacie) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2092,7 +2017,6 @@ func (m *awsRestjson1_serializeOpEnableMacie) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableMacieInput(v *EnableMacieInput, encoder *httpbinding.Encoder) error {
@@ -2135,10 +2059,8 @@ func (m *awsRestjson1_serializeOpEnableOrganizationAdm
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
@@ -2183,7 +2105,6 @@ func (m *awsRestjson1_serializeOpEnableOrganizationAdm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableOrganizationAdminAccountInput(v *EnableOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -2221,10 +2142,8 @@ func (m *awsRestjson1_serializeOpGetAdministratorAccou
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
@@ -2258,7 +2177,6 @@ func (m *awsRestjson1_serializeOpGetAdministratorAccou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAdministratorAccountInput(v *GetAdministratorAccountInput, encoder *httpbinding.Encoder) error {
@@ -2279,10 +2197,8 @@ func (m *awsRestjson1_serializeOpGetAllowList) HandleS
 func (m *awsRestjson1_serializeOpGetAllowList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2320,7 +2236,6 @@ func (m *awsRestjson1_serializeOpGetAllowList) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAllowListInput(v *GetAllowListInput, encoder *httpbinding.Encoder) error {
@@ -2350,10 +2265,8 @@ func (m *awsRestjson1_serializeOpGetAutomatedDiscovery
 func (m *awsRestjson1_serializeOpGetAutomatedDiscoveryConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2387,7 +2300,6 @@ func (m *awsRestjson1_serializeOpGetAutomatedDiscovery
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAutomatedDiscoveryConfigurationInput(v *GetAutomatedDiscoveryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2408,10 +2320,8 @@ func (m *awsRestjson1_serializeOpGetBucketStatistics) 
 func (m *awsRestjson1_serializeOpGetBucketStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2456,7 +2366,6 @@ func (m *awsRestjson1_serializeOpGetBucketStatistics) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBucketStatisticsInput(v *GetBucketStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -2489,10 +2398,8 @@ func (m *awsRestjson1_serializeOpGetClassificationExpo
 func (m *awsRestjson1_serializeOpGetClassificationExportConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2526,7 +2433,6 @@ func (m *awsRestjson1_serializeOpGetClassificationExpo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetClassificationExportConfigurationInput(v *GetClassificationExportConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2547,10 +2453,8 @@ func (m *awsRestjson1_serializeOpGetClassificationScop
 func (m *awsRestjson1_serializeOpGetClassificationScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2588,7 +2492,6 @@ func (m *awsRestjson1_serializeOpGetClassificationScop
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetClassificationScopeInput(v *GetClassificationScopeInput, encoder *httpbinding.Encoder) error {
@@ -2618,10 +2521,8 @@ func (m *awsRestjson1_serializeOpGetCustomDataIdentifi
 func (m *awsRestjson1_serializeOpGetCustomDataIdentifier) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2659,7 +2560,6 @@ func (m *awsRestjson1_serializeOpGetCustomDataIdentifi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCustomDataIdentifierInput(v *GetCustomDataIdentifierInput, encoder *httpbinding.Encoder) error {
@@ -2689,10 +2589,8 @@ func (m *awsRestjson1_serializeOpGetFindings) HandleSe
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
@@ -2737,7 +2635,6 @@ func (m *awsRestjson1_serializeOpGetFindings) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsInput(v *GetFindingsInput, encoder *httpbinding.Encoder) error {
@@ -2779,10 +2676,8 @@ func (m *awsRestjson1_serializeOpGetFindingsFilter) Ha
 func (m *awsRestjson1_serializeOpGetFindingsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2820,7 +2715,6 @@ func (m *awsRestjson1_serializeOpGetFindingsFilter) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsFilterInput(v *GetFindingsFilterInput, encoder *httpbinding.Encoder) error {
@@ -2850,10 +2744,8 @@ func (m *awsRestjson1_serializeOpGetFindingsPublicatio
 func (m *awsRestjson1_serializeOpGetFindingsPublicationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2887,7 +2779,6 @@ func (m *awsRestjson1_serializeOpGetFindingsPublicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsPublicationConfigurationInput(v *GetFindingsPublicationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2908,10 +2799,8 @@ func (m *awsRestjson1_serializeOpGetFindingStatistics)
 func (m *awsRestjson1_serializeOpGetFindingStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2956,7 +2845,6 @@ func (m *awsRestjson1_serializeOpGetFindingStatistics)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingStatisticsInput(v *GetFindingStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -3008,10 +2896,8 @@ func (m *awsRestjson1_serializeOpGetInvitationsCount) 
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
@@ -3045,7 +2931,6 @@ func (m *awsRestjson1_serializeOpGetInvitationsCount) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInvitationsCountInput(v *GetInvitationsCountInput, encoder *httpbinding.Encoder) error {
@@ -3066,10 +2951,8 @@ func (m *awsRestjson1_serializeOpGetMacieSession) Hand
 func (m *awsRestjson1_serializeOpGetMacieSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3103,7 +2986,6 @@ func (m *awsRestjson1_serializeOpGetMacieSession) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMacieSessionInput(v *GetMacieSessionInput, encoder *httpbinding.Encoder) error {
@@ -3124,10 +3006,8 @@ func (m *awsRestjson1_serializeOpGetMasterAccount) Han
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
@@ -3161,7 +3041,6 @@ func (m *awsRestjson1_serializeOpGetMasterAccount) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMasterAccountInput(v *GetMasterAccountInput, encoder *httpbinding.Encoder) error {
@@ -3182,10 +3061,8 @@ func (m *awsRestjson1_serializeOpGetMember) HandleSeri
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
@@ -3223,7 +3100,6 @@ func (m *awsRestjson1_serializeOpGetMember) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMemberInput(v *GetMemberInput, encoder *httpbinding.Encoder) error {
@@ -3253,10 +3129,8 @@ func (m *awsRestjson1_serializeOpGetResourceProfile) H
 func (m *awsRestjson1_serializeOpGetResourceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3294,7 +3168,6 @@ func (m *awsRestjson1_serializeOpGetResourceProfile) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceProfileInput(v *GetResourceProfileInput, encoder *httpbinding.Encoder) error {
@@ -3319,10 +3192,8 @@ func (m *awsRestjson1_serializeOpGetRevealConfiguratio
 func (m *awsRestjson1_serializeOpGetRevealConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3356,7 +3227,6 @@ func (m *awsRestjson1_serializeOpGetRevealConfiguratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRevealConfigurationInput(v *GetRevealConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3377,10 +3247,8 @@ func (m *awsRestjson1_serializeOpGetSensitiveDataOccur
 func (m *awsRestjson1_serializeOpGetSensitiveDataOccurrences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3418,7 +3286,6 @@ func (m *awsRestjson1_serializeOpGetSensitiveDataOccur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSensitiveDataOccurrencesInput(v *GetSensitiveDataOccurrencesInput, encoder *httpbinding.Encoder) error {
@@ -3448,10 +3315,8 @@ func (m *awsRestjson1_serializeOpGetSensitiveDataOccur
 func (m *awsRestjson1_serializeOpGetSensitiveDataOccurrencesAvailability) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3489,7 +3354,6 @@ func (m *awsRestjson1_serializeOpGetSensitiveDataOccur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSensitiveDataOccurrencesAvailabilityInput(v *GetSensitiveDataOccurrencesAvailabilityInput, encoder *httpbinding.Encoder) error {
@@ -3519,10 +3383,8 @@ func (m *awsRestjson1_serializeOpGetSensitivityInspect
 func (m *awsRestjson1_serializeOpGetSensitivityInspectionTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3560,7 +3422,6 @@ func (m *awsRestjson1_serializeOpGetSensitivityInspect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSensitivityInspectionTemplateInput(v *GetSensitivityInspectionTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3590,10 +3451,8 @@ func (m *awsRestjson1_serializeOpGetUsageStatistics) H
 func (m *awsRestjson1_serializeOpGetUsageStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3638,7 +3497,6 @@ func (m *awsRestjson1_serializeOpGetUsageStatistics) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsageStatisticsInput(v *GetUsageStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -3695,10 +3553,8 @@ func (m *awsRestjson1_serializeOpGetUsageTotals) Handl
 func (m *awsRestjson1_serializeOpGetUsageTotals) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3736,7 +3592,6 @@ func (m *awsRestjson1_serializeOpGetUsageTotals) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsageTotalsInput(v *GetUsageTotalsInput, encoder *httpbinding.Encoder) error {
@@ -3761,10 +3616,8 @@ func (m *awsRestjson1_serializeOpListAllowLists) Handl
 func (m *awsRestjson1_serializeOpListAllowLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3802,7 +3655,6 @@ func (m *awsRestjson1_serializeOpListAllowLists) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAllowListsInput(v *ListAllowListsInput, encoder *httpbinding.Encoder) error {
@@ -3831,10 +3683,8 @@ func (m *awsRestjson1_serializeOpListAutomatedDiscover
 func (m *awsRestjson1_serializeOpListAutomatedDiscoveryAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3872,7 +3722,6 @@ func (m *awsRestjson1_serializeOpListAutomatedDiscover
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAutomatedDiscoveryAccountsInput(v *ListAutomatedDiscoveryAccountsInput, encoder *httpbinding.Encoder) error {
@@ -3907,10 +3756,8 @@ func (m *awsRestjson1_serializeOpListClassificationJob
 func (m *awsRestjson1_serializeOpListClassificationJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3955,7 +3802,6 @@ func (m *awsRestjson1_serializeOpListClassificationJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListClassificationJobsInput(v *ListClassificationJobsInput, encoder *httpbinding.Encoder) error {
@@ -4007,10 +3853,8 @@ func (m *awsRestjson1_serializeOpListClassificationSco
 func (m *awsRestjson1_serializeOpListClassificationScopes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4048,7 +3892,6 @@ func (m *awsRestjson1_serializeOpListClassificationSco
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListClassificationScopesInput(v *ListClassificationScopesInput, encoder *httpbinding.Encoder) error {
@@ -4077,10 +3920,8 @@ func (m *awsRestjson1_serializeOpListCustomDataIdentif
 func (m *awsRestjson1_serializeOpListCustomDataIdentifiers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4125,7 +3966,6 @@ func (m *awsRestjson1_serializeOpListCustomDataIdentif
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomDataIdentifiersInput(v *ListCustomDataIdentifiersInput, encoder *httpbinding.Encoder) error {
@@ -4163,10 +4003,8 @@ func (m *awsRestjson1_serializeOpListFindings) HandleS
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
@@ -4211,7 +4049,6 @@ func (m *awsRestjson1_serializeOpListFindings) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFindingsInput(v *ListFindingsInput, encoder *httpbinding.Encoder) error {
@@ -4263,10 +4100,8 @@ func (m *awsRestjson1_serializeOpListFindingsFilters) 
 func (m *awsRestjson1_serializeOpListFindingsFilters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4304,7 +4139,6 @@ func (m *awsRestjson1_serializeOpListFindingsFilters) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFindingsFiltersInput(v *ListFindingsFiltersInput, encoder *httpbinding.Encoder) error {
@@ -4333,10 +4167,8 @@ func (m *awsRestjson1_serializeOpListInvitations) Hand
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
@@ -4374,7 +4206,6 @@ func (m *awsRestjson1_serializeOpListInvitations) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInvitationsInput(v *ListInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -4403,10 +4234,8 @@ func (m *awsRestjson1_serializeOpListManagedDataIdenti
 func (m *awsRestjson1_serializeOpListManagedDataIdentifiers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4451,7 +4280,6 @@ func (m *awsRestjson1_serializeOpListManagedDataIdenti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListManagedDataIdentifiersInput(v *ListManagedDataIdentifiersInput, encoder *httpbinding.Encoder) error {
@@ -4484,10 +4312,8 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
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
@@ -4525,7 +4351,6 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMembersInput(v *ListMembersInput, encoder *httpbinding.Encoder) error {
@@ -4558,10 +4383,8 @@ func (m *awsRestjson1_serializeOpListOrganizationAdmin
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
@@ -4599,7 +4422,6 @@ func (m *awsRestjson1_serializeOpListOrganizationAdmin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationAdminAccountsInput(v *ListOrganizationAdminAccountsInput, encoder *httpbinding.Encoder) error {
@@ -4628,10 +4450,8 @@ func (m *awsRestjson1_serializeOpListResourceProfileAr
 func (m *awsRestjson1_serializeOpListResourceProfileArtifacts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4669,7 +4489,6 @@ func (m *awsRestjson1_serializeOpListResourceProfileAr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceProfileArtifactsInput(v *ListResourceProfileArtifactsInput, encoder *httpbinding.Encoder) error {
@@ -4698,10 +4517,8 @@ func (m *awsRestjson1_serializeOpListResourceProfileDe
 func (m *awsRestjson1_serializeOpListResourceProfileDetections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4739,7 +4556,6 @@ func (m *awsRestjson1_serializeOpListResourceProfileDe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceProfileDetectionsInput(v *ListResourceProfileDetectionsInput, encoder *httpbinding.Encoder) error {
@@ -4772,10 +4588,8 @@ func (m *awsRestjson1_serializeOpListSensitivityInspec
 func (m *awsRestjson1_serializeOpListSensitivityInspectionTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4813,7 +4627,6 @@ func (m *awsRestjson1_serializeOpListSensitivityInspec
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSensitivityInspectionTemplatesInput(v *ListSensitivityInspectionTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -4842,10 +4655,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -4883,7 +4694,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4913,10 +4723,8 @@ func (m *awsRestjson1_serializeOpPutClassificationExpo
 func (m *awsRestjson1_serializeOpPutClassificationExportConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4961,7 +4769,6 @@ func (m *awsRestjson1_serializeOpPutClassificationExpo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutClassificationExportConfigurationInput(v *PutClassificationExportConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4996,10 +4803,8 @@ func (m *awsRestjson1_serializeOpPutFindingsPublicatio
 func (m *awsRestjson1_serializeOpPutFindingsPublicationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5044,7 +4849,6 @@ func (m *awsRestjson1_serializeOpPutFindingsPublicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutFindingsPublicationConfigurationInput(v *PutFindingsPublicationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5084,10 +4888,8 @@ func (m *awsRestjson1_serializeOpSearchResources) Hand
 func (m *awsRestjson1_serializeOpSearchResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5132,7 +4934,6 @@ func (m *awsRestjson1_serializeOpSearchResources) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchResourcesInput(v *SearchResourcesInput, encoder *httpbinding.Encoder) error {
@@ -5184,10 +4985,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5236,7 +5035,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5280,10 +5078,8 @@ func (m *awsRestjson1_serializeOpTestCustomDataIdentif
 func (m *awsRestjson1_serializeOpTestCustomDataIdentifier) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5328,7 +5124,6 @@ func (m *awsRestjson1_serializeOpTestCustomDataIdentif
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestCustomDataIdentifierInput(v *TestCustomDataIdentifierInput, encoder *httpbinding.Encoder) error {
@@ -5385,10 +5180,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5426,7 +5219,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5462,10 +5254,8 @@ func (m *awsRestjson1_serializeOpUpdateAllowList) Hand
 func (m *awsRestjson1_serializeOpUpdateAllowList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5514,7 +5304,6 @@ func (m *awsRestjson1_serializeOpUpdateAllowList) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAllowListInput(v *UpdateAllowListInput, encoder *httpbinding.Encoder) error {
@@ -5568,10 +5357,8 @@ func (m *awsRestjson1_serializeOpUpdateAutomatedDiscov
 func (m *awsRestjson1_serializeOpUpdateAutomatedDiscoveryConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5616,7 +5403,6 @@ func (m *awsRestjson1_serializeOpUpdateAutomatedDiscov
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAutomatedDiscoveryConfigurationInput(v *UpdateAutomatedDiscoveryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5654,10 +5440,8 @@ func (m *awsRestjson1_serializeOpUpdateClassificationJ
 func (m *awsRestjson1_serializeOpUpdateClassificationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5706,7 +5490,6 @@ func (m *awsRestjson1_serializeOpUpdateClassificationJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateClassificationJobInput(v *UpdateClassificationJobInput, encoder *httpbinding.Encoder) error {
@@ -5748,10 +5531,8 @@ func (m *awsRestjson1_serializeOpUpdateClassificationS
 func (m *awsRestjson1_serializeOpUpdateClassificationScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5800,7 +5581,6 @@ func (m *awsRestjson1_serializeOpUpdateClassificationS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateClassificationScopeInput(v *UpdateClassificationScopeInput, encoder *httpbinding.Encoder) error {
@@ -5844,10 +5624,8 @@ func (m *awsRestjson1_serializeOpUpdateFindingsFilter)
 func (m *awsRestjson1_serializeOpUpdateFindingsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5896,7 +5674,6 @@ func (m *awsRestjson1_serializeOpUpdateFindingsFilter)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFindingsFilterInput(v *UpdateFindingsFilterInput, encoder *httpbinding.Encoder) error {
@@ -5965,10 +5742,8 @@ func (m *awsRestjson1_serializeOpUpdateMacieSession) H
 func (m *awsRestjson1_serializeOpUpdateMacieSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6013,7 +5788,6 @@ func (m *awsRestjson1_serializeOpUpdateMacieSession) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMacieSessionInput(v *UpdateMacieSessionInput, encoder *httpbinding.Encoder) error {
@@ -6051,10 +5825,8 @@ func (m *awsRestjson1_serializeOpUpdateMemberSession) 
 func (m *awsRestjson1_serializeOpUpdateMemberSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6103,7 +5875,6 @@ func (m *awsRestjson1_serializeOpUpdateMemberSession) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMemberSessionInput(v *UpdateMemberSessionInput, encoder *httpbinding.Encoder) error {
@@ -6145,10 +5916,8 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
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
@@ -6193,7 +5962,6 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOrganizationConfigurationInput(v *UpdateOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6226,10 +5994,8 @@ func (m *awsRestjson1_serializeOpUpdateResourceProfile
 func (m *awsRestjson1_serializeOpUpdateResourceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6278,7 +6044,6 @@ func (m *awsRestjson1_serializeOpUpdateResourceProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceProfileInput(v *UpdateResourceProfileInput, encoder *httpbinding.Encoder) error {
@@ -6315,10 +6080,8 @@ func (m *awsRestjson1_serializeOpUpdateResourceProfile
 func (m *awsRestjson1_serializeOpUpdateResourceProfileDetections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6367,7 +6130,6 @@ func (m *awsRestjson1_serializeOpUpdateResourceProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceProfileDetectionsInput(v *UpdateResourceProfileDetectionsInput, encoder *httpbinding.Encoder) error {
@@ -6406,10 +6168,8 @@ func (m *awsRestjson1_serializeOpUpdateRevealConfigura
 func (m *awsRestjson1_serializeOpUpdateRevealConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6454,7 +6214,6 @@ func (m *awsRestjson1_serializeOpUpdateRevealConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRevealConfigurationInput(v *UpdateRevealConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6496,10 +6255,8 @@ func (m *awsRestjson1_serializeOpUpdateSensitivityInsp
 func (m *awsRestjson1_serializeOpUpdateSensitivityInspectionTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6548,7 +6305,6 @@ func (m *awsRestjson1_serializeOpUpdateSensitivityInsp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSensitivityInspectionTemplateInput(v *UpdateSensitivityInspectionTemplateInput, encoder *httpbinding.Encoder) error {
