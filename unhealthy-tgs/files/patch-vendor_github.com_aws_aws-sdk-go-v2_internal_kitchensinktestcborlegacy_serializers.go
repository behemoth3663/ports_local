--- vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestcborlegacy/serializers.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestcborlegacy/serializers.go
@@ -9,7 +9,6 @@ import (
 	smithy "github.com/aws/smithy-go"
 	smithycbor "github.com/aws/smithy-go/encoding/cbor"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -24,10 +23,8 @@ func (m *smithyRpcv2cbor_serializeOpCborGetItem) Handl
 func (m *smithyRpcv2cbor_serializeOpCborGetItem) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*CborGetItemInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -59,7 +56,6 @@ func (m *smithyRpcv2cbor_serializeOpCborGetItem) Handl
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -74,10 +70,8 @@ func (m *smithyRpcv2cbor_serializeOpCborGetLists) Hand
 func (m *smithyRpcv2cbor_serializeOpCborGetLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*CborGetListsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -109,7 +103,6 @@ func (m *smithyRpcv2cbor_serializeOpCborGetLists) Hand
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -124,10 +117,8 @@ func (m *smithyRpcv2cbor_serializeOpCborGetMaps) Handl
 func (m *smithyRpcv2cbor_serializeOpCborGetMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*CborGetMapsInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -159,7 +150,6 @@ func (m *smithyRpcv2cbor_serializeOpCborGetMaps) Handl
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -174,10 +164,8 @@ func (m *smithyRpcv2cbor_serializeOpCborPutCompressedD
 func (m *smithyRpcv2cbor_serializeOpCborPutCompressedData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*CborPutCompressedDataInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -209,7 +197,6 @@ func (m *smithyRpcv2cbor_serializeOpCborPutCompressedD
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
