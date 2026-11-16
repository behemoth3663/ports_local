--- vendor/github.com/aws/aws-sdk-go-v2/service/georoutes/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/georoutes/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCalculateIsolines) Ha
 func (m *awsRestjson1_serializeOpCalculateIsolines) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -78,7 +75,6 @@ func (m *awsRestjson1_serializeOpCalculateIsolines) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCalculateIsolinesInput(v *CalculateIsolinesInput, encoder *httpbinding.Encoder) error {
@@ -215,10 +211,8 @@ func (m *awsRestjson1_serializeOpCalculateRouteMatrix)
 func (m *awsRestjson1_serializeOpCalculateRouteMatrix) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -267,7 +261,6 @@ func (m *awsRestjson1_serializeOpCalculateRouteMatrix)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCalculateRouteMatrixInput(v *CalculateRouteMatrixInput, encoder *httpbinding.Encoder) error {
@@ -375,10 +368,8 @@ func (m *awsRestjson1_serializeOpCalculateRoutes) Hand
 func (m *awsRestjson1_serializeOpCalculateRoutes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -427,7 +418,6 @@ func (m *awsRestjson1_serializeOpCalculateRoutes) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCalculateRoutesInput(v *CalculateRoutesInput, encoder *httpbinding.Encoder) error {
@@ -609,10 +599,8 @@ func (m *awsRestjson1_serializeOpOptimizeWaypoints) Ha
 func (m *awsRestjson1_serializeOpOptimizeWaypoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -661,7 +649,6 @@ func (m *awsRestjson1_serializeOpOptimizeWaypoints) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsOptimizeWaypointsInput(v *OptimizeWaypointsInput, encoder *httpbinding.Encoder) error {
@@ -785,10 +772,8 @@ func (m *awsRestjson1_serializeOpSnapToRoads) HandleSe
 func (m *awsRestjson1_serializeOpSnapToRoads) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -837,7 +822,6 @@ func (m *awsRestjson1_serializeOpSnapToRoads) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSnapToRoadsInput(v *SnapToRoadsInput, encoder *httpbinding.Encoder) error {
