--- vendor/github.com/aws/aws-sdk-go-v2/service/networkflowmonitor/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/networkflowmonitor/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCreateMonitor) Handle
 func (m *awsRestjson1_serializeOpCreateMonitor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpCreateMonitor) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMonitorInput(v *CreateMonitorInput, encoder *httpbinding.Encoder) error {
@@ -138,10 +134,8 @@ func (m *awsRestjson1_serializeOpCreateScope) HandleSe
 func (m *awsRestjson1_serializeOpCreateScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -186,7 +180,6 @@ func (m *awsRestjson1_serializeOpCreateScope) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateScopeInput(v *CreateScopeInput, encoder *httpbinding.Encoder) error {
@@ -233,10 +226,8 @@ func (m *awsRestjson1_serializeOpDeleteMonitor) Handle
 func (m *awsRestjson1_serializeOpDeleteMonitor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -274,7 +265,6 @@ func (m *awsRestjson1_serializeOpDeleteMonitor) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMonitorInput(v *DeleteMonitorInput, encoder *httpbinding.Encoder) error {
@@ -304,10 +294,8 @@ func (m *awsRestjson1_serializeOpDeleteScope) HandleSe
 func (m *awsRestjson1_serializeOpDeleteScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -345,7 +333,6 @@ func (m *awsRestjson1_serializeOpDeleteScope) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteScopeInput(v *DeleteScopeInput, encoder *httpbinding.Encoder) error {
@@ -375,10 +362,8 @@ func (m *awsRestjson1_serializeOpGetMonitor) HandleSer
 func (m *awsRestjson1_serializeOpGetMonitor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -416,7 +401,6 @@ func (m *awsRestjson1_serializeOpGetMonitor) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMonitorInput(v *GetMonitorInput, encoder *httpbinding.Encoder) error {
@@ -446,10 +430,8 @@ func (m *awsRestjson1_serializeOpGetQueryResultsMonito
 func (m *awsRestjson1_serializeOpGetQueryResultsMonitorTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -487,7 +469,6 @@ func (m *awsRestjson1_serializeOpGetQueryResultsMonito
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryResultsMonitorTopContributorsInput(v *GetQueryResultsMonitorTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -534,10 +515,8 @@ func (m *awsRestjson1_serializeOpGetQueryResultsWorklo
 func (m *awsRestjson1_serializeOpGetQueryResultsWorkloadInsightsTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -575,7 +554,6 @@ func (m *awsRestjson1_serializeOpGetQueryResultsWorklo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryResultsWorkloadInsightsTopContributorsInput(v *GetQueryResultsWorkloadInsightsTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -622,10 +600,8 @@ func (m *awsRestjson1_serializeOpGetQueryResultsWorklo
 func (m *awsRestjson1_serializeOpGetQueryResultsWorkloadInsightsTopContributorsData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -663,7 +639,6 @@ func (m *awsRestjson1_serializeOpGetQueryResultsWorklo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryResultsWorkloadInsightsTopContributorsDataInput(v *GetQueryResultsWorkloadInsightsTopContributorsDataInput, encoder *httpbinding.Encoder) error {
@@ -710,10 +685,8 @@ func (m *awsRestjson1_serializeOpGetQueryStatusMonitor
 func (m *awsRestjson1_serializeOpGetQueryStatusMonitorTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -751,7 +724,6 @@ func (m *awsRestjson1_serializeOpGetQueryStatusMonitor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryStatusMonitorTopContributorsInput(v *GetQueryStatusMonitorTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -790,10 +762,8 @@ func (m *awsRestjson1_serializeOpGetQueryStatusWorkloa
 func (m *awsRestjson1_serializeOpGetQueryStatusWorkloadInsightsTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -831,7 +801,6 @@ func (m *awsRestjson1_serializeOpGetQueryStatusWorkloa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryStatusWorkloadInsightsTopContributorsInput(v *GetQueryStatusWorkloadInsightsTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -870,10 +839,8 @@ func (m *awsRestjson1_serializeOpGetQueryStatusWorkloa
 func (m *awsRestjson1_serializeOpGetQueryStatusWorkloadInsightsTopContributorsData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -911,7 +878,6 @@ func (m *awsRestjson1_serializeOpGetQueryStatusWorkloa
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueryStatusWorkloadInsightsTopContributorsDataInput(v *GetQueryStatusWorkloadInsightsTopContributorsDataInput, encoder *httpbinding.Encoder) error {
@@ -950,10 +916,8 @@ func (m *awsRestjson1_serializeOpGetScope) HandleSeria
 func (m *awsRestjson1_serializeOpGetScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -991,7 +955,6 @@ func (m *awsRestjson1_serializeOpGetScope) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetScopeInput(v *GetScopeInput, encoder *httpbinding.Encoder) error {
@@ -1021,10 +984,8 @@ func (m *awsRestjson1_serializeOpListMonitors) HandleS
 func (m *awsRestjson1_serializeOpListMonitors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1062,7 +1023,6 @@ func (m *awsRestjson1_serializeOpListMonitors) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMonitorsInput(v *ListMonitorsInput, encoder *httpbinding.Encoder) error {
@@ -1095,10 +1055,8 @@ func (m *awsRestjson1_serializeOpListScopes) HandleSer
 func (m *awsRestjson1_serializeOpListScopes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1136,7 +1094,6 @@ func (m *awsRestjson1_serializeOpListScopes) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListScopesInput(v *ListScopesInput, encoder *httpbinding.Encoder) error {
@@ -1165,10 +1122,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1206,7 +1161,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -1236,10 +1190,8 @@ func (m *awsRestjson1_serializeOpStartQueryMonitorTopC
 func (m *awsRestjson1_serializeOpStartQueryMonitorTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1288,7 +1240,6 @@ func (m *awsRestjson1_serializeOpStartQueryMonitorTopC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartQueryMonitorTopContributorsInput(v *StartQueryMonitorTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -1350,10 +1301,8 @@ func (m *awsRestjson1_serializeOpStartQueryWorkloadIns
 func (m *awsRestjson1_serializeOpStartQueryWorkloadInsightsTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1402,7 +1351,6 @@ func (m *awsRestjson1_serializeOpStartQueryWorkloadIns
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartQueryWorkloadInsightsTopContributorsInput(v *StartQueryWorkloadInsightsTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -1464,10 +1412,8 @@ func (m *awsRestjson1_serializeOpStartQueryWorkloadIns
 func (m *awsRestjson1_serializeOpStartQueryWorkloadInsightsTopContributorsData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1516,7 +1462,6 @@ func (m *awsRestjson1_serializeOpStartQueryWorkloadIns
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartQueryWorkloadInsightsTopContributorsDataInput(v *StartQueryWorkloadInsightsTopContributorsDataInput, encoder *httpbinding.Encoder) error {
@@ -1573,10 +1518,8 @@ func (m *awsRestjson1_serializeOpStopQueryMonitorTopCo
 func (m *awsRestjson1_serializeOpStopQueryMonitorTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1614,7 +1557,6 @@ func (m *awsRestjson1_serializeOpStopQueryMonitorTopCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopQueryMonitorTopContributorsInput(v *StopQueryMonitorTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -1653,10 +1595,8 @@ func (m *awsRestjson1_serializeOpStopQueryWorkloadInsi
 func (m *awsRestjson1_serializeOpStopQueryWorkloadInsightsTopContributors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1694,7 +1634,6 @@ func (m *awsRestjson1_serializeOpStopQueryWorkloadInsi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopQueryWorkloadInsightsTopContributorsInput(v *StopQueryWorkloadInsightsTopContributorsInput, encoder *httpbinding.Encoder) error {
@@ -1733,10 +1672,8 @@ func (m *awsRestjson1_serializeOpStopQueryWorkloadInsi
 func (m *awsRestjson1_serializeOpStopQueryWorkloadInsightsTopContributorsData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1774,7 +1711,6 @@ func (m *awsRestjson1_serializeOpStopQueryWorkloadInsi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopQueryWorkloadInsightsTopContributorsDataInput(v *StopQueryWorkloadInsightsTopContributorsDataInput, encoder *httpbinding.Encoder) error {
@@ -1813,10 +1749,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1865,7 +1799,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -1909,10 +1842,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1950,7 +1881,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -1986,10 +1916,8 @@ func (m *awsRestjson1_serializeOpUpdateMonitor) Handle
 func (m *awsRestjson1_serializeOpUpdateMonitor) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2038,7 +1966,6 @@ func (m *awsRestjson1_serializeOpUpdateMonitor) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMonitorInput(v *UpdateMonitorInput, encoder *httpbinding.Encoder) error {
@@ -2108,10 +2035,8 @@ func (m *awsRestjson1_serializeOpUpdateScope) HandleSe
 func (m *awsRestjson1_serializeOpUpdateScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2160,7 +2085,6 @@ func (m *awsRestjson1_serializeOpUpdateScope) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateScopeInput(v *UpdateScopeInput, encoder *httpbinding.Encoder) error {
