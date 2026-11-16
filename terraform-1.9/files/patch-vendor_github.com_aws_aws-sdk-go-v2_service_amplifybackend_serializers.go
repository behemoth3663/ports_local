--- vendor/github.com/aws/aws-sdk-go-v2/service/amplifybackend/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/amplifybackend/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCloneBackend) HandleS
 func (m *awsRestjson1_serializeOpCloneBackend) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -78,7 +75,6 @@ func (m *awsRestjson1_serializeOpCloneBackend) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCloneBackendInput(v *CloneBackendInput, encoder *httpbinding.Encoder) error {
@@ -129,10 +125,8 @@ func (m *awsRestjson1_serializeOpCreateBackend) Handle
 func (m *awsRestjson1_serializeOpCreateBackend) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -177,7 +171,6 @@ func (m *awsRestjson1_serializeOpCreateBackend) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackendInput(v *CreateBackendInput, encoder *httpbinding.Encoder) error {
@@ -232,10 +225,8 @@ func (m *awsRestjson1_serializeOpCreateBackendAPI) Han
 func (m *awsRestjson1_serializeOpCreateBackendAPI) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -284,7 +275,6 @@ func (m *awsRestjson1_serializeOpCreateBackendAPI) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackendAPIInput(v *CreateBackendAPIInput, encoder *httpbinding.Encoder) error {
@@ -338,10 +328,8 @@ func (m *awsRestjson1_serializeOpCreateBackendAuth) Ha
 func (m *awsRestjson1_serializeOpCreateBackendAuth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -390,7 +378,6 @@ func (m *awsRestjson1_serializeOpCreateBackendAuth) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackendAuthInput(v *CreateBackendAuthInput, encoder *httpbinding.Encoder) error {
@@ -444,10 +431,8 @@ func (m *awsRestjson1_serializeOpCreateBackendConfig) 
 func (m *awsRestjson1_serializeOpCreateBackendConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -496,7 +481,6 @@ func (m *awsRestjson1_serializeOpCreateBackendConfig) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackendConfigInput(v *CreateBackendConfigInput, encoder *httpbinding.Encoder) error {
@@ -538,10 +522,8 @@ func (m *awsRestjson1_serializeOpCreateBackendStorage)
 func (m *awsRestjson1_serializeOpCreateBackendStorage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -590,7 +572,6 @@ func (m *awsRestjson1_serializeOpCreateBackendStorage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBackendStorageInput(v *CreateBackendStorageInput, encoder *httpbinding.Encoder) error {
@@ -644,10 +625,8 @@ func (m *awsRestjson1_serializeOpCreateToken) HandleSe
 func (m *awsRestjson1_serializeOpCreateToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -685,7 +664,6 @@ func (m *awsRestjson1_serializeOpCreateToken) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTokenInput(v *CreateTokenInput, encoder *httpbinding.Encoder) error {
@@ -715,10 +693,8 @@ func (m *awsRestjson1_serializeOpDeleteBackend) Handle
 func (m *awsRestjson1_serializeOpDeleteBackend) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -756,7 +732,6 @@ func (m *awsRestjson1_serializeOpDeleteBackend) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackendInput(v *DeleteBackendInput, encoder *httpbinding.Encoder) error {
@@ -795,10 +770,8 @@ func (m *awsRestjson1_serializeOpDeleteBackendAPI) Han
 func (m *awsRestjson1_serializeOpDeleteBackendAPI) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -847,7 +820,6 @@ func (m *awsRestjson1_serializeOpDeleteBackendAPI) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackendAPIInput(v *DeleteBackendAPIInput, encoder *httpbinding.Encoder) error {
@@ -905,10 +877,8 @@ func (m *awsRestjson1_serializeOpDeleteBackendAuth) Ha
 func (m *awsRestjson1_serializeOpDeleteBackendAuth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -957,7 +927,6 @@ func (m *awsRestjson1_serializeOpDeleteBackendAuth) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackendAuthInput(v *DeleteBackendAuthInput, encoder *httpbinding.Encoder) error {
@@ -1008,10 +977,8 @@ func (m *awsRestjson1_serializeOpDeleteBackendStorage)
 func (m *awsRestjson1_serializeOpDeleteBackendStorage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1060,7 +1027,6 @@ func (m *awsRestjson1_serializeOpDeleteBackendStorage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBackendStorageInput(v *DeleteBackendStorageInput, encoder *httpbinding.Encoder) error {
@@ -1116,10 +1082,8 @@ func (m *awsRestjson1_serializeOpDeleteToken) HandleSe
 func (m *awsRestjson1_serializeOpDeleteToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1157,7 +1121,6 @@ func (m *awsRestjson1_serializeOpDeleteToken) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTokenInput(v *DeleteTokenInput, encoder *httpbinding.Encoder) error {
@@ -1196,10 +1159,8 @@ func (m *awsRestjson1_serializeOpGenerateBackendAPIMod
 func (m *awsRestjson1_serializeOpGenerateBackendAPIModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1248,7 +1209,6 @@ func (m *awsRestjson1_serializeOpGenerateBackendAPIMod
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateBackendAPIModelsInput(v *GenerateBackendAPIModelsInput, encoder *httpbinding.Encoder) error {
@@ -1299,10 +1259,8 @@ func (m *awsRestjson1_serializeOpGetBackend) HandleSer
 func (m *awsRestjson1_serializeOpGetBackend) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1351,7 +1309,6 @@ func (m *awsRestjson1_serializeOpGetBackend) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackendInput(v *GetBackendInput, encoder *httpbinding.Encoder) error {
@@ -1393,10 +1350,8 @@ func (m *awsRestjson1_serializeOpGetBackendAPI) Handle
 func (m *awsRestjson1_serializeOpGetBackendAPI) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1445,7 +1400,6 @@ func (m *awsRestjson1_serializeOpGetBackendAPI) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackendAPIInput(v *GetBackendAPIInput, encoder *httpbinding.Encoder) error {
@@ -1503,10 +1457,8 @@ func (m *awsRestjson1_serializeOpGetBackendAPIModels) 
 func (m *awsRestjson1_serializeOpGetBackendAPIModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1555,7 +1507,6 @@ func (m *awsRestjson1_serializeOpGetBackendAPIModels) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackendAPIModelsInput(v *GetBackendAPIModelsInput, encoder *httpbinding.Encoder) error {
@@ -1606,10 +1557,8 @@ func (m *awsRestjson1_serializeOpGetBackendAuth) Handl
 func (m *awsRestjson1_serializeOpGetBackendAuth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1658,7 +1607,6 @@ func (m *awsRestjson1_serializeOpGetBackendAuth) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackendAuthInput(v *GetBackendAuthInput, encoder *httpbinding.Encoder) error {
@@ -1709,10 +1657,8 @@ func (m *awsRestjson1_serializeOpGetBackendJob) Handle
 func (m *awsRestjson1_serializeOpGetBackendJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1750,7 +1696,6 @@ func (m *awsRestjson1_serializeOpGetBackendJob) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackendJobInput(v *GetBackendJobInput, encoder *httpbinding.Encoder) error {
@@ -1798,10 +1743,8 @@ func (m *awsRestjson1_serializeOpGetBackendStorage) Ha
 func (m *awsRestjson1_serializeOpGetBackendStorage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1850,7 +1793,6 @@ func (m *awsRestjson1_serializeOpGetBackendStorage) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBackendStorageInput(v *GetBackendStorageInput, encoder *httpbinding.Encoder) error {
@@ -1901,10 +1843,8 @@ func (m *awsRestjson1_serializeOpGetToken) HandleSeria
 func (m *awsRestjson1_serializeOpGetToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1942,7 +1882,6 @@ func (m *awsRestjson1_serializeOpGetToken) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTokenInput(v *GetTokenInput, encoder *httpbinding.Encoder) error {
@@ -1981,10 +1920,8 @@ func (m *awsRestjson1_serializeOpImportBackendAuth) Ha
 func (m *awsRestjson1_serializeOpImportBackendAuth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2033,7 +1970,6 @@ func (m *awsRestjson1_serializeOpImportBackendAuth) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportBackendAuthInput(v *ImportBackendAuthInput, encoder *httpbinding.Encoder) error {
@@ -2099,10 +2035,8 @@ func (m *awsRestjson1_serializeOpImportBackendStorage)
 func (m *awsRestjson1_serializeOpImportBackendStorage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2151,7 +2085,6 @@ func (m *awsRestjson1_serializeOpImportBackendStorage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportBackendStorageInput(v *ImportBackendStorageInput, encoder *httpbinding.Encoder) error {
@@ -2207,10 +2140,8 @@ func (m *awsRestjson1_serializeOpListBackendJobs) Hand
 func (m *awsRestjson1_serializeOpListBackendJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2259,7 +2190,6 @@ func (m *awsRestjson1_serializeOpListBackendJobs) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBackendJobsInput(v *ListBackendJobsInput, encoder *httpbinding.Encoder) error {
@@ -2330,10 +2260,8 @@ func (m *awsRestjson1_serializeOpListS3Buckets) Handle
 func (m *awsRestjson1_serializeOpListS3Buckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2378,7 +2306,6 @@ func (m *awsRestjson1_serializeOpListS3Buckets) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListS3BucketsInput(v *ListS3BucketsInput, encoder *httpbinding.Encoder) error {
@@ -2411,10 +2338,8 @@ func (m *awsRestjson1_serializeOpRemoveAllBackends) Ha
 func (m *awsRestjson1_serializeOpRemoveAllBackends) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2463,7 +2388,6 @@ func (m *awsRestjson1_serializeOpRemoveAllBackends) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveAllBackendsInput(v *RemoveAllBackendsInput, encoder *httpbinding.Encoder) error {
@@ -2505,10 +2429,8 @@ func (m *awsRestjson1_serializeOpRemoveBackendConfig) 
 func (m *awsRestjson1_serializeOpRemoveBackendConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2546,7 +2468,6 @@ func (m *awsRestjson1_serializeOpRemoveBackendConfig) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveBackendConfigInput(v *RemoveBackendConfigInput, encoder *httpbinding.Encoder) error {
@@ -2576,10 +2497,8 @@ func (m *awsRestjson1_serializeOpUpdateBackendAPI) Han
 func (m *awsRestjson1_serializeOpUpdateBackendAPI) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2628,7 +2547,6 @@ func (m *awsRestjson1_serializeOpUpdateBackendAPI) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBackendAPIInput(v *UpdateBackendAPIInput, encoder *httpbinding.Encoder) error {
@@ -2686,10 +2604,8 @@ func (m *awsRestjson1_serializeOpUpdateBackendAuth) Ha
 func (m *awsRestjson1_serializeOpUpdateBackendAuth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2738,7 +2654,6 @@ func (m *awsRestjson1_serializeOpUpdateBackendAuth) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBackendAuthInput(v *UpdateBackendAuthInput, encoder *httpbinding.Encoder) error {
@@ -2796,10 +2711,8 @@ func (m *awsRestjson1_serializeOpUpdateBackendConfig) 
 func (m *awsRestjson1_serializeOpUpdateBackendConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2848,7 +2761,6 @@ func (m *awsRestjson1_serializeOpUpdateBackendConfig) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBackendConfigInput(v *UpdateBackendConfigInput, encoder *httpbinding.Encoder) error {
@@ -2892,10 +2804,8 @@ func (m *awsRestjson1_serializeOpUpdateBackendJob) Han
 func (m *awsRestjson1_serializeOpUpdateBackendJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2944,7 +2854,6 @@ func (m *awsRestjson1_serializeOpUpdateBackendJob) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBackendJobInput(v *UpdateBackendJobInput, encoder *httpbinding.Encoder) error {
@@ -3009,10 +2918,8 @@ func (m *awsRestjson1_serializeOpUpdateBackendStorage)
 func (m *awsRestjson1_serializeOpUpdateBackendStorage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3061,7 +2968,6 @@ func (m *awsRestjson1_serializeOpUpdateBackendStorage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBackendStorageInput(v *UpdateBackendStorageInput, encoder *httpbinding.Encoder) error {
