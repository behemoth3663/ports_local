--- vendor/github.com/aws/aws-sdk-go-v2/service/imagebuilder/serializers.go.orig	2026-11-21 14:46:20 UTC
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
@@ -338,10 +328,8 @@ func (m *awsRestjson1_serializeOpCreateContainerRecipe
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
@@ -386,7 +374,6 @@ func (m *awsRestjson1_serializeOpCreateContainerRecipe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContainerRecipeInput(v *CreateContainerRecipeInput, encoder *httpbinding.Encoder) error {
@@ -507,10 +494,8 @@ func (m *awsRestjson1_serializeOpCreateDistributionCon
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
@@ -555,7 +540,6 @@ func (m *awsRestjson1_serializeOpCreateDistributionCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDistributionConfigurationInput(v *CreateDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -617,10 +601,8 @@ func (m *awsRestjson1_serializeOpCreateImage) HandleSe
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
@@ -665,7 +647,6 @@ func (m *awsRestjson1_serializeOpCreateImage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImageInput(v *CreateImageInput, encoder *httpbinding.Encoder) error {
@@ -763,10 +744,8 @@ func (m *awsRestjson1_serializeOpCreateImagePipeline) 
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
@@ -811,7 +790,6 @@ func (m *awsRestjson1_serializeOpCreateImagePipeline) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImagePipelineInput(v *CreateImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -943,10 +921,8 @@ func (m *awsRestjson1_serializeOpCreateImageRecipe) Ha
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
@@ -991,7 +967,6 @@ func (m *awsRestjson1_serializeOpCreateImageRecipe) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateImageRecipeInput(v *CreateImageRecipeInput, encoder *httpbinding.Encoder) error {
@@ -1096,10 +1071,8 @@ func (m *awsRestjson1_serializeOpCreateInfrastructureC
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
@@ -1144,7 +1117,6 @@ func (m *awsRestjson1_serializeOpCreateInfrastructureC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateInfrastructureConfigurationInput(v *CreateInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1266,10 +1238,8 @@ func (m *awsRestjson1_serializeOpCreateLifecyclePolicy
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
@@ -1314,7 +1284,6 @@ func (m *awsRestjson1_serializeOpCreateLifecyclePolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLifecyclePolicyInput(v *CreateLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
@@ -1398,10 +1367,8 @@ func (m *awsRestjson1_serializeOpCreateWorkflow) Handl
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
@@ -1446,7 +1413,6 @@ func (m *awsRestjson1_serializeOpCreateWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWorkflowInput(v *CreateWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -1531,10 +1497,8 @@ func (m *awsRestjson1_serializeOpDeleteComponent) Hand
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
@@ -1572,7 +1536,6 @@ func (m *awsRestjson1_serializeOpDeleteComponent) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteComponentInput(v *DeleteComponentInput, encoder *httpbinding.Encoder) error {
@@ -1597,10 +1560,8 @@ func (m *awsRestjson1_serializeOpDeleteContainerRecipe
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
@@ -1638,7 +1599,6 @@ func (m *awsRestjson1_serializeOpDeleteContainerRecipe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContainerRecipeInput(v *DeleteContainerRecipeInput, encoder *httpbinding.Encoder) error {
@@ -1663,10 +1623,8 @@ func (m *awsRestjson1_serializeOpDeleteDistributionCon
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
@@ -1704,7 +1662,6 @@ func (m *awsRestjson1_serializeOpDeleteDistributionCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDistributionConfigurationInput(v *DeleteDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1729,10 +1686,8 @@ func (m *awsRestjson1_serializeOpDeleteImage) HandleSe
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
@@ -1770,7 +1725,6 @@ func (m *awsRestjson1_serializeOpDeleteImage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImageInput(v *DeleteImageInput, encoder *httpbinding.Encoder) error {
@@ -1795,10 +1749,8 @@ func (m *awsRestjson1_serializeOpDeleteImagePipeline) 
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
@@ -1836,7 +1788,6 @@ func (m *awsRestjson1_serializeOpDeleteImagePipeline) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImagePipelineInput(v *DeleteImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -1861,10 +1812,8 @@ func (m *awsRestjson1_serializeOpDeleteImageRecipe) Ha
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
@@ -1902,7 +1851,6 @@ func (m *awsRestjson1_serializeOpDeleteImageRecipe) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImageRecipeInput(v *DeleteImageRecipeInput, encoder *httpbinding.Encoder) error {
@@ -1927,10 +1875,8 @@ func (m *awsRestjson1_serializeOpDeleteInfrastructureC
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
@@ -1968,7 +1914,6 @@ func (m *awsRestjson1_serializeOpDeleteInfrastructureC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInfrastructureConfigurationInput(v *DeleteInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1993,10 +1938,8 @@ func (m *awsRestjson1_serializeOpDeleteLifecyclePolicy
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
@@ -2034,7 +1977,6 @@ func (m *awsRestjson1_serializeOpDeleteLifecyclePolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLifecyclePolicyInput(v *DeleteLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2059,10 +2001,8 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
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
@@ -2100,7 +2040,6 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWorkflowInput(v *DeleteWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -2125,10 +2064,8 @@ func (m *awsRestjson1_serializeOpDistributeImage) Hand
 func (m *awsRestjson1_serializeOpDistributeImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2173,7 +2110,6 @@ func (m *awsRestjson1_serializeOpDistributeImage) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDistributeImageInput(v *DistributeImageInput, encoder *httpbinding.Encoder) error {
@@ -2235,10 +2171,8 @@ func (m *awsRestjson1_serializeOpGetComponent) HandleS
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
@@ -2276,7 +2210,6 @@ func (m *awsRestjson1_serializeOpGetComponent) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetComponentInput(v *GetComponentInput, encoder *httpbinding.Encoder) error {
@@ -2301,10 +2234,8 @@ func (m *awsRestjson1_serializeOpGetComponentPolicy) H
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
@@ -2342,7 +2273,6 @@ func (m *awsRestjson1_serializeOpGetComponentPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetComponentPolicyInput(v *GetComponentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2367,10 +2297,8 @@ func (m *awsRestjson1_serializeOpGetContainerRecipe) H
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
@@ -2408,7 +2336,6 @@ func (m *awsRestjson1_serializeOpGetContainerRecipe) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContainerRecipeInput(v *GetContainerRecipeInput, encoder *httpbinding.Encoder) error {
@@ -2433,10 +2360,8 @@ func (m *awsRestjson1_serializeOpGetContainerRecipePol
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
@@ -2474,7 +2399,6 @@ func (m *awsRestjson1_serializeOpGetContainerRecipePol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContainerRecipePolicyInput(v *GetContainerRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2499,10 +2423,8 @@ func (m *awsRestjson1_serializeOpGetDistributionConfig
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
@@ -2540,7 +2462,6 @@ func (m *awsRestjson1_serializeOpGetDistributionConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDistributionConfigurationInput(v *GetDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2565,10 +2486,8 @@ func (m *awsRestjson1_serializeOpGetImage) HandleSeria
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
@@ -2606,7 +2525,6 @@ func (m *awsRestjson1_serializeOpGetImage) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImageInput(v *GetImageInput, encoder *httpbinding.Encoder) error {
@@ -2631,10 +2549,8 @@ func (m *awsRestjson1_serializeOpGetImagePipeline) Han
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
@@ -2672,7 +2588,6 @@ func (m *awsRestjson1_serializeOpGetImagePipeline) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImagePipelineInput(v *GetImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -2697,10 +2612,8 @@ func (m *awsRestjson1_serializeOpGetImagePolicy) Handl
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
@@ -2738,7 +2651,6 @@ func (m *awsRestjson1_serializeOpGetImagePolicy) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImagePolicyInput(v *GetImagePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2763,10 +2675,8 @@ func (m *awsRestjson1_serializeOpGetImageRecipe) Handl
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
@@ -2804,7 +2714,6 @@ func (m *awsRestjson1_serializeOpGetImageRecipe) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImageRecipeInput(v *GetImageRecipeInput, encoder *httpbinding.Encoder) error {
@@ -2829,10 +2738,8 @@ func (m *awsRestjson1_serializeOpGetImageRecipePolicy)
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
@@ -2870,7 +2777,6 @@ func (m *awsRestjson1_serializeOpGetImageRecipePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImageRecipePolicyInput(v *GetImageRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2895,10 +2801,8 @@ func (m *awsRestjson1_serializeOpGetInfrastructureConf
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
@@ -2936,7 +2840,6 @@ func (m *awsRestjson1_serializeOpGetInfrastructureConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInfrastructureConfigurationInput(v *GetInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2961,10 +2864,8 @@ func (m *awsRestjson1_serializeOpGetLifecycleExecution
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
@@ -3002,7 +2903,6 @@ func (m *awsRestjson1_serializeOpGetLifecycleExecution
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLifecycleExecutionInput(v *GetLifecycleExecutionInput, encoder *httpbinding.Encoder) error {
@@ -3027,10 +2927,8 @@ func (m *awsRestjson1_serializeOpGetLifecyclePolicy) H
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
@@ -3068,7 +2966,6 @@ func (m *awsRestjson1_serializeOpGetLifecyclePolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLifecyclePolicyInput(v *GetLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
@@ -3093,10 +2990,8 @@ func (m *awsRestjson1_serializeOpGetMarketplaceResourc
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
@@ -3141,7 +3036,6 @@ func (m *awsRestjson1_serializeOpGetMarketplaceResourc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMarketplaceResourceInput(v *GetMarketplaceResourceInput, encoder *httpbinding.Encoder) error {
@@ -3184,10 +3078,8 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
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
@@ -3225,7 +3117,6 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowInput(v *GetWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -3250,10 +3141,8 @@ func (m *awsRestjson1_serializeOpGetWorkflowExecution)
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
@@ -3291,7 +3180,6 @@ func (m *awsRestjson1_serializeOpGetWorkflowExecution)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowExecutionInput(v *GetWorkflowExecutionInput, encoder *httpbinding.Encoder) error {
@@ -3316,10 +3204,8 @@ func (m *awsRestjson1_serializeOpGetWorkflowStepExecut
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
@@ -3357,7 +3243,6 @@ func (m *awsRestjson1_serializeOpGetWorkflowStepExecut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowStepExecutionInput(v *GetWorkflowStepExecutionInput, encoder *httpbinding.Encoder) error {
@@ -3382,10 +3267,8 @@ func (m *awsRestjson1_serializeOpImportComponent) Hand
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
@@ -3430,7 +3313,6 @@ func (m *awsRestjson1_serializeOpImportComponent) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportComponentInput(v *ImportComponentInput, encoder *httpbinding.Encoder) error {
@@ -3520,10 +3402,8 @@ func (m *awsRestjson1_serializeOpImportDiskImage) Hand
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
@@ -3568,7 +3448,6 @@ func (m *awsRestjson1_serializeOpImportDiskImage) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportDiskImageInput(v *ImportDiskImageInput, encoder *httpbinding.Encoder) error {
@@ -3669,10 +3548,8 @@ func (m *awsRestjson1_serializeOpImportVmImage) Handle
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
@@ -3717,7 +3594,6 @@ func (m *awsRestjson1_serializeOpImportVmImage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportVmImageInput(v *ImportVmImageInput, encoder *httpbinding.Encoder) error {
@@ -3794,10 +3670,8 @@ func (m *awsRestjson1_serializeOpListComponentBuildVer
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
@@ -3842,7 +3716,6 @@ func (m *awsRestjson1_serializeOpListComponentBuildVer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListComponentBuildVersionsInput(v *ListComponentBuildVersionsInput, encoder *httpbinding.Encoder) error {
@@ -3885,10 +3758,8 @@ func (m *awsRestjson1_serializeOpListComponents) Handl
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
@@ -3933,7 +3804,6 @@ func (m *awsRestjson1_serializeOpListComponents) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListComponentsInput(v *ListComponentsInput, encoder *httpbinding.Encoder) error {
@@ -3988,10 +3858,8 @@ func (m *awsRestjson1_serializeOpListContainerRecipes)
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
@@ -4036,7 +3904,6 @@ func (m *awsRestjson1_serializeOpListContainerRecipes)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContainerRecipesInput(v *ListContainerRecipesInput, encoder *httpbinding.Encoder) error {
@@ -4086,10 +3953,8 @@ func (m *awsRestjson1_serializeOpListDistributionConfi
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
@@ -4134,7 +3999,6 @@ func (m *awsRestjson1_serializeOpListDistributionConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDistributionConfigurationsInput(v *ListDistributionConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -4179,10 +4043,8 @@ func (m *awsRestjson1_serializeOpListImageBuildVersion
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
@@ -4227,7 +4089,6 @@ func (m *awsRestjson1_serializeOpListImageBuildVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageBuildVersionsInput(v *ListImageBuildVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4277,10 +4138,8 @@ func (m *awsRestjson1_serializeOpListImagePackages) Ha
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
@@ -4325,7 +4184,6 @@ func (m *awsRestjson1_serializeOpListImagePackages) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagePackagesInput(v *ListImagePackagesInput, encoder *httpbinding.Encoder) error {
@@ -4368,10 +4226,8 @@ func (m *awsRestjson1_serializeOpListImagePipelineImag
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
@@ -4416,7 +4272,6 @@ func (m *awsRestjson1_serializeOpListImagePipelineImag
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagePipelineImagesInput(v *ListImagePipelineImagesInput, encoder *httpbinding.Encoder) error {
@@ -4466,10 +4321,8 @@ func (m *awsRestjson1_serializeOpListImagePipelines) H
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
@@ -4514,7 +4367,6 @@ func (m *awsRestjson1_serializeOpListImagePipelines) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagePipelinesInput(v *ListImagePipelinesInput, encoder *httpbinding.Encoder) error {
@@ -4559,10 +4411,8 @@ func (m *awsRestjson1_serializeOpListImageRecipes) Han
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
@@ -4607,7 +4457,6 @@ func (m *awsRestjson1_serializeOpListImageRecipes) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageRecipesInput(v *ListImageRecipesInput, encoder *httpbinding.Encoder) error {
@@ -4657,10 +4506,8 @@ func (m *awsRestjson1_serializeOpListImages) HandleSer
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
@@ -4705,7 +4552,6 @@ func (m *awsRestjson1_serializeOpListImages) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImagesInput(v *ListImagesInput, encoder *httpbinding.Encoder) error {
@@ -4765,10 +4611,8 @@ func (m *awsRestjson1_serializeOpListImageScanFindingA
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
@@ -4813,7 +4657,6 @@ func (m *awsRestjson1_serializeOpListImageScanFindingA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageScanFindingAggregationsInput(v *ListImageScanFindingAggregationsInput, encoder *httpbinding.Encoder) error {
@@ -4853,10 +4696,8 @@ func (m *awsRestjson1_serializeOpListImageScanFindings
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
@@ -4901,7 +4742,6 @@ func (m *awsRestjson1_serializeOpListImageScanFindings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImageScanFindingsInput(v *ListImageScanFindingsInput, encoder *httpbinding.Encoder) error {
@@ -4946,10 +4786,8 @@ func (m *awsRestjson1_serializeOpListInfrastructureCon
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
@@ -4994,7 +4832,6 @@ func (m *awsRestjson1_serializeOpListInfrastructureCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInfrastructureConfigurationsInput(v *ListInfrastructureConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -5039,10 +4876,8 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
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
@@ -5087,7 +4922,6 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLifecycleExecutionResourcesInput(v *ListLifecycleExecutionResourcesInput, encoder *httpbinding.Encoder) error {
@@ -5135,10 +4969,8 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
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
@@ -5183,7 +5015,6 @@ func (m *awsRestjson1_serializeOpListLifecycleExecutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLifecycleExecutionsInput(v *ListLifecycleExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -5226,10 +5057,8 @@ func (m *awsRestjson1_serializeOpListLifecyclePolicies
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
@@ -5274,7 +5103,6 @@ func (m *awsRestjson1_serializeOpListLifecyclePolicies
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLifecyclePoliciesInput(v *ListLifecyclePoliciesInput, encoder *httpbinding.Encoder) error {
@@ -5319,10 +5147,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -5360,7 +5186,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -5390,10 +5215,8 @@ func (m *awsRestjson1_serializeOpListWaitingWorkflowSt
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
@@ -5438,7 +5261,6 @@ func (m *awsRestjson1_serializeOpListWaitingWorkflowSt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWaitingWorkflowStepsInput(v *ListWaitingWorkflowStepsInput, encoder *httpbinding.Encoder) error {
@@ -5476,10 +5298,8 @@ func (m *awsRestjson1_serializeOpListWorkflowBuildVers
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
@@ -5524,7 +5344,6 @@ func (m *awsRestjson1_serializeOpListWorkflowBuildVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowBuildVersionsInput(v *ListWorkflowBuildVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5567,10 +5386,8 @@ func (m *awsRestjson1_serializeOpListWorkflowExecution
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
@@ -5615,7 +5432,6 @@ func (m *awsRestjson1_serializeOpListWorkflowExecution
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowExecutionsInput(v *ListWorkflowExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -5658,10 +5474,8 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
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
@@ -5706,7 +5520,6 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowsInput(v *ListWorkflowsInput, encoder *httpbinding.Encoder) error {
@@ -5761,10 +5574,8 @@ func (m *awsRestjson1_serializeOpListWorkflowStepExecu
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
@@ -5809,7 +5620,6 @@ func (m *awsRestjson1_serializeOpListWorkflowStepExecu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowStepExecutionsInput(v *ListWorkflowStepExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -5852,10 +5662,8 @@ func (m *awsRestjson1_serializeOpPutComponentPolicy) H
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
@@ -5900,7 +5708,6 @@ func (m *awsRestjson1_serializeOpPutComponentPolicy) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutComponentPolicyInput(v *PutComponentPolicyInput, encoder *httpbinding.Encoder) error {
@@ -5938,10 +5745,8 @@ func (m *awsRestjson1_serializeOpPutContainerRecipePol
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
@@ -5986,7 +5791,6 @@ func (m *awsRestjson1_serializeOpPutContainerRecipePol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutContainerRecipePolicyInput(v *PutContainerRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -6024,10 +5828,8 @@ func (m *awsRestjson1_serializeOpPutImagePolicy) Handl
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
@@ -6072,7 +5874,6 @@ func (m *awsRestjson1_serializeOpPutImagePolicy) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutImagePolicyInput(v *PutImagePolicyInput, encoder *httpbinding.Encoder) error {
@@ -6110,10 +5911,8 @@ func (m *awsRestjson1_serializeOpPutImageRecipePolicy)
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
@@ -6158,7 +5957,6 @@ func (m *awsRestjson1_serializeOpPutImageRecipePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutImageRecipePolicyInput(v *PutImageRecipePolicyInput, encoder *httpbinding.Encoder) error {
@@ -6196,10 +5994,8 @@ func (m *awsRestjson1_serializeOpRetryImage) HandleSer
 func (m *awsRestjson1_serializeOpRetryImage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6244,7 +6040,6 @@ func (m *awsRestjson1_serializeOpRetryImage) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRetryImageInput(v *RetryImageInput, encoder *httpbinding.Encoder) error {
@@ -6282,10 +6077,8 @@ func (m *awsRestjson1_serializeOpSendWorkflowStepActio
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
@@ -6330,7 +6123,6 @@ func (m *awsRestjson1_serializeOpSendWorkflowStepActio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendWorkflowStepActionInput(v *SendWorkflowStepActionInput, encoder *httpbinding.Encoder) error {
@@ -6383,10 +6175,8 @@ func (m *awsRestjson1_serializeOpStartImagePipelineExe
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
@@ -6431,7 +6221,6 @@ func (m *awsRestjson1_serializeOpStartImagePipelineExe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartImagePipelineExecutionInput(v *StartImagePipelineExecutionInput, encoder *httpbinding.Encoder) error {
@@ -6476,10 +6265,8 @@ func (m *awsRestjson1_serializeOpStartResourceStateUpd
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
@@ -6524,7 +6311,6 @@ func (m *awsRestjson1_serializeOpStartResourceStateUpd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartResourceStateUpdateInput(v *StartResourceStateUpdateInput, encoder *httpbinding.Encoder) error {
@@ -6593,10 +6379,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -6645,7 +6429,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6689,10 +6472,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -6730,7 +6511,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6766,10 +6546,8 @@ func (m *awsRestjson1_serializeOpUpdateDistributionCon
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
@@ -6814,7 +6592,6 @@ func (m *awsRestjson1_serializeOpUpdateDistributionCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDistributionConfigurationInput(v *UpdateDistributionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6864,10 +6641,8 @@ func (m *awsRestjson1_serializeOpUpdateImagePipeline) 
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
@@ -6912,7 +6687,6 @@ func (m *awsRestjson1_serializeOpUpdateImagePipeline) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateImagePipelineInput(v *UpdateImagePipelineInput, encoder *httpbinding.Encoder) error {
@@ -7032,10 +6806,8 @@ func (m *awsRestjson1_serializeOpUpdateInfrastructureC
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
@@ -7080,7 +6852,6 @@ func (m *awsRestjson1_serializeOpUpdateInfrastructureC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateInfrastructureConfigurationInput(v *UpdateInfrastructureConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7190,10 +6961,8 @@ func (m *awsRestjson1_serializeOpUpdateLifecyclePolicy
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
@@ -7238,7 +7007,6 @@ func (m *awsRestjson1_serializeOpUpdateLifecyclePolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLifecyclePolicyInput(v *UpdateLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
