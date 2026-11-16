--- vendor/github.com/aws/aws-sdk-go-v2/service/deadline/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/deadline/serializers.go
@@ -12,7 +12,6 @@ import (
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"math"
 )
@@ -27,10 +26,8 @@ func (m *awsRestjson1_serializeOpAssociateMemberToFarm
 func (m *awsRestjson1_serializeOpAssociateMemberToFarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -79,7 +76,6 @@ func (m *awsRestjson1_serializeOpAssociateMemberToFarm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateMemberToFarmInput(v *AssociateMemberToFarmInput, encoder *httpbinding.Encoder) error {
@@ -140,10 +136,8 @@ func (m *awsRestjson1_serializeOpAssociateMemberToFlee
 func (m *awsRestjson1_serializeOpAssociateMemberToFleet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -192,7 +186,6 @@ func (m *awsRestjson1_serializeOpAssociateMemberToFlee
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateMemberToFleetInput(v *AssociateMemberToFleetInput, encoder *httpbinding.Encoder) error {
@@ -262,10 +255,8 @@ func (m *awsRestjson1_serializeOpAssociateMemberToJob)
 func (m *awsRestjson1_serializeOpAssociateMemberToJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -314,7 +305,6 @@ func (m *awsRestjson1_serializeOpAssociateMemberToJob)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateMemberToJobInput(v *AssociateMemberToJobInput, encoder *httpbinding.Encoder) error {
@@ -393,10 +383,8 @@ func (m *awsRestjson1_serializeOpAssociateMemberToQueu
 func (m *awsRestjson1_serializeOpAssociateMemberToQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -445,7 +433,6 @@ func (m *awsRestjson1_serializeOpAssociateMemberToQueu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateMemberToQueueInput(v *AssociateMemberToQueueInput, encoder *httpbinding.Encoder) error {
@@ -515,10 +502,8 @@ func (m *awsRestjson1_serializeOpAssumeFleetRoleForRea
 func (m *awsRestjson1_serializeOpAssumeFleetRoleForRead) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -556,7 +541,6 @@ func (m *awsRestjson1_serializeOpAssumeFleetRoleForRea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssumeFleetRoleForReadInput(v *AssumeFleetRoleForReadInput, encoder *httpbinding.Encoder) error {
@@ -595,10 +579,8 @@ func (m *awsRestjson1_serializeOpAssumeFleetRoleForWor
 func (m *awsRestjson1_serializeOpAssumeFleetRoleForWorker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -636,7 +618,6 @@ func (m *awsRestjson1_serializeOpAssumeFleetRoleForWor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssumeFleetRoleForWorkerInput(v *AssumeFleetRoleForWorkerInput, encoder *httpbinding.Encoder) error {
@@ -684,10 +665,8 @@ func (m *awsRestjson1_serializeOpAssumeQueueRoleForRea
 func (m *awsRestjson1_serializeOpAssumeQueueRoleForRead) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -725,7 +704,6 @@ func (m *awsRestjson1_serializeOpAssumeQueueRoleForRea
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssumeQueueRoleForReadInput(v *AssumeQueueRoleForReadInput, encoder *httpbinding.Encoder) error {
@@ -764,10 +742,8 @@ func (m *awsRestjson1_serializeOpAssumeQueueRoleForUse
 func (m *awsRestjson1_serializeOpAssumeQueueRoleForUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -805,7 +781,6 @@ func (m *awsRestjson1_serializeOpAssumeQueueRoleForUse
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssumeQueueRoleForUserInput(v *AssumeQueueRoleForUserInput, encoder *httpbinding.Encoder) error {
@@ -844,10 +819,8 @@ func (m *awsRestjson1_serializeOpAssumeQueueRoleForWor
 func (m *awsRestjson1_serializeOpAssumeQueueRoleForWorker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -885,7 +858,6 @@ func (m *awsRestjson1_serializeOpAssumeQueueRoleForWor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssumeQueueRoleForWorkerInput(v *AssumeQueueRoleForWorkerInput, encoder *httpbinding.Encoder) error {
@@ -937,10 +909,8 @@ func (m *awsRestjson1_serializeOpBatchGetJobEntity) Ha
 func (m *awsRestjson1_serializeOpBatchGetJobEntity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -989,7 +959,6 @@ func (m *awsRestjson1_serializeOpBatchGetJobEntity) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchGetJobEntityInput(v *BatchGetJobEntityInput, encoder *httpbinding.Encoder) error {
@@ -1051,10 +1020,8 @@ func (m *awsRestjson1_serializeOpCopyJobTemplate) Hand
 func (m *awsRestjson1_serializeOpCopyJobTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1103,7 +1070,6 @@ func (m *awsRestjson1_serializeOpCopyJobTemplate) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCopyJobTemplateInput(v *CopyJobTemplateInput, encoder *httpbinding.Encoder) error {
@@ -1165,10 +1131,8 @@ func (m *awsRestjson1_serializeOpCreateBudget) HandleS
 func (m *awsRestjson1_serializeOpCreateBudget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1217,7 +1181,6 @@ func (m *awsRestjson1_serializeOpCreateBudget) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBudgetInput(v *CreateBudgetInput, encoder *httpbinding.Encoder) error {
@@ -1308,10 +1271,8 @@ func (m *awsRestjson1_serializeOpCreateFarm) HandleSer
 func (m *awsRestjson1_serializeOpCreateFarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1360,7 +1321,6 @@ func (m *awsRestjson1_serializeOpCreateFarm) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFarmInput(v *CreateFarmInput, encoder *httpbinding.Encoder) error {
@@ -1415,10 +1375,8 @@ func (m *awsRestjson1_serializeOpCreateFleet) HandleSe
 func (m *awsRestjson1_serializeOpCreateFleet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1467,7 +1425,6 @@ func (m *awsRestjson1_serializeOpCreateFleet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateFleetInput(v *CreateFleetInput, encoder *httpbinding.Encoder) error {
@@ -1555,10 +1512,8 @@ func (m *awsRestjson1_serializeOpCreateJob) HandleSeri
 func (m *awsRestjson1_serializeOpCreateJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1607,7 +1562,6 @@ func (m *awsRestjson1_serializeOpCreateJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateJobInput(v *CreateJobInput, encoder *httpbinding.Encoder) error {
@@ -1717,10 +1671,8 @@ func (m *awsRestjson1_serializeOpCreateLicenseEndpoint
 func (m *awsRestjson1_serializeOpCreateLicenseEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1769,7 +1721,6 @@ func (m *awsRestjson1_serializeOpCreateLicenseEndpoint
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLicenseEndpointInput(v *CreateLicenseEndpointInput, encoder *httpbinding.Encoder) error {
@@ -1828,10 +1779,8 @@ func (m *awsRestjson1_serializeOpCreateLimit) HandleSe
 func (m *awsRestjson1_serializeOpCreateLimit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1880,7 +1829,6 @@ func (m *awsRestjson1_serializeOpCreateLimit) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLimitInput(v *CreateLimitInput, encoder *httpbinding.Encoder) error {
@@ -1942,10 +1890,8 @@ func (m *awsRestjson1_serializeOpCreateMonitor) Handle
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
@@ -1994,7 +1940,6 @@ func (m *awsRestjson1_serializeOpCreateMonitor) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMonitorInput(v *CreateMonitorInput, encoder *httpbinding.Encoder) error {
@@ -2047,10 +1992,8 @@ func (m *awsRestjson1_serializeOpCreateQueue) HandleSe
 func (m *awsRestjson1_serializeOpCreateQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2099,7 +2042,6 @@ func (m *awsRestjson1_serializeOpCreateQueue) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateQueueInput(v *CreateQueueInput, encoder *httpbinding.Encoder) error {
@@ -2196,10 +2138,8 @@ func (m *awsRestjson1_serializeOpCreateQueueEnvironmen
 func (m *awsRestjson1_serializeOpCreateQueueEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2248,7 +2188,6 @@ func (m *awsRestjson1_serializeOpCreateQueueEnvironmen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateQueueEnvironmentInput(v *CreateQueueEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -2314,10 +2253,8 @@ func (m *awsRestjson1_serializeOpCreateQueueFleetAssoc
 func (m *awsRestjson1_serializeOpCreateQueueFleetAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2366,7 +2303,6 @@ func (m *awsRestjson1_serializeOpCreateQueueFleetAssoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateQueueFleetAssociationInput(v *CreateQueueFleetAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2413,10 +2349,8 @@ func (m *awsRestjson1_serializeOpCreateQueueLimitAssoc
 func (m *awsRestjson1_serializeOpCreateQueueLimitAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2465,7 +2399,6 @@ func (m *awsRestjson1_serializeOpCreateQueueLimitAssoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateQueueLimitAssociationInput(v *CreateQueueLimitAssociationInput, encoder *httpbinding.Encoder) error {
@@ -2512,10 +2445,8 @@ func (m *awsRestjson1_serializeOpCreateStorageProfile)
 func (m *awsRestjson1_serializeOpCreateStorageProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2564,7 +2495,6 @@ func (m *awsRestjson1_serializeOpCreateStorageProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateStorageProfileInput(v *CreateStorageProfileInput, encoder *httpbinding.Encoder) error {
@@ -2623,10 +2553,8 @@ func (m *awsRestjson1_serializeOpCreateWorker) HandleS
 func (m *awsRestjson1_serializeOpCreateWorker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2675,7 +2603,6 @@ func (m *awsRestjson1_serializeOpCreateWorker) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateWorkerInput(v *CreateWorkerInput, encoder *httpbinding.Encoder) error {
@@ -2740,10 +2667,8 @@ func (m *awsRestjson1_serializeOpDeleteBudget) HandleS
 func (m *awsRestjson1_serializeOpDeleteBudget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2781,7 +2706,6 @@ func (m *awsRestjson1_serializeOpDeleteBudget) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteBudgetInput(v *DeleteBudgetInput, encoder *httpbinding.Encoder) error {
@@ -2820,10 +2744,8 @@ func (m *awsRestjson1_serializeOpDeleteFarm) HandleSer
 func (m *awsRestjson1_serializeOpDeleteFarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2861,7 +2783,6 @@ func (m *awsRestjson1_serializeOpDeleteFarm) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFarmInput(v *DeleteFarmInput, encoder *httpbinding.Encoder) error {
@@ -2891,10 +2812,8 @@ func (m *awsRestjson1_serializeOpDeleteFleet) HandleSe
 func (m *awsRestjson1_serializeOpDeleteFleet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2932,7 +2851,6 @@ func (m *awsRestjson1_serializeOpDeleteFleet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteFleetInput(v *DeleteFleetInput, encoder *httpbinding.Encoder) error {
@@ -2976,10 +2894,8 @@ func (m *awsRestjson1_serializeOpDeleteLicenseEndpoint
 func (m *awsRestjson1_serializeOpDeleteLicenseEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3017,7 +2933,6 @@ func (m *awsRestjson1_serializeOpDeleteLicenseEndpoint
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLicenseEndpointInput(v *DeleteLicenseEndpointInput, encoder *httpbinding.Encoder) error {
@@ -3047,10 +2962,8 @@ func (m *awsRestjson1_serializeOpDeleteLimit) HandleSe
 func (m *awsRestjson1_serializeOpDeleteLimit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3088,7 +3001,6 @@ func (m *awsRestjson1_serializeOpDeleteLimit) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLimitInput(v *DeleteLimitInput, encoder *httpbinding.Encoder) error {
@@ -3127,10 +3039,8 @@ func (m *awsRestjson1_serializeOpDeleteMeteredProduct)
 func (m *awsRestjson1_serializeOpDeleteMeteredProduct) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3168,7 +3078,6 @@ func (m *awsRestjson1_serializeOpDeleteMeteredProduct)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMeteredProductInput(v *DeleteMeteredProductInput, encoder *httpbinding.Encoder) error {
@@ -3207,10 +3116,8 @@ func (m *awsRestjson1_serializeOpDeleteMonitor) Handle
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
@@ -3248,7 +3155,6 @@ func (m *awsRestjson1_serializeOpDeleteMonitor) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteMonitorInput(v *DeleteMonitorInput, encoder *httpbinding.Encoder) error {
@@ -3278,10 +3184,8 @@ func (m *awsRestjson1_serializeOpDeleteQueue) HandleSe
 func (m *awsRestjson1_serializeOpDeleteQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3319,7 +3223,6 @@ func (m *awsRestjson1_serializeOpDeleteQueue) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQueueInput(v *DeleteQueueInput, encoder *httpbinding.Encoder) error {
@@ -3358,10 +3261,8 @@ func (m *awsRestjson1_serializeOpDeleteQueueEnvironmen
 func (m *awsRestjson1_serializeOpDeleteQueueEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3399,7 +3300,6 @@ func (m *awsRestjson1_serializeOpDeleteQueueEnvironmen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQueueEnvironmentInput(v *DeleteQueueEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -3447,10 +3347,8 @@ func (m *awsRestjson1_serializeOpDeleteQueueFleetAssoc
 func (m *awsRestjson1_serializeOpDeleteQueueFleetAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3488,7 +3386,6 @@ func (m *awsRestjson1_serializeOpDeleteQueueFleetAssoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQueueFleetAssociationInput(v *DeleteQueueFleetAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3536,10 +3433,8 @@ func (m *awsRestjson1_serializeOpDeleteQueueLimitAssoc
 func (m *awsRestjson1_serializeOpDeleteQueueLimitAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3577,7 +3472,6 @@ func (m *awsRestjson1_serializeOpDeleteQueueLimitAssoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteQueueLimitAssociationInput(v *DeleteQueueLimitAssociationInput, encoder *httpbinding.Encoder) error {
@@ -3625,10 +3519,8 @@ func (m *awsRestjson1_serializeOpDeleteStorageProfile)
 func (m *awsRestjson1_serializeOpDeleteStorageProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3666,7 +3558,6 @@ func (m *awsRestjson1_serializeOpDeleteStorageProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteStorageProfileInput(v *DeleteStorageProfileInput, encoder *httpbinding.Encoder) error {
@@ -3705,10 +3596,8 @@ func (m *awsRestjson1_serializeOpDeleteWorker) HandleS
 func (m *awsRestjson1_serializeOpDeleteWorker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3746,7 +3635,6 @@ func (m *awsRestjson1_serializeOpDeleteWorker) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteWorkerInput(v *DeleteWorkerInput, encoder *httpbinding.Encoder) error {
@@ -3794,10 +3682,8 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 func (m *awsRestjson1_serializeOpDisassociateMemberFromFarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3835,7 +3721,6 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMemberFromFarmInput(v *DisassociateMemberFromFarmInput, encoder *httpbinding.Encoder) error {
@@ -3874,10 +3759,8 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 func (m *awsRestjson1_serializeOpDisassociateMemberFromFleet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3915,7 +3798,6 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMemberFromFleetInput(v *DisassociateMemberFromFleetInput, encoder *httpbinding.Encoder) error {
@@ -3963,10 +3845,8 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 func (m *awsRestjson1_serializeOpDisassociateMemberFromJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4004,7 +3884,6 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMemberFromJobInput(v *DisassociateMemberFromJobInput, encoder *httpbinding.Encoder) error {
@@ -4061,10 +3940,8 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 func (m *awsRestjson1_serializeOpDisassociateMemberFromQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4102,7 +3979,6 @@ func (m *awsRestjson1_serializeOpDisassociateMemberFro
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateMemberFromQueueInput(v *DisassociateMemberFromQueueInput, encoder *httpbinding.Encoder) error {
@@ -4150,10 +4026,8 @@ func (m *awsRestjson1_serializeOpGetBudget) HandleSeri
 func (m *awsRestjson1_serializeOpGetBudget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4191,7 +4065,6 @@ func (m *awsRestjson1_serializeOpGetBudget) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBudgetInput(v *GetBudgetInput, encoder *httpbinding.Encoder) error {
@@ -4230,10 +4103,8 @@ func (m *awsRestjson1_serializeOpGetFarm) HandleSerial
 func (m *awsRestjson1_serializeOpGetFarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4271,7 +4142,6 @@ func (m *awsRestjson1_serializeOpGetFarm) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFarmInput(v *GetFarmInput, encoder *httpbinding.Encoder) error {
@@ -4301,10 +4171,8 @@ func (m *awsRestjson1_serializeOpGetFleet) HandleSeria
 func (m *awsRestjson1_serializeOpGetFleet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4342,7 +4210,6 @@ func (m *awsRestjson1_serializeOpGetFleet) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFleetInput(v *GetFleetInput, encoder *httpbinding.Encoder) error {
@@ -4381,10 +4248,8 @@ func (m *awsRestjson1_serializeOpGetJob) HandleSeriali
 func (m *awsRestjson1_serializeOpGetJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4422,7 +4287,6 @@ func (m *awsRestjson1_serializeOpGetJob) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetJobInput(v *GetJobInput, encoder *httpbinding.Encoder) error {
@@ -4470,10 +4334,8 @@ func (m *awsRestjson1_serializeOpGetLicenseEndpoint) H
 func (m *awsRestjson1_serializeOpGetLicenseEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4511,7 +4373,6 @@ func (m *awsRestjson1_serializeOpGetLicenseEndpoint) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLicenseEndpointInput(v *GetLicenseEndpointInput, encoder *httpbinding.Encoder) error {
@@ -4541,10 +4402,8 @@ func (m *awsRestjson1_serializeOpGetLimit) HandleSeria
 func (m *awsRestjson1_serializeOpGetLimit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4582,7 +4441,6 @@ func (m *awsRestjson1_serializeOpGetLimit) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLimitInput(v *GetLimitInput, encoder *httpbinding.Encoder) error {
@@ -4621,10 +4479,8 @@ func (m *awsRestjson1_serializeOpGetMonitor) HandleSer
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
@@ -4662,7 +4518,6 @@ func (m *awsRestjson1_serializeOpGetMonitor) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetMonitorInput(v *GetMonitorInput, encoder *httpbinding.Encoder) error {
@@ -4692,10 +4547,8 @@ func (m *awsRestjson1_serializeOpGetQueue) HandleSeria
 func (m *awsRestjson1_serializeOpGetQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4733,7 +4586,6 @@ func (m *awsRestjson1_serializeOpGetQueue) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueueInput(v *GetQueueInput, encoder *httpbinding.Encoder) error {
@@ -4772,10 +4624,8 @@ func (m *awsRestjson1_serializeOpGetQueueEnvironment) 
 func (m *awsRestjson1_serializeOpGetQueueEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4813,7 +4663,6 @@ func (m *awsRestjson1_serializeOpGetQueueEnvironment) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueueEnvironmentInput(v *GetQueueEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -4861,10 +4710,8 @@ func (m *awsRestjson1_serializeOpGetQueueFleetAssociat
 func (m *awsRestjson1_serializeOpGetQueueFleetAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4902,7 +4749,6 @@ func (m *awsRestjson1_serializeOpGetQueueFleetAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueueFleetAssociationInput(v *GetQueueFleetAssociationInput, encoder *httpbinding.Encoder) error {
@@ -4950,10 +4796,8 @@ func (m *awsRestjson1_serializeOpGetQueueLimitAssociat
 func (m *awsRestjson1_serializeOpGetQueueLimitAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4991,7 +4835,6 @@ func (m *awsRestjson1_serializeOpGetQueueLimitAssociat
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetQueueLimitAssociationInput(v *GetQueueLimitAssociationInput, encoder *httpbinding.Encoder) error {
@@ -5039,10 +4882,8 @@ func (m *awsRestjson1_serializeOpGetSession) HandleSer
 func (m *awsRestjson1_serializeOpGetSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5080,7 +4921,6 @@ func (m *awsRestjson1_serializeOpGetSession) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSessionInput(v *GetSessionInput, encoder *httpbinding.Encoder) error {
@@ -5137,10 +4977,8 @@ func (m *awsRestjson1_serializeOpGetSessionAction) Han
 func (m *awsRestjson1_serializeOpGetSessionAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5178,7 +5016,6 @@ func (m *awsRestjson1_serializeOpGetSessionAction) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSessionActionInput(v *GetSessionActionInput, encoder *httpbinding.Encoder) error {
@@ -5235,10 +5072,8 @@ func (m *awsRestjson1_serializeOpGetSessionsStatistics
 func (m *awsRestjson1_serializeOpGetSessionsStatisticsAggregation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5276,7 +5111,6 @@ func (m *awsRestjson1_serializeOpGetSessionsStatistics
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSessionsStatisticsAggregationInput(v *GetSessionsStatisticsAggregationInput, encoder *httpbinding.Encoder) error {
@@ -5318,10 +5152,8 @@ func (m *awsRestjson1_serializeOpGetStep) HandleSerial
 func (m *awsRestjson1_serializeOpGetStep) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5359,7 +5191,6 @@ func (m *awsRestjson1_serializeOpGetStep) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStepInput(v *GetStepInput, encoder *httpbinding.Encoder) error {
@@ -5416,10 +5247,8 @@ func (m *awsRestjson1_serializeOpGetStorageProfile) Ha
 func (m *awsRestjson1_serializeOpGetStorageProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5457,7 +5286,6 @@ func (m *awsRestjson1_serializeOpGetStorageProfile) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStorageProfileInput(v *GetStorageProfileInput, encoder *httpbinding.Encoder) error {
@@ -5496,10 +5324,8 @@ func (m *awsRestjson1_serializeOpGetStorageProfileForQ
 func (m *awsRestjson1_serializeOpGetStorageProfileForQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5537,7 +5363,6 @@ func (m *awsRestjson1_serializeOpGetStorageProfileForQ
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetStorageProfileForQueueInput(v *GetStorageProfileForQueueInput, encoder *httpbinding.Encoder) error {
@@ -5585,10 +5410,8 @@ func (m *awsRestjson1_serializeOpGetTask) HandleSerial
 func (m *awsRestjson1_serializeOpGetTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5626,7 +5449,6 @@ func (m *awsRestjson1_serializeOpGetTask) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetTaskInput(v *GetTaskInput, encoder *httpbinding.Encoder) error {
@@ -5692,10 +5514,8 @@ func (m *awsRestjson1_serializeOpGetWorker) HandleSeri
 func (m *awsRestjson1_serializeOpGetWorker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5733,7 +5553,6 @@ func (m *awsRestjson1_serializeOpGetWorker) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetWorkerInput(v *GetWorkerInput, encoder *httpbinding.Encoder) error {
@@ -5781,10 +5600,8 @@ func (m *awsRestjson1_serializeOpListAvailableMeteredP
 func (m *awsRestjson1_serializeOpListAvailableMeteredProducts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5822,7 +5639,6 @@ func (m *awsRestjson1_serializeOpListAvailableMeteredP
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAvailableMeteredProductsInput(v *ListAvailableMeteredProductsInput, encoder *httpbinding.Encoder) error {
@@ -5851,10 +5667,8 @@ func (m *awsRestjson1_serializeOpListBudgets) HandleSe
 func (m *awsRestjson1_serializeOpListBudgets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5892,7 +5706,6 @@ func (m *awsRestjson1_serializeOpListBudgets) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBudgetsInput(v *ListBudgetsInput, encoder *httpbinding.Encoder) error {
@@ -5934,10 +5747,8 @@ func (m *awsRestjson1_serializeOpListFarmMembers) Hand
 func (m *awsRestjson1_serializeOpListFarmMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5975,7 +5786,6 @@ func (m *awsRestjson1_serializeOpListFarmMembers) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFarmMembersInput(v *ListFarmMembersInput, encoder *httpbinding.Encoder) error {
@@ -6013,10 +5823,8 @@ func (m *awsRestjson1_serializeOpListFarms) HandleSeri
 func (m *awsRestjson1_serializeOpListFarms) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6054,7 +5862,6 @@ func (m *awsRestjson1_serializeOpListFarms) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFarmsInput(v *ListFarmsInput, encoder *httpbinding.Encoder) error {
@@ -6087,10 +5894,8 @@ func (m *awsRestjson1_serializeOpListFleetMembers) Han
 func (m *awsRestjson1_serializeOpListFleetMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6128,7 +5933,6 @@ func (m *awsRestjson1_serializeOpListFleetMembers) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFleetMembersInput(v *ListFleetMembersInput, encoder *httpbinding.Encoder) error {
@@ -6175,10 +5979,8 @@ func (m *awsRestjson1_serializeOpListFleets) HandleSer
 func (m *awsRestjson1_serializeOpListFleets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6216,7 +6018,6 @@ func (m *awsRestjson1_serializeOpListFleets) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListFleetsInput(v *ListFleetsInput, encoder *httpbinding.Encoder) error {
@@ -6266,10 +6067,8 @@ func (m *awsRestjson1_serializeOpListJobMembers) Handl
 func (m *awsRestjson1_serializeOpListJobMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6307,7 +6106,6 @@ func (m *awsRestjson1_serializeOpListJobMembers) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobMembersInput(v *ListJobMembersInput, encoder *httpbinding.Encoder) error {
@@ -6363,10 +6161,8 @@ func (m *awsRestjson1_serializeOpListJobParameterDefin
 func (m *awsRestjson1_serializeOpListJobParameterDefinitions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6404,7 +6200,6 @@ func (m *awsRestjson1_serializeOpListJobParameterDefin
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobParameterDefinitionsInput(v *ListJobParameterDefinitionsInput, encoder *httpbinding.Encoder) error {
@@ -6460,10 +6255,8 @@ func (m *awsRestjson1_serializeOpListJobs) HandleSeria
 func (m *awsRestjson1_serializeOpListJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6501,7 +6294,6 @@ func (m *awsRestjson1_serializeOpListJobs) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListJobsInput(v *ListJobsInput, encoder *httpbinding.Encoder) error {
@@ -6552,10 +6344,8 @@ func (m *awsRestjson1_serializeOpListLicenseEndpoints)
 func (m *awsRestjson1_serializeOpListLicenseEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6593,7 +6383,6 @@ func (m *awsRestjson1_serializeOpListLicenseEndpoints)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLicenseEndpointsInput(v *ListLicenseEndpointsInput, encoder *httpbinding.Encoder) error {
@@ -6622,10 +6411,8 @@ func (m *awsRestjson1_serializeOpListLimits) HandleSer
 func (m *awsRestjson1_serializeOpListLimits) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6663,7 +6450,6 @@ func (m *awsRestjson1_serializeOpListLimits) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLimitsInput(v *ListLimitsInput, encoder *httpbinding.Encoder) error {
@@ -6701,10 +6487,8 @@ func (m *awsRestjson1_serializeOpListMeteredProducts) 
 func (m *awsRestjson1_serializeOpListMeteredProducts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6742,7 +6526,6 @@ func (m *awsRestjson1_serializeOpListMeteredProducts) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMeteredProductsInput(v *ListMeteredProductsInput, encoder *httpbinding.Encoder) error {
@@ -6780,10 +6563,8 @@ func (m *awsRestjson1_serializeOpListMonitors) HandleS
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
@@ -6821,7 +6602,6 @@ func (m *awsRestjson1_serializeOpListMonitors) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListMonitorsInput(v *ListMonitorsInput, encoder *httpbinding.Encoder) error {
@@ -6850,10 +6630,8 @@ func (m *awsRestjson1_serializeOpListQueueEnvironments
 func (m *awsRestjson1_serializeOpListQueueEnvironments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6891,7 +6669,6 @@ func (m *awsRestjson1_serializeOpListQueueEnvironments
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueueEnvironmentsInput(v *ListQueueEnvironmentsInput, encoder *httpbinding.Encoder) error {
@@ -6938,10 +6715,8 @@ func (m *awsRestjson1_serializeOpListQueueFleetAssocia
 func (m *awsRestjson1_serializeOpListQueueFleetAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -6979,7 +6754,6 @@ func (m *awsRestjson1_serializeOpListQueueFleetAssocia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueueFleetAssociationsInput(v *ListQueueFleetAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -7025,10 +6799,8 @@ func (m *awsRestjson1_serializeOpListQueueLimitAssocia
 func (m *awsRestjson1_serializeOpListQueueLimitAssociations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7066,7 +6838,6 @@ func (m *awsRestjson1_serializeOpListQueueLimitAssocia
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueueLimitAssociationsInput(v *ListQueueLimitAssociationsInput, encoder *httpbinding.Encoder) error {
@@ -7112,10 +6883,8 @@ func (m *awsRestjson1_serializeOpListQueueMembers) Han
 func (m *awsRestjson1_serializeOpListQueueMembers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7153,7 +6922,6 @@ func (m *awsRestjson1_serializeOpListQueueMembers) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueueMembersInput(v *ListQueueMembersInput, encoder *httpbinding.Encoder) error {
@@ -7200,10 +6968,8 @@ func (m *awsRestjson1_serializeOpListQueues) HandleSer
 func (m *awsRestjson1_serializeOpListQueues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7241,7 +7007,6 @@ func (m *awsRestjson1_serializeOpListQueues) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListQueuesInput(v *ListQueuesInput, encoder *httpbinding.Encoder) error {
@@ -7287,10 +7052,8 @@ func (m *awsRestjson1_serializeOpListSessionActions) H
 func (m *awsRestjson1_serializeOpListSessionActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7328,7 +7091,6 @@ func (m *awsRestjson1_serializeOpListSessionActions) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSessionActionsInput(v *ListSessionActionsInput, encoder *httpbinding.Encoder) error {
@@ -7392,10 +7154,8 @@ func (m *awsRestjson1_serializeOpListSessions) HandleS
 func (m *awsRestjson1_serializeOpListSessions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7433,7 +7193,6 @@ func (m *awsRestjson1_serializeOpListSessions) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSessionsInput(v *ListSessionsInput, encoder *httpbinding.Encoder) error {
@@ -7489,10 +7248,8 @@ func (m *awsRestjson1_serializeOpListSessionsForWorker
 func (m *awsRestjson1_serializeOpListSessionsForWorker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7530,7 +7287,6 @@ func (m *awsRestjson1_serializeOpListSessionsForWorker
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSessionsForWorkerInput(v *ListSessionsForWorkerInput, encoder *httpbinding.Encoder) error {
@@ -7586,10 +7342,8 @@ func (m *awsRestjson1_serializeOpListStepConsumers) Ha
 func (m *awsRestjson1_serializeOpListStepConsumers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7627,7 +7381,6 @@ func (m *awsRestjson1_serializeOpListStepConsumers) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStepConsumersInput(v *ListStepConsumersInput, encoder *httpbinding.Encoder) error {
@@ -7692,10 +7445,8 @@ func (m *awsRestjson1_serializeOpListStepDependencies)
 func (m *awsRestjson1_serializeOpListStepDependencies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7733,7 +7484,6 @@ func (m *awsRestjson1_serializeOpListStepDependencies)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStepDependenciesInput(v *ListStepDependenciesInput, encoder *httpbinding.Encoder) error {
@@ -7798,10 +7548,8 @@ func (m *awsRestjson1_serializeOpListSteps) HandleSeri
 func (m *awsRestjson1_serializeOpListSteps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7839,7 +7587,6 @@ func (m *awsRestjson1_serializeOpListSteps) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStepsInput(v *ListStepsInput, encoder *httpbinding.Encoder) error {
@@ -7895,10 +7642,8 @@ func (m *awsRestjson1_serializeOpListStorageProfiles) 
 func (m *awsRestjson1_serializeOpListStorageProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -7936,7 +7681,6 @@ func (m *awsRestjson1_serializeOpListStorageProfiles) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStorageProfilesInput(v *ListStorageProfilesInput, encoder *httpbinding.Encoder) error {
@@ -7974,10 +7718,8 @@ func (m *awsRestjson1_serializeOpListStorageProfilesFo
 func (m *awsRestjson1_serializeOpListStorageProfilesForQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8015,7 +7757,6 @@ func (m *awsRestjson1_serializeOpListStorageProfilesFo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStorageProfilesForQueueInput(v *ListStorageProfilesForQueueInput, encoder *httpbinding.Encoder) error {
@@ -8062,10 +7803,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -8103,7 +7842,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -8133,10 +7871,8 @@ func (m *awsRestjson1_serializeOpListTasks) HandleSeri
 func (m *awsRestjson1_serializeOpListTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8174,7 +7910,6 @@ func (m *awsRestjson1_serializeOpListTasks) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTasksInput(v *ListTasksInput, encoder *httpbinding.Encoder) error {
@@ -8239,10 +7974,8 @@ func (m *awsRestjson1_serializeOpListWorkers) HandleSe
 func (m *awsRestjson1_serializeOpListWorkers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8280,7 +8013,6 @@ func (m *awsRestjson1_serializeOpListWorkers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListWorkersInput(v *ListWorkersInput, encoder *httpbinding.Encoder) error {
@@ -8327,10 +8059,8 @@ func (m *awsRestjson1_serializeOpPutMeteredProduct) Ha
 func (m *awsRestjson1_serializeOpPutMeteredProduct) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8368,7 +8098,6 @@ func (m *awsRestjson1_serializeOpPutMeteredProduct) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutMeteredProductInput(v *PutMeteredProductInput, encoder *httpbinding.Encoder) error {
@@ -8407,10 +8136,8 @@ func (m *awsRestjson1_serializeOpSearchJobs) HandleSer
 func (m *awsRestjson1_serializeOpSearchJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8459,7 +8186,6 @@ func (m *awsRestjson1_serializeOpSearchJobs) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchJobsInput(v *SearchJobsInput, encoder *httpbinding.Encoder) error {
@@ -8527,10 +8253,8 @@ func (m *awsRestjson1_serializeOpSearchSteps) HandleSe
 func (m *awsRestjson1_serializeOpSearchSteps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8579,7 +8303,6 @@ func (m *awsRestjson1_serializeOpSearchSteps) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchStepsInput(v *SearchStepsInput, encoder *httpbinding.Encoder) error {
@@ -8652,10 +8375,8 @@ func (m *awsRestjson1_serializeOpSearchTasks) HandleSe
 func (m *awsRestjson1_serializeOpSearchTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8704,7 +8425,6 @@ func (m *awsRestjson1_serializeOpSearchTasks) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchTasksInput(v *SearchTasksInput, encoder *httpbinding.Encoder) error {
@@ -8777,10 +8497,8 @@ func (m *awsRestjson1_serializeOpSearchWorkers) Handle
 func (m *awsRestjson1_serializeOpSearchWorkers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8829,7 +8547,6 @@ func (m *awsRestjson1_serializeOpSearchWorkers) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchWorkersInput(v *SearchWorkersInput, encoder *httpbinding.Encoder) error {
@@ -8897,10 +8614,8 @@ func (m *awsRestjson1_serializeOpStartSessionsStatisti
 func (m *awsRestjson1_serializeOpStartSessionsStatisticsAggregation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -8949,7 +8664,6 @@ func (m *awsRestjson1_serializeOpStartSessionsStatisti
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartSessionsStatisticsAggregationInput(v *StartSessionsStatisticsAggregationInput, encoder *httpbinding.Encoder) error {
@@ -9027,10 +8741,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -9079,7 +8791,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -9123,10 +8834,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -9164,7 +8873,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -9200,10 +8908,8 @@ func (m *awsRestjson1_serializeOpUpdateBudget) HandleS
 func (m *awsRestjson1_serializeOpUpdateBudget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9252,7 +8958,6 @@ func (m *awsRestjson1_serializeOpUpdateBudget) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBudgetInput(v *UpdateBudgetInput, encoder *httpbinding.Encoder) error {
@@ -9357,10 +9062,8 @@ func (m *awsRestjson1_serializeOpUpdateFarm) HandleSer
 func (m *awsRestjson1_serializeOpUpdateFarm) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9409,7 +9112,6 @@ func (m *awsRestjson1_serializeOpUpdateFarm) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFarmInput(v *UpdateFarmInput, encoder *httpbinding.Encoder) error {
@@ -9456,10 +9158,8 @@ func (m *awsRestjson1_serializeOpUpdateFleet) HandleSe
 func (m *awsRestjson1_serializeOpUpdateFleet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9508,7 +9208,6 @@ func (m *awsRestjson1_serializeOpUpdateFleet) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFleetInput(v *UpdateFleetInput, encoder *httpbinding.Encoder) error {
@@ -9598,10 +9297,8 @@ func (m *awsRestjson1_serializeOpUpdateJob) HandleSeri
 func (m *awsRestjson1_serializeOpUpdateJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9650,7 +9347,6 @@ func (m *awsRestjson1_serializeOpUpdateJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateJobInput(v *UpdateJobInput, encoder *httpbinding.Encoder) error {
@@ -9740,10 +9436,8 @@ func (m *awsRestjson1_serializeOpUpdateLimit) HandleSe
 func (m *awsRestjson1_serializeOpUpdateLimit) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -9792,7 +9486,6 @@ func (m *awsRestjson1_serializeOpUpdateLimit) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLimitInput(v *UpdateLimitInput, encoder *httpbinding.Encoder) error {
@@ -9853,10 +9546,8 @@ func (m *awsRestjson1_serializeOpUpdateMonitor) Handle
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
@@ -9905,7 +9596,6 @@ func (m *awsRestjson1_serializeOpUpdateMonitor) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateMonitorInput(v *UpdateMonitorInput, encoder *httpbinding.Encoder) error {
@@ -9957,10 +9647,8 @@ func (m *awsRestjson1_serializeOpUpdateQueue) HandleSe
 func (m *awsRestjson1_serializeOpUpdateQueue) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10009,7 +9697,6 @@ func (m *awsRestjson1_serializeOpUpdateQueue) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueInput(v *UpdateQueueInput, encoder *httpbinding.Encoder) error {
@@ -10122,10 +9809,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueEnvironmen
 func (m *awsRestjson1_serializeOpUpdateQueueEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10174,7 +9859,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueEnvironmen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueEnvironmentInput(v *UpdateQueueEnvironmentInput, encoder *httpbinding.Encoder) error {
@@ -10249,10 +9933,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueFleetAssoc
 func (m *awsRestjson1_serializeOpUpdateQueueFleetAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10301,7 +9983,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueFleetAssoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueFleetAssociationInput(v *UpdateQueueFleetAssociationInput, encoder *httpbinding.Encoder) error {
@@ -10361,10 +10042,8 @@ func (m *awsRestjson1_serializeOpUpdateQueueLimitAssoc
 func (m *awsRestjson1_serializeOpUpdateQueueLimitAssociation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10413,7 +10092,6 @@ func (m *awsRestjson1_serializeOpUpdateQueueLimitAssoc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateQueueLimitAssociationInput(v *UpdateQueueLimitAssociationInput, encoder *httpbinding.Encoder) error {
@@ -10473,10 +10151,8 @@ func (m *awsRestjson1_serializeOpUpdateSession) Handle
 func (m *awsRestjson1_serializeOpUpdateSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10525,7 +10201,6 @@ func (m *awsRestjson1_serializeOpUpdateSession) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSessionInput(v *UpdateSessionInput, encoder *httpbinding.Encoder) error {
@@ -10599,10 +10274,8 @@ func (m *awsRestjson1_serializeOpUpdateStep) HandleSer
 func (m *awsRestjson1_serializeOpUpdateStep) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10651,7 +10324,6 @@ func (m *awsRestjson1_serializeOpUpdateStep) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateStepInput(v *UpdateStepInput, encoder *httpbinding.Encoder) error {
@@ -10725,10 +10397,8 @@ func (m *awsRestjson1_serializeOpUpdateStorageProfile)
 func (m *awsRestjson1_serializeOpUpdateStorageProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10777,7 +10447,6 @@ func (m *awsRestjson1_serializeOpUpdateStorageProfile)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateStorageProfileInput(v *UpdateStorageProfileInput, encoder *httpbinding.Encoder) error {
@@ -10852,10 +10521,8 @@ func (m *awsRestjson1_serializeOpUpdateTask) HandleSer
 func (m *awsRestjson1_serializeOpUpdateTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -10904,7 +10571,6 @@ func (m *awsRestjson1_serializeOpUpdateTask) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateTaskInput(v *UpdateTaskInput, encoder *httpbinding.Encoder) error {
@@ -10987,10 +10653,8 @@ func (m *awsRestjson1_serializeOpUpdateWorker) HandleS
 func (m *awsRestjson1_serializeOpUpdateWorker) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11039,7 +10703,6 @@ func (m *awsRestjson1_serializeOpUpdateWorker) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWorkerInput(v *UpdateWorkerInput, encoder *httpbinding.Encoder) error {
@@ -11113,10 +10776,8 @@ func (m *awsRestjson1_serializeOpUpdateWorkerSchedule)
 func (m *awsRestjson1_serializeOpUpdateWorkerSchedule) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -11165,7 +10826,6 @@ func (m *awsRestjson1_serializeOpUpdateWorkerSchedule)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateWorkerScheduleInput(v *UpdateWorkerScheduleInput, encoder *httpbinding.Encoder) error {
