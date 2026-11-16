--- vendor/github.com/aws/aws-sdk-go-v2/service/observabilityadmin/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/observabilityadmin/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpGetTelemetryEvaluatio
 func (m *awsRestjson1_serializeOpGetTelemetryEvaluationStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -62,7 +59,6 @@ func (m *awsRestjson1_serializeOpGetTelemetryEvaluatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTelemetryEvaluationStatusInput(v *GetTelemetryEvaluationStatusInput, encoder *httpbinding.Encoder) error {
@@ -83,10 +79,8 @@ func (m *awsRestjson1_serializeOpGetTelemetryEvaluatio
 func (m *awsRestjson1_serializeOpGetTelemetryEvaluationStatusForOrganization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -120,7 +114,6 @@ func (m *awsRestjson1_serializeOpGetTelemetryEvaluatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTelemetryEvaluationStatusForOrganizationInput(v *GetTelemetryEvaluationStatusForOrganizationInput, encoder *httpbinding.Encoder) error {
@@ -141,10 +134,8 @@ func (m *awsRestjson1_serializeOpListResourceTelemetry
 func (m *awsRestjson1_serializeOpListResourceTelemetry) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -189,7 +180,6 @@ func (m *awsRestjson1_serializeOpListResourceTelemetry
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceTelemetryInput(v *ListResourceTelemetryInput, encoder *httpbinding.Encoder) error {
@@ -253,10 +243,8 @@ func (m *awsRestjson1_serializeOpListResourceTelemetry
 func (m *awsRestjson1_serializeOpListResourceTelemetryForOrganization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -301,7 +289,6 @@ func (m *awsRestjson1_serializeOpListResourceTelemetry
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListResourceTelemetryForOrganizationInput(v *ListResourceTelemetryForOrganizationInput, encoder *httpbinding.Encoder) error {
@@ -372,10 +359,8 @@ func (m *awsRestjson1_serializeOpStartTelemetryEvaluat
 func (m *awsRestjson1_serializeOpStartTelemetryEvaluation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -409,7 +394,6 @@ func (m *awsRestjson1_serializeOpStartTelemetryEvaluat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTelemetryEvaluationInput(v *StartTelemetryEvaluationInput, encoder *httpbinding.Encoder) error {
@@ -430,10 +414,8 @@ func (m *awsRestjson1_serializeOpStartTelemetryEvaluat
 func (m *awsRestjson1_serializeOpStartTelemetryEvaluationForOrganization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -467,7 +449,6 @@ func (m *awsRestjson1_serializeOpStartTelemetryEvaluat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartTelemetryEvaluationForOrganizationInput(v *StartTelemetryEvaluationForOrganizationInput, encoder *httpbinding.Encoder) error {
@@ -488,10 +469,8 @@ func (m *awsRestjson1_serializeOpStopTelemetryEvaluati
 func (m *awsRestjson1_serializeOpStopTelemetryEvaluation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -525,7 +504,6 @@ func (m *awsRestjson1_serializeOpStopTelemetryEvaluati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopTelemetryEvaluationInput(v *StopTelemetryEvaluationInput, encoder *httpbinding.Encoder) error {
@@ -546,10 +524,8 @@ func (m *awsRestjson1_serializeOpStopTelemetryEvaluati
 func (m *awsRestjson1_serializeOpStopTelemetryEvaluationForOrganization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -583,7 +559,6 @@ func (m *awsRestjson1_serializeOpStopTelemetryEvaluati
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopTelemetryEvaluationForOrganizationInput(v *StopTelemetryEvaluationForOrganizationInput, encoder *httpbinding.Encoder) error {
