--- vendor/github.com/aws/aws-sdk-go-v2/service/guardduty/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/guardduty/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAcceptAdministratorIn
 func (m *awsRestjson1_serializeOpAcceptAdministratorInvitation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpAcceptAdministratorIn
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptAdministratorInvitationInput(v *AcceptAdministratorInvitationInput, encoder *httpbinding.Encoder) error {
@@ -124,10 +120,8 @@ func (m *awsRestjson1_serializeOpAcceptInvitation) Han
 func (m *awsRestjson1_serializeOpAcceptInvitation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -176,7 +170,6 @@ func (m *awsRestjson1_serializeOpAcceptInvitation) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAcceptInvitationInput(v *AcceptInvitationInput, encoder *httpbinding.Encoder) error {
@@ -223,10 +216,8 @@ func (m *awsRestjson1_serializeOpArchiveFindings) Hand
 func (m *awsRestjson1_serializeOpArchiveFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -275,7 +266,6 @@ func (m *awsRestjson1_serializeOpArchiveFindings) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsArchiveFindingsInput(v *ArchiveFindingsInput, encoder *httpbinding.Encoder) error {
@@ -319,10 +309,8 @@ func (m *awsRestjson1_serializeOpCreateDetector) Handl
 func (m *awsRestjson1_serializeOpCreateDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -367,7 +355,6 @@ func (m *awsRestjson1_serializeOpCreateDetector) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateDetectorInput(v *CreateDetectorInput, encoder *httpbinding.Encoder) error {
@@ -431,10 +418,8 @@ func (m *awsRestjson1_serializeOpCreateFilter) HandleS
 func (m *awsRestjson1_serializeOpCreateFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -483,7 +468,6 @@ func (m *awsRestjson1_serializeOpCreateFilter) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFilterInput(v *CreateFilterInput, encoder *httpbinding.Encoder) error {
@@ -559,10 +543,8 @@ func (m *awsRestjson1_serializeOpCreateIPSet) HandleSe
 func (m *awsRestjson1_serializeOpCreateIPSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -611,7 +593,6 @@ func (m *awsRestjson1_serializeOpCreateIPSet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateIPSetInput(v *CreateIPSetInput, encoder *httpbinding.Encoder) error {
@@ -685,10 +666,8 @@ func (m *awsRestjson1_serializeOpCreateMalwareProtecti
 func (m *awsRestjson1_serializeOpCreateMalwareProtectionPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -733,7 +712,6 @@ func (m *awsRestjson1_serializeOpCreateMalwareProtecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMalwareProtectionPlanInput(v *CreateMalwareProtectionPlanInput, encoder *httpbinding.Encoder) error {
@@ -792,10 +770,8 @@ func (m *awsRestjson1_serializeOpCreateMembers) Handle
 func (m *awsRestjson1_serializeOpCreateMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -844,7 +820,6 @@ func (m *awsRestjson1_serializeOpCreateMembers) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMembersInput(v *CreateMembersInput, encoder *httpbinding.Encoder) error {
@@ -888,10 +863,8 @@ func (m *awsRestjson1_serializeOpCreatePublishingDesti
 func (m *awsRestjson1_serializeOpCreatePublishingDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -940,7 +913,6 @@ func (m *awsRestjson1_serializeOpCreatePublishingDesti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePublishingDestinationInput(v *CreatePublishingDestinationInput, encoder *httpbinding.Encoder) error {
@@ -994,10 +966,8 @@ func (m *awsRestjson1_serializeOpCreateSampleFindings)
 func (m *awsRestjson1_serializeOpCreateSampleFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1046,7 +1016,6 @@ func (m *awsRestjson1_serializeOpCreateSampleFindings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSampleFindingsInput(v *CreateSampleFindingsInput, encoder *httpbinding.Encoder) error {
@@ -1090,10 +1059,8 @@ func (m *awsRestjson1_serializeOpCreateThreatIntelSet)
 func (m *awsRestjson1_serializeOpCreateThreatIntelSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1142,7 +1109,6 @@ func (m *awsRestjson1_serializeOpCreateThreatIntelSet)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateThreatIntelSetInput(v *CreateThreatIntelSetInput, encoder *httpbinding.Encoder) error {
@@ -1216,10 +1182,8 @@ func (m *awsRestjson1_serializeOpDeclineInvitations) H
 func (m *awsRestjson1_serializeOpDeclineInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1264,7 +1228,6 @@ func (m *awsRestjson1_serializeOpDeclineInvitations) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeclineInvitationsInput(v *DeclineInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -1299,10 +1262,8 @@ func (m *awsRestjson1_serializeOpDeleteDetector) Handl
 func (m *awsRestjson1_serializeOpDeleteDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1340,7 +1301,6 @@ func (m *awsRestjson1_serializeOpDeleteDetector) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteDetectorInput(v *DeleteDetectorInput, encoder *httpbinding.Encoder) error {
@@ -1370,10 +1330,8 @@ func (m *awsRestjson1_serializeOpDeleteFilter) HandleS
 func (m *awsRestjson1_serializeOpDeleteFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1411,7 +1369,6 @@ func (m *awsRestjson1_serializeOpDeleteFilter) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFilterInput(v *DeleteFilterInput, encoder *httpbinding.Encoder) error {
@@ -1450,10 +1407,8 @@ func (m *awsRestjson1_serializeOpDeleteInvitations) Ha
 func (m *awsRestjson1_serializeOpDeleteInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1498,7 +1453,6 @@ func (m *awsRestjson1_serializeOpDeleteInvitations) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteInvitationsInput(v *DeleteInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -1533,10 +1487,8 @@ func (m *awsRestjson1_serializeOpDeleteIPSet) HandleSe
 func (m *awsRestjson1_serializeOpDeleteIPSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1574,7 +1526,6 @@ func (m *awsRestjson1_serializeOpDeleteIPSet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteIPSetInput(v *DeleteIPSetInput, encoder *httpbinding.Encoder) error {
@@ -1613,10 +1564,8 @@ func (m *awsRestjson1_serializeOpDeleteMalwareProtecti
 func (m *awsRestjson1_serializeOpDeleteMalwareProtectionPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1654,7 +1603,6 @@ func (m *awsRestjson1_serializeOpDeleteMalwareProtecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMalwareProtectionPlanInput(v *DeleteMalwareProtectionPlanInput, encoder *httpbinding.Encoder) error {
@@ -1684,10 +1632,8 @@ func (m *awsRestjson1_serializeOpDeleteMembers) Handle
 func (m *awsRestjson1_serializeOpDeleteMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1736,7 +1682,6 @@ func (m *awsRestjson1_serializeOpDeleteMembers) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMembersInput(v *DeleteMembersInput, encoder *httpbinding.Encoder) error {
@@ -1780,10 +1725,8 @@ func (m *awsRestjson1_serializeOpDeletePublishingDesti
 func (m *awsRestjson1_serializeOpDeletePublishingDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1821,7 +1764,6 @@ func (m *awsRestjson1_serializeOpDeletePublishingDesti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePublishingDestinationInput(v *DeletePublishingDestinationInput, encoder *httpbinding.Encoder) error {
@@ -1860,10 +1802,8 @@ func (m *awsRestjson1_serializeOpDeleteThreatIntelSet)
 func (m *awsRestjson1_serializeOpDeleteThreatIntelSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1901,7 +1841,6 @@ func (m *awsRestjson1_serializeOpDeleteThreatIntelSet)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteThreatIntelSetInput(v *DeleteThreatIntelSetInput, encoder *httpbinding.Encoder) error {
@@ -1940,10 +1879,8 @@ func (m *awsRestjson1_serializeOpDescribeMalwareScans)
 func (m *awsRestjson1_serializeOpDescribeMalwareScans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1992,7 +1929,6 @@ func (m *awsRestjson1_serializeOpDescribeMalwareScans)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeMalwareScansInput(v *DescribeMalwareScansInput, encoder *httpbinding.Encoder) error {
@@ -2053,10 +1989,8 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
 func (m *awsRestjson1_serializeOpDescribeOrganizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2094,7 +2028,6 @@ func (m *awsRestjson1_serializeOpDescribeOrganizationC
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeOrganizationConfigurationInput(v *DescribeOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2132,10 +2065,8 @@ func (m *awsRestjson1_serializeOpDescribePublishingDes
 func (m *awsRestjson1_serializeOpDescribePublishingDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2173,7 +2104,6 @@ func (m *awsRestjson1_serializeOpDescribePublishingDes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribePublishingDestinationInput(v *DescribePublishingDestinationInput, encoder *httpbinding.Encoder) error {
@@ -2212,10 +2142,8 @@ func (m *awsRestjson1_serializeOpDisableOrganizationAd
 func (m *awsRestjson1_serializeOpDisableOrganizationAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2260,7 +2188,6 @@ func (m *awsRestjson1_serializeOpDisableOrganizationAd
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisableOrganizationAdminAccountInput(v *DisableOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -2293,10 +2220,8 @@ func (m *awsRestjson1_serializeOpDisassociateFromAdmin
 func (m *awsRestjson1_serializeOpDisassociateFromAdministratorAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2334,7 +2259,6 @@ func (m *awsRestjson1_serializeOpDisassociateFromAdmin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFromAdministratorAccountInput(v *DisassociateFromAdministratorAccountInput, encoder *httpbinding.Encoder) error {
@@ -2364,10 +2288,8 @@ func (m *awsRestjson1_serializeOpDisassociateFromMaste
 func (m *awsRestjson1_serializeOpDisassociateFromMasterAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2405,7 +2327,6 @@ func (m *awsRestjson1_serializeOpDisassociateFromMaste
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateFromMasterAccountInput(v *DisassociateFromMasterAccountInput, encoder *httpbinding.Encoder) error {
@@ -2435,10 +2356,8 @@ func (m *awsRestjson1_serializeOpDisassociateMembers) 
 func (m *awsRestjson1_serializeOpDisassociateMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2487,7 +2406,6 @@ func (m *awsRestjson1_serializeOpDisassociateMembers) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMembersInput(v *DisassociateMembersInput, encoder *httpbinding.Encoder) error {
@@ -2531,10 +2449,8 @@ func (m *awsRestjson1_serializeOpEnableOrganizationAdm
 func (m *awsRestjson1_serializeOpEnableOrganizationAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2579,7 +2495,6 @@ func (m *awsRestjson1_serializeOpEnableOrganizationAdm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsEnableOrganizationAdminAccountInput(v *EnableOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -2612,10 +2527,8 @@ func (m *awsRestjson1_serializeOpGetAdministratorAccou
 func (m *awsRestjson1_serializeOpGetAdministratorAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2653,7 +2566,6 @@ func (m *awsRestjson1_serializeOpGetAdministratorAccou
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAdministratorAccountInput(v *GetAdministratorAccountInput, encoder *httpbinding.Encoder) error {
@@ -2683,10 +2595,8 @@ func (m *awsRestjson1_serializeOpGetCoverageStatistics
 func (m *awsRestjson1_serializeOpGetCoverageStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2735,7 +2645,6 @@ func (m *awsRestjson1_serializeOpGetCoverageStatistics
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetCoverageStatisticsInput(v *GetCoverageStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -2786,10 +2695,8 @@ func (m *awsRestjson1_serializeOpGetDetector) HandleSe
 func (m *awsRestjson1_serializeOpGetDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2827,7 +2734,6 @@ func (m *awsRestjson1_serializeOpGetDetector) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDetectorInput(v *GetDetectorInput, encoder *httpbinding.Encoder) error {
@@ -2857,10 +2763,8 @@ func (m *awsRestjson1_serializeOpGetFilter) HandleSeri
 func (m *awsRestjson1_serializeOpGetFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2898,7 +2802,6 @@ func (m *awsRestjson1_serializeOpGetFilter) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFilterInput(v *GetFilterInput, encoder *httpbinding.Encoder) error {
@@ -2937,10 +2840,8 @@ func (m *awsRestjson1_serializeOpGetFindings) HandleSe
 func (m *awsRestjson1_serializeOpGetFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2989,7 +2890,6 @@ func (m *awsRestjson1_serializeOpGetFindings) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsInput(v *GetFindingsInput, encoder *httpbinding.Encoder) error {
@@ -3040,10 +2940,8 @@ func (m *awsRestjson1_serializeOpGetFindingsStatistics
 func (m *awsRestjson1_serializeOpGetFindingsStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3092,7 +2990,6 @@ func (m *awsRestjson1_serializeOpGetFindingsStatistics
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFindingsStatisticsInput(v *GetFindingsStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -3158,10 +3055,8 @@ func (m *awsRestjson1_serializeOpGetInvitationsCount) 
 func (m *awsRestjson1_serializeOpGetInvitationsCount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3195,7 +3090,6 @@ func (m *awsRestjson1_serializeOpGetInvitationsCount) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInvitationsCountInput(v *GetInvitationsCountInput, encoder *httpbinding.Encoder) error {
@@ -3216,10 +3110,8 @@ func (m *awsRestjson1_serializeOpGetIPSet) HandleSeria
 func (m *awsRestjson1_serializeOpGetIPSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3257,7 +3149,6 @@ func (m *awsRestjson1_serializeOpGetIPSet) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetIPSetInput(v *GetIPSetInput, encoder *httpbinding.Encoder) error {
@@ -3296,10 +3187,8 @@ func (m *awsRestjson1_serializeOpGetMalwareProtectionP
 func (m *awsRestjson1_serializeOpGetMalwareProtectionPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3337,7 +3226,6 @@ func (m *awsRestjson1_serializeOpGetMalwareProtectionP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMalwareProtectionPlanInput(v *GetMalwareProtectionPlanInput, encoder *httpbinding.Encoder) error {
@@ -3367,10 +3255,8 @@ func (m *awsRestjson1_serializeOpGetMalwareScanSetting
 func (m *awsRestjson1_serializeOpGetMalwareScanSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3408,7 +3294,6 @@ func (m *awsRestjson1_serializeOpGetMalwareScanSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMalwareScanSettingsInput(v *GetMalwareScanSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3438,10 +3323,8 @@ func (m *awsRestjson1_serializeOpGetMasterAccount) Han
 func (m *awsRestjson1_serializeOpGetMasterAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3479,7 +3362,6 @@ func (m *awsRestjson1_serializeOpGetMasterAccount) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMasterAccountInput(v *GetMasterAccountInput, encoder *httpbinding.Encoder) error {
@@ -3509,10 +3391,8 @@ func (m *awsRestjson1_serializeOpGetMemberDetectors) H
 func (m *awsRestjson1_serializeOpGetMemberDetectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3561,7 +3441,6 @@ func (m *awsRestjson1_serializeOpGetMemberDetectors) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMemberDetectorsInput(v *GetMemberDetectorsInput, encoder *httpbinding.Encoder) error {
@@ -3605,10 +3484,8 @@ func (m *awsRestjson1_serializeOpGetMembers) HandleSer
 func (m *awsRestjson1_serializeOpGetMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3657,7 +3534,6 @@ func (m *awsRestjson1_serializeOpGetMembers) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMembersInput(v *GetMembersInput, encoder *httpbinding.Encoder) error {
@@ -3701,10 +3577,8 @@ func (m *awsRestjson1_serializeOpGetOrganizationStatis
 func (m *awsRestjson1_serializeOpGetOrganizationStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3738,7 +3612,6 @@ func (m *awsRestjson1_serializeOpGetOrganizationStatis
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOrganizationStatisticsInput(v *GetOrganizationStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -3759,10 +3632,8 @@ func (m *awsRestjson1_serializeOpGetRemainingFreeTrial
 func (m *awsRestjson1_serializeOpGetRemainingFreeTrialDays) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3811,7 +3682,6 @@ func (m *awsRestjson1_serializeOpGetRemainingFreeTrial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRemainingFreeTrialDaysInput(v *GetRemainingFreeTrialDaysInput, encoder *httpbinding.Encoder) error {
@@ -3855,10 +3725,8 @@ func (m *awsRestjson1_serializeOpGetThreatIntelSet) Ha
 func (m *awsRestjson1_serializeOpGetThreatIntelSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3896,7 +3764,6 @@ func (m *awsRestjson1_serializeOpGetThreatIntelSet) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetThreatIntelSetInput(v *GetThreatIntelSetInput, encoder *httpbinding.Encoder) error {
@@ -3935,10 +3802,8 @@ func (m *awsRestjson1_serializeOpGetUsageStatistics) H
 func (m *awsRestjson1_serializeOpGetUsageStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3987,7 +3852,6 @@ func (m *awsRestjson1_serializeOpGetUsageStatistics) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUsageStatisticsInput(v *GetUsageStatisticsInput, encoder *httpbinding.Encoder) error {
@@ -4051,10 +3915,8 @@ func (m *awsRestjson1_serializeOpInviteMembers) Handle
 func (m *awsRestjson1_serializeOpInviteMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4103,7 +3965,6 @@ func (m *awsRestjson1_serializeOpInviteMembers) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInviteMembersInput(v *InviteMembersInput, encoder *httpbinding.Encoder) error {
@@ -4157,10 +4018,8 @@ func (m *awsRestjson1_serializeOpListCoverage) HandleS
 func (m *awsRestjson1_serializeOpListCoverage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4209,7 +4068,6 @@ func (m *awsRestjson1_serializeOpListCoverage) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListCoverageInput(v *ListCoverageInput, encoder *httpbinding.Encoder) error {
@@ -4270,10 +4128,8 @@ func (m *awsRestjson1_serializeOpListDetectors) Handle
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
@@ -4311,7 +4167,6 @@ func (m *awsRestjson1_serializeOpListDetectors) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListDetectorsInput(v *ListDetectorsInput, encoder *httpbinding.Encoder) error {
@@ -4340,10 +4195,8 @@ func (m *awsRestjson1_serializeOpListFilters) HandleSe
 func (m *awsRestjson1_serializeOpListFilters) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4381,7 +4234,6 @@ func (m *awsRestjson1_serializeOpListFilters) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFiltersInput(v *ListFiltersInput, encoder *httpbinding.Encoder) error {
@@ -4419,10 +4271,8 @@ func (m *awsRestjson1_serializeOpListFindings) HandleS
 func (m *awsRestjson1_serializeOpListFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4471,7 +4321,6 @@ func (m *awsRestjson1_serializeOpListFindings) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFindingsInput(v *ListFindingsInput, encoder *httpbinding.Encoder) error {
@@ -4532,10 +4381,8 @@ func (m *awsRestjson1_serializeOpListInvitations) Hand
 func (m *awsRestjson1_serializeOpListInvitations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4573,7 +4420,6 @@ func (m *awsRestjson1_serializeOpListInvitations) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListInvitationsInput(v *ListInvitationsInput, encoder *httpbinding.Encoder) error {
@@ -4602,10 +4448,8 @@ func (m *awsRestjson1_serializeOpListIPSets) HandleSer
 func (m *awsRestjson1_serializeOpListIPSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4643,7 +4487,6 @@ func (m *awsRestjson1_serializeOpListIPSets) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListIPSetsInput(v *ListIPSetsInput, encoder *httpbinding.Encoder) error {
@@ -4681,10 +4524,8 @@ func (m *awsRestjson1_serializeOpListMalwareProtection
 func (m *awsRestjson1_serializeOpListMalwareProtectionPlans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4722,7 +4563,6 @@ func (m *awsRestjson1_serializeOpListMalwareProtection
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMalwareProtectionPlansInput(v *ListMalwareProtectionPlansInput, encoder *httpbinding.Encoder) error {
@@ -4747,10 +4587,8 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
 func (m *awsRestjson1_serializeOpListMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4788,7 +4626,6 @@ func (m *awsRestjson1_serializeOpListMembers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMembersInput(v *ListMembersInput, encoder *httpbinding.Encoder) error {
@@ -4830,10 +4667,8 @@ func (m *awsRestjson1_serializeOpListOrganizationAdmin
 func (m *awsRestjson1_serializeOpListOrganizationAdminAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4871,7 +4706,6 @@ func (m *awsRestjson1_serializeOpListOrganizationAdmin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListOrganizationAdminAccountsInput(v *ListOrganizationAdminAccountsInput, encoder *httpbinding.Encoder) error {
@@ -4900,10 +4734,8 @@ func (m *awsRestjson1_serializeOpListPublishingDestina
 func (m *awsRestjson1_serializeOpListPublishingDestinations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4941,7 +4773,6 @@ func (m *awsRestjson1_serializeOpListPublishingDestina
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPublishingDestinationsInput(v *ListPublishingDestinationsInput, encoder *httpbinding.Encoder) error {
@@ -4979,10 +4810,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -5020,7 +4849,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -5050,10 +4878,8 @@ func (m *awsRestjson1_serializeOpListThreatIntelSets) 
 func (m *awsRestjson1_serializeOpListThreatIntelSets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5091,7 +4917,6 @@ func (m *awsRestjson1_serializeOpListThreatIntelSets) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListThreatIntelSetsInput(v *ListThreatIntelSetsInput, encoder *httpbinding.Encoder) error {
@@ -5129,10 +4954,8 @@ func (m *awsRestjson1_serializeOpStartMalwareScan) Han
 func (m *awsRestjson1_serializeOpStartMalwareScan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5177,7 +5000,6 @@ func (m *awsRestjson1_serializeOpStartMalwareScan) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMalwareScanInput(v *StartMalwareScanInput, encoder *httpbinding.Encoder) error {
@@ -5210,10 +5032,8 @@ func (m *awsRestjson1_serializeOpStartMonitoringMember
 func (m *awsRestjson1_serializeOpStartMonitoringMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5262,7 +5082,6 @@ func (m *awsRestjson1_serializeOpStartMonitoringMember
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartMonitoringMembersInput(v *StartMonitoringMembersInput, encoder *httpbinding.Encoder) error {
@@ -5306,10 +5125,8 @@ func (m *awsRestjson1_serializeOpStopMonitoringMembers
 func (m *awsRestjson1_serializeOpStopMonitoringMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5358,7 +5175,6 @@ func (m *awsRestjson1_serializeOpStopMonitoringMembers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopMonitoringMembersInput(v *StopMonitoringMembersInput, encoder *httpbinding.Encoder) error {
@@ -5402,10 +5218,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -5454,7 +5268,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5498,10 +5311,8 @@ func (m *awsRestjson1_serializeOpUnarchiveFindings) Ha
 func (m *awsRestjson1_serializeOpUnarchiveFindings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5550,7 +5361,6 @@ func (m *awsRestjson1_serializeOpUnarchiveFindings) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUnarchiveFindingsInput(v *UnarchiveFindingsInput, encoder *httpbinding.Encoder) error {
@@ -5594,10 +5404,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -5635,7 +5443,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -5671,10 +5478,8 @@ func (m *awsRestjson1_serializeOpUpdateDetector) Handl
 func (m *awsRestjson1_serializeOpUpdateDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5723,7 +5528,6 @@ func (m *awsRestjson1_serializeOpUpdateDetector) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateDetectorInput(v *UpdateDetectorInput, encoder *httpbinding.Encoder) error {
@@ -5784,10 +5588,8 @@ func (m *awsRestjson1_serializeOpUpdateFilter) HandleS
 func (m *awsRestjson1_serializeOpUpdateFilter) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5836,7 +5638,6 @@ func (m *awsRestjson1_serializeOpUpdateFilter) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFilterInput(v *UpdateFilterInput, encoder *httpbinding.Encoder) error {
@@ -5904,10 +5705,8 @@ func (m *awsRestjson1_serializeOpUpdateFindingsFeedbac
 func (m *awsRestjson1_serializeOpUpdateFindingsFeedback) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5956,7 +5755,6 @@ func (m *awsRestjson1_serializeOpUpdateFindingsFeedbac
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFindingsFeedbackInput(v *UpdateFindingsFeedbackInput, encoder *httpbinding.Encoder) error {
@@ -6010,10 +5808,8 @@ func (m *awsRestjson1_serializeOpUpdateIPSet) HandleSe
 func (m *awsRestjson1_serializeOpUpdateIPSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6062,7 +5858,6 @@ func (m *awsRestjson1_serializeOpUpdateIPSet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateIPSetInput(v *UpdateIPSetInput, encoder *httpbinding.Encoder) error {
@@ -6128,10 +5923,8 @@ func (m *awsRestjson1_serializeOpUpdateMalwareProtecti
 func (m *awsRestjson1_serializeOpUpdateMalwareProtectionPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6180,7 +5973,6 @@ func (m *awsRestjson1_serializeOpUpdateMalwareProtecti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMalwareProtectionPlanInput(v *UpdateMalwareProtectionPlanInput, encoder *httpbinding.Encoder) error {
@@ -6236,10 +6028,8 @@ func (m *awsRestjson1_serializeOpUpdateMalwareScanSett
 func (m *awsRestjson1_serializeOpUpdateMalwareScanSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6288,7 +6078,6 @@ func (m *awsRestjson1_serializeOpUpdateMalwareScanSett
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMalwareScanSettingsInput(v *UpdateMalwareScanSettingsInput, encoder *httpbinding.Encoder) error {
@@ -6337,10 +6126,8 @@ func (m *awsRestjson1_serializeOpUpdateMemberDetectors
 func (m *awsRestjson1_serializeOpUpdateMemberDetectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6389,7 +6176,6 @@ func (m *awsRestjson1_serializeOpUpdateMemberDetectors
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMemberDetectorsInput(v *UpdateMemberDetectorsInput, encoder *httpbinding.Encoder) error {
@@ -6447,10 +6233,8 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
 func (m *awsRestjson1_serializeOpUpdateOrganizationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6499,7 +6283,6 @@ func (m *awsRestjson1_serializeOpUpdateOrganizationCon
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateOrganizationConfigurationInput(v *UpdateOrganizationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -6560,10 +6343,8 @@ func (m *awsRestjson1_serializeOpUpdatePublishingDesti
 func (m *awsRestjson1_serializeOpUpdatePublishingDestination) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6612,7 +6393,6 @@ func (m *awsRestjson1_serializeOpUpdatePublishingDesti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePublishingDestinationInput(v *UpdatePublishingDestinationInput, encoder *httpbinding.Encoder) error {
@@ -6665,10 +6445,8 @@ func (m *awsRestjson1_serializeOpUpdateThreatIntelSet)
 func (m *awsRestjson1_serializeOpUpdateThreatIntelSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6717,7 +6495,6 @@ func (m *awsRestjson1_serializeOpUpdateThreatIntelSet)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateThreatIntelSetInput(v *UpdateThreatIntelSetInput, encoder *httpbinding.Encoder) error {
