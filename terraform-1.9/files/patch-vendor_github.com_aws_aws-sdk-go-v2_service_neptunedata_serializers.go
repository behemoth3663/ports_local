--- vendor/github.com/aws/aws-sdk-go-v2/service/neptunedata/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/neptunedata/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpCancelGremlinQuery) H
 func (m *awsRestjson1_serializeOpCancelGremlinQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpCancelGremlinQuery) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelGremlinQueryInput(v *CancelGremlinQueryInput, encoder *httpbinding.Encoder) error {
@@ -96,10 +92,8 @@ func (m *awsRestjson1_serializeOpCancelLoaderJob) Hand
 func (m *awsRestjson1_serializeOpCancelLoaderJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -137,7 +131,6 @@ func (m *awsRestjson1_serializeOpCancelLoaderJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelLoaderJobInput(v *CancelLoaderJobInput, encoder *httpbinding.Encoder) error {
@@ -167,10 +160,8 @@ func (m *awsRestjson1_serializeOpCancelMLDataProcessin
 func (m *awsRestjson1_serializeOpCancelMLDataProcessingJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -208,7 +199,6 @@ func (m *awsRestjson1_serializeOpCancelMLDataProcessin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelMLDataProcessingJobInput(v *CancelMLDataProcessingJobInput, encoder *httpbinding.Encoder) error {
@@ -246,10 +236,8 @@ func (m *awsRestjson1_serializeOpCancelMLModelTraining
 func (m *awsRestjson1_serializeOpCancelMLModelTrainingJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -287,7 +275,6 @@ func (m *awsRestjson1_serializeOpCancelMLModelTraining
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelMLModelTrainingJobInput(v *CancelMLModelTrainingJobInput, encoder *httpbinding.Encoder) error {
@@ -325,10 +312,8 @@ func (m *awsRestjson1_serializeOpCancelMLModelTransfor
 func (m *awsRestjson1_serializeOpCancelMLModelTransformJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -366,7 +351,6 @@ func (m *awsRestjson1_serializeOpCancelMLModelTransfor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelMLModelTransformJobInput(v *CancelMLModelTransformJobInput, encoder *httpbinding.Encoder) error {
@@ -404,10 +388,8 @@ func (m *awsRestjson1_serializeOpCancelOpenCypherQuery
 func (m *awsRestjson1_serializeOpCancelOpenCypherQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -445,7 +427,6 @@ func (m *awsRestjson1_serializeOpCancelOpenCypherQuery
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelOpenCypherQueryInput(v *CancelOpenCypherQueryInput, encoder *httpbinding.Encoder) error {
@@ -479,10 +460,8 @@ func (m *awsRestjson1_serializeOpCreateMLEndpoint) Han
 func (m *awsRestjson1_serializeOpCreateMLEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -527,7 +506,6 @@ func (m *awsRestjson1_serializeOpCreateMLEndpoint) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMLEndpointInput(v *CreateMLEndpointInput, encoder *httpbinding.Encoder) error {
@@ -600,10 +578,8 @@ func (m *awsRestjson1_serializeOpDeleteMLEndpoint) Han
 func (m *awsRestjson1_serializeOpDeleteMLEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -641,7 +617,6 @@ func (m *awsRestjson1_serializeOpDeleteMLEndpoint) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMLEndpointInput(v *DeleteMLEndpointInput, encoder *httpbinding.Encoder) error {
@@ -679,10 +654,8 @@ func (m *awsRestjson1_serializeOpDeletePropertygraphSt
 func (m *awsRestjson1_serializeOpDeletePropertygraphStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -716,7 +689,6 @@ func (m *awsRestjson1_serializeOpDeletePropertygraphSt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePropertygraphStatisticsInput(v *DeletePropertygraphStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -737,10 +709,8 @@ func (m *awsRestjson1_serializeOpDeleteSparqlStatistic
 func (m *awsRestjson1_serializeOpDeleteSparqlStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -774,7 +744,6 @@ func (m *awsRestjson1_serializeOpDeleteSparqlStatistic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSparqlStatisticsInput(v *DeleteSparqlStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -795,10 +764,8 @@ func (m *awsRestjson1_serializeOpExecuteFastReset) Han
 func (m *awsRestjson1_serializeOpExecuteFastReset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -843,7 +810,6 @@ func (m *awsRestjson1_serializeOpExecuteFastReset) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteFastResetInput(v *ExecuteFastResetInput, encoder *httpbinding.Encoder) error {
@@ -881,10 +847,8 @@ func (m *awsRestjson1_serializeOpExecuteGremlinExplain
 func (m *awsRestjson1_serializeOpExecuteGremlinExplainQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -929,7 +893,6 @@ func (m *awsRestjson1_serializeOpExecuteGremlinExplain
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteGremlinExplainQueryInput(v *ExecuteGremlinExplainQueryInput, encoder *httpbinding.Encoder) error {
@@ -962,10 +925,8 @@ func (m *awsRestjson1_serializeOpExecuteGremlinProfile
 func (m *awsRestjson1_serializeOpExecuteGremlinProfileQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1010,7 +971,6 @@ func (m *awsRestjson1_serializeOpExecuteGremlinProfile
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteGremlinProfileQueryInput(v *ExecuteGremlinProfileQueryInput, encoder *httpbinding.Encoder) error {
@@ -1063,10 +1023,8 @@ func (m *awsRestjson1_serializeOpExecuteGremlinQuery) 
 func (m *awsRestjson1_serializeOpExecuteGremlinQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1115,7 +1073,6 @@ func (m *awsRestjson1_serializeOpExecuteGremlinQuery) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteGremlinQueryInput(v *ExecuteGremlinQueryInput, encoder *httpbinding.Encoder) error {
@@ -1153,10 +1110,8 @@ func (m *awsRestjson1_serializeOpExecuteOpenCypherExpl
 func (m *awsRestjson1_serializeOpExecuteOpenCypherExplainQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1201,7 +1156,6 @@ func (m *awsRestjson1_serializeOpExecuteOpenCypherExpl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteOpenCypherExplainQueryInput(v *ExecuteOpenCypherExplainQueryInput, encoder *httpbinding.Encoder) error {
@@ -1244,10 +1198,8 @@ func (m *awsRestjson1_serializeOpExecuteOpenCypherQuer
 func (m *awsRestjson1_serializeOpExecuteOpenCypherQuery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1292,7 +1244,6 @@ func (m *awsRestjson1_serializeOpExecuteOpenCypherQuer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteOpenCypherQueryInput(v *ExecuteOpenCypherQueryInput, encoder *httpbinding.Encoder) error {
@@ -1330,10 +1281,8 @@ func (m *awsRestjson1_serializeOpGetEngineStatus) Hand
 func (m *awsRestjson1_serializeOpGetEngineStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1367,7 +1316,6 @@ func (m *awsRestjson1_serializeOpGetEngineStatus) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEngineStatusInput(v *GetEngineStatusInput, encoder *httpbinding.Encoder) error {
@@ -1388,10 +1336,8 @@ func (m *awsRestjson1_serializeOpGetGremlinQueryStatus
 func (m *awsRestjson1_serializeOpGetGremlinQueryStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1429,7 +1375,6 @@ func (m *awsRestjson1_serializeOpGetGremlinQueryStatus
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGremlinQueryStatusInput(v *GetGremlinQueryStatusInput, encoder *httpbinding.Encoder) error {
@@ -1459,10 +1404,8 @@ func (m *awsRestjson1_serializeOpGetLoaderJobStatus) H
 func (m *awsRestjson1_serializeOpGetLoaderJobStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1500,7 +1443,6 @@ func (m *awsRestjson1_serializeOpGetLoaderJobStatus) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLoaderJobStatusInput(v *GetLoaderJobStatusInput, encoder *httpbinding.Encoder) error {
@@ -1546,10 +1488,8 @@ func (m *awsRestjson1_serializeOpGetMLDataProcessingJo
 func (m *awsRestjson1_serializeOpGetMLDataProcessingJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1587,7 +1527,6 @@ func (m *awsRestjson1_serializeOpGetMLDataProcessingJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMLDataProcessingJobInput(v *GetMLDataProcessingJobInput, encoder *httpbinding.Encoder) error {
@@ -1621,10 +1560,8 @@ func (m *awsRestjson1_serializeOpGetMLEndpoint) Handle
 func (m *awsRestjson1_serializeOpGetMLEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1662,7 +1599,6 @@ func (m *awsRestjson1_serializeOpGetMLEndpoint) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMLEndpointInput(v *GetMLEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1696,10 +1632,8 @@ func (m *awsRestjson1_serializeOpGetMLModelTrainingJob
 func (m *awsRestjson1_serializeOpGetMLModelTrainingJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1737,7 +1671,6 @@ func (m *awsRestjson1_serializeOpGetMLModelTrainingJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMLModelTrainingJobInput(v *GetMLModelTrainingJobInput, encoder *httpbinding.Encoder) error {
@@ -1771,10 +1704,8 @@ func (m *awsRestjson1_serializeOpGetMLModelTransformJo
 func (m *awsRestjson1_serializeOpGetMLModelTransformJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1812,7 +1743,6 @@ func (m *awsRestjson1_serializeOpGetMLModelTransformJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMLModelTransformJobInput(v *GetMLModelTransformJobInput, encoder *httpbinding.Encoder) error {
@@ -1846,10 +1776,8 @@ func (m *awsRestjson1_serializeOpGetOpenCypherQuerySta
 func (m *awsRestjson1_serializeOpGetOpenCypherQueryStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1887,7 +1815,6 @@ func (m *awsRestjson1_serializeOpGetOpenCypherQuerySta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOpenCypherQueryStatusInput(v *GetOpenCypherQueryStatusInput, encoder *httpbinding.Encoder) error {
@@ -1917,10 +1844,8 @@ func (m *awsRestjson1_serializeOpGetPropertygraphStati
 func (m *awsRestjson1_serializeOpGetPropertygraphStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1954,7 +1879,6 @@ func (m *awsRestjson1_serializeOpGetPropertygraphStati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPropertygraphStatisticsInput(v *GetPropertygraphStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -1975,10 +1899,8 @@ func (m *awsRestjson1_serializeOpGetPropertygraphStrea
 func (m *awsRestjson1_serializeOpGetPropertygraphStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2016,7 +1938,6 @@ func (m *awsRestjson1_serializeOpGetPropertygraphStrea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPropertygraphStreamInput(v *GetPropertygraphStreamInput, encoder *httpbinding.Encoder) error {
@@ -2058,10 +1979,8 @@ func (m *awsRestjson1_serializeOpGetPropertygraphSumma
 func (m *awsRestjson1_serializeOpGetPropertygraphSummary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2099,7 +2018,6 @@ func (m *awsRestjson1_serializeOpGetPropertygraphSumma
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPropertygraphSummaryInput(v *GetPropertygraphSummaryInput, encoder *httpbinding.Encoder) error {
@@ -2124,10 +2042,8 @@ func (m *awsRestjson1_serializeOpGetRDFGraphSummary) H
 func (m *awsRestjson1_serializeOpGetRDFGraphSummary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2165,7 +2081,6 @@ func (m *awsRestjson1_serializeOpGetRDFGraphSummary) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRDFGraphSummaryInput(v *GetRDFGraphSummaryInput, encoder *httpbinding.Encoder) error {
@@ -2190,10 +2105,8 @@ func (m *awsRestjson1_serializeOpGetSparqlStatistics) 
 func (m *awsRestjson1_serializeOpGetSparqlStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2227,7 +2140,6 @@ func (m *awsRestjson1_serializeOpGetSparqlStatistics) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSparqlStatisticsInput(v *GetSparqlStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -2248,10 +2160,8 @@ func (m *awsRestjson1_serializeOpGetSparqlStream) Hand
 func (m *awsRestjson1_serializeOpGetSparqlStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2289,7 +2199,6 @@ func (m *awsRestjson1_serializeOpGetSparqlStream) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSparqlStreamInput(v *GetSparqlStreamInput, encoder *httpbinding.Encoder) error {
@@ -2331,10 +2240,8 @@ func (m *awsRestjson1_serializeOpListGremlinQueries) H
 func (m *awsRestjson1_serializeOpListGremlinQueries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2372,7 +2279,6 @@ func (m *awsRestjson1_serializeOpListGremlinQueries) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGremlinQueriesInput(v *ListGremlinQueriesInput, encoder *httpbinding.Encoder) error {
@@ -2397,10 +2303,8 @@ func (m *awsRestjson1_serializeOpListLoaderJobs) Handl
 func (m *awsRestjson1_serializeOpListLoaderJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2438,7 +2342,6 @@ func (m *awsRestjson1_serializeOpListLoaderJobs) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLoaderJobsInput(v *ListLoaderJobsInput, encoder *httpbinding.Encoder) error {
@@ -2467,10 +2370,8 @@ func (m *awsRestjson1_serializeOpListMLDataProcessingJ
 func (m *awsRestjson1_serializeOpListMLDataProcessingJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2508,7 +2409,6 @@ func (m *awsRestjson1_serializeOpListMLDataProcessingJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMLDataProcessingJobsInput(v *ListMLDataProcessingJobsInput, encoder *httpbinding.Encoder) error {
@@ -2537,10 +2437,8 @@ func (m *awsRestjson1_serializeOpListMLEndpoints) Hand
 func (m *awsRestjson1_serializeOpListMLEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2578,7 +2476,6 @@ func (m *awsRestjson1_serializeOpListMLEndpoints) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMLEndpointsInput(v *ListMLEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -2607,10 +2504,8 @@ func (m *awsRestjson1_serializeOpListMLModelTrainingJo
 func (m *awsRestjson1_serializeOpListMLModelTrainingJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2648,7 +2543,6 @@ func (m *awsRestjson1_serializeOpListMLModelTrainingJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMLModelTrainingJobsInput(v *ListMLModelTrainingJobsInput, encoder *httpbinding.Encoder) error {
@@ -2677,10 +2571,8 @@ func (m *awsRestjson1_serializeOpListMLModelTransformJ
 func (m *awsRestjson1_serializeOpListMLModelTransformJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2718,7 +2610,6 @@ func (m *awsRestjson1_serializeOpListMLModelTransformJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMLModelTransformJobsInput(v *ListMLModelTransformJobsInput, encoder *httpbinding.Encoder) error {
@@ -2747,10 +2638,8 @@ func (m *awsRestjson1_serializeOpListOpenCypherQueries
 func (m *awsRestjson1_serializeOpListOpenCypherQueries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2788,7 +2677,6 @@ func (m *awsRestjson1_serializeOpListOpenCypherQueries
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOpenCypherQueriesInput(v *ListOpenCypherQueriesInput, encoder *httpbinding.Encoder) error {
@@ -2813,10 +2701,8 @@ func (m *awsRestjson1_serializeOpManagePropertygraphSt
 func (m *awsRestjson1_serializeOpManagePropertygraphStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2861,7 +2747,6 @@ func (m *awsRestjson1_serializeOpManagePropertygraphSt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsManagePropertygraphStatisticsInput(v *ManagePropertygraphStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -2894,10 +2779,8 @@ func (m *awsRestjson1_serializeOpManageSparqlStatistic
 func (m *awsRestjson1_serializeOpManageSparqlStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2942,7 +2825,6 @@ func (m *awsRestjson1_serializeOpManageSparqlStatistic
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsManageSparqlStatisticsInput(v *ManageSparqlStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -2975,10 +2857,8 @@ func (m *awsRestjson1_serializeOpStartLoaderJob) Handl
 func (m *awsRestjson1_serializeOpStartLoaderJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3023,7 +2903,6 @@ func (m *awsRestjson1_serializeOpStartLoaderJob) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartLoaderJobInput(v *StartLoaderJobInput, encoder *httpbinding.Encoder) error {
@@ -3115,10 +2994,8 @@ func (m *awsRestjson1_serializeOpStartMLDataProcessing
 func (m *awsRestjson1_serializeOpStartMLDataProcessingJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3163,7 +3040,6 @@ func (m *awsRestjson1_serializeOpStartMLDataProcessing
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMLDataProcessingJobInput(v *StartMLDataProcessingJobInput, encoder *httpbinding.Encoder) error {
@@ -3270,10 +3146,8 @@ func (m *awsRestjson1_serializeOpStartMLModelTrainingJ
 func (m *awsRestjson1_serializeOpStartMLModelTrainingJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3318,7 +3192,6 @@ func (m *awsRestjson1_serializeOpStartMLModelTrainingJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMLModelTrainingJobInput(v *StartMLModelTrainingJobInput, encoder *httpbinding.Encoder) error {
@@ -3442,10 +3315,8 @@ func (m *awsRestjson1_serializeOpStartMLModelTransform
 func (m *awsRestjson1_serializeOpStartMLModelTransformJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3490,7 +3361,6 @@ func (m *awsRestjson1_serializeOpStartMLModelTransform
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMLModelTransformJobInput(v *StartMLModelTransformJobInput, encoder *httpbinding.Encoder) error {
