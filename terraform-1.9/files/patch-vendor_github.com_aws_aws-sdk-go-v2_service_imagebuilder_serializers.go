--- vendor/github.com/aws/aws-sdk-go-v2/service/imagebuilder/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/imagebuilder/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCancelImageCreation) 
 func (m *awsRestjson1_serializeOpCancelImageCreation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpCancelImageCreation) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelImageCreationInput(v *CancelImageCreationInput, encoder *httpbinding.Encoder) error {
@@ -112,10 +108,8 @@ func (m *awsRestjson1_serializeOpCancelLifecycleExecut
 func (m *awsRestjson1_serializeOpCancelLifecycleExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -160,7 +154,6 @@ func (m *awsRestjson1_serializeOpCancelLifecycleExecut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelLifecycleExecutionInput(v *CancelLifecycleExecutionInput, encoder *httpbinding.Encoder) error {
@@ -198,10 +191,8 @@ func (m *awsRestjson1_serializeOpCreateComponent) Hand
 func (m *awsRestjson1_serializeOpCreateComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -246,7 +237,6 @@ func (m *awsRestjson1_serializeOpCreateComponent) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateComponentInput(v *CreateComponentInput, encoder *httpbinding.Encoder) error {
@@ -333,10 +323,8 @@ func (m *awsRestjson1_serializeOpCreateContainerRecipe
 func (m *awsRestjson1_serializeOpCreateContainerRecipe) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -381,7 +369,6 @@ func (m *awsRestjson1_serializeOpCreateContainerRecipe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContainerRecipeInput(v *CreateContainerRecipeInput, encoder *httpbinding.Encoder) error {
@@ -497,10 +484,8 @@ func (m *awsRestjson1_serializeOpCreateDistributionCon
 func (m *awsRestjson1_serializeOpCreateDistributionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -545,7 +530,6 @@ func (m *awsRestjson1_serializeOpCreateDistributionCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDistributionConfigurationInput(v *CreateDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -602,10 +586,8 @@ func (m *awsRestjson1_serializeOpCreateImage) HandleSe
 func (m *awsRestjson1_serializeOpCreateImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -650,7 +632,6 @@ func (m *awsRestjson1_serializeOpCreateImage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImageInput(v *CreateImageInput, encoder *httpbinding.Encoder) error {
@@ -741,10 +722,8 @@ func (m *awsRestjson1_serializeOpCreateImagePipeline) 
 func (m *awsRestjson1_serializeOpCreateImagePipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -789,7 +768,6 @@ func (m *awsRestjson1_serializeOpCreateImagePipeline) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImagePipelineInput(v *CreateImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -902,10 +880,8 @@ func (m *awsRestjson1_serializeOpCreateImageRecipe) Ha
 func (m *awsRestjson1_serializeOpCreateImageRecipe) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -950,7 +926,6 @@ func (m *awsRestjson1_serializeOpCreateImageRecipe) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImageRecipeInput(v *CreateImageRecipeInput, encoder *httpbinding.Encoder) error {
@@ -1036,10 +1011,8 @@ func (m *awsRestjson1_serializeOpCreateInfrastructureC
 func (m *awsRestjson1_serializeOpCreateInfrastructureConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1084,7 +1057,6 @@ func (m *awsRestjson1_serializeOpCreateInfrastructureC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInfrastructureConfigurationInput(v *CreateInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1201,10 +1173,8 @@ func (m *awsRestjson1_serializeOpCreateLifecyclePolicy
 func (m *awsRestjson1_serializeOpCreateLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1249,7 +1219,6 @@ func (m *awsRestjson1_serializeOpCreateLifecyclePolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLifecyclePolicyInput(v *CreateLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
@@ -1328,10 +1297,8 @@ func (m *awsRestjson1_serializeOpCreateWorkflow) Handl
 func (m *awsRestjson1_serializeOpCreateWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1376,7 +1343,6 @@ func (m *awsRestjson1_serializeOpCreateWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWorkflowInput(v *CreateWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -1456,10 +1422,8 @@ func (m *awsRestjson1_serializeOpDeleteComponent) Hand
 func (m *awsRestjson1_serializeOpDeleteComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1497,7 +1461,6 @@ func (m *awsRestjson1_serializeOpDeleteComponent) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteComponentInput(v *DeleteComponentInput, encoder *httpbinding.Encoder) error {
@@ -1522,10 +1485,8 @@ func (m *awsRestjson1_serializeOpDeleteContainerRecipe
 func (m *awsRestjson1_serializeOpDeleteContainerRecipe) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1563,7 +1524,6 @@ func (m *awsRestjson1_serializeOpDeleteContainerRecipe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContainerRecipeInput(v *DeleteContainerRecipeInput, encoder *httpbinding.Encoder) error {
@@ -1588,10 +1548,8 @@ func (m *awsRestjson1_serializeOpDeleteDistributionCon
 func (m *awsRestjson1_serializeOpDeleteDistributionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1629,7 +1587,6 @@ func (m *awsRestjson1_serializeOpDeleteDistributionCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDistributionConfigurationInput(v *DeleteDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1654,10 +1611,8 @@ func (m *awsRestjson1_serializeOpDeleteImage) HandleSe
 func (m *awsRestjson1_serializeOpDeleteImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1695,7 +1650,6 @@ func (m *awsRestjson1_serializeOpDeleteImage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImageInput(v *DeleteImageInput, encoder *httpbinding.Encoder) error {
@@ -1720,10 +1674,8 @@ func (m *awsRestjson1_serializeOpDeleteImagePipeline) 
 func (m *awsRestjson1_serializeOpDeleteImagePipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1761,7 +1713,6 @@ func (m *awsRestjson1_serializeOpDeleteImagePipeline) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImagePipelineInput(v *DeleteImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -1786,10 +1737,8 @@ func (m *awsRestjson1_serializeOpDeleteImageRecipe) Ha
 func (m *awsRestjson1_serializeOpDeleteImageRecipe) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1827,7 +1776,6 @@ func (m *awsRestjson1_serializeOpDeleteImageRecipe) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImageRecipeInput(v *DeleteImageRecipeInput, encoder *httpbinding.Encoder) error {
@@ -1852,10 +1800,8 @@ func (m *awsRestjson1_serializeOpDeleteInfrastructureC
 func (m *awsRestjson1_serializeOpDeleteInfrastructureConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1893,7 +1839,6 @@ func (m *awsRestjson1_serializeOpDeleteInfrastructureC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInfrastructureConfigurationInput(v *DeleteInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1918,10 +1863,8 @@ func (m *awsRestjson1_serializeOpDeleteLifecyclePolicy
 func (m *awsRestjson1_serializeOpDeleteLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1959,7 +1902,6 @@ func (m *awsRestjson1_serializeOpDeleteLifecyclePolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLifecyclePolicyInput(v *DeleteLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
@@ -1984,10 +1926,8 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
 func (m *awsRestjson1_serializeOpDeleteWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2025,7 +1965,6 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWorkflowInput(v *DeleteWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -2050,10 +1989,8 @@ func (m *awsRestjson1_serializeOpGetComponent) HandleS
 func (m *awsRestjson1_serializeOpGetComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2091,7 +2028,6 @@ func (m *awsRestjson1_serializeOpGetComponent) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetComponentInput(v *GetComponentInput, encoder *httpbinding.Encoder) error {
@@ -2116,10 +2052,8 @@ func (m *awsRestjson1_serializeOpGetComponentPolicy) H
 func (m *awsRestjson1_serializeOpGetComponentPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2157,7 +2091,6 @@ func (m *awsRestjson1_serializeOpGetComponentPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetComponentPolicyInput(v *GetComponentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2182,10 +2115,8 @@ func (m *awsRestjson1_serializeOpGetContainerRecipe) H
 func (m *awsRestjson1_serializeOpGetContainerRecipe) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2223,7 +2154,6 @@ func (m *awsRestjson1_serializeOpGetContainerRecipe) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContainerRecipeInput(v *GetContainerRecipeInput, encoder *httpbinding.Encoder) error {
@@ -2248,10 +2178,8 @@ func (m *awsRestjson1_serializeOpGetContainerRecipePol
 func (m *awsRestjson1_serializeOpGetContainerRecipePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2289,7 +2217,6 @@ func (m *awsRestjson1_serializeOpGetContainerRecipePol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContainerRecipePolicyInput(v *GetContainerRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2314,10 +2241,8 @@ func (m *awsRestjson1_serializeOpGetDistributionConfig
 func (m *awsRestjson1_serializeOpGetDistributionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2355,7 +2280,6 @@ func (m *awsRestjson1_serializeOpGetDistributionConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDistributionConfigurationInput(v *GetDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2380,10 +2304,8 @@ func (m *awsRestjson1_serializeOpGetImage) HandleSeria
 func (m *awsRestjson1_serializeOpGetImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2421,7 +2343,6 @@ func (m *awsRestjson1_serializeOpGetImage) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImageInput(v *GetImageInput, encoder *httpbinding.Encoder) error {
@@ -2446,10 +2367,8 @@ func (m *awsRestjson1_serializeOpGetImagePipeline) Han
 func (m *awsRestjson1_serializeOpGetImagePipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2487,7 +2406,6 @@ func (m *awsRestjson1_serializeOpGetImagePipeline) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImagePipelineInput(v *GetImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -2512,10 +2430,8 @@ func (m *awsRestjson1_serializeOpGetImagePolicy) Handl
 func (m *awsRestjson1_serializeOpGetImagePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2553,7 +2469,6 @@ func (m *awsRestjson1_serializeOpGetImagePolicy) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImagePolicyInput(v *GetImagePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2578,10 +2493,8 @@ func (m *awsRestjson1_serializeOpGetImageRecipe) Handl
 func (m *awsRestjson1_serializeOpGetImageRecipe) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2619,7 +2532,6 @@ func (m *awsRestjson1_serializeOpGetImageRecipe) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImageRecipeInput(v *GetImageRecipeInput, encoder *httpbinding.Encoder) error {
@@ -2644,10 +2556,8 @@ func (m *awsRestjson1_serializeOpGetImageRecipePolicy)
 func (m *awsRestjson1_serializeOpGetImageRecipePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2685,7 +2595,6 @@ func (m *awsRestjson1_serializeOpGetImageRecipePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImageRecipePolicyInput(v *GetImageRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2710,10 +2619,8 @@ func (m *awsRestjson1_serializeOpGetInfrastructureConf
 func (m *awsRestjson1_serializeOpGetInfrastructureConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2751,7 +2658,6 @@ func (m *awsRestjson1_serializeOpGetInfrastructureConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInfrastructureConfigurationInput(v *GetInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2776,10 +2682,8 @@ func (m *awsRestjson1_serializeOpGetLifecycleExecution
 func (m *awsRestjson1_serializeOpGetLifecycleExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2817,7 +2721,6 @@ func (m *awsRestjson1_serializeOpGetLifecycleExecution
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLifecycleExecutionInput(v *GetLifecycleExecutionInput, encoder *httpbinding.Encoder) error {
@@ -2842,10 +2745,8 @@ func (m *awsRestjson1_serializeOpGetLifecyclePolicy) H
 func (m *awsRestjson1_serializeOpGetLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2883,7 +2784,6 @@ func (m *awsRestjson1_serializeOpGetLifecyclePolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLifecyclePolicyInput(v *GetLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2908,10 +2808,8 @@ func (m *awsRestjson1_serializeOpGetMarketplaceResourc
 func (m *awsRestjson1_serializeOpGetMarketplaceResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2956,7 +2854,6 @@ func (m *awsRestjson1_serializeOpGetMarketplaceResourc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMarketplaceResourceInput(v *GetMarketplaceResourceInput, encoder *httpbinding.Encoder) error {
@@ -2999,10 +2896,8 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
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
@@ -3040,7 +2935,6 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowInput(v *GetWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -3065,10 +2959,8 @@ func (m *awsRestjson1_serializeOpGetWorkflowExecution)
 func (m *awsRestjson1_serializeOpGetWorkflowExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3106,7 +2998,6 @@ func (m *awsRestjson1_serializeOpGetWorkflowExecution)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowExecutionInput(v *GetWorkflowExecutionInput, encoder *httpbinding.Encoder) error {
@@ -3131,10 +3022,8 @@ func (m *awsRestjson1_serializeOpGetWorkflowStepExecut
 func (m *awsRestjson1_serializeOpGetWorkflowStepExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3172,7 +3061,6 @@ func (m *awsRestjson1_serializeOpGetWorkflowStepExecut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowStepExecutionInput(v *GetWorkflowStepExecutionInput, encoder *httpbinding.Encoder) error {
@@ -3197,10 +3085,8 @@ func (m *awsRestjson1_serializeOpImportComponent) Hand
 func (m *awsRestjson1_serializeOpImportComponent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3245,7 +3131,6 @@ func (m *awsRestjson1_serializeOpImportComponent) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportComponentInput(v *ImportComponentInput, encoder *httpbinding.Encoder) error {
@@ -3335,10 +3220,8 @@ func (m *awsRestjson1_serializeOpImportDiskImage) Hand
 func (m *awsRestjson1_serializeOpImportDiskImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3383,7 +3266,6 @@ func (m *awsRestjson1_serializeOpImportDiskImage) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportDiskImageInput(v *ImportDiskImageInput, encoder *httpbinding.Encoder) error {
@@ -3463,10 +3345,8 @@ func (m *awsRestjson1_serializeOpImportVmImage) Handle
 func (m *awsRestjson1_serializeOpImportVmImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3511,7 +3391,6 @@ func (m *awsRestjson1_serializeOpImportVmImage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportVmImageInput(v *ImportVmImageInput, encoder *httpbinding.Encoder) error {
@@ -3581,10 +3460,8 @@ func (m *awsRestjson1_serializeOpListComponentBuildVer
 func (m *awsRestjson1_serializeOpListComponentBuildVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3629,7 +3506,6 @@ func (m *awsRestjson1_serializeOpListComponentBuildVer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListComponentBuildVersionsInput(v *ListComponentBuildVersionsInput, encoder *httpbinding.Encoder) error {
@@ -3672,10 +3548,8 @@ func (m *awsRestjson1_serializeOpListComponents) Handl
 func (m *awsRestjson1_serializeOpListComponents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3720,7 +3594,6 @@ func (m *awsRestjson1_serializeOpListComponents) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListComponentsInput(v *ListComponentsInput, encoder *httpbinding.Encoder) error {
@@ -3775,10 +3648,8 @@ func (m *awsRestjson1_serializeOpListContainerRecipes)
 func (m *awsRestjson1_serializeOpListContainerRecipes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3823,7 +3694,6 @@ func (m *awsRestjson1_serializeOpListContainerRecipes)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContainerRecipesInput(v *ListContainerRecipesInput, encoder *httpbinding.Encoder) error {
@@ -3873,10 +3743,8 @@ func (m *awsRestjson1_serializeOpListDistributionConfi
 func (m *awsRestjson1_serializeOpListDistributionConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3921,7 +3789,6 @@ func (m *awsRestjson1_serializeOpListDistributionConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDistributionConfigurationsInput(v *ListDistributionConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -3966,10 +3833,8 @@ func (m *awsRestjson1_serializeOpListImageBuildVersion
 func (m *awsRestjson1_serializeOpListImageBuildVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4014,7 +3879,6 @@ func (m *awsRestjson1_serializeOpListImageBuildVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageBuildVersionsInput(v *ListImageBuildVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4064,10 +3928,8 @@ func (m *awsRestjson1_serializeOpListImagePackages) Ha
 func (m *awsRestjson1_serializeOpListImagePackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4112,7 +3974,6 @@ func (m *awsRestjson1_serializeOpListImagePackages) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagePackagesInput(v *ListImagePackagesInput, encoder *httpbinding.Encoder) error {
@@ -4155,10 +4016,8 @@ func (m *awsRestjson1_serializeOpListImagePipelineImag
 func (m *awsRestjson1_serializeOpListImagePipelineImages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4203,7 +4062,6 @@ func (m *awsRestjson1_serializeOpListImagePipelineImag
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagePipelineImagesInput(v *ListImagePipelineImagesInput, encoder *httpbinding.Encoder) error {
@@ -4253,10 +4111,8 @@ func (m *awsRestjson1_serializeOpListImagePipelines) H
 func (m *awsRestjson1_serializeOpListImagePipelines) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4301,7 +4157,6 @@ func (m *awsRestjson1_serializeOpListImagePipelines) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagePipelinesInput(v *ListImagePipelinesInput, encoder *httpbinding.Encoder) error {
@@ -4346,10 +4201,8 @@ func (m *awsRestjson1_serializeOpListImageRecipes) Han
 func (m *awsRestjson1_serializeOpListImageRecipes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4394,7 +4247,6 @@ func (m *awsRestjson1_serializeOpListImageRecipes) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageRecipesInput(v *ListImageRecipesInput, encoder *httpbinding.Encoder) error {
@@ -4444,10 +4296,8 @@ func (m *awsRestjson1_serializeOpListImages) HandleSer
 func (m *awsRestjson1_serializeOpListImages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4492,7 +4342,6 @@ func (m *awsRestjson1_serializeOpListImages) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagesInput(v *ListImagesInput, encoder *httpbinding.Encoder) error {
@@ -4552,10 +4401,8 @@ func (m *awsRestjson1_serializeOpListImageScanFindingA
 func (m *awsRestjson1_serializeOpListImageScanFindingAggregations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4600,7 +4447,6 @@ func (m *awsRestjson1_serializeOpListImageScanFindingA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageScanFindingAggregationsInput(v *ListImageScanFindingAggregationsInput, encoder *httpbinding.Encoder) error {
@@ -4640,10 +4486,8 @@ func (m *awsRestjson1_serializeOpListImageScanFindings
 func (m *awsRestjson1_serializeOpListImageScanFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4688,7 +4532,6 @@ func (m *awsRestjson1_serializeOpListImageScanFindings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageScanFindingsInput(v *ListImageScanFindingsInput, encoder *httpbinding.Encoder) error {
@@ -4733,10 +4576,8 @@ func (m *awsRestjson1_serializeOpListInfrastructureCon
 func (m *awsRestjson1_serializeOpListInfrastructureConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4781,7 +4622,6 @@ func (m *awsRestjson1_serializeOpListInfrastructureCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInfrastructureConfigurationsInput(v *ListInfrastructureConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -4826,10 +4666,8 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
 func (m *awsRestjson1_serializeOpListLifecycleExecutionResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4874,7 +4712,6 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLifecycleExecutionResourcesInput(v *ListLifecycleExecutionResourcesInput, encoder *httpbinding.Encoder) error {
@@ -4922,10 +4759,8 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
 func (m *awsRestjson1_serializeOpListLifecycleExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4970,7 +4805,6 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLifecycleExecutionsInput(v *ListLifecycleExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -5013,10 +4847,8 @@ func (m *awsRestjson1_serializeOpListLifecyclePolicies
 func (m *awsRestjson1_serializeOpListLifecyclePolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5061,7 +4893,6 @@ func (m *awsRestjson1_serializeOpListLifecyclePolicies
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLifecyclePoliciesInput(v *ListLifecyclePoliciesInput, encoder *httpbinding.Encoder) error {
@@ -5106,10 +4937,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -5147,7 +4976,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -5177,10 +5005,8 @@ func (m *awsRestjson1_serializeOpListWaitingWorkflowSt
 func (m *awsRestjson1_serializeOpListWaitingWorkflowSteps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5225,7 +5051,6 @@ func (m *awsRestjson1_serializeOpListWaitingWorkflowSt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWaitingWorkflowStepsInput(v *ListWaitingWorkflowStepsInput, encoder *httpbinding.Encoder) error {
@@ -5263,10 +5088,8 @@ func (m *awsRestjson1_serializeOpListWorkflowBuildVers
 func (m *awsRestjson1_serializeOpListWorkflowBuildVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5311,7 +5134,6 @@ func (m *awsRestjson1_serializeOpListWorkflowBuildVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowBuildVersionsInput(v *ListWorkflowBuildVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5354,10 +5176,8 @@ func (m *awsRestjson1_serializeOpListWorkflowExecution
 func (m *awsRestjson1_serializeOpListWorkflowExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5402,7 +5222,6 @@ func (m *awsRestjson1_serializeOpListWorkflowExecution
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowExecutionsInput(v *ListWorkflowExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -5445,10 +5264,8 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
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
@@ -5493,7 +5310,6 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowsInput(v *ListWorkflowsInput, encoder *httpbinding.Encoder) error {
@@ -5548,10 +5364,8 @@ func (m *awsRestjson1_serializeOpListWorkflowStepExecu
 func (m *awsRestjson1_serializeOpListWorkflowStepExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5596,7 +5410,6 @@ func (m *awsRestjson1_serializeOpListWorkflowStepExecu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowStepExecutionsInput(v *ListWorkflowStepExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -5639,10 +5452,8 @@ func (m *awsRestjson1_serializeOpPutComponentPolicy) H
 func (m *awsRestjson1_serializeOpPutComponentPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5687,7 +5498,6 @@ func (m *awsRestjson1_serializeOpPutComponentPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutComponentPolicyInput(v *PutComponentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -5725,10 +5535,8 @@ func (m *awsRestjson1_serializeOpPutContainerRecipePol
 func (m *awsRestjson1_serializeOpPutContainerRecipePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5773,7 +5581,6 @@ func (m *awsRestjson1_serializeOpPutContainerRecipePol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutContainerRecipePolicyInput(v *PutContainerRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -5811,10 +5618,8 @@ func (m *awsRestjson1_serializeOpPutImagePolicy) Handl
 func (m *awsRestjson1_serializeOpPutImagePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5859,7 +5664,6 @@ func (m *awsRestjson1_serializeOpPutImagePolicy) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutImagePolicyInput(v *PutImagePolicyInput, encoder *httpbinding.Encoder) error {
@@ -5897,10 +5701,8 @@ func (m *awsRestjson1_serializeOpPutImageRecipePolicy)
 func (m *awsRestjson1_serializeOpPutImageRecipePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5945,7 +5747,6 @@ func (m *awsRestjson1_serializeOpPutImageRecipePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutImageRecipePolicyInput(v *PutImageRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -5983,10 +5784,8 @@ func (m *awsRestjson1_serializeOpSendWorkflowStepActio
 func (m *awsRestjson1_serializeOpSendWorkflowStepAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6031,7 +5830,6 @@ func (m *awsRestjson1_serializeOpSendWorkflowStepActio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendWorkflowStepActionInput(v *SendWorkflowStepActionInput, encoder *httpbinding.Encoder) error {
@@ -6084,10 +5882,8 @@ func (m *awsRestjson1_serializeOpStartImagePipelineExe
 func (m *awsRestjson1_serializeOpStartImagePipelineExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6132,7 +5928,6 @@ func (m *awsRestjson1_serializeOpStartImagePipelineExe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartImagePipelineExecutionInput(v *StartImagePipelineExecutionInput, encoder *httpbinding.Encoder) error {
@@ -6170,10 +5965,8 @@ func (m *awsRestjson1_serializeOpStartResourceStateUpd
 func (m *awsRestjson1_serializeOpStartResourceStateUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6218,7 +6011,6 @@ func (m *awsRestjson1_serializeOpStartResourceStateUpd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartResourceStateUpdateInput(v *StartResourceStateUpdateInput, encoder *httpbinding.Encoder) error {
@@ -6287,10 +6079,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -6339,7 +6129,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6383,10 +6172,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -6424,7 +6211,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6460,10 +6246,8 @@ func (m *awsRestjson1_serializeOpUpdateDistributionCon
 func (m *awsRestjson1_serializeOpUpdateDistributionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6508,7 +6292,6 @@ func (m *awsRestjson1_serializeOpUpdateDistributionCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDistributionConfigurationInput(v *UpdateDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6558,10 +6341,8 @@ func (m *awsRestjson1_serializeOpUpdateImagePipeline) 
 func (m *awsRestjson1_serializeOpUpdateImagePipeline) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6606,7 +6387,6 @@ func (m *awsRestjson1_serializeOpUpdateImagePipeline) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateImagePipelineInput(v *UpdateImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -6712,10 +6492,8 @@ func (m *awsRestjson1_serializeOpUpdateInfrastructureC
 func (m *awsRestjson1_serializeOpUpdateInfrastructureConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6760,7 +6538,6 @@ func (m *awsRestjson1_serializeOpUpdateInfrastructureC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInfrastructureConfigurationInput(v *UpdateInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6870,10 +6647,8 @@ func (m *awsRestjson1_serializeOpUpdateLifecyclePolicy
 func (m *awsRestjson1_serializeOpUpdateLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6918,7 +6693,6 @@ func (m *awsRestjson1_serializeOpUpdateLifecyclePolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLifecyclePolicyInput(v *UpdateLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
