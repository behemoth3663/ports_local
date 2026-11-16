--- vendor/github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations/serializers.go
@@ -13,7 +13,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -28,10 +27,8 @@ func (m *awsRestjson1_serializeOpCreateAccountAssociat
 func (m *awsRestjson1_serializeOpCreateAccountAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -76,7 +73,6 @@ func (m *awsRestjson1_serializeOpCreateAccountAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAccountAssociationInput(v *CreateAccountAssociationInput, encoder *httpbinding.Encoder) error {
@@ -131,10 +127,8 @@ func (m *awsRestjson1_serializeOpCreateCloudConnector)
 func (m *awsRestjson1_serializeOpCreateCloudConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -179,7 +173,6 @@ func (m *awsRestjson1_serializeOpCreateCloudConnector)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCloudConnectorInput(v *CreateCloudConnectorInput, encoder *httpbinding.Encoder) error {
@@ -234,10 +227,8 @@ func (m *awsRestjson1_serializeOpCreateConnectorDestin
 func (m *awsRestjson1_serializeOpCreateConnectorDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -282,7 +273,6 @@ func (m *awsRestjson1_serializeOpCreateConnectorDestin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectorDestinationInput(v *CreateConnectorDestinationInput, encoder *httpbinding.Encoder) error {
@@ -349,10 +339,8 @@ func (m *awsRestjson1_serializeOpCreateCredentialLocke
 func (m *awsRestjson1_serializeOpCreateCredentialLocker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -397,7 +385,6 @@ func (m *awsRestjson1_serializeOpCreateCredentialLocke
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCredentialLockerInput(v *CreateCredentialLockerInput, encoder *httpbinding.Encoder) error {
@@ -442,10 +429,8 @@ func (m *awsRestjson1_serializeOpCreateDestination) Ha
 func (m *awsRestjson1_serializeOpCreateDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -490,7 +475,6 @@ func (m *awsRestjson1_serializeOpCreateDestination) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDestinationInput(v *CreateDestinationInput, encoder *httpbinding.Encoder) error {
@@ -555,10 +539,8 @@ func (m *awsRestjson1_serializeOpCreateEventLogConfigu
 func (m *awsRestjson1_serializeOpCreateEventLogConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -603,7 +585,6 @@ func (m *awsRestjson1_serializeOpCreateEventLogConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEventLogConfigurationInput(v *CreateEventLogConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -651,10 +632,8 @@ func (m *awsRestjson1_serializeOpCreateManagedThing) H
 func (m *awsRestjson1_serializeOpCreateManagedThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -699,7 +678,6 @@ func (m *awsRestjson1_serializeOpCreateManagedThing) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateManagedThingInput(v *CreateManagedThingInput, encoder *httpbinding.Encoder) error {
@@ -815,10 +793,8 @@ func (m *awsRestjson1_serializeOpCreateNotificationCon
 func (m *awsRestjson1_serializeOpCreateNotificationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -863,7 +839,6 @@ func (m *awsRestjson1_serializeOpCreateNotificationCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNotificationConfigurationInput(v *CreateNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -913,10 +888,8 @@ func (m *awsRestjson1_serializeOpCreateOtaTask) Handle
 func (m *awsRestjson1_serializeOpCreateOtaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -961,7 +934,6 @@ func (m *awsRestjson1_serializeOpCreateOtaTask) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOtaTaskInput(v *CreateOtaTaskInput, encoder *httpbinding.Encoder) error {
@@ -1057,10 +1029,8 @@ func (m *awsRestjson1_serializeOpCreateOtaTaskConfigur
 func (m *awsRestjson1_serializeOpCreateOtaTaskConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1105,7 +1075,6 @@ func (m *awsRestjson1_serializeOpCreateOtaTaskConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOtaTaskConfigurationInput(v *CreateOtaTaskConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1155,10 +1124,8 @@ func (m *awsRestjson1_serializeOpCreateProvisioningPro
 func (m *awsRestjson1_serializeOpCreateProvisioningProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1203,7 +1170,6 @@ func (m *awsRestjson1_serializeOpCreateProvisioningPro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateProvisioningProfileInput(v *CreateProvisioningProfileInput, encoder *httpbinding.Encoder) error {
@@ -1258,10 +1224,8 @@ func (m *awsRestjson1_serializeOpDeleteAccountAssociat
 func (m *awsRestjson1_serializeOpDeleteAccountAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1299,7 +1263,6 @@ func (m *awsRestjson1_serializeOpDeleteAccountAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccountAssociationInput(v *DeleteAccountAssociationInput, encoder *httpbinding.Encoder) error {
@@ -1329,10 +1292,8 @@ func (m *awsRestjson1_serializeOpDeleteCloudConnector)
 func (m *awsRestjson1_serializeOpDeleteCloudConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1370,7 +1331,6 @@ func (m *awsRestjson1_serializeOpDeleteCloudConnector)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCloudConnectorInput(v *DeleteCloudConnectorInput, encoder *httpbinding.Encoder) error {
@@ -1400,10 +1360,8 @@ func (m *awsRestjson1_serializeOpDeleteConnectorDestin
 func (m *awsRestjson1_serializeOpDeleteConnectorDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1441,7 +1399,6 @@ func (m *awsRestjson1_serializeOpDeleteConnectorDestin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConnectorDestinationInput(v *DeleteConnectorDestinationInput, encoder *httpbinding.Encoder) error {
@@ -1471,10 +1428,8 @@ func (m *awsRestjson1_serializeOpDeleteCredentialLocke
 func (m *awsRestjson1_serializeOpDeleteCredentialLocker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1512,7 +1467,6 @@ func (m *awsRestjson1_serializeOpDeleteCredentialLocke
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCredentialLockerInput(v *DeleteCredentialLockerInput, encoder *httpbinding.Encoder) error {
@@ -1542,10 +1496,8 @@ func (m *awsRestjson1_serializeOpDeleteDestination) Ha
 func (m *awsRestjson1_serializeOpDeleteDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1583,7 +1535,6 @@ func (m *awsRestjson1_serializeOpDeleteDestination) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDestinationInput(v *DeleteDestinationInput, encoder *httpbinding.Encoder) error {
@@ -1613,10 +1564,8 @@ func (m *awsRestjson1_serializeOpDeleteEventLogConfigu
 func (m *awsRestjson1_serializeOpDeleteEventLogConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1654,7 +1603,6 @@ func (m *awsRestjson1_serializeOpDeleteEventLogConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventLogConfigurationInput(v *DeleteEventLogConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1684,10 +1632,8 @@ func (m *awsRestjson1_serializeOpDeleteManagedThing) H
 func (m *awsRestjson1_serializeOpDeleteManagedThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1725,7 +1671,6 @@ func (m *awsRestjson1_serializeOpDeleteManagedThing) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteManagedThingInput(v *DeleteManagedThingInput, encoder *httpbinding.Encoder) error {
@@ -1759,10 +1704,8 @@ func (m *awsRestjson1_serializeOpDeleteNotificationCon
 func (m *awsRestjson1_serializeOpDeleteNotificationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1800,7 +1743,6 @@ func (m *awsRestjson1_serializeOpDeleteNotificationCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteNotificationConfigurationInput(v *DeleteNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1830,10 +1772,8 @@ func (m *awsRestjson1_serializeOpDeleteOtaTask) Handle
 func (m *awsRestjson1_serializeOpDeleteOtaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1871,7 +1811,6 @@ func (m *awsRestjson1_serializeOpDeleteOtaTask) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteOtaTaskInput(v *DeleteOtaTaskInput, encoder *httpbinding.Encoder) error {
@@ -1901,10 +1840,8 @@ func (m *awsRestjson1_serializeOpDeleteOtaTaskConfigur
 func (m *awsRestjson1_serializeOpDeleteOtaTaskConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1942,7 +1879,6 @@ func (m *awsRestjson1_serializeOpDeleteOtaTaskConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteOtaTaskConfigurationInput(v *DeleteOtaTaskConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1972,10 +1908,8 @@ func (m *awsRestjson1_serializeOpDeleteProvisioningPro
 func (m *awsRestjson1_serializeOpDeleteProvisioningProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2013,7 +1947,6 @@ func (m *awsRestjson1_serializeOpDeleteProvisioningPro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteProvisioningProfileInput(v *DeleteProvisioningProfileInput, encoder *httpbinding.Encoder) error {
@@ -2043,10 +1976,8 @@ func (m *awsRestjson1_serializeOpDeregisterAccountAsso
 func (m *awsRestjson1_serializeOpDeregisterAccountAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2091,7 +2022,6 @@ func (m *awsRestjson1_serializeOpDeregisterAccountAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeregisterAccountAssociationInput(v *DeregisterAccountAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2129,10 +2059,8 @@ func (m *awsRestjson1_serializeOpGetAccountAssociation
 func (m *awsRestjson1_serializeOpGetAccountAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2170,7 +2098,6 @@ func (m *awsRestjson1_serializeOpGetAccountAssociation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountAssociationInput(v *GetAccountAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2200,10 +2127,8 @@ func (m *awsRestjson1_serializeOpGetCloudConnector) Ha
 func (m *awsRestjson1_serializeOpGetCloudConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2241,7 +2166,6 @@ func (m *awsRestjson1_serializeOpGetCloudConnector) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCloudConnectorInput(v *GetCloudConnectorInput, encoder *httpbinding.Encoder) error {
@@ -2271,10 +2195,8 @@ func (m *awsRestjson1_serializeOpGetConnectorDestinati
 func (m *awsRestjson1_serializeOpGetConnectorDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2312,7 +2234,6 @@ func (m *awsRestjson1_serializeOpGetConnectorDestinati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectorDestinationInput(v *GetConnectorDestinationInput, encoder *httpbinding.Encoder) error {
@@ -2342,10 +2263,8 @@ func (m *awsRestjson1_serializeOpGetCredentialLocker) 
 func (m *awsRestjson1_serializeOpGetCredentialLocker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2383,7 +2302,6 @@ func (m *awsRestjson1_serializeOpGetCredentialLocker) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCredentialLockerInput(v *GetCredentialLockerInput, encoder *httpbinding.Encoder) error {
@@ -2413,10 +2331,8 @@ func (m *awsRestjson1_serializeOpGetCustomEndpoint) Ha
 func (m *awsRestjson1_serializeOpGetCustomEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2450,7 +2366,6 @@ func (m *awsRestjson1_serializeOpGetCustomEndpoint) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCustomEndpointInput(v *GetCustomEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2471,10 +2386,8 @@ func (m *awsRestjson1_serializeOpGetDefaultEncryptionC
 func (m *awsRestjson1_serializeOpGetDefaultEncryptionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2508,7 +2421,6 @@ func (m *awsRestjson1_serializeOpGetDefaultEncryptionC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDefaultEncryptionConfigurationInput(v *GetDefaultEncryptionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2529,10 +2441,8 @@ func (m *awsRestjson1_serializeOpGetDestination) Handl
 func (m *awsRestjson1_serializeOpGetDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2570,7 +2480,6 @@ func (m *awsRestjson1_serializeOpGetDestination) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDestinationInput(v *GetDestinationInput, encoder *httpbinding.Encoder) error {
@@ -2600,10 +2509,8 @@ func (m *awsRestjson1_serializeOpGetDeviceDiscovery) H
 func (m *awsRestjson1_serializeOpGetDeviceDiscovery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2641,7 +2548,6 @@ func (m *awsRestjson1_serializeOpGetDeviceDiscovery) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeviceDiscoveryInput(v *GetDeviceDiscoveryInput, encoder *httpbinding.Encoder) error {
@@ -2671,10 +2577,8 @@ func (m *awsRestjson1_serializeOpGetEventLogConfigurat
 func (m *awsRestjson1_serializeOpGetEventLogConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2712,7 +2616,6 @@ func (m *awsRestjson1_serializeOpGetEventLogConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventLogConfigurationInput(v *GetEventLogConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2742,10 +2645,8 @@ func (m *awsRestjson1_serializeOpGetHubConfiguration) 
 func (m *awsRestjson1_serializeOpGetHubConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2779,7 +2680,6 @@ func (m *awsRestjson1_serializeOpGetHubConfiguration) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetHubConfigurationInput(v *GetHubConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2800,10 +2700,8 @@ func (m *awsRestjson1_serializeOpGetManagedThing) Hand
 func (m *awsRestjson1_serializeOpGetManagedThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2841,7 +2739,6 @@ func (m *awsRestjson1_serializeOpGetManagedThing) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetManagedThingInput(v *GetManagedThingInput, encoder *httpbinding.Encoder) error {
@@ -2871,10 +2768,8 @@ func (m *awsRestjson1_serializeOpGetManagedThingCapabi
 func (m *awsRestjson1_serializeOpGetManagedThingCapabilities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2912,7 +2807,6 @@ func (m *awsRestjson1_serializeOpGetManagedThingCapabi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetManagedThingCapabilitiesInput(v *GetManagedThingCapabilitiesInput, encoder *httpbinding.Encoder) error {
@@ -2942,10 +2836,8 @@ func (m *awsRestjson1_serializeOpGetManagedThingConnec
 func (m *awsRestjson1_serializeOpGetManagedThingConnectivityData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2983,7 +2875,6 @@ func (m *awsRestjson1_serializeOpGetManagedThingConnec
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetManagedThingConnectivityDataInput(v *GetManagedThingConnectivityDataInput, encoder *httpbinding.Encoder) error {
@@ -3013,10 +2904,8 @@ func (m *awsRestjson1_serializeOpGetManagedThingMetaDa
 func (m *awsRestjson1_serializeOpGetManagedThingMetaData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3054,7 +2943,6 @@ func (m *awsRestjson1_serializeOpGetManagedThingMetaDa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetManagedThingMetaDataInput(v *GetManagedThingMetaDataInput, encoder *httpbinding.Encoder) error {
@@ -3084,10 +2972,8 @@ func (m *awsRestjson1_serializeOpGetManagedThingState)
 func (m *awsRestjson1_serializeOpGetManagedThingState) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3125,7 +3011,6 @@ func (m *awsRestjson1_serializeOpGetManagedThingState)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetManagedThingStateInput(v *GetManagedThingStateInput, encoder *httpbinding.Encoder) error {
@@ -3155,10 +3040,8 @@ func (m *awsRestjson1_serializeOpGetNotificationConfig
 func (m *awsRestjson1_serializeOpGetNotificationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3196,7 +3079,6 @@ func (m *awsRestjson1_serializeOpGetNotificationConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNotificationConfigurationInput(v *GetNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3226,10 +3108,8 @@ func (m *awsRestjson1_serializeOpGetOtaTask) HandleSer
 func (m *awsRestjson1_serializeOpGetOtaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3267,7 +3147,6 @@ func (m *awsRestjson1_serializeOpGetOtaTask) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOtaTaskInput(v *GetOtaTaskInput, encoder *httpbinding.Encoder) error {
@@ -3297,10 +3176,8 @@ func (m *awsRestjson1_serializeOpGetOtaTaskConfigurati
 func (m *awsRestjson1_serializeOpGetOtaTaskConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3338,7 +3215,6 @@ func (m *awsRestjson1_serializeOpGetOtaTaskConfigurati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOtaTaskConfigurationInput(v *GetOtaTaskConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3368,10 +3244,8 @@ func (m *awsRestjson1_serializeOpGetProvisioningProfil
 func (m *awsRestjson1_serializeOpGetProvisioningProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3409,7 +3283,6 @@ func (m *awsRestjson1_serializeOpGetProvisioningProfil
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProvisioningProfileInput(v *GetProvisioningProfileInput, encoder *httpbinding.Encoder) error {
@@ -3439,10 +3312,8 @@ func (m *awsRestjson1_serializeOpGetRuntimeLogConfigur
 func (m *awsRestjson1_serializeOpGetRuntimeLogConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3480,7 +3351,6 @@ func (m *awsRestjson1_serializeOpGetRuntimeLogConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRuntimeLogConfigurationInput(v *GetRuntimeLogConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3510,10 +3380,8 @@ func (m *awsRestjson1_serializeOpGetSchemaVersion) Han
 func (m *awsRestjson1_serializeOpGetSchemaVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3551,7 +3419,6 @@ func (m *awsRestjson1_serializeOpGetSchemaVersion) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSchemaVersionInput(v *GetSchemaVersionInput, encoder *httpbinding.Encoder) error {
@@ -3594,10 +3461,8 @@ func (m *awsRestjson1_serializeOpListAccountAssociatio
 func (m *awsRestjson1_serializeOpListAccountAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3635,7 +3500,6 @@ func (m *awsRestjson1_serializeOpListAccountAssociatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAccountAssociationsInput(v *ListAccountAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -3668,10 +3532,8 @@ func (m *awsRestjson1_serializeOpListCloudConnectors) 
 func (m *awsRestjson1_serializeOpListCloudConnectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3709,7 +3571,6 @@ func (m *awsRestjson1_serializeOpListCloudConnectors) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCloudConnectorsInput(v *ListCloudConnectorsInput, encoder *httpbinding.Encoder) error {
@@ -3746,10 +3607,8 @@ func (m *awsRestjson1_serializeOpListConnectorDestinat
 func (m *awsRestjson1_serializeOpListConnectorDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3787,7 +3646,6 @@ func (m *awsRestjson1_serializeOpListConnectorDestinat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConnectorDestinationsInput(v *ListConnectorDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -3820,10 +3678,8 @@ func (m *awsRestjson1_serializeOpListCredentialLockers
 func (m *awsRestjson1_serializeOpListCredentialLockers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3861,7 +3717,6 @@ func (m *awsRestjson1_serializeOpListCredentialLockers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCredentialLockersInput(v *ListCredentialLockersInput, encoder *httpbinding.Encoder) error {
@@ -3890,10 +3745,8 @@ func (m *awsRestjson1_serializeOpListDestinations) Han
 func (m *awsRestjson1_serializeOpListDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3931,7 +3784,6 @@ func (m *awsRestjson1_serializeOpListDestinations) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDestinationsInput(v *ListDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -3960,10 +3812,8 @@ func (m *awsRestjson1_serializeOpListDeviceDiscoveries
 func (m *awsRestjson1_serializeOpListDeviceDiscoveries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4001,7 +3851,6 @@ func (m *awsRestjson1_serializeOpListDeviceDiscoveries
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDeviceDiscoveriesInput(v *ListDeviceDiscoveriesInput, encoder *httpbinding.Encoder) error {
@@ -4038,10 +3887,8 @@ func (m *awsRestjson1_serializeOpListDiscoveredDevices
 func (m *awsRestjson1_serializeOpListDiscoveredDevices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4079,7 +3926,6 @@ func (m *awsRestjson1_serializeOpListDiscoveredDevices
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDiscoveredDevicesInput(v *ListDiscoveredDevicesInput, encoder *httpbinding.Encoder) error {
@@ -4117,10 +3963,8 @@ func (m *awsRestjson1_serializeOpListEventLogConfigura
 func (m *awsRestjson1_serializeOpListEventLogConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4158,7 +4002,6 @@ func (m *awsRestjson1_serializeOpListEventLogConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEventLogConfigurationsInput(v *ListEventLogConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -4187,10 +4030,8 @@ func (m *awsRestjson1_serializeOpListManagedThingAccou
 func (m *awsRestjson1_serializeOpListManagedThingAccountAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4228,7 +4069,6 @@ func (m *awsRestjson1_serializeOpListManagedThingAccou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListManagedThingAccountAssociationsInput(v *ListManagedThingAccountAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4265,10 +4105,8 @@ func (m *awsRestjson1_serializeOpListManagedThings) Ha
 func (m *awsRestjson1_serializeOpListManagedThings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4306,7 +4144,6 @@ func (m *awsRestjson1_serializeOpListManagedThings) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListManagedThingsInput(v *ListManagedThingsInput, encoder *httpbinding.Encoder) error {
@@ -4371,10 +4208,8 @@ func (m *awsRestjson1_serializeOpListManagedThingSchem
 func (m *awsRestjson1_serializeOpListManagedThingSchemas) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4412,7 +4247,6 @@ func (m *awsRestjson1_serializeOpListManagedThingSchem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListManagedThingSchemasInput(v *ListManagedThingSchemasInput, encoder *httpbinding.Encoder) error {
@@ -4458,10 +4292,8 @@ func (m *awsRestjson1_serializeOpListNotificationConfi
 func (m *awsRestjson1_serializeOpListNotificationConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4499,7 +4331,6 @@ func (m *awsRestjson1_serializeOpListNotificationConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNotificationConfigurationsInput(v *ListNotificationConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -4528,10 +4359,8 @@ func (m *awsRestjson1_serializeOpListOtaTaskConfigurat
 func (m *awsRestjson1_serializeOpListOtaTaskConfigurations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4569,7 +4398,6 @@ func (m *awsRestjson1_serializeOpListOtaTaskConfigurat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOtaTaskConfigurationsInput(v *ListOtaTaskConfigurationsInput, encoder *httpbinding.Encoder) error {
@@ -4598,10 +4426,8 @@ func (m *awsRestjson1_serializeOpListOtaTaskExecutions
 func (m *awsRestjson1_serializeOpListOtaTaskExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4639,7 +4465,6 @@ func (m *awsRestjson1_serializeOpListOtaTaskExecutions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOtaTaskExecutionsInput(v *ListOtaTaskExecutionsInput, encoder *httpbinding.Encoder) error {
@@ -4677,10 +4502,8 @@ func (m *awsRestjson1_serializeOpListOtaTasks) HandleS
 func (m *awsRestjson1_serializeOpListOtaTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4718,7 +4541,6 @@ func (m *awsRestjson1_serializeOpListOtaTasks) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOtaTasksInput(v *ListOtaTasksInput, encoder *httpbinding.Encoder) error {
@@ -4747,10 +4569,8 @@ func (m *awsRestjson1_serializeOpListProvisioningProfi
 func (m *awsRestjson1_serializeOpListProvisioningProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4788,7 +4608,6 @@ func (m *awsRestjson1_serializeOpListProvisioningProfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListProvisioningProfilesInput(v *ListProvisioningProfilesInput, encoder *httpbinding.Encoder) error {
@@ -4817,10 +4636,8 @@ func (m *awsRestjson1_serializeOpListSchemaVersions) H
 func (m *awsRestjson1_serializeOpListSchemaVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4858,7 +4675,6 @@ func (m *awsRestjson1_serializeOpListSchemaVersions) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSchemaVersionsInput(v *ListSchemaVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4912,10 +4728,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -4953,7 +4767,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4983,10 +4796,8 @@ func (m *awsRestjson1_serializeOpPutDefaultEncryptionC
 func (m *awsRestjson1_serializeOpPutDefaultEncryptionConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5031,7 +4842,6 @@ func (m *awsRestjson1_serializeOpPutDefaultEncryptionC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDefaultEncryptionConfigurationInput(v *PutDefaultEncryptionConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5069,10 +4879,8 @@ func (m *awsRestjson1_serializeOpPutHubConfiguration) 
 func (m *awsRestjson1_serializeOpPutHubConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5117,7 +4925,6 @@ func (m *awsRestjson1_serializeOpPutHubConfiguration) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutHubConfigurationInput(v *PutHubConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5150,10 +4957,8 @@ func (m *awsRestjson1_serializeOpPutRuntimeLogConfigur
 func (m *awsRestjson1_serializeOpPutRuntimeLogConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5202,7 +5007,6 @@ func (m *awsRestjson1_serializeOpPutRuntimeLogConfigur
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutRuntimeLogConfigurationInput(v *PutRuntimeLogConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5246,10 +5050,8 @@ func (m *awsRestjson1_serializeOpRegisterAccountAssoci
 func (m *awsRestjson1_serializeOpRegisterAccountAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5294,7 +5096,6 @@ func (m *awsRestjson1_serializeOpRegisterAccountAssoci
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterAccountAssociationInput(v *RegisterAccountAssociationInput, encoder *httpbinding.Encoder) error {
@@ -5337,10 +5138,8 @@ func (m *awsRestjson1_serializeOpRegisterCustomEndpoin
 func (m *awsRestjson1_serializeOpRegisterCustomEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5374,7 +5173,6 @@ func (m *awsRestjson1_serializeOpRegisterCustomEndpoin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterCustomEndpointInput(v *RegisterCustomEndpointInput, encoder *httpbinding.Encoder) error {
@@ -5395,10 +5193,8 @@ func (m *awsRestjson1_serializeOpResetRuntimeLogConfig
 func (m *awsRestjson1_serializeOpResetRuntimeLogConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5436,7 +5232,6 @@ func (m *awsRestjson1_serializeOpResetRuntimeLogConfig
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResetRuntimeLogConfigurationInput(v *ResetRuntimeLogConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -5466,10 +5261,8 @@ func (m *awsRestjson1_serializeOpSendConnectorEvent) H
 func (m *awsRestjson1_serializeOpSendConnectorEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5518,7 +5311,6 @@ func (m *awsRestjson1_serializeOpSendConnectorEvent) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendConnectorEventInput(v *SendConnectorEventInput, encoder *httpbinding.Encoder) error {
@@ -5609,10 +5401,8 @@ func (m *awsRestjson1_serializeOpSendManagedThingComma
 func (m *awsRestjson1_serializeOpSendManagedThingCommand) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5661,7 +5451,6 @@ func (m *awsRestjson1_serializeOpSendManagedThingComma
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendManagedThingCommandInput(v *SendManagedThingCommandInput, encoder *httpbinding.Encoder) error {
@@ -5715,10 +5504,8 @@ func (m *awsRestjson1_serializeOpStartAccountAssociati
 func (m *awsRestjson1_serializeOpStartAccountAssociationRefresh) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5756,7 +5543,6 @@ func (m *awsRestjson1_serializeOpStartAccountAssociati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAccountAssociationRefreshInput(v *StartAccountAssociationRefreshInput, encoder *httpbinding.Encoder) error {
@@ -5786,10 +5572,8 @@ func (m *awsRestjson1_serializeOpStartDeviceDiscovery)
 func (m *awsRestjson1_serializeOpStartDeviceDiscovery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5834,7 +5618,6 @@ func (m *awsRestjson1_serializeOpStartDeviceDiscovery)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDeviceDiscoveryInput(v *StartDeviceDiscoveryInput, encoder *httpbinding.Encoder) error {
@@ -5911,10 +5694,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5963,7 +5744,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6007,10 +5787,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -6048,7 +5826,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6084,10 +5861,8 @@ func (m *awsRestjson1_serializeOpUpdateAccountAssociat
 func (m *awsRestjson1_serializeOpUpdateAccountAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6136,7 +5911,6 @@ func (m *awsRestjson1_serializeOpUpdateAccountAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountAssociationInput(v *UpdateAccountAssociationInput, encoder *httpbinding.Encoder) error {
@@ -6183,10 +5957,8 @@ func (m *awsRestjson1_serializeOpUpdateCloudConnector)
 func (m *awsRestjson1_serializeOpUpdateCloudConnector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6235,7 +6007,6 @@ func (m *awsRestjson1_serializeOpUpdateCloudConnector)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCloudConnectorInput(v *UpdateCloudConnectorInput, encoder *httpbinding.Encoder) error {
@@ -6282,10 +6053,8 @@ func (m *awsRestjson1_serializeOpUpdateConnectorDestin
 func (m *awsRestjson1_serializeOpUpdateConnectorDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6334,7 +6103,6 @@ func (m *awsRestjson1_serializeOpUpdateConnectorDestin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConnectorDestinationInput(v *UpdateConnectorDestinationInput, encoder *httpbinding.Encoder) error {
@@ -6400,10 +6168,8 @@ func (m *awsRestjson1_serializeOpUpdateDestination) Ha
 func (m *awsRestjson1_serializeOpUpdateDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6452,7 +6218,6 @@ func (m *awsRestjson1_serializeOpUpdateDestination) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDestinationInput(v *UpdateDestinationInput, encoder *httpbinding.Encoder) error {
@@ -6509,10 +6274,8 @@ func (m *awsRestjson1_serializeOpUpdateEventLogConfigu
 func (m *awsRestjson1_serializeOpUpdateEventLogConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6561,7 +6324,6 @@ func (m *awsRestjson1_serializeOpUpdateEventLogConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateEventLogConfigurationInput(v *UpdateEventLogConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6603,10 +6365,8 @@ func (m *awsRestjson1_serializeOpUpdateManagedThing) H
 func (m *awsRestjson1_serializeOpUpdateManagedThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6655,7 +6415,6 @@ func (m *awsRestjson1_serializeOpUpdateManagedThing) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateManagedThingInput(v *UpdateManagedThingInput, encoder *httpbinding.Encoder) error {
@@ -6758,10 +6517,8 @@ func (m *awsRestjson1_serializeOpUpdateNotificationCon
 func (m *awsRestjson1_serializeOpUpdateNotificationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6810,7 +6567,6 @@ func (m *awsRestjson1_serializeOpUpdateNotificationCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateNotificationConfigurationInput(v *UpdateNotificationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6852,10 +6608,8 @@ func (m *awsRestjson1_serializeOpUpdateOtaTask) Handle
 func (m *awsRestjson1_serializeOpUpdateOtaTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6904,7 +6658,6 @@ func (m *awsRestjson1_serializeOpUpdateOtaTask) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOtaTaskInput(v *UpdateOtaTaskInput, encoder *httpbinding.Encoder) error {
