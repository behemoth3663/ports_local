--- vendor/github.com/aws/aws-sdk-go-v2/service/devopsguru/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/devopsguru/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAddNotificationChanne
 func (m *awsRestjson1_serializeOpAddNotificationChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpAddNotificationChanne
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddNotificationChannelInput(v *AddNotificationChannelInput, encoder *httpbinding.Encoder) error {
@@ -109,10 +105,8 @@ func (m *awsRestjson1_serializeOpDeleteInsight) Handle
 func (m *awsRestjson1_serializeOpDeleteInsight) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -150,7 +144,6 @@ func (m *awsRestjson1_serializeOpDeleteInsight) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInsightInput(v *DeleteInsightInput, encoder *httpbinding.Encoder) error {
@@ -180,10 +173,8 @@ func (m *awsRestjson1_serializeOpDescribeAccountHealth
 func (m *awsRestjson1_serializeOpDescribeAccountHealth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -217,7 +208,6 @@ func (m *awsRestjson1_serializeOpDescribeAccountHealth
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccountHealthInput(v *DescribeAccountHealthInput, encoder *httpbinding.Encoder) error {
@@ -238,10 +228,8 @@ func (m *awsRestjson1_serializeOpDescribeAccountOvervi
 func (m *awsRestjson1_serializeOpDescribeAccountOverview) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -286,7 +274,6 @@ func (m *awsRestjson1_serializeOpDescribeAccountOvervi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAccountOverviewInput(v *DescribeAccountOverviewInput, encoder *httpbinding.Encoder) error {
@@ -324,10 +311,8 @@ func (m *awsRestjson1_serializeOpDescribeAnomaly) Hand
 func (m *awsRestjson1_serializeOpDescribeAnomaly) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -365,7 +350,6 @@ func (m *awsRestjson1_serializeOpDescribeAnomaly) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAnomalyInput(v *DescribeAnomalyInput, encoder *httpbinding.Encoder) error {
@@ -399,10 +383,8 @@ func (m *awsRestjson1_serializeOpDescribeEventSourcesC
 func (m *awsRestjson1_serializeOpDescribeEventSourcesConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -436,7 +418,6 @@ func (m *awsRestjson1_serializeOpDescribeEventSourcesC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeEventSourcesConfigInput(v *DescribeEventSourcesConfigInput, encoder *httpbinding.Encoder) error {
@@ -457,10 +438,8 @@ func (m *awsRestjson1_serializeOpDescribeFeedback) Han
 func (m *awsRestjson1_serializeOpDescribeFeedback) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -505,7 +484,6 @@ func (m *awsRestjson1_serializeOpDescribeFeedback) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeFeedbackInput(v *DescribeFeedbackInput, encoder *httpbinding.Encoder) error {
@@ -538,10 +516,8 @@ func (m *awsRestjson1_serializeOpDescribeInsight) Hand
 func (m *awsRestjson1_serializeOpDescribeInsight) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -579,7 +555,6 @@ func (m *awsRestjson1_serializeOpDescribeInsight) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInsightInput(v *DescribeInsightInput, encoder *httpbinding.Encoder) error {
@@ -613,10 +588,8 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationH
 func (m *awsRestjson1_serializeOpDescribeOrganizationHealth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -661,7 +634,6 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationH
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOrganizationHealthInput(v *DescribeOrganizationHealthInput, encoder *httpbinding.Encoder) error {
@@ -703,10 +675,8 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationO
 func (m *awsRestjson1_serializeOpDescribeOrganizationOverview) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -751,7 +721,6 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOrganizationOverviewInput(v *DescribeOrganizationOverviewInput, encoder *httpbinding.Encoder) error {
@@ -803,10 +772,8 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationR
 func (m *awsRestjson1_serializeOpDescribeOrganizationResourceCollectionHealth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -851,7 +818,6 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationR
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOrganizationResourceCollectionHealthInput(v *DescribeOrganizationResourceCollectionHealthInput, encoder *httpbinding.Encoder) error {
@@ -908,10 +874,8 @@ func (m *awsRestjson1_serializeOpDescribeResourceColle
 func (m *awsRestjson1_serializeOpDescribeResourceCollectionHealth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -949,7 +913,6 @@ func (m *awsRestjson1_serializeOpDescribeResourceColle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeResourceCollectionHealthInput(v *DescribeResourceCollectionHealthInput, encoder *httpbinding.Encoder) error {
@@ -983,10 +946,8 @@ func (m *awsRestjson1_serializeOpDescribeServiceIntegr
 func (m *awsRestjson1_serializeOpDescribeServiceIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1020,7 +981,6 @@ func (m *awsRestjson1_serializeOpDescribeServiceIntegr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeServiceIntegrationInput(v *DescribeServiceIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -1041,10 +1001,8 @@ func (m *awsRestjson1_serializeOpGetCostEstimation) Ha
 func (m *awsRestjson1_serializeOpGetCostEstimation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1082,7 +1040,6 @@ func (m *awsRestjson1_serializeOpGetCostEstimation) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCostEstimationInput(v *GetCostEstimationInput, encoder *httpbinding.Encoder) error {
@@ -1107,10 +1064,8 @@ func (m *awsRestjson1_serializeOpGetResourceCollection
 func (m *awsRestjson1_serializeOpGetResourceCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1148,7 +1103,6 @@ func (m *awsRestjson1_serializeOpGetResourceCollection
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceCollectionInput(v *GetResourceCollectionInput, encoder *httpbinding.Encoder) error {
@@ -1182,10 +1136,8 @@ func (m *awsRestjson1_serializeOpListAnomaliesForInsig
 func (m *awsRestjson1_serializeOpListAnomaliesForInsight) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1234,7 +1186,6 @@ func (m *awsRestjson1_serializeOpListAnomaliesForInsig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnomaliesForInsightInput(v *ListAnomaliesForInsightInput, encoder *httpbinding.Encoder) error {
@@ -1300,10 +1251,8 @@ func (m *awsRestjson1_serializeOpListAnomalousLogGroup
 func (m *awsRestjson1_serializeOpListAnomalousLogGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1348,7 +1297,6 @@ func (m *awsRestjson1_serializeOpListAnomalousLogGroup
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnomalousLogGroupsInput(v *ListAnomalousLogGroupsInput, encoder *httpbinding.Encoder) error {
@@ -1391,10 +1339,8 @@ func (m *awsRestjson1_serializeOpListEvents) HandleSer
 func (m *awsRestjson1_serializeOpListEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1439,7 +1385,6 @@ func (m *awsRestjson1_serializeOpListEvents) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventsInput(v *ListEventsInput, encoder *httpbinding.Encoder) error {
@@ -1489,10 +1434,8 @@ func (m *awsRestjson1_serializeOpListInsights) HandleS
 func (m *awsRestjson1_serializeOpListInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1537,7 +1480,6 @@ func (m *awsRestjson1_serializeOpListInsights) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInsightsInput(v *ListInsightsInput, encoder *httpbinding.Encoder) error {
@@ -1582,10 +1524,8 @@ func (m *awsRestjson1_serializeOpListMonitoredResource
 func (m *awsRestjson1_serializeOpListMonitoredResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1630,7 +1570,6 @@ func (m *awsRestjson1_serializeOpListMonitoredResource
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMonitoredResourcesInput(v *ListMonitoredResourcesInput, encoder *httpbinding.Encoder) error {
@@ -1675,10 +1614,8 @@ func (m *awsRestjson1_serializeOpListNotificationChann
 func (m *awsRestjson1_serializeOpListNotificationChannels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1723,7 +1660,6 @@ func (m *awsRestjson1_serializeOpListNotificationChann
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNotificationChannelsInput(v *ListNotificationChannelsInput, encoder *httpbinding.Encoder) error {
@@ -1756,10 +1692,8 @@ func (m *awsRestjson1_serializeOpListOrganizationInsig
 func (m *awsRestjson1_serializeOpListOrganizationInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1804,7 +1738,6 @@ func (m *awsRestjson1_serializeOpListOrganizationInsig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationInsightsInput(v *ListOrganizationInsightsInput, encoder *httpbinding.Encoder) error {
@@ -1863,10 +1796,8 @@ func (m *awsRestjson1_serializeOpListRecommendations) 
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
@@ -1911,7 +1842,6 @@ func (m *awsRestjson1_serializeOpListRecommendations) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRecommendationsInput(v *ListRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -1959,10 +1889,8 @@ func (m *awsRestjson1_serializeOpPutFeedback) HandleSe
 func (m *awsRestjson1_serializeOpPutFeedback) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2007,7 +1935,6 @@ func (m *awsRestjson1_serializeOpPutFeedback) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutFeedbackInput(v *PutFeedbackInput, encoder *httpbinding.Encoder) error {
@@ -2042,10 +1969,8 @@ func (m *awsRestjson1_serializeOpRemoveNotificationCha
 func (m *awsRestjson1_serializeOpRemoveNotificationChannel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2083,7 +2008,6 @@ func (m *awsRestjson1_serializeOpRemoveNotificationCha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveNotificationChannelInput(v *RemoveNotificationChannelInput, encoder *httpbinding.Encoder) error {
@@ -2113,10 +2037,8 @@ func (m *awsRestjson1_serializeOpSearchInsights) Handl
 func (m *awsRestjson1_serializeOpSearchInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2161,7 +2083,6 @@ func (m *awsRestjson1_serializeOpSearchInsights) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchInsightsInput(v *SearchInsightsInput, encoder *httpbinding.Encoder) error {
@@ -2218,10 +2139,8 @@ func (m *awsRestjson1_serializeOpSearchOrganizationIns
 func (m *awsRestjson1_serializeOpSearchOrganizationInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2266,7 +2185,6 @@ func (m *awsRestjson1_serializeOpSearchOrganizationIns
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchOrganizationInsightsInput(v *SearchOrganizationInsightsInput, encoder *httpbinding.Encoder) error {
@@ -2330,10 +2248,8 @@ func (m *awsRestjson1_serializeOpStartCostEstimation) 
 func (m *awsRestjson1_serializeOpStartCostEstimation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2378,7 +2294,6 @@ func (m *awsRestjson1_serializeOpStartCostEstimation) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartCostEstimationInput(v *StartCostEstimationInput, encoder *httpbinding.Encoder) error {
@@ -2418,10 +2333,8 @@ func (m *awsRestjson1_serializeOpUpdateEventSourcesCon
 func (m *awsRestjson1_serializeOpUpdateEventSourcesConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2466,7 +2379,6 @@ func (m *awsRestjson1_serializeOpUpdateEventSourcesCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventSourcesConfigInput(v *UpdateEventSourcesConfigInput, encoder *httpbinding.Encoder) error {
@@ -2501,10 +2413,8 @@ func (m *awsRestjson1_serializeOpUpdateResourceCollect
 func (m *awsRestjson1_serializeOpUpdateResourceCollection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2549,7 +2459,6 @@ func (m *awsRestjson1_serializeOpUpdateResourceCollect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceCollectionInput(v *UpdateResourceCollectionInput, encoder *httpbinding.Encoder) error {
@@ -2589,10 +2498,8 @@ func (m *awsRestjson1_serializeOpUpdateServiceIntegrat
 func (m *awsRestjson1_serializeOpUpdateServiceIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2637,7 +2544,6 @@ func (m *awsRestjson1_serializeOpUpdateServiceIntegrat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateServiceIntegrationInput(v *UpdateServiceIntegrationInput, encoder *httpbinding.Encoder) error {
