--- vendor/github.com/aws/aws-sdk-go-v2/service/cognitosync/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cognitosync/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpBulkPublish) HandleSe
 func (m *awsRestjson1_serializeOpBulkPublish) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_serializeOpBulkPublish) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBulkPublishInput(v *BulkPublishInput, encoder *httpbinding.Encoder) error {
@@ -97,10 +93,8 @@ func (m *awsRestjson1_serializeOpDeleteDataset) Handle
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
@@ -138,7 +132,6 @@ func (m *awsRestjson1_serializeOpDeleteDataset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDatasetInput(v *DeleteDatasetInput, encoder *httpbinding.Encoder) error {
@@ -186,10 +179,8 @@ func (m *awsRestjson1_serializeOpDescribeDataset) Hand
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
@@ -227,7 +218,6 @@ func (m *awsRestjson1_serializeOpDescribeDataset) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDatasetInput(v *DescribeDatasetInput, encoder *httpbinding.Encoder) error {
@@ -275,10 +265,8 @@ func (m *awsRestjson1_serializeOpDescribeIdentityPoolU
 func (m *awsRestjson1_serializeOpDescribeIdentityPoolUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -316,7 +304,6 @@ func (m *awsRestjson1_serializeOpDescribeIdentityPoolU
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeIdentityPoolUsageInput(v *DescribeIdentityPoolUsageInput, encoder *httpbinding.Encoder) error {
@@ -346,10 +333,8 @@ func (m *awsRestjson1_serializeOpDescribeIdentityUsage
 func (m *awsRestjson1_serializeOpDescribeIdentityUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -387,7 +372,6 @@ func (m *awsRestjson1_serializeOpDescribeIdentityUsage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeIdentityUsageInput(v *DescribeIdentityUsageInput, encoder *httpbinding.Encoder) error {
@@ -426,10 +410,8 @@ func (m *awsRestjson1_serializeOpGetBulkPublishDetails
 func (m *awsRestjson1_serializeOpGetBulkPublishDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -467,7 +449,6 @@ func (m *awsRestjson1_serializeOpGetBulkPublishDetails
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBulkPublishDetailsInput(v *GetBulkPublishDetailsInput, encoder *httpbinding.Encoder) error {
@@ -497,10 +478,8 @@ func (m *awsRestjson1_serializeOpGetCognitoEvents) Han
 func (m *awsRestjson1_serializeOpGetCognitoEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -538,7 +517,6 @@ func (m *awsRestjson1_serializeOpGetCognitoEvents) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCognitoEventsInput(v *GetCognitoEventsInput, encoder *httpbinding.Encoder) error {
@@ -568,10 +546,8 @@ func (m *awsRestjson1_serializeOpGetIdentityPoolConfig
 func (m *awsRestjson1_serializeOpGetIdentityPoolConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -609,7 +585,6 @@ func (m *awsRestjson1_serializeOpGetIdentityPoolConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIdentityPoolConfigurationInput(v *GetIdentityPoolConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -639,10 +614,8 @@ func (m *awsRestjson1_serializeOpListDatasets) HandleS
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
@@ -680,7 +653,6 @@ func (m *awsRestjson1_serializeOpListDatasets) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDatasetsInput(v *ListDatasetsInput, encoder *httpbinding.Encoder) error {
@@ -727,10 +699,8 @@ func (m *awsRestjson1_serializeOpListIdentityPoolUsage
 func (m *awsRestjson1_serializeOpListIdentityPoolUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -768,7 +738,6 @@ func (m *awsRestjson1_serializeOpListIdentityPoolUsage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIdentityPoolUsageInput(v *ListIdentityPoolUsageInput, encoder *httpbinding.Encoder) error {
@@ -797,10 +766,8 @@ func (m *awsRestjson1_serializeOpListRecords) HandleSe
 func (m *awsRestjson1_serializeOpListRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -838,7 +805,6 @@ func (m *awsRestjson1_serializeOpListRecords) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecordsInput(v *ListRecordsInput, encoder *httpbinding.Encoder) error {
@@ -902,10 +868,8 @@ func (m *awsRestjson1_serializeOpRegisterDevice) Handl
 func (m *awsRestjson1_serializeOpRegisterDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -954,7 +918,6 @@ func (m *awsRestjson1_serializeOpRegisterDevice) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterDeviceInput(v *RegisterDeviceInput, encoder *httpbinding.Encoder) error {
@@ -1010,10 +973,8 @@ func (m *awsRestjson1_serializeOpSetCognitoEvents) Han
 func (m *awsRestjson1_serializeOpSetCognitoEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1062,7 +1023,6 @@ func (m *awsRestjson1_serializeOpSetCognitoEvents) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetCognitoEventsInput(v *SetCognitoEventsInput, encoder *httpbinding.Encoder) error {
@@ -1106,10 +1066,8 @@ func (m *awsRestjson1_serializeOpSetIdentityPoolConfig
 func (m *awsRestjson1_serializeOpSetIdentityPoolConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1158,7 +1116,6 @@ func (m *awsRestjson1_serializeOpSetIdentityPoolConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetIdentityPoolConfigurationInput(v *SetIdentityPoolConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1209,10 +1166,8 @@ func (m *awsRestjson1_serializeOpSubscribeToDataset) H
 func (m *awsRestjson1_serializeOpSubscribeToDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1250,7 +1205,6 @@ func (m *awsRestjson1_serializeOpSubscribeToDataset) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSubscribeToDatasetInput(v *SubscribeToDatasetInput, encoder *httpbinding.Encoder) error {
@@ -1307,10 +1261,8 @@ func (m *awsRestjson1_serializeOpUnsubscribeFromDatase
 func (m *awsRestjson1_serializeOpUnsubscribeFromDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1348,7 +1300,6 @@ func (m *awsRestjson1_serializeOpUnsubscribeFromDatase
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUnsubscribeFromDatasetInput(v *UnsubscribeFromDatasetInput, encoder *httpbinding.Encoder) error {
@@ -1405,10 +1356,8 @@ func (m *awsRestjson1_serializeOpUpdateRecords) Handle
 func (m *awsRestjson1_serializeOpUpdateRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1457,7 +1406,6 @@ func (m *awsRestjson1_serializeOpUpdateRecords) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRecordsInput(v *UpdateRecordsInput, encoder *httpbinding.Encoder) error {
