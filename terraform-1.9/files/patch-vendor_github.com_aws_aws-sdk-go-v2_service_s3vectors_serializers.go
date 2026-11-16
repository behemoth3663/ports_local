--- vendor/github.com/aws/aws-sdk-go-v2/service/s3vectors/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3vectors/serializers.go
@@ -13,7 +13,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -28,10 +27,8 @@ func (m *awsRestjson1_serializeOpCreateIndex) HandleSe
 func (m *awsRestjson1_serializeOpCreateIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -76,7 +73,6 @@ func (m *awsRestjson1_serializeOpCreateIndex) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIndexInput(v *CreateIndexInput, encoder *httpbinding.Encoder) error {
@@ -141,10 +137,8 @@ func (m *awsRestjson1_serializeOpCreateVectorBucket) H
 func (m *awsRestjson1_serializeOpCreateVectorBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -189,7 +183,6 @@ func (m *awsRestjson1_serializeOpCreateVectorBucket) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVectorBucketInput(v *CreateVectorBucketInput, encoder *httpbinding.Encoder) error {
@@ -229,10 +222,8 @@ func (m *awsRestjson1_serializeOpDeleteIndex) HandleSe
 func (m *awsRestjson1_serializeOpDeleteIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -277,7 +268,6 @@ func (m *awsRestjson1_serializeOpDeleteIndex) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIndexInput(v *DeleteIndexInput, encoder *httpbinding.Encoder) error {
@@ -320,10 +310,8 @@ func (m *awsRestjson1_serializeOpDeleteVectorBucket) H
 func (m *awsRestjson1_serializeOpDeleteVectorBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -368,7 +356,6 @@ func (m *awsRestjson1_serializeOpDeleteVectorBucket) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVectorBucketInput(v *DeleteVectorBucketInput, encoder *httpbinding.Encoder) error {
@@ -406,10 +393,8 @@ func (m *awsRestjson1_serializeOpDeleteVectorBucketPol
 func (m *awsRestjson1_serializeOpDeleteVectorBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -454,7 +439,6 @@ func (m *awsRestjson1_serializeOpDeleteVectorBucketPol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVectorBucketPolicyInput(v *DeleteVectorBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -492,10 +476,8 @@ func (m *awsRestjson1_serializeOpDeleteVectors) Handle
 func (m *awsRestjson1_serializeOpDeleteVectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -540,7 +522,6 @@ func (m *awsRestjson1_serializeOpDeleteVectors) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVectorsInput(v *DeleteVectorsInput, encoder *httpbinding.Encoder) error {
@@ -590,10 +571,8 @@ func (m *awsRestjson1_serializeOpGetIndex) HandleSeria
 func (m *awsRestjson1_serializeOpGetIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -638,7 +617,6 @@ func (m *awsRestjson1_serializeOpGetIndex) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIndexInput(v *GetIndexInput, encoder *httpbinding.Encoder) error {
@@ -681,10 +659,8 @@ func (m *awsRestjson1_serializeOpGetVectorBucket) Hand
 func (m *awsRestjson1_serializeOpGetVectorBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -729,7 +705,6 @@ func (m *awsRestjson1_serializeOpGetVectorBucket) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVectorBucketInput(v *GetVectorBucketInput, encoder *httpbinding.Encoder) error {
@@ -767,10 +742,8 @@ func (m *awsRestjson1_serializeOpGetVectorBucketPolicy
 func (m *awsRestjson1_serializeOpGetVectorBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -815,7 +788,6 @@ func (m *awsRestjson1_serializeOpGetVectorBucketPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVectorBucketPolicyInput(v *GetVectorBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -853,10 +825,8 @@ func (m *awsRestjson1_serializeOpGetVectors) HandleSer
 func (m *awsRestjson1_serializeOpGetVectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -901,7 +871,6 @@ func (m *awsRestjson1_serializeOpGetVectors) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVectorsInput(v *GetVectorsInput, encoder *httpbinding.Encoder) error {
@@ -961,10 +930,8 @@ func (m *awsRestjson1_serializeOpListIndexes) HandleSe
 func (m *awsRestjson1_serializeOpListIndexes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1009,7 +976,6 @@ func (m *awsRestjson1_serializeOpListIndexes) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIndexesInput(v *ListIndexesInput, encoder *httpbinding.Encoder) error {
@@ -1062,10 +1028,8 @@ func (m *awsRestjson1_serializeOpListVectorBuckets) Ha
 func (m *awsRestjson1_serializeOpListVectorBuckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1110,7 +1074,6 @@ func (m *awsRestjson1_serializeOpListVectorBuckets) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVectorBucketsInput(v *ListVectorBucketsInput, encoder *httpbinding.Encoder) error {
@@ -1153,10 +1116,8 @@ func (m *awsRestjson1_serializeOpListVectors) HandleSe
 func (m *awsRestjson1_serializeOpListVectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1201,7 +1162,6 @@ func (m *awsRestjson1_serializeOpListVectors) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVectorsInput(v *ListVectorsInput, encoder *httpbinding.Encoder) error {
@@ -1274,10 +1234,8 @@ func (m *awsRestjson1_serializeOpPutVectorBucketPolicy
 func (m *awsRestjson1_serializeOpPutVectorBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1322,7 +1280,6 @@ func (m *awsRestjson1_serializeOpPutVectorBucketPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVectorBucketPolicyInput(v *PutVectorBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1365,10 +1322,8 @@ func (m *awsRestjson1_serializeOpPutVectors) HandleSer
 func (m *awsRestjson1_serializeOpPutVectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1413,7 +1368,6 @@ func (m *awsRestjson1_serializeOpPutVectors) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutVectorsInput(v *PutVectorsInput, encoder *httpbinding.Encoder) error {
@@ -1463,10 +1417,8 @@ func (m *awsRestjson1_serializeOpQueryVectors) HandleS
 func (m *awsRestjson1_serializeOpQueryVectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1511,7 +1463,6 @@ func (m *awsRestjson1_serializeOpQueryVectors) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsQueryVectorsInput(v *QueryVectorsInput, encoder *httpbinding.Encoder) error {
