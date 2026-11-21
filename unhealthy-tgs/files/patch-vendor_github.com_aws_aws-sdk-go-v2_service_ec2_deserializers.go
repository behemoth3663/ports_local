--- vendor/github.com/aws/aws-sdk-go-v2/service/ec2/deserializers.go.orig	2026-11-21 14:46:10 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ec2/deserializers.go
@@ -17,7 +17,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -51,10 +50,8 @@ func (m *awsEc2query_deserializeOpAcceptAddressTransfe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptAddressTransfer(response, &metadata)
@@ -154,10 +151,8 @@ func (m *awsEc2query_deserializeOpAcceptCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptCapacityReservationBillingOwnership(response, &metadata)
@@ -257,10 +252,8 @@ func (m *awsEc2query_deserializeOpAcceptReservedInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptReservedInstancesExchangeQuote(response, &metadata)
@@ -360,10 +353,8 @@ func (m *awsEc2query_deserializeOpAcceptTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptTransitGatewayClientVpnAttachment(response, &metadata)
@@ -463,10 +454,8 @@ func (m *awsEc2query_deserializeOpAcceptTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptTransitGatewayMulticastDomainAssociations(response, &metadata)
@@ -566,10 +555,8 @@ func (m *awsEc2query_deserializeOpAcceptTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptTransitGatewayPeeringAttachment(response, &metadata)
@@ -669,10 +656,8 @@ func (m *awsEc2query_deserializeOpAcceptTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptTransitGatewayVpcAttachment(response, &metadata)
@@ -772,10 +757,8 @@ func (m *awsEc2query_deserializeOpAcceptVpcEndpointCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptVpcEndpointConnections(response, &metadata)
@@ -875,10 +858,8 @@ func (m *awsEc2query_deserializeOpAcceptVpcPeeringConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAcceptVpcPeeringConnection(response, &metadata)
@@ -978,10 +959,8 @@ func (m *awsEc2query_deserializeOpAdvertiseByoipCidr) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAdvertiseByoipCidr(response, &metadata)
@@ -1081,10 +1060,8 @@ func (m *awsEc2query_deserializeOpAllocateAddress) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAllocateAddress(response, &metadata)
@@ -1184,10 +1161,8 @@ func (m *awsEc2query_deserializeOpAllocateHosts) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAllocateHosts(response, &metadata)
@@ -1287,10 +1262,8 @@ func (m *awsEc2query_deserializeOpAllocateIpamPoolCidr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAllocateIpamPoolCidr(response, &metadata)
@@ -1390,10 +1363,8 @@ func (m *awsEc2query_deserializeOpApplySecurityGroupsT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorApplySecurityGroupsToClientVpnTargetNetwork(response, &metadata)
@@ -1493,10 +1464,8 @@ func (m *awsEc2query_deserializeOpAssignIpv6Addresses)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssignIpv6Addresses(response, &metadata)
@@ -1596,10 +1565,8 @@ func (m *awsEc2query_deserializeOpAssignPrivateIpAddre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssignPrivateIpAddresses(response, &metadata)
@@ -1699,10 +1666,8 @@ func (m *awsEc2query_deserializeOpAssignPrivateNatGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssignPrivateNatGatewayAddress(response, &metadata)
@@ -1802,10 +1767,8 @@ func (m *awsEc2query_deserializeOpAssociateAddress) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateAddress(response, &metadata)
@@ -1905,10 +1868,8 @@ func (m *awsEc2query_deserializeOpAssociateApplication
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateApplicationStatusCheck(response, &metadata)
@@ -2008,10 +1969,8 @@ func (m *awsEc2query_deserializeOpAssociateCapacityRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateCapacityReservationBillingOwner(response, &metadata)
@@ -2111,10 +2070,8 @@ func (m *awsEc2query_deserializeOpAssociateClientVpnTa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateClientVpnTargetNetwork(response, &metadata)
@@ -2214,10 +2171,8 @@ func (m *awsEc2query_deserializeOpAssociateDhcpOptions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateDhcpOptions(response, &metadata)
@@ -2294,10 +2249,8 @@ func (m *awsEc2query_deserializeOpAssociateEnclaveCert
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateEnclaveCertificateIamRole(response, &metadata)
@@ -2397,10 +2350,8 @@ func (m *awsEc2query_deserializeOpAssociateIamInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateIamInstanceProfile(response, &metadata)
@@ -2500,10 +2451,8 @@ func (m *awsEc2query_deserializeOpAssociateInstanceEve
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateInstanceEventWindow(response, &metadata)
@@ -2603,10 +2552,8 @@ func (m *awsEc2query_deserializeOpAssociateIpamByoasn)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateIpamByoasn(response, &metadata)
@@ -2706,10 +2653,8 @@ func (m *awsEc2query_deserializeOpAssociateIpamResourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateIpamResourceDiscovery(response, &metadata)
@@ -2809,10 +2754,8 @@ func (m *awsEc2query_deserializeOpAssociateNatGatewayA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateNatGatewayAddress(response, &metadata)
@@ -2912,10 +2855,8 @@ func (m *awsEc2query_deserializeOpAssociateRouteServer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateRouteServer(response, &metadata)
@@ -3015,10 +2956,8 @@ func (m *awsEc2query_deserializeOpAssociateRouteTable)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateRouteTable(response, &metadata)
@@ -3118,10 +3057,8 @@ func (m *awsEc2query_deserializeOpAssociateSecurityGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateSecurityGroupVpc(response, &metadata)
@@ -3221,10 +3158,8 @@ func (m *awsEc2query_deserializeOpAssociateSubnetCidrB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateSubnetCidrBlock(response, &metadata)
@@ -3324,10 +3259,8 @@ func (m *awsEc2query_deserializeOpAssociateTransitGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateTransitGatewayMulticastDomain(response, &metadata)
@@ -3427,10 +3360,8 @@ func (m *awsEc2query_deserializeOpAssociateTransitGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateTransitGatewayPolicyTable(response, &metadata)
@@ -3530,10 +3461,8 @@ func (m *awsEc2query_deserializeOpAssociateTransitGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateTransitGatewayRouteTable(response, &metadata)
@@ -3633,10 +3562,8 @@ func (m *awsEc2query_deserializeOpAssociateTrunkInterf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateTrunkInterface(response, &metadata)
@@ -3736,10 +3663,8 @@ func (m *awsEc2query_deserializeOpAssociateVpcCidrBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAssociateVpcCidrBlock(response, &metadata)
@@ -3839,10 +3764,8 @@ func (m *awsEc2query_deserializeOpAttachClassicLinkVpc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAttachClassicLinkVpc(response, &metadata)
@@ -3942,10 +3865,8 @@ func (m *awsEc2query_deserializeOpAttachImageWatermark
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAttachImageWatermark(response, &metadata)
@@ -4045,10 +3966,8 @@ func (m *awsEc2query_deserializeOpAttachInternetGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAttachInternetGateway(response, &metadata)
@@ -4125,10 +4044,8 @@ func (m *awsEc2query_deserializeOpAttachNetworkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAttachNetworkInterface(response, &metadata)
@@ -4228,10 +4145,8 @@ func (m *awsEc2query_deserializeOpAttachVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAttachVerifiedAccessTrustProvider(response, &metadata)
@@ -4331,10 +4246,8 @@ func (m *awsEc2query_deserializeOpAttachVolume) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAttachVolume(response, &metadata)
@@ -4434,10 +4347,8 @@ func (m *awsEc2query_deserializeOpAttachVpnGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAttachVpnGateway(response, &metadata)
@@ -4537,10 +4448,8 @@ func (m *awsEc2query_deserializeOpAuthorizeClientVpnIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAuthorizeClientVpnIngress(response, &metadata)
@@ -4640,10 +4549,8 @@ func (m *awsEc2query_deserializeOpAuthorizeSecurityGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAuthorizeSecurityGroupEgress(response, &metadata)
@@ -4743,10 +4650,8 @@ func (m *awsEc2query_deserializeOpAuthorizeSecurityGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorAuthorizeSecurityGroupIngress(response, &metadata)
@@ -4846,10 +4751,8 @@ func (m *awsEc2query_deserializeOpBatchModifyIpamRouti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorBatchModifyIpamRoutingPolicyRegistrations(response, &metadata)
@@ -4949,10 +4852,8 @@ func (m *awsEc2query_deserializeOpBundleInstance) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorBundleInstance(response, &metadata)
@@ -5052,10 +4953,8 @@ func (m *awsEc2query_deserializeOpCancelBundleTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelBundleTask(response, &metadata)
@@ -5155,10 +5054,8 @@ func (m *awsEc2query_deserializeOpCancelCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelCapacityReservation(response, &metadata)
@@ -5258,10 +5155,8 @@ func (m *awsEc2query_deserializeOpCancelCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelCapacityReservationFleets(response, &metadata)
@@ -5361,10 +5256,8 @@ func (m *awsEc2query_deserializeOpCancelConversionTask
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelConversionTask(response, &metadata)
@@ -5441,10 +5334,8 @@ func (m *awsEc2query_deserializeOpCancelDeclarativePol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelDeclarativePoliciesReport(response, &metadata)
@@ -5544,10 +5435,8 @@ func (m *awsEc2query_deserializeOpCancelExportTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelExportTask(response, &metadata)
@@ -5624,10 +5513,8 @@ func (m *awsEc2query_deserializeOpCancelImageLaunchPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelImageLaunchPermission(response, &metadata)
@@ -5727,10 +5614,8 @@ func (m *awsEc2query_deserializeOpCancelImportTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelImportTask(response, &metadata)
@@ -5830,10 +5715,8 @@ func (m *awsEc2query_deserializeOpCancelReservedInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelReservedInstancesListing(response, &metadata)
@@ -5933,10 +5816,8 @@ func (m *awsEc2query_deserializeOpCancelSpotFleetReque
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelSpotFleetRequests(response, &metadata)
@@ -6036,10 +5917,8 @@ func (m *awsEc2query_deserializeOpCancelSpotInstanceRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCancelSpotInstanceRequests(response, &metadata)
@@ -6139,10 +6018,8 @@ func (m *awsEc2query_deserializeOpConfirmProductInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorConfirmProductInstance(response, &metadata)
@@ -6242,10 +6119,8 @@ func (m *awsEc2query_deserializeOpCopyFpgaImage) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCopyFpgaImage(response, &metadata)
@@ -6345,10 +6220,8 @@ func (m *awsEc2query_deserializeOpCopyImage) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCopyImage(response, &metadata)
@@ -6448,10 +6321,8 @@ func (m *awsEc2query_deserializeOpCopySnapshot) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCopySnapshot(response, &metadata)
@@ -6551,10 +6422,8 @@ func (m *awsEc2query_deserializeOpCopyVolumes) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCopyVolumes(response, &metadata)
@@ -6654,10 +6523,8 @@ func (m *awsEc2query_deserializeOpCreateApplicationSta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateApplicationStatusCheck(response, &metadata)
@@ -6757,10 +6624,8 @@ func (m *awsEc2query_deserializeOpCreateCapacityManage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCapacityManagerDataExport(response, &metadata)
@@ -6860,10 +6725,8 @@ func (m *awsEc2query_deserializeOpCreateCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCapacityReservation(response, &metadata)
@@ -6963,10 +6826,8 @@ func (m *awsEc2query_deserializeOpCreateCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCapacityReservationBySplitting(response, &metadata)
@@ -7066,10 +6927,8 @@ func (m *awsEc2query_deserializeOpCreateCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCapacityReservationCancellationQuote(response, &metadata)
@@ -7169,10 +7028,8 @@ func (m *awsEc2query_deserializeOpCreateCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCapacityReservationDateChangeQuote(response, &metadata)
@@ -7272,10 +7129,8 @@ func (m *awsEc2query_deserializeOpCreateCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCapacityReservationFleet(response, &metadata)
@@ -7375,10 +7230,8 @@ func (m *awsEc2query_deserializeOpCreateCarrierGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCarrierGateway(response, &metadata)
@@ -7478,10 +7331,8 @@ func (m *awsEc2query_deserializeOpCreateClientVpnEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateClientVpnEndpoint(response, &metadata)
@@ -7581,10 +7432,8 @@ func (m *awsEc2query_deserializeOpCreateClientVpnRoute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateClientVpnRoute(response, &metadata)
@@ -7684,10 +7533,8 @@ func (m *awsEc2query_deserializeOpCreateCoipCidr) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCoipCidr(response, &metadata)
@@ -7787,10 +7634,8 @@ func (m *awsEc2query_deserializeOpCreateCoipPool) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCoipPool(response, &metadata)
@@ -7890,10 +7735,8 @@ func (m *awsEc2query_deserializeOpCreateCustomerGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateCustomerGateway(response, &metadata)
@@ -7993,10 +7836,8 @@ func (m *awsEc2query_deserializeOpCreateDefaultSubnet)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateDefaultSubnet(response, &metadata)
@@ -8096,10 +7937,8 @@ func (m *awsEc2query_deserializeOpCreateDefaultVpc) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateDefaultVpc(response, &metadata)
@@ -8199,10 +8038,8 @@ func (m *awsEc2query_deserializeOpCreateDelegateMacVol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateDelegateMacVolumeOwnershipTask(response, &metadata)
@@ -8302,10 +8139,8 @@ func (m *awsEc2query_deserializeOpCreateDhcpOptions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateDhcpOptions(response, &metadata)
@@ -8405,10 +8240,8 @@ func (m *awsEc2query_deserializeOpCreateEgressOnlyInte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateEgressOnlyInternetGateway(response, &metadata)
@@ -8508,10 +8341,8 @@ func (m *awsEc2query_deserializeOpCreateFleet) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateFleet(response, &metadata)
@@ -8611,10 +8442,8 @@ func (m *awsEc2query_deserializeOpCreateFlowLogs) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateFlowLogs(response, &metadata)
@@ -8714,10 +8543,8 @@ func (m *awsEc2query_deserializeOpCreateFpgaImage) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateFpgaImage(response, &metadata)
@@ -8817,10 +8644,8 @@ func (m *awsEc2query_deserializeOpCreateImage) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateImage(response, &metadata)
@@ -8920,10 +8745,8 @@ func (m *awsEc2query_deserializeOpCreateImageUsageRepo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateImageUsageReport(response, &metadata)
@@ -9023,10 +8846,8 @@ func (m *awsEc2query_deserializeOpCreateInstanceConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateInstanceConnectEndpoint(response, &metadata)
@@ -9126,10 +8947,8 @@ func (m *awsEc2query_deserializeOpCreateInstanceEventW
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateInstanceEventWindow(response, &metadata)
@@ -9229,10 +9048,8 @@ func (m *awsEc2query_deserializeOpCreateInstanceExport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateInstanceExportTask(response, &metadata)
@@ -9332,10 +9149,8 @@ func (m *awsEc2query_deserializeOpCreateInternetGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateInternetGateway(response, &metadata)
@@ -9435,10 +9250,8 @@ func (m *awsEc2query_deserializeOpCreateInterruptibleC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateInterruptibleCapacityReservationAllocation(response, &metadata)
@@ -9538,10 +9351,8 @@ func (m *awsEc2query_deserializeOpCreateIpam) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpam(response, &metadata)
@@ -9641,10 +9452,8 @@ func (m *awsEc2query_deserializeOpCreateIpamExternalRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamExternalResourceVerificationToken(response, &metadata)
@@ -9744,10 +9553,8 @@ func (m *awsEc2query_deserializeOpCreateIpamInternetRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamInternetRegistryAssociation(response, &metadata)
@@ -9847,10 +9654,8 @@ func (m *awsEc2query_deserializeOpCreateIpamPolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamPolicy(response, &metadata)
@@ -9950,10 +9755,8 @@ func (m *awsEc2query_deserializeOpCreateIpamPool) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamPool(response, &metadata)
@@ -10053,10 +9856,8 @@ func (m *awsEc2query_deserializeOpCreateIpamPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamPrefixListResolver(response, &metadata)
@@ -10156,10 +9957,8 @@ func (m *awsEc2query_deserializeOpCreateIpamPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamPrefixListResolverTarget(response, &metadata)
@@ -10259,10 +10058,8 @@ func (m *awsEc2query_deserializeOpCreateIpamResourceDi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamResourceDiscovery(response, &metadata)
@@ -10362,10 +10159,8 @@ func (m *awsEc2query_deserializeOpCreateIpamRoutingPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamRoutingPolicyRegistration(response, &metadata)
@@ -10465,10 +10260,8 @@ func (m *awsEc2query_deserializeOpCreateIpamScope) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateIpamScope(response, &metadata)
@@ -10568,10 +10361,8 @@ func (m *awsEc2query_deserializeOpCreateKeyPair) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateKeyPair(response, &metadata)
@@ -10671,10 +10462,8 @@ func (m *awsEc2query_deserializeOpCreateLaunchTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLaunchTemplate(response, &metadata)
@@ -10774,10 +10563,8 @@ func (m *awsEc2query_deserializeOpCreateLaunchTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLaunchTemplateVersion(response, &metadata)
@@ -10877,10 +10664,8 @@ func (m *awsEc2query_deserializeOpCreateLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLocalGatewayRoute(response, &metadata)
@@ -10980,10 +10765,8 @@ func (m *awsEc2query_deserializeOpCreateLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLocalGatewayRouteTable(response, &metadata)
@@ -11083,10 +10866,8 @@ func (m *awsEc2query_deserializeOpCreateLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(response, &metadata)
@@ -11186,10 +10967,8 @@ func (m *awsEc2query_deserializeOpCreateLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLocalGatewayRouteTableVpcAssociation(response, &metadata)
@@ -11289,10 +11068,8 @@ func (m *awsEc2query_deserializeOpCreateLocalGatewayVi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLocalGatewayVirtualInterface(response, &metadata)
@@ -11392,10 +11169,8 @@ func (m *awsEc2query_deserializeOpCreateLocalGatewayVi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateLocalGatewayVirtualInterfaceGroup(response, &metadata)
@@ -11495,10 +11270,8 @@ func (m *awsEc2query_deserializeOpCreateMacSystemInteg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateMacSystemIntegrityProtectionModificationTask(response, &metadata)
@@ -11598,10 +11371,8 @@ func (m *awsEc2query_deserializeOpCreateManagedPrefixL
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateManagedPrefixList(response, &metadata)
@@ -11701,10 +11472,8 @@ func (m *awsEc2query_deserializeOpCreateNatGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateNatGateway(response, &metadata)
@@ -11804,10 +11573,8 @@ func (m *awsEc2query_deserializeOpCreateNetworkAcl) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateNetworkAcl(response, &metadata)
@@ -11907,10 +11674,8 @@ func (m *awsEc2query_deserializeOpCreateNetworkAclEntr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateNetworkAclEntry(response, &metadata)
@@ -11987,10 +11752,8 @@ func (m *awsEc2query_deserializeOpCreateNetworkInsight
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateNetworkInsightsAccessScope(response, &metadata)
@@ -12090,10 +11853,8 @@ func (m *awsEc2query_deserializeOpCreateNetworkInsight
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateNetworkInsightsPath(response, &metadata)
@@ -12193,10 +11954,8 @@ func (m *awsEc2query_deserializeOpCreateNetworkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateNetworkInterface(response, &metadata)
@@ -12296,10 +12055,8 @@ func (m *awsEc2query_deserializeOpCreateNetworkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateNetworkInterfacePermission(response, &metadata)
@@ -12399,10 +12156,8 @@ func (m *awsEc2query_deserializeOpCreatePlacementGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreatePlacementGroup(response, &metadata)
@@ -12502,10 +12257,8 @@ func (m *awsEc2query_deserializeOpCreatePublicIpv4Pool
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreatePublicIpv4Pool(response, &metadata)
@@ -12605,10 +12358,8 @@ func (m *awsEc2query_deserializeOpCreateReplaceRootVol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateReplaceRootVolumeTask(response, &metadata)
@@ -12708,10 +12459,8 @@ func (m *awsEc2query_deserializeOpCreateReservedInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateReservedInstancesListing(response, &metadata)
@@ -12811,10 +12560,8 @@ func (m *awsEc2query_deserializeOpCreateRestoreImageTa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateRestoreImageTask(response, &metadata)
@@ -12914,10 +12661,8 @@ func (m *awsEc2query_deserializeOpCreateRoute) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateRoute(response, &metadata)
@@ -13017,10 +12762,8 @@ func (m *awsEc2query_deserializeOpCreateRouteServer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateRouteServer(response, &metadata)
@@ -13120,10 +12863,8 @@ func (m *awsEc2query_deserializeOpCreateRouteServerEnd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateRouteServerEndpoint(response, &metadata)
@@ -13223,10 +12964,8 @@ func (m *awsEc2query_deserializeOpCreateRouteServerPee
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateRouteServerPeer(response, &metadata)
@@ -13326,10 +13065,8 @@ func (m *awsEc2query_deserializeOpCreateRouteTable) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateRouteTable(response, &metadata)
@@ -13429,10 +13166,8 @@ func (m *awsEc2query_deserializeOpCreateSecondaryNetwo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSecondaryNetwork(response, &metadata)
@@ -13532,10 +13267,8 @@ func (m *awsEc2query_deserializeOpCreateSecondarySubne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSecondarySubnet(response, &metadata)
@@ -13635,10 +13368,8 @@ func (m *awsEc2query_deserializeOpCreateSecurityGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSecurityGroup(response, &metadata)
@@ -13738,10 +13469,8 @@ func (m *awsEc2query_deserializeOpCreateSnapshot) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSnapshot(response, &metadata)
@@ -13841,10 +13570,8 @@ func (m *awsEc2query_deserializeOpCreateSnapshots) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSnapshots(response, &metadata)
@@ -13944,10 +13671,8 @@ func (m *awsEc2query_deserializeOpCreateSpotDatafeedSu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSpotDatafeedSubscription(response, &metadata)
@@ -14047,10 +13772,8 @@ func (m *awsEc2query_deserializeOpCreateStoreImageTask
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateStoreImageTask(response, &metadata)
@@ -14150,10 +13873,8 @@ func (m *awsEc2query_deserializeOpCreateSubnet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSubnet(response, &metadata)
@@ -14253,10 +13974,8 @@ func (m *awsEc2query_deserializeOpCreateSubnetCidrRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateSubnetCidrReservation(response, &metadata)
@@ -14356,10 +14075,8 @@ func (m *awsEc2query_deserializeOpCreateTags) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTags(response, &metadata)
@@ -14436,10 +14153,8 @@ func (m *awsEc2query_deserializeOpCreateTrafficMirrorF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTrafficMirrorFilter(response, &metadata)
@@ -14539,10 +14254,8 @@ func (m *awsEc2query_deserializeOpCreateTrafficMirrorF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTrafficMirrorFilterRule(response, &metadata)
@@ -14642,10 +14355,8 @@ func (m *awsEc2query_deserializeOpCreateTrafficMirrorS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTrafficMirrorSession(response, &metadata)
@@ -14745,10 +14456,8 @@ func (m *awsEc2query_deserializeOpCreateTrafficMirrorT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTrafficMirrorTarget(response, &metadata)
@@ -14848,10 +14557,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGateway(response, &metadata)
@@ -14951,10 +14658,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayConnect(response, &metadata)
@@ -15054,10 +14759,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayConnectPeer(response, &metadata)
@@ -15157,10 +14860,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayMeteringPolicy(response, &metadata)
@@ -15260,10 +14961,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayMeteringPolicyEntry(response, &metadata)
@@ -15363,10 +15062,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayMulticastDomain(response, &metadata)
@@ -15466,10 +15163,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayPeeringAttachment(response, &metadata)
@@ -15569,10 +15264,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayPolicyTable(response, &metadata)
@@ -15672,10 +15365,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayPolicyTableEntry(response, &metadata)
@@ -15775,10 +15466,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayPrefixListReference(response, &metadata)
@@ -15878,10 +15567,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayRoute(response, &metadata)
@@ -15981,10 +15668,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayRouteTable(response, &metadata)
@@ -16084,10 +15769,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayRouteTableAnnouncement(response, &metadata)
@@ -16187,10 +15870,8 @@ func (m *awsEc2query_deserializeOpCreateTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateTransitGatewayVpcAttachment(response, &metadata)
@@ -16290,10 +15971,8 @@ func (m *awsEc2query_deserializeOpCreateVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVerifiedAccessEndpoint(response, &metadata)
@@ -16393,10 +16072,8 @@ func (m *awsEc2query_deserializeOpCreateVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVerifiedAccessGroup(response, &metadata)
@@ -16496,10 +16173,8 @@ func (m *awsEc2query_deserializeOpCreateVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVerifiedAccessInstance(response, &metadata)
@@ -16599,10 +16274,8 @@ func (m *awsEc2query_deserializeOpCreateVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVerifiedAccessTrustProvider(response, &metadata)
@@ -16702,10 +16375,8 @@ func (m *awsEc2query_deserializeOpCreateVolume) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVolume(response, &metadata)
@@ -16805,10 +16476,8 @@ func (m *awsEc2query_deserializeOpCreateVpc) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpc(response, &metadata)
@@ -16908,10 +16577,8 @@ func (m *awsEc2query_deserializeOpCreateVpcBlockPublic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpcBlockPublicAccessExclusion(response, &metadata)
@@ -17011,10 +16678,8 @@ func (m *awsEc2query_deserializeOpCreateVpcEncryptionC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpcEncryptionControl(response, &metadata)
@@ -17114,10 +16779,8 @@ func (m *awsEc2query_deserializeOpCreateVpcEndpoint) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpcEndpoint(response, &metadata)
@@ -17217,10 +16880,8 @@ func (m *awsEc2query_deserializeOpCreateVpcEndpointCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpcEndpointConnectionNotification(response, &metadata)
@@ -17320,10 +16981,8 @@ func (m *awsEc2query_deserializeOpCreateVpcEndpointSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpcEndpointServiceConfiguration(response, &metadata)
@@ -17423,10 +17082,8 @@ func (m *awsEc2query_deserializeOpCreateVpcPeeringConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpcPeeringConnection(response, &metadata)
@@ -17526,10 +17183,8 @@ func (m *awsEc2query_deserializeOpCreateVpnConcentrato
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpnConcentrator(response, &metadata)
@@ -17629,10 +17284,8 @@ func (m *awsEc2query_deserializeOpCreateVpnConnection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpnConnection(response, &metadata)
@@ -17732,10 +17385,8 @@ func (m *awsEc2query_deserializeOpCreateVpnConnectionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpnConnectionRoute(response, &metadata)
@@ -17812,10 +17463,8 @@ func (m *awsEc2query_deserializeOpCreateVpnGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorCreateVpnGateway(response, &metadata)
@@ -17915,10 +17564,8 @@ func (m *awsEc2query_deserializeOpDeleteApplicationSta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteApplicationStatusCheck(response, &metadata)
@@ -18018,10 +17665,8 @@ func (m *awsEc2query_deserializeOpDeleteCapacityManage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteCapacityManagerDataExport(response, &metadata)
@@ -18121,10 +17766,8 @@ func (m *awsEc2query_deserializeOpDeleteCarrierGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteCarrierGateway(response, &metadata)
@@ -18224,10 +17867,8 @@ func (m *awsEc2query_deserializeOpDeleteClientVpnEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteClientVpnEndpoint(response, &metadata)
@@ -18327,10 +17968,8 @@ func (m *awsEc2query_deserializeOpDeleteClientVpnEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteClientVpnEndpointAuthorizationPolicy(response, &metadata)
@@ -18430,10 +18069,8 @@ func (m *awsEc2query_deserializeOpDeleteClientVpnRoute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteClientVpnRoute(response, &metadata)
@@ -18533,10 +18170,8 @@ func (m *awsEc2query_deserializeOpDeleteCoipCidr) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteCoipCidr(response, &metadata)
@@ -18636,10 +18271,8 @@ func (m *awsEc2query_deserializeOpDeleteCoipPool) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteCoipPool(response, &metadata)
@@ -18739,10 +18372,8 @@ func (m *awsEc2query_deserializeOpDeleteCustomerGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteCustomerGateway(response, &metadata)
@@ -18819,10 +18450,8 @@ func (m *awsEc2query_deserializeOpDeleteDhcpOptions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteDhcpOptions(response, &metadata)
@@ -18899,10 +18528,8 @@ func (m *awsEc2query_deserializeOpDeleteEgressOnlyInte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteEgressOnlyInternetGateway(response, &metadata)
@@ -19002,10 +18629,8 @@ func (m *awsEc2query_deserializeOpDeleteFleets) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteFleets(response, &metadata)
@@ -19105,10 +18730,8 @@ func (m *awsEc2query_deserializeOpDeleteFlowLogs) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteFlowLogs(response, &metadata)
@@ -19208,10 +18831,8 @@ func (m *awsEc2query_deserializeOpDeleteFpgaImage) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteFpgaImage(response, &metadata)
@@ -19311,10 +18932,8 @@ func (m *awsEc2query_deserializeOpDeleteImageUsageRepo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteImageUsageReport(response, &metadata)
@@ -19414,10 +19033,8 @@ func (m *awsEc2query_deserializeOpDeleteInstanceConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteInstanceConnectEndpoint(response, &metadata)
@@ -19517,10 +19134,8 @@ func (m *awsEc2query_deserializeOpDeleteInstanceEventW
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteInstanceEventWindow(response, &metadata)
@@ -19620,10 +19235,8 @@ func (m *awsEc2query_deserializeOpDeleteInternetGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteInternetGateway(response, &metadata)
@@ -19700,10 +19313,8 @@ func (m *awsEc2query_deserializeOpDeleteIpam) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpam(response, &metadata)
@@ -19803,10 +19414,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamExternalRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamExternalResourceVerificationToken(response, &metadata)
@@ -19906,10 +19515,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamInternetRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamInternetRegistryAssociation(response, &metadata)
@@ -20009,10 +19616,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamPolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamPolicy(response, &metadata)
@@ -20112,10 +19717,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamPool) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamPool(response, &metadata)
@@ -20215,10 +19818,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamPrefixListResolver(response, &metadata)
@@ -20318,10 +19919,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamPrefixListResolverTarget(response, &metadata)
@@ -20421,10 +20020,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamResourceDi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamResourceDiscovery(response, &metadata)
@@ -20524,10 +20121,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamRoutingPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamRoutingPolicyRegistration(response, &metadata)
@@ -20627,10 +20222,8 @@ func (m *awsEc2query_deserializeOpDeleteIpamScope) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteIpamScope(response, &metadata)
@@ -20730,10 +20323,8 @@ func (m *awsEc2query_deserializeOpDeleteKeyPair) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteKeyPair(response, &metadata)
@@ -20833,10 +20424,8 @@ func (m *awsEc2query_deserializeOpDeleteLaunchTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLaunchTemplate(response, &metadata)
@@ -20936,10 +20525,8 @@ func (m *awsEc2query_deserializeOpDeleteLaunchTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLaunchTemplateVersions(response, &metadata)
@@ -21039,10 +20626,8 @@ func (m *awsEc2query_deserializeOpDeleteLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLocalGatewayRoute(response, &metadata)
@@ -21142,10 +20727,8 @@ func (m *awsEc2query_deserializeOpDeleteLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLocalGatewayRouteTable(response, &metadata)
@@ -21245,10 +20828,8 @@ func (m *awsEc2query_deserializeOpDeleteLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(response, &metadata)
@@ -21348,10 +20929,8 @@ func (m *awsEc2query_deserializeOpDeleteLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLocalGatewayRouteTableVpcAssociation(response, &metadata)
@@ -21451,10 +21030,8 @@ func (m *awsEc2query_deserializeOpDeleteLocalGatewayVi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLocalGatewayVirtualInterface(response, &metadata)
@@ -21554,10 +21131,8 @@ func (m *awsEc2query_deserializeOpDeleteLocalGatewayVi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteLocalGatewayVirtualInterfaceGroup(response, &metadata)
@@ -21657,10 +21232,8 @@ func (m *awsEc2query_deserializeOpDeleteManagedPrefixL
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteManagedPrefixList(response, &metadata)
@@ -21760,10 +21333,8 @@ func (m *awsEc2query_deserializeOpDeleteNatGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNatGateway(response, &metadata)
@@ -21863,10 +21434,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkAcl) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkAcl(response, &metadata)
@@ -21943,10 +21512,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkAclEntr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkAclEntry(response, &metadata)
@@ -22023,10 +21590,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkInsight
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkInsightsAccessScope(response, &metadata)
@@ -22126,10 +21691,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkInsight
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkInsightsAccessScopeAnalysis(response, &metadata)
@@ -22229,10 +21792,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkInsight
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkInsightsAnalysis(response, &metadata)
@@ -22332,10 +21893,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkInsight
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkInsightsPath(response, &metadata)
@@ -22435,10 +21994,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkInterface(response, &metadata)
@@ -22515,10 +22072,8 @@ func (m *awsEc2query_deserializeOpDeleteNetworkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteNetworkInterfacePermission(response, &metadata)
@@ -22618,10 +22173,8 @@ func (m *awsEc2query_deserializeOpDeletePlacementGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeletePlacementGroup(response, &metadata)
@@ -22698,10 +22251,8 @@ func (m *awsEc2query_deserializeOpDeletePublicIpv4Pool
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeletePublicIpv4Pool(response, &metadata)
@@ -22801,10 +22352,8 @@ func (m *awsEc2query_deserializeOpDeleteQueuedReserved
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteQueuedReservedInstances(response, &metadata)
@@ -22904,10 +22453,8 @@ func (m *awsEc2query_deserializeOpDeleteRoute) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteRoute(response, &metadata)
@@ -22984,10 +22531,8 @@ func (m *awsEc2query_deserializeOpDeleteRouteServer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteRouteServer(response, &metadata)
@@ -23087,10 +22632,8 @@ func (m *awsEc2query_deserializeOpDeleteRouteServerEnd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteRouteServerEndpoint(response, &metadata)
@@ -23190,10 +22733,8 @@ func (m *awsEc2query_deserializeOpDeleteRouteServerPee
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteRouteServerPeer(response, &metadata)
@@ -23293,10 +22834,8 @@ func (m *awsEc2query_deserializeOpDeleteRouteTable) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteRouteTable(response, &metadata)
@@ -23373,10 +22912,8 @@ func (m *awsEc2query_deserializeOpDeleteSecondaryNetwo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteSecondaryNetwork(response, &metadata)
@@ -23476,10 +23013,8 @@ func (m *awsEc2query_deserializeOpDeleteSecondarySubne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteSecondarySubnet(response, &metadata)
@@ -23579,10 +23114,8 @@ func (m *awsEc2query_deserializeOpDeleteSecurityGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteSecurityGroup(response, &metadata)
@@ -23682,10 +23215,8 @@ func (m *awsEc2query_deserializeOpDeleteSnapshot) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteSnapshot(response, &metadata)
@@ -23762,10 +23293,8 @@ func (m *awsEc2query_deserializeOpDeleteSpotDatafeedSu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteSpotDatafeedSubscription(response, &metadata)
@@ -23842,10 +23371,8 @@ func (m *awsEc2query_deserializeOpDeleteSubnet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteSubnet(response, &metadata)
@@ -23922,10 +23449,8 @@ func (m *awsEc2query_deserializeOpDeleteSubnetCidrRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteSubnetCidrReservation(response, &metadata)
@@ -24025,10 +23550,8 @@ func (m *awsEc2query_deserializeOpDeleteTags) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTags(response, &metadata)
@@ -24105,10 +23628,8 @@ func (m *awsEc2query_deserializeOpDeleteTrafficMirrorF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTrafficMirrorFilter(response, &metadata)
@@ -24208,10 +23729,8 @@ func (m *awsEc2query_deserializeOpDeleteTrafficMirrorF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTrafficMirrorFilterRule(response, &metadata)
@@ -24311,10 +23830,8 @@ func (m *awsEc2query_deserializeOpDeleteTrafficMirrorS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTrafficMirrorSession(response, &metadata)
@@ -24414,10 +23931,8 @@ func (m *awsEc2query_deserializeOpDeleteTrafficMirrorT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTrafficMirrorTarget(response, &metadata)
@@ -24517,10 +24032,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGateway(response, &metadata)
@@ -24620,10 +24133,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayClientVpnAttachment(response, &metadata)
@@ -24723,10 +24234,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayConnect(response, &metadata)
@@ -24826,10 +24335,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayConnectPeer(response, &metadata)
@@ -24929,10 +24436,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayMeteringPolicy(response, &metadata)
@@ -25032,10 +24537,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayMeteringPolicyEntry(response, &metadata)
@@ -25135,10 +24638,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayMulticastDomain(response, &metadata)
@@ -25238,10 +24739,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayPeeringAttachment(response, &metadata)
@@ -25341,10 +24840,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayPolicyTable(response, &metadata)
@@ -25444,10 +24941,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayPolicyTableEntry(response, &metadata)
@@ -25547,10 +25042,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayPrefixListReference(response, &metadata)
@@ -25650,10 +25143,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayRoute(response, &metadata)
@@ -25753,10 +25244,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayRouteTable(response, &metadata)
@@ -25856,10 +25345,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayRouteTableAnnouncement(response, &metadata)
@@ -25959,10 +25446,8 @@ func (m *awsEc2query_deserializeOpDeleteTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteTransitGatewayVpcAttachment(response, &metadata)
@@ -26062,10 +25547,8 @@ func (m *awsEc2query_deserializeOpDeleteVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVerifiedAccessEndpoint(response, &metadata)
@@ -26165,10 +25648,8 @@ func (m *awsEc2query_deserializeOpDeleteVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVerifiedAccessGroup(response, &metadata)
@@ -26268,10 +25749,8 @@ func (m *awsEc2query_deserializeOpDeleteVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVerifiedAccessInstance(response, &metadata)
@@ -26371,10 +25850,8 @@ func (m *awsEc2query_deserializeOpDeleteVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVerifiedAccessTrustProvider(response, &metadata)
@@ -26474,10 +25951,8 @@ func (m *awsEc2query_deserializeOpDeleteVolume) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVolume(response, &metadata)
@@ -26554,10 +26029,8 @@ func (m *awsEc2query_deserializeOpDeleteVpc) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpc(response, &metadata)
@@ -26634,10 +26107,8 @@ func (m *awsEc2query_deserializeOpDeleteVpcBlockPublic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpcBlockPublicAccessExclusion(response, &metadata)
@@ -26737,10 +26208,8 @@ func (m *awsEc2query_deserializeOpDeleteVpcEncryptionC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpcEncryptionControl(response, &metadata)
@@ -26840,10 +26309,8 @@ func (m *awsEc2query_deserializeOpDeleteVpcEndpointCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpcEndpointConnectionNotifications(response, &metadata)
@@ -26943,10 +26410,8 @@ func (m *awsEc2query_deserializeOpDeleteVpcEndpoints) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpcEndpoints(response, &metadata)
@@ -27046,10 +26511,8 @@ func (m *awsEc2query_deserializeOpDeleteVpcEndpointSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpcEndpointServiceConfigurations(response, &metadata)
@@ -27149,10 +26612,8 @@ func (m *awsEc2query_deserializeOpDeleteVpcPeeringConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpcPeeringConnection(response, &metadata)
@@ -27252,10 +26713,8 @@ func (m *awsEc2query_deserializeOpDeleteVpnConcentrato
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpnConcentrator(response, &metadata)
@@ -27355,10 +26814,8 @@ func (m *awsEc2query_deserializeOpDeleteVpnConnection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpnConnection(response, &metadata)
@@ -27435,10 +26892,8 @@ func (m *awsEc2query_deserializeOpDeleteVpnConnectionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpnConnectionRoute(response, &metadata)
@@ -27515,10 +26970,8 @@ func (m *awsEc2query_deserializeOpDeleteVpnGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeleteVpnGateway(response, &metadata)
@@ -27595,10 +27048,8 @@ func (m *awsEc2query_deserializeOpDeprovisionByoipCidr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeprovisionByoipCidr(response, &metadata)
@@ -27698,10 +27149,8 @@ func (m *awsEc2query_deserializeOpDeprovisionIpamByoas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeprovisionIpamByoasn(response, &metadata)
@@ -27801,10 +27250,8 @@ func (m *awsEc2query_deserializeOpDeprovisionIpamPoolC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeprovisionIpamPoolCidr(response, &metadata)
@@ -27904,10 +27351,8 @@ func (m *awsEc2query_deserializeOpDeprovisionPublicIpv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeprovisionPublicIpv4PoolCidr(response, &metadata)
@@ -28007,10 +27452,8 @@ func (m *awsEc2query_deserializeOpDeregisterImage) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeregisterImage(response, &metadata)
@@ -28110,10 +27553,8 @@ func (m *awsEc2query_deserializeOpDeregisterInstanceEv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeregisterInstanceEventNotificationAttributes(response, &metadata)
@@ -28213,10 +27654,8 @@ func (m *awsEc2query_deserializeOpDeregisterTransitGat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeregisterTransitGatewayMulticastGroupMembers(response, &metadata)
@@ -28316,10 +27755,8 @@ func (m *awsEc2query_deserializeOpDeregisterTransitGat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDeregisterTransitGatewayMulticastGroupSources(response, &metadata)
@@ -28419,10 +27856,8 @@ func (m *awsEc2query_deserializeOpDescribeAccountAttri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAccountAttributes(response, &metadata)
@@ -28522,10 +27957,8 @@ func (m *awsEc2query_deserializeOpDescribeAccountVpcEn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAccountVpcEncryptionControl(response, &metadata)
@@ -28625,10 +28058,8 @@ func (m *awsEc2query_deserializeOpDescribeAddresses) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAddresses(response, &metadata)
@@ -28728,10 +28159,8 @@ func (m *awsEc2query_deserializeOpDescribeAddressesAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAddressesAttribute(response, &metadata)
@@ -28831,10 +28260,8 @@ func (m *awsEc2query_deserializeOpDescribeAddressTrans
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAddressTransfers(response, &metadata)
@@ -28934,10 +28361,8 @@ func (m *awsEc2query_deserializeOpDescribeAggregateIdF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAggregateIdFormat(response, &metadata)
@@ -29037,10 +28462,8 @@ func (m *awsEc2query_deserializeOpDescribeApplicationS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeApplicationStatus(response, &metadata)
@@ -29140,10 +28563,8 @@ func (m *awsEc2query_deserializeOpDescribeApplicationS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeApplicationStatusCheckAssociations(response, &metadata)
@@ -29243,10 +28664,8 @@ func (m *awsEc2query_deserializeOpDescribeApplicationS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeApplicationStatusChecks(response, &metadata)
@@ -29346,10 +28765,8 @@ func (m *awsEc2query_deserializeOpDescribeAvailability
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAvailabilityZones(response, &metadata)
@@ -29449,10 +28866,8 @@ func (m *awsEc2query_deserializeOpDescribeAwsNetworkPe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeAwsNetworkPerformanceMetricSubscriptions(response, &metadata)
@@ -29552,10 +28967,8 @@ func (m *awsEc2query_deserializeOpDescribeBundleTasks)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeBundleTasks(response, &metadata)
@@ -29655,10 +29068,8 @@ func (m *awsEc2query_deserializeOpDescribeByoipCidrs) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeByoipCidrs(response, &metadata)
@@ -29758,10 +29169,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityBlockExtensionHistory(response, &metadata)
@@ -29861,10 +29270,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityBlockExtensionOfferings(response, &metadata)
@@ -29964,10 +29371,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityBlockOfferings(response, &metadata)
@@ -30067,10 +29472,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityBlocks(response, &metadata)
@@ -30170,10 +29573,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityBlockStatus(response, &metadata)
@@ -30273,10 +29674,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityMana
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityManagerDataExports(response, &metadata)
@@ -30376,10 +29775,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityReservationBillingRequests(response, &metadata)
@@ -30479,10 +29876,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityReservationCancellationQuotes(response, &metadata)
@@ -30582,10 +29977,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityReservationDateChangeQuotes(response, &metadata)
@@ -30685,10 +30078,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityReservationFleets(response, &metadata)
@@ -30788,10 +30179,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityReservations(response, &metadata)
@@ -30891,10 +30280,8 @@ func (m *awsEc2query_deserializeOpDescribeCapacityRese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCapacityReservationTopology(response, &metadata)
@@ -30994,10 +30381,8 @@ func (m *awsEc2query_deserializeOpDescribeCarrierGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCarrierGateways(response, &metadata)
@@ -31097,10 +30482,8 @@ func (m *awsEc2query_deserializeOpDescribeClassicLinkI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeClassicLinkInstances(response, &metadata)
@@ -31200,10 +30583,8 @@ func (m *awsEc2query_deserializeOpDescribeClientVpnAut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeClientVpnAuthorizationRules(response, &metadata)
@@ -31303,10 +30684,8 @@ func (m *awsEc2query_deserializeOpDescribeClientVpnCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeClientVpnConnections(response, &metadata)
@@ -31406,10 +30785,8 @@ func (m *awsEc2query_deserializeOpDescribeClientVpnEnd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeClientVpnEndpoints(response, &metadata)
@@ -31509,10 +30886,8 @@ func (m *awsEc2query_deserializeOpDescribeClientVpnRou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeClientVpnRoutes(response, &metadata)
@@ -31612,10 +30987,8 @@ func (m *awsEc2query_deserializeOpDescribeClientVpnTar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeClientVpnTargetNetworks(response, &metadata)
@@ -31715,10 +31088,8 @@ func (m *awsEc2query_deserializeOpDescribeCoipPools) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCoipPools(response, &metadata)
@@ -31818,10 +31189,8 @@ func (m *awsEc2query_deserializeOpDescribeConversionTa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeConversionTasks(response, &metadata)
@@ -31921,10 +31290,8 @@ func (m *awsEc2query_deserializeOpDescribeCustomerGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeCustomerGateways(response, &metadata)
@@ -32024,10 +31391,8 @@ func (m *awsEc2query_deserializeOpDescribeDeclarativeP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeDeclarativePoliciesReports(response, &metadata)
@@ -32127,10 +31492,8 @@ func (m *awsEc2query_deserializeOpDescribeDhcpOptions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeDhcpOptions(response, &metadata)
@@ -32230,10 +31593,8 @@ func (m *awsEc2query_deserializeOpDescribeEgressOnlyIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeEgressOnlyInternetGateways(response, &metadata)
@@ -32333,10 +31694,8 @@ func (m *awsEc2query_deserializeOpDescribeElasticGpus)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeElasticGpus(response, &metadata)
@@ -32436,10 +31795,8 @@ func (m *awsEc2query_deserializeOpDescribeExportImageT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeExportImageTasks(response, &metadata)
@@ -32539,10 +31896,8 @@ func (m *awsEc2query_deserializeOpDescribeExportTasks)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeExportTasks(response, &metadata)
@@ -32642,10 +31997,8 @@ func (m *awsEc2query_deserializeOpDescribeFastLaunchIm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFastLaunchImages(response, &metadata)
@@ -32745,10 +32098,8 @@ func (m *awsEc2query_deserializeOpDescribeFastSnapshot
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFastSnapshotRestores(response, &metadata)
@@ -32848,10 +32199,8 @@ func (m *awsEc2query_deserializeOpDescribeFleetHistory
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFleetHistory(response, &metadata)
@@ -32951,10 +32300,8 @@ func (m *awsEc2query_deserializeOpDescribeFleetInstanc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFleetInstances(response, &metadata)
@@ -33054,10 +32401,8 @@ func (m *awsEc2query_deserializeOpDescribeFleets) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFleets(response, &metadata)
@@ -33157,10 +32502,8 @@ func (m *awsEc2query_deserializeOpDescribeFlowLogs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFlowLogs(response, &metadata)
@@ -33260,10 +32603,8 @@ func (m *awsEc2query_deserializeOpDescribeFpgaImageAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFpgaImageAttribute(response, &metadata)
@@ -33363,10 +32704,8 @@ func (m *awsEc2query_deserializeOpDescribeFpgaImages) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeFpgaImages(response, &metadata)
@@ -33466,10 +32805,8 @@ func (m *awsEc2query_deserializeOpDescribeHostReservat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeHostReservationOfferings(response, &metadata)
@@ -33569,10 +32906,8 @@ func (m *awsEc2query_deserializeOpDescribeHostReservat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeHostReservations(response, &metadata)
@@ -33672,10 +33007,8 @@ func (m *awsEc2query_deserializeOpDescribeHosts) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeHosts(response, &metadata)
@@ -33775,10 +33108,8 @@ func (m *awsEc2query_deserializeOpDescribeIamInstanceP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIamInstanceProfileAssociations(response, &metadata)
@@ -33878,10 +33209,8 @@ func (m *awsEc2query_deserializeOpDescribeIdentityIdFo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIdentityIdFormat(response, &metadata)
@@ -33981,10 +33310,8 @@ func (m *awsEc2query_deserializeOpDescribeIdFormat) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIdFormat(response, &metadata)
@@ -34084,10 +33411,8 @@ func (m *awsEc2query_deserializeOpDescribeImageAttribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeImageAttribute(response, &metadata)
@@ -34187,10 +33512,8 @@ func (m *awsEc2query_deserializeOpDescribeImageReferen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeImageReferences(response, &metadata)
@@ -34290,10 +33613,8 @@ func (m *awsEc2query_deserializeOpDescribeImages) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeImages(response, &metadata)
@@ -34393,10 +33714,8 @@ func (m *awsEc2query_deserializeOpDescribeImageUsageRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeImageUsageReportEntries(response, &metadata)
@@ -34496,10 +33815,8 @@ func (m *awsEc2query_deserializeOpDescribeImageUsageRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeImageUsageReports(response, &metadata)
@@ -34599,10 +33916,8 @@ func (m *awsEc2query_deserializeOpDescribeImportImageT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeImportImageTasks(response, &metadata)
@@ -34702,10 +34017,8 @@ func (m *awsEc2query_deserializeOpDescribeImportSnapsh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeImportSnapshotTasks(response, &metadata)
@@ -34805,10 +34118,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceAttr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceAttribute(response, &metadata)
@@ -34908,10 +34219,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceConnectEndpoints(response, &metadata)
@@ -35011,10 +34320,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceCred
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceCreditSpecifications(response, &metadata)
@@ -35114,10 +34421,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceEven
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceEventNotificationAttributes(response, &metadata)
@@ -35217,10 +34522,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceEven
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceEventWindows(response, &metadata)
@@ -35320,10 +34623,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceImag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceImageMetadata(response, &metadata)
@@ -35423,10 +34724,8 @@ func (m *awsEc2query_deserializeOpDescribeInstances) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstances(response, &metadata)
@@ -35526,10 +34825,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceSqlH
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceSqlHaHistoryStates(response, &metadata)
@@ -35629,10 +34926,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceSqlH
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceSqlHaStates(response, &metadata)
@@ -35732,10 +35027,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceStat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceStatus(response, &metadata)
@@ -35835,10 +35128,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceTopo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceTopology(response, &metadata)
@@ -35938,10 +35229,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceType
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceTypeOfferings(response, &metadata)
@@ -36041,10 +35330,8 @@ func (m *awsEc2query_deserializeOpDescribeInstanceType
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInstanceTypes(response, &metadata)
@@ -36144,10 +35431,8 @@ func (m *awsEc2query_deserializeOpDescribeInternetGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeInternetGateways(response, &metadata)
@@ -36247,10 +35532,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamByoasn) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamByoasn(response, &metadata)
@@ -36350,10 +35633,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamExternal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamExternalResourceVerificationTokens(response, &metadata)
@@ -36453,10 +35734,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamInternet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamInternetRegistryAssociations(response, &metadata)
@@ -36556,10 +35835,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamPolicies
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamPolicies(response, &metadata)
@@ -36659,10 +35936,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamPoolAllo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamPoolAllocations(response, &metadata)
@@ -36762,10 +36037,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamPools) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamPools(response, &metadata)
@@ -36865,10 +36138,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamPrefixLi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamPrefixListResolvers(response, &metadata)
@@ -36968,10 +36239,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamPrefixLi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamPrefixListResolverTargets(response, &metadata)
@@ -37071,10 +36340,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamResourceDiscoveries(response, &metadata)
@@ -37174,10 +36441,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamResourceDiscoveryAssociations(response, &metadata)
@@ -37277,10 +36542,8 @@ func (m *awsEc2query_deserializeOpDescribeIpams) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpams(response, &metadata)
@@ -37380,10 +36643,8 @@ func (m *awsEc2query_deserializeOpDescribeIpamScopes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpamScopes(response, &metadata)
@@ -37483,10 +36744,8 @@ func (m *awsEc2query_deserializeOpDescribeIpv6Pools) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeIpv6Pools(response, &metadata)
@@ -37586,10 +36845,8 @@ func (m *awsEc2query_deserializeOpDescribeKeyPairs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeKeyPairs(response, &metadata)
@@ -37689,10 +36946,8 @@ func (m *awsEc2query_deserializeOpDescribeLaunchTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLaunchTemplates(response, &metadata)
@@ -37792,10 +37047,8 @@ func (m *awsEc2query_deserializeOpDescribeLaunchTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLaunchTemplateVersions(response, &metadata)
@@ -37895,10 +37148,8 @@ func (m *awsEc2query_deserializeOpDescribeLocalGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLocalGatewayRouteTables(response, &metadata)
@@ -37998,10 +37249,8 @@ func (m *awsEc2query_deserializeOpDescribeLocalGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(response, &metadata)
@@ -38101,10 +37350,8 @@ func (m *awsEc2query_deserializeOpDescribeLocalGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLocalGatewayRouteTableVpcAssociations(response, &metadata)
@@ -38204,10 +37451,8 @@ func (m *awsEc2query_deserializeOpDescribeLocalGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLocalGateways(response, &metadata)
@@ -38307,10 +37552,8 @@ func (m *awsEc2query_deserializeOpDescribeLocalGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLocalGatewayVirtualInterfaceGroups(response, &metadata)
@@ -38410,10 +37653,8 @@ func (m *awsEc2query_deserializeOpDescribeLocalGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLocalGatewayVirtualInterfaces(response, &metadata)
@@ -38513,10 +37754,8 @@ func (m *awsEc2query_deserializeOpDescribeLockedSnapsh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeLockedSnapshots(response, &metadata)
@@ -38616,10 +37855,8 @@ func (m *awsEc2query_deserializeOpDescribeMacHosts) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeMacHosts(response, &metadata)
@@ -38719,10 +37956,8 @@ func (m *awsEc2query_deserializeOpDescribeMacModificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeMacModificationTasks(response, &metadata)
@@ -38822,10 +38057,8 @@ func (m *awsEc2query_deserializeOpDescribeManagedPrefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeManagedPrefixLists(response, &metadata)
@@ -38925,10 +38158,8 @@ func (m *awsEc2query_deserializeOpDescribeMovingAddres
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeMovingAddresses(response, &metadata)
@@ -39028,10 +38259,8 @@ func (m *awsEc2query_deserializeOpDescribeNatGateways)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNatGateways(response, &metadata)
@@ -39131,10 +38360,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkAcls)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkAcls(response, &metadata)
@@ -39234,10 +38461,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkInsig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkInsightsAccessScopeAnalyses(response, &metadata)
@@ -39337,10 +38562,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkInsig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkInsightsAccessScopes(response, &metadata)
@@ -39440,10 +38663,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkInsig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkInsightsAnalyses(response, &metadata)
@@ -39543,10 +38764,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkInsig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkInsightsPaths(response, &metadata)
@@ -39646,10 +38865,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkInter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkInterfaceAttribute(response, &metadata)
@@ -39749,10 +38966,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkInter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkInterfacePermissions(response, &metadata)
@@ -39852,10 +39067,8 @@ func (m *awsEc2query_deserializeOpDescribeNetworkInter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeNetworkInterfaces(response, &metadata)
@@ -39955,10 +39168,8 @@ func (m *awsEc2query_deserializeOpDescribeOutpostLags)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeOutpostLags(response, &metadata)
@@ -40058,10 +39269,8 @@ func (m *awsEc2query_deserializeOpDescribePlacementGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribePlacementGroups(response, &metadata)
@@ -40161,10 +39370,8 @@ func (m *awsEc2query_deserializeOpDescribePrefixLists)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribePrefixLists(response, &metadata)
@@ -40264,10 +39471,8 @@ func (m *awsEc2query_deserializeOpDescribePrincipalIdF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribePrincipalIdFormat(response, &metadata)
@@ -40367,10 +39572,8 @@ func (m *awsEc2query_deserializeOpDescribePublicIpv4Po
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribePublicIpv4Pools(response, &metadata)
@@ -40470,10 +39673,8 @@ func (m *awsEc2query_deserializeOpDescribeRegions) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeRegions(response, &metadata)
@@ -40573,10 +39774,8 @@ func (m *awsEc2query_deserializeOpDescribeReplaceRootV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeReplaceRootVolumeTasks(response, &metadata)
@@ -40676,10 +39875,8 @@ func (m *awsEc2query_deserializeOpDescribeReservedInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeReservedInstances(response, &metadata)
@@ -40779,10 +39976,8 @@ func (m *awsEc2query_deserializeOpDescribeReservedInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeReservedInstancesListings(response, &metadata)
@@ -40882,10 +40077,8 @@ func (m *awsEc2query_deserializeOpDescribeReservedInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeReservedInstancesModifications(response, &metadata)
@@ -40985,10 +40178,8 @@ func (m *awsEc2query_deserializeOpDescribeReservedInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeReservedInstancesOfferings(response, &metadata)
@@ -41088,10 +40279,8 @@ func (m *awsEc2query_deserializeOpDescribeRouteServerE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeRouteServerEndpoints(response, &metadata)
@@ -41191,10 +40380,8 @@ func (m *awsEc2query_deserializeOpDescribeRouteServerP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeRouteServerPeers(response, &metadata)
@@ -41294,10 +40481,8 @@ func (m *awsEc2query_deserializeOpDescribeRouteServers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeRouteServers(response, &metadata)
@@ -41397,10 +40582,8 @@ func (m *awsEc2query_deserializeOpDescribeRouteTables)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeRouteTables(response, &metadata)
@@ -41500,10 +40683,8 @@ func (m *awsEc2query_deserializeOpDescribeScheduledIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeScheduledInstanceAvailability(response, &metadata)
@@ -41603,10 +40784,8 @@ func (m *awsEc2query_deserializeOpDescribeScheduledIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeScheduledInstances(response, &metadata)
@@ -41706,10 +40885,8 @@ func (m *awsEc2query_deserializeOpDescribeSecondaryInt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSecondaryInterfaces(response, &metadata)
@@ -41809,10 +40986,8 @@ func (m *awsEc2query_deserializeOpDescribeSecondaryNet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSecondaryNetworks(response, &metadata)
@@ -41912,10 +41087,8 @@ func (m *awsEc2query_deserializeOpDescribeSecondarySub
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSecondarySubnets(response, &metadata)
@@ -42015,10 +41188,8 @@ func (m *awsEc2query_deserializeOpDescribeSecurityGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSecurityGroupReferences(response, &metadata)
@@ -42118,10 +41289,8 @@ func (m *awsEc2query_deserializeOpDescribeSecurityGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSecurityGroupRules(response, &metadata)
@@ -42221,10 +41390,8 @@ func (m *awsEc2query_deserializeOpDescribeSecurityGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSecurityGroups(response, &metadata)
@@ -42324,10 +41491,8 @@ func (m *awsEc2query_deserializeOpDescribeSecurityGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSecurityGroupVpcAssociations(response, &metadata)
@@ -42427,10 +41592,8 @@ func (m *awsEc2query_deserializeOpDescribeServiceLinkV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeServiceLinkVirtualInterfaces(response, &metadata)
@@ -42530,10 +41693,8 @@ func (m *awsEc2query_deserializeOpDescribeSnapshotAttr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSnapshotAttribute(response, &metadata)
@@ -42633,10 +41794,8 @@ func (m *awsEc2query_deserializeOpDescribeSnapshots) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSnapshots(response, &metadata)
@@ -42736,10 +41895,8 @@ func (m *awsEc2query_deserializeOpDescribeSnapshotTier
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSnapshotTierStatus(response, &metadata)
@@ -42839,10 +41996,8 @@ func (m *awsEc2query_deserializeOpDescribeSpotDatafeed
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSpotDatafeedSubscription(response, &metadata)
@@ -42942,10 +42097,8 @@ func (m *awsEc2query_deserializeOpDescribeSpotFleetIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSpotFleetInstances(response, &metadata)
@@ -43045,10 +42198,8 @@ func (m *awsEc2query_deserializeOpDescribeSpotFleetReq
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSpotFleetRequestHistory(response, &metadata)
@@ -43148,10 +42299,8 @@ func (m *awsEc2query_deserializeOpDescribeSpotFleetReq
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSpotFleetRequests(response, &metadata)
@@ -43251,10 +42400,8 @@ func (m *awsEc2query_deserializeOpDescribeSpotInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSpotInstanceRequests(response, &metadata)
@@ -43354,10 +42501,8 @@ func (m *awsEc2query_deserializeOpDescribeSpotPriceHis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSpotPriceHistory(response, &metadata)
@@ -43457,10 +42602,8 @@ func (m *awsEc2query_deserializeOpDescribeStaleSecurit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeStaleSecurityGroups(response, &metadata)
@@ -43560,10 +42703,8 @@ func (m *awsEc2query_deserializeOpDescribeStoreImageTa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeStoreImageTasks(response, &metadata)
@@ -43663,10 +42804,8 @@ func (m *awsEc2query_deserializeOpDescribeSubnets) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeSubnets(response, &metadata)
@@ -43766,10 +42905,8 @@ func (m *awsEc2query_deserializeOpDescribeTags) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTags(response, &metadata)
@@ -43869,10 +43006,8 @@ func (m *awsEc2query_deserializeOpDescribeTrafficMirro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTrafficMirrorFilterRules(response, &metadata)
@@ -43972,10 +43107,8 @@ func (m *awsEc2query_deserializeOpDescribeTrafficMirro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTrafficMirrorFilters(response, &metadata)
@@ -44075,10 +43208,8 @@ func (m *awsEc2query_deserializeOpDescribeTrafficMirro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTrafficMirrorSessions(response, &metadata)
@@ -44178,10 +43309,8 @@ func (m *awsEc2query_deserializeOpDescribeTrafficMirro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTrafficMirrorTargets(response, &metadata)
@@ -44281,10 +43410,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayAttachments(response, &metadata)
@@ -44384,10 +43511,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayConnectPeers(response, &metadata)
@@ -44487,10 +43612,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayConnects(response, &metadata)
@@ -44590,10 +43713,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayMeteringPolicies(response, &metadata)
@@ -44693,10 +43814,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayMulticastDomains(response, &metadata)
@@ -44796,10 +43915,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayPeeringAttachments(response, &metadata)
@@ -44899,10 +44016,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayPolicyTables(response, &metadata)
@@ -45002,10 +44117,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayRouteTableAnnouncements(response, &metadata)
@@ -45105,10 +44218,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayRouteTables(response, &metadata)
@@ -45208,10 +44319,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGateways(response, &metadata)
@@ -45311,10 +44420,8 @@ func (m *awsEc2query_deserializeOpDescribeTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTransitGatewayVpcAttachments(response, &metadata)
@@ -45414,10 +44521,8 @@ func (m *awsEc2query_deserializeOpDescribeTrunkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeTrunkInterfaceAssociations(response, &metadata)
@@ -45517,10 +44622,8 @@ func (m *awsEc2query_deserializeOpDescribeVerifiedAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVerifiedAccessEndpoints(response, &metadata)
@@ -45620,10 +44723,8 @@ func (m *awsEc2query_deserializeOpDescribeVerifiedAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVerifiedAccessGroups(response, &metadata)
@@ -45723,10 +44824,8 @@ func (m *awsEc2query_deserializeOpDescribeVerifiedAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVerifiedAccessInstanceLoggingConfigurations(response, &metadata)
@@ -45826,10 +44925,8 @@ func (m *awsEc2query_deserializeOpDescribeVerifiedAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVerifiedAccessInstances(response, &metadata)
@@ -45929,10 +45026,8 @@ func (m *awsEc2query_deserializeOpDescribeVerifiedAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVerifiedAccessTrustProviders(response, &metadata)
@@ -46032,10 +45127,8 @@ func (m *awsEc2query_deserializeOpDescribeVolumeAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVolumeAttribute(response, &metadata)
@@ -46135,10 +45228,8 @@ func (m *awsEc2query_deserializeOpDescribeVolumes) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVolumes(response, &metadata)
@@ -46238,10 +45329,8 @@ func (m *awsEc2query_deserializeOpDescribeVolumesModif
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVolumesModifications(response, &metadata)
@@ -46341,10 +45430,8 @@ func (m *awsEc2query_deserializeOpDescribeVolumeStatus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVolumeStatus(response, &metadata)
@@ -46444,10 +45531,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcAttribute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcAttribute(response, &metadata)
@@ -46547,10 +45632,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcBlockPubl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcBlockPublicAccessExclusions(response, &metadata)
@@ -46650,10 +45733,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcBlockPubl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcBlockPublicAccessOptions(response, &metadata)
@@ -46753,10 +45834,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcClassicLi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcClassicLink(response, &metadata)
@@ -46856,10 +45935,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcClassicLi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcClassicLinkDnsSupport(response, &metadata)
@@ -46959,10 +46036,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEncryptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEncryptionControls(response, &metadata)
@@ -47062,10 +46137,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEndpointA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEndpointAssociations(response, &metadata)
@@ -47165,10 +46238,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEndpointC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEndpointConnectionNotifications(response, &metadata)
@@ -47268,10 +46339,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEndpointC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEndpointConnections(response, &metadata)
@@ -47371,10 +46440,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEndpoints
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEndpoints(response, &metadata)
@@ -47474,10 +46541,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEndpointS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEndpointServiceConfigurations(response, &metadata)
@@ -47577,10 +46642,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEndpointS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEndpointServicePermissions(response, &metadata)
@@ -47680,10 +46743,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcEndpointS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcEndpointServices(response, &metadata)
@@ -47783,10 +46844,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcPeeringCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcPeeringConnections(response, &metadata)
@@ -47886,10 +46945,8 @@ func (m *awsEc2query_deserializeOpDescribeVpcs) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpcs(response, &metadata)
@@ -47989,10 +47046,8 @@ func (m *awsEc2query_deserializeOpDescribeVpnConcentra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpnConcentrators(response, &metadata)
@@ -48092,10 +47147,8 @@ func (m *awsEc2query_deserializeOpDescribeVpnConnectio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpnConnections(response, &metadata)
@@ -48195,10 +47248,8 @@ func (m *awsEc2query_deserializeOpDescribeVpnGateways)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDescribeVpnGateways(response, &metadata)
@@ -48298,10 +47349,8 @@ func (m *awsEc2query_deserializeOpDetachClassicLinkVpc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDetachClassicLinkVpc(response, &metadata)
@@ -48401,10 +47450,8 @@ func (m *awsEc2query_deserializeOpDetachImageWatermark
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDetachImageWatermark(response, &metadata)
@@ -48504,10 +47551,8 @@ func (m *awsEc2query_deserializeOpDetachInternetGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDetachInternetGateway(response, &metadata)
@@ -48584,10 +47629,8 @@ func (m *awsEc2query_deserializeOpDetachNetworkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDetachNetworkInterface(response, &metadata)
@@ -48664,10 +47707,8 @@ func (m *awsEc2query_deserializeOpDetachVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDetachVerifiedAccessTrustProvider(response, &metadata)
@@ -48767,10 +47808,8 @@ func (m *awsEc2query_deserializeOpDetachVolume) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDetachVolume(response, &metadata)
@@ -48870,10 +47909,8 @@ func (m *awsEc2query_deserializeOpDetachVpnGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDetachVpnGateway(response, &metadata)
@@ -48950,10 +47987,8 @@ func (m *awsEc2query_deserializeOpDisableAddressTransf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableAddressTransfer(response, &metadata)
@@ -49053,10 +48088,8 @@ func (m *awsEc2query_deserializeOpDisableAllowedImages
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableAllowedImagesSettings(response, &metadata)
@@ -49156,10 +48189,8 @@ func (m *awsEc2query_deserializeOpDisableApplicationSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableApplicationStatusCheckSuppression(response, &metadata)
@@ -49259,10 +48290,8 @@ func (m *awsEc2query_deserializeOpDisableAwsNetworkPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableAwsNetworkPerformanceMetricSubscription(response, &metadata)
@@ -49362,10 +48391,8 @@ func (m *awsEc2query_deserializeOpDisableCapacityManag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableCapacityManager(response, &metadata)
@@ -49465,10 +48492,8 @@ func (m *awsEc2query_deserializeOpDisableEbsEncryption
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableEbsEncryptionByDefault(response, &metadata)
@@ -49568,10 +48593,8 @@ func (m *awsEc2query_deserializeOpDisableFastLaunch) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableFastLaunch(response, &metadata)
@@ -49671,10 +48694,8 @@ func (m *awsEc2query_deserializeOpDisableFastSnapshotR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableFastSnapshotRestores(response, &metadata)
@@ -49774,10 +48795,8 @@ func (m *awsEc2query_deserializeOpDisableImage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableImage(response, &metadata)
@@ -49877,10 +48896,8 @@ func (m *awsEc2query_deserializeOpDisableImageBlockPub
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableImageBlockPublicAccess(response, &metadata)
@@ -49980,10 +48997,8 @@ func (m *awsEc2query_deserializeOpDisableImageDeprecat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableImageDeprecation(response, &metadata)
@@ -50083,10 +49098,8 @@ func (m *awsEc2query_deserializeOpDisableImageDeregist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableImageDeregistrationProtection(response, &metadata)
@@ -50186,10 +49199,8 @@ func (m *awsEc2query_deserializeOpDisableInstanceSqlHa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableInstanceSqlHaStandbyDetections(response, &metadata)
@@ -50289,10 +49300,8 @@ func (m *awsEc2query_deserializeOpDisableIpamOrganizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableIpamOrganizationAdminAccount(response, &metadata)
@@ -50392,10 +49401,8 @@ func (m *awsEc2query_deserializeOpDisableIpamPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableIpamPolicy(response, &metadata)
@@ -50495,10 +49502,8 @@ func (m *awsEc2query_deserializeOpDisableRouteServerPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableRouteServerPropagation(response, &metadata)
@@ -50598,10 +49603,8 @@ func (m *awsEc2query_deserializeOpDisableSerialConsole
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableSerialConsoleAccess(response, &metadata)
@@ -50701,10 +49704,8 @@ func (m *awsEc2query_deserializeOpDisableSnapshotBlock
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableSnapshotBlockPublicAccess(response, &metadata)
@@ -50804,10 +49805,8 @@ func (m *awsEc2query_deserializeOpDisableTransitGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableTransitGatewayRouteTablePropagation(response, &metadata)
@@ -50907,10 +49906,8 @@ func (m *awsEc2query_deserializeOpDisableVgwRoutePropa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableVgwRoutePropagation(response, &metadata)
@@ -50987,10 +49984,8 @@ func (m *awsEc2query_deserializeOpDisableVpcClassicLin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableVpcClassicLink(response, &metadata)
@@ -51090,10 +50085,8 @@ func (m *awsEc2query_deserializeOpDisableVpcClassicLin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisableVpcClassicLinkDnsSupport(response, &metadata)
@@ -51193,10 +50186,8 @@ func (m *awsEc2query_deserializeOpDisassociateAddress)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateAddress(response, &metadata)
@@ -51273,10 +50264,8 @@ func (m *awsEc2query_deserializeOpDisassociateApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateApplicationStatusCheck(response, &metadata)
@@ -51376,10 +50365,8 @@ func (m *awsEc2query_deserializeOpDisassociateCapacity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateCapacityReservationBillingOwner(response, &metadata)
@@ -51479,10 +50466,8 @@ func (m *awsEc2query_deserializeOpDisassociateClientVp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateClientVpnTargetNetwork(response, &metadata)
@@ -51582,10 +50567,8 @@ func (m *awsEc2query_deserializeOpDisassociateEnclaveC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateEnclaveCertificateIamRole(response, &metadata)
@@ -51685,10 +50668,8 @@ func (m *awsEc2query_deserializeOpDisassociateIamInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateIamInstanceProfile(response, &metadata)
@@ -51788,10 +50769,8 @@ func (m *awsEc2query_deserializeOpDisassociateInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateInstanceEventWindow(response, &metadata)
@@ -51891,10 +50870,8 @@ func (m *awsEc2query_deserializeOpDisassociateIpamByoa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateIpamByoasn(response, &metadata)
@@ -51994,10 +50971,8 @@ func (m *awsEc2query_deserializeOpDisassociateIpamReso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateIpamResourceDiscovery(response, &metadata)
@@ -52097,10 +51072,8 @@ func (m *awsEc2query_deserializeOpDisassociateNatGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateNatGatewayAddress(response, &metadata)
@@ -52200,10 +51173,8 @@ func (m *awsEc2query_deserializeOpDisassociateRouteSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateRouteServer(response, &metadata)
@@ -52303,10 +51274,8 @@ func (m *awsEc2query_deserializeOpDisassociateRouteTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateRouteTable(response, &metadata)
@@ -52383,10 +51352,8 @@ func (m *awsEc2query_deserializeOpDisassociateSecurity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateSecurityGroupVpc(response, &metadata)
@@ -52486,10 +51453,8 @@ func (m *awsEc2query_deserializeOpDisassociateSubnetCi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateSubnetCidrBlock(response, &metadata)
@@ -52589,10 +51554,8 @@ func (m *awsEc2query_deserializeOpDisassociateTransitG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateTransitGatewayMulticastDomain(response, &metadata)
@@ -52692,10 +51655,8 @@ func (m *awsEc2query_deserializeOpDisassociateTransitG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateTransitGatewayPolicyTable(response, &metadata)
@@ -52795,10 +51756,8 @@ func (m *awsEc2query_deserializeOpDisassociateTransitG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateTransitGatewayRouteTable(response, &metadata)
@@ -52898,10 +51857,8 @@ func (m *awsEc2query_deserializeOpDisassociateTrunkInt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateTrunkInterface(response, &metadata)
@@ -53001,10 +51958,8 @@ func (m *awsEc2query_deserializeOpDisassociateVpcCidrB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorDisassociateVpcCidrBlock(response, &metadata)
@@ -53104,10 +52059,8 @@ func (m *awsEc2query_deserializeOpEnableAddressTransfe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableAddressTransfer(response, &metadata)
@@ -53207,10 +52160,8 @@ func (m *awsEc2query_deserializeOpEnableAllowedImagesS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableAllowedImagesSettings(response, &metadata)
@@ -53310,10 +52261,8 @@ func (m *awsEc2query_deserializeOpEnableApplicationSta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableApplicationStatusCheckSuppression(response, &metadata)
@@ -53413,10 +52362,8 @@ func (m *awsEc2query_deserializeOpEnableAwsNetworkPerf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableAwsNetworkPerformanceMetricSubscription(response, &metadata)
@@ -53516,10 +52463,8 @@ func (m *awsEc2query_deserializeOpEnableCapacityManage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableCapacityManager(response, &metadata)
@@ -53619,10 +52564,8 @@ func (m *awsEc2query_deserializeOpEnableEbsEncryptionB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableEbsEncryptionByDefault(response, &metadata)
@@ -53722,10 +52665,8 @@ func (m *awsEc2query_deserializeOpEnableFastLaunch) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableFastLaunch(response, &metadata)
@@ -53825,10 +52766,8 @@ func (m *awsEc2query_deserializeOpEnableFastSnapshotRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableFastSnapshotRestores(response, &metadata)
@@ -53928,10 +52867,8 @@ func (m *awsEc2query_deserializeOpEnableImage) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableImage(response, &metadata)
@@ -54031,10 +52968,8 @@ func (m *awsEc2query_deserializeOpEnableImageBlockPubl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableImageBlockPublicAccess(response, &metadata)
@@ -54134,10 +53069,8 @@ func (m *awsEc2query_deserializeOpEnableImageDeprecati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableImageDeprecation(response, &metadata)
@@ -54237,10 +53170,8 @@ func (m *awsEc2query_deserializeOpEnableImageDeregistr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableImageDeregistrationProtection(response, &metadata)
@@ -54340,10 +53271,8 @@ func (m *awsEc2query_deserializeOpEnableInstanceSqlHaS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableInstanceSqlHaStandbyDetections(response, &metadata)
@@ -54443,10 +53372,8 @@ func (m *awsEc2query_deserializeOpEnableIpamInternetRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableIpamInternetRegistryAssociation(response, &metadata)
@@ -54546,10 +53473,8 @@ func (m *awsEc2query_deserializeOpEnableIpamOrganizati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableIpamOrganizationAdminAccount(response, &metadata)
@@ -54649,10 +53574,8 @@ func (m *awsEc2query_deserializeOpEnableIpamPolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableIpamPolicy(response, &metadata)
@@ -54752,10 +53675,8 @@ func (m *awsEc2query_deserializeOpEnableReachabilityAn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableReachabilityAnalyzerOrganizationSharing(response, &metadata)
@@ -54855,10 +53776,8 @@ func (m *awsEc2query_deserializeOpEnableRouteServerPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableRouteServerPropagation(response, &metadata)
@@ -54958,10 +53877,8 @@ func (m *awsEc2query_deserializeOpEnableSerialConsoleA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableSerialConsoleAccess(response, &metadata)
@@ -55061,10 +53978,8 @@ func (m *awsEc2query_deserializeOpEnableSnapshotBlockP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableSnapshotBlockPublicAccess(response, &metadata)
@@ -55164,10 +54079,8 @@ func (m *awsEc2query_deserializeOpEnableTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableTransitGatewayRouteTablePropagation(response, &metadata)
@@ -55267,10 +54180,8 @@ func (m *awsEc2query_deserializeOpEnableVgwRoutePropag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableVgwRoutePropagation(response, &metadata)
@@ -55347,10 +54258,8 @@ func (m *awsEc2query_deserializeOpEnableVolumeIO) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableVolumeIO(response, &metadata)
@@ -55427,10 +54336,8 @@ func (m *awsEc2query_deserializeOpEnableVpcClassicLink
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableVpcClassicLink(response, &metadata)
@@ -55530,10 +54437,8 @@ func (m *awsEc2query_deserializeOpEnableVpcClassicLink
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorEnableVpcClassicLinkDnsSupport(response, &metadata)
@@ -55633,10 +54538,8 @@ func (m *awsEc2query_deserializeOpExportClientVpnClien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorExportClientVpnClientCertificateRevocationList(response, &metadata)
@@ -55736,10 +54639,8 @@ func (m *awsEc2query_deserializeOpExportClientVpnClien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorExportClientVpnClientConfiguration(response, &metadata)
@@ -55839,10 +54740,8 @@ func (m *awsEc2query_deserializeOpExportImage) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorExportImage(response, &metadata)
@@ -55942,10 +54841,8 @@ func (m *awsEc2query_deserializeOpExportTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorExportTransitGatewayRoutes(response, &metadata)
@@ -56045,10 +54942,8 @@ func (m *awsEc2query_deserializeOpExportVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorExportVerifiedAccessInstanceClientConfiguration(response, &metadata)
@@ -56148,10 +55043,8 @@ func (m *awsEc2query_deserializeOpGetActiveVpnTunnelSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetActiveVpnTunnelStatus(response, &metadata)
@@ -56251,10 +55144,8 @@ func (m *awsEc2query_deserializeOpGetAllowedImagesSett
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetAllowedImagesSettings(response, &metadata)
@@ -56354,10 +55245,8 @@ func (m *awsEc2query_deserializeOpGetAssociatedEnclave
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetAssociatedEnclaveCertificateIamRoles(response, &metadata)
@@ -56457,10 +55346,8 @@ func (m *awsEc2query_deserializeOpGetAssociatedIpv6Poo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetAssociatedIpv6PoolCidrs(response, &metadata)
@@ -56560,10 +55447,8 @@ func (m *awsEc2query_deserializeOpGetAwsNetworkPerform
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetAwsNetworkPerformanceData(response, &metadata)
@@ -56663,10 +55548,8 @@ func (m *awsEc2query_deserializeOpGetCapacityManagerAt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetCapacityManagerAttributes(response, &metadata)
@@ -56766,10 +55649,8 @@ func (m *awsEc2query_deserializeOpGetCapacityManagerMe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetCapacityManagerMetricData(response, &metadata)
@@ -56869,10 +55750,8 @@ func (m *awsEc2query_deserializeOpGetCapacityManagerMe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetCapacityManagerMetricDimensions(response, &metadata)
@@ -56972,10 +55851,8 @@ func (m *awsEc2query_deserializeOpGetCapacityManagerMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetCapacityManagerMonitoredTagKeys(response, &metadata)
@@ -57075,10 +55952,8 @@ func (m *awsEc2query_deserializeOpGetCapacityManagerSu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetCapacityManagerSupportabilityScores(response, &metadata)
@@ -57178,10 +56053,8 @@ func (m *awsEc2query_deserializeOpGetCapacityReservati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetCapacityReservationUsage(response, &metadata)
@@ -57281,10 +56154,8 @@ func (m *awsEc2query_deserializeOpGetClientVpnEndpoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetClientVpnEndpointAuthorizationPolicy(response, &metadata)
@@ -57384,10 +56255,8 @@ func (m *awsEc2query_deserializeOpGetCoipPoolUsage) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetCoipPoolUsage(response, &metadata)
@@ -57487,10 +56356,8 @@ func (m *awsEc2query_deserializeOpGetConsoleOutput) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetConsoleOutput(response, &metadata)
@@ -57590,10 +56457,8 @@ func (m *awsEc2query_deserializeOpGetConsoleScreenshot
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetConsoleScreenshot(response, &metadata)
@@ -57693,10 +56558,8 @@ func (m *awsEc2query_deserializeOpGetDeclarativePolici
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetDeclarativePoliciesReportSummary(response, &metadata)
@@ -57796,10 +56659,8 @@ func (m *awsEc2query_deserializeOpGetDefaultCreditSpec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetDefaultCreditSpecification(response, &metadata)
@@ -57899,10 +56760,8 @@ func (m *awsEc2query_deserializeOpGetEbsDefaultKmsKeyI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetEbsDefaultKmsKeyId(response, &metadata)
@@ -58002,10 +56861,8 @@ func (m *awsEc2query_deserializeOpGetEbsEncryptionByDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetEbsEncryptionByDefault(response, &metadata)
@@ -58105,10 +56962,8 @@ func (m *awsEc2query_deserializeOpGetEnabledIpamPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetEnabledIpamPolicy(response, &metadata)
@@ -58208,10 +57063,8 @@ func (m *awsEc2query_deserializeOpGetFlowLogsIntegrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetFlowLogsIntegrationTemplate(response, &metadata)
@@ -58311,10 +57164,8 @@ func (m *awsEc2query_deserializeOpGetGroupsForCapacity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetGroupsForCapacityReservation(response, &metadata)
@@ -58414,10 +57265,8 @@ func (m *awsEc2query_deserializeOpGetHostReservationPu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetHostReservationPurchasePreview(response, &metadata)
@@ -58517,10 +57366,8 @@ func (m *awsEc2query_deserializeOpGetImageAncestry) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetImageAncestry(response, &metadata)
@@ -58620,10 +57467,8 @@ func (m *awsEc2query_deserializeOpGetImageBlockPublicA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetImageBlockPublicAccessState(response, &metadata)
@@ -58723,10 +57568,8 @@ func (m *awsEc2query_deserializeOpGetImageFindings) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetImageFindings(response, &metadata)
@@ -58826,10 +57669,8 @@ func (m *awsEc2query_deserializeOpGetInstanceMetadataD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetInstanceMetadataDefaults(response, &metadata)
@@ -58929,10 +57770,8 @@ func (m *awsEc2query_deserializeOpGetInstanceTpmEkPub)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetInstanceTpmEkPub(response, &metadata)
@@ -59032,10 +57871,8 @@ func (m *awsEc2query_deserializeOpGetInstanceTypesFrom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetInstanceTypesFromInstanceRequirements(response, &metadata)
@@ -59135,10 +57972,8 @@ func (m *awsEc2query_deserializeOpGetInstanceUefiData)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetInstanceUefiData(response, &metadata)
@@ -59238,10 +58073,8 @@ func (m *awsEc2query_deserializeOpGetIpamAddressHistor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamAddressHistory(response, &metadata)
@@ -59341,10 +58174,8 @@ func (m *awsEc2query_deserializeOpGetIpamDiscoveredAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamDiscoveredAccounts(response, &metadata)
@@ -59444,10 +58275,8 @@ func (m *awsEc2query_deserializeOpGetIpamDiscoveredPub
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamDiscoveredPublicAddresses(response, &metadata)
@@ -59547,10 +58376,8 @@ func (m *awsEc2query_deserializeOpGetIpamDiscoveredRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamDiscoveredResourceCidrs(response, &metadata)
@@ -59650,10 +58477,8 @@ func (m *awsEc2query_deserializeOpGetIpamDiscoveredRou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamDiscoveredRoutes(response, &metadata)
@@ -59753,10 +58578,8 @@ func (m *awsEc2query_deserializeOpGetIpamInternetRegis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamInternetRegistryAssociationAsns(response, &metadata)
@@ -59856,10 +58679,8 @@ func (m *awsEc2query_deserializeOpGetIpamInternetRegis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamInternetRegistryAssociationCidrs(response, &metadata)
@@ -59959,10 +58780,8 @@ func (m *awsEc2query_deserializeOpGetIpamPolicyAllocat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamPolicyAllocationRules(response, &metadata)
@@ -60062,10 +58881,8 @@ func (m *awsEc2query_deserializeOpGetIpamPolicyOrganiz
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamPolicyOrganizationTargets(response, &metadata)
@@ -60165,10 +58982,8 @@ func (m *awsEc2query_deserializeOpGetIpamPoolAllocatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamPoolAllocations(response, &metadata)
@@ -60268,10 +59083,8 @@ func (m *awsEc2query_deserializeOpGetIpamPoolCidrs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamPoolCidrs(response, &metadata)
@@ -60371,10 +59184,8 @@ func (m *awsEc2query_deserializeOpGetIpamPrefixListRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamPrefixListResolverRules(response, &metadata)
@@ -60474,10 +59285,8 @@ func (m *awsEc2query_deserializeOpGetIpamPrefixListRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamPrefixListResolverVersionEntries(response, &metadata)
@@ -60577,10 +59386,8 @@ func (m *awsEc2query_deserializeOpGetIpamPrefixListRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamPrefixListResolverVersions(response, &metadata)
@@ -60680,10 +59487,8 @@ func (m *awsEc2query_deserializeOpGetIpamResourceCidrs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamResourceCidrs(response, &metadata)
@@ -60783,10 +59588,8 @@ func (m *awsEc2query_deserializeOpGetIpamRouteOriginAu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamRouteOriginAuthorizations(response, &metadata)
@@ -60886,10 +59689,8 @@ func (m *awsEc2query_deserializeOpGetIpamRouteProtecti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamRouteProtectionFindings(response, &metadata)
@@ -60989,10 +59790,8 @@ func (m *awsEc2query_deserializeOpGetIpamRoutingPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamRoutingPolicyRegistrationDeltas(response, &metadata)
@@ -61092,10 +59891,8 @@ func (m *awsEc2query_deserializeOpGetIpamRoutingPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetIpamRoutingPolicyRegistrations(response, &metadata)
@@ -61195,10 +59992,8 @@ func (m *awsEc2query_deserializeOpGetLaunchTemplateDat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetLaunchTemplateData(response, &metadata)
@@ -61298,10 +60093,8 @@ func (m *awsEc2query_deserializeOpGetManagedPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetManagedPrefixListAssociations(response, &metadata)
@@ -61401,10 +60194,8 @@ func (m *awsEc2query_deserializeOpGetManagedPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetManagedPrefixListEntries(response, &metadata)
@@ -61504,10 +60295,8 @@ func (m *awsEc2query_deserializeOpGetManagedResourceVi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetManagedResourceVisibility(response, &metadata)
@@ -61607,10 +60396,8 @@ func (m *awsEc2query_deserializeOpGetNetworkInsightsAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetNetworkInsightsAccessScopeAnalysisFindings(response, &metadata)
@@ -61710,10 +60497,8 @@ func (m *awsEc2query_deserializeOpGetNetworkInsightsAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetNetworkInsightsAccessScopeContent(response, &metadata)
@@ -61813,10 +60598,8 @@ func (m *awsEc2query_deserializeOpGetPasswordData) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetPasswordData(response, &metadata)
@@ -61916,10 +60699,8 @@ func (m *awsEc2query_deserializeOpGetReservedInstances
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetReservedInstancesExchangeQuote(response, &metadata)
@@ -62019,10 +60800,8 @@ func (m *awsEc2query_deserializeOpGetRouteServerAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetRouteServerAssociations(response, &metadata)
@@ -62122,10 +60901,8 @@ func (m *awsEc2query_deserializeOpGetRouteServerPropag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetRouteServerPropagations(response, &metadata)
@@ -62225,10 +61002,8 @@ func (m *awsEc2query_deserializeOpGetRouteServerRoutin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetRouteServerRoutingDatabase(response, &metadata)
@@ -62328,10 +61103,8 @@ func (m *awsEc2query_deserializeOpGetSecurityGroupsFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetSecurityGroupsForVpc(response, &metadata)
@@ -62431,10 +61204,8 @@ func (m *awsEc2query_deserializeOpGetSerialConsoleAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetSerialConsoleAccessStatus(response, &metadata)
@@ -62534,10 +61305,8 @@ func (m *awsEc2query_deserializeOpGetSnapshotBlockPubl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetSnapshotBlockPublicAccessState(response, &metadata)
@@ -62637,10 +61406,8 @@ func (m *awsEc2query_deserializeOpGetSpotPlacementScor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetSpotPlacementScores(response, &metadata)
@@ -62740,10 +61507,8 @@ func (m *awsEc2query_deserializeOpGetSubnetCidrReserva
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetSubnetCidrReservations(response, &metadata)
@@ -62843,10 +61608,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayAttachmentNatConfiguration(response, &metadata)
@@ -62946,10 +61709,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayAttachmentPropagations(response, &metadata)
@@ -63049,10 +61810,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayMet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayMeteringPolicyEntries(response, &metadata)
@@ -63152,10 +61911,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayMul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayMulticastDomainAssociations(response, &metadata)
@@ -63255,10 +62012,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayPolicyTableAssociations(response, &metadata)
@@ -63358,10 +62113,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayPolicyTableEntries(response, &metadata)
@@ -63461,10 +62214,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayPre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayPrefixListReferences(response, &metadata)
@@ -63564,10 +62315,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayRou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayRouteTableAssociations(response, &metadata)
@@ -63667,10 +62416,8 @@ func (m *awsEc2query_deserializeOpGetTransitGatewayRou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetTransitGatewayRouteTablePropagations(response, &metadata)
@@ -63770,10 +62517,8 @@ func (m *awsEc2query_deserializeOpGetVerifiedAccessEnd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetVerifiedAccessEndpointPolicy(response, &metadata)
@@ -63873,10 +62618,8 @@ func (m *awsEc2query_deserializeOpGetVerifiedAccessEnd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetVerifiedAccessEndpointTargets(response, &metadata)
@@ -63976,10 +62719,8 @@ func (m *awsEc2query_deserializeOpGetVerifiedAccessGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetVerifiedAccessGroupPolicy(response, &metadata)
@@ -64079,10 +62820,8 @@ func (m *awsEc2query_deserializeOpGetVpcResourcesBlock
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetVpcResourcesBlockingEncryptionEnforcement(response, &metadata)
@@ -64182,10 +62921,8 @@ func (m *awsEc2query_deserializeOpGetVpnConnectionDevi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetVpnConnectionDeviceSampleConfiguration(response, &metadata)
@@ -64285,10 +63022,8 @@ func (m *awsEc2query_deserializeOpGetVpnConnectionDevi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetVpnConnectionDeviceTypes(response, &metadata)
@@ -64388,10 +63123,8 @@ func (m *awsEc2query_deserializeOpGetVpnTunnelReplacem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorGetVpnTunnelReplacementStatus(response, &metadata)
@@ -64491,10 +63224,8 @@ func (m *awsEc2query_deserializeOpImportClientVpnClien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorImportClientVpnClientCertificateRevocationList(response, &metadata)
@@ -64594,10 +63325,8 @@ func (m *awsEc2query_deserializeOpImportImage) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorImportImage(response, &metadata)
@@ -64697,10 +63426,8 @@ func (m *awsEc2query_deserializeOpImportInstance) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorImportInstance(response, &metadata)
@@ -64800,10 +63527,8 @@ func (m *awsEc2query_deserializeOpImportKeyPair) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorImportKeyPair(response, &metadata)
@@ -64903,10 +63628,8 @@ func (m *awsEc2query_deserializeOpImportSnapshot) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorImportSnapshot(response, &metadata)
@@ -65006,10 +63729,8 @@ func (m *awsEc2query_deserializeOpImportVolume) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorImportVolume(response, &metadata)
@@ -65109,10 +63830,8 @@ func (m *awsEc2query_deserializeOpListImagesInRecycleB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorListImagesInRecycleBin(response, &metadata)
@@ -65212,10 +63931,8 @@ func (m *awsEc2query_deserializeOpListSnapshotsInRecyc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorListSnapshotsInRecycleBin(response, &metadata)
@@ -65315,10 +64032,8 @@ func (m *awsEc2query_deserializeOpListVolumesInRecycle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorListVolumesInRecycleBin(response, &metadata)
@@ -65418,10 +64133,8 @@ func (m *awsEc2query_deserializeOpLockSnapshot) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorLockSnapshot(response, &metadata)
@@ -65521,10 +64234,8 @@ func (m *awsEc2query_deserializeOpModifyAccountVpcEncr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyAccountVpcEncryptionControl(response, &metadata)
@@ -65624,10 +64335,8 @@ func (m *awsEc2query_deserializeOpModifyAddressAttribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyAddressAttribute(response, &metadata)
@@ -65727,10 +64436,8 @@ func (m *awsEc2query_deserializeOpModifyApplicationSta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyApplicationStatusCheck(response, &metadata)
@@ -65830,10 +64537,8 @@ func (m *awsEc2query_deserializeOpModifyAvailabilityZo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyAvailabilityZoneGroup(response, &metadata)
@@ -65933,10 +64638,8 @@ func (m *awsEc2query_deserializeOpModifyCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyCapacityReservation(response, &metadata)
@@ -66036,10 +64739,8 @@ func (m *awsEc2query_deserializeOpModifyCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyCapacityReservationFleet(response, &metadata)
@@ -66139,10 +64840,8 @@ func (m *awsEc2query_deserializeOpModifyClientVpnEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyClientVpnEndpoint(response, &metadata)
@@ -66242,10 +64941,8 @@ func (m *awsEc2query_deserializeOpModifyClientVpnEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyClientVpnEndpointAuthorizationPolicy(response, &metadata)
@@ -66345,10 +65042,8 @@ func (m *awsEc2query_deserializeOpModifyDefaultCreditS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyDefaultCreditSpecification(response, &metadata)
@@ -66448,10 +65143,8 @@ func (m *awsEc2query_deserializeOpModifyEbsDefaultKmsK
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyEbsDefaultKmsKeyId(response, &metadata)
@@ -66551,10 +65244,8 @@ func (m *awsEc2query_deserializeOpModifyFleet) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyFleet(response, &metadata)
@@ -66654,10 +65345,8 @@ func (m *awsEc2query_deserializeOpModifyFpgaImageAttri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyFpgaImageAttribute(response, &metadata)
@@ -66757,10 +65446,8 @@ func (m *awsEc2query_deserializeOpModifyHosts) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyHosts(response, &metadata)
@@ -66860,10 +65547,8 @@ func (m *awsEc2query_deserializeOpModifyIdentityIdForm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIdentityIdFormat(response, &metadata)
@@ -66940,10 +65625,8 @@ func (m *awsEc2query_deserializeOpModifyIdFormat) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIdFormat(response, &metadata)
@@ -67020,10 +65703,8 @@ func (m *awsEc2query_deserializeOpModifyImageAttribute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyImageAttribute(response, &metadata)
@@ -67100,10 +65781,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceAttribute(response, &metadata)
@@ -67180,10 +65859,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceCapaci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceCapacityReservationAttributes(response, &metadata)
@@ -67283,10 +65960,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceConnectEndpoint(response, &metadata)
@@ -67386,10 +66061,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceCpuOpt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceCpuOptions(response, &metadata)
@@ -67489,10 +66162,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceCredit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceCreditSpecification(response, &metadata)
@@ -67592,10 +66263,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceEventS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceEventStartTime(response, &metadata)
@@ -67695,10 +66364,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceEventW
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceEventWindow(response, &metadata)
@@ -67798,10 +66465,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceMainte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceMaintenanceOptions(response, &metadata)
@@ -67901,10 +66566,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceMetada
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceMetadataDefaults(response, &metadata)
@@ -68004,10 +66667,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceMetada
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceMetadataOptions(response, &metadata)
@@ -68107,10 +66768,8 @@ func (m *awsEc2query_deserializeOpModifyInstanceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstanceNetworkPerformanceOptions(response, &metadata)
@@ -68210,10 +66869,8 @@ func (m *awsEc2query_deserializeOpModifyInstancePlacem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyInstancePlacement(response, &metadata)
@@ -68313,10 +66970,8 @@ func (m *awsEc2query_deserializeOpModifyIpam) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpam(response, &metadata)
@@ -68416,10 +67071,8 @@ func (m *awsEc2query_deserializeOpModifyIpamPolicyAllo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamPolicyAllocationRules(response, &metadata)
@@ -68519,10 +67172,8 @@ func (m *awsEc2query_deserializeOpModifyIpamPool) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamPool(response, &metadata)
@@ -68622,10 +67273,8 @@ func (m *awsEc2query_deserializeOpModifyIpamPoolAlloca
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamPoolAllocation(response, &metadata)
@@ -68725,10 +67374,8 @@ func (m *awsEc2query_deserializeOpModifyIpamPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamPrefixListResolver(response, &metadata)
@@ -68828,10 +67475,8 @@ func (m *awsEc2query_deserializeOpModifyIpamPrefixList
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamPrefixListResolverTarget(response, &metadata)
@@ -68931,10 +67576,8 @@ func (m *awsEc2query_deserializeOpModifyIpamResourceCi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamResourceCidr(response, &metadata)
@@ -69034,10 +67677,8 @@ func (m *awsEc2query_deserializeOpModifyIpamResourceDi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamResourceDiscovery(response, &metadata)
@@ -69137,10 +67778,8 @@ func (m *awsEc2query_deserializeOpModifyIpamRoutingPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamRoutingPolicyRegistration(response, &metadata)
@@ -69240,10 +67879,8 @@ func (m *awsEc2query_deserializeOpModifyIpamScope) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyIpamScope(response, &metadata)
@@ -69343,10 +67980,8 @@ func (m *awsEc2query_deserializeOpModifyLaunchTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyLaunchTemplate(response, &metadata)
@@ -69446,10 +68081,8 @@ func (m *awsEc2query_deserializeOpModifyLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyLocalGatewayRoute(response, &metadata)
@@ -69549,10 +68182,8 @@ func (m *awsEc2query_deserializeOpModifyManagedPrefixL
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyManagedPrefixList(response, &metadata)
@@ -69652,10 +68283,8 @@ func (m *awsEc2query_deserializeOpModifyManagedResourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyManagedResourceVisibility(response, &metadata)
@@ -69755,10 +68384,8 @@ func (m *awsEc2query_deserializeOpModifyNetworkInterfa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyNetworkInterfaceAttribute(response, &metadata)
@@ -69835,10 +68462,8 @@ func (m *awsEc2query_deserializeOpModifyPrivateDnsName
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyPrivateDnsNameOptions(response, &metadata)
@@ -69938,10 +68563,8 @@ func (m *awsEc2query_deserializeOpModifyPublicIpDnsNam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyPublicIpDnsNameOptions(response, &metadata)
@@ -70041,10 +68664,8 @@ func (m *awsEc2query_deserializeOpModifyReservedInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyReservedInstances(response, &metadata)
@@ -70144,10 +68765,8 @@ func (m *awsEc2query_deserializeOpModifyRouteServer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyRouteServer(response, &metadata)
@@ -70247,10 +68866,8 @@ func (m *awsEc2query_deserializeOpModifySecurityGroupR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifySecurityGroupRules(response, &metadata)
@@ -70350,10 +68967,8 @@ func (m *awsEc2query_deserializeOpModifySnapshotAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifySnapshotAttribute(response, &metadata)
@@ -70430,10 +69045,8 @@ func (m *awsEc2query_deserializeOpModifySnapshotTier) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifySnapshotTier(response, &metadata)
@@ -70533,10 +69146,8 @@ func (m *awsEc2query_deserializeOpModifySpotFleetReque
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifySpotFleetRequest(response, &metadata)
@@ -70636,10 +69247,8 @@ func (m *awsEc2query_deserializeOpModifySubnetAttribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifySubnetAttribute(response, &metadata)
@@ -70716,10 +69325,8 @@ func (m *awsEc2query_deserializeOpModifyTrafficMirrorF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTrafficMirrorFilterNetworkServices(response, &metadata)
@@ -70819,10 +69426,8 @@ func (m *awsEc2query_deserializeOpModifyTrafficMirrorF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTrafficMirrorFilterRule(response, &metadata)
@@ -70922,10 +69527,8 @@ func (m *awsEc2query_deserializeOpModifyTrafficMirrorS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTrafficMirrorSession(response, &metadata)
@@ -71025,10 +69628,8 @@ func (m *awsEc2query_deserializeOpModifyTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTransitGateway(response, &metadata)
@@ -71128,10 +69729,8 @@ func (m *awsEc2query_deserializeOpModifyTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTransitGatewayAttachmentNatConfiguration(response, &metadata)
@@ -71231,10 +69830,8 @@ func (m *awsEc2query_deserializeOpModifyTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTransitGatewayMeteringPolicy(response, &metadata)
@@ -71334,10 +69931,8 @@ func (m *awsEc2query_deserializeOpModifyTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTransitGatewayPolicyTableEntry(response, &metadata)
@@ -71437,10 +70032,8 @@ func (m *awsEc2query_deserializeOpModifyTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTransitGatewayPrefixListReference(response, &metadata)
@@ -71540,10 +70133,8 @@ func (m *awsEc2query_deserializeOpModifyTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyTransitGatewayVpcAttachment(response, &metadata)
@@ -71643,10 +70234,8 @@ func (m *awsEc2query_deserializeOpModifyVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVerifiedAccessEndpoint(response, &metadata)
@@ -71746,10 +70335,8 @@ func (m *awsEc2query_deserializeOpModifyVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVerifiedAccessEndpointPolicy(response, &metadata)
@@ -71849,10 +70436,8 @@ func (m *awsEc2query_deserializeOpModifyVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVerifiedAccessGroup(response, &metadata)
@@ -71952,10 +70537,8 @@ func (m *awsEc2query_deserializeOpModifyVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVerifiedAccessGroupPolicy(response, &metadata)
@@ -72055,10 +70638,8 @@ func (m *awsEc2query_deserializeOpModifyVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVerifiedAccessInstance(response, &metadata)
@@ -72158,10 +70739,8 @@ func (m *awsEc2query_deserializeOpModifyVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVerifiedAccessInstanceLoggingConfiguration(response, &metadata)
@@ -72261,10 +70840,8 @@ func (m *awsEc2query_deserializeOpModifyVerifiedAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVerifiedAccessTrustProvider(response, &metadata)
@@ -72364,10 +70941,8 @@ func (m *awsEc2query_deserializeOpModifyVolume) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVolume(response, &metadata)
@@ -72467,10 +71042,8 @@ func (m *awsEc2query_deserializeOpModifyVolumeAttribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVolumeAttribute(response, &metadata)
@@ -72547,10 +71120,8 @@ func (m *awsEc2query_deserializeOpModifyVpcAttribute) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcAttribute(response, &metadata)
@@ -72627,10 +71198,8 @@ func (m *awsEc2query_deserializeOpModifyVpcBlockPublic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcBlockPublicAccessExclusion(response, &metadata)
@@ -72730,10 +71299,8 @@ func (m *awsEc2query_deserializeOpModifyVpcBlockPublic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcBlockPublicAccessOptions(response, &metadata)
@@ -72833,10 +71400,8 @@ func (m *awsEc2query_deserializeOpModifyVpcEncryptionC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcEncryptionControl(response, &metadata)
@@ -72936,10 +71501,8 @@ func (m *awsEc2query_deserializeOpModifyVpcEndpoint) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcEndpoint(response, &metadata)
@@ -73039,10 +71602,8 @@ func (m *awsEc2query_deserializeOpModifyVpcEndpointCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcEndpointConnectionNotification(response, &metadata)
@@ -73142,10 +71703,8 @@ func (m *awsEc2query_deserializeOpModifyVpcEndpointPay
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcEndpointPayerResponsibility(response, &metadata)
@@ -73245,10 +71804,8 @@ func (m *awsEc2query_deserializeOpModifyVpcEndpointSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcEndpointServiceConfiguration(response, &metadata)
@@ -73348,10 +71905,8 @@ func (m *awsEc2query_deserializeOpModifyVpcEndpointSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcEndpointServicePayerResponsibility(response, &metadata)
@@ -73451,10 +72006,8 @@ func (m *awsEc2query_deserializeOpModifyVpcEndpointSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcEndpointServicePermissions(response, &metadata)
@@ -73554,10 +72107,8 @@ func (m *awsEc2query_deserializeOpModifyVpcPeeringConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcPeeringConnectionOptions(response, &metadata)
@@ -73657,10 +72208,8 @@ func (m *awsEc2query_deserializeOpModifyVpcTenancy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpcTenancy(response, &metadata)
@@ -73760,10 +72309,8 @@ func (m *awsEc2query_deserializeOpModifyVpnConnection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpnConnection(response, &metadata)
@@ -73863,10 +72410,8 @@ func (m *awsEc2query_deserializeOpModifyVpnConnectionO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpnConnectionOptions(response, &metadata)
@@ -73966,10 +72511,8 @@ func (m *awsEc2query_deserializeOpModifyVpnTunnelCerti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpnTunnelCertificate(response, &metadata)
@@ -74069,10 +72612,8 @@ func (m *awsEc2query_deserializeOpModifyVpnTunnelOptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorModifyVpnTunnelOptions(response, &metadata)
@@ -74172,10 +72713,8 @@ func (m *awsEc2query_deserializeOpMonitorInstances) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorMonitorInstances(response, &metadata)
@@ -74275,10 +72814,8 @@ func (m *awsEc2query_deserializeOpMoveAddressToVpc) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorMoveAddressToVpc(response, &metadata)
@@ -74378,10 +72915,8 @@ func (m *awsEc2query_deserializeOpMoveByoipCidrToIpam)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorMoveByoipCidrToIpam(response, &metadata)
@@ -74481,10 +73016,8 @@ func (m *awsEc2query_deserializeOpMoveCapacityReservat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorMoveCapacityReservationInstances(response, &metadata)
@@ -74584,10 +73117,8 @@ func (m *awsEc2query_deserializeOpProvisionByoipCidr) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorProvisionByoipCidr(response, &metadata)
@@ -74687,10 +73218,8 @@ func (m *awsEc2query_deserializeOpProvisionIpamByoasn)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorProvisionIpamByoasn(response, &metadata)
@@ -74790,10 +73319,8 @@ func (m *awsEc2query_deserializeOpProvisionIpamPoolCid
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorProvisionIpamPoolCidr(response, &metadata)
@@ -74893,10 +73420,8 @@ func (m *awsEc2query_deserializeOpProvisionPublicIpv4P
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorProvisionPublicIpv4PoolCidr(response, &metadata)
@@ -74996,10 +73521,8 @@ func (m *awsEc2query_deserializeOpPurchaseCapacityBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorPurchaseCapacityBlock(response, &metadata)
@@ -75099,10 +73622,8 @@ func (m *awsEc2query_deserializeOpPurchaseCapacityBloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorPurchaseCapacityBlockExtension(response, &metadata)
@@ -75202,10 +73723,8 @@ func (m *awsEc2query_deserializeOpPurchaseHostReservat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorPurchaseHostReservation(response, &metadata)
@@ -75305,10 +73824,8 @@ func (m *awsEc2query_deserializeOpPurchaseReservedInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorPurchaseReservedInstancesOffering(response, &metadata)
@@ -75408,10 +73925,8 @@ func (m *awsEc2query_deserializeOpPurchaseScheduledIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorPurchaseScheduledInstances(response, &metadata)
@@ -75511,10 +74026,8 @@ func (m *awsEc2query_deserializeOpRebootInstances) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRebootInstances(response, &metadata)
@@ -75591,10 +74104,8 @@ func (m *awsEc2query_deserializeOpRegisterImage) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRegisterImage(response, &metadata)
@@ -75694,10 +74205,8 @@ func (m *awsEc2query_deserializeOpRegisterInstanceEven
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRegisterInstanceEventNotificationAttributes(response, &metadata)
@@ -75797,10 +74306,8 @@ func (m *awsEc2query_deserializeOpRegisterTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRegisterTransitGatewayMulticastGroupMembers(response, &metadata)
@@ -75900,10 +74407,8 @@ func (m *awsEc2query_deserializeOpRegisterTransitGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRegisterTransitGatewayMulticastGroupSources(response, &metadata)
@@ -76003,10 +74508,8 @@ func (m *awsEc2query_deserializeOpRejectCapacityReserv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRejectCapacityReservationBillingOwnership(response, &metadata)
@@ -76106,10 +74609,8 @@ func (m *awsEc2query_deserializeOpRejectTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRejectTransitGatewayClientVpnAttachment(response, &metadata)
@@ -76209,10 +74710,8 @@ func (m *awsEc2query_deserializeOpRejectTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRejectTransitGatewayMulticastDomainAssociations(response, &metadata)
@@ -76312,10 +74811,8 @@ func (m *awsEc2query_deserializeOpRejectTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRejectTransitGatewayPeeringAttachment(response, &metadata)
@@ -76415,10 +74912,8 @@ func (m *awsEc2query_deserializeOpRejectTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRejectTransitGatewayVpcAttachment(response, &metadata)
@@ -76518,10 +75013,8 @@ func (m *awsEc2query_deserializeOpRejectVpcEndpointCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRejectVpcEndpointConnections(response, &metadata)
@@ -76621,10 +75114,8 @@ func (m *awsEc2query_deserializeOpRejectVpcPeeringConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRejectVpcPeeringConnection(response, &metadata)
@@ -76724,10 +75215,8 @@ func (m *awsEc2query_deserializeOpReleaseAddress) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReleaseAddress(response, &metadata)
@@ -76804,10 +75293,8 @@ func (m *awsEc2query_deserializeOpReleaseHosts) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReleaseHosts(response, &metadata)
@@ -76907,10 +75394,8 @@ func (m *awsEc2query_deserializeOpReleaseIpamPoolAlloc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReleaseIpamPoolAllocation(response, &metadata)
@@ -77010,10 +75495,8 @@ func (m *awsEc2query_deserializeOpReplaceIamInstancePr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceIamInstanceProfileAssociation(response, &metadata)
@@ -77113,10 +75596,8 @@ func (m *awsEc2query_deserializeOpReplaceImageCriteria
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceImageCriteriaInAllowedImagesSettings(response, &metadata)
@@ -77216,10 +75697,8 @@ func (m *awsEc2query_deserializeOpReplaceImageInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceImageInstanceTypeSpecification(response, &metadata)
@@ -77319,10 +75798,8 @@ func (m *awsEc2query_deserializeOpReplaceNetworkAclAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceNetworkAclAssociation(response, &metadata)
@@ -77422,10 +75899,8 @@ func (m *awsEc2query_deserializeOpReplaceNetworkAclEnt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceNetworkAclEntry(response, &metadata)
@@ -77502,10 +75977,8 @@ func (m *awsEc2query_deserializeOpReplaceRoute) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceRoute(response, &metadata)
@@ -77582,10 +76055,8 @@ func (m *awsEc2query_deserializeOpReplaceRouteTableAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceRouteTableAssociation(response, &metadata)
@@ -77685,10 +76156,8 @@ func (m *awsEc2query_deserializeOpReplaceTransitGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceTransitGatewayRoute(response, &metadata)
@@ -77788,10 +76257,8 @@ func (m *awsEc2query_deserializeOpReplaceVpnTunnel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReplaceVpnTunnel(response, &metadata)
@@ -77891,10 +76358,8 @@ func (m *awsEc2query_deserializeOpReportInstanceStatus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorReportInstanceStatus(response, &metadata)
@@ -77971,10 +76436,8 @@ func (m *awsEc2query_deserializeOpRequestSpotFleet) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRequestSpotFleet(response, &metadata)
@@ -78074,10 +76537,8 @@ func (m *awsEc2query_deserializeOpRequestSpotInstances
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRequestSpotInstances(response, &metadata)
@@ -78177,10 +76638,8 @@ func (m *awsEc2query_deserializeOpResetAddressAttribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorResetAddressAttribute(response, &metadata)
@@ -78280,10 +76739,8 @@ func (m *awsEc2query_deserializeOpResetEbsDefaultKmsKe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorResetEbsDefaultKmsKeyId(response, &metadata)
@@ -78383,10 +76840,8 @@ func (m *awsEc2query_deserializeOpResetFpgaImageAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorResetFpgaImageAttribute(response, &metadata)
@@ -78486,10 +76941,8 @@ func (m *awsEc2query_deserializeOpResetImageAttribute)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorResetImageAttribute(response, &metadata)
@@ -78566,10 +77019,8 @@ func (m *awsEc2query_deserializeOpResetInstanceAttribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorResetInstanceAttribute(response, &metadata)
@@ -78646,10 +77097,8 @@ func (m *awsEc2query_deserializeOpResetNetworkInterfac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorResetNetworkInterfaceAttribute(response, &metadata)
@@ -78726,10 +77175,8 @@ func (m *awsEc2query_deserializeOpResetSnapshotAttribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorResetSnapshotAttribute(response, &metadata)
@@ -78806,10 +77253,8 @@ func (m *awsEc2query_deserializeOpRestoreAddressToClas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRestoreAddressToClassic(response, &metadata)
@@ -78909,10 +77354,8 @@ func (m *awsEc2query_deserializeOpRestoreImageFromRecy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRestoreImageFromRecycleBin(response, &metadata)
@@ -79012,10 +77455,8 @@ func (m *awsEc2query_deserializeOpRestoreManagedPrefix
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRestoreManagedPrefixListVersion(response, &metadata)
@@ -79115,10 +77556,8 @@ func (m *awsEc2query_deserializeOpRestoreSnapshotFromR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRestoreSnapshotFromRecycleBin(response, &metadata)
@@ -79218,10 +77657,8 @@ func (m *awsEc2query_deserializeOpRestoreSnapshotTier)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRestoreSnapshotTier(response, &metadata)
@@ -79321,10 +77758,8 @@ func (m *awsEc2query_deserializeOpRestoreVolumeFromRec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRestoreVolumeFromRecycleBin(response, &metadata)
@@ -79424,10 +77859,8 @@ func (m *awsEc2query_deserializeOpRevokeClientVpnIngre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRevokeClientVpnIngress(response, &metadata)
@@ -79527,10 +77960,8 @@ func (m *awsEc2query_deserializeOpRevokeSecurityGroupE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRevokeSecurityGroupEgress(response, &metadata)
@@ -79630,10 +78061,8 @@ func (m *awsEc2query_deserializeOpRevokeSecurityGroupI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRevokeSecurityGroupIngress(response, &metadata)
@@ -79733,10 +78162,8 @@ func (m *awsEc2query_deserializeOpRunInstances) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRunInstances(response, &metadata)
@@ -79836,10 +78263,8 @@ func (m *awsEc2query_deserializeOpRunScheduledInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorRunScheduledInstances(response, &metadata)
@@ -79939,10 +78364,8 @@ func (m *awsEc2query_deserializeOpSearchLocalGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorSearchLocalGatewayRoutes(response, &metadata)
@@ -80042,10 +78465,8 @@ func (m *awsEc2query_deserializeOpSearchTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorSearchTransitGatewayMulticastGroups(response, &metadata)
@@ -80145,10 +78566,8 @@ func (m *awsEc2query_deserializeOpSearchTransitGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorSearchTransitGatewayRoutes(response, &metadata)
@@ -80248,10 +78667,8 @@ func (m *awsEc2query_deserializeOpSendDiagnosticInterr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorSendDiagnosticInterrupt(response, &metadata)
@@ -80328,10 +78745,8 @@ func (m *awsEc2query_deserializeOpStartDeclarativePoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorStartDeclarativePoliciesReport(response, &metadata)
@@ -80431,10 +78846,8 @@ func (m *awsEc2query_deserializeOpStartInstances) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorStartInstances(response, &metadata)
@@ -80534,10 +78947,8 @@ func (m *awsEc2query_deserializeOpStartNetworkInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorStartNetworkInsightsAccessScopeAnalysis(response, &metadata)
@@ -80637,10 +79048,8 @@ func (m *awsEc2query_deserializeOpStartNetworkInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorStartNetworkInsightsAnalysis(response, &metadata)
@@ -80740,10 +79149,8 @@ func (m *awsEc2query_deserializeOpStartVpcEndpointServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorStartVpcEndpointServicePrivateDnsVerification(response, &metadata)
@@ -80843,10 +79250,8 @@ func (m *awsEc2query_deserializeOpStopInstances) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorStopInstances(response, &metadata)
@@ -80946,10 +79351,8 @@ func (m *awsEc2query_deserializeOpTerminateClientVpnCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorTerminateClientVpnConnections(response, &metadata)
@@ -81049,10 +79452,8 @@ func (m *awsEc2query_deserializeOpTerminateInstances) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorTerminateInstances(response, &metadata)
@@ -81152,10 +79553,8 @@ func (m *awsEc2query_deserializeOpUnassignIpv6Addresse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUnassignIpv6Addresses(response, &metadata)
@@ -81255,10 +79654,8 @@ func (m *awsEc2query_deserializeOpUnassignPrivateIpAdd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUnassignPrivateIpAddresses(response, &metadata)
@@ -81335,10 +79732,8 @@ func (m *awsEc2query_deserializeOpUnassignPrivateNatGa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUnassignPrivateNatGatewayAddress(response, &metadata)
@@ -81438,10 +79833,8 @@ func (m *awsEc2query_deserializeOpUnlockSnapshot) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUnlockSnapshot(response, &metadata)
@@ -81541,10 +79934,8 @@ func (m *awsEc2query_deserializeOpUnmonitorInstances) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUnmonitorInstances(response, &metadata)
@@ -81644,10 +80035,8 @@ func (m *awsEc2query_deserializeOpUpdateCapacityManage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUpdateCapacityManagerMonitoredTagKeys(response, &metadata)
@@ -81747,10 +80136,8 @@ func (m *awsEc2query_deserializeOpUpdateCapacityManage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUpdateCapacityManagerOrganizationsAccess(response, &metadata)
@@ -81850,10 +80237,8 @@ func (m *awsEc2query_deserializeOpUpdateInterruptibleC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUpdateInterruptibleCapacityReservationAllocation(response, &metadata)
@@ -81953,10 +80338,8 @@ func (m *awsEc2query_deserializeOpUpdateSecurityGroupR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUpdateSecurityGroupRuleDescriptionsEgress(response, &metadata)
@@ -82056,10 +80439,8 @@ func (m *awsEc2query_deserializeOpUpdateSecurityGroupR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorUpdateSecurityGroupRuleDescriptionsIngress(response, &metadata)
@@ -82159,10 +80540,8 @@ func (m *awsEc2query_deserializeOpValidateSecurityGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorValidateSecurityGroupQuotasForInterface(response, &metadata)
@@ -82262,10 +80641,8 @@ func (m *awsEc2query_deserializeOpWithdrawByoipCidr) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsEc2query_deserializeOpErrorWithdrawByoipCidr(response, &metadata)
