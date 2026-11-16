--- vendor/github.com/aws/aws-sdk-go-v2/service/lambda/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lambda/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAddLayerVersionPermis
 func (m *awsRestjson1_serializeOpAddLayerVersionPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpAddLayerVersionPermis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddLayerVersionPermissionInput(v *AddLayerVersionPermissionInput, encoder *httpbinding.Encoder) error {
@@ -149,10 +145,8 @@ func (m *awsRestjson1_serializeOpAddPermission) Handle
 func (m *awsRestjson1_serializeOpAddPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -201,7 +195,6 @@ func (m *awsRestjson1_serializeOpAddPermission) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddPermissionInput(v *AddPermissionInput, encoder *httpbinding.Encoder) error {
@@ -287,10 +280,8 @@ func (m *awsRestjson1_serializeOpCreateAlias) HandleSe
 func (m *awsRestjson1_serializeOpCreateAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -339,7 +330,6 @@ func (m *awsRestjson1_serializeOpCreateAlias) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAliasInput(v *CreateAliasInput, encoder *httpbinding.Encoder) error {
@@ -398,10 +388,8 @@ func (m *awsRestjson1_serializeOpCreateCodeSigningConf
 func (m *awsRestjson1_serializeOpCreateCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -446,7 +434,6 @@ func (m *awsRestjson1_serializeOpCreateCodeSigningConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCodeSigningConfigInput(v *CreateCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -500,10 +487,8 @@ func (m *awsRestjson1_serializeOpCreateEventSourceMapp
 func (m *awsRestjson1_serializeOpCreateEventSourceMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -548,7 +533,6 @@ func (m *awsRestjson1_serializeOpCreateEventSourceMapp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEventSourceMappingInput(v *CreateEventSourceMappingInput, encoder *httpbinding.Encoder) error {
@@ -739,10 +723,8 @@ func (m *awsRestjson1_serializeOpCreateFunction) Handl
 func (m *awsRestjson1_serializeOpCreateFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -787,7 +769,6 @@ func (m *awsRestjson1_serializeOpCreateFunction) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFunctionInput(v *CreateFunctionInput, encoder *httpbinding.Encoder) error {
@@ -961,10 +942,8 @@ func (m *awsRestjson1_serializeOpCreateFunctionUrlConf
 func (m *awsRestjson1_serializeOpCreateFunctionUrlConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1013,7 +992,6 @@ func (m *awsRestjson1_serializeOpCreateFunctionUrlConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFunctionUrlConfigInput(v *CreateFunctionUrlConfigInput, encoder *httpbinding.Encoder) error {
@@ -1071,10 +1049,8 @@ func (m *awsRestjson1_serializeOpDeleteAlias) HandleSe
 func (m *awsRestjson1_serializeOpDeleteAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1112,7 +1088,6 @@ func (m *awsRestjson1_serializeOpDeleteAlias) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAliasInput(v *DeleteAliasInput, encoder *httpbinding.Encoder) error {
@@ -1151,10 +1126,8 @@ func (m *awsRestjson1_serializeOpDeleteCodeSigningConf
 func (m *awsRestjson1_serializeOpDeleteCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1192,7 +1165,6 @@ func (m *awsRestjson1_serializeOpDeleteCodeSigningConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCodeSigningConfigInput(v *DeleteCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -1222,10 +1194,8 @@ func (m *awsRestjson1_serializeOpDeleteEventSourceMapp
 func (m *awsRestjson1_serializeOpDeleteEventSourceMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1263,7 +1233,6 @@ func (m *awsRestjson1_serializeOpDeleteEventSourceMapp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventSourceMappingInput(v *DeleteEventSourceMappingInput, encoder *httpbinding.Encoder) error {
@@ -1293,10 +1262,8 @@ func (m *awsRestjson1_serializeOpDeleteFunction) Handl
 func (m *awsRestjson1_serializeOpDeleteFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1334,7 +1301,6 @@ func (m *awsRestjson1_serializeOpDeleteFunction) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFunctionInput(v *DeleteFunctionInput, encoder *httpbinding.Encoder) error {
@@ -1368,10 +1334,8 @@ func (m *awsRestjson1_serializeOpDeleteFunctionCodeSig
 func (m *awsRestjson1_serializeOpDeleteFunctionCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1409,7 +1373,6 @@ func (m *awsRestjson1_serializeOpDeleteFunctionCodeSig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFunctionCodeSigningConfigInput(v *DeleteFunctionCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -1439,10 +1402,8 @@ func (m *awsRestjson1_serializeOpDeleteFunctionConcurr
 func (m *awsRestjson1_serializeOpDeleteFunctionConcurrency) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1480,7 +1441,6 @@ func (m *awsRestjson1_serializeOpDeleteFunctionConcurr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFunctionConcurrencyInput(v *DeleteFunctionConcurrencyInput, encoder *httpbinding.Encoder) error {
@@ -1510,10 +1470,8 @@ func (m *awsRestjson1_serializeOpDeleteFunctionEventIn
 func (m *awsRestjson1_serializeOpDeleteFunctionEventInvokeConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1551,7 +1509,6 @@ func (m *awsRestjson1_serializeOpDeleteFunctionEventIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFunctionEventInvokeConfigInput(v *DeleteFunctionEventInvokeConfigInput, encoder *httpbinding.Encoder) error {
@@ -1585,10 +1542,8 @@ func (m *awsRestjson1_serializeOpDeleteFunctionUrlConf
 func (m *awsRestjson1_serializeOpDeleteFunctionUrlConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1626,7 +1581,6 @@ func (m *awsRestjson1_serializeOpDeleteFunctionUrlConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFunctionUrlConfigInput(v *DeleteFunctionUrlConfigInput, encoder *httpbinding.Encoder) error {
@@ -1660,10 +1614,8 @@ func (m *awsRestjson1_serializeOpDeleteLayerVersion) H
 func (m *awsRestjson1_serializeOpDeleteLayerVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1701,7 +1653,6 @@ func (m *awsRestjson1_serializeOpDeleteLayerVersion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLayerVersionInput(v *DeleteLayerVersionInput, encoder *httpbinding.Encoder) error {
@@ -1740,10 +1691,8 @@ func (m *awsRestjson1_serializeOpDeleteProvisionedConc
 func (m *awsRestjson1_serializeOpDeleteProvisionedConcurrencyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1781,7 +1730,6 @@ func (m *awsRestjson1_serializeOpDeleteProvisionedConc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProvisionedConcurrencyConfigInput(v *DeleteProvisionedConcurrencyConfigInput, encoder *httpbinding.Encoder) error {
@@ -1815,10 +1763,8 @@ func (m *awsRestjson1_serializeOpGetAccountSettings) H
 func (m *awsRestjson1_serializeOpGetAccountSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1852,7 +1798,6 @@ func (m *awsRestjson1_serializeOpGetAccountSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountSettingsInput(v *GetAccountSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1873,10 +1818,8 @@ func (m *awsRestjson1_serializeOpGetAlias) HandleSeria
 func (m *awsRestjson1_serializeOpGetAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1914,7 +1857,6 @@ func (m *awsRestjson1_serializeOpGetAlias) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAliasInput(v *GetAliasInput, encoder *httpbinding.Encoder) error {
@@ -1953,10 +1895,8 @@ func (m *awsRestjson1_serializeOpGetCodeSigningConfig)
 func (m *awsRestjson1_serializeOpGetCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1994,7 +1934,6 @@ func (m *awsRestjson1_serializeOpGetCodeSigningConfig)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCodeSigningConfigInput(v *GetCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -2024,10 +1963,8 @@ func (m *awsRestjson1_serializeOpGetEventSourceMapping
 func (m *awsRestjson1_serializeOpGetEventSourceMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2065,7 +2002,6 @@ func (m *awsRestjson1_serializeOpGetEventSourceMapping
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventSourceMappingInput(v *GetEventSourceMappingInput, encoder *httpbinding.Encoder) error {
@@ -2095,10 +2031,8 @@ func (m *awsRestjson1_serializeOpGetFunction) HandleSe
 func (m *awsRestjson1_serializeOpGetFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2136,7 +2070,6 @@ func (m *awsRestjson1_serializeOpGetFunction) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionInput(v *GetFunctionInput, encoder *httpbinding.Encoder) error {
@@ -2170,10 +2103,8 @@ func (m *awsRestjson1_serializeOpGetFunctionCodeSignin
 func (m *awsRestjson1_serializeOpGetFunctionCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2211,7 +2142,6 @@ func (m *awsRestjson1_serializeOpGetFunctionCodeSignin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionCodeSigningConfigInput(v *GetFunctionCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -2241,10 +2171,8 @@ func (m *awsRestjson1_serializeOpGetFunctionConcurrenc
 func (m *awsRestjson1_serializeOpGetFunctionConcurrency) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2282,7 +2210,6 @@ func (m *awsRestjson1_serializeOpGetFunctionConcurrenc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionConcurrencyInput(v *GetFunctionConcurrencyInput, encoder *httpbinding.Encoder) error {
@@ -2312,10 +2239,8 @@ func (m *awsRestjson1_serializeOpGetFunctionConfigurat
 func (m *awsRestjson1_serializeOpGetFunctionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2353,7 +2278,6 @@ func (m *awsRestjson1_serializeOpGetFunctionConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionConfigurationInput(v *GetFunctionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2387,10 +2311,8 @@ func (m *awsRestjson1_serializeOpGetFunctionEventInvok
 func (m *awsRestjson1_serializeOpGetFunctionEventInvokeConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2428,7 +2350,6 @@ func (m *awsRestjson1_serializeOpGetFunctionEventInvok
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionEventInvokeConfigInput(v *GetFunctionEventInvokeConfigInput, encoder *httpbinding.Encoder) error {
@@ -2462,10 +2383,8 @@ func (m *awsRestjson1_serializeOpGetFunctionRecursionC
 func (m *awsRestjson1_serializeOpGetFunctionRecursionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2503,7 +2422,6 @@ func (m *awsRestjson1_serializeOpGetFunctionRecursionC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionRecursionConfigInput(v *GetFunctionRecursionConfigInput, encoder *httpbinding.Encoder) error {
@@ -2533,10 +2451,8 @@ func (m *awsRestjson1_serializeOpGetFunctionUrlConfig)
 func (m *awsRestjson1_serializeOpGetFunctionUrlConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2574,7 +2490,6 @@ func (m *awsRestjson1_serializeOpGetFunctionUrlConfig)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionUrlConfigInput(v *GetFunctionUrlConfigInput, encoder *httpbinding.Encoder) error {
@@ -2608,10 +2523,8 @@ func (m *awsRestjson1_serializeOpGetLayerVersion) Hand
 func (m *awsRestjson1_serializeOpGetLayerVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2649,7 +2562,6 @@ func (m *awsRestjson1_serializeOpGetLayerVersion) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLayerVersionInput(v *GetLayerVersionInput, encoder *httpbinding.Encoder) error {
@@ -2688,10 +2600,8 @@ func (m *awsRestjson1_serializeOpGetLayerVersionByArn)
 func (m *awsRestjson1_serializeOpGetLayerVersionByArn) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2729,7 +2639,6 @@ func (m *awsRestjson1_serializeOpGetLayerVersionByArn)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLayerVersionByArnInput(v *GetLayerVersionByArnInput, encoder *httpbinding.Encoder) error {
@@ -2754,10 +2663,8 @@ func (m *awsRestjson1_serializeOpGetLayerVersionPolicy
 func (m *awsRestjson1_serializeOpGetLayerVersionPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2795,7 +2702,6 @@ func (m *awsRestjson1_serializeOpGetLayerVersionPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLayerVersionPolicyInput(v *GetLayerVersionPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2834,10 +2740,8 @@ func (m *awsRestjson1_serializeOpGetPolicy) HandleSeri
 func (m *awsRestjson1_serializeOpGetPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2875,7 +2779,6 @@ func (m *awsRestjson1_serializeOpGetPolicy) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPolicyInput(v *GetPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2909,10 +2812,8 @@ func (m *awsRestjson1_serializeOpGetProvisionedConcurr
 func (m *awsRestjson1_serializeOpGetProvisionedConcurrencyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2950,7 +2851,6 @@ func (m *awsRestjson1_serializeOpGetProvisionedConcurr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProvisionedConcurrencyConfigInput(v *GetProvisionedConcurrencyConfigInput, encoder *httpbinding.Encoder) error {
@@ -2984,10 +2884,8 @@ func (m *awsRestjson1_serializeOpGetRuntimeManagementC
 func (m *awsRestjson1_serializeOpGetRuntimeManagementConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3025,7 +2923,6 @@ func (m *awsRestjson1_serializeOpGetRuntimeManagementC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRuntimeManagementConfigInput(v *GetRuntimeManagementConfigInput, encoder *httpbinding.Encoder) error {
@@ -3059,10 +2956,8 @@ func (m *awsRestjson1_serializeOpInvoke) HandleSeriali
 func (m *awsRestjson1_serializeOpInvoke) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3112,7 +3007,6 @@ func (m *awsRestjson1_serializeOpInvoke) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeInput(v *InvokeInput, encoder *httpbinding.Encoder) error {
@@ -3161,10 +3055,8 @@ func (m *awsRestjson1_serializeOpInvokeAsync) HandleSe
 func (m *awsRestjson1_serializeOpInvokeAsync) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3214,7 +3106,6 @@ func (m *awsRestjson1_serializeOpInvokeAsync) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeAsyncInput(v *InvokeAsyncInput, encoder *httpbinding.Encoder) error {
@@ -3244,10 +3135,8 @@ func (m *awsRestjson1_serializeOpInvokeWithResponseStr
 func (m *awsRestjson1_serializeOpInvokeWithResponseStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3297,7 +3186,6 @@ func (m *awsRestjson1_serializeOpInvokeWithResponseStr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInvokeWithResponseStreamInput(v *InvokeWithResponseStreamInput, encoder *httpbinding.Encoder) error {
@@ -3346,10 +3234,8 @@ func (m *awsRestjson1_serializeOpListAliases) HandleSe
 func (m *awsRestjson1_serializeOpListAliases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3387,7 +3273,6 @@ func (m *awsRestjson1_serializeOpListAliases) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAliasesInput(v *ListAliasesInput, encoder *httpbinding.Encoder) error {
@@ -3429,10 +3314,8 @@ func (m *awsRestjson1_serializeOpListCodeSigningConfig
 func (m *awsRestjson1_serializeOpListCodeSigningConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3470,7 +3353,6 @@ func (m *awsRestjson1_serializeOpListCodeSigningConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCodeSigningConfigsInput(v *ListCodeSigningConfigsInput, encoder *httpbinding.Encoder) error {
@@ -3499,10 +3381,8 @@ func (m *awsRestjson1_serializeOpListEventSourceMappin
 func (m *awsRestjson1_serializeOpListEventSourceMappings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3540,7 +3420,6 @@ func (m *awsRestjson1_serializeOpListEventSourceMappin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventSourceMappingsInput(v *ListEventSourceMappingsInput, encoder *httpbinding.Encoder) error {
@@ -3577,10 +3456,8 @@ func (m *awsRestjson1_serializeOpListFunctionEventInvo
 func (m *awsRestjson1_serializeOpListFunctionEventInvokeConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3618,7 +3495,6 @@ func (m *awsRestjson1_serializeOpListFunctionEventInvo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFunctionEventInvokeConfigsInput(v *ListFunctionEventInvokeConfigsInput, encoder *httpbinding.Encoder) error {
@@ -3656,10 +3532,8 @@ func (m *awsRestjson1_serializeOpListFunctions) Handle
 func (m *awsRestjson1_serializeOpListFunctions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3697,7 +3571,6 @@ func (m *awsRestjson1_serializeOpListFunctions) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFunctionsInput(v *ListFunctionsInput, encoder *httpbinding.Encoder) error {
@@ -3734,10 +3607,8 @@ func (m *awsRestjson1_serializeOpListFunctionsByCodeSi
 func (m *awsRestjson1_serializeOpListFunctionsByCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3775,7 +3646,6 @@ func (m *awsRestjson1_serializeOpListFunctionsByCodeSi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFunctionsByCodeSigningConfigInput(v *ListFunctionsByCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -3813,10 +3683,8 @@ func (m *awsRestjson1_serializeOpListFunctionUrlConfig
 func (m *awsRestjson1_serializeOpListFunctionUrlConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3854,7 +3722,6 @@ func (m *awsRestjson1_serializeOpListFunctionUrlConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFunctionUrlConfigsInput(v *ListFunctionUrlConfigsInput, encoder *httpbinding.Encoder) error {
@@ -3892,10 +3759,8 @@ func (m *awsRestjson1_serializeOpListLayers) HandleSer
 func (m *awsRestjson1_serializeOpListLayers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3933,7 +3798,6 @@ func (m *awsRestjson1_serializeOpListLayers) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLayersInput(v *ListLayersInput, encoder *httpbinding.Encoder) error {
@@ -3970,10 +3834,8 @@ func (m *awsRestjson1_serializeOpListLayerVersions) Ha
 func (m *awsRestjson1_serializeOpListLayerVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4011,7 +3873,6 @@ func (m *awsRestjson1_serializeOpListLayerVersions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLayerVersionsInput(v *ListLayerVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4057,10 +3918,8 @@ func (m *awsRestjson1_serializeOpListProvisionedConcur
 func (m *awsRestjson1_serializeOpListProvisionedConcurrencyConfigs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4098,7 +3957,6 @@ func (m *awsRestjson1_serializeOpListProvisionedConcur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProvisionedConcurrencyConfigsInput(v *ListProvisionedConcurrencyConfigsInput, encoder *httpbinding.Encoder) error {
@@ -4136,10 +3994,8 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
 func (m *awsRestjson1_serializeOpListTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4177,7 +4033,6 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsInput(v *ListTagsInput, encoder *httpbinding.Encoder) error {
@@ -4207,10 +4062,8 @@ func (m *awsRestjson1_serializeOpListVersionsByFunctio
 func (m *awsRestjson1_serializeOpListVersionsByFunction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4248,7 +4101,6 @@ func (m *awsRestjson1_serializeOpListVersionsByFunctio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVersionsByFunctionInput(v *ListVersionsByFunctionInput, encoder *httpbinding.Encoder) error {
@@ -4286,10 +4138,8 @@ func (m *awsRestjson1_serializeOpPublishLayerVersion) 
 func (m *awsRestjson1_serializeOpPublishLayerVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4338,7 +4188,6 @@ func (m *awsRestjson1_serializeOpPublishLayerVersion) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPublishLayerVersionInput(v *PublishLayerVersionInput, encoder *httpbinding.Encoder) error {
@@ -4406,10 +4255,8 @@ func (m *awsRestjson1_serializeOpPublishVersion) Handl
 func (m *awsRestjson1_serializeOpPublishVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4458,7 +4305,6 @@ func (m *awsRestjson1_serializeOpPublishVersion) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPublishVersionInput(v *PublishVersionInput, encoder *httpbinding.Encoder) error {
@@ -4510,10 +4356,8 @@ func (m *awsRestjson1_serializeOpPutFunctionCodeSignin
 func (m *awsRestjson1_serializeOpPutFunctionCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4562,7 +4406,6 @@ func (m *awsRestjson1_serializeOpPutFunctionCodeSignin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutFunctionCodeSigningConfigInput(v *PutFunctionCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -4604,10 +4447,8 @@ func (m *awsRestjson1_serializeOpPutFunctionConcurrenc
 func (m *awsRestjson1_serializeOpPutFunctionConcurrency) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4656,7 +4497,6 @@ func (m *awsRestjson1_serializeOpPutFunctionConcurrenc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutFunctionConcurrencyInput(v *PutFunctionConcurrencyInput, encoder *httpbinding.Encoder) error {
@@ -4698,10 +4538,8 @@ func (m *awsRestjson1_serializeOpPutFunctionEventInvok
 func (m *awsRestjson1_serializeOpPutFunctionEventInvokeConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4750,7 +4588,6 @@ func (m *awsRestjson1_serializeOpPutFunctionEventInvok
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutFunctionEventInvokeConfigInput(v *PutFunctionEventInvokeConfigInput, encoder *httpbinding.Encoder) error {
@@ -4808,10 +4645,8 @@ func (m *awsRestjson1_serializeOpPutFunctionRecursionC
 func (m *awsRestjson1_serializeOpPutFunctionRecursionConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4860,7 +4695,6 @@ func (m *awsRestjson1_serializeOpPutFunctionRecursionC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutFunctionRecursionConfigInput(v *PutFunctionRecursionConfigInput, encoder *httpbinding.Encoder) error {
@@ -4902,10 +4736,8 @@ func (m *awsRestjson1_serializeOpPutProvisionedConcurr
 func (m *awsRestjson1_serializeOpPutProvisionedConcurrencyConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4954,7 +4786,6 @@ func (m *awsRestjson1_serializeOpPutProvisionedConcurr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutProvisionedConcurrencyConfigInput(v *PutProvisionedConcurrencyConfigInput, encoder *httpbinding.Encoder) error {
@@ -5000,10 +4831,8 @@ func (m *awsRestjson1_serializeOpPutRuntimeManagementC
 func (m *awsRestjson1_serializeOpPutRuntimeManagementConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5052,7 +4881,6 @@ func (m *awsRestjson1_serializeOpPutRuntimeManagementC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutRuntimeManagementConfigInput(v *PutRuntimeManagementConfigInput, encoder *httpbinding.Encoder) error {
@@ -5103,10 +4931,8 @@ func (m *awsRestjson1_serializeOpRemoveLayerVersionPer
 func (m *awsRestjson1_serializeOpRemoveLayerVersionPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5144,7 +4970,6 @@ func (m *awsRestjson1_serializeOpRemoveLayerVersionPer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveLayerVersionPermissionInput(v *RemoveLayerVersionPermissionInput, encoder *httpbinding.Encoder) error {
@@ -5196,10 +5021,8 @@ func (m *awsRestjson1_serializeOpRemovePermission) Han
 func (m *awsRestjson1_serializeOpRemovePermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5237,7 +5060,6 @@ func (m *awsRestjson1_serializeOpRemovePermission) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemovePermissionInput(v *RemovePermissionInput, encoder *httpbinding.Encoder) error {
@@ -5284,10 +5106,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5336,7 +5156,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5380,10 +5199,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5421,7 +5238,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5457,10 +5273,8 @@ func (m *awsRestjson1_serializeOpUpdateAlias) HandleSe
 func (m *awsRestjson1_serializeOpUpdateAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5509,7 +5323,6 @@ func (m *awsRestjson1_serializeOpUpdateAlias) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAliasInput(v *UpdateAliasInput, encoder *httpbinding.Encoder) error {
@@ -5577,10 +5390,8 @@ func (m *awsRestjson1_serializeOpUpdateCodeSigningConf
 func (m *awsRestjson1_serializeOpUpdateCodeSigningConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5629,7 +5440,6 @@ func (m *awsRestjson1_serializeOpUpdateCodeSigningConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCodeSigningConfigInput(v *UpdateCodeSigningConfigInput, encoder *httpbinding.Encoder) error {
@@ -5685,10 +5495,8 @@ func (m *awsRestjson1_serializeOpUpdateEventSourceMapp
 func (m *awsRestjson1_serializeOpUpdateEventSourceMapping) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5737,7 +5545,6 @@ func (m *awsRestjson1_serializeOpUpdateEventSourceMapp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventSourceMappingInput(v *UpdateEventSourceMappingInput, encoder *httpbinding.Encoder) error {
@@ -5894,10 +5701,8 @@ func (m *awsRestjson1_serializeOpUpdateFunctionCode) H
 func (m *awsRestjson1_serializeOpUpdateFunctionCode) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5946,7 +5751,6 @@ func (m *awsRestjson1_serializeOpUpdateFunctionCode) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFunctionCodeInput(v *UpdateFunctionCodeInput, encoder *httpbinding.Encoder) error {
@@ -6035,10 +5839,8 @@ func (m *awsRestjson1_serializeOpUpdateFunctionConfigu
 func (m *awsRestjson1_serializeOpUpdateFunctionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6087,7 +5889,6 @@ func (m *awsRestjson1_serializeOpUpdateFunctionConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFunctionConfigurationInput(v *UpdateFunctionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6234,10 +6035,8 @@ func (m *awsRestjson1_serializeOpUpdateFunctionEventIn
 func (m *awsRestjson1_serializeOpUpdateFunctionEventInvokeConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6286,7 +6085,6 @@ func (m *awsRestjson1_serializeOpUpdateFunctionEventIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFunctionEventInvokeConfigInput(v *UpdateFunctionEventInvokeConfigInput, encoder *httpbinding.Encoder) error {
@@ -6344,10 +6142,8 @@ func (m *awsRestjson1_serializeOpUpdateFunctionUrlConf
 func (m *awsRestjson1_serializeOpUpdateFunctionUrlConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6396,7 +6192,6 @@ func (m *awsRestjson1_serializeOpUpdateFunctionUrlConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFunctionUrlConfigInput(v *UpdateFunctionUrlConfigInput, encoder *httpbinding.Encoder) error {
