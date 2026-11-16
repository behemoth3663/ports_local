--- vendor/github.com/aws/aws-sdk-go-v2/service/codecatalyst/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/codecatalyst/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCreateAccessToken) Ha
 func (m *awsRestjson1_serializeOpCreateAccessToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpCreateAccessToken) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAccessTokenInput(v *CreateAccessTokenInput, encoder *httpbinding.Encoder) error {
@@ -112,10 +108,8 @@ func (m *awsRestjson1_serializeOpCreateDevEnvironment)
 func (m *awsRestjson1_serializeOpCreateDevEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -164,7 +158,6 @@ func (m *awsRestjson1_serializeOpCreateDevEnvironment)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDevEnvironmentInput(v *CreateDevEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -256,10 +249,8 @@ func (m *awsRestjson1_serializeOpCreateProject) Handle
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
@@ -308,7 +299,6 @@ func (m *awsRestjson1_serializeOpCreateProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProjectInput(v *CreateProjectInput, encoder *httpbinding.Encoder) error {
@@ -355,10 +345,8 @@ func (m *awsRestjson1_serializeOpCreateSourceRepositor
 func (m *awsRestjson1_serializeOpCreateSourceRepository) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -407,7 +395,6 @@ func (m *awsRestjson1_serializeOpCreateSourceRepositor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSourceRepositoryInput(v *CreateSourceRepositoryInput, encoder *httpbinding.Encoder) error {
@@ -467,10 +454,8 @@ func (m *awsRestjson1_serializeOpCreateSourceRepositor
 func (m *awsRestjson1_serializeOpCreateSourceRepositoryBranch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -519,7 +504,6 @@ func (m *awsRestjson1_serializeOpCreateSourceRepositor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSourceRepositoryBranchInput(v *CreateSourceRepositoryBranchInput, encoder *httpbinding.Encoder) error {
@@ -588,10 +572,8 @@ func (m *awsRestjson1_serializeOpDeleteAccessToken) Ha
 func (m *awsRestjson1_serializeOpDeleteAccessToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -629,7 +611,6 @@ func (m *awsRestjson1_serializeOpDeleteAccessToken) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccessTokenInput(v *DeleteAccessTokenInput, encoder *httpbinding.Encoder) error {
@@ -659,10 +640,8 @@ func (m *awsRestjson1_serializeOpDeleteDevEnvironment)
 func (m *awsRestjson1_serializeOpDeleteDevEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -700,7 +679,6 @@ func (m *awsRestjson1_serializeOpDeleteDevEnvironment)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDevEnvironmentInput(v *DeleteDevEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -748,10 +726,8 @@ func (m *awsRestjson1_serializeOpDeleteProject) Handle
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
@@ -789,7 +765,6 @@ func (m *awsRestjson1_serializeOpDeleteProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProjectInput(v *DeleteProjectInput, encoder *httpbinding.Encoder) error {
@@ -828,10 +803,8 @@ func (m *awsRestjson1_serializeOpDeleteSourceRepositor
 func (m *awsRestjson1_serializeOpDeleteSourceRepository) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -869,7 +842,6 @@ func (m *awsRestjson1_serializeOpDeleteSourceRepositor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSourceRepositoryInput(v *DeleteSourceRepositoryInput, encoder *httpbinding.Encoder) error {
@@ -917,10 +889,8 @@ func (m *awsRestjson1_serializeOpDeleteSpace) HandleSe
 func (m *awsRestjson1_serializeOpDeleteSpace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -958,7 +928,6 @@ func (m *awsRestjson1_serializeOpDeleteSpace) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSpaceInput(v *DeleteSpaceInput, encoder *httpbinding.Encoder) error {
@@ -988,10 +957,8 @@ func (m *awsRestjson1_serializeOpGetDevEnvironment) Ha
 func (m *awsRestjson1_serializeOpGetDevEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1029,7 +996,6 @@ func (m *awsRestjson1_serializeOpGetDevEnvironment) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDevEnvironmentInput(v *GetDevEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -1077,10 +1043,8 @@ func (m *awsRestjson1_serializeOpGetProject) HandleSer
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
@@ -1118,7 +1082,6 @@ func (m *awsRestjson1_serializeOpGetProject) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProjectInput(v *GetProjectInput, encoder *httpbinding.Encoder) error {
@@ -1157,10 +1120,8 @@ func (m *awsRestjson1_serializeOpGetSourceRepository) 
 func (m *awsRestjson1_serializeOpGetSourceRepository) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1198,7 +1159,6 @@ func (m *awsRestjson1_serializeOpGetSourceRepository) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSourceRepositoryInput(v *GetSourceRepositoryInput, encoder *httpbinding.Encoder) error {
@@ -1246,10 +1206,8 @@ func (m *awsRestjson1_serializeOpGetSourceRepositoryCl
 func (m *awsRestjson1_serializeOpGetSourceRepositoryCloneUrls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1287,7 +1245,6 @@ func (m *awsRestjson1_serializeOpGetSourceRepositoryCl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSourceRepositoryCloneUrlsInput(v *GetSourceRepositoryCloneUrlsInput, encoder *httpbinding.Encoder) error {
@@ -1335,10 +1292,8 @@ func (m *awsRestjson1_serializeOpGetSpace) HandleSeria
 func (m *awsRestjson1_serializeOpGetSpace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1376,7 +1331,6 @@ func (m *awsRestjson1_serializeOpGetSpace) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSpaceInput(v *GetSpaceInput, encoder *httpbinding.Encoder) error {
@@ -1406,10 +1360,8 @@ func (m *awsRestjson1_serializeOpGetSubscription) Hand
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
@@ -1447,7 +1399,6 @@ func (m *awsRestjson1_serializeOpGetSubscription) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSubscriptionInput(v *GetSubscriptionInput, encoder *httpbinding.Encoder) error {
@@ -1477,10 +1428,8 @@ func (m *awsRestjson1_serializeOpGetUserDetails) Handl
 func (m *awsRestjson1_serializeOpGetUserDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1518,7 +1467,6 @@ func (m *awsRestjson1_serializeOpGetUserDetails) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserDetailsInput(v *GetUserDetailsInput, encoder *httpbinding.Encoder) error {
@@ -1547,10 +1495,8 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
 func (m *awsRestjson1_serializeOpGetWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1588,7 +1534,6 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowInput(v *GetWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -1636,10 +1581,8 @@ func (m *awsRestjson1_serializeOpGetWorkflowRun) Handl
 func (m *awsRestjson1_serializeOpGetWorkflowRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1677,7 +1620,6 @@ func (m *awsRestjson1_serializeOpGetWorkflowRun) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowRunInput(v *GetWorkflowRunInput, encoder *httpbinding.Encoder) error {
@@ -1725,10 +1667,8 @@ func (m *awsRestjson1_serializeOpListAccessTokens) Han
 func (m *awsRestjson1_serializeOpListAccessTokens) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1773,7 +1713,6 @@ func (m *awsRestjson1_serializeOpListAccessTokens) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAccessTokensInput(v *ListAccessTokensInput, encoder *httpbinding.Encoder) error {
@@ -1811,10 +1750,8 @@ func (m *awsRestjson1_serializeOpListDevEnvironments) 
 func (m *awsRestjson1_serializeOpListDevEnvironments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1863,7 +1800,6 @@ func (m *awsRestjson1_serializeOpListDevEnvironments) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDevEnvironmentsInput(v *ListDevEnvironmentsInput, encoder *httpbinding.Encoder) error {
@@ -1922,10 +1858,8 @@ func (m *awsRestjson1_serializeOpListDevEnvironmentSes
 func (m *awsRestjson1_serializeOpListDevEnvironmentSessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1974,7 +1908,6 @@ func (m *awsRestjson1_serializeOpListDevEnvironmentSes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDevEnvironmentSessionsInput(v *ListDevEnvironmentSessionsInput, encoder *httpbinding.Encoder) error {
@@ -2039,10 +1972,8 @@ func (m *awsRestjson1_serializeOpListEventLogs) Handle
 func (m *awsRestjson1_serializeOpListEventLogs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2091,7 +2022,6 @@ func (m *awsRestjson1_serializeOpListEventLogs) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventLogsInput(v *ListEventLogsInput, encoder *httpbinding.Encoder) error {
@@ -2153,10 +2083,8 @@ func (m *awsRestjson1_serializeOpListProjects) HandleS
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
@@ -2205,7 +2133,6 @@ func (m *awsRestjson1_serializeOpListProjects) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProjectsInput(v *ListProjectsInput, encoder *httpbinding.Encoder) error {
@@ -2259,10 +2186,8 @@ func (m *awsRestjson1_serializeOpListSourceRepositorie
 func (m *awsRestjson1_serializeOpListSourceRepositories) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2311,7 +2236,6 @@ func (m *awsRestjson1_serializeOpListSourceRepositorie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSourceRepositoriesInput(v *ListSourceRepositoriesInput, encoder *httpbinding.Encoder) error {
@@ -2367,10 +2291,8 @@ func (m *awsRestjson1_serializeOpListSourceRepositoryB
 func (m *awsRestjson1_serializeOpListSourceRepositoryBranches) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2419,7 +2341,6 @@ func (m *awsRestjson1_serializeOpListSourceRepositoryB
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSourceRepositoryBranchesInput(v *ListSourceRepositoryBranchesInput, encoder *httpbinding.Encoder) error {
@@ -2484,10 +2405,8 @@ func (m *awsRestjson1_serializeOpListSpaces) HandleSer
 func (m *awsRestjson1_serializeOpListSpaces) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2532,7 +2451,6 @@ func (m *awsRestjson1_serializeOpListSpaces) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSpacesInput(v *ListSpacesInput, encoder *httpbinding.Encoder) error {
@@ -2565,10 +2483,8 @@ func (m *awsRestjson1_serializeOpListWorkflowRuns) Han
 func (m *awsRestjson1_serializeOpListWorkflowRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2617,7 +2533,6 @@ func (m *awsRestjson1_serializeOpListWorkflowRuns) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowRunsInput(v *ListWorkflowRunsInput, encoder *httpbinding.Encoder) error {
@@ -2682,10 +2597,8 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
 func (m *awsRestjson1_serializeOpListWorkflows) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2734,7 +2647,6 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowsInput(v *ListWorkflowsInput, encoder *httpbinding.Encoder) error {
@@ -2795,10 +2707,8 @@ func (m *awsRestjson1_serializeOpStartDevEnvironment) 
 func (m *awsRestjson1_serializeOpStartDevEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2847,7 +2757,6 @@ func (m *awsRestjson1_serializeOpStartDevEnvironment) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDevEnvironmentInput(v *StartDevEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -2919,10 +2828,8 @@ func (m *awsRestjson1_serializeOpStartDevEnvironmentSe
 func (m *awsRestjson1_serializeOpStartDevEnvironmentSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2971,7 +2878,6 @@ func (m *awsRestjson1_serializeOpStartDevEnvironmentSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDevEnvironmentSessionInput(v *StartDevEnvironmentSessionInput, encoder *httpbinding.Encoder) error {
@@ -3033,10 +2939,8 @@ func (m *awsRestjson1_serializeOpStartWorkflowRun) Han
 func (m *awsRestjson1_serializeOpStartWorkflowRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3085,7 +2989,6 @@ func (m *awsRestjson1_serializeOpStartWorkflowRun) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartWorkflowRunInput(v *StartWorkflowRunInput, encoder *httpbinding.Encoder) error {
@@ -3140,10 +3043,8 @@ func (m *awsRestjson1_serializeOpStopDevEnvironment) H
 func (m *awsRestjson1_serializeOpStopDevEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3181,7 +3082,6 @@ func (m *awsRestjson1_serializeOpStopDevEnvironment) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopDevEnvironmentInput(v *StopDevEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -3229,10 +3129,8 @@ func (m *awsRestjson1_serializeOpStopDevEnvironmentSes
 func (m *awsRestjson1_serializeOpStopDevEnvironmentSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3270,7 +3168,6 @@ func (m *awsRestjson1_serializeOpStopDevEnvironmentSes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopDevEnvironmentSessionInput(v *StopDevEnvironmentSessionInput, encoder *httpbinding.Encoder) error {
@@ -3327,10 +3224,8 @@ func (m *awsRestjson1_serializeOpUpdateDevEnvironment)
 func (m *awsRestjson1_serializeOpUpdateDevEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3379,7 +3274,6 @@ func (m *awsRestjson1_serializeOpUpdateDevEnvironment)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDevEnvironmentInput(v *UpdateDevEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -3461,10 +3355,8 @@ func (m *awsRestjson1_serializeOpUpdateProject) Handle
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
@@ -3513,7 +3405,6 @@ func (m *awsRestjson1_serializeOpUpdateProject) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProjectInput(v *UpdateProjectInput, encoder *httpbinding.Encoder) error {
@@ -3564,10 +3455,8 @@ func (m *awsRestjson1_serializeOpUpdateSpace) HandleSe
 func (m *awsRestjson1_serializeOpUpdateSpace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3616,7 +3505,6 @@ func (m *awsRestjson1_serializeOpUpdateSpace) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSpaceInput(v *UpdateSpaceInput, encoder *httpbinding.Encoder) error {
@@ -3658,10 +3546,8 @@ func (m *awsRestjson1_serializeOpVerifySession) Handle
 func (m *awsRestjson1_serializeOpVerifySession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3695,7 +3581,6 @@ func (m *awsRestjson1_serializeOpVerifySession) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsVerifySessionInput(v *VerifySessionInput, encoder *httpbinding.Encoder) error {
