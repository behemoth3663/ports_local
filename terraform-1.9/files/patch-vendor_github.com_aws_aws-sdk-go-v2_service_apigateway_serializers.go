--- vendor/github.com/aws/aws-sdk-go-v2/service/apigateway/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/apigateway/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCreateApiKey) HandleS
 func (m *awsRestjson1_serializeOpCreateApiKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpCreateApiKey) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateApiKeyInput(v *CreateApiKeyInput, encoder *httpbinding.Encoder) error {
@@ -146,10 +142,8 @@ func (m *awsRestjson1_serializeOpCreateAuthorizer) Han
 func (m *awsRestjson1_serializeOpCreateAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -198,7 +192,6 @@ func (m *awsRestjson1_serializeOpCreateAuthorizer) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAuthorizerInput(v *CreateAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -282,10 +275,8 @@ func (m *awsRestjson1_serializeOpCreateBasePathMapping
 func (m *awsRestjson1_serializeOpCreateBasePathMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -334,7 +325,6 @@ func (m *awsRestjson1_serializeOpCreateBasePathMapping
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBasePathMappingInput(v *CreateBasePathMappingInput, encoder *httpbinding.Encoder) error {
@@ -390,10 +380,8 @@ func (m *awsRestjson1_serializeOpCreateDeployment) Han
 func (m *awsRestjson1_serializeOpCreateDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -442,7 +430,6 @@ func (m *awsRestjson1_serializeOpCreateDeployment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeploymentInput(v *CreateDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -523,10 +510,8 @@ func (m *awsRestjson1_serializeOpCreateDocumentationPa
 func (m *awsRestjson1_serializeOpCreateDocumentationPart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -575,7 +560,6 @@ func (m *awsRestjson1_serializeOpCreateDocumentationPa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDocumentationPartInput(v *CreateDocumentationPartInput, encoder *httpbinding.Encoder) error {
@@ -624,10 +608,8 @@ func (m *awsRestjson1_serializeOpCreateDocumentationVe
 func (m *awsRestjson1_serializeOpCreateDocumentationVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -676,7 +658,6 @@ func (m *awsRestjson1_serializeOpCreateDocumentationVe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDocumentationVersionInput(v *CreateDocumentationVersionInput, encoder *httpbinding.Encoder) error {
@@ -728,10 +709,8 @@ func (m *awsRestjson1_serializeOpCreateDomainName) Han
 func (m *awsRestjson1_serializeOpCreateDomainName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -776,7 +755,6 @@ func (m *awsRestjson1_serializeOpCreateDomainName) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainNameInput(v *CreateDomainNameInput, encoder *httpbinding.Encoder) error {
@@ -885,10 +863,8 @@ func (m *awsRestjson1_serializeOpCreateDomainNameAcces
 func (m *awsRestjson1_serializeOpCreateDomainNameAccessAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -933,7 +909,6 @@ func (m *awsRestjson1_serializeOpCreateDomainNameAcces
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainNameAccessAssociationInput(v *CreateDomainNameAccessAssociationInput, encoder *httpbinding.Encoder) error {
@@ -983,10 +958,8 @@ func (m *awsRestjson1_serializeOpCreateModel) HandleSe
 func (m *awsRestjson1_serializeOpCreateModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1035,7 +1008,6 @@ func (m *awsRestjson1_serializeOpCreateModel) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateModelInput(v *CreateModelInput, encoder *httpbinding.Encoder) error {
@@ -1092,10 +1064,8 @@ func (m *awsRestjson1_serializeOpCreateRequestValidato
 func (m *awsRestjson1_serializeOpCreateRequestValidator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1144,7 +1114,6 @@ func (m *awsRestjson1_serializeOpCreateRequestValidato
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRequestValidatorInput(v *CreateRequestValidatorInput, encoder *httpbinding.Encoder) error {
@@ -1196,10 +1165,8 @@ func (m *awsRestjson1_serializeOpCreateResource) Handl
 func (m *awsRestjson1_serializeOpCreateResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1248,7 +1215,6 @@ func (m *awsRestjson1_serializeOpCreateResource) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateResourceInput(v *CreateResourceInput, encoder *httpbinding.Encoder) error {
@@ -1299,10 +1265,8 @@ func (m *awsRestjson1_serializeOpCreateRestApi) Handle
 func (m *awsRestjson1_serializeOpCreateRestApi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1347,7 +1311,6 @@ func (m *awsRestjson1_serializeOpCreateRestApi) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRestApiInput(v *CreateRestApiInput, encoder *httpbinding.Encoder) error {
@@ -1436,10 +1399,8 @@ func (m *awsRestjson1_serializeOpCreateStage) HandleSe
 func (m *awsRestjson1_serializeOpCreateStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1488,7 +1449,6 @@ func (m *awsRestjson1_serializeOpCreateStage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateStageInput(v *CreateStageInput, encoder *httpbinding.Encoder) error {
@@ -1581,10 +1541,8 @@ func (m *awsRestjson1_serializeOpCreateUsagePlan) Hand
 func (m *awsRestjson1_serializeOpCreateUsagePlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1629,7 +1587,6 @@ func (m *awsRestjson1_serializeOpCreateUsagePlan) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUsagePlanInput(v *CreateUsagePlanInput, encoder *httpbinding.Encoder) error {
@@ -1695,10 +1652,8 @@ func (m *awsRestjson1_serializeOpCreateUsagePlanKey) H
 func (m *awsRestjson1_serializeOpCreateUsagePlanKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1747,7 +1702,6 @@ func (m *awsRestjson1_serializeOpCreateUsagePlanKey) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUsagePlanKeyInput(v *CreateUsagePlanKeyInput, encoder *httpbinding.Encoder) error {
@@ -1794,10 +1748,8 @@ func (m *awsRestjson1_serializeOpCreateVpcLink) Handle
 func (m *awsRestjson1_serializeOpCreateVpcLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1842,7 +1794,6 @@ func (m *awsRestjson1_serializeOpCreateVpcLink) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVpcLinkInput(v *CreateVpcLinkInput, encoder *httpbinding.Encoder) error {
@@ -1894,10 +1845,8 @@ func (m *awsRestjson1_serializeOpDeleteApiKey) HandleS
 func (m *awsRestjson1_serializeOpDeleteApiKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1935,7 +1884,6 @@ func (m *awsRestjson1_serializeOpDeleteApiKey) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteApiKeyInput(v *DeleteApiKeyInput, encoder *httpbinding.Encoder) error {
@@ -1965,10 +1913,8 @@ func (m *awsRestjson1_serializeOpDeleteAuthorizer) Han
 func (m *awsRestjson1_serializeOpDeleteAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2006,7 +1952,6 @@ func (m *awsRestjson1_serializeOpDeleteAuthorizer) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAuthorizerInput(v *DeleteAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -2045,10 +1990,8 @@ func (m *awsRestjson1_serializeOpDeleteBasePathMapping
 func (m *awsRestjson1_serializeOpDeleteBasePathMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2086,7 +2029,6 @@ func (m *awsRestjson1_serializeOpDeleteBasePathMapping
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBasePathMappingInput(v *DeleteBasePathMappingInput, encoder *httpbinding.Encoder) error {
@@ -2129,10 +2071,8 @@ func (m *awsRestjson1_serializeOpDeleteClientCertifica
 func (m *awsRestjson1_serializeOpDeleteClientCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2170,7 +2110,6 @@ func (m *awsRestjson1_serializeOpDeleteClientCertifica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteClientCertificateInput(v *DeleteClientCertificateInput, encoder *httpbinding.Encoder) error {
@@ -2200,10 +2139,8 @@ func (m *awsRestjson1_serializeOpDeleteDeployment) Han
 func (m *awsRestjson1_serializeOpDeleteDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2241,7 +2178,6 @@ func (m *awsRestjson1_serializeOpDeleteDeployment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDeploymentInput(v *DeleteDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -2280,10 +2216,8 @@ func (m *awsRestjson1_serializeOpDeleteDocumentationPa
 func (m *awsRestjson1_serializeOpDeleteDocumentationPart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2321,7 +2255,6 @@ func (m *awsRestjson1_serializeOpDeleteDocumentationPa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDocumentationPartInput(v *DeleteDocumentationPartInput, encoder *httpbinding.Encoder) error {
@@ -2360,10 +2293,8 @@ func (m *awsRestjson1_serializeOpDeleteDocumentationVe
 func (m *awsRestjson1_serializeOpDeleteDocumentationVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2401,7 +2332,6 @@ func (m *awsRestjson1_serializeOpDeleteDocumentationVe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDocumentationVersionInput(v *DeleteDocumentationVersionInput, encoder *httpbinding.Encoder) error {
@@ -2440,10 +2370,8 @@ func (m *awsRestjson1_serializeOpDeleteDomainName) Han
 func (m *awsRestjson1_serializeOpDeleteDomainName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2481,7 +2409,6 @@ func (m *awsRestjson1_serializeOpDeleteDomainName) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainNameInput(v *DeleteDomainNameInput, encoder *httpbinding.Encoder) error {
@@ -2515,10 +2442,8 @@ func (m *awsRestjson1_serializeOpDeleteDomainNameAcces
 func (m *awsRestjson1_serializeOpDeleteDomainNameAccessAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2556,7 +2481,6 @@ func (m *awsRestjson1_serializeOpDeleteDomainNameAcces
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainNameAccessAssociationInput(v *DeleteDomainNameAccessAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2586,10 +2510,8 @@ func (m *awsRestjson1_serializeOpDeleteGatewayResponse
 func (m *awsRestjson1_serializeOpDeleteGatewayResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2627,7 +2549,6 @@ func (m *awsRestjson1_serializeOpDeleteGatewayResponse
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGatewayResponseInput(v *DeleteGatewayResponseInput, encoder *httpbinding.Encoder) error {
@@ -2666,10 +2587,8 @@ func (m *awsRestjson1_serializeOpDeleteIntegration) Ha
 func (m *awsRestjson1_serializeOpDeleteIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2707,7 +2626,6 @@ func (m *awsRestjson1_serializeOpDeleteIntegration) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIntegrationInput(v *DeleteIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -2755,10 +2673,8 @@ func (m *awsRestjson1_serializeOpDeleteIntegrationResp
 func (m *awsRestjson1_serializeOpDeleteIntegrationResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2796,7 +2712,6 @@ func (m *awsRestjson1_serializeOpDeleteIntegrationResp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIntegrationResponseInput(v *DeleteIntegrationResponseInput, encoder *httpbinding.Encoder) error {
@@ -2853,10 +2768,8 @@ func (m *awsRestjson1_serializeOpDeleteMethod) HandleS
 func (m *awsRestjson1_serializeOpDeleteMethod) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2894,7 +2807,6 @@ func (m *awsRestjson1_serializeOpDeleteMethod) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMethodInput(v *DeleteMethodInput, encoder *httpbinding.Encoder) error {
@@ -2942,10 +2854,8 @@ func (m *awsRestjson1_serializeOpDeleteMethodResponse)
 func (m *awsRestjson1_serializeOpDeleteMethodResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2983,7 +2893,6 @@ func (m *awsRestjson1_serializeOpDeleteMethodResponse)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMethodResponseInput(v *DeleteMethodResponseInput, encoder *httpbinding.Encoder) error {
@@ -3040,10 +2949,8 @@ func (m *awsRestjson1_serializeOpDeleteModel) HandleSe
 func (m *awsRestjson1_serializeOpDeleteModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3081,7 +2988,6 @@ func (m *awsRestjson1_serializeOpDeleteModel) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteModelInput(v *DeleteModelInput, encoder *httpbinding.Encoder) error {
@@ -3120,10 +3026,8 @@ func (m *awsRestjson1_serializeOpDeleteRequestValidato
 func (m *awsRestjson1_serializeOpDeleteRequestValidator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3161,7 +3065,6 @@ func (m *awsRestjson1_serializeOpDeleteRequestValidato
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRequestValidatorInput(v *DeleteRequestValidatorInput, encoder *httpbinding.Encoder) error {
@@ -3200,10 +3103,8 @@ func (m *awsRestjson1_serializeOpDeleteResource) Handl
 func (m *awsRestjson1_serializeOpDeleteResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3241,7 +3142,6 @@ func (m *awsRestjson1_serializeOpDeleteResource) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResourceInput(v *DeleteResourceInput, encoder *httpbinding.Encoder) error {
@@ -3280,10 +3180,8 @@ func (m *awsRestjson1_serializeOpDeleteRestApi) Handle
 func (m *awsRestjson1_serializeOpDeleteRestApi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3321,7 +3219,6 @@ func (m *awsRestjson1_serializeOpDeleteRestApi) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRestApiInput(v *DeleteRestApiInput, encoder *httpbinding.Encoder) error {
@@ -3351,10 +3248,8 @@ func (m *awsRestjson1_serializeOpDeleteStage) HandleSe
 func (m *awsRestjson1_serializeOpDeleteStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3392,7 +3287,6 @@ func (m *awsRestjson1_serializeOpDeleteStage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteStageInput(v *DeleteStageInput, encoder *httpbinding.Encoder) error {
@@ -3431,10 +3325,8 @@ func (m *awsRestjson1_serializeOpDeleteUsagePlan) Hand
 func (m *awsRestjson1_serializeOpDeleteUsagePlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3472,7 +3364,6 @@ func (m *awsRestjson1_serializeOpDeleteUsagePlan) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUsagePlanInput(v *DeleteUsagePlanInput, encoder *httpbinding.Encoder) error {
@@ -3502,10 +3393,8 @@ func (m *awsRestjson1_serializeOpDeleteUsagePlanKey) H
 func (m *awsRestjson1_serializeOpDeleteUsagePlanKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3543,7 +3432,6 @@ func (m *awsRestjson1_serializeOpDeleteUsagePlanKey) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUsagePlanKeyInput(v *DeleteUsagePlanKeyInput, encoder *httpbinding.Encoder) error {
@@ -3582,10 +3470,8 @@ func (m *awsRestjson1_serializeOpDeleteVpcLink) Handle
 func (m *awsRestjson1_serializeOpDeleteVpcLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3623,7 +3509,6 @@ func (m *awsRestjson1_serializeOpDeleteVpcLink) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVpcLinkInput(v *DeleteVpcLinkInput, encoder *httpbinding.Encoder) error {
@@ -3653,10 +3538,8 @@ func (m *awsRestjson1_serializeOpFlushStageAuthorizers
 func (m *awsRestjson1_serializeOpFlushStageAuthorizersCache) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3694,7 +3577,6 @@ func (m *awsRestjson1_serializeOpFlushStageAuthorizers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsFlushStageAuthorizersCacheInput(v *FlushStageAuthorizersCacheInput, encoder *httpbinding.Encoder) error {
@@ -3733,10 +3615,8 @@ func (m *awsRestjson1_serializeOpFlushStageCache) Hand
 func (m *awsRestjson1_serializeOpFlushStageCache) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3774,7 +3654,6 @@ func (m *awsRestjson1_serializeOpFlushStageCache) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsFlushStageCacheInput(v *FlushStageCacheInput, encoder *httpbinding.Encoder) error {
@@ -3813,10 +3692,8 @@ func (m *awsRestjson1_serializeOpGenerateClientCertifi
 func (m *awsRestjson1_serializeOpGenerateClientCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3861,7 +3738,6 @@ func (m *awsRestjson1_serializeOpGenerateClientCertifi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGenerateClientCertificateInput(v *GenerateClientCertificateInput, encoder *httpbinding.Encoder) error {
@@ -3901,10 +3777,8 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 func (m *awsRestjson1_serializeOpGetAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3938,7 +3812,6 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountInput(v *GetAccountInput, encoder *httpbinding.Encoder) error {
@@ -3959,10 +3832,8 @@ func (m *awsRestjson1_serializeOpGetApiKey) HandleSeri
 func (m *awsRestjson1_serializeOpGetApiKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4000,7 +3871,6 @@ func (m *awsRestjson1_serializeOpGetApiKey) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApiKeyInput(v *GetApiKeyInput, encoder *httpbinding.Encoder) error {
@@ -4034,10 +3904,8 @@ func (m *awsRestjson1_serializeOpGetApiKeys) HandleSer
 func (m *awsRestjson1_serializeOpGetApiKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4075,7 +3943,6 @@ func (m *awsRestjson1_serializeOpGetApiKeys) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApiKeysInput(v *GetApiKeysInput, encoder *httpbinding.Encoder) error {
@@ -4116,10 +3983,8 @@ func (m *awsRestjson1_serializeOpGetAuthorizer) Handle
 func (m *awsRestjson1_serializeOpGetAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4157,7 +4022,6 @@ func (m *awsRestjson1_serializeOpGetAuthorizer) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAuthorizerInput(v *GetAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -4196,10 +4060,8 @@ func (m *awsRestjson1_serializeOpGetAuthorizers) Handl
 func (m *awsRestjson1_serializeOpGetAuthorizers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4237,7 +4099,6 @@ func (m *awsRestjson1_serializeOpGetAuthorizers) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAuthorizersInput(v *GetAuthorizersInput, encoder *httpbinding.Encoder) error {
@@ -4275,10 +4136,8 @@ func (m *awsRestjson1_serializeOpGetBasePathMapping) H
 func (m *awsRestjson1_serializeOpGetBasePathMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4316,7 +4175,6 @@ func (m *awsRestjson1_serializeOpGetBasePathMapping) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBasePathMappingInput(v *GetBasePathMappingInput, encoder *httpbinding.Encoder) error {
@@ -4359,10 +4217,8 @@ func (m *awsRestjson1_serializeOpGetBasePathMappings) 
 func (m *awsRestjson1_serializeOpGetBasePathMappings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4400,7 +4256,6 @@ func (m *awsRestjson1_serializeOpGetBasePathMappings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBasePathMappingsInput(v *GetBasePathMappingsInput, encoder *httpbinding.Encoder) error {
@@ -4442,10 +4297,8 @@ func (m *awsRestjson1_serializeOpGetClientCertificate)
 func (m *awsRestjson1_serializeOpGetClientCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4483,7 +4336,6 @@ func (m *awsRestjson1_serializeOpGetClientCertificate)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetClientCertificateInput(v *GetClientCertificateInput, encoder *httpbinding.Encoder) error {
@@ -4513,10 +4365,8 @@ func (m *awsRestjson1_serializeOpGetClientCertificates
 func (m *awsRestjson1_serializeOpGetClientCertificates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4554,7 +4404,6 @@ func (m *awsRestjson1_serializeOpGetClientCertificates
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetClientCertificatesInput(v *GetClientCertificatesInput, encoder *httpbinding.Encoder) error {
@@ -4583,10 +4432,8 @@ func (m *awsRestjson1_serializeOpGetDeployment) Handle
 func (m *awsRestjson1_serializeOpGetDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4624,7 +4471,6 @@ func (m *awsRestjson1_serializeOpGetDeployment) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeploymentInput(v *GetDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -4669,10 +4515,8 @@ func (m *awsRestjson1_serializeOpGetDeployments) Handl
 func (m *awsRestjson1_serializeOpGetDeployments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4710,7 +4554,6 @@ func (m *awsRestjson1_serializeOpGetDeployments) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeploymentsInput(v *GetDeploymentsInput, encoder *httpbinding.Encoder) error {
@@ -4748,10 +4591,8 @@ func (m *awsRestjson1_serializeOpGetDocumentationPart)
 func (m *awsRestjson1_serializeOpGetDocumentationPart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4789,7 +4630,6 @@ func (m *awsRestjson1_serializeOpGetDocumentationPart)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDocumentationPartInput(v *GetDocumentationPartInput, encoder *httpbinding.Encoder) error {
@@ -4828,10 +4668,8 @@ func (m *awsRestjson1_serializeOpGetDocumentationParts
 func (m *awsRestjson1_serializeOpGetDocumentationParts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4869,7 +4707,6 @@ func (m *awsRestjson1_serializeOpGetDocumentationParts
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDocumentationPartsInput(v *GetDocumentationPartsInput, encoder *httpbinding.Encoder) error {
@@ -4923,10 +4760,8 @@ func (m *awsRestjson1_serializeOpGetDocumentationVersi
 func (m *awsRestjson1_serializeOpGetDocumentationVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4964,7 +4799,6 @@ func (m *awsRestjson1_serializeOpGetDocumentationVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDocumentationVersionInput(v *GetDocumentationVersionInput, encoder *httpbinding.Encoder) error {
@@ -5003,10 +4837,8 @@ func (m *awsRestjson1_serializeOpGetDocumentationVersi
 func (m *awsRestjson1_serializeOpGetDocumentationVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5044,7 +4876,6 @@ func (m *awsRestjson1_serializeOpGetDocumentationVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDocumentationVersionsInput(v *GetDocumentationVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5082,10 +4913,8 @@ func (m *awsRestjson1_serializeOpGetDomainName) Handle
 func (m *awsRestjson1_serializeOpGetDomainName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5123,7 +4952,6 @@ func (m *awsRestjson1_serializeOpGetDomainName) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainNameInput(v *GetDomainNameInput, encoder *httpbinding.Encoder) error {
@@ -5157,10 +4985,8 @@ func (m *awsRestjson1_serializeOpGetDomainNameAccessAs
 func (m *awsRestjson1_serializeOpGetDomainNameAccessAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5198,7 +5024,6 @@ func (m *awsRestjson1_serializeOpGetDomainNameAccessAs
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainNameAccessAssociationsInput(v *GetDomainNameAccessAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -5231,10 +5056,8 @@ func (m *awsRestjson1_serializeOpGetDomainNames) Handl
 func (m *awsRestjson1_serializeOpGetDomainNames) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5272,7 +5095,6 @@ func (m *awsRestjson1_serializeOpGetDomainNames) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainNamesInput(v *GetDomainNamesInput, encoder *httpbinding.Encoder) error {
@@ -5305,10 +5127,8 @@ func (m *awsRestjson1_serializeOpGetExport) HandleSeri
 func (m *awsRestjson1_serializeOpGetExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5346,7 +5166,6 @@ func (m *awsRestjson1_serializeOpGetExport) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetExportInput(v *GetExportInput, encoder *httpbinding.Encoder) error {
@@ -5408,10 +5227,8 @@ func (m *awsRestjson1_serializeOpGetGatewayResponse) H
 func (m *awsRestjson1_serializeOpGetGatewayResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5449,7 +5266,6 @@ func (m *awsRestjson1_serializeOpGetGatewayResponse) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGatewayResponseInput(v *GetGatewayResponseInput, encoder *httpbinding.Encoder) error {
@@ -5488,10 +5304,8 @@ func (m *awsRestjson1_serializeOpGetGatewayResponses) 
 func (m *awsRestjson1_serializeOpGetGatewayResponses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5529,7 +5343,6 @@ func (m *awsRestjson1_serializeOpGetGatewayResponses) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGatewayResponsesInput(v *GetGatewayResponsesInput, encoder *httpbinding.Encoder) error {
@@ -5567,10 +5380,8 @@ func (m *awsRestjson1_serializeOpGetIntegration) Handl
 func (m *awsRestjson1_serializeOpGetIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5608,7 +5419,6 @@ func (m *awsRestjson1_serializeOpGetIntegration) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIntegrationInput(v *GetIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -5656,10 +5466,8 @@ func (m *awsRestjson1_serializeOpGetIntegrationRespons
 func (m *awsRestjson1_serializeOpGetIntegrationResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5697,7 +5505,6 @@ func (m *awsRestjson1_serializeOpGetIntegrationRespons
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIntegrationResponseInput(v *GetIntegrationResponseInput, encoder *httpbinding.Encoder) error {
@@ -5754,10 +5561,8 @@ func (m *awsRestjson1_serializeOpGetMethod) HandleSeri
 func (m *awsRestjson1_serializeOpGetMethod) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5795,7 +5600,6 @@ func (m *awsRestjson1_serializeOpGetMethod) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMethodInput(v *GetMethodInput, encoder *httpbinding.Encoder) error {
@@ -5843,10 +5647,8 @@ func (m *awsRestjson1_serializeOpGetMethodResponse) Ha
 func (m *awsRestjson1_serializeOpGetMethodResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5884,7 +5686,6 @@ func (m *awsRestjson1_serializeOpGetMethodResponse) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMethodResponseInput(v *GetMethodResponseInput, encoder *httpbinding.Encoder) error {
@@ -5941,10 +5742,8 @@ func (m *awsRestjson1_serializeOpGetModel) HandleSeria
 func (m *awsRestjson1_serializeOpGetModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5982,7 +5781,6 @@ func (m *awsRestjson1_serializeOpGetModel) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelInput(v *GetModelInput, encoder *httpbinding.Encoder) error {
@@ -6025,10 +5823,8 @@ func (m *awsRestjson1_serializeOpGetModels) HandleSeri
 func (m *awsRestjson1_serializeOpGetModels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6066,7 +5862,6 @@ func (m *awsRestjson1_serializeOpGetModels) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelsInput(v *GetModelsInput, encoder *httpbinding.Encoder) error {
@@ -6104,10 +5899,8 @@ func (m *awsRestjson1_serializeOpGetModelTemplate) Han
 func (m *awsRestjson1_serializeOpGetModelTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6145,7 +5938,6 @@ func (m *awsRestjson1_serializeOpGetModelTemplate) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetModelTemplateInput(v *GetModelTemplateInput, encoder *httpbinding.Encoder) error {
@@ -6184,10 +5976,8 @@ func (m *awsRestjson1_serializeOpGetRequestValidator) 
 func (m *awsRestjson1_serializeOpGetRequestValidator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6225,7 +6015,6 @@ func (m *awsRestjson1_serializeOpGetRequestValidator) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRequestValidatorInput(v *GetRequestValidatorInput, encoder *httpbinding.Encoder) error {
@@ -6264,10 +6053,8 @@ func (m *awsRestjson1_serializeOpGetRequestValidators)
 func (m *awsRestjson1_serializeOpGetRequestValidators) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6305,7 +6092,6 @@ func (m *awsRestjson1_serializeOpGetRequestValidators)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRequestValidatorsInput(v *GetRequestValidatorsInput, encoder *httpbinding.Encoder) error {
@@ -6343,10 +6129,8 @@ func (m *awsRestjson1_serializeOpGetResource) HandleSe
 func (m *awsRestjson1_serializeOpGetResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6384,7 +6168,6 @@ func (m *awsRestjson1_serializeOpGetResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceInput(v *GetResourceInput, encoder *httpbinding.Encoder) error {
@@ -6429,10 +6212,8 @@ func (m *awsRestjson1_serializeOpGetResources) HandleS
 func (m *awsRestjson1_serializeOpGetResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6470,7 +6251,6 @@ func (m *awsRestjson1_serializeOpGetResources) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourcesInput(v *GetResourcesInput, encoder *httpbinding.Encoder) error {
@@ -6514,10 +6294,8 @@ func (m *awsRestjson1_serializeOpGetRestApi) HandleSer
 func (m *awsRestjson1_serializeOpGetRestApi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6555,7 +6333,6 @@ func (m *awsRestjson1_serializeOpGetRestApi) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRestApiInput(v *GetRestApiInput, encoder *httpbinding.Encoder) error {
@@ -6585,10 +6362,8 @@ func (m *awsRestjson1_serializeOpGetRestApis) HandleSe
 func (m *awsRestjson1_serializeOpGetRestApis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6626,7 +6401,6 @@ func (m *awsRestjson1_serializeOpGetRestApis) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRestApisInput(v *GetRestApisInput, encoder *httpbinding.Encoder) error {
@@ -6655,10 +6429,8 @@ func (m *awsRestjson1_serializeOpGetSdk) HandleSeriali
 func (m *awsRestjson1_serializeOpGetSdk) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6696,7 +6468,6 @@ func (m *awsRestjson1_serializeOpGetSdk) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSdkInput(v *GetSdkInput, encoder *httpbinding.Encoder) error {
@@ -6753,10 +6524,8 @@ func (m *awsRestjson1_serializeOpGetSdkType) HandleSer
 func (m *awsRestjson1_serializeOpGetSdkType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6794,7 +6563,6 @@ func (m *awsRestjson1_serializeOpGetSdkType) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSdkTypeInput(v *GetSdkTypeInput, encoder *httpbinding.Encoder) error {
@@ -6824,10 +6592,8 @@ func (m *awsRestjson1_serializeOpGetSdkTypes) HandleSe
 func (m *awsRestjson1_serializeOpGetSdkTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6865,7 +6631,6 @@ func (m *awsRestjson1_serializeOpGetSdkTypes) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSdkTypesInput(v *GetSdkTypesInput, encoder *httpbinding.Encoder) error {
@@ -6894,10 +6659,8 @@ func (m *awsRestjson1_serializeOpGetStage) HandleSeria
 func (m *awsRestjson1_serializeOpGetStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6935,7 +6698,6 @@ func (m *awsRestjson1_serializeOpGetStage) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStageInput(v *GetStageInput, encoder *httpbinding.Encoder) error {
@@ -6974,10 +6736,8 @@ func (m *awsRestjson1_serializeOpGetStages) HandleSeri
 func (m *awsRestjson1_serializeOpGetStages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7015,7 +6775,6 @@ func (m *awsRestjson1_serializeOpGetStages) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStagesInput(v *GetStagesInput, encoder *httpbinding.Encoder) error {
@@ -7049,10 +6808,8 @@ func (m *awsRestjson1_serializeOpGetTags) HandleSerial
 func (m *awsRestjson1_serializeOpGetTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7090,7 +6847,6 @@ func (m *awsRestjson1_serializeOpGetTags) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTagsInput(v *GetTagsInput, encoder *httpbinding.Encoder) error {
@@ -7128,10 +6884,8 @@ func (m *awsRestjson1_serializeOpGetUsage) HandleSeria
 func (m *awsRestjson1_serializeOpGetUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7169,7 +6923,6 @@ func (m *awsRestjson1_serializeOpGetUsage) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsageInput(v *GetUsageInput, encoder *httpbinding.Encoder) error {
@@ -7219,10 +6972,8 @@ func (m *awsRestjson1_serializeOpGetUsagePlan) HandleS
 func (m *awsRestjson1_serializeOpGetUsagePlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7260,7 +7011,6 @@ func (m *awsRestjson1_serializeOpGetUsagePlan) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsagePlanInput(v *GetUsagePlanInput, encoder *httpbinding.Encoder) error {
@@ -7290,10 +7040,8 @@ func (m *awsRestjson1_serializeOpGetUsagePlanKey) Hand
 func (m *awsRestjson1_serializeOpGetUsagePlanKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7331,7 +7079,6 @@ func (m *awsRestjson1_serializeOpGetUsagePlanKey) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsagePlanKeyInput(v *GetUsagePlanKeyInput, encoder *httpbinding.Encoder) error {
@@ -7370,10 +7117,8 @@ func (m *awsRestjson1_serializeOpGetUsagePlanKeys) Han
 func (m *awsRestjson1_serializeOpGetUsagePlanKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7411,7 +7156,6 @@ func (m *awsRestjson1_serializeOpGetUsagePlanKeys) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsagePlanKeysInput(v *GetUsagePlanKeysInput, encoder *httpbinding.Encoder) error {
@@ -7453,10 +7197,8 @@ func (m *awsRestjson1_serializeOpGetUsagePlans) Handle
 func (m *awsRestjson1_serializeOpGetUsagePlans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7494,7 +7236,6 @@ func (m *awsRestjson1_serializeOpGetUsagePlans) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsagePlansInput(v *GetUsagePlansInput, encoder *httpbinding.Encoder) error {
@@ -7527,10 +7268,8 @@ func (m *awsRestjson1_serializeOpGetVpcLink) HandleSer
 func (m *awsRestjson1_serializeOpGetVpcLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7568,7 +7307,6 @@ func (m *awsRestjson1_serializeOpGetVpcLink) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVpcLinkInput(v *GetVpcLinkInput, encoder *httpbinding.Encoder) error {
@@ -7598,10 +7336,8 @@ func (m *awsRestjson1_serializeOpGetVpcLinks) HandleSe
 func (m *awsRestjson1_serializeOpGetVpcLinks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7639,7 +7375,6 @@ func (m *awsRestjson1_serializeOpGetVpcLinks) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVpcLinksInput(v *GetVpcLinksInput, encoder *httpbinding.Encoder) error {
@@ -7668,10 +7403,8 @@ func (m *awsRestjson1_serializeOpImportApiKeys) Handle
 func (m *awsRestjson1_serializeOpImportApiKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7721,7 +7454,6 @@ func (m *awsRestjson1_serializeOpImportApiKeys) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportApiKeysInput(v *ImportApiKeysInput, encoder *httpbinding.Encoder) error {
@@ -7750,10 +7482,8 @@ func (m *awsRestjson1_serializeOpImportDocumentationPa
 func (m *awsRestjson1_serializeOpImportDocumentationParts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7803,7 +7533,6 @@ func (m *awsRestjson1_serializeOpImportDocumentationPa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportDocumentationPartsInput(v *ImportDocumentationPartsInput, encoder *httpbinding.Encoder) error {
@@ -7841,10 +7570,8 @@ func (m *awsRestjson1_serializeOpImportRestApi) Handle
 func (m *awsRestjson1_serializeOpImportRestApi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7894,7 +7621,6 @@ func (m *awsRestjson1_serializeOpImportRestApi) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportRestApiInput(v *ImportRestApiInput, encoder *httpbinding.Encoder) error {
@@ -7928,10 +7654,8 @@ func (m *awsRestjson1_serializeOpPutGatewayResponse) H
 func (m *awsRestjson1_serializeOpPutGatewayResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7980,7 +7704,6 @@ func (m *awsRestjson1_serializeOpPutGatewayResponse) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutGatewayResponseInput(v *PutGatewayResponseInput, encoder *httpbinding.Encoder) error {
@@ -8045,10 +7768,8 @@ func (m *awsRestjson1_serializeOpPutIntegration) Handl
 func (m *awsRestjson1_serializeOpPutIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8097,7 +7818,6 @@ func (m *awsRestjson1_serializeOpPutIntegration) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutIntegrationInput(v *PutIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -8230,10 +7950,8 @@ func (m *awsRestjson1_serializeOpPutIntegrationRespons
 func (m *awsRestjson1_serializeOpPutIntegrationResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8282,7 +8000,6 @@ func (m *awsRestjson1_serializeOpPutIntegrationRespons
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutIntegrationResponseInput(v *PutIntegrationResponseInput, encoder *httpbinding.Encoder) error {
@@ -8370,10 +8087,8 @@ func (m *awsRestjson1_serializeOpPutMethod) HandleSeri
 func (m *awsRestjson1_serializeOpPutMethod) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8422,7 +8137,6 @@ func (m *awsRestjson1_serializeOpPutMethod) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutMethodInput(v *PutMethodInput, encoder *httpbinding.Encoder) error {
@@ -8523,10 +8237,8 @@ func (m *awsRestjson1_serializeOpPutMethodResponse) Ha
 func (m *awsRestjson1_serializeOpPutMethodResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8575,7 +8287,6 @@ func (m *awsRestjson1_serializeOpPutMethodResponse) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutMethodResponseInput(v *PutMethodResponseInput, encoder *httpbinding.Encoder) error {
@@ -8653,10 +8364,8 @@ func (m *awsRestjson1_serializeOpPutRestApi) HandleSer
 func (m *awsRestjson1_serializeOpPutRestApi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8706,7 +8415,6 @@ func (m *awsRestjson1_serializeOpPutRestApi) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutRestApiInput(v *PutRestApiInput, encoder *httpbinding.Encoder) error {
@@ -8753,10 +8461,8 @@ func (m *awsRestjson1_serializeOpRejectDomainNameAcces
 func (m *awsRestjson1_serializeOpRejectDomainNameAccessAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8794,7 +8500,6 @@ func (m *awsRestjson1_serializeOpRejectDomainNameAcces
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectDomainNameAccessAssociationInput(v *RejectDomainNameAccessAssociationInput, encoder *httpbinding.Encoder) error {
@@ -8823,10 +8528,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -8875,7 +8578,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -8919,10 +8621,8 @@ func (m *awsRestjson1_serializeOpTestInvokeAuthorizer)
 func (m *awsRestjson1_serializeOpTestInvokeAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8971,7 +8671,6 @@ func (m *awsRestjson1_serializeOpTestInvokeAuthorizer)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestInvokeAuthorizerInput(v *TestInvokeAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -9055,10 +8754,8 @@ func (m *awsRestjson1_serializeOpTestInvokeMethod) Han
 func (m *awsRestjson1_serializeOpTestInvokeMethod) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9107,7 +8804,6 @@ func (m *awsRestjson1_serializeOpTestInvokeMethod) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTestInvokeMethodInput(v *TestInvokeMethodInput, encoder *httpbinding.Encoder) error {
@@ -9198,10 +8894,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -9239,7 +8933,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -9275,10 +8968,8 @@ func (m *awsRestjson1_serializeOpUpdateAccount) Handle
 func (m *awsRestjson1_serializeOpUpdateAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9323,7 +9014,6 @@ func (m *awsRestjson1_serializeOpUpdateAccount) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountInput(v *UpdateAccountInput, encoder *httpbinding.Encoder) error {
@@ -9358,10 +9048,8 @@ func (m *awsRestjson1_serializeOpUpdateApiKey) HandleS
 func (m *awsRestjson1_serializeOpUpdateApiKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9410,7 +9098,6 @@ func (m *awsRestjson1_serializeOpUpdateApiKey) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApiKeyInput(v *UpdateApiKeyInput, encoder *httpbinding.Encoder) error {
@@ -9454,10 +9141,8 @@ func (m *awsRestjson1_serializeOpUpdateAuthorizer) Han
 func (m *awsRestjson1_serializeOpUpdateAuthorizer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9506,7 +9191,6 @@ func (m *awsRestjson1_serializeOpUpdateAuthorizer) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAuthorizerInput(v *UpdateAuthorizerInput, encoder *httpbinding.Encoder) error {
@@ -9559,10 +9243,8 @@ func (m *awsRestjson1_serializeOpUpdateBasePathMapping
 func (m *awsRestjson1_serializeOpUpdateBasePathMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9611,7 +9293,6 @@ func (m *awsRestjson1_serializeOpUpdateBasePathMapping
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBasePathMappingInput(v *UpdateBasePathMappingInput, encoder *httpbinding.Encoder) error {
@@ -9668,10 +9349,8 @@ func (m *awsRestjson1_serializeOpUpdateClientCertifica
 func (m *awsRestjson1_serializeOpUpdateClientCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9720,7 +9399,6 @@ func (m *awsRestjson1_serializeOpUpdateClientCertifica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateClientCertificateInput(v *UpdateClientCertificateInput, encoder *httpbinding.Encoder) error {
@@ -9764,10 +9442,8 @@ func (m *awsRestjson1_serializeOpUpdateDeployment) Han
 func (m *awsRestjson1_serializeOpUpdateDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9816,7 +9492,6 @@ func (m *awsRestjson1_serializeOpUpdateDeployment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDeploymentInput(v *UpdateDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -9869,10 +9544,8 @@ func (m *awsRestjson1_serializeOpUpdateDocumentationPa
 func (m *awsRestjson1_serializeOpUpdateDocumentationPart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9921,7 +9594,6 @@ func (m *awsRestjson1_serializeOpUpdateDocumentationPa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDocumentationPartInput(v *UpdateDocumentationPartInput, encoder *httpbinding.Encoder) error {
@@ -9974,10 +9646,8 @@ func (m *awsRestjson1_serializeOpUpdateDocumentationVe
 func (m *awsRestjson1_serializeOpUpdateDocumentationVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10026,7 +9696,6 @@ func (m *awsRestjson1_serializeOpUpdateDocumentationVe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDocumentationVersionInput(v *UpdateDocumentationVersionInput, encoder *httpbinding.Encoder) error {
@@ -10079,10 +9748,8 @@ func (m *awsRestjson1_serializeOpUpdateDomainName) Han
 func (m *awsRestjson1_serializeOpUpdateDomainName) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10131,7 +9798,6 @@ func (m *awsRestjson1_serializeOpUpdateDomainName) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDomainNameInput(v *UpdateDomainNameInput, encoder *httpbinding.Encoder) error {
@@ -10179,10 +9845,8 @@ func (m *awsRestjson1_serializeOpUpdateGatewayResponse
 func (m *awsRestjson1_serializeOpUpdateGatewayResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10231,7 +9895,6 @@ func (m *awsRestjson1_serializeOpUpdateGatewayResponse
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGatewayResponseInput(v *UpdateGatewayResponseInput, encoder *httpbinding.Encoder) error {
@@ -10284,10 +9947,8 @@ func (m *awsRestjson1_serializeOpUpdateIntegration) Ha
 func (m *awsRestjson1_serializeOpUpdateIntegration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10336,7 +9997,6 @@ func (m *awsRestjson1_serializeOpUpdateIntegration) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIntegrationInput(v *UpdateIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -10398,10 +10058,8 @@ func (m *awsRestjson1_serializeOpUpdateIntegrationResp
 func (m *awsRestjson1_serializeOpUpdateIntegrationResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10450,7 +10108,6 @@ func (m *awsRestjson1_serializeOpUpdateIntegrationResp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIntegrationResponseInput(v *UpdateIntegrationResponseInput, encoder *httpbinding.Encoder) error {
@@ -10521,10 +10178,8 @@ func (m *awsRestjson1_serializeOpUpdateMethod) HandleS
 func (m *awsRestjson1_serializeOpUpdateMethod) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10573,7 +10228,6 @@ func (m *awsRestjson1_serializeOpUpdateMethod) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMethodInput(v *UpdateMethodInput, encoder *httpbinding.Encoder) error {
@@ -10635,10 +10289,8 @@ func (m *awsRestjson1_serializeOpUpdateMethodResponse)
 func (m *awsRestjson1_serializeOpUpdateMethodResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10687,7 +10339,6 @@ func (m *awsRestjson1_serializeOpUpdateMethodResponse)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMethodResponseInput(v *UpdateMethodResponseInput, encoder *httpbinding.Encoder) error {
@@ -10758,10 +10409,8 @@ func (m *awsRestjson1_serializeOpUpdateModel) HandleSe
 func (m *awsRestjson1_serializeOpUpdateModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10810,7 +10459,6 @@ func (m *awsRestjson1_serializeOpUpdateModel) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateModelInput(v *UpdateModelInput, encoder *httpbinding.Encoder) error {
@@ -10863,10 +10511,8 @@ func (m *awsRestjson1_serializeOpUpdateRequestValidato
 func (m *awsRestjson1_serializeOpUpdateRequestValidator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10915,7 +10561,6 @@ func (m *awsRestjson1_serializeOpUpdateRequestValidato
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRequestValidatorInput(v *UpdateRequestValidatorInput, encoder *httpbinding.Encoder) error {
@@ -10968,10 +10613,8 @@ func (m *awsRestjson1_serializeOpUpdateResource) Handl
 func (m *awsRestjson1_serializeOpUpdateResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11020,7 +10663,6 @@ func (m *awsRestjson1_serializeOpUpdateResource) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceInput(v *UpdateResourceInput, encoder *httpbinding.Encoder) error {
@@ -11073,10 +10715,8 @@ func (m *awsRestjson1_serializeOpUpdateRestApi) Handle
 func (m *awsRestjson1_serializeOpUpdateRestApi) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11125,7 +10765,6 @@ func (m *awsRestjson1_serializeOpUpdateRestApi) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRestApiInput(v *UpdateRestApiInput, encoder *httpbinding.Encoder) error {
@@ -11169,10 +10808,8 @@ func (m *awsRestjson1_serializeOpUpdateStage) HandleSe
 func (m *awsRestjson1_serializeOpUpdateStage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11221,7 +10858,6 @@ func (m *awsRestjson1_serializeOpUpdateStage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateStageInput(v *UpdateStageInput, encoder *httpbinding.Encoder) error {
@@ -11274,10 +10910,8 @@ func (m *awsRestjson1_serializeOpUpdateUsage) HandleSe
 func (m *awsRestjson1_serializeOpUpdateUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11326,7 +10960,6 @@ func (m *awsRestjson1_serializeOpUpdateUsage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUsageInput(v *UpdateUsageInput, encoder *httpbinding.Encoder) error {
@@ -11379,10 +11012,8 @@ func (m *awsRestjson1_serializeOpUpdateUsagePlan) Hand
 func (m *awsRestjson1_serializeOpUpdateUsagePlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11431,7 +11062,6 @@ func (m *awsRestjson1_serializeOpUpdateUsagePlan) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUsagePlanInput(v *UpdateUsagePlanInput, encoder *httpbinding.Encoder) error {
@@ -11475,10 +11105,8 @@ func (m *awsRestjson1_serializeOpUpdateVpcLink) Handle
 func (m *awsRestjson1_serializeOpUpdateVpcLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11527,7 +11155,6 @@ func (m *awsRestjson1_serializeOpUpdateVpcLink) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVpcLinkInput(v *UpdateVpcLinkInput, encoder *httpbinding.Encoder) error {
