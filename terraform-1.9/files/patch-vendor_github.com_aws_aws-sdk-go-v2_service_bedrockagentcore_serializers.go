--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/serializers.go
@@ -14,7 +14,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -28,10 +27,8 @@ func (m *awsRestjson1_serializeOpCreateEvent) HandleSe
 func (m *awsRestjson1_serializeOpCreateEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -80,7 +77,6 @@ func (m *awsRestjson1_serializeOpCreateEvent) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEventInput(v *CreateEventInput, encoder *httpbinding.Encoder) error {
@@ -151,10 +147,8 @@ func (m *awsRestjson1_serializeOpDeleteEvent) HandleSe
 func (m *awsRestjson1_serializeOpDeleteEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -192,7 +186,6 @@ func (m *awsRestjson1_serializeOpDeleteEvent) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventInput(v *DeleteEventInput, encoder *httpbinding.Encoder) error {
@@ -249,10 +242,8 @@ func (m *awsRestjson1_serializeOpDeleteMemoryRecord) H
 func (m *awsRestjson1_serializeOpDeleteMemoryRecord) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -290,7 +281,6 @@ func (m *awsRestjson1_serializeOpDeleteMemoryRecord) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMemoryRecordInput(v *DeleteMemoryRecordInput, encoder *httpbinding.Encoder) error {
@@ -329,10 +319,8 @@ func (m *awsRestjson1_serializeOpGetBrowserSession) Ha
 func (m *awsRestjson1_serializeOpGetBrowserSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -370,7 +358,6 @@ func (m *awsRestjson1_serializeOpGetBrowserSession) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBrowserSessionInput(v *GetBrowserSessionInput, encoder *httpbinding.Encoder) error {
@@ -404,10 +391,8 @@ func (m *awsRestjson1_serializeOpGetCodeInterpreterSes
 func (m *awsRestjson1_serializeOpGetCodeInterpreterSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -445,7 +430,6 @@ func (m *awsRestjson1_serializeOpGetCodeInterpreterSes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCodeInterpreterSessionInput(v *GetCodeInterpreterSessionInput, encoder *httpbinding.Encoder) error {
@@ -479,10 +463,8 @@ func (m *awsRestjson1_serializeOpGetEvent) HandleSeria
 func (m *awsRestjson1_serializeOpGetEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -520,7 +502,6 @@ func (m *awsRestjson1_serializeOpGetEvent) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventInput(v *GetEventInput, encoder *httpbinding.Encoder) error {
@@ -577,10 +558,8 @@ func (m *awsRestjson1_serializeOpGetMemoryRecord) Hand
 func (m *awsRestjson1_serializeOpGetMemoryRecord) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -618,7 +597,6 @@ func (m *awsRestjson1_serializeOpGetMemoryRecord) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMemoryRecordInput(v *GetMemoryRecordInput, encoder *httpbinding.Encoder) error {
@@ -657,10 +635,8 @@ func (m *awsRestjson1_serializeOpGetResourceApiKey) Ha
 func (m *awsRestjson1_serializeOpGetResourceApiKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -705,7 +681,6 @@ func (m *awsRestjson1_serializeOpGetResourceApiKey) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceApiKeyInput(v *GetResourceApiKeyInput, encoder *httpbinding.Encoder) error {
@@ -743,10 +718,8 @@ func (m *awsRestjson1_serializeOpGetResourceOauth2Toke
 func (m *awsRestjson1_serializeOpGetResourceOauth2Token) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -791,7 +764,6 @@ func (m *awsRestjson1_serializeOpGetResourceOauth2Toke
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceOauth2TokenInput(v *GetResourceOauth2TokenInput, encoder *httpbinding.Encoder) error {
@@ -863,10 +835,8 @@ func (m *awsRestjson1_serializeOpGetWorkloadAccessToke
 func (m *awsRestjson1_serializeOpGetWorkloadAccessToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -911,7 +881,6 @@ func (m *awsRestjson1_serializeOpGetWorkloadAccessToke
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkloadAccessTokenInput(v *GetWorkloadAccessTokenInput, encoder *httpbinding.Encoder) error {
@@ -944,10 +913,8 @@ func (m *awsRestjson1_serializeOpGetWorkloadAccessToke
 func (m *awsRestjson1_serializeOpGetWorkloadAccessTokenForJWT) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -992,7 +959,6 @@ func (m *awsRestjson1_serializeOpGetWorkloadAccessToke
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkloadAccessTokenForJWTInput(v *GetWorkloadAccessTokenForJWTInput, encoder *httpbinding.Encoder) error {
@@ -1030,10 +996,8 @@ func (m *awsRestjson1_serializeOpGetWorkloadAccessToke
 func (m *awsRestjson1_serializeOpGetWorkloadAccessTokenForUserId) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1078,7 +1042,6 @@ func (m *awsRestjson1_serializeOpGetWorkloadAccessToke
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkloadAccessTokenForUserIdInput(v *GetWorkloadAccessTokenForUserIdInput, encoder *httpbinding.Encoder) error {
@@ -1116,10 +1079,8 @@ func (m *awsRestjson1_serializeOpInvokeAgentRuntime) H
 func (m *awsRestjson1_serializeOpInvokeAgentRuntime) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1169,7 +1130,6 @@ func (m *awsRestjson1_serializeOpInvokeAgentRuntime) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeAgentRuntimeInput(v *InvokeAgentRuntimeInput, encoder *httpbinding.Encoder) error {
@@ -1253,10 +1213,8 @@ func (m *awsRestjson1_serializeOpInvokeCodeInterpreter
 func (m *awsRestjson1_serializeOpInvokeCodeInterpreter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1305,7 +1263,6 @@ func (m *awsRestjson1_serializeOpInvokeCodeInterpreter
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeCodeInterpreterInput(v *InvokeCodeInterpreterInput, encoder *httpbinding.Encoder) error {
@@ -1359,10 +1316,8 @@ func (m *awsRestjson1_serializeOpListActors) HandleSer
 func (m *awsRestjson1_serializeOpListActors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1411,7 +1366,6 @@ func (m *awsRestjson1_serializeOpListActors) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListActorsInput(v *ListActorsInput, encoder *httpbinding.Encoder) error {
@@ -1458,10 +1412,8 @@ func (m *awsRestjson1_serializeOpListBrowserSessions) 
 func (m *awsRestjson1_serializeOpListBrowserSessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1510,7 +1462,6 @@ func (m *awsRestjson1_serializeOpListBrowserSessions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBrowserSessionsInput(v *ListBrowserSessionsInput, encoder *httpbinding.Encoder) error {
@@ -1562,10 +1513,8 @@ func (m *awsRestjson1_serializeOpListCodeInterpreterSe
 func (m *awsRestjson1_serializeOpListCodeInterpreterSessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1614,7 +1563,6 @@ func (m *awsRestjson1_serializeOpListCodeInterpreterSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCodeInterpreterSessionsInput(v *ListCodeInterpreterSessionsInput, encoder *httpbinding.Encoder) error {
@@ -1666,10 +1614,8 @@ func (m *awsRestjson1_serializeOpListEvents) HandleSer
 func (m *awsRestjson1_serializeOpListEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1718,7 +1664,6 @@ func (m *awsRestjson1_serializeOpListEvents) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventsInput(v *ListEventsInput, encoder *httpbinding.Encoder) error {
@@ -1795,10 +1740,8 @@ func (m *awsRestjson1_serializeOpListMemoryRecords) Ha
 func (m *awsRestjson1_serializeOpListMemoryRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1847,7 +1790,6 @@ func (m *awsRestjson1_serializeOpListMemoryRecords) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMemoryRecordsInput(v *ListMemoryRecordsInput, encoder *httpbinding.Encoder) error {
@@ -1904,10 +1846,8 @@ func (m *awsRestjson1_serializeOpListSessions) HandleS
 func (m *awsRestjson1_serializeOpListSessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1956,7 +1896,6 @@ func (m *awsRestjson1_serializeOpListSessions) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSessionsInput(v *ListSessionsInput, encoder *httpbinding.Encoder) error {
@@ -2012,10 +1951,8 @@ func (m *awsRestjson1_serializeOpRetrieveMemoryRecords
 func (m *awsRestjson1_serializeOpRetrieveMemoryRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2064,7 +2001,6 @@ func (m *awsRestjson1_serializeOpRetrieveMemoryRecords
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRetrieveMemoryRecordsInput(v *RetrieveMemoryRecordsInput, encoder *httpbinding.Encoder) error {
@@ -2123,10 +2059,8 @@ func (m *awsRestjson1_serializeOpStartBrowserSession) 
 func (m *awsRestjson1_serializeOpStartBrowserSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2175,7 +2109,6 @@ func (m *awsRestjson1_serializeOpStartBrowserSession) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartBrowserSessionInput(v *StartBrowserSessionInput, encoder *httpbinding.Encoder) error {
@@ -2234,10 +2167,8 @@ func (m *awsRestjson1_serializeOpStartCodeInterpreterS
 func (m *awsRestjson1_serializeOpStartCodeInterpreterSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2286,7 +2217,6 @@ func (m *awsRestjson1_serializeOpStartCodeInterpreterS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartCodeInterpreterSessionInput(v *StartCodeInterpreterSessionInput, encoder *httpbinding.Encoder) error {
@@ -2338,10 +2268,8 @@ func (m *awsRestjson1_serializeOpStopBrowserSession) H
 func (m *awsRestjson1_serializeOpStopBrowserSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2390,7 +2318,6 @@ func (m *awsRestjson1_serializeOpStopBrowserSession) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopBrowserSessionInput(v *StopBrowserSessionInput, encoder *httpbinding.Encoder) error {
@@ -2436,10 +2363,8 @@ func (m *awsRestjson1_serializeOpStopCodeInterpreterSe
 func (m *awsRestjson1_serializeOpStopCodeInterpreterSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2488,7 +2413,6 @@ func (m *awsRestjson1_serializeOpStopCodeInterpreterSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopCodeInterpreterSessionInput(v *StopCodeInterpreterSessionInput, encoder *httpbinding.Encoder) error {
@@ -2534,10 +2458,8 @@ func (m *awsRestjson1_serializeOpUpdateBrowserStream) 
 func (m *awsRestjson1_serializeOpUpdateBrowserStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2586,7 +2508,6 @@ func (m *awsRestjson1_serializeOpUpdateBrowserStream) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBrowserStreamInput(v *UpdateBrowserStreamInput, encoder *httpbinding.Encoder) error {
