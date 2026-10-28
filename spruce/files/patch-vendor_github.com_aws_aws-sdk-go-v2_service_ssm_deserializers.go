--- vendor/github.com/aws/aws-sdk-go-v2/service/ssm/deserializers.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ssm/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -47,11 +46,8 @@ func (m *awsAwsjson11_deserializeOpAddTagsToResource) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorAddTagsToResource(response, &metadata)
@@ -178,11 +174,8 @@ func (m *awsAwsjson11_deserializeOpAssociateOpsItemRel
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorAssociateOpsItemRelatedItem(response, &metadata)
@@ -312,11 +305,8 @@ func (m *awsAwsjson11_deserializeOpCancelCommand) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCancelCommand(response, &metadata)
@@ -440,11 +430,8 @@ func (m *awsAwsjson11_deserializeOpCancelMaintenanceWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCancelMaintenanceWindowExecution(response, &metadata)
@@ -562,11 +549,8 @@ func (m *awsAwsjson11_deserializeOpCreateActivation) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateActivation(response, &metadata)
@@ -684,11 +668,8 @@ func (m *awsAwsjson11_deserializeOpCreateAssociation) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateAssociation(response, &metadata)
@@ -839,11 +820,8 @@ func (m *awsAwsjson11_deserializeOpCreateAssociationBa
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateAssociationBatch(response, &metadata)
@@ -991,11 +969,8 @@ func (m *awsAwsjson11_deserializeOpCreateCloudConnecto
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateCloudConnector(response, &metadata)
@@ -1116,11 +1091,8 @@ func (m *awsAwsjson11_deserializeOpCreateDocument) Han
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateDocument(response, &metadata)
@@ -1256,11 +1228,8 @@ func (m *awsAwsjson11_deserializeOpCreateMaintenanceWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateMaintenanceWindow(response, &metadata)
@@ -1381,11 +1350,8 @@ func (m *awsAwsjson11_deserializeOpCreateOpsItem) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateOpsItem(response, &metadata)
@@ -1512,11 +1478,8 @@ func (m *awsAwsjson11_deserializeOpCreateOpsMetadata) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateOpsMetadata(response, &metadata)
@@ -1643,11 +1606,8 @@ func (m *awsAwsjson11_deserializeOpCreatePatchBaseline
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreatePatchBaseline(response, &metadata)
@@ -1768,11 +1728,8 @@ func (m *awsAwsjson11_deserializeOpCreateResourceDataS
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorCreateResourceDataSync(response, &metadata)
@@ -1896,11 +1853,8 @@ func (m *awsAwsjson11_deserializeOpDeleteActivation) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteActivation(response, &metadata)
@@ -2024,11 +1978,8 @@ func (m *awsAwsjson11_deserializeOpDeleteAssociation) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteAssociation(response, &metadata)
@@ -2155,11 +2106,8 @@ func (m *awsAwsjson11_deserializeOpDeleteCloudConnecto
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteCloudConnector(response, &metadata)
@@ -2280,11 +2228,8 @@ func (m *awsAwsjson11_deserializeOpDeleteDocument) Han
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteDocument(response, &metadata)
@@ -2411,11 +2356,8 @@ func (m *awsAwsjson11_deserializeOpDeleteInventory) Ha
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteInventory(response, &metadata)
@@ -2542,11 +2484,8 @@ func (m *awsAwsjson11_deserializeOpDeleteMaintenanceWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteMaintenanceWindow(response, &metadata)
@@ -2661,11 +2600,8 @@ func (m *awsAwsjson11_deserializeOpDeleteOpsItem) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteOpsItem(response, &metadata)
@@ -2783,11 +2719,8 @@ func (m *awsAwsjson11_deserializeOpDeleteOpsMetadata) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteOpsMetadata(response, &metadata)
@@ -2908,11 +2841,8 @@ func (m *awsAwsjson11_deserializeOpDeleteParameter) Ha
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteParameter(response, &metadata)
@@ -3030,11 +2960,8 @@ func (m *awsAwsjson11_deserializeOpDeleteParameters) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteParameters(response, &metadata)
@@ -3149,11 +3076,8 @@ func (m *awsAwsjson11_deserializeOpDeletePatchBaseline
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeletePatchBaseline(response, &metadata)
@@ -3271,11 +3195,8 @@ func (m *awsAwsjson11_deserializeOpDeleteResourceDataS
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteResourceDataSync(response, &metadata)
@@ -3396,11 +3317,8 @@ func (m *awsAwsjson11_deserializeOpDeleteResourcePolic
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteResourcePolicy(response, &metadata)
@@ -3530,11 +3448,8 @@ func (m *awsAwsjson11_deserializeOpDeregisterManagedIn
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeregisterManagedInstance(response, &metadata)
@@ -3652,11 +3567,8 @@ func (m *awsAwsjson11_deserializeOpDeregisterPatchBase
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeregisterPatchBaselineForPatchGroup(response, &metadata)
@@ -3774,11 +3686,8 @@ func (m *awsAwsjson11_deserializeOpDeregisterTargetFro
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeregisterTargetFromMaintenanceWindow(response, &metadata)
@@ -3899,11 +3808,8 @@ func (m *awsAwsjson11_deserializeOpDeregisterTaskFromM
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDeregisterTaskFromMaintenanceWindow(response, &metadata)
@@ -4021,11 +3927,8 @@ func (m *awsAwsjson11_deserializeOpDescribeActivations
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeActivations(response, &metadata)
@@ -4146,11 +4049,8 @@ func (m *awsAwsjson11_deserializeOpDescribeAssociation
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeAssociation(response, &metadata)
@@ -4277,11 +4177,8 @@ func (m *awsAwsjson11_deserializeOpDescribeAssociation
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeAssociationExecutions(response, &metadata)
@@ -4402,11 +4299,8 @@ func (m *awsAwsjson11_deserializeOpDescribeAssociation
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeAssociationExecutionTargets(response, &metadata)
@@ -4530,11 +4424,8 @@ func (m *awsAwsjson11_deserializeOpDescribeAutomationE
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeAutomationExecutions(response, &metadata)
@@ -4658,11 +4549,8 @@ func (m *awsAwsjson11_deserializeOpDescribeAutomationS
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeAutomationStepExecutions(response, &metadata)
@@ -4789,11 +4677,8 @@ func (m *awsAwsjson11_deserializeOpDescribeAvailablePa
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeAvailablePatches(response, &metadata)
@@ -4908,11 +4793,8 @@ func (m *awsAwsjson11_deserializeOpDescribeDocument) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeDocument(response, &metadata)
@@ -5033,11 +4915,8 @@ func (m *awsAwsjson11_deserializeOpDescribeDocumentPer
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeDocumentPermission(response, &metadata)
@@ -5164,11 +5043,8 @@ func (m *awsAwsjson11_deserializeOpDescribeEffectiveIn
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeEffectiveInstanceAssociations(response, &metadata)
@@ -5289,11 +5165,8 @@ func (m *awsAwsjson11_deserializeOpDescribeEffectivePa
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeEffectivePatchesForPatchBaseline(response, &metadata)
@@ -5417,11 +5290,8 @@ func (m *awsAwsjson11_deserializeOpDescribeInstanceAss
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeInstanceAssociationsStatus(response, &metadata)
@@ -5542,11 +5412,8 @@ func (m *awsAwsjson11_deserializeOpDescribeInstanceInf
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeInstanceInformation(response, &metadata)
@@ -5673,11 +5540,8 @@ func (m *awsAwsjson11_deserializeOpDescribeInstancePat
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeInstancePatches(response, &metadata)
@@ -5801,11 +5665,8 @@ func (m *awsAwsjson11_deserializeOpDescribeInstancePat
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeInstancePatchStates(response, &metadata)
@@ -5923,11 +5784,8 @@ func (m *awsAwsjson11_deserializeOpDescribeInstancePat
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeInstancePatchStatesForPatchGroup(response, &metadata)
@@ -6048,11 +5906,8 @@ func (m *awsAwsjson11_deserializeOpDescribeInstancePro
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeInstanceProperties(response, &metadata)
@@ -6185,11 +6040,8 @@ func (m *awsAwsjson11_deserializeOpDescribeInventoryDe
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeInventoryDeletions(response, &metadata)
@@ -6310,11 +6162,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindowExecutions(response, &metadata)
@@ -6429,11 +6278,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindowExecutionTaskInvocations(response, &metadata)
@@ -6551,11 +6397,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindowExecutionTasks(response, &metadata)
@@ -6673,11 +6516,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindows(response, &metadata)
@@ -6792,11 +6632,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindowSchedule(response, &metadata)
@@ -6914,11 +6751,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindowsForTarget(response, &metadata)
@@ -7033,11 +6867,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindowTargets(response, &metadata)
@@ -7155,11 +6986,8 @@ func (m *awsAwsjson11_deserializeOpDescribeMaintenance
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeMaintenanceWindowTasks(response, &metadata)
@@ -7277,11 +7105,8 @@ func (m *awsAwsjson11_deserializeOpDescribeOpsItems) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeOpsItems(response, &metadata)
@@ -7396,11 +7221,8 @@ func (m *awsAwsjson11_deserializeOpDescribeParameters)
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeParameters(response, &metadata)
@@ -7527,11 +7349,8 @@ func (m *awsAwsjson11_deserializeOpDescribePatchBaseli
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribePatchBaselines(response, &metadata)
@@ -7646,11 +7465,8 @@ func (m *awsAwsjson11_deserializeOpDescribePatchGroups
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribePatchGroups(response, &metadata)
@@ -7765,11 +7581,8 @@ func (m *awsAwsjson11_deserializeOpDescribePatchGroupS
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribePatchGroupState(response, &metadata)
@@ -7887,11 +7700,8 @@ func (m *awsAwsjson11_deserializeOpDescribePatchProper
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribePatchProperties(response, &metadata)
@@ -8006,11 +7816,8 @@ func (m *awsAwsjson11_deserializeOpDescribeSessions) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeSessions(response, &metadata)
@@ -8131,11 +7938,8 @@ func (m *awsAwsjson11_deserializeOpDisassociateOpsItem
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorDisassociateOpsItemRelatedItem(response, &metadata)
@@ -8262,11 +8066,8 @@ func (m *awsAwsjson11_deserializeOpGetAccessToken) Han
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetAccessToken(response, &metadata)
@@ -8393,11 +8194,8 @@ func (m *awsAwsjson11_deserializeOpGetAutomationExecut
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetAutomationExecution(response, &metadata)
@@ -8515,11 +8313,8 @@ func (m *awsAwsjson11_deserializeOpGetCalendarState) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetCalendarState(response, &metadata)
@@ -8643,11 +8438,8 @@ func (m *awsAwsjson11_deserializeOpGetCloudConnector) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetCloudConnector(response, &metadata)
@@ -8765,11 +8557,8 @@ func (m *awsAwsjson11_deserializeOpGetCommandInvocatio
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetCommandInvocation(response, &metadata)
@@ -8896,11 +8685,8 @@ func (m *awsAwsjson11_deserializeOpGetConnectionStatus
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetConnectionStatus(response, &metadata)
@@ -9015,11 +8801,8 @@ func (m *awsAwsjson11_deserializeOpGetDefaultPatchBase
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetDefaultPatchBaseline(response, &metadata)
@@ -9134,11 +8917,8 @@ func (m *awsAwsjson11_deserializeOpGetDeployablePatchS
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetDeployablePatchSnapshotForInstance(response, &metadata)
@@ -9259,11 +9039,8 @@ func (m *awsAwsjson11_deserializeOpGetDocument) Handle
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetDocument(response, &metadata)
@@ -9384,11 +9161,8 @@ func (m *awsAwsjson11_deserializeOpGetExecutionPreview
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetExecutionPreview(response, &metadata)
@@ -9506,11 +9280,8 @@ func (m *awsAwsjson11_deserializeOpGetInventory) Handl
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetInventory(response, &metadata)
@@ -9643,11 +9414,8 @@ func (m *awsAwsjson11_deserializeOpGetInventorySchema)
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetInventorySchema(response, &metadata)
@@ -9768,11 +9536,8 @@ func (m *awsAwsjson11_deserializeOpGetMaintenanceWindo
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetMaintenanceWindow(response, &metadata)
@@ -9890,11 +9655,8 @@ func (m *awsAwsjson11_deserializeOpGetMaintenanceWindo
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetMaintenanceWindowExecution(response, &metadata)
@@ -10012,11 +9774,8 @@ func (m *awsAwsjson11_deserializeOpGetMaintenanceWindo
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetMaintenanceWindowExecutionTask(response, &metadata)
@@ -10134,11 +9893,8 @@ func (m *awsAwsjson11_deserializeOpGetMaintenanceWindo
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetMaintenanceWindowExecutionTaskInvocation(response, &metadata)
@@ -10256,11 +10012,8 @@ func (m *awsAwsjson11_deserializeOpGetMaintenanceWindo
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetMaintenanceWindowTask(response, &metadata)
@@ -10378,11 +10131,8 @@ func (m *awsAwsjson11_deserializeOpGetOpsItem) HandleD
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetOpsItem(response, &metadata)
@@ -10503,11 +10253,8 @@ func (m *awsAwsjson11_deserializeOpGetOpsMetadata) Han
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetOpsMetadata(response, &metadata)
@@ -10628,11 +10375,8 @@ func (m *awsAwsjson11_deserializeOpGetOpsSummary) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetOpsSummary(response, &metadata)
@@ -10762,11 +10506,8 @@ func (m *awsAwsjson11_deserializeOpGetParameter) Handl
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetParameter(response, &metadata)
@@ -10890,11 +10631,8 @@ func (m *awsAwsjson11_deserializeOpGetParameterHistory
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetParameterHistory(response, &metadata)
@@ -11018,11 +10756,8 @@ func (m *awsAwsjson11_deserializeOpGetParameters) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetParameters(response, &metadata)
@@ -11140,11 +10875,8 @@ func (m *awsAwsjson11_deserializeOpGetParametersByPath
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetParametersByPath(response, &metadata)
@@ -11274,11 +11006,8 @@ func (m *awsAwsjson11_deserializeOpGetPatchBaseline) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetPatchBaseline(response, &metadata)
@@ -11399,11 +11128,8 @@ func (m *awsAwsjson11_deserializeOpGetPatchBaselineFor
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetPatchBaselineForPatchGroup(response, &metadata)
@@ -11518,11 +11244,8 @@ func (m *awsAwsjson11_deserializeOpGetResourcePolicies
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetResourcePolicies(response, &metadata)
@@ -11643,11 +11366,8 @@ func (m *awsAwsjson11_deserializeOpGetServiceSetting) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorGetServiceSetting(response, &metadata)
@@ -11765,11 +11485,8 @@ func (m *awsAwsjson11_deserializeOpLabelParameterVersi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorLabelParameterVersion(response, &metadata)
@@ -11896,11 +11613,8 @@ func (m *awsAwsjson11_deserializeOpListAssociations) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListAssociations(response, &metadata)
@@ -12018,11 +11732,8 @@ func (m *awsAwsjson11_deserializeOpListAssociationVers
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListAssociationVersions(response, &metadata)
@@ -12143,11 +11854,8 @@ func (m *awsAwsjson11_deserializeOpListCloudConnectors
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListCloudConnectors(response, &metadata)
@@ -12262,11 +11970,8 @@ func (m *awsAwsjson11_deserializeOpListCommandInvocati
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListCommandInvocations(response, &metadata)
@@ -12393,11 +12098,8 @@ func (m *awsAwsjson11_deserializeOpListCommands) Handl
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListCommands(response, &metadata)
@@ -12524,11 +12226,8 @@ func (m *awsAwsjson11_deserializeOpListComplianceItems
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListComplianceItems(response, &metadata)
@@ -12655,11 +12354,8 @@ func (m *awsAwsjson11_deserializeOpListComplianceSumma
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListComplianceSummaries(response, &metadata)
@@ -12780,11 +12476,8 @@ func (m *awsAwsjson11_deserializeOpListDocumentMetadat
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListDocumentMetadataHistory(response, &metadata)
@@ -12908,11 +12601,8 @@ func (m *awsAwsjson11_deserializeOpListDocuments) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListDocuments(response, &metadata)
@@ -13033,11 +12723,8 @@ func (m *awsAwsjson11_deserializeOpListDocumentVersion
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListDocumentVersions(response, &metadata)
@@ -13158,11 +12845,8 @@ func (m *awsAwsjson11_deserializeOpListInventoryEntrie
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListInventoryEntries(response, &metadata)
@@ -13289,11 +12973,8 @@ func (m *awsAwsjson11_deserializeOpListNodes) HandleDe
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListNodes(response, &metadata)
@@ -13420,11 +13101,8 @@ func (m *awsAwsjson11_deserializeOpListNodesSummary) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListNodesSummary(response, &metadata)
@@ -13554,11 +13232,8 @@ func (m *awsAwsjson11_deserializeOpListOpsItemEvents) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListOpsItemEvents(response, &metadata)
@@ -13682,11 +13357,8 @@ func (m *awsAwsjson11_deserializeOpListOpsItemRelatedI
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListOpsItemRelatedItems(response, &metadata)
@@ -13804,11 +13476,8 @@ func (m *awsAwsjson11_deserializeOpListOpsMetadata) Ha
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListOpsMetadata(response, &metadata)
@@ -13926,11 +13595,8 @@ func (m *awsAwsjson11_deserializeOpListResourceComplia
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListResourceComplianceSummaries(response, &metadata)
@@ -14051,11 +13717,8 @@ func (m *awsAwsjson11_deserializeOpListResourceDataSyn
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListResourceDataSync(response, &metadata)
@@ -14176,11 +13839,8 @@ func (m *awsAwsjson11_deserializeOpListTagsForResource
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorListTagsForResource(response, &metadata)
@@ -14301,11 +13961,8 @@ func (m *awsAwsjson11_deserializeOpModifyDocumentPermi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorModifyDocumentPermission(response, &metadata)
@@ -14432,11 +14089,8 @@ func (m *awsAwsjson11_deserializeOpPutComplianceItems)
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorPutComplianceItems(response, &metadata)
@@ -14569,11 +14223,8 @@ func (m *awsAwsjson11_deserializeOpPutInventory) Handl
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorPutInventory(response, &metadata)
@@ -14721,11 +14372,8 @@ func (m *awsAwsjson11_deserializeOpPutParameter) Handl
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorPutParameter(response, &metadata)
@@ -14882,11 +14530,8 @@ func (m *awsAwsjson11_deserializeOpPutResourcePolicy) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorPutResourcePolicy(response, &metadata)
@@ -15019,11 +14664,8 @@ func (m *awsAwsjson11_deserializeOpRegisterDefaultPatc
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorRegisterDefaultPatchBaseline(response, &metadata)
@@ -15144,11 +14786,8 @@ func (m *awsAwsjson11_deserializeOpRegisterPatchBaseli
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorRegisterPatchBaselineForPatchGroup(response, &metadata)
@@ -15275,11 +14914,8 @@ func (m *awsAwsjson11_deserializeOpRegisterTargetWithM
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorRegisterTargetWithMaintenanceWindow(response, &metadata)
@@ -15403,11 +15039,8 @@ func (m *awsAwsjson11_deserializeOpRegisterTaskWithMai
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorRegisterTaskWithMaintenanceWindow(response, &metadata)
@@ -15534,11 +15167,8 @@ func (m *awsAwsjson11_deserializeOpRemoveTagsFromResou
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorRemoveTagsFromResource(response, &metadata)
@@ -15662,11 +15292,8 @@ func (m *awsAwsjson11_deserializeOpResetServiceSetting
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorResetServiceSetting(response, &metadata)
@@ -15787,11 +15414,8 @@ func (m *awsAwsjson11_deserializeOpResumeSession) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorResumeSession(response, &metadata)
@@ -15909,11 +15533,8 @@ func (m *awsAwsjson11_deserializeOpSendAutomationSigna
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorSendAutomationSignal(response, &metadata)
@@ -16037,11 +15658,8 @@ func (m *awsAwsjson11_deserializeOpSendCommand) Handle
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorSendCommand(response, &metadata)
@@ -16186,11 +15804,8 @@ func (m *awsAwsjson11_deserializeOpStartAccessRequest)
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorStartAccessRequest(response, &metadata)
@@ -16320,11 +15935,8 @@ func (m *awsAwsjson11_deserializeOpStartAssociationsOn
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorStartAssociationsOnce(response, &metadata)
@@ -16442,11 +16054,8 @@ func (m *awsAwsjson11_deserializeOpStartAutomationExec
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorStartAutomationExecution(response, &metadata)
@@ -16579,11 +16188,8 @@ func (m *awsAwsjson11_deserializeOpStartChangeRequestE
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorStartChangeRequestExecution(response, &metadata)
@@ -16719,11 +16325,8 @@ func (m *awsAwsjson11_deserializeOpStartExecutionPrevi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorStartExecutionPreview(response, &metadata)
@@ -16841,11 +16444,8 @@ func (m *awsAwsjson11_deserializeOpStartSession) Handl
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorStartSession(response, &metadata)
@@ -16966,11 +16566,8 @@ func (m *awsAwsjson11_deserializeOpStopAutomationExecu
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorStopAutomationExecution(response, &metadata)
@@ -17091,11 +16688,8 @@ func (m *awsAwsjson11_deserializeOpTerminateSession) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorTerminateSession(response, &metadata)
@@ -17210,11 +16804,8 @@ func (m *awsAwsjson11_deserializeOpUnlabelParameterVer
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUnlabelParameterVersion(response, &metadata)
@@ -17338,11 +16929,8 @@ func (m *awsAwsjson11_deserializeOpUpdateAssociation) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateAssociation(response, &metadata)
@@ -17493,11 +17081,8 @@ func (m *awsAwsjson11_deserializeOpUpdateAssociationSt
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateAssociationStatus(response, &metadata)
@@ -17627,11 +17212,8 @@ func (m *awsAwsjson11_deserializeOpUpdateCloudConnecto
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateCloudConnector(response, &metadata)
@@ -17752,11 +17334,8 @@ func (m *awsAwsjson11_deserializeOpUpdateDocument) Han
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateDocument(response, &metadata)
@@ -17898,11 +17477,8 @@ func (m *awsAwsjson11_deserializeOpUpdateDocumentDefau
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateDocumentDefaultVersion(response, &metadata)
@@ -18026,11 +17602,8 @@ func (m *awsAwsjson11_deserializeOpUpdateDocumentMetad
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateDocumentMetadata(response, &metadata)
@@ -18157,11 +17730,8 @@ func (m *awsAwsjson11_deserializeOpUpdateMaintenanceWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateMaintenanceWindow(response, &metadata)
@@ -18279,11 +17849,8 @@ func (m *awsAwsjson11_deserializeOpUpdateMaintenanceWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateMaintenanceWindowTarget(response, &metadata)
@@ -18401,11 +17968,8 @@ func (m *awsAwsjson11_deserializeOpUpdateMaintenanceWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateMaintenanceWindowTask(response, &metadata)
@@ -18523,11 +18087,8 @@ func (m *awsAwsjson11_deserializeOpUpdateManagedInstan
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateManagedInstanceRole(response, &metadata)
@@ -18645,11 +18206,8 @@ func (m *awsAwsjson11_deserializeOpUpdateOpsItem) Hand
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateOpsItem(response, &metadata)
@@ -18782,11 +18340,8 @@ func (m *awsAwsjson11_deserializeOpUpdateOpsMetadata) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateOpsMetadata(response, &metadata)
@@ -18913,11 +18468,8 @@ func (m *awsAwsjson11_deserializeOpUpdatePatchBaseline
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdatePatchBaseline(response, &metadata)
@@ -19035,11 +18587,8 @@ func (m *awsAwsjson11_deserializeOpUpdateResourceDataS
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateResourceDataSync(response, &metadata)
@@ -19163,11 +18712,8 @@ func (m *awsAwsjson11_deserializeOpUpdateServiceSettin
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorUpdateServiceSetting(response, &metadata)
@@ -19288,11 +18834,8 @@ func (m *awsAwsjson11_deserializeOpValidateCloudConnec
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsjson11_deserializeOpErrorValidateCloudConnector(response, &metadata)
