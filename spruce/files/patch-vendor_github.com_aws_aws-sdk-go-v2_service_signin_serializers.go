--- vendor/github.com/aws/aws-sdk-go-v2/service/signin/serializers.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/signin/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpCreateOAuth2Token) Ha
 func (m *awsRestjson1_serializeOpCreateOAuth2Token) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -86,7 +83,6 @@ func (m *awsRestjson1_serializeOpCreateOAuth2Token) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOAuth2TokenInput(v *CreateOAuth2TokenInput, encoder *httpbinding.Encoder) error {
@@ -107,10 +103,8 @@ func (m *awsRestjson1_serializeOpCreateOAuth2TokenWith
 func (m *awsRestjson1_serializeOpCreateOAuth2TokenWithIAM) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -155,7 +149,6 @@ func (m *awsRestjson1_serializeOpCreateOAuth2TokenWith
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOAuth2TokenWithIAMInput(v *CreateOAuth2TokenWithIAMInput, encoder *httpbinding.Encoder) error {
@@ -193,10 +186,8 @@ func (m *awsRestjson1_serializeOpDeleteConsoleAuthoriz
 func (m *awsRestjson1_serializeOpDeleteConsoleAuthorizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -241,7 +232,6 @@ func (m *awsRestjson1_serializeOpDeleteConsoleAuthoriz
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConsoleAuthorizationConfigurationInput(v *DeleteConsoleAuthorizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -274,10 +264,8 @@ func (m *awsRestjson1_serializeOpDeleteResourcePermiss
 func (m *awsRestjson1_serializeOpDeleteResourcePermissionStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -322,7 +310,6 @@ func (m *awsRestjson1_serializeOpDeleteResourcePermiss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResourcePermissionStatementInput(v *DeleteResourcePermissionStatementInput, encoder *httpbinding.Encoder) error {
@@ -360,10 +347,8 @@ func (m *awsRestjson1_serializeOpGetConsoleAuthorizati
 func (m *awsRestjson1_serializeOpGetConsoleAuthorizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -408,7 +393,6 @@ func (m *awsRestjson1_serializeOpGetConsoleAuthorizati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConsoleAuthorizationConfigurationInput(v *GetConsoleAuthorizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -441,10 +425,8 @@ func (m *awsRestjson1_serializeOpGetResourcePolicy) Ha
 func (m *awsRestjson1_serializeOpGetResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -478,7 +460,6 @@ func (m *awsRestjson1_serializeOpGetResourcePolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourcePolicyInput(v *GetResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -499,10 +480,8 @@ func (m *awsRestjson1_serializeOpIntrospectOAuth2Token
 func (m *awsRestjson1_serializeOpIntrospectOAuth2TokenWithIAM) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -547,7 +526,6 @@ func (m *awsRestjson1_serializeOpIntrospectOAuth2Token
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsIntrospectOAuth2TokenWithIAMInput(v *IntrospectOAuth2TokenWithIAMInput, encoder *httpbinding.Encoder) error {
@@ -585,10 +563,8 @@ func (m *awsRestjson1_serializeOpListResourcePermissio
 func (m *awsRestjson1_serializeOpListResourcePermissionStatements) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -633,7 +609,6 @@ func (m *awsRestjson1_serializeOpListResourcePermissio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourcePermissionStatementsInput(v *ListResourcePermissionStatementsInput, encoder *httpbinding.Encoder) error {
@@ -671,10 +646,8 @@ func (m *awsRestjson1_serializeOpPutConsoleAuthorizati
 func (m *awsRestjson1_serializeOpPutConsoleAuthorizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -719,7 +692,6 @@ func (m *awsRestjson1_serializeOpPutConsoleAuthorizati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConsoleAuthorizationConfigurationInput(v *PutConsoleAuthorizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -752,10 +724,8 @@ func (m *awsRestjson1_serializeOpPutResourcePermission
 func (m *awsRestjson1_serializeOpPutResourcePermissionStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -800,7 +770,6 @@ func (m *awsRestjson1_serializeOpPutResourcePermission
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutResourcePermissionStatementInput(v *PutResourcePermissionStatementInput, encoder *httpbinding.Encoder) error {
@@ -868,10 +837,8 @@ func (m *awsRestjson1_serializeOpRevokeOAuth2TokenWith
 func (m *awsRestjson1_serializeOpRevokeOAuth2TokenWithIAM) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -916,7 +883,6 @@ func (m *awsRestjson1_serializeOpRevokeOAuth2TokenWith
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRevokeOAuth2TokenWithIAMInput(v *RevokeOAuth2TokenWithIAMInput, encoder *httpbinding.Encoder) error {
