--- vendor/github.com/aws/aws-sdk-go-v2/service/s3tables/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3tables/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpCreateNamespace) Hand
 func (m *awsRestjson1_serializeOpCreateNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpCreateNamespace) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateNamespaceInput(v *CreateNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -121,10 +117,8 @@ func (m *awsRestjson1_serializeOpCreateTable) HandleSe
 func (m *awsRestjson1_serializeOpCreateTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -173,7 +167,6 @@ func (m *awsRestjson1_serializeOpCreateTable) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTableInput(v *CreateTableInput, encoder *httpbinding.Encoder) error {
@@ -243,10 +236,8 @@ func (m *awsRestjson1_serializeOpCreateTableBucket) Ha
 func (m *awsRestjson1_serializeOpCreateTableBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -291,7 +282,6 @@ func (m *awsRestjson1_serializeOpCreateTableBucket) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTableBucketInput(v *CreateTableBucketInput, encoder *httpbinding.Encoder) error {
@@ -331,10 +321,8 @@ func (m *awsRestjson1_serializeOpDeleteNamespace) Hand
 func (m *awsRestjson1_serializeOpDeleteNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -372,7 +360,6 @@ func (m *awsRestjson1_serializeOpDeleteNamespace) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteNamespaceInput(v *DeleteNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -411,10 +398,8 @@ func (m *awsRestjson1_serializeOpDeleteTable) HandleSe
 func (m *awsRestjson1_serializeOpDeleteTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -452,7 +437,6 @@ func (m *awsRestjson1_serializeOpDeleteTable) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTableInput(v *DeleteTableInput, encoder *httpbinding.Encoder) error {
@@ -504,10 +488,8 @@ func (m *awsRestjson1_serializeOpDeleteTableBucket) Ha
 func (m *awsRestjson1_serializeOpDeleteTableBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -545,7 +527,6 @@ func (m *awsRestjson1_serializeOpDeleteTableBucket) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTableBucketInput(v *DeleteTableBucketInput, encoder *httpbinding.Encoder) error {
@@ -575,10 +556,8 @@ func (m *awsRestjson1_serializeOpDeleteTableBucketEncr
 func (m *awsRestjson1_serializeOpDeleteTableBucketEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -616,7 +595,6 @@ func (m *awsRestjson1_serializeOpDeleteTableBucketEncr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTableBucketEncryptionInput(v *DeleteTableBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -646,10 +624,8 @@ func (m *awsRestjson1_serializeOpDeleteTableBucketPoli
 func (m *awsRestjson1_serializeOpDeleteTableBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -687,7 +663,6 @@ func (m *awsRestjson1_serializeOpDeleteTableBucketPoli
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTableBucketPolicyInput(v *DeleteTableBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -717,10 +692,8 @@ func (m *awsRestjson1_serializeOpDeleteTablePolicy) Ha
 func (m *awsRestjson1_serializeOpDeleteTablePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -758,7 +731,6 @@ func (m *awsRestjson1_serializeOpDeleteTablePolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTablePolicyInput(v *DeleteTablePolicyInput, encoder *httpbinding.Encoder) error {
@@ -806,10 +778,8 @@ func (m *awsRestjson1_serializeOpGetNamespace) HandleS
 func (m *awsRestjson1_serializeOpGetNamespace) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -847,7 +817,6 @@ func (m *awsRestjson1_serializeOpGetNamespace) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetNamespaceInput(v *GetNamespaceInput, encoder *httpbinding.Encoder) error {
@@ -886,10 +855,8 @@ func (m *awsRestjson1_serializeOpGetTable) HandleSeria
 func (m *awsRestjson1_serializeOpGetTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -927,7 +894,6 @@ func (m *awsRestjson1_serializeOpGetTable) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableInput(v *GetTableInput, encoder *httpbinding.Encoder) error {
@@ -964,10 +930,8 @@ func (m *awsRestjson1_serializeOpGetTableBucket) Handl
 func (m *awsRestjson1_serializeOpGetTableBucket) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1005,7 +969,6 @@ func (m *awsRestjson1_serializeOpGetTableBucket) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableBucketInput(v *GetTableBucketInput, encoder *httpbinding.Encoder) error {
@@ -1035,10 +998,8 @@ func (m *awsRestjson1_serializeOpGetTableBucketEncrypt
 func (m *awsRestjson1_serializeOpGetTableBucketEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1076,7 +1037,6 @@ func (m *awsRestjson1_serializeOpGetTableBucketEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableBucketEncryptionInput(v *GetTableBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -1106,10 +1066,8 @@ func (m *awsRestjson1_serializeOpGetTableBucketMainten
 func (m *awsRestjson1_serializeOpGetTableBucketMaintenanceConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1147,7 +1105,6 @@ func (m *awsRestjson1_serializeOpGetTableBucketMainten
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableBucketMaintenanceConfigurationInput(v *GetTableBucketMaintenanceConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1177,10 +1134,8 @@ func (m *awsRestjson1_serializeOpGetTableBucketPolicy)
 func (m *awsRestjson1_serializeOpGetTableBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1218,7 +1173,6 @@ func (m *awsRestjson1_serializeOpGetTableBucketPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableBucketPolicyInput(v *GetTableBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -1248,10 +1202,8 @@ func (m *awsRestjson1_serializeOpGetTableEncryption) H
 func (m *awsRestjson1_serializeOpGetTableEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1289,7 +1241,6 @@ func (m *awsRestjson1_serializeOpGetTableEncryption) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableEncryptionInput(v *GetTableEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -1337,10 +1288,8 @@ func (m *awsRestjson1_serializeOpGetTableMaintenanceCo
 func (m *awsRestjson1_serializeOpGetTableMaintenanceConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1378,7 +1327,6 @@ func (m *awsRestjson1_serializeOpGetTableMaintenanceCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableMaintenanceConfigurationInput(v *GetTableMaintenanceConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1426,10 +1374,8 @@ func (m *awsRestjson1_serializeOpGetTableMaintenanceJo
 func (m *awsRestjson1_serializeOpGetTableMaintenanceJobStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1467,7 +1413,6 @@ func (m *awsRestjson1_serializeOpGetTableMaintenanceJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableMaintenanceJobStatusInput(v *GetTableMaintenanceJobStatusInput, encoder *httpbinding.Encoder) error {
@@ -1515,10 +1460,8 @@ func (m *awsRestjson1_serializeOpGetTableMetadataLocat
 func (m *awsRestjson1_serializeOpGetTableMetadataLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1556,7 +1499,6 @@ func (m *awsRestjson1_serializeOpGetTableMetadataLocat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTableMetadataLocationInput(v *GetTableMetadataLocationInput, encoder *httpbinding.Encoder) error {
@@ -1604,10 +1546,8 @@ func (m *awsRestjson1_serializeOpGetTablePolicy) Handl
 func (m *awsRestjson1_serializeOpGetTablePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1645,7 +1585,6 @@ func (m *awsRestjson1_serializeOpGetTablePolicy) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTablePolicyInput(v *GetTablePolicyInput, encoder *httpbinding.Encoder) error {
@@ -1693,10 +1632,8 @@ func (m *awsRestjson1_serializeOpListNamespaces) Handl
 func (m *awsRestjson1_serializeOpListNamespaces) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1734,7 +1671,6 @@ func (m *awsRestjson1_serializeOpListNamespaces) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNamespacesInput(v *ListNamespacesInput, encoder *httpbinding.Encoder) error {
@@ -1776,10 +1712,8 @@ func (m *awsRestjson1_serializeOpListTableBuckets) Han
 func (m *awsRestjson1_serializeOpListTableBuckets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1817,7 +1751,6 @@ func (m *awsRestjson1_serializeOpListTableBuckets) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTableBucketsInput(v *ListTableBucketsInput, encoder *httpbinding.Encoder) error {
@@ -1854,10 +1787,8 @@ func (m *awsRestjson1_serializeOpListTables) HandleSer
 func (m *awsRestjson1_serializeOpListTables) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1895,7 +1826,6 @@ func (m *awsRestjson1_serializeOpListTables) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTablesInput(v *ListTablesInput, encoder *httpbinding.Encoder) error {
@@ -1941,10 +1871,8 @@ func (m *awsRestjson1_serializeOpPutTableBucketEncrypt
 func (m *awsRestjson1_serializeOpPutTableBucketEncryption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1993,7 +1921,6 @@ func (m *awsRestjson1_serializeOpPutTableBucketEncrypt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTableBucketEncryptionInput(v *PutTableBucketEncryptionInput, encoder *httpbinding.Encoder) error {
@@ -2037,10 +1964,8 @@ func (m *awsRestjson1_serializeOpPutTableBucketMainten
 func (m *awsRestjson1_serializeOpPutTableBucketMaintenanceConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2089,7 +2014,6 @@ func (m *awsRestjson1_serializeOpPutTableBucketMainten
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTableBucketMaintenanceConfigurationInput(v *PutTableBucketMaintenanceConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2142,10 +2066,8 @@ func (m *awsRestjson1_serializeOpPutTableBucketPolicy)
 func (m *awsRestjson1_serializeOpPutTableBucketPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2194,7 +2116,6 @@ func (m *awsRestjson1_serializeOpPutTableBucketPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTableBucketPolicyInput(v *PutTableBucketPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2236,10 +2157,8 @@ func (m *awsRestjson1_serializeOpPutTableMaintenanceCo
 func (m *awsRestjson1_serializeOpPutTableMaintenanceConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2288,7 +2207,6 @@ func (m *awsRestjson1_serializeOpPutTableMaintenanceCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTableMaintenanceConfigurationInput(v *PutTableMaintenanceConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2359,10 +2277,8 @@ func (m *awsRestjson1_serializeOpPutTablePolicy) Handl
 func (m *awsRestjson1_serializeOpPutTablePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2411,7 +2327,6 @@ func (m *awsRestjson1_serializeOpPutTablePolicy) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutTablePolicyInput(v *PutTablePolicyInput, encoder *httpbinding.Encoder) error {
@@ -2471,10 +2386,8 @@ func (m *awsRestjson1_serializeOpRenameTable) HandleSe
 func (m *awsRestjson1_serializeOpRenameTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2523,7 +2436,6 @@ func (m *awsRestjson1_serializeOpRenameTable) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRenameTableInput(v *RenameTableInput, encoder *httpbinding.Encoder) error {
@@ -2593,10 +2505,8 @@ func (m *awsRestjson1_serializeOpUpdateTableMetadataLo
 func (m *awsRestjson1_serializeOpUpdateTableMetadataLocation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2645,7 +2555,6 @@ func (m *awsRestjson1_serializeOpUpdateTableMetadataLo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTableMetadataLocationInput(v *UpdateTableMetadataLocationInput, encoder *httpbinding.Encoder) error {
