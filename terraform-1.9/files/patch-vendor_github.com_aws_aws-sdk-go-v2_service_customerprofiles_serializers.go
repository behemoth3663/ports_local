--- vendor/github.com/aws/aws-sdk-go-v2/service/customerprofiles/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/customerprofiles/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAddProfileKey) Handle
 func (m *awsRestjson1_serializeOpAddProfileKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpAddProfileKey) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddProfileKeyInput(v *AddProfileKeyInput, encoder *httpbinding.Encoder) error {
@@ -133,10 +129,8 @@ func (m *awsRestjson1_serializeOpBatchGetCalculatedAtt
 func (m *awsRestjson1_serializeOpBatchGetCalculatedAttributeForProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -185,7 +179,6 @@ func (m *awsRestjson1_serializeOpBatchGetCalculatedAtt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetCalculatedAttributeForProfileInput(v *BatchGetCalculatedAttributeForProfileInput, encoder *httpbinding.Encoder) error {
@@ -245,10 +238,8 @@ func (m *awsRestjson1_serializeOpBatchGetProfile) Hand
 func (m *awsRestjson1_serializeOpBatchGetProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -297,7 +288,6 @@ func (m *awsRestjson1_serializeOpBatchGetProfile) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetProfileInput(v *BatchGetProfileInput, encoder *httpbinding.Encoder) error {
@@ -341,10 +331,8 @@ func (m *awsRestjson1_serializeOpCreateCalculatedAttri
 func (m *awsRestjson1_serializeOpCreateCalculatedAttributeDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -393,7 +381,6 @@ func (m *awsRestjson1_serializeOpCreateCalculatedAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCalculatedAttributeDefinitionInput(v *CreateCalculatedAttributeDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -487,10 +474,8 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
 func (m *awsRestjson1_serializeOpCreateDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -539,7 +524,6 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainInput(v *CreateDomainInput, encoder *httpbinding.Encoder) error {
@@ -612,10 +596,8 @@ func (m *awsRestjson1_serializeOpCreateDomainLayout) H
 func (m *awsRestjson1_serializeOpCreateDomainLayout) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -664,7 +646,6 @@ func (m *awsRestjson1_serializeOpCreateDomainLayout) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainLayoutInput(v *CreateDomainLayoutInput, encoder *httpbinding.Encoder) error {
@@ -742,10 +723,8 @@ func (m *awsRestjson1_serializeOpCreateEventStream) Ha
 func (m *awsRestjson1_serializeOpCreateEventStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -794,7 +773,6 @@ func (m *awsRestjson1_serializeOpCreateEventStream) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEventStreamInput(v *CreateEventStreamInput, encoder *httpbinding.Encoder) error {
@@ -852,10 +830,8 @@ func (m *awsRestjson1_serializeOpCreateEventTrigger) H
 func (m *awsRestjson1_serializeOpCreateEventTrigger) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -904,7 +880,6 @@ func (m *awsRestjson1_serializeOpCreateEventTrigger) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEventTriggerInput(v *CreateEventTriggerInput, encoder *httpbinding.Encoder) error {
@@ -986,10 +961,8 @@ func (m *awsRestjson1_serializeOpCreateIntegrationWork
 func (m *awsRestjson1_serializeOpCreateIntegrationWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1038,7 +1011,6 @@ func (m *awsRestjson1_serializeOpCreateIntegrationWork
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIntegrationWorkflowInput(v *CreateIntegrationWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -1104,10 +1076,8 @@ func (m *awsRestjson1_serializeOpCreateProfile) Handle
 func (m *awsRestjson1_serializeOpCreateProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1156,7 +1126,6 @@ func (m *awsRestjson1_serializeOpCreateProfile) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProfileInput(v *CreateProfileInput, encoder *httpbinding.Encoder) error {
@@ -1318,10 +1287,8 @@ func (m *awsRestjson1_serializeOpCreateSegmentDefiniti
 func (m *awsRestjson1_serializeOpCreateSegmentDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1370,7 +1337,6 @@ func (m *awsRestjson1_serializeOpCreateSegmentDefiniti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSegmentDefinitionInput(v *CreateSegmentDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -1440,10 +1406,8 @@ func (m *awsRestjson1_serializeOpCreateSegmentEstimate
 func (m *awsRestjson1_serializeOpCreateSegmentEstimate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1492,7 +1456,6 @@ func (m *awsRestjson1_serializeOpCreateSegmentEstimate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSegmentEstimateInput(v *CreateSegmentEstimateInput, encoder *httpbinding.Encoder) error {
@@ -1536,10 +1499,8 @@ func (m *awsRestjson1_serializeOpCreateSegmentSnapshot
 func (m *awsRestjson1_serializeOpCreateSegmentSnapshot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1588,7 +1549,6 @@ func (m *awsRestjson1_serializeOpCreateSegmentSnapshot
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSegmentSnapshotInput(v *CreateSegmentSnapshotInput, encoder *httpbinding.Encoder) error {
@@ -1654,10 +1614,8 @@ func (m *awsRestjson1_serializeOpCreateUploadJob) Hand
 func (m *awsRestjson1_serializeOpCreateUploadJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1706,7 +1664,6 @@ func (m *awsRestjson1_serializeOpCreateUploadJob) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUploadJobInput(v *CreateUploadJobInput, encoder *httpbinding.Encoder) error {
@@ -1765,10 +1722,8 @@ func (m *awsRestjson1_serializeOpDeleteCalculatedAttri
 func (m *awsRestjson1_serializeOpDeleteCalculatedAttributeDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1806,7 +1761,6 @@ func (m *awsRestjson1_serializeOpDeleteCalculatedAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCalculatedAttributeDefinitionInput(v *DeleteCalculatedAttributeDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -1845,10 +1799,8 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
 func (m *awsRestjson1_serializeOpDeleteDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1886,7 +1838,6 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainInput(v *DeleteDomainInput, encoder *httpbinding.Encoder) error {
@@ -1916,10 +1867,8 @@ func (m *awsRestjson1_serializeOpDeleteDomainLayout) H
 func (m *awsRestjson1_serializeOpDeleteDomainLayout) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1957,7 +1906,6 @@ func (m *awsRestjson1_serializeOpDeleteDomainLayout) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainLayoutInput(v *DeleteDomainLayoutInput, encoder *httpbinding.Encoder) error {
@@ -1996,10 +1944,8 @@ func (m *awsRestjson1_serializeOpDeleteEventStream) Ha
 func (m *awsRestjson1_serializeOpDeleteEventStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2037,7 +1983,6 @@ func (m *awsRestjson1_serializeOpDeleteEventStream) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventStreamInput(v *DeleteEventStreamInput, encoder *httpbinding.Encoder) error {
@@ -2076,10 +2021,8 @@ func (m *awsRestjson1_serializeOpDeleteEventTrigger) H
 func (m *awsRestjson1_serializeOpDeleteEventTrigger) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2117,7 +2060,6 @@ func (m *awsRestjson1_serializeOpDeleteEventTrigger) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventTriggerInput(v *DeleteEventTriggerInput, encoder *httpbinding.Encoder) error {
@@ -2156,10 +2098,8 @@ func (m *awsRestjson1_serializeOpDeleteIntegration) Ha
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
@@ -2208,7 +2148,6 @@ func (m *awsRestjson1_serializeOpDeleteIntegration) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIntegrationInput(v *DeleteIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -2250,10 +2189,8 @@ func (m *awsRestjson1_serializeOpDeleteProfile) Handle
 func (m *awsRestjson1_serializeOpDeleteProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2302,7 +2239,6 @@ func (m *awsRestjson1_serializeOpDeleteProfile) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProfileInput(v *DeleteProfileInput, encoder *httpbinding.Encoder) error {
@@ -2344,10 +2280,8 @@ func (m *awsRestjson1_serializeOpDeleteProfileKey) Han
 func (m *awsRestjson1_serializeOpDeleteProfileKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2396,7 +2330,6 @@ func (m *awsRestjson1_serializeOpDeleteProfileKey) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProfileKeyInput(v *DeleteProfileKeyInput, encoder *httpbinding.Encoder) error {
@@ -2450,10 +2383,8 @@ func (m *awsRestjson1_serializeOpDeleteProfileObject) 
 func (m *awsRestjson1_serializeOpDeleteProfileObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2502,7 +2433,6 @@ func (m *awsRestjson1_serializeOpDeleteProfileObject) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProfileObjectInput(v *DeleteProfileObjectInput, encoder *httpbinding.Encoder) error {
@@ -2554,10 +2484,8 @@ func (m *awsRestjson1_serializeOpDeleteProfileObjectTy
 func (m *awsRestjson1_serializeOpDeleteProfileObjectType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2595,7 +2523,6 @@ func (m *awsRestjson1_serializeOpDeleteProfileObjectTy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProfileObjectTypeInput(v *DeleteProfileObjectTypeInput, encoder *httpbinding.Encoder) error {
@@ -2634,10 +2561,8 @@ func (m *awsRestjson1_serializeOpDeleteSegmentDefiniti
 func (m *awsRestjson1_serializeOpDeleteSegmentDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2675,7 +2600,6 @@ func (m *awsRestjson1_serializeOpDeleteSegmentDefiniti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSegmentDefinitionInput(v *DeleteSegmentDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2714,10 +2638,8 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
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
@@ -2755,7 +2677,6 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWorkflowInput(v *DeleteWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -2794,10 +2715,8 @@ func (m *awsRestjson1_serializeOpDetectProfileObjectTy
 func (m *awsRestjson1_serializeOpDetectProfileObjectType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2846,7 +2765,6 @@ func (m *awsRestjson1_serializeOpDetectProfileObjectTy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetectProfileObjectTypeInput(v *DetectProfileObjectTypeInput, encoder *httpbinding.Encoder) error {
@@ -2890,10 +2808,8 @@ func (m *awsRestjson1_serializeOpGetAutoMergingPreview
 func (m *awsRestjson1_serializeOpGetAutoMergingPreview) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2942,7 +2858,6 @@ func (m *awsRestjson1_serializeOpGetAutoMergingPreview
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAutoMergingPreviewInput(v *GetAutoMergingPreviewInput, encoder *httpbinding.Encoder) error {
@@ -3011,10 +2926,8 @@ func (m *awsRestjson1_serializeOpGetCalculatedAttribut
 func (m *awsRestjson1_serializeOpGetCalculatedAttributeDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3052,7 +2965,6 @@ func (m *awsRestjson1_serializeOpGetCalculatedAttribut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCalculatedAttributeDefinitionInput(v *GetCalculatedAttributeDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -3091,10 +3003,8 @@ func (m *awsRestjson1_serializeOpGetCalculatedAttribut
 func (m *awsRestjson1_serializeOpGetCalculatedAttributeForProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3132,7 +3042,6 @@ func (m *awsRestjson1_serializeOpGetCalculatedAttribut
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCalculatedAttributeForProfileInput(v *GetCalculatedAttributeForProfileInput, encoder *httpbinding.Encoder) error {
@@ -3180,10 +3089,8 @@ func (m *awsRestjson1_serializeOpGetDomain) HandleSeri
 func (m *awsRestjson1_serializeOpGetDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3221,7 +3128,6 @@ func (m *awsRestjson1_serializeOpGetDomain) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainInput(v *GetDomainInput, encoder *httpbinding.Encoder) error {
@@ -3251,10 +3157,8 @@ func (m *awsRestjson1_serializeOpGetDomainLayout) Hand
 func (m *awsRestjson1_serializeOpGetDomainLayout) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3292,7 +3196,6 @@ func (m *awsRestjson1_serializeOpGetDomainLayout) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainLayoutInput(v *GetDomainLayoutInput, encoder *httpbinding.Encoder) error {
@@ -3331,10 +3234,8 @@ func (m *awsRestjson1_serializeOpGetEventStream) Handl
 func (m *awsRestjson1_serializeOpGetEventStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3372,7 +3273,6 @@ func (m *awsRestjson1_serializeOpGetEventStream) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventStreamInput(v *GetEventStreamInput, encoder *httpbinding.Encoder) error {
@@ -3411,10 +3311,8 @@ func (m *awsRestjson1_serializeOpGetEventTrigger) Hand
 func (m *awsRestjson1_serializeOpGetEventTrigger) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3452,7 +3350,6 @@ func (m *awsRestjson1_serializeOpGetEventTrigger) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventTriggerInput(v *GetEventTriggerInput, encoder *httpbinding.Encoder) error {
@@ -3491,10 +3388,8 @@ func (m *awsRestjson1_serializeOpGetIdentityResolution
 func (m *awsRestjson1_serializeOpGetIdentityResolutionJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3532,7 +3427,6 @@ func (m *awsRestjson1_serializeOpGetIdentityResolution
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIdentityResolutionJobInput(v *GetIdentityResolutionJobInput, encoder *httpbinding.Encoder) error {
@@ -3571,10 +3465,8 @@ func (m *awsRestjson1_serializeOpGetIntegration) Handl
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
@@ -3623,7 +3515,6 @@ func (m *awsRestjson1_serializeOpGetIntegration) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIntegrationInput(v *GetIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -3665,10 +3556,8 @@ func (m *awsRestjson1_serializeOpGetMatches) HandleSer
 func (m *awsRestjson1_serializeOpGetMatches) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3706,7 +3595,6 @@ func (m *awsRestjson1_serializeOpGetMatches) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMatchesInput(v *GetMatchesInput, encoder *httpbinding.Encoder) error {
@@ -3744,10 +3632,8 @@ func (m *awsRestjson1_serializeOpGetProfileObjectType)
 func (m *awsRestjson1_serializeOpGetProfileObjectType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3785,7 +3671,6 @@ func (m *awsRestjson1_serializeOpGetProfileObjectType)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProfileObjectTypeInput(v *GetProfileObjectTypeInput, encoder *httpbinding.Encoder) error {
@@ -3824,10 +3709,8 @@ func (m *awsRestjson1_serializeOpGetProfileObjectTypeT
 func (m *awsRestjson1_serializeOpGetProfileObjectTypeTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3865,7 +3748,6 @@ func (m *awsRestjson1_serializeOpGetProfileObjectTypeT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProfileObjectTypeTemplateInput(v *GetProfileObjectTypeTemplateInput, encoder *httpbinding.Encoder) error {
@@ -3895,10 +3777,8 @@ func (m *awsRestjson1_serializeOpGetSegmentDefinition)
 func (m *awsRestjson1_serializeOpGetSegmentDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3936,7 +3816,6 @@ func (m *awsRestjson1_serializeOpGetSegmentDefinition)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentDefinitionInput(v *GetSegmentDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -3975,10 +3854,8 @@ func (m *awsRestjson1_serializeOpGetSegmentEstimate) H
 func (m *awsRestjson1_serializeOpGetSegmentEstimate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4016,7 +3893,6 @@ func (m *awsRestjson1_serializeOpGetSegmentEstimate) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentEstimateInput(v *GetSegmentEstimateInput, encoder *httpbinding.Encoder) error {
@@ -4055,10 +3931,8 @@ func (m *awsRestjson1_serializeOpGetSegmentMembership)
 func (m *awsRestjson1_serializeOpGetSegmentMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4107,7 +3981,6 @@ func (m *awsRestjson1_serializeOpGetSegmentMembership)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentMembershipInput(v *GetSegmentMembershipInput, encoder *httpbinding.Encoder) error {
@@ -4160,10 +4033,8 @@ func (m *awsRestjson1_serializeOpGetSegmentSnapshot) H
 func (m *awsRestjson1_serializeOpGetSegmentSnapshot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4201,7 +4072,6 @@ func (m *awsRestjson1_serializeOpGetSegmentSnapshot) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSegmentSnapshotInput(v *GetSegmentSnapshotInput, encoder *httpbinding.Encoder) error {
@@ -4249,10 +4119,8 @@ func (m *awsRestjson1_serializeOpGetSimilarProfiles) H
 func (m *awsRestjson1_serializeOpGetSimilarProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4301,7 +4169,6 @@ func (m *awsRestjson1_serializeOpGetSimilarProfiles) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSimilarProfilesInput(v *GetSimilarProfilesInput, encoder *httpbinding.Encoder) error {
@@ -4361,10 +4228,8 @@ func (m *awsRestjson1_serializeOpGetUploadJob) HandleS
 func (m *awsRestjson1_serializeOpGetUploadJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4402,7 +4267,6 @@ func (m *awsRestjson1_serializeOpGetUploadJob) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUploadJobInput(v *GetUploadJobInput, encoder *httpbinding.Encoder) error {
@@ -4441,10 +4305,8 @@ func (m *awsRestjson1_serializeOpGetUploadJobPath) Han
 func (m *awsRestjson1_serializeOpGetUploadJobPath) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4482,7 +4344,6 @@ func (m *awsRestjson1_serializeOpGetUploadJobPath) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUploadJobPathInput(v *GetUploadJobPathInput, encoder *httpbinding.Encoder) error {
@@ -4521,10 +4382,8 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
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
@@ -4562,7 +4421,6 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowInput(v *GetWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -4601,10 +4459,8 @@ func (m *awsRestjson1_serializeOpGetWorkflowSteps) Han
 func (m *awsRestjson1_serializeOpGetWorkflowSteps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4642,7 +4498,6 @@ func (m *awsRestjson1_serializeOpGetWorkflowSteps) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowStepsInput(v *GetWorkflowStepsInput, encoder *httpbinding.Encoder) error {
@@ -4689,10 +4544,8 @@ func (m *awsRestjson1_serializeOpListAccountIntegratio
 func (m *awsRestjson1_serializeOpListAccountIntegrations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4741,7 +4594,6 @@ func (m *awsRestjson1_serializeOpListAccountIntegratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAccountIntegrationsInput(v *ListAccountIntegrationsInput, encoder *httpbinding.Encoder) error {
@@ -4786,10 +4638,8 @@ func (m *awsRestjson1_serializeOpListCalculatedAttribu
 func (m *awsRestjson1_serializeOpListCalculatedAttributeDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4827,7 +4677,6 @@ func (m *awsRestjson1_serializeOpListCalculatedAttribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCalculatedAttributeDefinitionsInput(v *ListCalculatedAttributeDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -4865,10 +4714,8 @@ func (m *awsRestjson1_serializeOpListCalculatedAttribu
 func (m *awsRestjson1_serializeOpListCalculatedAttributesForProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4906,7 +4753,6 @@ func (m *awsRestjson1_serializeOpListCalculatedAttribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCalculatedAttributesForProfileInput(v *ListCalculatedAttributesForProfileInput, encoder *httpbinding.Encoder) error {
@@ -4953,10 +4799,8 @@ func (m *awsRestjson1_serializeOpListDomainLayouts) Ha
 func (m *awsRestjson1_serializeOpListDomainLayouts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4994,7 +4838,6 @@ func (m *awsRestjson1_serializeOpListDomainLayouts) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainLayoutsInput(v *ListDomainLayoutsInput, encoder *httpbinding.Encoder) error {
@@ -5032,10 +4875,8 @@ func (m *awsRestjson1_serializeOpListDomains) HandleSe
 func (m *awsRestjson1_serializeOpListDomains) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5073,7 +4914,6 @@ func (m *awsRestjson1_serializeOpListDomains) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainsInput(v *ListDomainsInput, encoder *httpbinding.Encoder) error {
@@ -5102,10 +4942,8 @@ func (m *awsRestjson1_serializeOpListEventStreams) Han
 func (m *awsRestjson1_serializeOpListEventStreams) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5143,7 +4981,6 @@ func (m *awsRestjson1_serializeOpListEventStreams) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventStreamsInput(v *ListEventStreamsInput, encoder *httpbinding.Encoder) error {
@@ -5181,10 +5018,8 @@ func (m *awsRestjson1_serializeOpListEventTriggers) Ha
 func (m *awsRestjson1_serializeOpListEventTriggers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5222,7 +5057,6 @@ func (m *awsRestjson1_serializeOpListEventTriggers) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventTriggersInput(v *ListEventTriggersInput, encoder *httpbinding.Encoder) error {
@@ -5260,10 +5094,8 @@ func (m *awsRestjson1_serializeOpListIdentityResolutio
 func (m *awsRestjson1_serializeOpListIdentityResolutionJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5301,7 +5133,6 @@ func (m *awsRestjson1_serializeOpListIdentityResolutio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIdentityResolutionJobsInput(v *ListIdentityResolutionJobsInput, encoder *httpbinding.Encoder) error {
@@ -5339,10 +5170,8 @@ func (m *awsRestjson1_serializeOpListIntegrations) Han
 func (m *awsRestjson1_serializeOpListIntegrations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5380,7 +5209,6 @@ func (m *awsRestjson1_serializeOpListIntegrations) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIntegrationsInput(v *ListIntegrationsInput, encoder *httpbinding.Encoder) error {
@@ -5422,10 +5250,8 @@ func (m *awsRestjson1_serializeOpListObjectTypeAttribu
 func (m *awsRestjson1_serializeOpListObjectTypeAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5463,7 +5289,6 @@ func (m *awsRestjson1_serializeOpListObjectTypeAttribu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListObjectTypeAttributesInput(v *ListObjectTypeAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5510,10 +5335,8 @@ func (m *awsRestjson1_serializeOpListProfileAttributeV
 func (m *awsRestjson1_serializeOpListProfileAttributeValues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5551,7 +5374,6 @@ func (m *awsRestjson1_serializeOpListProfileAttributeV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProfileAttributeValuesInput(v *ListProfileAttributeValuesInput, encoder *httpbinding.Encoder) error {
@@ -5590,10 +5412,8 @@ func (m *awsRestjson1_serializeOpListProfileObjects) H
 func (m *awsRestjson1_serializeOpListProfileObjects) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5642,7 +5462,6 @@ func (m *awsRestjson1_serializeOpListProfileObjects) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProfileObjectsInput(v *ListProfileObjectsInput, encoder *httpbinding.Encoder) error {
@@ -5704,10 +5523,8 @@ func (m *awsRestjson1_serializeOpListProfileObjectType
 func (m *awsRestjson1_serializeOpListProfileObjectTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5745,7 +5562,6 @@ func (m *awsRestjson1_serializeOpListProfileObjectType
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProfileObjectTypesInput(v *ListProfileObjectTypesInput, encoder *httpbinding.Encoder) error {
@@ -5783,10 +5599,8 @@ func (m *awsRestjson1_serializeOpListProfileObjectType
 func (m *awsRestjson1_serializeOpListProfileObjectTypeTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5824,7 +5638,6 @@ func (m *awsRestjson1_serializeOpListProfileObjectType
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProfileObjectTypeTemplatesInput(v *ListProfileObjectTypeTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -5853,10 +5666,8 @@ func (m *awsRestjson1_serializeOpListRuleBasedMatches)
 func (m *awsRestjson1_serializeOpListRuleBasedMatches) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5894,7 +5705,6 @@ func (m *awsRestjson1_serializeOpListRuleBasedMatches)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRuleBasedMatchesInput(v *ListRuleBasedMatchesInput, encoder *httpbinding.Encoder) error {
@@ -5932,10 +5742,8 @@ func (m *awsRestjson1_serializeOpListSegmentDefinition
 func (m *awsRestjson1_serializeOpListSegmentDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5973,7 +5781,6 @@ func (m *awsRestjson1_serializeOpListSegmentDefinition
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSegmentDefinitionsInput(v *ListSegmentDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -6011,10 +5818,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6052,7 +5857,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6082,10 +5886,8 @@ func (m *awsRestjson1_serializeOpListUploadJobs) Handl
 func (m *awsRestjson1_serializeOpListUploadJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6123,7 +5925,6 @@ func (m *awsRestjson1_serializeOpListUploadJobs) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUploadJobsInput(v *ListUploadJobsInput, encoder *httpbinding.Encoder) error {
@@ -6161,10 +5962,8 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
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
@@ -6213,7 +6012,6 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowsInput(v *ListWorkflowsInput, encoder *httpbinding.Encoder) error {
@@ -6278,10 +6076,8 @@ func (m *awsRestjson1_serializeOpMergeProfiles) Handle
 func (m *awsRestjson1_serializeOpMergeProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6330,7 +6126,6 @@ func (m *awsRestjson1_serializeOpMergeProfiles) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsMergeProfilesInput(v *MergeProfilesInput, encoder *httpbinding.Encoder) error {
@@ -6386,10 +6181,8 @@ func (m *awsRestjson1_serializeOpPutIntegration) Handl
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
@@ -6438,7 +6231,6 @@ func (m *awsRestjson1_serializeOpPutIntegration) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutIntegrationInput(v *PutIntegrationInput, encoder *httpbinding.Encoder) error {
@@ -6518,10 +6310,8 @@ func (m *awsRestjson1_serializeOpPutProfileObject) Han
 func (m *awsRestjson1_serializeOpPutProfileObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6570,7 +6360,6 @@ func (m *awsRestjson1_serializeOpPutProfileObject) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutProfileObjectInput(v *PutProfileObjectInput, encoder *httpbinding.Encoder) error {
@@ -6617,10 +6406,8 @@ func (m *awsRestjson1_serializeOpPutProfileObjectType)
 func (m *awsRestjson1_serializeOpPutProfileObjectType) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6669,7 +6456,6 @@ func (m *awsRestjson1_serializeOpPutProfileObjectType)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutProfileObjectTypeInput(v *PutProfileObjectTypeInput, encoder *httpbinding.Encoder) error {
@@ -6771,10 +6557,8 @@ func (m *awsRestjson1_serializeOpSearchProfiles) Handl
 func (m *awsRestjson1_serializeOpSearchProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6823,7 +6607,6 @@ func (m *awsRestjson1_serializeOpSearchProfiles) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchProfilesInput(v *SearchProfilesInput, encoder *httpbinding.Encoder) error {
@@ -6892,10 +6675,8 @@ func (m *awsRestjson1_serializeOpStartUploadJob) Handl
 func (m *awsRestjson1_serializeOpStartUploadJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6933,7 +6714,6 @@ func (m *awsRestjson1_serializeOpStartUploadJob) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartUploadJobInput(v *StartUploadJobInput, encoder *httpbinding.Encoder) error {
@@ -6972,10 +6752,8 @@ func (m *awsRestjson1_serializeOpStopUploadJob) Handle
 func (m *awsRestjson1_serializeOpStopUploadJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7013,7 +6791,6 @@ func (m *awsRestjson1_serializeOpStopUploadJob) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopUploadJobInput(v *StopUploadJobInput, encoder *httpbinding.Encoder) error {
@@ -7052,10 +6829,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7104,7 +6879,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7148,10 +6922,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7189,7 +6961,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7225,10 +6996,8 @@ func (m *awsRestjson1_serializeOpUpdateCalculatedAttri
 func (m *awsRestjson1_serializeOpUpdateCalculatedAttributeDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7277,7 +7046,6 @@ func (m *awsRestjson1_serializeOpUpdateCalculatedAttri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCalculatedAttributeDefinitionInput(v *UpdateCalculatedAttributeDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -7340,10 +7108,8 @@ func (m *awsRestjson1_serializeOpUpdateDomain) HandleS
 func (m *awsRestjson1_serializeOpUpdateDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7392,7 +7158,6 @@ func (m *awsRestjson1_serializeOpUpdateDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDomainInput(v *UpdateDomainInput, encoder *httpbinding.Encoder) error {
@@ -7465,10 +7230,8 @@ func (m *awsRestjson1_serializeOpUpdateDomainLayout) H
 func (m *awsRestjson1_serializeOpUpdateDomainLayout) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7517,7 +7280,6 @@ func (m *awsRestjson1_serializeOpUpdateDomainLayout) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDomainLayoutInput(v *UpdateDomainLayoutInput, encoder *httpbinding.Encoder) error {
@@ -7588,10 +7350,8 @@ func (m *awsRestjson1_serializeOpUpdateEventTrigger) H
 func (m *awsRestjson1_serializeOpUpdateEventTrigger) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7640,7 +7400,6 @@ func (m *awsRestjson1_serializeOpUpdateEventTrigger) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventTriggerInput(v *UpdateEventTriggerInput, encoder *httpbinding.Encoder) error {
@@ -7715,10 +7474,8 @@ func (m *awsRestjson1_serializeOpUpdateProfile) Handle
 func (m *awsRestjson1_serializeOpUpdateProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7767,7 +7524,6 @@ func (m *awsRestjson1_serializeOpUpdateProfile) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateProfileInput(v *UpdateProfileInput, encoder *httpbinding.Encoder) error {
