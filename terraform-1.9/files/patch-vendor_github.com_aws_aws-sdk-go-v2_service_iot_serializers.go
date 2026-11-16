--- vendor/github.com/aws/aws-sdk-go-v2/service/iot/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iot/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAcceptCertificateTran
 func (m *awsRestjson1_serializeOpAcceptCertificateTransfer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -68,7 +65,6 @@ func (m *awsRestjson1_serializeOpAcceptCertificateTran
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptCertificateTransferInput(v *AcceptCertificateTransferInput, encoder *httpbinding.Encoder) error {
@@ -102,10 +98,8 @@ func (m *awsRestjson1_serializeOpAddThingToBillingGrou
 func (m *awsRestjson1_serializeOpAddThingToBillingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -150,7 +144,6 @@ func (m *awsRestjson1_serializeOpAddThingToBillingGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddThingToBillingGroupInput(v *AddThingToBillingGroupInput, encoder *httpbinding.Encoder) error {
@@ -198,10 +191,8 @@ func (m *awsRestjson1_serializeOpAddThingToThingGroup)
 func (m *awsRestjson1_serializeOpAddThingToThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -246,7 +237,6 @@ func (m *awsRestjson1_serializeOpAddThingToThingGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddThingToThingGroupInput(v *AddThingToThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -299,10 +289,8 @@ func (m *awsRestjson1_serializeOpAssociateSbomWithPack
 func (m *awsRestjson1_serializeOpAssociateSbomWithPackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -351,7 +339,6 @@ func (m *awsRestjson1_serializeOpAssociateSbomWithPack
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateSbomWithPackageVersionInput(v *AssociateSbomWithPackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -408,10 +395,8 @@ func (m *awsRestjson1_serializeOpAssociateTargetsWithJ
 func (m *awsRestjson1_serializeOpAssociateTargetsWithJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -460,7 +445,6 @@ func (m *awsRestjson1_serializeOpAssociateTargetsWithJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateTargetsWithJobInput(v *AssociateTargetsWithJobInput, encoder *httpbinding.Encoder) error {
@@ -513,10 +497,8 @@ func (m *awsRestjson1_serializeOpAttachPolicy) HandleS
 func (m *awsRestjson1_serializeOpAttachPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -565,7 +547,6 @@ func (m *awsRestjson1_serializeOpAttachPolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachPolicyInput(v *AttachPolicyInput, encoder *httpbinding.Encoder) error {
@@ -607,10 +588,8 @@ func (m *awsRestjson1_serializeOpAttachPrincipalPolicy
 func (m *awsRestjson1_serializeOpAttachPrincipalPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -648,7 +627,6 @@ func (m *awsRestjson1_serializeOpAttachPrincipalPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachPrincipalPolicyInput(v *AttachPrincipalPolicyInput, encoder *httpbinding.Encoder) error {
@@ -683,10 +661,8 @@ func (m *awsRestjson1_serializeOpAttachSecurityProfile
 func (m *awsRestjson1_serializeOpAttachSecurityProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -724,7 +700,6 @@ func (m *awsRestjson1_serializeOpAttachSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachSecurityProfileInput(v *AttachSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -758,10 +733,8 @@ func (m *awsRestjson1_serializeOpAttachThingPrincipal)
 func (m *awsRestjson1_serializeOpAttachThingPrincipal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -799,7 +772,6 @@ func (m *awsRestjson1_serializeOpAttachThingPrincipal)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachThingPrincipalInput(v *AttachThingPrincipalInput, encoder *httpbinding.Encoder) error {
@@ -838,10 +810,8 @@ func (m *awsRestjson1_serializeOpCancelAuditMitigation
 func (m *awsRestjson1_serializeOpCancelAuditMitigationActionsTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -879,7 +849,6 @@ func (m *awsRestjson1_serializeOpCancelAuditMitigation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelAuditMitigationActionsTaskInput(v *CancelAuditMitigationActionsTaskInput, encoder *httpbinding.Encoder) error {
@@ -909,10 +878,8 @@ func (m *awsRestjson1_serializeOpCancelAuditTask) Hand
 func (m *awsRestjson1_serializeOpCancelAuditTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -950,7 +917,6 @@ func (m *awsRestjson1_serializeOpCancelAuditTask) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelAuditTaskInput(v *CancelAuditTaskInput, encoder *httpbinding.Encoder) error {
@@ -980,10 +946,8 @@ func (m *awsRestjson1_serializeOpCancelCertificateTran
 func (m *awsRestjson1_serializeOpCancelCertificateTransfer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1021,7 +985,6 @@ func (m *awsRestjson1_serializeOpCancelCertificateTran
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelCertificateTransferInput(v *CancelCertificateTransferInput, encoder *httpbinding.Encoder) error {
@@ -1051,10 +1014,8 @@ func (m *awsRestjson1_serializeOpCancelDetectMitigatio
 func (m *awsRestjson1_serializeOpCancelDetectMitigationActionsTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1092,7 +1053,6 @@ func (m *awsRestjson1_serializeOpCancelDetectMitigatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelDetectMitigationActionsTaskInput(v *CancelDetectMitigationActionsTaskInput, encoder *httpbinding.Encoder) error {
@@ -1122,10 +1082,8 @@ func (m *awsRestjson1_serializeOpCancelJob) HandleSeri
 func (m *awsRestjson1_serializeOpCancelJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1174,7 +1132,6 @@ func (m *awsRestjson1_serializeOpCancelJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelJobInput(v *CancelJobInput, encoder *httpbinding.Encoder) error {
@@ -1225,10 +1182,8 @@ func (m *awsRestjson1_serializeOpCancelJobExecution) H
 func (m *awsRestjson1_serializeOpCancelJobExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1277,7 +1232,6 @@ func (m *awsRestjson1_serializeOpCancelJobExecution) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelJobExecutionInput(v *CancelJobExecutionInput, encoder *httpbinding.Encoder) error {
@@ -1339,10 +1293,8 @@ func (m *awsRestjson1_serializeOpClearDefaultAuthorize
 func (m *awsRestjson1_serializeOpClearDefaultAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1376,7 +1328,6 @@ func (m *awsRestjson1_serializeOpClearDefaultAuthorize
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsClearDefaultAuthorizerInput(v *ClearDefaultAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -1397,10 +1348,8 @@ func (m *awsRestjson1_serializeOpConfirmTopicRuleDesti
 func (m *awsRestjson1_serializeOpConfirmTopicRuleDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1438,7 +1387,6 @@ func (m *awsRestjson1_serializeOpConfirmTopicRuleDesti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsConfirmTopicRuleDestinationInput(v *ConfirmTopicRuleDestinationInput, encoder *httpbinding.Encoder) error {
@@ -1468,10 +1416,8 @@ func (m *awsRestjson1_serializeOpCreateAuditSuppressio
 func (m *awsRestjson1_serializeOpCreateAuditSuppression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1516,7 +1462,6 @@ func (m *awsRestjson1_serializeOpCreateAuditSuppressio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAuditSuppressionInput(v *CreateAuditSuppressionInput, encoder *httpbinding.Encoder) error {
@@ -1576,10 +1521,8 @@ func (m *awsRestjson1_serializeOpCreateAuthorizer) Han
 func (m *awsRestjson1_serializeOpCreateAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1628,7 +1571,6 @@ func (m *awsRestjson1_serializeOpCreateAuthorizer) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAuthorizerInput(v *CreateAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -1704,10 +1646,8 @@ func (m *awsRestjson1_serializeOpCreateBillingGroup) H
 func (m *awsRestjson1_serializeOpCreateBillingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1756,7 +1696,6 @@ func (m *awsRestjson1_serializeOpCreateBillingGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBillingGroupInput(v *CreateBillingGroupInput, encoder *httpbinding.Encoder) error {
@@ -1807,10 +1746,8 @@ func (m *awsRestjson1_serializeOpCreateCertificateFrom
 func (m *awsRestjson1_serializeOpCreateCertificateFromCsr) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1859,7 +1796,6 @@ func (m *awsRestjson1_serializeOpCreateCertificateFrom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCertificateFromCsrInput(v *CreateCertificateFromCsrInput, encoder *httpbinding.Encoder) error {
@@ -1896,10 +1832,8 @@ func (m *awsRestjson1_serializeOpCreateCertificateProv
 func (m *awsRestjson1_serializeOpCreateCertificateProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1948,7 +1882,6 @@ func (m *awsRestjson1_serializeOpCreateCertificateProv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCertificateProviderInput(v *CreateCertificateProviderInput, encoder *httpbinding.Encoder) error {
@@ -2009,10 +1942,8 @@ func (m *awsRestjson1_serializeOpCreateCommand) Handle
 func (m *awsRestjson1_serializeOpCreateCommand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2061,7 +1992,6 @@ func (m *awsRestjson1_serializeOpCreateCommand) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCommandInput(v *CreateCommandInput, encoder *httpbinding.Encoder) error {
@@ -2139,10 +2069,8 @@ func (m *awsRestjson1_serializeOpCreateCustomMetric) H
 func (m *awsRestjson1_serializeOpCreateCustomMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2191,7 +2119,6 @@ func (m *awsRestjson1_serializeOpCreateCustomMetric) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCustomMetricInput(v *CreateCustomMetricInput, encoder *httpbinding.Encoder) error {
@@ -2250,10 +2177,8 @@ func (m *awsRestjson1_serializeOpCreateDimension) Hand
 func (m *awsRestjson1_serializeOpCreateDimension) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2302,7 +2227,6 @@ func (m *awsRestjson1_serializeOpCreateDimension) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDimensionInput(v *CreateDimensionInput, encoder *httpbinding.Encoder) error {
@@ -2363,10 +2287,8 @@ func (m *awsRestjson1_serializeOpCreateDomainConfigura
 func (m *awsRestjson1_serializeOpCreateDomainConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2415,7 +2337,6 @@ func (m *awsRestjson1_serializeOpCreateDomainConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainConfigurationInput(v *CreateDomainConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2519,10 +2440,8 @@ func (m *awsRestjson1_serializeOpCreateDynamicThingGro
 func (m *awsRestjson1_serializeOpCreateDynamicThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2571,7 +2490,6 @@ func (m *awsRestjson1_serializeOpCreateDynamicThingGro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDynamicThingGroupInput(v *CreateDynamicThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -2637,10 +2555,8 @@ func (m *awsRestjson1_serializeOpCreateFleetMetric) Ha
 func (m *awsRestjson1_serializeOpCreateFleetMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2689,7 +2605,6 @@ func (m *awsRestjson1_serializeOpCreateFleetMetric) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFleetMetricInput(v *CreateFleetMetricInput, encoder *httpbinding.Encoder) error {
@@ -2775,10 +2690,8 @@ func (m *awsRestjson1_serializeOpCreateJob) HandleSeri
 func (m *awsRestjson1_serializeOpCreateJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2827,7 +2740,6 @@ func (m *awsRestjson1_serializeOpCreateJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateJobInput(v *CreateJobInput, encoder *httpbinding.Encoder) error {
@@ -2964,10 +2876,8 @@ func (m *awsRestjson1_serializeOpCreateJobTemplate) Ha
 func (m *awsRestjson1_serializeOpCreateJobTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3016,7 +2926,6 @@ func (m *awsRestjson1_serializeOpCreateJobTemplate) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateJobTemplateInput(v *CreateJobTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3129,10 +3038,8 @@ func (m *awsRestjson1_serializeOpCreateKeysAndCertific
 func (m *awsRestjson1_serializeOpCreateKeysAndCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3170,7 +3077,6 @@ func (m *awsRestjson1_serializeOpCreateKeysAndCertific
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateKeysAndCertificateInput(v *CreateKeysAndCertificateInput, encoder *httpbinding.Encoder) error {
@@ -3195,10 +3101,8 @@ func (m *awsRestjson1_serializeOpCreateMitigationActio
 func (m *awsRestjson1_serializeOpCreateMitigationAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3247,7 +3151,6 @@ func (m *awsRestjson1_serializeOpCreateMitigationActio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMitigationActionInput(v *CreateMitigationActionInput, encoder *httpbinding.Encoder) error {
@@ -3303,10 +3206,8 @@ func (m *awsRestjson1_serializeOpCreateOTAUpdate) Hand
 func (m *awsRestjson1_serializeOpCreateOTAUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3355,7 +3256,6 @@ func (m *awsRestjson1_serializeOpCreateOTAUpdate) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOTAUpdateInput(v *CreateOTAUpdateInput, encoder *httpbinding.Encoder) error {
@@ -3470,10 +3370,8 @@ func (m *awsRestjson1_serializeOpCreatePackage) Handle
 func (m *awsRestjson1_serializeOpCreatePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3522,7 +3420,6 @@ func (m *awsRestjson1_serializeOpCreatePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePackageInput(v *CreatePackageInput, encoder *httpbinding.Encoder) error {
@@ -3575,10 +3472,8 @@ func (m *awsRestjson1_serializeOpCreatePackageVersion)
 func (m *awsRestjson1_serializeOpCreatePackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3627,7 +3522,6 @@ func (m *awsRestjson1_serializeOpCreatePackageVersion)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePackageVersionInput(v *CreatePackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -3708,10 +3602,8 @@ func (m *awsRestjson1_serializeOpCreatePolicy) HandleS
 func (m *awsRestjson1_serializeOpCreatePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3760,7 +3652,6 @@ func (m *awsRestjson1_serializeOpCreatePolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePolicyInput(v *CreatePolicyInput, encoder *httpbinding.Encoder) error {
@@ -3809,10 +3700,8 @@ func (m *awsRestjson1_serializeOpCreatePolicyVersion) 
 func (m *awsRestjson1_serializeOpCreatePolicyVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3861,7 +3750,6 @@ func (m *awsRestjson1_serializeOpCreatePolicyVersion) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePolicyVersionInput(v *CreatePolicyVersionInput, encoder *httpbinding.Encoder) error {
@@ -3907,10 +3795,8 @@ func (m *awsRestjson1_serializeOpCreateProvisioningCla
 func (m *awsRestjson1_serializeOpCreateProvisioningClaim) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3948,7 +3834,6 @@ func (m *awsRestjson1_serializeOpCreateProvisioningCla
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProvisioningClaimInput(v *CreateProvisioningClaimInput, encoder *httpbinding.Encoder) error {
@@ -3978,10 +3863,8 @@ func (m *awsRestjson1_serializeOpCreateProvisioningTem
 func (m *awsRestjson1_serializeOpCreateProvisioningTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4026,7 +3909,6 @@ func (m *awsRestjson1_serializeOpCreateProvisioningTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProvisioningTemplateInput(v *CreateProvisioningTemplateInput, encoder *httpbinding.Encoder) error {
@@ -4098,10 +3980,8 @@ func (m *awsRestjson1_serializeOpCreateProvisioningTem
 func (m *awsRestjson1_serializeOpCreateProvisioningTemplateVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4150,7 +4030,6 @@ func (m *awsRestjson1_serializeOpCreateProvisioningTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProvisioningTemplateVersionInput(v *CreateProvisioningTemplateVersionInput, encoder *httpbinding.Encoder) error {
@@ -4196,10 +4075,8 @@ func (m *awsRestjson1_serializeOpCreateRoleAlias) Hand
 func (m *awsRestjson1_serializeOpCreateRoleAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4248,7 +4125,6 @@ func (m *awsRestjson1_serializeOpCreateRoleAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRoleAliasInput(v *CreateRoleAliasInput, encoder *httpbinding.Encoder) error {
@@ -4302,10 +4178,8 @@ func (m *awsRestjson1_serializeOpCreateScheduledAudit)
 func (m *awsRestjson1_serializeOpCreateScheduledAudit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4354,7 +4228,6 @@ func (m *awsRestjson1_serializeOpCreateScheduledAudit)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateScheduledAuditInput(v *CreateScheduledAuditInput, encoder *httpbinding.Encoder) error {
@@ -4420,10 +4293,8 @@ func (m *awsRestjson1_serializeOpCreateSecurityProfile
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
@@ -4472,7 +4343,6 @@ func (m *awsRestjson1_serializeOpCreateSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSecurityProfileInput(v *CreateSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -4556,10 +4426,8 @@ func (m *awsRestjson1_serializeOpCreateStream) HandleS
 func (m *awsRestjson1_serializeOpCreateStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4608,7 +4476,6 @@ func (m *awsRestjson1_serializeOpCreateStream) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateStreamInput(v *CreateStreamInput, encoder *httpbinding.Encoder) error {
@@ -4669,10 +4536,8 @@ func (m *awsRestjson1_serializeOpCreateThing) HandleSe
 func (m *awsRestjson1_serializeOpCreateThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4721,7 +4586,6 @@ func (m *awsRestjson1_serializeOpCreateThing) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateThingInput(v *CreateThingInput, encoder *httpbinding.Encoder) error {
@@ -4775,10 +4639,8 @@ func (m *awsRestjson1_serializeOpCreateThingGroup) Han
 func (m *awsRestjson1_serializeOpCreateThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4827,7 +4689,6 @@ func (m *awsRestjson1_serializeOpCreateThingGroup) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateThingGroupInput(v *CreateThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -4883,10 +4744,8 @@ func (m *awsRestjson1_serializeOpCreateThingType) Hand
 func (m *awsRestjson1_serializeOpCreateThingType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4935,7 +4794,6 @@ func (m *awsRestjson1_serializeOpCreateThingType) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateThingTypeInput(v *CreateThingTypeInput, encoder *httpbinding.Encoder) error {
@@ -4986,10 +4844,8 @@ func (m *awsRestjson1_serializeOpCreateTopicRule) Hand
 func (m *awsRestjson1_serializeOpCreateTopicRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5051,7 +4907,6 @@ func (m *awsRestjson1_serializeOpCreateTopicRule) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTopicRuleInput(v *CreateTopicRuleInput, encoder *httpbinding.Encoder) error {
@@ -5086,10 +4941,8 @@ func (m *awsRestjson1_serializeOpCreateTopicRuleDestin
 func (m *awsRestjson1_serializeOpCreateTopicRuleDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5134,7 +4987,6 @@ func (m *awsRestjson1_serializeOpCreateTopicRuleDestin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTopicRuleDestinationInput(v *CreateTopicRuleDestinationInput, encoder *httpbinding.Encoder) error {
@@ -5169,10 +5021,8 @@ func (m *awsRestjson1_serializeOpDeleteAccountAuditCon
 func (m *awsRestjson1_serializeOpDeleteAccountAuditConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5210,7 +5060,6 @@ func (m *awsRestjson1_serializeOpDeleteAccountAuditCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccountAuditConfigurationInput(v *DeleteAccountAuditConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5235,10 +5084,8 @@ func (m *awsRestjson1_serializeOpDeleteAuditSuppressio
 func (m *awsRestjson1_serializeOpDeleteAuditSuppression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5283,7 +5130,6 @@ func (m *awsRestjson1_serializeOpDeleteAuditSuppressio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAuditSuppressionInput(v *DeleteAuditSuppressionInput, encoder *httpbinding.Encoder) error {
@@ -5323,10 +5169,8 @@ func (m *awsRestjson1_serializeOpDeleteAuthorizer) Han
 func (m *awsRestjson1_serializeOpDeleteAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5364,7 +5208,6 @@ func (m *awsRestjson1_serializeOpDeleteAuthorizer) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAuthorizerInput(v *DeleteAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -5394,10 +5237,8 @@ func (m *awsRestjson1_serializeOpDeleteBillingGroup) H
 func (m *awsRestjson1_serializeOpDeleteBillingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5435,7 +5276,6 @@ func (m *awsRestjson1_serializeOpDeleteBillingGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBillingGroupInput(v *DeleteBillingGroupInput, encoder *httpbinding.Encoder) error {
@@ -5469,10 +5309,8 @@ func (m *awsRestjson1_serializeOpDeleteCACertificate) 
 func (m *awsRestjson1_serializeOpDeleteCACertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5510,7 +5348,6 @@ func (m *awsRestjson1_serializeOpDeleteCACertificate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCACertificateInput(v *DeleteCACertificateInput, encoder *httpbinding.Encoder) error {
@@ -5540,10 +5377,8 @@ func (m *awsRestjson1_serializeOpDeleteCertificate) Ha
 func (m *awsRestjson1_serializeOpDeleteCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5581,7 +5416,6 @@ func (m *awsRestjson1_serializeOpDeleteCertificate) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCertificateInput(v *DeleteCertificateInput, encoder *httpbinding.Encoder) error {
@@ -5615,10 +5449,8 @@ func (m *awsRestjson1_serializeOpDeleteCertificateProv
 func (m *awsRestjson1_serializeOpDeleteCertificateProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5656,7 +5488,6 @@ func (m *awsRestjson1_serializeOpDeleteCertificateProv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCertificateProviderInput(v *DeleteCertificateProviderInput, encoder *httpbinding.Encoder) error {
@@ -5686,10 +5517,8 @@ func (m *awsRestjson1_serializeOpDeleteCommand) Handle
 func (m *awsRestjson1_serializeOpDeleteCommand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5727,7 +5556,6 @@ func (m *awsRestjson1_serializeOpDeleteCommand) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCommandInput(v *DeleteCommandInput, encoder *httpbinding.Encoder) error {
@@ -5757,10 +5585,8 @@ func (m *awsRestjson1_serializeOpDeleteCommandExecutio
 func (m *awsRestjson1_serializeOpDeleteCommandExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5798,7 +5624,6 @@ func (m *awsRestjson1_serializeOpDeleteCommandExecutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCommandExecutionInput(v *DeleteCommandExecutionInput, encoder *httpbinding.Encoder) error {
@@ -5832,10 +5657,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomMetric) H
 func (m *awsRestjson1_serializeOpDeleteCustomMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5873,7 +5696,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomMetric) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomMetricInput(v *DeleteCustomMetricInput, encoder *httpbinding.Encoder) error {
@@ -5903,10 +5725,8 @@ func (m *awsRestjson1_serializeOpDeleteDimension) Hand
 func (m *awsRestjson1_serializeOpDeleteDimension) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5944,7 +5764,6 @@ func (m *awsRestjson1_serializeOpDeleteDimension) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDimensionInput(v *DeleteDimensionInput, encoder *httpbinding.Encoder) error {
@@ -5974,10 +5793,8 @@ func (m *awsRestjson1_serializeOpDeleteDomainConfigura
 func (m *awsRestjson1_serializeOpDeleteDomainConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6015,7 +5832,6 @@ func (m *awsRestjson1_serializeOpDeleteDomainConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainConfigurationInput(v *DeleteDomainConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6045,10 +5861,8 @@ func (m *awsRestjson1_serializeOpDeleteDynamicThingGro
 func (m *awsRestjson1_serializeOpDeleteDynamicThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6086,7 +5900,6 @@ func (m *awsRestjson1_serializeOpDeleteDynamicThingGro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDynamicThingGroupInput(v *DeleteDynamicThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -6120,10 +5933,8 @@ func (m *awsRestjson1_serializeOpDeleteFleetMetric) Ha
 func (m *awsRestjson1_serializeOpDeleteFleetMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6161,7 +5972,6 @@ func (m *awsRestjson1_serializeOpDeleteFleetMetric) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFleetMetricInput(v *DeleteFleetMetricInput, encoder *httpbinding.Encoder) error {
@@ -6195,10 +6005,8 @@ func (m *awsRestjson1_serializeOpDeleteJob) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6236,7 +6044,6 @@ func (m *awsRestjson1_serializeOpDeleteJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteJobInput(v *DeleteJobInput, encoder *httpbinding.Encoder) error {
@@ -6274,10 +6081,8 @@ func (m *awsRestjson1_serializeOpDeleteJobExecution) H
 func (m *awsRestjson1_serializeOpDeleteJobExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6315,7 +6120,6 @@ func (m *awsRestjson1_serializeOpDeleteJobExecution) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteJobExecutionInput(v *DeleteJobExecutionInput, encoder *httpbinding.Encoder) error {
@@ -6371,10 +6175,8 @@ func (m *awsRestjson1_serializeOpDeleteJobTemplate) Ha
 func (m *awsRestjson1_serializeOpDeleteJobTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6412,7 +6214,6 @@ func (m *awsRestjson1_serializeOpDeleteJobTemplate) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteJobTemplateInput(v *DeleteJobTemplateInput, encoder *httpbinding.Encoder) error {
@@ -6442,10 +6243,8 @@ func (m *awsRestjson1_serializeOpDeleteMitigationActio
 func (m *awsRestjson1_serializeOpDeleteMitigationAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6483,7 +6282,6 @@ func (m *awsRestjson1_serializeOpDeleteMitigationActio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMitigationActionInput(v *DeleteMitigationActionInput, encoder *httpbinding.Encoder) error {
@@ -6513,10 +6311,8 @@ func (m *awsRestjson1_serializeOpDeleteOTAUpdate) Hand
 func (m *awsRestjson1_serializeOpDeleteOTAUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6554,7 +6350,6 @@ func (m *awsRestjson1_serializeOpDeleteOTAUpdate) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteOTAUpdateInput(v *DeleteOTAUpdateInput, encoder *httpbinding.Encoder) error {
@@ -6592,10 +6387,8 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
 func (m *awsRestjson1_serializeOpDeletePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6633,7 +6426,6 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePackageInput(v *DeletePackageInput, encoder *httpbinding.Encoder) error {
@@ -6667,10 +6459,8 @@ func (m *awsRestjson1_serializeOpDeletePackageVersion)
 func (m *awsRestjson1_serializeOpDeletePackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6708,7 +6498,6 @@ func (m *awsRestjson1_serializeOpDeletePackageVersion)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePackageVersionInput(v *DeletePackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -6751,10 +6540,8 @@ func (m *awsRestjson1_serializeOpDeletePolicy) HandleS
 func (m *awsRestjson1_serializeOpDeletePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6792,7 +6579,6 @@ func (m *awsRestjson1_serializeOpDeletePolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePolicyInput(v *DeletePolicyInput, encoder *httpbinding.Encoder) error {
@@ -6822,10 +6608,8 @@ func (m *awsRestjson1_serializeOpDeletePolicyVersion) 
 func (m *awsRestjson1_serializeOpDeletePolicyVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6863,7 +6647,6 @@ func (m *awsRestjson1_serializeOpDeletePolicyVersion) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePolicyVersionInput(v *DeletePolicyVersionInput, encoder *httpbinding.Encoder) error {
@@ -6902,10 +6685,8 @@ func (m *awsRestjson1_serializeOpDeleteProvisioningTem
 func (m *awsRestjson1_serializeOpDeleteProvisioningTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6943,7 +6724,6 @@ func (m *awsRestjson1_serializeOpDeleteProvisioningTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProvisioningTemplateInput(v *DeleteProvisioningTemplateInput, encoder *httpbinding.Encoder) error {
@@ -6973,10 +6753,8 @@ func (m *awsRestjson1_serializeOpDeleteProvisioningTem
 func (m *awsRestjson1_serializeOpDeleteProvisioningTemplateVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7014,7 +6792,6 @@ func (m *awsRestjson1_serializeOpDeleteProvisioningTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProvisioningTemplateVersionInput(v *DeleteProvisioningTemplateVersionInput, encoder *httpbinding.Encoder) error {
@@ -7053,10 +6830,8 @@ func (m *awsRestjson1_serializeOpDeleteRegistrationCod
 func (m *awsRestjson1_serializeOpDeleteRegistrationCode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7090,7 +6865,6 @@ func (m *awsRestjson1_serializeOpDeleteRegistrationCod
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRegistrationCodeInput(v *DeleteRegistrationCodeInput, encoder *httpbinding.Encoder) error {
@@ -7111,10 +6885,8 @@ func (m *awsRestjson1_serializeOpDeleteRoleAlias) Hand
 func (m *awsRestjson1_serializeOpDeleteRoleAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7152,7 +6924,6 @@ func (m *awsRestjson1_serializeOpDeleteRoleAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRoleAliasInput(v *DeleteRoleAliasInput, encoder *httpbinding.Encoder) error {
@@ -7182,10 +6953,8 @@ func (m *awsRestjson1_serializeOpDeleteScheduledAudit)
 func (m *awsRestjson1_serializeOpDeleteScheduledAudit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7223,7 +6992,6 @@ func (m *awsRestjson1_serializeOpDeleteScheduledAudit)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteScheduledAuditInput(v *DeleteScheduledAuditInput, encoder *httpbinding.Encoder) error {
@@ -7253,10 +7021,8 @@ func (m *awsRestjson1_serializeOpDeleteSecurityProfile
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
@@ -7294,7 +7060,6 @@ func (m *awsRestjson1_serializeOpDeleteSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSecurityProfileInput(v *DeleteSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -7328,10 +7093,8 @@ func (m *awsRestjson1_serializeOpDeleteStream) HandleS
 func (m *awsRestjson1_serializeOpDeleteStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7369,7 +7132,6 @@ func (m *awsRestjson1_serializeOpDeleteStream) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteStreamInput(v *DeleteStreamInput, encoder *httpbinding.Encoder) error {
@@ -7399,10 +7161,8 @@ func (m *awsRestjson1_serializeOpDeleteThing) HandleSe
 func (m *awsRestjson1_serializeOpDeleteThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7440,7 +7200,6 @@ func (m *awsRestjson1_serializeOpDeleteThing) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteThingInput(v *DeleteThingInput, encoder *httpbinding.Encoder) error {
@@ -7474,10 +7233,8 @@ func (m *awsRestjson1_serializeOpDeleteThingGroup) Han
 func (m *awsRestjson1_serializeOpDeleteThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7515,7 +7272,6 @@ func (m *awsRestjson1_serializeOpDeleteThingGroup) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteThingGroupInput(v *DeleteThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -7549,10 +7305,8 @@ func (m *awsRestjson1_serializeOpDeleteThingType) Hand
 func (m *awsRestjson1_serializeOpDeleteThingType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7590,7 +7344,6 @@ func (m *awsRestjson1_serializeOpDeleteThingType) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteThingTypeInput(v *DeleteThingTypeInput, encoder *httpbinding.Encoder) error {
@@ -7620,10 +7373,8 @@ func (m *awsRestjson1_serializeOpDeleteTopicRule) Hand
 func (m *awsRestjson1_serializeOpDeleteTopicRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7661,7 +7412,6 @@ func (m *awsRestjson1_serializeOpDeleteTopicRule) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTopicRuleInput(v *DeleteTopicRuleInput, encoder *httpbinding.Encoder) error {
@@ -7691,10 +7441,8 @@ func (m *awsRestjson1_serializeOpDeleteTopicRuleDestin
 func (m *awsRestjson1_serializeOpDeleteTopicRuleDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7732,7 +7480,6 @@ func (m *awsRestjson1_serializeOpDeleteTopicRuleDestin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTopicRuleDestinationInput(v *DeleteTopicRuleDestinationInput, encoder *httpbinding.Encoder) error {
@@ -7762,10 +7509,8 @@ func (m *awsRestjson1_serializeOpDeleteV2LoggingLevel)
 func (m *awsRestjson1_serializeOpDeleteV2LoggingLevel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7803,7 +7548,6 @@ func (m *awsRestjson1_serializeOpDeleteV2LoggingLevel)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteV2LoggingLevelInput(v *DeleteV2LoggingLevelInput, encoder *httpbinding.Encoder) error {
@@ -7832,10 +7576,8 @@ func (m *awsRestjson1_serializeOpDeprecateThingType) H
 func (m *awsRestjson1_serializeOpDeprecateThingType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7884,7 +7626,6 @@ func (m *awsRestjson1_serializeOpDeprecateThingType) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeprecateThingTypeInput(v *DeprecateThingTypeInput, encoder *httpbinding.Encoder) error {
@@ -7926,10 +7667,8 @@ func (m *awsRestjson1_serializeOpDescribeAccountAuditC
 func (m *awsRestjson1_serializeOpDescribeAccountAuditConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7963,7 +7702,6 @@ func (m *awsRestjson1_serializeOpDescribeAccountAuditC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccountAuditConfigurationInput(v *DescribeAccountAuditConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7984,10 +7722,8 @@ func (m *awsRestjson1_serializeOpDescribeAuditFinding)
 func (m *awsRestjson1_serializeOpDescribeAuditFinding) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8025,7 +7761,6 @@ func (m *awsRestjson1_serializeOpDescribeAuditFinding)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAuditFindingInput(v *DescribeAuditFindingInput, encoder *httpbinding.Encoder) error {
@@ -8055,10 +7790,8 @@ func (m *awsRestjson1_serializeOpDescribeAuditMitigati
 func (m *awsRestjson1_serializeOpDescribeAuditMitigationActionsTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8096,7 +7829,6 @@ func (m *awsRestjson1_serializeOpDescribeAuditMitigati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAuditMitigationActionsTaskInput(v *DescribeAuditMitigationActionsTaskInput, encoder *httpbinding.Encoder) error {
@@ -8126,10 +7858,8 @@ func (m *awsRestjson1_serializeOpDescribeAuditSuppress
 func (m *awsRestjson1_serializeOpDescribeAuditSuppression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8174,7 +7904,6 @@ func (m *awsRestjson1_serializeOpDescribeAuditSuppress
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAuditSuppressionInput(v *DescribeAuditSuppressionInput, encoder *httpbinding.Encoder) error {
@@ -8214,10 +7943,8 @@ func (m *awsRestjson1_serializeOpDescribeAuditTask) Ha
 func (m *awsRestjson1_serializeOpDescribeAuditTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8255,7 +7982,6 @@ func (m *awsRestjson1_serializeOpDescribeAuditTask) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAuditTaskInput(v *DescribeAuditTaskInput, encoder *httpbinding.Encoder) error {
@@ -8285,10 +8011,8 @@ func (m *awsRestjson1_serializeOpDescribeAuthorizer) H
 func (m *awsRestjson1_serializeOpDescribeAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8326,7 +8050,6 @@ func (m *awsRestjson1_serializeOpDescribeAuthorizer) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAuthorizerInput(v *DescribeAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -8356,10 +8079,8 @@ func (m *awsRestjson1_serializeOpDescribeBillingGroup)
 func (m *awsRestjson1_serializeOpDescribeBillingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8397,7 +8118,6 @@ func (m *awsRestjson1_serializeOpDescribeBillingGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBillingGroupInput(v *DescribeBillingGroupInput, encoder *httpbinding.Encoder) error {
@@ -8427,10 +8147,8 @@ func (m *awsRestjson1_serializeOpDescribeCACertificate
 func (m *awsRestjson1_serializeOpDescribeCACertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8468,7 +8186,6 @@ func (m *awsRestjson1_serializeOpDescribeCACertificate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeCACertificateInput(v *DescribeCACertificateInput, encoder *httpbinding.Encoder) error {
@@ -8498,10 +8215,8 @@ func (m *awsRestjson1_serializeOpDescribeCertificate) 
 func (m *awsRestjson1_serializeOpDescribeCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8539,7 +8254,6 @@ func (m *awsRestjson1_serializeOpDescribeCertificate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeCertificateInput(v *DescribeCertificateInput, encoder *httpbinding.Encoder) error {
@@ -8569,10 +8283,8 @@ func (m *awsRestjson1_serializeOpDescribeCertificatePr
 func (m *awsRestjson1_serializeOpDescribeCertificateProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8610,7 +8322,6 @@ func (m *awsRestjson1_serializeOpDescribeCertificatePr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeCertificateProviderInput(v *DescribeCertificateProviderInput, encoder *httpbinding.Encoder) error {
@@ -8640,10 +8351,8 @@ func (m *awsRestjson1_serializeOpDescribeCustomMetric)
 func (m *awsRestjson1_serializeOpDescribeCustomMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8681,7 +8390,6 @@ func (m *awsRestjson1_serializeOpDescribeCustomMetric)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeCustomMetricInput(v *DescribeCustomMetricInput, encoder *httpbinding.Encoder) error {
@@ -8711,10 +8419,8 @@ func (m *awsRestjson1_serializeOpDescribeDefaultAuthor
 func (m *awsRestjson1_serializeOpDescribeDefaultAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8748,7 +8454,6 @@ func (m *awsRestjson1_serializeOpDescribeDefaultAuthor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDefaultAuthorizerInput(v *DescribeDefaultAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -8769,10 +8474,8 @@ func (m *awsRestjson1_serializeOpDescribeDetectMitigat
 func (m *awsRestjson1_serializeOpDescribeDetectMitigationActionsTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8810,7 +8513,6 @@ func (m *awsRestjson1_serializeOpDescribeDetectMitigat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDetectMitigationActionsTaskInput(v *DescribeDetectMitigationActionsTaskInput, encoder *httpbinding.Encoder) error {
@@ -8840,10 +8542,8 @@ func (m *awsRestjson1_serializeOpDescribeDimension) Ha
 func (m *awsRestjson1_serializeOpDescribeDimension) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8881,7 +8581,6 @@ func (m *awsRestjson1_serializeOpDescribeDimension) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDimensionInput(v *DescribeDimensionInput, encoder *httpbinding.Encoder) error {
@@ -8911,10 +8610,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainConfigu
 func (m *awsRestjson1_serializeOpDescribeDomainConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8952,7 +8649,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainConfigurationInput(v *DescribeDomainConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -8982,10 +8678,8 @@ func (m *awsRestjson1_serializeOpDescribeEndpoint) Han
 func (m *awsRestjson1_serializeOpDescribeEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9023,7 +8717,6 @@ func (m *awsRestjson1_serializeOpDescribeEndpoint) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeEndpointInput(v *DescribeEndpointInput, encoder *httpbinding.Encoder) error {
@@ -9048,10 +8741,8 @@ func (m *awsRestjson1_serializeOpDescribeEventConfigur
 func (m *awsRestjson1_serializeOpDescribeEventConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9085,7 +8776,6 @@ func (m *awsRestjson1_serializeOpDescribeEventConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeEventConfigurationsInput(v *DescribeEventConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -9106,10 +8796,8 @@ func (m *awsRestjson1_serializeOpDescribeFleetMetric) 
 func (m *awsRestjson1_serializeOpDescribeFleetMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9147,7 +8835,6 @@ func (m *awsRestjson1_serializeOpDescribeFleetMetric) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeFleetMetricInput(v *DescribeFleetMetricInput, encoder *httpbinding.Encoder) error {
@@ -9177,10 +8864,8 @@ func (m *awsRestjson1_serializeOpDescribeIndex) Handle
 func (m *awsRestjson1_serializeOpDescribeIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9218,7 +8903,6 @@ func (m *awsRestjson1_serializeOpDescribeIndex) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeIndexInput(v *DescribeIndexInput, encoder *httpbinding.Encoder) error {
@@ -9248,10 +8932,8 @@ func (m *awsRestjson1_serializeOpDescribeJob) HandleSe
 func (m *awsRestjson1_serializeOpDescribeJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9289,7 +8971,6 @@ func (m *awsRestjson1_serializeOpDescribeJob) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeJobInput(v *DescribeJobInput, encoder *httpbinding.Encoder) error {
@@ -9323,10 +9004,8 @@ func (m *awsRestjson1_serializeOpDescribeJobExecution)
 func (m *awsRestjson1_serializeOpDescribeJobExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9364,7 +9043,6 @@ func (m *awsRestjson1_serializeOpDescribeJobExecution)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeJobExecutionInput(v *DescribeJobExecutionInput, encoder *httpbinding.Encoder) error {
@@ -9407,10 +9085,8 @@ func (m *awsRestjson1_serializeOpDescribeJobTemplate) 
 func (m *awsRestjson1_serializeOpDescribeJobTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9448,7 +9124,6 @@ func (m *awsRestjson1_serializeOpDescribeJobTemplate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeJobTemplateInput(v *DescribeJobTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9478,10 +9153,8 @@ func (m *awsRestjson1_serializeOpDescribeManagedJobTem
 func (m *awsRestjson1_serializeOpDescribeManagedJobTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9519,7 +9192,6 @@ func (m *awsRestjson1_serializeOpDescribeManagedJobTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeManagedJobTemplateInput(v *DescribeManagedJobTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9553,10 +9225,8 @@ func (m *awsRestjson1_serializeOpDescribeMitigationAct
 func (m *awsRestjson1_serializeOpDescribeMitigationAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9594,7 +9264,6 @@ func (m *awsRestjson1_serializeOpDescribeMitigationAct
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeMitigationActionInput(v *DescribeMitigationActionInput, encoder *httpbinding.Encoder) error {
@@ -9624,10 +9293,8 @@ func (m *awsRestjson1_serializeOpDescribeProvisioningT
 func (m *awsRestjson1_serializeOpDescribeProvisioningTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9665,7 +9332,6 @@ func (m *awsRestjson1_serializeOpDescribeProvisioningT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeProvisioningTemplateInput(v *DescribeProvisioningTemplateInput, encoder *httpbinding.Encoder) error {
@@ -9695,10 +9361,8 @@ func (m *awsRestjson1_serializeOpDescribeProvisioningT
 func (m *awsRestjson1_serializeOpDescribeProvisioningTemplateVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9736,7 +9400,6 @@ func (m *awsRestjson1_serializeOpDescribeProvisioningT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeProvisioningTemplateVersionInput(v *DescribeProvisioningTemplateVersionInput, encoder *httpbinding.Encoder) error {
@@ -9775,10 +9438,8 @@ func (m *awsRestjson1_serializeOpDescribeRoleAlias) Ha
 func (m *awsRestjson1_serializeOpDescribeRoleAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9816,7 +9477,6 @@ func (m *awsRestjson1_serializeOpDescribeRoleAlias) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRoleAliasInput(v *DescribeRoleAliasInput, encoder *httpbinding.Encoder) error {
@@ -9846,10 +9506,8 @@ func (m *awsRestjson1_serializeOpDescribeScheduledAudi
 func (m *awsRestjson1_serializeOpDescribeScheduledAudit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9887,7 +9545,6 @@ func (m *awsRestjson1_serializeOpDescribeScheduledAudi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeScheduledAuditInput(v *DescribeScheduledAuditInput, encoder *httpbinding.Encoder) error {
@@ -9917,10 +9574,8 @@ func (m *awsRestjson1_serializeOpDescribeSecurityProfi
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
@@ -9958,7 +9613,6 @@ func (m *awsRestjson1_serializeOpDescribeSecurityProfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSecurityProfileInput(v *DescribeSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -9988,10 +9642,8 @@ func (m *awsRestjson1_serializeOpDescribeStream) Handl
 func (m *awsRestjson1_serializeOpDescribeStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10029,7 +9681,6 @@ func (m *awsRestjson1_serializeOpDescribeStream) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeStreamInput(v *DescribeStreamInput, encoder *httpbinding.Encoder) error {
@@ -10059,10 +9710,8 @@ func (m *awsRestjson1_serializeOpDescribeThing) Handle
 func (m *awsRestjson1_serializeOpDescribeThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10100,7 +9749,6 @@ func (m *awsRestjson1_serializeOpDescribeThing) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThingInput(v *DescribeThingInput, encoder *httpbinding.Encoder) error {
@@ -10130,10 +9778,8 @@ func (m *awsRestjson1_serializeOpDescribeThingGroup) H
 func (m *awsRestjson1_serializeOpDescribeThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10171,7 +9817,6 @@ func (m *awsRestjson1_serializeOpDescribeThingGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThingGroupInput(v *DescribeThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -10201,10 +9846,8 @@ func (m *awsRestjson1_serializeOpDescribeThingRegistra
 func (m *awsRestjson1_serializeOpDescribeThingRegistrationTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10242,7 +9885,6 @@ func (m *awsRestjson1_serializeOpDescribeThingRegistra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThingRegistrationTaskInput(v *DescribeThingRegistrationTaskInput, encoder *httpbinding.Encoder) error {
@@ -10272,10 +9914,8 @@ func (m *awsRestjson1_serializeOpDescribeThingType) Ha
 func (m *awsRestjson1_serializeOpDescribeThingType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10313,7 +9953,6 @@ func (m *awsRestjson1_serializeOpDescribeThingType) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeThingTypeInput(v *DescribeThingTypeInput, encoder *httpbinding.Encoder) error {
@@ -10343,10 +9982,8 @@ func (m *awsRestjson1_serializeOpDetachPolicy) HandleS
 func (m *awsRestjson1_serializeOpDetachPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10395,7 +10032,6 @@ func (m *awsRestjson1_serializeOpDetachPolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachPolicyInput(v *DetachPolicyInput, encoder *httpbinding.Encoder) error {
@@ -10437,10 +10073,8 @@ func (m *awsRestjson1_serializeOpDetachPrincipalPolicy
 func (m *awsRestjson1_serializeOpDetachPrincipalPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10478,7 +10112,6 @@ func (m *awsRestjson1_serializeOpDetachPrincipalPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachPrincipalPolicyInput(v *DetachPrincipalPolicyInput, encoder *httpbinding.Encoder) error {
@@ -10513,10 +10146,8 @@ func (m *awsRestjson1_serializeOpDetachSecurityProfile
 func (m *awsRestjson1_serializeOpDetachSecurityProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10554,7 +10185,6 @@ func (m *awsRestjson1_serializeOpDetachSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachSecurityProfileInput(v *DetachSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -10588,10 +10218,8 @@ func (m *awsRestjson1_serializeOpDetachThingPrincipal)
 func (m *awsRestjson1_serializeOpDetachThingPrincipal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10629,7 +10257,6 @@ func (m *awsRestjson1_serializeOpDetachThingPrincipal)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachThingPrincipalInput(v *DetachThingPrincipalInput, encoder *httpbinding.Encoder) error {
@@ -10664,10 +10291,8 @@ func (m *awsRestjson1_serializeOpDisableTopicRule) Han
 func (m *awsRestjson1_serializeOpDisableTopicRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10705,7 +10330,6 @@ func (m *awsRestjson1_serializeOpDisableTopicRule) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableTopicRuleInput(v *DisableTopicRuleInput, encoder *httpbinding.Encoder) error {
@@ -10735,10 +10359,8 @@ func (m *awsRestjson1_serializeOpDisassociateSbomFromP
 func (m *awsRestjson1_serializeOpDisassociateSbomFromPackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10776,7 +10398,6 @@ func (m *awsRestjson1_serializeOpDisassociateSbomFromP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateSbomFromPackageVersionInput(v *DisassociateSbomFromPackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -10819,10 +10440,8 @@ func (m *awsRestjson1_serializeOpEnableTopicRule) Hand
 func (m *awsRestjson1_serializeOpEnableTopicRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10860,7 +10479,6 @@ func (m *awsRestjson1_serializeOpEnableTopicRule) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableTopicRuleInput(v *EnableTopicRuleInput, encoder *httpbinding.Encoder) error {
@@ -10890,10 +10508,8 @@ func (m *awsRestjson1_serializeOpGetBehaviorModelTrain
 func (m *awsRestjson1_serializeOpGetBehaviorModelTrainingSummaries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10931,7 +10547,6 @@ func (m *awsRestjson1_serializeOpGetBehaviorModelTrain
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBehaviorModelTrainingSummariesInput(v *GetBehaviorModelTrainingSummariesInput, encoder *httpbinding.Encoder) error {
@@ -10964,10 +10579,8 @@ func (m *awsRestjson1_serializeOpGetBucketsAggregation
 func (m *awsRestjson1_serializeOpGetBucketsAggregation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11012,7 +10625,6 @@ func (m *awsRestjson1_serializeOpGetBucketsAggregation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBucketsAggregationInput(v *GetBucketsAggregationInput, encoder *httpbinding.Encoder) error {
@@ -11067,10 +10679,8 @@ func (m *awsRestjson1_serializeOpGetCardinality) Handl
 func (m *awsRestjson1_serializeOpGetCardinality) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11115,7 +10725,6 @@ func (m *awsRestjson1_serializeOpGetCardinality) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCardinalityInput(v *GetCardinalityInput, encoder *httpbinding.Encoder) error {
@@ -11163,10 +10772,8 @@ func (m *awsRestjson1_serializeOpGetCommand) HandleSer
 func (m *awsRestjson1_serializeOpGetCommand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11204,7 +10811,6 @@ func (m *awsRestjson1_serializeOpGetCommand) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCommandInput(v *GetCommandInput, encoder *httpbinding.Encoder) error {
@@ -11234,10 +10840,8 @@ func (m *awsRestjson1_serializeOpGetCommandExecution) 
 func (m *awsRestjson1_serializeOpGetCommandExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11275,7 +10879,6 @@ func (m *awsRestjson1_serializeOpGetCommandExecution) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCommandExecutionInput(v *GetCommandExecutionInput, encoder *httpbinding.Encoder) error {
@@ -11313,10 +10916,8 @@ func (m *awsRestjson1_serializeOpGetEffectivePolicies)
 func (m *awsRestjson1_serializeOpGetEffectivePolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11365,7 +10966,6 @@ func (m *awsRestjson1_serializeOpGetEffectivePolicies)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEffectivePoliciesInput(v *GetEffectivePoliciesInput, encoder *httpbinding.Encoder) error {
@@ -11407,10 +11007,8 @@ func (m *awsRestjson1_serializeOpGetIndexingConfigurat
 func (m *awsRestjson1_serializeOpGetIndexingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11444,7 +11042,6 @@ func (m *awsRestjson1_serializeOpGetIndexingConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIndexingConfigurationInput(v *GetIndexingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -11465,10 +11062,8 @@ func (m *awsRestjson1_serializeOpGetJobDocument) Handl
 func (m *awsRestjson1_serializeOpGetJobDocument) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11506,7 +11101,6 @@ func (m *awsRestjson1_serializeOpGetJobDocument) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJobDocumentInput(v *GetJobDocumentInput, encoder *httpbinding.Encoder) error {
@@ -11540,10 +11134,8 @@ func (m *awsRestjson1_serializeOpGetLoggingOptions) Ha
 func (m *awsRestjson1_serializeOpGetLoggingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11577,7 +11169,6 @@ func (m *awsRestjson1_serializeOpGetLoggingOptions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLoggingOptionsInput(v *GetLoggingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -11598,10 +11189,8 @@ func (m *awsRestjson1_serializeOpGetOTAUpdate) HandleS
 func (m *awsRestjson1_serializeOpGetOTAUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11639,7 +11228,6 @@ func (m *awsRestjson1_serializeOpGetOTAUpdate) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOTAUpdateInput(v *GetOTAUpdateInput, encoder *httpbinding.Encoder) error {
@@ -11669,10 +11257,8 @@ func (m *awsRestjson1_serializeOpGetPackage) HandleSer
 func (m *awsRestjson1_serializeOpGetPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11710,7 +11296,6 @@ func (m *awsRestjson1_serializeOpGetPackage) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPackageInput(v *GetPackageInput, encoder *httpbinding.Encoder) error {
@@ -11740,10 +11325,8 @@ func (m *awsRestjson1_serializeOpGetPackageConfigurati
 func (m *awsRestjson1_serializeOpGetPackageConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11777,7 +11360,6 @@ func (m *awsRestjson1_serializeOpGetPackageConfigurati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPackageConfigurationInput(v *GetPackageConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -11798,10 +11380,8 @@ func (m *awsRestjson1_serializeOpGetPackageVersion) Ha
 func (m *awsRestjson1_serializeOpGetPackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11839,7 +11419,6 @@ func (m *awsRestjson1_serializeOpGetPackageVersion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPackageVersionInput(v *GetPackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -11878,10 +11457,8 @@ func (m *awsRestjson1_serializeOpGetPercentiles) Handl
 func (m *awsRestjson1_serializeOpGetPercentiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11926,7 +11503,6 @@ func (m *awsRestjson1_serializeOpGetPercentiles) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPercentilesInput(v *GetPercentilesInput, encoder *httpbinding.Encoder) error {
@@ -11981,10 +11557,8 @@ func (m *awsRestjson1_serializeOpGetPolicy) HandleSeri
 func (m *awsRestjson1_serializeOpGetPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12022,7 +11596,6 @@ func (m *awsRestjson1_serializeOpGetPolicy) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPolicyInput(v *GetPolicyInput, encoder *httpbinding.Encoder) error {
@@ -12052,10 +11625,8 @@ func (m *awsRestjson1_serializeOpGetPolicyVersion) Han
 func (m *awsRestjson1_serializeOpGetPolicyVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12093,7 +11664,6 @@ func (m *awsRestjson1_serializeOpGetPolicyVersion) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPolicyVersionInput(v *GetPolicyVersionInput, encoder *httpbinding.Encoder) error {
@@ -12132,10 +11702,8 @@ func (m *awsRestjson1_serializeOpGetRegistrationCode) 
 func (m *awsRestjson1_serializeOpGetRegistrationCode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12169,7 +11737,6 @@ func (m *awsRestjson1_serializeOpGetRegistrationCode) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRegistrationCodeInput(v *GetRegistrationCodeInput, encoder *httpbinding.Encoder) error {
@@ -12190,10 +11757,8 @@ func (m *awsRestjson1_serializeOpGetStatistics) Handle
 func (m *awsRestjson1_serializeOpGetStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12238,7 +11803,6 @@ func (m *awsRestjson1_serializeOpGetStatistics) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStatisticsInput(v *GetStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -12286,10 +11850,8 @@ func (m *awsRestjson1_serializeOpGetThingConnectivityD
 func (m *awsRestjson1_serializeOpGetThingConnectivityData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12327,7 +11889,6 @@ func (m *awsRestjson1_serializeOpGetThingConnectivityD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetThingConnectivityDataInput(v *GetThingConnectivityDataInput, encoder *httpbinding.Encoder) error {
@@ -12357,10 +11918,8 @@ func (m *awsRestjson1_serializeOpGetTopicRule) HandleS
 func (m *awsRestjson1_serializeOpGetTopicRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12398,7 +11957,6 @@ func (m *awsRestjson1_serializeOpGetTopicRule) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTopicRuleInput(v *GetTopicRuleInput, encoder *httpbinding.Encoder) error {
@@ -12428,10 +11986,8 @@ func (m *awsRestjson1_serializeOpGetTopicRuleDestinati
 func (m *awsRestjson1_serializeOpGetTopicRuleDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12469,7 +12025,6 @@ func (m *awsRestjson1_serializeOpGetTopicRuleDestinati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTopicRuleDestinationInput(v *GetTopicRuleDestinationInput, encoder *httpbinding.Encoder) error {
@@ -12499,10 +12054,8 @@ func (m *awsRestjson1_serializeOpGetV2LoggingOptions) 
 func (m *awsRestjson1_serializeOpGetV2LoggingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12536,7 +12089,6 @@ func (m *awsRestjson1_serializeOpGetV2LoggingOptions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetV2LoggingOptionsInput(v *GetV2LoggingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -12557,10 +12109,8 @@ func (m *awsRestjson1_serializeOpListActiveViolations)
 func (m *awsRestjson1_serializeOpListActiveViolations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12598,7 +12148,6 @@ func (m *awsRestjson1_serializeOpListActiveViolations)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListActiveViolationsInput(v *ListActiveViolationsInput, encoder *httpbinding.Encoder) error {
@@ -12647,10 +12196,8 @@ func (m *awsRestjson1_serializeOpListAttachedPolicies)
 func (m *awsRestjson1_serializeOpListAttachedPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12688,7 +12235,6 @@ func (m *awsRestjson1_serializeOpListAttachedPolicies)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAttachedPoliciesInput(v *ListAttachedPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -12730,10 +12276,8 @@ func (m *awsRestjson1_serializeOpListAuditFindings) Ha
 func (m *awsRestjson1_serializeOpListAuditFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12778,7 +12322,6 @@ func (m *awsRestjson1_serializeOpListAuditFindings) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAuditFindingsInput(v *ListAuditFindingsInput, encoder *httpbinding.Encoder) error {
@@ -12848,10 +12391,8 @@ func (m *awsRestjson1_serializeOpListAuditMitigationAc
 func (m *awsRestjson1_serializeOpListAuditMitigationActionsExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12889,7 +12430,6 @@ func (m *awsRestjson1_serializeOpListAuditMitigationAc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAuditMitigationActionsExecutionsInput(v *ListAuditMitigationActionsExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -12930,10 +12470,8 @@ func (m *awsRestjson1_serializeOpListAuditMitigationAc
 func (m *awsRestjson1_serializeOpListAuditMitigationActionsTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12971,7 +12509,6 @@ func (m *awsRestjson1_serializeOpListAuditMitigationAc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAuditMitigationActionsTasksInput(v *ListAuditMitigationActionsTasksInput, encoder *httpbinding.Encoder) error {
@@ -13020,10 +12557,8 @@ func (m *awsRestjson1_serializeOpListAuditSuppressions
 func (m *awsRestjson1_serializeOpListAuditSuppressions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13068,7 +12603,6 @@ func (m *awsRestjson1_serializeOpListAuditSuppressions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAuditSuppressionsInput(v *ListAuditSuppressionsInput, encoder *httpbinding.Encoder) error {
@@ -13123,10 +12657,8 @@ func (m *awsRestjson1_serializeOpListAuditTasks) Handl
 func (m *awsRestjson1_serializeOpListAuditTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13164,7 +12696,6 @@ func (m *awsRestjson1_serializeOpListAuditTasks) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAuditTasksInput(v *ListAuditTasksInput, encoder *httpbinding.Encoder) error {
@@ -13209,10 +12740,8 @@ func (m *awsRestjson1_serializeOpListAuthorizers) Hand
 func (m *awsRestjson1_serializeOpListAuthorizers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13250,7 +12779,6 @@ func (m *awsRestjson1_serializeOpListAuthorizers) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAuthorizersInput(v *ListAuthorizersInput, encoder *httpbinding.Encoder) error {
@@ -13287,10 +12815,8 @@ func (m *awsRestjson1_serializeOpListBillingGroups) Ha
 func (m *awsRestjson1_serializeOpListBillingGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13328,7 +12854,6 @@ func (m *awsRestjson1_serializeOpListBillingGroups) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBillingGroupsInput(v *ListBillingGroupsInput, encoder *httpbinding.Encoder) error {
@@ -13361,10 +12886,8 @@ func (m *awsRestjson1_serializeOpListCACertificates) H
 func (m *awsRestjson1_serializeOpListCACertificates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13402,7 +12925,6 @@ func (m *awsRestjson1_serializeOpListCACertificates) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCACertificatesInput(v *ListCACertificatesInput, encoder *httpbinding.Encoder) error {
@@ -13439,10 +12961,8 @@ func (m *awsRestjson1_serializeOpListCertificateProvid
 func (m *awsRestjson1_serializeOpListCertificateProviders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13480,7 +13000,6 @@ func (m *awsRestjson1_serializeOpListCertificateProvid
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCertificateProvidersInput(v *ListCertificateProvidersInput, encoder *httpbinding.Encoder) error {
@@ -13509,10 +13028,8 @@ func (m *awsRestjson1_serializeOpListCertificates) Han
 func (m *awsRestjson1_serializeOpListCertificates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13550,7 +13067,6 @@ func (m *awsRestjson1_serializeOpListCertificates) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCertificatesInput(v *ListCertificatesInput, encoder *httpbinding.Encoder) error {
@@ -13583,10 +13099,8 @@ func (m *awsRestjson1_serializeOpListCertificatesByCA)
 func (m *awsRestjson1_serializeOpListCertificatesByCA) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13624,7 +13138,6 @@ func (m *awsRestjson1_serializeOpListCertificatesByCA)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCertificatesByCAInput(v *ListCertificatesByCAInput, encoder *httpbinding.Encoder) error {
@@ -13666,10 +13179,8 @@ func (m *awsRestjson1_serializeOpListCommandExecutions
 func (m *awsRestjson1_serializeOpListCommandExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13718,7 +13229,6 @@ func (m *awsRestjson1_serializeOpListCommandExecutions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCommandExecutionsInput(v *ListCommandExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -13793,10 +13303,8 @@ func (m *awsRestjson1_serializeOpListCommands) HandleS
 func (m *awsRestjson1_serializeOpListCommands) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13834,7 +13342,6 @@ func (m *awsRestjson1_serializeOpListCommands) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCommandsInput(v *ListCommandsInput, encoder *httpbinding.Encoder) error {
@@ -13875,10 +13382,8 @@ func (m *awsRestjson1_serializeOpListCustomMetrics) Ha
 func (m *awsRestjson1_serializeOpListCustomMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13916,7 +13421,6 @@ func (m *awsRestjson1_serializeOpListCustomMetrics) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomMetricsInput(v *ListCustomMetricsInput, encoder *httpbinding.Encoder) error {
@@ -13945,10 +13449,8 @@ func (m *awsRestjson1_serializeOpListDetectMitigationA
 func (m *awsRestjson1_serializeOpListDetectMitigationActionsExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13986,7 +13488,6 @@ func (m *awsRestjson1_serializeOpListDetectMitigationA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDetectMitigationActionsExecutionsInput(v *ListDetectMitigationActionsExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -14035,10 +13536,8 @@ func (m *awsRestjson1_serializeOpListDetectMitigationA
 func (m *awsRestjson1_serializeOpListDetectMitigationActionsTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14076,7 +13575,6 @@ func (m *awsRestjson1_serializeOpListDetectMitigationA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDetectMitigationActionsTasksInput(v *ListDetectMitigationActionsTasksInput, encoder *httpbinding.Encoder) error {
@@ -14113,10 +13611,8 @@ func (m *awsRestjson1_serializeOpListDimensions) Handl
 func (m *awsRestjson1_serializeOpListDimensions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14154,7 +13650,6 @@ func (m *awsRestjson1_serializeOpListDimensions) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDimensionsInput(v *ListDimensionsInput, encoder *httpbinding.Encoder) error {
@@ -14183,10 +13678,8 @@ func (m *awsRestjson1_serializeOpListDomainConfigurati
 func (m *awsRestjson1_serializeOpListDomainConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14224,7 +13717,6 @@ func (m *awsRestjson1_serializeOpListDomainConfigurati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainConfigurationsInput(v *ListDomainConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -14257,10 +13749,8 @@ func (m *awsRestjson1_serializeOpListFleetMetrics) Han
 func (m *awsRestjson1_serializeOpListFleetMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14298,7 +13788,6 @@ func (m *awsRestjson1_serializeOpListFleetMetrics) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFleetMetricsInput(v *ListFleetMetricsInput, encoder *httpbinding.Encoder) error {
@@ -14327,10 +13816,8 @@ func (m *awsRestjson1_serializeOpListIndices) HandleSe
 func (m *awsRestjson1_serializeOpListIndices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14368,7 +13855,6 @@ func (m *awsRestjson1_serializeOpListIndices) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIndicesInput(v *ListIndicesInput, encoder *httpbinding.Encoder) error {
@@ -14397,10 +13883,8 @@ func (m *awsRestjson1_serializeOpListJobExecutionsForJ
 func (m *awsRestjson1_serializeOpListJobExecutionsForJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14438,7 +13922,6 @@ func (m *awsRestjson1_serializeOpListJobExecutionsForJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobExecutionsForJobInput(v *ListJobExecutionsForJobInput, encoder *httpbinding.Encoder) error {
@@ -14480,10 +13963,8 @@ func (m *awsRestjson1_serializeOpListJobExecutionsForT
 func (m *awsRestjson1_serializeOpListJobExecutionsForThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14521,7 +14002,6 @@ func (m *awsRestjson1_serializeOpListJobExecutionsForT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobExecutionsForThingInput(v *ListJobExecutionsForThingInput, encoder *httpbinding.Encoder) error {
@@ -14571,10 +14051,8 @@ func (m *awsRestjson1_serializeOpListJobs) HandleSeria
 func (m *awsRestjson1_serializeOpListJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14612,7 +14090,6 @@ func (m *awsRestjson1_serializeOpListJobs) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobsInput(v *ListJobsInput, encoder *httpbinding.Encoder) error {
@@ -14661,10 +14138,8 @@ func (m *awsRestjson1_serializeOpListJobTemplates) Han
 func (m *awsRestjson1_serializeOpListJobTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14702,7 +14177,6 @@ func (m *awsRestjson1_serializeOpListJobTemplates) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobTemplatesInput(v *ListJobTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -14731,10 +14205,8 @@ func (m *awsRestjson1_serializeOpListManagedJobTemplat
 func (m *awsRestjson1_serializeOpListManagedJobTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14772,7 +14244,6 @@ func (m *awsRestjson1_serializeOpListManagedJobTemplat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListManagedJobTemplatesInput(v *ListManagedJobTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -14805,10 +14276,8 @@ func (m *awsRestjson1_serializeOpListMetricValues) Han
 func (m *awsRestjson1_serializeOpListMetricValues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14846,7 +14315,6 @@ func (m *awsRestjson1_serializeOpListMetricValues) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMetricValuesInput(v *ListMetricValuesInput, encoder *httpbinding.Encoder) error {
@@ -14899,10 +14367,8 @@ func (m *awsRestjson1_serializeOpListMitigationActions
 func (m *awsRestjson1_serializeOpListMitigationActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14940,7 +14406,6 @@ func (m *awsRestjson1_serializeOpListMitigationActions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMitigationActionsInput(v *ListMitigationActionsInput, encoder *httpbinding.Encoder) error {
@@ -14973,10 +14438,8 @@ func (m *awsRestjson1_serializeOpListOTAUpdates) Handl
 func (m *awsRestjson1_serializeOpListOTAUpdates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15014,7 +14477,6 @@ func (m *awsRestjson1_serializeOpListOTAUpdates) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOTAUpdatesInput(v *ListOTAUpdatesInput, encoder *httpbinding.Encoder) error {
@@ -15047,10 +14509,8 @@ func (m *awsRestjson1_serializeOpListOutgoingCertifica
 func (m *awsRestjson1_serializeOpListOutgoingCertificates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15088,7 +14548,6 @@ func (m *awsRestjson1_serializeOpListOutgoingCertifica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOutgoingCertificatesInput(v *ListOutgoingCertificatesInput, encoder *httpbinding.Encoder) error {
@@ -15121,10 +14580,8 @@ func (m *awsRestjson1_serializeOpListPackages) HandleS
 func (m *awsRestjson1_serializeOpListPackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15162,7 +14619,6 @@ func (m *awsRestjson1_serializeOpListPackages) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackagesInput(v *ListPackagesInput, encoder *httpbinding.Encoder) error {
@@ -15191,10 +14647,8 @@ func (m *awsRestjson1_serializeOpListPackageVersions) 
 func (m *awsRestjson1_serializeOpListPackageVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15232,7 +14686,6 @@ func (m *awsRestjson1_serializeOpListPackageVersions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackageVersionsInput(v *ListPackageVersionsInput, encoder *httpbinding.Encoder) error {
@@ -15274,10 +14727,8 @@ func (m *awsRestjson1_serializeOpListPolicies) HandleS
 func (m *awsRestjson1_serializeOpListPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15315,7 +14766,6 @@ func (m *awsRestjson1_serializeOpListPolicies) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPoliciesInput(v *ListPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -15348,10 +14798,8 @@ func (m *awsRestjson1_serializeOpListPolicyPrincipals)
 func (m *awsRestjson1_serializeOpListPolicyPrincipals) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15389,7 +14837,6 @@ func (m *awsRestjson1_serializeOpListPolicyPrincipals)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPolicyPrincipalsInput(v *ListPolicyPrincipalsInput, encoder *httpbinding.Encoder) error {
@@ -15427,10 +14874,8 @@ func (m *awsRestjson1_serializeOpListPolicyVersions) H
 func (m *awsRestjson1_serializeOpListPolicyVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15468,7 +14913,6 @@ func (m *awsRestjson1_serializeOpListPolicyVersions) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPolicyVersionsInput(v *ListPolicyVersionsInput, encoder *httpbinding.Encoder) error {
@@ -15498,10 +14942,8 @@ func (m *awsRestjson1_serializeOpListPrincipalPolicies
 func (m *awsRestjson1_serializeOpListPrincipalPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15539,7 +14981,6 @@ func (m *awsRestjson1_serializeOpListPrincipalPolicies
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPrincipalPoliciesInput(v *ListPrincipalPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -15577,10 +15018,8 @@ func (m *awsRestjson1_serializeOpListPrincipalThings) 
 func (m *awsRestjson1_serializeOpListPrincipalThings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15618,7 +15057,6 @@ func (m *awsRestjson1_serializeOpListPrincipalThings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPrincipalThingsInput(v *ListPrincipalThingsInput, encoder *httpbinding.Encoder) error {
@@ -15652,10 +15090,8 @@ func (m *awsRestjson1_serializeOpListPrincipalThingsV2
 func (m *awsRestjson1_serializeOpListPrincipalThingsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15693,7 +15129,6 @@ func (m *awsRestjson1_serializeOpListPrincipalThingsV2
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPrincipalThingsV2Input(v *ListPrincipalThingsV2Input, encoder *httpbinding.Encoder) error {
@@ -15731,10 +15166,8 @@ func (m *awsRestjson1_serializeOpListProvisioningTempl
 func (m *awsRestjson1_serializeOpListProvisioningTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15772,7 +15205,6 @@ func (m *awsRestjson1_serializeOpListProvisioningTempl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProvisioningTemplatesInput(v *ListProvisioningTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -15801,10 +15233,8 @@ func (m *awsRestjson1_serializeOpListProvisioningTempl
 func (m *awsRestjson1_serializeOpListProvisioningTemplateVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15842,7 +15272,6 @@ func (m *awsRestjson1_serializeOpListProvisioningTempl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProvisioningTemplateVersionsInput(v *ListProvisioningTemplateVersionsInput, encoder *httpbinding.Encoder) error {
@@ -15880,10 +15309,8 @@ func (m *awsRestjson1_serializeOpListRelatedResourcesF
 func (m *awsRestjson1_serializeOpListRelatedResourcesForAuditFinding) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15921,7 +15348,6 @@ func (m *awsRestjson1_serializeOpListRelatedResourcesF
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRelatedResourcesForAuditFindingInput(v *ListRelatedResourcesForAuditFindingInput, encoder *httpbinding.Encoder) error {
@@ -15954,10 +15380,8 @@ func (m *awsRestjson1_serializeOpListRoleAliases) Hand
 func (m *awsRestjson1_serializeOpListRoleAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15995,7 +15419,6 @@ func (m *awsRestjson1_serializeOpListRoleAliases) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRoleAliasesInput(v *ListRoleAliasesInput, encoder *httpbinding.Encoder) error {
@@ -16028,10 +15451,8 @@ func (m *awsRestjson1_serializeOpListSbomValidationRes
 func (m *awsRestjson1_serializeOpListSbomValidationResults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16069,7 +15490,6 @@ func (m *awsRestjson1_serializeOpListSbomValidationRes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSbomValidationResultsInput(v *ListSbomValidationResultsInput, encoder *httpbinding.Encoder) error {
@@ -16120,10 +15540,8 @@ func (m *awsRestjson1_serializeOpListScheduledAudits) 
 func (m *awsRestjson1_serializeOpListScheduledAudits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16161,7 +15579,6 @@ func (m *awsRestjson1_serializeOpListScheduledAudits) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListScheduledAuditsInput(v *ListScheduledAuditsInput, encoder *httpbinding.Encoder) error {
@@ -16190,10 +15607,8 @@ func (m *awsRestjson1_serializeOpListSecurityProfiles)
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
@@ -16231,7 +15646,6 @@ func (m *awsRestjson1_serializeOpListSecurityProfiles)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSecurityProfilesInput(v *ListSecurityProfilesInput, encoder *httpbinding.Encoder) error {
@@ -16268,10 +15682,8 @@ func (m *awsRestjson1_serializeOpListSecurityProfilesF
 func (m *awsRestjson1_serializeOpListSecurityProfilesForTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16309,7 +15721,6 @@ func (m *awsRestjson1_serializeOpListSecurityProfilesF
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSecurityProfilesForTargetInput(v *ListSecurityProfilesForTargetInput, encoder *httpbinding.Encoder) error {
@@ -16346,10 +15757,8 @@ func (m *awsRestjson1_serializeOpListStreams) HandleSe
 func (m *awsRestjson1_serializeOpListStreams) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16387,7 +15796,6 @@ func (m *awsRestjson1_serializeOpListStreams) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStreamsInput(v *ListStreamsInput, encoder *httpbinding.Encoder) error {
@@ -16420,10 +15828,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -16461,7 +15867,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -16490,10 +15895,8 @@ func (m *awsRestjson1_serializeOpListTargetsForPolicy)
 func (m *awsRestjson1_serializeOpListTargetsForPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16531,7 +15934,6 @@ func (m *awsRestjson1_serializeOpListTargetsForPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTargetsForPolicyInput(v *ListTargetsForPolicyInput, encoder *httpbinding.Encoder) error {
@@ -16569,10 +15971,8 @@ func (m *awsRestjson1_serializeOpListTargetsForSecurit
 func (m *awsRestjson1_serializeOpListTargetsForSecurityProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16610,7 +16010,6 @@ func (m *awsRestjson1_serializeOpListTargetsForSecurit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTargetsForSecurityProfileInput(v *ListTargetsForSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -16648,10 +16047,8 @@ func (m *awsRestjson1_serializeOpListThingGroups) Hand
 func (m *awsRestjson1_serializeOpListThingGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16689,7 +16086,6 @@ func (m *awsRestjson1_serializeOpListThingGroups) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingGroupsInput(v *ListThingGroupsInput, encoder *httpbinding.Encoder) error {
@@ -16730,10 +16126,8 @@ func (m *awsRestjson1_serializeOpListThingGroupsForThi
 func (m *awsRestjson1_serializeOpListThingGroupsForThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16771,7 +16165,6 @@ func (m *awsRestjson1_serializeOpListThingGroupsForThi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingGroupsForThingInput(v *ListThingGroupsForThingInput, encoder *httpbinding.Encoder) error {
@@ -16809,10 +16202,8 @@ func (m *awsRestjson1_serializeOpListThingPrincipals) 
 func (m *awsRestjson1_serializeOpListThingPrincipals) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16850,7 +16241,6 @@ func (m *awsRestjson1_serializeOpListThingPrincipals) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingPrincipalsInput(v *ListThingPrincipalsInput, encoder *httpbinding.Encoder) error {
@@ -16888,10 +16278,8 @@ func (m *awsRestjson1_serializeOpListThingPrincipalsV2
 func (m *awsRestjson1_serializeOpListThingPrincipalsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -16929,7 +16317,6 @@ func (m *awsRestjson1_serializeOpListThingPrincipalsV2
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingPrincipalsV2Input(v *ListThingPrincipalsV2Input, encoder *httpbinding.Encoder) error {
@@ -16971,10 +16358,8 @@ func (m *awsRestjson1_serializeOpListThingRegistration
 func (m *awsRestjson1_serializeOpListThingRegistrationTaskReports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17012,7 +16397,6 @@ func (m *awsRestjson1_serializeOpListThingRegistration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingRegistrationTaskReportsInput(v *ListThingRegistrationTaskReportsInput, encoder *httpbinding.Encoder) error {
@@ -17054,10 +16438,8 @@ func (m *awsRestjson1_serializeOpListThingRegistration
 func (m *awsRestjson1_serializeOpListThingRegistrationTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17095,7 +16477,6 @@ func (m *awsRestjson1_serializeOpListThingRegistration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingRegistrationTasksInput(v *ListThingRegistrationTasksInput, encoder *httpbinding.Encoder) error {
@@ -17128,10 +16509,8 @@ func (m *awsRestjson1_serializeOpListThings) HandleSer
 func (m *awsRestjson1_serializeOpListThings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17169,7 +16548,6 @@ func (m *awsRestjson1_serializeOpListThings) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingsInput(v *ListThingsInput, encoder *httpbinding.Encoder) error {
@@ -17214,10 +16592,8 @@ func (m *awsRestjson1_serializeOpListThingsInBillingGr
 func (m *awsRestjson1_serializeOpListThingsInBillingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17255,7 +16631,6 @@ func (m *awsRestjson1_serializeOpListThingsInBillingGr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingsInBillingGroupInput(v *ListThingsInBillingGroupInput, encoder *httpbinding.Encoder) error {
@@ -17293,10 +16668,8 @@ func (m *awsRestjson1_serializeOpListThingsInThingGrou
 func (m *awsRestjson1_serializeOpListThingsInThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17334,7 +16707,6 @@ func (m *awsRestjson1_serializeOpListThingsInThingGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingsInThingGroupInput(v *ListThingsInThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -17376,10 +16748,8 @@ func (m *awsRestjson1_serializeOpListThingTypes) Handl
 func (m *awsRestjson1_serializeOpListThingTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17417,7 +16787,6 @@ func (m *awsRestjson1_serializeOpListThingTypes) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThingTypesInput(v *ListThingTypesInput, encoder *httpbinding.Encoder) error {
@@ -17450,10 +16819,8 @@ func (m *awsRestjson1_serializeOpListTopicRuleDestinat
 func (m *awsRestjson1_serializeOpListTopicRuleDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17491,7 +16858,6 @@ func (m *awsRestjson1_serializeOpListTopicRuleDestinat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTopicRuleDestinationsInput(v *ListTopicRuleDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -17520,10 +16886,8 @@ func (m *awsRestjson1_serializeOpListTopicRules) Handl
 func (m *awsRestjson1_serializeOpListTopicRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17561,7 +16925,6 @@ func (m *awsRestjson1_serializeOpListTopicRules) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTopicRulesInput(v *ListTopicRulesInput, encoder *httpbinding.Encoder) error {
@@ -17598,10 +16961,8 @@ func (m *awsRestjson1_serializeOpListV2LoggingLevels) 
 func (m *awsRestjson1_serializeOpListV2LoggingLevels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17639,7 +17000,6 @@ func (m *awsRestjson1_serializeOpListV2LoggingLevels) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListV2LoggingLevelsInput(v *ListV2LoggingLevelsInput, encoder *httpbinding.Encoder) error {
@@ -17672,10 +17032,8 @@ func (m *awsRestjson1_serializeOpListViolationEvents) 
 func (m *awsRestjson1_serializeOpListViolationEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17713,7 +17071,6 @@ func (m *awsRestjson1_serializeOpListViolationEvents) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListViolationEventsInput(v *ListViolationEventsInput, encoder *httpbinding.Encoder) error {
@@ -17770,10 +17127,8 @@ func (m *awsRestjson1_serializeOpPutVerificationStateO
 func (m *awsRestjson1_serializeOpPutVerificationStateOnViolation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17822,7 +17177,6 @@ func (m *awsRestjson1_serializeOpPutVerificationStateO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVerificationStateOnViolationInput(v *PutVerificationStateOnViolationInput, encoder *httpbinding.Encoder) error {
@@ -17869,10 +17223,8 @@ func (m *awsRestjson1_serializeOpRegisterCACertificate
 func (m *awsRestjson1_serializeOpRegisterCACertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -17921,7 +17273,6 @@ func (m *awsRestjson1_serializeOpRegisterCACertificate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterCACertificateInput(v *RegisterCACertificateInput, encoder *httpbinding.Encoder) error {
@@ -17986,10 +17337,8 @@ func (m *awsRestjson1_serializeOpRegisterCertificate) 
 func (m *awsRestjson1_serializeOpRegisterCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18038,7 +17387,6 @@ func (m *awsRestjson1_serializeOpRegisterCertificate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterCertificateInput(v *RegisterCertificateInput, encoder *httpbinding.Encoder) error {
@@ -18085,10 +17433,8 @@ func (m *awsRestjson1_serializeOpRegisterCertificateWi
 func (m *awsRestjson1_serializeOpRegisterCertificateWithoutCA) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18133,7 +17479,6 @@ func (m *awsRestjson1_serializeOpRegisterCertificateWi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterCertificateWithoutCAInput(v *RegisterCertificateWithoutCAInput, encoder *httpbinding.Encoder) error {
@@ -18171,10 +17516,8 @@ func (m *awsRestjson1_serializeOpRegisterThing) Handle
 func (m *awsRestjson1_serializeOpRegisterThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18219,7 +17562,6 @@ func (m *awsRestjson1_serializeOpRegisterThing) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterThingInput(v *RegisterThingInput, encoder *httpbinding.Encoder) error {
@@ -18259,10 +17601,8 @@ func (m *awsRestjson1_serializeOpRejectCertificateTran
 func (m *awsRestjson1_serializeOpRejectCertificateTransfer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18311,7 +17651,6 @@ func (m *awsRestjson1_serializeOpRejectCertificateTran
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectCertificateTransferInput(v *RejectCertificateTransferInput, encoder *httpbinding.Encoder) error {
@@ -18353,10 +17692,8 @@ func (m *awsRestjson1_serializeOpRemoveThingFromBillin
 func (m *awsRestjson1_serializeOpRemoveThingFromBillingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18401,7 +17738,6 @@ func (m *awsRestjson1_serializeOpRemoveThingFromBillin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveThingFromBillingGroupInput(v *RemoveThingFromBillingGroupInput, encoder *httpbinding.Encoder) error {
@@ -18449,10 +17785,8 @@ func (m *awsRestjson1_serializeOpRemoveThingFromThingG
 func (m *awsRestjson1_serializeOpRemoveThingFromThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18497,7 +17831,6 @@ func (m *awsRestjson1_serializeOpRemoveThingFromThingG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveThingFromThingGroupInput(v *RemoveThingFromThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -18545,10 +17878,8 @@ func (m *awsRestjson1_serializeOpReplaceTopicRule) Han
 func (m *awsRestjson1_serializeOpReplaceTopicRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18610,7 +17941,6 @@ func (m *awsRestjson1_serializeOpReplaceTopicRule) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsReplaceTopicRuleInput(v *ReplaceTopicRuleInput, encoder *httpbinding.Encoder) error {
@@ -18640,10 +17970,8 @@ func (m *awsRestjson1_serializeOpSearchIndex) HandleSe
 func (m *awsRestjson1_serializeOpSearchIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18688,7 +18016,6 @@ func (m *awsRestjson1_serializeOpSearchIndex) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchIndexInput(v *SearchIndexInput, encoder *httpbinding.Encoder) error {
@@ -18741,10 +18068,8 @@ func (m *awsRestjson1_serializeOpSetDefaultAuthorizer)
 func (m *awsRestjson1_serializeOpSetDefaultAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18789,7 +18114,6 @@ func (m *awsRestjson1_serializeOpSetDefaultAuthorizer)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetDefaultAuthorizerInput(v *SetDefaultAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -18822,10 +18146,8 @@ func (m *awsRestjson1_serializeOpSetDefaultPolicyVersi
 func (m *awsRestjson1_serializeOpSetDefaultPolicyVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18863,7 +18185,6 @@ func (m *awsRestjson1_serializeOpSetDefaultPolicyVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetDefaultPolicyVersionInput(v *SetDefaultPolicyVersionInput, encoder *httpbinding.Encoder) error {
@@ -18902,10 +18223,8 @@ func (m *awsRestjson1_serializeOpSetLoggingOptions) Ha
 func (m *awsRestjson1_serializeOpSetLoggingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -18963,7 +18282,6 @@ func (m *awsRestjson1_serializeOpSetLoggingOptions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetLoggingOptionsInput(v *SetLoggingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -18984,10 +18302,8 @@ func (m *awsRestjson1_serializeOpSetV2LoggingLevel) Ha
 func (m *awsRestjson1_serializeOpSetV2LoggingLevel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19032,7 +18348,6 @@ func (m *awsRestjson1_serializeOpSetV2LoggingLevel) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetV2LoggingLevelInput(v *SetV2LoggingLevelInput, encoder *httpbinding.Encoder) error {
@@ -19072,10 +18387,8 @@ func (m *awsRestjson1_serializeOpSetV2LoggingOptions) 
 func (m *awsRestjson1_serializeOpSetV2LoggingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19120,7 +18433,6 @@ func (m *awsRestjson1_serializeOpSetV2LoggingOptions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetV2LoggingOptionsInput(v *SetV2LoggingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -19163,10 +18475,8 @@ func (m *awsRestjson1_serializeOpStartAuditMitigationA
 func (m *awsRestjson1_serializeOpStartAuditMitigationActionsTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19215,7 +18525,6 @@ func (m *awsRestjson1_serializeOpStartAuditMitigationA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAuditMitigationActionsTaskInput(v *StartAuditMitigationActionsTaskInput, encoder *httpbinding.Encoder) error {
@@ -19271,10 +18580,8 @@ func (m *awsRestjson1_serializeOpStartDetectMitigation
 func (m *awsRestjson1_serializeOpStartDetectMitigationActionsTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19323,7 +18630,6 @@ func (m *awsRestjson1_serializeOpStartDetectMitigation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDetectMitigationActionsTaskInput(v *StartDetectMitigationActionsTaskInput, encoder *httpbinding.Encoder) error {
@@ -19396,10 +18702,8 @@ func (m *awsRestjson1_serializeOpStartOnDemandAuditTas
 func (m *awsRestjson1_serializeOpStartOnDemandAuditTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19444,7 +18748,6 @@ func (m *awsRestjson1_serializeOpStartOnDemandAuditTas
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartOnDemandAuditTaskInput(v *StartOnDemandAuditTaskInput, encoder *httpbinding.Encoder) error {
@@ -19479,10 +18782,8 @@ func (m *awsRestjson1_serializeOpStartThingRegistratio
 func (m *awsRestjson1_serializeOpStartThingRegistrationTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19527,7 +18828,6 @@ func (m *awsRestjson1_serializeOpStartThingRegistratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartThingRegistrationTaskInput(v *StartThingRegistrationTaskInput, encoder *httpbinding.Encoder) error {
@@ -19575,10 +18875,8 @@ func (m *awsRestjson1_serializeOpStopThingRegistration
 func (m *awsRestjson1_serializeOpStopThingRegistrationTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19616,7 +18914,6 @@ func (m *awsRestjson1_serializeOpStopThingRegistration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopThingRegistrationTaskInput(v *StopThingRegistrationTaskInput, encoder *httpbinding.Encoder) error {
@@ -19646,10 +18943,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -19694,7 +18989,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -19734,10 +19028,8 @@ func (m *awsRestjson1_serializeOpTestAuthorization) Ha
 func (m *awsRestjson1_serializeOpTestAuthorization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19786,7 +19078,6 @@ func (m *awsRestjson1_serializeOpTestAuthorization) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestAuthorizationInput(v *TestAuthorizationInput, encoder *httpbinding.Encoder) error {
@@ -19849,10 +19140,8 @@ func (m *awsRestjson1_serializeOpTestInvokeAuthorizer)
 func (m *awsRestjson1_serializeOpTestInvokeAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -19901,7 +19190,6 @@ func (m *awsRestjson1_serializeOpTestInvokeAuthorizer)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestInvokeAuthorizerInput(v *TestInvokeAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -19969,10 +19257,8 @@ func (m *awsRestjson1_serializeOpTransferCertificate) 
 func (m *awsRestjson1_serializeOpTransferCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20021,7 +19307,6 @@ func (m *awsRestjson1_serializeOpTransferCertificate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTransferCertificateInput(v *TransferCertificateInput, encoder *httpbinding.Encoder) error {
@@ -20067,10 +19352,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -20115,7 +19398,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -20155,10 +19437,8 @@ func (m *awsRestjson1_serializeOpUpdateAccountAuditCon
 func (m *awsRestjson1_serializeOpUpdateAccountAuditConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20203,7 +19483,6 @@ func (m *awsRestjson1_serializeOpUpdateAccountAuditCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountAuditConfigurationInput(v *UpdateAccountAuditConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -20250,10 +19529,8 @@ func (m *awsRestjson1_serializeOpUpdateAuditSuppressio
 func (m *awsRestjson1_serializeOpUpdateAuditSuppression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20298,7 +19575,6 @@ func (m *awsRestjson1_serializeOpUpdateAuditSuppressio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAuditSuppressionInput(v *UpdateAuditSuppressionInput, encoder *httpbinding.Encoder) error {
@@ -20353,10 +19629,8 @@ func (m *awsRestjson1_serializeOpUpdateAuthorizer) Han
 func (m *awsRestjson1_serializeOpUpdateAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20405,7 +19679,6 @@ func (m *awsRestjson1_serializeOpUpdateAuthorizer) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAuthorizerInput(v *UpdateAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -20469,10 +19742,8 @@ func (m *awsRestjson1_serializeOpUpdateBillingGroup) H
 func (m *awsRestjson1_serializeOpUpdateBillingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20521,7 +19792,6 @@ func (m *awsRestjson1_serializeOpUpdateBillingGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBillingGroupInput(v *UpdateBillingGroupInput, encoder *httpbinding.Encoder) error {
@@ -20570,10 +19840,8 @@ func (m *awsRestjson1_serializeOpUpdateCACertificate) 
 func (m *awsRestjson1_serializeOpUpdateCACertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20622,7 +19890,6 @@ func (m *awsRestjson1_serializeOpUpdateCACertificate) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCACertificateInput(v *UpdateCACertificateInput, encoder *httpbinding.Encoder) error {
@@ -20679,10 +19946,8 @@ func (m *awsRestjson1_serializeOpUpdateCertificate) Ha
 func (m *awsRestjson1_serializeOpUpdateCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20720,7 +19985,6 @@ func (m *awsRestjson1_serializeOpUpdateCertificate) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCertificateInput(v *UpdateCertificateInput, encoder *httpbinding.Encoder) error {
@@ -20754,10 +20018,8 @@ func (m *awsRestjson1_serializeOpUpdateCertificateProv
 func (m *awsRestjson1_serializeOpUpdateCertificateProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20806,7 +20068,6 @@ func (m *awsRestjson1_serializeOpUpdateCertificateProv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCertificateProviderInput(v *UpdateCertificateProviderInput, encoder *httpbinding.Encoder) error {
@@ -20855,10 +20116,8 @@ func (m *awsRestjson1_serializeOpUpdateCommand) Handle
 func (m *awsRestjson1_serializeOpUpdateCommand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -20907,7 +20166,6 @@ func (m *awsRestjson1_serializeOpUpdateCommand) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCommandInput(v *UpdateCommandInput, encoder *httpbinding.Encoder) error {
@@ -20959,10 +20217,8 @@ func (m *awsRestjson1_serializeOpUpdateCustomMetric) H
 func (m *awsRestjson1_serializeOpUpdateCustomMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21011,7 +20267,6 @@ func (m *awsRestjson1_serializeOpUpdateCustomMetric) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCustomMetricInput(v *UpdateCustomMetricInput, encoder *httpbinding.Encoder) error {
@@ -21053,10 +20308,8 @@ func (m *awsRestjson1_serializeOpUpdateDimension) Hand
 func (m *awsRestjson1_serializeOpUpdateDimension) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21105,7 +20358,6 @@ func (m *awsRestjson1_serializeOpUpdateDimension) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDimensionInput(v *UpdateDimensionInput, encoder *httpbinding.Encoder) error {
@@ -21149,10 +20401,8 @@ func (m *awsRestjson1_serializeOpUpdateDomainConfigura
 func (m *awsRestjson1_serializeOpUpdateDomainConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21201,7 +20451,6 @@ func (m *awsRestjson1_serializeOpUpdateDomainConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDomainConfigurationInput(v *UpdateDomainConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -21286,10 +20535,8 @@ func (m *awsRestjson1_serializeOpUpdateDynamicThingGro
 func (m *awsRestjson1_serializeOpUpdateDynamicThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21338,7 +20585,6 @@ func (m *awsRestjson1_serializeOpUpdateDynamicThingGro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDynamicThingGroupInput(v *UpdateDynamicThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -21402,10 +20648,8 @@ func (m *awsRestjson1_serializeOpUpdateEventConfigurat
 func (m *awsRestjson1_serializeOpUpdateEventConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21450,7 +20694,6 @@ func (m *awsRestjson1_serializeOpUpdateEventConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventConfigurationsInput(v *UpdateEventConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -21485,10 +20728,8 @@ func (m *awsRestjson1_serializeOpUpdateFleetMetric) Ha
 func (m *awsRestjson1_serializeOpUpdateFleetMetric) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21537,7 +20778,6 @@ func (m *awsRestjson1_serializeOpUpdateFleetMetric) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFleetMetricInput(v *UpdateFleetMetricInput, encoder *httpbinding.Encoder) error {
@@ -21621,10 +20861,8 @@ func (m *awsRestjson1_serializeOpUpdateIndexingConfigu
 func (m *awsRestjson1_serializeOpUpdateIndexingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21669,7 +20907,6 @@ func (m *awsRestjson1_serializeOpUpdateIndexingConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIndexingConfigurationInput(v *UpdateIndexingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -21711,10 +20948,8 @@ func (m *awsRestjson1_serializeOpUpdateJob) HandleSeri
 func (m *awsRestjson1_serializeOpUpdateJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21763,7 +20998,6 @@ func (m *awsRestjson1_serializeOpUpdateJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateJobInput(v *UpdateJobInput, encoder *httpbinding.Encoder) error {
@@ -21844,10 +21078,8 @@ func (m *awsRestjson1_serializeOpUpdateMitigationActio
 func (m *awsRestjson1_serializeOpUpdateMitigationAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21896,7 +21128,6 @@ func (m *awsRestjson1_serializeOpUpdateMitigationActio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMitigationActionInput(v *UpdateMitigationActionInput, encoder *httpbinding.Encoder) error {
@@ -21945,10 +21176,8 @@ func (m *awsRestjson1_serializeOpUpdatePackage) Handle
 func (m *awsRestjson1_serializeOpUpdatePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -21997,7 +21226,6 @@ func (m *awsRestjson1_serializeOpUpdatePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageInput(v *UpdatePackageInput, encoder *httpbinding.Encoder) error {
@@ -22053,10 +21281,8 @@ func (m *awsRestjson1_serializeOpUpdatePackageConfigur
 func (m *awsRestjson1_serializeOpUpdatePackageConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22105,7 +21331,6 @@ func (m *awsRestjson1_serializeOpUpdatePackageConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageConfigurationInput(v *UpdatePackageConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -22144,10 +21369,8 @@ func (m *awsRestjson1_serializeOpUpdatePackageVersion)
 func (m *awsRestjson1_serializeOpUpdatePackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22196,7 +21419,6 @@ func (m *awsRestjson1_serializeOpUpdatePackageVersion)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageVersionInput(v *UpdatePackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -22275,10 +21497,8 @@ func (m *awsRestjson1_serializeOpUpdateProvisioningTem
 func (m *awsRestjson1_serializeOpUpdateProvisioningTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22327,7 +21547,6 @@ func (m *awsRestjson1_serializeOpUpdateProvisioningTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProvisioningTemplateInput(v *UpdateProvisioningTemplateInput, encoder *httpbinding.Encoder) error {
@@ -22396,10 +21615,8 @@ func (m *awsRestjson1_serializeOpUpdateRoleAlias) Hand
 func (m *awsRestjson1_serializeOpUpdateRoleAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22448,7 +21665,6 @@ func (m *awsRestjson1_serializeOpUpdateRoleAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoleAliasInput(v *UpdateRoleAliasInput, encoder *httpbinding.Encoder) error {
@@ -22495,10 +21711,8 @@ func (m *awsRestjson1_serializeOpUpdateScheduledAudit)
 func (m *awsRestjson1_serializeOpUpdateScheduledAudit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22547,7 +21761,6 @@ func (m *awsRestjson1_serializeOpUpdateScheduledAudit)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateScheduledAuditInput(v *UpdateScheduledAuditInput, encoder *httpbinding.Encoder) error {
@@ -22606,10 +21819,8 @@ func (m *awsRestjson1_serializeOpUpdateSecurityProfile
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
@@ -22658,7 +21869,6 @@ func (m *awsRestjson1_serializeOpUpdateSecurityProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSecurityProfileInput(v *UpdateSecurityProfileInput, encoder *httpbinding.Encoder) error {
@@ -22759,10 +21969,8 @@ func (m *awsRestjson1_serializeOpUpdateStream) HandleS
 func (m *awsRestjson1_serializeOpUpdateStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22811,7 +22019,6 @@ func (m *awsRestjson1_serializeOpUpdateStream) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateStreamInput(v *UpdateStreamInput, encoder *httpbinding.Encoder) error {
@@ -22865,10 +22072,8 @@ func (m *awsRestjson1_serializeOpUpdateThing) HandleSe
 func (m *awsRestjson1_serializeOpUpdateThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -22917,7 +22122,6 @@ func (m *awsRestjson1_serializeOpUpdateThing) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThingInput(v *UpdateThingInput, encoder *httpbinding.Encoder) error {
@@ -22976,10 +22180,8 @@ func (m *awsRestjson1_serializeOpUpdateThingGroup) Han
 func (m *awsRestjson1_serializeOpUpdateThingGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23028,7 +22230,6 @@ func (m *awsRestjson1_serializeOpUpdateThingGroup) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThingGroupInput(v *UpdateThingGroupInput, encoder *httpbinding.Encoder) error {
@@ -23077,10 +22278,8 @@ func (m *awsRestjson1_serializeOpUpdateThingGroupsForT
 func (m *awsRestjson1_serializeOpUpdateThingGroupsForThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23125,7 +22324,6 @@ func (m *awsRestjson1_serializeOpUpdateThingGroupsForT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThingGroupsForThingInput(v *UpdateThingGroupsForThingInput, encoder *httpbinding.Encoder) error {
@@ -23177,10 +22375,8 @@ func (m *awsRestjson1_serializeOpUpdateThingType) Hand
 func (m *awsRestjson1_serializeOpUpdateThingType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23229,7 +22425,6 @@ func (m *awsRestjson1_serializeOpUpdateThingType) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThingTypeInput(v *UpdateThingTypeInput, encoder *httpbinding.Encoder) error {
@@ -23273,10 +22468,8 @@ func (m *awsRestjson1_serializeOpUpdateTopicRuleDestin
 func (m *awsRestjson1_serializeOpUpdateTopicRuleDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23321,7 +22514,6 @@ func (m *awsRestjson1_serializeOpUpdateTopicRuleDestin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTopicRuleDestinationInput(v *UpdateTopicRuleDestinationInput, encoder *httpbinding.Encoder) error {
@@ -23359,10 +22551,8 @@ func (m *awsRestjson1_serializeOpValidateSecurityProfi
 func (m *awsRestjson1_serializeOpValidateSecurityProfileBehaviors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -23407,7 +22597,6 @@ func (m *awsRestjson1_serializeOpValidateSecurityProfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsValidateSecurityProfileBehaviorsInput(v *ValidateSecurityProfileBehaviorsInput, encoder *httpbinding.Encoder) error {
