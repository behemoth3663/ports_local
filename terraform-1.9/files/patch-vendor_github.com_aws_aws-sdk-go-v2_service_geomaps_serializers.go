--- vendor/github.com/aws/aws-sdk-go-v2/service/geomaps/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/geomaps/serializers.go
@@ -8,7 +8,6 @@ import (
 	smithy "github.com/aws/smithy-go"
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -22,10 +21,8 @@ func (m *awsRestjson1_serializeOpGetGlyphs) HandleSeri
 func (m *awsRestjson1_serializeOpGetGlyphs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -63,7 +60,6 @@ func (m *awsRestjson1_serializeOpGetGlyphs) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGlyphsInput(v *GetGlyphsInput, encoder *httpbinding.Encoder) error {
@@ -102,10 +98,8 @@ func (m *awsRestjson1_serializeOpGetSprites) HandleSer
 func (m *awsRestjson1_serializeOpGetSprites) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -143,7 +137,6 @@ func (m *awsRestjson1_serializeOpGetSprites) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSpritesInput(v *GetSpritesInput, encoder *httpbinding.Encoder) error {
@@ -200,10 +193,8 @@ func (m *awsRestjson1_serializeOpGetStaticMap) HandleS
 func (m *awsRestjson1_serializeOpGetStaticMap) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -241,7 +232,6 @@ func (m *awsRestjson1_serializeOpGetStaticMap) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStaticMapInput(v *GetStaticMapInput, encoder *httpbinding.Encoder) error {
@@ -347,10 +337,8 @@ func (m *awsRestjson1_serializeOpGetStyleDescriptor) H
 func (m *awsRestjson1_serializeOpGetStyleDescriptor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -388,7 +376,6 @@ func (m *awsRestjson1_serializeOpGetStyleDescriptor) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStyleDescriptorInput(v *GetStyleDescriptorInput, encoder *httpbinding.Encoder) error {
@@ -430,10 +417,8 @@ func (m *awsRestjson1_serializeOpGetTile) HandleSerial
 func (m *awsRestjson1_serializeOpGetTile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -471,7 +456,6 @@ func (m *awsRestjson1_serializeOpGetTile) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTileInput(v *GetTileInput, encoder *httpbinding.Encoder) error {
