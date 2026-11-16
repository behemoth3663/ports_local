--- vendor/github.com/aws/aws-sdk-go-v2/service/finspacedata/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/finspacedata/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociateUserToPermis
 func (m *awsRestjson1_serializeOpAssociateUserToPermissionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpAssociateUserToPermis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateUserToPermissionGroupInput(v *AssociateUserToPermissionGroupInput, encoder *httpbinding.Encoder) error {
@@ -128,10 +124,8 @@ func (m *awsRestjson1_serializeOpCreateChangeset) Hand
 func (m *awsRestjson1_serializeOpCreateChangeset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -180,7 +174,6 @@ func (m *awsRestjson1_serializeOpCreateChangeset) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateChangesetInput(v *CreateChangesetInput, encoder *httpbinding.Encoder) error {
@@ -241,10 +234,8 @@ func (m *awsRestjson1_serializeOpCreateDataset) Handle
 func (m *awsRestjson1_serializeOpCreateDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -289,7 +280,6 @@ func (m *awsRestjson1_serializeOpCreateDataset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDatasetInput(v *CreateDatasetInput, encoder *httpbinding.Encoder) error {
@@ -363,10 +353,8 @@ func (m *awsRestjson1_serializeOpCreateDataView) Handl
 func (m *awsRestjson1_serializeOpCreateDataView) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -415,7 +403,6 @@ func (m *awsRestjson1_serializeOpCreateDataView) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDataViewInput(v *CreateDataViewInput, encoder *httpbinding.Encoder) error {
@@ -488,10 +475,8 @@ func (m *awsRestjson1_serializeOpCreatePermissionGroup
 func (m *awsRestjson1_serializeOpCreatePermissionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -536,7 +521,6 @@ func (m *awsRestjson1_serializeOpCreatePermissionGroup
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePermissionGroupInput(v *CreatePermissionGroupInput, encoder *httpbinding.Encoder) error {
@@ -586,10 +570,8 @@ func (m *awsRestjson1_serializeOpCreateUser) HandleSer
 func (m *awsRestjson1_serializeOpCreateUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -634,7 +616,6 @@ func (m *awsRestjson1_serializeOpCreateUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUserInput(v *CreateUserInput, encoder *httpbinding.Encoder) error {
@@ -697,10 +678,8 @@ func (m *awsRestjson1_serializeOpDeleteDataset) Handle
 func (m *awsRestjson1_serializeOpDeleteDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -738,7 +717,6 @@ func (m *awsRestjson1_serializeOpDeleteDataset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDatasetInput(v *DeleteDatasetInput, encoder *httpbinding.Encoder) error {
@@ -772,10 +750,8 @@ func (m *awsRestjson1_serializeOpDeletePermissionGroup
 func (m *awsRestjson1_serializeOpDeletePermissionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -813,7 +789,6 @@ func (m *awsRestjson1_serializeOpDeletePermissionGroup
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePermissionGroupInput(v *DeletePermissionGroupInput, encoder *httpbinding.Encoder) error {
@@ -847,10 +822,8 @@ func (m *awsRestjson1_serializeOpDisableUser) HandleSe
 func (m *awsRestjson1_serializeOpDisableUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -899,7 +872,6 @@ func (m *awsRestjson1_serializeOpDisableUser) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableUserInput(v *DisableUserInput, encoder *httpbinding.Encoder) error {
@@ -941,10 +913,8 @@ func (m *awsRestjson1_serializeOpDisassociateUserFromP
 func (m *awsRestjson1_serializeOpDisassociateUserFromPermissionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -982,7 +952,6 @@ func (m *awsRestjson1_serializeOpDisassociateUserFromP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateUserFromPermissionGroupInput(v *DisassociateUserFromPermissionGroupInput, encoder *httpbinding.Encoder) error {
@@ -1025,10 +994,8 @@ func (m *awsRestjson1_serializeOpEnableUser) HandleSer
 func (m *awsRestjson1_serializeOpEnableUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1077,7 +1044,6 @@ func (m *awsRestjson1_serializeOpEnableUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableUserInput(v *EnableUserInput, encoder *httpbinding.Encoder) error {
@@ -1119,10 +1085,8 @@ func (m *awsRestjson1_serializeOpGetChangeset) HandleS
 func (m *awsRestjson1_serializeOpGetChangeset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1160,7 +1124,6 @@ func (m *awsRestjson1_serializeOpGetChangeset) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetChangesetInput(v *GetChangesetInput, encoder *httpbinding.Encoder) error {
@@ -1199,10 +1162,8 @@ func (m *awsRestjson1_serializeOpGetDataset) HandleSer
 func (m *awsRestjson1_serializeOpGetDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1240,7 +1201,6 @@ func (m *awsRestjson1_serializeOpGetDataset) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDatasetInput(v *GetDatasetInput, encoder *httpbinding.Encoder) error {
@@ -1270,10 +1230,8 @@ func (m *awsRestjson1_serializeOpGetDataView) HandleSe
 func (m *awsRestjson1_serializeOpGetDataView) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1311,7 +1269,6 @@ func (m *awsRestjson1_serializeOpGetDataView) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDataViewInput(v *GetDataViewInput, encoder *httpbinding.Encoder) error {
@@ -1350,10 +1307,8 @@ func (m *awsRestjson1_serializeOpGetExternalDataViewAc
 func (m *awsRestjson1_serializeOpGetExternalDataViewAccessDetails) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1391,7 +1346,6 @@ func (m *awsRestjson1_serializeOpGetExternalDataViewAc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetExternalDataViewAccessDetailsInput(v *GetExternalDataViewAccessDetailsInput, encoder *httpbinding.Encoder) error {
@@ -1430,10 +1384,8 @@ func (m *awsRestjson1_serializeOpGetPermissionGroup) H
 func (m *awsRestjson1_serializeOpGetPermissionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1471,7 +1423,6 @@ func (m *awsRestjson1_serializeOpGetPermissionGroup) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPermissionGroupInput(v *GetPermissionGroupInput, encoder *httpbinding.Encoder) error {
@@ -1501,10 +1452,8 @@ func (m *awsRestjson1_serializeOpGetProgrammaticAccess
 func (m *awsRestjson1_serializeOpGetProgrammaticAccessCredentials) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1542,7 +1491,6 @@ func (m *awsRestjson1_serializeOpGetProgrammaticAccess
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetProgrammaticAccessCredentialsInput(v *GetProgrammaticAccessCredentialsInput, encoder *httpbinding.Encoder) error {
@@ -1571,10 +1519,8 @@ func (m *awsRestjson1_serializeOpGetUser) HandleSerial
 func (m *awsRestjson1_serializeOpGetUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1612,7 +1558,6 @@ func (m *awsRestjson1_serializeOpGetUser) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserInput(v *GetUserInput, encoder *httpbinding.Encoder) error {
@@ -1642,10 +1587,8 @@ func (m *awsRestjson1_serializeOpGetWorkingLocation) H
 func (m *awsRestjson1_serializeOpGetWorkingLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1690,7 +1633,6 @@ func (m *awsRestjson1_serializeOpGetWorkingLocation) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkingLocationInput(v *GetWorkingLocationInput, encoder *httpbinding.Encoder) error {
@@ -1723,10 +1665,8 @@ func (m *awsRestjson1_serializeOpListChangesets) Handl
 func (m *awsRestjson1_serializeOpListChangesets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1764,7 +1704,6 @@ func (m *awsRestjson1_serializeOpListChangesets) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListChangesetsInput(v *ListChangesetsInput, encoder *httpbinding.Encoder) error {
@@ -1802,10 +1741,8 @@ func (m *awsRestjson1_serializeOpListDatasets) HandleS
 func (m *awsRestjson1_serializeOpListDatasets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1843,7 +1780,6 @@ func (m *awsRestjson1_serializeOpListDatasets) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDatasetsInput(v *ListDatasetsInput, encoder *httpbinding.Encoder) error {
@@ -1872,10 +1808,8 @@ func (m *awsRestjson1_serializeOpListDataViews) Handle
 func (m *awsRestjson1_serializeOpListDataViews) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1913,7 +1847,6 @@ func (m *awsRestjson1_serializeOpListDataViews) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDataViewsInput(v *ListDataViewsInput, encoder *httpbinding.Encoder) error {
@@ -1951,10 +1884,8 @@ func (m *awsRestjson1_serializeOpListPermissionGroups)
 func (m *awsRestjson1_serializeOpListPermissionGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1992,7 +1923,6 @@ func (m *awsRestjson1_serializeOpListPermissionGroups)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPermissionGroupsInput(v *ListPermissionGroupsInput, encoder *httpbinding.Encoder) error {
@@ -2021,10 +1951,8 @@ func (m *awsRestjson1_serializeOpListPermissionGroupsB
 func (m *awsRestjson1_serializeOpListPermissionGroupsByUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2062,7 +1990,6 @@ func (m *awsRestjson1_serializeOpListPermissionGroupsB
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPermissionGroupsByUserInput(v *ListPermissionGroupsByUserInput, encoder *httpbinding.Encoder) error {
@@ -2100,10 +2027,8 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
 func (m *awsRestjson1_serializeOpListUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2141,7 +2066,6 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUsersInput(v *ListUsersInput, encoder *httpbinding.Encoder) error {
@@ -2170,10 +2094,8 @@ func (m *awsRestjson1_serializeOpListUsersByPermission
 func (m *awsRestjson1_serializeOpListUsersByPermissionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2211,7 +2133,6 @@ func (m *awsRestjson1_serializeOpListUsersByPermission
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUsersByPermissionGroupInput(v *ListUsersByPermissionGroupInput, encoder *httpbinding.Encoder) error {
@@ -2249,10 +2170,8 @@ func (m *awsRestjson1_serializeOpResetUserPassword) Ha
 func (m *awsRestjson1_serializeOpResetUserPassword) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2301,7 +2220,6 @@ func (m *awsRestjson1_serializeOpResetUserPassword) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResetUserPasswordInput(v *ResetUserPasswordInput, encoder *httpbinding.Encoder) error {
@@ -2343,10 +2261,8 @@ func (m *awsRestjson1_serializeOpUpdateChangeset) Hand
 func (m *awsRestjson1_serializeOpUpdateChangeset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2395,7 +2311,6 @@ func (m *awsRestjson1_serializeOpUpdateChangeset) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateChangesetInput(v *UpdateChangesetInput, encoder *httpbinding.Encoder) error {
@@ -2460,10 +2375,8 @@ func (m *awsRestjson1_serializeOpUpdateDataset) Handle
 func (m *awsRestjson1_serializeOpUpdateDataset) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2512,7 +2425,6 @@ func (m *awsRestjson1_serializeOpUpdateDataset) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDatasetInput(v *UpdateDatasetInput, encoder *httpbinding.Encoder) error {
@@ -2581,10 +2493,8 @@ func (m *awsRestjson1_serializeOpUpdatePermissionGroup
 func (m *awsRestjson1_serializeOpUpdatePermissionGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2633,7 +2543,6 @@ func (m *awsRestjson1_serializeOpUpdatePermissionGroup
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePermissionGroupInput(v *UpdatePermissionGroupInput, encoder *httpbinding.Encoder) error {
@@ -2692,10 +2601,8 @@ func (m *awsRestjson1_serializeOpUpdateUser) HandleSer
 func (m *awsRestjson1_serializeOpUpdateUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2744,7 +2651,6 @@ func (m *awsRestjson1_serializeOpUpdateUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserInput(v *UpdateUserInput, encoder *httpbinding.Encoder) error {
