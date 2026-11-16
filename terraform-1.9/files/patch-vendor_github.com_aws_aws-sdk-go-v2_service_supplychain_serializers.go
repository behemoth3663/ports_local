--- vendor/github.com/aws/aws-sdk-go-v2/service/supplychain/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/supplychain/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCreateBillOfMaterials
 func (m *awsRestjson1_serializeOpCreateBillOfMaterialsImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -78,7 +75,6 @@ func (m *awsRestjson1_serializeOpCreateBillOfMaterials
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBillOfMaterialsImportJobInput(v *CreateBillOfMaterialsImportJobInput, encoder *httpbinding.Encoder) error {
@@ -125,10 +121,8 @@ func (m *awsRestjson1_serializeOpCreateDataIntegration
 func (m *awsRestjson1_serializeOpCreateDataIntegrationFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -177,7 +171,6 @@ func (m *awsRestjson1_serializeOpCreateDataIntegration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataIntegrationFlowInput(v *CreateDataIntegrationFlowInput, encoder *httpbinding.Encoder) error {
@@ -251,10 +244,8 @@ func (m *awsRestjson1_serializeOpCreateDataLakeDataset
 func (m *awsRestjson1_serializeOpCreateDataLakeDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -303,7 +294,6 @@ func (m *awsRestjson1_serializeOpCreateDataLakeDataset
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataLakeDatasetInput(v *CreateDataLakeDatasetInput, encoder *httpbinding.Encoder) error {
@@ -384,10 +374,8 @@ func (m *awsRestjson1_serializeOpCreateDataLakeNamespa
 func (m *awsRestjson1_serializeOpCreateDataLakeNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -436,7 +424,6 @@ func (m *awsRestjson1_serializeOpCreateDataLakeNamespa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataLakeNamespaceInput(v *CreateDataLakeNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -494,10 +481,8 @@ func (m *awsRestjson1_serializeOpCreateInstance) Handl
 func (m *awsRestjson1_serializeOpCreateInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -542,7 +527,6 @@ func (m *awsRestjson1_serializeOpCreateInstance) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInstanceInput(v *CreateInstanceInput, encoder *httpbinding.Encoder) error {
@@ -602,10 +586,8 @@ func (m *awsRestjson1_serializeOpDeleteDataIntegration
 func (m *awsRestjson1_serializeOpDeleteDataIntegrationFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -643,7 +625,6 @@ func (m *awsRestjson1_serializeOpDeleteDataIntegration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataIntegrationFlowInput(v *DeleteDataIntegrationFlowInput, encoder *httpbinding.Encoder) error {
@@ -682,10 +663,8 @@ func (m *awsRestjson1_serializeOpDeleteDataLakeDataset
 func (m *awsRestjson1_serializeOpDeleteDataLakeDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -723,7 +702,6 @@ func (m *awsRestjson1_serializeOpDeleteDataLakeDataset
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataLakeDatasetInput(v *DeleteDataLakeDatasetInput, encoder *httpbinding.Encoder) error {
@@ -771,10 +749,8 @@ func (m *awsRestjson1_serializeOpDeleteDataLakeNamespa
 func (m *awsRestjson1_serializeOpDeleteDataLakeNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -812,7 +788,6 @@ func (m *awsRestjson1_serializeOpDeleteDataLakeNamespa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataLakeNamespaceInput(v *DeleteDataLakeNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -851,10 +826,8 @@ func (m *awsRestjson1_serializeOpDeleteInstance) Handl
 func (m *awsRestjson1_serializeOpDeleteInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -892,7 +865,6 @@ func (m *awsRestjson1_serializeOpDeleteInstance) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInstanceInput(v *DeleteInstanceInput, encoder *httpbinding.Encoder) error {
@@ -922,10 +894,8 @@ func (m *awsRestjson1_serializeOpGetBillOfMaterialsImp
 func (m *awsRestjson1_serializeOpGetBillOfMaterialsImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -963,7 +933,6 @@ func (m *awsRestjson1_serializeOpGetBillOfMaterialsImp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBillOfMaterialsImportJobInput(v *GetBillOfMaterialsImportJobInput, encoder *httpbinding.Encoder) error {
@@ -1002,10 +971,8 @@ func (m *awsRestjson1_serializeOpGetDataIntegrationEve
 func (m *awsRestjson1_serializeOpGetDataIntegrationEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1043,7 +1010,6 @@ func (m *awsRestjson1_serializeOpGetDataIntegrationEve
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataIntegrationEventInput(v *GetDataIntegrationEventInput, encoder *httpbinding.Encoder) error {
@@ -1082,10 +1048,8 @@ func (m *awsRestjson1_serializeOpGetDataIntegrationFlo
 func (m *awsRestjson1_serializeOpGetDataIntegrationFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1123,7 +1087,6 @@ func (m *awsRestjson1_serializeOpGetDataIntegrationFlo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataIntegrationFlowInput(v *GetDataIntegrationFlowInput, encoder *httpbinding.Encoder) error {
@@ -1162,10 +1125,8 @@ func (m *awsRestjson1_serializeOpGetDataIntegrationFlo
 func (m *awsRestjson1_serializeOpGetDataIntegrationFlowExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1203,7 +1164,6 @@ func (m *awsRestjson1_serializeOpGetDataIntegrationFlo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataIntegrationFlowExecutionInput(v *GetDataIntegrationFlowExecutionInput, encoder *httpbinding.Encoder) error {
@@ -1251,10 +1211,8 @@ func (m *awsRestjson1_serializeOpGetDataLakeDataset) H
 func (m *awsRestjson1_serializeOpGetDataLakeDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1292,7 +1250,6 @@ func (m *awsRestjson1_serializeOpGetDataLakeDataset) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataLakeDatasetInput(v *GetDataLakeDatasetInput, encoder *httpbinding.Encoder) error {
@@ -1340,10 +1297,8 @@ func (m *awsRestjson1_serializeOpGetDataLakeNamespace)
 func (m *awsRestjson1_serializeOpGetDataLakeNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1381,7 +1336,6 @@ func (m *awsRestjson1_serializeOpGetDataLakeNamespace)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataLakeNamespaceInput(v *GetDataLakeNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -1420,10 +1374,8 @@ func (m *awsRestjson1_serializeOpGetInstance) HandleSe
 func (m *awsRestjson1_serializeOpGetInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1461,7 +1413,6 @@ func (m *awsRestjson1_serializeOpGetInstance) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInstanceInput(v *GetInstanceInput, encoder *httpbinding.Encoder) error {
@@ -1491,10 +1442,8 @@ func (m *awsRestjson1_serializeOpListDataIntegrationEv
 func (m *awsRestjson1_serializeOpListDataIntegrationEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1532,7 +1481,6 @@ func (m *awsRestjson1_serializeOpListDataIntegrationEv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataIntegrationEventsInput(v *ListDataIntegrationEventsInput, encoder *httpbinding.Encoder) error {
@@ -1574,10 +1522,8 @@ func (m *awsRestjson1_serializeOpListDataIntegrationFl
 func (m *awsRestjson1_serializeOpListDataIntegrationFlowExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1615,7 +1561,6 @@ func (m *awsRestjson1_serializeOpListDataIntegrationFl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataIntegrationFlowExecutionsInput(v *ListDataIntegrationFlowExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -1662,10 +1607,8 @@ func (m *awsRestjson1_serializeOpListDataIntegrationFl
 func (m *awsRestjson1_serializeOpListDataIntegrationFlows) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1703,7 +1646,6 @@ func (m *awsRestjson1_serializeOpListDataIntegrationFl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataIntegrationFlowsInput(v *ListDataIntegrationFlowsInput, encoder *httpbinding.Encoder) error {
@@ -1741,10 +1683,8 @@ func (m *awsRestjson1_serializeOpListDataLakeDatasets)
 func (m *awsRestjson1_serializeOpListDataLakeDatasets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1782,7 +1722,6 @@ func (m *awsRestjson1_serializeOpListDataLakeDatasets)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataLakeDatasetsInput(v *ListDataLakeDatasetsInput, encoder *httpbinding.Encoder) error {
@@ -1829,10 +1768,8 @@ func (m *awsRestjson1_serializeOpListDataLakeNamespace
 func (m *awsRestjson1_serializeOpListDataLakeNamespaces) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1870,7 +1807,6 @@ func (m *awsRestjson1_serializeOpListDataLakeNamespace
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataLakeNamespacesInput(v *ListDataLakeNamespacesInput, encoder *httpbinding.Encoder) error {
@@ -1908,10 +1844,8 @@ func (m *awsRestjson1_serializeOpListInstances) Handle
 func (m *awsRestjson1_serializeOpListInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1949,7 +1883,6 @@ func (m *awsRestjson1_serializeOpListInstances) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInstancesInput(v *ListInstancesInput, encoder *httpbinding.Encoder) error {
@@ -1990,10 +1923,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -2031,7 +1962,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -2061,10 +1991,8 @@ func (m *awsRestjson1_serializeOpSendDataIntegrationEv
 func (m *awsRestjson1_serializeOpSendDataIntegrationEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2113,7 +2041,6 @@ func (m *awsRestjson1_serializeOpSendDataIntegrationEv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendDataIntegrationEventInput(v *SendDataIntegrationEventInput, encoder *httpbinding.Encoder) error {
@@ -2182,10 +2109,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -2234,7 +2159,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2278,10 +2202,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -2319,7 +2241,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2355,10 +2276,8 @@ func (m *awsRestjson1_serializeOpUpdateDataIntegration
 func (m *awsRestjson1_serializeOpUpdateDataIntegrationFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2407,7 +2326,6 @@ func (m *awsRestjson1_serializeOpUpdateDataIntegration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataIntegrationFlowInput(v *UpdateDataIntegrationFlowInput, encoder *httpbinding.Encoder) error {
@@ -2474,10 +2392,8 @@ func (m *awsRestjson1_serializeOpUpdateDataLakeDataset
 func (m *awsRestjson1_serializeOpUpdateDataLakeDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2526,7 +2442,6 @@ func (m *awsRestjson1_serializeOpUpdateDataLakeDataset
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataLakeDatasetInput(v *UpdateDataLakeDatasetInput, encoder *httpbinding.Encoder) error {
@@ -2586,10 +2501,8 @@ func (m *awsRestjson1_serializeOpUpdateDataLakeNamespa
 func (m *awsRestjson1_serializeOpUpdateDataLakeNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2638,7 +2551,6 @@ func (m *awsRestjson1_serializeOpUpdateDataLakeNamespa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataLakeNamespaceInput(v *UpdateDataLakeNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -2689,10 +2601,8 @@ func (m *awsRestjson1_serializeOpUpdateInstance) Handl
 func (m *awsRestjson1_serializeOpUpdateInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2741,7 +2651,6 @@ func (m *awsRestjson1_serializeOpUpdateInstance) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInstanceInput(v *UpdateInstanceInput, encoder *httpbinding.Encoder) error {
