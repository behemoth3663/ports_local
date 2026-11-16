--- vendor/github.com/aws/aws-sdk-go-v2/service/trustedadvisor/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/trustedadvisor/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateRecommenda
 func (m *awsRestjson1_serializeOpBatchUpdateRecommendationResourceExclusion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateRecommenda
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateRecommendationResourceExclusionInput(v *BatchUpdateRecommendationResourceExclusionInput, encoder *httpbinding.Encoder) error {
@@ -109,10 +105,8 @@ func (m *awsRestjson1_serializeOpGetOrganizationRecomm
 func (m *awsRestjson1_serializeOpGetOrganizationRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -150,7 +144,6 @@ func (m *awsRestjson1_serializeOpGetOrganizationRecomm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOrganizationRecommendationInput(v *GetOrganizationRecommendationInput, encoder *httpbinding.Encoder) error {
@@ -180,10 +173,8 @@ func (m *awsRestjson1_serializeOpGetRecommendation) Ha
 func (m *awsRestjson1_serializeOpGetRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -221,7 +212,6 @@ func (m *awsRestjson1_serializeOpGetRecommendation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRecommendationInput(v *GetRecommendationInput, encoder *httpbinding.Encoder) error {
@@ -251,10 +241,8 @@ func (m *awsRestjson1_serializeOpListChecks) HandleSer
 func (m *awsRestjson1_serializeOpListChecks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -292,7 +280,6 @@ func (m *awsRestjson1_serializeOpListChecks) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChecksInput(v *ListChecksInput, encoder *httpbinding.Encoder) error {
@@ -337,10 +324,8 @@ func (m *awsRestjson1_serializeOpListOrganizationRecom
 func (m *awsRestjson1_serializeOpListOrganizationRecommendationAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -378,7 +363,6 @@ func (m *awsRestjson1_serializeOpListOrganizationRecom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationAccountsInput(v *ListOrganizationRecommendationAccountsInput, encoder *httpbinding.Encoder) error {
@@ -420,10 +404,8 @@ func (m *awsRestjson1_serializeOpListOrganizationRecom
 func (m *awsRestjson1_serializeOpListOrganizationRecommendationResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -461,7 +443,6 @@ func (m *awsRestjson1_serializeOpListOrganizationRecom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationResourcesInput(v *ListOrganizationRecommendationResourcesInput, encoder *httpbinding.Encoder) error {
@@ -515,10 +496,8 @@ func (m *awsRestjson1_serializeOpListOrganizationRecom
 func (m *awsRestjson1_serializeOpListOrganizationRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -556,7 +535,6 @@ func (m *awsRestjson1_serializeOpListOrganizationRecom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationsInput(v *ListOrganizationRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -617,10 +595,8 @@ func (m *awsRestjson1_serializeOpListRecommendationRes
 func (m *awsRestjson1_serializeOpListRecommendationResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -658,7 +634,6 @@ func (m *awsRestjson1_serializeOpListRecommendationRes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecommendationResourcesInput(v *ListRecommendationResourcesInput, encoder *httpbinding.Encoder) error {
@@ -708,10 +683,8 @@ func (m *awsRestjson1_serializeOpListRecommendations) 
 func (m *awsRestjson1_serializeOpListRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -749,7 +722,6 @@ func (m *awsRestjson1_serializeOpListRecommendations) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecommendationsInput(v *ListRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -810,10 +782,8 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationRec
 func (m *awsRestjson1_serializeOpUpdateOrganizationRecommendationLifecycle) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -862,7 +832,6 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationRec
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOrganizationRecommendationLifecycleInput(v *UpdateOrganizationRecommendationLifecycleInput, encoder *httpbinding.Encoder) error {
@@ -914,10 +883,8 @@ func (m *awsRestjson1_serializeOpUpdateRecommendationL
 func (m *awsRestjson1_serializeOpUpdateRecommendationLifecycle) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -966,7 +933,6 @@ func (m *awsRestjson1_serializeOpUpdateRecommendationL
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRecommendationLifecycleInput(v *UpdateRecommendationLifecycleInput, encoder *httpbinding.Encoder) error {
