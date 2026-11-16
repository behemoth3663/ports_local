--- vendor/github.com/aws/aws-sdk-go-v2/service/greengrass/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/greengrass/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociateRoleToGroup)
 func (m *awsRestjson1_serializeOpAssociateRoleToGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpAssociateRoleToGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateRoleToGroupInput(v *AssociateRoleToGroupInput, encoder *httpbinding.Encoder) error {
@@ -119,10 +115,8 @@ func (m *awsRestjson1_serializeOpAssociateServiceRoleT
 func (m *awsRestjson1_serializeOpAssociateServiceRoleToAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -167,7 +161,6 @@ func (m *awsRestjson1_serializeOpAssociateServiceRoleT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateServiceRoleToAccountInput(v *AssociateServiceRoleToAccountInput, encoder *httpbinding.Encoder) error {
@@ -200,10 +193,8 @@ func (m *awsRestjson1_serializeOpCreateConnectorDefini
 func (m *awsRestjson1_serializeOpCreateConnectorDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -252,7 +243,6 @@ func (m *awsRestjson1_serializeOpCreateConnectorDefini
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectorDefinitionInput(v *CreateConnectorDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -304,10 +294,8 @@ func (m *awsRestjson1_serializeOpCreateConnectorDefini
 func (m *awsRestjson1_serializeOpCreateConnectorDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -356,7 +344,6 @@ func (m *awsRestjson1_serializeOpCreateConnectorDefini
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectorDefinitionVersionInput(v *CreateConnectorDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -405,10 +392,8 @@ func (m *awsRestjson1_serializeOpCreateCoreDefinition)
 func (m *awsRestjson1_serializeOpCreateCoreDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -457,7 +442,6 @@ func (m *awsRestjson1_serializeOpCreateCoreDefinition)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCoreDefinitionInput(v *CreateCoreDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -509,10 +493,8 @@ func (m *awsRestjson1_serializeOpCreateCoreDefinitionV
 func (m *awsRestjson1_serializeOpCreateCoreDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -561,7 +543,6 @@ func (m *awsRestjson1_serializeOpCreateCoreDefinitionV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCoreDefinitionVersionInput(v *CreateCoreDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -610,10 +591,8 @@ func (m *awsRestjson1_serializeOpCreateDeployment) Han
 func (m *awsRestjson1_serializeOpCreateDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -662,7 +641,6 @@ func (m *awsRestjson1_serializeOpCreateDeployment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeploymentInput(v *CreateDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -719,10 +697,8 @@ func (m *awsRestjson1_serializeOpCreateDeviceDefinitio
 func (m *awsRestjson1_serializeOpCreateDeviceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -771,7 +747,6 @@ func (m *awsRestjson1_serializeOpCreateDeviceDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeviceDefinitionInput(v *CreateDeviceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -823,10 +798,8 @@ func (m *awsRestjson1_serializeOpCreateDeviceDefinitio
 func (m *awsRestjson1_serializeOpCreateDeviceDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -875,7 +848,6 @@ func (m *awsRestjson1_serializeOpCreateDeviceDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeviceDefinitionVersionInput(v *CreateDeviceDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -924,10 +896,8 @@ func (m *awsRestjson1_serializeOpCreateFunctionDefinit
 func (m *awsRestjson1_serializeOpCreateFunctionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -976,7 +946,6 @@ func (m *awsRestjson1_serializeOpCreateFunctionDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFunctionDefinitionInput(v *CreateFunctionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -1028,10 +997,8 @@ func (m *awsRestjson1_serializeOpCreateFunctionDefinit
 func (m *awsRestjson1_serializeOpCreateFunctionDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1080,7 +1047,6 @@ func (m *awsRestjson1_serializeOpCreateFunctionDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFunctionDefinitionVersionInput(v *CreateFunctionDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -1136,10 +1102,8 @@ func (m *awsRestjson1_serializeOpCreateGroup) HandleSe
 func (m *awsRestjson1_serializeOpCreateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1188,7 +1152,6 @@ func (m *awsRestjson1_serializeOpCreateGroup) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGroupInput(v *CreateGroupInput, encoder *httpbinding.Encoder) error {
@@ -1240,10 +1203,8 @@ func (m *awsRestjson1_serializeOpCreateGroupCertificat
 func (m *awsRestjson1_serializeOpCreateGroupCertificateAuthority) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1281,7 +1242,6 @@ func (m *awsRestjson1_serializeOpCreateGroupCertificat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGroupCertificateAuthorityInput(v *CreateGroupCertificateAuthorityInput, encoder *httpbinding.Encoder) error {
@@ -1316,10 +1276,8 @@ func (m *awsRestjson1_serializeOpCreateGroupVersion) H
 func (m *awsRestjson1_serializeOpCreateGroupVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1368,7 +1326,6 @@ func (m *awsRestjson1_serializeOpCreateGroupVersion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGroupVersionInput(v *CreateGroupVersionInput, encoder *httpbinding.Encoder) error {
@@ -1445,10 +1402,8 @@ func (m *awsRestjson1_serializeOpCreateLoggerDefinitio
 func (m *awsRestjson1_serializeOpCreateLoggerDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1497,7 +1452,6 @@ func (m *awsRestjson1_serializeOpCreateLoggerDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLoggerDefinitionInput(v *CreateLoggerDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -1549,10 +1503,8 @@ func (m *awsRestjson1_serializeOpCreateLoggerDefinitio
 func (m *awsRestjson1_serializeOpCreateLoggerDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1601,7 +1553,6 @@ func (m *awsRestjson1_serializeOpCreateLoggerDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLoggerDefinitionVersionInput(v *CreateLoggerDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -1650,10 +1601,8 @@ func (m *awsRestjson1_serializeOpCreateResourceDefinit
 func (m *awsRestjson1_serializeOpCreateResourceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1702,7 +1651,6 @@ func (m *awsRestjson1_serializeOpCreateResourceDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateResourceDefinitionInput(v *CreateResourceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -1754,10 +1702,8 @@ func (m *awsRestjson1_serializeOpCreateResourceDefinit
 func (m *awsRestjson1_serializeOpCreateResourceDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1806,7 +1752,6 @@ func (m *awsRestjson1_serializeOpCreateResourceDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateResourceDefinitionVersionInput(v *CreateResourceDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -1855,10 +1800,8 @@ func (m *awsRestjson1_serializeOpCreateSoftwareUpdateJ
 func (m *awsRestjson1_serializeOpCreateSoftwareUpdateJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1907,7 +1850,6 @@ func (m *awsRestjson1_serializeOpCreateSoftwareUpdateJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSoftwareUpdateJobInput(v *CreateSoftwareUpdateJobInput, encoder *httpbinding.Encoder) error {
@@ -1972,10 +1914,8 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionDef
 func (m *awsRestjson1_serializeOpCreateSubscriptionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2024,7 +1964,6 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionDef
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSubscriptionDefinitionInput(v *CreateSubscriptionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2076,10 +2015,8 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionDef
 func (m *awsRestjson1_serializeOpCreateSubscriptionDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2128,7 +2065,6 @@ func (m *awsRestjson1_serializeOpCreateSubscriptionDef
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSubscriptionDefinitionVersionInput(v *CreateSubscriptionDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -2177,10 +2113,8 @@ func (m *awsRestjson1_serializeOpDeleteConnectorDefini
 func (m *awsRestjson1_serializeOpDeleteConnectorDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2218,7 +2152,6 @@ func (m *awsRestjson1_serializeOpDeleteConnectorDefini
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConnectorDefinitionInput(v *DeleteConnectorDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2248,10 +2181,8 @@ func (m *awsRestjson1_serializeOpDeleteCoreDefinition)
 func (m *awsRestjson1_serializeOpDeleteCoreDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2289,7 +2220,6 @@ func (m *awsRestjson1_serializeOpDeleteCoreDefinition)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCoreDefinitionInput(v *DeleteCoreDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2319,10 +2249,8 @@ func (m *awsRestjson1_serializeOpDeleteDeviceDefinitio
 func (m *awsRestjson1_serializeOpDeleteDeviceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2360,7 +2288,6 @@ func (m *awsRestjson1_serializeOpDeleteDeviceDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDeviceDefinitionInput(v *DeleteDeviceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2390,10 +2317,8 @@ func (m *awsRestjson1_serializeOpDeleteFunctionDefinit
 func (m *awsRestjson1_serializeOpDeleteFunctionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2431,7 +2356,6 @@ func (m *awsRestjson1_serializeOpDeleteFunctionDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFunctionDefinitionInput(v *DeleteFunctionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2461,10 +2385,8 @@ func (m *awsRestjson1_serializeOpDeleteGroup) HandleSe
 func (m *awsRestjson1_serializeOpDeleteGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2502,7 +2424,6 @@ func (m *awsRestjson1_serializeOpDeleteGroup) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGroupInput(v *DeleteGroupInput, encoder *httpbinding.Encoder) error {
@@ -2532,10 +2453,8 @@ func (m *awsRestjson1_serializeOpDeleteLoggerDefinitio
 func (m *awsRestjson1_serializeOpDeleteLoggerDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2573,7 +2492,6 @@ func (m *awsRestjson1_serializeOpDeleteLoggerDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLoggerDefinitionInput(v *DeleteLoggerDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2603,10 +2521,8 @@ func (m *awsRestjson1_serializeOpDeleteResourceDefinit
 func (m *awsRestjson1_serializeOpDeleteResourceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2644,7 +2560,6 @@ func (m *awsRestjson1_serializeOpDeleteResourceDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResourceDefinitionInput(v *DeleteResourceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2674,10 +2589,8 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionDef
 func (m *awsRestjson1_serializeOpDeleteSubscriptionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2715,7 +2628,6 @@ func (m *awsRestjson1_serializeOpDeleteSubscriptionDef
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSubscriptionDefinitionInput(v *DeleteSubscriptionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -2745,10 +2657,8 @@ func (m *awsRestjson1_serializeOpDisassociateRoleFromG
 func (m *awsRestjson1_serializeOpDisassociateRoleFromGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2786,7 +2696,6 @@ func (m *awsRestjson1_serializeOpDisassociateRoleFromG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateRoleFromGroupInput(v *DisassociateRoleFromGroupInput, encoder *httpbinding.Encoder) error {
@@ -2816,10 +2725,8 @@ func (m *awsRestjson1_serializeOpDisassociateServiceRo
 func (m *awsRestjson1_serializeOpDisassociateServiceRoleFromAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2853,7 +2760,6 @@ func (m *awsRestjson1_serializeOpDisassociateServiceRo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateServiceRoleFromAccountInput(v *DisassociateServiceRoleFromAccountInput, encoder *httpbinding.Encoder) error {
@@ -2874,10 +2780,8 @@ func (m *awsRestjson1_serializeOpGetAssociatedRole) Ha
 func (m *awsRestjson1_serializeOpGetAssociatedRole) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2915,7 +2819,6 @@ func (m *awsRestjson1_serializeOpGetAssociatedRole) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssociatedRoleInput(v *GetAssociatedRoleInput, encoder *httpbinding.Encoder) error {
@@ -2945,10 +2848,8 @@ func (m *awsRestjson1_serializeOpGetBulkDeploymentStat
 func (m *awsRestjson1_serializeOpGetBulkDeploymentStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2986,7 +2887,6 @@ func (m *awsRestjson1_serializeOpGetBulkDeploymentStat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBulkDeploymentStatusInput(v *GetBulkDeploymentStatusInput, encoder *httpbinding.Encoder) error {
@@ -3016,10 +2916,8 @@ func (m *awsRestjson1_serializeOpGetConnectivityInfo) 
 func (m *awsRestjson1_serializeOpGetConnectivityInfo) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3057,7 +2955,6 @@ func (m *awsRestjson1_serializeOpGetConnectivityInfo) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectivityInfoInput(v *GetConnectivityInfoInput, encoder *httpbinding.Encoder) error {
@@ -3087,10 +2984,8 @@ func (m *awsRestjson1_serializeOpGetConnectorDefinitio
 func (m *awsRestjson1_serializeOpGetConnectorDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3128,7 +3023,6 @@ func (m *awsRestjson1_serializeOpGetConnectorDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectorDefinitionInput(v *GetConnectorDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -3158,10 +3052,8 @@ func (m *awsRestjson1_serializeOpGetConnectorDefinitio
 func (m *awsRestjson1_serializeOpGetConnectorDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3199,7 +3091,6 @@ func (m *awsRestjson1_serializeOpGetConnectorDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectorDefinitionVersionInput(v *GetConnectorDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -3242,10 +3133,8 @@ func (m *awsRestjson1_serializeOpGetCoreDefinition) Ha
 func (m *awsRestjson1_serializeOpGetCoreDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3283,7 +3172,6 @@ func (m *awsRestjson1_serializeOpGetCoreDefinition) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCoreDefinitionInput(v *GetCoreDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -3313,10 +3201,8 @@ func (m *awsRestjson1_serializeOpGetCoreDefinitionVers
 func (m *awsRestjson1_serializeOpGetCoreDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3354,7 +3240,6 @@ func (m *awsRestjson1_serializeOpGetCoreDefinitionVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCoreDefinitionVersionInput(v *GetCoreDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -3393,10 +3278,8 @@ func (m *awsRestjson1_serializeOpGetDeploymentStatus) 
 func (m *awsRestjson1_serializeOpGetDeploymentStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3434,7 +3317,6 @@ func (m *awsRestjson1_serializeOpGetDeploymentStatus) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeploymentStatusInput(v *GetDeploymentStatusInput, encoder *httpbinding.Encoder) error {
@@ -3473,10 +3355,8 @@ func (m *awsRestjson1_serializeOpGetDeviceDefinition) 
 func (m *awsRestjson1_serializeOpGetDeviceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3514,7 +3394,6 @@ func (m *awsRestjson1_serializeOpGetDeviceDefinition) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeviceDefinitionInput(v *GetDeviceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -3544,10 +3423,8 @@ func (m *awsRestjson1_serializeOpGetDeviceDefinitionVe
 func (m *awsRestjson1_serializeOpGetDeviceDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3585,7 +3462,6 @@ func (m *awsRestjson1_serializeOpGetDeviceDefinitionVe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeviceDefinitionVersionInput(v *GetDeviceDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -3628,10 +3504,8 @@ func (m *awsRestjson1_serializeOpGetFunctionDefinition
 func (m *awsRestjson1_serializeOpGetFunctionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3669,7 +3543,6 @@ func (m *awsRestjson1_serializeOpGetFunctionDefinition
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionDefinitionInput(v *GetFunctionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -3699,10 +3572,8 @@ func (m *awsRestjson1_serializeOpGetFunctionDefinition
 func (m *awsRestjson1_serializeOpGetFunctionDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3740,7 +3611,6 @@ func (m *awsRestjson1_serializeOpGetFunctionDefinition
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFunctionDefinitionVersionInput(v *GetFunctionDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -3783,10 +3653,8 @@ func (m *awsRestjson1_serializeOpGetGroup) HandleSeria
 func (m *awsRestjson1_serializeOpGetGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3824,7 +3692,6 @@ func (m *awsRestjson1_serializeOpGetGroup) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGroupInput(v *GetGroupInput, encoder *httpbinding.Encoder) error {
@@ -3854,10 +3721,8 @@ func (m *awsRestjson1_serializeOpGetGroupCertificateAu
 func (m *awsRestjson1_serializeOpGetGroupCertificateAuthority) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3895,7 +3760,6 @@ func (m *awsRestjson1_serializeOpGetGroupCertificateAu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGroupCertificateAuthorityInput(v *GetGroupCertificateAuthorityInput, encoder *httpbinding.Encoder) error {
@@ -3934,10 +3798,8 @@ func (m *awsRestjson1_serializeOpGetGroupCertificateCo
 func (m *awsRestjson1_serializeOpGetGroupCertificateConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3975,7 +3837,6 @@ func (m *awsRestjson1_serializeOpGetGroupCertificateCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGroupCertificateConfigurationInput(v *GetGroupCertificateConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4005,10 +3866,8 @@ func (m *awsRestjson1_serializeOpGetGroupVersion) Hand
 func (m *awsRestjson1_serializeOpGetGroupVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4046,7 +3905,6 @@ func (m *awsRestjson1_serializeOpGetGroupVersion) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGroupVersionInput(v *GetGroupVersionInput, encoder *httpbinding.Encoder) error {
@@ -4085,10 +3943,8 @@ func (m *awsRestjson1_serializeOpGetLoggerDefinition) 
 func (m *awsRestjson1_serializeOpGetLoggerDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4126,7 +3982,6 @@ func (m *awsRestjson1_serializeOpGetLoggerDefinition) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLoggerDefinitionInput(v *GetLoggerDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -4156,10 +4011,8 @@ func (m *awsRestjson1_serializeOpGetLoggerDefinitionVe
 func (m *awsRestjson1_serializeOpGetLoggerDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4197,7 +4050,6 @@ func (m *awsRestjson1_serializeOpGetLoggerDefinitionVe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLoggerDefinitionVersionInput(v *GetLoggerDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -4240,10 +4092,8 @@ func (m *awsRestjson1_serializeOpGetResourceDefinition
 func (m *awsRestjson1_serializeOpGetResourceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4281,7 +4131,6 @@ func (m *awsRestjson1_serializeOpGetResourceDefinition
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceDefinitionInput(v *GetResourceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -4311,10 +4160,8 @@ func (m *awsRestjson1_serializeOpGetResourceDefinition
 func (m *awsRestjson1_serializeOpGetResourceDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4352,7 +4199,6 @@ func (m *awsRestjson1_serializeOpGetResourceDefinition
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceDefinitionVersionInput(v *GetResourceDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -4391,10 +4237,8 @@ func (m *awsRestjson1_serializeOpGetServiceRoleForAcco
 func (m *awsRestjson1_serializeOpGetServiceRoleForAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4428,7 +4272,6 @@ func (m *awsRestjson1_serializeOpGetServiceRoleForAcco
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetServiceRoleForAccountInput(v *GetServiceRoleForAccountInput, encoder *httpbinding.Encoder) error {
@@ -4449,10 +4292,8 @@ func (m *awsRestjson1_serializeOpGetSubscriptionDefini
 func (m *awsRestjson1_serializeOpGetSubscriptionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4490,7 +4331,6 @@ func (m *awsRestjson1_serializeOpGetSubscriptionDefini
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSubscriptionDefinitionInput(v *GetSubscriptionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -4520,10 +4360,8 @@ func (m *awsRestjson1_serializeOpGetSubscriptionDefini
 func (m *awsRestjson1_serializeOpGetSubscriptionDefinitionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4561,7 +4399,6 @@ func (m *awsRestjson1_serializeOpGetSubscriptionDefini
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSubscriptionDefinitionVersionInput(v *GetSubscriptionDefinitionVersionInput, encoder *httpbinding.Encoder) error {
@@ -4604,10 +4441,8 @@ func (m *awsRestjson1_serializeOpGetThingRuntimeConfig
 func (m *awsRestjson1_serializeOpGetThingRuntimeConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4645,7 +4480,6 @@ func (m *awsRestjson1_serializeOpGetThingRuntimeConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetThingRuntimeConfigurationInput(v *GetThingRuntimeConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4675,10 +4509,8 @@ func (m *awsRestjson1_serializeOpListBulkDeploymentDet
 func (m *awsRestjson1_serializeOpListBulkDeploymentDetailedReports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4716,7 +4548,6 @@ func (m *awsRestjson1_serializeOpListBulkDeploymentDet
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBulkDeploymentDetailedReportsInput(v *ListBulkDeploymentDetailedReportsInput, encoder *httpbinding.Encoder) error {
@@ -4754,10 +4585,8 @@ func (m *awsRestjson1_serializeOpListBulkDeployments) 
 func (m *awsRestjson1_serializeOpListBulkDeployments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4795,7 +4624,6 @@ func (m *awsRestjson1_serializeOpListBulkDeployments) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBulkDeploymentsInput(v *ListBulkDeploymentsInput, encoder *httpbinding.Encoder) error {
@@ -4824,10 +4652,8 @@ func (m *awsRestjson1_serializeOpListConnectorDefiniti
 func (m *awsRestjson1_serializeOpListConnectorDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4865,7 +4691,6 @@ func (m *awsRestjson1_serializeOpListConnectorDefiniti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConnectorDefinitionsInput(v *ListConnectorDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -4894,10 +4719,8 @@ func (m *awsRestjson1_serializeOpListConnectorDefiniti
 func (m *awsRestjson1_serializeOpListConnectorDefinitionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4935,7 +4758,6 @@ func (m *awsRestjson1_serializeOpListConnectorDefiniti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConnectorDefinitionVersionsInput(v *ListConnectorDefinitionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4973,10 +4795,8 @@ func (m *awsRestjson1_serializeOpListCoreDefinitions) 
 func (m *awsRestjson1_serializeOpListCoreDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5014,7 +4834,6 @@ func (m *awsRestjson1_serializeOpListCoreDefinitions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCoreDefinitionsInput(v *ListCoreDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -5043,10 +4862,8 @@ func (m *awsRestjson1_serializeOpListCoreDefinitionVer
 func (m *awsRestjson1_serializeOpListCoreDefinitionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5084,7 +4901,6 @@ func (m *awsRestjson1_serializeOpListCoreDefinitionVer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCoreDefinitionVersionsInput(v *ListCoreDefinitionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5122,10 +4938,8 @@ func (m *awsRestjson1_serializeOpListDeployments) Hand
 func (m *awsRestjson1_serializeOpListDeployments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5163,7 +4977,6 @@ func (m *awsRestjson1_serializeOpListDeployments) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDeploymentsInput(v *ListDeploymentsInput, encoder *httpbinding.Encoder) error {
@@ -5201,10 +5014,8 @@ func (m *awsRestjson1_serializeOpListDeviceDefinitions
 func (m *awsRestjson1_serializeOpListDeviceDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5242,7 +5053,6 @@ func (m *awsRestjson1_serializeOpListDeviceDefinitions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDeviceDefinitionsInput(v *ListDeviceDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -5271,10 +5081,8 @@ func (m *awsRestjson1_serializeOpListDeviceDefinitionV
 func (m *awsRestjson1_serializeOpListDeviceDefinitionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5312,7 +5120,6 @@ func (m *awsRestjson1_serializeOpListDeviceDefinitionV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDeviceDefinitionVersionsInput(v *ListDeviceDefinitionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5350,10 +5157,8 @@ func (m *awsRestjson1_serializeOpListFunctionDefinitio
 func (m *awsRestjson1_serializeOpListFunctionDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5391,7 +5196,6 @@ func (m *awsRestjson1_serializeOpListFunctionDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFunctionDefinitionsInput(v *ListFunctionDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -5420,10 +5224,8 @@ func (m *awsRestjson1_serializeOpListFunctionDefinitio
 func (m *awsRestjson1_serializeOpListFunctionDefinitionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5461,7 +5263,6 @@ func (m *awsRestjson1_serializeOpListFunctionDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFunctionDefinitionVersionsInput(v *ListFunctionDefinitionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5499,10 +5300,8 @@ func (m *awsRestjson1_serializeOpListGroupCertificateA
 func (m *awsRestjson1_serializeOpListGroupCertificateAuthorities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5540,7 +5339,6 @@ func (m *awsRestjson1_serializeOpListGroupCertificateA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGroupCertificateAuthoritiesInput(v *ListGroupCertificateAuthoritiesInput, encoder *httpbinding.Encoder) error {
@@ -5570,10 +5368,8 @@ func (m *awsRestjson1_serializeOpListGroups) HandleSer
 func (m *awsRestjson1_serializeOpListGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5611,7 +5407,6 @@ func (m *awsRestjson1_serializeOpListGroups) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGroupsInput(v *ListGroupsInput, encoder *httpbinding.Encoder) error {
@@ -5640,10 +5435,8 @@ func (m *awsRestjson1_serializeOpListGroupVersions) Ha
 func (m *awsRestjson1_serializeOpListGroupVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5681,7 +5474,6 @@ func (m *awsRestjson1_serializeOpListGroupVersions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListGroupVersionsInput(v *ListGroupVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5719,10 +5511,8 @@ func (m *awsRestjson1_serializeOpListLoggerDefinitions
 func (m *awsRestjson1_serializeOpListLoggerDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5760,7 +5550,6 @@ func (m *awsRestjson1_serializeOpListLoggerDefinitions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLoggerDefinitionsInput(v *ListLoggerDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -5789,10 +5578,8 @@ func (m *awsRestjson1_serializeOpListLoggerDefinitionV
 func (m *awsRestjson1_serializeOpListLoggerDefinitionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5830,7 +5617,6 @@ func (m *awsRestjson1_serializeOpListLoggerDefinitionV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLoggerDefinitionVersionsInput(v *ListLoggerDefinitionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5868,10 +5654,8 @@ func (m *awsRestjson1_serializeOpListResourceDefinitio
 func (m *awsRestjson1_serializeOpListResourceDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5909,7 +5693,6 @@ func (m *awsRestjson1_serializeOpListResourceDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceDefinitionsInput(v *ListResourceDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -5938,10 +5721,8 @@ func (m *awsRestjson1_serializeOpListResourceDefinitio
 func (m *awsRestjson1_serializeOpListResourceDefinitionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5979,7 +5760,6 @@ func (m *awsRestjson1_serializeOpListResourceDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceDefinitionVersionsInput(v *ListResourceDefinitionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -6017,10 +5797,8 @@ func (m *awsRestjson1_serializeOpListSubscriptionDefin
 func (m *awsRestjson1_serializeOpListSubscriptionDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6058,7 +5836,6 @@ func (m *awsRestjson1_serializeOpListSubscriptionDefin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubscriptionDefinitionsInput(v *ListSubscriptionDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -6087,10 +5864,8 @@ func (m *awsRestjson1_serializeOpListSubscriptionDefin
 func (m *awsRestjson1_serializeOpListSubscriptionDefinitionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6128,7 +5903,6 @@ func (m *awsRestjson1_serializeOpListSubscriptionDefin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubscriptionDefinitionVersionsInput(v *ListSubscriptionDefinitionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -6166,10 +5940,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6207,7 +5979,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6237,10 +6008,8 @@ func (m *awsRestjson1_serializeOpResetDeployments) Han
 func (m *awsRestjson1_serializeOpResetDeployments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6289,7 +6058,6 @@ func (m *awsRestjson1_serializeOpResetDeployments) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResetDeploymentsInput(v *ResetDeploymentsInput, encoder *httpbinding.Encoder) error {
@@ -6336,10 +6104,8 @@ func (m *awsRestjson1_serializeOpStartBulkDeployment) 
 func (m *awsRestjson1_serializeOpStartBulkDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6388,7 +6154,6 @@ func (m *awsRestjson1_serializeOpStartBulkDeployment) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartBulkDeploymentInput(v *StartBulkDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -6438,10 +6203,8 @@ func (m *awsRestjson1_serializeOpStopBulkDeployment) H
 func (m *awsRestjson1_serializeOpStopBulkDeployment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6479,7 +6242,6 @@ func (m *awsRestjson1_serializeOpStopBulkDeployment) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopBulkDeploymentInput(v *StopBulkDeploymentInput, encoder *httpbinding.Encoder) error {
@@ -6509,10 +6271,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -6561,7 +6321,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6605,10 +6364,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -6646,7 +6403,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6682,10 +6438,8 @@ func (m *awsRestjson1_serializeOpUpdateConnectivityInf
 func (m *awsRestjson1_serializeOpUpdateConnectivityInfo) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6734,7 +6488,6 @@ func (m *awsRestjson1_serializeOpUpdateConnectivityInf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConnectivityInfoInput(v *UpdateConnectivityInfoInput, encoder *httpbinding.Encoder) error {
@@ -6778,10 +6531,8 @@ func (m *awsRestjson1_serializeOpUpdateConnectorDefini
 func (m *awsRestjson1_serializeOpUpdateConnectorDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6830,7 +6581,6 @@ func (m *awsRestjson1_serializeOpUpdateConnectorDefini
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConnectorDefinitionInput(v *UpdateConnectorDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -6872,10 +6622,8 @@ func (m *awsRestjson1_serializeOpUpdateCoreDefinition)
 func (m *awsRestjson1_serializeOpUpdateCoreDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6924,7 +6672,6 @@ func (m *awsRestjson1_serializeOpUpdateCoreDefinition)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCoreDefinitionInput(v *UpdateCoreDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -6966,10 +6713,8 @@ func (m *awsRestjson1_serializeOpUpdateDeviceDefinitio
 func (m *awsRestjson1_serializeOpUpdateDeviceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7018,7 +6763,6 @@ func (m *awsRestjson1_serializeOpUpdateDeviceDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDeviceDefinitionInput(v *UpdateDeviceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -7060,10 +6804,8 @@ func (m *awsRestjson1_serializeOpUpdateFunctionDefinit
 func (m *awsRestjson1_serializeOpUpdateFunctionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7112,7 +6854,6 @@ func (m *awsRestjson1_serializeOpUpdateFunctionDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFunctionDefinitionInput(v *UpdateFunctionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -7154,10 +6895,8 @@ func (m *awsRestjson1_serializeOpUpdateGroup) HandleSe
 func (m *awsRestjson1_serializeOpUpdateGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7206,7 +6945,6 @@ func (m *awsRestjson1_serializeOpUpdateGroup) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGroupInput(v *UpdateGroupInput, encoder *httpbinding.Encoder) error {
@@ -7248,10 +6986,8 @@ func (m *awsRestjson1_serializeOpUpdateGroupCertificat
 func (m *awsRestjson1_serializeOpUpdateGroupCertificateConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7300,7 +7036,6 @@ func (m *awsRestjson1_serializeOpUpdateGroupCertificat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGroupCertificateConfigurationInput(v *UpdateGroupCertificateConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -7342,10 +7077,8 @@ func (m *awsRestjson1_serializeOpUpdateLoggerDefinitio
 func (m *awsRestjson1_serializeOpUpdateLoggerDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7394,7 +7127,6 @@ func (m *awsRestjson1_serializeOpUpdateLoggerDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLoggerDefinitionInput(v *UpdateLoggerDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -7436,10 +7168,8 @@ func (m *awsRestjson1_serializeOpUpdateResourceDefinit
 func (m *awsRestjson1_serializeOpUpdateResourceDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7488,7 +7218,6 @@ func (m *awsRestjson1_serializeOpUpdateResourceDefinit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceDefinitionInput(v *UpdateResourceDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -7530,10 +7259,8 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionDef
 func (m *awsRestjson1_serializeOpUpdateSubscriptionDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7582,7 +7309,6 @@ func (m *awsRestjson1_serializeOpUpdateSubscriptionDef
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSubscriptionDefinitionInput(v *UpdateSubscriptionDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -7624,10 +7350,8 @@ func (m *awsRestjson1_serializeOpUpdateThingRuntimeCon
 func (m *awsRestjson1_serializeOpUpdateThingRuntimeConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7676,7 +7400,6 @@ func (m *awsRestjson1_serializeOpUpdateThingRuntimeCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThingRuntimeConfigurationInput(v *UpdateThingRuntimeConfigurationInput, encoder *httpbinding.Encoder) error {
