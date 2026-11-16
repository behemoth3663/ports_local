--- vendor/github.com/aws/aws-sdk-go-v2/service/backup/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/backup/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAssociateBackupVaultM
 func (m *awsRestjson1_serializeOpAssociateBackupVaultMpaApprovalTeam) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -78,7 +75,6 @@ func (m *awsRestjson1_serializeOpAssociateBackupVaultM
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateBackupVaultMpaApprovalTeamInput(v *AssociateBackupVaultMpaApprovalTeamInput, encoder *httpbinding.Encoder) error {
@@ -125,10 +121,8 @@ func (m *awsRestjson1_serializeOpCancelLegalHold) Hand
 func (m *awsRestjson1_serializeOpCancelLegalHold) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -166,7 +160,6 @@ func (m *awsRestjson1_serializeOpCancelLegalHold) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelLegalHoldInput(v *CancelLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -204,10 +197,8 @@ func (m *awsRestjson1_serializeOpCreateBackupPlan) Han
 func (m *awsRestjson1_serializeOpCreateBackupPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -252,7 +243,6 @@ func (m *awsRestjson1_serializeOpCreateBackupPlan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackupPlanInput(v *CreateBackupPlanInput, encoder *httpbinding.Encoder) error {
@@ -299,10 +289,8 @@ func (m *awsRestjson1_serializeOpCreateBackupSelection
 func (m *awsRestjson1_serializeOpCreateBackupSelection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -351,7 +339,6 @@ func (m *awsRestjson1_serializeOpCreateBackupSelection
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackupSelectionInput(v *CreateBackupSelectionInput, encoder *httpbinding.Encoder) error {
@@ -400,10 +387,8 @@ func (m *awsRestjson1_serializeOpCreateBackupVault) Ha
 func (m *awsRestjson1_serializeOpCreateBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -452,7 +437,6 @@ func (m *awsRestjson1_serializeOpCreateBackupVault) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackupVaultInput(v *CreateBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -506,10 +490,8 @@ func (m *awsRestjson1_serializeOpCreateFramework) Hand
 func (m *awsRestjson1_serializeOpCreateFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -554,7 +536,6 @@ func (m *awsRestjson1_serializeOpCreateFramework) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFrameworkInput(v *CreateFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -611,10 +592,8 @@ func (m *awsRestjson1_serializeOpCreateLegalHold) Hand
 func (m *awsRestjson1_serializeOpCreateLegalHold) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -659,7 +638,6 @@ func (m *awsRestjson1_serializeOpCreateLegalHold) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLegalHoldInput(v *CreateLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -716,10 +694,8 @@ func (m *awsRestjson1_serializeOpCreateLogicallyAirGap
 func (m *awsRestjson1_serializeOpCreateLogicallyAirGappedBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -768,7 +744,6 @@ func (m *awsRestjson1_serializeOpCreateLogicallyAirGap
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLogicallyAirGappedBackupVaultInput(v *CreateLogicallyAirGappedBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -827,10 +802,8 @@ func (m *awsRestjson1_serializeOpCreateReportPlan) Han
 func (m *awsRestjson1_serializeOpCreateReportPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -875,7 +848,6 @@ func (m *awsRestjson1_serializeOpCreateReportPlan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateReportPlanInput(v *CreateReportPlanInput, encoder *httpbinding.Encoder) error {
@@ -939,10 +911,8 @@ func (m *awsRestjson1_serializeOpCreateRestoreAccessBa
 func (m *awsRestjson1_serializeOpCreateRestoreAccessBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -987,7 +957,6 @@ func (m *awsRestjson1_serializeOpCreateRestoreAccessBa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRestoreAccessBackupVaultInput(v *CreateRestoreAccessBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -1042,10 +1011,8 @@ func (m *awsRestjson1_serializeOpCreateRestoreTestingP
 func (m *awsRestjson1_serializeOpCreateRestoreTestingPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1090,7 +1057,6 @@ func (m *awsRestjson1_serializeOpCreateRestoreTestingP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRestoreTestingPlanInput(v *CreateRestoreTestingPlanInput, encoder *httpbinding.Encoder) error {
@@ -1137,10 +1103,8 @@ func (m *awsRestjson1_serializeOpCreateRestoreTestingS
 func (m *awsRestjson1_serializeOpCreateRestoreTestingSelection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1189,7 +1153,6 @@ func (m *awsRestjson1_serializeOpCreateRestoreTestingS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRestoreTestingSelectionInput(v *CreateRestoreTestingSelectionInput, encoder *httpbinding.Encoder) error {
@@ -1238,10 +1201,8 @@ func (m *awsRestjson1_serializeOpDeleteBackupPlan) Han
 func (m *awsRestjson1_serializeOpDeleteBackupPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1279,7 +1240,6 @@ func (m *awsRestjson1_serializeOpDeleteBackupPlan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackupPlanInput(v *DeleteBackupPlanInput, encoder *httpbinding.Encoder) error {
@@ -1309,10 +1269,8 @@ func (m *awsRestjson1_serializeOpDeleteBackupSelection
 func (m *awsRestjson1_serializeOpDeleteBackupSelection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1350,7 +1308,6 @@ func (m *awsRestjson1_serializeOpDeleteBackupSelection
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackupSelectionInput(v *DeleteBackupSelectionInput, encoder *httpbinding.Encoder) error {
@@ -1389,10 +1346,8 @@ func (m *awsRestjson1_serializeOpDeleteBackupVault) Ha
 func (m *awsRestjson1_serializeOpDeleteBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1430,7 +1385,6 @@ func (m *awsRestjson1_serializeOpDeleteBackupVault) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackupVaultInput(v *DeleteBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -1460,10 +1414,8 @@ func (m *awsRestjson1_serializeOpDeleteBackupVaultAcce
 func (m *awsRestjson1_serializeOpDeleteBackupVaultAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1501,7 +1453,6 @@ func (m *awsRestjson1_serializeOpDeleteBackupVaultAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackupVaultAccessPolicyInput(v *DeleteBackupVaultAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1531,10 +1482,8 @@ func (m *awsRestjson1_serializeOpDeleteBackupVaultLock
 func (m *awsRestjson1_serializeOpDeleteBackupVaultLockConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1572,7 +1521,6 @@ func (m *awsRestjson1_serializeOpDeleteBackupVaultLock
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackupVaultLockConfigurationInput(v *DeleteBackupVaultLockConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1602,10 +1550,8 @@ func (m *awsRestjson1_serializeOpDeleteBackupVaultNoti
 func (m *awsRestjson1_serializeOpDeleteBackupVaultNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1643,7 +1589,6 @@ func (m *awsRestjson1_serializeOpDeleteBackupVaultNoti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackupVaultNotificationsInput(v *DeleteBackupVaultNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -1673,10 +1618,8 @@ func (m *awsRestjson1_serializeOpDeleteFramework) Hand
 func (m *awsRestjson1_serializeOpDeleteFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1714,7 +1657,6 @@ func (m *awsRestjson1_serializeOpDeleteFramework) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFrameworkInput(v *DeleteFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -1744,10 +1686,8 @@ func (m *awsRestjson1_serializeOpDeleteRecoveryPoint) 
 func (m *awsRestjson1_serializeOpDeleteRecoveryPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1785,7 +1725,6 @@ func (m *awsRestjson1_serializeOpDeleteRecoveryPoint) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRecoveryPointInput(v *DeleteRecoveryPointInput, encoder *httpbinding.Encoder) error {
@@ -1824,10 +1763,8 @@ func (m *awsRestjson1_serializeOpDeleteReportPlan) Han
 func (m *awsRestjson1_serializeOpDeleteReportPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1865,7 +1802,6 @@ func (m *awsRestjson1_serializeOpDeleteReportPlan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteReportPlanInput(v *DeleteReportPlanInput, encoder *httpbinding.Encoder) error {
@@ -1895,10 +1831,8 @@ func (m *awsRestjson1_serializeOpDeleteRestoreTestingP
 func (m *awsRestjson1_serializeOpDeleteRestoreTestingPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1936,7 +1870,6 @@ func (m *awsRestjson1_serializeOpDeleteRestoreTestingP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRestoreTestingPlanInput(v *DeleteRestoreTestingPlanInput, encoder *httpbinding.Encoder) error {
@@ -1966,10 +1899,8 @@ func (m *awsRestjson1_serializeOpDeleteRestoreTestingS
 func (m *awsRestjson1_serializeOpDeleteRestoreTestingSelection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2007,7 +1938,6 @@ func (m *awsRestjson1_serializeOpDeleteRestoreTestingS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRestoreTestingSelectionInput(v *DeleteRestoreTestingSelectionInput, encoder *httpbinding.Encoder) error {
@@ -2046,10 +1976,8 @@ func (m *awsRestjson1_serializeOpDescribeBackupJob) Ha
 func (m *awsRestjson1_serializeOpDescribeBackupJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2087,7 +2015,6 @@ func (m *awsRestjson1_serializeOpDescribeBackupJob) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBackupJobInput(v *DescribeBackupJobInput, encoder *httpbinding.Encoder) error {
@@ -2117,10 +2044,8 @@ func (m *awsRestjson1_serializeOpDescribeBackupVault) 
 func (m *awsRestjson1_serializeOpDescribeBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2158,7 +2083,6 @@ func (m *awsRestjson1_serializeOpDescribeBackupVault) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBackupVaultInput(v *DescribeBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -2192,10 +2116,8 @@ func (m *awsRestjson1_serializeOpDescribeCopyJob) Hand
 func (m *awsRestjson1_serializeOpDescribeCopyJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2233,7 +2155,6 @@ func (m *awsRestjson1_serializeOpDescribeCopyJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeCopyJobInput(v *DescribeCopyJobInput, encoder *httpbinding.Encoder) error {
@@ -2263,10 +2184,8 @@ func (m *awsRestjson1_serializeOpDescribeFramework) Ha
 func (m *awsRestjson1_serializeOpDescribeFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2304,7 +2223,6 @@ func (m *awsRestjson1_serializeOpDescribeFramework) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeFrameworkInput(v *DescribeFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -2334,10 +2252,8 @@ func (m *awsRestjson1_serializeOpDescribeGlobalSetting
 func (m *awsRestjson1_serializeOpDescribeGlobalSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2371,7 +2287,6 @@ func (m *awsRestjson1_serializeOpDescribeGlobalSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeGlobalSettingsInput(v *DescribeGlobalSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2392,10 +2307,8 @@ func (m *awsRestjson1_serializeOpDescribeProtectedReso
 func (m *awsRestjson1_serializeOpDescribeProtectedResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2433,7 +2346,6 @@ func (m *awsRestjson1_serializeOpDescribeProtectedReso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeProtectedResourceInput(v *DescribeProtectedResourceInput, encoder *httpbinding.Encoder) error {
@@ -2463,10 +2375,8 @@ func (m *awsRestjson1_serializeOpDescribeRecoveryPoint
 func (m *awsRestjson1_serializeOpDescribeRecoveryPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2504,7 +2414,6 @@ func (m *awsRestjson1_serializeOpDescribeRecoveryPoint
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRecoveryPointInput(v *DescribeRecoveryPointInput, encoder *httpbinding.Encoder) error {
@@ -2547,10 +2456,8 @@ func (m *awsRestjson1_serializeOpDescribeRegionSetting
 func (m *awsRestjson1_serializeOpDescribeRegionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2584,7 +2491,6 @@ func (m *awsRestjson1_serializeOpDescribeRegionSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRegionSettingsInput(v *DescribeRegionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2605,10 +2511,8 @@ func (m *awsRestjson1_serializeOpDescribeReportJob) Ha
 func (m *awsRestjson1_serializeOpDescribeReportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2646,7 +2550,6 @@ func (m *awsRestjson1_serializeOpDescribeReportJob) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReportJobInput(v *DescribeReportJobInput, encoder *httpbinding.Encoder) error {
@@ -2676,10 +2579,8 @@ func (m *awsRestjson1_serializeOpDescribeReportPlan) H
 func (m *awsRestjson1_serializeOpDescribeReportPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2717,7 +2618,6 @@ func (m *awsRestjson1_serializeOpDescribeReportPlan) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReportPlanInput(v *DescribeReportPlanInput, encoder *httpbinding.Encoder) error {
@@ -2747,10 +2647,8 @@ func (m *awsRestjson1_serializeOpDescribeRestoreJob) H
 func (m *awsRestjson1_serializeOpDescribeRestoreJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2788,7 +2686,6 @@ func (m *awsRestjson1_serializeOpDescribeRestoreJob) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRestoreJobInput(v *DescribeRestoreJobInput, encoder *httpbinding.Encoder) error {
@@ -2818,10 +2715,8 @@ func (m *awsRestjson1_serializeOpDisassociateBackupVau
 func (m *awsRestjson1_serializeOpDisassociateBackupVaultMpaApprovalTeam) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2870,7 +2765,6 @@ func (m *awsRestjson1_serializeOpDisassociateBackupVau
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateBackupVaultMpaApprovalTeamInput(v *DisassociateBackupVaultMpaApprovalTeamInput, encoder *httpbinding.Encoder) error {
@@ -2912,10 +2806,8 @@ func (m *awsRestjson1_serializeOpDisassociateRecoveryP
 func (m *awsRestjson1_serializeOpDisassociateRecoveryPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2953,7 +2845,6 @@ func (m *awsRestjson1_serializeOpDisassociateRecoveryP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateRecoveryPointInput(v *DisassociateRecoveryPointInput, encoder *httpbinding.Encoder) error {
@@ -2992,10 +2883,8 @@ func (m *awsRestjson1_serializeOpDisassociateRecoveryP
 func (m *awsRestjson1_serializeOpDisassociateRecoveryPointFromParent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3033,7 +2922,6 @@ func (m *awsRestjson1_serializeOpDisassociateRecoveryP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateRecoveryPointFromParentInput(v *DisassociateRecoveryPointFromParentInput, encoder *httpbinding.Encoder) error {
@@ -3072,10 +2960,8 @@ func (m *awsRestjson1_serializeOpExportBackupPlanTempl
 func (m *awsRestjson1_serializeOpExportBackupPlanTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3113,7 +2999,6 @@ func (m *awsRestjson1_serializeOpExportBackupPlanTempl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExportBackupPlanTemplateInput(v *ExportBackupPlanTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3143,10 +3028,8 @@ func (m *awsRestjson1_serializeOpGetBackupPlan) Handle
 func (m *awsRestjson1_serializeOpGetBackupPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3184,7 +3067,6 @@ func (m *awsRestjson1_serializeOpGetBackupPlan) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackupPlanInput(v *GetBackupPlanInput, encoder *httpbinding.Encoder) error {
@@ -3218,10 +3100,8 @@ func (m *awsRestjson1_serializeOpGetBackupPlanFromJSON
 func (m *awsRestjson1_serializeOpGetBackupPlanFromJSON) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3266,7 +3146,6 @@ func (m *awsRestjson1_serializeOpGetBackupPlanFromJSON
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackupPlanFromJSONInput(v *GetBackupPlanFromJSONInput, encoder *httpbinding.Encoder) error {
@@ -3299,10 +3178,8 @@ func (m *awsRestjson1_serializeOpGetBackupPlanFromTemp
 func (m *awsRestjson1_serializeOpGetBackupPlanFromTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3340,7 +3217,6 @@ func (m *awsRestjson1_serializeOpGetBackupPlanFromTemp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackupPlanFromTemplateInput(v *GetBackupPlanFromTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3370,10 +3246,8 @@ func (m *awsRestjson1_serializeOpGetBackupSelection) H
 func (m *awsRestjson1_serializeOpGetBackupSelection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3411,7 +3285,6 @@ func (m *awsRestjson1_serializeOpGetBackupSelection) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackupSelectionInput(v *GetBackupSelectionInput, encoder *httpbinding.Encoder) error {
@@ -3450,10 +3323,8 @@ func (m *awsRestjson1_serializeOpGetBackupVaultAccessP
 func (m *awsRestjson1_serializeOpGetBackupVaultAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3491,7 +3362,6 @@ func (m *awsRestjson1_serializeOpGetBackupVaultAccessP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackupVaultAccessPolicyInput(v *GetBackupVaultAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3521,10 +3391,8 @@ func (m *awsRestjson1_serializeOpGetBackupVaultNotific
 func (m *awsRestjson1_serializeOpGetBackupVaultNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3562,7 +3430,6 @@ func (m *awsRestjson1_serializeOpGetBackupVaultNotific
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackupVaultNotificationsInput(v *GetBackupVaultNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -3592,10 +3459,8 @@ func (m *awsRestjson1_serializeOpGetLegalHold) HandleS
 func (m *awsRestjson1_serializeOpGetLegalHold) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3633,7 +3498,6 @@ func (m *awsRestjson1_serializeOpGetLegalHold) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLegalHoldInput(v *GetLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -3663,10 +3527,8 @@ func (m *awsRestjson1_serializeOpGetRecoveryPointIndex
 func (m *awsRestjson1_serializeOpGetRecoveryPointIndexDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3704,7 +3566,6 @@ func (m *awsRestjson1_serializeOpGetRecoveryPointIndex
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRecoveryPointIndexDetailsInput(v *GetRecoveryPointIndexDetailsInput, encoder *httpbinding.Encoder) error {
@@ -3743,10 +3604,8 @@ func (m *awsRestjson1_serializeOpGetRecoveryPointResto
 func (m *awsRestjson1_serializeOpGetRecoveryPointRestoreMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3784,7 +3643,6 @@ func (m *awsRestjson1_serializeOpGetRecoveryPointResto
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRecoveryPointRestoreMetadataInput(v *GetRecoveryPointRestoreMetadataInput, encoder *httpbinding.Encoder) error {
@@ -3827,10 +3685,8 @@ func (m *awsRestjson1_serializeOpGetRestoreJobMetadata
 func (m *awsRestjson1_serializeOpGetRestoreJobMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3868,7 +3724,6 @@ func (m *awsRestjson1_serializeOpGetRestoreJobMetadata
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRestoreJobMetadataInput(v *GetRestoreJobMetadataInput, encoder *httpbinding.Encoder) error {
@@ -3898,10 +3753,8 @@ func (m *awsRestjson1_serializeOpGetRestoreTestingInfe
 func (m *awsRestjson1_serializeOpGetRestoreTestingInferredMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3939,7 +3792,6 @@ func (m *awsRestjson1_serializeOpGetRestoreTestingInfe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRestoreTestingInferredMetadataInput(v *GetRestoreTestingInferredMetadataInput, encoder *httpbinding.Encoder) error {
@@ -3972,10 +3824,8 @@ func (m *awsRestjson1_serializeOpGetRestoreTestingPlan
 func (m *awsRestjson1_serializeOpGetRestoreTestingPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4013,7 +3863,6 @@ func (m *awsRestjson1_serializeOpGetRestoreTestingPlan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRestoreTestingPlanInput(v *GetRestoreTestingPlanInput, encoder *httpbinding.Encoder) error {
@@ -4043,10 +3892,8 @@ func (m *awsRestjson1_serializeOpGetRestoreTestingSele
 func (m *awsRestjson1_serializeOpGetRestoreTestingSelection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4084,7 +3931,6 @@ func (m *awsRestjson1_serializeOpGetRestoreTestingSele
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRestoreTestingSelectionInput(v *GetRestoreTestingSelectionInput, encoder *httpbinding.Encoder) error {
@@ -4123,10 +3969,8 @@ func (m *awsRestjson1_serializeOpGetSupportedResourceT
 func (m *awsRestjson1_serializeOpGetSupportedResourceTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4160,7 +4004,6 @@ func (m *awsRestjson1_serializeOpGetSupportedResourceT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSupportedResourceTypesInput(v *GetSupportedResourceTypesInput, encoder *httpbinding.Encoder) error {
@@ -4181,10 +4024,8 @@ func (m *awsRestjson1_serializeOpListBackupJobs) Handl
 func (m *awsRestjson1_serializeOpListBackupJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4222,7 +4063,6 @@ func (m *awsRestjson1_serializeOpListBackupJobs) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackupJobsInput(v *ListBackupJobsInput, encoder *httpbinding.Encoder) error {
@@ -4295,10 +4135,8 @@ func (m *awsRestjson1_serializeOpListBackupJobSummarie
 func (m *awsRestjson1_serializeOpListBackupJobSummaries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4336,7 +4174,6 @@ func (m *awsRestjson1_serializeOpListBackupJobSummarie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackupJobSummariesInput(v *ListBackupJobSummariesInput, encoder *httpbinding.Encoder) error {
@@ -4385,10 +4222,8 @@ func (m *awsRestjson1_serializeOpListBackupPlans) Hand
 func (m *awsRestjson1_serializeOpListBackupPlans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4426,7 +4261,6 @@ func (m *awsRestjson1_serializeOpListBackupPlans) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackupPlansInput(v *ListBackupPlansInput, encoder *httpbinding.Encoder) error {
@@ -4459,10 +4293,8 @@ func (m *awsRestjson1_serializeOpListBackupPlanTemplat
 func (m *awsRestjson1_serializeOpListBackupPlanTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4500,7 +4332,6 @@ func (m *awsRestjson1_serializeOpListBackupPlanTemplat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackupPlanTemplatesInput(v *ListBackupPlanTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -4529,10 +4360,8 @@ func (m *awsRestjson1_serializeOpListBackupPlanVersion
 func (m *awsRestjson1_serializeOpListBackupPlanVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4570,7 +4399,6 @@ func (m *awsRestjson1_serializeOpListBackupPlanVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackupPlanVersionsInput(v *ListBackupPlanVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4608,10 +4436,8 @@ func (m *awsRestjson1_serializeOpListBackupSelections)
 func (m *awsRestjson1_serializeOpListBackupSelections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4649,7 +4475,6 @@ func (m *awsRestjson1_serializeOpListBackupSelections)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackupSelectionsInput(v *ListBackupSelectionsInput, encoder *httpbinding.Encoder) error {
@@ -4687,10 +4512,8 @@ func (m *awsRestjson1_serializeOpListBackupVaults) Han
 func (m *awsRestjson1_serializeOpListBackupVaults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4728,7 +4551,6 @@ func (m *awsRestjson1_serializeOpListBackupVaults) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackupVaultsInput(v *ListBackupVaultsInput, encoder *httpbinding.Encoder) error {
@@ -4765,10 +4587,8 @@ func (m *awsRestjson1_serializeOpListCopyJobs) HandleS
 func (m *awsRestjson1_serializeOpListCopyJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4806,7 +4626,6 @@ func (m *awsRestjson1_serializeOpListCopyJobs) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCopyJobsInput(v *ListCopyJobsInput, encoder *httpbinding.Encoder) error {
@@ -4879,10 +4698,8 @@ func (m *awsRestjson1_serializeOpListCopyJobSummaries)
 func (m *awsRestjson1_serializeOpListCopyJobSummaries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4920,7 +4737,6 @@ func (m *awsRestjson1_serializeOpListCopyJobSummaries)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCopyJobSummariesInput(v *ListCopyJobSummariesInput, encoder *httpbinding.Encoder) error {
@@ -4969,10 +4785,8 @@ func (m *awsRestjson1_serializeOpListFrameworks) Handl
 func (m *awsRestjson1_serializeOpListFrameworks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5010,7 +4824,6 @@ func (m *awsRestjson1_serializeOpListFrameworks) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFrameworksInput(v *ListFrameworksInput, encoder *httpbinding.Encoder) error {
@@ -5039,10 +4852,8 @@ func (m *awsRestjson1_serializeOpListIndexedRecoveryPo
 func (m *awsRestjson1_serializeOpListIndexedRecoveryPoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5080,7 +4891,6 @@ func (m *awsRestjson1_serializeOpListIndexedRecoveryPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIndexedRecoveryPointsInput(v *ListIndexedRecoveryPointsInput, encoder *httpbinding.Encoder) error {
@@ -5129,10 +4939,8 @@ func (m *awsRestjson1_serializeOpListLegalHolds) Handl
 func (m *awsRestjson1_serializeOpListLegalHolds) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5170,7 +4978,6 @@ func (m *awsRestjson1_serializeOpListLegalHolds) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLegalHoldsInput(v *ListLegalHoldsInput, encoder *httpbinding.Encoder) error {
@@ -5199,10 +5006,8 @@ func (m *awsRestjson1_serializeOpListProtectedResource
 func (m *awsRestjson1_serializeOpListProtectedResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5240,7 +5045,6 @@ func (m *awsRestjson1_serializeOpListProtectedResource
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProtectedResourcesInput(v *ListProtectedResourcesInput, encoder *httpbinding.Encoder) error {
@@ -5269,10 +5073,8 @@ func (m *awsRestjson1_serializeOpListProtectedResource
 func (m *awsRestjson1_serializeOpListProtectedResourcesByBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5310,7 +5112,6 @@ func (m *awsRestjson1_serializeOpListProtectedResource
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProtectedResourcesByBackupVaultInput(v *ListProtectedResourcesByBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -5352,10 +5153,8 @@ func (m *awsRestjson1_serializeOpListRecoveryPointsByB
 func (m *awsRestjson1_serializeOpListRecoveryPointsByBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5393,7 +5192,6 @@ func (m *awsRestjson1_serializeOpListRecoveryPointsByB
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecoveryPointsByBackupVaultInput(v *ListRecoveryPointsByBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -5459,10 +5257,8 @@ func (m *awsRestjson1_serializeOpListRecoveryPointsByL
 func (m *awsRestjson1_serializeOpListRecoveryPointsByLegalHold) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5500,7 +5296,6 @@ func (m *awsRestjson1_serializeOpListRecoveryPointsByL
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecoveryPointsByLegalHoldInput(v *ListRecoveryPointsByLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -5538,10 +5333,8 @@ func (m *awsRestjson1_serializeOpListRecoveryPointsByR
 func (m *awsRestjson1_serializeOpListRecoveryPointsByResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5579,7 +5372,6 @@ func (m *awsRestjson1_serializeOpListRecoveryPointsByR
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecoveryPointsByResourceInput(v *ListRecoveryPointsByResourceInput, encoder *httpbinding.Encoder) error {
@@ -5621,10 +5413,8 @@ func (m *awsRestjson1_serializeOpListReportJobs) Handl
 func (m *awsRestjson1_serializeOpListReportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5662,7 +5452,6 @@ func (m *awsRestjson1_serializeOpListReportJobs) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReportJobsInput(v *ListReportJobsInput, encoder *httpbinding.Encoder) error {
@@ -5707,10 +5496,8 @@ func (m *awsRestjson1_serializeOpListReportPlans) Hand
 func (m *awsRestjson1_serializeOpListReportPlans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5748,7 +5535,6 @@ func (m *awsRestjson1_serializeOpListReportPlans) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReportPlansInput(v *ListReportPlansInput, encoder *httpbinding.Encoder) error {
@@ -5777,10 +5563,8 @@ func (m *awsRestjson1_serializeOpListRestoreAccessBack
 func (m *awsRestjson1_serializeOpListRestoreAccessBackupVaults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5818,7 +5602,6 @@ func (m *awsRestjson1_serializeOpListRestoreAccessBack
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRestoreAccessBackupVaultsInput(v *ListRestoreAccessBackupVaultsInput, encoder *httpbinding.Encoder) error {
@@ -5856,10 +5639,8 @@ func (m *awsRestjson1_serializeOpListRestoreJobs) Hand
 func (m *awsRestjson1_serializeOpListRestoreJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5897,7 +5678,6 @@ func (m *awsRestjson1_serializeOpListRestoreJobs) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRestoreJobsInput(v *ListRestoreJobsInput, encoder *httpbinding.Encoder) error {
@@ -5958,10 +5738,8 @@ func (m *awsRestjson1_serializeOpListRestoreJobsByProt
 func (m *awsRestjson1_serializeOpListRestoreJobsByProtectedResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5999,7 +5777,6 @@ func (m *awsRestjson1_serializeOpListRestoreJobsByProt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRestoreJobsByProtectedResourceInput(v *ListRestoreJobsByProtectedResourceInput, encoder *httpbinding.Encoder) error {
@@ -6049,10 +5826,8 @@ func (m *awsRestjson1_serializeOpListRestoreJobSummari
 func (m *awsRestjson1_serializeOpListRestoreJobSummaries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6090,7 +5865,6 @@ func (m *awsRestjson1_serializeOpListRestoreJobSummari
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRestoreJobSummariesInput(v *ListRestoreJobSummariesInput, encoder *httpbinding.Encoder) error {
@@ -6135,10 +5909,8 @@ func (m *awsRestjson1_serializeOpListRestoreTestingPla
 func (m *awsRestjson1_serializeOpListRestoreTestingPlans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6176,7 +5948,6 @@ func (m *awsRestjson1_serializeOpListRestoreTestingPla
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRestoreTestingPlansInput(v *ListRestoreTestingPlansInput, encoder *httpbinding.Encoder) error {
@@ -6205,10 +5976,8 @@ func (m *awsRestjson1_serializeOpListRestoreTestingSel
 func (m *awsRestjson1_serializeOpListRestoreTestingSelections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6246,7 +6015,6 @@ func (m *awsRestjson1_serializeOpListRestoreTestingSel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRestoreTestingSelectionsInput(v *ListRestoreTestingSelectionsInput, encoder *httpbinding.Encoder) error {
@@ -6284,10 +6052,8 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
 func (m *awsRestjson1_serializeOpListTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6325,7 +6091,6 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsInput(v *ListTagsInput, encoder *httpbinding.Encoder) error {
@@ -6363,10 +6128,8 @@ func (m *awsRestjson1_serializeOpPutBackupVaultAccessP
 func (m *awsRestjson1_serializeOpPutBackupVaultAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6415,7 +6178,6 @@ func (m *awsRestjson1_serializeOpPutBackupVaultAccessP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutBackupVaultAccessPolicyInput(v *PutBackupVaultAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -6457,10 +6219,8 @@ func (m *awsRestjson1_serializeOpPutBackupVaultLockCon
 func (m *awsRestjson1_serializeOpPutBackupVaultLockConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6509,7 +6269,6 @@ func (m *awsRestjson1_serializeOpPutBackupVaultLockCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutBackupVaultLockConfigurationInput(v *PutBackupVaultLockConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6561,10 +6320,8 @@ func (m *awsRestjson1_serializeOpPutBackupVaultNotific
 func (m *awsRestjson1_serializeOpPutBackupVaultNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6613,7 +6370,6 @@ func (m *awsRestjson1_serializeOpPutBackupVaultNotific
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutBackupVaultNotificationsInput(v *PutBackupVaultNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -6662,10 +6418,8 @@ func (m *awsRestjson1_serializeOpPutRestoreValidationR
 func (m *awsRestjson1_serializeOpPutRestoreValidationResult) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6714,7 +6468,6 @@ func (m *awsRestjson1_serializeOpPutRestoreValidationR
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutRestoreValidationResultInput(v *PutRestoreValidationResultInput, encoder *httpbinding.Encoder) error {
@@ -6761,10 +6514,8 @@ func (m *awsRestjson1_serializeOpRevokeRestoreAccessBa
 func (m *awsRestjson1_serializeOpRevokeRestoreAccessBackupVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6802,7 +6553,6 @@ func (m *awsRestjson1_serializeOpRevokeRestoreAccessBa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRevokeRestoreAccessBackupVaultInput(v *RevokeRestoreAccessBackupVaultInput, encoder *httpbinding.Encoder) error {
@@ -6845,10 +6595,8 @@ func (m *awsRestjson1_serializeOpStartBackupJob) Handl
 func (m *awsRestjson1_serializeOpStartBackupJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6893,7 +6641,6 @@ func (m *awsRestjson1_serializeOpStartBackupJob) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartBackupJobInput(v *StartBackupJobInput, encoder *httpbinding.Encoder) error {
@@ -6977,10 +6724,8 @@ func (m *awsRestjson1_serializeOpStartCopyJob) HandleS
 func (m *awsRestjson1_serializeOpStartCopyJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7025,7 +6770,6 @@ func (m *awsRestjson1_serializeOpStartCopyJob) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartCopyJobInput(v *StartCopyJobInput, encoder *httpbinding.Encoder) error {
@@ -7085,10 +6829,8 @@ func (m *awsRestjson1_serializeOpStartReportJob) Handl
 func (m *awsRestjson1_serializeOpStartReportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7137,7 +6879,6 @@ func (m *awsRestjson1_serializeOpStartReportJob) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartReportJobInput(v *StartReportJobInput, encoder *httpbinding.Encoder) error {
@@ -7179,10 +6920,8 @@ func (m *awsRestjson1_serializeOpStartRestoreJob) Hand
 func (m *awsRestjson1_serializeOpStartRestoreJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7227,7 +6966,6 @@ func (m *awsRestjson1_serializeOpStartRestoreJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartRestoreJobInput(v *StartRestoreJobInput, encoder *httpbinding.Encoder) error {
@@ -7287,10 +7025,8 @@ func (m *awsRestjson1_serializeOpStopBackupJob) Handle
 func (m *awsRestjson1_serializeOpStopBackupJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7328,7 +7064,6 @@ func (m *awsRestjson1_serializeOpStopBackupJob) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopBackupJobInput(v *StopBackupJobInput, encoder *httpbinding.Encoder) error {
@@ -7358,10 +7093,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7410,7 +7143,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7454,10 +7186,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7506,7 +7236,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7550,10 +7279,8 @@ func (m *awsRestjson1_serializeOpUpdateBackupPlan) Han
 func (m *awsRestjson1_serializeOpUpdateBackupPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7602,7 +7329,6 @@ func (m *awsRestjson1_serializeOpUpdateBackupPlan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBackupPlanInput(v *UpdateBackupPlanInput, encoder *httpbinding.Encoder) error {
@@ -7646,10 +7372,8 @@ func (m *awsRestjson1_serializeOpUpdateFramework) Hand
 func (m *awsRestjson1_serializeOpUpdateFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7698,7 +7422,6 @@ func (m *awsRestjson1_serializeOpUpdateFramework) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFrameworkInput(v *UpdateFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -7752,10 +7475,8 @@ func (m *awsRestjson1_serializeOpUpdateGlobalSettings)
 func (m *awsRestjson1_serializeOpUpdateGlobalSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7800,7 +7521,6 @@ func (m *awsRestjson1_serializeOpUpdateGlobalSettings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGlobalSettingsInput(v *UpdateGlobalSettingsInput, encoder *httpbinding.Encoder) error {
@@ -7835,10 +7555,8 @@ func (m *awsRestjson1_serializeOpUpdateRecoveryPointIn
 func (m *awsRestjson1_serializeOpUpdateRecoveryPointIndexSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7887,7 +7605,6 @@ func (m *awsRestjson1_serializeOpUpdateRecoveryPointIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRecoveryPointIndexSettingsInput(v *UpdateRecoveryPointIndexSettingsInput, encoder *httpbinding.Encoder) error {
@@ -7943,10 +7660,8 @@ func (m *awsRestjson1_serializeOpUpdateRecoveryPointLi
 func (m *awsRestjson1_serializeOpUpdateRecoveryPointLifecycle) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7995,7 +7710,6 @@ func (m *awsRestjson1_serializeOpUpdateRecoveryPointLi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRecoveryPointLifecycleInput(v *UpdateRecoveryPointLifecycleInput, encoder *httpbinding.Encoder) error {
@@ -8048,10 +7762,8 @@ func (m *awsRestjson1_serializeOpUpdateRegionSettings)
 func (m *awsRestjson1_serializeOpUpdateRegionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8096,7 +7808,6 @@ func (m *awsRestjson1_serializeOpUpdateRegionSettings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRegionSettingsInput(v *UpdateRegionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -8138,10 +7849,8 @@ func (m *awsRestjson1_serializeOpUpdateReportPlan) Han
 func (m *awsRestjson1_serializeOpUpdateReportPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8190,7 +7899,6 @@ func (m *awsRestjson1_serializeOpUpdateReportPlan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateReportPlanInput(v *UpdateReportPlanInput, encoder *httpbinding.Encoder) error {
@@ -8251,10 +7959,8 @@ func (m *awsRestjson1_serializeOpUpdateRestoreTestingP
 func (m *awsRestjson1_serializeOpUpdateRestoreTestingPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8303,7 +8009,6 @@ func (m *awsRestjson1_serializeOpUpdateRestoreTestingP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRestoreTestingPlanInput(v *UpdateRestoreTestingPlanInput, encoder *httpbinding.Encoder) error {
@@ -8347,10 +8052,8 @@ func (m *awsRestjson1_serializeOpUpdateRestoreTestingS
 func (m *awsRestjson1_serializeOpUpdateRestoreTestingSelection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8399,7 +8102,6 @@ func (m *awsRestjson1_serializeOpUpdateRestoreTestingS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRestoreTestingSelectionInput(v *UpdateRestoreTestingSelectionInput, encoder *httpbinding.Encoder) error {
