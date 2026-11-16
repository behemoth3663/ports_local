--- vendor/github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpDeleteReportDefinitio
 func (m *awsRestjson1_serializeOpDeleteReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpDeleteReportDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteReportDefinitionInput(v *DeleteReportDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -96,10 +92,8 @@ func (m *awsRestjson1_serializeOpGetReportDefinition) 
 func (m *awsRestjson1_serializeOpGetReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -137,7 +131,6 @@ func (m *awsRestjson1_serializeOpGetReportDefinition) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReportDefinitionInput(v *GetReportDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -167,10 +160,8 @@ func (m *awsRestjson1_serializeOpImportApplicationUsag
 func (m *awsRestjson1_serializeOpImportApplicationUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -215,7 +206,6 @@ func (m *awsRestjson1_serializeOpImportApplicationUsag
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsImportApplicationUsageInput(v *ImportApplicationUsageInput, encoder *httpbinding.Encoder) error {
@@ -250,10 +240,8 @@ func (m *awsRestjson1_serializeOpListReportDefinitions
 func (m *awsRestjson1_serializeOpListReportDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -291,7 +279,6 @@ func (m *awsRestjson1_serializeOpListReportDefinitions
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListReportDefinitionsInput(v *ListReportDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -320,10 +307,8 @@ func (m *awsRestjson1_serializeOpPutReportDefinition) 
 func (m *awsRestjson1_serializeOpPutReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -368,7 +353,6 @@ func (m *awsRestjson1_serializeOpPutReportDefinition) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutReportDefinitionInput(v *PutReportDefinitionInput, encoder *httpbinding.Encoder) error {
@@ -423,10 +407,8 @@ func (m *awsRestjson1_serializeOpUpdateReportDefinitio
 func (m *awsRestjson1_serializeOpUpdateReportDefinition) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -475,7 +457,6 @@ func (m *awsRestjson1_serializeOpUpdateReportDefinitio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateReportDefinitionInput(v *UpdateReportDefinitionInput, encoder *httpbinding.Encoder) error {
