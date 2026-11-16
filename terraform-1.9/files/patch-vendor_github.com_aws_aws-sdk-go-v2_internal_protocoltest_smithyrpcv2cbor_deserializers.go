--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cbor/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cbor/deserializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -39,10 +38,8 @@ func (m *smithyRpcv2cbor_deserializeOpEmptyInputOutput
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -88,10 +85,8 @@ func (m *smithyRpcv2cbor_deserializeOpFloat16) HandleD
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -153,10 +148,8 @@ func (m *smithyRpcv2cbor_deserializeOpFractionalSecond
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -218,10 +211,8 @@ func (m *smithyRpcv2cbor_deserializeOpGreetingWithErro
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -283,10 +274,8 @@ func (m *smithyRpcv2cbor_deserializeOpNoInputOutput) H
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -332,10 +321,8 @@ func (m *smithyRpcv2cbor_deserializeOpOperationWithDef
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -397,10 +384,8 @@ func (m *smithyRpcv2cbor_deserializeOpOptionalInputOut
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -462,10 +447,8 @@ func (m *smithyRpcv2cbor_deserializeOpRecursiveShapes)
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -527,10 +510,8 @@ func (m *smithyRpcv2cbor_deserializeOpRpcV2CborDenseMa
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -592,10 +573,8 @@ func (m *smithyRpcv2cbor_deserializeOpRpcV2CborLists) 
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -657,10 +636,8 @@ func (m *smithyRpcv2cbor_deserializeOpRpcV2CborSparseM
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -722,10 +699,8 @@ func (m *smithyRpcv2cbor_deserializeOpSimpleScalarProp
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
@@ -787,10 +762,8 @@ func (m *smithyRpcv2cbor_deserializeOpSparseNullsOpera
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
