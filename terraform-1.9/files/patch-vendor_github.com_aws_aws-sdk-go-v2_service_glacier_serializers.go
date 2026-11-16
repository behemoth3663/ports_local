--- vendor/github.com/aws/aws-sdk-go-v2/service/glacier/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/glacier/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAbortMultipartUpload)
 func (m *awsRestjson1_serializeOpAbortMultipartUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpAbortMultipartUpload)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAbortMultipartUploadInput(v *AbortMultipartUploadInput, encoder *httpbinding.Encoder) error {
@@ -114,10 +110,8 @@ func (m *awsRestjson1_serializeOpAbortVaultLock) Handl
 func (m *awsRestjson1_serializeOpAbortVaultLock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -155,7 +149,6 @@ func (m *awsRestjson1_serializeOpAbortVaultLock) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAbortVaultLockInput(v *AbortVaultLockInput, encoder *httpbinding.Encoder) error {
@@ -194,10 +187,8 @@ func (m *awsRestjson1_serializeOpAddTagsToVault) Handl
 func (m *awsRestjson1_serializeOpAddTagsToVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -246,7 +237,6 @@ func (m *awsRestjson1_serializeOpAddTagsToVault) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddTagsToVaultInput(v *AddTagsToVaultInput, encoder *httpbinding.Encoder) error {
@@ -299,10 +289,8 @@ func (m *awsRestjson1_serializeOpCompleteMultipartUplo
 func (m *awsRestjson1_serializeOpCompleteMultipartUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -340,7 +328,6 @@ func (m *awsRestjson1_serializeOpCompleteMultipartUplo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCompleteMultipartUploadInput(v *CompleteMultipartUploadInput, encoder *httpbinding.Encoder) error {
@@ -398,10 +385,8 @@ func (m *awsRestjson1_serializeOpCompleteVaultLock) Ha
 func (m *awsRestjson1_serializeOpCompleteVaultLock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -439,7 +424,6 @@ func (m *awsRestjson1_serializeOpCompleteVaultLock) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCompleteVaultLockInput(v *CompleteVaultLockInput, encoder *httpbinding.Encoder) error {
@@ -487,10 +471,8 @@ func (m *awsRestjson1_serializeOpCreateVault) HandleSe
 func (m *awsRestjson1_serializeOpCreateVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -528,7 +510,6 @@ func (m *awsRestjson1_serializeOpCreateVault) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVaultInput(v *CreateVaultInput, encoder *httpbinding.Encoder) error {
@@ -567,10 +548,8 @@ func (m *awsRestjson1_serializeOpDeleteArchive) Handle
 func (m *awsRestjson1_serializeOpDeleteArchive) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -608,7 +587,6 @@ func (m *awsRestjson1_serializeOpDeleteArchive) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteArchiveInput(v *DeleteArchiveInput, encoder *httpbinding.Encoder) error {
@@ -656,10 +634,8 @@ func (m *awsRestjson1_serializeOpDeleteVault) HandleSe
 func (m *awsRestjson1_serializeOpDeleteVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -697,7 +673,6 @@ func (m *awsRestjson1_serializeOpDeleteVault) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVaultInput(v *DeleteVaultInput, encoder *httpbinding.Encoder) error {
@@ -736,10 +711,8 @@ func (m *awsRestjson1_serializeOpDeleteVaultAccessPoli
 func (m *awsRestjson1_serializeOpDeleteVaultAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -777,7 +750,6 @@ func (m *awsRestjson1_serializeOpDeleteVaultAccessPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVaultAccessPolicyInput(v *DeleteVaultAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -816,10 +788,8 @@ func (m *awsRestjson1_serializeOpDeleteVaultNotificati
 func (m *awsRestjson1_serializeOpDeleteVaultNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -857,7 +827,6 @@ func (m *awsRestjson1_serializeOpDeleteVaultNotificati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVaultNotificationsInput(v *DeleteVaultNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -896,10 +865,8 @@ func (m *awsRestjson1_serializeOpDescribeJob) HandleSe
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
@@ -937,7 +904,6 @@ func (m *awsRestjson1_serializeOpDescribeJob) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeJobInput(v *DescribeJobInput, encoder *httpbinding.Encoder) error {
@@ -985,10 +951,8 @@ func (m *awsRestjson1_serializeOpDescribeVault) Handle
 func (m *awsRestjson1_serializeOpDescribeVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1026,7 +990,6 @@ func (m *awsRestjson1_serializeOpDescribeVault) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeVaultInput(v *DescribeVaultInput, encoder *httpbinding.Encoder) error {
@@ -1065,10 +1028,8 @@ func (m *awsRestjson1_serializeOpGetDataRetrievalPolic
 func (m *awsRestjson1_serializeOpGetDataRetrievalPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1106,7 +1067,6 @@ func (m *awsRestjson1_serializeOpGetDataRetrievalPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataRetrievalPolicyInput(v *GetDataRetrievalPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1136,10 +1096,8 @@ func (m *awsRestjson1_serializeOpGetJobOutput) HandleS
 func (m *awsRestjson1_serializeOpGetJobOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1177,7 +1135,6 @@ func (m *awsRestjson1_serializeOpGetJobOutput) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJobOutputInput(v *GetJobOutputInput, encoder *httpbinding.Encoder) error {
@@ -1230,10 +1187,8 @@ func (m *awsRestjson1_serializeOpGetVaultAccessPolicy)
 func (m *awsRestjson1_serializeOpGetVaultAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1271,7 +1226,6 @@ func (m *awsRestjson1_serializeOpGetVaultAccessPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVaultAccessPolicyInput(v *GetVaultAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1310,10 +1264,8 @@ func (m *awsRestjson1_serializeOpGetVaultLock) HandleS
 func (m *awsRestjson1_serializeOpGetVaultLock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1351,7 +1303,6 @@ func (m *awsRestjson1_serializeOpGetVaultLock) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVaultLockInput(v *GetVaultLockInput, encoder *httpbinding.Encoder) error {
@@ -1390,10 +1341,8 @@ func (m *awsRestjson1_serializeOpGetVaultNotifications
 func (m *awsRestjson1_serializeOpGetVaultNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1431,7 +1380,6 @@ func (m *awsRestjson1_serializeOpGetVaultNotifications
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVaultNotificationsInput(v *GetVaultNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -1470,10 +1418,8 @@ func (m *awsRestjson1_serializeOpInitiateJob) HandleSe
 func (m *awsRestjson1_serializeOpInitiateJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1535,7 +1481,6 @@ func (m *awsRestjson1_serializeOpInitiateJob) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInitiateJobInput(v *InitiateJobInput, encoder *httpbinding.Encoder) error {
@@ -1574,10 +1519,8 @@ func (m *awsRestjson1_serializeOpInitiateMultipartUplo
 func (m *awsRestjson1_serializeOpInitiateMultipartUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1615,7 +1558,6 @@ func (m *awsRestjson1_serializeOpInitiateMultipartUplo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInitiateMultipartUploadInput(v *InitiateMultipartUploadInput, encoder *httpbinding.Encoder) error {
@@ -1664,10 +1606,8 @@ func (m *awsRestjson1_serializeOpInitiateVaultLock) Ha
 func (m *awsRestjson1_serializeOpInitiateVaultLock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1729,7 +1669,6 @@ func (m *awsRestjson1_serializeOpInitiateVaultLock) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInitiateVaultLockInput(v *InitiateVaultLockInput, encoder *httpbinding.Encoder) error {
@@ -1768,10 +1707,8 @@ func (m *awsRestjson1_serializeOpListJobs) HandleSeria
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
@@ -1809,7 +1746,6 @@ func (m *awsRestjson1_serializeOpListJobs) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobsInput(v *ListJobsInput, encoder *httpbinding.Encoder) error {
@@ -1864,10 +1800,8 @@ func (m *awsRestjson1_serializeOpListMultipartUploads)
 func (m *awsRestjson1_serializeOpListMultipartUploads) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1905,7 +1839,6 @@ func (m *awsRestjson1_serializeOpListMultipartUploads)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMultipartUploadsInput(v *ListMultipartUploadsInput, encoder *httpbinding.Encoder) error {
@@ -1952,10 +1885,8 @@ func (m *awsRestjson1_serializeOpListParts) HandleSeri
 func (m *awsRestjson1_serializeOpListParts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1993,7 +1924,6 @@ func (m *awsRestjson1_serializeOpListParts) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPartsInput(v *ListPartsInput, encoder *httpbinding.Encoder) error {
@@ -2049,10 +1979,8 @@ func (m *awsRestjson1_serializeOpListProvisionedCapaci
 func (m *awsRestjson1_serializeOpListProvisionedCapacity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2090,7 +2018,6 @@ func (m *awsRestjson1_serializeOpListProvisionedCapaci
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProvisionedCapacityInput(v *ListProvisionedCapacityInput, encoder *httpbinding.Encoder) error {
@@ -2120,10 +2047,8 @@ func (m *awsRestjson1_serializeOpListTagsForVault) Han
 func (m *awsRestjson1_serializeOpListTagsForVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2161,7 +2086,6 @@ func (m *awsRestjson1_serializeOpListTagsForVault) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForVaultInput(v *ListTagsForVaultInput, encoder *httpbinding.Encoder) error {
@@ -2200,10 +2124,8 @@ func (m *awsRestjson1_serializeOpListVaults) HandleSer
 func (m *awsRestjson1_serializeOpListVaults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2241,7 +2163,6 @@ func (m *awsRestjson1_serializeOpListVaults) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVaultsInput(v *ListVaultsInput, encoder *httpbinding.Encoder) error {
@@ -2279,10 +2200,8 @@ func (m *awsRestjson1_serializeOpPurchaseProvisionedCa
 func (m *awsRestjson1_serializeOpPurchaseProvisionedCapacity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2320,7 +2239,6 @@ func (m *awsRestjson1_serializeOpPurchaseProvisionedCa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPurchaseProvisionedCapacityInput(v *PurchaseProvisionedCapacityInput, encoder *httpbinding.Encoder) error {
@@ -2350,10 +2268,8 @@ func (m *awsRestjson1_serializeOpRemoveTagsFromVault) 
 func (m *awsRestjson1_serializeOpRemoveTagsFromVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2402,7 +2318,6 @@ func (m *awsRestjson1_serializeOpRemoveTagsFromVault) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveTagsFromVaultInput(v *RemoveTagsFromVaultInput, encoder *httpbinding.Encoder) error {
@@ -2455,10 +2370,8 @@ func (m *awsRestjson1_serializeOpSetDataRetrievalPolic
 func (m *awsRestjson1_serializeOpSetDataRetrievalPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2507,7 +2420,6 @@ func (m *awsRestjson1_serializeOpSetDataRetrievalPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetDataRetrievalPolicyInput(v *SetDataRetrievalPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2551,10 +2463,8 @@ func (m *awsRestjson1_serializeOpSetVaultAccessPolicy)
 func (m *awsRestjson1_serializeOpSetVaultAccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2616,7 +2526,6 @@ func (m *awsRestjson1_serializeOpSetVaultAccessPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetVaultAccessPolicyInput(v *SetVaultAccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2655,10 +2564,8 @@ func (m *awsRestjson1_serializeOpSetVaultNotifications
 func (m *awsRestjson1_serializeOpSetVaultNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2720,7 +2627,6 @@ func (m *awsRestjson1_serializeOpSetVaultNotifications
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetVaultNotificationsInput(v *SetVaultNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -2759,10 +2665,8 @@ func (m *awsRestjson1_serializeOpUploadArchive) Handle
 func (m *awsRestjson1_serializeOpUploadArchive) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2812,7 +2716,6 @@ func (m *awsRestjson1_serializeOpUploadArchive) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUploadArchiveInput(v *UploadArchiveInput, encoder *httpbinding.Encoder) error {
@@ -2861,10 +2764,8 @@ func (m *awsRestjson1_serializeOpUploadMultipartPart) 
 func (m *awsRestjson1_serializeOpUploadMultipartPart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2914,7 +2815,6 @@ func (m *awsRestjson1_serializeOpUploadMultipartPart) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUploadMultipartPartInput(v *UploadMultipartPartInput, encoder *httpbinding.Encoder) error {
