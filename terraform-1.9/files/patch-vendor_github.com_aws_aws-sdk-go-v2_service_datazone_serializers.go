--- vendor/github.com/aws/aws-sdk-go-v2/service/datazone/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/datazone/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAcceptPredictions) Ha
 func (m *awsRestjson1_serializeOpAcceptPredictions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpAcceptPredictions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptPredictionsInput(v *AcceptPredictionsInput, encoder *httpbinding.Encoder) error {
@@ -148,10 +144,8 @@ func (m *awsRestjson1_serializeOpAcceptSubscriptionReq
 func (m *awsRestjson1_serializeOpAcceptSubscriptionRequest) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -200,7 +194,6 @@ func (m *awsRestjson1_serializeOpAcceptSubscriptionReq
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptSubscriptionRequestInput(v *AcceptSubscriptionRequestInput, encoder *httpbinding.Encoder) error {
@@ -258,10 +251,8 @@ func (m *awsRestjson1_serializeOpAddEntityOwner) Handl
 func (m *awsRestjson1_serializeOpAddEntityOwner) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -310,7 +301,6 @@ func (m *awsRestjson1_serializeOpAddEntityOwner) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddEntityOwnerInput(v *AddEntityOwnerInput, encoder *httpbinding.Encoder) error {
@@ -377,10 +367,8 @@ func (m *awsRestjson1_serializeOpAddPolicyGrant) Handl
 func (m *awsRestjson1_serializeOpAddPolicyGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -429,7 +417,6 @@ func (m *awsRestjson1_serializeOpAddPolicyGrant) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddPolicyGrantInput(v *AddPolicyGrantInput, encoder *httpbinding.Encoder) error {
@@ -508,10 +495,8 @@ func (m *awsRestjson1_serializeOpAssociateEnvironmentR
 func (m *awsRestjson1_serializeOpAssociateEnvironmentRole) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -549,7 +534,6 @@ func (m *awsRestjson1_serializeOpAssociateEnvironmentR
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateEnvironmentRoleInput(v *AssociateEnvironmentRoleInput, encoder *httpbinding.Encoder) error {
@@ -597,10 +581,8 @@ func (m *awsRestjson1_serializeOpCancelMetadataGenerat
 func (m *awsRestjson1_serializeOpCancelMetadataGenerationRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -638,7 +620,6 @@ func (m *awsRestjson1_serializeOpCancelMetadataGenerat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelMetadataGenerationRunInput(v *CancelMetadataGenerationRunInput, encoder *httpbinding.Encoder) error {
@@ -677,10 +658,8 @@ func (m *awsRestjson1_serializeOpCancelSubscription) H
 func (m *awsRestjson1_serializeOpCancelSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -718,7 +697,6 @@ func (m *awsRestjson1_serializeOpCancelSubscription) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelSubscriptionInput(v *CancelSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -757,10 +735,8 @@ func (m *awsRestjson1_serializeOpCreateAsset) HandleSe
 func (m *awsRestjson1_serializeOpCreateAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -809,7 +785,6 @@ func (m *awsRestjson1_serializeOpCreateAsset) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssetInput(v *CreateAssetInput, encoder *httpbinding.Encoder) error {
@@ -902,10 +877,8 @@ func (m *awsRestjson1_serializeOpCreateAssetFilter) Ha
 func (m *awsRestjson1_serializeOpCreateAssetFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -954,7 +927,6 @@ func (m *awsRestjson1_serializeOpCreateAssetFilter) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssetFilterInput(v *CreateAssetFilterInput, encoder *httpbinding.Encoder) error {
@@ -1022,10 +994,8 @@ func (m *awsRestjson1_serializeOpCreateAssetRevision) 
 func (m *awsRestjson1_serializeOpCreateAssetRevision) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1074,7 +1044,6 @@ func (m *awsRestjson1_serializeOpCreateAssetRevision) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssetRevisionInput(v *CreateAssetRevisionInput, encoder *httpbinding.Encoder) error {
@@ -1161,10 +1130,8 @@ func (m *awsRestjson1_serializeOpCreateAssetType) Hand
 func (m *awsRestjson1_serializeOpCreateAssetType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1213,7 +1180,6 @@ func (m *awsRestjson1_serializeOpCreateAssetType) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssetTypeInput(v *CreateAssetTypeInput, encoder *httpbinding.Encoder) error {
@@ -1272,10 +1238,8 @@ func (m *awsRestjson1_serializeOpCreateConnection) Han
 func (m *awsRestjson1_serializeOpCreateConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1324,7 +1288,6 @@ func (m *awsRestjson1_serializeOpCreateConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectionInput(v *CreateConnectionInput, encoder *httpbinding.Encoder) error {
@@ -1395,10 +1358,8 @@ func (m *awsRestjson1_serializeOpCreateDataProduct) Ha
 func (m *awsRestjson1_serializeOpCreateDataProduct) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1447,7 +1408,6 @@ func (m *awsRestjson1_serializeOpCreateDataProduct) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataProductInput(v *CreateDataProductInput, encoder *httpbinding.Encoder) error {
@@ -1525,10 +1485,8 @@ func (m *awsRestjson1_serializeOpCreateDataProductRevi
 func (m *awsRestjson1_serializeOpCreateDataProductRevision) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1577,7 +1535,6 @@ func (m *awsRestjson1_serializeOpCreateDataProductRevi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataProductRevisionInput(v *CreateDataProductRevisionInput, encoder *httpbinding.Encoder) error {
@@ -1659,10 +1616,8 @@ func (m *awsRestjson1_serializeOpCreateDataSource) Han
 func (m *awsRestjson1_serializeOpCreateDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1711,7 +1666,6 @@ func (m *awsRestjson1_serializeOpCreateDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataSourceInput(v *CreateDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -1821,10 +1775,8 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
 func (m *awsRestjson1_serializeOpCreateDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1869,7 +1821,6 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainInput(v *CreateDomainInput, encoder *httpbinding.Encoder) error {
@@ -1946,10 +1897,8 @@ func (m *awsRestjson1_serializeOpCreateDomainUnit) Han
 func (m *awsRestjson1_serializeOpCreateDomainUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1998,7 +1947,6 @@ func (m *awsRestjson1_serializeOpCreateDomainUnit) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainUnitInput(v *CreateDomainUnitInput, encoder *httpbinding.Encoder) error {
@@ -2055,10 +2003,8 @@ func (m *awsRestjson1_serializeOpCreateEnvironment) Ha
 func (m *awsRestjson1_serializeOpCreateEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2107,7 +2053,6 @@ func (m *awsRestjson1_serializeOpCreateEnvironment) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEnvironmentInput(v *CreateEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -2203,10 +2148,8 @@ func (m *awsRestjson1_serializeOpCreateEnvironmentActi
 func (m *awsRestjson1_serializeOpCreateEnvironmentAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2255,7 +2198,6 @@ func (m *awsRestjson1_serializeOpCreateEnvironmentActi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEnvironmentActionInput(v *CreateEnvironmentActionInput, encoder *httpbinding.Encoder) error {
@@ -2318,10 +2260,8 @@ func (m *awsRestjson1_serializeOpCreateEnvironmentProf
 func (m *awsRestjson1_serializeOpCreateEnvironmentProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2370,7 +2310,6 @@ func (m *awsRestjson1_serializeOpCreateEnvironmentProf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEnvironmentProfileInput(v *CreateEnvironmentProfileInput, encoder *httpbinding.Encoder) error {
@@ -2444,10 +2383,8 @@ func (m *awsRestjson1_serializeOpCreateFormType) Handl
 func (m *awsRestjson1_serializeOpCreateFormType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2496,7 +2433,6 @@ func (m *awsRestjson1_serializeOpCreateFormType) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFormTypeInput(v *CreateFormTypeInput, encoder *httpbinding.Encoder) error {
@@ -2560,10 +2496,8 @@ func (m *awsRestjson1_serializeOpCreateGlossary) Handl
 func (m *awsRestjson1_serializeOpCreateGlossary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2612,7 +2546,6 @@ func (m *awsRestjson1_serializeOpCreateGlossary) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGlossaryInput(v *CreateGlossaryInput, encoder *httpbinding.Encoder) error {
@@ -2674,10 +2607,8 @@ func (m *awsRestjson1_serializeOpCreateGlossaryTerm) H
 func (m *awsRestjson1_serializeOpCreateGlossaryTerm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2726,7 +2657,6 @@ func (m *awsRestjson1_serializeOpCreateGlossaryTerm) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGlossaryTermInput(v *CreateGlossaryTermInput, encoder *httpbinding.Encoder) error {
@@ -2800,10 +2730,8 @@ func (m *awsRestjson1_serializeOpCreateGroupProfile) H
 func (m *awsRestjson1_serializeOpCreateGroupProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2852,7 +2780,6 @@ func (m *awsRestjson1_serializeOpCreateGroupProfile) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGroupProfileInput(v *CreateGroupProfileInput, encoder *httpbinding.Encoder) error {
@@ -2899,10 +2826,8 @@ func (m *awsRestjson1_serializeOpCreateListingChangeSe
 func (m *awsRestjson1_serializeOpCreateListingChangeSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2951,7 +2876,6 @@ func (m *awsRestjson1_serializeOpCreateListingChangeSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateListingChangeSetInput(v *CreateListingChangeSetInput, encoder *httpbinding.Encoder) error {
@@ -3013,10 +2937,8 @@ func (m *awsRestjson1_serializeOpCreateProject) Handle
 func (m *awsRestjson1_serializeOpCreateProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3065,7 +2987,6 @@ func (m *awsRestjson1_serializeOpCreateProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProjectInput(v *CreateProjectInput, encoder *httpbinding.Encoder) error {
@@ -3136,10 +3057,8 @@ func (m *awsRestjson1_serializeOpCreateProjectMembersh
 func (m *awsRestjson1_serializeOpCreateProjectMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3188,7 +3107,6 @@ func (m *awsRestjson1_serializeOpCreateProjectMembersh
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProjectMembershipInput(v *CreateProjectMembershipInput, encoder *httpbinding.Encoder) error {
@@ -3246,10 +3164,8 @@ func (m *awsRestjson1_serializeOpCreateProjectProfile)
 func (m *awsRestjson1_serializeOpCreateProjectProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3298,7 +3214,6 @@ func (m *awsRestjson1_serializeOpCreateProjectProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProjectProfileInput(v *CreateProjectProfileInput, encoder *httpbinding.Encoder) error {
@@ -3362,10 +3277,8 @@ func (m *awsRestjson1_serializeOpCreateRule) HandleSer
 func (m *awsRestjson1_serializeOpCreateRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3414,7 +3327,6 @@ func (m *awsRestjson1_serializeOpCreateRule) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRuleInput(v *CreateRuleInput, encoder *httpbinding.Encoder) error {
@@ -3492,10 +3404,8 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionGra
 func (m *awsRestjson1_serializeOpCreateSubscriptionGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3544,7 +3454,6 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionGra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSubscriptionGrantInput(v *CreateSubscriptionGrantInput, encoder *httpbinding.Encoder) error {
@@ -3610,10 +3519,8 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionReq
 func (m *awsRestjson1_serializeOpCreateSubscriptionRequest) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3662,7 +3569,6 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionReq
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSubscriptionRequestInput(v *CreateSubscriptionRequestInput, encoder *httpbinding.Encoder) error {
@@ -3730,10 +3636,8 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionTar
 func (m *awsRestjson1_serializeOpCreateSubscriptionTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3782,7 +3686,6 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionTar
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSubscriptionTargetInput(v *CreateSubscriptionTargetInput, encoder *httpbinding.Encoder) error {
@@ -3874,10 +3777,8 @@ func (m *awsRestjson1_serializeOpCreateUserProfile) Ha
 func (m *awsRestjson1_serializeOpCreateUserProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3926,7 +3827,6 @@ func (m *awsRestjson1_serializeOpCreateUserProfile) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUserProfileInput(v *CreateUserProfileInput, encoder *httpbinding.Encoder) error {
@@ -3978,10 +3878,8 @@ func (m *awsRestjson1_serializeOpDeleteAsset) HandleSe
 func (m *awsRestjson1_serializeOpDeleteAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4019,7 +3917,6 @@ func (m *awsRestjson1_serializeOpDeleteAsset) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssetInput(v *DeleteAssetInput, encoder *httpbinding.Encoder) error {
@@ -4058,10 +3955,8 @@ func (m *awsRestjson1_serializeOpDeleteAssetFilter) Ha
 func (m *awsRestjson1_serializeOpDeleteAssetFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4099,7 +3994,6 @@ func (m *awsRestjson1_serializeOpDeleteAssetFilter) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssetFilterInput(v *DeleteAssetFilterInput, encoder *httpbinding.Encoder) error {
@@ -4147,10 +4041,8 @@ func (m *awsRestjson1_serializeOpDeleteAssetType) Hand
 func (m *awsRestjson1_serializeOpDeleteAssetType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4188,7 +4080,6 @@ func (m *awsRestjson1_serializeOpDeleteAssetType) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssetTypeInput(v *DeleteAssetTypeInput, encoder *httpbinding.Encoder) error {
@@ -4227,10 +4118,8 @@ func (m *awsRestjson1_serializeOpDeleteConnection) Han
 func (m *awsRestjson1_serializeOpDeleteConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4268,7 +4157,6 @@ func (m *awsRestjson1_serializeOpDeleteConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConnectionInput(v *DeleteConnectionInput, encoder *httpbinding.Encoder) error {
@@ -4307,10 +4195,8 @@ func (m *awsRestjson1_serializeOpDeleteDataProduct) Ha
 func (m *awsRestjson1_serializeOpDeleteDataProduct) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4348,7 +4234,6 @@ func (m *awsRestjson1_serializeOpDeleteDataProduct) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataProductInput(v *DeleteDataProductInput, encoder *httpbinding.Encoder) error {
@@ -4387,10 +4272,8 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 func (m *awsRestjson1_serializeOpDeleteDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4428,7 +4311,6 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataSourceInput(v *DeleteDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -4475,10 +4357,8 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
 func (m *awsRestjson1_serializeOpDeleteDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4516,7 +4396,6 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainInput(v *DeleteDomainInput, encoder *httpbinding.Encoder) error {
@@ -4554,10 +4433,8 @@ func (m *awsRestjson1_serializeOpDeleteDomainUnit) Han
 func (m *awsRestjson1_serializeOpDeleteDomainUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4595,7 +4472,6 @@ func (m *awsRestjson1_serializeOpDeleteDomainUnit) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainUnitInput(v *DeleteDomainUnitInput, encoder *httpbinding.Encoder) error {
@@ -4634,10 +4510,8 @@ func (m *awsRestjson1_serializeOpDeleteEnvironment) Ha
 func (m *awsRestjson1_serializeOpDeleteEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4675,7 +4549,6 @@ func (m *awsRestjson1_serializeOpDeleteEnvironment) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEnvironmentInput(v *DeleteEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -4714,10 +4587,8 @@ func (m *awsRestjson1_serializeOpDeleteEnvironmentActi
 func (m *awsRestjson1_serializeOpDeleteEnvironmentAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4755,7 +4626,6 @@ func (m *awsRestjson1_serializeOpDeleteEnvironmentActi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEnvironmentActionInput(v *DeleteEnvironmentActionInput, encoder *httpbinding.Encoder) error {
@@ -4803,10 +4673,8 @@ func (m *awsRestjson1_serializeOpDeleteEnvironmentBlue
 func (m *awsRestjson1_serializeOpDeleteEnvironmentBlueprintConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4844,7 +4712,6 @@ func (m *awsRestjson1_serializeOpDeleteEnvironmentBlue
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEnvironmentBlueprintConfigurationInput(v *DeleteEnvironmentBlueprintConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4883,10 +4750,8 @@ func (m *awsRestjson1_serializeOpDeleteEnvironmentProf
 func (m *awsRestjson1_serializeOpDeleteEnvironmentProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4924,7 +4789,6 @@ func (m *awsRestjson1_serializeOpDeleteEnvironmentProf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEnvironmentProfileInput(v *DeleteEnvironmentProfileInput, encoder *httpbinding.Encoder) error {
@@ -4963,10 +4827,8 @@ func (m *awsRestjson1_serializeOpDeleteFormType) Handl
 func (m *awsRestjson1_serializeOpDeleteFormType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5004,7 +4866,6 @@ func (m *awsRestjson1_serializeOpDeleteFormType) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFormTypeInput(v *DeleteFormTypeInput, encoder *httpbinding.Encoder) error {
@@ -5043,10 +4904,8 @@ func (m *awsRestjson1_serializeOpDeleteGlossary) Handl
 func (m *awsRestjson1_serializeOpDeleteGlossary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5084,7 +4943,6 @@ func (m *awsRestjson1_serializeOpDeleteGlossary) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGlossaryInput(v *DeleteGlossaryInput, encoder *httpbinding.Encoder) error {
@@ -5123,10 +4981,8 @@ func (m *awsRestjson1_serializeOpDeleteGlossaryTerm) H
 func (m *awsRestjson1_serializeOpDeleteGlossaryTerm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5164,7 +5020,6 @@ func (m *awsRestjson1_serializeOpDeleteGlossaryTerm) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGlossaryTermInput(v *DeleteGlossaryTermInput, encoder *httpbinding.Encoder) error {
@@ -5203,10 +5058,8 @@ func (m *awsRestjson1_serializeOpDeleteListing) Handle
 func (m *awsRestjson1_serializeOpDeleteListing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5244,7 +5097,6 @@ func (m *awsRestjson1_serializeOpDeleteListing) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteListingInput(v *DeleteListingInput, encoder *httpbinding.Encoder) error {
@@ -5283,10 +5135,8 @@ func (m *awsRestjson1_serializeOpDeleteProject) Handle
 func (m *awsRestjson1_serializeOpDeleteProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5324,7 +5174,6 @@ func (m *awsRestjson1_serializeOpDeleteProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProjectInput(v *DeleteProjectInput, encoder *httpbinding.Encoder) error {
@@ -5367,10 +5216,8 @@ func (m *awsRestjson1_serializeOpDeleteProjectMembersh
 func (m *awsRestjson1_serializeOpDeleteProjectMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5419,7 +5266,6 @@ func (m *awsRestjson1_serializeOpDeleteProjectMembersh
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProjectMembershipInput(v *DeleteProjectMembershipInput, encoder *httpbinding.Encoder) error {
@@ -5472,10 +5318,8 @@ func (m *awsRestjson1_serializeOpDeleteProjectProfile)
 func (m *awsRestjson1_serializeOpDeleteProjectProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5513,7 +5357,6 @@ func (m *awsRestjson1_serializeOpDeleteProjectProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProjectProfileInput(v *DeleteProjectProfileInput, encoder *httpbinding.Encoder) error {
@@ -5552,10 +5395,8 @@ func (m *awsRestjson1_serializeOpDeleteRule) HandleSer
 func (m *awsRestjson1_serializeOpDeleteRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5593,7 +5434,6 @@ func (m *awsRestjson1_serializeOpDeleteRule) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRuleInput(v *DeleteRuleInput, encoder *httpbinding.Encoder) error {
@@ -5632,10 +5472,8 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionGra
 func (m *awsRestjson1_serializeOpDeleteSubscriptionGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5673,7 +5511,6 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionGra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSubscriptionGrantInput(v *DeleteSubscriptionGrantInput, encoder *httpbinding.Encoder) error {
@@ -5712,10 +5549,8 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionReq
 func (m *awsRestjson1_serializeOpDeleteSubscriptionRequest) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5753,7 +5588,6 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionReq
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSubscriptionRequestInput(v *DeleteSubscriptionRequestInput, encoder *httpbinding.Encoder) error {
@@ -5792,10 +5626,8 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionTar
 func (m *awsRestjson1_serializeOpDeleteSubscriptionTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5833,7 +5665,6 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionTar
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSubscriptionTargetInput(v *DeleteSubscriptionTargetInput, encoder *httpbinding.Encoder) error {
@@ -5881,10 +5712,8 @@ func (m *awsRestjson1_serializeOpDeleteTimeSeriesDataP
 func (m *awsRestjson1_serializeOpDeleteTimeSeriesDataPoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5922,7 +5751,6 @@ func (m *awsRestjson1_serializeOpDeleteTimeSeriesDataP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTimeSeriesDataPointsInput(v *DeleteTimeSeriesDataPointsInput, encoder *httpbinding.Encoder) error {
@@ -5978,10 +5806,8 @@ func (m *awsRestjson1_serializeOpDisassociateEnvironme
 func (m *awsRestjson1_serializeOpDisassociateEnvironmentRole) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6019,7 +5845,6 @@ func (m *awsRestjson1_serializeOpDisassociateEnvironme
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateEnvironmentRoleInput(v *DisassociateEnvironmentRoleInput, encoder *httpbinding.Encoder) error {
@@ -6067,10 +5892,8 @@ func (m *awsRestjson1_serializeOpGetAsset) HandleSeria
 func (m *awsRestjson1_serializeOpGetAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6108,7 +5931,6 @@ func (m *awsRestjson1_serializeOpGetAsset) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssetInput(v *GetAssetInput, encoder *httpbinding.Encoder) error {
@@ -6151,10 +5973,8 @@ func (m *awsRestjson1_serializeOpGetAssetFilter) Handl
 func (m *awsRestjson1_serializeOpGetAssetFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6192,7 +6012,6 @@ func (m *awsRestjson1_serializeOpGetAssetFilter) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssetFilterInput(v *GetAssetFilterInput, encoder *httpbinding.Encoder) error {
@@ -6240,10 +6059,8 @@ func (m *awsRestjson1_serializeOpGetAssetType) HandleS
 func (m *awsRestjson1_serializeOpGetAssetType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6281,7 +6098,6 @@ func (m *awsRestjson1_serializeOpGetAssetType) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssetTypeInput(v *GetAssetTypeInput, encoder *httpbinding.Encoder) error {
@@ -6324,10 +6140,8 @@ func (m *awsRestjson1_serializeOpGetConnection) Handle
 func (m *awsRestjson1_serializeOpGetConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6365,7 +6179,6 @@ func (m *awsRestjson1_serializeOpGetConnection) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectionInput(v *GetConnectionInput, encoder *httpbinding.Encoder) error {
@@ -6408,10 +6221,8 @@ func (m *awsRestjson1_serializeOpGetDataProduct) Handl
 func (m *awsRestjson1_serializeOpGetDataProduct) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6449,7 +6260,6 @@ func (m *awsRestjson1_serializeOpGetDataProduct) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataProductInput(v *GetDataProductInput, encoder *httpbinding.Encoder) error {
@@ -6492,10 +6302,8 @@ func (m *awsRestjson1_serializeOpGetDataSource) Handle
 func (m *awsRestjson1_serializeOpGetDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6533,7 +6341,6 @@ func (m *awsRestjson1_serializeOpGetDataSource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataSourceInput(v *GetDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -6572,10 +6379,8 @@ func (m *awsRestjson1_serializeOpGetDataSourceRun) Han
 func (m *awsRestjson1_serializeOpGetDataSourceRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6613,7 +6418,6 @@ func (m *awsRestjson1_serializeOpGetDataSourceRun) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataSourceRunInput(v *GetDataSourceRunInput, encoder *httpbinding.Encoder) error {
@@ -6652,10 +6456,8 @@ func (m *awsRestjson1_serializeOpGetDomain) HandleSeri
 func (m *awsRestjson1_serializeOpGetDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6693,7 +6495,6 @@ func (m *awsRestjson1_serializeOpGetDomain) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainInput(v *GetDomainInput, encoder *httpbinding.Encoder) error {
@@ -6723,10 +6524,8 @@ func (m *awsRestjson1_serializeOpGetDomainUnit) Handle
 func (m *awsRestjson1_serializeOpGetDomainUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6764,7 +6563,6 @@ func (m *awsRestjson1_serializeOpGetDomainUnit) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainUnitInput(v *GetDomainUnitInput, encoder *httpbinding.Encoder) error {
@@ -6803,10 +6601,8 @@ func (m *awsRestjson1_serializeOpGetEnvironment) Handl
 func (m *awsRestjson1_serializeOpGetEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6844,7 +6640,6 @@ func (m *awsRestjson1_serializeOpGetEnvironment) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEnvironmentInput(v *GetEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -6883,10 +6678,8 @@ func (m *awsRestjson1_serializeOpGetEnvironmentAction)
 func (m *awsRestjson1_serializeOpGetEnvironmentAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6924,7 +6717,6 @@ func (m *awsRestjson1_serializeOpGetEnvironmentAction)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEnvironmentActionInput(v *GetEnvironmentActionInput, encoder *httpbinding.Encoder) error {
@@ -6972,10 +6764,8 @@ func (m *awsRestjson1_serializeOpGetEnvironmentBluepri
 func (m *awsRestjson1_serializeOpGetEnvironmentBlueprint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7013,7 +6803,6 @@ func (m *awsRestjson1_serializeOpGetEnvironmentBluepri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEnvironmentBlueprintInput(v *GetEnvironmentBlueprintInput, encoder *httpbinding.Encoder) error {
@@ -7052,10 +6841,8 @@ func (m *awsRestjson1_serializeOpGetEnvironmentBluepri
 func (m *awsRestjson1_serializeOpGetEnvironmentBlueprintConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7093,7 +6880,6 @@ func (m *awsRestjson1_serializeOpGetEnvironmentBluepri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEnvironmentBlueprintConfigurationInput(v *GetEnvironmentBlueprintConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7132,10 +6918,8 @@ func (m *awsRestjson1_serializeOpGetEnvironmentCredent
 func (m *awsRestjson1_serializeOpGetEnvironmentCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7173,7 +6957,6 @@ func (m *awsRestjson1_serializeOpGetEnvironmentCredent
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEnvironmentCredentialsInput(v *GetEnvironmentCredentialsInput, encoder *httpbinding.Encoder) error {
@@ -7212,10 +6995,8 @@ func (m *awsRestjson1_serializeOpGetEnvironmentProfile
 func (m *awsRestjson1_serializeOpGetEnvironmentProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7253,7 +7034,6 @@ func (m *awsRestjson1_serializeOpGetEnvironmentProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEnvironmentProfileInput(v *GetEnvironmentProfileInput, encoder *httpbinding.Encoder) error {
@@ -7292,10 +7072,8 @@ func (m *awsRestjson1_serializeOpGetFormType) HandleSe
 func (m *awsRestjson1_serializeOpGetFormType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7333,7 +7111,6 @@ func (m *awsRestjson1_serializeOpGetFormType) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFormTypeInput(v *GetFormTypeInput, encoder *httpbinding.Encoder) error {
@@ -7376,10 +7153,8 @@ func (m *awsRestjson1_serializeOpGetGlossary) HandleSe
 func (m *awsRestjson1_serializeOpGetGlossary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7417,7 +7192,6 @@ func (m *awsRestjson1_serializeOpGetGlossary) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGlossaryInput(v *GetGlossaryInput, encoder *httpbinding.Encoder) error {
@@ -7456,10 +7230,8 @@ func (m *awsRestjson1_serializeOpGetGlossaryTerm) Hand
 func (m *awsRestjson1_serializeOpGetGlossaryTerm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7497,7 +7269,6 @@ func (m *awsRestjson1_serializeOpGetGlossaryTerm) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGlossaryTermInput(v *GetGlossaryTermInput, encoder *httpbinding.Encoder) error {
@@ -7536,10 +7307,8 @@ func (m *awsRestjson1_serializeOpGetGroupProfile) Hand
 func (m *awsRestjson1_serializeOpGetGroupProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7577,7 +7346,6 @@ func (m *awsRestjson1_serializeOpGetGroupProfile) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGroupProfileInput(v *GetGroupProfileInput, encoder *httpbinding.Encoder) error {
@@ -7616,10 +7384,8 @@ func (m *awsRestjson1_serializeOpGetIamPortalLoginUrl)
 func (m *awsRestjson1_serializeOpGetIamPortalLoginUrl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7657,7 +7423,6 @@ func (m *awsRestjson1_serializeOpGetIamPortalLoginUrl)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIamPortalLoginUrlInput(v *GetIamPortalLoginUrlInput, encoder *httpbinding.Encoder) error {
@@ -7687,10 +7452,8 @@ func (m *awsRestjson1_serializeOpGetJobRun) HandleSeri
 func (m *awsRestjson1_serializeOpGetJobRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7728,7 +7491,6 @@ func (m *awsRestjson1_serializeOpGetJobRun) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJobRunInput(v *GetJobRunInput, encoder *httpbinding.Encoder) error {
@@ -7767,10 +7529,8 @@ func (m *awsRestjson1_serializeOpGetLineageEvent) Hand
 func (m *awsRestjson1_serializeOpGetLineageEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7808,7 +7568,6 @@ func (m *awsRestjson1_serializeOpGetLineageEvent) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLineageEventInput(v *GetLineageEventInput, encoder *httpbinding.Encoder) error {
@@ -7847,10 +7606,8 @@ func (m *awsRestjson1_serializeOpGetLineageNode) Handl
 func (m *awsRestjson1_serializeOpGetLineageNode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7888,7 +7645,6 @@ func (m *awsRestjson1_serializeOpGetLineageNode) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLineageNodeInput(v *GetLineageNodeInput, encoder *httpbinding.Encoder) error {
@@ -7931,10 +7687,8 @@ func (m *awsRestjson1_serializeOpGetListing) HandleSer
 func (m *awsRestjson1_serializeOpGetListing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7972,7 +7726,6 @@ func (m *awsRestjson1_serializeOpGetListing) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetListingInput(v *GetListingInput, encoder *httpbinding.Encoder) error {
@@ -8015,10 +7768,8 @@ func (m *awsRestjson1_serializeOpGetMetadataGeneration
 func (m *awsRestjson1_serializeOpGetMetadataGenerationRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8056,7 +7807,6 @@ func (m *awsRestjson1_serializeOpGetMetadataGeneration
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMetadataGenerationRunInput(v *GetMetadataGenerationRunInput, encoder *httpbinding.Encoder) error {
@@ -8095,10 +7845,8 @@ func (m *awsRestjson1_serializeOpGetProject) HandleSer
 func (m *awsRestjson1_serializeOpGetProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8136,7 +7884,6 @@ func (m *awsRestjson1_serializeOpGetProject) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProjectInput(v *GetProjectInput, encoder *httpbinding.Encoder) error {
@@ -8175,10 +7922,8 @@ func (m *awsRestjson1_serializeOpGetProjectProfile) Ha
 func (m *awsRestjson1_serializeOpGetProjectProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8216,7 +7961,6 @@ func (m *awsRestjson1_serializeOpGetProjectProfile) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProjectProfileInput(v *GetProjectProfileInput, encoder *httpbinding.Encoder) error {
@@ -8255,10 +7999,8 @@ func (m *awsRestjson1_serializeOpGetRule) HandleSerial
 func (m *awsRestjson1_serializeOpGetRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8296,7 +8038,6 @@ func (m *awsRestjson1_serializeOpGetRule) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRuleInput(v *GetRuleInput, encoder *httpbinding.Encoder) error {
@@ -8339,10 +8080,8 @@ func (m *awsRestjson1_serializeOpGetSubscription) Hand
 func (m *awsRestjson1_serializeOpGetSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8380,7 +8119,6 @@ func (m *awsRestjson1_serializeOpGetSubscription) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSubscriptionInput(v *GetSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -8419,10 +8157,8 @@ func (m *awsRestjson1_serializeOpGetSubscriptionGrant)
 func (m *awsRestjson1_serializeOpGetSubscriptionGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8460,7 +8196,6 @@ func (m *awsRestjson1_serializeOpGetSubscriptionGrant)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSubscriptionGrantInput(v *GetSubscriptionGrantInput, encoder *httpbinding.Encoder) error {
@@ -8499,10 +8234,8 @@ func (m *awsRestjson1_serializeOpGetSubscriptionReques
 func (m *awsRestjson1_serializeOpGetSubscriptionRequestDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8540,7 +8273,6 @@ func (m *awsRestjson1_serializeOpGetSubscriptionReques
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSubscriptionRequestDetailsInput(v *GetSubscriptionRequestDetailsInput, encoder *httpbinding.Encoder) error {
@@ -8579,10 +8311,8 @@ func (m *awsRestjson1_serializeOpGetSubscriptionTarget
 func (m *awsRestjson1_serializeOpGetSubscriptionTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8620,7 +8350,6 @@ func (m *awsRestjson1_serializeOpGetSubscriptionTarget
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSubscriptionTargetInput(v *GetSubscriptionTargetInput, encoder *httpbinding.Encoder) error {
@@ -8668,10 +8397,8 @@ func (m *awsRestjson1_serializeOpGetTimeSeriesDataPoin
 func (m *awsRestjson1_serializeOpGetTimeSeriesDataPoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8709,7 +8436,6 @@ func (m *awsRestjson1_serializeOpGetTimeSeriesDataPoin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTimeSeriesDataPointInput(v *GetTimeSeriesDataPointInput, encoder *httpbinding.Encoder) error {
@@ -8770,10 +8496,8 @@ func (m *awsRestjson1_serializeOpGetUserProfile) Handl
 func (m *awsRestjson1_serializeOpGetUserProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8811,7 +8535,6 @@ func (m *awsRestjson1_serializeOpGetUserProfile) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserProfileInput(v *GetUserProfileInput, encoder *httpbinding.Encoder) error {
@@ -8854,10 +8577,8 @@ func (m *awsRestjson1_serializeOpListAssetFilters) Han
 func (m *awsRestjson1_serializeOpListAssetFilters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8895,7 +8616,6 @@ func (m *awsRestjson1_serializeOpListAssetFilters) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetFiltersInput(v *ListAssetFiltersInput, encoder *httpbinding.Encoder) error {
@@ -8946,10 +8666,8 @@ func (m *awsRestjson1_serializeOpListAssetRevisions) H
 func (m *awsRestjson1_serializeOpListAssetRevisions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8987,7 +8705,6 @@ func (m *awsRestjson1_serializeOpListAssetRevisions) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssetRevisionsInput(v *ListAssetRevisionsInput, encoder *httpbinding.Encoder) error {
@@ -9034,10 +8751,8 @@ func (m *awsRestjson1_serializeOpListConnections) Hand
 func (m *awsRestjson1_serializeOpListConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9075,7 +8790,6 @@ func (m *awsRestjson1_serializeOpListConnections) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConnectionsInput(v *ListConnectionsInput, encoder *httpbinding.Encoder) error {
@@ -9137,10 +8851,8 @@ func (m *awsRestjson1_serializeOpListDataProductRevisi
 func (m *awsRestjson1_serializeOpListDataProductRevisions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9178,7 +8890,6 @@ func (m *awsRestjson1_serializeOpListDataProductRevisi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataProductRevisionsInput(v *ListDataProductRevisionsInput, encoder *httpbinding.Encoder) error {
@@ -9225,10 +8936,8 @@ func (m *awsRestjson1_serializeOpListDataSourceRunActi
 func (m *awsRestjson1_serializeOpListDataSourceRunActivities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9266,7 +8975,6 @@ func (m *awsRestjson1_serializeOpListDataSourceRunActi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataSourceRunActivitiesInput(v *ListDataSourceRunActivitiesInput, encoder *httpbinding.Encoder) error {
@@ -9317,10 +9025,8 @@ func (m *awsRestjson1_serializeOpListDataSourceRuns) H
 func (m *awsRestjson1_serializeOpListDataSourceRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9358,7 +9064,6 @@ func (m *awsRestjson1_serializeOpListDataSourceRuns) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataSourceRunsInput(v *ListDataSourceRunsInput, encoder *httpbinding.Encoder) error {
@@ -9409,10 +9114,8 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 func (m *awsRestjson1_serializeOpListDataSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9450,7 +9153,6 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataSourcesInput(v *ListDataSourcesInput, encoder *httpbinding.Encoder) error {
@@ -9512,10 +9214,8 @@ func (m *awsRestjson1_serializeOpListDomains) HandleSe
 func (m *awsRestjson1_serializeOpListDomains) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9553,7 +9253,6 @@ func (m *awsRestjson1_serializeOpListDomains) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainsInput(v *ListDomainsInput, encoder *httpbinding.Encoder) error {
@@ -9586,10 +9285,8 @@ func (m *awsRestjson1_serializeOpListDomainUnitsForPar
 func (m *awsRestjson1_serializeOpListDomainUnitsForParent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9627,7 +9324,6 @@ func (m *awsRestjson1_serializeOpListDomainUnitsForPar
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainUnitsForParentInput(v *ListDomainUnitsForParentInput, encoder *httpbinding.Encoder) error {
@@ -9669,10 +9365,8 @@ func (m *awsRestjson1_serializeOpListEntityOwners) Han
 func (m *awsRestjson1_serializeOpListEntityOwners) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9710,7 +9404,6 @@ func (m *awsRestjson1_serializeOpListEntityOwners) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEntityOwnersInput(v *ListEntityOwnersInput, encoder *httpbinding.Encoder) error {
@@ -9766,10 +9459,8 @@ func (m *awsRestjson1_serializeOpListEnvironmentAction
 func (m *awsRestjson1_serializeOpListEnvironmentActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9807,7 +9498,6 @@ func (m *awsRestjson1_serializeOpListEnvironmentAction
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEnvironmentActionsInput(v *ListEnvironmentActionsInput, encoder *httpbinding.Encoder) error {
@@ -9854,10 +9544,8 @@ func (m *awsRestjson1_serializeOpListEnvironmentBluepr
 func (m *awsRestjson1_serializeOpListEnvironmentBlueprintConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9895,7 +9583,6 @@ func (m *awsRestjson1_serializeOpListEnvironmentBluepr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEnvironmentBlueprintConfigurationsInput(v *ListEnvironmentBlueprintConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -9933,10 +9620,8 @@ func (m *awsRestjson1_serializeOpListEnvironmentBluepr
 func (m *awsRestjson1_serializeOpListEnvironmentBlueprints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9974,7 +9659,6 @@ func (m *awsRestjson1_serializeOpListEnvironmentBluepr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEnvironmentBlueprintsInput(v *ListEnvironmentBlueprintsInput, encoder *httpbinding.Encoder) error {
@@ -10020,10 +9704,8 @@ func (m *awsRestjson1_serializeOpListEnvironmentProfil
 func (m *awsRestjson1_serializeOpListEnvironmentProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10061,7 +9743,6 @@ func (m *awsRestjson1_serializeOpListEnvironmentProfil
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEnvironmentProfilesInput(v *ListEnvironmentProfilesInput, encoder *httpbinding.Encoder) error {
@@ -10119,10 +9800,8 @@ func (m *awsRestjson1_serializeOpListEnvironments) Han
 func (m *awsRestjson1_serializeOpListEnvironments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10160,7 +9839,6 @@ func (m *awsRestjson1_serializeOpListEnvironments) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEnvironmentsInput(v *ListEnvironmentsInput, encoder *httpbinding.Encoder) error {
@@ -10230,10 +9908,8 @@ func (m *awsRestjson1_serializeOpListJobRuns) HandleSe
 func (m *awsRestjson1_serializeOpListJobRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10271,7 +9947,6 @@ func (m *awsRestjson1_serializeOpListJobRuns) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobRunsInput(v *ListJobRunsInput, encoder *httpbinding.Encoder) error {
@@ -10326,10 +10001,8 @@ func (m *awsRestjson1_serializeOpListLineageEvents) Ha
 func (m *awsRestjson1_serializeOpListLineageEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10367,7 +10040,6 @@ func (m *awsRestjson1_serializeOpListLineageEvents) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLineageEventsInput(v *ListLineageEventsInput, encoder *httpbinding.Encoder) error {
@@ -10421,10 +10093,8 @@ func (m *awsRestjson1_serializeOpListLineageNodeHistor
 func (m *awsRestjson1_serializeOpListLineageNodeHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10462,7 +10132,6 @@ func (m *awsRestjson1_serializeOpListLineageNodeHistor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLineageNodeHistoryInput(v *ListLineageNodeHistoryInput, encoder *httpbinding.Encoder) error {
@@ -10525,10 +10194,8 @@ func (m *awsRestjson1_serializeOpListMetadataGeneratio
 func (m *awsRestjson1_serializeOpListMetadataGenerationRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10566,7 +10233,6 @@ func (m *awsRestjson1_serializeOpListMetadataGeneratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMetadataGenerationRunsInput(v *ListMetadataGenerationRunsInput, encoder *httpbinding.Encoder) error {
@@ -10612,10 +10278,8 @@ func (m *awsRestjson1_serializeOpListNotifications) Ha
 func (m *awsRestjson1_serializeOpListNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10653,7 +10317,6 @@ func (m *awsRestjson1_serializeOpListNotifications) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNotificationsInput(v *ListNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -10713,10 +10376,8 @@ func (m *awsRestjson1_serializeOpListPolicyGrants) Han
 func (m *awsRestjson1_serializeOpListPolicyGrants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10754,7 +10415,6 @@ func (m *awsRestjson1_serializeOpListPolicyGrants) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPolicyGrantsInput(v *ListPolicyGrantsInput, encoder *httpbinding.Encoder) error {
@@ -10814,10 +10474,8 @@ func (m *awsRestjson1_serializeOpListProjectMembership
 func (m *awsRestjson1_serializeOpListProjectMemberships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10855,7 +10513,6 @@ func (m *awsRestjson1_serializeOpListProjectMembership
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProjectMembershipsInput(v *ListProjectMembershipsInput, encoder *httpbinding.Encoder) error {
@@ -10910,10 +10567,8 @@ func (m *awsRestjson1_serializeOpListProjectProfiles) 
 func (m *awsRestjson1_serializeOpListProjectProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10951,7 +10606,6 @@ func (m *awsRestjson1_serializeOpListProjectProfiles) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProjectProfilesInput(v *ListProjectProfilesInput, encoder *httpbinding.Encoder) error {
@@ -11001,10 +10655,8 @@ func (m *awsRestjson1_serializeOpListProjects) HandleS
 func (m *awsRestjson1_serializeOpListProjects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11042,7 +10694,6 @@ func (m *awsRestjson1_serializeOpListProjects) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProjectsInput(v *ListProjectsInput, encoder *httpbinding.Encoder) error {
@@ -11092,10 +10743,8 @@ func (m *awsRestjson1_serializeOpListRules) HandleSeri
 func (m *awsRestjson1_serializeOpListRules) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11133,7 +10782,6 @@ func (m *awsRestjson1_serializeOpListRules) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRulesInput(v *ListRulesInput, encoder *httpbinding.Encoder) error {
@@ -11217,10 +10865,8 @@ func (m *awsRestjson1_serializeOpListSubscriptionGrant
 func (m *awsRestjson1_serializeOpListSubscriptionGrants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11258,7 +10904,6 @@ func (m *awsRestjson1_serializeOpListSubscriptionGrant
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubscriptionGrantsInput(v *ListSubscriptionGrantsInput, encoder *httpbinding.Encoder) error {
@@ -11324,10 +10969,8 @@ func (m *awsRestjson1_serializeOpListSubscriptionReque
 func (m *awsRestjson1_serializeOpListSubscriptionRequests) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11365,7 +11008,6 @@ func (m *awsRestjson1_serializeOpListSubscriptionReque
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubscriptionRequestsInput(v *ListSubscriptionRequestsInput, encoder *httpbinding.Encoder) error {
@@ -11427,10 +11069,8 @@ func (m *awsRestjson1_serializeOpListSubscriptions) Ha
 func (m *awsRestjson1_serializeOpListSubscriptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11468,7 +11108,6 @@ func (m *awsRestjson1_serializeOpListSubscriptions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubscriptionsInput(v *ListSubscriptionsInput, encoder *httpbinding.Encoder) error {
@@ -11534,10 +11173,8 @@ func (m *awsRestjson1_serializeOpListSubscriptionTarge
 func (m *awsRestjson1_serializeOpListSubscriptionTargets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11575,7 +11212,6 @@ func (m *awsRestjson1_serializeOpListSubscriptionTarge
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubscriptionTargetsInput(v *ListSubscriptionTargetsInput, encoder *httpbinding.Encoder) error {
@@ -11630,10 +11266,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11671,7 +11305,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -11701,10 +11334,8 @@ func (m *awsRestjson1_serializeOpListTimeSeriesDataPoi
 func (m *awsRestjson1_serializeOpListTimeSeriesDataPoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11742,7 +11373,6 @@ func (m *awsRestjson1_serializeOpListTimeSeriesDataPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTimeSeriesDataPointsInput(v *ListTimeSeriesDataPointsInput, encoder *httpbinding.Encoder) error {
@@ -11810,10 +11440,8 @@ func (m *awsRestjson1_serializeOpPostLineageEvent) Han
 func (m *awsRestjson1_serializeOpPostLineageEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11863,7 +11491,6 @@ func (m *awsRestjson1_serializeOpPostLineageEvent) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPostLineageEventInput(v *PostLineageEventInput, encoder *httpbinding.Encoder) error {
@@ -11898,10 +11525,8 @@ func (m *awsRestjson1_serializeOpPostTimeSeriesDataPoi
 func (m *awsRestjson1_serializeOpPostTimeSeriesDataPoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11950,7 +11575,6 @@ func (m *awsRestjson1_serializeOpPostTimeSeriesDataPoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPostTimeSeriesDataPointsInput(v *PostTimeSeriesDataPointsInput, encoder *httpbinding.Encoder) error {
@@ -12017,10 +11641,8 @@ func (m *awsRestjson1_serializeOpPutEnvironmentBluepri
 func (m *awsRestjson1_serializeOpPutEnvironmentBlueprintConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12069,7 +11691,6 @@ func (m *awsRestjson1_serializeOpPutEnvironmentBluepri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEnvironmentBlueprintConfigurationInput(v *PutEnvironmentBlueprintConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -12151,10 +11772,8 @@ func (m *awsRestjson1_serializeOpRejectPredictions) Ha
 func (m *awsRestjson1_serializeOpRejectPredictions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12203,7 +11822,6 @@ func (m *awsRestjson1_serializeOpRejectPredictions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectPredictionsInput(v *RejectPredictionsInput, encoder *httpbinding.Encoder) error {
@@ -12272,10 +11890,8 @@ func (m *awsRestjson1_serializeOpRejectSubscriptionReq
 func (m *awsRestjson1_serializeOpRejectSubscriptionRequest) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12324,7 +11940,6 @@ func (m *awsRestjson1_serializeOpRejectSubscriptionReq
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectSubscriptionRequestInput(v *RejectSubscriptionRequestInput, encoder *httpbinding.Encoder) error {
@@ -12375,10 +11990,8 @@ func (m *awsRestjson1_serializeOpRemoveEntityOwner) Ha
 func (m *awsRestjson1_serializeOpRemoveEntityOwner) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12427,7 +12040,6 @@ func (m *awsRestjson1_serializeOpRemoveEntityOwner) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveEntityOwnerInput(v *RemoveEntityOwnerInput, encoder *httpbinding.Encoder) error {
@@ -12494,10 +12106,8 @@ func (m *awsRestjson1_serializeOpRemovePolicyGrant) Ha
 func (m *awsRestjson1_serializeOpRemovePolicyGrant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12546,7 +12156,6 @@ func (m *awsRestjson1_serializeOpRemovePolicyGrant) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemovePolicyGrantInput(v *RemovePolicyGrantInput, encoder *httpbinding.Encoder) error {
@@ -12618,10 +12227,8 @@ func (m *awsRestjson1_serializeOpRevokeSubscription) H
 func (m *awsRestjson1_serializeOpRevokeSubscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12670,7 +12277,6 @@ func (m *awsRestjson1_serializeOpRevokeSubscription) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRevokeSubscriptionInput(v *RevokeSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -12721,10 +12327,8 @@ func (m *awsRestjson1_serializeOpSearch) HandleSeriali
 func (m *awsRestjson1_serializeOpSearch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12773,7 +12377,6 @@ func (m *awsRestjson1_serializeOpSearch) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchInput(v *SearchInput, encoder *httpbinding.Encoder) error {
@@ -12863,10 +12466,8 @@ func (m *awsRestjson1_serializeOpSearchGroupProfiles) 
 func (m *awsRestjson1_serializeOpSearchGroupProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -12915,7 +12516,6 @@ func (m *awsRestjson1_serializeOpSearchGroupProfiles) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchGroupProfilesInput(v *SearchGroupProfilesInput, encoder *httpbinding.Encoder) error {
@@ -12972,10 +12572,8 @@ func (m *awsRestjson1_serializeOpSearchListings) Handl
 func (m *awsRestjson1_serializeOpSearchListings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13024,7 +12622,6 @@ func (m *awsRestjson1_serializeOpSearchListings) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchListingsInput(v *SearchListingsInput, encoder *httpbinding.Encoder) error {
@@ -13104,10 +12701,8 @@ func (m *awsRestjson1_serializeOpSearchTypes) HandleSe
 func (m *awsRestjson1_serializeOpSearchTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13156,7 +12751,6 @@ func (m *awsRestjson1_serializeOpSearchTypes) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchTypesInput(v *SearchTypesInput, encoder *httpbinding.Encoder) error {
@@ -13239,10 +12833,8 @@ func (m *awsRestjson1_serializeOpSearchUserProfiles) H
 func (m *awsRestjson1_serializeOpSearchUserProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13291,7 +12883,6 @@ func (m *awsRestjson1_serializeOpSearchUserProfiles) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchUserProfilesInput(v *SearchUserProfilesInput, encoder *httpbinding.Encoder) error {
@@ -13348,10 +12939,8 @@ func (m *awsRestjson1_serializeOpStartDataSourceRun) H
 func (m *awsRestjson1_serializeOpStartDataSourceRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13400,7 +12989,6 @@ func (m *awsRestjson1_serializeOpStartDataSourceRun) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDataSourceRunInput(v *StartDataSourceRunInput, encoder *httpbinding.Encoder) error {
@@ -13451,10 +13039,8 @@ func (m *awsRestjson1_serializeOpStartMetadataGenerati
 func (m *awsRestjson1_serializeOpStartMetadataGenerationRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13503,7 +13089,6 @@ func (m *awsRestjson1_serializeOpStartMetadataGenerati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMetadataGenerationRunInput(v *StartMetadataGenerationRunInput, encoder *httpbinding.Encoder) error {
@@ -13562,10 +13147,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13614,7 +13197,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -13658,10 +13240,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13699,7 +13279,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -13735,10 +13314,8 @@ func (m *awsRestjson1_serializeOpUpdateAssetFilter) Ha
 func (m *awsRestjson1_serializeOpUpdateAssetFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13787,7 +13364,6 @@ func (m *awsRestjson1_serializeOpUpdateAssetFilter) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssetFilterInput(v *UpdateAssetFilterInput, encoder *httpbinding.Encoder) error {
@@ -13859,10 +13435,8 @@ func (m *awsRestjson1_serializeOpUpdateConnection) Han
 func (m *awsRestjson1_serializeOpUpdateConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -13911,7 +13485,6 @@ func (m *awsRestjson1_serializeOpUpdateConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConnectionInput(v *UpdateConnectionInput, encoder *httpbinding.Encoder) error {
@@ -13976,10 +13549,8 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 func (m *awsRestjson1_serializeOpUpdateDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14028,7 +13599,6 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataSourceInput(v *UpdateDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -14127,10 +13697,8 @@ func (m *awsRestjson1_serializeOpUpdateDomain) HandleS
 func (m *awsRestjson1_serializeOpUpdateDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14179,7 +13747,6 @@ func (m *awsRestjson1_serializeOpUpdateDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDomainInput(v *UpdateDomainInput, encoder *httpbinding.Encoder) error {
@@ -14247,10 +13814,8 @@ func (m *awsRestjson1_serializeOpUpdateDomainUnit) Han
 func (m *awsRestjson1_serializeOpUpdateDomainUnit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14299,7 +13864,6 @@ func (m *awsRestjson1_serializeOpUpdateDomainUnit) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDomainUnitInput(v *UpdateDomainUnitInput, encoder *httpbinding.Encoder) error {
@@ -14355,10 +13919,8 @@ func (m *awsRestjson1_serializeOpUpdateEnvironment) Ha
 func (m *awsRestjson1_serializeOpUpdateEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14407,7 +13969,6 @@ func (m *awsRestjson1_serializeOpUpdateEnvironment) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEnvironmentInput(v *UpdateEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -14482,10 +14043,8 @@ func (m *awsRestjson1_serializeOpUpdateEnvironmentActi
 func (m *awsRestjson1_serializeOpUpdateEnvironmentAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14534,7 +14093,6 @@ func (m *awsRestjson1_serializeOpUpdateEnvironmentActi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEnvironmentActionInput(v *UpdateEnvironmentActionInput, encoder *httpbinding.Encoder) error {
@@ -14606,10 +14164,8 @@ func (m *awsRestjson1_serializeOpUpdateEnvironmentProf
 func (m *awsRestjson1_serializeOpUpdateEnvironmentProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14658,7 +14214,6 @@ func (m *awsRestjson1_serializeOpUpdateEnvironmentProf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEnvironmentProfileInput(v *UpdateEnvironmentProfileInput, encoder *httpbinding.Encoder) error {
@@ -14731,10 +14286,8 @@ func (m *awsRestjson1_serializeOpUpdateGlossary) Handl
 func (m *awsRestjson1_serializeOpUpdateGlossary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14783,7 +14336,6 @@ func (m *awsRestjson1_serializeOpUpdateGlossary) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGlossaryInput(v *UpdateGlossaryInput, encoder *httpbinding.Encoder) error {
@@ -14849,10 +14401,8 @@ func (m *awsRestjson1_serializeOpUpdateGlossaryTerm) H
 func (m *awsRestjson1_serializeOpUpdateGlossaryTerm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -14901,7 +14451,6 @@ func (m *awsRestjson1_serializeOpUpdateGlossaryTerm) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGlossaryTermInput(v *UpdateGlossaryTermInput, encoder *httpbinding.Encoder) error {
@@ -14979,10 +14528,8 @@ func (m *awsRestjson1_serializeOpUpdateGroupProfile) H
 func (m *awsRestjson1_serializeOpUpdateGroupProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15031,7 +14578,6 @@ func (m *awsRestjson1_serializeOpUpdateGroupProfile) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGroupProfileInput(v *UpdateGroupProfileInput, encoder *httpbinding.Encoder) error {
@@ -15082,10 +14628,8 @@ func (m *awsRestjson1_serializeOpUpdateProject) Handle
 func (m *awsRestjson1_serializeOpUpdateProject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15134,7 +14678,6 @@ func (m *awsRestjson1_serializeOpUpdateProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProjectInput(v *UpdateProjectInput, encoder *httpbinding.Encoder) error {
@@ -15221,10 +14764,8 @@ func (m *awsRestjson1_serializeOpUpdateProjectProfile)
 func (m *awsRestjson1_serializeOpUpdateProjectProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15273,7 +14814,6 @@ func (m *awsRestjson1_serializeOpUpdateProjectProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProjectProfileInput(v *UpdateProjectProfileInput, encoder *httpbinding.Encoder) error {
@@ -15346,10 +14886,8 @@ func (m *awsRestjson1_serializeOpUpdateRule) HandleSer
 func (m *awsRestjson1_serializeOpUpdateRule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15398,7 +14936,6 @@ func (m *awsRestjson1_serializeOpUpdateRule) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRuleInput(v *UpdateRuleInput, encoder *httpbinding.Encoder) error {
@@ -15473,10 +15010,8 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionGra
 func (m *awsRestjson1_serializeOpUpdateSubscriptionGrantStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15525,7 +15060,6 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionGra
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSubscriptionGrantStatusInput(v *UpdateSubscriptionGrantStatusInput, encoder *httpbinding.Encoder) error {
@@ -15597,10 +15131,8 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionReq
 func (m *awsRestjson1_serializeOpUpdateSubscriptionRequest) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15649,7 +15181,6 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionReq
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSubscriptionRequestInput(v *UpdateSubscriptionRequestInput, encoder *httpbinding.Encoder) error {
@@ -15700,10 +15231,8 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionTar
 func (m *awsRestjson1_serializeOpUpdateSubscriptionTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15752,7 +15281,6 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionTar
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSubscriptionTargetInput(v *UpdateSubscriptionTargetInput, encoder *httpbinding.Encoder) error {
@@ -15843,10 +15371,8 @@ func (m *awsRestjson1_serializeOpUpdateUserProfile) Ha
 func (m *awsRestjson1_serializeOpUpdateUserProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -15895,7 +15421,6 @@ func (m *awsRestjson1_serializeOpUpdateUserProfile) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserProfileInput(v *UpdateUserProfileInput, encoder *httpbinding.Encoder) error {
