--- vendor/github.com/aws/aws-sdk-go-v2/service/tnb/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/tnb/serializers.go
@@ -13,7 +13,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpCancelSolNetworkOpera
 func (m *awsRestjson1_serializeOpCancelSolNetworkOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -68,7 +65,6 @@ func (m *awsRestjson1_serializeOpCancelSolNetworkOpera
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelSolNetworkOperationInput(v *CancelSolNetworkOperationInput, encoder *httpbinding.Encoder) error {
@@ -98,10 +94,8 @@ func (m *awsRestjson1_serializeOpCreateSolFunctionPack
 func (m *awsRestjson1_serializeOpCreateSolFunctionPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -146,7 +140,6 @@ func (m *awsRestjson1_serializeOpCreateSolFunctionPack
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSolFunctionPackageInput(v *CreateSolFunctionPackageInput, encoder *httpbinding.Encoder) error {
@@ -181,10 +174,8 @@ func (m *awsRestjson1_serializeOpCreateSolNetworkInsta
 func (m *awsRestjson1_serializeOpCreateSolNetworkInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -229,7 +220,6 @@ func (m *awsRestjson1_serializeOpCreateSolNetworkInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSolNetworkInstanceInput(v *CreateSolNetworkInstanceInput, encoder *httpbinding.Encoder) error {
@@ -279,10 +269,8 @@ func (m *awsRestjson1_serializeOpCreateSolNetworkPacka
 func (m *awsRestjson1_serializeOpCreateSolNetworkPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -327,7 +315,6 @@ func (m *awsRestjson1_serializeOpCreateSolNetworkPacka
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSolNetworkPackageInput(v *CreateSolNetworkPackageInput, encoder *httpbinding.Encoder) error {
@@ -362,10 +349,8 @@ func (m *awsRestjson1_serializeOpDeleteSolFunctionPack
 func (m *awsRestjson1_serializeOpDeleteSolFunctionPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -403,7 +388,6 @@ func (m *awsRestjson1_serializeOpDeleteSolFunctionPack
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSolFunctionPackageInput(v *DeleteSolFunctionPackageInput, encoder *httpbinding.Encoder) error {
@@ -433,10 +417,8 @@ func (m *awsRestjson1_serializeOpDeleteSolNetworkInsta
 func (m *awsRestjson1_serializeOpDeleteSolNetworkInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -474,7 +456,6 @@ func (m *awsRestjson1_serializeOpDeleteSolNetworkInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSolNetworkInstanceInput(v *DeleteSolNetworkInstanceInput, encoder *httpbinding.Encoder) error {
@@ -504,10 +485,8 @@ func (m *awsRestjson1_serializeOpDeleteSolNetworkPacka
 func (m *awsRestjson1_serializeOpDeleteSolNetworkPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -545,7 +524,6 @@ func (m *awsRestjson1_serializeOpDeleteSolNetworkPacka
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSolNetworkPackageInput(v *DeleteSolNetworkPackageInput, encoder *httpbinding.Encoder) error {
@@ -575,10 +553,8 @@ func (m *awsRestjson1_serializeOpGetSolFunctionInstanc
 func (m *awsRestjson1_serializeOpGetSolFunctionInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -616,7 +592,6 @@ func (m *awsRestjson1_serializeOpGetSolFunctionInstanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolFunctionInstanceInput(v *GetSolFunctionInstanceInput, encoder *httpbinding.Encoder) error {
@@ -646,10 +621,8 @@ func (m *awsRestjson1_serializeOpGetSolFunctionPackage
 func (m *awsRestjson1_serializeOpGetSolFunctionPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -687,7 +660,6 @@ func (m *awsRestjson1_serializeOpGetSolFunctionPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolFunctionPackageInput(v *GetSolFunctionPackageInput, encoder *httpbinding.Encoder) error {
@@ -717,10 +689,8 @@ func (m *awsRestjson1_serializeOpGetSolFunctionPackage
 func (m *awsRestjson1_serializeOpGetSolFunctionPackageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -758,7 +728,6 @@ func (m *awsRestjson1_serializeOpGetSolFunctionPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolFunctionPackageContentInput(v *GetSolFunctionPackageContentInput, encoder *httpbinding.Encoder) error {
@@ -793,10 +762,8 @@ func (m *awsRestjson1_serializeOpGetSolFunctionPackage
 func (m *awsRestjson1_serializeOpGetSolFunctionPackageDescriptor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -834,7 +801,6 @@ func (m *awsRestjson1_serializeOpGetSolFunctionPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolFunctionPackageDescriptorInput(v *GetSolFunctionPackageDescriptorInput, encoder *httpbinding.Encoder) error {
@@ -869,10 +835,8 @@ func (m *awsRestjson1_serializeOpGetSolNetworkInstance
 func (m *awsRestjson1_serializeOpGetSolNetworkInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -910,7 +874,6 @@ func (m *awsRestjson1_serializeOpGetSolNetworkInstance
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolNetworkInstanceInput(v *GetSolNetworkInstanceInput, encoder *httpbinding.Encoder) error {
@@ -940,10 +903,8 @@ func (m *awsRestjson1_serializeOpGetSolNetworkOperatio
 func (m *awsRestjson1_serializeOpGetSolNetworkOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -981,7 +942,6 @@ func (m *awsRestjson1_serializeOpGetSolNetworkOperatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolNetworkOperationInput(v *GetSolNetworkOperationInput, encoder *httpbinding.Encoder) error {
@@ -1011,10 +971,8 @@ func (m *awsRestjson1_serializeOpGetSolNetworkPackage)
 func (m *awsRestjson1_serializeOpGetSolNetworkPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1052,7 +1010,6 @@ func (m *awsRestjson1_serializeOpGetSolNetworkPackage)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolNetworkPackageInput(v *GetSolNetworkPackageInput, encoder *httpbinding.Encoder) error {
@@ -1082,10 +1039,8 @@ func (m *awsRestjson1_serializeOpGetSolNetworkPackageC
 func (m *awsRestjson1_serializeOpGetSolNetworkPackageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1123,7 +1078,6 @@ func (m *awsRestjson1_serializeOpGetSolNetworkPackageC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolNetworkPackageContentInput(v *GetSolNetworkPackageContentInput, encoder *httpbinding.Encoder) error {
@@ -1158,10 +1112,8 @@ func (m *awsRestjson1_serializeOpGetSolNetworkPackageD
 func (m *awsRestjson1_serializeOpGetSolNetworkPackageDescriptor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1199,7 +1151,6 @@ func (m *awsRestjson1_serializeOpGetSolNetworkPackageD
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSolNetworkPackageDescriptorInput(v *GetSolNetworkPackageDescriptorInput, encoder *httpbinding.Encoder) error {
@@ -1229,10 +1180,8 @@ func (m *awsRestjson1_serializeOpInstantiateSolNetwork
 func (m *awsRestjson1_serializeOpInstantiateSolNetworkInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1281,7 +1230,6 @@ func (m *awsRestjson1_serializeOpInstantiateSolNetwork
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInstantiateSolNetworkInstanceInput(v *InstantiateSolNetworkInstanceInput, encoder *httpbinding.Encoder) error {
@@ -1336,10 +1284,8 @@ func (m *awsRestjson1_serializeOpListSolFunctionInstan
 func (m *awsRestjson1_serializeOpListSolFunctionInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1377,7 +1323,6 @@ func (m *awsRestjson1_serializeOpListSolFunctionInstan
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSolFunctionInstancesInput(v *ListSolFunctionInstancesInput, encoder *httpbinding.Encoder) error {
@@ -1406,10 +1351,8 @@ func (m *awsRestjson1_serializeOpListSolFunctionPackag
 func (m *awsRestjson1_serializeOpListSolFunctionPackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1447,7 +1390,6 @@ func (m *awsRestjson1_serializeOpListSolFunctionPackag
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSolFunctionPackagesInput(v *ListSolFunctionPackagesInput, encoder *httpbinding.Encoder) error {
@@ -1476,10 +1418,8 @@ func (m *awsRestjson1_serializeOpListSolNetworkInstanc
 func (m *awsRestjson1_serializeOpListSolNetworkInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1517,7 +1457,6 @@ func (m *awsRestjson1_serializeOpListSolNetworkInstanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSolNetworkInstancesInput(v *ListSolNetworkInstancesInput, encoder *httpbinding.Encoder) error {
@@ -1546,10 +1485,8 @@ func (m *awsRestjson1_serializeOpListSolNetworkOperati
 func (m *awsRestjson1_serializeOpListSolNetworkOperations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1587,7 +1524,6 @@ func (m *awsRestjson1_serializeOpListSolNetworkOperati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSolNetworkOperationsInput(v *ListSolNetworkOperationsInput, encoder *httpbinding.Encoder) error {
@@ -1620,10 +1556,8 @@ func (m *awsRestjson1_serializeOpListSolNetworkPackage
 func (m *awsRestjson1_serializeOpListSolNetworkPackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1661,7 +1595,6 @@ func (m *awsRestjson1_serializeOpListSolNetworkPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSolNetworkPackagesInput(v *ListSolNetworkPackagesInput, encoder *httpbinding.Encoder) error {
@@ -1690,10 +1623,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -1731,7 +1662,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -1761,10 +1691,8 @@ func (m *awsRestjson1_serializeOpPutSolFunctionPackage
 func (m *awsRestjson1_serializeOpPutSolFunctionPackageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1814,7 +1742,6 @@ func (m *awsRestjson1_serializeOpPutSolFunctionPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutSolFunctionPackageContentInput(v *PutSolFunctionPackageContentInput, encoder *httpbinding.Encoder) error {
@@ -1849,10 +1776,8 @@ func (m *awsRestjson1_serializeOpPutSolNetworkPackageC
 func (m *awsRestjson1_serializeOpPutSolNetworkPackageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1902,7 +1827,6 @@ func (m *awsRestjson1_serializeOpPutSolNetworkPackageC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutSolNetworkPackageContentInput(v *PutSolNetworkPackageContentInput, encoder *httpbinding.Encoder) error {
@@ -1937,10 +1861,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -1989,7 +1911,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2033,10 +1954,8 @@ func (m *awsRestjson1_serializeOpTerminateSolNetworkIn
 func (m *awsRestjson1_serializeOpTerminateSolNetworkInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2085,7 +2004,6 @@ func (m *awsRestjson1_serializeOpTerminateSolNetworkIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTerminateSolNetworkInstanceInput(v *TerminateSolNetworkInstanceInput, encoder *httpbinding.Encoder) error {
@@ -2129,10 +2047,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -2170,7 +2086,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -2206,10 +2121,8 @@ func (m *awsRestjson1_serializeOpUpdateSolFunctionPack
 func (m *awsRestjson1_serializeOpUpdateSolFunctionPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2258,7 +2171,6 @@ func (m *awsRestjson1_serializeOpUpdateSolFunctionPack
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSolFunctionPackageInput(v *UpdateSolFunctionPackageInput, encoder *httpbinding.Encoder) error {
@@ -2300,10 +2212,8 @@ func (m *awsRestjson1_serializeOpUpdateSolNetworkInsta
 func (m *awsRestjson1_serializeOpUpdateSolNetworkInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2352,7 +2262,6 @@ func (m *awsRestjson1_serializeOpUpdateSolNetworkInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSolNetworkInstanceInput(v *UpdateSolNetworkInstanceInput, encoder *httpbinding.Encoder) error {
@@ -2415,10 +2324,8 @@ func (m *awsRestjson1_serializeOpUpdateSolNetworkPacka
 func (m *awsRestjson1_serializeOpUpdateSolNetworkPackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2467,7 +2374,6 @@ func (m *awsRestjson1_serializeOpUpdateSolNetworkPacka
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSolNetworkPackageInput(v *UpdateSolNetworkPackageInput, encoder *httpbinding.Encoder) error {
@@ -2509,10 +2415,8 @@ func (m *awsRestjson1_serializeOpValidateSolFunctionPa
 func (m *awsRestjson1_serializeOpValidateSolFunctionPackageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2562,7 +2466,6 @@ func (m *awsRestjson1_serializeOpValidateSolFunctionPa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsValidateSolFunctionPackageContentInput(v *ValidateSolFunctionPackageContentInput, encoder *httpbinding.Encoder) error {
@@ -2597,10 +2500,8 @@ func (m *awsRestjson1_serializeOpValidateSolNetworkPac
 func (m *awsRestjson1_serializeOpValidateSolNetworkPackageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2650,7 +2551,6 @@ func (m *awsRestjson1_serializeOpValidateSolNetworkPac
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsValidateSolNetworkPackageContentInput(v *ValidateSolNetworkPackageContentInput, encoder *httpbinding.Encoder) error {
