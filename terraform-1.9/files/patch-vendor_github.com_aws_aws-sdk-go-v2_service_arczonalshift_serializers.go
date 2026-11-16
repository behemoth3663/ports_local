--- vendor/github.com/aws/aws-sdk-go-v2/service/arczonalshift/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/arczonalshift/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpCancelPracticeRun) Ha
 func (m *awsRestjson1_serializeOpCancelPracticeRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_serializeOpCancelPracticeRun) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelPracticeRunInput(v *CancelPracticeRunInput, encoder *httpbinding.Encoder) error {
@@ -96,10 +92,8 @@ func (m *awsRestjson1_serializeOpCancelZonalShift) Han
 func (m *awsRestjson1_serializeOpCancelZonalShift) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -137,7 +131,6 @@ func (m *awsRestjson1_serializeOpCancelZonalShift) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCancelZonalShiftInput(v *CancelZonalShiftInput, encoder *httpbinding.Encoder) error {
@@ -167,10 +160,8 @@ func (m *awsRestjson1_serializeOpCreatePracticeRunConf
 func (m *awsRestjson1_serializeOpCreatePracticeRunConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -215,7 +206,6 @@ func (m *awsRestjson1_serializeOpCreatePracticeRunConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePracticeRunConfigurationInput(v *CreatePracticeRunConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -276,10 +266,8 @@ func (m *awsRestjson1_serializeOpDeletePracticeRunConf
 func (m *awsRestjson1_serializeOpDeletePracticeRunConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -317,7 +305,6 @@ func (m *awsRestjson1_serializeOpDeletePracticeRunConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePracticeRunConfigurationInput(v *DeletePracticeRunConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -347,10 +334,8 @@ func (m *awsRestjson1_serializeOpGetAutoshiftObserverN
 func (m *awsRestjson1_serializeOpGetAutoshiftObserverNotificationStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -384,7 +369,6 @@ func (m *awsRestjson1_serializeOpGetAutoshiftObserverN
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAutoshiftObserverNotificationStatusInput(v *GetAutoshiftObserverNotificationStatusInput, encoder *httpbinding.Encoder) error {
@@ -405,10 +389,8 @@ func (m *awsRestjson1_serializeOpGetManagedResource) H
 func (m *awsRestjson1_serializeOpGetManagedResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -446,7 +428,6 @@ func (m *awsRestjson1_serializeOpGetManagedResource) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetManagedResourceInput(v *GetManagedResourceInput, encoder *httpbinding.Encoder) error {
@@ -476,10 +457,8 @@ func (m *awsRestjson1_serializeOpListAutoshifts) Handl
 func (m *awsRestjson1_serializeOpListAutoshifts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -517,7 +496,6 @@ func (m *awsRestjson1_serializeOpListAutoshifts) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAutoshiftsInput(v *ListAutoshiftsInput, encoder *httpbinding.Encoder) error {
@@ -550,10 +528,8 @@ func (m *awsRestjson1_serializeOpListManagedResources)
 func (m *awsRestjson1_serializeOpListManagedResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -591,7 +567,6 @@ func (m *awsRestjson1_serializeOpListManagedResources)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListManagedResourcesInput(v *ListManagedResourcesInput, encoder *httpbinding.Encoder) error {
@@ -620,10 +595,8 @@ func (m *awsRestjson1_serializeOpListZonalShifts) Hand
 func (m *awsRestjson1_serializeOpListZonalShifts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -661,7 +634,6 @@ func (m *awsRestjson1_serializeOpListZonalShifts) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListZonalShiftsInput(v *ListZonalShiftsInput, encoder *httpbinding.Encoder) error {
@@ -698,10 +670,8 @@ func (m *awsRestjson1_serializeOpStartPracticeRun) Han
 func (m *awsRestjson1_serializeOpStartPracticeRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -746,7 +716,6 @@ func (m *awsRestjson1_serializeOpStartPracticeRun) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartPracticeRunInput(v *StartPracticeRunInput, encoder *httpbinding.Encoder) error {
@@ -789,10 +758,8 @@ func (m *awsRestjson1_serializeOpStartZonalShift) Hand
 func (m *awsRestjson1_serializeOpStartZonalShift) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -837,7 +804,6 @@ func (m *awsRestjson1_serializeOpStartZonalShift) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartZonalShiftInput(v *StartZonalShiftInput, encoder *httpbinding.Encoder) error {
@@ -885,10 +851,8 @@ func (m *awsRestjson1_serializeOpUpdateAutoshiftObserv
 func (m *awsRestjson1_serializeOpUpdateAutoshiftObserverNotificationStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -933,7 +897,6 @@ func (m *awsRestjson1_serializeOpUpdateAutoshiftObserv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAutoshiftObserverNotificationStatusInput(v *UpdateAutoshiftObserverNotificationStatusInput, encoder *httpbinding.Encoder) error {
@@ -966,10 +929,8 @@ func (m *awsRestjson1_serializeOpUpdatePracticeRunConf
 func (m *awsRestjson1_serializeOpUpdatePracticeRunConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1018,7 +979,6 @@ func (m *awsRestjson1_serializeOpUpdatePracticeRunConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePracticeRunConfigurationInput(v *UpdatePracticeRunConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1083,10 +1043,8 @@ func (m *awsRestjson1_serializeOpUpdateZonalAutoshiftC
 func (m *awsRestjson1_serializeOpUpdateZonalAutoshiftConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1135,7 +1093,6 @@ func (m *awsRestjson1_serializeOpUpdateZonalAutoshiftC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateZonalAutoshiftConfigurationInput(v *UpdateZonalAutoshiftConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1177,10 +1134,8 @@ func (m *awsRestjson1_serializeOpUpdateZonalShift) Han
 func (m *awsRestjson1_serializeOpUpdateZonalShift) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1229,7 +1184,6 @@ func (m *awsRestjson1_serializeOpUpdateZonalShift) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateZonalShiftInput(v *UpdateZonalShiftInput, encoder *httpbinding.Encoder) error {
