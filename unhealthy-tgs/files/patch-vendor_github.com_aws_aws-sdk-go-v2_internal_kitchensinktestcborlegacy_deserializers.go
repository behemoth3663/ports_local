--- vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestcborlegacy/deserializers.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/kitchensinktestcborlegacy/deserializers.go
@@ -10,7 +10,6 @@ import (
 	smithycbor "github.com/aws/smithy-go/encoding/cbor"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -48,10 +47,8 @@ func (m *smithyRpcv2cbor_deserializeOpCborGetItem) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
 		return out, metadata, &smithy.DeserializationError{
@@ -108,10 +105,8 @@ func (m *smithyRpcv2cbor_deserializeOpCborGetLists) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
 		return out, metadata, &smithy.DeserializationError{
@@ -184,10 +179,8 @@ func (m *smithyRpcv2cbor_deserializeOpCborGetMaps) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
 		return out, metadata, &smithy.DeserializationError{
@@ -260,10 +253,8 @@ func (m *smithyRpcv2cbor_deserializeOpCborPutCompresse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
 		return out, metadata, &smithy.DeserializationError{
