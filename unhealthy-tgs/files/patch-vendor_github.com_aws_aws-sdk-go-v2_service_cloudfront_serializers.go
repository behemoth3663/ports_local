--- vendor/github.com/aws/aws-sdk-go-v2/service/cloudfront/serializers.go.orig	2026-11-21 14:46:26 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cloudfront/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyxml "github.com/aws/smithy-go/encoding/xml"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -26,10 +25,8 @@ func (m *awsRestxml_serializeOpAssociateAlias) HandleS
 func (m *awsRestxml_serializeOpAssociateAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -67,7 +64,6 @@ func (m *awsRestxml_serializeOpAssociateAlias) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsAssociateAliasInput(v *AssociateAliasInput, encoder *httpbinding.Encoder) error {
@@ -101,10 +97,8 @@ func (m *awsRestxml_serializeOpAssociateDistributionTe
 func (m *awsRestxml_serializeOpAssociateDistributionTenantWebACL) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -160,7 +154,6 @@ func (m *awsRestxml_serializeOpAssociateDistributionTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsAssociateDistributionTenantWebACLInput(v *AssociateDistributionTenantWebACLInput, encoder *httpbinding.Encoder) error {
@@ -211,10 +204,8 @@ func (m *awsRestxml_serializeOpAssociateDistributionWe
 func (m *awsRestxml_serializeOpAssociateDistributionWebACL) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -270,7 +261,6 @@ func (m *awsRestxml_serializeOpAssociateDistributionWe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsAssociateDistributionWebACLInput(v *AssociateDistributionWebACLInput, encoder *httpbinding.Encoder) error {
@@ -321,10 +311,8 @@ func (m *awsRestxml_serializeOpCopyDistribution) Handl
 func (m *awsRestxml_serializeOpCopyDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -380,7 +368,6 @@ func (m *awsRestxml_serializeOpCopyDistribution) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCopyDistributionInput(v *CopyDistributionInput, encoder *httpbinding.Encoder) error {
@@ -447,10 +434,8 @@ func (m *awsRestxml_serializeOpCreateAnycastIpList) Ha
 func (m *awsRestxml_serializeOpCreateAnycastIpList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -502,7 +487,6 @@ func (m *awsRestxml_serializeOpCreateAnycastIpList) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateAnycastIpListInput(v *CreateAnycastIpListInput, encoder *httpbinding.Encoder) error {
@@ -587,10 +571,8 @@ func (m *awsRestxml_serializeOpCreateCachePolicy) Hand
 func (m *awsRestxml_serializeOpCreateCachePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -648,7 +630,6 @@ func (m *awsRestxml_serializeOpCreateCachePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateCachePolicyInput(v *CreateCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -669,10 +650,8 @@ func (m *awsRestxml_serializeOpCreateCloudFrontOriginA
 func (m *awsRestxml_serializeOpCreateCloudFrontOriginAccessIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -730,7 +709,6 @@ func (m *awsRestxml_serializeOpCreateCloudFrontOriginA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateCloudFrontOriginAccessIdentityInput(v *CreateCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -751,10 +729,8 @@ func (m *awsRestxml_serializeOpCreateConnectionFunctio
 func (m *awsRestxml_serializeOpCreateConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -806,7 +782,6 @@ func (m *awsRestxml_serializeOpCreateConnectionFunctio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateConnectionFunctionInput(v *CreateConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -880,10 +855,8 @@ func (m *awsRestxml_serializeOpCreateConnectionGroup) 
 func (m *awsRestxml_serializeOpCreateConnectionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -935,7 +908,6 @@ func (m *awsRestxml_serializeOpCreateConnectionGroup) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateConnectionGroupInput(v *CreateConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -1018,10 +990,8 @@ func (m *awsRestxml_serializeOpCreateContinuousDeploym
 func (m *awsRestxml_serializeOpCreateContinuousDeploymentPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1079,7 +1049,6 @@ func (m *awsRestxml_serializeOpCreateContinuousDeploym
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateContinuousDeploymentPolicyInput(v *CreateContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1100,10 +1069,8 @@ func (m *awsRestxml_serializeOpCreateDistribution) Han
 func (m *awsRestxml_serializeOpCreateDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1161,7 +1128,6 @@ func (m *awsRestxml_serializeOpCreateDistribution) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateDistributionInput(v *CreateDistributionInput, encoder *httpbinding.Encoder) error {
@@ -1182,10 +1148,8 @@ func (m *awsRestxml_serializeOpCreateDistributionTenan
 func (m *awsRestxml_serializeOpCreateDistributionTenant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1237,7 +1201,6 @@ func (m *awsRestxml_serializeOpCreateDistributionTenan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateDistributionTenantInput(v *CreateDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -1372,10 +1335,8 @@ func (m *awsRestxml_serializeOpCreateDistributionWithT
 func (m *awsRestxml_serializeOpCreateDistributionWithTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1433,7 +1394,6 @@ func (m *awsRestxml_serializeOpCreateDistributionWithT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateDistributionWithTagsInput(v *CreateDistributionWithTagsInput, encoder *httpbinding.Encoder) error {
@@ -1454,10 +1414,8 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
 func (m *awsRestxml_serializeOpCreateFieldLevelEncryptionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1515,7 +1473,6 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateFieldLevelEncryptionConfigInput(v *CreateFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -1536,10 +1493,8 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
 func (m *awsRestxml_serializeOpCreateFieldLevelEncryptionProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1597,7 +1552,6 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateFieldLevelEncryptionProfileInput(v *CreateFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -1618,10 +1572,8 @@ func (m *awsRestxml_serializeOpCreateFunction) HandleS
 func (m *awsRestxml_serializeOpCreateFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1673,7 +1625,6 @@ func (m *awsRestxml_serializeOpCreateFunction) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateFunctionInput(v *CreateFunctionInput, encoder *httpbinding.Encoder) error {
@@ -1747,10 +1698,8 @@ func (m *awsRestxml_serializeOpCreateInvalidation) Han
 func (m *awsRestxml_serializeOpCreateInvalidation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1812,7 +1761,6 @@ func (m *awsRestxml_serializeOpCreateInvalidation) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateInvalidationInput(v *CreateInvalidationInput, encoder *httpbinding.Encoder) error {
@@ -1842,10 +1790,8 @@ func (m *awsRestxml_serializeOpCreateInvalidationForDi
 func (m *awsRestxml_serializeOpCreateInvalidationForDistributionTenant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1907,7 +1853,6 @@ func (m *awsRestxml_serializeOpCreateInvalidationForDi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateInvalidationForDistributionTenantInput(v *CreateInvalidationForDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -1937,10 +1882,8 @@ func (m *awsRestxml_serializeOpCreateKeyGroup) HandleS
 func (m *awsRestxml_serializeOpCreateKeyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1998,7 +1941,6 @@ func (m *awsRestxml_serializeOpCreateKeyGroup) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateKeyGroupInput(v *CreateKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -2019,10 +1961,8 @@ func (m *awsRestxml_serializeOpCreateKeyValueStore) Ha
 func (m *awsRestxml_serializeOpCreateKeyValueStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2074,7 +2014,6 @@ func (m *awsRestxml_serializeOpCreateKeyValueStore) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateKeyValueStoreInput(v *CreateKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -2148,10 +2087,8 @@ func (m *awsRestxml_serializeOpCreateMonitoringSubscri
 func (m *awsRestxml_serializeOpCreateMonitoringSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2213,7 +2150,6 @@ func (m *awsRestxml_serializeOpCreateMonitoringSubscri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateMonitoringSubscriptionInput(v *CreateMonitoringSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -2243,10 +2179,8 @@ func (m *awsRestxml_serializeOpCreateOriginAccessContr
 func (m *awsRestxml_serializeOpCreateOriginAccessControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2304,7 +2238,6 @@ func (m *awsRestxml_serializeOpCreateOriginAccessContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateOriginAccessControlInput(v *CreateOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -2325,10 +2258,8 @@ func (m *awsRestxml_serializeOpCreateOriginRequestPoli
 func (m *awsRestxml_serializeOpCreateOriginRequestPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2386,7 +2317,6 @@ func (m *awsRestxml_serializeOpCreateOriginRequestPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateOriginRequestPolicyInput(v *CreateOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2407,10 +2337,8 @@ func (m *awsRestxml_serializeOpCreatePublicKey) Handle
 func (m *awsRestxml_serializeOpCreatePublicKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2468,7 +2396,6 @@ func (m *awsRestxml_serializeOpCreatePublicKey) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreatePublicKeyInput(v *CreatePublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -2489,10 +2416,8 @@ func (m *awsRestxml_serializeOpCreateRealtimeLogConfig
 func (m *awsRestxml_serializeOpCreateRealtimeLogConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2544,7 +2469,6 @@ func (m *awsRestxml_serializeOpCreateRealtimeLogConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateRealtimeLogConfigInput(v *CreateRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -2618,10 +2542,8 @@ func (m *awsRestxml_serializeOpCreateResponseHeadersPo
 func (m *awsRestxml_serializeOpCreateResponseHeadersPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2679,7 +2601,6 @@ func (m *awsRestxml_serializeOpCreateResponseHeadersPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateResponseHeadersPolicyInput(v *CreateResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2700,10 +2621,8 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
 func (m *awsRestxml_serializeOpCreateStreamingDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2761,7 +2680,6 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateStreamingDistributionInput(v *CreateStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -2782,10 +2700,8 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
 func (m *awsRestxml_serializeOpCreateStreamingDistributionWithTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2843,7 +2759,6 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateStreamingDistributionWithTagsInput(v *CreateStreamingDistributionWithTagsInput, encoder *httpbinding.Encoder) error {
@@ -2864,10 +2779,8 @@ func (m *awsRestxml_serializeOpCreateTrustStore) Handl
 func (m *awsRestxml_serializeOpCreateTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2919,7 +2832,6 @@ func (m *awsRestxml_serializeOpCreateTrustStore) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateTrustStoreInput(v *CreateTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -2993,10 +2905,8 @@ func (m *awsRestxml_serializeOpCreateVpcOrigin) Handle
 func (m *awsRestxml_serializeOpCreateVpcOrigin) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3048,7 +2958,6 @@ func (m *awsRestxml_serializeOpCreateVpcOrigin) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateVpcOriginInput(v *CreateVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -3100,10 +3009,8 @@ func (m *awsRestxml_serializeOpDeleteAnycastIpList) Ha
 func (m *awsRestxml_serializeOpDeleteAnycastIpList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3141,7 +3048,6 @@ func (m *awsRestxml_serializeOpDeleteAnycastIpList) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAnycastIpListInput(v *DeleteAnycastIpListInput, encoder *httpbinding.Encoder) error {
@@ -3176,10 +3082,8 @@ func (m *awsRestxml_serializeOpDeleteCachePolicy) Hand
 func (m *awsRestxml_serializeOpDeleteCachePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3217,7 +3121,6 @@ func (m *awsRestxml_serializeOpDeleteCachePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteCachePolicyInput(v *DeleteCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -3252,10 +3155,8 @@ func (m *awsRestxml_serializeOpDeleteCloudFrontOriginA
 func (m *awsRestxml_serializeOpDeleteCloudFrontOriginAccessIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3293,7 +3194,6 @@ func (m *awsRestxml_serializeOpDeleteCloudFrontOriginA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteCloudFrontOriginAccessIdentityInput(v *DeleteCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -3328,10 +3228,8 @@ func (m *awsRestxml_serializeOpDeleteConnectionFunctio
 func (m *awsRestxml_serializeOpDeleteConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3369,7 +3267,6 @@ func (m *awsRestxml_serializeOpDeleteConnectionFunctio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteConnectionFunctionInput(v *DeleteConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -3404,10 +3301,8 @@ func (m *awsRestxml_serializeOpDeleteConnectionGroup) 
 func (m *awsRestxml_serializeOpDeleteConnectionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3445,7 +3340,6 @@ func (m *awsRestxml_serializeOpDeleteConnectionGroup) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteConnectionGroupInput(v *DeleteConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -3480,10 +3374,8 @@ func (m *awsRestxml_serializeOpDeleteContinuousDeploym
 func (m *awsRestxml_serializeOpDeleteContinuousDeploymentPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3521,7 +3413,6 @@ func (m *awsRestxml_serializeOpDeleteContinuousDeploym
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteContinuousDeploymentPolicyInput(v *DeleteContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3556,10 +3447,8 @@ func (m *awsRestxml_serializeOpDeleteDistribution) Han
 func (m *awsRestxml_serializeOpDeleteDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3597,7 +3486,6 @@ func (m *awsRestxml_serializeOpDeleteDistribution) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteDistributionInput(v *DeleteDistributionInput, encoder *httpbinding.Encoder) error {
@@ -3632,10 +3520,8 @@ func (m *awsRestxml_serializeOpDeleteDistributionTenan
 func (m *awsRestxml_serializeOpDeleteDistributionTenant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3673,7 +3559,6 @@ func (m *awsRestxml_serializeOpDeleteDistributionTenan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteDistributionTenantInput(v *DeleteDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -3708,10 +3593,8 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
 func (m *awsRestxml_serializeOpDeleteFieldLevelEncryptionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3749,7 +3632,6 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteFieldLevelEncryptionConfigInput(v *DeleteFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -3784,10 +3666,8 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
 func (m *awsRestxml_serializeOpDeleteFieldLevelEncryptionProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3825,7 +3705,6 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteFieldLevelEncryptionProfileInput(v *DeleteFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -3860,10 +3739,8 @@ func (m *awsRestxml_serializeOpDeleteFunction) HandleS
 func (m *awsRestxml_serializeOpDeleteFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3901,7 +3778,6 @@ func (m *awsRestxml_serializeOpDeleteFunction) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteFunctionInput(v *DeleteFunctionInput, encoder *httpbinding.Encoder) error {
@@ -3936,10 +3812,8 @@ func (m *awsRestxml_serializeOpDeleteKeyGroup) HandleS
 func (m *awsRestxml_serializeOpDeleteKeyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3977,7 +3851,6 @@ func (m *awsRestxml_serializeOpDeleteKeyGroup) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteKeyGroupInput(v *DeleteKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -4012,10 +3885,8 @@ func (m *awsRestxml_serializeOpDeleteKeyValueStore) Ha
 func (m *awsRestxml_serializeOpDeleteKeyValueStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4053,7 +3924,6 @@ func (m *awsRestxml_serializeOpDeleteKeyValueStore) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteKeyValueStoreInput(v *DeleteKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -4088,10 +3958,8 @@ func (m *awsRestxml_serializeOpDeleteMonitoringSubscri
 func (m *awsRestxml_serializeOpDeleteMonitoringSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4129,7 +3997,6 @@ func (m *awsRestxml_serializeOpDeleteMonitoringSubscri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteMonitoringSubscriptionInput(v *DeleteMonitoringSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -4159,10 +4026,8 @@ func (m *awsRestxml_serializeOpDeleteOriginAccessContr
 func (m *awsRestxml_serializeOpDeleteOriginAccessControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4200,7 +4065,6 @@ func (m *awsRestxml_serializeOpDeleteOriginAccessContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteOriginAccessControlInput(v *DeleteOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -4235,10 +4099,8 @@ func (m *awsRestxml_serializeOpDeleteOriginRequestPoli
 func (m *awsRestxml_serializeOpDeleteOriginRequestPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4276,7 +4138,6 @@ func (m *awsRestxml_serializeOpDeleteOriginRequestPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteOriginRequestPolicyInput(v *DeleteOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4311,10 +4172,8 @@ func (m *awsRestxml_serializeOpDeletePublicKey) Handle
 func (m *awsRestxml_serializeOpDeletePublicKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4352,7 +4211,6 @@ func (m *awsRestxml_serializeOpDeletePublicKey) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeletePublicKeyInput(v *DeletePublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -4387,10 +4245,8 @@ func (m *awsRestxml_serializeOpDeleteRealtimeLogConfig
 func (m *awsRestxml_serializeOpDeleteRealtimeLogConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4442,7 +4298,6 @@ func (m *awsRestxml_serializeOpDeleteRealtimeLogConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteRealtimeLogConfigInput(v *DeleteRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -4490,10 +4345,8 @@ func (m *awsRestxml_serializeOpDeleteResourcePolicy) H
 func (m *awsRestxml_serializeOpDeleteResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4545,7 +4398,6 @@ func (m *awsRestxml_serializeOpDeleteResourcePolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteResourcePolicyInput(v *DeleteResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -4582,10 +4434,8 @@ func (m *awsRestxml_serializeOpDeleteResponseHeadersPo
 func (m *awsRestxml_serializeOpDeleteResponseHeadersPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4623,7 +4473,6 @@ func (m *awsRestxml_serializeOpDeleteResponseHeadersPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteResponseHeadersPolicyInput(v *DeleteResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4658,10 +4507,8 @@ func (m *awsRestxml_serializeOpDeleteStreamingDistribu
 func (m *awsRestxml_serializeOpDeleteStreamingDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4699,7 +4546,6 @@ func (m *awsRestxml_serializeOpDeleteStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStreamingDistributionInput(v *DeleteStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -4734,10 +4580,8 @@ func (m *awsRestxml_serializeOpDeleteTrustStore) Handl
 func (m *awsRestxml_serializeOpDeleteTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4775,7 +4619,6 @@ func (m *awsRestxml_serializeOpDeleteTrustStore) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteTrustStoreInput(v *DeleteTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -4810,10 +4653,8 @@ func (m *awsRestxml_serializeOpDeleteVpcOrigin) Handle
 func (m *awsRestxml_serializeOpDeleteVpcOrigin) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4851,7 +4692,6 @@ func (m *awsRestxml_serializeOpDeleteVpcOrigin) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteVpcOriginInput(v *DeleteVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -4886,10 +4726,8 @@ func (m *awsRestxml_serializeOpDescribeConnectionFunct
 func (m *awsRestxml_serializeOpDescribeConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4927,7 +4765,6 @@ func (m *awsRestxml_serializeOpDescribeConnectionFunct
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeConnectionFunctionInput(v *DescribeConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -4961,10 +4798,8 @@ func (m *awsRestxml_serializeOpDescribeFunction) Handl
 func (m *awsRestxml_serializeOpDescribeFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5002,7 +4837,6 @@ func (m *awsRestxml_serializeOpDescribeFunction) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeFunctionInput(v *DescribeFunctionInput, encoder *httpbinding.Encoder) error {
@@ -5036,10 +4870,8 @@ func (m *awsRestxml_serializeOpDescribeKeyValueStore) 
 func (m *awsRestxml_serializeOpDescribeKeyValueStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5077,7 +4909,6 @@ func (m *awsRestxml_serializeOpDescribeKeyValueStore) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeKeyValueStoreInput(v *DescribeKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -5107,10 +4938,8 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
 func (m *awsRestxml_serializeOpDisassociateDistributionTenantWebACL) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5148,7 +4977,6 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDisassociateDistributionTenantWebACLInput(v *DisassociateDistributionTenantWebACLInput, encoder *httpbinding.Encoder) error {
@@ -5183,10 +5011,8 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
 func (m *awsRestxml_serializeOpDisassociateDistributionWebACL) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5224,7 +5050,6 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDisassociateDistributionWebACLInput(v *DisassociateDistributionWebACLInput, encoder *httpbinding.Encoder) error {
@@ -5259,10 +5084,8 @@ func (m *awsRestxml_serializeOpGetAnycastIpList) Handl
 func (m *awsRestxml_serializeOpGetAnycastIpList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5300,7 +5123,6 @@ func (m *awsRestxml_serializeOpGetAnycastIpList) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAnycastIpListInput(v *GetAnycastIpListInput, encoder *httpbinding.Encoder) error {
@@ -5330,10 +5152,8 @@ func (m *awsRestxml_serializeOpGetCachePolicy) HandleS
 func (m *awsRestxml_serializeOpGetCachePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5371,7 +5191,6 @@ func (m *awsRestxml_serializeOpGetCachePolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCachePolicyInput(v *GetCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -5401,10 +5220,8 @@ func (m *awsRestxml_serializeOpGetCachePolicyConfig) H
 func (m *awsRestxml_serializeOpGetCachePolicyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5442,7 +5259,6 @@ func (m *awsRestxml_serializeOpGetCachePolicyConfig) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCachePolicyConfigInput(v *GetCachePolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -5472,10 +5288,8 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
 func (m *awsRestxml_serializeOpGetCloudFrontOriginAccessIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5513,7 +5327,6 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCloudFrontOriginAccessIdentityInput(v *GetCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -5543,10 +5356,8 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
 func (m *awsRestxml_serializeOpGetCloudFrontOriginAccessIdentityConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5584,7 +5395,6 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCloudFrontOriginAccessIdentityConfigInput(v *GetCloudFrontOriginAccessIdentityConfigInput, encoder *httpbinding.Encoder) error {
@@ -5614,10 +5424,8 @@ func (m *awsRestxml_serializeOpGetConnectionFunction) 
 func (m *awsRestxml_serializeOpGetConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5655,7 +5463,6 @@ func (m *awsRestxml_serializeOpGetConnectionFunction) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetConnectionFunctionInput(v *GetConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -5689,10 +5496,8 @@ func (m *awsRestxml_serializeOpGetConnectionGroup) Han
 func (m *awsRestxml_serializeOpGetConnectionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5730,7 +5535,6 @@ func (m *awsRestxml_serializeOpGetConnectionGroup) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetConnectionGroupInput(v *GetConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -5760,10 +5564,8 @@ func (m *awsRestxml_serializeOpGetConnectionGroupByRou
 func (m *awsRestxml_serializeOpGetConnectionGroupByRoutingEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5801,7 +5603,6 @@ func (m *awsRestxml_serializeOpGetConnectionGroupByRou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetConnectionGroupByRoutingEndpointInput(v *GetConnectionGroupByRoutingEndpointInput, encoder *httpbinding.Encoder) error {
@@ -5826,10 +5627,8 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
 func (m *awsRestxml_serializeOpGetContinuousDeploymentPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5867,7 +5666,6 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetContinuousDeploymentPolicyInput(v *GetContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -5897,10 +5695,8 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
 func (m *awsRestxml_serializeOpGetContinuousDeploymentPolicyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5938,7 +5734,6 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetContinuousDeploymentPolicyConfigInput(v *GetContinuousDeploymentPolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -5968,10 +5763,8 @@ func (m *awsRestxml_serializeOpGetDistribution) Handle
 func (m *awsRestxml_serializeOpGetDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6009,7 +5802,6 @@ func (m *awsRestxml_serializeOpGetDistribution) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionInput(v *GetDistributionInput, encoder *httpbinding.Encoder) error {
@@ -6039,10 +5831,8 @@ func (m *awsRestxml_serializeOpGetDistributionConfig) 
 func (m *awsRestxml_serializeOpGetDistributionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6080,7 +5870,6 @@ func (m *awsRestxml_serializeOpGetDistributionConfig) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionConfigInput(v *GetDistributionConfigInput, encoder *httpbinding.Encoder) error {
@@ -6110,10 +5899,8 @@ func (m *awsRestxml_serializeOpGetDistributionTenant) 
 func (m *awsRestxml_serializeOpGetDistributionTenant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6151,7 +5938,6 @@ func (m *awsRestxml_serializeOpGetDistributionTenant) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionTenantInput(v *GetDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -6181,10 +5967,8 @@ func (m *awsRestxml_serializeOpGetDistributionTenantBy
 func (m *awsRestxml_serializeOpGetDistributionTenantByDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6222,7 +6006,6 @@ func (m *awsRestxml_serializeOpGetDistributionTenantBy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionTenantByDomainInput(v *GetDistributionTenantByDomainInput, encoder *httpbinding.Encoder) error {
@@ -6247,10 +6030,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 func (m *awsRestxml_serializeOpGetFieldLevelEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6288,7 +6069,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionInput(v *GetFieldLevelEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -6318,10 +6098,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 func (m *awsRestxml_serializeOpGetFieldLevelEncryptionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6359,7 +6137,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionConfigInput(v *GetFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -6389,10 +6166,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 func (m *awsRestxml_serializeOpGetFieldLevelEncryptionProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6430,7 +6205,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionProfileInput(v *GetFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -6460,10 +6234,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 func (m *awsRestxml_serializeOpGetFieldLevelEncryptionProfileConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6501,7 +6273,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionProfileConfigInput(v *GetFieldLevelEncryptionProfileConfigInput, encoder *httpbinding.Encoder) error {
@@ -6531,10 +6302,8 @@ func (m *awsRestxml_serializeOpGetFunction) HandleSeri
 func (m *awsRestxml_serializeOpGetFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6572,7 +6341,6 @@ func (m *awsRestxml_serializeOpGetFunction) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFunctionInput(v *GetFunctionInput, encoder *httpbinding.Encoder) error {
@@ -6606,10 +6374,8 @@ func (m *awsRestxml_serializeOpGetInvalidation) Handle
 func (m *awsRestxml_serializeOpGetInvalidation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6647,7 +6413,6 @@ func (m *awsRestxml_serializeOpGetInvalidation) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetInvalidationInput(v *GetInvalidationInput, encoder *httpbinding.Encoder) error {
@@ -6686,10 +6451,8 @@ func (m *awsRestxml_serializeOpGetInvalidationForDistr
 func (m *awsRestxml_serializeOpGetInvalidationForDistributionTenant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6727,7 +6490,6 @@ func (m *awsRestxml_serializeOpGetInvalidationForDistr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetInvalidationForDistributionTenantInput(v *GetInvalidationForDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -6766,10 +6528,8 @@ func (m *awsRestxml_serializeOpGetKeyGroup) HandleSeri
 func (m *awsRestxml_serializeOpGetKeyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6807,7 +6567,6 @@ func (m *awsRestxml_serializeOpGetKeyGroup) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetKeyGroupInput(v *GetKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -6837,10 +6596,8 @@ func (m *awsRestxml_serializeOpGetKeyGroupConfig) Hand
 func (m *awsRestxml_serializeOpGetKeyGroupConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6878,7 +6635,6 @@ func (m *awsRestxml_serializeOpGetKeyGroupConfig) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetKeyGroupConfigInput(v *GetKeyGroupConfigInput, encoder *httpbinding.Encoder) error {
@@ -6908,10 +6664,8 @@ func (m *awsRestxml_serializeOpGetManagedCertificateDe
 func (m *awsRestxml_serializeOpGetManagedCertificateDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6949,7 +6703,6 @@ func (m *awsRestxml_serializeOpGetManagedCertificateDe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetManagedCertificateDetailsInput(v *GetManagedCertificateDetailsInput, encoder *httpbinding.Encoder) error {
@@ -6979,10 +6732,8 @@ func (m *awsRestxml_serializeOpGetMonitoringSubscripti
 func (m *awsRestxml_serializeOpGetMonitoringSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7020,7 +6771,6 @@ func (m *awsRestxml_serializeOpGetMonitoringSubscripti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMonitoringSubscriptionInput(v *GetMonitoringSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -7050,10 +6800,8 @@ func (m *awsRestxml_serializeOpGetOriginAccessControl)
 func (m *awsRestxml_serializeOpGetOriginAccessControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7091,7 +6839,6 @@ func (m *awsRestxml_serializeOpGetOriginAccessControl)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginAccessControlInput(v *GetOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -7121,10 +6868,8 @@ func (m *awsRestxml_serializeOpGetOriginAccessControlC
 func (m *awsRestxml_serializeOpGetOriginAccessControlConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7162,7 +6907,6 @@ func (m *awsRestxml_serializeOpGetOriginAccessControlC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginAccessControlConfigInput(v *GetOriginAccessControlConfigInput, encoder *httpbinding.Encoder) error {
@@ -7192,10 +6936,8 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicy)
 func (m *awsRestxml_serializeOpGetOriginRequestPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7233,7 +6975,6 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginRequestPolicyInput(v *GetOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7263,10 +7004,8 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicyC
 func (m *awsRestxml_serializeOpGetOriginRequestPolicyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7304,7 +7043,6 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicyC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginRequestPolicyConfigInput(v *GetOriginRequestPolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -7334,10 +7072,8 @@ func (m *awsRestxml_serializeOpGetPublicKey) HandleSer
 func (m *awsRestxml_serializeOpGetPublicKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7375,7 +7111,6 @@ func (m *awsRestxml_serializeOpGetPublicKey) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicKeyInput(v *GetPublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -7405,10 +7140,8 @@ func (m *awsRestxml_serializeOpGetPublicKeyConfig) Han
 func (m *awsRestxml_serializeOpGetPublicKeyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7446,7 +7179,6 @@ func (m *awsRestxml_serializeOpGetPublicKeyConfig) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicKeyConfigInput(v *GetPublicKeyConfigInput, encoder *httpbinding.Encoder) error {
@@ -7476,10 +7208,8 @@ func (m *awsRestxml_serializeOpGetRealtimeLogConfig) H
 func (m *awsRestxml_serializeOpGetRealtimeLogConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7531,7 +7261,6 @@ func (m *awsRestxml_serializeOpGetRealtimeLogConfig) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetRealtimeLogConfigInput(v *GetRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -7579,10 +7308,8 @@ func (m *awsRestxml_serializeOpGetResourcePolicy) Hand
 func (m *awsRestxml_serializeOpGetResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7634,7 +7361,6 @@ func (m *awsRestxml_serializeOpGetResourcePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetResourcePolicyInput(v *GetResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -7671,10 +7397,8 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
 func (m *awsRestxml_serializeOpGetResponseHeadersPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7712,7 +7436,6 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetResponseHeadersPolicyInput(v *GetResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -7742,10 +7465,8 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
 func (m *awsRestxml_serializeOpGetResponseHeadersPolicyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7783,7 +7504,6 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetResponseHeadersPolicyConfigInput(v *GetResponseHeadersPolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -7813,10 +7533,8 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
 func (m *awsRestxml_serializeOpGetStreamingDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7854,7 +7572,6 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStreamingDistributionInput(v *GetStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -7884,10 +7601,8 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
 func (m *awsRestxml_serializeOpGetStreamingDistributionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7925,7 +7640,6 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStreamingDistributionConfigInput(v *GetStreamingDistributionConfigInput, encoder *httpbinding.Encoder) error {
@@ -7955,10 +7669,8 @@ func (m *awsRestxml_serializeOpGetTrustStore) HandleSe
 func (m *awsRestxml_serializeOpGetTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7996,7 +7708,6 @@ func (m *awsRestxml_serializeOpGetTrustStore) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetTrustStoreInput(v *GetTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -8026,10 +7737,8 @@ func (m *awsRestxml_serializeOpGetVpcOrigin) HandleSer
 func (m *awsRestxml_serializeOpGetVpcOrigin) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8067,7 +7776,6 @@ func (m *awsRestxml_serializeOpGetVpcOrigin) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetVpcOriginInput(v *GetVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -8097,10 +7805,8 @@ func (m *awsRestxml_serializeOpListAnycastIpLists) Han
 func (m *awsRestxml_serializeOpListAnycastIpLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8138,7 +7844,6 @@ func (m *awsRestxml_serializeOpListAnycastIpLists) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAnycastIpListsInput(v *ListAnycastIpListsInput, encoder *httpbinding.Encoder) error {
@@ -8167,10 +7872,8 @@ func (m *awsRestxml_serializeOpListCachePolicies) Hand
 func (m *awsRestxml_serializeOpListCachePolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8208,7 +7911,6 @@ func (m *awsRestxml_serializeOpListCachePolicies) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCachePoliciesInput(v *ListCachePoliciesInput, encoder *httpbinding.Encoder) error {
@@ -8241,10 +7943,8 @@ func (m *awsRestxml_serializeOpListCloudFrontOriginAcc
 func (m *awsRestxml_serializeOpListCloudFrontOriginAccessIdentities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8282,7 +7982,6 @@ func (m *awsRestxml_serializeOpListCloudFrontOriginAcc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCloudFrontOriginAccessIdentitiesInput(v *ListCloudFrontOriginAccessIdentitiesInput, encoder *httpbinding.Encoder) error {
@@ -8311,10 +8010,8 @@ func (m *awsRestxml_serializeOpListConflictingAliases)
 func (m *awsRestxml_serializeOpListConflictingAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8352,7 +8049,6 @@ func (m *awsRestxml_serializeOpListConflictingAliases)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListConflictingAliasesInput(v *ListConflictingAliasesInput, encoder *httpbinding.Encoder) error {
@@ -8389,10 +8085,8 @@ func (m *awsRestxml_serializeOpListConnectionFunctions
 func (m *awsRestxml_serializeOpListConnectionFunctions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8444,7 +8138,6 @@ func (m *awsRestxml_serializeOpListConnectionFunctions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListConnectionFunctionsInput(v *ListConnectionFunctionsInput, encoder *httpbinding.Encoder) error {
@@ -8503,10 +8196,8 @@ func (m *awsRestxml_serializeOpListConnectionGroups) H
 func (m *awsRestxml_serializeOpListConnectionGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8558,7 +8249,6 @@ func (m *awsRestxml_serializeOpListConnectionGroups) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListConnectionGroupsInput(v *ListConnectionGroupsInput, encoder *httpbinding.Encoder) error {
@@ -8619,10 +8309,8 @@ func (m *awsRestxml_serializeOpListContinuousDeploymen
 func (m *awsRestxml_serializeOpListContinuousDeploymentPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8660,7 +8348,6 @@ func (m *awsRestxml_serializeOpListContinuousDeploymen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListContinuousDeploymentPoliciesInput(v *ListContinuousDeploymentPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -8689,10 +8376,8 @@ func (m *awsRestxml_serializeOpListDistributions) Hand
 func (m *awsRestxml_serializeOpListDistributions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8730,7 +8415,6 @@ func (m *awsRestxml_serializeOpListDistributions) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsInput(v *ListDistributionsInput, encoder *httpbinding.Encoder) error {
@@ -8759,10 +8443,8 @@ func (m *awsRestxml_serializeOpListDistributionsByAnyc
 func (m *awsRestxml_serializeOpListDistributionsByAnycastIpListId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8800,7 +8482,6 @@ func (m *awsRestxml_serializeOpListDistributionsByAnyc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByAnycastIpListIdInput(v *ListDistributionsByAnycastIpListIdInput, encoder *httpbinding.Encoder) error {
@@ -8838,10 +8519,8 @@ func (m *awsRestxml_serializeOpListDistributionsByCach
 func (m *awsRestxml_serializeOpListDistributionsByCachePolicyId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8879,7 +8558,6 @@ func (m *awsRestxml_serializeOpListDistributionsByCach
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByCachePolicyIdInput(v *ListDistributionsByCachePolicyIdInput, encoder *httpbinding.Encoder) error {
@@ -8917,10 +8595,8 @@ func (m *awsRestxml_serializeOpListDistributionsByConn
 func (m *awsRestxml_serializeOpListDistributionsByConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8958,7 +8634,6 @@ func (m *awsRestxml_serializeOpListDistributionsByConn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByConnectionFunctionInput(v *ListDistributionsByConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -8991,10 +8666,8 @@ func (m *awsRestxml_serializeOpListDistributionsByConn
 func (m *awsRestxml_serializeOpListDistributionsByConnectionMode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9032,7 +8705,6 @@ func (m *awsRestxml_serializeOpListDistributionsByConn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByConnectionModeInput(v *ListDistributionsByConnectionModeInput, encoder *httpbinding.Encoder) error {
@@ -9070,10 +8742,8 @@ func (m *awsRestxml_serializeOpListDistributionsByKeyG
 func (m *awsRestxml_serializeOpListDistributionsByKeyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9111,7 +8781,6 @@ func (m *awsRestxml_serializeOpListDistributionsByKeyG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByKeyGroupInput(v *ListDistributionsByKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -9149,10 +8818,8 @@ func (m *awsRestxml_serializeOpListDistributionsByOrig
 func (m *awsRestxml_serializeOpListDistributionsByOriginRequestPolicyId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9190,7 +8857,6 @@ func (m *awsRestxml_serializeOpListDistributionsByOrig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByOriginRequestPolicyIdInput(v *ListDistributionsByOriginRequestPolicyIdInput, encoder *httpbinding.Encoder) error {
@@ -9228,10 +8894,8 @@ func (m *awsRestxml_serializeOpListDistributionsByOwne
 func (m *awsRestxml_serializeOpListDistributionsByOwnedResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9269,7 +8933,6 @@ func (m *awsRestxml_serializeOpListDistributionsByOwne
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByOwnedResourceInput(v *ListDistributionsByOwnedResourceInput, encoder *httpbinding.Encoder) error {
@@ -9307,10 +8970,8 @@ func (m *awsRestxml_serializeOpListDistributionsByReal
 func (m *awsRestxml_serializeOpListDistributionsByRealtimeLogConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9362,7 +9023,6 @@ func (m *awsRestxml_serializeOpListDistributionsByReal
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByRealtimeLogConfigInput(v *ListDistributionsByRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -9432,10 +9092,8 @@ func (m *awsRestxml_serializeOpListDistributionsByResp
 func (m *awsRestxml_serializeOpListDistributionsByResponseHeadersPolicyId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9473,7 +9131,6 @@ func (m *awsRestxml_serializeOpListDistributionsByResp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByResponseHeadersPolicyIdInput(v *ListDistributionsByResponseHeadersPolicyIdInput, encoder *httpbinding.Encoder) error {
@@ -9511,10 +9168,8 @@ func (m *awsRestxml_serializeOpListDistributionsByTrus
 func (m *awsRestxml_serializeOpListDistributionsByTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9552,7 +9207,6 @@ func (m *awsRestxml_serializeOpListDistributionsByTrus
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByTrustStoreInput(v *ListDistributionsByTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -9585,10 +9239,8 @@ func (m *awsRestxml_serializeOpListDistributionsByVpcO
 func (m *awsRestxml_serializeOpListDistributionsByVpcOriginId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9626,7 +9278,6 @@ func (m *awsRestxml_serializeOpListDistributionsByVpcO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByVpcOriginIdInput(v *ListDistributionsByVpcOriginIdInput, encoder *httpbinding.Encoder) error {
@@ -9664,10 +9315,8 @@ func (m *awsRestxml_serializeOpListDistributionsByWebA
 func (m *awsRestxml_serializeOpListDistributionsByWebACLId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9705,7 +9354,6 @@ func (m *awsRestxml_serializeOpListDistributionsByWebA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByWebACLIdInput(v *ListDistributionsByWebACLIdInput, encoder *httpbinding.Encoder) error {
@@ -9743,10 +9391,8 @@ func (m *awsRestxml_serializeOpListDistributionTenants
 func (m *awsRestxml_serializeOpListDistributionTenants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9798,7 +9444,6 @@ func (m *awsRestxml_serializeOpListDistributionTenants
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionTenantsInput(v *ListDistributionTenantsInput, encoder *httpbinding.Encoder) error {
@@ -9859,10 +9504,8 @@ func (m *awsRestxml_serializeOpListDistributionTenants
 func (m *awsRestxml_serializeOpListDistributionTenantsByCustomization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9914,7 +9557,6 @@ func (m *awsRestxml_serializeOpListDistributionTenants
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionTenantsByCustomizationInput(v *ListDistributionTenantsByCustomizationInput, encoder *httpbinding.Encoder) error {
@@ -9984,10 +9626,8 @@ func (m *awsRestxml_serializeOpListDomainConflicts) Ha
 func (m *awsRestxml_serializeOpListDomainConflicts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10039,7 +9679,6 @@ func (m *awsRestxml_serializeOpListDomainConflicts) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDomainConflictsInput(v *ListDomainConflictsInput, encoder *httpbinding.Encoder) error {
@@ -10111,10 +9750,8 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
 func (m *awsRestxml_serializeOpListFieldLevelEncryptionConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10152,7 +9789,6 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListFieldLevelEncryptionConfigsInput(v *ListFieldLevelEncryptionConfigsInput, encoder *httpbinding.Encoder) error {
@@ -10181,10 +9817,8 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
 func (m *awsRestxml_serializeOpListFieldLevelEncryptionProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10222,7 +9856,6 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListFieldLevelEncryptionProfilesInput(v *ListFieldLevelEncryptionProfilesInput, encoder *httpbinding.Encoder) error {
@@ -10251,10 +9884,8 @@ func (m *awsRestxml_serializeOpListFunctions) HandleSe
 func (m *awsRestxml_serializeOpListFunctions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10292,7 +9923,6 @@ func (m *awsRestxml_serializeOpListFunctions) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListFunctionsInput(v *ListFunctionsInput, encoder *httpbinding.Encoder) error {
@@ -10325,10 +9955,8 @@ func (m *awsRestxml_serializeOpListInvalidations) Hand
 func (m *awsRestxml_serializeOpListInvalidations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10366,7 +9994,6 @@ func (m *awsRestxml_serializeOpListInvalidations) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListInvalidationsInput(v *ListInvalidationsInput, encoder *httpbinding.Encoder) error {
@@ -10404,10 +10031,8 @@ func (m *awsRestxml_serializeOpListInvalidationsForDis
 func (m *awsRestxml_serializeOpListInvalidationsForDistributionTenant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10445,7 +10070,6 @@ func (m *awsRestxml_serializeOpListInvalidationsForDis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListInvalidationsForDistributionTenantInput(v *ListInvalidationsForDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -10483,10 +10107,8 @@ func (m *awsRestxml_serializeOpListKeyGroups) HandleSe
 func (m *awsRestxml_serializeOpListKeyGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10524,7 +10146,6 @@ func (m *awsRestxml_serializeOpListKeyGroups) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListKeyGroupsInput(v *ListKeyGroupsInput, encoder *httpbinding.Encoder) error {
@@ -10553,10 +10174,8 @@ func (m *awsRestxml_serializeOpListKeyValueStores) Han
 func (m *awsRestxml_serializeOpListKeyValueStores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10594,7 +10213,6 @@ func (m *awsRestxml_serializeOpListKeyValueStores) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListKeyValueStoresInput(v *ListKeyValueStoresInput, encoder *httpbinding.Encoder) error {
@@ -10627,10 +10245,8 @@ func (m *awsRestxml_serializeOpListOriginAccessControl
 func (m *awsRestxml_serializeOpListOriginAccessControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10668,7 +10284,6 @@ func (m *awsRestxml_serializeOpListOriginAccessControl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListOriginAccessControlsInput(v *ListOriginAccessControlsInput, encoder *httpbinding.Encoder) error {
@@ -10697,10 +10312,8 @@ func (m *awsRestxml_serializeOpListOriginRequestPolici
 func (m *awsRestxml_serializeOpListOriginRequestPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10738,7 +10351,6 @@ func (m *awsRestxml_serializeOpListOriginRequestPolici
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListOriginRequestPoliciesInput(v *ListOriginRequestPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -10771,10 +10383,8 @@ func (m *awsRestxml_serializeOpListPublicKeys) HandleS
 func (m *awsRestxml_serializeOpListPublicKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10812,7 +10422,6 @@ func (m *awsRestxml_serializeOpListPublicKeys) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListPublicKeysInput(v *ListPublicKeysInput, encoder *httpbinding.Encoder) error {
@@ -10841,10 +10450,8 @@ func (m *awsRestxml_serializeOpListRealtimeLogConfigs)
 func (m *awsRestxml_serializeOpListRealtimeLogConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10882,7 +10489,6 @@ func (m *awsRestxml_serializeOpListRealtimeLogConfigs)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListRealtimeLogConfigsInput(v *ListRealtimeLogConfigsInput, encoder *httpbinding.Encoder) error {
@@ -10911,10 +10517,8 @@ func (m *awsRestxml_serializeOpListResponseHeadersPoli
 func (m *awsRestxml_serializeOpListResponseHeadersPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10952,7 +10556,6 @@ func (m *awsRestxml_serializeOpListResponseHeadersPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListResponseHeadersPoliciesInput(v *ListResponseHeadersPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -10985,10 +10588,8 @@ func (m *awsRestxml_serializeOpListStreamingDistributi
 func (m *awsRestxml_serializeOpListStreamingDistributions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11026,7 +10627,6 @@ func (m *awsRestxml_serializeOpListStreamingDistributi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListStreamingDistributionsInput(v *ListStreamingDistributionsInput, encoder *httpbinding.Encoder) error {
@@ -11055,10 +10655,8 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
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
@@ -11096,7 +10694,6 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -11121,10 +10718,8 @@ func (m *awsRestxml_serializeOpListTrustStores) Handle
 func (m *awsRestxml_serializeOpListTrustStores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11176,7 +10771,6 @@ func (m *awsRestxml_serializeOpListTrustStores) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTrustStoresInput(v *ListTrustStoresInput, encoder *httpbinding.Encoder) error {
@@ -11224,10 +10818,8 @@ func (m *awsRestxml_serializeOpListVpcOrigins) HandleS
 func (m *awsRestxml_serializeOpListVpcOrigins) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11265,7 +10857,6 @@ func (m *awsRestxml_serializeOpListVpcOrigins) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListVpcOriginsInput(v *ListVpcOriginsInput, encoder *httpbinding.Encoder) error {
@@ -11294,10 +10885,8 @@ func (m *awsRestxml_serializeOpPublishConnectionFuncti
 func (m *awsRestxml_serializeOpPublishConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11335,7 +10924,6 @@ func (m *awsRestxml_serializeOpPublishConnectionFuncti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPublishConnectionFunctionInput(v *PublishConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -11370,10 +10958,8 @@ func (m *awsRestxml_serializeOpPublishFunction) Handle
 func (m *awsRestxml_serializeOpPublishFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11411,7 +10997,6 @@ func (m *awsRestxml_serializeOpPublishFunction) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPublishFunctionInput(v *PublishFunctionInput, encoder *httpbinding.Encoder) error {
@@ -11446,10 +11031,8 @@ func (m *awsRestxml_serializeOpPutResourcePolicy) Hand
 func (m *awsRestxml_serializeOpPutResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11501,7 +11084,6 @@ func (m *awsRestxml_serializeOpPutResourcePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPutResourcePolicyInput(v *PutResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -11549,10 +11131,8 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
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
@@ -11614,7 +11194,6 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -11639,10 +11218,8 @@ func (m *awsRestxml_serializeOpTestConnectionFunction)
 func (m *awsRestxml_serializeOpTestConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11698,7 +11275,6 @@ func (m *awsRestxml_serializeOpTestConnectionFunction)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTestConnectionFunctionInput(v *TestConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -11760,10 +11336,8 @@ func (m *awsRestxml_serializeOpTestFunction) HandleSer
 func (m *awsRestxml_serializeOpTestFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11819,7 +11393,6 @@ func (m *awsRestxml_serializeOpTestFunction) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTestFunctionInput(v *TestFunctionInput, encoder *httpbinding.Encoder) error {
@@ -11881,10 +11454,8 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
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
@@ -11946,7 +11517,6 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -11971,10 +11541,8 @@ func (m *awsRestxml_serializeOpUpdateAnycastIpList) Ha
 func (m *awsRestxml_serializeOpUpdateAnycastIpList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12030,7 +11598,6 @@ func (m *awsRestxml_serializeOpUpdateAnycastIpList) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateAnycastIpListInput(v *UpdateAnycastIpListInput, encoder *httpbinding.Encoder) error {
@@ -12094,10 +11661,8 @@ func (m *awsRestxml_serializeOpUpdateCachePolicy) Hand
 func (m *awsRestxml_serializeOpUpdateCachePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12159,7 +11724,6 @@ func (m *awsRestxml_serializeOpUpdateCachePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateCachePolicyInput(v *UpdateCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -12194,10 +11758,8 @@ func (m *awsRestxml_serializeOpUpdateCloudFrontOriginA
 func (m *awsRestxml_serializeOpUpdateCloudFrontOriginAccessIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12259,7 +11821,6 @@ func (m *awsRestxml_serializeOpUpdateCloudFrontOriginA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateCloudFrontOriginAccessIdentityInput(v *UpdateCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -12294,10 +11855,8 @@ func (m *awsRestxml_serializeOpUpdateConnectionFunctio
 func (m *awsRestxml_serializeOpUpdateConnectionFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12353,7 +11912,6 @@ func (m *awsRestxml_serializeOpUpdateConnectionFunctio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateConnectionFunctionInput(v *UpdateConnectionFunctionInput, encoder *httpbinding.Encoder) error {
@@ -12417,10 +11975,8 @@ func (m *awsRestxml_serializeOpUpdateConnectionGroup) 
 func (m *awsRestxml_serializeOpUpdateConnectionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12476,7 +12032,6 @@ func (m *awsRestxml_serializeOpUpdateConnectionGroup) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateConnectionGroupInput(v *UpdateConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -12549,10 +12104,8 @@ func (m *awsRestxml_serializeOpUpdateContinuousDeploym
 func (m *awsRestxml_serializeOpUpdateContinuousDeploymentPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12614,7 +12167,6 @@ func (m *awsRestxml_serializeOpUpdateContinuousDeploym
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateContinuousDeploymentPolicyInput(v *UpdateContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -12649,10 +12201,8 @@ func (m *awsRestxml_serializeOpUpdateDistribution) Han
 func (m *awsRestxml_serializeOpUpdateDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12714,7 +12264,6 @@ func (m *awsRestxml_serializeOpUpdateDistribution) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDistributionInput(v *UpdateDistributionInput, encoder *httpbinding.Encoder) error {
@@ -12749,10 +12298,8 @@ func (m *awsRestxml_serializeOpUpdateDistributionTenan
 func (m *awsRestxml_serializeOpUpdateDistributionTenant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12808,7 +12355,6 @@ func (m *awsRestxml_serializeOpUpdateDistributionTenan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDistributionTenantInput(v *UpdateDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -12933,10 +12479,8 @@ func (m *awsRestxml_serializeOpUpdateDistributionWithS
 func (m *awsRestxml_serializeOpUpdateDistributionWithStagingConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12974,7 +12518,6 @@ func (m *awsRestxml_serializeOpUpdateDistributionWithS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDistributionWithStagingConfigInput(v *UpdateDistributionWithStagingConfigInput, encoder *httpbinding.Encoder) error {
@@ -13013,10 +12556,8 @@ func (m *awsRestxml_serializeOpUpdateDomainAssociation
 func (m *awsRestxml_serializeOpUpdateDomainAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13072,7 +12613,6 @@ func (m *awsRestxml_serializeOpUpdateDomainAssociation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDomainAssociationInput(v *UpdateDomainAssociationInput, encoder *httpbinding.Encoder) error {
@@ -13127,10 +12667,8 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
 func (m *awsRestxml_serializeOpUpdateFieldLevelEncryptionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13192,7 +12730,6 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateFieldLevelEncryptionConfigInput(v *UpdateFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -13227,10 +12764,8 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
 func (m *awsRestxml_serializeOpUpdateFieldLevelEncryptionProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13292,7 +12827,6 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateFieldLevelEncryptionProfileInput(v *UpdateFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -13327,10 +12861,8 @@ func (m *awsRestxml_serializeOpUpdateFunction) HandleS
 func (m *awsRestxml_serializeOpUpdateFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13386,7 +12918,6 @@ func (m *awsRestxml_serializeOpUpdateFunction) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateFunctionInput(v *UpdateFunctionInput, encoder *httpbinding.Encoder) error {
@@ -13450,10 +12981,8 @@ func (m *awsRestxml_serializeOpUpdateKeyGroup) HandleS
 func (m *awsRestxml_serializeOpUpdateKeyGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13515,7 +13044,6 @@ func (m *awsRestxml_serializeOpUpdateKeyGroup) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateKeyGroupInput(v *UpdateKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -13550,10 +13078,8 @@ func (m *awsRestxml_serializeOpUpdateKeyValueStore) Ha
 func (m *awsRestxml_serializeOpUpdateKeyValueStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13609,7 +13135,6 @@ func (m *awsRestxml_serializeOpUpdateKeyValueStore) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateKeyValueStoreInput(v *UpdateKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -13660,10 +13185,8 @@ func (m *awsRestxml_serializeOpUpdateOriginAccessContr
 func (m *awsRestxml_serializeOpUpdateOriginAccessControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13725,7 +13248,6 @@ func (m *awsRestxml_serializeOpUpdateOriginAccessContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateOriginAccessControlInput(v *UpdateOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -13760,10 +13282,8 @@ func (m *awsRestxml_serializeOpUpdateOriginRequestPoli
 func (m *awsRestxml_serializeOpUpdateOriginRequestPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13825,7 +13345,6 @@ func (m *awsRestxml_serializeOpUpdateOriginRequestPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateOriginRequestPolicyInput(v *UpdateOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -13860,10 +13379,8 @@ func (m *awsRestxml_serializeOpUpdatePublicKey) Handle
 func (m *awsRestxml_serializeOpUpdatePublicKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13925,7 +13442,6 @@ func (m *awsRestxml_serializeOpUpdatePublicKey) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdatePublicKeyInput(v *UpdatePublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -13960,10 +13476,8 @@ func (m *awsRestxml_serializeOpUpdateRealtimeLogConfig
 func (m *awsRestxml_serializeOpUpdateRealtimeLogConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14015,7 +13529,6 @@ func (m *awsRestxml_serializeOpUpdateRealtimeLogConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateRealtimeLogConfigInput(v *UpdateRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -14100,10 +13613,8 @@ func (m *awsRestxml_serializeOpUpdateResponseHeadersPo
 func (m *awsRestxml_serializeOpUpdateResponseHeadersPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14165,7 +13676,6 @@ func (m *awsRestxml_serializeOpUpdateResponseHeadersPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateResponseHeadersPolicyInput(v *UpdateResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -14200,10 +13710,8 @@ func (m *awsRestxml_serializeOpUpdateStreamingDistribu
 func (m *awsRestxml_serializeOpUpdateStreamingDistribution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14265,7 +13773,6 @@ func (m *awsRestxml_serializeOpUpdateStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateStreamingDistributionInput(v *UpdateStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -14300,10 +13807,8 @@ func (m *awsRestxml_serializeOpUpdateTrustStore) Handl
 func (m *awsRestxml_serializeOpUpdateTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14365,7 +13870,6 @@ func (m *awsRestxml_serializeOpUpdateTrustStore) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateTrustStoreInput(v *UpdateTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -14405,10 +13909,8 @@ func (m *awsRestxml_serializeOpUpdateVpcOrigin) Handle
 func (m *awsRestxml_serializeOpUpdateVpcOrigin) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14470,7 +13972,6 @@ func (m *awsRestxml_serializeOpUpdateVpcOrigin) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateVpcOriginInput(v *UpdateVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -14505,10 +14006,8 @@ func (m *awsRestxml_serializeOpVerifyDnsConfiguration)
 func (m *awsRestxml_serializeOpVerifyDnsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14560,7 +14059,6 @@ func (m *awsRestxml_serializeOpVerifyDnsConfiguration)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsVerifyDnsConfigurationInput(v *VerifyDnsConfigurationInput, encoder *httpbinding.Encoder) error {
