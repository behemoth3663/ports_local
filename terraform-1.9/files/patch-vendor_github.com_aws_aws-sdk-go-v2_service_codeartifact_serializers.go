--- vendor/github.com/aws/aws-sdk-go-v2/service/codeartifact/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/codeartifact/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociateExternalConn
 func (m *awsRestjson1_serializeOpAssociateExternalConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpAssociateExternalConn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateExternalConnectionInput(v *AssociateExternalConnectionInput, encoder *httpbinding.Encoder) error {
@@ -103,10 +99,8 @@ func (m *awsRestjson1_serializeOpCopyPackageVersions) 
 func (m *awsRestjson1_serializeOpCopyPackageVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -155,7 +149,6 @@ func (m *awsRestjson1_serializeOpCopyPackageVersions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCopyPackageVersionsInput(v *CopyPackageVersionsInput, encoder *httpbinding.Encoder) error {
@@ -235,10 +228,8 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
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
@@ -287,7 +278,6 @@ func (m *awsRestjson1_serializeOpCreateDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDomainInput(v *CreateDomainInput, encoder *httpbinding.Encoder) error {
@@ -331,10 +321,8 @@ func (m *awsRestjson1_serializeOpCreatePackageGroup) H
 func (m *awsRestjson1_serializeOpCreatePackageGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -383,7 +371,6 @@ func (m *awsRestjson1_serializeOpCreatePackageGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePackageGroupInput(v *CreatePackageGroupInput, encoder *httpbinding.Encoder) error {
@@ -441,10 +428,8 @@ func (m *awsRestjson1_serializeOpCreateRepository) Han
 func (m *awsRestjson1_serializeOpCreateRepository) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -493,7 +478,6 @@ func (m *awsRestjson1_serializeOpCreateRepository) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRepositoryInput(v *CreateRepositoryInput, encoder *httpbinding.Encoder) error {
@@ -552,10 +536,8 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
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
@@ -593,7 +575,6 @@ func (m *awsRestjson1_serializeOpDeleteDomain) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainInput(v *DeleteDomainInput, encoder *httpbinding.Encoder) error {
@@ -622,10 +603,8 @@ func (m *awsRestjson1_serializeOpDeleteDomainPermissio
 func (m *awsRestjson1_serializeOpDeleteDomainPermissionsPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -663,7 +642,6 @@ func (m *awsRestjson1_serializeOpDeleteDomainPermissio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDomainPermissionsPolicyInput(v *DeleteDomainPermissionsPolicyInput, encoder *httpbinding.Encoder) error {
@@ -696,10 +674,8 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
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
@@ -737,7 +713,6 @@ func (m *awsRestjson1_serializeOpDeletePackage) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePackageInput(v *DeletePackageInput, encoder *httpbinding.Encoder) error {
@@ -782,10 +757,8 @@ func (m *awsRestjson1_serializeOpDeletePackageGroup) H
 func (m *awsRestjson1_serializeOpDeletePackageGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -823,7 +796,6 @@ func (m *awsRestjson1_serializeOpDeletePackageGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePackageGroupInput(v *DeletePackageGroupInput, encoder *httpbinding.Encoder) error {
@@ -856,10 +828,8 @@ func (m *awsRestjson1_serializeOpDeletePackageVersions
 func (m *awsRestjson1_serializeOpDeletePackageVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -908,7 +878,6 @@ func (m *awsRestjson1_serializeOpDeletePackageVersions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePackageVersionsInput(v *DeletePackageVersionsInput, encoder *httpbinding.Encoder) error {
@@ -972,10 +941,8 @@ func (m *awsRestjson1_serializeOpDeleteRepository) Han
 func (m *awsRestjson1_serializeOpDeleteRepository) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1013,7 +980,6 @@ func (m *awsRestjson1_serializeOpDeleteRepository) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRepositoryInput(v *DeleteRepositoryInput, encoder *httpbinding.Encoder) error {
@@ -1046,10 +1012,8 @@ func (m *awsRestjson1_serializeOpDeleteRepositoryPermi
 func (m *awsRestjson1_serializeOpDeleteRepositoryPermissionsPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1087,7 +1051,6 @@ func (m *awsRestjson1_serializeOpDeleteRepositoryPermi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRepositoryPermissionsPolicyInput(v *DeleteRepositoryPermissionsPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1124,10 +1087,8 @@ func (m *awsRestjson1_serializeOpDescribeDomain) Handl
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
@@ -1165,7 +1126,6 @@ func (m *awsRestjson1_serializeOpDescribeDomain) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDomainInput(v *DescribeDomainInput, encoder *httpbinding.Encoder) error {
@@ -1194,10 +1154,8 @@ func (m *awsRestjson1_serializeOpDescribePackage) Hand
 func (m *awsRestjson1_serializeOpDescribePackage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1235,7 +1193,6 @@ func (m *awsRestjson1_serializeOpDescribePackage) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePackageInput(v *DescribePackageInput, encoder *httpbinding.Encoder) error {
@@ -1280,10 +1237,8 @@ func (m *awsRestjson1_serializeOpDescribePackageGroup)
 func (m *awsRestjson1_serializeOpDescribePackageGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1321,7 +1276,6 @@ func (m *awsRestjson1_serializeOpDescribePackageGroup)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePackageGroupInput(v *DescribePackageGroupInput, encoder *httpbinding.Encoder) error {
@@ -1354,10 +1308,8 @@ func (m *awsRestjson1_serializeOpDescribePackageVersio
 func (m *awsRestjson1_serializeOpDescribePackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1395,7 +1347,6 @@ func (m *awsRestjson1_serializeOpDescribePackageVersio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePackageVersionInput(v *DescribePackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -1444,10 +1395,8 @@ func (m *awsRestjson1_serializeOpDescribeRepository) H
 func (m *awsRestjson1_serializeOpDescribeRepository) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1485,7 +1434,6 @@ func (m *awsRestjson1_serializeOpDescribeRepository) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRepositoryInput(v *DescribeRepositoryInput, encoder *httpbinding.Encoder) error {
@@ -1518,10 +1466,8 @@ func (m *awsRestjson1_serializeOpDisassociateExternalC
 func (m *awsRestjson1_serializeOpDisassociateExternalConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1559,7 +1505,6 @@ func (m *awsRestjson1_serializeOpDisassociateExternalC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateExternalConnectionInput(v *DisassociateExternalConnectionInput, encoder *httpbinding.Encoder) error {
@@ -1596,10 +1541,8 @@ func (m *awsRestjson1_serializeOpDisposePackageVersion
 func (m *awsRestjson1_serializeOpDisposePackageVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1648,7 +1591,6 @@ func (m *awsRestjson1_serializeOpDisposePackageVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisposePackageVersionsInput(v *DisposePackageVersionsInput, encoder *httpbinding.Encoder) error {
@@ -1719,10 +1661,8 @@ func (m *awsRestjson1_serializeOpGetAssociatedPackageG
 func (m *awsRestjson1_serializeOpGetAssociatedPackageGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1760,7 +1700,6 @@ func (m *awsRestjson1_serializeOpGetAssociatedPackageG
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssociatedPackageGroupInput(v *GetAssociatedPackageGroupInput, encoder *httpbinding.Encoder) error {
@@ -1801,10 +1740,8 @@ func (m *awsRestjson1_serializeOpGetAuthorizationToken
 func (m *awsRestjson1_serializeOpGetAuthorizationToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1842,7 +1779,6 @@ func (m *awsRestjson1_serializeOpGetAuthorizationToken
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAuthorizationTokenInput(v *GetAuthorizationTokenInput, encoder *httpbinding.Encoder) error {
@@ -1875,10 +1811,8 @@ func (m *awsRestjson1_serializeOpGetDomainPermissionsP
 func (m *awsRestjson1_serializeOpGetDomainPermissionsPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1916,7 +1850,6 @@ func (m *awsRestjson1_serializeOpGetDomainPermissionsP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainPermissionsPolicyInput(v *GetDomainPermissionsPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1945,10 +1878,8 @@ func (m *awsRestjson1_serializeOpGetPackageVersionAsse
 func (m *awsRestjson1_serializeOpGetPackageVersionAsset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1986,7 +1917,6 @@ func (m *awsRestjson1_serializeOpGetPackageVersionAsse
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPackageVersionAssetInput(v *GetPackageVersionAssetInput, encoder *httpbinding.Encoder) error {
@@ -2043,10 +1973,8 @@ func (m *awsRestjson1_serializeOpGetPackageVersionRead
 func (m *awsRestjson1_serializeOpGetPackageVersionReadme) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2084,7 +2012,6 @@ func (m *awsRestjson1_serializeOpGetPackageVersionRead
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPackageVersionReadmeInput(v *GetPackageVersionReadmeInput, encoder *httpbinding.Encoder) error {
@@ -2133,10 +2060,8 @@ func (m *awsRestjson1_serializeOpGetRepositoryEndpoint
 func (m *awsRestjson1_serializeOpGetRepositoryEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2174,7 +2099,6 @@ func (m *awsRestjson1_serializeOpGetRepositoryEndpoint
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRepositoryEndpointInput(v *GetRepositoryEndpointInput, encoder *httpbinding.Encoder) error {
@@ -2215,10 +2139,8 @@ func (m *awsRestjson1_serializeOpGetRepositoryPermissi
 func (m *awsRestjson1_serializeOpGetRepositoryPermissionsPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2256,7 +2178,6 @@ func (m *awsRestjson1_serializeOpGetRepositoryPermissi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRepositoryPermissionsPolicyInput(v *GetRepositoryPermissionsPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2289,10 +2210,8 @@ func (m *awsRestjson1_serializeOpListAllowedRepositori
 func (m *awsRestjson1_serializeOpListAllowedRepositoriesForGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2330,7 +2249,6 @@ func (m *awsRestjson1_serializeOpListAllowedRepositori
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAllowedRepositoriesForGroupInput(v *ListAllowedRepositoriesForGroupInput, encoder *httpbinding.Encoder) error {
@@ -2375,10 +2293,8 @@ func (m *awsRestjson1_serializeOpListAssociatedPackage
 func (m *awsRestjson1_serializeOpListAssociatedPackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2416,7 +2332,6 @@ func (m *awsRestjson1_serializeOpListAssociatedPackage
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssociatedPackagesInput(v *ListAssociatedPackagesInput, encoder *httpbinding.Encoder) error {
@@ -2461,10 +2376,8 @@ func (m *awsRestjson1_serializeOpListDomains) HandleSe
 func (m *awsRestjson1_serializeOpListDomains) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2509,7 +2422,6 @@ func (m *awsRestjson1_serializeOpListDomains) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainsInput(v *ListDomainsInput, encoder *httpbinding.Encoder) error {
@@ -2547,10 +2459,8 @@ func (m *awsRestjson1_serializeOpListPackageGroups) Ha
 func (m *awsRestjson1_serializeOpListPackageGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2588,7 +2498,6 @@ func (m *awsRestjson1_serializeOpListPackageGroups) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackageGroupsInput(v *ListPackageGroupsInput, encoder *httpbinding.Encoder) error {
@@ -2629,10 +2538,8 @@ func (m *awsRestjson1_serializeOpListPackages) HandleS
 func (m *awsRestjson1_serializeOpListPackages) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2670,7 +2577,6 @@ func (m *awsRestjson1_serializeOpListPackages) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackagesInput(v *ListPackagesInput, encoder *httpbinding.Encoder) error {
@@ -2731,10 +2637,8 @@ func (m *awsRestjson1_serializeOpListPackageVersionAss
 func (m *awsRestjson1_serializeOpListPackageVersionAssets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2772,7 +2676,6 @@ func (m *awsRestjson1_serializeOpListPackageVersionAss
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackageVersionAssetsInput(v *ListPackageVersionAssetsInput, encoder *httpbinding.Encoder) error {
@@ -2829,10 +2732,8 @@ func (m *awsRestjson1_serializeOpListPackageVersionDep
 func (m *awsRestjson1_serializeOpListPackageVersionDependencies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2870,7 +2771,6 @@ func (m *awsRestjson1_serializeOpListPackageVersionDep
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackageVersionDependenciesInput(v *ListPackageVersionDependenciesInput, encoder *httpbinding.Encoder) error {
@@ -2923,10 +2823,8 @@ func (m *awsRestjson1_serializeOpListPackageVersions) 
 func (m *awsRestjson1_serializeOpListPackageVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2964,7 +2862,6 @@ func (m *awsRestjson1_serializeOpListPackageVersions) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPackageVersionsInput(v *ListPackageVersionsInput, encoder *httpbinding.Encoder) error {
@@ -3029,10 +2926,8 @@ func (m *awsRestjson1_serializeOpListRepositories) Han
 func (m *awsRestjson1_serializeOpListRepositories) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3070,7 +2965,6 @@ func (m *awsRestjson1_serializeOpListRepositories) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRepositoriesInput(v *ListRepositoriesInput, encoder *httpbinding.Encoder) error {
@@ -3103,10 +2997,8 @@ func (m *awsRestjson1_serializeOpListRepositoriesInDom
 func (m *awsRestjson1_serializeOpListRepositoriesInDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3144,7 +3036,6 @@ func (m *awsRestjson1_serializeOpListRepositoriesInDom
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRepositoriesInDomainInput(v *ListRepositoriesInDomainInput, encoder *httpbinding.Encoder) error {
@@ -3189,10 +3080,8 @@ func (m *awsRestjson1_serializeOpListSubPackageGroups)
 func (m *awsRestjson1_serializeOpListSubPackageGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3230,7 +3119,6 @@ func (m *awsRestjson1_serializeOpListSubPackageGroups)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSubPackageGroupsInput(v *ListSubPackageGroupsInput, encoder *httpbinding.Encoder) error {
@@ -3271,10 +3159,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -3312,7 +3198,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -3337,10 +3222,8 @@ func (m *awsRestjson1_serializeOpPublishPackageVersion
 func (m *awsRestjson1_serializeOpPublishPackageVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3390,7 +3273,6 @@ func (m *awsRestjson1_serializeOpPublishPackageVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPublishPackageVersionInput(v *PublishPackageVersionInput, encoder *httpbinding.Encoder) error {
@@ -3452,10 +3334,8 @@ func (m *awsRestjson1_serializeOpPutDomainPermissionsP
 func (m *awsRestjson1_serializeOpPutDomainPermissionsPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3500,7 +3380,6 @@ func (m *awsRestjson1_serializeOpPutDomainPermissionsP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDomainPermissionsPolicyInput(v *PutDomainPermissionsPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3548,10 +3427,8 @@ func (m *awsRestjson1_serializeOpPutPackageOriginConfi
 func (m *awsRestjson1_serializeOpPutPackageOriginConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3600,7 +3477,6 @@ func (m *awsRestjson1_serializeOpPutPackageOriginConfi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutPackageOriginConfigurationInput(v *PutPackageOriginConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3659,10 +3535,8 @@ func (m *awsRestjson1_serializeOpPutRepositoryPermissi
 func (m *awsRestjson1_serializeOpPutRepositoryPermissionsPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3711,7 +3585,6 @@ func (m *awsRestjson1_serializeOpPutRepositoryPermissi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutRepositoryPermissionsPolicyInput(v *PutRepositoryPermissionsPolicyInput, encoder *httpbinding.Encoder) error {
@@ -3761,10 +3634,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -3813,7 +3684,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3852,10 +3722,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -3904,7 +3772,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3943,10 +3810,8 @@ func (m *awsRestjson1_serializeOpUpdatePackageGroup) H
 func (m *awsRestjson1_serializeOpUpdatePackageGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3995,7 +3860,6 @@ func (m *awsRestjson1_serializeOpUpdatePackageGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageGroupInput(v *UpdatePackageGroupInput, encoder *httpbinding.Encoder) error {
@@ -4046,10 +3910,8 @@ func (m *awsRestjson1_serializeOpUpdatePackageGroupOri
 func (m *awsRestjson1_serializeOpUpdatePackageGroupOriginConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4098,7 +3960,6 @@ func (m *awsRestjson1_serializeOpUpdatePackageGroupOri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageGroupOriginConfigurationInput(v *UpdatePackageGroupOriginConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4159,10 +4020,8 @@ func (m *awsRestjson1_serializeOpUpdatePackageVersions
 func (m *awsRestjson1_serializeOpUpdatePackageVersionsStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4211,7 +4070,6 @@ func (m *awsRestjson1_serializeOpUpdatePackageVersions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePackageVersionsStatusInput(v *UpdatePackageVersionsStatusInput, encoder *httpbinding.Encoder) error {
@@ -4287,10 +4145,8 @@ func (m *awsRestjson1_serializeOpUpdateRepository) Han
 func (m *awsRestjson1_serializeOpUpdateRepository) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4339,7 +4195,6 @@ func (m *awsRestjson1_serializeOpUpdateRepository) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRepositoryInput(v *UpdateRepositoryInput, encoder *httpbinding.Encoder) error {
