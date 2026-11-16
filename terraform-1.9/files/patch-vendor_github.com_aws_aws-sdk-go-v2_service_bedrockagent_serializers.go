--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagent/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagent/serializers.go
@@ -13,7 +13,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -28,10 +27,8 @@ func (m *awsRestjson1_serializeOpAssociateAgentCollabo
 func (m *awsRestjson1_serializeOpAssociateAgentCollaborator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -80,7 +77,6 @@ func (m *awsRestjson1_serializeOpAssociateAgentCollabo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateAgentCollaboratorInput(v *AssociateAgentCollaboratorInput, encoder *httpbinding.Encoder) error {
@@ -153,10 +149,8 @@ func (m *awsRestjson1_serializeOpAssociateAgentKnowled
 func (m *awsRestjson1_serializeOpAssociateAgentKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -205,7 +199,6 @@ func (m *awsRestjson1_serializeOpAssociateAgentKnowled
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateAgentKnowledgeBaseInput(v *AssociateAgentKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -266,10 +259,8 @@ func (m *awsRestjson1_serializeOpCreateAgent) HandleSe
 func (m *awsRestjson1_serializeOpCreateAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -314,7 +305,6 @@ func (m *awsRestjson1_serializeOpCreateAgent) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAgentInput(v *CreateAgentInput, encoder *httpbinding.Encoder) error {
@@ -427,10 +417,8 @@ func (m *awsRestjson1_serializeOpCreateAgentActionGrou
 func (m *awsRestjson1_serializeOpCreateAgentActionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -479,7 +467,6 @@ func (m *awsRestjson1_serializeOpCreateAgentActionGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAgentActionGroupInput(v *CreateAgentActionGroupInput, encoder *httpbinding.Encoder) error {
@@ -578,10 +565,8 @@ func (m *awsRestjson1_serializeOpCreateAgentAlias) Han
 func (m *awsRestjson1_serializeOpCreateAgentAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -630,7 +615,6 @@ func (m *awsRestjson1_serializeOpCreateAgentAlias) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAgentAliasInput(v *CreateAgentAliasInput, encoder *httpbinding.Encoder) error {
@@ -696,10 +680,8 @@ func (m *awsRestjson1_serializeOpCreateDataSource) Han
 func (m *awsRestjson1_serializeOpCreateDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -748,7 +730,6 @@ func (m *awsRestjson1_serializeOpCreateDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataSourceInput(v *CreateDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -826,10 +807,8 @@ func (m *awsRestjson1_serializeOpCreateFlow) HandleSer
 func (m *awsRestjson1_serializeOpCreateFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -874,7 +853,6 @@ func (m *awsRestjson1_serializeOpCreateFlow) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFlowInput(v *CreateFlowInput, encoder *httpbinding.Encoder) error {
@@ -941,10 +919,8 @@ func (m *awsRestjson1_serializeOpCreateFlowAlias) Hand
 func (m *awsRestjson1_serializeOpCreateFlowAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -993,7 +969,6 @@ func (m *awsRestjson1_serializeOpCreateFlowAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFlowAliasInput(v *CreateFlowAliasInput, encoder *httpbinding.Encoder) error {
@@ -1066,10 +1041,8 @@ func (m *awsRestjson1_serializeOpCreateFlowVersion) Ha
 func (m *awsRestjson1_serializeOpCreateFlowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1118,7 +1091,6 @@ func (m *awsRestjson1_serializeOpCreateFlowVersion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFlowVersionInput(v *CreateFlowVersionInput, encoder *httpbinding.Encoder) error {
@@ -1165,10 +1137,8 @@ func (m *awsRestjson1_serializeOpCreateKnowledgeBase) 
 func (m *awsRestjson1_serializeOpCreateKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1213,7 +1183,6 @@ func (m *awsRestjson1_serializeOpCreateKnowledgeBase) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateKnowledgeBaseInput(v *CreateKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -1282,10 +1251,8 @@ func (m *awsRestjson1_serializeOpCreatePrompt) HandleS
 func (m *awsRestjson1_serializeOpCreatePrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1330,7 +1297,6 @@ func (m *awsRestjson1_serializeOpCreatePrompt) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePromptInput(v *CreatePromptInput, encoder *httpbinding.Encoder) error {
@@ -1397,10 +1363,8 @@ func (m *awsRestjson1_serializeOpCreatePromptVersion) 
 func (m *awsRestjson1_serializeOpCreatePromptVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1449,7 +1413,6 @@ func (m *awsRestjson1_serializeOpCreatePromptVersion) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePromptVersionInput(v *CreatePromptVersionInput, encoder *httpbinding.Encoder) error {
@@ -1503,10 +1466,8 @@ func (m *awsRestjson1_serializeOpDeleteAgent) HandleSe
 func (m *awsRestjson1_serializeOpDeleteAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1544,7 +1505,6 @@ func (m *awsRestjson1_serializeOpDeleteAgent) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAgentInput(v *DeleteAgentInput, encoder *httpbinding.Encoder) error {
@@ -1578,10 +1538,8 @@ func (m *awsRestjson1_serializeOpDeleteAgentActionGrou
 func (m *awsRestjson1_serializeOpDeleteAgentActionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1619,7 +1577,6 @@ func (m *awsRestjson1_serializeOpDeleteAgentActionGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAgentActionGroupInput(v *DeleteAgentActionGroupInput, encoder *httpbinding.Encoder) error {
@@ -1671,10 +1628,8 @@ func (m *awsRestjson1_serializeOpDeleteAgentAlias) Han
 func (m *awsRestjson1_serializeOpDeleteAgentAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1712,7 +1667,6 @@ func (m *awsRestjson1_serializeOpDeleteAgentAlias) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAgentAliasInput(v *DeleteAgentAliasInput, encoder *httpbinding.Encoder) error {
@@ -1751,10 +1705,8 @@ func (m *awsRestjson1_serializeOpDeleteAgentVersion) H
 func (m *awsRestjson1_serializeOpDeleteAgentVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1792,7 +1744,6 @@ func (m *awsRestjson1_serializeOpDeleteAgentVersion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAgentVersionInput(v *DeleteAgentVersionInput, encoder *httpbinding.Encoder) error {
@@ -1835,10 +1786,8 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 func (m *awsRestjson1_serializeOpDeleteDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1876,7 +1825,6 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataSourceInput(v *DeleteDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -1915,10 +1863,8 @@ func (m *awsRestjson1_serializeOpDeleteFlow) HandleSer
 func (m *awsRestjson1_serializeOpDeleteFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1956,7 +1902,6 @@ func (m *awsRestjson1_serializeOpDeleteFlow) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFlowInput(v *DeleteFlowInput, encoder *httpbinding.Encoder) error {
@@ -1990,10 +1935,8 @@ func (m *awsRestjson1_serializeOpDeleteFlowAlias) Hand
 func (m *awsRestjson1_serializeOpDeleteFlowAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2031,7 +1974,6 @@ func (m *awsRestjson1_serializeOpDeleteFlowAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFlowAliasInput(v *DeleteFlowAliasInput, encoder *httpbinding.Encoder) error {
@@ -2070,10 +2012,8 @@ func (m *awsRestjson1_serializeOpDeleteFlowVersion) Ha
 func (m *awsRestjson1_serializeOpDeleteFlowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2111,7 +2051,6 @@ func (m *awsRestjson1_serializeOpDeleteFlowVersion) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFlowVersionInput(v *DeleteFlowVersionInput, encoder *httpbinding.Encoder) error {
@@ -2154,10 +2093,8 @@ func (m *awsRestjson1_serializeOpDeleteKnowledgeBase) 
 func (m *awsRestjson1_serializeOpDeleteKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2195,7 +2132,6 @@ func (m *awsRestjson1_serializeOpDeleteKnowledgeBase) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteKnowledgeBaseInput(v *DeleteKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -2225,10 +2161,8 @@ func (m *awsRestjson1_serializeOpDeleteKnowledgeBaseDo
 func (m *awsRestjson1_serializeOpDeleteKnowledgeBaseDocuments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2277,7 +2211,6 @@ func (m *awsRestjson1_serializeOpDeleteKnowledgeBaseDo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteKnowledgeBaseDocumentsInput(v *DeleteKnowledgeBaseDocumentsInput, encoder *httpbinding.Encoder) error {
@@ -2335,10 +2268,8 @@ func (m *awsRestjson1_serializeOpDeletePrompt) HandleS
 func (m *awsRestjson1_serializeOpDeletePrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2376,7 +2307,6 @@ func (m *awsRestjson1_serializeOpDeletePrompt) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePromptInput(v *DeletePromptInput, encoder *httpbinding.Encoder) error {
@@ -2410,10 +2340,8 @@ func (m *awsRestjson1_serializeOpDisassociateAgentColl
 func (m *awsRestjson1_serializeOpDisassociateAgentCollaborator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2451,7 +2379,6 @@ func (m *awsRestjson1_serializeOpDisassociateAgentColl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateAgentCollaboratorInput(v *DisassociateAgentCollaboratorInput, encoder *httpbinding.Encoder) error {
@@ -2499,10 +2426,8 @@ func (m *awsRestjson1_serializeOpDisassociateAgentKnow
 func (m *awsRestjson1_serializeOpDisassociateAgentKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2540,7 +2465,6 @@ func (m *awsRestjson1_serializeOpDisassociateAgentKnow
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateAgentKnowledgeBaseInput(v *DisassociateAgentKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -2588,10 +2512,8 @@ func (m *awsRestjson1_serializeOpGetAgent) HandleSeria
 func (m *awsRestjson1_serializeOpGetAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2629,7 +2551,6 @@ func (m *awsRestjson1_serializeOpGetAgent) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentInput(v *GetAgentInput, encoder *httpbinding.Encoder) error {
@@ -2659,10 +2580,8 @@ func (m *awsRestjson1_serializeOpGetAgentActionGroup) 
 func (m *awsRestjson1_serializeOpGetAgentActionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2700,7 +2619,6 @@ func (m *awsRestjson1_serializeOpGetAgentActionGroup) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentActionGroupInput(v *GetAgentActionGroupInput, encoder *httpbinding.Encoder) error {
@@ -2748,10 +2666,8 @@ func (m *awsRestjson1_serializeOpGetAgentAlias) Handle
 func (m *awsRestjson1_serializeOpGetAgentAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2789,7 +2705,6 @@ func (m *awsRestjson1_serializeOpGetAgentAlias) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentAliasInput(v *GetAgentAliasInput, encoder *httpbinding.Encoder) error {
@@ -2828,10 +2743,8 @@ func (m *awsRestjson1_serializeOpGetAgentCollaborator)
 func (m *awsRestjson1_serializeOpGetAgentCollaborator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2869,7 +2782,6 @@ func (m *awsRestjson1_serializeOpGetAgentCollaborator)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentCollaboratorInput(v *GetAgentCollaboratorInput, encoder *httpbinding.Encoder) error {
@@ -2917,10 +2829,8 @@ func (m *awsRestjson1_serializeOpGetAgentKnowledgeBase
 func (m *awsRestjson1_serializeOpGetAgentKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2958,7 +2868,6 @@ func (m *awsRestjson1_serializeOpGetAgentKnowledgeBase
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentKnowledgeBaseInput(v *GetAgentKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -3006,10 +2915,8 @@ func (m *awsRestjson1_serializeOpGetAgentVersion) Hand
 func (m *awsRestjson1_serializeOpGetAgentVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3047,7 +2954,6 @@ func (m *awsRestjson1_serializeOpGetAgentVersion) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAgentVersionInput(v *GetAgentVersionInput, encoder *httpbinding.Encoder) error {
@@ -3086,10 +2992,8 @@ func (m *awsRestjson1_serializeOpGetDataSource) Handle
 func (m *awsRestjson1_serializeOpGetDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3127,7 +3031,6 @@ func (m *awsRestjson1_serializeOpGetDataSource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataSourceInput(v *GetDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -3166,10 +3069,8 @@ func (m *awsRestjson1_serializeOpGetFlow) HandleSerial
 func (m *awsRestjson1_serializeOpGetFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3207,7 +3108,6 @@ func (m *awsRestjson1_serializeOpGetFlow) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFlowInput(v *GetFlowInput, encoder *httpbinding.Encoder) error {
@@ -3237,10 +3137,8 @@ func (m *awsRestjson1_serializeOpGetFlowAlias) HandleS
 func (m *awsRestjson1_serializeOpGetFlowAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3278,7 +3176,6 @@ func (m *awsRestjson1_serializeOpGetFlowAlias) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFlowAliasInput(v *GetFlowAliasInput, encoder *httpbinding.Encoder) error {
@@ -3317,10 +3214,8 @@ func (m *awsRestjson1_serializeOpGetFlowVersion) Handl
 func (m *awsRestjson1_serializeOpGetFlowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3358,7 +3253,6 @@ func (m *awsRestjson1_serializeOpGetFlowVersion) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFlowVersionInput(v *GetFlowVersionInput, encoder *httpbinding.Encoder) error {
@@ -3397,10 +3291,8 @@ func (m *awsRestjson1_serializeOpGetIngestionJob) Hand
 func (m *awsRestjson1_serializeOpGetIngestionJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3438,7 +3330,6 @@ func (m *awsRestjson1_serializeOpGetIngestionJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIngestionJobInput(v *GetIngestionJobInput, encoder *httpbinding.Encoder) error {
@@ -3486,10 +3377,8 @@ func (m *awsRestjson1_serializeOpGetKnowledgeBase) Han
 func (m *awsRestjson1_serializeOpGetKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3527,7 +3416,6 @@ func (m *awsRestjson1_serializeOpGetKnowledgeBase) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetKnowledgeBaseInput(v *GetKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -3557,10 +3445,8 @@ func (m *awsRestjson1_serializeOpGetKnowledgeBaseDocum
 func (m *awsRestjson1_serializeOpGetKnowledgeBaseDocuments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3609,7 +3495,6 @@ func (m *awsRestjson1_serializeOpGetKnowledgeBaseDocum
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetKnowledgeBaseDocumentsInput(v *GetKnowledgeBaseDocumentsInput, encoder *httpbinding.Encoder) error {
@@ -3662,10 +3547,8 @@ func (m *awsRestjson1_serializeOpGetPrompt) HandleSeri
 func (m *awsRestjson1_serializeOpGetPrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3703,7 +3586,6 @@ func (m *awsRestjson1_serializeOpGetPrompt) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPromptInput(v *GetPromptInput, encoder *httpbinding.Encoder) error {
@@ -3737,10 +3619,8 @@ func (m *awsRestjson1_serializeOpIngestKnowledgeBaseDo
 func (m *awsRestjson1_serializeOpIngestKnowledgeBaseDocuments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3789,7 +3669,6 @@ func (m *awsRestjson1_serializeOpIngestKnowledgeBaseDo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsIngestKnowledgeBaseDocumentsInput(v *IngestKnowledgeBaseDocumentsInput, encoder *httpbinding.Encoder) error {
@@ -3847,10 +3726,8 @@ func (m *awsRestjson1_serializeOpListAgentActionGroups
 func (m *awsRestjson1_serializeOpListAgentActionGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3899,7 +3776,6 @@ func (m *awsRestjson1_serializeOpListAgentActionGroups
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentActionGroupsInput(v *ListAgentActionGroupsInput, encoder *httpbinding.Encoder) error {
@@ -3955,10 +3831,8 @@ func (m *awsRestjson1_serializeOpListAgentAliases) Han
 func (m *awsRestjson1_serializeOpListAgentAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4007,7 +3881,6 @@ func (m *awsRestjson1_serializeOpListAgentAliases) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentAliasesInput(v *ListAgentAliasesInput, encoder *httpbinding.Encoder) error {
@@ -4054,10 +3927,8 @@ func (m *awsRestjson1_serializeOpListAgentCollaborator
 func (m *awsRestjson1_serializeOpListAgentCollaborators) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4106,7 +3977,6 @@ func (m *awsRestjson1_serializeOpListAgentCollaborator
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentCollaboratorsInput(v *ListAgentCollaboratorsInput, encoder *httpbinding.Encoder) error {
@@ -4162,10 +4032,8 @@ func (m *awsRestjson1_serializeOpListAgentKnowledgeBas
 func (m *awsRestjson1_serializeOpListAgentKnowledgeBases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4214,7 +4082,6 @@ func (m *awsRestjson1_serializeOpListAgentKnowledgeBas
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentKnowledgeBasesInput(v *ListAgentKnowledgeBasesInput, encoder *httpbinding.Encoder) error {
@@ -4270,10 +4137,8 @@ func (m *awsRestjson1_serializeOpListAgents) HandleSer
 func (m *awsRestjson1_serializeOpListAgents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4318,7 +4183,6 @@ func (m *awsRestjson1_serializeOpListAgents) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentsInput(v *ListAgentsInput, encoder *httpbinding.Encoder) error {
@@ -4356,10 +4220,8 @@ func (m *awsRestjson1_serializeOpListAgentVersions) Ha
 func (m *awsRestjson1_serializeOpListAgentVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4408,7 +4270,6 @@ func (m *awsRestjson1_serializeOpListAgentVersions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAgentVersionsInput(v *ListAgentVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4455,10 +4316,8 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 func (m *awsRestjson1_serializeOpListDataSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4507,7 +4366,6 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataSourcesInput(v *ListDataSourcesInput, encoder *httpbinding.Encoder) error {
@@ -4554,10 +4412,8 @@ func (m *awsRestjson1_serializeOpListFlowAliases) Hand
 func (m *awsRestjson1_serializeOpListFlowAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4595,7 +4451,6 @@ func (m *awsRestjson1_serializeOpListFlowAliases) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFlowAliasesInput(v *ListFlowAliasesInput, encoder *httpbinding.Encoder) error {
@@ -4633,10 +4488,8 @@ func (m *awsRestjson1_serializeOpListFlows) HandleSeri
 func (m *awsRestjson1_serializeOpListFlows) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4674,7 +4527,6 @@ func (m *awsRestjson1_serializeOpListFlows) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFlowsInput(v *ListFlowsInput, encoder *httpbinding.Encoder) error {
@@ -4703,10 +4555,8 @@ func (m *awsRestjson1_serializeOpListFlowVersions) Han
 func (m *awsRestjson1_serializeOpListFlowVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4744,7 +4594,6 @@ func (m *awsRestjson1_serializeOpListFlowVersions) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFlowVersionsInput(v *ListFlowVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4782,10 +4631,8 @@ func (m *awsRestjson1_serializeOpListIngestionJobs) Ha
 func (m *awsRestjson1_serializeOpListIngestionJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4834,7 +4681,6 @@ func (m *awsRestjson1_serializeOpListIngestionJobs) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIngestionJobsInput(v *ListIngestionJobsInput, encoder *httpbinding.Encoder) error {
@@ -4904,10 +4750,8 @@ func (m *awsRestjson1_serializeOpListKnowledgeBaseDocu
 func (m *awsRestjson1_serializeOpListKnowledgeBaseDocuments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4956,7 +4800,6 @@ func (m *awsRestjson1_serializeOpListKnowledgeBaseDocu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListKnowledgeBaseDocumentsInput(v *ListKnowledgeBaseDocumentsInput, encoder *httpbinding.Encoder) error {
@@ -5012,10 +4855,8 @@ func (m *awsRestjson1_serializeOpListKnowledgeBases) H
 func (m *awsRestjson1_serializeOpListKnowledgeBases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5060,7 +4901,6 @@ func (m *awsRestjson1_serializeOpListKnowledgeBases) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListKnowledgeBasesInput(v *ListKnowledgeBasesInput, encoder *httpbinding.Encoder) error {
@@ -5098,10 +4938,8 @@ func (m *awsRestjson1_serializeOpListPrompts) HandleSe
 func (m *awsRestjson1_serializeOpListPrompts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5139,7 +4977,6 @@ func (m *awsRestjson1_serializeOpListPrompts) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPromptsInput(v *ListPromptsInput, encoder *httpbinding.Encoder) error {
@@ -5172,10 +5009,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -5213,7 +5048,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -5243,10 +5077,8 @@ func (m *awsRestjson1_serializeOpPrepareAgent) HandleS
 func (m *awsRestjson1_serializeOpPrepareAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5284,7 +5116,6 @@ func (m *awsRestjson1_serializeOpPrepareAgent) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPrepareAgentInput(v *PrepareAgentInput, encoder *httpbinding.Encoder) error {
@@ -5314,10 +5145,8 @@ func (m *awsRestjson1_serializeOpPrepareFlow) HandleSe
 func (m *awsRestjson1_serializeOpPrepareFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5355,7 +5184,6 @@ func (m *awsRestjson1_serializeOpPrepareFlow) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPrepareFlowInput(v *PrepareFlowInput, encoder *httpbinding.Encoder) error {
@@ -5385,10 +5213,8 @@ func (m *awsRestjson1_serializeOpStartIngestionJob) Ha
 func (m *awsRestjson1_serializeOpStartIngestionJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5437,7 +5263,6 @@ func (m *awsRestjson1_serializeOpStartIngestionJob) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartIngestionJobInput(v *StartIngestionJobInput, encoder *httpbinding.Encoder) error {
@@ -5493,10 +5318,8 @@ func (m *awsRestjson1_serializeOpStopIngestionJob) Han
 func (m *awsRestjson1_serializeOpStopIngestionJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5534,7 +5357,6 @@ func (m *awsRestjson1_serializeOpStopIngestionJob) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopIngestionJobInput(v *StopIngestionJobInput, encoder *httpbinding.Encoder) error {
@@ -5582,10 +5404,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5634,7 +5454,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5678,10 +5497,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5719,7 +5536,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5755,10 +5571,8 @@ func (m *awsRestjson1_serializeOpUpdateAgent) HandleSe
 func (m *awsRestjson1_serializeOpUpdateAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5807,7 +5621,6 @@ func (m *awsRestjson1_serializeOpUpdateAgent) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentInput(v *UpdateAgentInput, encoder *httpbinding.Encoder) error {
@@ -5917,10 +5730,8 @@ func (m *awsRestjson1_serializeOpUpdateAgentActionGrou
 func (m *awsRestjson1_serializeOpUpdateAgentActionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5969,7 +5780,6 @@ func (m *awsRestjson1_serializeOpUpdateAgentActionGrou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentActionGroupInput(v *UpdateAgentActionGroupInput, encoder *httpbinding.Encoder) error {
@@ -6072,10 +5882,8 @@ func (m *awsRestjson1_serializeOpUpdateAgentAlias) Han
 func (m *awsRestjson1_serializeOpUpdateAgentAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6124,7 +5932,6 @@ func (m *awsRestjson1_serializeOpUpdateAgentAlias) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentAliasInput(v *UpdateAgentAliasInput, encoder *httpbinding.Encoder) error {
@@ -6192,10 +5999,8 @@ func (m *awsRestjson1_serializeOpUpdateAgentCollaborat
 func (m *awsRestjson1_serializeOpUpdateAgentCollaborator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6244,7 +6049,6 @@ func (m *awsRestjson1_serializeOpUpdateAgentCollaborat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentCollaboratorInput(v *UpdateAgentCollaboratorInput, encoder *httpbinding.Encoder) error {
@@ -6321,10 +6125,8 @@ func (m *awsRestjson1_serializeOpUpdateAgentKnowledgeB
 func (m *awsRestjson1_serializeOpUpdateAgentKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6373,7 +6175,6 @@ func (m *awsRestjson1_serializeOpUpdateAgentKnowledgeB
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAgentKnowledgeBaseInput(v *UpdateAgentKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -6438,10 +6239,8 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 func (m *awsRestjson1_serializeOpUpdateDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6490,7 +6289,6 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataSourceInput(v *UpdateDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -6572,10 +6370,8 @@ func (m *awsRestjson1_serializeOpUpdateFlow) HandleSer
 func (m *awsRestjson1_serializeOpUpdateFlow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6624,7 +6420,6 @@ func (m *awsRestjson1_serializeOpUpdateFlow) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFlowInput(v *UpdateFlowInput, encoder *httpbinding.Encoder) error {
@@ -6688,10 +6483,8 @@ func (m *awsRestjson1_serializeOpUpdateFlowAlias) Hand
 func (m *awsRestjson1_serializeOpUpdateFlowAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6740,7 +6533,6 @@ func (m *awsRestjson1_serializeOpUpdateFlowAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFlowAliasInput(v *UpdateFlowAliasInput, encoder *httpbinding.Encoder) error {
@@ -6810,10 +6602,8 @@ func (m *awsRestjson1_serializeOpUpdateKnowledgeBase) 
 func (m *awsRestjson1_serializeOpUpdateKnowledgeBase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6862,7 +6652,6 @@ func (m *awsRestjson1_serializeOpUpdateKnowledgeBase) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateKnowledgeBaseInput(v *UpdateKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -6928,10 +6717,8 @@ func (m *awsRestjson1_serializeOpUpdatePrompt) HandleS
 func (m *awsRestjson1_serializeOpUpdatePrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6980,7 +6767,6 @@ func (m *awsRestjson1_serializeOpUpdatePrompt) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePromptInput(v *UpdatePromptInput, encoder *httpbinding.Encoder) error {
@@ -7044,10 +6830,8 @@ func (m *awsRestjson1_serializeOpValidateFlowDefinitio
 func (m *awsRestjson1_serializeOpValidateFlowDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7092,7 +6876,6 @@ func (m *awsRestjson1_serializeOpValidateFlowDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsValidateFlowDefinitionInput(v *ValidateFlowDefinitionInput, encoder *httpbinding.Encoder) error {
