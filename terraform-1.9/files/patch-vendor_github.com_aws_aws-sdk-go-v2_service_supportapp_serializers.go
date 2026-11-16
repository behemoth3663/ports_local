--- vendor/github.com/aws/aws-sdk-go-v2/service/supportapp/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/supportapp/serializers.go
@@ -10,7 +10,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -24,10 +23,8 @@ func (m *awsRestjson1_serializeOpCreateSlackChannelCon
 func (m *awsRestjson1_serializeOpCreateSlackChannelConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -72,7 +69,6 @@ func (m *awsRestjson1_serializeOpCreateSlackChannelCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSlackChannelConfigurationInput(v *CreateSlackChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -140,10 +136,8 @@ func (m *awsRestjson1_serializeOpDeleteAccountAlias) H
 func (m *awsRestjson1_serializeOpDeleteAccountAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -177,7 +171,6 @@ func (m *awsRestjson1_serializeOpDeleteAccountAlias) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccountAliasInput(v *DeleteAccountAliasInput, encoder *httpbinding.Encoder) error {
@@ -198,10 +191,8 @@ func (m *awsRestjson1_serializeOpDeleteSlackChannelCon
 func (m *awsRestjson1_serializeOpDeleteSlackChannelConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -246,7 +237,6 @@ func (m *awsRestjson1_serializeOpDeleteSlackChannelCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSlackChannelConfigurationInput(v *DeleteSlackChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -284,10 +274,8 @@ func (m *awsRestjson1_serializeOpDeleteSlackWorkspaceC
 func (m *awsRestjson1_serializeOpDeleteSlackWorkspaceConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -332,7 +320,6 @@ func (m *awsRestjson1_serializeOpDeleteSlackWorkspaceC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSlackWorkspaceConfigurationInput(v *DeleteSlackWorkspaceConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -365,10 +352,8 @@ func (m *awsRestjson1_serializeOpGetAccountAlias) Hand
 func (m *awsRestjson1_serializeOpGetAccountAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -402,7 +387,6 @@ func (m *awsRestjson1_serializeOpGetAccountAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountAliasInput(v *GetAccountAliasInput, encoder *httpbinding.Encoder) error {
@@ -423,10 +407,8 @@ func (m *awsRestjson1_serializeOpListSlackChannelConfi
 func (m *awsRestjson1_serializeOpListSlackChannelConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -471,7 +453,6 @@ func (m *awsRestjson1_serializeOpListSlackChannelConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSlackChannelConfigurationsInput(v *ListSlackChannelConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -504,10 +485,8 @@ func (m *awsRestjson1_serializeOpListSlackWorkspaceCon
 func (m *awsRestjson1_serializeOpListSlackWorkspaceConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -552,7 +531,6 @@ func (m *awsRestjson1_serializeOpListSlackWorkspaceCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSlackWorkspaceConfigurationsInput(v *ListSlackWorkspaceConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -585,10 +563,8 @@ func (m *awsRestjson1_serializeOpPutAccountAlias) Hand
 func (m *awsRestjson1_serializeOpPutAccountAlias) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -633,7 +609,6 @@ func (m *awsRestjson1_serializeOpPutAccountAlias) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountAliasInput(v *PutAccountAliasInput, encoder *httpbinding.Encoder) error {
@@ -666,10 +641,8 @@ func (m *awsRestjson1_serializeOpRegisterSlackWorkspac
 func (m *awsRestjson1_serializeOpRegisterSlackWorkspaceForOrganization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -714,7 +687,6 @@ func (m *awsRestjson1_serializeOpRegisterSlackWorkspac
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterSlackWorkspaceForOrganizationInput(v *RegisterSlackWorkspaceForOrganizationInput, encoder *httpbinding.Encoder) error {
@@ -747,10 +719,8 @@ func (m *awsRestjson1_serializeOpUpdateSlackChannelCon
 func (m *awsRestjson1_serializeOpUpdateSlackChannelConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -795,7 +765,6 @@ func (m *awsRestjson1_serializeOpUpdateSlackChannelCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSlackChannelConfigurationInput(v *UpdateSlackChannelConfigurationInput, encoder *httpbinding.Encoder) error {
