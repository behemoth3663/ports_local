--- vendor/github.com/aws/aws-sdk-go-v2/service/qconnect/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/qconnect/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpActivateMessageTempla
 func (m *awsRestjson1_serializeOpActivateMessageTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpActivateMessageTempla
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsActivateMessageTemplateInput(v *ActivateMessageTemplateInput, encoder *httpbinding.Encoder) error {
@@ -130,10 +126,8 @@ func (m *awsRestjson1_serializeOpCreateAIAgent) Handle
 func (m *awsRestjson1_serializeOpCreateAIAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -182,7 +176,6 @@ func (m *awsRestjson1_serializeOpCreateAIAgent) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAIAgentInput(v *CreateAIAgentInput, encoder *httpbinding.Encoder) error {
@@ -258,10 +251,8 @@ func (m *awsRestjson1_serializeOpCreateAIAgentVersion)
 func (m *awsRestjson1_serializeOpCreateAIAgentVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -310,7 +301,6 @@ func (m *awsRestjson1_serializeOpCreateAIAgentVersion)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAIAgentVersionInput(v *CreateAIAgentVersionInput, encoder *httpbinding.Encoder) error {
@@ -366,10 +356,8 @@ func (m *awsRestjson1_serializeOpCreateAIGuardrail) Ha
 func (m *awsRestjson1_serializeOpCreateAIGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -418,7 +406,6 @@ func (m *awsRestjson1_serializeOpCreateAIGuardrail) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAIGuardrailInput(v *CreateAIGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -527,10 +514,8 @@ func (m *awsRestjson1_serializeOpCreateAIGuardrailVers
 func (m *awsRestjson1_serializeOpCreateAIGuardrailVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -579,7 +564,6 @@ func (m *awsRestjson1_serializeOpCreateAIGuardrailVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAIGuardrailVersionInput(v *CreateAIGuardrailVersionInput, encoder *httpbinding.Encoder) error {
@@ -635,10 +619,8 @@ func (m *awsRestjson1_serializeOpCreateAIPrompt) Handl
 func (m *awsRestjson1_serializeOpCreateAIPrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -687,7 +669,6 @@ func (m *awsRestjson1_serializeOpCreateAIPrompt) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAIPromptInput(v *CreateAIPromptInput, encoder *httpbinding.Encoder) error {
@@ -778,10 +759,8 @@ func (m *awsRestjson1_serializeOpCreateAIPromptVersion
 func (m *awsRestjson1_serializeOpCreateAIPromptVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -830,7 +809,6 @@ func (m *awsRestjson1_serializeOpCreateAIPromptVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAIPromptVersionInput(v *CreateAIPromptVersionInput, encoder *httpbinding.Encoder) error {
@@ -886,10 +864,8 @@ func (m *awsRestjson1_serializeOpCreateAssistant) Hand
 func (m *awsRestjson1_serializeOpCreateAssistant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -934,7 +910,6 @@ func (m *awsRestjson1_serializeOpCreateAssistant) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssistantInput(v *CreateAssistantInput, encoder *httpbinding.Encoder) error {
@@ -996,10 +971,8 @@ func (m *awsRestjson1_serializeOpCreateAssistantAssoci
 func (m *awsRestjson1_serializeOpCreateAssistantAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1048,7 +1021,6 @@ func (m *awsRestjson1_serializeOpCreateAssistantAssoci
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssistantAssociationInput(v *CreateAssistantAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1109,10 +1081,8 @@ func (m *awsRestjson1_serializeOpCreateContent) Handle
 func (m *awsRestjson1_serializeOpCreateContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1161,7 +1131,6 @@ func (m *awsRestjson1_serializeOpCreateContent) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContentInput(v *CreateContentInput, encoder *httpbinding.Encoder) error {
@@ -1237,10 +1206,8 @@ func (m *awsRestjson1_serializeOpCreateContentAssociat
 func (m *awsRestjson1_serializeOpCreateContentAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1289,7 +1256,6 @@ func (m *awsRestjson1_serializeOpCreateContentAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateContentAssociationInput(v *CreateContentAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1359,10 +1325,8 @@ func (m *awsRestjson1_serializeOpCreateKnowledgeBase) 
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
@@ -1407,7 +1371,6 @@ func (m *awsRestjson1_serializeOpCreateKnowledgeBase) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateKnowledgeBaseInput(v *CreateKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -1490,10 +1453,8 @@ func (m *awsRestjson1_serializeOpCreateMessageTemplate
 func (m *awsRestjson1_serializeOpCreateMessageTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1542,7 +1503,6 @@ func (m *awsRestjson1_serializeOpCreateMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMessageTemplateInput(v *CreateMessageTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1632,10 +1592,8 @@ func (m *awsRestjson1_serializeOpCreateMessageTemplate
 func (m *awsRestjson1_serializeOpCreateMessageTemplateAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1684,7 +1642,6 @@ func (m *awsRestjson1_serializeOpCreateMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMessageTemplateAttachmentInput(v *CreateMessageTemplateAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -1750,10 +1707,8 @@ func (m *awsRestjson1_serializeOpCreateMessageTemplate
 func (m *awsRestjson1_serializeOpCreateMessageTemplateVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1802,7 +1757,6 @@ func (m *awsRestjson1_serializeOpCreateMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMessageTemplateVersionInput(v *CreateMessageTemplateVersionInput, encoder *httpbinding.Encoder) error {
@@ -1853,10 +1807,8 @@ func (m *awsRestjson1_serializeOpCreateQuickResponse) 
 func (m *awsRestjson1_serializeOpCreateQuickResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1905,7 +1857,6 @@ func (m *awsRestjson1_serializeOpCreateQuickResponse) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateQuickResponseInput(v *CreateQuickResponseInput, encoder *httpbinding.Encoder) error {
@@ -2005,10 +1956,8 @@ func (m *awsRestjson1_serializeOpCreateSession) Handle
 func (m *awsRestjson1_serializeOpCreateSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2057,7 +2006,6 @@ func (m *awsRestjson1_serializeOpCreateSession) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSessionInput(v *CreateSessionInput, encoder *httpbinding.Encoder) error {
@@ -2130,10 +2078,8 @@ func (m *awsRestjson1_serializeOpDeactivateMessageTemp
 func (m *awsRestjson1_serializeOpDeactivateMessageTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2182,7 +2128,6 @@ func (m *awsRestjson1_serializeOpDeactivateMessageTemp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeactivateMessageTemplateInput(v *DeactivateMessageTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2233,10 +2178,8 @@ func (m *awsRestjson1_serializeOpDeleteAIAgent) Handle
 func (m *awsRestjson1_serializeOpDeleteAIAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2274,7 +2217,6 @@ func (m *awsRestjson1_serializeOpDeleteAIAgent) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAIAgentInput(v *DeleteAIAgentInput, encoder *httpbinding.Encoder) error {
@@ -2313,10 +2255,8 @@ func (m *awsRestjson1_serializeOpDeleteAIAgentVersion)
 func (m *awsRestjson1_serializeOpDeleteAIAgentVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2354,7 +2294,6 @@ func (m *awsRestjson1_serializeOpDeleteAIAgentVersion)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAIAgentVersionInput(v *DeleteAIAgentVersionInput, encoder *httpbinding.Encoder) error {
@@ -2402,10 +2341,8 @@ func (m *awsRestjson1_serializeOpDeleteAIGuardrail) Ha
 func (m *awsRestjson1_serializeOpDeleteAIGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2443,7 +2380,6 @@ func (m *awsRestjson1_serializeOpDeleteAIGuardrail) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAIGuardrailInput(v *DeleteAIGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -2482,10 +2418,8 @@ func (m *awsRestjson1_serializeOpDeleteAIGuardrailVers
 func (m *awsRestjson1_serializeOpDeleteAIGuardrailVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2523,7 +2457,6 @@ func (m *awsRestjson1_serializeOpDeleteAIGuardrailVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAIGuardrailVersionInput(v *DeleteAIGuardrailVersionInput, encoder *httpbinding.Encoder) error {
@@ -2571,10 +2504,8 @@ func (m *awsRestjson1_serializeOpDeleteAIPrompt) Handl
 func (m *awsRestjson1_serializeOpDeleteAIPrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2612,7 +2543,6 @@ func (m *awsRestjson1_serializeOpDeleteAIPrompt) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAIPromptInput(v *DeleteAIPromptInput, encoder *httpbinding.Encoder) error {
@@ -2651,10 +2581,8 @@ func (m *awsRestjson1_serializeOpDeleteAIPromptVersion
 func (m *awsRestjson1_serializeOpDeleteAIPromptVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2692,7 +2620,6 @@ func (m *awsRestjson1_serializeOpDeleteAIPromptVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAIPromptVersionInput(v *DeleteAIPromptVersionInput, encoder *httpbinding.Encoder) error {
@@ -2740,10 +2667,8 @@ func (m *awsRestjson1_serializeOpDeleteAssistant) Hand
 func (m *awsRestjson1_serializeOpDeleteAssistant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2781,7 +2706,6 @@ func (m *awsRestjson1_serializeOpDeleteAssistant) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssistantInput(v *DeleteAssistantInput, encoder *httpbinding.Encoder) error {
@@ -2811,10 +2735,8 @@ func (m *awsRestjson1_serializeOpDeleteAssistantAssoci
 func (m *awsRestjson1_serializeOpDeleteAssistantAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2852,7 +2774,6 @@ func (m *awsRestjson1_serializeOpDeleteAssistantAssoci
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssistantAssociationInput(v *DeleteAssistantAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2891,10 +2812,8 @@ func (m *awsRestjson1_serializeOpDeleteContent) Handle
 func (m *awsRestjson1_serializeOpDeleteContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2932,7 +2851,6 @@ func (m *awsRestjson1_serializeOpDeleteContent) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContentInput(v *DeleteContentInput, encoder *httpbinding.Encoder) error {
@@ -2971,10 +2889,8 @@ func (m *awsRestjson1_serializeOpDeleteContentAssociat
 func (m *awsRestjson1_serializeOpDeleteContentAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3012,7 +2928,6 @@ func (m *awsRestjson1_serializeOpDeleteContentAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteContentAssociationInput(v *DeleteContentAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3060,10 +2975,8 @@ func (m *awsRestjson1_serializeOpDeleteImportJob) Hand
 func (m *awsRestjson1_serializeOpDeleteImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3101,7 +3014,6 @@ func (m *awsRestjson1_serializeOpDeleteImportJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteImportJobInput(v *DeleteImportJobInput, encoder *httpbinding.Encoder) error {
@@ -3140,10 +3052,8 @@ func (m *awsRestjson1_serializeOpDeleteKnowledgeBase) 
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
@@ -3181,7 +3091,6 @@ func (m *awsRestjson1_serializeOpDeleteKnowledgeBase) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteKnowledgeBaseInput(v *DeleteKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -3211,10 +3120,8 @@ func (m *awsRestjson1_serializeOpDeleteMessageTemplate
 func (m *awsRestjson1_serializeOpDeleteMessageTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3252,7 +3159,6 @@ func (m *awsRestjson1_serializeOpDeleteMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMessageTemplateInput(v *DeleteMessageTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3291,10 +3197,8 @@ func (m *awsRestjson1_serializeOpDeleteMessageTemplate
 func (m *awsRestjson1_serializeOpDeleteMessageTemplateAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3332,7 +3236,6 @@ func (m *awsRestjson1_serializeOpDeleteMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMessageTemplateAttachmentInput(v *DeleteMessageTemplateAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -3380,10 +3283,8 @@ func (m *awsRestjson1_serializeOpDeleteQuickResponse) 
 func (m *awsRestjson1_serializeOpDeleteQuickResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3421,7 +3322,6 @@ func (m *awsRestjson1_serializeOpDeleteQuickResponse) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQuickResponseInput(v *DeleteQuickResponseInput, encoder *httpbinding.Encoder) error {
@@ -3460,10 +3360,8 @@ func (m *awsRestjson1_serializeOpGetAIAgent) HandleSer
 func (m *awsRestjson1_serializeOpGetAIAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3501,7 +3399,6 @@ func (m *awsRestjson1_serializeOpGetAIAgent) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAIAgentInput(v *GetAIAgentInput, encoder *httpbinding.Encoder) error {
@@ -3540,10 +3437,8 @@ func (m *awsRestjson1_serializeOpGetAIGuardrail) Handl
 func (m *awsRestjson1_serializeOpGetAIGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3581,7 +3476,6 @@ func (m *awsRestjson1_serializeOpGetAIGuardrail) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAIGuardrailInput(v *GetAIGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -3620,10 +3514,8 @@ func (m *awsRestjson1_serializeOpGetAIPrompt) HandleSe
 func (m *awsRestjson1_serializeOpGetAIPrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3661,7 +3553,6 @@ func (m *awsRestjson1_serializeOpGetAIPrompt) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAIPromptInput(v *GetAIPromptInput, encoder *httpbinding.Encoder) error {
@@ -3700,10 +3591,8 @@ func (m *awsRestjson1_serializeOpGetAssistant) HandleS
 func (m *awsRestjson1_serializeOpGetAssistant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3741,7 +3630,6 @@ func (m *awsRestjson1_serializeOpGetAssistant) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssistantInput(v *GetAssistantInput, encoder *httpbinding.Encoder) error {
@@ -3771,10 +3659,8 @@ func (m *awsRestjson1_serializeOpGetAssistantAssociati
 func (m *awsRestjson1_serializeOpGetAssistantAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3812,7 +3698,6 @@ func (m *awsRestjson1_serializeOpGetAssistantAssociati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssistantAssociationInput(v *GetAssistantAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3851,10 +3736,8 @@ func (m *awsRestjson1_serializeOpGetContent) HandleSer
 func (m *awsRestjson1_serializeOpGetContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3892,7 +3775,6 @@ func (m *awsRestjson1_serializeOpGetContent) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContentInput(v *GetContentInput, encoder *httpbinding.Encoder) error {
@@ -3931,10 +3813,8 @@ func (m *awsRestjson1_serializeOpGetContentAssociation
 func (m *awsRestjson1_serializeOpGetContentAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3972,7 +3852,6 @@ func (m *awsRestjson1_serializeOpGetContentAssociation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContentAssociationInput(v *GetContentAssociationInput, encoder *httpbinding.Encoder) error {
@@ -4020,10 +3899,8 @@ func (m *awsRestjson1_serializeOpGetContentSummary) Ha
 func (m *awsRestjson1_serializeOpGetContentSummary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4061,7 +3938,6 @@ func (m *awsRestjson1_serializeOpGetContentSummary) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetContentSummaryInput(v *GetContentSummaryInput, encoder *httpbinding.Encoder) error {
@@ -4100,10 +3976,8 @@ func (m *awsRestjson1_serializeOpGetImportJob) HandleS
 func (m *awsRestjson1_serializeOpGetImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4141,7 +4015,6 @@ func (m *awsRestjson1_serializeOpGetImportJob) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetImportJobInput(v *GetImportJobInput, encoder *httpbinding.Encoder) error {
@@ -4180,10 +4053,8 @@ func (m *awsRestjson1_serializeOpGetKnowledgeBase) Han
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
@@ -4221,7 +4092,6 @@ func (m *awsRestjson1_serializeOpGetKnowledgeBase) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetKnowledgeBaseInput(v *GetKnowledgeBaseInput, encoder *httpbinding.Encoder) error {
@@ -4251,10 +4121,8 @@ func (m *awsRestjson1_serializeOpGetMessageTemplate) H
 func (m *awsRestjson1_serializeOpGetMessageTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4292,7 +4160,6 @@ func (m *awsRestjson1_serializeOpGetMessageTemplate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMessageTemplateInput(v *GetMessageTemplateInput, encoder *httpbinding.Encoder) error {
@@ -4331,10 +4198,8 @@ func (m *awsRestjson1_serializeOpGetNextMessage) Handl
 func (m *awsRestjson1_serializeOpGetNextMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4372,7 +4237,6 @@ func (m *awsRestjson1_serializeOpGetNextMessage) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNextMessageInput(v *GetNextMessageInput, encoder *httpbinding.Encoder) error {
@@ -4415,10 +4279,8 @@ func (m *awsRestjson1_serializeOpGetQuickResponse) Han
 func (m *awsRestjson1_serializeOpGetQuickResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4456,7 +4318,6 @@ func (m *awsRestjson1_serializeOpGetQuickResponse) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQuickResponseInput(v *GetQuickResponseInput, encoder *httpbinding.Encoder) error {
@@ -4495,10 +4356,8 @@ func (m *awsRestjson1_serializeOpGetRecommendations) H
 func (m *awsRestjson1_serializeOpGetRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4536,7 +4395,6 @@ func (m *awsRestjson1_serializeOpGetRecommendations) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRecommendationsInput(v *GetRecommendationsInput, encoder *httpbinding.Encoder) error {
@@ -4587,10 +4445,8 @@ func (m *awsRestjson1_serializeOpGetSession) HandleSer
 func (m *awsRestjson1_serializeOpGetSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4628,7 +4484,6 @@ func (m *awsRestjson1_serializeOpGetSession) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSessionInput(v *GetSessionInput, encoder *httpbinding.Encoder) error {
@@ -4667,10 +4522,8 @@ func (m *awsRestjson1_serializeOpListAIAgents) HandleS
 func (m *awsRestjson1_serializeOpListAIAgents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4708,7 +4561,6 @@ func (m *awsRestjson1_serializeOpListAIAgents) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAIAgentsInput(v *ListAIAgentsInput, encoder *httpbinding.Encoder) error {
@@ -4750,10 +4602,8 @@ func (m *awsRestjson1_serializeOpListAIAgentVersions) 
 func (m *awsRestjson1_serializeOpListAIAgentVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4791,7 +4641,6 @@ func (m *awsRestjson1_serializeOpListAIAgentVersions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAIAgentVersionsInput(v *ListAIAgentVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4842,10 +4691,8 @@ func (m *awsRestjson1_serializeOpListAIGuardrails) Han
 func (m *awsRestjson1_serializeOpListAIGuardrails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4883,7 +4730,6 @@ func (m *awsRestjson1_serializeOpListAIGuardrails) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAIGuardrailsInput(v *ListAIGuardrailsInput, encoder *httpbinding.Encoder) error {
@@ -4921,10 +4767,8 @@ func (m *awsRestjson1_serializeOpListAIGuardrailVersio
 func (m *awsRestjson1_serializeOpListAIGuardrailVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4962,7 +4806,6 @@ func (m *awsRestjson1_serializeOpListAIGuardrailVersio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAIGuardrailVersionsInput(v *ListAIGuardrailVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5009,10 +4852,8 @@ func (m *awsRestjson1_serializeOpListAIPrompts) Handle
 func (m *awsRestjson1_serializeOpListAIPrompts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5050,7 +4891,6 @@ func (m *awsRestjson1_serializeOpListAIPrompts) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAIPromptsInput(v *ListAIPromptsInput, encoder *httpbinding.Encoder) error {
@@ -5092,10 +4932,8 @@ func (m *awsRestjson1_serializeOpListAIPromptVersions)
 func (m *awsRestjson1_serializeOpListAIPromptVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5133,7 +4971,6 @@ func (m *awsRestjson1_serializeOpListAIPromptVersions)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAIPromptVersionsInput(v *ListAIPromptVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5184,10 +5021,8 @@ func (m *awsRestjson1_serializeOpListAssistantAssociat
 func (m *awsRestjson1_serializeOpListAssistantAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5225,7 +5060,6 @@ func (m *awsRestjson1_serializeOpListAssistantAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssistantAssociationsInput(v *ListAssistantAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -5263,10 +5097,8 @@ func (m *awsRestjson1_serializeOpListAssistants) Handl
 func (m *awsRestjson1_serializeOpListAssistants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5304,7 +5136,6 @@ func (m *awsRestjson1_serializeOpListAssistants) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssistantsInput(v *ListAssistantsInput, encoder *httpbinding.Encoder) error {
@@ -5333,10 +5164,8 @@ func (m *awsRestjson1_serializeOpListContentAssociatio
 func (m *awsRestjson1_serializeOpListContentAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5374,7 +5203,6 @@ func (m *awsRestjson1_serializeOpListContentAssociatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContentAssociationsInput(v *ListContentAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -5421,10 +5249,8 @@ func (m *awsRestjson1_serializeOpListContents) HandleS
 func (m *awsRestjson1_serializeOpListContents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5462,7 +5288,6 @@ func (m *awsRestjson1_serializeOpListContents) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListContentsInput(v *ListContentsInput, encoder *httpbinding.Encoder) error {
@@ -5500,10 +5325,8 @@ func (m *awsRestjson1_serializeOpListImportJobs) Handl
 func (m *awsRestjson1_serializeOpListImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5541,7 +5364,6 @@ func (m *awsRestjson1_serializeOpListImportJobs) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListImportJobsInput(v *ListImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -5579,10 +5401,8 @@ func (m *awsRestjson1_serializeOpListKnowledgeBases) H
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
@@ -5620,7 +5440,6 @@ func (m *awsRestjson1_serializeOpListKnowledgeBases) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListKnowledgeBasesInput(v *ListKnowledgeBasesInput, encoder *httpbinding.Encoder) error {
@@ -5649,10 +5468,8 @@ func (m *awsRestjson1_serializeOpListMessages) HandleS
 func (m *awsRestjson1_serializeOpListMessages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5690,7 +5507,6 @@ func (m *awsRestjson1_serializeOpListMessages) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMessagesInput(v *ListMessagesInput, encoder *httpbinding.Encoder) error {
@@ -5737,10 +5553,8 @@ func (m *awsRestjson1_serializeOpListMessageTemplates)
 func (m *awsRestjson1_serializeOpListMessageTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5778,7 +5592,6 @@ func (m *awsRestjson1_serializeOpListMessageTemplates)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMessageTemplatesInput(v *ListMessageTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -5816,10 +5629,8 @@ func (m *awsRestjson1_serializeOpListMessageTemplateVe
 func (m *awsRestjson1_serializeOpListMessageTemplateVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5857,7 +5668,6 @@ func (m *awsRestjson1_serializeOpListMessageTemplateVe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMessageTemplateVersionsInput(v *ListMessageTemplateVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5904,10 +5714,8 @@ func (m *awsRestjson1_serializeOpListQuickResponses) H
 func (m *awsRestjson1_serializeOpListQuickResponses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5945,7 +5753,6 @@ func (m *awsRestjson1_serializeOpListQuickResponses) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQuickResponsesInput(v *ListQuickResponsesInput, encoder *httpbinding.Encoder) error {
@@ -5983,10 +5790,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6024,7 +5829,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6054,10 +5858,8 @@ func (m *awsRestjson1_serializeOpNotifyRecommendations
 func (m *awsRestjson1_serializeOpNotifyRecommendationsReceived) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6106,7 +5908,6 @@ func (m *awsRestjson1_serializeOpNotifyRecommendations
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsNotifyRecommendationsReceivedInput(v *NotifyRecommendationsReceivedInput, encoder *httpbinding.Encoder) error {
@@ -6159,10 +5960,8 @@ func (m *awsRestjson1_serializeOpPutFeedback) HandleSe
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
@@ -6211,7 +6010,6 @@ func (m *awsRestjson1_serializeOpPutFeedback) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutFeedbackInput(v *PutFeedbackInput, encoder *httpbinding.Encoder) error {
@@ -6265,10 +6063,8 @@ func (m *awsRestjson1_serializeOpQueryAssistant) Handl
 func (m *awsRestjson1_serializeOpQueryAssistant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6317,7 +6113,6 @@ func (m *awsRestjson1_serializeOpQueryAssistant) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsQueryAssistantInput(v *QueryAssistantInput, encoder *httpbinding.Encoder) error {
@@ -6393,10 +6188,8 @@ func (m *awsRestjson1_serializeOpRemoveAssistantAIAgen
 func (m *awsRestjson1_serializeOpRemoveAssistantAIAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6434,7 +6227,6 @@ func (m *awsRestjson1_serializeOpRemoveAssistantAIAgen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveAssistantAIAgentInput(v *RemoveAssistantAIAgentInput, encoder *httpbinding.Encoder) error {
@@ -6468,10 +6260,8 @@ func (m *awsRestjson1_serializeOpRemoveKnowledgeBaseTe
 func (m *awsRestjson1_serializeOpRemoveKnowledgeBaseTemplateUri) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6509,7 +6299,6 @@ func (m *awsRestjson1_serializeOpRemoveKnowledgeBaseTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveKnowledgeBaseTemplateUriInput(v *RemoveKnowledgeBaseTemplateUriInput, encoder *httpbinding.Encoder) error {
@@ -6539,10 +6328,8 @@ func (m *awsRestjson1_serializeOpRenderMessageTemplate
 func (m *awsRestjson1_serializeOpRenderMessageTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6591,7 +6378,6 @@ func (m *awsRestjson1_serializeOpRenderMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRenderMessageTemplateInput(v *RenderMessageTemplateInput, encoder *httpbinding.Encoder) error {
@@ -6644,10 +6430,8 @@ func (m *awsRestjson1_serializeOpSearchContent) Handle
 func (m *awsRestjson1_serializeOpSearchContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6696,7 +6480,6 @@ func (m *awsRestjson1_serializeOpSearchContent) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchContentInput(v *SearchContentInput, encoder *httpbinding.Encoder) error {
@@ -6748,10 +6531,8 @@ func (m *awsRestjson1_serializeOpSearchMessageTemplate
 func (m *awsRestjson1_serializeOpSearchMessageTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6800,7 +6581,6 @@ func (m *awsRestjson1_serializeOpSearchMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchMessageTemplatesInput(v *SearchMessageTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -6852,10 +6632,8 @@ func (m *awsRestjson1_serializeOpSearchQuickResponses)
 func (m *awsRestjson1_serializeOpSearchQuickResponses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6904,7 +6682,6 @@ func (m *awsRestjson1_serializeOpSearchQuickResponses)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchQuickResponsesInput(v *SearchQuickResponsesInput, encoder *httpbinding.Encoder) error {
@@ -6963,10 +6740,8 @@ func (m *awsRestjson1_serializeOpSearchSessions) Handl
 func (m *awsRestjson1_serializeOpSearchSessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7015,7 +6790,6 @@ func (m *awsRestjson1_serializeOpSearchSessions) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchSessionsInput(v *SearchSessionsInput, encoder *httpbinding.Encoder) error {
@@ -7067,10 +6841,8 @@ func (m *awsRestjson1_serializeOpSendMessage) HandleSe
 func (m *awsRestjson1_serializeOpSendMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7119,7 +6891,6 @@ func (m *awsRestjson1_serializeOpSendMessage) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendMessageInput(v *SendMessageInput, encoder *httpbinding.Encoder) error {
@@ -7196,10 +6967,8 @@ func (m *awsRestjson1_serializeOpStartContentUpload) H
 func (m *awsRestjson1_serializeOpStartContentUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7248,7 +7017,6 @@ func (m *awsRestjson1_serializeOpStartContentUpload) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartContentUploadInput(v *StartContentUploadInput, encoder *httpbinding.Encoder) error {
@@ -7295,10 +7063,8 @@ func (m *awsRestjson1_serializeOpStartImportJob) Handl
 func (m *awsRestjson1_serializeOpStartImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7347,7 +7113,6 @@ func (m *awsRestjson1_serializeOpStartImportJob) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartImportJobInput(v *StartImportJobInput, encoder *httpbinding.Encoder) error {
@@ -7413,10 +7178,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7465,7 +7228,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7509,10 +7271,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7550,7 +7310,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7586,10 +7345,8 @@ func (m *awsRestjson1_serializeOpUpdateAIAgent) Handle
 func (m *awsRestjson1_serializeOpUpdateAIAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7638,7 +7395,6 @@ func (m *awsRestjson1_serializeOpUpdateAIAgent) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAIAgentInput(v *UpdateAIAgentInput, encoder *httpbinding.Encoder) error {
@@ -7706,10 +7462,8 @@ func (m *awsRestjson1_serializeOpUpdateAIGuardrail) Ha
 func (m *awsRestjson1_serializeOpUpdateAIGuardrail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7758,7 +7512,6 @@ func (m *awsRestjson1_serializeOpUpdateAIGuardrail) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAIGuardrailInput(v *UpdateAIGuardrailInput, encoder *httpbinding.Encoder) error {
@@ -7864,10 +7617,8 @@ func (m *awsRestjson1_serializeOpUpdateAIPrompt) Handl
 func (m *awsRestjson1_serializeOpUpdateAIPrompt) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7916,7 +7667,6 @@ func (m *awsRestjson1_serializeOpUpdateAIPrompt) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAIPromptInput(v *UpdateAIPromptInput, encoder *httpbinding.Encoder) error {
@@ -7984,10 +7734,8 @@ func (m *awsRestjson1_serializeOpUpdateAssistantAIAgen
 func (m *awsRestjson1_serializeOpUpdateAssistantAIAgent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8036,7 +7784,6 @@ func (m *awsRestjson1_serializeOpUpdateAssistantAIAgen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssistantAIAgentInput(v *UpdateAssistantAIAgentInput, encoder *httpbinding.Encoder) error {
@@ -8085,10 +7832,8 @@ func (m *awsRestjson1_serializeOpUpdateContent) Handle
 func (m *awsRestjson1_serializeOpUpdateContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8137,7 +7882,6 @@ func (m *awsRestjson1_serializeOpUpdateContent) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateContentInput(v *UpdateContentInput, encoder *httpbinding.Encoder) error {
@@ -8215,10 +7959,8 @@ func (m *awsRestjson1_serializeOpUpdateKnowledgeBaseTe
 func (m *awsRestjson1_serializeOpUpdateKnowledgeBaseTemplateUri) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8267,7 +8009,6 @@ func (m *awsRestjson1_serializeOpUpdateKnowledgeBaseTe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateKnowledgeBaseTemplateUriInput(v *UpdateKnowledgeBaseTemplateUriInput, encoder *httpbinding.Encoder) error {
@@ -8309,10 +8050,8 @@ func (m *awsRestjson1_serializeOpUpdateMessageTemplate
 func (m *awsRestjson1_serializeOpUpdateMessageTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8361,7 +8100,6 @@ func (m *awsRestjson1_serializeOpUpdateMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMessageTemplateInput(v *UpdateMessageTemplateInput, encoder *httpbinding.Encoder) error {
@@ -8426,10 +8164,8 @@ func (m *awsRestjson1_serializeOpUpdateMessageTemplate
 func (m *awsRestjson1_serializeOpUpdateMessageTemplateMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8478,7 +8214,6 @@ func (m *awsRestjson1_serializeOpUpdateMessageTemplate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMessageTemplateMetadataInput(v *UpdateMessageTemplateMetadataInput, encoder *httpbinding.Encoder) error {
@@ -8541,10 +8276,8 @@ func (m *awsRestjson1_serializeOpUpdateQuickResponse) 
 func (m *awsRestjson1_serializeOpUpdateQuickResponse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8593,7 +8326,6 @@ func (m *awsRestjson1_serializeOpUpdateQuickResponse) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQuickResponseInput(v *UpdateQuickResponseInput, encoder *httpbinding.Encoder) error {
@@ -8705,10 +8437,8 @@ func (m *awsRestjson1_serializeOpUpdateSession) Handle
 func (m *awsRestjson1_serializeOpUpdateSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8757,7 +8487,6 @@ func (m *awsRestjson1_serializeOpUpdateSession) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSessionInput(v *UpdateSessionInput, encoder *httpbinding.Encoder) error {
@@ -8822,10 +8551,8 @@ func (m *awsRestjson1_serializeOpUpdateSessionData) Ha
 func (m *awsRestjson1_serializeOpUpdateSessionData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8874,7 +8601,6 @@ func (m *awsRestjson1_serializeOpUpdateSessionData) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSessionDataInput(v *UpdateSessionDataInput, encoder *httpbinding.Encoder) error {
