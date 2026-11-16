--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cbor/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cbor/serializers.go
@@ -10,7 +10,6 @@ import (
 	smithy "github.com/aws/smithy-go"
 	smithycbor "github.com/aws/smithy-go/encoding/cbor"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 	"time"
@@ -26,10 +25,8 @@ func (m *smithyRpcv2cbor_serializeOpEmptyInputOutput) 
 func (m *smithyRpcv2cbor_serializeOpEmptyInputOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*EmptyInputOutputInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -61,7 +58,6 @@ func (m *smithyRpcv2cbor_serializeOpEmptyInputOutput) 
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -76,10 +72,8 @@ func (m *smithyRpcv2cbor_serializeOpFloat16) HandleSer
 func (m *smithyRpcv2cbor_serializeOpFloat16) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*Float16Input)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -98,7 +92,6 @@ func (m *smithyRpcv2cbor_serializeOpFloat16) HandleSer
 	req.Header.Set("Accept", "application/cbor")
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -113,10 +106,8 @@ func (m *smithyRpcv2cbor_serializeOpFractionalSeconds)
 func (m *smithyRpcv2cbor_serializeOpFractionalSeconds) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*FractionalSecondsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -135,7 +126,6 @@ func (m *smithyRpcv2cbor_serializeOpFractionalSeconds)
 	req.Header.Set("Accept", "application/cbor")
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -150,10 +140,8 @@ func (m *smithyRpcv2cbor_serializeOpGreetingWithErrors
 func (m *smithyRpcv2cbor_serializeOpGreetingWithErrors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*GreetingWithErrorsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -172,7 +160,6 @@ func (m *smithyRpcv2cbor_serializeOpGreetingWithErrors
 	req.Header.Set("Accept", "application/cbor")
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -187,10 +174,8 @@ func (m *smithyRpcv2cbor_serializeOpNoInputOutput) Han
 func (m *smithyRpcv2cbor_serializeOpNoInputOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*NoInputOutputInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -209,7 +194,6 @@ func (m *smithyRpcv2cbor_serializeOpNoInputOutput) Han
 	req.Header.Set("Accept", "application/cbor")
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -224,10 +208,8 @@ func (m *smithyRpcv2cbor_serializeOpOperationWithDefau
 func (m *smithyRpcv2cbor_serializeOpOperationWithDefaults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*OperationWithDefaultsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -259,7 +241,6 @@ func (m *smithyRpcv2cbor_serializeOpOperationWithDefau
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -274,10 +255,8 @@ func (m *smithyRpcv2cbor_serializeOpOptionalInputOutpu
 func (m *smithyRpcv2cbor_serializeOpOptionalInputOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*OptionalInputOutputInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -309,7 +288,6 @@ func (m *smithyRpcv2cbor_serializeOpOptionalInputOutpu
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -324,10 +302,8 @@ func (m *smithyRpcv2cbor_serializeOpRecursiveShapes) H
 func (m *smithyRpcv2cbor_serializeOpRecursiveShapes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*RecursiveShapesInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -359,7 +335,6 @@ func (m *smithyRpcv2cbor_serializeOpRecursiveShapes) H
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -374,10 +349,8 @@ func (m *smithyRpcv2cbor_serializeOpRpcV2CborDenseMaps
 func (m *smithyRpcv2cbor_serializeOpRpcV2CborDenseMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*RpcV2CborDenseMapsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -409,7 +382,6 @@ func (m *smithyRpcv2cbor_serializeOpRpcV2CborDenseMaps
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -424,10 +396,8 @@ func (m *smithyRpcv2cbor_serializeOpRpcV2CborLists) Ha
 func (m *smithyRpcv2cbor_serializeOpRpcV2CborLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*RpcV2CborListsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -459,7 +429,6 @@ func (m *smithyRpcv2cbor_serializeOpRpcV2CborLists) Ha
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -474,10 +443,8 @@ func (m *smithyRpcv2cbor_serializeOpRpcV2CborSparseMap
 func (m *smithyRpcv2cbor_serializeOpRpcV2CborSparseMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*RpcV2CborSparseMapsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -509,7 +476,6 @@ func (m *smithyRpcv2cbor_serializeOpRpcV2CborSparseMap
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -524,10 +490,8 @@ func (m *smithyRpcv2cbor_serializeOpSimpleScalarProper
 func (m *smithyRpcv2cbor_serializeOpSimpleScalarProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*SimpleScalarPropertiesInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -559,7 +523,6 @@ func (m *smithyRpcv2cbor_serializeOpSimpleScalarProper
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -574,10 +537,8 @@ func (m *smithyRpcv2cbor_serializeOpSparseNullsOperati
 func (m *smithyRpcv2cbor_serializeOpSparseNullsOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*SparseNullsOperationInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -609,7 +570,6 @@ func (m *smithyRpcv2cbor_serializeOpSparseNullsOperati
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
