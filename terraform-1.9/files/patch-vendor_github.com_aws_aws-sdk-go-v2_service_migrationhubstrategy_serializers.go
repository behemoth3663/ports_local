--- vendor/github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpGetApplicationCompone
 func (m *awsRestjson1_serializeOpGetApplicationComponentDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpGetApplicationCompone
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApplicationComponentDetailsInput(v *GetApplicationComponentDetailsInput, encoder *httpbinding.Encoder) error {
@@ -96,10 +92,8 @@ func (m *awsRestjson1_serializeOpGetApplicationCompone
 func (m *awsRestjson1_serializeOpGetApplicationComponentStrategies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -137,7 +131,6 @@ func (m *awsRestjson1_serializeOpGetApplicationCompone
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApplicationComponentStrategiesInput(v *GetApplicationComponentStrategiesInput, encoder *httpbinding.Encoder) error {
@@ -167,10 +160,8 @@ func (m *awsRestjson1_serializeOpGetAssessment) Handle
 func (m *awsRestjson1_serializeOpGetAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -208,7 +199,6 @@ func (m *awsRestjson1_serializeOpGetAssessment) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssessmentInput(v *GetAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -238,10 +228,8 @@ func (m *awsRestjson1_serializeOpGetImportFileTask) Ha
 func (m *awsRestjson1_serializeOpGetImportFileTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -279,7 +267,6 @@ func (m *awsRestjson1_serializeOpGetImportFileTask) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImportFileTaskInput(v *GetImportFileTaskInput, encoder *httpbinding.Encoder) error {
@@ -309,10 +296,8 @@ func (m *awsRestjson1_serializeOpGetLatestAssessmentId
 func (m *awsRestjson1_serializeOpGetLatestAssessmentId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -346,7 +331,6 @@ func (m *awsRestjson1_serializeOpGetLatestAssessmentId
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLatestAssessmentIdInput(v *GetLatestAssessmentIdInput, encoder *httpbinding.Encoder) error {
@@ -367,10 +351,8 @@ func (m *awsRestjson1_serializeOpGetPortfolioPreferenc
 func (m *awsRestjson1_serializeOpGetPortfolioPreferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -404,7 +386,6 @@ func (m *awsRestjson1_serializeOpGetPortfolioPreferenc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPortfolioPreferencesInput(v *GetPortfolioPreferencesInput, encoder *httpbinding.Encoder) error {
@@ -425,10 +406,8 @@ func (m *awsRestjson1_serializeOpGetPortfolioSummary) 
 func (m *awsRestjson1_serializeOpGetPortfolioSummary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -462,7 +441,6 @@ func (m *awsRestjson1_serializeOpGetPortfolioSummary) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPortfolioSummaryInput(v *GetPortfolioSummaryInput, encoder *httpbinding.Encoder) error {
@@ -483,10 +461,8 @@ func (m *awsRestjson1_serializeOpGetRecommendationRepo
 func (m *awsRestjson1_serializeOpGetRecommendationReportDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -524,7 +500,6 @@ func (m *awsRestjson1_serializeOpGetRecommendationRepo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRecommendationReportDetailsInput(v *GetRecommendationReportDetailsInput, encoder *httpbinding.Encoder) error {
@@ -554,10 +529,8 @@ func (m *awsRestjson1_serializeOpGetServerDetails) Han
 func (m *awsRestjson1_serializeOpGetServerDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -595,7 +568,6 @@ func (m *awsRestjson1_serializeOpGetServerDetails) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetServerDetailsInput(v *GetServerDetailsInput, encoder *httpbinding.Encoder) error {
@@ -633,10 +605,8 @@ func (m *awsRestjson1_serializeOpGetServerStrategies) 
 func (m *awsRestjson1_serializeOpGetServerStrategies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -674,7 +644,6 @@ func (m *awsRestjson1_serializeOpGetServerStrategies) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetServerStrategiesInput(v *GetServerStrategiesInput, encoder *httpbinding.Encoder) error {
@@ -704,10 +673,8 @@ func (m *awsRestjson1_serializeOpListAnalyzableServers
 func (m *awsRestjson1_serializeOpListAnalyzableServers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -752,7 +719,6 @@ func (m *awsRestjson1_serializeOpListAnalyzableServers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnalyzableServersInput(v *ListAnalyzableServersInput, encoder *httpbinding.Encoder) error {
@@ -795,10 +761,8 @@ func (m *awsRestjson1_serializeOpListApplicationCompon
 func (m *awsRestjson1_serializeOpListApplicationComponents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -843,7 +807,6 @@ func (m *awsRestjson1_serializeOpListApplicationCompon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListApplicationComponentsInput(v *ListApplicationComponentsInput, encoder *httpbinding.Encoder) error {
@@ -903,10 +866,8 @@ func (m *awsRestjson1_serializeOpListCollectors) Handl
 func (m *awsRestjson1_serializeOpListCollectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -944,7 +905,6 @@ func (m *awsRestjson1_serializeOpListCollectors) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCollectorsInput(v *ListCollectorsInput, encoder *httpbinding.Encoder) error {
@@ -973,10 +933,8 @@ func (m *awsRestjson1_serializeOpListImportFileTask) H
 func (m *awsRestjson1_serializeOpListImportFileTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1014,7 +972,6 @@ func (m *awsRestjson1_serializeOpListImportFileTask) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImportFileTaskInput(v *ListImportFileTaskInput, encoder *httpbinding.Encoder) error {
@@ -1043,10 +1000,8 @@ func (m *awsRestjson1_serializeOpListServers) HandleSe
 func (m *awsRestjson1_serializeOpListServers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1091,7 +1046,6 @@ func (m *awsRestjson1_serializeOpListServers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListServersInput(v *ListServersInput, encoder *httpbinding.Encoder) error {
@@ -1151,10 +1105,8 @@ func (m *awsRestjson1_serializeOpPutPortfolioPreferenc
 func (m *awsRestjson1_serializeOpPutPortfolioPreferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1199,7 +1151,6 @@ func (m *awsRestjson1_serializeOpPutPortfolioPreferenc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutPortfolioPreferencesInput(v *PutPortfolioPreferencesInput, encoder *httpbinding.Encoder) error {
@@ -1253,10 +1204,8 @@ func (m *awsRestjson1_serializeOpStartAssessment) Hand
 func (m *awsRestjson1_serializeOpStartAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1301,7 +1250,6 @@ func (m *awsRestjson1_serializeOpStartAssessment) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAssessmentInput(v *StartAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -1351,10 +1299,8 @@ func (m *awsRestjson1_serializeOpStartImportFileTask) 
 func (m *awsRestjson1_serializeOpStartImportFileTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1399,7 +1345,6 @@ func (m *awsRestjson1_serializeOpStartImportFileTask) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartImportFileTaskInput(v *StartImportFileTaskInput, encoder *httpbinding.Encoder) error {
@@ -1459,10 +1404,8 @@ func (m *awsRestjson1_serializeOpStartRecommendationRe
 func (m *awsRestjson1_serializeOpStartRecommendationReportGeneration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1507,7 +1450,6 @@ func (m *awsRestjson1_serializeOpStartRecommendationRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartRecommendationReportGenerationInput(v *StartRecommendationReportGenerationInput, encoder *httpbinding.Encoder) error {
@@ -1547,10 +1489,8 @@ func (m *awsRestjson1_serializeOpStopAssessment) Handl
 func (m *awsRestjson1_serializeOpStopAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1595,7 +1535,6 @@ func (m *awsRestjson1_serializeOpStopAssessment) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopAssessmentInput(v *StopAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -1628,10 +1567,8 @@ func (m *awsRestjson1_serializeOpUpdateApplicationComp
 func (m *awsRestjson1_serializeOpUpdateApplicationComponentConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1676,7 +1613,6 @@ func (m *awsRestjson1_serializeOpUpdateApplicationComp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApplicationComponentConfigInput(v *UpdateApplicationComponentConfigInput, encoder *httpbinding.Encoder) error {
@@ -1743,10 +1679,8 @@ func (m *awsRestjson1_serializeOpUpdateServerConfig) H
 func (m *awsRestjson1_serializeOpUpdateServerConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1791,7 +1725,6 @@ func (m *awsRestjson1_serializeOpUpdateServerConfig) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateServerConfigInput(v *UpdateServerConfigInput, encoder *httpbinding.Encoder) error {
