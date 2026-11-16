--- vendor/github.com/aws/aws-sdk-go-v2/service/omics/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/omics/serializers.go
@@ -14,7 +14,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -28,10 +27,8 @@ func (m *awsRestjson1_serializeOpAbortMultipartReadSet
 func (m *awsRestjson1_serializeOpAbortMultipartReadSetUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -69,7 +66,6 @@ func (m *awsRestjson1_serializeOpAbortMultipartReadSet
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAbortMultipartReadSetUploadInput(v *AbortMultipartReadSetUploadInput, encoder *httpbinding.Encoder) error {
@@ -108,10 +104,8 @@ func (m *awsRestjson1_serializeOpAcceptShare) HandleSe
 func (m *awsRestjson1_serializeOpAcceptShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -149,7 +143,6 @@ func (m *awsRestjson1_serializeOpAcceptShare) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptShareInput(v *AcceptShareInput, encoder *httpbinding.Encoder) error {
@@ -179,10 +172,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteReadSet) H
 func (m *awsRestjson1_serializeOpBatchDeleteReadSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -231,7 +222,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteReadSet) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteReadSetInput(v *BatchDeleteReadSetInput, encoder *httpbinding.Encoder) error {
@@ -275,10 +265,8 @@ func (m *awsRestjson1_serializeOpCancelAnnotationImpor
 func (m *awsRestjson1_serializeOpCancelAnnotationImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -316,7 +304,6 @@ func (m *awsRestjson1_serializeOpCancelAnnotationImpor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelAnnotationImportJobInput(v *CancelAnnotationImportJobInput, encoder *httpbinding.Encoder) error {
@@ -346,10 +333,8 @@ func (m *awsRestjson1_serializeOpCancelRun) HandleSeri
 func (m *awsRestjson1_serializeOpCancelRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -387,7 +372,6 @@ func (m *awsRestjson1_serializeOpCancelRun) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelRunInput(v *CancelRunInput, encoder *httpbinding.Encoder) error {
@@ -417,10 +401,8 @@ func (m *awsRestjson1_serializeOpCancelVariantImportJo
 func (m *awsRestjson1_serializeOpCancelVariantImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -458,7 +440,6 @@ func (m *awsRestjson1_serializeOpCancelVariantImportJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelVariantImportJobInput(v *CancelVariantImportJobInput, encoder *httpbinding.Encoder) error {
@@ -488,10 +469,8 @@ func (m *awsRestjson1_serializeOpCompleteMultipartRead
 func (m *awsRestjson1_serializeOpCompleteMultipartReadSetUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -540,7 +519,6 @@ func (m *awsRestjson1_serializeOpCompleteMultipartRead
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCompleteMultipartReadSetUploadInput(v *CompleteMultipartReadSetUploadInput, encoder *httpbinding.Encoder) error {
@@ -593,10 +571,8 @@ func (m *awsRestjson1_serializeOpCreateAnnotationStore
 func (m *awsRestjson1_serializeOpCreateAnnotationStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -641,7 +617,6 @@ func (m *awsRestjson1_serializeOpCreateAnnotationStore
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAnnotationStoreInput(v *CreateAnnotationStoreInput, encoder *httpbinding.Encoder) error {
@@ -717,10 +692,8 @@ func (m *awsRestjson1_serializeOpCreateAnnotationStore
 func (m *awsRestjson1_serializeOpCreateAnnotationStoreVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -769,7 +742,6 @@ func (m *awsRestjson1_serializeOpCreateAnnotationStore
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAnnotationStoreVersionInput(v *CreateAnnotationStoreVersionInput, encoder *httpbinding.Encoder) error {
@@ -830,10 +802,8 @@ func (m *awsRestjson1_serializeOpCreateMultipartReadSe
 func (m *awsRestjson1_serializeOpCreateMultipartReadSetUpload) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -882,7 +852,6 @@ func (m *awsRestjson1_serializeOpCreateMultipartReadSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMultipartReadSetUploadInput(v *CreateMultipartReadSetUploadInput, encoder *httpbinding.Encoder) error {
@@ -966,10 +935,8 @@ func (m *awsRestjson1_serializeOpCreateReferenceStore)
 func (m *awsRestjson1_serializeOpCreateReferenceStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1014,7 +981,6 @@ func (m *awsRestjson1_serializeOpCreateReferenceStore)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateReferenceStoreInput(v *CreateReferenceStoreInput, encoder *httpbinding.Encoder) error {
@@ -1071,10 +1037,8 @@ func (m *awsRestjson1_serializeOpCreateRunCache) Handl
 func (m *awsRestjson1_serializeOpCreateRunCache) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1119,7 +1083,6 @@ func (m *awsRestjson1_serializeOpCreateRunCache) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRunCacheInput(v *CreateRunCacheInput, encoder *httpbinding.Encoder) error {
@@ -1184,10 +1147,8 @@ func (m *awsRestjson1_serializeOpCreateRunGroup) Handl
 func (m *awsRestjson1_serializeOpCreateRunGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1232,7 +1193,6 @@ func (m *awsRestjson1_serializeOpCreateRunGroup) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRunGroupInput(v *CreateRunGroupInput, encoder *httpbinding.Encoder) error {
@@ -1297,10 +1257,8 @@ func (m *awsRestjson1_serializeOpCreateSequenceStore) 
 func (m *awsRestjson1_serializeOpCreateSequenceStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1345,7 +1303,6 @@ func (m *awsRestjson1_serializeOpCreateSequenceStore) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSequenceStoreInput(v *CreateSequenceStoreInput, encoder *httpbinding.Encoder) error {
@@ -1426,10 +1383,8 @@ func (m *awsRestjson1_serializeOpCreateShare) HandleSe
 func (m *awsRestjson1_serializeOpCreateShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1474,7 +1429,6 @@ func (m *awsRestjson1_serializeOpCreateShare) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateShareInput(v *CreateShareInput, encoder *httpbinding.Encoder) error {
@@ -1517,10 +1471,8 @@ func (m *awsRestjson1_serializeOpCreateVariantStore) H
 func (m *awsRestjson1_serializeOpCreateVariantStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1565,7 +1517,6 @@ func (m *awsRestjson1_serializeOpCreateVariantStore) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateVariantStoreInput(v *CreateVariantStoreInput, encoder *httpbinding.Encoder) error {
@@ -1624,10 +1575,8 @@ func (m *awsRestjson1_serializeOpCreateWorkflow) Handl
 func (m *awsRestjson1_serializeOpCreateWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1672,7 +1621,6 @@ func (m *awsRestjson1_serializeOpCreateWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWorkflowInput(v *CreateWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -1764,10 +1712,8 @@ func (m *awsRestjson1_serializeOpCreateWorkflowVersion
 func (m *awsRestjson1_serializeOpCreateWorkflowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1816,7 +1762,6 @@ func (m *awsRestjson1_serializeOpCreateWorkflowVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWorkflowVersionInput(v *CreateWorkflowVersionInput, encoder *httpbinding.Encoder) error {
@@ -1922,10 +1867,8 @@ func (m *awsRestjson1_serializeOpDeleteAnnotationStore
 func (m *awsRestjson1_serializeOpDeleteAnnotationStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1963,7 +1906,6 @@ func (m *awsRestjson1_serializeOpDeleteAnnotationStore
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAnnotationStoreInput(v *DeleteAnnotationStoreInput, encoder *httpbinding.Encoder) error {
@@ -1997,10 +1939,8 @@ func (m *awsRestjson1_serializeOpDeleteAnnotationStore
 func (m *awsRestjson1_serializeOpDeleteAnnotationStoreVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2049,7 +1989,6 @@ func (m *awsRestjson1_serializeOpDeleteAnnotationStore
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAnnotationStoreVersionsInput(v *DeleteAnnotationStoreVersionsInput, encoder *httpbinding.Encoder) error {
@@ -2097,10 +2036,8 @@ func (m *awsRestjson1_serializeOpDeleteReference) Hand
 func (m *awsRestjson1_serializeOpDeleteReference) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2138,7 +2075,6 @@ func (m *awsRestjson1_serializeOpDeleteReference) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteReferenceInput(v *DeleteReferenceInput, encoder *httpbinding.Encoder) error {
@@ -2177,10 +2113,8 @@ func (m *awsRestjson1_serializeOpDeleteReferenceStore)
 func (m *awsRestjson1_serializeOpDeleteReferenceStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2218,7 +2152,6 @@ func (m *awsRestjson1_serializeOpDeleteReferenceStore)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteReferenceStoreInput(v *DeleteReferenceStoreInput, encoder *httpbinding.Encoder) error {
@@ -2248,10 +2181,8 @@ func (m *awsRestjson1_serializeOpDeleteRun) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2289,7 +2220,6 @@ func (m *awsRestjson1_serializeOpDeleteRun) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRunInput(v *DeleteRunInput, encoder *httpbinding.Encoder) error {
@@ -2319,10 +2249,8 @@ func (m *awsRestjson1_serializeOpDeleteRunCache) Handl
 func (m *awsRestjson1_serializeOpDeleteRunCache) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2360,7 +2288,6 @@ func (m *awsRestjson1_serializeOpDeleteRunCache) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRunCacheInput(v *DeleteRunCacheInput, encoder *httpbinding.Encoder) error {
@@ -2390,10 +2317,8 @@ func (m *awsRestjson1_serializeOpDeleteRunGroup) Handl
 func (m *awsRestjson1_serializeOpDeleteRunGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2431,7 +2356,6 @@ func (m *awsRestjson1_serializeOpDeleteRunGroup) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRunGroupInput(v *DeleteRunGroupInput, encoder *httpbinding.Encoder) error {
@@ -2461,10 +2385,8 @@ func (m *awsRestjson1_serializeOpDeleteS3AccessPolicy)
 func (m *awsRestjson1_serializeOpDeleteS3AccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2502,7 +2424,6 @@ func (m *awsRestjson1_serializeOpDeleteS3AccessPolicy)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteS3AccessPolicyInput(v *DeleteS3AccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -2532,10 +2453,8 @@ func (m *awsRestjson1_serializeOpDeleteSequenceStore) 
 func (m *awsRestjson1_serializeOpDeleteSequenceStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2573,7 +2492,6 @@ func (m *awsRestjson1_serializeOpDeleteSequenceStore) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSequenceStoreInput(v *DeleteSequenceStoreInput, encoder *httpbinding.Encoder) error {
@@ -2603,10 +2521,8 @@ func (m *awsRestjson1_serializeOpDeleteShare) HandleSe
 func (m *awsRestjson1_serializeOpDeleteShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2644,7 +2560,6 @@ func (m *awsRestjson1_serializeOpDeleteShare) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteShareInput(v *DeleteShareInput, encoder *httpbinding.Encoder) error {
@@ -2674,10 +2589,8 @@ func (m *awsRestjson1_serializeOpDeleteVariantStore) H
 func (m *awsRestjson1_serializeOpDeleteVariantStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2715,7 +2628,6 @@ func (m *awsRestjson1_serializeOpDeleteVariantStore) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteVariantStoreInput(v *DeleteVariantStoreInput, encoder *httpbinding.Encoder) error {
@@ -2749,10 +2661,8 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
 func (m *awsRestjson1_serializeOpDeleteWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2790,7 +2700,6 @@ func (m *awsRestjson1_serializeOpDeleteWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWorkflowInput(v *DeleteWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -2820,10 +2729,8 @@ func (m *awsRestjson1_serializeOpDeleteWorkflowVersion
 func (m *awsRestjson1_serializeOpDeleteWorkflowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2861,7 +2768,6 @@ func (m *awsRestjson1_serializeOpDeleteWorkflowVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWorkflowVersionInput(v *DeleteWorkflowVersionInput, encoder *httpbinding.Encoder) error {
@@ -2900,10 +2806,8 @@ func (m *awsRestjson1_serializeOpGetAnnotationImportJo
 func (m *awsRestjson1_serializeOpGetAnnotationImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2941,7 +2845,6 @@ func (m *awsRestjson1_serializeOpGetAnnotationImportJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAnnotationImportJobInput(v *GetAnnotationImportJobInput, encoder *httpbinding.Encoder) error {
@@ -2971,10 +2874,8 @@ func (m *awsRestjson1_serializeOpGetAnnotationStore) H
 func (m *awsRestjson1_serializeOpGetAnnotationStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3012,7 +2913,6 @@ func (m *awsRestjson1_serializeOpGetAnnotationStore) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAnnotationStoreInput(v *GetAnnotationStoreInput, encoder *httpbinding.Encoder) error {
@@ -3042,10 +2942,8 @@ func (m *awsRestjson1_serializeOpGetAnnotationStoreVer
 func (m *awsRestjson1_serializeOpGetAnnotationStoreVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3083,7 +2981,6 @@ func (m *awsRestjson1_serializeOpGetAnnotationStoreVer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAnnotationStoreVersionInput(v *GetAnnotationStoreVersionInput, encoder *httpbinding.Encoder) error {
@@ -3122,10 +3019,8 @@ func (m *awsRestjson1_serializeOpGetReadSet) HandleSer
 func (m *awsRestjson1_serializeOpGetReadSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3163,7 +3058,6 @@ func (m *awsRestjson1_serializeOpGetReadSet) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReadSetInput(v *GetReadSetInput, encoder *httpbinding.Encoder) error {
@@ -3210,10 +3104,8 @@ func (m *awsRestjson1_serializeOpGetReadSetActivationJ
 func (m *awsRestjson1_serializeOpGetReadSetActivationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3251,7 +3143,6 @@ func (m *awsRestjson1_serializeOpGetReadSetActivationJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReadSetActivationJobInput(v *GetReadSetActivationJobInput, encoder *httpbinding.Encoder) error {
@@ -3290,10 +3181,8 @@ func (m *awsRestjson1_serializeOpGetReadSetExportJob) 
 func (m *awsRestjson1_serializeOpGetReadSetExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3331,7 +3220,6 @@ func (m *awsRestjson1_serializeOpGetReadSetExportJob) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReadSetExportJobInput(v *GetReadSetExportJobInput, encoder *httpbinding.Encoder) error {
@@ -3370,10 +3258,8 @@ func (m *awsRestjson1_serializeOpGetReadSetImportJob) 
 func (m *awsRestjson1_serializeOpGetReadSetImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3411,7 +3297,6 @@ func (m *awsRestjson1_serializeOpGetReadSetImportJob) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReadSetImportJobInput(v *GetReadSetImportJobInput, encoder *httpbinding.Encoder) error {
@@ -3450,10 +3335,8 @@ func (m *awsRestjson1_serializeOpGetReadSetMetadata) H
 func (m *awsRestjson1_serializeOpGetReadSetMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3491,7 +3374,6 @@ func (m *awsRestjson1_serializeOpGetReadSetMetadata) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReadSetMetadataInput(v *GetReadSetMetadataInput, encoder *httpbinding.Encoder) error {
@@ -3530,10 +3412,8 @@ func (m *awsRestjson1_serializeOpGetReference) HandleS
 func (m *awsRestjson1_serializeOpGetReference) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3571,7 +3451,6 @@ func (m *awsRestjson1_serializeOpGetReference) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReferenceInput(v *GetReferenceInput, encoder *httpbinding.Encoder) error {
@@ -3623,10 +3502,8 @@ func (m *awsRestjson1_serializeOpGetReferenceImportJob
 func (m *awsRestjson1_serializeOpGetReferenceImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3664,7 +3541,6 @@ func (m *awsRestjson1_serializeOpGetReferenceImportJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReferenceImportJobInput(v *GetReferenceImportJobInput, encoder *httpbinding.Encoder) error {
@@ -3703,10 +3579,8 @@ func (m *awsRestjson1_serializeOpGetReferenceMetadata)
 func (m *awsRestjson1_serializeOpGetReferenceMetadata) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3744,7 +3618,6 @@ func (m *awsRestjson1_serializeOpGetReferenceMetadata)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReferenceMetadataInput(v *GetReferenceMetadataInput, encoder *httpbinding.Encoder) error {
@@ -3783,10 +3656,8 @@ func (m *awsRestjson1_serializeOpGetReferenceStore) Ha
 func (m *awsRestjson1_serializeOpGetReferenceStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3824,7 +3695,6 @@ func (m *awsRestjson1_serializeOpGetReferenceStore) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReferenceStoreInput(v *GetReferenceStoreInput, encoder *httpbinding.Encoder) error {
@@ -3854,10 +3724,8 @@ func (m *awsRestjson1_serializeOpGetRun) HandleSeriali
 func (m *awsRestjson1_serializeOpGetRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3895,7 +3763,6 @@ func (m *awsRestjson1_serializeOpGetRun) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRunInput(v *GetRunInput, encoder *httpbinding.Encoder) error {
@@ -3931,10 +3798,8 @@ func (m *awsRestjson1_serializeOpGetRunCache) HandleSe
 func (m *awsRestjson1_serializeOpGetRunCache) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3972,7 +3837,6 @@ func (m *awsRestjson1_serializeOpGetRunCache) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRunCacheInput(v *GetRunCacheInput, encoder *httpbinding.Encoder) error {
@@ -4002,10 +3866,8 @@ func (m *awsRestjson1_serializeOpGetRunGroup) HandleSe
 func (m *awsRestjson1_serializeOpGetRunGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4043,7 +3905,6 @@ func (m *awsRestjson1_serializeOpGetRunGroup) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRunGroupInput(v *GetRunGroupInput, encoder *httpbinding.Encoder) error {
@@ -4073,10 +3934,8 @@ func (m *awsRestjson1_serializeOpGetRunTask) HandleSer
 func (m *awsRestjson1_serializeOpGetRunTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4114,7 +3973,6 @@ func (m *awsRestjson1_serializeOpGetRunTask) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRunTaskInput(v *GetRunTaskInput, encoder *httpbinding.Encoder) error {
@@ -4153,10 +4011,8 @@ func (m *awsRestjson1_serializeOpGetS3AccessPolicy) Ha
 func (m *awsRestjson1_serializeOpGetS3AccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4194,7 +4050,6 @@ func (m *awsRestjson1_serializeOpGetS3AccessPolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetS3AccessPolicyInput(v *GetS3AccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -4224,10 +4079,8 @@ func (m *awsRestjson1_serializeOpGetSequenceStore) Han
 func (m *awsRestjson1_serializeOpGetSequenceStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4265,7 +4118,6 @@ func (m *awsRestjson1_serializeOpGetSequenceStore) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSequenceStoreInput(v *GetSequenceStoreInput, encoder *httpbinding.Encoder) error {
@@ -4295,10 +4147,8 @@ func (m *awsRestjson1_serializeOpGetShare) HandleSeria
 func (m *awsRestjson1_serializeOpGetShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4336,7 +4186,6 @@ func (m *awsRestjson1_serializeOpGetShare) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetShareInput(v *GetShareInput, encoder *httpbinding.Encoder) error {
@@ -4366,10 +4215,8 @@ func (m *awsRestjson1_serializeOpGetVariantImportJob) 
 func (m *awsRestjson1_serializeOpGetVariantImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4407,7 +4254,6 @@ func (m *awsRestjson1_serializeOpGetVariantImportJob) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVariantImportJobInput(v *GetVariantImportJobInput, encoder *httpbinding.Encoder) error {
@@ -4437,10 +4283,8 @@ func (m *awsRestjson1_serializeOpGetVariantStore) Hand
 func (m *awsRestjson1_serializeOpGetVariantStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4478,7 +4322,6 @@ func (m *awsRestjson1_serializeOpGetVariantStore) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetVariantStoreInput(v *GetVariantStoreInput, encoder *httpbinding.Encoder) error {
@@ -4508,10 +4351,8 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
 func (m *awsRestjson1_serializeOpGetWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4549,7 +4390,6 @@ func (m *awsRestjson1_serializeOpGetWorkflow) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowInput(v *GetWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -4593,10 +4433,8 @@ func (m *awsRestjson1_serializeOpGetWorkflowVersion) H
 func (m *awsRestjson1_serializeOpGetWorkflowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4634,7 +4472,6 @@ func (m *awsRestjson1_serializeOpGetWorkflowVersion) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkflowVersionInput(v *GetWorkflowVersionInput, encoder *httpbinding.Encoder) error {
@@ -4687,10 +4524,8 @@ func (m *awsRestjson1_serializeOpListAnnotationImportJ
 func (m *awsRestjson1_serializeOpListAnnotationImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4739,7 +4574,6 @@ func (m *awsRestjson1_serializeOpListAnnotationImportJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnnotationImportJobsInput(v *ListAnnotationImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -4789,10 +4623,8 @@ func (m *awsRestjson1_serializeOpListAnnotationStores)
 func (m *awsRestjson1_serializeOpListAnnotationStores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4841,7 +4673,6 @@ func (m *awsRestjson1_serializeOpListAnnotationStores)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnnotationStoresInput(v *ListAnnotationStoresInput, encoder *httpbinding.Encoder) error {
@@ -4891,10 +4722,8 @@ func (m *awsRestjson1_serializeOpListAnnotationStoreVe
 func (m *awsRestjson1_serializeOpListAnnotationStoreVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4943,7 +4772,6 @@ func (m *awsRestjson1_serializeOpListAnnotationStoreVe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAnnotationStoreVersionsInput(v *ListAnnotationStoreVersionsInput, encoder *httpbinding.Encoder) error {
@@ -4995,10 +4823,8 @@ func (m *awsRestjson1_serializeOpListMultipartReadSetU
 func (m *awsRestjson1_serializeOpListMultipartReadSetUploads) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5036,7 +4862,6 @@ func (m *awsRestjson1_serializeOpListMultipartReadSetU
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMultipartReadSetUploadsInput(v *ListMultipartReadSetUploadsInput, encoder *httpbinding.Encoder) error {
@@ -5074,10 +4899,8 @@ func (m *awsRestjson1_serializeOpListReadSetActivation
 func (m *awsRestjson1_serializeOpListReadSetActivationJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5126,7 +4949,6 @@ func (m *awsRestjson1_serializeOpListReadSetActivation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReadSetActivationJobsInput(v *ListReadSetActivationJobsInput, encoder *httpbinding.Encoder) error {
@@ -5178,10 +5000,8 @@ func (m *awsRestjson1_serializeOpListReadSetExportJobs
 func (m *awsRestjson1_serializeOpListReadSetExportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5230,7 +5050,6 @@ func (m *awsRestjson1_serializeOpListReadSetExportJobs
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReadSetExportJobsInput(v *ListReadSetExportJobsInput, encoder *httpbinding.Encoder) error {
@@ -5282,10 +5101,8 @@ func (m *awsRestjson1_serializeOpListReadSetImportJobs
 func (m *awsRestjson1_serializeOpListReadSetImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5334,7 +5151,6 @@ func (m *awsRestjson1_serializeOpListReadSetImportJobs
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReadSetImportJobsInput(v *ListReadSetImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -5386,10 +5202,8 @@ func (m *awsRestjson1_serializeOpListReadSets) HandleS
 func (m *awsRestjson1_serializeOpListReadSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5438,7 +5252,6 @@ func (m *awsRestjson1_serializeOpListReadSets) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReadSetsInput(v *ListReadSetsInput, encoder *httpbinding.Encoder) error {
@@ -5490,10 +5303,8 @@ func (m *awsRestjson1_serializeOpListReadSetUploadPart
 func (m *awsRestjson1_serializeOpListReadSetUploadParts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5542,7 +5353,6 @@ func (m *awsRestjson1_serializeOpListReadSetUploadPart
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReadSetUploadPartsInput(v *ListReadSetUploadPartsInput, encoder *httpbinding.Encoder) error {
@@ -5608,10 +5418,8 @@ func (m *awsRestjson1_serializeOpListReferenceImportJo
 func (m *awsRestjson1_serializeOpListReferenceImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5660,7 +5468,6 @@ func (m *awsRestjson1_serializeOpListReferenceImportJo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReferenceImportJobsInput(v *ListReferenceImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -5712,10 +5519,8 @@ func (m *awsRestjson1_serializeOpListReferences) Handl
 func (m *awsRestjson1_serializeOpListReferences) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5764,7 +5569,6 @@ func (m *awsRestjson1_serializeOpListReferences) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReferencesInput(v *ListReferencesInput, encoder *httpbinding.Encoder) error {
@@ -5816,10 +5620,8 @@ func (m *awsRestjson1_serializeOpListReferenceStores) 
 func (m *awsRestjson1_serializeOpListReferenceStores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5868,7 +5670,6 @@ func (m *awsRestjson1_serializeOpListReferenceStores) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReferenceStoresInput(v *ListReferenceStoresInput, encoder *httpbinding.Encoder) error {
@@ -5911,10 +5712,8 @@ func (m *awsRestjson1_serializeOpListRunCaches) Handle
 func (m *awsRestjson1_serializeOpListRunCaches) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5952,7 +5751,6 @@ func (m *awsRestjson1_serializeOpListRunCaches) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRunCachesInput(v *ListRunCachesInput, encoder *httpbinding.Encoder) error {
@@ -5981,10 +5779,8 @@ func (m *awsRestjson1_serializeOpListRunGroups) Handle
 func (m *awsRestjson1_serializeOpListRunGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6022,7 +5818,6 @@ func (m *awsRestjson1_serializeOpListRunGroups) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRunGroupsInput(v *ListRunGroupsInput, encoder *httpbinding.Encoder) error {
@@ -6055,10 +5850,8 @@ func (m *awsRestjson1_serializeOpListRuns) HandleSeria
 func (m *awsRestjson1_serializeOpListRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6096,7 +5889,6 @@ func (m *awsRestjson1_serializeOpListRuns) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRunsInput(v *ListRunsInput, encoder *httpbinding.Encoder) error {
@@ -6137,10 +5929,8 @@ func (m *awsRestjson1_serializeOpListRunTasks) HandleS
 func (m *awsRestjson1_serializeOpListRunTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6178,7 +5968,6 @@ func (m *awsRestjson1_serializeOpListRunTasks) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRunTasksInput(v *ListRunTasksInput, encoder *httpbinding.Encoder) error {
@@ -6220,10 +6009,8 @@ func (m *awsRestjson1_serializeOpListSequenceStores) H
 func (m *awsRestjson1_serializeOpListSequenceStores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6272,7 +6059,6 @@ func (m *awsRestjson1_serializeOpListSequenceStores) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSequenceStoresInput(v *ListSequenceStoresInput, encoder *httpbinding.Encoder) error {
@@ -6315,10 +6101,8 @@ func (m *awsRestjson1_serializeOpListShares) HandleSer
 func (m *awsRestjson1_serializeOpListShares) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6367,7 +6151,6 @@ func (m *awsRestjson1_serializeOpListShares) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSharesInput(v *ListSharesInput, encoder *httpbinding.Encoder) error {
@@ -6415,10 +6198,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -6456,7 +6237,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -6486,10 +6266,8 @@ func (m *awsRestjson1_serializeOpListVariantImportJobs
 func (m *awsRestjson1_serializeOpListVariantImportJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6538,7 +6316,6 @@ func (m *awsRestjson1_serializeOpListVariantImportJobs
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVariantImportJobsInput(v *ListVariantImportJobsInput, encoder *httpbinding.Encoder) error {
@@ -6588,10 +6365,8 @@ func (m *awsRestjson1_serializeOpListVariantStores) Ha
 func (m *awsRestjson1_serializeOpListVariantStores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6640,7 +6415,6 @@ func (m *awsRestjson1_serializeOpListVariantStores) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListVariantStoresInput(v *ListVariantStoresInput, encoder *httpbinding.Encoder) error {
@@ -6690,10 +6464,8 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
 func (m *awsRestjson1_serializeOpListWorkflows) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6731,7 +6503,6 @@ func (m *awsRestjson1_serializeOpListWorkflows) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowsInput(v *ListWorkflowsInput, encoder *httpbinding.Encoder) error {
@@ -6768,10 +6539,8 @@ func (m *awsRestjson1_serializeOpListWorkflowVersions)
 func (m *awsRestjson1_serializeOpListWorkflowVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6809,7 +6578,6 @@ func (m *awsRestjson1_serializeOpListWorkflowVersions)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkflowVersionsInput(v *ListWorkflowVersionsInput, encoder *httpbinding.Encoder) error {
@@ -6855,10 +6623,8 @@ func (m *awsRestjson1_serializeOpPutS3AccessPolicy) Ha
 func (m *awsRestjson1_serializeOpPutS3AccessPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6907,7 +6673,6 @@ func (m *awsRestjson1_serializeOpPutS3AccessPolicy) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutS3AccessPolicyInput(v *PutS3AccessPolicyInput, encoder *httpbinding.Encoder) error {
@@ -6949,10 +6714,8 @@ func (m *awsRestjson1_serializeOpStartAnnotationImport
 func (m *awsRestjson1_serializeOpStartAnnotationImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6997,7 +6760,6 @@ func (m *awsRestjson1_serializeOpStartAnnotationImport
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAnnotationImportJobInput(v *StartAnnotationImportJobInput, encoder *httpbinding.Encoder) error {
@@ -7066,10 +6828,8 @@ func (m *awsRestjson1_serializeOpStartReadSetActivatio
 func (m *awsRestjson1_serializeOpStartReadSetActivationJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7118,7 +6878,6 @@ func (m *awsRestjson1_serializeOpStartReadSetActivatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartReadSetActivationJobInput(v *StartReadSetActivationJobInput, encoder *httpbinding.Encoder) error {
@@ -7167,10 +6926,8 @@ func (m *awsRestjson1_serializeOpStartReadSetExportJob
 func (m *awsRestjson1_serializeOpStartReadSetExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7219,7 +6976,6 @@ func (m *awsRestjson1_serializeOpStartReadSetExportJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartReadSetExportJobInput(v *StartReadSetExportJobInput, encoder *httpbinding.Encoder) error {
@@ -7278,10 +7034,8 @@ func (m *awsRestjson1_serializeOpStartReadSetImportJob
 func (m *awsRestjson1_serializeOpStartReadSetImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7330,7 +7084,6 @@ func (m *awsRestjson1_serializeOpStartReadSetImportJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartReadSetImportJobInput(v *StartReadSetImportJobInput, encoder *httpbinding.Encoder) error {
@@ -7384,10 +7137,8 @@ func (m *awsRestjson1_serializeOpStartReferenceImportJ
 func (m *awsRestjson1_serializeOpStartReferenceImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7436,7 +7187,6 @@ func (m *awsRestjson1_serializeOpStartReferenceImportJ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartReferenceImportJobInput(v *StartReferenceImportJobInput, encoder *httpbinding.Encoder) error {
@@ -7490,10 +7240,8 @@ func (m *awsRestjson1_serializeOpStartRun) HandleSeria
 func (m *awsRestjson1_serializeOpStartRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7538,7 +7286,6 @@ func (m *awsRestjson1_serializeOpStartRun) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartRunInput(v *StartRunInput, encoder *httpbinding.Encoder) error {
@@ -7665,10 +7412,8 @@ func (m *awsRestjson1_serializeOpStartVariantImportJob
 func (m *awsRestjson1_serializeOpStartVariantImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7713,7 +7458,6 @@ func (m *awsRestjson1_serializeOpStartVariantImportJob
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartVariantImportJobInput(v *StartVariantImportJobInput, encoder *httpbinding.Encoder) error {
@@ -7770,10 +7514,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -7822,7 +7564,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7866,10 +7607,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -7907,7 +7646,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -7943,10 +7681,8 @@ func (m *awsRestjson1_serializeOpUpdateAnnotationStore
 func (m *awsRestjson1_serializeOpUpdateAnnotationStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7995,7 +7731,6 @@ func (m *awsRestjson1_serializeOpUpdateAnnotationStore
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAnnotationStoreInput(v *UpdateAnnotationStoreInput, encoder *httpbinding.Encoder) error {
@@ -8037,10 +7772,8 @@ func (m *awsRestjson1_serializeOpUpdateAnnotationStore
 func (m *awsRestjson1_serializeOpUpdateAnnotationStoreVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8089,7 +7822,6 @@ func (m *awsRestjson1_serializeOpUpdateAnnotationStore
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAnnotationStoreVersionInput(v *UpdateAnnotationStoreVersionInput, encoder *httpbinding.Encoder) error {
@@ -8140,10 +7872,8 @@ func (m *awsRestjson1_serializeOpUpdateRunCache) Handl
 func (m *awsRestjson1_serializeOpUpdateRunCache) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8192,7 +7922,6 @@ func (m *awsRestjson1_serializeOpUpdateRunCache) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRunCacheInput(v *UpdateRunCacheInput, encoder *httpbinding.Encoder) error {
@@ -8244,10 +7973,8 @@ func (m *awsRestjson1_serializeOpUpdateRunGroup) Handl
 func (m *awsRestjson1_serializeOpUpdateRunGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8296,7 +8023,6 @@ func (m *awsRestjson1_serializeOpUpdateRunGroup) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRunGroupInput(v *UpdateRunGroupInput, encoder *httpbinding.Encoder) error {
@@ -8358,10 +8084,8 @@ func (m *awsRestjson1_serializeOpUpdateSequenceStore) 
 func (m *awsRestjson1_serializeOpUpdateSequenceStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8410,7 +8134,6 @@ func (m *awsRestjson1_serializeOpUpdateSequenceStore) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSequenceStoreInput(v *UpdateSequenceStoreInput, encoder *httpbinding.Encoder) error {
@@ -8481,10 +8204,8 @@ func (m *awsRestjson1_serializeOpUpdateVariantStore) H
 func (m *awsRestjson1_serializeOpUpdateVariantStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8533,7 +8254,6 @@ func (m *awsRestjson1_serializeOpUpdateVariantStore) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateVariantStoreInput(v *UpdateVariantStoreInput, encoder *httpbinding.Encoder) error {
@@ -8575,10 +8295,8 @@ func (m *awsRestjson1_serializeOpUpdateWorkflow) Handl
 func (m *awsRestjson1_serializeOpUpdateWorkflow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8627,7 +8345,6 @@ func (m *awsRestjson1_serializeOpUpdateWorkflow) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWorkflowInput(v *UpdateWorkflowInput, encoder *httpbinding.Encoder) error {
@@ -8684,10 +8401,8 @@ func (m *awsRestjson1_serializeOpUpdateWorkflowVersion
 func (m *awsRestjson1_serializeOpUpdateWorkflowVersion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8736,7 +8451,6 @@ func (m *awsRestjson1_serializeOpUpdateWorkflowVersion
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWorkflowVersionInput(v *UpdateWorkflowVersionInput, encoder *httpbinding.Encoder) error {
@@ -8797,10 +8511,8 @@ func (m *awsRestjson1_serializeOpUploadReadSetPart) Ha
 func (m *awsRestjson1_serializeOpUploadReadSetPart) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8850,7 +8562,6 @@ func (m *awsRestjson1_serializeOpUploadReadSetPart) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUploadReadSetPartInput(v *UploadReadSetPartInput, encoder *httpbinding.Encoder) error {
