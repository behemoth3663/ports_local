--- vendor/github.com/aws/aws-sdk-go-v2/service/workspacesweb/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/workspacesweb/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociateBrowserSetti
 func (m *awsRestjson1_serializeOpAssociateBrowserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpAssociateBrowserSetti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateBrowserSettingsInput(v *AssociateBrowserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -100,10 +96,8 @@ func (m *awsRestjson1_serializeOpAssociateDataProtecti
 func (m *awsRestjson1_serializeOpAssociateDataProtectionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -141,7 +135,6 @@ func (m *awsRestjson1_serializeOpAssociateDataProtecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateDataProtectionSettingsInput(v *AssociateDataProtectionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -175,10 +168,8 @@ func (m *awsRestjson1_serializeOpAssociateIpAccessSett
 func (m *awsRestjson1_serializeOpAssociateIpAccessSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -216,7 +207,6 @@ func (m *awsRestjson1_serializeOpAssociateIpAccessSett
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateIpAccessSettingsInput(v *AssociateIpAccessSettingsInput, encoder *httpbinding.Encoder) error {
@@ -250,10 +240,8 @@ func (m *awsRestjson1_serializeOpAssociateNetworkSetti
 func (m *awsRestjson1_serializeOpAssociateNetworkSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -291,7 +279,6 @@ func (m *awsRestjson1_serializeOpAssociateNetworkSetti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateNetworkSettingsInput(v *AssociateNetworkSettingsInput, encoder *httpbinding.Encoder) error {
@@ -325,10 +312,8 @@ func (m *awsRestjson1_serializeOpAssociateTrustStore) 
 func (m *awsRestjson1_serializeOpAssociateTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -366,7 +351,6 @@ func (m *awsRestjson1_serializeOpAssociateTrustStore) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateTrustStoreInput(v *AssociateTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -400,10 +384,8 @@ func (m *awsRestjson1_serializeOpAssociateUserAccessLo
 func (m *awsRestjson1_serializeOpAssociateUserAccessLoggingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -441,7 +423,6 @@ func (m *awsRestjson1_serializeOpAssociateUserAccessLo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateUserAccessLoggingSettingsInput(v *AssociateUserAccessLoggingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -475,10 +456,8 @@ func (m *awsRestjson1_serializeOpAssociateUserSettings
 func (m *awsRestjson1_serializeOpAssociateUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -516,7 +495,6 @@ func (m *awsRestjson1_serializeOpAssociateUserSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateUserSettingsInput(v *AssociateUserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -550,10 +528,8 @@ func (m *awsRestjson1_serializeOpCreateBrowserSettings
 func (m *awsRestjson1_serializeOpCreateBrowserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -598,7 +574,6 @@ func (m *awsRestjson1_serializeOpCreateBrowserSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBrowserSettingsInput(v *CreateBrowserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -655,10 +630,8 @@ func (m *awsRestjson1_serializeOpCreateDataProtectionS
 func (m *awsRestjson1_serializeOpCreateDataProtectionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -703,7 +676,6 @@ func (m *awsRestjson1_serializeOpCreateDataProtectionS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataProtectionSettingsInput(v *CreateDataProtectionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -772,10 +744,8 @@ func (m *awsRestjson1_serializeOpCreateIdentityProvide
 func (m *awsRestjson1_serializeOpCreateIdentityProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -820,7 +790,6 @@ func (m *awsRestjson1_serializeOpCreateIdentityProvide
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIdentityProviderInput(v *CreateIdentityProviderInput, encoder *httpbinding.Encoder) error {
@@ -882,10 +851,8 @@ func (m *awsRestjson1_serializeOpCreateIpAccessSetting
 func (m *awsRestjson1_serializeOpCreateIpAccessSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -930,7 +897,6 @@ func (m *awsRestjson1_serializeOpCreateIpAccessSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIpAccessSettingsInput(v *CreateIpAccessSettingsInput, encoder *httpbinding.Encoder) error {
@@ -999,10 +965,8 @@ func (m *awsRestjson1_serializeOpCreateNetworkSettings
 func (m *awsRestjson1_serializeOpCreateNetworkSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1047,7 +1011,6 @@ func (m *awsRestjson1_serializeOpCreateNetworkSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNetworkSettingsInput(v *CreateNetworkSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1106,10 +1069,8 @@ func (m *awsRestjson1_serializeOpCreatePortal) HandleS
 func (m *awsRestjson1_serializeOpCreatePortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1154,7 +1115,6 @@ func (m *awsRestjson1_serializeOpCreatePortal) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePortalInput(v *CreatePortalInput, encoder *httpbinding.Encoder) error {
@@ -1226,10 +1186,8 @@ func (m *awsRestjson1_serializeOpCreateTrustStore) Han
 func (m *awsRestjson1_serializeOpCreateTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1274,7 +1232,6 @@ func (m *awsRestjson1_serializeOpCreateTrustStore) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTrustStoreInput(v *CreateTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -1321,10 +1278,8 @@ func (m *awsRestjson1_serializeOpCreateUserAccessLoggi
 func (m *awsRestjson1_serializeOpCreateUserAccessLoggingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1369,7 +1324,6 @@ func (m *awsRestjson1_serializeOpCreateUserAccessLoggi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUserAccessLoggingSettingsInput(v *CreateUserAccessLoggingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1414,10 +1368,8 @@ func (m *awsRestjson1_serializeOpCreateUserSettings) H
 func (m *awsRestjson1_serializeOpCreateUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1462,7 +1414,6 @@ func (m *awsRestjson1_serializeOpCreateUserSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUserSettingsInput(v *CreateUserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1568,10 +1519,8 @@ func (m *awsRestjson1_serializeOpDeleteBrowserSettings
 func (m *awsRestjson1_serializeOpDeleteBrowserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1609,7 +1558,6 @@ func (m *awsRestjson1_serializeOpDeleteBrowserSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBrowserSettingsInput(v *DeleteBrowserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1639,10 +1587,8 @@ func (m *awsRestjson1_serializeOpDeleteDataProtectionS
 func (m *awsRestjson1_serializeOpDeleteDataProtectionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1680,7 +1626,6 @@ func (m *awsRestjson1_serializeOpDeleteDataProtectionS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataProtectionSettingsInput(v *DeleteDataProtectionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1710,10 +1655,8 @@ func (m *awsRestjson1_serializeOpDeleteIdentityProvide
 func (m *awsRestjson1_serializeOpDeleteIdentityProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1751,7 +1694,6 @@ func (m *awsRestjson1_serializeOpDeleteIdentityProvide
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIdentityProviderInput(v *DeleteIdentityProviderInput, encoder *httpbinding.Encoder) error {
@@ -1781,10 +1723,8 @@ func (m *awsRestjson1_serializeOpDeleteIpAccessSetting
 func (m *awsRestjson1_serializeOpDeleteIpAccessSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1822,7 +1762,6 @@ func (m *awsRestjson1_serializeOpDeleteIpAccessSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIpAccessSettingsInput(v *DeleteIpAccessSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1852,10 +1791,8 @@ func (m *awsRestjson1_serializeOpDeleteNetworkSettings
 func (m *awsRestjson1_serializeOpDeleteNetworkSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1893,7 +1830,6 @@ func (m *awsRestjson1_serializeOpDeleteNetworkSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteNetworkSettingsInput(v *DeleteNetworkSettingsInput, encoder *httpbinding.Encoder) error {
@@ -1923,10 +1859,8 @@ func (m *awsRestjson1_serializeOpDeletePortal) HandleS
 func (m *awsRestjson1_serializeOpDeletePortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1964,7 +1898,6 @@ func (m *awsRestjson1_serializeOpDeletePortal) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePortalInput(v *DeletePortalInput, encoder *httpbinding.Encoder) error {
@@ -1994,10 +1927,8 @@ func (m *awsRestjson1_serializeOpDeleteTrustStore) Han
 func (m *awsRestjson1_serializeOpDeleteTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2035,7 +1966,6 @@ func (m *awsRestjson1_serializeOpDeleteTrustStore) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTrustStoreInput(v *DeleteTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -2065,10 +1995,8 @@ func (m *awsRestjson1_serializeOpDeleteUserAccessLoggi
 func (m *awsRestjson1_serializeOpDeleteUserAccessLoggingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2106,7 +2034,6 @@ func (m *awsRestjson1_serializeOpDeleteUserAccessLoggi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserAccessLoggingSettingsInput(v *DeleteUserAccessLoggingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2136,10 +2063,8 @@ func (m *awsRestjson1_serializeOpDeleteUserSettings) H
 func (m *awsRestjson1_serializeOpDeleteUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2177,7 +2102,6 @@ func (m *awsRestjson1_serializeOpDeleteUserSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteUserSettingsInput(v *DeleteUserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2207,10 +2131,8 @@ func (m *awsRestjson1_serializeOpDisassociateBrowserSe
 func (m *awsRestjson1_serializeOpDisassociateBrowserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2248,7 +2170,6 @@ func (m *awsRestjson1_serializeOpDisassociateBrowserSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateBrowserSettingsInput(v *DisassociateBrowserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2278,10 +2199,8 @@ func (m *awsRestjson1_serializeOpDisassociateDataProte
 func (m *awsRestjson1_serializeOpDisassociateDataProtectionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2319,7 +2238,6 @@ func (m *awsRestjson1_serializeOpDisassociateDataProte
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateDataProtectionSettingsInput(v *DisassociateDataProtectionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2349,10 +2267,8 @@ func (m *awsRestjson1_serializeOpDisassociateIpAccessS
 func (m *awsRestjson1_serializeOpDisassociateIpAccessSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2390,7 +2306,6 @@ func (m *awsRestjson1_serializeOpDisassociateIpAccessS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateIpAccessSettingsInput(v *DisassociateIpAccessSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2420,10 +2335,8 @@ func (m *awsRestjson1_serializeOpDisassociateNetworkSe
 func (m *awsRestjson1_serializeOpDisassociateNetworkSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2461,7 +2374,6 @@ func (m *awsRestjson1_serializeOpDisassociateNetworkSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateNetworkSettingsInput(v *DisassociateNetworkSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2491,10 +2403,8 @@ func (m *awsRestjson1_serializeOpDisassociateTrustStor
 func (m *awsRestjson1_serializeOpDisassociateTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2532,7 +2442,6 @@ func (m *awsRestjson1_serializeOpDisassociateTrustStor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateTrustStoreInput(v *DisassociateTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -2562,10 +2471,8 @@ func (m *awsRestjson1_serializeOpDisassociateUserAcces
 func (m *awsRestjson1_serializeOpDisassociateUserAccessLoggingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2603,7 +2510,6 @@ func (m *awsRestjson1_serializeOpDisassociateUserAcces
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateUserAccessLoggingSettingsInput(v *DisassociateUserAccessLoggingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2633,10 +2539,8 @@ func (m *awsRestjson1_serializeOpDisassociateUserSetti
 func (m *awsRestjson1_serializeOpDisassociateUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2674,7 +2578,6 @@ func (m *awsRestjson1_serializeOpDisassociateUserSetti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateUserSettingsInput(v *DisassociateUserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2704,10 +2607,8 @@ func (m *awsRestjson1_serializeOpExpireSession) Handle
 func (m *awsRestjson1_serializeOpExpireSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2745,7 +2646,6 @@ func (m *awsRestjson1_serializeOpExpireSession) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExpireSessionInput(v *ExpireSessionInput, encoder *httpbinding.Encoder) error {
@@ -2784,10 +2684,8 @@ func (m *awsRestjson1_serializeOpGetBrowserSettings) H
 func (m *awsRestjson1_serializeOpGetBrowserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2825,7 +2723,6 @@ func (m *awsRestjson1_serializeOpGetBrowserSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBrowserSettingsInput(v *GetBrowserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2855,10 +2752,8 @@ func (m *awsRestjson1_serializeOpGetDataProtectionSett
 func (m *awsRestjson1_serializeOpGetDataProtectionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2896,7 +2791,6 @@ func (m *awsRestjson1_serializeOpGetDataProtectionSett
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataProtectionSettingsInput(v *GetDataProtectionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2926,10 +2820,8 @@ func (m *awsRestjson1_serializeOpGetIdentityProvider) 
 func (m *awsRestjson1_serializeOpGetIdentityProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2967,7 +2859,6 @@ func (m *awsRestjson1_serializeOpGetIdentityProvider) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIdentityProviderInput(v *GetIdentityProviderInput, encoder *httpbinding.Encoder) error {
@@ -2997,10 +2888,8 @@ func (m *awsRestjson1_serializeOpGetIpAccessSettings) 
 func (m *awsRestjson1_serializeOpGetIpAccessSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3038,7 +2927,6 @@ func (m *awsRestjson1_serializeOpGetIpAccessSettings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIpAccessSettingsInput(v *GetIpAccessSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3068,10 +2956,8 @@ func (m *awsRestjson1_serializeOpGetNetworkSettings) H
 func (m *awsRestjson1_serializeOpGetNetworkSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3109,7 +2995,6 @@ func (m *awsRestjson1_serializeOpGetNetworkSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNetworkSettingsInput(v *GetNetworkSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3139,10 +3024,8 @@ func (m *awsRestjson1_serializeOpGetPortal) HandleSeri
 func (m *awsRestjson1_serializeOpGetPortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3180,7 +3063,6 @@ func (m *awsRestjson1_serializeOpGetPortal) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPortalInput(v *GetPortalInput, encoder *httpbinding.Encoder) error {
@@ -3210,10 +3092,8 @@ func (m *awsRestjson1_serializeOpGetPortalServiceProvi
 func (m *awsRestjson1_serializeOpGetPortalServiceProviderMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3251,7 +3131,6 @@ func (m *awsRestjson1_serializeOpGetPortalServiceProvi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPortalServiceProviderMetadataInput(v *GetPortalServiceProviderMetadataInput, encoder *httpbinding.Encoder) error {
@@ -3281,10 +3160,8 @@ func (m *awsRestjson1_serializeOpGetSession) HandleSer
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
@@ -3322,7 +3199,6 @@ func (m *awsRestjson1_serializeOpGetSession) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSessionInput(v *GetSessionInput, encoder *httpbinding.Encoder) error {
@@ -3361,10 +3237,8 @@ func (m *awsRestjson1_serializeOpGetTrustStore) Handle
 func (m *awsRestjson1_serializeOpGetTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3402,7 +3276,6 @@ func (m *awsRestjson1_serializeOpGetTrustStore) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTrustStoreInput(v *GetTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -3432,10 +3305,8 @@ func (m *awsRestjson1_serializeOpGetTrustStoreCertific
 func (m *awsRestjson1_serializeOpGetTrustStoreCertificate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3473,7 +3344,6 @@ func (m *awsRestjson1_serializeOpGetTrustStoreCertific
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTrustStoreCertificateInput(v *GetTrustStoreCertificateInput, encoder *httpbinding.Encoder) error {
@@ -3507,10 +3377,8 @@ func (m *awsRestjson1_serializeOpGetUserAccessLoggingS
 func (m *awsRestjson1_serializeOpGetUserAccessLoggingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3548,7 +3416,6 @@ func (m *awsRestjson1_serializeOpGetUserAccessLoggingS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserAccessLoggingSettingsInput(v *GetUserAccessLoggingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3578,10 +3445,8 @@ func (m *awsRestjson1_serializeOpGetUserSettings) Hand
 func (m *awsRestjson1_serializeOpGetUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3619,7 +3484,6 @@ func (m *awsRestjson1_serializeOpGetUserSettings) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserSettingsInput(v *GetUserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3649,10 +3513,8 @@ func (m *awsRestjson1_serializeOpListBrowserSettings) 
 func (m *awsRestjson1_serializeOpListBrowserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3690,7 +3552,6 @@ func (m *awsRestjson1_serializeOpListBrowserSettings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBrowserSettingsInput(v *ListBrowserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3719,10 +3580,8 @@ func (m *awsRestjson1_serializeOpListDataProtectionSet
 func (m *awsRestjson1_serializeOpListDataProtectionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3760,7 +3619,6 @@ func (m *awsRestjson1_serializeOpListDataProtectionSet
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataProtectionSettingsInput(v *ListDataProtectionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3789,10 +3647,8 @@ func (m *awsRestjson1_serializeOpListIdentityProviders
 func (m *awsRestjson1_serializeOpListIdentityProviders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3830,7 +3686,6 @@ func (m *awsRestjson1_serializeOpListIdentityProviders
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIdentityProvidersInput(v *ListIdentityProvidersInput, encoder *httpbinding.Encoder) error {
@@ -3868,10 +3723,8 @@ func (m *awsRestjson1_serializeOpListIpAccessSettings)
 func (m *awsRestjson1_serializeOpListIpAccessSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3909,7 +3762,6 @@ func (m *awsRestjson1_serializeOpListIpAccessSettings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIpAccessSettingsInput(v *ListIpAccessSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3938,10 +3790,8 @@ func (m *awsRestjson1_serializeOpListNetworkSettings) 
 func (m *awsRestjson1_serializeOpListNetworkSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3979,7 +3829,6 @@ func (m *awsRestjson1_serializeOpListNetworkSettings) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNetworkSettingsInput(v *ListNetworkSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4008,10 +3857,8 @@ func (m *awsRestjson1_serializeOpListPortals) HandleSe
 func (m *awsRestjson1_serializeOpListPortals) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4049,7 +3896,6 @@ func (m *awsRestjson1_serializeOpListPortals) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPortalsInput(v *ListPortalsInput, encoder *httpbinding.Encoder) error {
@@ -4078,10 +3924,8 @@ func (m *awsRestjson1_serializeOpListSessions) HandleS
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
@@ -4119,7 +3963,6 @@ func (m *awsRestjson1_serializeOpListSessions) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSessionsInput(v *ListSessionsInput, encoder *httpbinding.Encoder) error {
@@ -4173,10 +4016,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -4214,7 +4055,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4244,10 +4084,8 @@ func (m *awsRestjson1_serializeOpListTrustStoreCertifi
 func (m *awsRestjson1_serializeOpListTrustStoreCertificates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4285,7 +4123,6 @@ func (m *awsRestjson1_serializeOpListTrustStoreCertifi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrustStoreCertificatesInput(v *ListTrustStoreCertificatesInput, encoder *httpbinding.Encoder) error {
@@ -4323,10 +4160,8 @@ func (m *awsRestjson1_serializeOpListTrustStores) Hand
 func (m *awsRestjson1_serializeOpListTrustStores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4364,7 +4199,6 @@ func (m *awsRestjson1_serializeOpListTrustStores) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTrustStoresInput(v *ListTrustStoresInput, encoder *httpbinding.Encoder) error {
@@ -4393,10 +4227,8 @@ func (m *awsRestjson1_serializeOpListUserAccessLogging
 func (m *awsRestjson1_serializeOpListUserAccessLoggingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4434,7 +4266,6 @@ func (m *awsRestjson1_serializeOpListUserAccessLogging
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUserAccessLoggingSettingsInput(v *ListUserAccessLoggingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4463,10 +4294,8 @@ func (m *awsRestjson1_serializeOpListUserSettings) Han
 func (m *awsRestjson1_serializeOpListUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4504,7 +4333,6 @@ func (m *awsRestjson1_serializeOpListUserSettings) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUserSettingsInput(v *ListUserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4533,10 +4361,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -4585,7 +4411,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4634,10 +4459,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -4675,7 +4498,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4711,10 +4533,8 @@ func (m *awsRestjson1_serializeOpUpdateBrowserSettings
 func (m *awsRestjson1_serializeOpUpdateBrowserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4763,7 +4583,6 @@ func (m *awsRestjson1_serializeOpUpdateBrowserSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBrowserSettingsInput(v *UpdateBrowserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4810,10 +4629,8 @@ func (m *awsRestjson1_serializeOpUpdateDataProtectionS
 func (m *awsRestjson1_serializeOpUpdateDataProtectionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4862,7 +4679,6 @@ func (m *awsRestjson1_serializeOpUpdateDataProtectionS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataProtectionSettingsInput(v *UpdateDataProtectionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4921,10 +4737,8 @@ func (m *awsRestjson1_serializeOpUpdateIdentityProvide
 func (m *awsRestjson1_serializeOpUpdateIdentityProvider) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4973,7 +4787,6 @@ func (m *awsRestjson1_serializeOpUpdateIdentityProvide
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIdentityProviderInput(v *UpdateIdentityProviderInput, encoder *httpbinding.Encoder) error {
@@ -5032,10 +4845,8 @@ func (m *awsRestjson1_serializeOpUpdateIpAccessSetting
 func (m *awsRestjson1_serializeOpUpdateIpAccessSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5084,7 +4895,6 @@ func (m *awsRestjson1_serializeOpUpdateIpAccessSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIpAccessSettingsInput(v *UpdateIpAccessSettingsInput, encoder *httpbinding.Encoder) error {
@@ -5143,10 +4953,8 @@ func (m *awsRestjson1_serializeOpUpdateNetworkSettings
 func (m *awsRestjson1_serializeOpUpdateNetworkSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5195,7 +5003,6 @@ func (m *awsRestjson1_serializeOpUpdateNetworkSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateNetworkSettingsInput(v *UpdateNetworkSettingsInput, encoder *httpbinding.Encoder) error {
@@ -5256,10 +5063,8 @@ func (m *awsRestjson1_serializeOpUpdatePortal) HandleS
 func (m *awsRestjson1_serializeOpUpdatePortal) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5308,7 +5113,6 @@ func (m *awsRestjson1_serializeOpUpdatePortal) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePortalInput(v *UpdatePortalInput, encoder *httpbinding.Encoder) error {
@@ -5365,10 +5169,8 @@ func (m *awsRestjson1_serializeOpUpdateTrustStore) Han
 func (m *awsRestjson1_serializeOpUpdateTrustStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5417,7 +5219,6 @@ func (m *awsRestjson1_serializeOpUpdateTrustStore) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTrustStoreInput(v *UpdateTrustStoreInput, encoder *httpbinding.Encoder) error {
@@ -5473,10 +5274,8 @@ func (m *awsRestjson1_serializeOpUpdateUserAccessLoggi
 func (m *awsRestjson1_serializeOpUpdateUserAccessLoggingSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5525,7 +5324,6 @@ func (m *awsRestjson1_serializeOpUpdateUserAccessLoggi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserAccessLoggingSettingsInput(v *UpdateUserAccessLoggingSettingsInput, encoder *httpbinding.Encoder) error {
@@ -5572,10 +5370,8 @@ func (m *awsRestjson1_serializeOpUpdateUserSettings) H
 func (m *awsRestjson1_serializeOpUpdateUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5624,7 +5420,6 @@ func (m *awsRestjson1_serializeOpUpdateUserSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserSettingsInput(v *UpdateUserSettingsInput, encoder *httpbinding.Encoder) error {
