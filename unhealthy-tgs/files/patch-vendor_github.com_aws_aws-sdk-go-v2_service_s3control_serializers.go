--- vendor/github.com/aws/aws-sdk-go-v2/service/s3control/serializers.go.orig	2026-11-21 14:46:05 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3control/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyxml "github.com/aws/smithy-go/encoding/xml"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestxml_serializeOpAssociateAccessGrantsId
 func (m *awsRestxml_serializeOpAssociateAccessGrantsIdentityCenter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -86,7 +83,6 @@ func (m *awsRestxml_serializeOpAssociateAccessGrantsId
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsAssociateAccessGrantsIdentityCenterInput(v *AssociateAccessGrantsIdentityCenterInput, encoder *httpbinding.Encoder) error {
@@ -128,10 +124,8 @@ func (m *awsRestxml_serializeOpCreateAccessGrant) Hand
 func (m *awsRestxml_serializeOpCreateAccessGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -187,7 +181,6 @@ func (m *awsRestxml_serializeOpCreateAccessGrant) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateAccessGrantInput(v *CreateAccessGrantInput, encoder *httpbinding.Encoder) error {
@@ -301,10 +294,8 @@ func (m *awsRestxml_serializeOpCreateAccessGrantsInsta
 func (m *awsRestxml_serializeOpCreateAccessGrantsInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -360,7 +351,6 @@ func (m *awsRestxml_serializeOpCreateAccessGrantsInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateAccessGrantsInstanceInput(v *CreateAccessGrantsInstanceInput, encoder *httpbinding.Encoder) error {
@@ -415,10 +405,8 @@ func (m *awsRestxml_serializeOpCreateAccessGrantsLocat
 func (m *awsRestxml_serializeOpCreateAccessGrantsLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -474,7 +462,6 @@ func (m *awsRestxml_serializeOpCreateAccessGrantsLocat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateAccessGrantsLocationInput(v *CreateAccessGrantsLocationInput, encoder *httpbinding.Encoder) error {
@@ -540,10 +527,8 @@ func (m *awsRestxml_serializeOpCreateAccessPoint) Hand
 func (m *awsRestxml_serializeOpCreateAccessPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -599,7 +584,6 @@ func (m *awsRestxml_serializeOpCreateAccessPoint) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateAccessPointInput(v *CreateAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -713,10 +697,8 @@ func (m *awsRestxml_serializeOpCreateAccessPointForObj
 func (m *awsRestxml_serializeOpCreateAccessPointForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -772,7 +754,6 @@ func (m *awsRestxml_serializeOpCreateAccessPointForObj
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateAccessPointForObjectLambdaInput(v *CreateAccessPointForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -825,10 +806,8 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
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
@@ -890,7 +869,6 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketInput(v *CreateBucketInput, encoder *httpbinding.Encoder) error {
@@ -960,10 +938,8 @@ func (m *awsRestxml_serializeOpCreateJob) HandleSerial
 func (m *awsRestxml_serializeOpCreateJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1019,7 +995,6 @@ func (m *awsRestxml_serializeOpCreateJob) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateJobInput(v *CreateJobInput, encoder *httpbinding.Encoder) error {
@@ -1170,10 +1145,8 @@ func (m *awsRestxml_serializeOpCreateMultiRegionAccess
 func (m *awsRestxml_serializeOpCreateMultiRegionAccessPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1229,7 +1202,6 @@ func (m *awsRestxml_serializeOpCreateMultiRegionAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateMultiRegionAccessPointInput(v *CreateMultiRegionAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -1284,10 +1256,8 @@ func (m *awsRestxml_serializeOpCreateStorageLensGroup)
 func (m *awsRestxml_serializeOpCreateStorageLensGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1343,7 +1313,6 @@ func (m *awsRestxml_serializeOpCreateStorageLensGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateStorageLensGroupInput(v *CreateStorageLensGroupInput, encoder *httpbinding.Encoder) error {
@@ -1400,10 +1369,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrant) Hand
 func (m *awsRestxml_serializeOpDeleteAccessGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1441,7 +1408,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrant) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantInput(v *DeleteAccessGrantInput, encoder *httpbinding.Encoder) error {
@@ -1476,10 +1442,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
 func (m *awsRestxml_serializeOpDeleteAccessGrantsInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1517,7 +1481,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantsInstanceInput(v *DeleteAccessGrantsInstanceInput, encoder *httpbinding.Encoder) error {
@@ -1543,10 +1506,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
 func (m *awsRestxml_serializeOpDeleteAccessGrantsInstanceResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1584,7 +1545,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantsInstanceResourcePolicyInput(v *DeleteAccessGrantsInstanceResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -1610,10 +1570,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsLocat
 func (m *awsRestxml_serializeOpDeleteAccessGrantsLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1651,7 +1609,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsLocat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantsLocationInput(v *DeleteAccessGrantsLocationInput, encoder *httpbinding.Encoder) error {
@@ -1686,10 +1643,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPoint) Hand
 func (m *awsRestxml_serializeOpDeleteAccessPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1727,7 +1682,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPoint) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointInput(v *DeleteAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -1762,10 +1716,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointForObj
 func (m *awsRestxml_serializeOpDeleteAccessPointForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1803,7 +1755,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointForObj
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointForObjectLambdaInput(v *DeleteAccessPointForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -1838,10 +1789,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
 func (m *awsRestxml_serializeOpDeleteAccessPointPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1879,7 +1828,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointPolicyInput(v *DeleteAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1914,10 +1862,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
 func (m *awsRestxml_serializeOpDeleteAccessPointPolicyForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1955,7 +1901,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointPolicyForObjectLambdaInput(v *DeleteAccessPointPolicyForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -1990,10 +1935,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointScope)
 func (m *awsRestxml_serializeOpDeleteAccessPointScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2031,7 +1974,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointScope)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointScopeInput(v *DeleteAccessPointScopeInput, encoder *httpbinding.Encoder) error {
@@ -2066,10 +2008,8 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
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
@@ -2107,7 +2047,6 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketInput(v *DeleteBucketInput, encoder *httpbinding.Encoder) error {
@@ -2142,10 +2081,8 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycleCo
 func (m *awsRestxml_serializeOpDeleteBucketLifecycleConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2183,7 +2120,6 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycleCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketLifecycleConfigurationInput(v *DeleteBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2218,10 +2154,8 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
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
@@ -2259,7 +2193,6 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketPolicyInput(v *DeleteBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2294,10 +2227,8 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
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
@@ -2335,7 +2266,6 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketReplicationInput(v *DeleteBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -2370,10 +2300,8 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
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
@@ -2411,7 +2339,6 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketTaggingInput(v *DeleteBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2446,10 +2373,8 @@ func (m *awsRestxml_serializeOpDeleteJobTagging) Handl
 func (m *awsRestxml_serializeOpDeleteJobTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2487,7 +2412,6 @@ func (m *awsRestxml_serializeOpDeleteJobTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteJobTaggingInput(v *DeleteJobTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2522,10 +2446,8 @@ func (m *awsRestxml_serializeOpDeleteMultiRegionAccess
 func (m *awsRestxml_serializeOpDeleteMultiRegionAccessPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2581,7 +2503,6 @@ func (m *awsRestxml_serializeOpDeleteMultiRegionAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteMultiRegionAccessPointInput(v *DeleteMultiRegionAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -2636,10 +2557,8 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
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
@@ -2677,7 +2596,6 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeletePublicAccessBlockInput(v *DeletePublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -2703,10 +2621,8 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
 func (m *awsRestxml_serializeOpDeleteStorageLensConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2744,7 +2660,6 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStorageLensConfigurationInput(v *DeleteStorageLensConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2779,10 +2694,8 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
 func (m *awsRestxml_serializeOpDeleteStorageLensConfigurationTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2820,7 +2733,6 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStorageLensConfigurationTaggingInput(v *DeleteStorageLensConfigurationTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2855,10 +2767,8 @@ func (m *awsRestxml_serializeOpDeleteStorageLensGroup)
 func (m *awsRestxml_serializeOpDeleteStorageLensGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2896,7 +2806,6 @@ func (m *awsRestxml_serializeOpDeleteStorageLensGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStorageLensGroupInput(v *DeleteStorageLensGroupInput, encoder *httpbinding.Encoder) error {
@@ -2931,10 +2840,8 @@ func (m *awsRestxml_serializeOpDescribeJob) HandleSeri
 func (m *awsRestxml_serializeOpDescribeJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2972,7 +2879,6 @@ func (m *awsRestxml_serializeOpDescribeJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeJobInput(v *DescribeJobInput, encoder *httpbinding.Encoder) error {
@@ -3007,10 +2913,8 @@ func (m *awsRestxml_serializeOpDescribeMultiRegionAcce
 func (m *awsRestxml_serializeOpDescribeMultiRegionAccessPointOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3048,7 +2952,6 @@ func (m *awsRestxml_serializeOpDescribeMultiRegionAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeMultiRegionAccessPointOperationInput(v *DescribeMultiRegionAccessPointOperationInput, encoder *httpbinding.Encoder) error {
@@ -3083,10 +2986,8 @@ func (m *awsRestxml_serializeOpDissociateAccessGrantsI
 func (m *awsRestxml_serializeOpDissociateAccessGrantsIdentityCenter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3124,7 +3025,6 @@ func (m *awsRestxml_serializeOpDissociateAccessGrantsI
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDissociateAccessGrantsIdentityCenterInput(v *DissociateAccessGrantsIdentityCenterInput, encoder *httpbinding.Encoder) error {
@@ -3150,10 +3050,8 @@ func (m *awsRestxml_serializeOpGetAccessGrant) HandleS
 func (m *awsRestxml_serializeOpGetAccessGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3191,7 +3089,6 @@ func (m *awsRestxml_serializeOpGetAccessGrant) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantInput(v *GetAccessGrantInput, encoder *httpbinding.Encoder) error {
@@ -3226,10 +3123,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 func (m *awsRestxml_serializeOpGetAccessGrantsInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3267,7 +3162,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsInstanceInput(v *GetAccessGrantsInstanceInput, encoder *httpbinding.Encoder) error {
@@ -3293,10 +3187,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 func (m *awsRestxml_serializeOpGetAccessGrantsInstanceForPrefix) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3334,7 +3226,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsInstanceForPrefixInput(v *GetAccessGrantsInstanceForPrefixInput, encoder *httpbinding.Encoder) error {
@@ -3364,10 +3255,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 func (m *awsRestxml_serializeOpGetAccessGrantsInstanceResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3405,7 +3294,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsInstanceResourcePolicyInput(v *GetAccessGrantsInstanceResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -3431,10 +3319,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsLocation
 func (m *awsRestxml_serializeOpGetAccessGrantsLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3472,7 +3358,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsLocation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsLocationInput(v *GetAccessGrantsLocationInput, encoder *httpbinding.Encoder) error {
@@ -3507,10 +3392,8 @@ func (m *awsRestxml_serializeOpGetAccessPoint) HandleS
 func (m *awsRestxml_serializeOpGetAccessPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3548,7 +3431,6 @@ func (m *awsRestxml_serializeOpGetAccessPoint) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointInput(v *GetAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -3583,10 +3465,8 @@ func (m *awsRestxml_serializeOpGetAccessPointConfigura
 func (m *awsRestxml_serializeOpGetAccessPointConfigurationForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3624,7 +3504,6 @@ func (m *awsRestxml_serializeOpGetAccessPointConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointConfigurationForObjectLambdaInput(v *GetAccessPointConfigurationForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -3659,10 +3538,8 @@ func (m *awsRestxml_serializeOpGetAccessPointForObject
 func (m *awsRestxml_serializeOpGetAccessPointForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3700,7 +3577,6 @@ func (m *awsRestxml_serializeOpGetAccessPointForObject
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointForObjectLambdaInput(v *GetAccessPointForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -3735,10 +3611,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicy) H
 func (m *awsRestxml_serializeOpGetAccessPointPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3776,7 +3650,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyInput(v *GetAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3811,10 +3684,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicyFor
 func (m *awsRestxml_serializeOpGetAccessPointPolicyForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3852,7 +3723,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicyFor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyForObjectLambdaInput(v *GetAccessPointPolicyForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -3887,10 +3757,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
 func (m *awsRestxml_serializeOpGetAccessPointPolicyStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3928,7 +3796,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyStatusInput(v *GetAccessPointPolicyStatusInput, encoder *httpbinding.Encoder) error {
@@ -3963,10 +3830,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
 func (m *awsRestxml_serializeOpGetAccessPointPolicyStatusForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4004,7 +3869,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyStatusForObjectLambdaInput(v *GetAccessPointPolicyStatusForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -4039,10 +3903,8 @@ func (m *awsRestxml_serializeOpGetAccessPointScope) Ha
 func (m *awsRestxml_serializeOpGetAccessPointScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4080,7 +3942,6 @@ func (m *awsRestxml_serializeOpGetAccessPointScope) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointScopeInput(v *GetAccessPointScopeInput, encoder *httpbinding.Encoder) error {
@@ -4115,10 +3976,8 @@ func (m *awsRestxml_serializeOpGetBucket) HandleSerial
 func (m *awsRestxml_serializeOpGetBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4156,7 +4015,6 @@ func (m *awsRestxml_serializeOpGetBucket) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketInput(v *GetBucketInput, encoder *httpbinding.Encoder) error {
@@ -4191,10 +4049,8 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
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
@@ -4232,7 +4088,6 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLifecycleConfigurationInput(v *GetBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4267,10 +4122,8 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
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
@@ -4308,7 +4161,6 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketPolicyInput(v *GetBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4343,10 +4195,8 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
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
@@ -4384,7 +4234,6 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketReplicationInput(v *GetBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -4419,10 +4268,8 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
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
@@ -4460,7 +4307,6 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketTaggingInput(v *GetBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -4495,10 +4341,8 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
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
@@ -4536,7 +4380,6 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketVersioningInput(v *GetBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -4571,10 +4414,8 @@ func (m *awsRestxml_serializeOpGetDataAccess) HandleSe
 func (m *awsRestxml_serializeOpGetDataAccess) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4612,7 +4453,6 @@ func (m *awsRestxml_serializeOpGetDataAccess) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDataAccessInput(v *GetDataAccessInput, encoder *httpbinding.Encoder) error {
@@ -4666,10 +4506,8 @@ func (m *awsRestxml_serializeOpGetJobTagging) HandleSe
 func (m *awsRestxml_serializeOpGetJobTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4707,7 +4545,6 @@ func (m *awsRestxml_serializeOpGetJobTagging) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetJobTaggingInput(v *GetJobTaggingInput, encoder *httpbinding.Encoder) error {
@@ -4742,10 +4579,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 func (m *awsRestxml_serializeOpGetMultiRegionAccessPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4783,7 +4618,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointInput(v *GetMultiRegionAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -4818,10 +4652,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 func (m *awsRestxml_serializeOpGetMultiRegionAccessPointPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4859,7 +4691,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointPolicyInput(v *GetMultiRegionAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4894,10 +4725,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 func (m *awsRestxml_serializeOpGetMultiRegionAccessPointPolicyStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4935,7 +4764,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointPolicyStatusInput(v *GetMultiRegionAccessPointPolicyStatusInput, encoder *httpbinding.Encoder) error {
@@ -4970,10 +4798,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 func (m *awsRestxml_serializeOpGetMultiRegionAccessPointRoutes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5011,7 +4837,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointRoutesInput(v *GetMultiRegionAccessPointRoutesInput, encoder *httpbinding.Encoder) error {
@@ -5046,10 +4871,8 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
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
@@ -5087,7 +4910,6 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicAccessBlockInput(v *GetPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -5113,10 +4935,8 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
 func (m *awsRestxml_serializeOpGetStorageLensConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5154,7 +4974,6 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStorageLensConfigurationInput(v *GetStorageLensConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5189,10 +5008,8 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
 func (m *awsRestxml_serializeOpGetStorageLensConfigurationTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5230,7 +5047,6 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStorageLensConfigurationTaggingInput(v *GetStorageLensConfigurationTaggingInput, encoder *httpbinding.Encoder) error {
@@ -5265,10 +5081,8 @@ func (m *awsRestxml_serializeOpGetStorageLensGroup) Ha
 func (m *awsRestxml_serializeOpGetStorageLensGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5306,7 +5120,6 @@ func (m *awsRestxml_serializeOpGetStorageLensGroup) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStorageLensGroupInput(v *GetStorageLensGroupInput, encoder *httpbinding.Encoder) error {
@@ -5341,10 +5154,8 @@ func (m *awsRestxml_serializeOpListAccessGrants) Handl
 func (m *awsRestxml_serializeOpListAccessGrants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5382,7 +5193,6 @@ func (m *awsRestxml_serializeOpListAccessGrants) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessGrantsInput(v *ListAccessGrantsInput, encoder *httpbinding.Encoder) error {
@@ -5436,10 +5246,8 @@ func (m *awsRestxml_serializeOpListAccessGrantsInstanc
 func (m *awsRestxml_serializeOpListAccessGrantsInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5477,7 +5285,6 @@ func (m *awsRestxml_serializeOpListAccessGrantsInstanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessGrantsInstancesInput(v *ListAccessGrantsInstancesInput, encoder *httpbinding.Encoder) error {
@@ -5511,10 +5318,8 @@ func (m *awsRestxml_serializeOpListAccessGrantsLocatio
 func (m *awsRestxml_serializeOpListAccessGrantsLocations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5552,7 +5357,6 @@ func (m *awsRestxml_serializeOpListAccessGrantsLocatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessGrantsLocationsInput(v *ListAccessGrantsLocationsInput, encoder *httpbinding.Encoder) error {
@@ -5590,10 +5394,8 @@ func (m *awsRestxml_serializeOpListAccessPoints) Handl
 func (m *awsRestxml_serializeOpListAccessPoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5631,7 +5433,6 @@ func (m *awsRestxml_serializeOpListAccessPoints) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessPointsInput(v *ListAccessPointsInput, encoder *httpbinding.Encoder) error {
@@ -5677,10 +5478,8 @@ func (m *awsRestxml_serializeOpListAccessPointsForDire
 func (m *awsRestxml_serializeOpListAccessPointsForDirectoryBuckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5718,7 +5517,6 @@ func (m *awsRestxml_serializeOpListAccessPointsForDire
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessPointsForDirectoryBucketsInput(v *ListAccessPointsForDirectoryBucketsInput, encoder *httpbinding.Encoder) error {
@@ -5756,10 +5554,8 @@ func (m *awsRestxml_serializeOpListAccessPointsForObje
 func (m *awsRestxml_serializeOpListAccessPointsForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5797,7 +5593,6 @@ func (m *awsRestxml_serializeOpListAccessPointsForObje
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessPointsForObjectLambdaInput(v *ListAccessPointsForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -5831,10 +5626,8 @@ func (m *awsRestxml_serializeOpListCallerAccessGrants)
 func (m *awsRestxml_serializeOpListCallerAccessGrants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5872,7 +5665,6 @@ func (m *awsRestxml_serializeOpListCallerAccessGrants)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCallerAccessGrantsInput(v *ListCallerAccessGrantsInput, encoder *httpbinding.Encoder) error {
@@ -5918,10 +5710,8 @@ func (m *awsRestxml_serializeOpListJobs) HandleSeriali
 func (m *awsRestxml_serializeOpListJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5959,7 +5749,6 @@ func (m *awsRestxml_serializeOpListJobs) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListJobsInput(v *ListJobsInput, encoder *httpbinding.Encoder) error {
@@ -5999,10 +5788,8 @@ func (m *awsRestxml_serializeOpListMultiRegionAccessPo
 func (m *awsRestxml_serializeOpListMultiRegionAccessPoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6040,7 +5827,6 @@ func (m *awsRestxml_serializeOpListMultiRegionAccessPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListMultiRegionAccessPointsInput(v *ListMultiRegionAccessPointsInput, encoder *httpbinding.Encoder) error {
@@ -6074,10 +5860,8 @@ func (m *awsRestxml_serializeOpListRegionalBuckets) Ha
 func (m *awsRestxml_serializeOpListRegionalBuckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6115,7 +5899,6 @@ func (m *awsRestxml_serializeOpListRegionalBuckets) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListRegionalBucketsInput(v *ListRegionalBucketsInput, encoder *httpbinding.Encoder) error {
@@ -6154,10 +5937,8 @@ func (m *awsRestxml_serializeOpListStorageLensConfigur
 func (m *awsRestxml_serializeOpListStorageLensConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6195,7 +5976,6 @@ func (m *awsRestxml_serializeOpListStorageLensConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListStorageLensConfigurationsInput(v *ListStorageLensConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -6225,10 +6005,8 @@ func (m *awsRestxml_serializeOpListStorageLensGroups) 
 func (m *awsRestxml_serializeOpListStorageLensGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6266,7 +6044,6 @@ func (m *awsRestxml_serializeOpListStorageLensGroups) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListStorageLensGroupsInput(v *ListStorageLensGroupsInput, encoder *httpbinding.Encoder) error {
@@ -6296,10 +6073,8 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
 func (m *awsRestxml_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6337,7 +6112,6 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6372,10 +6146,8 @@ func (m *awsRestxml_serializeOpPutAccessGrantsInstance
 func (m *awsRestxml_serializeOpPutAccessGrantsInstanceResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6431,7 +6203,6 @@ func (m *awsRestxml_serializeOpPutAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessGrantsInstanceResourcePolicyInput(v *PutAccessGrantsInstanceResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -6484,10 +6255,8 @@ func (m *awsRestxml_serializeOpPutAccessPointConfigura
 func (m *awsRestxml_serializeOpPutAccessPointConfigurationForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6543,7 +6312,6 @@ func (m *awsRestxml_serializeOpPutAccessPointConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointConfigurationForObjectLambdaInput(v *PutAccessPointConfigurationForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -6596,10 +6364,8 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicy) H
 func (m *awsRestxml_serializeOpPutAccessPointPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6655,7 +6421,6 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointPolicyInput(v *PutAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -6706,10 +6471,8 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicyFor
 func (m *awsRestxml_serializeOpPutAccessPointPolicyForObjectLambda) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6765,7 +6528,6 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicyFor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointPolicyForObjectLambdaInput(v *PutAccessPointPolicyForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -6816,10 +6578,8 @@ func (m *awsRestxml_serializeOpPutAccessPointScope) Ha
 func (m *awsRestxml_serializeOpPutAccessPointScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6875,7 +6635,6 @@ func (m *awsRestxml_serializeOpPutAccessPointScope) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointScopeInput(v *PutAccessPointScopeInput, encoder *httpbinding.Encoder) error {
@@ -6928,10 +6687,8 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
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
@@ -6993,7 +6750,6 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketLifecycleConfigurationInput(v *PutBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7028,10 +6784,8 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
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
@@ -7087,7 +6841,6 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketPolicyInput(v *PutBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7143,10 +6896,8 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
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
@@ -7208,7 +6959,6 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketReplicationInput(v *PutBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -7243,10 +6993,8 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
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
@@ -7308,7 +7056,6 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketTaggingInput(v *PutBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -7343,10 +7090,8 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
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
@@ -7408,7 +7153,6 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketVersioningInput(v *PutBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -7448,10 +7192,8 @@ func (m *awsRestxml_serializeOpPutJobTagging) HandleSe
 func (m *awsRestxml_serializeOpPutJobTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7507,7 +7249,6 @@ func (m *awsRestxml_serializeOpPutJobTagging) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutJobTaggingInput(v *PutJobTaggingInput, encoder *httpbinding.Encoder) error {
@@ -7560,10 +7301,8 @@ func (m *awsRestxml_serializeOpPutMultiRegionAccessPoi
 func (m *awsRestxml_serializeOpPutMultiRegionAccessPointPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7619,7 +7358,6 @@ func (m *awsRestxml_serializeOpPutMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutMultiRegionAccessPointPolicyInput(v *PutMultiRegionAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7674,10 +7412,8 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
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
@@ -7739,7 +7475,6 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutPublicAccessBlockInput(v *PutPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -7765,10 +7500,8 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
 func (m *awsRestxml_serializeOpPutStorageLensConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7824,7 +7557,6 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutStorageLensConfigurationInput(v *PutStorageLensConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7890,10 +7622,8 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
 func (m *awsRestxml_serializeOpPutStorageLensConfigurationTagging) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7949,7 +7679,6 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutStorageLensConfigurationTaggingInput(v *PutStorageLensConfigurationTaggingInput, encoder *httpbinding.Encoder) error {
@@ -8002,10 +7731,8 @@ func (m *awsRestxml_serializeOpSubmitMultiRegionAccess
 func (m *awsRestxml_serializeOpSubmitMultiRegionAccessPointRoutes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8061,7 +7788,6 @@ func (m *awsRestxml_serializeOpSubmitMultiRegionAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsSubmitMultiRegionAccessPointRoutesInput(v *SubmitMultiRegionAccessPointRoutesInput, encoder *httpbinding.Encoder) error {
@@ -8114,10 +7840,8 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
 func (m *awsRestxml_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8173,7 +7897,6 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -8226,10 +7949,8 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
 func (m *awsRestxml_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8267,7 +7988,6 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -8308,10 +8028,8 @@ func (m *awsRestxml_serializeOpUpdateAccessGrantsLocat
 func (m *awsRestxml_serializeOpUpdateAccessGrantsLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8367,7 +8085,6 @@ func (m *awsRestxml_serializeOpUpdateAccessGrantsLocat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateAccessGrantsLocationInput(v *UpdateAccessGrantsLocationInput, encoder *httpbinding.Encoder) error {
@@ -8418,10 +8135,8 @@ func (m *awsRestxml_serializeOpUpdateJobPriority) Hand
 func (m *awsRestxml_serializeOpUpdateJobPriority) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8459,7 +8174,6 @@ func (m *awsRestxml_serializeOpUpdateJobPriority) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateJobPriorityInput(v *UpdateJobPriorityInput, encoder *httpbinding.Encoder) error {
@@ -8498,10 +8212,8 @@ func (m *awsRestxml_serializeOpUpdateJobStatus) Handle
 func (m *awsRestxml_serializeOpUpdateJobStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8539,7 +8251,6 @@ func (m *awsRestxml_serializeOpUpdateJobStatus) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateJobStatusInput(v *UpdateJobStatusInput, encoder *httpbinding.Encoder) error {
@@ -8582,10 +8293,8 @@ func (m *awsRestxml_serializeOpUpdateStorageLensGroup)
 func (m *awsRestxml_serializeOpUpdateStorageLensGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8641,7 +8350,6 @@ func (m *awsRestxml_serializeOpUpdateStorageLensGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateStorageLensGroupInput(v *UpdateStorageLensGroupInput, encoder *httpbinding.Encoder) error {
