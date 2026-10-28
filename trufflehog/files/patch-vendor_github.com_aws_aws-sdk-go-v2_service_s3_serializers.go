--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/serializers.go.orig	2026-10-28 15:34:01 UTC
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
@@ -596,10 +586,8 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
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
@@ -661,7 +649,6 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketInput(v *CreateBucketInput, encoder *httpbinding.Encoder) error {
@@ -727,10 +714,8 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataCon
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
@@ -792,7 +777,6 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketMetadataConfigurationInput(v *CreateBucketMetadataConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -828,10 +812,8 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataTab
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
@@ -893,7 +875,6 @@ func (m *awsRestxml_serializeOpCreateBucketMetadataTab
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketMetadataTableConfigurationInput(v *CreateBucketMetadataTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -929,10 +910,8 @@ func (m *awsRestxml_serializeOpCreateMultipartUpload) 
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
@@ -970,7 +949,6 @@ func (m *awsRestxml_serializeOpCreateMultipartUpload) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateMultipartUploadInput(v *CreateMultipartUploadInput, encoder *httpbinding.Encoder) error {
@@ -1147,10 +1125,8 @@ func (m *awsRestxml_serializeOpCreateSession) HandleSe
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
@@ -1188,7 +1164,6 @@ func (m *awsRestxml_serializeOpCreateSession) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateSessionInput(v *CreateSessionInput, encoder *httpbinding.Encoder) error {
@@ -1234,10 +1209,8 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
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
@@ -1275,7 +1248,6 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketInput(v *DeleteBucketInput, encoder *httpbinding.Encoder) error {
@@ -1301,10 +1273,8 @@ func (m *awsRestxml_serializeOpDeleteBucketAnalyticsCo
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
@@ -1342,7 +1312,6 @@ func (m *awsRestxml_serializeOpDeleteBucketAnalyticsCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketAnalyticsConfigurationInput(v *DeleteBucketAnalyticsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1372,10 +1341,8 @@ func (m *awsRestxml_serializeOpDeleteBucketCors) Handl
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
@@ -1413,7 +1380,6 @@ func (m *awsRestxml_serializeOpDeleteBucketCors) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketCorsInput(v *DeleteBucketCorsInput, encoder *httpbinding.Encoder) error {
@@ -1439,10 +1405,8 @@ func (m *awsRestxml_serializeOpDeleteBucketEncryption)
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
@@ -1480,7 +1444,6 @@ func (m *awsRestxml_serializeOpDeleteBucketEncryption)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketEncryptionInput(v *DeleteBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -1506,10 +1469,8 @@ func (m *awsRestxml_serializeOpDeleteBucketIntelligent
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
@@ -1547,7 +1508,6 @@ func (m *awsRestxml_serializeOpDeleteBucketIntelligent
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketIntelligentTieringConfigurationInput(v *DeleteBucketIntelligentTieringConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1577,10 +1537,8 @@ func (m *awsRestxml_serializeOpDeleteBucketInventoryCo
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
@@ -1618,7 +1576,6 @@ func (m *awsRestxml_serializeOpDeleteBucketInventoryCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketInventoryConfigurationInput(v *DeleteBucketInventoryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1648,10 +1605,8 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycle) 
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
@@ -1689,7 +1644,6 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycle) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketLifecycleInput(v *DeleteBucketLifecycleInput, encoder *httpbinding.Encoder) error {
@@ -1715,10 +1669,8 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataCon
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
@@ -1756,7 +1708,6 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketMetadataConfigurationInput(v *DeleteBucketMetadataConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1782,10 +1733,8 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataTab
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
@@ -1823,7 +1772,6 @@ func (m *awsRestxml_serializeOpDeleteBucketMetadataTab
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketMetadataTableConfigurationInput(v *DeleteBucketMetadataTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1849,10 +1797,8 @@ func (m *awsRestxml_serializeOpDeleteBucketMetricsConf
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
@@ -1890,7 +1836,6 @@ func (m *awsRestxml_serializeOpDeleteBucketMetricsConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketMetricsConfigurationInput(v *DeleteBucketMetricsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1920,10 +1865,8 @@ func (m *awsRestxml_serializeOpDeleteBucketOwnershipCo
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
@@ -1961,7 +1904,6 @@ func (m *awsRestxml_serializeOpDeleteBucketOwnershipCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketOwnershipControlsInput(v *DeleteBucketOwnershipControlsInput, encoder *httpbinding.Encoder) error {
@@ -1987,10 +1929,8 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
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
@@ -2028,7 +1968,6 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketPolicyInput(v *DeleteBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2054,10 +1993,8 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
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
@@ -2095,7 +2032,6 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketReplicationInput(v *DeleteBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -2121,10 +2057,8 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
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
@@ -2162,7 +2096,6 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketTaggingInput(v *DeleteBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2188,10 +2121,8 @@ func (m *awsRestxml_serializeOpDeleteBucketWebsite) Ha
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
@@ -2229,7 +2160,6 @@ func (m *awsRestxml_serializeOpDeleteBucketWebsite) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketWebsiteInput(v *DeleteBucketWebsiteInput, encoder *httpbinding.Encoder) error {
@@ -2255,10 +2185,8 @@ func (m *awsRestxml_serializeOpDeleteObject) HandleSer
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
@@ -2296,7 +2224,6 @@ func (m *awsRestxml_serializeOpDeleteObject) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectInput(v *DeleteObjectInput, encoder *httpbinding.Encoder) error {
@@ -2365,10 +2292,8 @@ func (m *awsRestxml_serializeOpDeleteObjectAnnotation)
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
@@ -2406,7 +2331,6 @@ func (m *awsRestxml_serializeOpDeleteObjectAnnotation)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectAnnotationInput(v *DeleteObjectAnnotationInput, encoder *httpbinding.Encoder) error {
@@ -2459,10 +2383,8 @@ func (m *awsRestxml_serializeOpDeleteObjects) HandleSe
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
@@ -2524,7 +2446,6 @@ func (m *awsRestxml_serializeOpDeleteObjects) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectsInput(v *DeleteObjectsInput, encoder *httpbinding.Encoder) error {
@@ -2570,10 +2491,8 @@ func (m *awsRestxml_serializeOpDeleteObjectTagging) Ha
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
@@ -2611,7 +2530,6 @@ func (m *awsRestxml_serializeOpDeleteObjectTagging) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteObjectTaggingInput(v *DeleteObjectTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2650,10 +2568,8 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
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
@@ -2691,7 +2607,6 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeletePublicAccessBlockInput(v *DeletePublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -2717,10 +2632,8 @@ func (m *awsRestxml_serializeOpGetBucketAbac) HandleSe
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
@@ -2758,7 +2671,6 @@ func (m *awsRestxml_serializeOpGetBucketAbac) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAbacInput(v *GetBucketAbacInput, encoder *httpbinding.Encoder) error {
@@ -2784,10 +2696,8 @@ func (m *awsRestxml_serializeOpGetBucketAccelerateConf
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
@@ -2825,7 +2735,6 @@ func (m *awsRestxml_serializeOpGetBucketAccelerateConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAccelerateConfigurationInput(v *GetBucketAccelerateConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2856,10 +2765,8 @@ func (m *awsRestxml_serializeOpGetBucketAcl) HandleSer
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
@@ -2897,7 +2804,6 @@ func (m *awsRestxml_serializeOpGetBucketAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAclInput(v *GetBucketAclInput, encoder *httpbinding.Encoder) error {
@@ -2923,10 +2829,8 @@ func (m *awsRestxml_serializeOpGetBucketAnalyticsConfi
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
@@ -2964,7 +2868,6 @@ func (m *awsRestxml_serializeOpGetBucketAnalyticsConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketAnalyticsConfigurationInput(v *GetBucketAnalyticsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2994,10 +2897,8 @@ func (m *awsRestxml_serializeOpGetBucketCors) HandleSe
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
@@ -3035,7 +2936,6 @@ func (m *awsRestxml_serializeOpGetBucketCors) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketCorsInput(v *GetBucketCorsInput, encoder *httpbinding.Encoder) error {
@@ -3061,10 +2961,8 @@ func (m *awsRestxml_serializeOpGetBucketEncryption) Ha
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
@@ -3102,7 +3000,6 @@ func (m *awsRestxml_serializeOpGetBucketEncryption) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketEncryptionInput(v *GetBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -3128,10 +3025,8 @@ func (m *awsRestxml_serializeOpGetBucketIntelligentTie
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
@@ -3169,7 +3064,6 @@ func (m *awsRestxml_serializeOpGetBucketIntelligentTie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketIntelligentTieringConfigurationInput(v *GetBucketIntelligentTieringConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3199,10 +3093,8 @@ func (m *awsRestxml_serializeOpGetBucketInventoryConfi
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
@@ -3240,7 +3132,6 @@ func (m *awsRestxml_serializeOpGetBucketInventoryConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketInventoryConfigurationInput(v *GetBucketInventoryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3270,10 +3161,8 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
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
@@ -3311,7 +3200,6 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLifecycleConfigurationInput(v *GetBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3337,10 +3225,8 @@ func (m *awsRestxml_serializeOpGetBucketLocation) Hand
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
@@ -3378,7 +3264,6 @@ func (m *awsRestxml_serializeOpGetBucketLocation) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLocationInput(v *GetBucketLocationInput, encoder *httpbinding.Encoder) error {
@@ -3404,10 +3289,8 @@ func (m *awsRestxml_serializeOpGetBucketLogging) Handl
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
@@ -3445,7 +3328,6 @@ func (m *awsRestxml_serializeOpGetBucketLogging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLoggingInput(v *GetBucketLoggingInput, encoder *httpbinding.Encoder) error {
@@ -3471,10 +3353,8 @@ func (m *awsRestxml_serializeOpGetBucketMetadataConfig
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
@@ -3512,7 +3392,6 @@ func (m *awsRestxml_serializeOpGetBucketMetadataConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketMetadataConfigurationInput(v *GetBucketMetadataConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3538,10 +3417,8 @@ func (m *awsRestxml_serializeOpGetBucketMetadataTableC
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
@@ -3579,7 +3456,6 @@ func (m *awsRestxml_serializeOpGetBucketMetadataTableC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketMetadataTableConfigurationInput(v *GetBucketMetadataTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3605,10 +3481,8 @@ func (m *awsRestxml_serializeOpGetBucketMetricsConfigu
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
@@ -3646,7 +3520,6 @@ func (m *awsRestxml_serializeOpGetBucketMetricsConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketMetricsConfigurationInput(v *GetBucketMetricsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3676,10 +3549,8 @@ func (m *awsRestxml_serializeOpGetBucketNotificationCo
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
@@ -3717,7 +3588,6 @@ func (m *awsRestxml_serializeOpGetBucketNotificationCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketNotificationConfigurationInput(v *GetBucketNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3743,10 +3613,8 @@ func (m *awsRestxml_serializeOpGetBucketOwnershipContr
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
@@ -3784,7 +3652,6 @@ func (m *awsRestxml_serializeOpGetBucketOwnershipContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketOwnershipControlsInput(v *GetBucketOwnershipControlsInput, encoder *httpbinding.Encoder) error {
@@ -3810,10 +3677,8 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
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
@@ -3851,7 +3716,6 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketPolicyInput(v *GetBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3877,10 +3741,8 @@ func (m *awsRestxml_serializeOpGetBucketPolicyStatus) 
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
@@ -3918,7 +3780,6 @@ func (m *awsRestxml_serializeOpGetBucketPolicyStatus) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketPolicyStatusInput(v *GetBucketPolicyStatusInput, encoder *httpbinding.Encoder) error {
@@ -3944,10 +3805,8 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
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
@@ -3985,7 +3844,6 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketReplicationInput(v *GetBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -4011,10 +3869,8 @@ func (m *awsRestxml_serializeOpGetBucketRequestPayment
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
@@ -4052,7 +3908,6 @@ func (m *awsRestxml_serializeOpGetBucketRequestPayment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketRequestPaymentInput(v *GetBucketRequestPaymentInput, encoder *httpbinding.Encoder) error {
@@ -4078,10 +3933,8 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
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
@@ -4119,7 +3972,6 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketTaggingInput(v *GetBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -4145,10 +3997,8 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
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
@@ -4186,7 +4036,6 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketVersioningInput(v *GetBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -4212,10 +4061,8 @@ func (m *awsRestxml_serializeOpGetBucketWebsite) Handl
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
@@ -4253,7 +4100,6 @@ func (m *awsRestxml_serializeOpGetBucketWebsite) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketWebsiteInput(v *GetBucketWebsiteInput, encoder *httpbinding.Encoder) error {
@@ -4279,10 +4125,8 @@ func (m *awsRestxml_serializeOpGetObject) HandleSerial
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
@@ -4320,7 +4164,6 @@ func (m *awsRestxml_serializeOpGetObject) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectInput(v *GetObjectInput, encoder *httpbinding.Encoder) error {
@@ -4437,10 +4280,8 @@ func (m *awsRestxml_serializeOpGetObjectAcl) HandleSer
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
@@ -4478,7 +4319,6 @@ func (m *awsRestxml_serializeOpGetObjectAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectAclInput(v *GetObjectAclInput, encoder *httpbinding.Encoder) error {
@@ -4522,10 +4362,8 @@ func (m *awsRestxml_serializeOpGetObjectAnnotation) Ha
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
@@ -4563,7 +4401,6 @@ func (m *awsRestxml_serializeOpGetObjectAnnotation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectAnnotationInput(v *GetObjectAnnotationInput, encoder *httpbinding.Encoder) error {
@@ -4616,10 +4453,8 @@ func (m *awsRestxml_serializeOpGetObjectAttributes) Ha
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
@@ -4657,7 +4492,6 @@ func (m *awsRestxml_serializeOpGetObjectAttributes) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectAttributesInput(v *GetObjectAttributesInput, encoder *httpbinding.Encoder) error {
@@ -4743,10 +4577,8 @@ func (m *awsRestxml_serializeOpGetObjectLegalHold) Han
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
@@ -4784,7 +4616,6 @@ func (m *awsRestxml_serializeOpGetObjectLegalHold) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectLegalHoldInput(v *GetObjectLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -4828,10 +4659,8 @@ func (m *awsRestxml_serializeOpGetObjectLockConfigurat
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
@@ -4869,7 +4698,6 @@ func (m *awsRestxml_serializeOpGetObjectLockConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectLockConfigurationInput(v *GetObjectLockConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4895,10 +4723,8 @@ func (m *awsRestxml_serializeOpGetObjectRetention) Han
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
@@ -4936,7 +4762,6 @@ func (m *awsRestxml_serializeOpGetObjectRetention) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectRetentionInput(v *GetObjectRetentionInput, encoder *httpbinding.Encoder) error {
@@ -4980,10 +4805,8 @@ func (m *awsRestxml_serializeOpGetObjectTagging) Handl
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
@@ -5021,7 +4844,6 @@ func (m *awsRestxml_serializeOpGetObjectTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectTaggingInput(v *GetObjectTaggingInput, encoder *httpbinding.Encoder) error {
@@ -5065,10 +4887,8 @@ func (m *awsRestxml_serializeOpGetObjectTorrent) Handl
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
@@ -5106,7 +4926,6 @@ func (m *awsRestxml_serializeOpGetObjectTorrent) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetObjectTorrentInput(v *GetObjectTorrentInput, encoder *httpbinding.Encoder) error {
@@ -5146,10 +4965,8 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
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
@@ -5187,7 +5004,6 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicAccessBlockInput(v *GetPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -5213,10 +5029,8 @@ func (m *awsRestxml_serializeOpHeadBucket) HandleSeria
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
@@ -5254,7 +5068,6 @@ func (m *awsRestxml_serializeOpHeadBucket) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHeadBucketInput(v *HeadBucketInput, encoder *httpbinding.Encoder) error {
@@ -5280,10 +5093,8 @@ func (m *awsRestxml_serializeOpHeadObject) HandleSeria
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
@@ -5321,7 +5132,6 @@ func (m *awsRestxml_serializeOpHeadObject) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsHeadObjectInput(v *HeadObjectInput, encoder *httpbinding.Encoder) error {
@@ -5438,10 +5248,8 @@ func (m *awsRestxml_serializeOpListBucketAnalyticsConf
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
@@ -5479,7 +5287,6 @@ func (m *awsRestxml_serializeOpListBucketAnalyticsConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketAnalyticsConfigurationsInput(v *ListBucketAnalyticsConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5509,10 +5316,8 @@ func (m *awsRestxml_serializeOpListBucketIntelligentTi
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
@@ -5550,7 +5355,6 @@ func (m *awsRestxml_serializeOpListBucketIntelligentTi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketIntelligentTieringConfigurationsInput(v *ListBucketIntelligentTieringConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5580,10 +5384,8 @@ func (m *awsRestxml_serializeOpListBucketInventoryConf
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
@@ -5621,7 +5423,6 @@ func (m *awsRestxml_serializeOpListBucketInventoryConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketInventoryConfigurationsInput(v *ListBucketInventoryConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5651,10 +5452,8 @@ func (m *awsRestxml_serializeOpListBucketMetricsConfig
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
@@ -5692,7 +5491,6 @@ func (m *awsRestxml_serializeOpListBucketMetricsConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketMetricsConfigurationsInput(v *ListBucketMetricsConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5722,10 +5520,8 @@ func (m *awsRestxml_serializeOpListBuckets) HandleSeri
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
@@ -5763,7 +5559,6 @@ func (m *awsRestxml_serializeOpListBuckets) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListBucketsInput(v *ListBucketsInput, encoder *httpbinding.Encoder) error {
@@ -5800,10 +5595,8 @@ func (m *awsRestxml_serializeOpListDirectoryBuckets) H
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
@@ -5841,7 +5634,6 @@ func (m *awsRestxml_serializeOpListDirectoryBuckets) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDirectoryBucketsInput(v *ListDirectoryBucketsInput, encoder *httpbinding.Encoder) error {
@@ -5870,10 +5662,8 @@ func (m *awsRestxml_serializeOpListMultipartUploads) H
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
@@ -5911,7 +5701,6 @@ func (m *awsRestxml_serializeOpListMultipartUploads) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListMultipartUploadsInput(v *ListMultipartUploadsInput, encoder *httpbinding.Encoder) error {
@@ -5966,10 +5755,8 @@ func (m *awsRestxml_serializeOpListObjectAnnotations) 
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
@@ -6007,7 +5794,6 @@ func (m *awsRestxml_serializeOpListObjectAnnotations) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectAnnotationsInput(v *ListObjectAnnotationsInput, encoder *httpbinding.Encoder) error {
@@ -6063,10 +5849,8 @@ func (m *awsRestxml_serializeOpListObjects) HandleSeri
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
@@ -6104,7 +5888,6 @@ func (m *awsRestxml_serializeOpListObjects) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectsInput(v *ListObjectsInput, encoder *httpbinding.Encoder) error {
@@ -6172,10 +5955,8 @@ func (m *awsRestxml_serializeOpListObjectsV2) HandleSe
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
@@ -6213,7 +5994,6 @@ func (m *awsRestxml_serializeOpListObjectsV2) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectsV2Input(v *ListObjectsV2Input, encoder *httpbinding.Encoder) error {
@@ -6289,10 +6069,8 @@ func (m *awsRestxml_serializeOpListObjectVersions) Han
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
@@ -6330,7 +6108,6 @@ func (m *awsRestxml_serializeOpListObjectVersions) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListObjectVersionsInput(v *ListObjectVersionsInput, encoder *httpbinding.Encoder) error {
@@ -6402,10 +6179,8 @@ func (m *awsRestxml_serializeOpListParts) HandleSerial
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
@@ -6443,7 +6218,6 @@ func (m *awsRestxml_serializeOpListParts) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListPartsInput(v *ListPartsInput, encoder *httpbinding.Encoder) error {
@@ -6510,10 +6284,8 @@ func (m *awsRestxml_serializeOpPutBucketAbac) HandleSe
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
@@ -6575,7 +6347,6 @@ func (m *awsRestxml_serializeOpPutBucketAbac) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAbacInput(v *PutBucketAbacInput, encoder *httpbinding.Encoder) error {
@@ -6611,10 +6382,8 @@ func (m *awsRestxml_serializeOpPutBucketAccelerateConf
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
@@ -6676,7 +6445,6 @@ func (m *awsRestxml_serializeOpPutBucketAccelerateConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAccelerateConfigurationInput(v *PutBucketAccelerateConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6707,10 +6475,8 @@ func (m *awsRestxml_serializeOpPutBucketAcl) HandleSer
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
@@ -6772,7 +6538,6 @@ func (m *awsRestxml_serializeOpPutBucketAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAclInput(v *PutBucketAclInput, encoder *httpbinding.Encoder) error {
@@ -6838,10 +6603,8 @@ func (m *awsRestxml_serializeOpPutBucketAnalyticsConfi
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
@@ -6903,7 +6666,6 @@ func (m *awsRestxml_serializeOpPutBucketAnalyticsConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketAnalyticsConfigurationInput(v *PutBucketAnalyticsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6933,10 +6695,8 @@ func (m *awsRestxml_serializeOpPutBucketCors) HandleSe
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
@@ -6998,7 +6758,6 @@ func (m *awsRestxml_serializeOpPutBucketCors) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketCorsInput(v *PutBucketCorsInput, encoder *httpbinding.Encoder) error {
@@ -7034,10 +6793,8 @@ func (m *awsRestxml_serializeOpPutBucketEncryption) Ha
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
@@ -7099,7 +6856,6 @@ func (m *awsRestxml_serializeOpPutBucketEncryption) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketEncryptionInput(v *PutBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -7135,10 +6891,8 @@ func (m *awsRestxml_serializeOpPutBucketIntelligentTie
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
@@ -7200,7 +6954,6 @@ func (m *awsRestxml_serializeOpPutBucketIntelligentTie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketIntelligentTieringConfigurationInput(v *PutBucketIntelligentTieringConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7230,10 +6983,8 @@ func (m *awsRestxml_serializeOpPutBucketInventoryConfi
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
@@ -7295,7 +7046,6 @@ func (m *awsRestxml_serializeOpPutBucketInventoryConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketInventoryConfigurationInput(v *PutBucketInventoryConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7325,10 +7075,8 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
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
@@ -7390,7 +7138,6 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketLifecycleConfigurationInput(v *PutBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7426,10 +7173,8 @@ func (m *awsRestxml_serializeOpPutBucketLogging) Handl
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
@@ -7491,7 +7236,6 @@ func (m *awsRestxml_serializeOpPutBucketLogging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketLoggingInput(v *PutBucketLoggingInput, encoder *httpbinding.Encoder) error {
@@ -7527,10 +7271,8 @@ func (m *awsRestxml_serializeOpPutBucketMetricsConfigu
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
@@ -7592,7 +7334,6 @@ func (m *awsRestxml_serializeOpPutBucketMetricsConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketMetricsConfigurationInput(v *PutBucketMetricsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7622,10 +7363,8 @@ func (m *awsRestxml_serializeOpPutBucketNotificationCo
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
@@ -7687,7 +7426,6 @@ func (m *awsRestxml_serializeOpPutBucketNotificationCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketNotificationConfigurationInput(v *PutBucketNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7718,10 +7456,8 @@ func (m *awsRestxml_serializeOpPutBucketOwnershipContr
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
@@ -7783,7 +7519,6 @@ func (m *awsRestxml_serializeOpPutBucketOwnershipContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketOwnershipControlsInput(v *PutBucketOwnershipControlsInput, encoder *httpbinding.Encoder) error {
@@ -7819,10 +7554,8 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
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
@@ -7872,7 +7605,6 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketPolicyInput(v *PutBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7913,10 +7645,8 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
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
@@ -7978,7 +7708,6 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketReplicationInput(v *PutBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -8019,10 +7748,8 @@ func (m *awsRestxml_serializeOpPutBucketRequestPayment
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
@@ -8084,7 +7811,6 @@ func (m *awsRestxml_serializeOpPutBucketRequestPayment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketRequestPaymentInput(v *PutBucketRequestPaymentInput, encoder *httpbinding.Encoder) error {
@@ -8120,10 +7846,8 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
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
@@ -8185,7 +7909,6 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketTaggingInput(v *PutBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -8221,10 +7944,8 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
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
@@ -8286,7 +8007,6 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketVersioningInput(v *PutBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -8327,10 +8047,8 @@ func (m *awsRestxml_serializeOpPutBucketWebsite) Handl
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
@@ -8392,7 +8110,6 @@ func (m *awsRestxml_serializeOpPutBucketWebsite) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketWebsiteInput(v *PutBucketWebsiteInput, encoder *httpbinding.Encoder) error {
@@ -8428,10 +8145,8 @@ func (m *awsRestxml_serializeOpPutObject) HandleSerial
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
@@ -8481,7 +8196,6 @@ func (m *awsRestxml_serializeOpPutObject) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectInput(v *PutObjectInput, encoder *httpbinding.Encoder) error {
@@ -8728,10 +8442,8 @@ func (m *awsRestxml_serializeOpPutObjectAcl) HandleSer
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
@@ -8793,7 +8505,6 @@ func (m *awsRestxml_serializeOpPutObjectAcl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectAclInput(v *PutObjectAclInput, encoder *httpbinding.Encoder) error {
@@ -8877,10 +8588,8 @@ func (m *awsRestxml_serializeOpPutObjectAnnotation) Ha
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
@@ -8930,7 +8639,6 @@ func (m *awsRestxml_serializeOpPutObjectAnnotation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectAnnotationInput(v *PutObjectAnnotationInput, encoder *httpbinding.Encoder) error {
@@ -9043,10 +8751,8 @@ func (m *awsRestxml_serializeOpPutObjectLegalHold) Han
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
@@ -9108,7 +8814,6 @@ func (m *awsRestxml_serializeOpPutObjectLegalHold) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectLegalHoldInput(v *PutObjectLegalHoldInput, encoder *httpbinding.Encoder) error {
@@ -9162,10 +8867,8 @@ func (m *awsRestxml_serializeOpPutObjectLockConfigurat
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
@@ -9227,7 +8930,6 @@ func (m *awsRestxml_serializeOpPutObjectLockConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectLockConfigurationInput(v *PutObjectLockConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -9273,10 +8975,8 @@ func (m *awsRestxml_serializeOpPutObjectRetention) Han
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
@@ -9338,7 +9038,6 @@ func (m *awsRestxml_serializeOpPutObjectRetention) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectRetentionInput(v *PutObjectRetentionInput, encoder *httpbinding.Encoder) error {
@@ -9397,10 +9096,8 @@ func (m *awsRestxml_serializeOpPutObjectTagging) Handl
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
@@ -9462,7 +9159,6 @@ func (m *awsRestxml_serializeOpPutObjectTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutObjectTaggingInput(v *PutObjectTaggingInput, encoder *httpbinding.Encoder) error {
@@ -9516,10 +9212,8 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
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
@@ -9581,7 +9275,6 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutPublicAccessBlockInput(v *PutPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -9617,10 +9310,8 @@ func (m *awsRestxml_serializeOpRenameObject) HandleSer
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
@@ -9658,7 +9349,6 @@ func (m *awsRestxml_serializeOpRenameObject) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsRenameObjectInput(v *RenameObjectInput, encoder *httpbinding.Encoder) error {
@@ -9738,10 +9428,8 @@ func (m *awsRestxml_serializeOpRestoreObject) HandleSe
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
@@ -9803,7 +9491,6 @@ func (m *awsRestxml_serializeOpRestoreObject) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsRestoreObjectInput(v *RestoreObjectInput, encoder *httpbinding.Encoder) error {
@@ -9852,10 +9539,8 @@ func (m *awsRestxml_serializeOpSelectObjectContent) Ha
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
@@ -9911,7 +9596,6 @@ func (m *awsRestxml_serializeOpSelectObjectContent) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsSelectObjectContentInput(v *SelectObjectContentInput, encoder *httpbinding.Encoder) error {
@@ -10040,10 +9724,8 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataAnn
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
@@ -10105,7 +9787,6 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataAnn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateBucketMetadataAnnotationTableConfigurationInput(v *UpdateBucketMetadataAnnotationTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -10141,10 +9822,8 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataInv
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
@@ -10206,7 +9885,6 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataInv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateBucketMetadataInventoryTableConfigurationInput(v *UpdateBucketMetadataInventoryTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -10242,10 +9920,8 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataJou
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
@@ -10307,7 +9983,6 @@ func (m *awsRestxml_serializeOpUpdateBucketMetadataJou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateBucketMetadataJournalTableConfigurationInput(v *UpdateBucketMetadataJournalTableConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -10343,10 +10018,8 @@ func (m *awsRestxml_serializeOpUpdateObjectEncryption)
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
@@ -10408,7 +10081,6 @@ func (m *awsRestxml_serializeOpUpdateObjectEncryption)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateObjectEncryptionInput(v *UpdateObjectEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -10462,10 +10134,8 @@ func (m *awsRestxml_serializeOpUploadPart) HandleSeria
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
@@ -10515,7 +10185,6 @@ func (m *awsRestxml_serializeOpUploadPart) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUploadPartInput(v *UploadPartInput, encoder *httpbinding.Encoder) error {
@@ -10643,10 +10312,8 @@ func (m *awsRestxml_serializeOpUploadPartCopy) HandleS
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
@@ -10684,7 +10351,6 @@ func (m *awsRestxml_serializeOpUploadPartCopy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUploadPartCopyInput(v *UploadPartCopyInput, encoder *httpbinding.Encoder) error {
@@ -10797,10 +10463,8 @@ func (m *awsRestxml_serializeOpWriteGetObjectResponse)
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
@@ -10850,7 +10514,6 @@ func (m *awsRestxml_serializeOpWriteGetObjectResponse)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsWriteGetObjectResponseInput(v *WriteGetObjectResponseInput, encoder *httpbinding.Encoder) error {
