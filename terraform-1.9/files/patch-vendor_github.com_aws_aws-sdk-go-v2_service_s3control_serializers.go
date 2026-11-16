--- vendor/github.com/aws/aws-sdk-go-v2/service/s3control/serializers.go.orig	2025-07-19 01:28:18 UTC
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
@@ -700,10 +684,8 @@ func (m *awsRestxml_serializeOpCreateAccessPointForObj
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
@@ -759,7 +741,6 @@ func (m *awsRestxml_serializeOpCreateAccessPointForObj
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateAccessPointForObjectLambdaInput(v *CreateAccessPointForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -812,10 +793,8 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
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
@@ -877,7 +856,6 @@ func (m *awsRestxml_serializeOpCreateBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateBucketInput(v *CreateBucketInput, encoder *httpbinding.Encoder) error {
@@ -947,10 +925,8 @@ func (m *awsRestxml_serializeOpCreateJob) HandleSerial
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
@@ -1006,7 +982,6 @@ func (m *awsRestxml_serializeOpCreateJob) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateJobInput(v *CreateJobInput, encoder *httpbinding.Encoder) error {
@@ -1157,10 +1132,8 @@ func (m *awsRestxml_serializeOpCreateMultiRegionAccess
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
@@ -1216,7 +1189,6 @@ func (m *awsRestxml_serializeOpCreateMultiRegionAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateMultiRegionAccessPointInput(v *CreateMultiRegionAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -1271,10 +1243,8 @@ func (m *awsRestxml_serializeOpCreateStorageLensGroup)
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
@@ -1330,7 +1300,6 @@ func (m *awsRestxml_serializeOpCreateStorageLensGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateStorageLensGroupInput(v *CreateStorageLensGroupInput, encoder *httpbinding.Encoder) error {
@@ -1387,10 +1356,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrant) Hand
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
@@ -1428,7 +1395,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrant) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantInput(v *DeleteAccessGrantInput, encoder *httpbinding.Encoder) error {
@@ -1463,10 +1429,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
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
@@ -1504,7 +1468,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantsInstanceInput(v *DeleteAccessGrantsInstanceInput, encoder *httpbinding.Encoder) error {
@@ -1530,10 +1493,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
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
@@ -1571,7 +1532,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantsInstanceResourcePolicyInput(v *DeleteAccessGrantsInstanceResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -1597,10 +1557,8 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsLocat
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
@@ -1638,7 +1596,6 @@ func (m *awsRestxml_serializeOpDeleteAccessGrantsLocat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessGrantsLocationInput(v *DeleteAccessGrantsLocationInput, encoder *httpbinding.Encoder) error {
@@ -1673,10 +1630,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPoint) Hand
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
@@ -1714,7 +1669,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPoint) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointInput(v *DeleteAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -1749,10 +1703,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointForObj
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
@@ -1790,7 +1742,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointForObj
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointForObjectLambdaInput(v *DeleteAccessPointForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -1825,10 +1776,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
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
@@ -1866,7 +1815,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointPolicyInput(v *DeleteAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1901,10 +1849,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
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
@@ -1942,7 +1888,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointPolicyForObjectLambdaInput(v *DeleteAccessPointPolicyForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -1977,10 +1922,8 @@ func (m *awsRestxml_serializeOpDeleteAccessPointScope)
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
@@ -2018,7 +1961,6 @@ func (m *awsRestxml_serializeOpDeleteAccessPointScope)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAccessPointScopeInput(v *DeleteAccessPointScopeInput, encoder *httpbinding.Encoder) error {
@@ -2053,10 +1995,8 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
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
@@ -2094,7 +2034,6 @@ func (m *awsRestxml_serializeOpDeleteBucket) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketInput(v *DeleteBucketInput, encoder *httpbinding.Encoder) error {
@@ -2129,10 +2068,8 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycleCo
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
@@ -2170,7 +2107,6 @@ func (m *awsRestxml_serializeOpDeleteBucketLifecycleCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketLifecycleConfigurationInput(v *DeleteBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2205,10 +2141,8 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
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
@@ -2246,7 +2180,6 @@ func (m *awsRestxml_serializeOpDeleteBucketPolicy) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketPolicyInput(v *DeleteBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2281,10 +2214,8 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
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
@@ -2322,7 +2253,6 @@ func (m *awsRestxml_serializeOpDeleteBucketReplication
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketReplicationInput(v *DeleteBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -2357,10 +2287,8 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
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
@@ -2398,7 +2326,6 @@ func (m *awsRestxml_serializeOpDeleteBucketTagging) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteBucketTaggingInput(v *DeleteBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2433,10 +2360,8 @@ func (m *awsRestxml_serializeOpDeleteJobTagging) Handl
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
@@ -2474,7 +2399,6 @@ func (m *awsRestxml_serializeOpDeleteJobTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteJobTaggingInput(v *DeleteJobTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2509,10 +2433,8 @@ func (m *awsRestxml_serializeOpDeleteMultiRegionAccess
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
@@ -2568,7 +2490,6 @@ func (m *awsRestxml_serializeOpDeleteMultiRegionAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteMultiRegionAccessPointInput(v *DeleteMultiRegionAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -2623,10 +2544,8 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
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
@@ -2664,7 +2583,6 @@ func (m *awsRestxml_serializeOpDeletePublicAccessBlock
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeletePublicAccessBlockInput(v *DeletePublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -2690,10 +2608,8 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
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
@@ -2731,7 +2647,6 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStorageLensConfigurationInput(v *DeleteStorageLensConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2766,10 +2681,8 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
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
@@ -2807,7 +2720,6 @@ func (m *awsRestxml_serializeOpDeleteStorageLensConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStorageLensConfigurationTaggingInput(v *DeleteStorageLensConfigurationTaggingInput, encoder *httpbinding.Encoder) error {
@@ -2842,10 +2754,8 @@ func (m *awsRestxml_serializeOpDeleteStorageLensGroup)
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
@@ -2883,7 +2793,6 @@ func (m *awsRestxml_serializeOpDeleteStorageLensGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStorageLensGroupInput(v *DeleteStorageLensGroupInput, encoder *httpbinding.Encoder) error {
@@ -2918,10 +2827,8 @@ func (m *awsRestxml_serializeOpDescribeJob) HandleSeri
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
@@ -2959,7 +2866,6 @@ func (m *awsRestxml_serializeOpDescribeJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeJobInput(v *DescribeJobInput, encoder *httpbinding.Encoder) error {
@@ -2994,10 +2900,8 @@ func (m *awsRestxml_serializeOpDescribeMultiRegionAcce
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
@@ -3035,7 +2939,6 @@ func (m *awsRestxml_serializeOpDescribeMultiRegionAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeMultiRegionAccessPointOperationInput(v *DescribeMultiRegionAccessPointOperationInput, encoder *httpbinding.Encoder) error {
@@ -3070,10 +2973,8 @@ func (m *awsRestxml_serializeOpDissociateAccessGrantsI
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
@@ -3111,7 +3012,6 @@ func (m *awsRestxml_serializeOpDissociateAccessGrantsI
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDissociateAccessGrantsIdentityCenterInput(v *DissociateAccessGrantsIdentityCenterInput, encoder *httpbinding.Encoder) error {
@@ -3137,10 +3037,8 @@ func (m *awsRestxml_serializeOpGetAccessGrant) HandleS
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
@@ -3178,7 +3076,6 @@ func (m *awsRestxml_serializeOpGetAccessGrant) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantInput(v *GetAccessGrantInput, encoder *httpbinding.Encoder) error {
@@ -3213,10 +3110,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
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
@@ -3254,7 +3149,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsInstanceInput(v *GetAccessGrantsInstanceInput, encoder *httpbinding.Encoder) error {
@@ -3280,10 +3174,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
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
@@ -3321,7 +3213,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsInstanceForPrefixInput(v *GetAccessGrantsInstanceForPrefixInput, encoder *httpbinding.Encoder) error {
@@ -3351,10 +3242,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
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
@@ -3392,7 +3281,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsInstanceResourcePolicyInput(v *GetAccessGrantsInstanceResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -3418,10 +3306,8 @@ func (m *awsRestxml_serializeOpGetAccessGrantsLocation
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
@@ -3459,7 +3345,6 @@ func (m *awsRestxml_serializeOpGetAccessGrantsLocation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessGrantsLocationInput(v *GetAccessGrantsLocationInput, encoder *httpbinding.Encoder) error {
@@ -3494,10 +3379,8 @@ func (m *awsRestxml_serializeOpGetAccessPoint) HandleS
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
@@ -3535,7 +3418,6 @@ func (m *awsRestxml_serializeOpGetAccessPoint) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointInput(v *GetAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -3570,10 +3452,8 @@ func (m *awsRestxml_serializeOpGetAccessPointConfigura
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
@@ -3611,7 +3491,6 @@ func (m *awsRestxml_serializeOpGetAccessPointConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointConfigurationForObjectLambdaInput(v *GetAccessPointConfigurationForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -3646,10 +3525,8 @@ func (m *awsRestxml_serializeOpGetAccessPointForObject
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
@@ -3687,7 +3564,6 @@ func (m *awsRestxml_serializeOpGetAccessPointForObject
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointForObjectLambdaInput(v *GetAccessPointForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -3722,10 +3598,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicy) H
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
@@ -3763,7 +3637,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyInput(v *GetAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3798,10 +3671,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicyFor
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
@@ -3839,7 +3710,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicyFor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyForObjectLambdaInput(v *GetAccessPointPolicyForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -3874,10 +3744,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
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
@@ -3915,7 +3783,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyStatusInput(v *GetAccessPointPolicyStatusInput, encoder *httpbinding.Encoder) error {
@@ -3950,10 +3817,8 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
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
@@ -3991,7 +3856,6 @@ func (m *awsRestxml_serializeOpGetAccessPointPolicySta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointPolicyStatusForObjectLambdaInput(v *GetAccessPointPolicyStatusForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -4026,10 +3890,8 @@ func (m *awsRestxml_serializeOpGetAccessPointScope) Ha
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
@@ -4067,7 +3929,6 @@ func (m *awsRestxml_serializeOpGetAccessPointScope) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccessPointScopeInput(v *GetAccessPointScopeInput, encoder *httpbinding.Encoder) error {
@@ -4102,10 +3963,8 @@ func (m *awsRestxml_serializeOpGetBucket) HandleSerial
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
@@ -4143,7 +4002,6 @@ func (m *awsRestxml_serializeOpGetBucket) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketInput(v *GetBucketInput, encoder *httpbinding.Encoder) error {
@@ -4178,10 +4036,8 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
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
@@ -4219,7 +4075,6 @@ func (m *awsRestxml_serializeOpGetBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketLifecycleConfigurationInput(v *GetBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4254,10 +4109,8 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
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
@@ -4295,7 +4148,6 @@ func (m *awsRestxml_serializeOpGetBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketPolicyInput(v *GetBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4330,10 +4182,8 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
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
@@ -4371,7 +4221,6 @@ func (m *awsRestxml_serializeOpGetBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketReplicationInput(v *GetBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -4406,10 +4255,8 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
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
@@ -4447,7 +4294,6 @@ func (m *awsRestxml_serializeOpGetBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketTaggingInput(v *GetBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -4482,10 +4328,8 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
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
@@ -4523,7 +4367,6 @@ func (m *awsRestxml_serializeOpGetBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetBucketVersioningInput(v *GetBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -4558,10 +4401,8 @@ func (m *awsRestxml_serializeOpGetDataAccess) HandleSe
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
@@ -4599,7 +4440,6 @@ func (m *awsRestxml_serializeOpGetDataAccess) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDataAccessInput(v *GetDataAccessInput, encoder *httpbinding.Encoder) error {
@@ -4645,10 +4485,8 @@ func (m *awsRestxml_serializeOpGetJobTagging) HandleSe
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
@@ -4686,7 +4524,6 @@ func (m *awsRestxml_serializeOpGetJobTagging) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetJobTaggingInput(v *GetJobTaggingInput, encoder *httpbinding.Encoder) error {
@@ -4721,10 +4558,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
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
@@ -4762,7 +4597,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointInput(v *GetMultiRegionAccessPointInput, encoder *httpbinding.Encoder) error {
@@ -4797,10 +4631,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
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
@@ -4838,7 +4670,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointPolicyInput(v *GetMultiRegionAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4873,10 +4704,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
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
@@ -4914,7 +4743,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointPolicyStatusInput(v *GetMultiRegionAccessPointPolicyStatusInput, encoder *httpbinding.Encoder) error {
@@ -4949,10 +4777,8 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
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
@@ -4990,7 +4816,6 @@ func (m *awsRestxml_serializeOpGetMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMultiRegionAccessPointRoutesInput(v *GetMultiRegionAccessPointRoutesInput, encoder *httpbinding.Encoder) error {
@@ -5025,10 +4850,8 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
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
@@ -5066,7 +4889,6 @@ func (m *awsRestxml_serializeOpGetPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicAccessBlockInput(v *GetPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -5092,10 +4914,8 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
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
@@ -5133,7 +4953,6 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStorageLensConfigurationInput(v *GetStorageLensConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5168,10 +4987,8 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
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
@@ -5209,7 +5026,6 @@ func (m *awsRestxml_serializeOpGetStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStorageLensConfigurationTaggingInput(v *GetStorageLensConfigurationTaggingInput, encoder *httpbinding.Encoder) error {
@@ -5244,10 +5060,8 @@ func (m *awsRestxml_serializeOpGetStorageLensGroup) Ha
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
@@ -5285,7 +5099,6 @@ func (m *awsRestxml_serializeOpGetStorageLensGroup) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStorageLensGroupInput(v *GetStorageLensGroupInput, encoder *httpbinding.Encoder) error {
@@ -5320,10 +5133,8 @@ func (m *awsRestxml_serializeOpListAccessGrants) Handl
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
@@ -5361,7 +5172,6 @@ func (m *awsRestxml_serializeOpListAccessGrants) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessGrantsInput(v *ListAccessGrantsInput, encoder *httpbinding.Encoder) error {
@@ -5415,10 +5225,8 @@ func (m *awsRestxml_serializeOpListAccessGrantsInstanc
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
@@ -5456,7 +5264,6 @@ func (m *awsRestxml_serializeOpListAccessGrantsInstanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessGrantsInstancesInput(v *ListAccessGrantsInstancesInput, encoder *httpbinding.Encoder) error {
@@ -5490,10 +5297,8 @@ func (m *awsRestxml_serializeOpListAccessGrantsLocatio
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
@@ -5531,7 +5336,6 @@ func (m *awsRestxml_serializeOpListAccessGrantsLocatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessGrantsLocationsInput(v *ListAccessGrantsLocationsInput, encoder *httpbinding.Encoder) error {
@@ -5569,10 +5373,8 @@ func (m *awsRestxml_serializeOpListAccessPoints) Handl
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
@@ -5610,7 +5412,6 @@ func (m *awsRestxml_serializeOpListAccessPoints) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessPointsInput(v *ListAccessPointsInput, encoder *httpbinding.Encoder) error {
@@ -5656,10 +5457,8 @@ func (m *awsRestxml_serializeOpListAccessPointsForDire
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
@@ -5697,7 +5496,6 @@ func (m *awsRestxml_serializeOpListAccessPointsForDire
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessPointsForDirectoryBucketsInput(v *ListAccessPointsForDirectoryBucketsInput, encoder *httpbinding.Encoder) error {
@@ -5735,10 +5533,8 @@ func (m *awsRestxml_serializeOpListAccessPointsForObje
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
@@ -5776,7 +5572,6 @@ func (m *awsRestxml_serializeOpListAccessPointsForObje
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAccessPointsForObjectLambdaInput(v *ListAccessPointsForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -5810,10 +5605,8 @@ func (m *awsRestxml_serializeOpListCallerAccessGrants)
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
@@ -5851,7 +5644,6 @@ func (m *awsRestxml_serializeOpListCallerAccessGrants)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCallerAccessGrantsInput(v *ListCallerAccessGrantsInput, encoder *httpbinding.Encoder) error {
@@ -5893,10 +5685,8 @@ func (m *awsRestxml_serializeOpListJobs) HandleSeriali
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
@@ -5934,7 +5724,6 @@ func (m *awsRestxml_serializeOpListJobs) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListJobsInput(v *ListJobsInput, encoder *httpbinding.Encoder) error {
@@ -5974,10 +5763,8 @@ func (m *awsRestxml_serializeOpListMultiRegionAccessPo
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
@@ -6015,7 +5802,6 @@ func (m *awsRestxml_serializeOpListMultiRegionAccessPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListMultiRegionAccessPointsInput(v *ListMultiRegionAccessPointsInput, encoder *httpbinding.Encoder) error {
@@ -6049,10 +5835,8 @@ func (m *awsRestxml_serializeOpListRegionalBuckets) Ha
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
@@ -6090,7 +5874,6 @@ func (m *awsRestxml_serializeOpListRegionalBuckets) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListRegionalBucketsInput(v *ListRegionalBucketsInput, encoder *httpbinding.Encoder) error {
@@ -6129,10 +5912,8 @@ func (m *awsRestxml_serializeOpListStorageLensConfigur
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
@@ -6170,7 +5951,6 @@ func (m *awsRestxml_serializeOpListStorageLensConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListStorageLensConfigurationsInput(v *ListStorageLensConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -6200,10 +5980,8 @@ func (m *awsRestxml_serializeOpListStorageLensGroups) 
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
@@ -6241,7 +6019,6 @@ func (m *awsRestxml_serializeOpListStorageLensGroups) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListStorageLensGroupsInput(v *ListStorageLensGroupsInput, encoder *httpbinding.Encoder) error {
@@ -6271,10 +6048,8 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
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
@@ -6312,7 +6087,6 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6347,10 +6121,8 @@ func (m *awsRestxml_serializeOpPutAccessGrantsInstance
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
@@ -6406,7 +6178,6 @@ func (m *awsRestxml_serializeOpPutAccessGrantsInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessGrantsInstanceResourcePolicyInput(v *PutAccessGrantsInstanceResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -6459,10 +6230,8 @@ func (m *awsRestxml_serializeOpPutAccessPointConfigura
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
@@ -6518,7 +6287,6 @@ func (m *awsRestxml_serializeOpPutAccessPointConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointConfigurationForObjectLambdaInput(v *PutAccessPointConfigurationForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -6571,10 +6339,8 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicy) H
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
@@ -6630,7 +6396,6 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointPolicyInput(v *PutAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -6681,10 +6446,8 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicyFor
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
@@ -6740,7 +6503,6 @@ func (m *awsRestxml_serializeOpPutAccessPointPolicyFor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointPolicyForObjectLambdaInput(v *PutAccessPointPolicyForObjectLambdaInput, encoder *httpbinding.Encoder) error {
@@ -6791,10 +6553,8 @@ func (m *awsRestxml_serializeOpPutAccessPointScope) Ha
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
@@ -6850,7 +6610,6 @@ func (m *awsRestxml_serializeOpPutAccessPointScope) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutAccessPointScopeInput(v *PutAccessPointScopeInput, encoder *httpbinding.Encoder) error {
@@ -6903,10 +6662,8 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
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
@@ -6968,7 +6725,6 @@ func (m *awsRestxml_serializeOpPutBucketLifecycleConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketLifecycleConfigurationInput(v *PutBucketLifecycleConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7003,10 +6759,8 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
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
@@ -7062,7 +6816,6 @@ func (m *awsRestxml_serializeOpPutBucketPolicy) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketPolicyInput(v *PutBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7118,10 +6871,8 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
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
@@ -7183,7 +6934,6 @@ func (m *awsRestxml_serializeOpPutBucketReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketReplicationInput(v *PutBucketReplicationInput, encoder *httpbinding.Encoder) error {
@@ -7218,10 +6968,8 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
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
@@ -7283,7 +7031,6 @@ func (m *awsRestxml_serializeOpPutBucketTagging) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketTaggingInput(v *PutBucketTaggingInput, encoder *httpbinding.Encoder) error {
@@ -7318,10 +7065,8 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
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
@@ -7383,7 +7128,6 @@ func (m *awsRestxml_serializeOpPutBucketVersioning) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutBucketVersioningInput(v *PutBucketVersioningInput, encoder *httpbinding.Encoder) error {
@@ -7423,10 +7167,8 @@ func (m *awsRestxml_serializeOpPutJobTagging) HandleSe
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
@@ -7482,7 +7224,6 @@ func (m *awsRestxml_serializeOpPutJobTagging) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutJobTaggingInput(v *PutJobTaggingInput, encoder *httpbinding.Encoder) error {
@@ -7535,10 +7276,8 @@ func (m *awsRestxml_serializeOpPutMultiRegionAccessPoi
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
@@ -7594,7 +7333,6 @@ func (m *awsRestxml_serializeOpPutMultiRegionAccessPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutMultiRegionAccessPointPolicyInput(v *PutMultiRegionAccessPointPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7649,10 +7387,8 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
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
@@ -7714,7 +7450,6 @@ func (m *awsRestxml_serializeOpPutPublicAccessBlock) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutPublicAccessBlockInput(v *PutPublicAccessBlockInput, encoder *httpbinding.Encoder) error {
@@ -7740,10 +7475,8 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
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
@@ -7799,7 +7532,6 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutStorageLensConfigurationInput(v *PutStorageLensConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7865,10 +7597,8 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
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
@@ -7924,7 +7654,6 @@ func (m *awsRestxml_serializeOpPutStorageLensConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutStorageLensConfigurationTaggingInput(v *PutStorageLensConfigurationTaggingInput, encoder *httpbinding.Encoder) error {
@@ -7977,10 +7706,8 @@ func (m *awsRestxml_serializeOpSubmitMultiRegionAccess
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
@@ -8036,7 +7763,6 @@ func (m *awsRestxml_serializeOpSubmitMultiRegionAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsSubmitMultiRegionAccessPointRoutesInput(v *SubmitMultiRegionAccessPointRoutesInput, encoder *httpbinding.Encoder) error {
@@ -8089,10 +7815,8 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
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
@@ -8148,7 +7872,6 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -8201,10 +7924,8 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
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
@@ -8242,7 +7963,6 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -8283,10 +8003,8 @@ func (m *awsRestxml_serializeOpUpdateAccessGrantsLocat
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
@@ -8342,7 +8060,6 @@ func (m *awsRestxml_serializeOpUpdateAccessGrantsLocat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateAccessGrantsLocationInput(v *UpdateAccessGrantsLocationInput, encoder *httpbinding.Encoder) error {
@@ -8393,10 +8110,8 @@ func (m *awsRestxml_serializeOpUpdateJobPriority) Hand
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
@@ -8434,7 +8149,6 @@ func (m *awsRestxml_serializeOpUpdateJobPriority) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateJobPriorityInput(v *UpdateJobPriorityInput, encoder *httpbinding.Encoder) error {
@@ -8473,10 +8187,8 @@ func (m *awsRestxml_serializeOpUpdateJobStatus) Handle
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
@@ -8514,7 +8226,6 @@ func (m *awsRestxml_serializeOpUpdateJobStatus) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateJobStatusInput(v *UpdateJobStatusInput, encoder *httpbinding.Encoder) error {
@@ -8557,10 +8268,8 @@ func (m *awsRestxml_serializeOpUpdateStorageLensGroup)
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
@@ -8616,7 +8325,6 @@ func (m *awsRestxml_serializeOpUpdateStorageLensGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateStorageLensGroupInput(v *UpdateStorageLensGroupInput, encoder *httpbinding.Encoder) error {
