--- vendor/github.com/aws/aws-sdk-go-v2/service/clouddirectory/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/clouddirectory/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpAddFacetToObject) Han
 func (m *awsRestjson1_serializeOpAddFacetToObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -78,7 +75,6 @@ func (m *awsRestjson1_serializeOpAddFacetToObject) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAddFacetToObjectInput(v *AddFacetToObjectInput, encoder *httpbinding.Encoder) error {
@@ -132,10 +128,8 @@ func (m *awsRestjson1_serializeOpApplySchema) HandleSe
 func (m *awsRestjson1_serializeOpApplySchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -184,7 +178,6 @@ func (m *awsRestjson1_serializeOpApplySchema) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsApplySchemaInput(v *ApplySchemaInput, encoder *httpbinding.Encoder) error {
@@ -222,10 +215,8 @@ func (m *awsRestjson1_serializeOpAttachObject) HandleS
 func (m *awsRestjson1_serializeOpAttachObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -274,7 +265,6 @@ func (m *awsRestjson1_serializeOpAttachObject) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachObjectInput(v *AttachObjectInput, encoder *httpbinding.Encoder) error {
@@ -326,10 +316,8 @@ func (m *awsRestjson1_serializeOpAttachPolicy) HandleS
 func (m *awsRestjson1_serializeOpAttachPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -378,7 +366,6 @@ func (m *awsRestjson1_serializeOpAttachPolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachPolicyInput(v *AttachPolicyInput, encoder *httpbinding.Encoder) error {
@@ -425,10 +412,8 @@ func (m *awsRestjson1_serializeOpAttachToIndex) Handle
 func (m *awsRestjson1_serializeOpAttachToIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -477,7 +462,6 @@ func (m *awsRestjson1_serializeOpAttachToIndex) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachToIndexInput(v *AttachToIndexInput, encoder *httpbinding.Encoder) error {
@@ -524,10 +508,8 @@ func (m *awsRestjson1_serializeOpAttachTypedLink) Hand
 func (m *awsRestjson1_serializeOpAttachTypedLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -576,7 +558,6 @@ func (m *awsRestjson1_serializeOpAttachTypedLink) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAttachTypedLinkInput(v *AttachTypedLinkInput, encoder *httpbinding.Encoder) error {
@@ -637,10 +618,8 @@ func (m *awsRestjson1_serializeOpBatchRead) HandleSeri
 func (m *awsRestjson1_serializeOpBatchRead) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -689,7 +668,6 @@ func (m *awsRestjson1_serializeOpBatchRead) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchReadInput(v *BatchReadInput, encoder *httpbinding.Encoder) error {
@@ -734,10 +712,8 @@ func (m *awsRestjson1_serializeOpBatchWrite) HandleSer
 func (m *awsRestjson1_serializeOpBatchWrite) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -786,7 +762,6 @@ func (m *awsRestjson1_serializeOpBatchWrite) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchWriteInput(v *BatchWriteInput, encoder *httpbinding.Encoder) error {
@@ -826,10 +801,8 @@ func (m *awsRestjson1_serializeOpCreateDirectory) Hand
 func (m *awsRestjson1_serializeOpCreateDirectory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -878,7 +851,6 @@ func (m *awsRestjson1_serializeOpCreateDirectory) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDirectoryInput(v *CreateDirectoryInput, encoder *httpbinding.Encoder) error {
@@ -916,10 +888,8 @@ func (m *awsRestjson1_serializeOpCreateFacet) HandleSe
 func (m *awsRestjson1_serializeOpCreateFacet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -968,7 +938,6 @@ func (m *awsRestjson1_serializeOpCreateFacet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFacetInput(v *CreateFacetInput, encoder *httpbinding.Encoder) error {
@@ -1023,10 +992,8 @@ func (m *awsRestjson1_serializeOpCreateIndex) HandleSe
 func (m *awsRestjson1_serializeOpCreateIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1075,7 +1042,6 @@ func (m *awsRestjson1_serializeOpCreateIndex) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIndexInput(v *CreateIndexInput, encoder *httpbinding.Encoder) error {
@@ -1132,10 +1098,8 @@ func (m *awsRestjson1_serializeOpCreateObject) HandleS
 func (m *awsRestjson1_serializeOpCreateObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1184,7 +1148,6 @@ func (m *awsRestjson1_serializeOpCreateObject) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateObjectInput(v *CreateObjectInput, encoder *httpbinding.Encoder) error {
@@ -1243,10 +1206,8 @@ func (m *awsRestjson1_serializeOpCreateSchema) HandleS
 func (m *awsRestjson1_serializeOpCreateSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1291,7 +1252,6 @@ func (m *awsRestjson1_serializeOpCreateSchema) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSchemaInput(v *CreateSchemaInput, encoder *httpbinding.Encoder) error {
@@ -1324,10 +1284,8 @@ func (m *awsRestjson1_serializeOpCreateTypedLinkFacet)
 func (m *awsRestjson1_serializeOpCreateTypedLinkFacet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1376,7 +1334,6 @@ func (m *awsRestjson1_serializeOpCreateTypedLinkFacet)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateTypedLinkFacetInput(v *CreateTypedLinkFacetInput, encoder *httpbinding.Encoder) error {
@@ -1416,10 +1373,8 @@ func (m *awsRestjson1_serializeOpDeleteDirectory) Hand
 func (m *awsRestjson1_serializeOpDeleteDirectory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1457,7 +1412,6 @@ func (m *awsRestjson1_serializeOpDeleteDirectory) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDirectoryInput(v *DeleteDirectoryInput, encoder *httpbinding.Encoder) error {
@@ -1483,10 +1437,8 @@ func (m *awsRestjson1_serializeOpDeleteFacet) HandleSe
 func (m *awsRestjson1_serializeOpDeleteFacet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1535,7 +1487,6 @@ func (m *awsRestjson1_serializeOpDeleteFacet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFacetInput(v *DeleteFacetInput, encoder *httpbinding.Encoder) error {
@@ -1573,10 +1524,8 @@ func (m *awsRestjson1_serializeOpDeleteObject) HandleS
 func (m *awsRestjson1_serializeOpDeleteObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1625,7 +1574,6 @@ func (m *awsRestjson1_serializeOpDeleteObject) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteObjectInput(v *DeleteObjectInput, encoder *httpbinding.Encoder) error {
@@ -1665,10 +1613,8 @@ func (m *awsRestjson1_serializeOpDeleteSchema) HandleS
 func (m *awsRestjson1_serializeOpDeleteSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1706,7 +1652,6 @@ func (m *awsRestjson1_serializeOpDeleteSchema) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSchemaInput(v *DeleteSchemaInput, encoder *httpbinding.Encoder) error {
@@ -1732,10 +1677,8 @@ func (m *awsRestjson1_serializeOpDeleteTypedLinkFacet)
 func (m *awsRestjson1_serializeOpDeleteTypedLinkFacet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1784,7 +1727,6 @@ func (m *awsRestjson1_serializeOpDeleteTypedLinkFacet)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteTypedLinkFacetInput(v *DeleteTypedLinkFacetInput, encoder *httpbinding.Encoder) error {
@@ -1822,10 +1764,8 @@ func (m *awsRestjson1_serializeOpDetachFromIndex) Hand
 func (m *awsRestjson1_serializeOpDetachFromIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1874,7 +1814,6 @@ func (m *awsRestjson1_serializeOpDetachFromIndex) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachFromIndexInput(v *DetachFromIndexInput, encoder *httpbinding.Encoder) error {
@@ -1921,10 +1860,8 @@ func (m *awsRestjson1_serializeOpDetachObject) HandleS
 func (m *awsRestjson1_serializeOpDetachObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1973,7 +1910,6 @@ func (m *awsRestjson1_serializeOpDetachObject) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachObjectInput(v *DetachObjectInput, encoder *httpbinding.Encoder) error {
@@ -2018,10 +1954,8 @@ func (m *awsRestjson1_serializeOpDetachPolicy) HandleS
 func (m *awsRestjson1_serializeOpDetachPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2070,7 +2004,6 @@ func (m *awsRestjson1_serializeOpDetachPolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachPolicyInput(v *DetachPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2117,10 +2050,8 @@ func (m *awsRestjson1_serializeOpDetachTypedLink) Hand
 func (m *awsRestjson1_serializeOpDetachTypedLink) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2169,7 +2100,6 @@ func (m *awsRestjson1_serializeOpDetachTypedLink) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDetachTypedLinkInput(v *DetachTypedLinkInput, encoder *httpbinding.Encoder) error {
@@ -2209,10 +2139,8 @@ func (m *awsRestjson1_serializeOpDisableDirectory) Han
 func (m *awsRestjson1_serializeOpDisableDirectory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2250,7 +2178,6 @@ func (m *awsRestjson1_serializeOpDisableDirectory) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableDirectoryInput(v *DisableDirectoryInput, encoder *httpbinding.Encoder) error {
@@ -2276,10 +2203,8 @@ func (m *awsRestjson1_serializeOpEnableDirectory) Hand
 func (m *awsRestjson1_serializeOpEnableDirectory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2317,7 +2242,6 @@ func (m *awsRestjson1_serializeOpEnableDirectory) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableDirectoryInput(v *EnableDirectoryInput, encoder *httpbinding.Encoder) error {
@@ -2343,10 +2267,8 @@ func (m *awsRestjson1_serializeOpGetAppliedSchemaVersi
 func (m *awsRestjson1_serializeOpGetAppliedSchemaVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2391,7 +2313,6 @@ func (m *awsRestjson1_serializeOpGetAppliedSchemaVersi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAppliedSchemaVersionInput(v *GetAppliedSchemaVersionInput, encoder *httpbinding.Encoder) error {
@@ -2424,10 +2345,8 @@ func (m *awsRestjson1_serializeOpGetDirectory) HandleS
 func (m *awsRestjson1_serializeOpGetDirectory) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2465,7 +2384,6 @@ func (m *awsRestjson1_serializeOpGetDirectory) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDirectoryInput(v *GetDirectoryInput, encoder *httpbinding.Encoder) error {
@@ -2491,10 +2409,8 @@ func (m *awsRestjson1_serializeOpGetFacet) HandleSeria
 func (m *awsRestjson1_serializeOpGetFacet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2543,7 +2459,6 @@ func (m *awsRestjson1_serializeOpGetFacet) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFacetInput(v *GetFacetInput, encoder *httpbinding.Encoder) error {
@@ -2581,10 +2496,8 @@ func (m *awsRestjson1_serializeOpGetLinkAttributes) Ha
 func (m *awsRestjson1_serializeOpGetLinkAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2633,7 +2546,6 @@ func (m *awsRestjson1_serializeOpGetLinkAttributes) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLinkAttributesInput(v *GetLinkAttributesInput, encoder *httpbinding.Encoder) error {
@@ -2685,10 +2597,8 @@ func (m *awsRestjson1_serializeOpGetObjectAttributes) 
 func (m *awsRestjson1_serializeOpGetObjectAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2737,7 +2647,6 @@ func (m *awsRestjson1_serializeOpGetObjectAttributes) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetObjectAttributesInput(v *GetObjectAttributesInput, encoder *httpbinding.Encoder) error {
@@ -2796,10 +2705,8 @@ func (m *awsRestjson1_serializeOpGetObjectInformation)
 func (m *awsRestjson1_serializeOpGetObjectInformation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2848,7 +2755,6 @@ func (m *awsRestjson1_serializeOpGetObjectInformation)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetObjectInformationInput(v *GetObjectInformationInput, encoder *httpbinding.Encoder) error {
@@ -2893,10 +2799,8 @@ func (m *awsRestjson1_serializeOpGetSchemaAsJson) Hand
 func (m *awsRestjson1_serializeOpGetSchemaAsJson) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2934,7 +2838,6 @@ func (m *awsRestjson1_serializeOpGetSchemaAsJson) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSchemaAsJsonInput(v *GetSchemaAsJsonInput, encoder *httpbinding.Encoder) error {
@@ -2960,10 +2863,8 @@ func (m *awsRestjson1_serializeOpGetTypedLinkFacetInfo
 func (m *awsRestjson1_serializeOpGetTypedLinkFacetInformation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3012,7 +2913,6 @@ func (m *awsRestjson1_serializeOpGetTypedLinkFacetInfo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTypedLinkFacetInformationInput(v *GetTypedLinkFacetInformationInput, encoder *httpbinding.Encoder) error {
@@ -3050,10 +2950,8 @@ func (m *awsRestjson1_serializeOpListAppliedSchemaArns
 func (m *awsRestjson1_serializeOpListAppliedSchemaArns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3098,7 +2996,6 @@ func (m *awsRestjson1_serializeOpListAppliedSchemaArns
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAppliedSchemaArnsInput(v *ListAppliedSchemaArnsInput, encoder *httpbinding.Encoder) error {
@@ -3146,10 +3043,8 @@ func (m *awsRestjson1_serializeOpListAttachedIndices) 
 func (m *awsRestjson1_serializeOpListAttachedIndices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3198,7 +3093,6 @@ func (m *awsRestjson1_serializeOpListAttachedIndices) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAttachedIndicesInput(v *ListAttachedIndicesInput, encoder *httpbinding.Encoder) error {
@@ -3253,10 +3147,8 @@ func (m *awsRestjson1_serializeOpListDevelopmentSchema
 func (m *awsRestjson1_serializeOpListDevelopmentSchemaArns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3301,7 +3193,6 @@ func (m *awsRestjson1_serializeOpListDevelopmentSchema
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDevelopmentSchemaArnsInput(v *ListDevelopmentSchemaArnsInput, encoder *httpbinding.Encoder) error {
@@ -3339,10 +3230,8 @@ func (m *awsRestjson1_serializeOpListDirectories) Hand
 func (m *awsRestjson1_serializeOpListDirectories) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3387,7 +3276,6 @@ func (m *awsRestjson1_serializeOpListDirectories) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDirectoriesInput(v *ListDirectoriesInput, encoder *httpbinding.Encoder) error {
@@ -3430,10 +3318,8 @@ func (m *awsRestjson1_serializeOpListFacetAttributes) 
 func (m *awsRestjson1_serializeOpListFacetAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3482,7 +3368,6 @@ func (m *awsRestjson1_serializeOpListFacetAttributes) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFacetAttributesInput(v *ListFacetAttributesInput, encoder *httpbinding.Encoder) error {
@@ -3530,10 +3415,8 @@ func (m *awsRestjson1_serializeOpListFacetNames) Handl
 func (m *awsRestjson1_serializeOpListFacetNames) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3582,7 +3465,6 @@ func (m *awsRestjson1_serializeOpListFacetNames) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFacetNamesInput(v *ListFacetNamesInput, encoder *httpbinding.Encoder) error {
@@ -3625,10 +3507,8 @@ func (m *awsRestjson1_serializeOpListIncomingTypedLink
 func (m *awsRestjson1_serializeOpListIncomingTypedLinks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3677,7 +3557,6 @@ func (m *awsRestjson1_serializeOpListIncomingTypedLink
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIncomingTypedLinksInput(v *ListIncomingTypedLinksInput, encoder *httpbinding.Encoder) error {
@@ -3746,10 +3625,8 @@ func (m *awsRestjson1_serializeOpListIndex) HandleSeri
 func (m *awsRestjson1_serializeOpListIndex) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3798,7 +3675,6 @@ func (m *awsRestjson1_serializeOpListIndex) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIndexInput(v *ListIndexInput, encoder *httpbinding.Encoder) error {
@@ -3860,10 +3736,8 @@ func (m *awsRestjson1_serializeOpListManagedSchemaArns
 func (m *awsRestjson1_serializeOpListManagedSchemaArns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3908,7 +3782,6 @@ func (m *awsRestjson1_serializeOpListManagedSchemaArns
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListManagedSchemaArnsInput(v *ListManagedSchemaArnsInput, encoder *httpbinding.Encoder) error {
@@ -3951,10 +3824,8 @@ func (m *awsRestjson1_serializeOpListObjectAttributes)
 func (m *awsRestjson1_serializeOpListObjectAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4003,7 +3874,6 @@ func (m *awsRestjson1_serializeOpListObjectAttributes)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListObjectAttributesInput(v *ListObjectAttributesInput, encoder *httpbinding.Encoder) error {
@@ -4065,10 +3935,8 @@ func (m *awsRestjson1_serializeOpListObjectChildren) H
 func (m *awsRestjson1_serializeOpListObjectChildren) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4117,7 +3985,6 @@ func (m *awsRestjson1_serializeOpListObjectChildren) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListObjectChildrenInput(v *ListObjectChildrenInput, encoder *httpbinding.Encoder) error {
@@ -4172,10 +4039,8 @@ func (m *awsRestjson1_serializeOpListObjectParentPaths
 func (m *awsRestjson1_serializeOpListObjectParentPaths) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4224,7 +4089,6 @@ func (m *awsRestjson1_serializeOpListObjectParentPaths
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListObjectParentPathsInput(v *ListObjectParentPathsInput, encoder *httpbinding.Encoder) error {
@@ -4274,10 +4138,8 @@ func (m *awsRestjson1_serializeOpListObjectParents) Ha
 func (m *awsRestjson1_serializeOpListObjectParents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4326,7 +4188,6 @@ func (m *awsRestjson1_serializeOpListObjectParents) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListObjectParentsInput(v *ListObjectParentsInput, encoder *httpbinding.Encoder) error {
@@ -4386,10 +4247,8 @@ func (m *awsRestjson1_serializeOpListObjectPolicies) H
 func (m *awsRestjson1_serializeOpListObjectPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4438,7 +4297,6 @@ func (m *awsRestjson1_serializeOpListObjectPolicies) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListObjectPoliciesInput(v *ListObjectPoliciesInput, encoder *httpbinding.Encoder) error {
@@ -4493,10 +4351,8 @@ func (m *awsRestjson1_serializeOpListOutgoingTypedLink
 func (m *awsRestjson1_serializeOpListOutgoingTypedLinks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4545,7 +4401,6 @@ func (m *awsRestjson1_serializeOpListOutgoingTypedLink
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOutgoingTypedLinksInput(v *ListOutgoingTypedLinksInput, encoder *httpbinding.Encoder) error {
@@ -4614,10 +4469,8 @@ func (m *awsRestjson1_serializeOpListPolicyAttachments
 func (m *awsRestjson1_serializeOpListPolicyAttachments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4666,7 +4519,6 @@ func (m *awsRestjson1_serializeOpListPolicyAttachments
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPolicyAttachmentsInput(v *ListPolicyAttachmentsInput, encoder *httpbinding.Encoder) error {
@@ -4721,10 +4573,8 @@ func (m *awsRestjson1_serializeOpListPublishedSchemaAr
 func (m *awsRestjson1_serializeOpListPublishedSchemaArns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4769,7 +4619,6 @@ func (m *awsRestjson1_serializeOpListPublishedSchemaAr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPublishedSchemaArnsInput(v *ListPublishedSchemaArnsInput, encoder *httpbinding.Encoder) error {
@@ -4812,10 +4661,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -4860,7 +4707,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -4903,10 +4749,8 @@ func (m *awsRestjson1_serializeOpListTypedLinkFacetAtt
 func (m *awsRestjson1_serializeOpListTypedLinkFacetAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4955,7 +4799,6 @@ func (m *awsRestjson1_serializeOpListTypedLinkFacetAtt
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTypedLinkFacetAttributesInput(v *ListTypedLinkFacetAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5003,10 +4846,8 @@ func (m *awsRestjson1_serializeOpListTypedLinkFacetNam
 func (m *awsRestjson1_serializeOpListTypedLinkFacetNames) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5055,7 +4896,6 @@ func (m *awsRestjson1_serializeOpListTypedLinkFacetNam
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTypedLinkFacetNamesInput(v *ListTypedLinkFacetNamesInput, encoder *httpbinding.Encoder) error {
@@ -5098,10 +4938,8 @@ func (m *awsRestjson1_serializeOpLookupPolicy) HandleS
 func (m *awsRestjson1_serializeOpLookupPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5150,7 +4988,6 @@ func (m *awsRestjson1_serializeOpLookupPolicy) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsLookupPolicyInput(v *LookupPolicyInput, encoder *httpbinding.Encoder) error {
@@ -5200,10 +5037,8 @@ func (m *awsRestjson1_serializeOpPublishSchema) Handle
 func (m *awsRestjson1_serializeOpPublishSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5252,7 +5087,6 @@ func (m *awsRestjson1_serializeOpPublishSchema) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPublishSchemaInput(v *PublishSchemaInput, encoder *httpbinding.Encoder) error {
@@ -5300,10 +5134,8 @@ func (m *awsRestjson1_serializeOpPutSchemaFromJson) Ha
 func (m *awsRestjson1_serializeOpPutSchemaFromJson) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5352,7 +5184,6 @@ func (m *awsRestjson1_serializeOpPutSchemaFromJson) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutSchemaFromJsonInput(v *PutSchemaFromJsonInput, encoder *httpbinding.Encoder) error {
@@ -5390,10 +5221,8 @@ func (m *awsRestjson1_serializeOpRemoveFacetFromObject
 func (m *awsRestjson1_serializeOpRemoveFacetFromObject) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5442,7 +5271,6 @@ func (m *awsRestjson1_serializeOpRemoveFacetFromObject
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRemoveFacetFromObjectInput(v *RemoveFacetFromObjectInput, encoder *httpbinding.Encoder) error {
@@ -5489,10 +5317,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5537,7 +5363,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5577,10 +5402,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5625,7 +5448,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5665,10 +5487,8 @@ func (m *awsRestjson1_serializeOpUpdateFacet) HandleSe
 func (m *awsRestjson1_serializeOpUpdateFacet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5717,7 +5537,6 @@ func (m *awsRestjson1_serializeOpUpdateFacet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFacetInput(v *UpdateFacetInput, encoder *httpbinding.Encoder) error {
@@ -5767,10 +5586,8 @@ func (m *awsRestjson1_serializeOpUpdateLinkAttributes)
 func (m *awsRestjson1_serializeOpUpdateLinkAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5819,7 +5636,6 @@ func (m *awsRestjson1_serializeOpUpdateLinkAttributes)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLinkAttributesInput(v *UpdateLinkAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5866,10 +5682,8 @@ func (m *awsRestjson1_serializeOpUpdateObjectAttribute
 func (m *awsRestjson1_serializeOpUpdateObjectAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5918,7 +5732,6 @@ func (m *awsRestjson1_serializeOpUpdateObjectAttribute
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateObjectAttributesInput(v *UpdateObjectAttributesInput, encoder *httpbinding.Encoder) error {
@@ -5965,10 +5778,8 @@ func (m *awsRestjson1_serializeOpUpdateSchema) HandleS
 func (m *awsRestjson1_serializeOpUpdateSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6017,7 +5828,6 @@ func (m *awsRestjson1_serializeOpUpdateSchema) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSchemaInput(v *UpdateSchemaInput, encoder *httpbinding.Encoder) error {
@@ -6055,10 +5865,8 @@ func (m *awsRestjson1_serializeOpUpdateTypedLinkFacet)
 func (m *awsRestjson1_serializeOpUpdateTypedLinkFacet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6107,7 +5915,6 @@ func (m *awsRestjson1_serializeOpUpdateTypedLinkFacet)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTypedLinkFacetInput(v *UpdateTypedLinkFacetInput, encoder *httpbinding.Encoder) error {
@@ -6159,10 +5966,8 @@ func (m *awsRestjson1_serializeOpUpgradeAppliedSchema)
 func (m *awsRestjson1_serializeOpUpgradeAppliedSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6207,7 +6012,6 @@ func (m *awsRestjson1_serializeOpUpgradeAppliedSchema)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpgradeAppliedSchemaInput(v *UpgradeAppliedSchemaInput, encoder *httpbinding.Encoder) error {
@@ -6250,10 +6054,8 @@ func (m *awsRestjson1_serializeOpUpgradePublishedSchem
 func (m *awsRestjson1_serializeOpUpgradePublishedSchema) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6298,7 +6100,6 @@ func (m *awsRestjson1_serializeOpUpgradePublishedSchem
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpgradePublishedSchemaInput(v *UpgradePublishedSchemaInput, encoder *httpbinding.Encoder) error {
