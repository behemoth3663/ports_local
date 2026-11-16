--- vendor/github.com/aws/aws-sdk-go-v2/service/pinpointemail/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/pinpointemail/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -26,10 +25,8 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 func (m *awsRestjson1_serializeOpCreateConfigurationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -74,7 +71,6 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfigurationSetInput(v *CreateConfigurationSetInput, encoder *httpbinding.Encoder) error {
@@ -142,10 +138,8 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 func (m *awsRestjson1_serializeOpCreateConfigurationSetEventDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -194,7 +188,6 @@ func (m *awsRestjson1_serializeOpCreateConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateConfigurationSetEventDestinationInput(v *CreateConfigurationSetEventDestinationInput, encoder *httpbinding.Encoder) error {
@@ -243,10 +236,8 @@ func (m *awsRestjson1_serializeOpCreateDedicatedIpPool
 func (m *awsRestjson1_serializeOpCreateDedicatedIpPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -291,7 +282,6 @@ func (m *awsRestjson1_serializeOpCreateDedicatedIpPool
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDedicatedIpPoolInput(v *CreateDedicatedIpPoolInput, encoder *httpbinding.Encoder) error {
@@ -331,10 +321,8 @@ func (m *awsRestjson1_serializeOpCreateDeliverabilityT
 func (m *awsRestjson1_serializeOpCreateDeliverabilityTestReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -379,7 +367,6 @@ func (m *awsRestjson1_serializeOpCreateDeliverabilityT
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDeliverabilityTestReportInput(v *CreateDeliverabilityTestReportInput, encoder *httpbinding.Encoder) error {
@@ -431,10 +418,8 @@ func (m *awsRestjson1_serializeOpCreateEmailIdentity) 
 func (m *awsRestjson1_serializeOpCreateEmailIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -479,7 +464,6 @@ func (m *awsRestjson1_serializeOpCreateEmailIdentity) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateEmailIdentityInput(v *CreateEmailIdentityInput, encoder *httpbinding.Encoder) error {
@@ -519,10 +503,8 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 func (m *awsRestjson1_serializeOpDeleteConfigurationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -560,7 +542,6 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfigurationSetInput(v *DeleteConfigurationSetInput, encoder *httpbinding.Encoder) error {
@@ -590,10 +571,8 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 func (m *awsRestjson1_serializeOpDeleteConfigurationSetEventDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -631,7 +610,6 @@ func (m *awsRestjson1_serializeOpDeleteConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteConfigurationSetEventDestinationInput(v *DeleteConfigurationSetEventDestinationInput, encoder *httpbinding.Encoder) error {
@@ -670,10 +648,8 @@ func (m *awsRestjson1_serializeOpDeleteDedicatedIpPool
 func (m *awsRestjson1_serializeOpDeleteDedicatedIpPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -711,7 +687,6 @@ func (m *awsRestjson1_serializeOpDeleteDedicatedIpPool
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDedicatedIpPoolInput(v *DeleteDedicatedIpPoolInput, encoder *httpbinding.Encoder) error {
@@ -741,10 +716,8 @@ func (m *awsRestjson1_serializeOpDeleteEmailIdentity) 
 func (m *awsRestjson1_serializeOpDeleteEmailIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -782,7 +755,6 @@ func (m *awsRestjson1_serializeOpDeleteEmailIdentity) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEmailIdentityInput(v *DeleteEmailIdentityInput, encoder *httpbinding.Encoder) error {
@@ -812,10 +784,8 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 func (m *awsRestjson1_serializeOpGetAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -849,7 +819,6 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountInput(v *GetAccountInput, encoder *httpbinding.Encoder) error {
@@ -870,10 +839,8 @@ func (m *awsRestjson1_serializeOpGetBlacklistReports) 
 func (m *awsRestjson1_serializeOpGetBlacklistReports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -911,7 +878,6 @@ func (m *awsRestjson1_serializeOpGetBlacklistReports) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBlacklistReportsInput(v *GetBlacklistReportsInput, encoder *httpbinding.Encoder) error {
@@ -938,10 +904,8 @@ func (m *awsRestjson1_serializeOpGetConfigurationSet) 
 func (m *awsRestjson1_serializeOpGetConfigurationSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -979,7 +943,6 @@ func (m *awsRestjson1_serializeOpGetConfigurationSet) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfigurationSetInput(v *GetConfigurationSetInput, encoder *httpbinding.Encoder) error {
@@ -1009,10 +972,8 @@ func (m *awsRestjson1_serializeOpGetConfigurationSetEv
 func (m *awsRestjson1_serializeOpGetConfigurationSetEventDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1050,7 +1011,6 @@ func (m *awsRestjson1_serializeOpGetConfigurationSetEv
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetConfigurationSetEventDestinationsInput(v *GetConfigurationSetEventDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -1080,10 +1040,8 @@ func (m *awsRestjson1_serializeOpGetDedicatedIp) Handl
 func (m *awsRestjson1_serializeOpGetDedicatedIp) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1121,7 +1079,6 @@ func (m *awsRestjson1_serializeOpGetDedicatedIp) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDedicatedIpInput(v *GetDedicatedIpInput, encoder *httpbinding.Encoder) error {
@@ -1151,10 +1108,8 @@ func (m *awsRestjson1_serializeOpGetDedicatedIps) Hand
 func (m *awsRestjson1_serializeOpGetDedicatedIps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1192,7 +1147,6 @@ func (m *awsRestjson1_serializeOpGetDedicatedIps) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDedicatedIpsInput(v *GetDedicatedIpsInput, encoder *httpbinding.Encoder) error {
@@ -1225,10 +1179,8 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityDash
 func (m *awsRestjson1_serializeOpGetDeliverabilityDashboardOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1262,7 +1214,6 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityDash
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeliverabilityDashboardOptionsInput(v *GetDeliverabilityDashboardOptionsInput, encoder *httpbinding.Encoder) error {
@@ -1283,10 +1234,8 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityTest
 func (m *awsRestjson1_serializeOpGetDeliverabilityTestReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1324,7 +1273,6 @@ func (m *awsRestjson1_serializeOpGetDeliverabilityTest
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDeliverabilityTestReportInput(v *GetDeliverabilityTestReportInput, encoder *httpbinding.Encoder) error {
@@ -1354,10 +1302,8 @@ func (m *awsRestjson1_serializeOpGetDomainDeliverabili
 func (m *awsRestjson1_serializeOpGetDomainDeliverabilityCampaign) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1395,7 +1341,6 @@ func (m *awsRestjson1_serializeOpGetDomainDeliverabili
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainDeliverabilityCampaignInput(v *GetDomainDeliverabilityCampaignInput, encoder *httpbinding.Encoder) error {
@@ -1425,10 +1370,8 @@ func (m *awsRestjson1_serializeOpGetDomainStatisticsRe
 func (m *awsRestjson1_serializeOpGetDomainStatisticsReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1466,7 +1409,6 @@ func (m *awsRestjson1_serializeOpGetDomainStatisticsRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDomainStatisticsReportInput(v *GetDomainStatisticsReportInput, encoder *httpbinding.Encoder) error {
@@ -1504,10 +1446,8 @@ func (m *awsRestjson1_serializeOpGetEmailIdentity) Han
 func (m *awsRestjson1_serializeOpGetEmailIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1545,7 +1485,6 @@ func (m *awsRestjson1_serializeOpGetEmailIdentity) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEmailIdentityInput(v *GetEmailIdentityInput, encoder *httpbinding.Encoder) error {
@@ -1575,10 +1514,8 @@ func (m *awsRestjson1_serializeOpListConfigurationSets
 func (m *awsRestjson1_serializeOpListConfigurationSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1616,7 +1553,6 @@ func (m *awsRestjson1_serializeOpListConfigurationSets
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListConfigurationSetsInput(v *ListConfigurationSetsInput, encoder *httpbinding.Encoder) error {
@@ -1645,10 +1581,8 @@ func (m *awsRestjson1_serializeOpListDedicatedIpPools)
 func (m *awsRestjson1_serializeOpListDedicatedIpPools) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1686,7 +1620,6 @@ func (m *awsRestjson1_serializeOpListDedicatedIpPools)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDedicatedIpPoolsInput(v *ListDedicatedIpPoolsInput, encoder *httpbinding.Encoder) error {
@@ -1715,10 +1648,8 @@ func (m *awsRestjson1_serializeOpListDeliverabilityTes
 func (m *awsRestjson1_serializeOpListDeliverabilityTestReports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1756,7 +1687,6 @@ func (m *awsRestjson1_serializeOpListDeliverabilityTes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDeliverabilityTestReportsInput(v *ListDeliverabilityTestReportsInput, encoder *httpbinding.Encoder) error {
@@ -1785,10 +1715,8 @@ func (m *awsRestjson1_serializeOpListDomainDeliverabil
 func (m *awsRestjson1_serializeOpListDomainDeliverabilityCampaigns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1826,7 +1754,6 @@ func (m *awsRestjson1_serializeOpListDomainDeliverabil
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDomainDeliverabilityCampaignsInput(v *ListDomainDeliverabilityCampaignsInput, encoder *httpbinding.Encoder) error {
@@ -1872,10 +1799,8 @@ func (m *awsRestjson1_serializeOpListEmailIdentities) 
 func (m *awsRestjson1_serializeOpListEmailIdentities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1913,7 +1838,6 @@ func (m *awsRestjson1_serializeOpListEmailIdentities) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListEmailIdentitiesInput(v *ListEmailIdentitiesInput, encoder *httpbinding.Encoder) error {
@@ -1942,10 +1866,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -1983,7 +1905,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -2008,10 +1929,8 @@ func (m *awsRestjson1_serializeOpPutAccountDedicatedIp
 func (m *awsRestjson1_serializeOpPutAccountDedicatedIpWarmupAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2056,7 +1975,6 @@ func (m *awsRestjson1_serializeOpPutAccountDedicatedIp
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountDedicatedIpWarmupAttributesInput(v *PutAccountDedicatedIpWarmupAttributesInput, encoder *httpbinding.Encoder) error {
@@ -2089,10 +2007,8 @@ func (m *awsRestjson1_serializeOpPutAccountSendingAttr
 func (m *awsRestjson1_serializeOpPutAccountSendingAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2137,7 +2053,6 @@ func (m *awsRestjson1_serializeOpPutAccountSendingAttr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutAccountSendingAttributesInput(v *PutAccountSendingAttributesInput, encoder *httpbinding.Encoder) error {
@@ -2170,10 +2085,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetDe
 func (m *awsRestjson1_serializeOpPutConfigurationSetDeliveryOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2222,7 +2135,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetDe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetDeliveryOptionsInput(v *PutConfigurationSetDeliveryOptionsInput, encoder *httpbinding.Encoder) error {
@@ -2269,10 +2181,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetRe
 func (m *awsRestjson1_serializeOpPutConfigurationSetReputationOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2321,7 +2231,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetReputationOptionsInput(v *PutConfigurationSetReputationOptionsInput, encoder *httpbinding.Encoder) error {
@@ -2363,10 +2272,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetSe
 func (m *awsRestjson1_serializeOpPutConfigurationSetSendingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2415,7 +2322,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetSendingOptionsInput(v *PutConfigurationSetSendingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -2457,10 +2363,8 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetTr
 func (m *awsRestjson1_serializeOpPutConfigurationSetTrackingOptions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2509,7 +2413,6 @@ func (m *awsRestjson1_serializeOpPutConfigurationSetTr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutConfigurationSetTrackingOptionsInput(v *PutConfigurationSetTrackingOptionsInput, encoder *httpbinding.Encoder) error {
@@ -2551,10 +2454,8 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpInPool)
 func (m *awsRestjson1_serializeOpPutDedicatedIpInPool) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2603,7 +2504,6 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpInPool)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDedicatedIpInPoolInput(v *PutDedicatedIpInPoolInput, encoder *httpbinding.Encoder) error {
@@ -2645,10 +2545,8 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpWarmupA
 func (m *awsRestjson1_serializeOpPutDedicatedIpWarmupAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2697,7 +2595,6 @@ func (m *awsRestjson1_serializeOpPutDedicatedIpWarmupA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDedicatedIpWarmupAttributesInput(v *PutDedicatedIpWarmupAttributesInput, encoder *httpbinding.Encoder) error {
@@ -2739,10 +2636,8 @@ func (m *awsRestjson1_serializeOpPutDeliverabilityDash
 func (m *awsRestjson1_serializeOpPutDeliverabilityDashboardOption) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2787,7 +2682,6 @@ func (m *awsRestjson1_serializeOpPutDeliverabilityDash
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutDeliverabilityDashboardOptionInput(v *PutDeliverabilityDashboardOptionInput, encoder *httpbinding.Encoder) error {
@@ -2827,10 +2721,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityDkimA
 func (m *awsRestjson1_serializeOpPutEmailIdentityDkimAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2879,7 +2771,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityDkimA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityDkimAttributesInput(v *PutEmailIdentityDkimAttributesInput, encoder *httpbinding.Encoder) error {
@@ -2921,10 +2812,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityFeedb
 func (m *awsRestjson1_serializeOpPutEmailIdentityFeedbackAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2973,7 +2862,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityFeedb
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityFeedbackAttributesInput(v *PutEmailIdentityFeedbackAttributesInput, encoder *httpbinding.Encoder) error {
@@ -3015,10 +2903,8 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityMailF
 func (m *awsRestjson1_serializeOpPutEmailIdentityMailFromAttributes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3067,7 +2953,6 @@ func (m *awsRestjson1_serializeOpPutEmailIdentityMailF
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEmailIdentityMailFromAttributesInput(v *PutEmailIdentityMailFromAttributesInput, encoder *httpbinding.Encoder) error {
@@ -3114,10 +2999,8 @@ func (m *awsRestjson1_serializeOpSendEmail) HandleSeri
 func (m *awsRestjson1_serializeOpSendEmail) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3162,7 +3045,6 @@ func (m *awsRestjson1_serializeOpSendEmail) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSendEmailInput(v *SendEmailInput, encoder *httpbinding.Encoder) error {
@@ -3233,10 +3115,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -3281,7 +3161,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3321,10 +3200,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -3362,7 +3239,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3393,10 +3269,8 @@ func (m *awsRestjson1_serializeOpUpdateConfigurationSe
 func (m *awsRestjson1_serializeOpUpdateConfigurationSetEventDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3445,7 +3319,6 @@ func (m *awsRestjson1_serializeOpUpdateConfigurationSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateConfigurationSetEventDestinationInput(v *UpdateConfigurationSetEventDestinationInput, encoder *httpbinding.Encoder) error {
