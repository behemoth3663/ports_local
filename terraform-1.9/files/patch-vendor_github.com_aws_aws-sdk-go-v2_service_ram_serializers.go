--- vendor/github.com/aws/aws-sdk-go-v2/service/ram/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ram/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAcceptResourceShareIn
 func (m *awsRestjson1_serializeOpAcceptResourceShareInvitation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -73,7 +70,6 @@ func (m *awsRestjson1_serializeOpAcceptResourceShareIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptResourceShareInvitationInput(v *AcceptResourceShareInvitationInput, encoder *httpbinding.Encoder) error {
@@ -111,10 +107,8 @@ func (m *awsRestjson1_serializeOpAssociateResourceShar
 func (m *awsRestjson1_serializeOpAssociateResourceShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -159,7 +153,6 @@ func (m *awsRestjson1_serializeOpAssociateResourceShar
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateResourceShareInput(v *AssociateResourceShareInput, encoder *httpbinding.Encoder) error {
@@ -218,10 +211,8 @@ func (m *awsRestjson1_serializeOpAssociateResourceShar
 func (m *awsRestjson1_serializeOpAssociateResourceSharePermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -266,7 +257,6 @@ func (m *awsRestjson1_serializeOpAssociateResourceShar
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateResourceSharePermissionInput(v *AssociateResourceSharePermissionInput, encoder *httpbinding.Encoder) error {
@@ -319,10 +309,8 @@ func (m *awsRestjson1_serializeOpCreatePermission) Han
 func (m *awsRestjson1_serializeOpCreatePermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -367,7 +355,6 @@ func (m *awsRestjson1_serializeOpCreatePermission) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePermissionInput(v *CreatePermissionInput, encoder *httpbinding.Encoder) error {
@@ -422,10 +409,8 @@ func (m *awsRestjson1_serializeOpCreatePermissionVersi
 func (m *awsRestjson1_serializeOpCreatePermissionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -470,7 +455,6 @@ func (m *awsRestjson1_serializeOpCreatePermissionVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePermissionVersionInput(v *CreatePermissionVersionInput, encoder *httpbinding.Encoder) error {
@@ -513,10 +497,8 @@ func (m *awsRestjson1_serializeOpCreateResourceShare) 
 func (m *awsRestjson1_serializeOpCreateResourceShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -561,7 +543,6 @@ func (m *awsRestjson1_serializeOpCreateResourceShare) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateResourceShareInput(v *CreateResourceShareInput, encoder *httpbinding.Encoder) error {
@@ -639,10 +620,8 @@ func (m *awsRestjson1_serializeOpDeletePermission) Han
 func (m *awsRestjson1_serializeOpDeletePermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -680,7 +659,6 @@ func (m *awsRestjson1_serializeOpDeletePermission) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePermissionInput(v *DeletePermissionInput, encoder *httpbinding.Encoder) error {
@@ -709,10 +687,8 @@ func (m *awsRestjson1_serializeOpDeletePermissionVersi
 func (m *awsRestjson1_serializeOpDeletePermissionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -750,7 +726,6 @@ func (m *awsRestjson1_serializeOpDeletePermissionVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePermissionVersionInput(v *DeletePermissionVersionInput, encoder *httpbinding.Encoder) error {
@@ -783,10 +758,8 @@ func (m *awsRestjson1_serializeOpDeleteResourceShare) 
 func (m *awsRestjson1_serializeOpDeleteResourceShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -824,7 +797,6 @@ func (m *awsRestjson1_serializeOpDeleteResourceShare) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResourceShareInput(v *DeleteResourceShareInput, encoder *httpbinding.Encoder) error {
@@ -853,10 +825,8 @@ func (m *awsRestjson1_serializeOpDisassociateResourceS
 func (m *awsRestjson1_serializeOpDisassociateResourceShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -901,7 +871,6 @@ func (m *awsRestjson1_serializeOpDisassociateResourceS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateResourceShareInput(v *DisassociateResourceShareInput, encoder *httpbinding.Encoder) error {
@@ -960,10 +929,8 @@ func (m *awsRestjson1_serializeOpDisassociateResourceS
 func (m *awsRestjson1_serializeOpDisassociateResourceSharePermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1008,7 +975,6 @@ func (m *awsRestjson1_serializeOpDisassociateResourceS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateResourceSharePermissionInput(v *DisassociateResourceSharePermissionInput, encoder *httpbinding.Encoder) error {
@@ -1051,10 +1017,8 @@ func (m *awsRestjson1_serializeOpEnableSharingWithAwsO
 func (m *awsRestjson1_serializeOpEnableSharingWithAwsOrganization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1088,7 +1052,6 @@ func (m *awsRestjson1_serializeOpEnableSharingWithAwsO
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableSharingWithAwsOrganizationInput(v *EnableSharingWithAwsOrganizationInput, encoder *httpbinding.Encoder) error {
@@ -1109,10 +1072,8 @@ func (m *awsRestjson1_serializeOpGetPermission) Handle
 func (m *awsRestjson1_serializeOpGetPermission) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1157,7 +1118,6 @@ func (m *awsRestjson1_serializeOpGetPermission) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPermissionInput(v *GetPermissionInput, encoder *httpbinding.Encoder) error {
@@ -1195,10 +1155,8 @@ func (m *awsRestjson1_serializeOpGetResourcePolicies) 
 func (m *awsRestjson1_serializeOpGetResourcePolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1243,7 +1201,6 @@ func (m *awsRestjson1_serializeOpGetResourcePolicies) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourcePoliciesInput(v *GetResourcePoliciesInput, encoder *httpbinding.Encoder) error {
@@ -1293,10 +1250,8 @@ func (m *awsRestjson1_serializeOpGetResourceShareAssoc
 func (m *awsRestjson1_serializeOpGetResourceShareAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1341,7 +1296,6 @@ func (m *awsRestjson1_serializeOpGetResourceShareAssoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceShareAssociationsInput(v *GetResourceShareAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -1406,10 +1360,8 @@ func (m *awsRestjson1_serializeOpGetResourceShareInvit
 func (m *awsRestjson1_serializeOpGetResourceShareInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1454,7 +1406,6 @@ func (m *awsRestjson1_serializeOpGetResourceShareInvit
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceShareInvitationsInput(v *GetResourceShareInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -1506,10 +1457,8 @@ func (m *awsRestjson1_serializeOpGetResourceShares) Ha
 func (m *awsRestjson1_serializeOpGetResourceShares) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1554,7 +1503,6 @@ func (m *awsRestjson1_serializeOpGetResourceShares) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourceSharesInput(v *GetResourceSharesInput, encoder *httpbinding.Encoder) error {
@@ -1631,10 +1579,8 @@ func (m *awsRestjson1_serializeOpListPendingInvitation
 func (m *awsRestjson1_serializeOpListPendingInvitationResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1679,7 +1625,6 @@ func (m *awsRestjson1_serializeOpListPendingInvitation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPendingInvitationResourcesInput(v *ListPendingInvitationResourcesInput, encoder *httpbinding.Encoder) error {
@@ -1727,10 +1672,8 @@ func (m *awsRestjson1_serializeOpListPermissionAssocia
 func (m *awsRestjson1_serializeOpListPermissionAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1775,7 +1718,6 @@ func (m *awsRestjson1_serializeOpListPermissionAssocia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPermissionAssociationsInput(v *ListPermissionAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -1843,10 +1785,8 @@ func (m *awsRestjson1_serializeOpListPermissions) Hand
 func (m *awsRestjson1_serializeOpListPermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1891,7 +1831,6 @@ func (m *awsRestjson1_serializeOpListPermissions) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPermissionsInput(v *ListPermissionsInput, encoder *httpbinding.Encoder) error {
@@ -1939,10 +1878,8 @@ func (m *awsRestjson1_serializeOpListPermissionVersion
 func (m *awsRestjson1_serializeOpListPermissionVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1987,7 +1924,6 @@ func (m *awsRestjson1_serializeOpListPermissionVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPermissionVersionsInput(v *ListPermissionVersionsInput, encoder *httpbinding.Encoder) error {
@@ -2030,10 +1966,8 @@ func (m *awsRestjson1_serializeOpListPrincipals) Handl
 func (m *awsRestjson1_serializeOpListPrincipals) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2078,7 +2012,6 @@ func (m *awsRestjson1_serializeOpListPrincipals) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPrincipalsInput(v *ListPrincipalsInput, encoder *httpbinding.Encoder) error {
@@ -2145,10 +2078,8 @@ func (m *awsRestjson1_serializeOpListReplacePermission
 func (m *awsRestjson1_serializeOpListReplacePermissionAssociationsWork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2193,7 +2124,6 @@ func (m *awsRestjson1_serializeOpListReplacePermission
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReplacePermissionAssociationsWorkInput(v *ListReplacePermissionAssociationsWorkInput, encoder *httpbinding.Encoder) error {
@@ -2243,10 +2173,8 @@ func (m *awsRestjson1_serializeOpListResources) Handle
 func (m *awsRestjson1_serializeOpListResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2291,7 +2219,6 @@ func (m *awsRestjson1_serializeOpListResources) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourcesInput(v *ListResourcesInput, encoder *httpbinding.Encoder) error {
@@ -2363,10 +2290,8 @@ func (m *awsRestjson1_serializeOpListResourceSharePerm
 func (m *awsRestjson1_serializeOpListResourceSharePermissions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2411,7 +2336,6 @@ func (m *awsRestjson1_serializeOpListResourceSharePerm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceSharePermissionsInput(v *ListResourceSharePermissionsInput, encoder *httpbinding.Encoder) error {
@@ -2454,10 +2378,8 @@ func (m *awsRestjson1_serializeOpListResourceTypes) Ha
 func (m *awsRestjson1_serializeOpListResourceTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2502,7 +2424,6 @@ func (m *awsRestjson1_serializeOpListResourceTypes) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceTypesInput(v *ListResourceTypesInput, encoder *httpbinding.Encoder) error {
@@ -2545,10 +2466,8 @@ func (m *awsRestjson1_serializeOpPromotePermissionCrea
 func (m *awsRestjson1_serializeOpPromotePermissionCreatedFromPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2593,7 +2512,6 @@ func (m *awsRestjson1_serializeOpPromotePermissionCrea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPromotePermissionCreatedFromPolicyInput(v *PromotePermissionCreatedFromPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2636,10 +2554,8 @@ func (m *awsRestjson1_serializeOpPromoteResourceShareC
 func (m *awsRestjson1_serializeOpPromoteResourceShareCreatedFromPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2677,7 +2593,6 @@ func (m *awsRestjson1_serializeOpPromoteResourceShareC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPromoteResourceShareCreatedFromPolicyInput(v *PromoteResourceShareCreatedFromPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2702,10 +2617,8 @@ func (m *awsRestjson1_serializeOpRejectResourceShareIn
 func (m *awsRestjson1_serializeOpRejectResourceShareInvitation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2750,7 +2663,6 @@ func (m *awsRestjson1_serializeOpRejectResourceShareIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectResourceShareInvitationInput(v *RejectResourceShareInvitationInput, encoder *httpbinding.Encoder) error {
@@ -2788,10 +2700,8 @@ func (m *awsRestjson1_serializeOpReplacePermissionAsso
 func (m *awsRestjson1_serializeOpReplacePermissionAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2836,7 +2746,6 @@ func (m *awsRestjson1_serializeOpReplacePermissionAsso
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsReplacePermissionAssociationsInput(v *ReplacePermissionAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -2884,10 +2793,8 @@ func (m *awsRestjson1_serializeOpSetDefaultPermissionV
 func (m *awsRestjson1_serializeOpSetDefaultPermissionVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2932,7 +2839,6 @@ func (m *awsRestjson1_serializeOpSetDefaultPermissionV
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSetDefaultPermissionVersionInput(v *SetDefaultPermissionVersionInput, encoder *httpbinding.Encoder) error {
@@ -2975,10 +2881,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -3023,7 +2927,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3068,10 +2971,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -3116,7 +3017,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3161,10 +3061,8 @@ func (m *awsRestjson1_serializeOpUpdateResourceShare) 
 func (m *awsRestjson1_serializeOpUpdateResourceShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3209,7 +3107,6 @@ func (m *awsRestjson1_serializeOpUpdateResourceShare) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateResourceShareInput(v *UpdateResourceShareInput, encoder *httpbinding.Encoder) error {
