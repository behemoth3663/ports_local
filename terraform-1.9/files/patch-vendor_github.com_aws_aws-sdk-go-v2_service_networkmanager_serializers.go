--- vendor/github.com/aws/aws-sdk-go-v2/service/networkmanager/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/networkmanager/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAcceptAttachment) Han
 func (m *awsRestjson1_serializeOpAcceptAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpAcceptAttachment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptAttachmentInput(v *AcceptAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -96,10 +92,8 @@ func (m *awsRestjson1_serializeOpAssociateConnectPeer)
 func (m *awsRestjson1_serializeOpAssociateConnectPeer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -148,7 +142,6 @@ func (m *awsRestjson1_serializeOpAssociateConnectPeer)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateConnectPeerInput(v *AssociateConnectPeerInput, encoder *httpbinding.Encoder) error {
@@ -200,10 +193,8 @@ func (m *awsRestjson1_serializeOpAssociateCustomerGate
 func (m *awsRestjson1_serializeOpAssociateCustomerGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -252,7 +243,6 @@ func (m *awsRestjson1_serializeOpAssociateCustomerGate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateCustomerGatewayInput(v *AssociateCustomerGatewayInput, encoder *httpbinding.Encoder) error {
@@ -304,10 +294,8 @@ func (m *awsRestjson1_serializeOpAssociateLink) Handle
 func (m *awsRestjson1_serializeOpAssociateLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -356,7 +344,6 @@ func (m *awsRestjson1_serializeOpAssociateLink) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateLinkInput(v *AssociateLinkInput, encoder *httpbinding.Encoder) error {
@@ -403,10 +390,8 @@ func (m *awsRestjson1_serializeOpAssociateTransitGatew
 func (m *awsRestjson1_serializeOpAssociateTransitGatewayConnectPeer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -455,7 +440,6 @@ func (m *awsRestjson1_serializeOpAssociateTransitGatew
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateTransitGatewayConnectPeerInput(v *AssociateTransitGatewayConnectPeerInput, encoder *httpbinding.Encoder) error {
@@ -507,10 +491,8 @@ func (m *awsRestjson1_serializeOpCreateConnectAttachme
 func (m *awsRestjson1_serializeOpCreateConnectAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -555,7 +537,6 @@ func (m *awsRestjson1_serializeOpCreateConnectAttachme
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectAttachmentInput(v *CreateConnectAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -617,10 +598,8 @@ func (m *awsRestjson1_serializeOpCreateConnection) Han
 func (m *awsRestjson1_serializeOpCreateConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -669,7 +648,6 @@ func (m *awsRestjson1_serializeOpCreateConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectionInput(v *CreateConnectionInput, encoder *httpbinding.Encoder) error {
@@ -738,10 +716,8 @@ func (m *awsRestjson1_serializeOpCreateConnectPeer) Ha
 func (m *awsRestjson1_serializeOpCreateConnectPeer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -786,7 +762,6 @@ func (m *awsRestjson1_serializeOpCreateConnectPeer) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConnectPeerInput(v *CreateConnectPeerInput, encoder *httpbinding.Encoder) error {
@@ -860,10 +835,8 @@ func (m *awsRestjson1_serializeOpCreateCoreNetwork) Ha
 func (m *awsRestjson1_serializeOpCreateCoreNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -908,7 +881,6 @@ func (m *awsRestjson1_serializeOpCreateCoreNetwork) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateCoreNetworkInput(v *CreateCoreNetworkInput, encoder *httpbinding.Encoder) error {
@@ -963,10 +935,8 @@ func (m *awsRestjson1_serializeOpCreateDevice) HandleS
 func (m *awsRestjson1_serializeOpCreateDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1015,7 +985,6 @@ func (m *awsRestjson1_serializeOpCreateDevice) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeviceInput(v *CreateDeviceInput, encoder *httpbinding.Encoder) error {
@@ -1103,10 +1072,8 @@ func (m *awsRestjson1_serializeOpCreateDirectConnectGa
 func (m *awsRestjson1_serializeOpCreateDirectConnectGatewayAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1151,7 +1118,6 @@ func (m *awsRestjson1_serializeOpCreateDirectConnectGa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDirectConnectGatewayAttachmentInput(v *CreateDirectConnectGatewayAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -1208,10 +1174,8 @@ func (m *awsRestjson1_serializeOpCreateGlobalNetwork) 
 func (m *awsRestjson1_serializeOpCreateGlobalNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1256,7 +1220,6 @@ func (m *awsRestjson1_serializeOpCreateGlobalNetwork) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateGlobalNetworkInput(v *CreateGlobalNetworkInput, encoder *httpbinding.Encoder) error {
@@ -1296,10 +1259,8 @@ func (m *awsRestjson1_serializeOpCreateLink) HandleSer
 func (m *awsRestjson1_serializeOpCreateLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1348,7 +1309,6 @@ func (m *awsRestjson1_serializeOpCreateLink) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLinkInput(v *CreateLinkInput, encoder *httpbinding.Encoder) error {
@@ -1419,10 +1379,8 @@ func (m *awsRestjson1_serializeOpCreateSite) HandleSer
 func (m *awsRestjson1_serializeOpCreateSite) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1471,7 +1429,6 @@ func (m *awsRestjson1_serializeOpCreateSite) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSiteInput(v *CreateSiteInput, encoder *httpbinding.Encoder) error {
@@ -1527,10 +1484,8 @@ func (m *awsRestjson1_serializeOpCreateSiteToSiteVpnAt
 func (m *awsRestjson1_serializeOpCreateSiteToSiteVpnAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1575,7 +1530,6 @@ func (m *awsRestjson1_serializeOpCreateSiteToSiteVpnAt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSiteToSiteVpnAttachmentInput(v *CreateSiteToSiteVpnAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -1625,10 +1579,8 @@ func (m *awsRestjson1_serializeOpCreateTransitGatewayP
 func (m *awsRestjson1_serializeOpCreateTransitGatewayPeering) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1673,7 +1625,6 @@ func (m *awsRestjson1_serializeOpCreateTransitGatewayP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTransitGatewayPeeringInput(v *CreateTransitGatewayPeeringInput, encoder *httpbinding.Encoder) error {
@@ -1723,10 +1674,8 @@ func (m *awsRestjson1_serializeOpCreateTransitGatewayR
 func (m *awsRestjson1_serializeOpCreateTransitGatewayRouteTableAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1771,7 +1720,6 @@ func (m *awsRestjson1_serializeOpCreateTransitGatewayR
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTransitGatewayRouteTableAttachmentInput(v *CreateTransitGatewayRouteTableAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -1821,10 +1769,8 @@ func (m *awsRestjson1_serializeOpCreateVpcAttachment) 
 func (m *awsRestjson1_serializeOpCreateVpcAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1869,7 +1815,6 @@ func (m *awsRestjson1_serializeOpCreateVpcAttachment) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVpcAttachmentInput(v *CreateVpcAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -1933,10 +1878,8 @@ func (m *awsRestjson1_serializeOpDeleteAttachment) Han
 func (m *awsRestjson1_serializeOpDeleteAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1974,7 +1917,6 @@ func (m *awsRestjson1_serializeOpDeleteAttachment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAttachmentInput(v *DeleteAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -2004,10 +1946,8 @@ func (m *awsRestjson1_serializeOpDeleteConnection) Han
 func (m *awsRestjson1_serializeOpDeleteConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2045,7 +1985,6 @@ func (m *awsRestjson1_serializeOpDeleteConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConnectionInput(v *DeleteConnectionInput, encoder *httpbinding.Encoder) error {
@@ -2084,10 +2023,8 @@ func (m *awsRestjson1_serializeOpDeleteConnectPeer) Ha
 func (m *awsRestjson1_serializeOpDeleteConnectPeer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2125,7 +2062,6 @@ func (m *awsRestjson1_serializeOpDeleteConnectPeer) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConnectPeerInput(v *DeleteConnectPeerInput, encoder *httpbinding.Encoder) error {
@@ -2155,10 +2091,8 @@ func (m *awsRestjson1_serializeOpDeleteCoreNetwork) Ha
 func (m *awsRestjson1_serializeOpDeleteCoreNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2196,7 +2130,6 @@ func (m *awsRestjson1_serializeOpDeleteCoreNetwork) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCoreNetworkInput(v *DeleteCoreNetworkInput, encoder *httpbinding.Encoder) error {
@@ -2226,10 +2159,8 @@ func (m *awsRestjson1_serializeOpDeleteCoreNetworkPoli
 func (m *awsRestjson1_serializeOpDeleteCoreNetworkPolicyVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2267,7 +2198,6 @@ func (m *awsRestjson1_serializeOpDeleteCoreNetworkPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteCoreNetworkPolicyVersionInput(v *DeleteCoreNetworkPolicyVersionInput, encoder *httpbinding.Encoder) error {
@@ -2306,10 +2236,8 @@ func (m *awsRestjson1_serializeOpDeleteDevice) HandleS
 func (m *awsRestjson1_serializeOpDeleteDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2347,7 +2275,6 @@ func (m *awsRestjson1_serializeOpDeleteDevice) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDeviceInput(v *DeleteDeviceInput, encoder *httpbinding.Encoder) error {
@@ -2386,10 +2313,8 @@ func (m *awsRestjson1_serializeOpDeleteGlobalNetwork) 
 func (m *awsRestjson1_serializeOpDeleteGlobalNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2427,7 +2352,6 @@ func (m *awsRestjson1_serializeOpDeleteGlobalNetwork) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteGlobalNetworkInput(v *DeleteGlobalNetworkInput, encoder *httpbinding.Encoder) error {
@@ -2457,10 +2381,8 @@ func (m *awsRestjson1_serializeOpDeleteLink) HandleSer
 func (m *awsRestjson1_serializeOpDeleteLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2498,7 +2420,6 @@ func (m *awsRestjson1_serializeOpDeleteLink) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLinkInput(v *DeleteLinkInput, encoder *httpbinding.Encoder) error {
@@ -2537,10 +2458,8 @@ func (m *awsRestjson1_serializeOpDeletePeering) Handle
 func (m *awsRestjson1_serializeOpDeletePeering) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2578,7 +2497,6 @@ func (m *awsRestjson1_serializeOpDeletePeering) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePeeringInput(v *DeletePeeringInput, encoder *httpbinding.Encoder) error {
@@ -2608,10 +2526,8 @@ func (m *awsRestjson1_serializeOpDeleteResourcePolicy)
 func (m *awsRestjson1_serializeOpDeleteResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2649,7 +2565,6 @@ func (m *awsRestjson1_serializeOpDeleteResourcePolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteResourcePolicyInput(v *DeleteResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2679,10 +2594,8 @@ func (m *awsRestjson1_serializeOpDeleteSite) HandleSer
 func (m *awsRestjson1_serializeOpDeleteSite) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2720,7 +2633,6 @@ func (m *awsRestjson1_serializeOpDeleteSite) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSiteInput(v *DeleteSiteInput, encoder *httpbinding.Encoder) error {
@@ -2759,10 +2671,8 @@ func (m *awsRestjson1_serializeOpDeregisterTransitGate
 func (m *awsRestjson1_serializeOpDeregisterTransitGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2800,7 +2710,6 @@ func (m *awsRestjson1_serializeOpDeregisterTransitGate
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeregisterTransitGatewayInput(v *DeregisterTransitGatewayInput, encoder *httpbinding.Encoder) error {
@@ -2839,10 +2748,8 @@ func (m *awsRestjson1_serializeOpDescribeGlobalNetwork
 func (m *awsRestjson1_serializeOpDescribeGlobalNetworks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2880,7 +2787,6 @@ func (m *awsRestjson1_serializeOpDescribeGlobalNetwork
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeGlobalNetworksInput(v *DescribeGlobalNetworksInput, encoder *httpbinding.Encoder) error {
@@ -2915,10 +2821,8 @@ func (m *awsRestjson1_serializeOpDisassociateConnectPe
 func (m *awsRestjson1_serializeOpDisassociateConnectPeer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2956,7 +2860,6 @@ func (m *awsRestjson1_serializeOpDisassociateConnectPe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateConnectPeerInput(v *DisassociateConnectPeerInput, encoder *httpbinding.Encoder) error {
@@ -2995,10 +2898,8 @@ func (m *awsRestjson1_serializeOpDisassociateCustomerG
 func (m *awsRestjson1_serializeOpDisassociateCustomerGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3036,7 +2937,6 @@ func (m *awsRestjson1_serializeOpDisassociateCustomerG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateCustomerGatewayInput(v *DisassociateCustomerGatewayInput, encoder *httpbinding.Encoder) error {
@@ -3075,10 +2975,8 @@ func (m *awsRestjson1_serializeOpDisassociateLink) Han
 func (m *awsRestjson1_serializeOpDisassociateLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3116,7 +3014,6 @@ func (m *awsRestjson1_serializeOpDisassociateLink) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateLinkInput(v *DisassociateLinkInput, encoder *httpbinding.Encoder) error {
@@ -3154,10 +3051,8 @@ func (m *awsRestjson1_serializeOpDisassociateTransitGa
 func (m *awsRestjson1_serializeOpDisassociateTransitGatewayConnectPeer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3195,7 +3090,6 @@ func (m *awsRestjson1_serializeOpDisassociateTransitGa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateTransitGatewayConnectPeerInput(v *DisassociateTransitGatewayConnectPeerInput, encoder *httpbinding.Encoder) error {
@@ -3234,10 +3128,8 @@ func (m *awsRestjson1_serializeOpExecuteCoreNetworkCha
 func (m *awsRestjson1_serializeOpExecuteCoreNetworkChangeSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3275,7 +3167,6 @@ func (m *awsRestjson1_serializeOpExecuteCoreNetworkCha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExecuteCoreNetworkChangeSetInput(v *ExecuteCoreNetworkChangeSetInput, encoder *httpbinding.Encoder) error {
@@ -3314,10 +3205,8 @@ func (m *awsRestjson1_serializeOpGetConnectAttachment)
 func (m *awsRestjson1_serializeOpGetConnectAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3355,7 +3244,6 @@ func (m *awsRestjson1_serializeOpGetConnectAttachment)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectAttachmentInput(v *GetConnectAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -3385,10 +3273,8 @@ func (m *awsRestjson1_serializeOpGetConnections) Handl
 func (m *awsRestjson1_serializeOpGetConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3426,7 +3312,6 @@ func (m *awsRestjson1_serializeOpGetConnections) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectionsInput(v *GetConnectionsInput, encoder *httpbinding.Encoder) error {
@@ -3474,10 +3359,8 @@ func (m *awsRestjson1_serializeOpGetConnectPeer) Handl
 func (m *awsRestjson1_serializeOpGetConnectPeer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3515,7 +3398,6 @@ func (m *awsRestjson1_serializeOpGetConnectPeer) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectPeerInput(v *GetConnectPeerInput, encoder *httpbinding.Encoder) error {
@@ -3545,10 +3427,8 @@ func (m *awsRestjson1_serializeOpGetConnectPeerAssocia
 func (m *awsRestjson1_serializeOpGetConnectPeerAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3586,7 +3466,6 @@ func (m *awsRestjson1_serializeOpGetConnectPeerAssocia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConnectPeerAssociationsInput(v *GetConnectPeerAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -3630,10 +3509,8 @@ func (m *awsRestjson1_serializeOpGetCoreNetwork) Handl
 func (m *awsRestjson1_serializeOpGetCoreNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3671,7 +3548,6 @@ func (m *awsRestjson1_serializeOpGetCoreNetwork) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCoreNetworkInput(v *GetCoreNetworkInput, encoder *httpbinding.Encoder) error {
@@ -3701,10 +3577,8 @@ func (m *awsRestjson1_serializeOpGetCoreNetworkChangeE
 func (m *awsRestjson1_serializeOpGetCoreNetworkChangeEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3742,7 +3616,6 @@ func (m *awsRestjson1_serializeOpGetCoreNetworkChangeE
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCoreNetworkChangeEventsInput(v *GetCoreNetworkChangeEventsInput, encoder *httpbinding.Encoder) error {
@@ -3789,10 +3662,8 @@ func (m *awsRestjson1_serializeOpGetCoreNetworkChangeS
 func (m *awsRestjson1_serializeOpGetCoreNetworkChangeSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3830,7 +3701,6 @@ func (m *awsRestjson1_serializeOpGetCoreNetworkChangeS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCoreNetworkChangeSetInput(v *GetCoreNetworkChangeSetInput, encoder *httpbinding.Encoder) error {
@@ -3877,10 +3747,8 @@ func (m *awsRestjson1_serializeOpGetCoreNetworkPolicy)
 func (m *awsRestjson1_serializeOpGetCoreNetworkPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3918,7 +3786,6 @@ func (m *awsRestjson1_serializeOpGetCoreNetworkPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCoreNetworkPolicyInput(v *GetCoreNetworkPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3956,10 +3823,8 @@ func (m *awsRestjson1_serializeOpGetCustomerGatewayAss
 func (m *awsRestjson1_serializeOpGetCustomerGatewayAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3997,7 +3862,6 @@ func (m *awsRestjson1_serializeOpGetCustomerGatewayAss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCustomerGatewayAssociationsInput(v *GetCustomerGatewayAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4041,10 +3905,8 @@ func (m *awsRestjson1_serializeOpGetDevices) HandleSer
 func (m *awsRestjson1_serializeOpGetDevices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4082,7 +3944,6 @@ func (m *awsRestjson1_serializeOpGetDevices) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDevicesInput(v *GetDevicesInput, encoder *httpbinding.Encoder) error {
@@ -4130,10 +3991,8 @@ func (m *awsRestjson1_serializeOpGetDirectConnectGatew
 func (m *awsRestjson1_serializeOpGetDirectConnectGatewayAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4171,7 +4030,6 @@ func (m *awsRestjson1_serializeOpGetDirectConnectGatew
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDirectConnectGatewayAttachmentInput(v *GetDirectConnectGatewayAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -4201,10 +4059,8 @@ func (m *awsRestjson1_serializeOpGetLinkAssociations) 
 func (m *awsRestjson1_serializeOpGetLinkAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4242,7 +4098,6 @@ func (m *awsRestjson1_serializeOpGetLinkAssociations) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLinkAssociationsInput(v *GetLinkAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -4288,10 +4143,8 @@ func (m *awsRestjson1_serializeOpGetLinks) HandleSeria
 func (m *awsRestjson1_serializeOpGetLinks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4329,7 +4182,6 @@ func (m *awsRestjson1_serializeOpGetLinks) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLinksInput(v *GetLinksInput, encoder *httpbinding.Encoder) error {
@@ -4385,10 +4237,8 @@ func (m *awsRestjson1_serializeOpGetNetworkResourceCou
 func (m *awsRestjson1_serializeOpGetNetworkResourceCounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4426,7 +4276,6 @@ func (m *awsRestjson1_serializeOpGetNetworkResourceCou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNetworkResourceCountsInput(v *GetNetworkResourceCountsInput, encoder *httpbinding.Encoder) error {
@@ -4468,10 +4317,8 @@ func (m *awsRestjson1_serializeOpGetNetworkResourceRel
 func (m *awsRestjson1_serializeOpGetNetworkResourceRelationships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4509,7 +4356,6 @@ func (m *awsRestjson1_serializeOpGetNetworkResourceRel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNetworkResourceRelationshipsInput(v *GetNetworkResourceRelationshipsInput, encoder *httpbinding.Encoder) error {
@@ -4571,10 +4417,8 @@ func (m *awsRestjson1_serializeOpGetNetworkResources) 
 func (m *awsRestjson1_serializeOpGetNetworkResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4612,7 +4456,6 @@ func (m *awsRestjson1_serializeOpGetNetworkResources) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNetworkResourcesInput(v *GetNetworkResourcesInput, encoder *httpbinding.Encoder) error {
@@ -4674,10 +4517,8 @@ func (m *awsRestjson1_serializeOpGetNetworkRoutes) Han
 func (m *awsRestjson1_serializeOpGetNetworkRoutes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4726,7 +4567,6 @@ func (m *awsRestjson1_serializeOpGetNetworkRoutes) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNetworkRoutesInput(v *GetNetworkRoutesInput, encoder *httpbinding.Encoder) error {
@@ -4826,10 +4666,8 @@ func (m *awsRestjson1_serializeOpGetNetworkTelemetry) 
 func (m *awsRestjson1_serializeOpGetNetworkTelemetry) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4867,7 +4705,6 @@ func (m *awsRestjson1_serializeOpGetNetworkTelemetry) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNetworkTelemetryInput(v *GetNetworkTelemetryInput, encoder *httpbinding.Encoder) error {
@@ -4929,10 +4766,8 @@ func (m *awsRestjson1_serializeOpGetResourcePolicy) Ha
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
@@ -4970,7 +4805,6 @@ func (m *awsRestjson1_serializeOpGetResourcePolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetResourcePolicyInput(v *GetResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -5000,10 +4834,8 @@ func (m *awsRestjson1_serializeOpGetRouteAnalysis) Han
 func (m *awsRestjson1_serializeOpGetRouteAnalysis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5041,7 +4873,6 @@ func (m *awsRestjson1_serializeOpGetRouteAnalysis) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRouteAnalysisInput(v *GetRouteAnalysisInput, encoder *httpbinding.Encoder) error {
@@ -5080,10 +4911,8 @@ func (m *awsRestjson1_serializeOpGetSites) HandleSeria
 func (m *awsRestjson1_serializeOpGetSites) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5121,7 +4950,6 @@ func (m *awsRestjson1_serializeOpGetSites) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSitesInput(v *GetSitesInput, encoder *httpbinding.Encoder) error {
@@ -5165,10 +4993,8 @@ func (m *awsRestjson1_serializeOpGetSiteToSiteVpnAttac
 func (m *awsRestjson1_serializeOpGetSiteToSiteVpnAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5206,7 +5032,6 @@ func (m *awsRestjson1_serializeOpGetSiteToSiteVpnAttac
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSiteToSiteVpnAttachmentInput(v *GetSiteToSiteVpnAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -5236,10 +5061,8 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayConn
 func (m *awsRestjson1_serializeOpGetTransitGatewayConnectPeerAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5277,7 +5100,6 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayConn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTransitGatewayConnectPeerAssociationsInput(v *GetTransitGatewayConnectPeerAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -5321,10 +5143,8 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayPeer
 func (m *awsRestjson1_serializeOpGetTransitGatewayPeering) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5362,7 +5182,6 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayPeer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTransitGatewayPeeringInput(v *GetTransitGatewayPeeringInput, encoder *httpbinding.Encoder) error {
@@ -5392,10 +5211,8 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayRegi
 func (m *awsRestjson1_serializeOpGetTransitGatewayRegistrations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5433,7 +5250,6 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayRegi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTransitGatewayRegistrationsInput(v *GetTransitGatewayRegistrationsInput, encoder *httpbinding.Encoder) error {
@@ -5477,10 +5293,8 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayRout
 func (m *awsRestjson1_serializeOpGetTransitGatewayRouteTableAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5518,7 +5332,6 @@ func (m *awsRestjson1_serializeOpGetTransitGatewayRout
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTransitGatewayRouteTableAttachmentInput(v *GetTransitGatewayRouteTableAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -5548,10 +5361,8 @@ func (m *awsRestjson1_serializeOpGetVpcAttachment) Han
 func (m *awsRestjson1_serializeOpGetVpcAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5589,7 +5400,6 @@ func (m *awsRestjson1_serializeOpGetVpcAttachment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVpcAttachmentInput(v *GetVpcAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -5619,10 +5429,8 @@ func (m *awsRestjson1_serializeOpListAttachments) Hand
 func (m *awsRestjson1_serializeOpListAttachments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5660,7 +5468,6 @@ func (m *awsRestjson1_serializeOpListAttachments) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAttachmentsInput(v *ListAttachmentsInput, encoder *httpbinding.Encoder) error {
@@ -5705,10 +5512,8 @@ func (m *awsRestjson1_serializeOpListConnectPeers) Han
 func (m *awsRestjson1_serializeOpListConnectPeers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5746,7 +5551,6 @@ func (m *awsRestjson1_serializeOpListConnectPeers) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConnectPeersInput(v *ListConnectPeersInput, encoder *httpbinding.Encoder) error {
@@ -5783,10 +5587,8 @@ func (m *awsRestjson1_serializeOpListCoreNetworkPolicy
 func (m *awsRestjson1_serializeOpListCoreNetworkPolicyVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5824,7 +5626,6 @@ func (m *awsRestjson1_serializeOpListCoreNetworkPolicy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCoreNetworkPolicyVersionsInput(v *ListCoreNetworkPolicyVersionsInput, encoder *httpbinding.Encoder) error {
@@ -5862,10 +5663,8 @@ func (m *awsRestjson1_serializeOpListCoreNetworks) Han
 func (m *awsRestjson1_serializeOpListCoreNetworks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5903,7 +5702,6 @@ func (m *awsRestjson1_serializeOpListCoreNetworks) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCoreNetworksInput(v *ListCoreNetworksInput, encoder *httpbinding.Encoder) error {
@@ -5932,10 +5730,8 @@ func (m *awsRestjson1_serializeOpListOrganizationServi
 func (m *awsRestjson1_serializeOpListOrganizationServiceAccessStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5973,7 +5769,6 @@ func (m *awsRestjson1_serializeOpListOrganizationServi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationServiceAccessStatusInput(v *ListOrganizationServiceAccessStatusInput, encoder *httpbinding.Encoder) error {
@@ -6002,10 +5797,8 @@ func (m *awsRestjson1_serializeOpListPeerings) HandleS
 func (m *awsRestjson1_serializeOpListPeerings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6043,7 +5836,6 @@ func (m *awsRestjson1_serializeOpListPeerings) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPeeringsInput(v *ListPeeringsInput, encoder *httpbinding.Encoder) error {
@@ -6088,10 +5880,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6129,7 +5919,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6159,10 +5948,8 @@ func (m *awsRestjson1_serializeOpPutCoreNetworkPolicy)
 func (m *awsRestjson1_serializeOpPutCoreNetworkPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6211,7 +5998,6 @@ func (m *awsRestjson1_serializeOpPutCoreNetworkPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutCoreNetworkPolicyInput(v *PutCoreNetworkPolicyInput, encoder *httpbinding.Encoder) error {
@@ -6268,10 +6054,8 @@ func (m *awsRestjson1_serializeOpPutResourcePolicy) Ha
 func (m *awsRestjson1_serializeOpPutResourcePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6320,7 +6104,6 @@ func (m *awsRestjson1_serializeOpPutResourcePolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutResourcePolicyInput(v *PutResourcePolicyInput, encoder *httpbinding.Encoder) error {
@@ -6362,10 +6145,8 @@ func (m *awsRestjson1_serializeOpRegisterTransitGatewa
 func (m *awsRestjson1_serializeOpRegisterTransitGateway) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6414,7 +6195,6 @@ func (m *awsRestjson1_serializeOpRegisterTransitGatewa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterTransitGatewayInput(v *RegisterTransitGatewayInput, encoder *httpbinding.Encoder) error {
@@ -6456,10 +6236,8 @@ func (m *awsRestjson1_serializeOpRejectAttachment) Han
 func (m *awsRestjson1_serializeOpRejectAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6497,7 +6275,6 @@ func (m *awsRestjson1_serializeOpRejectAttachment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectAttachmentInput(v *RejectAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -6527,10 +6304,8 @@ func (m *awsRestjson1_serializeOpRestoreCoreNetworkPol
 func (m *awsRestjson1_serializeOpRestoreCoreNetworkPolicyVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6568,7 +6343,6 @@ func (m *awsRestjson1_serializeOpRestoreCoreNetworkPol
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRestoreCoreNetworkPolicyVersionInput(v *RestoreCoreNetworkPolicyVersionInput, encoder *httpbinding.Encoder) error {
@@ -6607,10 +6381,8 @@ func (m *awsRestjson1_serializeOpStartOrganizationServ
 func (m *awsRestjson1_serializeOpStartOrganizationServiceAccessUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6655,7 +6427,6 @@ func (m *awsRestjson1_serializeOpStartOrganizationServ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartOrganizationServiceAccessUpdateInput(v *StartOrganizationServiceAccessUpdateInput, encoder *httpbinding.Encoder) error {
@@ -6688,10 +6459,8 @@ func (m *awsRestjson1_serializeOpStartRouteAnalysis) H
 func (m *awsRestjson1_serializeOpStartRouteAnalysis) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6740,7 +6509,6 @@ func (m *awsRestjson1_serializeOpStartRouteAnalysis) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartRouteAnalysisInput(v *StartRouteAnalysisInput, encoder *httpbinding.Encoder) error {
@@ -6801,10 +6569,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -6853,7 +6619,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6897,10 +6662,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -6938,7 +6701,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -6974,10 +6736,8 @@ func (m *awsRestjson1_serializeOpUpdateConnection) Han
 func (m *awsRestjson1_serializeOpUpdateConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7026,7 +6786,6 @@ func (m *awsRestjson1_serializeOpUpdateConnection) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConnectionInput(v *UpdateConnectionInput, encoder *httpbinding.Encoder) error {
@@ -7087,10 +6846,8 @@ func (m *awsRestjson1_serializeOpUpdateCoreNetwork) Ha
 func (m *awsRestjson1_serializeOpUpdateCoreNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7139,7 +6896,6 @@ func (m *awsRestjson1_serializeOpUpdateCoreNetwork) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateCoreNetworkInput(v *UpdateCoreNetworkInput, encoder *httpbinding.Encoder) error {
@@ -7181,10 +6937,8 @@ func (m *awsRestjson1_serializeOpUpdateDevice) HandleS
 func (m *awsRestjson1_serializeOpUpdateDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7233,7 +6987,6 @@ func (m *awsRestjson1_serializeOpUpdateDevice) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDeviceInput(v *UpdateDeviceInput, encoder *httpbinding.Encoder) error {
@@ -7323,10 +7076,8 @@ func (m *awsRestjson1_serializeOpUpdateDirectConnectGa
 func (m *awsRestjson1_serializeOpUpdateDirectConnectGatewayAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7375,7 +7126,6 @@ func (m *awsRestjson1_serializeOpUpdateDirectConnectGa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDirectConnectGatewayAttachmentInput(v *UpdateDirectConnectGatewayAttachmentInput, encoder *httpbinding.Encoder) error {
@@ -7419,10 +7169,8 @@ func (m *awsRestjson1_serializeOpUpdateGlobalNetwork) 
 func (m *awsRestjson1_serializeOpUpdateGlobalNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7471,7 +7219,6 @@ func (m *awsRestjson1_serializeOpUpdateGlobalNetwork) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGlobalNetworkInput(v *UpdateGlobalNetworkInput, encoder *httpbinding.Encoder) error {
@@ -7513,10 +7260,8 @@ func (m *awsRestjson1_serializeOpUpdateLink) HandleSer
 func (m *awsRestjson1_serializeOpUpdateLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7565,7 +7310,6 @@ func (m *awsRestjson1_serializeOpUpdateLink) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLinkInput(v *UpdateLinkInput, encoder *httpbinding.Encoder) error {
@@ -7633,10 +7377,8 @@ func (m *awsRestjson1_serializeOpUpdateNetworkResource
 func (m *awsRestjson1_serializeOpUpdateNetworkResourceMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7685,7 +7427,6 @@ func (m *awsRestjson1_serializeOpUpdateNetworkResource
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateNetworkResourceMetadataInput(v *UpdateNetworkResourceMetadataInput, encoder *httpbinding.Encoder) error {
@@ -7738,10 +7479,8 @@ func (m *awsRestjson1_serializeOpUpdateSite) HandleSer
 func (m *awsRestjson1_serializeOpUpdateSite) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7790,7 +7529,6 @@ func (m *awsRestjson1_serializeOpUpdateSite) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSiteInput(v *UpdateSiteInput, encoder *httpbinding.Encoder) error {
@@ -7848,10 +7586,8 @@ func (m *awsRestjson1_serializeOpUpdateVpcAttachment) 
 func (m *awsRestjson1_serializeOpUpdateVpcAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7900,7 +7636,6 @@ func (m *awsRestjson1_serializeOpUpdateVpcAttachment) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVpcAttachmentInput(v *UpdateVpcAttachmentInput, encoder *httpbinding.Encoder) error {
