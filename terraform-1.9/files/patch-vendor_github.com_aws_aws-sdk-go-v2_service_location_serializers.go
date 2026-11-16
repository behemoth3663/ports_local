--- vendor/github.com/aws/aws-sdk-go-v2/service/location/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/location/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAssociateTrackerConsu
 func (m *awsRestjson1_serializeOpAssociateTrackerConsumer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpAssociateTrackerConsu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateTrackerConsumerInput(v *AssociateTrackerConsumerInput, encoder *httpbinding.Encoder) error {
@@ -121,10 +117,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteDevicePosi
 func (m *awsRestjson1_serializeOpBatchDeleteDevicePositionHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -173,7 +167,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteDevicePosi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteDevicePositionHistoryInput(v *BatchDeleteDevicePositionHistoryInput, encoder *httpbinding.Encoder) error {
@@ -217,10 +210,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteGeofence) 
 func (m *awsRestjson1_serializeOpBatchDeleteGeofence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -269,7 +260,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteGeofence) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteGeofenceInput(v *BatchDeleteGeofenceInput, encoder *httpbinding.Encoder) error {
@@ -313,10 +303,8 @@ func (m *awsRestjson1_serializeOpBatchEvaluateGeofence
 func (m *awsRestjson1_serializeOpBatchEvaluateGeofences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -365,7 +353,6 @@ func (m *awsRestjson1_serializeOpBatchEvaluateGeofence
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchEvaluateGeofencesInput(v *BatchEvaluateGeofencesInput, encoder *httpbinding.Encoder) error {
@@ -409,10 +396,8 @@ func (m *awsRestjson1_serializeOpBatchGetDevicePositio
 func (m *awsRestjson1_serializeOpBatchGetDevicePosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -461,7 +446,6 @@ func (m *awsRestjson1_serializeOpBatchGetDevicePositio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetDevicePositionInput(v *BatchGetDevicePositionInput, encoder *httpbinding.Encoder) error {
@@ -505,10 +489,8 @@ func (m *awsRestjson1_serializeOpBatchPutGeofence) Han
 func (m *awsRestjson1_serializeOpBatchPutGeofence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -557,7 +539,6 @@ func (m *awsRestjson1_serializeOpBatchPutGeofence) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchPutGeofenceInput(v *BatchPutGeofenceInput, encoder *httpbinding.Encoder) error {
@@ -601,10 +582,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateDevicePosi
 func (m *awsRestjson1_serializeOpBatchUpdateDevicePosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -653,7 +632,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateDevicePosi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateDevicePositionInput(v *BatchUpdateDevicePositionInput, encoder *httpbinding.Encoder) error {
@@ -697,10 +675,8 @@ func (m *awsRestjson1_serializeOpCalculateRoute) Handl
 func (m *awsRestjson1_serializeOpCalculateRoute) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -749,7 +725,6 @@ func (m *awsRestjson1_serializeOpCalculateRoute) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCalculateRouteInput(v *CalculateRouteInput, encoder *httpbinding.Encoder) error {
@@ -860,10 +835,8 @@ func (m *awsRestjson1_serializeOpCalculateRouteMatrix)
 func (m *awsRestjson1_serializeOpCalculateRouteMatrix) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -912,7 +885,6 @@ func (m *awsRestjson1_serializeOpCalculateRouteMatrix)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCalculateRouteMatrixInput(v *CalculateRouteMatrixInput, encoder *httpbinding.Encoder) error {
@@ -1001,10 +973,8 @@ func (m *awsRestjson1_serializeOpCreateGeofenceCollect
 func (m *awsRestjson1_serializeOpCreateGeofenceCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1049,7 +1019,6 @@ func (m *awsRestjson1_serializeOpCreateGeofenceCollect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGeofenceCollectionInput(v *CreateGeofenceCollectionInput, encoder *httpbinding.Encoder) error {
@@ -1109,10 +1078,8 @@ func (m *awsRestjson1_serializeOpCreateKey) HandleSeri
 func (m *awsRestjson1_serializeOpCreateKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1157,7 +1124,6 @@ func (m *awsRestjson1_serializeOpCreateKey) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateKeyInput(v *CreateKeyInput, encoder *httpbinding.Encoder) error {
@@ -1219,10 +1185,8 @@ func (m *awsRestjson1_serializeOpCreateMap) HandleSeri
 func (m *awsRestjson1_serializeOpCreateMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1267,7 +1231,6 @@ func (m *awsRestjson1_serializeOpCreateMap) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMapInput(v *CreateMapInput, encoder *httpbinding.Encoder) error {
@@ -1324,10 +1287,8 @@ func (m *awsRestjson1_serializeOpCreatePlaceIndex) Han
 func (m *awsRestjson1_serializeOpCreatePlaceIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1372,7 +1333,6 @@ func (m *awsRestjson1_serializeOpCreatePlaceIndex) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePlaceIndexInput(v *CreatePlaceIndexInput, encoder *httpbinding.Encoder) error {
@@ -1434,10 +1394,8 @@ func (m *awsRestjson1_serializeOpCreateRouteCalculator
 func (m *awsRestjson1_serializeOpCreateRouteCalculator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1482,7 +1440,6 @@ func (m *awsRestjson1_serializeOpCreateRouteCalculator
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRouteCalculatorInput(v *CreateRouteCalculatorInput, encoder *httpbinding.Encoder) error {
@@ -1537,10 +1494,8 @@ func (m *awsRestjson1_serializeOpCreateTracker) Handle
 func (m *awsRestjson1_serializeOpCreateTracker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1585,7 +1540,6 @@ func (m *awsRestjson1_serializeOpCreateTracker) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTrackerInput(v *CreateTrackerInput, encoder *httpbinding.Encoder) error {
@@ -1660,10 +1614,8 @@ func (m *awsRestjson1_serializeOpDeleteGeofenceCollect
 func (m *awsRestjson1_serializeOpDeleteGeofenceCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1701,7 +1653,6 @@ func (m *awsRestjson1_serializeOpDeleteGeofenceCollect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGeofenceCollectionInput(v *DeleteGeofenceCollectionInput, encoder *httpbinding.Encoder) error {
@@ -1731,10 +1682,8 @@ func (m *awsRestjson1_serializeOpDeleteKey) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1772,7 +1721,6 @@ func (m *awsRestjson1_serializeOpDeleteKey) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteKeyInput(v *DeleteKeyInput, encoder *httpbinding.Encoder) error {
@@ -1806,10 +1754,8 @@ func (m *awsRestjson1_serializeOpDeleteMap) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1847,7 +1793,6 @@ func (m *awsRestjson1_serializeOpDeleteMap) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMapInput(v *DeleteMapInput, encoder *httpbinding.Encoder) error {
@@ -1877,10 +1822,8 @@ func (m *awsRestjson1_serializeOpDeletePlaceIndex) Han
 func (m *awsRestjson1_serializeOpDeletePlaceIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1918,7 +1861,6 @@ func (m *awsRestjson1_serializeOpDeletePlaceIndex) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePlaceIndexInput(v *DeletePlaceIndexInput, encoder *httpbinding.Encoder) error {
@@ -1948,10 +1890,8 @@ func (m *awsRestjson1_serializeOpDeleteRouteCalculator
 func (m *awsRestjson1_serializeOpDeleteRouteCalculator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1989,7 +1929,6 @@ func (m *awsRestjson1_serializeOpDeleteRouteCalculator
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRouteCalculatorInput(v *DeleteRouteCalculatorInput, encoder *httpbinding.Encoder) error {
@@ -2019,10 +1958,8 @@ func (m *awsRestjson1_serializeOpDeleteTracker) Handle
 func (m *awsRestjson1_serializeOpDeleteTracker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2060,7 +1997,6 @@ func (m *awsRestjson1_serializeOpDeleteTracker) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTrackerInput(v *DeleteTrackerInput, encoder *httpbinding.Encoder) error {
@@ -2090,10 +2026,8 @@ func (m *awsRestjson1_serializeOpDescribeGeofenceColle
 func (m *awsRestjson1_serializeOpDescribeGeofenceCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2131,7 +2065,6 @@ func (m *awsRestjson1_serializeOpDescribeGeofenceColle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeGeofenceCollectionInput(v *DescribeGeofenceCollectionInput, encoder *httpbinding.Encoder) error {
@@ -2161,10 +2094,8 @@ func (m *awsRestjson1_serializeOpDescribeKey) HandleSe
 func (m *awsRestjson1_serializeOpDescribeKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2202,7 +2133,6 @@ func (m *awsRestjson1_serializeOpDescribeKey) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeKeyInput(v *DescribeKeyInput, encoder *httpbinding.Encoder) error {
@@ -2232,10 +2162,8 @@ func (m *awsRestjson1_serializeOpDescribeMap) HandleSe
 func (m *awsRestjson1_serializeOpDescribeMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2273,7 +2201,6 @@ func (m *awsRestjson1_serializeOpDescribeMap) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeMapInput(v *DescribeMapInput, encoder *httpbinding.Encoder) error {
@@ -2303,10 +2230,8 @@ func (m *awsRestjson1_serializeOpDescribePlaceIndex) H
 func (m *awsRestjson1_serializeOpDescribePlaceIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2344,7 +2269,6 @@ func (m *awsRestjson1_serializeOpDescribePlaceIndex) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePlaceIndexInput(v *DescribePlaceIndexInput, encoder *httpbinding.Encoder) error {
@@ -2374,10 +2298,8 @@ func (m *awsRestjson1_serializeOpDescribeRouteCalculat
 func (m *awsRestjson1_serializeOpDescribeRouteCalculator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2415,7 +2337,6 @@ func (m *awsRestjson1_serializeOpDescribeRouteCalculat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRouteCalculatorInput(v *DescribeRouteCalculatorInput, encoder *httpbinding.Encoder) error {
@@ -2445,10 +2366,8 @@ func (m *awsRestjson1_serializeOpDescribeTracker) Hand
 func (m *awsRestjson1_serializeOpDescribeTracker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2486,7 +2405,6 @@ func (m *awsRestjson1_serializeOpDescribeTracker) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeTrackerInput(v *DescribeTrackerInput, encoder *httpbinding.Encoder) error {
@@ -2516,10 +2434,8 @@ func (m *awsRestjson1_serializeOpDisassociateTrackerCo
 func (m *awsRestjson1_serializeOpDisassociateTrackerConsumer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2557,7 +2473,6 @@ func (m *awsRestjson1_serializeOpDisassociateTrackerCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateTrackerConsumerInput(v *DisassociateTrackerConsumerInput, encoder *httpbinding.Encoder) error {
@@ -2596,10 +2511,8 @@ func (m *awsRestjson1_serializeOpForecastGeofenceEvent
 func (m *awsRestjson1_serializeOpForecastGeofenceEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2648,7 +2561,6 @@ func (m *awsRestjson1_serializeOpForecastGeofenceEvent
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsForecastGeofenceEventsInput(v *ForecastGeofenceEventsInput, encoder *httpbinding.Encoder) error {
@@ -2730,10 +2642,8 @@ func (m *awsRestjson1_serializeOpGetDevicePosition) Ha
 func (m *awsRestjson1_serializeOpGetDevicePosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2771,7 +2681,6 @@ func (m *awsRestjson1_serializeOpGetDevicePosition) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDevicePositionInput(v *GetDevicePositionInput, encoder *httpbinding.Encoder) error {
@@ -2810,10 +2719,8 @@ func (m *awsRestjson1_serializeOpGetDevicePositionHist
 func (m *awsRestjson1_serializeOpGetDevicePositionHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2862,7 +2769,6 @@ func (m *awsRestjson1_serializeOpGetDevicePositionHist
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDevicePositionHistoryInput(v *GetDevicePositionHistoryInput, encoder *httpbinding.Encoder) error {
@@ -2928,10 +2834,8 @@ func (m *awsRestjson1_serializeOpGetGeofence) HandleSe
 func (m *awsRestjson1_serializeOpGetGeofence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2969,7 +2873,6 @@ func (m *awsRestjson1_serializeOpGetGeofence) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGeofenceInput(v *GetGeofenceInput, encoder *httpbinding.Encoder) error {
@@ -3008,10 +2911,8 @@ func (m *awsRestjson1_serializeOpGetMapGlyphs) HandleS
 func (m *awsRestjson1_serializeOpGetMapGlyphs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3049,7 +2950,6 @@ func (m *awsRestjson1_serializeOpGetMapGlyphs) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMapGlyphsInput(v *GetMapGlyphsInput, encoder *httpbinding.Encoder) error {
@@ -3101,10 +3001,8 @@ func (m *awsRestjson1_serializeOpGetMapSprites) Handle
 func (m *awsRestjson1_serializeOpGetMapSprites) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3142,7 +3040,6 @@ func (m *awsRestjson1_serializeOpGetMapSprites) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMapSpritesInput(v *GetMapSpritesInput, encoder *httpbinding.Encoder) error {
@@ -3185,10 +3082,8 @@ func (m *awsRestjson1_serializeOpGetMapStyleDescriptor
 func (m *awsRestjson1_serializeOpGetMapStyleDescriptor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3226,7 +3121,6 @@ func (m *awsRestjson1_serializeOpGetMapStyleDescriptor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMapStyleDescriptorInput(v *GetMapStyleDescriptorInput, encoder *httpbinding.Encoder) error {
@@ -3260,10 +3154,8 @@ func (m *awsRestjson1_serializeOpGetMapTile) HandleSer
 func (m *awsRestjson1_serializeOpGetMapTile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3301,7 +3193,6 @@ func (m *awsRestjson1_serializeOpGetMapTile) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMapTileInput(v *GetMapTileInput, encoder *httpbinding.Encoder) error {
@@ -3362,10 +3253,8 @@ func (m *awsRestjson1_serializeOpGetPlace) HandleSeria
 func (m *awsRestjson1_serializeOpGetPlace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3403,7 +3292,6 @@ func (m *awsRestjson1_serializeOpGetPlace) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPlaceInput(v *GetPlaceInput, encoder *httpbinding.Encoder) error {
@@ -3450,10 +3338,8 @@ func (m *awsRestjson1_serializeOpListDevicePositions) 
 func (m *awsRestjson1_serializeOpListDevicePositions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3502,7 +3388,6 @@ func (m *awsRestjson1_serializeOpListDevicePositions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDevicePositionsInput(v *ListDevicePositionsInput, encoder *httpbinding.Encoder) error {
@@ -3556,10 +3441,8 @@ func (m *awsRestjson1_serializeOpListGeofenceCollectio
 func (m *awsRestjson1_serializeOpListGeofenceCollections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3604,7 +3487,6 @@ func (m *awsRestjson1_serializeOpListGeofenceCollectio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGeofenceCollectionsInput(v *ListGeofenceCollectionsInput, encoder *httpbinding.Encoder) error {
@@ -3642,10 +3524,8 @@ func (m *awsRestjson1_serializeOpListGeofences) Handle
 func (m *awsRestjson1_serializeOpListGeofences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3694,7 +3574,6 @@ func (m *awsRestjson1_serializeOpListGeofences) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGeofencesInput(v *ListGeofencesInput, encoder *httpbinding.Encoder) error {
@@ -3741,10 +3620,8 @@ func (m *awsRestjson1_serializeOpListKeys) HandleSeria
 func (m *awsRestjson1_serializeOpListKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3789,7 +3666,6 @@ func (m *awsRestjson1_serializeOpListKeys) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListKeysInput(v *ListKeysInput, encoder *httpbinding.Encoder) error {
@@ -3834,10 +3710,8 @@ func (m *awsRestjson1_serializeOpListMaps) HandleSeria
 func (m *awsRestjson1_serializeOpListMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3882,7 +3756,6 @@ func (m *awsRestjson1_serializeOpListMaps) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMapsInput(v *ListMapsInput, encoder *httpbinding.Encoder) error {
@@ -3920,10 +3793,8 @@ func (m *awsRestjson1_serializeOpListPlaceIndexes) Han
 func (m *awsRestjson1_serializeOpListPlaceIndexes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3968,7 +3839,6 @@ func (m *awsRestjson1_serializeOpListPlaceIndexes) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPlaceIndexesInput(v *ListPlaceIndexesInput, encoder *httpbinding.Encoder) error {
@@ -4006,10 +3876,8 @@ func (m *awsRestjson1_serializeOpListRouteCalculators)
 func (m *awsRestjson1_serializeOpListRouteCalculators) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4054,7 +3922,6 @@ func (m *awsRestjson1_serializeOpListRouteCalculators)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRouteCalculatorsInput(v *ListRouteCalculatorsInput, encoder *httpbinding.Encoder) error {
@@ -4092,10 +3959,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -4133,7 +3998,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4163,10 +4027,8 @@ func (m *awsRestjson1_serializeOpListTrackerConsumers)
 func (m *awsRestjson1_serializeOpListTrackerConsumers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4215,7 +4077,6 @@ func (m *awsRestjson1_serializeOpListTrackerConsumers)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrackerConsumersInput(v *ListTrackerConsumersInput, encoder *httpbinding.Encoder) error {
@@ -4262,10 +4123,8 @@ func (m *awsRestjson1_serializeOpListTrackers) HandleS
 func (m *awsRestjson1_serializeOpListTrackers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4310,7 +4169,6 @@ func (m *awsRestjson1_serializeOpListTrackers) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrackersInput(v *ListTrackersInput, encoder *httpbinding.Encoder) error {
@@ -4348,10 +4206,8 @@ func (m *awsRestjson1_serializeOpPutGeofence) HandleSe
 func (m *awsRestjson1_serializeOpPutGeofence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4400,7 +4256,6 @@ func (m *awsRestjson1_serializeOpPutGeofence) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutGeofenceInput(v *PutGeofenceInput, encoder *httpbinding.Encoder) error {
@@ -4460,10 +4315,8 @@ func (m *awsRestjson1_serializeOpSearchPlaceIndexForPo
 func (m *awsRestjson1_serializeOpSearchPlaceIndexForPosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4512,7 +4365,6 @@ func (m *awsRestjson1_serializeOpSearchPlaceIndexForPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchPlaceIndexForPositionInput(v *SearchPlaceIndexForPositionInput, encoder *httpbinding.Encoder) error {
@@ -4570,10 +4422,8 @@ func (m *awsRestjson1_serializeOpSearchPlaceIndexForSu
 func (m *awsRestjson1_serializeOpSearchPlaceIndexForSuggestions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4622,7 +4472,6 @@ func (m *awsRestjson1_serializeOpSearchPlaceIndexForSu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchPlaceIndexForSuggestionsInput(v *SearchPlaceIndexForSuggestionsInput, encoder *httpbinding.Encoder) error {
@@ -4706,10 +4555,8 @@ func (m *awsRestjson1_serializeOpSearchPlaceIndexForTe
 func (m *awsRestjson1_serializeOpSearchPlaceIndexForText) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4758,7 +4605,6 @@ func (m *awsRestjson1_serializeOpSearchPlaceIndexForTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchPlaceIndexForTextInput(v *SearchPlaceIndexForTextInput, encoder *httpbinding.Encoder) error {
@@ -4842,10 +4688,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -4894,7 +4738,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4938,10 +4781,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -4979,7 +4820,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5015,10 +4855,8 @@ func (m *awsRestjson1_serializeOpUpdateGeofenceCollect
 func (m *awsRestjson1_serializeOpUpdateGeofenceCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5067,7 +4905,6 @@ func (m *awsRestjson1_serializeOpUpdateGeofenceCollect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGeofenceCollectionInput(v *UpdateGeofenceCollectionInput, encoder *httpbinding.Encoder) error {
@@ -5119,10 +4956,8 @@ func (m *awsRestjson1_serializeOpUpdateKey) HandleSeri
 func (m *awsRestjson1_serializeOpUpdateKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5171,7 +5006,6 @@ func (m *awsRestjson1_serializeOpUpdateKey) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateKeyInput(v *UpdateKeyInput, encoder *httpbinding.Encoder) error {
@@ -5235,10 +5069,8 @@ func (m *awsRestjson1_serializeOpUpdateMap) HandleSeri
 func (m *awsRestjson1_serializeOpUpdateMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5287,7 +5119,6 @@ func (m *awsRestjson1_serializeOpUpdateMap) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMapInput(v *UpdateMapInput, encoder *httpbinding.Encoder) error {
@@ -5341,10 +5172,8 @@ func (m *awsRestjson1_serializeOpUpdatePlaceIndex) Han
 func (m *awsRestjson1_serializeOpUpdatePlaceIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5393,7 +5222,6 @@ func (m *awsRestjson1_serializeOpUpdatePlaceIndex) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePlaceIndexInput(v *UpdatePlaceIndexInput, encoder *httpbinding.Encoder) error {
@@ -5447,10 +5275,8 @@ func (m *awsRestjson1_serializeOpUpdateRouteCalculator
 func (m *awsRestjson1_serializeOpUpdateRouteCalculator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5499,7 +5325,6 @@ func (m *awsRestjson1_serializeOpUpdateRouteCalculator
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRouteCalculatorInput(v *UpdateRouteCalculatorInput, encoder *httpbinding.Encoder) error {
@@ -5546,10 +5371,8 @@ func (m *awsRestjson1_serializeOpUpdateTracker) Handle
 func (m *awsRestjson1_serializeOpUpdateTracker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5598,7 +5421,6 @@ func (m *awsRestjson1_serializeOpUpdateTracker) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTrackerInput(v *UpdateTrackerInput, encoder *httpbinding.Encoder) error {
@@ -5665,10 +5487,8 @@ func (m *awsRestjson1_serializeOpVerifyDevicePosition)
 func (m *awsRestjson1_serializeOpVerifyDevicePosition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5717,7 +5537,6 @@ func (m *awsRestjson1_serializeOpVerifyDevicePosition)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsVerifyDevicePositionInput(v *VerifyDevicePositionInput, encoder *httpbinding.Encoder) error {
