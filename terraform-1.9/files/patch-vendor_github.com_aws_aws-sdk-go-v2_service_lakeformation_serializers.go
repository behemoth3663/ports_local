--- vendor/github.com/aws/aws-sdk-go-v2/service/lakeformation/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lakeformation/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAddLFTagsToResource) 
 func (m *awsRestjson1_serializeOpAddLFTagsToResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpAddLFTagsToResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddLFTagsToResourceInput(v *AddLFTagsToResourceInput, encoder *httpbinding.Encoder) error {
@@ -121,10 +117,8 @@ func (m *awsRestjson1_serializeOpAssumeDecoratedRoleWi
 func (m *awsRestjson1_serializeOpAssumeDecoratedRoleWithSAML) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -169,7 +163,6 @@ func (m *awsRestjson1_serializeOpAssumeDecoratedRoleWi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssumeDecoratedRoleWithSAMLInput(v *AssumeDecoratedRoleWithSAMLInput, encoder *httpbinding.Encoder) error {
@@ -217,10 +210,8 @@ func (m *awsRestjson1_serializeOpBatchGrantPermissions
 func (m *awsRestjson1_serializeOpBatchGrantPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -265,7 +256,6 @@ func (m *awsRestjson1_serializeOpBatchGrantPermissions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGrantPermissionsInput(v *BatchGrantPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -305,10 +295,8 @@ func (m *awsRestjson1_serializeOpBatchRevokePermission
 func (m *awsRestjson1_serializeOpBatchRevokePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -353,7 +341,6 @@ func (m *awsRestjson1_serializeOpBatchRevokePermission
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchRevokePermissionsInput(v *BatchRevokePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -393,10 +380,8 @@ func (m *awsRestjson1_serializeOpCancelTransaction) Ha
 func (m *awsRestjson1_serializeOpCancelTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -441,7 +426,6 @@ func (m *awsRestjson1_serializeOpCancelTransaction) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelTransactionInput(v *CancelTransactionInput, encoder *httpbinding.Encoder) error {
@@ -474,10 +458,8 @@ func (m *awsRestjson1_serializeOpCommitTransaction) Ha
 func (m *awsRestjson1_serializeOpCommitTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -522,7 +504,6 @@ func (m *awsRestjson1_serializeOpCommitTransaction) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCommitTransactionInput(v *CommitTransactionInput, encoder *httpbinding.Encoder) error {
@@ -555,10 +536,8 @@ func (m *awsRestjson1_serializeOpCreateDataCellsFilter
 func (m *awsRestjson1_serializeOpCreateDataCellsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -603,7 +582,6 @@ func (m *awsRestjson1_serializeOpCreateDataCellsFilter
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataCellsFilterInput(v *CreateDataCellsFilterInput, encoder *httpbinding.Encoder) error {
@@ -638,10 +616,8 @@ func (m *awsRestjson1_serializeOpCreateLakeFormationId
 func (m *awsRestjson1_serializeOpCreateLakeFormationIdentityCenterConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -686,7 +662,6 @@ func (m *awsRestjson1_serializeOpCreateLakeFormationId
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLakeFormationIdentityCenterConfigurationInput(v *CreateLakeFormationIdentityCenterConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -738,10 +713,8 @@ func (m *awsRestjson1_serializeOpCreateLakeFormationOp
 func (m *awsRestjson1_serializeOpCreateLakeFormationOptIn) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -786,7 +759,6 @@ func (m *awsRestjson1_serializeOpCreateLakeFormationOp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLakeFormationOptInInput(v *CreateLakeFormationOptInInput, encoder *httpbinding.Encoder) error {
@@ -835,10 +807,8 @@ func (m *awsRestjson1_serializeOpCreateLFTag) HandleSe
 func (m *awsRestjson1_serializeOpCreateLFTag) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -883,7 +853,6 @@ func (m *awsRestjson1_serializeOpCreateLFTag) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLFTagInput(v *CreateLFTagInput, encoder *httpbinding.Encoder) error {
@@ -928,10 +897,8 @@ func (m *awsRestjson1_serializeOpCreateLFTagExpression
 func (m *awsRestjson1_serializeOpCreateLFTagExpression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -976,7 +943,6 @@ func (m *awsRestjson1_serializeOpCreateLFTagExpression
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLFTagExpressionInput(v *CreateLFTagExpressionInput, encoder *httpbinding.Encoder) error {
@@ -1026,10 +992,8 @@ func (m *awsRestjson1_serializeOpDeleteDataCellsFilter
 func (m *awsRestjson1_serializeOpDeleteDataCellsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1074,7 +1038,6 @@ func (m *awsRestjson1_serializeOpDeleteDataCellsFilter
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataCellsFilterInput(v *DeleteDataCellsFilterInput, encoder *httpbinding.Encoder) error {
@@ -1122,10 +1085,8 @@ func (m *awsRestjson1_serializeOpDeleteLakeFormationId
 func (m *awsRestjson1_serializeOpDeleteLakeFormationIdentityCenterConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1170,7 +1131,6 @@ func (m *awsRestjson1_serializeOpDeleteLakeFormationId
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLakeFormationIdentityCenterConfigurationInput(v *DeleteLakeFormationIdentityCenterConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1203,10 +1163,8 @@ func (m *awsRestjson1_serializeOpDeleteLakeFormationOp
 func (m *awsRestjson1_serializeOpDeleteLakeFormationOptIn) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1251,7 +1209,6 @@ func (m *awsRestjson1_serializeOpDeleteLakeFormationOp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLakeFormationOptInInput(v *DeleteLakeFormationOptInInput, encoder *httpbinding.Encoder) error {
@@ -1300,10 +1257,8 @@ func (m *awsRestjson1_serializeOpDeleteLFTag) HandleSe
 func (m *awsRestjson1_serializeOpDeleteLFTag) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1348,7 +1303,6 @@ func (m *awsRestjson1_serializeOpDeleteLFTag) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLFTagInput(v *DeleteLFTagInput, encoder *httpbinding.Encoder) error {
@@ -1386,10 +1340,8 @@ func (m *awsRestjson1_serializeOpDeleteLFTagExpression
 func (m *awsRestjson1_serializeOpDeleteLFTagExpression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1434,7 +1386,6 @@ func (m *awsRestjson1_serializeOpDeleteLFTagExpression
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLFTagExpressionInput(v *DeleteLFTagExpressionInput, encoder *httpbinding.Encoder) error {
@@ -1472,10 +1423,8 @@ func (m *awsRestjson1_serializeOpDeleteObjectsOnCancel
 func (m *awsRestjson1_serializeOpDeleteObjectsOnCancel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1520,7 +1469,6 @@ func (m *awsRestjson1_serializeOpDeleteObjectsOnCancel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteObjectsOnCancelInput(v *DeleteObjectsOnCancelInput, encoder *httpbinding.Encoder) error {
@@ -1575,10 +1523,8 @@ func (m *awsRestjson1_serializeOpDeregisterResource) H
 func (m *awsRestjson1_serializeOpDeregisterResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1623,7 +1569,6 @@ func (m *awsRestjson1_serializeOpDeregisterResource) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeregisterResourceInput(v *DeregisterResourceInput, encoder *httpbinding.Encoder) error {
@@ -1656,10 +1601,8 @@ func (m *awsRestjson1_serializeOpDescribeLakeFormation
 func (m *awsRestjson1_serializeOpDescribeLakeFormationIdentityCenterConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1704,7 +1647,6 @@ func (m *awsRestjson1_serializeOpDescribeLakeFormation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeLakeFormationIdentityCenterConfigurationInput(v *DescribeLakeFormationIdentityCenterConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1737,10 +1679,8 @@ func (m *awsRestjson1_serializeOpDescribeResource) Han
 func (m *awsRestjson1_serializeOpDescribeResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1785,7 +1725,6 @@ func (m *awsRestjson1_serializeOpDescribeResource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeResourceInput(v *DescribeResourceInput, encoder *httpbinding.Encoder) error {
@@ -1818,10 +1757,8 @@ func (m *awsRestjson1_serializeOpDescribeTransaction) 
 func (m *awsRestjson1_serializeOpDescribeTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1866,7 +1803,6 @@ func (m *awsRestjson1_serializeOpDescribeTransaction) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTransactionInput(v *DescribeTransactionInput, encoder *httpbinding.Encoder) error {
@@ -1899,10 +1835,8 @@ func (m *awsRestjson1_serializeOpExtendTransaction) Ha
 func (m *awsRestjson1_serializeOpExtendTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1947,7 +1881,6 @@ func (m *awsRestjson1_serializeOpExtendTransaction) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExtendTransactionInput(v *ExtendTransactionInput, encoder *httpbinding.Encoder) error {
@@ -1980,10 +1913,8 @@ func (m *awsRestjson1_serializeOpGetDataCellsFilter) H
 func (m *awsRestjson1_serializeOpGetDataCellsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2028,7 +1959,6 @@ func (m *awsRestjson1_serializeOpGetDataCellsFilter) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataCellsFilterInput(v *GetDataCellsFilterInput, encoder *httpbinding.Encoder) error {
@@ -2076,10 +2006,8 @@ func (m *awsRestjson1_serializeOpGetDataLakePrincipal)
 func (m *awsRestjson1_serializeOpGetDataLakePrincipal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2113,7 +2041,6 @@ func (m *awsRestjson1_serializeOpGetDataLakePrincipal)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataLakePrincipalInput(v *GetDataLakePrincipalInput, encoder *httpbinding.Encoder) error {
@@ -2134,10 +2061,8 @@ func (m *awsRestjson1_serializeOpGetDataLakeSettings) 
 func (m *awsRestjson1_serializeOpGetDataLakeSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2182,7 +2107,6 @@ func (m *awsRestjson1_serializeOpGetDataLakeSettings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataLakeSettingsInput(v *GetDataLakeSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2215,10 +2139,8 @@ func (m *awsRestjson1_serializeOpGetEffectivePermissio
 func (m *awsRestjson1_serializeOpGetEffectivePermissionsForPath) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2263,7 +2185,6 @@ func (m *awsRestjson1_serializeOpGetEffectivePermissio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEffectivePermissionsForPathInput(v *GetEffectivePermissionsForPathInput, encoder *httpbinding.Encoder) error {
@@ -2311,10 +2232,8 @@ func (m *awsRestjson1_serializeOpGetLFTag) HandleSeria
 func (m *awsRestjson1_serializeOpGetLFTag) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2359,7 +2278,6 @@ func (m *awsRestjson1_serializeOpGetLFTag) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLFTagInput(v *GetLFTagInput, encoder *httpbinding.Encoder) error {
@@ -2397,10 +2315,8 @@ func (m *awsRestjson1_serializeOpGetLFTagExpression) H
 func (m *awsRestjson1_serializeOpGetLFTagExpression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2445,7 +2361,6 @@ func (m *awsRestjson1_serializeOpGetLFTagExpression) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLFTagExpressionInput(v *GetLFTagExpressionInput, encoder *httpbinding.Encoder) error {
@@ -2483,10 +2398,8 @@ func (m *awsRestjson1_serializeOpGetQueryState) Handle
 func (m *awsRestjson1_serializeOpGetQueryState) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2531,7 +2444,6 @@ func (m *awsRestjson1_serializeOpGetQueryState) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryStateInput(v *GetQueryStateInput, encoder *httpbinding.Encoder) error {
@@ -2564,10 +2476,8 @@ func (m *awsRestjson1_serializeOpGetQueryStatistics) H
 func (m *awsRestjson1_serializeOpGetQueryStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2612,7 +2522,6 @@ func (m *awsRestjson1_serializeOpGetQueryStatistics) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryStatisticsInput(v *GetQueryStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -2645,10 +2554,8 @@ func (m *awsRestjson1_serializeOpGetResourceLFTags) Ha
 func (m *awsRestjson1_serializeOpGetResourceLFTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2693,7 +2600,6 @@ func (m *awsRestjson1_serializeOpGetResourceLFTags) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceLFTagsInput(v *GetResourceLFTagsInput, encoder *httpbinding.Encoder) error {
@@ -2738,10 +2644,8 @@ func (m *awsRestjson1_serializeOpGetTableObjects) Hand
 func (m *awsRestjson1_serializeOpGetTableObjects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2786,7 +2690,6 @@ func (m *awsRestjson1_serializeOpGetTableObjects) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableObjectsInput(v *GetTableObjectsInput, encoder *httpbinding.Encoder) error {
@@ -2854,10 +2757,8 @@ func (m *awsRestjson1_serializeOpGetTemporaryGlueParti
 func (m *awsRestjson1_serializeOpGetTemporaryGluePartitionCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2902,7 +2803,6 @@ func (m *awsRestjson1_serializeOpGetTemporaryGlueParti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTemporaryGluePartitionCredentialsInput(v *GetTemporaryGluePartitionCredentialsInput, encoder *httpbinding.Encoder) error {
@@ -2968,10 +2868,8 @@ func (m *awsRestjson1_serializeOpGetTemporaryGlueTable
 func (m *awsRestjson1_serializeOpGetTemporaryGlueTableCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3016,7 +2914,6 @@ func (m *awsRestjson1_serializeOpGetTemporaryGlueTable
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTemporaryGlueTableCredentialsInput(v *GetTemporaryGlueTableCredentialsInput, encoder *httpbinding.Encoder) error {
@@ -3087,10 +2984,8 @@ func (m *awsRestjson1_serializeOpGetWorkUnitResults) H
 func (m *awsRestjson1_serializeOpGetWorkUnitResults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3135,7 +3030,6 @@ func (m *awsRestjson1_serializeOpGetWorkUnitResults) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkUnitResultsInput(v *GetWorkUnitResultsInput, encoder *httpbinding.Encoder) error {
@@ -3178,10 +3072,8 @@ func (m *awsRestjson1_serializeOpGetWorkUnits) HandleS
 func (m *awsRestjson1_serializeOpGetWorkUnits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3226,7 +3118,6 @@ func (m *awsRestjson1_serializeOpGetWorkUnits) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkUnitsInput(v *GetWorkUnitsInput, encoder *httpbinding.Encoder) error {
@@ -3269,10 +3160,8 @@ func (m *awsRestjson1_serializeOpGrantPermissions) Han
 func (m *awsRestjson1_serializeOpGrantPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3317,7 +3206,6 @@ func (m *awsRestjson1_serializeOpGrantPermissions) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGrantPermissionsInput(v *GrantPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -3385,10 +3273,8 @@ func (m *awsRestjson1_serializeOpListDataCellsFilter) 
 func (m *awsRestjson1_serializeOpListDataCellsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3433,7 +3319,6 @@ func (m *awsRestjson1_serializeOpListDataCellsFilter) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataCellsFilterInput(v *ListDataCellsFilterInput, encoder *httpbinding.Encoder) error {
@@ -3478,10 +3363,8 @@ func (m *awsRestjson1_serializeOpListLakeFormationOptI
 func (m *awsRestjson1_serializeOpListLakeFormationOptIns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3526,7 +3409,6 @@ func (m *awsRestjson1_serializeOpListLakeFormationOptI
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLakeFormationOptInsInput(v *ListLakeFormationOptInsInput, encoder *httpbinding.Encoder) error {
@@ -3578,10 +3460,8 @@ func (m *awsRestjson1_serializeOpListLFTagExpressions)
 func (m *awsRestjson1_serializeOpListLFTagExpressions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3626,7 +3506,6 @@ func (m *awsRestjson1_serializeOpListLFTagExpressions)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLFTagExpressionsInput(v *ListLFTagExpressionsInput, encoder *httpbinding.Encoder) error {
@@ -3669,10 +3548,8 @@ func (m *awsRestjson1_serializeOpListLFTags) HandleSer
 func (m *awsRestjson1_serializeOpListLFTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3717,7 +3594,6 @@ func (m *awsRestjson1_serializeOpListLFTags) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLFTagsInput(v *ListLFTagsInput, encoder *httpbinding.Encoder) error {
@@ -3765,10 +3641,8 @@ func (m *awsRestjson1_serializeOpListPermissions) Hand
 func (m *awsRestjson1_serializeOpListPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3813,7 +3687,6 @@ func (m *awsRestjson1_serializeOpListPermissions) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPermissionsInput(v *ListPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -3880,10 +3753,8 @@ func (m *awsRestjson1_serializeOpListResources) Handle
 func (m *awsRestjson1_serializeOpListResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3928,7 +3799,6 @@ func (m *awsRestjson1_serializeOpListResources) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourcesInput(v *ListResourcesInput, encoder *httpbinding.Encoder) error {
@@ -3973,10 +3843,8 @@ func (m *awsRestjson1_serializeOpListTableStorageOptim
 func (m *awsRestjson1_serializeOpListTableStorageOptimizers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4021,7 +3889,6 @@ func (m *awsRestjson1_serializeOpListTableStorageOptim
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTableStorageOptimizersInput(v *ListTableStorageOptimizersInput, encoder *httpbinding.Encoder) error {
@@ -4079,10 +3946,8 @@ func (m *awsRestjson1_serializeOpListTransactions) Han
 func (m *awsRestjson1_serializeOpListTransactions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4127,7 +3992,6 @@ func (m *awsRestjson1_serializeOpListTransactions) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTransactionsInput(v *ListTransactionsInput, encoder *httpbinding.Encoder) error {
@@ -4175,10 +4039,8 @@ func (m *awsRestjson1_serializeOpPutDataLakeSettings) 
 func (m *awsRestjson1_serializeOpPutDataLakeSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4223,7 +4085,6 @@ func (m *awsRestjson1_serializeOpPutDataLakeSettings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDataLakeSettingsInput(v *PutDataLakeSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4263,10 +4124,8 @@ func (m *awsRestjson1_serializeOpRegisterResource) Han
 func (m *awsRestjson1_serializeOpRegisterResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4311,7 +4170,6 @@ func (m *awsRestjson1_serializeOpRegisterResource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterResourceInput(v *RegisterResourceInput, encoder *httpbinding.Encoder) error {
@@ -4369,10 +4227,8 @@ func (m *awsRestjson1_serializeOpRemoveLFTagsFromResou
 func (m *awsRestjson1_serializeOpRemoveLFTagsFromResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4417,7 +4273,6 @@ func (m *awsRestjson1_serializeOpRemoveLFTagsFromResou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveLFTagsFromResourceInput(v *RemoveLFTagsFromResourceInput, encoder *httpbinding.Encoder) error {
@@ -4464,10 +4319,8 @@ func (m *awsRestjson1_serializeOpRevokePermissions) Ha
 func (m *awsRestjson1_serializeOpRevokePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4512,7 +4365,6 @@ func (m *awsRestjson1_serializeOpRevokePermissions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRevokePermissionsInput(v *RevokePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -4580,10 +4432,8 @@ func (m *awsRestjson1_serializeOpSearchDatabasesByLFTa
 func (m *awsRestjson1_serializeOpSearchDatabasesByLFTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4628,7 +4478,6 @@ func (m *awsRestjson1_serializeOpSearchDatabasesByLFTa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchDatabasesByLFTagsInput(v *SearchDatabasesByLFTagsInput, encoder *httpbinding.Encoder) error {
@@ -4678,10 +4527,8 @@ func (m *awsRestjson1_serializeOpSearchTablesByLFTags)
 func (m *awsRestjson1_serializeOpSearchTablesByLFTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4726,7 +4573,6 @@ func (m *awsRestjson1_serializeOpSearchTablesByLFTags)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchTablesByLFTagsInput(v *SearchTablesByLFTagsInput, encoder *httpbinding.Encoder) error {
@@ -4776,10 +4622,8 @@ func (m *awsRestjson1_serializeOpStartQueryPlanning) H
 func (m *awsRestjson1_serializeOpStartQueryPlanning) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4824,7 +4668,6 @@ func (m *awsRestjson1_serializeOpStartQueryPlanning) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartQueryPlanningInput(v *StartQueryPlanningInput, encoder *httpbinding.Encoder) error {
@@ -4864,10 +4707,8 @@ func (m *awsRestjson1_serializeOpStartTransaction) Han
 func (m *awsRestjson1_serializeOpStartTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4912,7 +4753,6 @@ func (m *awsRestjson1_serializeOpStartTransaction) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTransactionInput(v *StartTransactionInput, encoder *httpbinding.Encoder) error {
@@ -4945,10 +4785,8 @@ func (m *awsRestjson1_serializeOpUpdateDataCellsFilter
 func (m *awsRestjson1_serializeOpUpdateDataCellsFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4993,7 +4831,6 @@ func (m *awsRestjson1_serializeOpUpdateDataCellsFilter
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataCellsFilterInput(v *UpdateDataCellsFilterInput, encoder *httpbinding.Encoder) error {
@@ -5028,10 +4865,8 @@ func (m *awsRestjson1_serializeOpUpdateLakeFormationId
 func (m *awsRestjson1_serializeOpUpdateLakeFormationIdentityCenterConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5076,7 +4911,6 @@ func (m *awsRestjson1_serializeOpUpdateLakeFormationId
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLakeFormationIdentityCenterConfigurationInput(v *UpdateLakeFormationIdentityCenterConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5128,10 +4962,8 @@ func (m *awsRestjson1_serializeOpUpdateLFTag) HandleSe
 func (m *awsRestjson1_serializeOpUpdateLFTag) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5176,7 +5008,6 @@ func (m *awsRestjson1_serializeOpUpdateLFTag) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLFTagInput(v *UpdateLFTagInput, encoder *httpbinding.Encoder) error {
@@ -5228,10 +5059,8 @@ func (m *awsRestjson1_serializeOpUpdateLFTagExpression
 func (m *awsRestjson1_serializeOpUpdateLFTagExpression) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5276,7 +5105,6 @@ func (m *awsRestjson1_serializeOpUpdateLFTagExpression
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLFTagExpressionInput(v *UpdateLFTagExpressionInput, encoder *httpbinding.Encoder) error {
@@ -5326,10 +5154,8 @@ func (m *awsRestjson1_serializeOpUpdateResource) Handl
 func (m *awsRestjson1_serializeOpUpdateResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5374,7 +5200,6 @@ func (m *awsRestjson1_serializeOpUpdateResource) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceInput(v *UpdateResourceInput, encoder *httpbinding.Encoder) error {
@@ -5422,10 +5247,8 @@ func (m *awsRestjson1_serializeOpUpdateTableObjects) H
 func (m *awsRestjson1_serializeOpUpdateTableObjects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5470,7 +5293,6 @@ func (m *awsRestjson1_serializeOpUpdateTableObjects) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTableObjectsInput(v *UpdateTableObjectsInput, encoder *httpbinding.Encoder) error {
@@ -5525,10 +5347,8 @@ func (m *awsRestjson1_serializeOpUpdateTableStorageOpt
 func (m *awsRestjson1_serializeOpUpdateTableStorageOptimizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5573,7 +5393,6 @@ func (m *awsRestjson1_serializeOpUpdateTableStorageOpt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTableStorageOptimizerInput(v *UpdateTableStorageOptimizerInput, encoder *httpbinding.Encoder) error {
