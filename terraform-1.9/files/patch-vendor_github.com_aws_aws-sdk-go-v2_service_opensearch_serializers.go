--- vendor/github.com/aws/aws-sdk-go-v2/service/opensearch/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/opensearch/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAcceptInboundConnecti
 func (m *awsRestjson1_serializeOpAcceptInboundConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_serializeOpAcceptInboundConnecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptInboundConnectionInput(v *AcceptInboundConnectionInput, encoder *httpbinding.Encoder) error {
@@ -97,10 +93,8 @@ func (m *awsRestjson1_serializeOpAddDataSource) Handle
 func (m *awsRestjson1_serializeOpAddDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -149,7 +143,6 @@ func (m *awsRestjson1_serializeOpAddDataSource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddDataSourceInput(v *AddDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -203,10 +196,8 @@ func (m *awsRestjson1_serializeOpAddDirectQueryDataSou
 func (m *awsRestjson1_serializeOpAddDirectQueryDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -251,7 +242,6 @@ func (m *awsRestjson1_serializeOpAddDirectQueryDataSou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddDirectQueryDataSourceInput(v *AddDirectQueryDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -310,10 +300,8 @@ func (m *awsRestjson1_serializeOpAddTags) HandleSerial
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
@@ -358,7 +346,6 @@ func (m *awsRestjson1_serializeOpAddTags) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddTagsInput(v *AddTagsInput, encoder *httpbinding.Encoder) error {
@@ -398,10 +385,8 @@ func (m *awsRestjson1_serializeOpAssociatePackage) Han
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
@@ -450,7 +435,6 @@ func (m *awsRestjson1_serializeOpAssociatePackage) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociatePackageInput(v *AssociatePackageInput, encoder *httpbinding.Encoder) error {
@@ -510,10 +494,8 @@ func (m *awsRestjson1_serializeOpAssociatePackages) Ha
 func (m *awsRestjson1_serializeOpAssociatePackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -558,7 +540,6 @@ func (m *awsRestjson1_serializeOpAssociatePackages) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociatePackagesInput(v *AssociatePackagesInput, encoder *httpbinding.Encoder) error {
@@ -598,10 +579,8 @@ func (m *awsRestjson1_serializeOpAuthorizeVpcEndpointA
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
@@ -650,7 +629,6 @@ func (m *awsRestjson1_serializeOpAuthorizeVpcEndpointA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAuthorizeVpcEndpointAccessInput(v *AuthorizeVpcEndpointAccessInput, encoder *httpbinding.Encoder) error {
@@ -697,10 +675,8 @@ func (m *awsRestjson1_serializeOpCancelDomainConfigCha
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
@@ -749,7 +725,6 @@ func (m *awsRestjson1_serializeOpCancelDomainConfigCha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelDomainConfigChangeInput(v *CancelDomainConfigChangeInput, encoder *httpbinding.Encoder) error {
@@ -791,10 +766,8 @@ func (m *awsRestjson1_serializeOpCancelServiceSoftware
 func (m *awsRestjson1_serializeOpCancelServiceSoftwareUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -839,7 +812,6 @@ func (m *awsRestjson1_serializeOpCancelServiceSoftware
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelServiceSoftwareUpdateInput(v *CancelServiceSoftwareUpdateInput, encoder *httpbinding.Encoder) error {
@@ -872,10 +844,8 @@ func (m *awsRestjson1_serializeOpCreateApplication) Ha
 func (m *awsRestjson1_serializeOpCreateApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -920,7 +890,6 @@ func (m *awsRestjson1_serializeOpCreateApplication) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateApplicationInput(v *CreateApplicationInput, encoder *httpbinding.Encoder) error {
@@ -986,10 +955,8 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
 func (m *awsRestjson1_serializeOpCreateDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1034,7 +1001,6 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainInput(v *CreateDomainInput, encoder *httpbinding.Encoder) error {
@@ -1201,10 +1167,8 @@ func (m *awsRestjson1_serializeOpCreateOutboundConnect
 func (m *awsRestjson1_serializeOpCreateOutboundConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1249,7 +1213,6 @@ func (m *awsRestjson1_serializeOpCreateOutboundConnect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateOutboundConnectionInput(v *CreateOutboundConnectionInput, encoder *httpbinding.Encoder) error {
@@ -1308,10 +1271,8 @@ func (m *awsRestjson1_serializeOpCreatePackage) Handle
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
@@ -1356,7 +1317,6 @@ func (m *awsRestjson1_serializeOpCreatePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePackageInput(v *CreatePackageInput, encoder *httpbinding.Encoder) error {
@@ -1432,10 +1392,8 @@ func (m *awsRestjson1_serializeOpCreateVpcEndpoint) Ha
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
@@ -1480,7 +1438,6 @@ func (m *awsRestjson1_serializeOpCreateVpcEndpoint) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVpcEndpointInput(v *CreateVpcEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1525,10 +1482,8 @@ func (m *awsRestjson1_serializeOpDeleteApplication) Ha
 func (m *awsRestjson1_serializeOpDeleteApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1566,7 +1521,6 @@ func (m *awsRestjson1_serializeOpDeleteApplication) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteApplicationInput(v *DeleteApplicationInput, encoder *httpbinding.Encoder) error {
@@ -1596,10 +1550,8 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 func (m *awsRestjson1_serializeOpDeleteDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1637,7 +1589,6 @@ func (m *awsRestjson1_serializeOpDeleteDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDataSourceInput(v *DeleteDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -1676,10 +1627,8 @@ func (m *awsRestjson1_serializeOpDeleteDirectQueryData
 func (m *awsRestjson1_serializeOpDeleteDirectQueryDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1717,7 +1666,6 @@ func (m *awsRestjson1_serializeOpDeleteDirectQueryData
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDirectQueryDataSourceInput(v *DeleteDirectQueryDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -1747,10 +1695,8 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
 func (m *awsRestjson1_serializeOpDeleteDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1788,7 +1734,6 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainInput(v *DeleteDomainInput, encoder *httpbinding.Encoder) error {
@@ -1818,10 +1763,8 @@ func (m *awsRestjson1_serializeOpDeleteInboundConnecti
 func (m *awsRestjson1_serializeOpDeleteInboundConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1859,7 +1802,6 @@ func (m *awsRestjson1_serializeOpDeleteInboundConnecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInboundConnectionInput(v *DeleteInboundConnectionInput, encoder *httpbinding.Encoder) error {
@@ -1889,10 +1831,8 @@ func (m *awsRestjson1_serializeOpDeleteOutboundConnect
 func (m *awsRestjson1_serializeOpDeleteOutboundConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1930,7 +1870,6 @@ func (m *awsRestjson1_serializeOpDeleteOutboundConnect
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteOutboundConnectionInput(v *DeleteOutboundConnectionInput, encoder *httpbinding.Encoder) error {
@@ -1960,10 +1899,8 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
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
@@ -2001,7 +1938,6 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePackageInput(v *DeletePackageInput, encoder *httpbinding.Encoder) error {
@@ -2031,10 +1967,8 @@ func (m *awsRestjson1_serializeOpDeleteVpcEndpoint) Ha
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
@@ -2072,7 +2006,6 @@ func (m *awsRestjson1_serializeOpDeleteVpcEndpoint) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVpcEndpointInput(v *DeleteVpcEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2102,10 +2035,8 @@ func (m *awsRestjson1_serializeOpDescribeDomain) Handl
 func (m *awsRestjson1_serializeOpDescribeDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2143,7 +2074,6 @@ func (m *awsRestjson1_serializeOpDescribeDomain) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainInput(v *DescribeDomainInput, encoder *httpbinding.Encoder) error {
@@ -2173,10 +2103,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainAutoTun
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
@@ -2225,7 +2153,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainAutoTun
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainAutoTunesInput(v *DescribeDomainAutoTunesInput, encoder *httpbinding.Encoder) error {
@@ -2272,10 +2199,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainChangeP
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
@@ -2313,7 +2238,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainChangeP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainChangeProgressInput(v *DescribeDomainChangeProgressInput, encoder *httpbinding.Encoder) error {
@@ -2347,10 +2271,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainConfig)
 func (m *awsRestjson1_serializeOpDescribeDomainConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2388,7 +2310,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainConfig)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainConfigInput(v *DescribeDomainConfigInput, encoder *httpbinding.Encoder) error {
@@ -2418,10 +2339,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainHealth)
 func (m *awsRestjson1_serializeOpDescribeDomainHealth) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2459,7 +2378,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainHealth)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainHealthInput(v *DescribeDomainHealthInput, encoder *httpbinding.Encoder) error {
@@ -2489,10 +2407,8 @@ func (m *awsRestjson1_serializeOpDescribeDomainNodes) 
 func (m *awsRestjson1_serializeOpDescribeDomainNodes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2530,7 +2446,6 @@ func (m *awsRestjson1_serializeOpDescribeDomainNodes) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainNodesInput(v *DescribeDomainNodesInput, encoder *httpbinding.Encoder) error {
@@ -2560,10 +2475,8 @@ func (m *awsRestjson1_serializeOpDescribeDomains) Hand
 func (m *awsRestjson1_serializeOpDescribeDomains) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2608,7 +2521,6 @@ func (m *awsRestjson1_serializeOpDescribeDomains) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainsInput(v *DescribeDomainsInput, encoder *httpbinding.Encoder) error {
@@ -2643,10 +2555,8 @@ func (m *awsRestjson1_serializeOpDescribeDryRunProgres
 func (m *awsRestjson1_serializeOpDescribeDryRunProgress) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2684,7 +2594,6 @@ func (m *awsRestjson1_serializeOpDescribeDryRunProgres
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDryRunProgressInput(v *DescribeDryRunProgressInput, encoder *httpbinding.Encoder) error {
@@ -2722,10 +2631,8 @@ func (m *awsRestjson1_serializeOpDescribeInboundConnec
 func (m *awsRestjson1_serializeOpDescribeInboundConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2770,7 +2677,6 @@ func (m *awsRestjson1_serializeOpDescribeInboundConnec
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInboundConnectionsInput(v *DescribeInboundConnectionsInput, encoder *httpbinding.Encoder) error {
@@ -2815,10 +2721,8 @@ func (m *awsRestjson1_serializeOpDescribeInstanceTypeL
 func (m *awsRestjson1_serializeOpDescribeInstanceTypeLimits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2856,7 +2760,6 @@ func (m *awsRestjson1_serializeOpDescribeInstanceTypeL
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeInstanceTypeLimitsInput(v *DescribeInstanceTypeLimitsInput, encoder *httpbinding.Encoder) error {
@@ -2899,10 +2802,8 @@ func (m *awsRestjson1_serializeOpDescribeOutboundConne
 func (m *awsRestjson1_serializeOpDescribeOutboundConnections) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2947,7 +2848,6 @@ func (m *awsRestjson1_serializeOpDescribeOutboundConne
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOutboundConnectionsInput(v *DescribeOutboundConnectionsInput, encoder *httpbinding.Encoder) error {
@@ -2992,10 +2892,8 @@ func (m *awsRestjson1_serializeOpDescribePackages) Han
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
@@ -3040,7 +2938,6 @@ func (m *awsRestjson1_serializeOpDescribePackages) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePackagesInput(v *DescribePackagesInput, encoder *httpbinding.Encoder) error {
@@ -3085,10 +2982,8 @@ func (m *awsRestjson1_serializeOpDescribeReservedInsta
 func (m *awsRestjson1_serializeOpDescribeReservedInstanceOfferings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3126,7 +3021,6 @@ func (m *awsRestjson1_serializeOpDescribeReservedInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReservedInstanceOfferingsInput(v *DescribeReservedInstanceOfferingsInput, encoder *httpbinding.Encoder) error {
@@ -3159,10 +3053,8 @@ func (m *awsRestjson1_serializeOpDescribeReservedInsta
 func (m *awsRestjson1_serializeOpDescribeReservedInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3200,7 +3092,6 @@ func (m *awsRestjson1_serializeOpDescribeReservedInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReservedInstancesInput(v *DescribeReservedInstancesInput, encoder *httpbinding.Encoder) error {
@@ -3233,10 +3124,8 @@ func (m *awsRestjson1_serializeOpDescribeVpcEndpoints)
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
@@ -3281,7 +3170,6 @@ func (m *awsRestjson1_serializeOpDescribeVpcEndpoints)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeVpcEndpointsInput(v *DescribeVpcEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -3316,10 +3204,8 @@ func (m *awsRestjson1_serializeOpDissociatePackage) Ha
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
@@ -3357,7 +3243,6 @@ func (m *awsRestjson1_serializeOpDissociatePackage) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDissociatePackageInput(v *DissociatePackageInput, encoder *httpbinding.Encoder) error {
@@ -3396,10 +3281,8 @@ func (m *awsRestjson1_serializeOpDissociatePackages) H
 func (m *awsRestjson1_serializeOpDissociatePackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3444,7 +3327,6 @@ func (m *awsRestjson1_serializeOpDissociatePackages) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDissociatePackagesInput(v *DissociatePackagesInput, encoder *httpbinding.Encoder) error {
@@ -3484,10 +3366,8 @@ func (m *awsRestjson1_serializeOpGetApplication) Handl
 func (m *awsRestjson1_serializeOpGetApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3525,7 +3405,6 @@ func (m *awsRestjson1_serializeOpGetApplication) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetApplicationInput(v *GetApplicationInput, encoder *httpbinding.Encoder) error {
@@ -3555,10 +3434,8 @@ func (m *awsRestjson1_serializeOpGetCompatibleVersions
 func (m *awsRestjson1_serializeOpGetCompatibleVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3596,7 +3473,6 @@ func (m *awsRestjson1_serializeOpGetCompatibleVersions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCompatibleVersionsInput(v *GetCompatibleVersionsInput, encoder *httpbinding.Encoder) error {
@@ -3621,10 +3497,8 @@ func (m *awsRestjson1_serializeOpGetDataSource) Handle
 func (m *awsRestjson1_serializeOpGetDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3662,7 +3536,6 @@ func (m *awsRestjson1_serializeOpGetDataSource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataSourceInput(v *GetDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -3701,10 +3574,8 @@ func (m *awsRestjson1_serializeOpGetDirectQueryDataSou
 func (m *awsRestjson1_serializeOpGetDirectQueryDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3742,7 +3613,6 @@ func (m *awsRestjson1_serializeOpGetDirectQueryDataSou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDirectQueryDataSourceInput(v *GetDirectQueryDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -3772,10 +3642,8 @@ func (m *awsRestjson1_serializeOpGetDomainMaintenanceS
 func (m *awsRestjson1_serializeOpGetDomainMaintenanceStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3813,7 +3681,6 @@ func (m *awsRestjson1_serializeOpGetDomainMaintenanceS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainMaintenanceStatusInput(v *GetDomainMaintenanceStatusInput, encoder *httpbinding.Encoder) error {
@@ -3847,10 +3714,8 @@ func (m *awsRestjson1_serializeOpGetPackageVersionHist
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
@@ -3888,7 +3753,6 @@ func (m *awsRestjson1_serializeOpGetPackageVersionHist
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPackageVersionHistoryInput(v *GetPackageVersionHistoryInput, encoder *httpbinding.Encoder) error {
@@ -3926,10 +3790,8 @@ func (m *awsRestjson1_serializeOpGetUpgradeHistory) Ha
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
@@ -3967,7 +3829,6 @@ func (m *awsRestjson1_serializeOpGetUpgradeHistory) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUpgradeHistoryInput(v *GetUpgradeHistoryInput, encoder *httpbinding.Encoder) error {
@@ -4005,10 +3866,8 @@ func (m *awsRestjson1_serializeOpGetUpgradeStatus) Han
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
@@ -4046,7 +3905,6 @@ func (m *awsRestjson1_serializeOpGetUpgradeStatus) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUpgradeStatusInput(v *GetUpgradeStatusInput, encoder *httpbinding.Encoder) error {
@@ -4076,10 +3934,8 @@ func (m *awsRestjson1_serializeOpListApplications) Han
 func (m *awsRestjson1_serializeOpListApplications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4117,7 +3973,6 @@ func (m *awsRestjson1_serializeOpListApplications) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListApplicationsInput(v *ListApplicationsInput, encoder *httpbinding.Encoder) error {
@@ -4152,10 +4007,8 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 func (m *awsRestjson1_serializeOpListDataSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4193,7 +4046,6 @@ func (m *awsRestjson1_serializeOpListDataSources) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataSourcesInput(v *ListDataSourcesInput, encoder *httpbinding.Encoder) error {
@@ -4223,10 +4075,8 @@ func (m *awsRestjson1_serializeOpListDirectQueryDataSo
 func (m *awsRestjson1_serializeOpListDirectQueryDataSources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4264,7 +4114,6 @@ func (m *awsRestjson1_serializeOpListDirectQueryDataSo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDirectQueryDataSourcesInput(v *ListDirectQueryDataSourcesInput, encoder *httpbinding.Encoder) error {
@@ -4289,10 +4138,8 @@ func (m *awsRestjson1_serializeOpListDomainMaintenance
 func (m *awsRestjson1_serializeOpListDomainMaintenances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4330,7 +4177,6 @@ func (m *awsRestjson1_serializeOpListDomainMaintenance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainMaintenancesInput(v *ListDomainMaintenancesInput, encoder *httpbinding.Encoder) error {
@@ -4376,10 +4222,8 @@ func (m *awsRestjson1_serializeOpListDomainNames) Hand
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
@@ -4417,7 +4261,6 @@ func (m *awsRestjson1_serializeOpListDomainNames) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainNamesInput(v *ListDomainNamesInput, encoder *httpbinding.Encoder) error {
@@ -4442,10 +4285,8 @@ func (m *awsRestjson1_serializeOpListDomainsForPackage
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
@@ -4483,7 +4324,6 @@ func (m *awsRestjson1_serializeOpListDomainsForPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainsForPackageInput(v *ListDomainsForPackageInput, encoder *httpbinding.Encoder) error {
@@ -4521,10 +4361,8 @@ func (m *awsRestjson1_serializeOpListInstanceTypeDetai
 func (m *awsRestjson1_serializeOpListInstanceTypeDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4562,7 +4400,6 @@ func (m *awsRestjson1_serializeOpListInstanceTypeDetai
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInstanceTypeDetailsInput(v *ListInstanceTypeDetailsInput, encoder *httpbinding.Encoder) error {
@@ -4612,10 +4449,8 @@ func (m *awsRestjson1_serializeOpListPackagesForDomain
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
@@ -4653,7 +4488,6 @@ func (m *awsRestjson1_serializeOpListPackagesForDomain
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackagesForDomainInput(v *ListPackagesForDomainInput, encoder *httpbinding.Encoder) error {
@@ -4691,10 +4525,8 @@ func (m *awsRestjson1_serializeOpListScheduledActions)
 func (m *awsRestjson1_serializeOpListScheduledActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4732,7 +4564,6 @@ func (m *awsRestjson1_serializeOpListScheduledActions)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListScheduledActionsInput(v *ListScheduledActionsInput, encoder *httpbinding.Encoder) error {
@@ -4770,10 +4601,8 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
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
@@ -4811,7 +4640,6 @@ func (m *awsRestjson1_serializeOpListTags) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsInput(v *ListTagsInput, encoder *httpbinding.Encoder) error {
@@ -4836,10 +4664,8 @@ func (m *awsRestjson1_serializeOpListVersions) HandleS
 func (m *awsRestjson1_serializeOpListVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4877,7 +4703,6 @@ func (m *awsRestjson1_serializeOpListVersions) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVersionsInput(v *ListVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4906,10 +4731,8 @@ func (m *awsRestjson1_serializeOpListVpcEndpointAccess
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
@@ -4947,7 +4770,6 @@ func (m *awsRestjson1_serializeOpListVpcEndpointAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVpcEndpointAccessInput(v *ListVpcEndpointAccessInput, encoder *httpbinding.Encoder) error {
@@ -4981,10 +4803,8 @@ func (m *awsRestjson1_serializeOpListVpcEndpoints) Han
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
@@ -5022,7 +4842,6 @@ func (m *awsRestjson1_serializeOpListVpcEndpoints) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVpcEndpointsInput(v *ListVpcEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -5047,10 +4866,8 @@ func (m *awsRestjson1_serializeOpListVpcEndpointsForDo
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
@@ -5088,7 +4905,6 @@ func (m *awsRestjson1_serializeOpListVpcEndpointsForDo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVpcEndpointsForDomainInput(v *ListVpcEndpointsForDomainInput, encoder *httpbinding.Encoder) error {
@@ -5122,10 +4938,8 @@ func (m *awsRestjson1_serializeOpPurchaseReservedInsta
 func (m *awsRestjson1_serializeOpPurchaseReservedInstanceOffering) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5170,7 +4984,6 @@ func (m *awsRestjson1_serializeOpPurchaseReservedInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPurchaseReservedInstanceOfferingInput(v *PurchaseReservedInstanceOfferingInput, encoder *httpbinding.Encoder) error {
@@ -5213,10 +5026,8 @@ func (m *awsRestjson1_serializeOpRejectInboundConnecti
 func (m *awsRestjson1_serializeOpRejectInboundConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5254,7 +5065,6 @@ func (m *awsRestjson1_serializeOpRejectInboundConnecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRejectInboundConnectionInput(v *RejectInboundConnectionInput, encoder *httpbinding.Encoder) error {
@@ -5284,10 +5094,8 @@ func (m *awsRestjson1_serializeOpRemoveTags) HandleSer
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
@@ -5332,7 +5140,6 @@ func (m *awsRestjson1_serializeOpRemoveTags) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveTagsInput(v *RemoveTagsInput, encoder *httpbinding.Encoder) error {
@@ -5372,10 +5179,8 @@ func (m *awsRestjson1_serializeOpRevokeVpcEndpointAcce
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
@@ -5424,7 +5229,6 @@ func (m *awsRestjson1_serializeOpRevokeVpcEndpointAcce
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRevokeVpcEndpointAccessInput(v *RevokeVpcEndpointAccessInput, encoder *httpbinding.Encoder) error {
@@ -5471,10 +5275,8 @@ func (m *awsRestjson1_serializeOpStartDomainMaintenanc
 func (m *awsRestjson1_serializeOpStartDomainMaintenance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5523,7 +5325,6 @@ func (m *awsRestjson1_serializeOpStartDomainMaintenanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartDomainMaintenanceInput(v *StartDomainMaintenanceInput, encoder *httpbinding.Encoder) error {
@@ -5570,10 +5371,8 @@ func (m *awsRestjson1_serializeOpStartServiceSoftwareU
 func (m *awsRestjson1_serializeOpStartServiceSoftwareUpdate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5618,7 +5417,6 @@ func (m *awsRestjson1_serializeOpStartServiceSoftwareU
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartServiceSoftwareUpdateInput(v *StartServiceSoftwareUpdateInput, encoder *httpbinding.Encoder) error {
@@ -5661,10 +5459,8 @@ func (m *awsRestjson1_serializeOpUpdateApplication) Ha
 func (m *awsRestjson1_serializeOpUpdateApplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5713,7 +5509,6 @@ func (m *awsRestjson1_serializeOpUpdateApplication) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateApplicationInput(v *UpdateApplicationInput, encoder *httpbinding.Encoder) error {
@@ -5764,10 +5559,8 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 func (m *awsRestjson1_serializeOpUpdateDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5816,7 +5609,6 @@ func (m *awsRestjson1_serializeOpUpdateDataSource) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDataSourceInput(v *UpdateDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -5879,10 +5671,8 @@ func (m *awsRestjson1_serializeOpUpdateDirectQueryData
 func (m *awsRestjson1_serializeOpUpdateDirectQueryDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5931,7 +5721,6 @@ func (m *awsRestjson1_serializeOpUpdateDirectQueryData
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDirectQueryDataSourceInput(v *UpdateDirectQueryDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -5987,10 +5776,8 @@ func (m *awsRestjson1_serializeOpUpdateDomainConfig) H
 func (m *awsRestjson1_serializeOpUpdateDomainConfig) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6039,7 +5826,6 @@ func (m *awsRestjson1_serializeOpUpdateDomainConfig) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDomainConfigInput(v *UpdateDomainConfigInput, encoder *httpbinding.Encoder) error {
@@ -6208,10 +5994,8 @@ func (m *awsRestjson1_serializeOpUpdatePackage) Handle
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
@@ -6256,7 +6040,6 @@ func (m *awsRestjson1_serializeOpUpdatePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageInput(v *UpdatePackageInput, encoder *httpbinding.Encoder) error {
@@ -6320,10 +6103,8 @@ func (m *awsRestjson1_serializeOpUpdatePackageScope) H
 func (m *awsRestjson1_serializeOpUpdatePackageScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6368,7 +6149,6 @@ func (m *awsRestjson1_serializeOpUpdatePackageScope) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageScopeInput(v *UpdatePackageScopeInput, encoder *httpbinding.Encoder) error {
@@ -6413,10 +6193,8 @@ func (m *awsRestjson1_serializeOpUpdateScheduledAction
 func (m *awsRestjson1_serializeOpUpdateScheduledAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6465,7 +6243,6 @@ func (m *awsRestjson1_serializeOpUpdateScheduledAction
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateScheduledActionInput(v *UpdateScheduledActionInput, encoder *httpbinding.Encoder) error {
@@ -6522,10 +6299,8 @@ func (m *awsRestjson1_serializeOpUpdateVpcEndpoint) Ha
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
@@ -6570,7 +6345,6 @@ func (m *awsRestjson1_serializeOpUpdateVpcEndpoint) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVpcEndpointInput(v *UpdateVpcEndpointInput, encoder *httpbinding.Encoder) error {
@@ -6610,10 +6384,8 @@ func (m *awsRestjson1_serializeOpUpgradeDomain) Handle
 func (m *awsRestjson1_serializeOpUpgradeDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6658,7 +6430,6 @@ func (m *awsRestjson1_serializeOpUpgradeDomain) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpgradeDomainInput(v *UpgradeDomainInput, encoder *httpbinding.Encoder) error {
