--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrock/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrock/serializers.go
@@ -14,7 +14,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -29,10 +28,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteEvaluation
 func (m *awsRestjson1_serializeOpBatchDeleteEvaluationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteEvaluation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteEvaluationJobInput(v *BatchDeleteEvaluationJobInput, encoder *httpbinding.Encoder) error {
@@ -112,10 +108,8 @@ func (m *awsRestjson1_serializeOpCreateCustomModel) Ha
 func (m *awsRestjson1_serializeOpCreateCustomModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -160,7 +154,6 @@ func (m *awsRestjson1_serializeOpCreateCustomModel) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCustomModelInput(v *CreateCustomModelInput, encoder *httpbinding.Encoder) error {
@@ -222,10 +215,8 @@ func (m *awsRestjson1_serializeOpCreateCustomModelDepl
 func (m *awsRestjson1_serializeOpCreateCustomModelDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -270,7 +261,6 @@ func (m *awsRestjson1_serializeOpCreateCustomModelDepl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCustomModelDeploymentInput(v *CreateCustomModelDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -325,10 +315,8 @@ func (m *awsRestjson1_serializeOpCreateEvaluationJob) 
 func (m *awsRestjson1_serializeOpCreateEvaluationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -373,7 +361,6 @@ func (m *awsRestjson1_serializeOpCreateEvaluationJob) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEvaluationJobInput(v *CreateEvaluationJobInput, encoder *httpbinding.Encoder) error {
@@ -459,10 +446,8 @@ func (m *awsRestjson1_serializeOpCreateFoundationModel
 func (m *awsRestjson1_serializeOpCreateFoundationModelAgreement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -507,7 +492,6 @@ func (m *awsRestjson1_serializeOpCreateFoundationModel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFoundationModelAgreementInput(v *CreateFoundationModelAgreementInput, encoder *httpbinding.Encoder) error {
@@ -545,10 +529,8 @@ func (m *awsRestjson1_serializeOpCreateGuardrail) Hand
 func (m *awsRestjson1_serializeOpCreateGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -593,7 +575,6 @@ func (m *awsRestjson1_serializeOpCreateGuardrail) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGuardrailInput(v *CreateGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -700,10 +681,8 @@ func (m *awsRestjson1_serializeOpCreateGuardrailVersio
 func (m *awsRestjson1_serializeOpCreateGuardrailVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -752,7 +731,6 @@ func (m *awsRestjson1_serializeOpCreateGuardrailVersio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGuardrailVersionInput(v *CreateGuardrailVersionInput, encoder *httpbinding.Encoder) error {
@@ -799,10 +777,8 @@ func (m *awsRestjson1_serializeOpCreateInferenceProfil
 func (m *awsRestjson1_serializeOpCreateInferenceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -847,7 +823,6 @@ func (m *awsRestjson1_serializeOpCreateInferenceProfil
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInferenceProfileInput(v *CreateInferenceProfileInput, encoder *httpbinding.Encoder) error {
@@ -904,10 +879,8 @@ func (m *awsRestjson1_serializeOpCreateMarketplaceMode
 func (m *awsRestjson1_serializeOpCreateMarketplaceModelEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -952,7 +925,6 @@ func (m *awsRestjson1_serializeOpCreateMarketplaceMode
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMarketplaceModelEndpointInput(v *CreateMarketplaceModelEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1014,10 +986,8 @@ func (m *awsRestjson1_serializeOpCreateModelCopyJob) H
 func (m *awsRestjson1_serializeOpCreateModelCopyJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1062,7 +1032,6 @@ func (m *awsRestjson1_serializeOpCreateModelCopyJob) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateModelCopyJobInput(v *CreateModelCopyJobInput, encoder *httpbinding.Encoder) error {
@@ -1117,10 +1086,8 @@ func (m *awsRestjson1_serializeOpCreateModelCustomizat
 func (m *awsRestjson1_serializeOpCreateModelCustomizationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1165,7 +1132,6 @@ func (m *awsRestjson1_serializeOpCreateModelCustomizat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateModelCustomizationJobInput(v *CreateModelCustomizationJobInput, encoder *httpbinding.Encoder) error {
@@ -1284,10 +1250,8 @@ func (m *awsRestjson1_serializeOpCreateModelImportJob)
 func (m *awsRestjson1_serializeOpCreateModelImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1332,7 +1296,6 @@ func (m *awsRestjson1_serializeOpCreateModelImportJob)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateModelImportJobInput(v *CreateModelImportJobInput, encoder *httpbinding.Encoder) error {
@@ -1413,10 +1376,8 @@ func (m *awsRestjson1_serializeOpCreateModelInvocation
 func (m *awsRestjson1_serializeOpCreateModelInvocationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1461,7 +1422,6 @@ func (m *awsRestjson1_serializeOpCreateModelInvocation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateModelInvocationJobInput(v *CreateModelInvocationJobInput, encoder *httpbinding.Encoder) error {
@@ -1542,10 +1502,8 @@ func (m *awsRestjson1_serializeOpCreatePromptRouter) H
 func (m *awsRestjson1_serializeOpCreatePromptRouter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1590,7 +1548,6 @@ func (m *awsRestjson1_serializeOpCreatePromptRouter) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePromptRouterInput(v *CreatePromptRouterInput, encoder *httpbinding.Encoder) error {
@@ -1661,10 +1618,8 @@ func (m *awsRestjson1_serializeOpCreateProvisionedMode
 func (m *awsRestjson1_serializeOpCreateProvisionedModelThroughput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1709,7 +1664,6 @@ func (m *awsRestjson1_serializeOpCreateProvisionedMode
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProvisionedModelThroughputInput(v *CreateProvisionedModelThroughputInput, encoder *httpbinding.Encoder) error {
@@ -1769,10 +1723,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomModel) Ha
 func (m *awsRestjson1_serializeOpDeleteCustomModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1810,7 +1762,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomModel) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomModelInput(v *DeleteCustomModelInput, encoder *httpbinding.Encoder) error {
@@ -1840,10 +1791,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomModelDepl
 func (m *awsRestjson1_serializeOpDeleteCustomModelDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1881,7 +1830,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomModelDepl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomModelDeploymentInput(v *DeleteCustomModelDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -1911,10 +1859,8 @@ func (m *awsRestjson1_serializeOpDeleteFoundationModel
 func (m *awsRestjson1_serializeOpDeleteFoundationModelAgreement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1959,7 +1905,6 @@ func (m *awsRestjson1_serializeOpDeleteFoundationModel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFoundationModelAgreementInput(v *DeleteFoundationModelAgreementInput, encoder *httpbinding.Encoder) error {
@@ -1992,10 +1937,8 @@ func (m *awsRestjson1_serializeOpDeleteGuardrail) Hand
 func (m *awsRestjson1_serializeOpDeleteGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2033,7 +1976,6 @@ func (m *awsRestjson1_serializeOpDeleteGuardrail) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGuardrailInput(v *DeleteGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -2067,10 +2009,8 @@ func (m *awsRestjson1_serializeOpDeleteImportedModel) 
 func (m *awsRestjson1_serializeOpDeleteImportedModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2108,7 +2048,6 @@ func (m *awsRestjson1_serializeOpDeleteImportedModel) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImportedModelInput(v *DeleteImportedModelInput, encoder *httpbinding.Encoder) error {
@@ -2138,10 +2077,8 @@ func (m *awsRestjson1_serializeOpDeleteInferenceProfil
 func (m *awsRestjson1_serializeOpDeleteInferenceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2179,7 +2116,6 @@ func (m *awsRestjson1_serializeOpDeleteInferenceProfil
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInferenceProfileInput(v *DeleteInferenceProfileInput, encoder *httpbinding.Encoder) error {
@@ -2209,10 +2145,8 @@ func (m *awsRestjson1_serializeOpDeleteMarketplaceMode
 func (m *awsRestjson1_serializeOpDeleteMarketplaceModelEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2250,7 +2184,6 @@ func (m *awsRestjson1_serializeOpDeleteMarketplaceMode
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMarketplaceModelEndpointInput(v *DeleteMarketplaceModelEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2280,10 +2213,8 @@ func (m *awsRestjson1_serializeOpDeleteModelInvocation
 func (m *awsRestjson1_serializeOpDeleteModelInvocationLoggingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2317,7 +2248,6 @@ func (m *awsRestjson1_serializeOpDeleteModelInvocation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteModelInvocationLoggingConfigurationInput(v *DeleteModelInvocationLoggingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2338,10 +2268,8 @@ func (m *awsRestjson1_serializeOpDeletePromptRouter) H
 func (m *awsRestjson1_serializeOpDeletePromptRouter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2379,7 +2307,6 @@ func (m *awsRestjson1_serializeOpDeletePromptRouter) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePromptRouterInput(v *DeletePromptRouterInput, encoder *httpbinding.Encoder) error {
@@ -2409,10 +2336,8 @@ func (m *awsRestjson1_serializeOpDeleteProvisionedMode
 func (m *awsRestjson1_serializeOpDeleteProvisionedModelThroughput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2450,7 +2375,6 @@ func (m *awsRestjson1_serializeOpDeleteProvisionedMode
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProvisionedModelThroughputInput(v *DeleteProvisionedModelThroughputInput, encoder *httpbinding.Encoder) error {
@@ -2480,10 +2404,8 @@ func (m *awsRestjson1_serializeOpDeregisterMarketplace
 func (m *awsRestjson1_serializeOpDeregisterMarketplaceModelEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2521,7 +2443,6 @@ func (m *awsRestjson1_serializeOpDeregisterMarketplace
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeregisterMarketplaceModelEndpointInput(v *DeregisterMarketplaceModelEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2551,10 +2472,8 @@ func (m *awsRestjson1_serializeOpGetCustomModel) Handl
 func (m *awsRestjson1_serializeOpGetCustomModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2592,7 +2511,6 @@ func (m *awsRestjson1_serializeOpGetCustomModel) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCustomModelInput(v *GetCustomModelInput, encoder *httpbinding.Encoder) error {
@@ -2622,10 +2540,8 @@ func (m *awsRestjson1_serializeOpGetCustomModelDeploym
 func (m *awsRestjson1_serializeOpGetCustomModelDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2663,7 +2579,6 @@ func (m *awsRestjson1_serializeOpGetCustomModelDeploym
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCustomModelDeploymentInput(v *GetCustomModelDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -2693,10 +2608,8 @@ func (m *awsRestjson1_serializeOpGetEvaluationJob) Han
 func (m *awsRestjson1_serializeOpGetEvaluationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2734,7 +2647,6 @@ func (m *awsRestjson1_serializeOpGetEvaluationJob) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEvaluationJobInput(v *GetEvaluationJobInput, encoder *httpbinding.Encoder) error {
@@ -2764,10 +2676,8 @@ func (m *awsRestjson1_serializeOpGetFoundationModel) H
 func (m *awsRestjson1_serializeOpGetFoundationModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2805,7 +2715,6 @@ func (m *awsRestjson1_serializeOpGetFoundationModel) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFoundationModelInput(v *GetFoundationModelInput, encoder *httpbinding.Encoder) error {
@@ -2835,10 +2744,8 @@ func (m *awsRestjson1_serializeOpGetFoundationModelAva
 func (m *awsRestjson1_serializeOpGetFoundationModelAvailability) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2876,7 +2783,6 @@ func (m *awsRestjson1_serializeOpGetFoundationModelAva
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFoundationModelAvailabilityInput(v *GetFoundationModelAvailabilityInput, encoder *httpbinding.Encoder) error {
@@ -2906,10 +2812,8 @@ func (m *awsRestjson1_serializeOpGetGuardrail) HandleS
 func (m *awsRestjson1_serializeOpGetGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2947,7 +2851,6 @@ func (m *awsRestjson1_serializeOpGetGuardrail) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGuardrailInput(v *GetGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -2981,10 +2884,8 @@ func (m *awsRestjson1_serializeOpGetImportedModel) Han
 func (m *awsRestjson1_serializeOpGetImportedModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3022,7 +2923,6 @@ func (m *awsRestjson1_serializeOpGetImportedModel) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImportedModelInput(v *GetImportedModelInput, encoder *httpbinding.Encoder) error {
@@ -3052,10 +2952,8 @@ func (m *awsRestjson1_serializeOpGetInferenceProfile) 
 func (m *awsRestjson1_serializeOpGetInferenceProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3093,7 +2991,6 @@ func (m *awsRestjson1_serializeOpGetInferenceProfile) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInferenceProfileInput(v *GetInferenceProfileInput, encoder *httpbinding.Encoder) error {
@@ -3123,10 +3020,8 @@ func (m *awsRestjson1_serializeOpGetMarketplaceModelEn
 func (m *awsRestjson1_serializeOpGetMarketplaceModelEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3164,7 +3059,6 @@ func (m *awsRestjson1_serializeOpGetMarketplaceModelEn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMarketplaceModelEndpointInput(v *GetMarketplaceModelEndpointInput, encoder *httpbinding.Encoder) error {
@@ -3194,10 +3088,8 @@ func (m *awsRestjson1_serializeOpGetModelCopyJob) Hand
 func (m *awsRestjson1_serializeOpGetModelCopyJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3235,7 +3127,6 @@ func (m *awsRestjson1_serializeOpGetModelCopyJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelCopyJobInput(v *GetModelCopyJobInput, encoder *httpbinding.Encoder) error {
@@ -3265,10 +3156,8 @@ func (m *awsRestjson1_serializeOpGetModelCustomization
 func (m *awsRestjson1_serializeOpGetModelCustomizationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3306,7 +3195,6 @@ func (m *awsRestjson1_serializeOpGetModelCustomization
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelCustomizationJobInput(v *GetModelCustomizationJobInput, encoder *httpbinding.Encoder) error {
@@ -3336,10 +3224,8 @@ func (m *awsRestjson1_serializeOpGetModelImportJob) Ha
 func (m *awsRestjson1_serializeOpGetModelImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3377,7 +3263,6 @@ func (m *awsRestjson1_serializeOpGetModelImportJob) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelImportJobInput(v *GetModelImportJobInput, encoder *httpbinding.Encoder) error {
@@ -3407,10 +3292,8 @@ func (m *awsRestjson1_serializeOpGetModelInvocationJob
 func (m *awsRestjson1_serializeOpGetModelInvocationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3448,7 +3331,6 @@ func (m *awsRestjson1_serializeOpGetModelInvocationJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelInvocationJobInput(v *GetModelInvocationJobInput, encoder *httpbinding.Encoder) error {
@@ -3478,10 +3360,8 @@ func (m *awsRestjson1_serializeOpGetModelInvocationLog
 func (m *awsRestjson1_serializeOpGetModelInvocationLoggingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3515,7 +3395,6 @@ func (m *awsRestjson1_serializeOpGetModelInvocationLog
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelInvocationLoggingConfigurationInput(v *GetModelInvocationLoggingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3536,10 +3415,8 @@ func (m *awsRestjson1_serializeOpGetPromptRouter) Hand
 func (m *awsRestjson1_serializeOpGetPromptRouter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3577,7 +3454,6 @@ func (m *awsRestjson1_serializeOpGetPromptRouter) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPromptRouterInput(v *GetPromptRouterInput, encoder *httpbinding.Encoder) error {
@@ -3607,10 +3483,8 @@ func (m *awsRestjson1_serializeOpGetProvisionedModelTh
 func (m *awsRestjson1_serializeOpGetProvisionedModelThroughput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3648,7 +3522,6 @@ func (m *awsRestjson1_serializeOpGetProvisionedModelTh
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProvisionedModelThroughputInput(v *GetProvisionedModelThroughputInput, encoder *httpbinding.Encoder) error {
@@ -3678,10 +3551,8 @@ func (m *awsRestjson1_serializeOpGetUseCaseForModelAcc
 func (m *awsRestjson1_serializeOpGetUseCaseForModelAccess) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3715,7 +3586,6 @@ func (m *awsRestjson1_serializeOpGetUseCaseForModelAcc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUseCaseForModelAccessInput(v *GetUseCaseForModelAccessInput, encoder *httpbinding.Encoder) error {
@@ -3736,10 +3606,8 @@ func (m *awsRestjson1_serializeOpListCustomModelDeploy
 func (m *awsRestjson1_serializeOpListCustomModelDeployments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3777,7 +3645,6 @@ func (m *awsRestjson1_serializeOpListCustomModelDeploy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomModelDeploymentsInput(v *ListCustomModelDeploymentsInput, encoder *httpbinding.Encoder) error {
@@ -3834,10 +3701,8 @@ func (m *awsRestjson1_serializeOpListCustomModels) Han
 func (m *awsRestjson1_serializeOpListCustomModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3875,7 +3740,6 @@ func (m *awsRestjson1_serializeOpListCustomModels) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomModelsInput(v *ListCustomModelsInput, encoder *httpbinding.Encoder) error {
@@ -3940,10 +3804,8 @@ func (m *awsRestjson1_serializeOpListEvaluationJobs) H
 func (m *awsRestjson1_serializeOpListEvaluationJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3981,7 +3843,6 @@ func (m *awsRestjson1_serializeOpListEvaluationJobs) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEvaluationJobsInput(v *ListEvaluationJobsInput, encoder *httpbinding.Encoder) error {
@@ -4038,10 +3899,8 @@ func (m *awsRestjson1_serializeOpListFoundationModelAg
 func (m *awsRestjson1_serializeOpListFoundationModelAgreementOffers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4079,7 +3938,6 @@ func (m *awsRestjson1_serializeOpListFoundationModelAg
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFoundationModelAgreementOffersInput(v *ListFoundationModelAgreementOffersInput, encoder *httpbinding.Encoder) error {
@@ -4113,10 +3971,8 @@ func (m *awsRestjson1_serializeOpListFoundationModels)
 func (m *awsRestjson1_serializeOpListFoundationModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4154,7 +4010,6 @@ func (m *awsRestjson1_serializeOpListFoundationModels)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFoundationModelsInput(v *ListFoundationModelsInput, encoder *httpbinding.Encoder) error {
@@ -4191,10 +4046,8 @@ func (m *awsRestjson1_serializeOpListGuardrails) Handl
 func (m *awsRestjson1_serializeOpListGuardrails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4232,7 +4085,6 @@ func (m *awsRestjson1_serializeOpListGuardrails) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGuardrailsInput(v *ListGuardrailsInput, encoder *httpbinding.Encoder) error {
@@ -4265,10 +4117,8 @@ func (m *awsRestjson1_serializeOpListImportedModels) H
 func (m *awsRestjson1_serializeOpListImportedModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4306,7 +4156,6 @@ func (m *awsRestjson1_serializeOpListImportedModels) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImportedModelsInput(v *ListImportedModelsInput, encoder *httpbinding.Encoder) error {
@@ -4355,10 +4204,8 @@ func (m *awsRestjson1_serializeOpListInferenceProfiles
 func (m *awsRestjson1_serializeOpListInferenceProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4396,7 +4243,6 @@ func (m *awsRestjson1_serializeOpListInferenceProfiles
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInferenceProfilesInput(v *ListInferenceProfilesInput, encoder *httpbinding.Encoder) error {
@@ -4429,10 +4275,8 @@ func (m *awsRestjson1_serializeOpListMarketplaceModelE
 func (m *awsRestjson1_serializeOpListMarketplaceModelEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4470,7 +4314,6 @@ func (m *awsRestjson1_serializeOpListMarketplaceModelE
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMarketplaceModelEndpointsInput(v *ListMarketplaceModelEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -4503,10 +4346,8 @@ func (m *awsRestjson1_serializeOpListModelCopyJobs) Ha
 func (m *awsRestjson1_serializeOpListModelCopyJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4544,7 +4385,6 @@ func (m *awsRestjson1_serializeOpListModelCopyJobs) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListModelCopyJobsInput(v *ListModelCopyJobsInput, encoder *httpbinding.Encoder) error {
@@ -4605,10 +4445,8 @@ func (m *awsRestjson1_serializeOpListModelCustomizatio
 func (m *awsRestjson1_serializeOpListModelCustomizationJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4646,7 +4484,6 @@ func (m *awsRestjson1_serializeOpListModelCustomizatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListModelCustomizationJobsInput(v *ListModelCustomizationJobsInput, encoder *httpbinding.Encoder) error {
@@ -4699,10 +4536,8 @@ func (m *awsRestjson1_serializeOpListModelImportJobs) 
 func (m *awsRestjson1_serializeOpListModelImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4740,7 +4575,6 @@ func (m *awsRestjson1_serializeOpListModelImportJobs) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListModelImportJobsInput(v *ListModelImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -4793,10 +4627,8 @@ func (m *awsRestjson1_serializeOpListModelInvocationJo
 func (m *awsRestjson1_serializeOpListModelInvocationJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4834,7 +4666,6 @@ func (m *awsRestjson1_serializeOpListModelInvocationJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListModelInvocationJobsInput(v *ListModelInvocationJobsInput, encoder *httpbinding.Encoder) error {
@@ -4887,10 +4718,8 @@ func (m *awsRestjson1_serializeOpListPromptRouters) Ha
 func (m *awsRestjson1_serializeOpListPromptRouters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4928,7 +4757,6 @@ func (m *awsRestjson1_serializeOpListPromptRouters) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPromptRoutersInput(v *ListPromptRoutersInput, encoder *httpbinding.Encoder) error {
@@ -4961,10 +4789,8 @@ func (m *awsRestjson1_serializeOpListProvisionedModelT
 func (m *awsRestjson1_serializeOpListProvisionedModelThroughputs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5002,7 +4828,6 @@ func (m *awsRestjson1_serializeOpListProvisionedModelT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProvisionedModelThroughputsInput(v *ListProvisionedModelThroughputsInput, encoder *httpbinding.Encoder) error {
@@ -5059,10 +4884,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -5107,7 +4930,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -5140,10 +4962,8 @@ func (m *awsRestjson1_serializeOpPutModelInvocationLog
 func (m *awsRestjson1_serializeOpPutModelInvocationLoggingConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5188,7 +5008,6 @@ func (m *awsRestjson1_serializeOpPutModelInvocationLog
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutModelInvocationLoggingConfigurationInput(v *PutModelInvocationLoggingConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5223,10 +5042,8 @@ func (m *awsRestjson1_serializeOpPutUseCaseForModelAcc
 func (m *awsRestjson1_serializeOpPutUseCaseForModelAccess) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5271,7 +5088,6 @@ func (m *awsRestjson1_serializeOpPutUseCaseForModelAcc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutUseCaseForModelAccessInput(v *PutUseCaseForModelAccessInput, encoder *httpbinding.Encoder) error {
@@ -5304,10 +5120,8 @@ func (m *awsRestjson1_serializeOpRegisterMarketplaceMo
 func (m *awsRestjson1_serializeOpRegisterMarketplaceModelEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5356,7 +5170,6 @@ func (m *awsRestjson1_serializeOpRegisterMarketplaceMo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterMarketplaceModelEndpointInput(v *RegisterMarketplaceModelEndpointInput, encoder *httpbinding.Encoder) error {
@@ -5398,10 +5211,8 @@ func (m *awsRestjson1_serializeOpStopEvaluationJob) Ha
 func (m *awsRestjson1_serializeOpStopEvaluationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5439,7 +5250,6 @@ func (m *awsRestjson1_serializeOpStopEvaluationJob) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopEvaluationJobInput(v *StopEvaluationJobInput, encoder *httpbinding.Encoder) error {
@@ -5469,10 +5279,8 @@ func (m *awsRestjson1_serializeOpStopModelCustomizatio
 func (m *awsRestjson1_serializeOpStopModelCustomizationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5510,7 +5318,6 @@ func (m *awsRestjson1_serializeOpStopModelCustomizatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopModelCustomizationJobInput(v *StopModelCustomizationJobInput, encoder *httpbinding.Encoder) error {
@@ -5540,10 +5347,8 @@ func (m *awsRestjson1_serializeOpStopModelInvocationJo
 func (m *awsRestjson1_serializeOpStopModelInvocationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5581,7 +5386,6 @@ func (m *awsRestjson1_serializeOpStopModelInvocationJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopModelInvocationJobInput(v *StopModelInvocationJobInput, encoder *httpbinding.Encoder) error {
@@ -5611,10 +5415,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5659,7 +5461,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5699,10 +5500,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5747,7 +5546,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5787,10 +5585,8 @@ func (m *awsRestjson1_serializeOpUpdateGuardrail) Hand
 func (m *awsRestjson1_serializeOpUpdateGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5839,7 +5635,6 @@ func (m *awsRestjson1_serializeOpUpdateGuardrail) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGuardrailInput(v *UpdateGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -5943,10 +5738,8 @@ func (m *awsRestjson1_serializeOpUpdateMarketplaceMode
 func (m *awsRestjson1_serializeOpUpdateMarketplaceModelEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5995,7 +5788,6 @@ func (m *awsRestjson1_serializeOpUpdateMarketplaceMode
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMarketplaceModelEndpointInput(v *UpdateMarketplaceModelEndpointInput, encoder *httpbinding.Encoder) error {
@@ -6044,10 +5836,8 @@ func (m *awsRestjson1_serializeOpUpdateProvisionedMode
 func (m *awsRestjson1_serializeOpUpdateProvisionedModelThroughput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6096,7 +5886,6 @@ func (m *awsRestjson1_serializeOpUpdateProvisionedMode
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProvisionedModelThroughputInput(v *UpdateProvisionedModelThroughputInput, encoder *httpbinding.Encoder) error {
