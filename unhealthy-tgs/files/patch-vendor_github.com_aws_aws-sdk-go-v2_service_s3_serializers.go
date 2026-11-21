--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/serializers.go.orig	2026-11-21 14:46:24 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyxml "github.com/aws/smithy-go/encoding/xml"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 	"strconv"
@@ -29,10 +28,8 @@ func (m *awsRestxml_serializeOpAbortMultipartUpload) H
 func (m *awsRestxml_serializeOpAbortMultipartUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -70,7 +67,6 @@ func (m *awsRestxml_serializeOpAbortMultipartUpload) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsAbortMultipartUploadInput(v *AbortMultipartUploadInput, encoder *httpbinding.Encoder) error {
@@ -119,10 +115,8 @@ func (m *awsRestxml_serializeOpCompleteMultipartUpload
 func (m *awsRestxml_serializeOpCompleteMultipartUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -184,7 +178,6 @@ func (m *awsRestxml_serializeOpCompleteMultipartUpload
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCompleteMultipartUploadInput(v *CompleteMultipartUploadInput, encoder *httpbinding.Encoder) error {
@@ -313,10 +306,8 @@ func (m *awsRestxml_serializeOpCopyObject) HandleSeria
 func (m *awsRestxml_serializeOpCopyObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -354,7 +345,6 @@ func (m *awsRestxml_serializeOpCopyObject) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCopyObjectInput(v *CopyObjectInput, encoder *httpbinding.Encoder) error {
@@ -611,10 +601,8 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
 func (m *awsRestxml_serializeOpCreateBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -676,7 +664,6 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketInput(v *CreateBucketInput, encoder *httpbinding.Encoder) error {
@@ -742,10 +729,8 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataCon
 func (m *awsRestxml_serializeOpCreateBucketMetadataConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -807,7 +792,6 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketMetadataConfigurationInput(v *CreateBucketMetadataConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -843,10 +827,8 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataTab
 func (m *awsRestxml_serializeOpCreateBucketMetadataTableConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -908,7 +890,6 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataTab
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketMetadataTableConfigurationInput(v *CreateBucketMetadataTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -944,10 +925,8 @@ func (m *awsRestxml_serializeOpCreateMultipartUpload) 
 func (m *awsRestxml_serializeOpCreateMultipartUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -985,7 +964,6 @@ func (m *awsRestxml_serializeOpCreateMultipartUpload) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateMultipartUploadInput(v *CreateMultipartUploadInput, encoder *httpbinding.Encoder) error {
@@ -1177,10 +1155,8 @@ func (m *awsRestxml_serializeOpCreateSession) HandleSe
 func (m *awsRestxml_serializeOpCreateSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1218,7 +1194,6 @@ func (m *awsRestxml_serializeOpCreateSession) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateSessionInput(v *CreateSessionInput, encoder *httpbinding.Encoder) error {
@@ -1264,10 +1239,8 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
 func (m *awsRestxml_serializeOpDeleteBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1305,7 +1278,6 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketInput(v *DeleteBucketInput, encoder *httpbinding.Encoder) error {
@@ -1331,10 +1303,8 @@ func (m *awsRestxml_serializeOpDeleteBucketAnalyticsCo
 func (m *awsRestxml_serializeOpDeleteBucketAnalyticsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1372,7 +1342,6 @@ func (m *awsRestxml_serializeOpDeleteBucketAnalyticsCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketAnalyticsConfigurationInput(v *DeleteBucketAnalyticsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1402,10 +1371,8 @@ func (m *awsRestxml_serializeOpDeleteBucketCors) Handl
 func (m *awsRestxml_serializeOpDeleteBucketCors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1443,7 +1410,6 @@ func (m *awsRestxml_serializeOpDeleteBucketCors) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketCorsInput(v *DeleteBucketCorsInput, encoder *httpbinding.Encoder) error {
@@ -1469,10 +1435,8 @@ func (m *awsRestxml_serializeOpDeleteBucketEncryption)
 func (m *awsRestxml_serializeOpDeleteBucketEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1510,7 +1474,6 @@ func (m *awsRestxml_serializeOpDeleteBucketEncryption)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketEncryptionInput(v *DeleteBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -1536,10 +1499,8 @@ func (m *awsRestxml_serializeOpDeleteBucketIntelligent
 func (m *awsRestxml_serializeOpDeleteBucketIntelligentTieringConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1577,7 +1538,6 @@ func (m *awsRestxml_serializeOpDeleteBucketIntelligent
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketIntelligentTieringConfigurationInput(v *DeleteBucketIntelligentTieringConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1607,10 +1567,8 @@ func (m *awsRestxml_serializeOpDeleteBucketInventoryCo
 func (m *awsRestxml_serializeOpDeleteBucketInventoryConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1648,7 +1606,6 @@ func (m *awsRestxml_serializeOpDeleteBucketInventoryCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketInventoryConfigurationInput(v *DeleteBucketInventoryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1678,10 +1635,8 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycle) 
 func (m *awsRestxml_serializeOpDeleteBucketLifecycle) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1719,7 +1674,6 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycle) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketLifecycleInput(v *DeleteBucketLifecycleInput, encoder *httpbinding.Encoder) error {
@@ -1745,10 +1699,8 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataCon
 func (m *awsRestxml_serializeOpDeleteBucketMetadataConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1786,7 +1738,6 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketMetadataConfigurationInput(v *DeleteBucketMetadataConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1812,10 +1763,8 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataTab
 func (m *awsRestxml_serializeOpDeleteBucketMetadataTableConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1853,7 +1802,6 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataTab
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketMetadataTableConfigurationInput(v *DeleteBucketMetadataTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1879,10 +1827,8 @@ func (m *awsRestxml_serializeOpDeleteBucketMetricsConf
 func (m *awsRestxml_serializeOpDeleteBucketMetricsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1920,7 +1866,6 @@ func (m *awsRestxml_serializeOpDeleteBucketMetricsConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketMetricsConfigurationInput(v *DeleteBucketMetricsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1950,10 +1895,8 @@ func (m *awsRestxml_serializeOpDeleteBucketOwnershipCo
 func (m *awsRestxml_serializeOpDeleteBucketOwnershipControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1991,7 +1934,6 @@ func (m *awsRestxml_serializeOpDeleteBucketOwnershipCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketOwnershipControlsInput(v *DeleteBucketOwnershipControlsInput, encoder *httpbinding.Encoder) error {
@@ -2017,10 +1959,8 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
 func (m *awsRestxml_serializeOpDeleteBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2058,7 +1998,6 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketPolicyInput(v *DeleteBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2084,10 +2023,8 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
 func (m *awsRestxml_serializeOpDeleteBucketReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2125,7 +2062,6 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketReplicationInput(v *DeleteBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -2151,10 +2087,8 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
 func (m *awsRestxml_serializeOpDeleteBucketTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2192,7 +2126,6 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketTaggingInput(v *DeleteBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2218,10 +2151,8 @@ func (m *awsRestxml_serializeOpDeleteBucketWebsite) Ha
 func (m *awsRestxml_serializeOpDeleteBucketWebsite) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2259,7 +2190,6 @@ func (m *awsRestxml_serializeOpDeleteBucketWebsite) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketWebsiteInput(v *DeleteBucketWebsiteInput, encoder *httpbinding.Encoder) error {
@@ -2285,10 +2215,8 @@ func (m *awsRestxml_serializeOpDeleteObject) HandleSer
 func (m *awsRestxml_serializeOpDeleteObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2326,7 +2254,6 @@ func (m *awsRestxml_serializeOpDeleteObject) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectInput(v *DeleteObjectInput, encoder *httpbinding.Encoder) error {
@@ -2395,10 +2322,8 @@ func (m *awsRestxml_serializeOpDeleteObjectAnnotation)
 func (m *awsRestxml_serializeOpDeleteObjectAnnotation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2436,7 +2361,6 @@ func (m *awsRestxml_serializeOpDeleteObjectAnnotation)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectAnnotationInput(v *DeleteObjectAnnotationInput, encoder *httpbinding.Encoder) error {
@@ -2489,10 +2413,8 @@ func (m *awsRestxml_serializeOpDeleteObjects) HandleSe
 func (m *awsRestxml_serializeOpDeleteObjects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2554,7 +2476,6 @@ func (m *awsRestxml_serializeOpDeleteObjects) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectsInput(v *DeleteObjectsInput, encoder *httpbinding.Encoder) error {
@@ -2600,10 +2521,8 @@ func (m *awsRestxml_serializeOpDeleteObjectTagging) Ha
 func (m *awsRestxml_serializeOpDeleteObjectTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2641,7 +2560,6 @@ func (m *awsRestxml_serializeOpDeleteObjectTagging) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectTaggingInput(v *DeleteObjectTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2680,10 +2598,8 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
 func (m *awsRestxml_serializeOpDeletePublicAccessBlock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2721,7 +2637,6 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeletePublicAccessBlockInput(v *DeletePublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -2747,10 +2662,8 @@ func (m *awsRestxml_serializeOpGetBucketAbac) HandleSe
 func (m *awsRestxml_serializeOpGetBucketAbac) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2788,7 +2701,6 @@ func (m *awsRestxml_serializeOpGetBucketAbac) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAbacInput(v *GetBucketAbacInput, encoder *httpbinding.Encoder) error {
@@ -2814,10 +2726,8 @@ func (m *awsRestxml_serializeOpGetBucketAccelerateConf
 func (m *awsRestxml_serializeOpGetBucketAccelerateConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2855,7 +2765,6 @@ func (m *awsRestxml_serializeOpGetBucketAccelerateConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAccelerateConfigurationInput(v *GetBucketAccelerateConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2886,10 +2795,8 @@ func (m *awsRestxml_serializeOpGetBucketAcl) HandleSer
 func (m *awsRestxml_serializeOpGetBucketAcl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2927,7 +2834,6 @@ func (m *awsRestxml_serializeOpGetBucketAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAclInput(v *GetBucketAclInput, encoder *httpbinding.Encoder) error {
@@ -2953,10 +2859,8 @@ func (m *awsRestxml_serializeOpGetBucketAnalyticsConfi
 func (m *awsRestxml_serializeOpGetBucketAnalyticsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2994,7 +2898,6 @@ func (m *awsRestxml_serializeOpGetBucketAnalyticsConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAnalyticsConfigurationInput(v *GetBucketAnalyticsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3024,10 +2927,8 @@ func (m *awsRestxml_serializeOpGetBucketCors) HandleSe
 func (m *awsRestxml_serializeOpGetBucketCors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3065,7 +2966,6 @@ func (m *awsRestxml_serializeOpGetBucketCors) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketCorsInput(v *GetBucketCorsInput, encoder *httpbinding.Encoder) error {
@@ -3091,10 +2991,8 @@ func (m *awsRestxml_serializeOpGetBucketEncryption) Ha
 func (m *awsRestxml_serializeOpGetBucketEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3132,7 +3030,6 @@ func (m *awsRestxml_serializeOpGetBucketEncryption) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketEncryptionInput(v *GetBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -3158,10 +3055,8 @@ func (m *awsRestxml_serializeOpGetBucketIntelligentTie
 func (m *awsRestxml_serializeOpGetBucketIntelligentTieringConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3199,7 +3094,6 @@ func (m *awsRestxml_serializeOpGetBucketIntelligentTie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketIntelligentTieringConfigurationInput(v *GetBucketIntelligentTieringConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3229,10 +3123,8 @@ func (m *awsRestxml_serializeOpGetBucketInventoryConfi
 func (m *awsRestxml_serializeOpGetBucketInventoryConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3270,7 +3162,6 @@ func (m *awsRestxml_serializeOpGetBucketInventoryConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketInventoryConfigurationInput(v *GetBucketInventoryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3300,10 +3191,8 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
 func (m *awsRestxml_serializeOpGetBucketLifecycleConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3341,7 +3230,6 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLifecycleConfigurationInput(v *GetBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3367,10 +3255,8 @@ func (m *awsRestxml_serializeOpGetBucketLocation) Hand
 func (m *awsRestxml_serializeOpGetBucketLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3408,7 +3294,6 @@ func (m *awsRestxml_serializeOpGetBucketLocation) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLocationInput(v *GetBucketLocationInput, encoder *httpbinding.Encoder) error {
@@ -3434,10 +3319,8 @@ func (m *awsRestxml_serializeOpGetBucketLogging) Handl
 func (m *awsRestxml_serializeOpGetBucketLogging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3475,7 +3358,6 @@ func (m *awsRestxml_serializeOpGetBucketLogging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLoggingInput(v *GetBucketLoggingInput, encoder *httpbinding.Encoder) error {
@@ -3501,10 +3383,8 @@ func (m *awsRestxml_serializeOpGetBucketMetadataConfig
 func (m *awsRestxml_serializeOpGetBucketMetadataConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3542,7 +3422,6 @@ func (m *awsRestxml_serializeOpGetBucketMetadataConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketMetadataConfigurationInput(v *GetBucketMetadataConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3568,10 +3447,8 @@ func (m *awsRestxml_serializeOpGetBucketMetadataTableC
 func (m *awsRestxml_serializeOpGetBucketMetadataTableConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3609,7 +3486,6 @@ func (m *awsRestxml_serializeOpGetBucketMetadataTableC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketMetadataTableConfigurationInput(v *GetBucketMetadataTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3635,10 +3511,8 @@ func (m *awsRestxml_serializeOpGetBucketMetricsConfigu
 func (m *awsRestxml_serializeOpGetBucketMetricsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3676,7 +3550,6 @@ func (m *awsRestxml_serializeOpGetBucketMetricsConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketMetricsConfigurationInput(v *GetBucketMetricsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3706,10 +3579,8 @@ func (m *awsRestxml_serializeOpGetBucketNotificationCo
 func (m *awsRestxml_serializeOpGetBucketNotificationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3747,7 +3618,6 @@ func (m *awsRestxml_serializeOpGetBucketNotificationCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketNotificationConfigurationInput(v *GetBucketNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3773,10 +3643,8 @@ func (m *awsRestxml_serializeOpGetBucketOwnershipContr
 func (m *awsRestxml_serializeOpGetBucketOwnershipControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3814,7 +3682,6 @@ func (m *awsRestxml_serializeOpGetBucketOwnershipContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketOwnershipControlsInput(v *GetBucketOwnershipControlsInput, encoder *httpbinding.Encoder) error {
@@ -3840,10 +3707,8 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
 func (m *awsRestxml_serializeOpGetBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3881,7 +3746,6 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketPolicyInput(v *GetBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3907,10 +3771,8 @@ func (m *awsRestxml_serializeOpGetBucketPolicyStatus) 
 func (m *awsRestxml_serializeOpGetBucketPolicyStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3948,7 +3810,6 @@ func (m *awsRestxml_serializeOpGetBucketPolicyStatus) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketPolicyStatusInput(v *GetBucketPolicyStatusInput, encoder *httpbinding.Encoder) error {
@@ -3974,10 +3835,8 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
 func (m *awsRestxml_serializeOpGetBucketReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4015,7 +3874,6 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketReplicationInput(v *GetBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -4041,10 +3899,8 @@ func (m *awsRestxml_serializeOpGetBucketRequestPayment
 func (m *awsRestxml_serializeOpGetBucketRequestPayment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4082,7 +3938,6 @@ func (m *awsRestxml_serializeOpGetBucketRequestPayment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketRequestPaymentInput(v *GetBucketRequestPaymentInput, encoder *httpbinding.Encoder) error {
@@ -4108,10 +3963,8 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
 func (m *awsRestxml_serializeOpGetBucketTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4149,7 +4002,6 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketTaggingInput(v *GetBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -4175,10 +4027,8 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
 func (m *awsRestxml_serializeOpGetBucketVersioning) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4216,7 +4066,6 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketVersioningInput(v *GetBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -4242,10 +4091,8 @@ func (m *awsRestxml_serializeOpGetBucketWebsite) Handl
 func (m *awsRestxml_serializeOpGetBucketWebsite) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4283,7 +4130,6 @@ func (m *awsRestxml_serializeOpGetBucketWebsite) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketWebsiteInput(v *GetBucketWebsiteInput, encoder *httpbinding.Encoder) error {
@@ -4309,10 +4155,8 @@ func (m *awsRestxml_serializeOpGetObject) HandleSerial
 func (m *awsRestxml_serializeOpGetObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4350,7 +4194,6 @@ func (m *awsRestxml_serializeOpGetObject) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectInput(v *GetObjectInput, encoder *httpbinding.Encoder) error {
@@ -4467,10 +4310,8 @@ func (m *awsRestxml_serializeOpGetObjectAcl) HandleSer
 func (m *awsRestxml_serializeOpGetObjectAcl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4508,7 +4349,6 @@ func (m *awsRestxml_serializeOpGetObjectAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectAclInput(v *GetObjectAclInput, encoder *httpbinding.Encoder) error {
@@ -4552,10 +4392,8 @@ func (m *awsRestxml_serializeOpGetObjectAnnotation) Ha
 func (m *awsRestxml_serializeOpGetObjectAnnotation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4593,7 +4431,6 @@ func (m *awsRestxml_serializeOpGetObjectAnnotation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectAnnotationInput(v *GetObjectAnnotationInput, encoder *httpbinding.Encoder) error {
@@ -4646,10 +4483,8 @@ func (m *awsRestxml_serializeOpGetObjectAttributes) Ha
 func (m *awsRestxml_serializeOpGetObjectAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4687,7 +4522,6 @@ func (m *awsRestxml_serializeOpGetObjectAttributes) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectAttributesInput(v *GetObjectAttributesInput, encoder *httpbinding.Encoder) error {
@@ -4773,10 +4607,8 @@ func (m *awsRestxml_serializeOpGetObjectLegalHold) Han
 func (m *awsRestxml_serializeOpGetObjectLegalHold) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4814,7 +4646,6 @@ func (m *awsRestxml_serializeOpGetObjectLegalHold) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectLegalHoldInput(v *GetObjectLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -4858,10 +4689,8 @@ func (m *awsRestxml_serializeOpGetObjectLockConfigurat
 func (m *awsRestxml_serializeOpGetObjectLockConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4899,7 +4728,6 @@ func (m *awsRestxml_serializeOpGetObjectLockConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectLockConfigurationInput(v *GetObjectLockConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4925,10 +4753,8 @@ func (m *awsRestxml_serializeOpGetObjectRetention) Han
 func (m *awsRestxml_serializeOpGetObjectRetention) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4966,7 +4792,6 @@ func (m *awsRestxml_serializeOpGetObjectRetention) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectRetentionInput(v *GetObjectRetentionInput, encoder *httpbinding.Encoder) error {
@@ -5010,10 +4835,8 @@ func (m *awsRestxml_serializeOpGetObjectTagging) Handl
 func (m *awsRestxml_serializeOpGetObjectTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5051,7 +4874,6 @@ func (m *awsRestxml_serializeOpGetObjectTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectTaggingInput(v *GetObjectTaggingInput, encoder *httpbinding.Encoder) error {
@@ -5095,10 +4917,8 @@ func (m *awsRestxml_serializeOpGetObjectTorrent) Handl
 func (m *awsRestxml_serializeOpGetObjectTorrent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5136,7 +4956,6 @@ func (m *awsRestxml_serializeOpGetObjectTorrent) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectTorrentInput(v *GetObjectTorrentInput, encoder *httpbinding.Encoder) error {
@@ -5176,10 +4995,8 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
 func (m *awsRestxml_serializeOpGetPublicAccessBlock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5217,7 +5034,6 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicAccessBlockInput(v *GetPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -5243,10 +5059,8 @@ func (m *awsRestxml_serializeOpHeadBucket) HandleSeria
 func (m *awsRestxml_serializeOpHeadBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5284,7 +5098,6 @@ func (m *awsRestxml_serializeOpHeadBucket) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHeadBucketInput(v *HeadBucketInput, encoder *httpbinding.Encoder) error {
@@ -5310,10 +5123,8 @@ func (m *awsRestxml_serializeOpHeadObject) HandleSeria
 func (m *awsRestxml_serializeOpHeadObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5351,7 +5162,6 @@ func (m *awsRestxml_serializeOpHeadObject) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHeadObjectInput(v *HeadObjectInput, encoder *httpbinding.Encoder) error {
@@ -5468,10 +5278,8 @@ func (m *awsRestxml_serializeOpListBucketAnalyticsConf
 func (m *awsRestxml_serializeOpListBucketAnalyticsConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5509,7 +5317,6 @@ func (m *awsRestxml_serializeOpListBucketAnalyticsConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketAnalyticsConfigurationsInput(v *ListBucketAnalyticsConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5539,10 +5346,8 @@ func (m *awsRestxml_serializeOpListBucketIntelligentTi
 func (m *awsRestxml_serializeOpListBucketIntelligentTieringConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5580,7 +5385,6 @@ func (m *awsRestxml_serializeOpListBucketIntelligentTi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketIntelligentTieringConfigurationsInput(v *ListBucketIntelligentTieringConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5610,10 +5414,8 @@ func (m *awsRestxml_serializeOpListBucketInventoryConf
 func (m *awsRestxml_serializeOpListBucketInventoryConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5651,7 +5453,6 @@ func (m *awsRestxml_serializeOpListBucketInventoryConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketInventoryConfigurationsInput(v *ListBucketInventoryConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5681,10 +5482,8 @@ func (m *awsRestxml_serializeOpListBucketMetricsConfig
 func (m *awsRestxml_serializeOpListBucketMetricsConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5722,7 +5521,6 @@ func (m *awsRestxml_serializeOpListBucketMetricsConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketMetricsConfigurationsInput(v *ListBucketMetricsConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5752,10 +5550,8 @@ func (m *awsRestxml_serializeOpListBuckets) HandleSeri
 func (m *awsRestxml_serializeOpListBuckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5793,7 +5589,6 @@ func (m *awsRestxml_serializeOpListBuckets) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketsInput(v *ListBucketsInput, encoder *httpbinding.Encoder) error {
@@ -5830,10 +5625,8 @@ func (m *awsRestxml_serializeOpListDirectoryBuckets) H
 func (m *awsRestxml_serializeOpListDirectoryBuckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5871,7 +5664,6 @@ func (m *awsRestxml_serializeOpListDirectoryBuckets) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDirectoryBucketsInput(v *ListDirectoryBucketsInput, encoder *httpbinding.Encoder) error {
@@ -5900,10 +5692,8 @@ func (m *awsRestxml_serializeOpListMultipartUploads) H
 func (m *awsRestxml_serializeOpListMultipartUploads) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5941,7 +5731,6 @@ func (m *awsRestxml_serializeOpListMultipartUploads) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListMultipartUploadsInput(v *ListMultipartUploadsInput, encoder *httpbinding.Encoder) error {
@@ -5996,10 +5785,8 @@ func (m *awsRestxml_serializeOpListObjectAnnotations) 
 func (m *awsRestxml_serializeOpListObjectAnnotations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6037,7 +5824,6 @@ func (m *awsRestxml_serializeOpListObjectAnnotations) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectAnnotationsInput(v *ListObjectAnnotationsInput, encoder *httpbinding.Encoder) error {
@@ -6093,10 +5879,8 @@ func (m *awsRestxml_serializeOpListObjects) HandleSeri
 func (m *awsRestxml_serializeOpListObjects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6134,7 +5918,6 @@ func (m *awsRestxml_serializeOpListObjects) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectsInput(v *ListObjectsInput, encoder *httpbinding.Encoder) error {
@@ -6202,10 +5985,8 @@ func (m *awsRestxml_serializeOpListObjectsV2) HandleSe
 func (m *awsRestxml_serializeOpListObjectsV2) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6243,7 +6024,6 @@ func (m *awsRestxml_serializeOpListObjectsV2) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectsV2Input(v *ListObjectsV2Input, encoder *httpbinding.Encoder) error {
@@ -6319,10 +6099,8 @@ func (m *awsRestxml_serializeOpListObjectVersions) Han
 func (m *awsRestxml_serializeOpListObjectVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6360,7 +6138,6 @@ func (m *awsRestxml_serializeOpListObjectVersions) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectVersionsInput(v *ListObjectVersionsInput, encoder *httpbinding.Encoder) error {
@@ -6432,10 +6209,8 @@ func (m *awsRestxml_serializeOpListParts) HandleSerial
 func (m *awsRestxml_serializeOpListParts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6473,7 +6248,6 @@ func (m *awsRestxml_serializeOpListParts) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListPartsInput(v *ListPartsInput, encoder *httpbinding.Encoder) error {
@@ -6540,10 +6314,8 @@ func (m *awsRestxml_serializeOpPutBucketAbac) HandleSe
 func (m *awsRestxml_serializeOpPutBucketAbac) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6605,7 +6377,6 @@ func (m *awsRestxml_serializeOpPutBucketAbac) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAbacInput(v *PutBucketAbacInput, encoder *httpbinding.Encoder) error {
@@ -6641,10 +6412,8 @@ func (m *awsRestxml_serializeOpPutBucketAccelerateConf
 func (m *awsRestxml_serializeOpPutBucketAccelerateConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6706,7 +6475,6 @@ func (m *awsRestxml_serializeOpPutBucketAccelerateConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAccelerateConfigurationInput(v *PutBucketAccelerateConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6737,10 +6505,8 @@ func (m *awsRestxml_serializeOpPutBucketAcl) HandleSer
 func (m *awsRestxml_serializeOpPutBucketAcl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6802,7 +6568,6 @@ func (m *awsRestxml_serializeOpPutBucketAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAclInput(v *PutBucketAclInput, encoder *httpbinding.Encoder) error {
@@ -6868,10 +6633,8 @@ func (m *awsRestxml_serializeOpPutBucketAnalyticsConfi
 func (m *awsRestxml_serializeOpPutBucketAnalyticsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6933,7 +6696,6 @@ func (m *awsRestxml_serializeOpPutBucketAnalyticsConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAnalyticsConfigurationInput(v *PutBucketAnalyticsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6963,10 +6725,8 @@ func (m *awsRestxml_serializeOpPutBucketCors) HandleSe
 func (m *awsRestxml_serializeOpPutBucketCors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7028,7 +6788,6 @@ func (m *awsRestxml_serializeOpPutBucketCors) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketCorsInput(v *PutBucketCorsInput, encoder *httpbinding.Encoder) error {
@@ -7064,10 +6823,8 @@ func (m *awsRestxml_serializeOpPutBucketEncryption) Ha
 func (m *awsRestxml_serializeOpPutBucketEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7129,7 +6886,6 @@ func (m *awsRestxml_serializeOpPutBucketEncryption) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketEncryptionInput(v *PutBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -7165,10 +6921,8 @@ func (m *awsRestxml_serializeOpPutBucketIntelligentTie
 func (m *awsRestxml_serializeOpPutBucketIntelligentTieringConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7230,7 +6984,6 @@ func (m *awsRestxml_serializeOpPutBucketIntelligentTie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketIntelligentTieringConfigurationInput(v *PutBucketIntelligentTieringConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7260,10 +7013,8 @@ func (m *awsRestxml_serializeOpPutBucketInventoryConfi
 func (m *awsRestxml_serializeOpPutBucketInventoryConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7325,7 +7076,6 @@ func (m *awsRestxml_serializeOpPutBucketInventoryConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketInventoryConfigurationInput(v *PutBucketInventoryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7355,10 +7105,8 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
 func (m *awsRestxml_serializeOpPutBucketLifecycleConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7420,7 +7168,6 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketLifecycleConfigurationInput(v *PutBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7456,10 +7203,8 @@ func (m *awsRestxml_serializeOpPutBucketLogging) Handl
 func (m *awsRestxml_serializeOpPutBucketLogging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7521,7 +7266,6 @@ func (m *awsRestxml_serializeOpPutBucketLogging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketLoggingInput(v *PutBucketLoggingInput, encoder *httpbinding.Encoder) error {
@@ -7557,10 +7301,8 @@ func (m *awsRestxml_serializeOpPutBucketMetricsConfigu
 func (m *awsRestxml_serializeOpPutBucketMetricsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7622,7 +7364,6 @@ func (m *awsRestxml_serializeOpPutBucketMetricsConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketMetricsConfigurationInput(v *PutBucketMetricsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7652,10 +7393,8 @@ func (m *awsRestxml_serializeOpPutBucketNotificationCo
 func (m *awsRestxml_serializeOpPutBucketNotificationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7717,7 +7456,6 @@ func (m *awsRestxml_serializeOpPutBucketNotificationCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketNotificationConfigurationInput(v *PutBucketNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7748,10 +7486,8 @@ func (m *awsRestxml_serializeOpPutBucketOwnershipContr
 func (m *awsRestxml_serializeOpPutBucketOwnershipControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7813,7 +7549,6 @@ func (m *awsRestxml_serializeOpPutBucketOwnershipContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketOwnershipControlsInput(v *PutBucketOwnershipControlsInput, encoder *httpbinding.Encoder) error {
@@ -7849,10 +7584,8 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
 func (m *awsRestxml_serializeOpPutBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7902,7 +7635,6 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketPolicyInput(v *PutBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7943,10 +7675,8 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
 func (m *awsRestxml_serializeOpPutBucketReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8008,7 +7738,6 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketReplicationInput(v *PutBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -8049,10 +7778,8 @@ func (m *awsRestxml_serializeOpPutBucketRequestPayment
 func (m *awsRestxml_serializeOpPutBucketRequestPayment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8114,7 +7841,6 @@ func (m *awsRestxml_serializeOpPutBucketRequestPayment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketRequestPaymentInput(v *PutBucketRequestPaymentInput, encoder *httpbinding.Encoder) error {
@@ -8150,10 +7876,8 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
 func (m *awsRestxml_serializeOpPutBucketTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8215,7 +7939,6 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketTaggingInput(v *PutBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -8251,10 +7974,8 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
 func (m *awsRestxml_serializeOpPutBucketVersioning) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8316,7 +8037,6 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketVersioningInput(v *PutBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -8357,10 +8077,8 @@ func (m *awsRestxml_serializeOpPutBucketWebsite) Handl
 func (m *awsRestxml_serializeOpPutBucketWebsite) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8422,7 +8140,6 @@ func (m *awsRestxml_serializeOpPutBucketWebsite) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketWebsiteInput(v *PutBucketWebsiteInput, encoder *httpbinding.Encoder) error {
@@ -8458,10 +8175,8 @@ func (m *awsRestxml_serializeOpPutObject) HandleSerial
 func (m *awsRestxml_serializeOpPutObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8511,7 +8226,6 @@ func (m *awsRestxml_serializeOpPutObject) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectInput(v *PutObjectInput, encoder *httpbinding.Encoder) error {
@@ -8773,10 +8487,8 @@ func (m *awsRestxml_serializeOpPutObjectAcl) HandleSer
 func (m *awsRestxml_serializeOpPutObjectAcl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8838,7 +8550,6 @@ func (m *awsRestxml_serializeOpPutObjectAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectAclInput(v *PutObjectAclInput, encoder *httpbinding.Encoder) error {
@@ -8922,10 +8633,8 @@ func (m *awsRestxml_serializeOpPutObjectAnnotation) Ha
 func (m *awsRestxml_serializeOpPutObjectAnnotation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8975,7 +8684,6 @@ func (m *awsRestxml_serializeOpPutObjectAnnotation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectAnnotationInput(v *PutObjectAnnotationInput, encoder *httpbinding.Encoder) error {
@@ -9088,10 +8796,8 @@ func (m *awsRestxml_serializeOpPutObjectLegalHold) Han
 func (m *awsRestxml_serializeOpPutObjectLegalHold) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9153,7 +8859,6 @@ func (m *awsRestxml_serializeOpPutObjectLegalHold) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectLegalHoldInput(v *PutObjectLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -9207,10 +8912,8 @@ func (m *awsRestxml_serializeOpPutObjectLockConfigurat
 func (m *awsRestxml_serializeOpPutObjectLockConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9272,7 +8975,6 @@ func (m *awsRestxml_serializeOpPutObjectLockConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectLockConfigurationInput(v *PutObjectLockConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -9318,10 +9020,8 @@ func (m *awsRestxml_serializeOpPutObjectRetention) Han
 func (m *awsRestxml_serializeOpPutObjectRetention) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9383,7 +9083,6 @@ func (m *awsRestxml_serializeOpPutObjectRetention) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectRetentionInput(v *PutObjectRetentionInput, encoder *httpbinding.Encoder) error {
@@ -9442,10 +9141,8 @@ func (m *awsRestxml_serializeOpPutObjectTagging) Handl
 func (m *awsRestxml_serializeOpPutObjectTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9507,7 +9204,6 @@ func (m *awsRestxml_serializeOpPutObjectTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectTaggingInput(v *PutObjectTaggingInput, encoder *httpbinding.Encoder) error {
@@ -9561,10 +9257,8 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
 func (m *awsRestxml_serializeOpPutPublicAccessBlock) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9626,7 +9320,6 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutPublicAccessBlockInput(v *PutPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -9662,10 +9355,8 @@ func (m *awsRestxml_serializeOpRenameObject) HandleSer
 func (m *awsRestxml_serializeOpRenameObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9703,7 +9394,6 @@ func (m *awsRestxml_serializeOpRenameObject) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsRenameObjectInput(v *RenameObjectInput, encoder *httpbinding.Encoder) error {
@@ -9783,10 +9473,8 @@ func (m *awsRestxml_serializeOpRestoreObject) HandleSe
 func (m *awsRestxml_serializeOpRestoreObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9848,7 +9536,6 @@ func (m *awsRestxml_serializeOpRestoreObject) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsRestoreObjectInput(v *RestoreObjectInput, encoder *httpbinding.Encoder) error {
@@ -9897,10 +9584,8 @@ func (m *awsRestxml_serializeOpSelectObjectContent) Ha
 func (m *awsRestxml_serializeOpSelectObjectContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9956,7 +9641,6 @@ func (m *awsRestxml_serializeOpSelectObjectContent) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsSelectObjectContentInput(v *SelectObjectContentInput, encoder *httpbinding.Encoder) error {
@@ -10085,10 +9769,8 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataAnn
 func (m *awsRestxml_serializeOpUpdateBucketMetadataAnnotationTableConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10150,7 +9832,6 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataAnn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateBucketMetadataAnnotationTableConfigurationInput(v *UpdateBucketMetadataAnnotationTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -10186,10 +9867,8 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataInv
 func (m *awsRestxml_serializeOpUpdateBucketMetadataInventoryTableConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10251,7 +9930,6 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataInv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateBucketMetadataInventoryTableConfigurationInput(v *UpdateBucketMetadataInventoryTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -10287,10 +9965,8 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataJou
 func (m *awsRestxml_serializeOpUpdateBucketMetadataJournalTableConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10352,7 +10028,6 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataJou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateBucketMetadataJournalTableConfigurationInput(v *UpdateBucketMetadataJournalTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -10388,10 +10063,8 @@ func (m *awsRestxml_serializeOpUpdateObjectEncryption)
 func (m *awsRestxml_serializeOpUpdateObjectEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10453,7 +10126,6 @@ func (m *awsRestxml_serializeOpUpdateObjectEncryption)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateObjectEncryptionInput(v *UpdateObjectEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -10507,10 +10179,8 @@ func (m *awsRestxml_serializeOpUploadPart) HandleSeria
 func (m *awsRestxml_serializeOpUploadPart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10560,7 +10230,6 @@ func (m *awsRestxml_serializeOpUploadPart) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUploadPartInput(v *UploadPartInput, encoder *httpbinding.Encoder) error {
@@ -10688,10 +10357,8 @@ func (m *awsRestxml_serializeOpUploadPartCopy) HandleS
 func (m *awsRestxml_serializeOpUploadPartCopy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10729,7 +10396,6 @@ func (m *awsRestxml_serializeOpUploadPartCopy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUploadPartCopyInput(v *UploadPartCopyInput, encoder *httpbinding.Encoder) error {
@@ -10842,10 +10508,8 @@ func (m *awsRestxml_serializeOpWriteGetObjectResponse)
 func (m *awsRestxml_serializeOpWriteGetObjectResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10895,7 +10559,6 @@ func (m *awsRestxml_serializeOpWriteGetObjectResponse)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsWriteGetObjectResponseInput(v *WriteGetObjectResponseInput, encoder *httpbinding.Encoder) error {
