--- vendor/github.com/aws/aws-sdk-go-v2/service/cloudfront/serializers.go.orig	2025-07-19 01:28:18 UTC
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
@@ -563,10 +547,8 @@ func (m *awsRestxml_serializeOpCreateCachePolicy) Hand
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
@@ -624,7 +606,6 @@ func (m *awsRestxml_serializeOpCreateCachePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateCachePolicyInput(v *CreateCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -645,10 +626,8 @@ func (m *awsRestxml_serializeOpCreateCloudFrontOriginA
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
@@ -706,7 +685,6 @@ func (m *awsRestxml_serializeOpCreateCloudFrontOriginA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateCloudFrontOriginAccessIdentityInput(v *CreateCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -727,10 +705,8 @@ func (m *awsRestxml_serializeOpCreateConnectionGroup) 
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
@@ -782,7 +758,6 @@ func (m *awsRestxml_serializeOpCreateConnectionGroup) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateConnectionGroupInput(v *CreateConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -865,10 +840,8 @@ func (m *awsRestxml_serializeOpCreateContinuousDeploym
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
@@ -926,7 +899,6 @@ func (m *awsRestxml_serializeOpCreateContinuousDeploym
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateContinuousDeploymentPolicyInput(v *CreateContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -947,10 +919,8 @@ func (m *awsRestxml_serializeOpCreateDistribution) Han
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
@@ -1008,7 +978,6 @@ func (m *awsRestxml_serializeOpCreateDistribution) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateDistributionInput(v *CreateDistributionInput, encoder *httpbinding.Encoder) error {
@@ -1029,10 +998,8 @@ func (m *awsRestxml_serializeOpCreateDistributionTenan
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
@@ -1084,7 +1051,6 @@ func (m *awsRestxml_serializeOpCreateDistributionTenan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateDistributionTenantInput(v *CreateDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -1219,10 +1185,8 @@ func (m *awsRestxml_serializeOpCreateDistributionWithT
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
@@ -1280,7 +1244,6 @@ func (m *awsRestxml_serializeOpCreateDistributionWithT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateDistributionWithTagsInput(v *CreateDistributionWithTagsInput, encoder *httpbinding.Encoder) error {
@@ -1301,10 +1264,8 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
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
@@ -1362,7 +1323,6 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateFieldLevelEncryptionConfigInput(v *CreateFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -1383,10 +1343,8 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
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
@@ -1444,7 +1402,6 @@ func (m *awsRestxml_serializeOpCreateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateFieldLevelEncryptionProfileInput(v *CreateFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -1465,10 +1422,8 @@ func (m *awsRestxml_serializeOpCreateFunction) HandleS
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
@@ -1520,7 +1475,6 @@ func (m *awsRestxml_serializeOpCreateFunction) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateFunctionInput(v *CreateFunctionInput, encoder *httpbinding.Encoder) error {
@@ -1581,10 +1535,8 @@ func (m *awsRestxml_serializeOpCreateInvalidation) Han
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
@@ -1646,7 +1598,6 @@ func (m *awsRestxml_serializeOpCreateInvalidation) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateInvalidationInput(v *CreateInvalidationInput, encoder *httpbinding.Encoder) error {
@@ -1676,10 +1627,8 @@ func (m *awsRestxml_serializeOpCreateInvalidationForDi
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
@@ -1741,7 +1690,6 @@ func (m *awsRestxml_serializeOpCreateInvalidationForDi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateInvalidationForDistributionTenantInput(v *CreateInvalidationForDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -1771,10 +1719,8 @@ func (m *awsRestxml_serializeOpCreateKeyGroup) HandleS
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
@@ -1832,7 +1778,6 @@ func (m *awsRestxml_serializeOpCreateKeyGroup) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateKeyGroupInput(v *CreateKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -1853,10 +1798,8 @@ func (m *awsRestxml_serializeOpCreateKeyValueStore) Ha
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
@@ -1908,7 +1851,6 @@ func (m *awsRestxml_serializeOpCreateKeyValueStore) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateKeyValueStoreInput(v *CreateKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -1969,10 +1911,8 @@ func (m *awsRestxml_serializeOpCreateMonitoringSubscri
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
@@ -2034,7 +1974,6 @@ func (m *awsRestxml_serializeOpCreateMonitoringSubscri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateMonitoringSubscriptionInput(v *CreateMonitoringSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -2064,10 +2003,8 @@ func (m *awsRestxml_serializeOpCreateOriginAccessContr
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
@@ -2125,7 +2062,6 @@ func (m *awsRestxml_serializeOpCreateOriginAccessContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateOriginAccessControlInput(v *CreateOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -2146,10 +2082,8 @@ func (m *awsRestxml_serializeOpCreateOriginRequestPoli
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
@@ -2207,7 +2141,6 @@ func (m *awsRestxml_serializeOpCreateOriginRequestPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateOriginRequestPolicyInput(v *CreateOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2228,10 +2161,8 @@ func (m *awsRestxml_serializeOpCreatePublicKey) Handle
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
@@ -2289,7 +2220,6 @@ func (m *awsRestxml_serializeOpCreatePublicKey) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreatePublicKeyInput(v *CreatePublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -2310,10 +2240,8 @@ func (m *awsRestxml_serializeOpCreateRealtimeLogConfig
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
@@ -2365,7 +2293,6 @@ func (m *awsRestxml_serializeOpCreateRealtimeLogConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateRealtimeLogConfigInput(v *CreateRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -2439,10 +2366,8 @@ func (m *awsRestxml_serializeOpCreateResponseHeadersPo
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
@@ -2500,7 +2425,6 @@ func (m *awsRestxml_serializeOpCreateResponseHeadersPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateResponseHeadersPolicyInput(v *CreateResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2521,10 +2445,8 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
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
@@ -2582,7 +2504,6 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateStreamingDistributionInput(v *CreateStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -2603,10 +2524,8 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
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
@@ -2664,7 +2583,6 @@ func (m *awsRestxml_serializeOpCreateStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateStreamingDistributionWithTagsInput(v *CreateStreamingDistributionWithTagsInput, encoder *httpbinding.Encoder) error {
@@ -2685,10 +2603,8 @@ func (m *awsRestxml_serializeOpCreateVpcOrigin) Handle
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
@@ -2740,7 +2656,6 @@ func (m *awsRestxml_serializeOpCreateVpcOrigin) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateVpcOriginInput(v *CreateVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -2792,10 +2707,8 @@ func (m *awsRestxml_serializeOpDeleteAnycastIpList) Ha
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
@@ -2833,7 +2746,6 @@ func (m *awsRestxml_serializeOpDeleteAnycastIpList) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteAnycastIpListInput(v *DeleteAnycastIpListInput, encoder *httpbinding.Encoder) error {
@@ -2868,10 +2780,8 @@ func (m *awsRestxml_serializeOpDeleteCachePolicy) Hand
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
@@ -2909,7 +2819,6 @@ func (m *awsRestxml_serializeOpDeleteCachePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteCachePolicyInput(v *DeleteCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2944,10 +2853,8 @@ func (m *awsRestxml_serializeOpDeleteCloudFrontOriginA
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
@@ -2985,7 +2892,6 @@ func (m *awsRestxml_serializeOpDeleteCloudFrontOriginA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteCloudFrontOriginAccessIdentityInput(v *DeleteCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -3020,10 +2926,8 @@ func (m *awsRestxml_serializeOpDeleteConnectionGroup) 
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
@@ -3061,7 +2965,6 @@ func (m *awsRestxml_serializeOpDeleteConnectionGroup) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteConnectionGroupInput(v *DeleteConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -3096,10 +2999,8 @@ func (m *awsRestxml_serializeOpDeleteContinuousDeploym
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
@@ -3137,7 +3038,6 @@ func (m *awsRestxml_serializeOpDeleteContinuousDeploym
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteContinuousDeploymentPolicyInput(v *DeleteContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3172,10 +3072,8 @@ func (m *awsRestxml_serializeOpDeleteDistribution) Han
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
@@ -3213,7 +3111,6 @@ func (m *awsRestxml_serializeOpDeleteDistribution) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteDistributionInput(v *DeleteDistributionInput, encoder *httpbinding.Encoder) error {
@@ -3248,10 +3145,8 @@ func (m *awsRestxml_serializeOpDeleteDistributionTenan
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
@@ -3289,7 +3184,6 @@ func (m *awsRestxml_serializeOpDeleteDistributionTenan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteDistributionTenantInput(v *DeleteDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -3324,10 +3218,8 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
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
@@ -3365,7 +3257,6 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteFieldLevelEncryptionConfigInput(v *DeleteFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -3400,10 +3291,8 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
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
@@ -3441,7 +3330,6 @@ func (m *awsRestxml_serializeOpDeleteFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteFieldLevelEncryptionProfileInput(v *DeleteFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -3476,10 +3364,8 @@ func (m *awsRestxml_serializeOpDeleteFunction) HandleS
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
@@ -3517,7 +3403,6 @@ func (m *awsRestxml_serializeOpDeleteFunction) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteFunctionInput(v *DeleteFunctionInput, encoder *httpbinding.Encoder) error {
@@ -3552,10 +3437,8 @@ func (m *awsRestxml_serializeOpDeleteKeyGroup) HandleS
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
@@ -3593,7 +3476,6 @@ func (m *awsRestxml_serializeOpDeleteKeyGroup) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteKeyGroupInput(v *DeleteKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -3628,10 +3510,8 @@ func (m *awsRestxml_serializeOpDeleteKeyValueStore) Ha
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
@@ -3669,7 +3549,6 @@ func (m *awsRestxml_serializeOpDeleteKeyValueStore) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteKeyValueStoreInput(v *DeleteKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -3704,10 +3583,8 @@ func (m *awsRestxml_serializeOpDeleteMonitoringSubscri
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
@@ -3745,7 +3622,6 @@ func (m *awsRestxml_serializeOpDeleteMonitoringSubscri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteMonitoringSubscriptionInput(v *DeleteMonitoringSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -3775,10 +3651,8 @@ func (m *awsRestxml_serializeOpDeleteOriginAccessContr
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
@@ -3816,7 +3690,6 @@ func (m *awsRestxml_serializeOpDeleteOriginAccessContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteOriginAccessControlInput(v *DeleteOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -3851,10 +3724,8 @@ func (m *awsRestxml_serializeOpDeleteOriginRequestPoli
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
@@ -3892,7 +3763,6 @@ func (m *awsRestxml_serializeOpDeleteOriginRequestPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteOriginRequestPolicyInput(v *DeleteOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3927,10 +3797,8 @@ func (m *awsRestxml_serializeOpDeletePublicKey) Handle
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
@@ -3968,7 +3836,6 @@ func (m *awsRestxml_serializeOpDeletePublicKey) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeletePublicKeyInput(v *DeletePublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -4003,10 +3870,8 @@ func (m *awsRestxml_serializeOpDeleteRealtimeLogConfig
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
@@ -4058,7 +3923,6 @@ func (m *awsRestxml_serializeOpDeleteRealtimeLogConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteRealtimeLogConfigInput(v *DeleteRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -4106,10 +3970,8 @@ func (m *awsRestxml_serializeOpDeleteResponseHeadersPo
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
@@ -4147,7 +4009,6 @@ func (m *awsRestxml_serializeOpDeleteResponseHeadersPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteResponseHeadersPolicyInput(v *DeleteResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4182,10 +4043,8 @@ func (m *awsRestxml_serializeOpDeleteStreamingDistribu
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
@@ -4223,7 +4082,6 @@ func (m *awsRestxml_serializeOpDeleteStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteStreamingDistributionInput(v *DeleteStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -4258,10 +4116,8 @@ func (m *awsRestxml_serializeOpDeleteVpcOrigin) Handle
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
@@ -4299,7 +4155,6 @@ func (m *awsRestxml_serializeOpDeleteVpcOrigin) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteVpcOriginInput(v *DeleteVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -4334,10 +4189,8 @@ func (m *awsRestxml_serializeOpDescribeFunction) Handl
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
@@ -4375,7 +4228,6 @@ func (m *awsRestxml_serializeOpDescribeFunction) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeFunctionInput(v *DescribeFunctionInput, encoder *httpbinding.Encoder) error {
@@ -4409,10 +4261,8 @@ func (m *awsRestxml_serializeOpDescribeKeyValueStore) 
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
@@ -4450,7 +4300,6 @@ func (m *awsRestxml_serializeOpDescribeKeyValueStore) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDescribeKeyValueStoreInput(v *DescribeKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -4480,10 +4329,8 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
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
@@ -4521,7 +4368,6 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDisassociateDistributionTenantWebACLInput(v *DisassociateDistributionTenantWebACLInput, encoder *httpbinding.Encoder) error {
@@ -4556,10 +4402,8 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
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
@@ -4597,7 +4441,6 @@ func (m *awsRestxml_serializeOpDisassociateDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDisassociateDistributionWebACLInput(v *DisassociateDistributionWebACLInput, encoder *httpbinding.Encoder) error {
@@ -4632,10 +4475,8 @@ func (m *awsRestxml_serializeOpGetAnycastIpList) Handl
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
@@ -4673,7 +4514,6 @@ func (m *awsRestxml_serializeOpGetAnycastIpList) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAnycastIpListInput(v *GetAnycastIpListInput, encoder *httpbinding.Encoder) error {
@@ -4703,10 +4543,8 @@ func (m *awsRestxml_serializeOpGetCachePolicy) HandleS
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
@@ -4744,7 +4582,6 @@ func (m *awsRestxml_serializeOpGetCachePolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCachePolicyInput(v *GetCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -4774,10 +4611,8 @@ func (m *awsRestxml_serializeOpGetCachePolicyConfig) H
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
@@ -4815,7 +4650,6 @@ func (m *awsRestxml_serializeOpGetCachePolicyConfig) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCachePolicyConfigInput(v *GetCachePolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -4845,10 +4679,8 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
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
@@ -4886,7 +4718,6 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCloudFrontOriginAccessIdentityInput(v *GetCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -4916,10 +4747,8 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
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
@@ -4957,7 +4786,6 @@ func (m *awsRestxml_serializeOpGetCloudFrontOriginAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCloudFrontOriginAccessIdentityConfigInput(v *GetCloudFrontOriginAccessIdentityConfigInput, encoder *httpbinding.Encoder) error {
@@ -4987,10 +4815,8 @@ func (m *awsRestxml_serializeOpGetConnectionGroup) Han
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
@@ -5028,7 +4854,6 @@ func (m *awsRestxml_serializeOpGetConnectionGroup) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetConnectionGroupInput(v *GetConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -5058,10 +4883,8 @@ func (m *awsRestxml_serializeOpGetConnectionGroupByRou
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
@@ -5099,7 +4922,6 @@ func (m *awsRestxml_serializeOpGetConnectionGroupByRou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetConnectionGroupByRoutingEndpointInput(v *GetConnectionGroupByRoutingEndpointInput, encoder *httpbinding.Encoder) error {
@@ -5124,10 +4946,8 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
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
@@ -5165,7 +4985,6 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetContinuousDeploymentPolicyInput(v *GetContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -5195,10 +5014,8 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
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
@@ -5236,7 +5053,6 @@ func (m *awsRestxml_serializeOpGetContinuousDeployment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetContinuousDeploymentPolicyConfigInput(v *GetContinuousDeploymentPolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -5266,10 +5082,8 @@ func (m *awsRestxml_serializeOpGetDistribution) Handle
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
@@ -5307,7 +5121,6 @@ func (m *awsRestxml_serializeOpGetDistribution) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionInput(v *GetDistributionInput, encoder *httpbinding.Encoder) error {
@@ -5337,10 +5150,8 @@ func (m *awsRestxml_serializeOpGetDistributionConfig) 
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
@@ -5378,7 +5189,6 @@ func (m *awsRestxml_serializeOpGetDistributionConfig) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionConfigInput(v *GetDistributionConfigInput, encoder *httpbinding.Encoder) error {
@@ -5408,10 +5218,8 @@ func (m *awsRestxml_serializeOpGetDistributionTenant) 
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
@@ -5449,7 +5257,6 @@ func (m *awsRestxml_serializeOpGetDistributionTenant) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionTenantInput(v *GetDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -5479,10 +5286,8 @@ func (m *awsRestxml_serializeOpGetDistributionTenantBy
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
@@ -5520,7 +5325,6 @@ func (m *awsRestxml_serializeOpGetDistributionTenantBy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDistributionTenantByDomainInput(v *GetDistributionTenantByDomainInput, encoder *httpbinding.Encoder) error {
@@ -5545,10 +5349,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
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
@@ -5586,7 +5388,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionInput(v *GetFieldLevelEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -5616,10 +5417,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
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
@@ -5657,7 +5456,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionConfigInput(v *GetFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -5687,10 +5485,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
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
@@ -5728,7 +5524,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionProfileInput(v *GetFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -5758,10 +5553,8 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
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
@@ -5799,7 +5592,6 @@ func (m *awsRestxml_serializeOpGetFieldLevelEncryption
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFieldLevelEncryptionProfileConfigInput(v *GetFieldLevelEncryptionProfileConfigInput, encoder *httpbinding.Encoder) error {
@@ -5829,10 +5621,8 @@ func (m *awsRestxml_serializeOpGetFunction) HandleSeri
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
@@ -5870,7 +5660,6 @@ func (m *awsRestxml_serializeOpGetFunction) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetFunctionInput(v *GetFunctionInput, encoder *httpbinding.Encoder) error {
@@ -5904,10 +5693,8 @@ func (m *awsRestxml_serializeOpGetInvalidation) Handle
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
@@ -5945,7 +5732,6 @@ func (m *awsRestxml_serializeOpGetInvalidation) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetInvalidationInput(v *GetInvalidationInput, encoder *httpbinding.Encoder) error {
@@ -5984,10 +5770,8 @@ func (m *awsRestxml_serializeOpGetInvalidationForDistr
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
@@ -6025,7 +5809,6 @@ func (m *awsRestxml_serializeOpGetInvalidationForDistr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetInvalidationForDistributionTenantInput(v *GetInvalidationForDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -6064,10 +5847,8 @@ func (m *awsRestxml_serializeOpGetKeyGroup) HandleSeri
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
@@ -6105,7 +5886,6 @@ func (m *awsRestxml_serializeOpGetKeyGroup) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetKeyGroupInput(v *GetKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -6135,10 +5915,8 @@ func (m *awsRestxml_serializeOpGetKeyGroupConfig) Hand
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
@@ -6176,7 +5954,6 @@ func (m *awsRestxml_serializeOpGetKeyGroupConfig) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetKeyGroupConfigInput(v *GetKeyGroupConfigInput, encoder *httpbinding.Encoder) error {
@@ -6206,10 +5983,8 @@ func (m *awsRestxml_serializeOpGetManagedCertificateDe
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
@@ -6247,7 +6022,6 @@ func (m *awsRestxml_serializeOpGetManagedCertificateDe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetManagedCertificateDetailsInput(v *GetManagedCertificateDetailsInput, encoder *httpbinding.Encoder) error {
@@ -6277,10 +6051,8 @@ func (m *awsRestxml_serializeOpGetMonitoringSubscripti
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
@@ -6318,7 +6090,6 @@ func (m *awsRestxml_serializeOpGetMonitoringSubscripti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetMonitoringSubscriptionInput(v *GetMonitoringSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -6348,10 +6119,8 @@ func (m *awsRestxml_serializeOpGetOriginAccessControl)
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
@@ -6389,7 +6158,6 @@ func (m *awsRestxml_serializeOpGetOriginAccessControl)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginAccessControlInput(v *GetOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -6419,10 +6187,8 @@ func (m *awsRestxml_serializeOpGetOriginAccessControlC
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
@@ -6460,7 +6226,6 @@ func (m *awsRestxml_serializeOpGetOriginAccessControlC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginAccessControlConfigInput(v *GetOriginAccessControlConfigInput, encoder *httpbinding.Encoder) error {
@@ -6490,10 +6255,8 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicy)
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
@@ -6531,7 +6294,6 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginRequestPolicyInput(v *GetOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -6561,10 +6323,8 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicyC
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
@@ -6602,7 +6362,6 @@ func (m *awsRestxml_serializeOpGetOriginRequestPolicyC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetOriginRequestPolicyConfigInput(v *GetOriginRequestPolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -6632,10 +6391,8 @@ func (m *awsRestxml_serializeOpGetPublicKey) HandleSer
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
@@ -6673,7 +6430,6 @@ func (m *awsRestxml_serializeOpGetPublicKey) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicKeyInput(v *GetPublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -6703,10 +6459,8 @@ func (m *awsRestxml_serializeOpGetPublicKeyConfig) Han
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
@@ -6744,7 +6498,6 @@ func (m *awsRestxml_serializeOpGetPublicKeyConfig) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetPublicKeyConfigInput(v *GetPublicKeyConfigInput, encoder *httpbinding.Encoder) error {
@@ -6774,10 +6527,8 @@ func (m *awsRestxml_serializeOpGetRealtimeLogConfig) H
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
@@ -6829,7 +6580,6 @@ func (m *awsRestxml_serializeOpGetRealtimeLogConfig) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetRealtimeLogConfigInput(v *GetRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -6877,10 +6627,8 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
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
@@ -6918,7 +6666,6 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetResponseHeadersPolicyInput(v *GetResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -6948,10 +6695,8 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
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
@@ -6989,7 +6734,6 @@ func (m *awsRestxml_serializeOpGetResponseHeadersPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetResponseHeadersPolicyConfigInput(v *GetResponseHeadersPolicyConfigInput, encoder *httpbinding.Encoder) error {
@@ -7019,10 +6763,8 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
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
@@ -7060,7 +6802,6 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStreamingDistributionInput(v *GetStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -7090,10 +6831,8 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
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
@@ -7131,7 +6870,6 @@ func (m *awsRestxml_serializeOpGetStreamingDistributio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetStreamingDistributionConfigInput(v *GetStreamingDistributionConfigInput, encoder *httpbinding.Encoder) error {
@@ -7161,10 +6899,8 @@ func (m *awsRestxml_serializeOpGetVpcOrigin) HandleSer
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
@@ -7202,7 +6938,6 @@ func (m *awsRestxml_serializeOpGetVpcOrigin) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetVpcOriginInput(v *GetVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -7232,10 +6967,8 @@ func (m *awsRestxml_serializeOpListAnycastIpLists) Han
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
@@ -7273,7 +7006,6 @@ func (m *awsRestxml_serializeOpListAnycastIpLists) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListAnycastIpListsInput(v *ListAnycastIpListsInput, encoder *httpbinding.Encoder) error {
@@ -7302,10 +7034,8 @@ func (m *awsRestxml_serializeOpListCachePolicies) Hand
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
@@ -7343,7 +7073,6 @@ func (m *awsRestxml_serializeOpListCachePolicies) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCachePoliciesInput(v *ListCachePoliciesInput, encoder *httpbinding.Encoder) error {
@@ -7376,10 +7105,8 @@ func (m *awsRestxml_serializeOpListCloudFrontOriginAcc
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
@@ -7417,7 +7144,6 @@ func (m *awsRestxml_serializeOpListCloudFrontOriginAcc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCloudFrontOriginAccessIdentitiesInput(v *ListCloudFrontOriginAccessIdentitiesInput, encoder *httpbinding.Encoder) error {
@@ -7446,10 +7172,8 @@ func (m *awsRestxml_serializeOpListConflictingAliases)
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
@@ -7487,7 +7211,6 @@ func (m *awsRestxml_serializeOpListConflictingAliases)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListConflictingAliasesInput(v *ListConflictingAliasesInput, encoder *httpbinding.Encoder) error {
@@ -7524,10 +7247,8 @@ func (m *awsRestxml_serializeOpListConnectionGroups) H
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
@@ -7579,7 +7300,6 @@ func (m *awsRestxml_serializeOpListConnectionGroups) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListConnectionGroupsInput(v *ListConnectionGroupsInput, encoder *httpbinding.Encoder) error {
@@ -7640,10 +7360,8 @@ func (m *awsRestxml_serializeOpListContinuousDeploymen
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
@@ -7681,7 +7399,6 @@ func (m *awsRestxml_serializeOpListContinuousDeploymen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListContinuousDeploymentPoliciesInput(v *ListContinuousDeploymentPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -7710,10 +7427,8 @@ func (m *awsRestxml_serializeOpListDistributions) Hand
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
@@ -7751,7 +7466,6 @@ func (m *awsRestxml_serializeOpListDistributions) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsInput(v *ListDistributionsInput, encoder *httpbinding.Encoder) error {
@@ -7780,10 +7494,8 @@ func (m *awsRestxml_serializeOpListDistributionsByAnyc
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
@@ -7821,7 +7533,6 @@ func (m *awsRestxml_serializeOpListDistributionsByAnyc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByAnycastIpListIdInput(v *ListDistributionsByAnycastIpListIdInput, encoder *httpbinding.Encoder) error {
@@ -7859,10 +7570,8 @@ func (m *awsRestxml_serializeOpListDistributionsByCach
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
@@ -7900,7 +7609,6 @@ func (m *awsRestxml_serializeOpListDistributionsByCach
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByCachePolicyIdInput(v *ListDistributionsByCachePolicyIdInput, encoder *httpbinding.Encoder) error {
@@ -7938,10 +7646,8 @@ func (m *awsRestxml_serializeOpListDistributionsByConn
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
@@ -7979,7 +7685,6 @@ func (m *awsRestxml_serializeOpListDistributionsByConn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByConnectionModeInput(v *ListDistributionsByConnectionModeInput, encoder *httpbinding.Encoder) error {
@@ -8017,10 +7722,8 @@ func (m *awsRestxml_serializeOpListDistributionsByKeyG
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
@@ -8058,7 +7761,6 @@ func (m *awsRestxml_serializeOpListDistributionsByKeyG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByKeyGroupInput(v *ListDistributionsByKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -8096,10 +7798,8 @@ func (m *awsRestxml_serializeOpListDistributionsByOrig
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
@@ -8137,7 +7837,6 @@ func (m *awsRestxml_serializeOpListDistributionsByOrig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByOriginRequestPolicyIdInput(v *ListDistributionsByOriginRequestPolicyIdInput, encoder *httpbinding.Encoder) error {
@@ -8175,10 +7874,8 @@ func (m *awsRestxml_serializeOpListDistributionsByReal
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
@@ -8230,7 +7927,6 @@ func (m *awsRestxml_serializeOpListDistributionsByReal
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByRealtimeLogConfigInput(v *ListDistributionsByRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -8300,10 +7996,8 @@ func (m *awsRestxml_serializeOpListDistributionsByResp
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
@@ -8341,7 +8035,6 @@ func (m *awsRestxml_serializeOpListDistributionsByResp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByResponseHeadersPolicyIdInput(v *ListDistributionsByResponseHeadersPolicyIdInput, encoder *httpbinding.Encoder) error {
@@ -8379,10 +8072,8 @@ func (m *awsRestxml_serializeOpListDistributionsByVpcO
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
@@ -8420,7 +8111,6 @@ func (m *awsRestxml_serializeOpListDistributionsByVpcO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByVpcOriginIdInput(v *ListDistributionsByVpcOriginIdInput, encoder *httpbinding.Encoder) error {
@@ -8458,10 +8148,8 @@ func (m *awsRestxml_serializeOpListDistributionsByWebA
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
@@ -8499,7 +8187,6 @@ func (m *awsRestxml_serializeOpListDistributionsByWebA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionsByWebACLIdInput(v *ListDistributionsByWebACLIdInput, encoder *httpbinding.Encoder) error {
@@ -8537,10 +8224,8 @@ func (m *awsRestxml_serializeOpListDistributionTenants
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
@@ -8592,7 +8277,6 @@ func (m *awsRestxml_serializeOpListDistributionTenants
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionTenantsInput(v *ListDistributionTenantsInput, encoder *httpbinding.Encoder) error {
@@ -8653,10 +8337,8 @@ func (m *awsRestxml_serializeOpListDistributionTenants
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
@@ -8708,7 +8390,6 @@ func (m *awsRestxml_serializeOpListDistributionTenants
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDistributionTenantsByCustomizationInput(v *ListDistributionTenantsByCustomizationInput, encoder *httpbinding.Encoder) error {
@@ -8778,10 +8459,8 @@ func (m *awsRestxml_serializeOpListDomainConflicts) Ha
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
@@ -8833,7 +8512,6 @@ func (m *awsRestxml_serializeOpListDomainConflicts) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListDomainConflictsInput(v *ListDomainConflictsInput, encoder *httpbinding.Encoder) error {
@@ -8905,10 +8583,8 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
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
@@ -8946,7 +8622,6 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListFieldLevelEncryptionConfigsInput(v *ListFieldLevelEncryptionConfigsInput, encoder *httpbinding.Encoder) error {
@@ -8975,10 +8650,8 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
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
@@ -9016,7 +8689,6 @@ func (m *awsRestxml_serializeOpListFieldLevelEncryptio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListFieldLevelEncryptionProfilesInput(v *ListFieldLevelEncryptionProfilesInput, encoder *httpbinding.Encoder) error {
@@ -9045,10 +8717,8 @@ func (m *awsRestxml_serializeOpListFunctions) HandleSe
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
@@ -9086,7 +8756,6 @@ func (m *awsRestxml_serializeOpListFunctions) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListFunctionsInput(v *ListFunctionsInput, encoder *httpbinding.Encoder) error {
@@ -9119,10 +8788,8 @@ func (m *awsRestxml_serializeOpListInvalidations) Hand
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
@@ -9160,7 +8827,6 @@ func (m *awsRestxml_serializeOpListInvalidations) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListInvalidationsInput(v *ListInvalidationsInput, encoder *httpbinding.Encoder) error {
@@ -9198,10 +8864,8 @@ func (m *awsRestxml_serializeOpListInvalidationsForDis
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
@@ -9239,7 +8903,6 @@ func (m *awsRestxml_serializeOpListInvalidationsForDis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListInvalidationsForDistributionTenantInput(v *ListInvalidationsForDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -9277,10 +8940,8 @@ func (m *awsRestxml_serializeOpListKeyGroups) HandleSe
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
@@ -9318,7 +8979,6 @@ func (m *awsRestxml_serializeOpListKeyGroups) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListKeyGroupsInput(v *ListKeyGroupsInput, encoder *httpbinding.Encoder) error {
@@ -9347,10 +9007,8 @@ func (m *awsRestxml_serializeOpListKeyValueStores) Han
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
@@ -9388,7 +9046,6 @@ func (m *awsRestxml_serializeOpListKeyValueStores) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListKeyValueStoresInput(v *ListKeyValueStoresInput, encoder *httpbinding.Encoder) error {
@@ -9421,10 +9078,8 @@ func (m *awsRestxml_serializeOpListOriginAccessControl
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
@@ -9462,7 +9117,6 @@ func (m *awsRestxml_serializeOpListOriginAccessControl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListOriginAccessControlsInput(v *ListOriginAccessControlsInput, encoder *httpbinding.Encoder) error {
@@ -9491,10 +9145,8 @@ func (m *awsRestxml_serializeOpListOriginRequestPolici
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
@@ -9532,7 +9184,6 @@ func (m *awsRestxml_serializeOpListOriginRequestPolici
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListOriginRequestPoliciesInput(v *ListOriginRequestPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -9565,10 +9216,8 @@ func (m *awsRestxml_serializeOpListPublicKeys) HandleS
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
@@ -9606,7 +9255,6 @@ func (m *awsRestxml_serializeOpListPublicKeys) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListPublicKeysInput(v *ListPublicKeysInput, encoder *httpbinding.Encoder) error {
@@ -9635,10 +9283,8 @@ func (m *awsRestxml_serializeOpListRealtimeLogConfigs)
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
@@ -9676,7 +9322,6 @@ func (m *awsRestxml_serializeOpListRealtimeLogConfigs)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListRealtimeLogConfigsInput(v *ListRealtimeLogConfigsInput, encoder *httpbinding.Encoder) error {
@@ -9705,10 +9350,8 @@ func (m *awsRestxml_serializeOpListResponseHeadersPoli
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
@@ -9746,7 +9389,6 @@ func (m *awsRestxml_serializeOpListResponseHeadersPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListResponseHeadersPoliciesInput(v *ListResponseHeadersPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -9779,10 +9421,8 @@ func (m *awsRestxml_serializeOpListStreamingDistributi
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
@@ -9820,7 +9460,6 @@ func (m *awsRestxml_serializeOpListStreamingDistributi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListStreamingDistributionsInput(v *ListStreamingDistributionsInput, encoder *httpbinding.Encoder) error {
@@ -9849,10 +9488,8 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
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
@@ -9890,7 +9527,6 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -9915,10 +9551,8 @@ func (m *awsRestxml_serializeOpListVpcOrigins) HandleS
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
@@ -9956,7 +9590,6 @@ func (m *awsRestxml_serializeOpListVpcOrigins) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListVpcOriginsInput(v *ListVpcOriginsInput, encoder *httpbinding.Encoder) error {
@@ -9985,10 +9618,8 @@ func (m *awsRestxml_serializeOpPublishFunction) Handle
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
@@ -10026,7 +9657,6 @@ func (m *awsRestxml_serializeOpPublishFunction) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsPublishFunctionInput(v *PublishFunctionInput, encoder *httpbinding.Encoder) error {
@@ -10061,10 +9691,8 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
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
@@ -10126,7 +9754,6 @@ func (m *awsRestxml_serializeOpTagResource) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -10151,10 +9778,8 @@ func (m *awsRestxml_serializeOpTestFunction) HandleSer
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
@@ -10210,7 +9835,6 @@ func (m *awsRestxml_serializeOpTestFunction) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTestFunctionInput(v *TestFunctionInput, encoder *httpbinding.Encoder) error {
@@ -10272,10 +9896,8 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
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
@@ -10337,7 +9959,6 @@ func (m *awsRestxml_serializeOpUntagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -10362,10 +9983,8 @@ func (m *awsRestxml_serializeOpUpdateCachePolicy) Hand
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
@@ -10427,7 +10046,6 @@ func (m *awsRestxml_serializeOpUpdateCachePolicy) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateCachePolicyInput(v *UpdateCachePolicyInput, encoder *httpbinding.Encoder) error {
@@ -10462,10 +10080,8 @@ func (m *awsRestxml_serializeOpUpdateCloudFrontOriginA
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
@@ -10527,7 +10143,6 @@ func (m *awsRestxml_serializeOpUpdateCloudFrontOriginA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateCloudFrontOriginAccessIdentityInput(v *UpdateCloudFrontOriginAccessIdentityInput, encoder *httpbinding.Encoder) error {
@@ -10562,10 +10177,8 @@ func (m *awsRestxml_serializeOpUpdateConnectionGroup) 
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
@@ -10621,7 +10234,6 @@ func (m *awsRestxml_serializeOpUpdateConnectionGroup) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateConnectionGroupInput(v *UpdateConnectionGroupInput, encoder *httpbinding.Encoder) error {
@@ -10694,10 +10306,8 @@ func (m *awsRestxml_serializeOpUpdateContinuousDeploym
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
@@ -10759,7 +10369,6 @@ func (m *awsRestxml_serializeOpUpdateContinuousDeploym
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateContinuousDeploymentPolicyInput(v *UpdateContinuousDeploymentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -10794,10 +10403,8 @@ func (m *awsRestxml_serializeOpUpdateDistribution) Han
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
@@ -10859,7 +10466,6 @@ func (m *awsRestxml_serializeOpUpdateDistribution) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDistributionInput(v *UpdateDistributionInput, encoder *httpbinding.Encoder) error {
@@ -10894,10 +10500,8 @@ func (m *awsRestxml_serializeOpUpdateDistributionTenan
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
@@ -10953,7 +10557,6 @@ func (m *awsRestxml_serializeOpUpdateDistributionTenan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDistributionTenantInput(v *UpdateDistributionTenantInput, encoder *httpbinding.Encoder) error {
@@ -11078,10 +10681,8 @@ func (m *awsRestxml_serializeOpUpdateDistributionWithS
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
@@ -11119,7 +10720,6 @@ func (m *awsRestxml_serializeOpUpdateDistributionWithS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDistributionWithStagingConfigInput(v *UpdateDistributionWithStagingConfigInput, encoder *httpbinding.Encoder) error {
@@ -11158,10 +10758,8 @@ func (m *awsRestxml_serializeOpUpdateDomainAssociation
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
@@ -11217,7 +10815,6 @@ func (m *awsRestxml_serializeOpUpdateDomainAssociation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateDomainAssociationInput(v *UpdateDomainAssociationInput, encoder *httpbinding.Encoder) error {
@@ -11272,10 +10869,8 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
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
@@ -11337,7 +10932,6 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateFieldLevelEncryptionConfigInput(v *UpdateFieldLevelEncryptionConfigInput, encoder *httpbinding.Encoder) error {
@@ -11372,10 +10966,8 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
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
@@ -11437,7 +11029,6 @@ func (m *awsRestxml_serializeOpUpdateFieldLevelEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateFieldLevelEncryptionProfileInput(v *UpdateFieldLevelEncryptionProfileInput, encoder *httpbinding.Encoder) error {
@@ -11472,10 +11063,8 @@ func (m *awsRestxml_serializeOpUpdateFunction) HandleS
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
@@ -11531,7 +11120,6 @@ func (m *awsRestxml_serializeOpUpdateFunction) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateFunctionInput(v *UpdateFunctionInput, encoder *httpbinding.Encoder) error {
@@ -11595,10 +11183,8 @@ func (m *awsRestxml_serializeOpUpdateKeyGroup) HandleS
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
@@ -11660,7 +11246,6 @@ func (m *awsRestxml_serializeOpUpdateKeyGroup) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateKeyGroupInput(v *UpdateKeyGroupInput, encoder *httpbinding.Encoder) error {
@@ -11695,10 +11280,8 @@ func (m *awsRestxml_serializeOpUpdateKeyValueStore) Ha
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
@@ -11754,7 +11337,6 @@ func (m *awsRestxml_serializeOpUpdateKeyValueStore) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateKeyValueStoreInput(v *UpdateKeyValueStoreInput, encoder *httpbinding.Encoder) error {
@@ -11805,10 +11387,8 @@ func (m *awsRestxml_serializeOpUpdateOriginAccessContr
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
@@ -11870,7 +11450,6 @@ func (m *awsRestxml_serializeOpUpdateOriginAccessContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateOriginAccessControlInput(v *UpdateOriginAccessControlInput, encoder *httpbinding.Encoder) error {
@@ -11905,10 +11484,8 @@ func (m *awsRestxml_serializeOpUpdateOriginRequestPoli
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
@@ -11970,7 +11547,6 @@ func (m *awsRestxml_serializeOpUpdateOriginRequestPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateOriginRequestPolicyInput(v *UpdateOriginRequestPolicyInput, encoder *httpbinding.Encoder) error {
@@ -12005,10 +11581,8 @@ func (m *awsRestxml_serializeOpUpdatePublicKey) Handle
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
@@ -12070,7 +11644,6 @@ func (m *awsRestxml_serializeOpUpdatePublicKey) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdatePublicKeyInput(v *UpdatePublicKeyInput, encoder *httpbinding.Encoder) error {
@@ -12105,10 +11678,8 @@ func (m *awsRestxml_serializeOpUpdateRealtimeLogConfig
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
@@ -12160,7 +11731,6 @@ func (m *awsRestxml_serializeOpUpdateRealtimeLogConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateRealtimeLogConfigInput(v *UpdateRealtimeLogConfigInput, encoder *httpbinding.Encoder) error {
@@ -12245,10 +11815,8 @@ func (m *awsRestxml_serializeOpUpdateResponseHeadersPo
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
@@ -12310,7 +11878,6 @@ func (m *awsRestxml_serializeOpUpdateResponseHeadersPo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateResponseHeadersPolicyInput(v *UpdateResponseHeadersPolicyInput, encoder *httpbinding.Encoder) error {
@@ -12345,10 +11912,8 @@ func (m *awsRestxml_serializeOpUpdateStreamingDistribu
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
@@ -12410,7 +11975,6 @@ func (m *awsRestxml_serializeOpUpdateStreamingDistribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateStreamingDistributionInput(v *UpdateStreamingDistributionInput, encoder *httpbinding.Encoder) error {
@@ -12445,10 +12009,8 @@ func (m *awsRestxml_serializeOpUpdateVpcOrigin) Handle
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
@@ -12510,7 +12072,6 @@ func (m *awsRestxml_serializeOpUpdateVpcOrigin) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateVpcOriginInput(v *UpdateVpcOriginInput, encoder *httpbinding.Encoder) error {
@@ -12545,10 +12106,8 @@ func (m *awsRestxml_serializeOpVerifyDnsConfiguration)
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
@@ -12600,7 +12159,6 @@ func (m *awsRestxml_serializeOpVerifyDnsConfiguration)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsVerifyDnsConfigurationInput(v *VerifyDnsConfigurationInput, encoder *httpbinding.Encoder) error {
