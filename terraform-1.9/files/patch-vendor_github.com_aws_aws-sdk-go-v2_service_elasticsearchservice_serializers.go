--- vendor/github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAcceptInboundCrossClu
 func (m *awsRestjson1_serializeOpAcceptInboundCrossClusterSearchConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_serializeOpAcceptInboundCrossClu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptInboundCrossClusterSearchConnectionInput(v *AcceptInboundCrossClusterSearchConnectionInput, encoder *httpbinding.Encoder) error {
@@ -97,10 +93,8 @@ func (m *awsRestjson1_serializeOpAddTags) HandleSerial
 func (m *awsRestjson1_serializeOpAddTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -145,7 +139,6 @@ func (m *awsRestjson1_serializeOpAddTags) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddTagsInput(v *AddTagsInput, encoder *httpbinding.Encoder) error {
@@ -185,10 +178,8 @@ func (m *awsRestjson1_serializeOpAssociatePackage) Han
 func (m *awsRestjson1_serializeOpAssociatePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -226,7 +217,6 @@ func (m *awsRestjson1_serializeOpAssociatePackage) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociatePackageInput(v *AssociatePackageInput, encoder *httpbinding.Encoder) error {
@@ -265,10 +255,8 @@ func (m *awsRestjson1_serializeOpAuthorizeVpcEndpointA
 func (m *awsRestjson1_serializeOpAuthorizeVpcEndpointAccess) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -317,7 +305,6 @@ func (m *awsRestjson1_serializeOpAuthorizeVpcEndpointA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAuthorizeVpcEndpointAccessInput(v *AuthorizeVpcEndpointAccessInput, encoder *httpbinding.Encoder) error {
@@ -359,10 +346,8 @@ func (m *awsRestjson1_serializeOpCancelDomainConfigCha
 func (m *awsRestjson1_serializeOpCancelDomainConfigChange) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -411,7 +396,6 @@ func (m *awsRestjson1_serializeOpCancelDomainConfigCha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelDomainConfigChangeInput(v *CancelDomainConfigChangeInput, encoder *httpbinding.Encoder) error {
@@ -453,10 +437,8 @@ func (m *awsRestjson1_serializeOpCancelElasticsearchSe
 func (m *awsRestjson1_serializeOpCancelElasticsearchServiceSoftwareUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -501,7 +483,6 @@ func (m *awsRestjson1_serializeOpCancelElasticsearchSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelElasticsearchServiceSoftwareUpdateInput(v *CancelElasticsearchServiceSoftwareUpdateInput, encoder *httpbinding.Encoder) error {
@@ -534,10 +515,8 @@ func (m *awsRestjson1_serializeOpCreateElasticsearchDo
 func (m *awsRestjson1_serializeOpCreateElasticsearchDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -582,7 +561,6 @@ func (m *awsRestjson1_serializeOpCreateElasticsearchDo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateElasticsearchDomainInput(v *CreateElasticsearchDomainInput, encoder *httpbinding.Encoder) error {
@@ -716,10 +694,8 @@ func (m *awsRestjson1_serializeOpCreateOutboundCrossCl
 func (m *awsRestjson1_serializeOpCreateOutboundCrossClusterSearchConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -764,7 +740,6 @@ func (m *awsRestjson1_serializeOpCreateOutboundCrossCl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOutboundCrossClusterSearchConnectionInput(v *CreateOutboundCrossClusterSearchConnectionInput, encoder *httpbinding.Encoder) error {
@@ -811,10 +786,8 @@ func (m *awsRestjson1_serializeOpCreatePackage) Handle
 func (m *awsRestjson1_serializeOpCreatePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -859,7 +832,6 @@ func (m *awsRestjson1_serializeOpCreatePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePackageInput(v *CreatePackageInput, encoder *httpbinding.Encoder) error {
@@ -909,10 +881,8 @@ func (m *awsRestjson1_serializeOpCreateVpcEndpoint) Ha
 func (m *awsRestjson1_serializeOpCreateVpcEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -957,7 +927,6 @@ func (m *awsRestjson1_serializeOpCreateVpcEndpoint) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVpcEndpointInput(v *CreateVpcEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1002,10 +971,8 @@ func (m *awsRestjson1_serializeOpDeleteElasticsearchDo
 func (m *awsRestjson1_serializeOpDeleteElasticsearchDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1043,7 +1010,6 @@ func (m *awsRestjson1_serializeOpDeleteElasticsearchDo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteElasticsearchDomainInput(v *DeleteElasticsearchDomainInput, encoder *httpbinding.Encoder) error {
@@ -1073,10 +1039,8 @@ func (m *awsRestjson1_serializeOpDeleteElasticsearchSe
 func (m *awsRestjson1_serializeOpDeleteElasticsearchServiceRole) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1110,7 +1074,6 @@ func (m *awsRestjson1_serializeOpDeleteElasticsearchSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteElasticsearchServiceRoleInput(v *DeleteElasticsearchServiceRoleInput, encoder *httpbinding.Encoder) error {
@@ -1131,10 +1094,8 @@ func (m *awsRestjson1_serializeOpDeleteInboundCrossClu
 func (m *awsRestjson1_serializeOpDeleteInboundCrossClusterSearchConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1172,7 +1133,6 @@ func (m *awsRestjson1_serializeOpDeleteInboundCrossClu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInboundCrossClusterSearchConnectionInput(v *DeleteInboundCrossClusterSearchConnectionInput, encoder *httpbinding.Encoder) error {
@@ -1202,10 +1162,8 @@ func (m *awsRestjson1_serializeOpDeleteOutboundCrossCl
 func (m *awsRestjson1_serializeOpDeleteOutboundCrossClusterSearchConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1243,7 +1201,6 @@ func (m *awsRestjson1_serializeOpDeleteOutboundCrossCl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteOutboundCrossClusterSearchConnectionInput(v *DeleteOutboundCrossClusterSearchConnectionInput, encoder *httpbinding.Encoder) error {
@@ -1273,10 +1230,8 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
 func (m *awsRestjson1_serializeOpDeletePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1314,7 +1269,6 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePackageInput(v *DeletePackageInput, encoder *httpbinding.Encoder) error {
@@ -1344,10 +1298,8 @@ func (m *awsRestjson1_serializeOpDeleteVpcEndpoint) Ha
 func (m *awsRestjson1_serializeOpDeleteVpcEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1385,7 +1337,6 @@ func (m *awsRestjson1_serializeOpDeleteVpcEndpoint) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVpcEndpointInput(v *DeleteVpcEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1415,10 +1366,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainAutoTun
 func (m *awsRestjson1_serializeOpDescribeDomainAutoTunes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1467,7 +1416,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainAutoTun
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainAutoTunesInput(v *DescribeDomainAutoTunesInput, encoder *httpbinding.Encoder) error {
@@ -1514,10 +1462,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainChangeP
 func (m *awsRestjson1_serializeOpDescribeDomainChangeProgress) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1555,7 +1501,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainChangeP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainChangeProgressInput(v *DescribeDomainChangeProgressInput, encoder *httpbinding.Encoder) error {
@@ -1589,10 +1534,8 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 func (m *awsRestjson1_serializeOpDescribeElasticsearchDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1630,7 +1573,6 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeElasticsearchDomainInput(v *DescribeElasticsearchDomainInput, encoder *httpbinding.Encoder) error {
@@ -1660,10 +1602,8 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 func (m *awsRestjson1_serializeOpDescribeElasticsearchDomainConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1701,7 +1641,6 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeElasticsearchDomainConfigInput(v *DescribeElasticsearchDomainConfigInput, encoder *httpbinding.Encoder) error {
@@ -1731,10 +1670,8 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 func (m *awsRestjson1_serializeOpDescribeElasticsearchDomains) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1779,7 +1716,6 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeElasticsearchDomainsInput(v *DescribeElasticsearchDomainsInput, encoder *httpbinding.Encoder) error {
@@ -1814,10 +1750,8 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 func (m *awsRestjson1_serializeOpDescribeElasticsearchInstanceTypeLimits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1855,7 +1789,6 @@ func (m *awsRestjson1_serializeOpDescribeElasticsearch
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeElasticsearchInstanceTypeLimitsInput(v *DescribeElasticsearchInstanceTypeLimitsInput, encoder *httpbinding.Encoder) error {
@@ -1898,10 +1831,8 @@ func (m *awsRestjson1_serializeOpDescribeInboundCrossC
 func (m *awsRestjson1_serializeOpDescribeInboundCrossClusterSearchConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1946,7 +1877,6 @@ func (m *awsRestjson1_serializeOpDescribeInboundCrossC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInboundCrossClusterSearchConnectionsInput(v *DescribeInboundCrossClusterSearchConnectionsInput, encoder *httpbinding.Encoder) error {
@@ -1991,10 +1921,8 @@ func (m *awsRestjson1_serializeOpDescribeOutboundCross
 func (m *awsRestjson1_serializeOpDescribeOutboundCrossClusterSearchConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2039,7 +1967,6 @@ func (m *awsRestjson1_serializeOpDescribeOutboundCross
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOutboundCrossClusterSearchConnectionsInput(v *DescribeOutboundCrossClusterSearchConnectionsInput, encoder *httpbinding.Encoder) error {
@@ -2084,10 +2011,8 @@ func (m *awsRestjson1_serializeOpDescribePackages) Han
 func (m *awsRestjson1_serializeOpDescribePackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2132,7 +2057,6 @@ func (m *awsRestjson1_serializeOpDescribePackages) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePackagesInput(v *DescribePackagesInput, encoder *httpbinding.Encoder) error {
@@ -2177,10 +2101,8 @@ func (m *awsRestjson1_serializeOpDescribeReservedElast
 func (m *awsRestjson1_serializeOpDescribeReservedElasticsearchInstanceOfferings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2218,7 +2140,6 @@ func (m *awsRestjson1_serializeOpDescribeReservedElast
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReservedElasticsearchInstanceOfferingsInput(v *DescribeReservedElasticsearchInstanceOfferingsInput, encoder *httpbinding.Encoder) error {
@@ -2251,10 +2172,8 @@ func (m *awsRestjson1_serializeOpDescribeReservedElast
 func (m *awsRestjson1_serializeOpDescribeReservedElasticsearchInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2292,7 +2211,6 @@ func (m *awsRestjson1_serializeOpDescribeReservedElast
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReservedElasticsearchInstancesInput(v *DescribeReservedElasticsearchInstancesInput, encoder *httpbinding.Encoder) error {
@@ -2325,10 +2243,8 @@ func (m *awsRestjson1_serializeOpDescribeVpcEndpoints)
 func (m *awsRestjson1_serializeOpDescribeVpcEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2373,7 +2289,6 @@ func (m *awsRestjson1_serializeOpDescribeVpcEndpoints)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeVpcEndpointsInput(v *DescribeVpcEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -2408,10 +2323,8 @@ func (m *awsRestjson1_serializeOpDissociatePackage) Ha
 func (m *awsRestjson1_serializeOpDissociatePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2449,7 +2362,6 @@ func (m *awsRestjson1_serializeOpDissociatePackage) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDissociatePackageInput(v *DissociatePackageInput, encoder *httpbinding.Encoder) error {
@@ -2488,10 +2400,8 @@ func (m *awsRestjson1_serializeOpGetCompatibleElastics
 func (m *awsRestjson1_serializeOpGetCompatibleElasticsearchVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2529,7 +2439,6 @@ func (m *awsRestjson1_serializeOpGetCompatibleElastics
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCompatibleElasticsearchVersionsInput(v *GetCompatibleElasticsearchVersionsInput, encoder *httpbinding.Encoder) error {
@@ -2554,10 +2463,8 @@ func (m *awsRestjson1_serializeOpGetPackageVersionHist
 func (m *awsRestjson1_serializeOpGetPackageVersionHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2595,7 +2502,6 @@ func (m *awsRestjson1_serializeOpGetPackageVersionHist
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPackageVersionHistoryInput(v *GetPackageVersionHistoryInput, encoder *httpbinding.Encoder) error {
@@ -2633,10 +2539,8 @@ func (m *awsRestjson1_serializeOpGetUpgradeHistory) Ha
 func (m *awsRestjson1_serializeOpGetUpgradeHistory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2674,7 +2578,6 @@ func (m *awsRestjson1_serializeOpGetUpgradeHistory) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUpgradeHistoryInput(v *GetUpgradeHistoryInput, encoder *httpbinding.Encoder) error {
@@ -2712,10 +2615,8 @@ func (m *awsRestjson1_serializeOpGetUpgradeStatus) Han
 func (m *awsRestjson1_serializeOpGetUpgradeStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2753,7 +2654,6 @@ func (m *awsRestjson1_serializeOpGetUpgradeStatus) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUpgradeStatusInput(v *GetUpgradeStatusInput, encoder *httpbinding.Encoder) error {
@@ -2783,10 +2683,8 @@ func (m *awsRestjson1_serializeOpListDomainNames) Hand
 func (m *awsRestjson1_serializeOpListDomainNames) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2824,7 +2722,6 @@ func (m *awsRestjson1_serializeOpListDomainNames) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainNamesInput(v *ListDomainNamesInput, encoder *httpbinding.Encoder) error {
@@ -2849,10 +2746,8 @@ func (m *awsRestjson1_serializeOpListDomainsForPackage
 func (m *awsRestjson1_serializeOpListDomainsForPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2890,7 +2785,6 @@ func (m *awsRestjson1_serializeOpListDomainsForPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainsForPackageInput(v *ListDomainsForPackageInput, encoder *httpbinding.Encoder) error {
@@ -2928,10 +2822,8 @@ func (m *awsRestjson1_serializeOpListElasticsearchInst
 func (m *awsRestjson1_serializeOpListElasticsearchInstanceTypes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2969,7 +2861,6 @@ func (m *awsRestjson1_serializeOpListElasticsearchInst
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListElasticsearchInstanceTypesInput(v *ListElasticsearchInstanceTypesInput, encoder *httpbinding.Encoder) error {
@@ -3011,10 +2902,8 @@ func (m *awsRestjson1_serializeOpListElasticsearchVers
 func (m *awsRestjson1_serializeOpListElasticsearchVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3052,7 +2941,6 @@ func (m *awsRestjson1_serializeOpListElasticsearchVers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListElasticsearchVersionsInput(v *ListElasticsearchVersionsInput, encoder *httpbinding.Encoder) error {
@@ -3081,10 +2969,8 @@ func (m *awsRestjson1_serializeOpListPackagesForDomain
 func (m *awsRestjson1_serializeOpListPackagesForDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3122,7 +3008,6 @@ func (m *awsRestjson1_serializeOpListPackagesForDomain
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackagesForDomainInput(v *ListPackagesForDomainInput, encoder *httpbinding.Encoder) error {
@@ -3160,10 +3045,8 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
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
@@ -3201,7 +3084,6 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsInput(v *ListTagsInput, encoder *httpbinding.Encoder) error {
@@ -3226,10 +3108,8 @@ func (m *awsRestjson1_serializeOpListVpcEndpointAccess
 func (m *awsRestjson1_serializeOpListVpcEndpointAccess) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3267,7 +3147,6 @@ func (m *awsRestjson1_serializeOpListVpcEndpointAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVpcEndpointAccessInput(v *ListVpcEndpointAccessInput, encoder *httpbinding.Encoder) error {
@@ -3301,10 +3180,8 @@ func (m *awsRestjson1_serializeOpListVpcEndpoints) Han
 func (m *awsRestjson1_serializeOpListVpcEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3342,7 +3219,6 @@ func (m *awsRestjson1_serializeOpListVpcEndpoints) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVpcEndpointsInput(v *ListVpcEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -3367,10 +3243,8 @@ func (m *awsRestjson1_serializeOpListVpcEndpointsForDo
 func (m *awsRestjson1_serializeOpListVpcEndpointsForDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3408,7 +3282,6 @@ func (m *awsRestjson1_serializeOpListVpcEndpointsForDo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVpcEndpointsForDomainInput(v *ListVpcEndpointsForDomainInput, encoder *httpbinding.Encoder) error {
@@ -3442,10 +3315,8 @@ func (m *awsRestjson1_serializeOpPurchaseReservedElast
 func (m *awsRestjson1_serializeOpPurchaseReservedElasticsearchInstanceOffering) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3490,7 +3361,6 @@ func (m *awsRestjson1_serializeOpPurchaseReservedElast
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPurchaseReservedElasticsearchInstanceOfferingInput(v *PurchaseReservedElasticsearchInstanceOfferingInput, encoder *httpbinding.Encoder) error {
@@ -3533,10 +3403,8 @@ func (m *awsRestjson1_serializeOpRejectInboundCrossClu
 func (m *awsRestjson1_serializeOpRejectInboundCrossClusterSearchConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3574,7 +3442,6 @@ func (m *awsRestjson1_serializeOpRejectInboundCrossClu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectInboundCrossClusterSearchConnectionInput(v *RejectInboundCrossClusterSearchConnectionInput, encoder *httpbinding.Encoder) error {
@@ -3604,10 +3471,8 @@ func (m *awsRestjson1_serializeOpRemoveTags) HandleSer
 func (m *awsRestjson1_serializeOpRemoveTags) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3652,7 +3517,6 @@ func (m *awsRestjson1_serializeOpRemoveTags) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveTagsInput(v *RemoveTagsInput, encoder *httpbinding.Encoder) error {
@@ -3692,10 +3556,8 @@ func (m *awsRestjson1_serializeOpRevokeVpcEndpointAcce
 func (m *awsRestjson1_serializeOpRevokeVpcEndpointAccess) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3744,7 +3606,6 @@ func (m *awsRestjson1_serializeOpRevokeVpcEndpointAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRevokeVpcEndpointAccessInput(v *RevokeVpcEndpointAccessInput, encoder *httpbinding.Encoder) error {
@@ -3786,10 +3647,8 @@ func (m *awsRestjson1_serializeOpStartElasticsearchSer
 func (m *awsRestjson1_serializeOpStartElasticsearchServiceSoftwareUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3834,7 +3693,6 @@ func (m *awsRestjson1_serializeOpStartElasticsearchSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartElasticsearchServiceSoftwareUpdateInput(v *StartElasticsearchServiceSoftwareUpdateInput, encoder *httpbinding.Encoder) error {
@@ -3867,10 +3725,8 @@ func (m *awsRestjson1_serializeOpUpdateElasticsearchDo
 func (m *awsRestjson1_serializeOpUpdateElasticsearchDomainConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3919,7 +3775,6 @@ func (m *awsRestjson1_serializeOpUpdateElasticsearchDo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateElasticsearchDomainConfigInput(v *UpdateElasticsearchDomainConfigInput, encoder *httpbinding.Encoder) error {
@@ -4050,10 +3905,8 @@ func (m *awsRestjson1_serializeOpUpdatePackage) Handle
 func (m *awsRestjson1_serializeOpUpdatePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4098,7 +3951,6 @@ func (m *awsRestjson1_serializeOpUpdatePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageInput(v *UpdatePackageInput, encoder *httpbinding.Encoder) error {
@@ -4148,10 +4000,8 @@ func (m *awsRestjson1_serializeOpUpdateVpcEndpoint) Ha
 func (m *awsRestjson1_serializeOpUpdateVpcEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4196,7 +4046,6 @@ func (m *awsRestjson1_serializeOpUpdateVpcEndpoint) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVpcEndpointInput(v *UpdateVpcEndpointInput, encoder *httpbinding.Encoder) error {
@@ -4236,10 +4085,8 @@ func (m *awsRestjson1_serializeOpUpgradeElasticsearchD
 func (m *awsRestjson1_serializeOpUpgradeElasticsearchDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4284,7 +4131,6 @@ func (m *awsRestjson1_serializeOpUpgradeElasticsearchD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpgradeElasticsearchDomainInput(v *UpgradeElasticsearchDomainInput, encoder *httpbinding.Encoder) error {
