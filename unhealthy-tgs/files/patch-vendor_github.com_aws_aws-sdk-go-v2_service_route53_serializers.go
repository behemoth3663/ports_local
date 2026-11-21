--- vendor/github.com/aws/aws-sdk-go-v2/service/route53/serializers.go.orig	2026-11-21 14:46:19 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/route53/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyxml "github.com/aws/smithy-go/encoding/xml"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestxml_serializeOpActivateKeySigningKey) 
 func (m *awsRestxml_serializeOpActivateKeySigningKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestxml_serializeOpActivateKeySigningKey) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsActivateKeySigningKeyInput(v *ActivateKeySigningKeyInput, encoder *httpbinding.Encoder) error {
@@ -105,10 +101,8 @@ func (m *awsRestxml_serializeOpAssociateVPCWithHostedZ
 func (m *awsRestxml_serializeOpAssociateVPCWithHostedZone) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -164,7 +158,6 @@ func (m *awsRestxml_serializeOpAssociateVPCWithHostedZ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsAssociateVPCWithHostedZoneInput(v *AssociateVPCWithHostedZoneInput, encoder *httpbinding.Encoder) error {
@@ -223,10 +216,8 @@ func (m *awsRestxml_serializeOpChangeCidrCollection) H
 func (m *awsRestxml_serializeOpChangeCidrCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -282,7 +273,6 @@ func (m *awsRestxml_serializeOpChangeCidrCollection) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsChangeCidrCollectionInput(v *ChangeCidrCollectionInput, encoder *httpbinding.Encoder) error {
@@ -341,10 +331,8 @@ func (m *awsRestxml_serializeOpChangeResourceRecordSet
 func (m *awsRestxml_serializeOpChangeResourceRecordSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -400,7 +388,6 @@ func (m *awsRestxml_serializeOpChangeResourceRecordSet
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsChangeResourceRecordSetsInput(v *ChangeResourceRecordSetsInput, encoder *httpbinding.Encoder) error {
@@ -448,10 +435,8 @@ func (m *awsRestxml_serializeOpChangeTagsForResource) 
 func (m *awsRestxml_serializeOpChangeTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -507,7 +492,6 @@ func (m *awsRestxml_serializeOpChangeTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsChangeTagsForResourceInput(v *ChangeTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -577,10 +561,8 @@ func (m *awsRestxml_serializeOpCreateCidrCollection) H
 func (m *awsRestxml_serializeOpCreateCidrCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -632,7 +614,6 @@ func (m *awsRestxml_serializeOpCreateCidrCollection) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateCidrCollectionInput(v *CreateCidrCollectionInput, encoder *httpbinding.Encoder) error {
@@ -680,10 +661,8 @@ func (m *awsRestxml_serializeOpCreateHealthCheck) Hand
 func (m *awsRestxml_serializeOpCreateHealthCheck) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -735,7 +714,6 @@ func (m *awsRestxml_serializeOpCreateHealthCheck) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateHealthCheckInput(v *CreateHealthCheckInput, encoder *httpbinding.Encoder) error {
@@ -785,10 +763,8 @@ func (m *awsRestxml_serializeOpCreateHostedZone) Handl
 func (m *awsRestxml_serializeOpCreateHostedZone) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -840,7 +816,6 @@ func (m *awsRestxml_serializeOpCreateHostedZone) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateHostedZoneInput(v *CreateHostedZoneInput, encoder *httpbinding.Encoder) error {
@@ -925,10 +900,8 @@ func (m *awsRestxml_serializeOpCreateKeySigningKey) Ha
 func (m *awsRestxml_serializeOpCreateKeySigningKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -980,7 +953,6 @@ func (m *awsRestxml_serializeOpCreateKeySigningKey) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateKeySigningKeyInput(v *CreateKeySigningKeyInput, encoder *httpbinding.Encoder) error {
@@ -1061,10 +1033,8 @@ func (m *awsRestxml_serializeOpCreateQueryLoggingConfi
 func (m *awsRestxml_serializeOpCreateQueryLoggingConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1116,7 +1086,6 @@ func (m *awsRestxml_serializeOpCreateQueryLoggingConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateQueryLoggingConfigInput(v *CreateQueryLoggingConfigInput, encoder *httpbinding.Encoder) error {
@@ -1164,10 +1133,8 @@ func (m *awsRestxml_serializeOpCreateReusableDelegatio
 func (m *awsRestxml_serializeOpCreateReusableDelegationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1219,7 +1186,6 @@ func (m *awsRestxml_serializeOpCreateReusableDelegatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateReusableDelegationSetInput(v *CreateReusableDelegationSetInput, encoder *httpbinding.Encoder) error {
@@ -1267,10 +1233,8 @@ func (m *awsRestxml_serializeOpCreateTrafficPolicy) Ha
 func (m *awsRestxml_serializeOpCreateTrafficPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1322,7 +1286,6 @@ func (m *awsRestxml_serializeOpCreateTrafficPolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateTrafficPolicyInput(v *CreateTrafficPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1381,10 +1344,8 @@ func (m *awsRestxml_serializeOpCreateTrafficPolicyInst
 func (m *awsRestxml_serializeOpCreateTrafficPolicyInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1436,7 +1397,6 @@ func (m *awsRestxml_serializeOpCreateTrafficPolicyInst
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateTrafficPolicyInstanceInput(v *CreateTrafficPolicyInstanceInput, encoder *httpbinding.Encoder) error {
@@ -1517,10 +1477,8 @@ func (m *awsRestxml_serializeOpCreateTrafficPolicyVers
 func (m *awsRestxml_serializeOpCreateTrafficPolicyVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1576,7 +1534,6 @@ func (m *awsRestxml_serializeOpCreateTrafficPolicyVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateTrafficPolicyVersionInput(v *CreateTrafficPolicyVersionInput, encoder *httpbinding.Encoder) error {
@@ -1633,10 +1590,8 @@ func (m *awsRestxml_serializeOpCreateVPCAssociationAut
 func (m *awsRestxml_serializeOpCreateVPCAssociationAuthorization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1692,7 +1647,6 @@ func (m *awsRestxml_serializeOpCreateVPCAssociationAut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsCreateVPCAssociationAuthorizationInput(v *CreateVPCAssociationAuthorizationInput, encoder *httpbinding.Encoder) error {
@@ -1740,10 +1694,8 @@ func (m *awsRestxml_serializeOpDeactivateKeySigningKey
 func (m *awsRestxml_serializeOpDeactivateKeySigningKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1781,7 +1733,6 @@ func (m *awsRestxml_serializeOpDeactivateKeySigningKey
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeactivateKeySigningKeyInput(v *DeactivateKeySigningKeyInput, encoder *httpbinding.Encoder) error {
@@ -1820,10 +1771,8 @@ func (m *awsRestxml_serializeOpDeleteCidrCollection) H
 func (m *awsRestxml_serializeOpDeleteCidrCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1861,7 +1810,6 @@ func (m *awsRestxml_serializeOpDeleteCidrCollection) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteCidrCollectionInput(v *DeleteCidrCollectionInput, encoder *httpbinding.Encoder) error {
@@ -1891,10 +1839,8 @@ func (m *awsRestxml_serializeOpDeleteHealthCheck) Hand
 func (m *awsRestxml_serializeOpDeleteHealthCheck) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1932,7 +1878,6 @@ func (m *awsRestxml_serializeOpDeleteHealthCheck) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteHealthCheckInput(v *DeleteHealthCheckInput, encoder *httpbinding.Encoder) error {
@@ -1962,10 +1907,8 @@ func (m *awsRestxml_serializeOpDeleteHostedZone) Handl
 func (m *awsRestxml_serializeOpDeleteHostedZone) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2003,7 +1946,6 @@ func (m *awsRestxml_serializeOpDeleteHostedZone) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteHostedZoneInput(v *DeleteHostedZoneInput, encoder *httpbinding.Encoder) error {
@@ -2033,10 +1975,8 @@ func (m *awsRestxml_serializeOpDeleteKeySigningKey) Ha
 func (m *awsRestxml_serializeOpDeleteKeySigningKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2074,7 +2014,6 @@ func (m *awsRestxml_serializeOpDeleteKeySigningKey) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteKeySigningKeyInput(v *DeleteKeySigningKeyInput, encoder *httpbinding.Encoder) error {
@@ -2113,10 +2052,8 @@ func (m *awsRestxml_serializeOpDeleteQueryLoggingConfi
 func (m *awsRestxml_serializeOpDeleteQueryLoggingConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2154,7 +2091,6 @@ func (m *awsRestxml_serializeOpDeleteQueryLoggingConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteQueryLoggingConfigInput(v *DeleteQueryLoggingConfigInput, encoder *httpbinding.Encoder) error {
@@ -2184,10 +2120,8 @@ func (m *awsRestxml_serializeOpDeleteReusableDelegatio
 func (m *awsRestxml_serializeOpDeleteReusableDelegationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2225,7 +2159,6 @@ func (m *awsRestxml_serializeOpDeleteReusableDelegatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteReusableDelegationSetInput(v *DeleteReusableDelegationSetInput, encoder *httpbinding.Encoder) error {
@@ -2255,10 +2188,8 @@ func (m *awsRestxml_serializeOpDeleteTrafficPolicy) Ha
 func (m *awsRestxml_serializeOpDeleteTrafficPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2296,7 +2227,6 @@ func (m *awsRestxml_serializeOpDeleteTrafficPolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteTrafficPolicyInput(v *DeleteTrafficPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2335,10 +2265,8 @@ func (m *awsRestxml_serializeOpDeleteTrafficPolicyInst
 func (m *awsRestxml_serializeOpDeleteTrafficPolicyInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2376,7 +2304,6 @@ func (m *awsRestxml_serializeOpDeleteTrafficPolicyInst
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteTrafficPolicyInstanceInput(v *DeleteTrafficPolicyInstanceInput, encoder *httpbinding.Encoder) error {
@@ -2406,10 +2333,8 @@ func (m *awsRestxml_serializeOpDeleteVPCAssociationAut
 func (m *awsRestxml_serializeOpDeleteVPCAssociationAuthorization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2465,7 +2390,6 @@ func (m *awsRestxml_serializeOpDeleteVPCAssociationAut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDeleteVPCAssociationAuthorizationInput(v *DeleteVPCAssociationAuthorizationInput, encoder *httpbinding.Encoder) error {
@@ -2513,10 +2437,8 @@ func (m *awsRestxml_serializeOpDisableHostedZoneDNSSEC
 func (m *awsRestxml_serializeOpDisableHostedZoneDNSSEC) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2554,7 +2476,6 @@ func (m *awsRestxml_serializeOpDisableHostedZoneDNSSEC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDisableHostedZoneDNSSECInput(v *DisableHostedZoneDNSSECInput, encoder *httpbinding.Encoder) error {
@@ -2584,10 +2505,8 @@ func (m *awsRestxml_serializeOpDisassociateVPCFromHost
 func (m *awsRestxml_serializeOpDisassociateVPCFromHostedZone) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2643,7 +2562,6 @@ func (m *awsRestxml_serializeOpDisassociateVPCFromHost
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsDisassociateVPCFromHostedZoneInput(v *DisassociateVPCFromHostedZoneInput, encoder *httpbinding.Encoder) error {
@@ -2702,10 +2620,8 @@ func (m *awsRestxml_serializeOpEnableHostedZoneDNSSEC)
 func (m *awsRestxml_serializeOpEnableHostedZoneDNSSEC) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2743,7 +2659,6 @@ func (m *awsRestxml_serializeOpEnableHostedZoneDNSSEC)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsEnableHostedZoneDNSSECInput(v *EnableHostedZoneDNSSECInput, encoder *httpbinding.Encoder) error {
@@ -2773,10 +2688,8 @@ func (m *awsRestxml_serializeOpGetAccountLimit) Handle
 func (m *awsRestxml_serializeOpGetAccountLimit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2814,7 +2727,6 @@ func (m *awsRestxml_serializeOpGetAccountLimit) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetAccountLimitInput(v *GetAccountLimitInput, encoder *httpbinding.Encoder) error {
@@ -2844,10 +2756,8 @@ func (m *awsRestxml_serializeOpGetChange) HandleSerial
 func (m *awsRestxml_serializeOpGetChange) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2885,7 +2795,6 @@ func (m *awsRestxml_serializeOpGetChange) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetChangeInput(v *GetChangeInput, encoder *httpbinding.Encoder) error {
@@ -2915,10 +2824,8 @@ func (m *awsRestxml_serializeOpGetCheckerIpRanges) Han
 func (m *awsRestxml_serializeOpGetCheckerIpRanges) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2952,7 +2859,6 @@ func (m *awsRestxml_serializeOpGetCheckerIpRanges) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetCheckerIpRangesInput(v *GetCheckerIpRangesInput, encoder *httpbinding.Encoder) error {
@@ -2973,10 +2879,8 @@ func (m *awsRestxml_serializeOpGetDNSSEC) HandleSerial
 func (m *awsRestxml_serializeOpGetDNSSEC) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3014,7 +2918,6 @@ func (m *awsRestxml_serializeOpGetDNSSEC) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetDNSSECInput(v *GetDNSSECInput, encoder *httpbinding.Encoder) error {
@@ -3044,10 +2947,8 @@ func (m *awsRestxml_serializeOpGetGeoLocation) HandleS
 func (m *awsRestxml_serializeOpGetGeoLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3085,7 +2986,6 @@ func (m *awsRestxml_serializeOpGetGeoLocation) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetGeoLocationInput(v *GetGeoLocationInput, encoder *httpbinding.Encoder) error {
@@ -3118,10 +3018,8 @@ func (m *awsRestxml_serializeOpGetHealthCheck) HandleS
 func (m *awsRestxml_serializeOpGetHealthCheck) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3159,7 +3057,6 @@ func (m *awsRestxml_serializeOpGetHealthCheck) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetHealthCheckInput(v *GetHealthCheckInput, encoder *httpbinding.Encoder) error {
@@ -3189,10 +3086,8 @@ func (m *awsRestxml_serializeOpGetHealthCheckCount) Ha
 func (m *awsRestxml_serializeOpGetHealthCheckCount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3226,7 +3121,6 @@ func (m *awsRestxml_serializeOpGetHealthCheckCount) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetHealthCheckCountInput(v *GetHealthCheckCountInput, encoder *httpbinding.Encoder) error {
@@ -3247,10 +3141,8 @@ func (m *awsRestxml_serializeOpGetHealthCheckLastFailu
 func (m *awsRestxml_serializeOpGetHealthCheckLastFailureReason) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3288,7 +3180,6 @@ func (m *awsRestxml_serializeOpGetHealthCheckLastFailu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetHealthCheckLastFailureReasonInput(v *GetHealthCheckLastFailureReasonInput, encoder *httpbinding.Encoder) error {
@@ -3318,10 +3209,8 @@ func (m *awsRestxml_serializeOpGetHealthCheckStatus) H
 func (m *awsRestxml_serializeOpGetHealthCheckStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3359,7 +3248,6 @@ func (m *awsRestxml_serializeOpGetHealthCheckStatus) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetHealthCheckStatusInput(v *GetHealthCheckStatusInput, encoder *httpbinding.Encoder) error {
@@ -3389,10 +3277,8 @@ func (m *awsRestxml_serializeOpGetHostedZone) HandleSe
 func (m *awsRestxml_serializeOpGetHostedZone) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3430,7 +3316,6 @@ func (m *awsRestxml_serializeOpGetHostedZone) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetHostedZoneInput(v *GetHostedZoneInput, encoder *httpbinding.Encoder) error {
@@ -3460,10 +3345,8 @@ func (m *awsRestxml_serializeOpGetHostedZoneCount) Han
 func (m *awsRestxml_serializeOpGetHostedZoneCount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3497,7 +3380,6 @@ func (m *awsRestxml_serializeOpGetHostedZoneCount) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetHostedZoneCountInput(v *GetHostedZoneCountInput, encoder *httpbinding.Encoder) error {
@@ -3518,10 +3400,8 @@ func (m *awsRestxml_serializeOpGetHostedZoneLimit) Han
 func (m *awsRestxml_serializeOpGetHostedZoneLimit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3559,7 +3439,6 @@ func (m *awsRestxml_serializeOpGetHostedZoneLimit) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetHostedZoneLimitInput(v *GetHostedZoneLimitInput, encoder *httpbinding.Encoder) error {
@@ -3598,10 +3477,8 @@ func (m *awsRestxml_serializeOpGetQueryLoggingConfig) 
 func (m *awsRestxml_serializeOpGetQueryLoggingConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3639,7 +3516,6 @@ func (m *awsRestxml_serializeOpGetQueryLoggingConfig) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetQueryLoggingConfigInput(v *GetQueryLoggingConfigInput, encoder *httpbinding.Encoder) error {
@@ -3669,10 +3545,8 @@ func (m *awsRestxml_serializeOpGetReusableDelegationSe
 func (m *awsRestxml_serializeOpGetReusableDelegationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3710,7 +3584,6 @@ func (m *awsRestxml_serializeOpGetReusableDelegationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetReusableDelegationSetInput(v *GetReusableDelegationSetInput, encoder *httpbinding.Encoder) error {
@@ -3740,10 +3613,8 @@ func (m *awsRestxml_serializeOpGetReusableDelegationSe
 func (m *awsRestxml_serializeOpGetReusableDelegationSetLimit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3781,7 +3652,6 @@ func (m *awsRestxml_serializeOpGetReusableDelegationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetReusableDelegationSetLimitInput(v *GetReusableDelegationSetLimitInput, encoder *httpbinding.Encoder) error {
@@ -3820,10 +3690,8 @@ func (m *awsRestxml_serializeOpGetTrafficPolicy) Handl
 func (m *awsRestxml_serializeOpGetTrafficPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3861,7 +3729,6 @@ func (m *awsRestxml_serializeOpGetTrafficPolicy) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetTrafficPolicyInput(v *GetTrafficPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3900,10 +3767,8 @@ func (m *awsRestxml_serializeOpGetTrafficPolicyInstanc
 func (m *awsRestxml_serializeOpGetTrafficPolicyInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3941,7 +3806,6 @@ func (m *awsRestxml_serializeOpGetTrafficPolicyInstanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetTrafficPolicyInstanceInput(v *GetTrafficPolicyInstanceInput, encoder *httpbinding.Encoder) error {
@@ -3971,10 +3835,8 @@ func (m *awsRestxml_serializeOpGetTrafficPolicyInstanc
 func (m *awsRestxml_serializeOpGetTrafficPolicyInstanceCount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4008,7 +3870,6 @@ func (m *awsRestxml_serializeOpGetTrafficPolicyInstanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsGetTrafficPolicyInstanceCountInput(v *GetTrafficPolicyInstanceCountInput, encoder *httpbinding.Encoder) error {
@@ -4029,10 +3890,8 @@ func (m *awsRestxml_serializeOpListCidrBlocks) HandleS
 func (m *awsRestxml_serializeOpListCidrBlocks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4070,7 +3929,6 @@ func (m *awsRestxml_serializeOpListCidrBlocks) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCidrBlocksInput(v *ListCidrBlocksInput, encoder *httpbinding.Encoder) error {
@@ -4112,10 +3970,8 @@ func (m *awsRestxml_serializeOpListCidrCollections) Ha
 func (m *awsRestxml_serializeOpListCidrCollections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4153,7 +4009,6 @@ func (m *awsRestxml_serializeOpListCidrCollections) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCidrCollectionsInput(v *ListCidrCollectionsInput, encoder *httpbinding.Encoder) error {
@@ -4182,10 +4037,8 @@ func (m *awsRestxml_serializeOpListCidrLocations) Hand
 func (m *awsRestxml_serializeOpListCidrLocations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4223,7 +4076,6 @@ func (m *awsRestxml_serializeOpListCidrLocations) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListCidrLocationsInput(v *ListCidrLocationsInput, encoder *httpbinding.Encoder) error {
@@ -4261,10 +4113,8 @@ func (m *awsRestxml_serializeOpListGeoLocations) Handl
 func (m *awsRestxml_serializeOpListGeoLocations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4302,7 +4152,6 @@ func (m *awsRestxml_serializeOpListGeoLocations) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListGeoLocationsInput(v *ListGeoLocationsInput, encoder *httpbinding.Encoder) error {
@@ -4339,10 +4188,8 @@ func (m *awsRestxml_serializeOpListHealthChecks) Handl
 func (m *awsRestxml_serializeOpListHealthChecks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4380,7 +4227,6 @@ func (m *awsRestxml_serializeOpListHealthChecks) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListHealthChecksInput(v *ListHealthChecksInput, encoder *httpbinding.Encoder) error {
@@ -4409,10 +4255,8 @@ func (m *awsRestxml_serializeOpListHostedZones) Handle
 func (m *awsRestxml_serializeOpListHostedZones) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4450,7 +4294,6 @@ func (m *awsRestxml_serializeOpListHostedZones) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListHostedZonesInput(v *ListHostedZonesInput, encoder *httpbinding.Encoder) error {
@@ -4487,10 +4330,8 @@ func (m *awsRestxml_serializeOpListHostedZonesByName) 
 func (m *awsRestxml_serializeOpListHostedZonesByName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4528,7 +4369,6 @@ func (m *awsRestxml_serializeOpListHostedZonesByName) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListHostedZonesByNameInput(v *ListHostedZonesByNameInput, encoder *httpbinding.Encoder) error {
@@ -4561,10 +4401,8 @@ func (m *awsRestxml_serializeOpListHostedZonesByVPC) H
 func (m *awsRestxml_serializeOpListHostedZonesByVPC) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4602,7 +4440,6 @@ func (m *awsRestxml_serializeOpListHostedZonesByVPC) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListHostedZonesByVPCInput(v *ListHostedZonesByVPCInput, encoder *httpbinding.Encoder) error {
@@ -4639,10 +4476,8 @@ func (m *awsRestxml_serializeOpListQueryLoggingConfigs
 func (m *awsRestxml_serializeOpListQueryLoggingConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4680,7 +4515,6 @@ func (m *awsRestxml_serializeOpListQueryLoggingConfigs
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListQueryLoggingConfigsInput(v *ListQueryLoggingConfigsInput, encoder *httpbinding.Encoder) error {
@@ -4713,10 +4547,8 @@ func (m *awsRestxml_serializeOpListResourceRecordSets)
 func (m *awsRestxml_serializeOpListResourceRecordSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4754,7 +4586,6 @@ func (m *awsRestxml_serializeOpListResourceRecordSets)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListResourceRecordSetsInput(v *ListResourceRecordSetsInput, encoder *httpbinding.Encoder) error {
@@ -4800,10 +4631,8 @@ func (m *awsRestxml_serializeOpListReusableDelegationS
 func (m *awsRestxml_serializeOpListReusableDelegationSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4841,7 +4670,6 @@ func (m *awsRestxml_serializeOpListReusableDelegationS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListReusableDelegationSetsInput(v *ListReusableDelegationSetsInput, encoder *httpbinding.Encoder) error {
@@ -4870,10 +4698,8 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
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
@@ -4911,7 +4737,6 @@ func (m *awsRestxml_serializeOpListTagsForResource) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4950,10 +4775,8 @@ func (m *awsRestxml_serializeOpListTagsForResources) H
 func (m *awsRestxml_serializeOpListTagsForResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5009,7 +4832,6 @@ func (m *awsRestxml_serializeOpListTagsForResources) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTagsForResourcesInput(v *ListTagsForResourcesInput, encoder *httpbinding.Encoder) error {
@@ -5057,10 +4879,8 @@ func (m *awsRestxml_serializeOpListTrafficPolicies) Ha
 func (m *awsRestxml_serializeOpListTrafficPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5098,7 +4918,6 @@ func (m *awsRestxml_serializeOpListTrafficPolicies) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTrafficPoliciesInput(v *ListTrafficPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -5127,10 +4946,8 @@ func (m *awsRestxml_serializeOpListTrafficPolicyInstan
 func (m *awsRestxml_serializeOpListTrafficPolicyInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5168,7 +4985,6 @@ func (m *awsRestxml_serializeOpListTrafficPolicyInstan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTrafficPolicyInstancesInput(v *ListTrafficPolicyInstancesInput, encoder *httpbinding.Encoder) error {
@@ -5205,10 +5021,8 @@ func (m *awsRestxml_serializeOpListTrafficPolicyInstan
 func (m *awsRestxml_serializeOpListTrafficPolicyInstancesByHostedZone) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5246,7 +5060,6 @@ func (m *awsRestxml_serializeOpListTrafficPolicyInstan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTrafficPolicyInstancesByHostedZoneInput(v *ListTrafficPolicyInstancesByHostedZoneInput, encoder *httpbinding.Encoder) error {
@@ -5283,10 +5096,8 @@ func (m *awsRestxml_serializeOpListTrafficPolicyInstan
 func (m *awsRestxml_serializeOpListTrafficPolicyInstancesByPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5324,7 +5135,6 @@ func (m *awsRestxml_serializeOpListTrafficPolicyInstan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTrafficPolicyInstancesByPolicyInput(v *ListTrafficPolicyInstancesByPolicyInput, encoder *httpbinding.Encoder) error {
@@ -5369,10 +5179,8 @@ func (m *awsRestxml_serializeOpListTrafficPolicyVersio
 func (m *awsRestxml_serializeOpListTrafficPolicyVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5410,7 +5218,6 @@ func (m *awsRestxml_serializeOpListTrafficPolicyVersio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListTrafficPolicyVersionsInput(v *ListTrafficPolicyVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5448,10 +5255,8 @@ func (m *awsRestxml_serializeOpListVPCAssociationAutho
 func (m *awsRestxml_serializeOpListVPCAssociationAuthorizations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5489,7 +5294,6 @@ func (m *awsRestxml_serializeOpListVPCAssociationAutho
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsListVPCAssociationAuthorizationsInput(v *ListVPCAssociationAuthorizationsInput, encoder *httpbinding.Encoder) error {
@@ -5527,10 +5331,8 @@ func (m *awsRestxml_serializeOpTestDNSAnswer) HandleSe
 func (m *awsRestxml_serializeOpTestDNSAnswer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5568,7 +5370,6 @@ func (m *awsRestxml_serializeOpTestDNSAnswer) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsTestDNSAnswerInput(v *TestDNSAnswerInput, encoder *httpbinding.Encoder) error {
@@ -5613,10 +5414,8 @@ func (m *awsRestxml_serializeOpUpdateHealthCheck) Hand
 func (m *awsRestxml_serializeOpUpdateHealthCheck) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5672,7 +5471,6 @@ func (m *awsRestxml_serializeOpUpdateHealthCheck) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateHealthCheckInput(v *UpdateHealthCheckInput, encoder *httpbinding.Encoder) error {
@@ -5891,10 +5689,8 @@ func (m *awsRestxml_serializeOpUpdateHostedZoneComment
 func (m *awsRestxml_serializeOpUpdateHostedZoneComment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5950,7 +5746,6 @@ func (m *awsRestxml_serializeOpUpdateHostedZoneComment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateHostedZoneCommentInput(v *UpdateHostedZoneCommentInput, encoder *httpbinding.Encoder) error {
@@ -5996,10 +5791,8 @@ func (m *awsRestxml_serializeOpUpdateHostedZoneFeature
 func (m *awsRestxml_serializeOpUpdateHostedZoneFeatures) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6055,7 +5848,6 @@ func (m *awsRestxml_serializeOpUpdateHostedZoneFeature
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateHostedZoneFeaturesInput(v *UpdateHostedZoneFeaturesInput, encoder *httpbinding.Encoder) error {
@@ -6101,10 +5893,8 @@ func (m *awsRestxml_serializeOpUpdateTrafficPolicyComm
 func (m *awsRestxml_serializeOpUpdateTrafficPolicyComment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6160,7 +5950,6 @@ func (m *awsRestxml_serializeOpUpdateTrafficPolicyComm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateTrafficPolicyCommentInput(v *UpdateTrafficPolicyCommentInput, encoder *httpbinding.Encoder) error {
@@ -6215,10 +6004,8 @@ func (m *awsRestxml_serializeOpUpdateTrafficPolicyInst
 func (m *awsRestxml_serializeOpUpdateTrafficPolicyInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6274,7 +6061,6 @@ func (m *awsRestxml_serializeOpUpdateTrafficPolicyInst
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestxml_serializeOpHttpBindingsUpdateTrafficPolicyInstanceInput(v *UpdateTrafficPolicyInstanceInput, encoder *httpbinding.Encoder) error {
