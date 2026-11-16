--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpCreateAgentRuntime) H
 func (m *awsRestjson1_serializeOpCreateAgentRuntime) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -73,7 +70,6 @@ func (m *awsRestjson1_serializeOpCreateAgentRuntime) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAgentRuntimeInput(v *CreateAgentRuntimeInput, encoder *httpbinding.Encoder) error {
@@ -156,10 +152,8 @@ func (m *awsRestjson1_serializeOpCreateAgentRuntimeEnd
 func (m *awsRestjson1_serializeOpCreateAgentRuntimeEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -208,7 +202,6 @@ func (m *awsRestjson1_serializeOpCreateAgentRuntimeEnd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAgentRuntimeEndpointInput(v *CreateAgentRuntimeEndpointInput, encoder *httpbinding.Encoder) error {
@@ -265,10 +258,8 @@ func (m *awsRestjson1_serializeOpCreateApiKeyCredentia
 func (m *awsRestjson1_serializeOpCreateApiKeyCredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -313,7 +304,6 @@ func (m *awsRestjson1_serializeOpCreateApiKeyCredentia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateApiKeyCredentialProviderInput(v *CreateApiKeyCredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -351,10 +341,8 @@ func (m *awsRestjson1_serializeOpCreateBrowser) Handle
 func (m *awsRestjson1_serializeOpCreateBrowser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -399,7 +387,6 @@ func (m *awsRestjson1_serializeOpCreateBrowser) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBrowserInput(v *CreateBrowserInput, encoder *httpbinding.Encoder) error {
@@ -461,10 +448,8 @@ func (m *awsRestjson1_serializeOpCreateCodeInterpreter
 func (m *awsRestjson1_serializeOpCreateCodeInterpreter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -509,7 +494,6 @@ func (m *awsRestjson1_serializeOpCreateCodeInterpreter
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCodeInterpreterInput(v *CreateCodeInterpreterInput, encoder *httpbinding.Encoder) error {
@@ -564,10 +548,8 @@ func (m *awsRestjson1_serializeOpCreateGateway) Handle
 func (m *awsRestjson1_serializeOpCreateGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -612,7 +594,6 @@ func (m *awsRestjson1_serializeOpCreateGateway) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGatewayInput(v *CreateGatewayInput, encoder *httpbinding.Encoder) error {
@@ -694,10 +675,8 @@ func (m *awsRestjson1_serializeOpCreateGatewayTarget) 
 func (m *awsRestjson1_serializeOpCreateGatewayTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -746,7 +725,6 @@ func (m *awsRestjson1_serializeOpCreateGatewayTarget) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGatewayTargetInput(v *CreateGatewayTargetInput, encoder *httpbinding.Encoder) error {
@@ -812,10 +790,8 @@ func (m *awsRestjson1_serializeOpCreateMemory) HandleS
 func (m *awsRestjson1_serializeOpCreateMemory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -860,7 +836,6 @@ func (m *awsRestjson1_serializeOpCreateMemory) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMemoryInput(v *CreateMemoryInput, encoder *httpbinding.Encoder) error {
@@ -925,10 +900,8 @@ func (m *awsRestjson1_serializeOpCreateOauth2Credentia
 func (m *awsRestjson1_serializeOpCreateOauth2CredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -973,7 +946,6 @@ func (m *awsRestjson1_serializeOpCreateOauth2Credentia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOauth2CredentialProviderInput(v *CreateOauth2CredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -1018,10 +990,8 @@ func (m *awsRestjson1_serializeOpCreateWorkloadIdentit
 func (m *awsRestjson1_serializeOpCreateWorkloadIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1066,7 +1036,6 @@ func (m *awsRestjson1_serializeOpCreateWorkloadIdentit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWorkloadIdentityInput(v *CreateWorkloadIdentityInput, encoder *httpbinding.Encoder) error {
@@ -1106,10 +1075,8 @@ func (m *awsRestjson1_serializeOpDeleteAgentRuntime) H
 func (m *awsRestjson1_serializeOpDeleteAgentRuntime) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1147,7 +1114,6 @@ func (m *awsRestjson1_serializeOpDeleteAgentRuntime) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAgentRuntimeInput(v *DeleteAgentRuntimeInput, encoder *httpbinding.Encoder) error {
@@ -1177,10 +1143,8 @@ func (m *awsRestjson1_serializeOpDeleteAgentRuntimeEnd
 func (m *awsRestjson1_serializeOpDeleteAgentRuntimeEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1218,7 +1182,6 @@ func (m *awsRestjson1_serializeOpDeleteAgentRuntimeEnd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAgentRuntimeEndpointInput(v *DeleteAgentRuntimeEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1261,10 +1224,8 @@ func (m *awsRestjson1_serializeOpDeleteApiKeyCredentia
 func (m *awsRestjson1_serializeOpDeleteApiKeyCredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1309,7 +1270,6 @@ func (m *awsRestjson1_serializeOpDeleteApiKeyCredentia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteApiKeyCredentialProviderInput(v *DeleteApiKeyCredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -1342,10 +1302,8 @@ func (m *awsRestjson1_serializeOpDeleteBrowser) Handle
 func (m *awsRestjson1_serializeOpDeleteBrowser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1383,7 +1341,6 @@ func (m *awsRestjson1_serializeOpDeleteBrowser) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBrowserInput(v *DeleteBrowserInput, encoder *httpbinding.Encoder) error {
@@ -1417,10 +1374,8 @@ func (m *awsRestjson1_serializeOpDeleteCodeInterpreter
 func (m *awsRestjson1_serializeOpDeleteCodeInterpreter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1458,7 +1413,6 @@ func (m *awsRestjson1_serializeOpDeleteCodeInterpreter
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCodeInterpreterInput(v *DeleteCodeInterpreterInput, encoder *httpbinding.Encoder) error {
@@ -1492,10 +1446,8 @@ func (m *awsRestjson1_serializeOpDeleteGateway) Handle
 func (m *awsRestjson1_serializeOpDeleteGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1533,7 +1485,6 @@ func (m *awsRestjson1_serializeOpDeleteGateway) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGatewayInput(v *DeleteGatewayInput, encoder *httpbinding.Encoder) error {
@@ -1563,10 +1514,8 @@ func (m *awsRestjson1_serializeOpDeleteGatewayTarget) 
 func (m *awsRestjson1_serializeOpDeleteGatewayTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1604,7 +1553,6 @@ func (m *awsRestjson1_serializeOpDeleteGatewayTarget) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGatewayTargetInput(v *DeleteGatewayTargetInput, encoder *httpbinding.Encoder) error {
@@ -1643,10 +1591,8 @@ func (m *awsRestjson1_serializeOpDeleteMemory) HandleS
 func (m *awsRestjson1_serializeOpDeleteMemory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1684,7 +1630,6 @@ func (m *awsRestjson1_serializeOpDeleteMemory) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMemoryInput(v *DeleteMemoryInput, encoder *httpbinding.Encoder) error {
@@ -1718,10 +1663,8 @@ func (m *awsRestjson1_serializeOpDeleteOauth2Credentia
 func (m *awsRestjson1_serializeOpDeleteOauth2CredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1766,7 +1709,6 @@ func (m *awsRestjson1_serializeOpDeleteOauth2Credentia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteOauth2CredentialProviderInput(v *DeleteOauth2CredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -1799,10 +1741,8 @@ func (m *awsRestjson1_serializeOpDeleteWorkloadIdentit
 func (m *awsRestjson1_serializeOpDeleteWorkloadIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1847,7 +1787,6 @@ func (m *awsRestjson1_serializeOpDeleteWorkloadIdentit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWorkloadIdentityInput(v *DeleteWorkloadIdentityInput, encoder *httpbinding.Encoder) error {
@@ -1880,10 +1819,8 @@ func (m *awsRestjson1_serializeOpGetAgentRuntime) Hand
 func (m *awsRestjson1_serializeOpGetAgentRuntime) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1921,7 +1858,6 @@ func (m *awsRestjson1_serializeOpGetAgentRuntime) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentRuntimeInput(v *GetAgentRuntimeInput, encoder *httpbinding.Encoder) error {
@@ -1955,10 +1891,8 @@ func (m *awsRestjson1_serializeOpGetAgentRuntimeEndpoi
 func (m *awsRestjson1_serializeOpGetAgentRuntimeEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1996,7 +1930,6 @@ func (m *awsRestjson1_serializeOpGetAgentRuntimeEndpoi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentRuntimeEndpointInput(v *GetAgentRuntimeEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2035,10 +1968,8 @@ func (m *awsRestjson1_serializeOpGetApiKeyCredentialPr
 func (m *awsRestjson1_serializeOpGetApiKeyCredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2083,7 +2014,6 @@ func (m *awsRestjson1_serializeOpGetApiKeyCredentialPr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApiKeyCredentialProviderInput(v *GetApiKeyCredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -2116,10 +2046,8 @@ func (m *awsRestjson1_serializeOpGetBrowser) HandleSer
 func (m *awsRestjson1_serializeOpGetBrowser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2157,7 +2085,6 @@ func (m *awsRestjson1_serializeOpGetBrowser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBrowserInput(v *GetBrowserInput, encoder *httpbinding.Encoder) error {
@@ -2187,10 +2114,8 @@ func (m *awsRestjson1_serializeOpGetCodeInterpreter) H
 func (m *awsRestjson1_serializeOpGetCodeInterpreter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2228,7 +2153,6 @@ func (m *awsRestjson1_serializeOpGetCodeInterpreter) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCodeInterpreterInput(v *GetCodeInterpreterInput, encoder *httpbinding.Encoder) error {
@@ -2258,10 +2182,8 @@ func (m *awsRestjson1_serializeOpGetGateway) HandleSer
 func (m *awsRestjson1_serializeOpGetGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2299,7 +2221,6 @@ func (m *awsRestjson1_serializeOpGetGateway) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGatewayInput(v *GetGatewayInput, encoder *httpbinding.Encoder) error {
@@ -2329,10 +2250,8 @@ func (m *awsRestjson1_serializeOpGetGatewayTarget) Han
 func (m *awsRestjson1_serializeOpGetGatewayTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2370,7 +2289,6 @@ func (m *awsRestjson1_serializeOpGetGatewayTarget) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGatewayTargetInput(v *GetGatewayTargetInput, encoder *httpbinding.Encoder) error {
@@ -2409,10 +2327,8 @@ func (m *awsRestjson1_serializeOpGetMemory) HandleSeri
 func (m *awsRestjson1_serializeOpGetMemory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2450,7 +2366,6 @@ func (m *awsRestjson1_serializeOpGetMemory) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMemoryInput(v *GetMemoryInput, encoder *httpbinding.Encoder) error {
@@ -2480,10 +2395,8 @@ func (m *awsRestjson1_serializeOpGetOauth2CredentialPr
 func (m *awsRestjson1_serializeOpGetOauth2CredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2528,7 +2441,6 @@ func (m *awsRestjson1_serializeOpGetOauth2CredentialPr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOauth2CredentialProviderInput(v *GetOauth2CredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -2561,10 +2473,8 @@ func (m *awsRestjson1_serializeOpGetTokenVault) Handle
 func (m *awsRestjson1_serializeOpGetTokenVault) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2609,7 +2519,6 @@ func (m *awsRestjson1_serializeOpGetTokenVault) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTokenVaultInput(v *GetTokenVaultInput, encoder *httpbinding.Encoder) error {
@@ -2642,10 +2551,8 @@ func (m *awsRestjson1_serializeOpGetWorkloadIdentity) 
 func (m *awsRestjson1_serializeOpGetWorkloadIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2690,7 +2597,6 @@ func (m *awsRestjson1_serializeOpGetWorkloadIdentity) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkloadIdentityInput(v *GetWorkloadIdentityInput, encoder *httpbinding.Encoder) error {
@@ -2723,10 +2629,8 @@ func (m *awsRestjson1_serializeOpListAgentRuntimeEndpo
 func (m *awsRestjson1_serializeOpListAgentRuntimeEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2764,7 +2668,6 @@ func (m *awsRestjson1_serializeOpListAgentRuntimeEndpo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentRuntimeEndpointsInput(v *ListAgentRuntimeEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -2802,10 +2705,8 @@ func (m *awsRestjson1_serializeOpListAgentRuntimes) Ha
 func (m *awsRestjson1_serializeOpListAgentRuntimes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2843,7 +2744,6 @@ func (m *awsRestjson1_serializeOpListAgentRuntimes) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentRuntimesInput(v *ListAgentRuntimesInput, encoder *httpbinding.Encoder) error {
@@ -2872,10 +2772,8 @@ func (m *awsRestjson1_serializeOpListAgentRuntimeVersi
 func (m *awsRestjson1_serializeOpListAgentRuntimeVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2913,7 +2811,6 @@ func (m *awsRestjson1_serializeOpListAgentRuntimeVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentRuntimeVersionsInput(v *ListAgentRuntimeVersionsInput, encoder *httpbinding.Encoder) error {
@@ -2951,10 +2848,8 @@ func (m *awsRestjson1_serializeOpListApiKeyCredentialP
 func (m *awsRestjson1_serializeOpListApiKeyCredentialProviders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2999,7 +2894,6 @@ func (m *awsRestjson1_serializeOpListApiKeyCredentialP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListApiKeyCredentialProvidersInput(v *ListApiKeyCredentialProvidersInput, encoder *httpbinding.Encoder) error {
@@ -3037,10 +2931,8 @@ func (m *awsRestjson1_serializeOpListBrowsers) HandleS
 func (m *awsRestjson1_serializeOpListBrowsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3078,7 +2970,6 @@ func (m *awsRestjson1_serializeOpListBrowsers) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBrowsersInput(v *ListBrowsersInput, encoder *httpbinding.Encoder) error {
@@ -3111,10 +3002,8 @@ func (m *awsRestjson1_serializeOpListCodeInterpreters)
 func (m *awsRestjson1_serializeOpListCodeInterpreters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3152,7 +3041,6 @@ func (m *awsRestjson1_serializeOpListCodeInterpreters)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCodeInterpretersInput(v *ListCodeInterpretersInput, encoder *httpbinding.Encoder) error {
@@ -3185,10 +3073,8 @@ func (m *awsRestjson1_serializeOpListGateways) HandleS
 func (m *awsRestjson1_serializeOpListGateways) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3226,7 +3112,6 @@ func (m *awsRestjson1_serializeOpListGateways) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGatewaysInput(v *ListGatewaysInput, encoder *httpbinding.Encoder) error {
@@ -3255,10 +3140,8 @@ func (m *awsRestjson1_serializeOpListGatewayTargets) H
 func (m *awsRestjson1_serializeOpListGatewayTargets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3296,7 +3179,6 @@ func (m *awsRestjson1_serializeOpListGatewayTargets) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGatewayTargetsInput(v *ListGatewayTargetsInput, encoder *httpbinding.Encoder) error {
@@ -3334,10 +3216,8 @@ func (m *awsRestjson1_serializeOpListMemories) HandleS
 func (m *awsRestjson1_serializeOpListMemories) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3382,7 +3262,6 @@ func (m *awsRestjson1_serializeOpListMemories) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMemoriesInput(v *ListMemoriesInput, encoder *httpbinding.Encoder) error {
@@ -3420,10 +3299,8 @@ func (m *awsRestjson1_serializeOpListOauth2CredentialP
 func (m *awsRestjson1_serializeOpListOauth2CredentialProviders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3468,7 +3345,6 @@ func (m *awsRestjson1_serializeOpListOauth2CredentialP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOauth2CredentialProvidersInput(v *ListOauth2CredentialProvidersInput, encoder *httpbinding.Encoder) error {
@@ -3506,10 +3382,8 @@ func (m *awsRestjson1_serializeOpListWorkloadIdentitie
 func (m *awsRestjson1_serializeOpListWorkloadIdentities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3554,7 +3428,6 @@ func (m *awsRestjson1_serializeOpListWorkloadIdentitie
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkloadIdentitiesInput(v *ListWorkloadIdentitiesInput, encoder *httpbinding.Encoder) error {
@@ -3592,10 +3465,8 @@ func (m *awsRestjson1_serializeOpSetTokenVaultCMK) Han
 func (m *awsRestjson1_serializeOpSetTokenVaultCMK) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3640,7 +3511,6 @@ func (m *awsRestjson1_serializeOpSetTokenVaultCMK) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetTokenVaultCMKInput(v *SetTokenVaultCMKInput, encoder *httpbinding.Encoder) error {
@@ -3680,10 +3550,8 @@ func (m *awsRestjson1_serializeOpUpdateAgentRuntime) H
 func (m *awsRestjson1_serializeOpUpdateAgentRuntime) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3732,7 +3600,6 @@ func (m *awsRestjson1_serializeOpUpdateAgentRuntime) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentRuntimeInput(v *UpdateAgentRuntimeInput, encoder *httpbinding.Encoder) error {
@@ -3819,10 +3686,8 @@ func (m *awsRestjson1_serializeOpUpdateAgentRuntimeEnd
 func (m *awsRestjson1_serializeOpUpdateAgentRuntimeEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3871,7 +3736,6 @@ func (m *awsRestjson1_serializeOpUpdateAgentRuntimeEnd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentRuntimeEndpointInput(v *UpdateAgentRuntimeEndpointInput, encoder *httpbinding.Encoder) error {
@@ -3932,10 +3796,8 @@ func (m *awsRestjson1_serializeOpUpdateApiKeyCredentia
 func (m *awsRestjson1_serializeOpUpdateApiKeyCredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3980,7 +3842,6 @@ func (m *awsRestjson1_serializeOpUpdateApiKeyCredentia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApiKeyCredentialProviderInput(v *UpdateApiKeyCredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -4018,10 +3879,8 @@ func (m *awsRestjson1_serializeOpUpdateGateway) Handle
 func (m *awsRestjson1_serializeOpUpdateGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4070,7 +3929,6 @@ func (m *awsRestjson1_serializeOpUpdateGateway) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGatewayInput(v *UpdateGatewayInput, encoder *httpbinding.Encoder) error {
@@ -4156,10 +4014,8 @@ func (m *awsRestjson1_serializeOpUpdateGatewayTarget) 
 func (m *awsRestjson1_serializeOpUpdateGatewayTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4208,7 +4064,6 @@ func (m *awsRestjson1_serializeOpUpdateGatewayTarget) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGatewayTargetInput(v *UpdateGatewayTargetInput, encoder *httpbinding.Encoder) error {
@@ -4278,10 +4133,8 @@ func (m *awsRestjson1_serializeOpUpdateMemory) HandleS
 func (m *awsRestjson1_serializeOpUpdateMemory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4330,7 +4183,6 @@ func (m *awsRestjson1_serializeOpUpdateMemory) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMemoryInput(v *UpdateMemoryInput, encoder *httpbinding.Encoder) error {
@@ -4394,10 +4246,8 @@ func (m *awsRestjson1_serializeOpUpdateOauth2Credentia
 func (m *awsRestjson1_serializeOpUpdateOauth2CredentialProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4442,7 +4292,6 @@ func (m *awsRestjson1_serializeOpUpdateOauth2Credentia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOauth2CredentialProviderInput(v *UpdateOauth2CredentialProviderInput, encoder *httpbinding.Encoder) error {
@@ -4487,10 +4336,8 @@ func (m *awsRestjson1_serializeOpUpdateWorkloadIdentit
 func (m *awsRestjson1_serializeOpUpdateWorkloadIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4535,7 +4382,6 @@ func (m *awsRestjson1_serializeOpUpdateWorkloadIdentit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWorkloadIdentityInput(v *UpdateWorkloadIdentityInput, encoder *httpbinding.Encoder) error {
