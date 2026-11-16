--- vendor/github.com/aws/aws-sdk-go-v2/service/iotsitewise/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotsitewise/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAssociateAssets) Hand
 func (m *awsRestjson1_serializeOpAssociateAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpAssociateAssets) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateAssetsInput(v *AssociateAssetsInput, encoder *httpbinding.Encoder) error {
@@ -131,10 +127,8 @@ func (m *awsRestjson1_serializeOpAssociateTimeSeriesTo
 func (m *awsRestjson1_serializeOpAssociateTimeSeriesToAssetProperty) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -183,7 +177,6 @@ func (m *awsRestjson1_serializeOpAssociateTimeSeriesTo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateTimeSeriesToAssetPropertyInput(v *AssociateTimeSeriesToAssetPropertyInput, encoder *httpbinding.Encoder) error {
@@ -228,10 +221,8 @@ func (m *awsRestjson1_serializeOpBatchAssociateProject
 func (m *awsRestjson1_serializeOpBatchAssociateProjectAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -280,7 +271,6 @@ func (m *awsRestjson1_serializeOpBatchAssociateProject
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchAssociateProjectAssetsInput(v *BatchAssociateProjectAssetsInput, encoder *httpbinding.Encoder) error {
@@ -329,10 +319,8 @@ func (m *awsRestjson1_serializeOpBatchDisassociateProj
 func (m *awsRestjson1_serializeOpBatchDisassociateProjectAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -381,7 +369,6 @@ func (m *awsRestjson1_serializeOpBatchDisassociateProj
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDisassociateProjectAssetsInput(v *BatchDisassociateProjectAssetsInput, encoder *httpbinding.Encoder) error {
@@ -430,10 +417,8 @@ func (m *awsRestjson1_serializeOpBatchGetAssetProperty
 func (m *awsRestjson1_serializeOpBatchGetAssetPropertyAggregates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -478,7 +463,6 @@ func (m *awsRestjson1_serializeOpBatchGetAssetProperty
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetAssetPropertyAggregatesInput(v *BatchGetAssetPropertyAggregatesInput, encoder *httpbinding.Encoder) error {
@@ -523,10 +507,8 @@ func (m *awsRestjson1_serializeOpBatchGetAssetProperty
 func (m *awsRestjson1_serializeOpBatchGetAssetPropertyValue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -571,7 +553,6 @@ func (m *awsRestjson1_serializeOpBatchGetAssetProperty
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetAssetPropertyValueInput(v *BatchGetAssetPropertyValueInput, encoder *httpbinding.Encoder) error {
@@ -611,10 +592,8 @@ func (m *awsRestjson1_serializeOpBatchGetAssetProperty
 func (m *awsRestjson1_serializeOpBatchGetAssetPropertyValueHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -659,7 +638,6 @@ func (m *awsRestjson1_serializeOpBatchGetAssetProperty
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetAssetPropertyValueHistoryInput(v *BatchGetAssetPropertyValueHistoryInput, encoder *httpbinding.Encoder) error {
@@ -704,10 +682,8 @@ func (m *awsRestjson1_serializeOpBatchPutAssetProperty
 func (m *awsRestjson1_serializeOpBatchPutAssetPropertyValue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -752,7 +728,6 @@ func (m *awsRestjson1_serializeOpBatchPutAssetProperty
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchPutAssetPropertyValueInput(v *BatchPutAssetPropertyValueInput, encoder *httpbinding.Encoder) error {
@@ -792,10 +767,8 @@ func (m *awsRestjson1_serializeOpCreateAccessPolicy) H
 func (m *awsRestjson1_serializeOpCreateAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -840,7 +813,6 @@ func (m *awsRestjson1_serializeOpCreateAccessPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAccessPolicyInput(v *CreateAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -899,10 +871,8 @@ func (m *awsRestjson1_serializeOpCreateAsset) HandleSe
 func (m *awsRestjson1_serializeOpCreateAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -947,7 +917,6 @@ func (m *awsRestjson1_serializeOpCreateAsset) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssetInput(v *CreateAssetInput, encoder *httpbinding.Encoder) error {
@@ -1012,10 +981,8 @@ func (m *awsRestjson1_serializeOpCreateAssetModel) Han
 func (m *awsRestjson1_serializeOpCreateAssetModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1060,7 +1027,6 @@ func (m *awsRestjson1_serializeOpCreateAssetModel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssetModelInput(v *CreateAssetModelInput, encoder *httpbinding.Encoder) error {
@@ -1146,10 +1112,8 @@ func (m *awsRestjson1_serializeOpCreateAssetModelCompo
 func (m *awsRestjson1_serializeOpCreateAssetModelCompositeModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1198,7 +1162,6 @@ func (m *awsRestjson1_serializeOpCreateAssetModelCompo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssetModelCompositeModelInput(v *CreateAssetModelCompositeModelInput, encoder *httpbinding.Encoder) error {
@@ -1297,10 +1260,8 @@ func (m *awsRestjson1_serializeOpCreateBulkImportJob) 
 func (m *awsRestjson1_serializeOpCreateBulkImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1345,7 +1306,6 @@ func (m *awsRestjson1_serializeOpCreateBulkImportJob) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBulkImportJobInput(v *CreateBulkImportJobInput, encoder *httpbinding.Encoder) error {
@@ -1414,10 +1374,8 @@ func (m *awsRestjson1_serializeOpCreateDashboard) Hand
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
@@ -1462,7 +1420,6 @@ func (m *awsRestjson1_serializeOpCreateDashboard) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDashboardInput(v *CreateDashboardInput, encoder *httpbinding.Encoder) error {
@@ -1522,10 +1479,8 @@ func (m *awsRestjson1_serializeOpCreateDataset) Handle
 func (m *awsRestjson1_serializeOpCreateDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1570,7 +1525,6 @@ func (m *awsRestjson1_serializeOpCreateDataset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDatasetInput(v *CreateDatasetInput, encoder *httpbinding.Encoder) error {
@@ -1632,10 +1586,8 @@ func (m *awsRestjson1_serializeOpCreateGateway) Handle
 func (m *awsRestjson1_serializeOpCreateGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1680,7 +1632,6 @@ func (m *awsRestjson1_serializeOpCreateGateway) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGatewayInput(v *CreateGatewayInput, encoder *httpbinding.Encoder) error {
@@ -1732,10 +1683,8 @@ func (m *awsRestjson1_serializeOpCreatePortal) HandleS
 func (m *awsRestjson1_serializeOpCreatePortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1780,7 +1729,6 @@ func (m *awsRestjson1_serializeOpCreatePortal) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePortalInput(v *CreatePortalInput, encoder *httpbinding.Encoder) error {
@@ -1876,10 +1824,8 @@ func (m *awsRestjson1_serializeOpCreateProject) Handle
 func (m *awsRestjson1_serializeOpCreateProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1924,7 +1870,6 @@ func (m *awsRestjson1_serializeOpCreateProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProjectInput(v *CreateProjectInput, encoder *httpbinding.Encoder) error {
@@ -1979,10 +1924,8 @@ func (m *awsRestjson1_serializeOpDeleteAccessPolicy) H
 func (m *awsRestjson1_serializeOpDeleteAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2020,7 +1963,6 @@ func (m *awsRestjson1_serializeOpDeleteAccessPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccessPolicyInput(v *DeleteAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2054,10 +1996,8 @@ func (m *awsRestjson1_serializeOpDeleteAsset) HandleSe
 func (m *awsRestjson1_serializeOpDeleteAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2095,7 +2035,6 @@ func (m *awsRestjson1_serializeOpDeleteAsset) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssetInput(v *DeleteAssetInput, encoder *httpbinding.Encoder) error {
@@ -2129,10 +2068,8 @@ func (m *awsRestjson1_serializeOpDeleteAssetModel) Han
 func (m *awsRestjson1_serializeOpDeleteAssetModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2170,7 +2107,6 @@ func (m *awsRestjson1_serializeOpDeleteAssetModel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssetModelInput(v *DeleteAssetModelInput, encoder *httpbinding.Encoder) error {
@@ -2219,10 +2155,8 @@ func (m *awsRestjson1_serializeOpDeleteAssetModelCompo
 func (m *awsRestjson1_serializeOpDeleteAssetModelCompositeModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2260,7 +2194,6 @@ func (m *awsRestjson1_serializeOpDeleteAssetModelCompo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssetModelCompositeModelInput(v *DeleteAssetModelCompositeModelInput, encoder *httpbinding.Encoder) error {
@@ -2318,10 +2251,8 @@ func (m *awsRestjson1_serializeOpDeleteDashboard) Hand
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
@@ -2359,7 +2290,6 @@ func (m *awsRestjson1_serializeOpDeleteDashboard) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDashboardInput(v *DeleteDashboardInput, encoder *httpbinding.Encoder) error {
@@ -2393,10 +2323,8 @@ func (m *awsRestjson1_serializeOpDeleteDataset) Handle
 func (m *awsRestjson1_serializeOpDeleteDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2434,7 +2362,6 @@ func (m *awsRestjson1_serializeOpDeleteDataset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDatasetInput(v *DeleteDatasetInput, encoder *httpbinding.Encoder) error {
@@ -2468,10 +2395,8 @@ func (m *awsRestjson1_serializeOpDeleteGateway) Handle
 func (m *awsRestjson1_serializeOpDeleteGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2509,7 +2434,6 @@ func (m *awsRestjson1_serializeOpDeleteGateway) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGatewayInput(v *DeleteGatewayInput, encoder *httpbinding.Encoder) error {
@@ -2539,10 +2463,8 @@ func (m *awsRestjson1_serializeOpDeletePortal) HandleS
 func (m *awsRestjson1_serializeOpDeletePortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2580,7 +2502,6 @@ func (m *awsRestjson1_serializeOpDeletePortal) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePortalInput(v *DeletePortalInput, encoder *httpbinding.Encoder) error {
@@ -2614,10 +2535,8 @@ func (m *awsRestjson1_serializeOpDeleteProject) Handle
 func (m *awsRestjson1_serializeOpDeleteProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2655,7 +2574,6 @@ func (m *awsRestjson1_serializeOpDeleteProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProjectInput(v *DeleteProjectInput, encoder *httpbinding.Encoder) error {
@@ -2689,10 +2607,8 @@ func (m *awsRestjson1_serializeOpDeleteTimeSeries) Han
 func (m *awsRestjson1_serializeOpDeleteTimeSeries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2741,7 +2657,6 @@ func (m *awsRestjson1_serializeOpDeleteTimeSeries) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTimeSeriesInput(v *DeleteTimeSeriesInput, encoder *httpbinding.Encoder) error {
@@ -2786,10 +2701,8 @@ func (m *awsRestjson1_serializeOpDescribeAccessPolicy)
 func (m *awsRestjson1_serializeOpDescribeAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2827,7 +2740,6 @@ func (m *awsRestjson1_serializeOpDescribeAccessPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccessPolicyInput(v *DescribeAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2857,10 +2769,8 @@ func (m *awsRestjson1_serializeOpDescribeAction) Handl
 func (m *awsRestjson1_serializeOpDescribeAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2898,7 +2808,6 @@ func (m *awsRestjson1_serializeOpDescribeAction) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeActionInput(v *DescribeActionInput, encoder *httpbinding.Encoder) error {
@@ -2928,10 +2837,8 @@ func (m *awsRestjson1_serializeOpDescribeAsset) Handle
 func (m *awsRestjson1_serializeOpDescribeAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2969,7 +2876,6 @@ func (m *awsRestjson1_serializeOpDescribeAsset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAssetInput(v *DescribeAssetInput, encoder *httpbinding.Encoder) error {
@@ -3003,10 +2909,8 @@ func (m *awsRestjson1_serializeOpDescribeAssetComposit
 func (m *awsRestjson1_serializeOpDescribeAssetCompositeModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3044,7 +2948,6 @@ func (m *awsRestjson1_serializeOpDescribeAssetComposit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAssetCompositeModelInput(v *DescribeAssetCompositeModelInput, encoder *httpbinding.Encoder) error {
@@ -3083,10 +2986,8 @@ func (m *awsRestjson1_serializeOpDescribeAssetModel) H
 func (m *awsRestjson1_serializeOpDescribeAssetModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3124,7 +3025,6 @@ func (m *awsRestjson1_serializeOpDescribeAssetModel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAssetModelInput(v *DescribeAssetModelInput, encoder *httpbinding.Encoder) error {
@@ -3162,10 +3062,8 @@ func (m *awsRestjson1_serializeOpDescribeAssetModelCom
 func (m *awsRestjson1_serializeOpDescribeAssetModelCompositeModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3203,7 +3101,6 @@ func (m *awsRestjson1_serializeOpDescribeAssetModelCom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAssetModelCompositeModelInput(v *DescribeAssetModelCompositeModelInput, encoder *httpbinding.Encoder) error {
@@ -3246,10 +3143,8 @@ func (m *awsRestjson1_serializeOpDescribeAssetProperty
 func (m *awsRestjson1_serializeOpDescribeAssetProperty) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3287,7 +3182,6 @@ func (m *awsRestjson1_serializeOpDescribeAssetProperty
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAssetPropertyInput(v *DescribeAssetPropertyInput, encoder *httpbinding.Encoder) error {
@@ -3326,10 +3220,8 @@ func (m *awsRestjson1_serializeOpDescribeBulkImportJob
 func (m *awsRestjson1_serializeOpDescribeBulkImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3367,7 +3259,6 @@ func (m *awsRestjson1_serializeOpDescribeBulkImportJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeBulkImportJobInput(v *DescribeBulkImportJobInput, encoder *httpbinding.Encoder) error {
@@ -3397,10 +3288,8 @@ func (m *awsRestjson1_serializeOpDescribeDashboard) Ha
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
@@ -3438,7 +3327,6 @@ func (m *awsRestjson1_serializeOpDescribeDashboard) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDashboardInput(v *DescribeDashboardInput, encoder *httpbinding.Encoder) error {
@@ -3468,10 +3356,8 @@ func (m *awsRestjson1_serializeOpDescribeDataset) Hand
 func (m *awsRestjson1_serializeOpDescribeDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3509,7 +3395,6 @@ func (m *awsRestjson1_serializeOpDescribeDataset) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDatasetInput(v *DescribeDatasetInput, encoder *httpbinding.Encoder) error {
@@ -3539,10 +3424,8 @@ func (m *awsRestjson1_serializeOpDescribeDefaultEncryp
 func (m *awsRestjson1_serializeOpDescribeDefaultEncryptionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3576,7 +3459,6 @@ func (m *awsRestjson1_serializeOpDescribeDefaultEncryp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDefaultEncryptionConfigurationInput(v *DescribeDefaultEncryptionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3597,10 +3479,8 @@ func (m *awsRestjson1_serializeOpDescribeGateway) Hand
 func (m *awsRestjson1_serializeOpDescribeGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3638,7 +3518,6 @@ func (m *awsRestjson1_serializeOpDescribeGateway) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeGatewayInput(v *DescribeGatewayInput, encoder *httpbinding.Encoder) error {
@@ -3668,10 +3547,8 @@ func (m *awsRestjson1_serializeOpDescribeGatewayCapabi
 func (m *awsRestjson1_serializeOpDescribeGatewayCapabilityConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3709,7 +3586,6 @@ func (m *awsRestjson1_serializeOpDescribeGatewayCapabi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeGatewayCapabilityConfigurationInput(v *DescribeGatewayCapabilityConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3748,10 +3624,8 @@ func (m *awsRestjson1_serializeOpDescribeLoggingOption
 func (m *awsRestjson1_serializeOpDescribeLoggingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3785,7 +3659,6 @@ func (m *awsRestjson1_serializeOpDescribeLoggingOption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeLoggingOptionsInput(v *DescribeLoggingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -3806,10 +3679,8 @@ func (m *awsRestjson1_serializeOpDescribePortal) Handl
 func (m *awsRestjson1_serializeOpDescribePortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3847,7 +3718,6 @@ func (m *awsRestjson1_serializeOpDescribePortal) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePortalInput(v *DescribePortalInput, encoder *httpbinding.Encoder) error {
@@ -3877,10 +3747,8 @@ func (m *awsRestjson1_serializeOpDescribeProject) Hand
 func (m *awsRestjson1_serializeOpDescribeProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3918,7 +3786,6 @@ func (m *awsRestjson1_serializeOpDescribeProject) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeProjectInput(v *DescribeProjectInput, encoder *httpbinding.Encoder) error {
@@ -3948,10 +3815,8 @@ func (m *awsRestjson1_serializeOpDescribeStorageConfig
 func (m *awsRestjson1_serializeOpDescribeStorageConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3985,7 +3850,6 @@ func (m *awsRestjson1_serializeOpDescribeStorageConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeStorageConfigurationInput(v *DescribeStorageConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4006,10 +3870,8 @@ func (m *awsRestjson1_serializeOpDescribeTimeSeries) H
 func (m *awsRestjson1_serializeOpDescribeTimeSeries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4047,7 +3909,6 @@ func (m *awsRestjson1_serializeOpDescribeTimeSeries) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTimeSeriesInput(v *DescribeTimeSeriesInput, encoder *httpbinding.Encoder) error {
@@ -4080,10 +3941,8 @@ func (m *awsRestjson1_serializeOpDisassociateAssets) H
 func (m *awsRestjson1_serializeOpDisassociateAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4132,7 +3991,6 @@ func (m *awsRestjson1_serializeOpDisassociateAssets) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateAssetsInput(v *DisassociateAssetsInput, encoder *httpbinding.Encoder) error {
@@ -4184,10 +4042,8 @@ func (m *awsRestjson1_serializeOpDisassociateTimeSerie
 func (m *awsRestjson1_serializeOpDisassociateTimeSeriesFromAssetProperty) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4236,7 +4092,6 @@ func (m *awsRestjson1_serializeOpDisassociateTimeSerie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateTimeSeriesFromAssetPropertyInput(v *DisassociateTimeSeriesFromAssetPropertyInput, encoder *httpbinding.Encoder) error {
@@ -4281,10 +4136,8 @@ func (m *awsRestjson1_serializeOpExecuteAction) Handle
 func (m *awsRestjson1_serializeOpExecuteAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4329,7 +4182,6 @@ func (m *awsRestjson1_serializeOpExecuteAction) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteActionInput(v *ExecuteActionInput, encoder *httpbinding.Encoder) error {
@@ -4381,10 +4233,8 @@ func (m *awsRestjson1_serializeOpExecuteQuery) HandleS
 func (m *awsRestjson1_serializeOpExecuteQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4429,7 +4279,6 @@ func (m *awsRestjson1_serializeOpExecuteQuery) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteQueryInput(v *ExecuteQueryInput, encoder *httpbinding.Encoder) error {
@@ -4477,10 +4326,8 @@ func (m *awsRestjson1_serializeOpGetAssetPropertyAggre
 func (m *awsRestjson1_serializeOpGetAssetPropertyAggregates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4518,7 +4365,6 @@ func (m *awsRestjson1_serializeOpGetAssetPropertyAggre
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssetPropertyAggregatesInput(v *GetAssetPropertyAggregatesInput, encoder *httpbinding.Encoder) error {
@@ -4587,10 +4433,8 @@ func (m *awsRestjson1_serializeOpGetAssetPropertyValue
 func (m *awsRestjson1_serializeOpGetAssetPropertyValue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4628,7 +4472,6 @@ func (m *awsRestjson1_serializeOpGetAssetPropertyValue
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssetPropertyValueInput(v *GetAssetPropertyValueInput, encoder *httpbinding.Encoder) error {
@@ -4661,10 +4504,8 @@ func (m *awsRestjson1_serializeOpGetAssetPropertyValue
 func (m *awsRestjson1_serializeOpGetAssetPropertyValueHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4702,7 +4543,6 @@ func (m *awsRestjson1_serializeOpGetAssetPropertyValue
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssetPropertyValueHistoryInput(v *GetAssetPropertyValueHistoryInput, encoder *httpbinding.Encoder) error {
@@ -4761,10 +4601,8 @@ func (m *awsRestjson1_serializeOpGetInterpolatedAssetP
 func (m *awsRestjson1_serializeOpGetInterpolatedAssetPropertyValues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4802,7 +4640,6 @@ func (m *awsRestjson1_serializeOpGetInterpolatedAssetP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInterpolatedAssetPropertyValuesInput(v *GetInterpolatedAssetPropertyValuesInput, encoder *httpbinding.Encoder) error {
@@ -4875,10 +4712,8 @@ func (m *awsRestjson1_serializeOpInvokeAssistant) Hand
 func (m *awsRestjson1_serializeOpInvokeAssistant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4923,7 +4758,6 @@ func (m *awsRestjson1_serializeOpInvokeAssistant) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeAssistantInput(v *InvokeAssistantInput, encoder *httpbinding.Encoder) error {
@@ -4966,10 +4800,8 @@ func (m *awsRestjson1_serializeOpListAccessPolicies) H
 func (m *awsRestjson1_serializeOpListAccessPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5007,7 +4839,6 @@ func (m *awsRestjson1_serializeOpListAccessPolicies) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAccessPoliciesInput(v *ListAccessPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -5056,10 +4887,8 @@ func (m *awsRestjson1_serializeOpListActions) HandleSe
 func (m *awsRestjson1_serializeOpListActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5097,7 +4926,6 @@ func (m *awsRestjson1_serializeOpListActions) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListActionsInput(v *ListActionsInput, encoder *httpbinding.Encoder) error {
@@ -5134,10 +4962,8 @@ func (m *awsRestjson1_serializeOpListAssetModelComposi
 func (m *awsRestjson1_serializeOpListAssetModelCompositeModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5175,7 +5001,6 @@ func (m *awsRestjson1_serializeOpListAssetModelComposi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetModelCompositeModelsInput(v *ListAssetModelCompositeModelsInput, encoder *httpbinding.Encoder) error {
@@ -5217,10 +5042,8 @@ func (m *awsRestjson1_serializeOpListAssetModelPropert
 func (m *awsRestjson1_serializeOpListAssetModelProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5258,7 +5081,6 @@ func (m *awsRestjson1_serializeOpListAssetModelPropert
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetModelPropertiesInput(v *ListAssetModelPropertiesInput, encoder *httpbinding.Encoder) error {
@@ -5304,10 +5126,8 @@ func (m *awsRestjson1_serializeOpListAssetModels) Hand
 func (m *awsRestjson1_serializeOpListAssetModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5345,7 +5165,6 @@ func (m *awsRestjson1_serializeOpListAssetModels) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetModelsInput(v *ListAssetModelsInput, encoder *httpbinding.Encoder) error {
@@ -5384,10 +5203,8 @@ func (m *awsRestjson1_serializeOpListAssetProperties) 
 func (m *awsRestjson1_serializeOpListAssetProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5425,7 +5242,6 @@ func (m *awsRestjson1_serializeOpListAssetProperties) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetPropertiesInput(v *ListAssetPropertiesInput, encoder *httpbinding.Encoder) error {
@@ -5467,10 +5283,8 @@ func (m *awsRestjson1_serializeOpListAssetRelationship
 func (m *awsRestjson1_serializeOpListAssetRelationships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5508,7 +5322,6 @@ func (m *awsRestjson1_serializeOpListAssetRelationship
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetRelationshipsInput(v *ListAssetRelationshipsInput, encoder *httpbinding.Encoder) error {
@@ -5550,10 +5363,8 @@ func (m *awsRestjson1_serializeOpListAssets) HandleSer
 func (m *awsRestjson1_serializeOpListAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5591,7 +5402,6 @@ func (m *awsRestjson1_serializeOpListAssets) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetsInput(v *ListAssetsInput, encoder *httpbinding.Encoder) error {
@@ -5628,10 +5438,8 @@ func (m *awsRestjson1_serializeOpListAssociatedAssets)
 func (m *awsRestjson1_serializeOpListAssociatedAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5669,7 +5477,6 @@ func (m *awsRestjson1_serializeOpListAssociatedAssets)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssociatedAssetsInput(v *ListAssociatedAssetsInput, encoder *httpbinding.Encoder) error {
@@ -5715,10 +5522,8 @@ func (m *awsRestjson1_serializeOpListBulkImportJobs) H
 func (m *awsRestjson1_serializeOpListBulkImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5756,7 +5561,6 @@ func (m *awsRestjson1_serializeOpListBulkImportJobs) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBulkImportJobsInput(v *ListBulkImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -5789,10 +5593,8 @@ func (m *awsRestjson1_serializeOpListCompositionRelati
 func (m *awsRestjson1_serializeOpListCompositionRelationships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5830,7 +5632,6 @@ func (m *awsRestjson1_serializeOpListCompositionRelati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCompositionRelationshipsInput(v *ListCompositionRelationshipsInput, encoder *httpbinding.Encoder) error {
@@ -5868,10 +5669,8 @@ func (m *awsRestjson1_serializeOpListDashboards) Handl
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
@@ -5909,7 +5708,6 @@ func (m *awsRestjson1_serializeOpListDashboards) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDashboardsInput(v *ListDashboardsInput, encoder *httpbinding.Encoder) error {
@@ -5942,10 +5740,8 @@ func (m *awsRestjson1_serializeOpListDatasets) HandleS
 func (m *awsRestjson1_serializeOpListDatasets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5983,7 +5779,6 @@ func (m *awsRestjson1_serializeOpListDatasets) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDatasetsInput(v *ListDatasetsInput, encoder *httpbinding.Encoder) error {
@@ -6016,10 +5811,8 @@ func (m *awsRestjson1_serializeOpListGateways) HandleS
 func (m *awsRestjson1_serializeOpListGateways) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6057,7 +5850,6 @@ func (m *awsRestjson1_serializeOpListGateways) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGatewaysInput(v *ListGatewaysInput, encoder *httpbinding.Encoder) error {
@@ -6086,10 +5878,8 @@ func (m *awsRestjson1_serializeOpListPortals) HandleSe
 func (m *awsRestjson1_serializeOpListPortals) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6127,7 +5917,6 @@ func (m *awsRestjson1_serializeOpListPortals) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPortalsInput(v *ListPortalsInput, encoder *httpbinding.Encoder) error {
@@ -6156,10 +5945,8 @@ func (m *awsRestjson1_serializeOpListProjectAssets) Ha
 func (m *awsRestjson1_serializeOpListProjectAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6197,7 +5984,6 @@ func (m *awsRestjson1_serializeOpListProjectAssets) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProjectAssetsInput(v *ListProjectAssetsInput, encoder *httpbinding.Encoder) error {
@@ -6235,10 +6021,8 @@ func (m *awsRestjson1_serializeOpListProjects) HandleS
 func (m *awsRestjson1_serializeOpListProjects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6276,7 +6060,6 @@ func (m *awsRestjson1_serializeOpListProjects) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProjectsInput(v *ListProjectsInput, encoder *httpbinding.Encoder) error {
@@ -6309,10 +6092,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6350,7 +6131,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6375,10 +6155,8 @@ func (m *awsRestjson1_serializeOpListTimeSeries) Handl
 func (m *awsRestjson1_serializeOpListTimeSeries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6416,7 +6194,6 @@ func (m *awsRestjson1_serializeOpListTimeSeries) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTimeSeriesInput(v *ListTimeSeriesInput, encoder *httpbinding.Encoder) error {
@@ -6457,10 +6234,8 @@ func (m *awsRestjson1_serializeOpPutDefaultEncryptionC
 func (m *awsRestjson1_serializeOpPutDefaultEncryptionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6505,7 +6280,6 @@ func (m *awsRestjson1_serializeOpPutDefaultEncryptionC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDefaultEncryptionConfigurationInput(v *PutDefaultEncryptionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6543,10 +6317,8 @@ func (m *awsRestjson1_serializeOpPutLoggingOptions) Ha
 func (m *awsRestjson1_serializeOpPutLoggingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6591,7 +6363,6 @@ func (m *awsRestjson1_serializeOpPutLoggingOptions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutLoggingOptionsInput(v *PutLoggingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -6626,10 +6397,8 @@ func (m *awsRestjson1_serializeOpPutStorageConfigurati
 func (m *awsRestjson1_serializeOpPutStorageConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6674,7 +6443,6 @@ func (m *awsRestjson1_serializeOpPutStorageConfigurati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutStorageConfigurationInput(v *PutStorageConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6743,10 +6511,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -6795,7 +6561,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6834,10 +6599,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -6875,7 +6638,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6906,10 +6668,8 @@ func (m *awsRestjson1_serializeOpUpdateAccessPolicy) H
 func (m *awsRestjson1_serializeOpUpdateAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6958,7 +6718,6 @@ func (m *awsRestjson1_serializeOpUpdateAccessPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccessPolicyInput(v *UpdateAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7019,10 +6778,8 @@ func (m *awsRestjson1_serializeOpUpdateAsset) HandleSe
 func (m *awsRestjson1_serializeOpUpdateAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7071,7 +6828,6 @@ func (m *awsRestjson1_serializeOpUpdateAsset) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssetInput(v *UpdateAssetInput, encoder *httpbinding.Encoder) error {
@@ -7128,10 +6884,8 @@ func (m *awsRestjson1_serializeOpUpdateAssetModel) Han
 func (m *awsRestjson1_serializeOpUpdateAssetModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7180,7 +6934,6 @@ func (m *awsRestjson1_serializeOpUpdateAssetModel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssetModelInput(v *UpdateAssetModelInput, encoder *httpbinding.Encoder) error {
@@ -7273,10 +7026,8 @@ func (m *awsRestjson1_serializeOpUpdateAssetModelCompo
 func (m *awsRestjson1_serializeOpUpdateAssetModelCompositeModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7325,7 +7076,6 @@ func (m *awsRestjson1_serializeOpUpdateAssetModelCompo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssetModelCompositeModelInput(v *UpdateAssetModelCompositeModelInput, encoder *httpbinding.Encoder) error {
@@ -7413,10 +7163,8 @@ func (m *awsRestjson1_serializeOpUpdateAssetProperty) 
 func (m *awsRestjson1_serializeOpUpdateAssetProperty) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7465,7 +7213,6 @@ func (m *awsRestjson1_serializeOpUpdateAssetProperty) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssetPropertyInput(v *UpdateAssetPropertyInput, encoder *httpbinding.Encoder) error {
@@ -7531,10 +7278,8 @@ func (m *awsRestjson1_serializeOpUpdateDashboard) Hand
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
@@ -7583,7 +7328,6 @@ func (m *awsRestjson1_serializeOpUpdateDashboard) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDashboardInput(v *UpdateDashboardInput, encoder *httpbinding.Encoder) error {
@@ -7640,10 +7384,8 @@ func (m *awsRestjson1_serializeOpUpdateDataset) Handle
 func (m *awsRestjson1_serializeOpUpdateDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7692,7 +7434,6 @@ func (m *awsRestjson1_serializeOpUpdateDataset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDatasetInput(v *UpdateDatasetInput, encoder *httpbinding.Encoder) error {
@@ -7751,10 +7492,8 @@ func (m *awsRestjson1_serializeOpUpdateGateway) Handle
 func (m *awsRestjson1_serializeOpUpdateGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7803,7 +7542,6 @@ func (m *awsRestjson1_serializeOpUpdateGateway) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGatewayInput(v *UpdateGatewayInput, encoder *httpbinding.Encoder) error {
@@ -7845,10 +7583,8 @@ func (m *awsRestjson1_serializeOpUpdateGatewayCapabili
 func (m *awsRestjson1_serializeOpUpdateGatewayCapabilityConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7897,7 +7633,6 @@ func (m *awsRestjson1_serializeOpUpdateGatewayCapabili
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGatewayCapabilityConfigurationInput(v *UpdateGatewayCapabilityConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7944,10 +7679,8 @@ func (m *awsRestjson1_serializeOpUpdatePortal) HandleS
 func (m *awsRestjson1_serializeOpUpdatePortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7996,7 +7729,6 @@ func (m *awsRestjson1_serializeOpUpdatePortal) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePortalInput(v *UpdatePortalInput, encoder *httpbinding.Encoder) error {
@@ -8089,10 +7821,8 @@ func (m *awsRestjson1_serializeOpUpdateProject) Handle
 func (m *awsRestjson1_serializeOpUpdateProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8141,7 +7871,6 @@ func (m *awsRestjson1_serializeOpUpdateProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProjectInput(v *UpdateProjectInput, encoder *httpbinding.Encoder) error {
