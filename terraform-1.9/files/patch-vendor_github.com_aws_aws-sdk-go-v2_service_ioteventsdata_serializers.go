--- vendor/github.com/aws/aws-sdk-go-v2/service/ioteventsdata/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ioteventsdata/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpBatchAcknowledgeAlarm
 func (m *awsRestjson1_serializeOpBatchAcknowledgeAlarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -73,7 +70,6 @@ func (m *awsRestjson1_serializeOpBatchAcknowledgeAlarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchAcknowledgeAlarmInput(v *BatchAcknowledgeAlarmInput, encoder *httpbinding.Encoder) error {
@@ -108,10 +104,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteDetector) 
 func (m *awsRestjson1_serializeOpBatchDeleteDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -156,7 +150,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteDetector) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteDetectorInput(v *BatchDeleteDetectorInput, encoder *httpbinding.Encoder) error {
@@ -191,10 +184,8 @@ func (m *awsRestjson1_serializeOpBatchDisableAlarm) Ha
 func (m *awsRestjson1_serializeOpBatchDisableAlarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -239,7 +230,6 @@ func (m *awsRestjson1_serializeOpBatchDisableAlarm) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDisableAlarmInput(v *BatchDisableAlarmInput, encoder *httpbinding.Encoder) error {
@@ -274,10 +264,8 @@ func (m *awsRestjson1_serializeOpBatchEnableAlarm) Han
 func (m *awsRestjson1_serializeOpBatchEnableAlarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -322,7 +310,6 @@ func (m *awsRestjson1_serializeOpBatchEnableAlarm) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchEnableAlarmInput(v *BatchEnableAlarmInput, encoder *httpbinding.Encoder) error {
@@ -357,10 +344,8 @@ func (m *awsRestjson1_serializeOpBatchPutMessage) Hand
 func (m *awsRestjson1_serializeOpBatchPutMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -405,7 +390,6 @@ func (m *awsRestjson1_serializeOpBatchPutMessage) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchPutMessageInput(v *BatchPutMessageInput, encoder *httpbinding.Encoder) error {
@@ -440,10 +424,8 @@ func (m *awsRestjson1_serializeOpBatchResetAlarm) Hand
 func (m *awsRestjson1_serializeOpBatchResetAlarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -488,7 +470,6 @@ func (m *awsRestjson1_serializeOpBatchResetAlarm) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchResetAlarmInput(v *BatchResetAlarmInput, encoder *httpbinding.Encoder) error {
@@ -523,10 +504,8 @@ func (m *awsRestjson1_serializeOpBatchSnoozeAlarm) Han
 func (m *awsRestjson1_serializeOpBatchSnoozeAlarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -571,7 +550,6 @@ func (m *awsRestjson1_serializeOpBatchSnoozeAlarm) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchSnoozeAlarmInput(v *BatchSnoozeAlarmInput, encoder *httpbinding.Encoder) error {
@@ -606,10 +584,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateDetector) 
 func (m *awsRestjson1_serializeOpBatchUpdateDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -654,7 +630,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateDetector) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateDetectorInput(v *BatchUpdateDetectorInput, encoder *httpbinding.Encoder) error {
@@ -689,10 +664,8 @@ func (m *awsRestjson1_serializeOpDescribeAlarm) Handle
 func (m *awsRestjson1_serializeOpDescribeAlarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -730,7 +703,6 @@ func (m *awsRestjson1_serializeOpDescribeAlarm) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeAlarmInput(v *DescribeAlarmInput, encoder *httpbinding.Encoder) error {
@@ -764,10 +736,8 @@ func (m *awsRestjson1_serializeOpDescribeDetector) Han
 func (m *awsRestjson1_serializeOpDescribeDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -805,7 +775,6 @@ func (m *awsRestjson1_serializeOpDescribeDetector) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeDetectorInput(v *DescribeDetectorInput, encoder *httpbinding.Encoder) error {
@@ -839,10 +808,8 @@ func (m *awsRestjson1_serializeOpListAlarms) HandleSer
 func (m *awsRestjson1_serializeOpListAlarms) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -880,7 +847,6 @@ func (m *awsRestjson1_serializeOpListAlarms) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAlarmsInput(v *ListAlarmsInput, encoder *httpbinding.Encoder) error {
@@ -918,10 +884,8 @@ func (m *awsRestjson1_serializeOpListDetectors) Handle
 func (m *awsRestjson1_serializeOpListDetectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -959,7 +923,6 @@ func (m *awsRestjson1_serializeOpListDetectors) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDetectorsInput(v *ListDetectorsInput, encoder *httpbinding.Encoder) error {
