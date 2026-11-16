--- vendor/github.com/aws/aws-sdk-go-v2/service/chatbot/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chatbot/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociateToConfigurat
 func (m *awsRestjson1_serializeOpAssociateToConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -73,7 +70,6 @@ func (m *awsRestjson1_serializeOpAssociateToConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateToConfigurationInput(v *AssociateToConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -111,10 +107,8 @@ func (m *awsRestjson1_serializeOpCreateChimeWebhookCon
 func (m *awsRestjson1_serializeOpCreateChimeWebhookConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -159,7 +153,6 @@ func (m *awsRestjson1_serializeOpCreateChimeWebhookCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChimeWebhookConfigurationInput(v *CreateChimeWebhookConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -226,10 +219,8 @@ func (m *awsRestjson1_serializeOpCreateCustomAction) H
 func (m *awsRestjson1_serializeOpCreateCustomAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -274,7 +265,6 @@ func (m *awsRestjson1_serializeOpCreateCustomAction) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCustomActionInput(v *CreateCustomActionInput, encoder *httpbinding.Encoder) error {
@@ -338,10 +328,8 @@ func (m *awsRestjson1_serializeOpCreateMicrosoftTeamsC
 func (m *awsRestjson1_serializeOpCreateMicrosoftTeamsChannelConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -386,7 +374,6 @@ func (m *awsRestjson1_serializeOpCreateMicrosoftTeamsC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMicrosoftTeamsChannelConfigurationInput(v *CreateMicrosoftTeamsChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -480,10 +467,8 @@ func (m *awsRestjson1_serializeOpCreateSlackChannelCon
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
@@ -528,7 +513,6 @@ func (m *awsRestjson1_serializeOpCreateSlackChannelCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSlackChannelConfigurationInput(v *CreateSlackChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -612,10 +596,8 @@ func (m *awsRestjson1_serializeOpDeleteChimeWebhookCon
 func (m *awsRestjson1_serializeOpDeleteChimeWebhookConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -660,7 +642,6 @@ func (m *awsRestjson1_serializeOpDeleteChimeWebhookCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteChimeWebhookConfigurationInput(v *DeleteChimeWebhookConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -693,10 +674,8 @@ func (m *awsRestjson1_serializeOpDeleteCustomAction) H
 func (m *awsRestjson1_serializeOpDeleteCustomAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -741,7 +720,6 @@ func (m *awsRestjson1_serializeOpDeleteCustomAction) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCustomActionInput(v *DeleteCustomActionInput, encoder *httpbinding.Encoder) error {
@@ -774,10 +752,8 @@ func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsC
 func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsChannelConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -822,7 +798,6 @@ func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMicrosoftTeamsChannelConfigurationInput(v *DeleteMicrosoftTeamsChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -855,10 +830,8 @@ func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsC
 func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsConfiguredTeam) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -903,7 +876,6 @@ func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMicrosoftTeamsConfiguredTeamInput(v *DeleteMicrosoftTeamsConfiguredTeamInput, encoder *httpbinding.Encoder) error {
@@ -936,10 +908,8 @@ func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsU
 func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsUserIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -984,7 +954,6 @@ func (m *awsRestjson1_serializeOpDeleteMicrosoftTeamsU
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMicrosoftTeamsUserIdentityInput(v *DeleteMicrosoftTeamsUserIdentityInput, encoder *httpbinding.Encoder) error {
@@ -1022,10 +991,8 @@ func (m *awsRestjson1_serializeOpDeleteSlackChannelCon
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
@@ -1070,7 +1037,6 @@ func (m *awsRestjson1_serializeOpDeleteSlackChannelCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSlackChannelConfigurationInput(v *DeleteSlackChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1103,10 +1069,8 @@ func (m *awsRestjson1_serializeOpDeleteSlackUserIdenti
 func (m *awsRestjson1_serializeOpDeleteSlackUserIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1151,7 +1115,6 @@ func (m *awsRestjson1_serializeOpDeleteSlackUserIdenti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSlackUserIdentityInput(v *DeleteSlackUserIdentityInput, encoder *httpbinding.Encoder) error {
@@ -1194,10 +1157,8 @@ func (m *awsRestjson1_serializeOpDeleteSlackWorkspaceA
 func (m *awsRestjson1_serializeOpDeleteSlackWorkspaceAuthorization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1242,7 +1203,6 @@ func (m *awsRestjson1_serializeOpDeleteSlackWorkspaceA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSlackWorkspaceAuthorizationInput(v *DeleteSlackWorkspaceAuthorizationInput, encoder *httpbinding.Encoder) error {
@@ -1275,10 +1235,8 @@ func (m *awsRestjson1_serializeOpDescribeChimeWebhookC
 func (m *awsRestjson1_serializeOpDescribeChimeWebhookConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1323,7 +1281,6 @@ func (m *awsRestjson1_serializeOpDescribeChimeWebhookC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeChimeWebhookConfigurationsInput(v *DescribeChimeWebhookConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -1366,10 +1323,8 @@ func (m *awsRestjson1_serializeOpDescribeSlackChannelC
 func (m *awsRestjson1_serializeOpDescribeSlackChannelConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1414,7 +1369,6 @@ func (m *awsRestjson1_serializeOpDescribeSlackChannelC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSlackChannelConfigurationsInput(v *DescribeSlackChannelConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -1457,10 +1411,8 @@ func (m *awsRestjson1_serializeOpDescribeSlackUserIden
 func (m *awsRestjson1_serializeOpDescribeSlackUserIdentities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1505,7 +1457,6 @@ func (m *awsRestjson1_serializeOpDescribeSlackUserIden
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSlackUserIdentitiesInput(v *DescribeSlackUserIdentitiesInput, encoder *httpbinding.Encoder) error {
@@ -1548,10 +1499,8 @@ func (m *awsRestjson1_serializeOpDescribeSlackWorkspac
 func (m *awsRestjson1_serializeOpDescribeSlackWorkspaces) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1596,7 +1545,6 @@ func (m *awsRestjson1_serializeOpDescribeSlackWorkspac
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSlackWorkspacesInput(v *DescribeSlackWorkspacesInput, encoder *httpbinding.Encoder) error {
@@ -1634,10 +1582,8 @@ func (m *awsRestjson1_serializeOpDisassociateFromConfi
 func (m *awsRestjson1_serializeOpDisassociateFromConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1682,7 +1628,6 @@ func (m *awsRestjson1_serializeOpDisassociateFromConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFromConfigurationInput(v *DisassociateFromConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1720,10 +1665,8 @@ func (m *awsRestjson1_serializeOpGetAccountPreferences
 func (m *awsRestjson1_serializeOpGetAccountPreferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1757,7 +1700,6 @@ func (m *awsRestjson1_serializeOpGetAccountPreferences
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountPreferencesInput(v *GetAccountPreferencesInput, encoder *httpbinding.Encoder) error {
@@ -1778,10 +1720,8 @@ func (m *awsRestjson1_serializeOpGetCustomAction) Hand
 func (m *awsRestjson1_serializeOpGetCustomAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1826,7 +1766,6 @@ func (m *awsRestjson1_serializeOpGetCustomAction) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCustomActionInput(v *GetCustomActionInput, encoder *httpbinding.Encoder) error {
@@ -1859,10 +1798,8 @@ func (m *awsRestjson1_serializeOpGetMicrosoftTeamsChan
 func (m *awsRestjson1_serializeOpGetMicrosoftTeamsChannelConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1907,7 +1844,6 @@ func (m *awsRestjson1_serializeOpGetMicrosoftTeamsChan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMicrosoftTeamsChannelConfigurationInput(v *GetMicrosoftTeamsChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1940,10 +1876,8 @@ func (m *awsRestjson1_serializeOpListAssociations) Han
 func (m *awsRestjson1_serializeOpListAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1988,7 +1922,6 @@ func (m *awsRestjson1_serializeOpListAssociations) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssociationsInput(v *ListAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -2031,10 +1964,8 @@ func (m *awsRestjson1_serializeOpListCustomActions) Ha
 func (m *awsRestjson1_serializeOpListCustomActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2079,7 +2010,6 @@ func (m *awsRestjson1_serializeOpListCustomActions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCustomActionsInput(v *ListCustomActionsInput, encoder *httpbinding.Encoder) error {
@@ -2117,10 +2047,8 @@ func (m *awsRestjson1_serializeOpListMicrosoftTeamsCha
 func (m *awsRestjson1_serializeOpListMicrosoftTeamsChannelConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2165,7 +2093,6 @@ func (m *awsRestjson1_serializeOpListMicrosoftTeamsCha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMicrosoftTeamsChannelConfigurationsInput(v *ListMicrosoftTeamsChannelConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -2208,10 +2135,8 @@ func (m *awsRestjson1_serializeOpListMicrosoftTeamsCon
 func (m *awsRestjson1_serializeOpListMicrosoftTeamsConfiguredTeams) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2256,7 +2181,6 @@ func (m *awsRestjson1_serializeOpListMicrosoftTeamsCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMicrosoftTeamsConfiguredTeamsInput(v *ListMicrosoftTeamsConfiguredTeamsInput, encoder *httpbinding.Encoder) error {
@@ -2294,10 +2218,8 @@ func (m *awsRestjson1_serializeOpListMicrosoftTeamsUse
 func (m *awsRestjson1_serializeOpListMicrosoftTeamsUserIdentities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2342,7 +2264,6 @@ func (m *awsRestjson1_serializeOpListMicrosoftTeamsUse
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMicrosoftTeamsUserIdentitiesInput(v *ListMicrosoftTeamsUserIdentitiesInput, encoder *httpbinding.Encoder) error {
@@ -2385,10 +2306,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -2433,7 +2352,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -2466,10 +2384,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -2514,7 +2430,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2554,10 +2469,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -2602,7 +2515,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2642,10 +2554,8 @@ func (m *awsRestjson1_serializeOpUpdateAccountPreferen
 func (m *awsRestjson1_serializeOpUpdateAccountPreferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2690,7 +2600,6 @@ func (m *awsRestjson1_serializeOpUpdateAccountPreferen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountPreferencesInput(v *UpdateAccountPreferencesInput, encoder *httpbinding.Encoder) error {
@@ -2728,10 +2637,8 @@ func (m *awsRestjson1_serializeOpUpdateChimeWebhookCon
 func (m *awsRestjson1_serializeOpUpdateChimeWebhookConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2776,7 +2683,6 @@ func (m *awsRestjson1_serializeOpUpdateChimeWebhookCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChimeWebhookConfigurationInput(v *UpdateChimeWebhookConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2836,10 +2742,8 @@ func (m *awsRestjson1_serializeOpUpdateCustomAction) H
 func (m *awsRestjson1_serializeOpUpdateCustomAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2884,7 +2788,6 @@ func (m *awsRestjson1_serializeOpUpdateCustomAction) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCustomActionInput(v *UpdateCustomActionInput, encoder *httpbinding.Encoder) error {
@@ -2936,10 +2839,8 @@ func (m *awsRestjson1_serializeOpUpdateMicrosoftTeamsC
 func (m *awsRestjson1_serializeOpUpdateMicrosoftTeamsChannelConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2984,7 +2885,6 @@ func (m *awsRestjson1_serializeOpUpdateMicrosoftTeamsC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMicrosoftTeamsChannelConfigurationInput(v *UpdateMicrosoftTeamsChannelConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3056,10 +2956,8 @@ func (m *awsRestjson1_serializeOpUpdateSlackChannelCon
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
@@ -3104,7 +3002,6 @@ func (m *awsRestjson1_serializeOpUpdateSlackChannelCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSlackChannelConfigurationInput(v *UpdateSlackChannelConfigurationInput, encoder *httpbinding.Encoder) error {
