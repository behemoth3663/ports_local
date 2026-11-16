--- vendor/github.com/aws/aws-sdk-go-v2/service/resiliencehub/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/resiliencehub/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAcceptResourceGroupin
 func (m *awsRestjson1_serializeOpAcceptResourceGroupingRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpAcceptResourceGroupin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptResourceGroupingRecommendationsInput(v *AcceptResourceGroupingRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -114,10 +110,8 @@ func (m *awsRestjson1_serializeOpAddDraftAppVersionRes
 func (m *awsRestjson1_serializeOpAddDraftAppVersionResourceMappings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -162,7 +156,6 @@ func (m *awsRestjson1_serializeOpAddDraftAppVersionRes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddDraftAppVersionResourceMappingsInput(v *AddDraftAppVersionResourceMappingsInput, encoder *httpbinding.Encoder) error {
@@ -202,10 +195,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateRecommenda
 func (m *awsRestjson1_serializeOpBatchUpdateRecommendationStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -250,7 +241,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateRecommenda
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateRecommendationStatusInput(v *BatchUpdateRecommendationStatusInput, encoder *httpbinding.Encoder) error {
@@ -290,10 +280,8 @@ func (m *awsRestjson1_serializeOpCreateApp) HandleSeri
 func (m *awsRestjson1_serializeOpCreateApp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -338,7 +326,6 @@ func (m *awsRestjson1_serializeOpCreateApp) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAppInput(v *CreateAppInput, encoder *httpbinding.Encoder) error {
@@ -417,10 +404,8 @@ func (m *awsRestjson1_serializeOpCreateAppVersionAppCo
 func (m *awsRestjson1_serializeOpCreateAppVersionAppComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -465,7 +450,6 @@ func (m *awsRestjson1_serializeOpCreateAppVersionAppCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAppVersionAppComponentInput(v *CreateAppVersionAppComponentInput, encoder *httpbinding.Encoder) error {
@@ -525,10 +509,8 @@ func (m *awsRestjson1_serializeOpCreateAppVersionResou
 func (m *awsRestjson1_serializeOpCreateAppVersionResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -573,7 +555,6 @@ func (m *awsRestjson1_serializeOpCreateAppVersionResou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAppVersionResourceInput(v *CreateAppVersionResourceInput, encoder *httpbinding.Encoder) error {
@@ -657,10 +638,8 @@ func (m *awsRestjson1_serializeOpCreateRecommendationT
 func (m *awsRestjson1_serializeOpCreateRecommendationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -705,7 +684,6 @@ func (m *awsRestjson1_serializeOpCreateRecommendationT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRecommendationTemplateInput(v *CreateRecommendationTemplateInput, encoder *httpbinding.Encoder) error {
@@ -779,10 +757,8 @@ func (m *awsRestjson1_serializeOpCreateResiliencyPolic
 func (m *awsRestjson1_serializeOpCreateResiliencyPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -827,7 +803,6 @@ func (m *awsRestjson1_serializeOpCreateResiliencyPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateResiliencyPolicyInput(v *CreateResiliencyPolicyInput, encoder *httpbinding.Encoder) error {
@@ -894,10 +869,8 @@ func (m *awsRestjson1_serializeOpDeleteApp) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteApp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -942,7 +915,6 @@ func (m *awsRestjson1_serializeOpDeleteApp) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAppInput(v *DeleteAppInput, encoder *httpbinding.Encoder) error {
@@ -985,10 +957,8 @@ func (m *awsRestjson1_serializeOpDeleteAppAssessment) 
 func (m *awsRestjson1_serializeOpDeleteAppAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1033,7 +1003,6 @@ func (m *awsRestjson1_serializeOpDeleteAppAssessment) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAppAssessmentInput(v *DeleteAppAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -1071,10 +1040,8 @@ func (m *awsRestjson1_serializeOpDeleteAppInputSource)
 func (m *awsRestjson1_serializeOpDeleteAppInputSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1119,7 +1086,6 @@ func (m *awsRestjson1_serializeOpDeleteAppInputSource)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAppInputSourceInput(v *DeleteAppInputSourceInput, encoder *httpbinding.Encoder) error {
@@ -1176,10 +1142,8 @@ func (m *awsRestjson1_serializeOpDeleteAppVersionAppCo
 func (m *awsRestjson1_serializeOpDeleteAppVersionAppComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1224,7 +1188,6 @@ func (m *awsRestjson1_serializeOpDeleteAppVersionAppCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAppVersionAppComponentInput(v *DeleteAppVersionAppComponentInput, encoder *httpbinding.Encoder) error {
@@ -1267,10 +1230,8 @@ func (m *awsRestjson1_serializeOpDeleteAppVersionResou
 func (m *awsRestjson1_serializeOpDeleteAppVersionResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1315,7 +1276,6 @@ func (m *awsRestjson1_serializeOpDeleteAppVersionResou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAppVersionResourceInput(v *DeleteAppVersionResourceInput, encoder *httpbinding.Encoder) error {
@@ -1380,10 +1340,8 @@ func (m *awsRestjson1_serializeOpDeleteRecommendationT
 func (m *awsRestjson1_serializeOpDeleteRecommendationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1428,7 +1386,6 @@ func (m *awsRestjson1_serializeOpDeleteRecommendationT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRecommendationTemplateInput(v *DeleteRecommendationTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1466,10 +1423,8 @@ func (m *awsRestjson1_serializeOpDeleteResiliencyPolic
 func (m *awsRestjson1_serializeOpDeleteResiliencyPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1514,7 +1469,6 @@ func (m *awsRestjson1_serializeOpDeleteResiliencyPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResiliencyPolicyInput(v *DeleteResiliencyPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1552,10 +1506,8 @@ func (m *awsRestjson1_serializeOpDescribeApp) HandleSe
 func (m *awsRestjson1_serializeOpDescribeApp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1600,7 +1552,6 @@ func (m *awsRestjson1_serializeOpDescribeApp) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAppInput(v *DescribeAppInput, encoder *httpbinding.Encoder) error {
@@ -1633,10 +1584,8 @@ func (m *awsRestjson1_serializeOpDescribeAppAssessment
 func (m *awsRestjson1_serializeOpDescribeAppAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1681,7 +1630,6 @@ func (m *awsRestjson1_serializeOpDescribeAppAssessment
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAppAssessmentInput(v *DescribeAppAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -1714,10 +1662,8 @@ func (m *awsRestjson1_serializeOpDescribeAppVersion) H
 func (m *awsRestjson1_serializeOpDescribeAppVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1762,7 +1708,6 @@ func (m *awsRestjson1_serializeOpDescribeAppVersion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAppVersionInput(v *DescribeAppVersionInput, encoder *httpbinding.Encoder) error {
@@ -1800,10 +1745,8 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionApp
 func (m *awsRestjson1_serializeOpDescribeAppVersionAppComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1848,7 +1791,6 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionApp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAppVersionAppComponentInput(v *DescribeAppVersionAppComponentInput, encoder *httpbinding.Encoder) error {
@@ -1891,10 +1833,8 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionRes
 func (m *awsRestjson1_serializeOpDescribeAppVersionResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1939,7 +1879,6 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionRes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAppVersionResourceInput(v *DescribeAppVersionResourceInput, encoder *httpbinding.Encoder) error {
@@ -2004,10 +1943,8 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionRes
 func (m *awsRestjson1_serializeOpDescribeAppVersionResourcesResolutionStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2052,7 +1989,6 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionRes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAppVersionResourcesResolutionStatusInput(v *DescribeAppVersionResourcesResolutionStatusInput, encoder *httpbinding.Encoder) error {
@@ -2095,10 +2031,8 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionTem
 func (m *awsRestjson1_serializeOpDescribeAppVersionTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2143,7 +2077,6 @@ func (m *awsRestjson1_serializeOpDescribeAppVersionTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAppVersionTemplateInput(v *DescribeAppVersionTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2181,10 +2114,8 @@ func (m *awsRestjson1_serializeOpDescribeDraftAppVersi
 func (m *awsRestjson1_serializeOpDescribeDraftAppVersionResourcesImportStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2229,7 +2160,6 @@ func (m *awsRestjson1_serializeOpDescribeDraftAppVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDraftAppVersionResourcesImportStatusInput(v *DescribeDraftAppVersionResourcesImportStatusInput, encoder *httpbinding.Encoder) error {
@@ -2262,10 +2192,8 @@ func (m *awsRestjson1_serializeOpDescribeMetricsExport
 func (m *awsRestjson1_serializeOpDescribeMetricsExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2310,7 +2238,6 @@ func (m *awsRestjson1_serializeOpDescribeMetricsExport
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeMetricsExportInput(v *DescribeMetricsExportInput, encoder *httpbinding.Encoder) error {
@@ -2343,10 +2270,8 @@ func (m *awsRestjson1_serializeOpDescribeResiliencyPol
 func (m *awsRestjson1_serializeOpDescribeResiliencyPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2391,7 +2316,6 @@ func (m *awsRestjson1_serializeOpDescribeResiliencyPol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeResiliencyPolicyInput(v *DescribeResiliencyPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2424,10 +2348,8 @@ func (m *awsRestjson1_serializeOpDescribeResourceGroup
 func (m *awsRestjson1_serializeOpDescribeResourceGroupingRecommendationTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2472,7 +2394,6 @@ func (m *awsRestjson1_serializeOpDescribeResourceGroup
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeResourceGroupingRecommendationTaskInput(v *DescribeResourceGroupingRecommendationTaskInput, encoder *httpbinding.Encoder) error {
@@ -2510,10 +2431,8 @@ func (m *awsRestjson1_serializeOpImportResourcesToDraf
 func (m *awsRestjson1_serializeOpImportResourcesToDraftAppVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2558,7 +2477,6 @@ func (m *awsRestjson1_serializeOpImportResourcesToDraf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportResourcesToDraftAppVersionInput(v *ImportResourcesToDraftAppVersionInput, encoder *httpbinding.Encoder) error {
@@ -2617,10 +2535,8 @@ func (m *awsRestjson1_serializeOpListAlarmRecommendati
 func (m *awsRestjson1_serializeOpListAlarmRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2665,7 +2581,6 @@ func (m *awsRestjson1_serializeOpListAlarmRecommendati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAlarmRecommendationsInput(v *ListAlarmRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -2708,10 +2623,8 @@ func (m *awsRestjson1_serializeOpListAppAssessmentComp
 func (m *awsRestjson1_serializeOpListAppAssessmentComplianceDrifts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2756,7 +2669,6 @@ func (m *awsRestjson1_serializeOpListAppAssessmentComp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppAssessmentComplianceDriftsInput(v *ListAppAssessmentComplianceDriftsInput, encoder *httpbinding.Encoder) error {
@@ -2799,10 +2711,8 @@ func (m *awsRestjson1_serializeOpListAppAssessmentReso
 func (m *awsRestjson1_serializeOpListAppAssessmentResourceDrifts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2847,7 +2757,6 @@ func (m *awsRestjson1_serializeOpListAppAssessmentReso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppAssessmentResourceDriftsInput(v *ListAppAssessmentResourceDriftsInput, encoder *httpbinding.Encoder) error {
@@ -2890,10 +2799,8 @@ func (m *awsRestjson1_serializeOpListAppAssessments) H
 func (m *awsRestjson1_serializeOpListAppAssessments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2931,7 +2838,6 @@ func (m *awsRestjson1_serializeOpListAppAssessments) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppAssessmentsInput(v *ListAppAssessmentsInput, encoder *httpbinding.Encoder) error {
@@ -2986,10 +2892,8 @@ func (m *awsRestjson1_serializeOpListAppComponentCompl
 func (m *awsRestjson1_serializeOpListAppComponentCompliances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3034,7 +2938,6 @@ func (m *awsRestjson1_serializeOpListAppComponentCompl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppComponentCompliancesInput(v *ListAppComponentCompliancesInput, encoder *httpbinding.Encoder) error {
@@ -3077,10 +2980,8 @@ func (m *awsRestjson1_serializeOpListAppComponentRecom
 func (m *awsRestjson1_serializeOpListAppComponentRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3125,7 +3026,6 @@ func (m *awsRestjson1_serializeOpListAppComponentRecom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppComponentRecommendationsInput(v *ListAppComponentRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -3168,10 +3068,8 @@ func (m *awsRestjson1_serializeOpListAppInputSources) 
 func (m *awsRestjson1_serializeOpListAppInputSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3216,7 +3114,6 @@ func (m *awsRestjson1_serializeOpListAppInputSources) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppInputSourcesInput(v *ListAppInputSourcesInput, encoder *httpbinding.Encoder) error {
@@ -3264,10 +3161,8 @@ func (m *awsRestjson1_serializeOpListApps) HandleSeria
 func (m *awsRestjson1_serializeOpListApps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3305,7 +3200,6 @@ func (m *awsRestjson1_serializeOpListApps) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppsInput(v *ListAppsInput, encoder *httpbinding.Encoder) error {
@@ -3358,10 +3252,8 @@ func (m *awsRestjson1_serializeOpListAppVersionAppComp
 func (m *awsRestjson1_serializeOpListAppVersionAppComponents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3406,7 +3298,6 @@ func (m *awsRestjson1_serializeOpListAppVersionAppComp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppVersionAppComponentsInput(v *ListAppVersionAppComponentsInput, encoder *httpbinding.Encoder) error {
@@ -3454,10 +3345,8 @@ func (m *awsRestjson1_serializeOpListAppVersionResourc
 func (m *awsRestjson1_serializeOpListAppVersionResourceMappings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3502,7 +3391,6 @@ func (m *awsRestjson1_serializeOpListAppVersionResourc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppVersionResourceMappingsInput(v *ListAppVersionResourceMappingsInput, encoder *httpbinding.Encoder) error {
@@ -3550,10 +3438,8 @@ func (m *awsRestjson1_serializeOpListAppVersionResourc
 func (m *awsRestjson1_serializeOpListAppVersionResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3598,7 +3484,6 @@ func (m *awsRestjson1_serializeOpListAppVersionResourc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppVersionResourcesInput(v *ListAppVersionResourcesInput, encoder *httpbinding.Encoder) error {
@@ -3651,10 +3536,8 @@ func (m *awsRestjson1_serializeOpListAppVersions) Hand
 func (m *awsRestjson1_serializeOpListAppVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3699,7 +3582,6 @@ func (m *awsRestjson1_serializeOpListAppVersions) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppVersionsInput(v *ListAppVersionsInput, encoder *httpbinding.Encoder) error {
@@ -3752,10 +3634,8 @@ func (m *awsRestjson1_serializeOpListMetrics) HandleSe
 func (m *awsRestjson1_serializeOpListMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3800,7 +3680,6 @@ func (m *awsRestjson1_serializeOpListMetrics) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMetricsInput(v *ListMetricsInput, encoder *httpbinding.Encoder) error {
@@ -3864,10 +3743,8 @@ func (m *awsRestjson1_serializeOpListRecommendationTem
 func (m *awsRestjson1_serializeOpListRecommendationTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3905,7 +3782,6 @@ func (m *awsRestjson1_serializeOpListRecommendationTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecommendationTemplatesInput(v *ListRecommendationTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -3956,10 +3832,8 @@ func (m *awsRestjson1_serializeOpListResiliencyPolicie
 func (m *awsRestjson1_serializeOpListResiliencyPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3997,7 +3871,6 @@ func (m *awsRestjson1_serializeOpListResiliencyPolicie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResiliencyPoliciesInput(v *ListResiliencyPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -4030,10 +3903,8 @@ func (m *awsRestjson1_serializeOpListResourceGroupingR
 func (m *awsRestjson1_serializeOpListResourceGroupingRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4071,7 +3942,6 @@ func (m *awsRestjson1_serializeOpListResourceGroupingR
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceGroupingRecommendationsInput(v *ListResourceGroupingRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -4104,10 +3974,8 @@ func (m *awsRestjson1_serializeOpListSopRecommendation
 func (m *awsRestjson1_serializeOpListSopRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4152,7 +4020,6 @@ func (m *awsRestjson1_serializeOpListSopRecommendation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSopRecommendationsInput(v *ListSopRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -4195,10 +4062,8 @@ func (m *awsRestjson1_serializeOpListSuggestedResilien
 func (m *awsRestjson1_serializeOpListSuggestedResiliencyPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4236,7 +4101,6 @@ func (m *awsRestjson1_serializeOpListSuggestedResilien
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSuggestedResiliencyPoliciesInput(v *ListSuggestedResiliencyPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -4265,10 +4129,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -4306,7 +4168,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4336,10 +4197,8 @@ func (m *awsRestjson1_serializeOpListTestRecommendatio
 func (m *awsRestjson1_serializeOpListTestRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4384,7 +4243,6 @@ func (m *awsRestjson1_serializeOpListTestRecommendatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTestRecommendationsInput(v *ListTestRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -4427,10 +4285,8 @@ func (m *awsRestjson1_serializeOpListUnsupportedAppVer
 func (m *awsRestjson1_serializeOpListUnsupportedAppVersionResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4475,7 +4331,6 @@ func (m *awsRestjson1_serializeOpListUnsupportedAppVer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUnsupportedAppVersionResourcesInput(v *ListUnsupportedAppVersionResourcesInput, encoder *httpbinding.Encoder) error {
@@ -4528,10 +4383,8 @@ func (m *awsRestjson1_serializeOpPublishAppVersion) Ha
 func (m *awsRestjson1_serializeOpPublishAppVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4576,7 +4429,6 @@ func (m *awsRestjson1_serializeOpPublishAppVersion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPublishAppVersionInput(v *PublishAppVersionInput, encoder *httpbinding.Encoder) error {
@@ -4614,10 +4466,8 @@ func (m *awsRestjson1_serializeOpPutDraftAppVersionTem
 func (m *awsRestjson1_serializeOpPutDraftAppVersionTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4662,7 +4512,6 @@ func (m *awsRestjson1_serializeOpPutDraftAppVersionTem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDraftAppVersionTemplateInput(v *PutDraftAppVersionTemplateInput, encoder *httpbinding.Encoder) error {
@@ -4700,10 +4549,8 @@ func (m *awsRestjson1_serializeOpRejectResourceGroupin
 func (m *awsRestjson1_serializeOpRejectResourceGroupingRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4748,7 +4595,6 @@ func (m *awsRestjson1_serializeOpRejectResourceGroupin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectResourceGroupingRecommendationsInput(v *RejectResourceGroupingRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -4788,10 +4634,8 @@ func (m *awsRestjson1_serializeOpRemoveDraftAppVersion
 func (m *awsRestjson1_serializeOpRemoveDraftAppVersionResourceMappings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4836,7 +4680,6 @@ func (m *awsRestjson1_serializeOpRemoveDraftAppVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveDraftAppVersionResourceMappingsInput(v *RemoveDraftAppVersionResourceMappingsInput, encoder *httpbinding.Encoder) error {
@@ -4911,10 +4754,8 @@ func (m *awsRestjson1_serializeOpResolveAppVersionReso
 func (m *awsRestjson1_serializeOpResolveAppVersionResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4959,7 +4800,6 @@ func (m *awsRestjson1_serializeOpResolveAppVersionReso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResolveAppVersionResourcesInput(v *ResolveAppVersionResourcesInput, encoder *httpbinding.Encoder) error {
@@ -4997,10 +4837,8 @@ func (m *awsRestjson1_serializeOpStartAppAssessment) H
 func (m *awsRestjson1_serializeOpStartAppAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5045,7 +4883,6 @@ func (m *awsRestjson1_serializeOpStartAppAssessment) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAppAssessmentInput(v *StartAppAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -5100,10 +4937,8 @@ func (m *awsRestjson1_serializeOpStartMetricsExport) H
 func (m *awsRestjson1_serializeOpStartMetricsExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5148,7 +4983,6 @@ func (m *awsRestjson1_serializeOpStartMetricsExport) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMetricsExportInput(v *StartMetricsExportInput, encoder *httpbinding.Encoder) error {
@@ -5186,10 +5020,8 @@ func (m *awsRestjson1_serializeOpStartResourceGrouping
 func (m *awsRestjson1_serializeOpStartResourceGroupingRecommendationTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5234,7 +5066,6 @@ func (m *awsRestjson1_serializeOpStartResourceGrouping
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartResourceGroupingRecommendationTaskInput(v *StartResourceGroupingRecommendationTaskInput, encoder *httpbinding.Encoder) error {
@@ -5267,10 +5098,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5319,7 +5148,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5363,10 +5191,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5404,7 +5230,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5440,10 +5265,8 @@ func (m *awsRestjson1_serializeOpUpdateApp) HandleSeri
 func (m *awsRestjson1_serializeOpUpdateApp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5488,7 +5311,6 @@ func (m *awsRestjson1_serializeOpUpdateApp) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAppInput(v *UpdateAppInput, encoder *httpbinding.Encoder) error {
@@ -5555,10 +5377,8 @@ func (m *awsRestjson1_serializeOpUpdateAppVersion) Han
 func (m *awsRestjson1_serializeOpUpdateAppVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5603,7 +5423,6 @@ func (m *awsRestjson1_serializeOpUpdateAppVersion) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAppVersionInput(v *UpdateAppVersionInput, encoder *httpbinding.Encoder) error {
@@ -5643,10 +5462,8 @@ func (m *awsRestjson1_serializeOpUpdateAppVersionAppCo
 func (m *awsRestjson1_serializeOpUpdateAppVersionAppComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5691,7 +5508,6 @@ func (m *awsRestjson1_serializeOpUpdateAppVersionAppCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAppVersionAppComponentInput(v *UpdateAppVersionAppComponentInput, encoder *httpbinding.Encoder) error {
@@ -5746,10 +5562,8 @@ func (m *awsRestjson1_serializeOpUpdateAppVersionResou
 func (m *awsRestjson1_serializeOpUpdateAppVersionResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5794,7 +5608,6 @@ func (m *awsRestjson1_serializeOpUpdateAppVersionResou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAppVersionResourceInput(v *UpdateAppVersionResourceInput, encoder *httpbinding.Encoder) error {
@@ -5878,10 +5691,8 @@ func (m *awsRestjson1_serializeOpUpdateResiliencyPolic
 func (m *awsRestjson1_serializeOpUpdateResiliencyPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5926,7 +5737,6 @@ func (m *awsRestjson1_serializeOpUpdateResiliencyPolic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResiliencyPolicyInput(v *UpdateResiliencyPolicyInput, encoder *httpbinding.Encoder) error {
