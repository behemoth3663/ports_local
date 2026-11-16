--- vendor/github.com/aws/aws-sdk-go-v2/service/auditmanager/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/auditmanager/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociateAssessmentRe
 func (m *awsRestjson1_serializeOpAssociateAssessmentReportEvidenceFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpAssociateAssessmentRe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateAssessmentReportEvidenceFolderInput(v *AssociateAssessmentReportEvidenceFolderInput, encoder *httpbinding.Encoder) error {
@@ -119,10 +115,8 @@ func (m *awsRestjson1_serializeOpBatchAssociateAssessm
 func (m *awsRestjson1_serializeOpBatchAssociateAssessmentReportEvidence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -171,7 +165,6 @@ func (m *awsRestjson1_serializeOpBatchAssociateAssessm
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchAssociateAssessmentReportEvidenceInput(v *BatchAssociateAssessmentReportEvidenceInput, encoder *httpbinding.Encoder) error {
@@ -220,10 +213,8 @@ func (m *awsRestjson1_serializeOpBatchCreateDelegation
 func (m *awsRestjson1_serializeOpBatchCreateDelegationByAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -272,7 +263,6 @@ func (m *awsRestjson1_serializeOpBatchCreateDelegation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchCreateDelegationByAssessmentInput(v *BatchCreateDelegationByAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -316,10 +306,8 @@ func (m *awsRestjson1_serializeOpBatchDeleteDelegation
 func (m *awsRestjson1_serializeOpBatchDeleteDelegationByAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -368,7 +356,6 @@ func (m *awsRestjson1_serializeOpBatchDeleteDelegation
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeleteDelegationByAssessmentInput(v *BatchDeleteDelegationByAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -412,10 +399,8 @@ func (m *awsRestjson1_serializeOpBatchDisassociateAsse
 func (m *awsRestjson1_serializeOpBatchDisassociateAssessmentReportEvidence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -464,7 +449,6 @@ func (m *awsRestjson1_serializeOpBatchDisassociateAsse
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDisassociateAssessmentReportEvidenceInput(v *BatchDisassociateAssessmentReportEvidenceInput, encoder *httpbinding.Encoder) error {
@@ -513,10 +497,8 @@ func (m *awsRestjson1_serializeOpBatchImportEvidenceTo
 func (m *awsRestjson1_serializeOpBatchImportEvidenceToAssessmentControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -565,7 +547,6 @@ func (m *awsRestjson1_serializeOpBatchImportEvidenceTo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchImportEvidenceToAssessmentControlInput(v *BatchImportEvidenceToAssessmentControlInput, encoder *httpbinding.Encoder) error {
@@ -627,10 +608,8 @@ func (m *awsRestjson1_serializeOpCreateAssessment) Han
 func (m *awsRestjson1_serializeOpCreateAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -675,7 +654,6 @@ func (m *awsRestjson1_serializeOpCreateAssessment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssessmentInput(v *CreateAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -746,10 +724,8 @@ func (m *awsRestjson1_serializeOpCreateAssessmentFrame
 func (m *awsRestjson1_serializeOpCreateAssessmentFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -794,7 +770,6 @@ func (m *awsRestjson1_serializeOpCreateAssessmentFrame
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssessmentFrameworkInput(v *CreateAssessmentFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -851,10 +826,8 @@ func (m *awsRestjson1_serializeOpCreateAssessmentRepor
 func (m *awsRestjson1_serializeOpCreateAssessmentReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -903,7 +876,6 @@ func (m *awsRestjson1_serializeOpCreateAssessmentRepor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAssessmentReportInput(v *CreateAssessmentReportInput, encoder *httpbinding.Encoder) error {
@@ -955,10 +927,8 @@ func (m *awsRestjson1_serializeOpCreateControl) Handle
 func (m *awsRestjson1_serializeOpCreateControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1003,7 +973,6 @@ func (m *awsRestjson1_serializeOpCreateControl) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateControlInput(v *CreateControlInput, encoder *httpbinding.Encoder) error {
@@ -1070,10 +1039,8 @@ func (m *awsRestjson1_serializeOpDeleteAssessment) Han
 func (m *awsRestjson1_serializeOpDeleteAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1111,7 +1078,6 @@ func (m *awsRestjson1_serializeOpDeleteAssessment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssessmentInput(v *DeleteAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -1141,10 +1107,8 @@ func (m *awsRestjson1_serializeOpDeleteAssessmentFrame
 func (m *awsRestjson1_serializeOpDeleteAssessmentFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1182,7 +1146,6 @@ func (m *awsRestjson1_serializeOpDeleteAssessmentFrame
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssessmentFrameworkInput(v *DeleteAssessmentFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -1212,10 +1175,8 @@ func (m *awsRestjson1_serializeOpDeleteAssessmentFrame
 func (m *awsRestjson1_serializeOpDeleteAssessmentFrameworkShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1253,7 +1214,6 @@ func (m *awsRestjson1_serializeOpDeleteAssessmentFrame
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssessmentFrameworkShareInput(v *DeleteAssessmentFrameworkShareInput, encoder *httpbinding.Encoder) error {
@@ -1287,10 +1247,8 @@ func (m *awsRestjson1_serializeOpDeleteAssessmentRepor
 func (m *awsRestjson1_serializeOpDeleteAssessmentReport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1328,7 +1286,6 @@ func (m *awsRestjson1_serializeOpDeleteAssessmentRepor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAssessmentReportInput(v *DeleteAssessmentReportInput, encoder *httpbinding.Encoder) error {
@@ -1367,10 +1324,8 @@ func (m *awsRestjson1_serializeOpDeleteControl) Handle
 func (m *awsRestjson1_serializeOpDeleteControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1408,7 +1363,6 @@ func (m *awsRestjson1_serializeOpDeleteControl) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteControlInput(v *DeleteControlInput, encoder *httpbinding.Encoder) error {
@@ -1438,10 +1392,8 @@ func (m *awsRestjson1_serializeOpDeregisterAccount) Ha
 func (m *awsRestjson1_serializeOpDeregisterAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1475,7 +1427,6 @@ func (m *awsRestjson1_serializeOpDeregisterAccount) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeregisterAccountInput(v *DeregisterAccountInput, encoder *httpbinding.Encoder) error {
@@ -1496,10 +1447,8 @@ func (m *awsRestjson1_serializeOpDeregisterOrganizatio
 func (m *awsRestjson1_serializeOpDeregisterOrganizationAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1544,7 +1493,6 @@ func (m *awsRestjson1_serializeOpDeregisterOrganizatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeregisterOrganizationAdminAccountInput(v *DeregisterOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -1577,10 +1525,8 @@ func (m *awsRestjson1_serializeOpDisassociateAssessmen
 func (m *awsRestjson1_serializeOpDisassociateAssessmentReportEvidenceFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1629,7 +1575,6 @@ func (m *awsRestjson1_serializeOpDisassociateAssessmen
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateAssessmentReportEvidenceFolderInput(v *DisassociateAssessmentReportEvidenceFolderInput, encoder *httpbinding.Encoder) error {
@@ -1671,10 +1616,8 @@ func (m *awsRestjson1_serializeOpGetAccountStatus) Han
 func (m *awsRestjson1_serializeOpGetAccountStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1708,7 +1651,6 @@ func (m *awsRestjson1_serializeOpGetAccountStatus) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountStatusInput(v *GetAccountStatusInput, encoder *httpbinding.Encoder) error {
@@ -1729,10 +1671,8 @@ func (m *awsRestjson1_serializeOpGetAssessment) Handle
 func (m *awsRestjson1_serializeOpGetAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1770,7 +1710,6 @@ func (m *awsRestjson1_serializeOpGetAssessment) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssessmentInput(v *GetAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -1800,10 +1739,8 @@ func (m *awsRestjson1_serializeOpGetAssessmentFramewor
 func (m *awsRestjson1_serializeOpGetAssessmentFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1841,7 +1778,6 @@ func (m *awsRestjson1_serializeOpGetAssessmentFramewor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssessmentFrameworkInput(v *GetAssessmentFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -1871,10 +1807,8 @@ func (m *awsRestjson1_serializeOpGetAssessmentReportUr
 func (m *awsRestjson1_serializeOpGetAssessmentReportUrl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1912,7 +1846,6 @@ func (m *awsRestjson1_serializeOpGetAssessmentReportUr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAssessmentReportUrlInput(v *GetAssessmentReportUrlInput, encoder *httpbinding.Encoder) error {
@@ -1951,10 +1884,8 @@ func (m *awsRestjson1_serializeOpGetChangeLogs) Handle
 func (m *awsRestjson1_serializeOpGetChangeLogs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1992,7 +1923,6 @@ func (m *awsRestjson1_serializeOpGetChangeLogs) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetChangeLogsInput(v *GetChangeLogsInput, encoder *httpbinding.Encoder) error {
@@ -2038,10 +1968,8 @@ func (m *awsRestjson1_serializeOpGetControl) HandleSer
 func (m *awsRestjson1_serializeOpGetControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2079,7 +2007,6 @@ func (m *awsRestjson1_serializeOpGetControl) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetControlInput(v *GetControlInput, encoder *httpbinding.Encoder) error {
@@ -2109,10 +2036,8 @@ func (m *awsRestjson1_serializeOpGetDelegations) Handl
 func (m *awsRestjson1_serializeOpGetDelegations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2150,7 +2075,6 @@ func (m *awsRestjson1_serializeOpGetDelegations) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetDelegationsInput(v *GetDelegationsInput, encoder *httpbinding.Encoder) error {
@@ -2179,10 +2103,8 @@ func (m *awsRestjson1_serializeOpGetEvidence) HandleSe
 func (m *awsRestjson1_serializeOpGetEvidence) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2220,7 +2142,6 @@ func (m *awsRestjson1_serializeOpGetEvidence) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEvidenceInput(v *GetEvidenceInput, encoder *httpbinding.Encoder) error {
@@ -2277,10 +2198,8 @@ func (m *awsRestjson1_serializeOpGetEvidenceByEvidence
 func (m *awsRestjson1_serializeOpGetEvidenceByEvidenceFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2318,7 +2237,6 @@ func (m *awsRestjson1_serializeOpGetEvidenceByEvidence
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEvidenceByEvidenceFolderInput(v *GetEvidenceByEvidenceFolderInput, encoder *httpbinding.Encoder) error {
@@ -2374,10 +2292,8 @@ func (m *awsRestjson1_serializeOpGetEvidenceFileUpload
 func (m *awsRestjson1_serializeOpGetEvidenceFileUploadUrl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2415,7 +2331,6 @@ func (m *awsRestjson1_serializeOpGetEvidenceFileUpload
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEvidenceFileUploadUrlInput(v *GetEvidenceFileUploadUrlInput, encoder *httpbinding.Encoder) error {
@@ -2440,10 +2355,8 @@ func (m *awsRestjson1_serializeOpGetEvidenceFolder) Ha
 func (m *awsRestjson1_serializeOpGetEvidenceFolder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2481,7 +2394,6 @@ func (m *awsRestjson1_serializeOpGetEvidenceFolder) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEvidenceFolderInput(v *GetEvidenceFolderInput, encoder *httpbinding.Encoder) error {
@@ -2529,10 +2441,8 @@ func (m *awsRestjson1_serializeOpGetEvidenceFoldersByA
 func (m *awsRestjson1_serializeOpGetEvidenceFoldersByAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2570,7 +2480,6 @@ func (m *awsRestjson1_serializeOpGetEvidenceFoldersByA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEvidenceFoldersByAssessmentInput(v *GetEvidenceFoldersByAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -2608,10 +2517,8 @@ func (m *awsRestjson1_serializeOpGetEvidenceFoldersByA
 func (m *awsRestjson1_serializeOpGetEvidenceFoldersByAssessmentControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2649,7 +2556,6 @@ func (m *awsRestjson1_serializeOpGetEvidenceFoldersByA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEvidenceFoldersByAssessmentControlInput(v *GetEvidenceFoldersByAssessmentControlInput, encoder *httpbinding.Encoder) error {
@@ -2705,10 +2611,8 @@ func (m *awsRestjson1_serializeOpGetInsights) HandleSe
 func (m *awsRestjson1_serializeOpGetInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2742,7 +2646,6 @@ func (m *awsRestjson1_serializeOpGetInsights) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInsightsInput(v *GetInsightsInput, encoder *httpbinding.Encoder) error {
@@ -2763,10 +2666,8 @@ func (m *awsRestjson1_serializeOpGetInsightsByAssessme
 func (m *awsRestjson1_serializeOpGetInsightsByAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2804,7 +2705,6 @@ func (m *awsRestjson1_serializeOpGetInsightsByAssessme
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetInsightsByAssessmentInput(v *GetInsightsByAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -2834,10 +2734,8 @@ func (m *awsRestjson1_serializeOpGetOrganizationAdminA
 func (m *awsRestjson1_serializeOpGetOrganizationAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2871,7 +2769,6 @@ func (m *awsRestjson1_serializeOpGetOrganizationAdminA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetOrganizationAdminAccountInput(v *GetOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -2892,10 +2789,8 @@ func (m *awsRestjson1_serializeOpGetServicesInScope) H
 func (m *awsRestjson1_serializeOpGetServicesInScope) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2929,7 +2824,6 @@ func (m *awsRestjson1_serializeOpGetServicesInScope) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetServicesInScopeInput(v *GetServicesInScopeInput, encoder *httpbinding.Encoder) error {
@@ -2950,10 +2844,8 @@ func (m *awsRestjson1_serializeOpGetSettings) HandleSe
 func (m *awsRestjson1_serializeOpGetSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2991,7 +2883,6 @@ func (m *awsRestjson1_serializeOpGetSettings) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetSettingsInput(v *GetSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3021,10 +2912,8 @@ func (m *awsRestjson1_serializeOpListAssessmentControl
 func (m *awsRestjson1_serializeOpListAssessmentControlInsightsByControlDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3062,7 +2951,6 @@ func (m *awsRestjson1_serializeOpListAssessmentControl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssessmentControlInsightsByControlDomainInput(v *ListAssessmentControlInsightsByControlDomainInput, encoder *httpbinding.Encoder) error {
@@ -3099,10 +2987,8 @@ func (m *awsRestjson1_serializeOpListAssessmentFramewo
 func (m *awsRestjson1_serializeOpListAssessmentFrameworks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3140,7 +3026,6 @@ func (m *awsRestjson1_serializeOpListAssessmentFramewo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssessmentFrameworksInput(v *ListAssessmentFrameworksInput, encoder *httpbinding.Encoder) error {
@@ -3173,10 +3058,8 @@ func (m *awsRestjson1_serializeOpListAssessmentFramewo
 func (m *awsRestjson1_serializeOpListAssessmentFrameworkShareRequests) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3214,7 +3097,6 @@ func (m *awsRestjson1_serializeOpListAssessmentFramewo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssessmentFrameworkShareRequestsInput(v *ListAssessmentFrameworkShareRequestsInput, encoder *httpbinding.Encoder) error {
@@ -3247,10 +3129,8 @@ func (m *awsRestjson1_serializeOpListAssessmentReports
 func (m *awsRestjson1_serializeOpListAssessmentReports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3288,7 +3168,6 @@ func (m *awsRestjson1_serializeOpListAssessmentReports
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssessmentReportsInput(v *ListAssessmentReportsInput, encoder *httpbinding.Encoder) error {
@@ -3317,10 +3196,8 @@ func (m *awsRestjson1_serializeOpListAssessments) Hand
 func (m *awsRestjson1_serializeOpListAssessments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3358,7 +3235,6 @@ func (m *awsRestjson1_serializeOpListAssessments) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAssessmentsInput(v *ListAssessmentsInput, encoder *httpbinding.Encoder) error {
@@ -3391,10 +3267,8 @@ func (m *awsRestjson1_serializeOpListControlDomainInsi
 func (m *awsRestjson1_serializeOpListControlDomainInsights) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3432,7 +3306,6 @@ func (m *awsRestjson1_serializeOpListControlDomainInsi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListControlDomainInsightsInput(v *ListControlDomainInsightsInput, encoder *httpbinding.Encoder) error {
@@ -3461,10 +3334,8 @@ func (m *awsRestjson1_serializeOpListControlDomainInsi
 func (m *awsRestjson1_serializeOpListControlDomainInsightsByAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3502,7 +3373,6 @@ func (m *awsRestjson1_serializeOpListControlDomainInsi
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListControlDomainInsightsByAssessmentInput(v *ListControlDomainInsightsByAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -3535,10 +3405,8 @@ func (m *awsRestjson1_serializeOpListControlInsightsBy
 func (m *awsRestjson1_serializeOpListControlInsightsByControlDomain) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3576,7 +3444,6 @@ func (m *awsRestjson1_serializeOpListControlInsightsBy
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListControlInsightsByControlDomainInput(v *ListControlInsightsByControlDomainInput, encoder *httpbinding.Encoder) error {
@@ -3609,10 +3476,8 @@ func (m *awsRestjson1_serializeOpListControls) HandleS
 func (m *awsRestjson1_serializeOpListControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3650,7 +3515,6 @@ func (m *awsRestjson1_serializeOpListControls) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListControlsInput(v *ListControlsInput, encoder *httpbinding.Encoder) error {
@@ -3687,10 +3551,8 @@ func (m *awsRestjson1_serializeOpListKeywordsForDataSo
 func (m *awsRestjson1_serializeOpListKeywordsForDataSource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3728,7 +3590,6 @@ func (m *awsRestjson1_serializeOpListKeywordsForDataSo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListKeywordsForDataSourceInput(v *ListKeywordsForDataSourceInput, encoder *httpbinding.Encoder) error {
@@ -3761,10 +3622,8 @@ func (m *awsRestjson1_serializeOpListNotifications) Ha
 func (m *awsRestjson1_serializeOpListNotifications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3802,7 +3661,6 @@ func (m *awsRestjson1_serializeOpListNotifications) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListNotificationsInput(v *ListNotificationsInput, encoder *httpbinding.Encoder) error {
@@ -3831,10 +3689,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
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
@@ -3872,7 +3728,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -3902,10 +3757,8 @@ func (m *awsRestjson1_serializeOpRegisterAccount) Hand
 func (m *awsRestjson1_serializeOpRegisterAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3950,7 +3803,6 @@ func (m *awsRestjson1_serializeOpRegisterAccount) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterAccountInput(v *RegisterAccountInput, encoder *httpbinding.Encoder) error {
@@ -3988,10 +3840,8 @@ func (m *awsRestjson1_serializeOpRegisterOrganizationA
 func (m *awsRestjson1_serializeOpRegisterOrganizationAdminAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4036,7 +3886,6 @@ func (m *awsRestjson1_serializeOpRegisterOrganizationA
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegisterOrganizationAdminAccountInput(v *RegisterOrganizationAdminAccountInput, encoder *httpbinding.Encoder) error {
@@ -4069,10 +3918,8 @@ func (m *awsRestjson1_serializeOpStartAssessmentFramew
 func (m *awsRestjson1_serializeOpStartAssessmentFrameworkShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4121,7 +3968,6 @@ func (m *awsRestjson1_serializeOpStartAssessmentFramew
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartAssessmentFrameworkShareInput(v *StartAssessmentFrameworkShareInput, encoder *httpbinding.Encoder) error {
@@ -4173,10 +4019,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
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
@@ -4225,7 +4069,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4269,10 +4112,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
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
@@ -4310,7 +4151,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4346,10 +4186,8 @@ func (m *awsRestjson1_serializeOpUpdateAssessment) Han
 func (m *awsRestjson1_serializeOpUpdateAssessment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4398,7 +4236,6 @@ func (m *awsRestjson1_serializeOpUpdateAssessment) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssessmentInput(v *UpdateAssessmentInput, encoder *httpbinding.Encoder) error {
@@ -4466,10 +4303,8 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentContr
 func (m *awsRestjson1_serializeOpUpdateAssessmentControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4518,7 +4353,6 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssessmentControlInput(v *UpdateAssessmentControlInput, encoder *httpbinding.Encoder) error {
@@ -4583,10 +4417,8 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentContr
 func (m *awsRestjson1_serializeOpUpdateAssessmentControlSetStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4635,7 +4467,6 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentContr
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssessmentControlSetStatusInput(v *UpdateAssessmentControlSetStatusInput, encoder *httpbinding.Encoder) error {
@@ -4691,10 +4522,8 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentFrame
 func (m *awsRestjson1_serializeOpUpdateAssessmentFramework) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4743,7 +4572,6 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentFrame
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssessmentFrameworkInput(v *UpdateAssessmentFrameworkInput, encoder *httpbinding.Encoder) error {
@@ -4802,10 +4630,8 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentFrame
 func (m *awsRestjson1_serializeOpUpdateAssessmentFrameworkShare) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4854,7 +4680,6 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentFrame
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssessmentFrameworkShareInput(v *UpdateAssessmentFrameworkShareInput, encoder *httpbinding.Encoder) error {
@@ -4901,10 +4726,8 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentStatu
 func (m *awsRestjson1_serializeOpUpdateAssessmentStatus) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4953,7 +4776,6 @@ func (m *awsRestjson1_serializeOpUpdateAssessmentStatu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAssessmentStatusInput(v *UpdateAssessmentStatusInput, encoder *httpbinding.Encoder) error {
@@ -4995,10 +4817,8 @@ func (m *awsRestjson1_serializeOpUpdateControl) Handle
 func (m *awsRestjson1_serializeOpUpdateControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5047,7 +4867,6 @@ func (m *awsRestjson1_serializeOpUpdateControl) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateControlInput(v *UpdateControlInput, encoder *httpbinding.Encoder) error {
@@ -5116,10 +4935,8 @@ func (m *awsRestjson1_serializeOpUpdateSettings) Handl
 func (m *awsRestjson1_serializeOpUpdateSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5164,7 +4981,6 @@ func (m *awsRestjson1_serializeOpUpdateSettings) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateSettingsInput(v *UpdateSettingsInput, encoder *httpbinding.Encoder) error {
@@ -5235,10 +5051,8 @@ func (m *awsRestjson1_serializeOpValidateAssessmentRep
 func (m *awsRestjson1_serializeOpValidateAssessmentReportIntegrity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5283,7 +5097,6 @@ func (m *awsRestjson1_serializeOpValidateAssessmentRep
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsValidateAssessmentReportIntegrityInput(v *ValidateAssessmentReportIntegrityInput, encoder *httpbinding.Encoder) error {
