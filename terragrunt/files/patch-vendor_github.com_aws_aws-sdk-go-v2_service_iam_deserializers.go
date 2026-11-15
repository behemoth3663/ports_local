--- vendor/github.com/aws/aws-sdk-go-v2/service/iam/deserializers.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iam/deserializers.go
@@ -17,7 +17,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -51,11 +50,9 @@ func (m *awsAwsquery_deserializeOpAcceptDelegationRequ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAcceptDelegationRequest(response, &metadata)
 	}
@@ -142,11 +139,9 @@ func (m *awsAwsquery_deserializeOpAcquireRole) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAcquireRole(response, &metadata)
 	}
@@ -289,11 +284,9 @@ func (m *awsAwsquery_deserializeOpAddClientIDToOpenIDC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAddClientIDToOpenIDConnectProvider(response, &metadata)
 	}
@@ -386,11 +379,9 @@ func (m *awsAwsquery_deserializeOpAddRoleToInstancePro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAddRoleToInstanceProfile(response, &metadata)
 	}
@@ -483,11 +474,9 @@ func (m *awsAwsquery_deserializeOpAddUserToGroup) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAddUserToGroup(response, &metadata)
 	}
@@ -574,11 +563,9 @@ func (m *awsAwsquery_deserializeOpAssociateDelegationR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssociateDelegationRequest(response, &metadata)
 	}
@@ -668,11 +655,9 @@ func (m *awsAwsquery_deserializeOpAttachGroupPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAttachGroupPolicy(response, &metadata)
 	}
@@ -765,11 +750,9 @@ func (m *awsAwsquery_deserializeOpAttachRolePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAttachRolePolicy(response, &metadata)
 	}
@@ -865,11 +848,9 @@ func (m *awsAwsquery_deserializeOpAttachUserPolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAttachUserPolicy(response, &metadata)
 	}
@@ -962,11 +943,9 @@ func (m *awsAwsquery_deserializeOpChangePassword) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorChangePassword(response, &metadata)
 	}
@@ -1062,11 +1041,9 @@ func (m *awsAwsquery_deserializeOpCreateAccessKey) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateAccessKey(response, &metadata)
 	}
@@ -1188,11 +1165,9 @@ func (m *awsAwsquery_deserializeOpCreateAccountAlias) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateAccountAlias(response, &metadata)
 	}
@@ -1282,11 +1257,9 @@ func (m *awsAwsquery_deserializeOpCreateDelegationRequ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateDelegationRequest(response, &metadata)
 	}
@@ -1414,11 +1387,9 @@ func (m *awsAwsquery_deserializeOpCreateGroup) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateGroup(response, &metadata)
 	}
@@ -1543,11 +1514,9 @@ func (m *awsAwsquery_deserializeOpCreateInstanceProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateInstanceProfile(response, &metadata)
 	}
@@ -1675,11 +1644,9 @@ func (m *awsAwsquery_deserializeOpCreateLoginProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateLoginProfile(response, &metadata)
 	}
@@ -1807,11 +1774,9 @@ func (m *awsAwsquery_deserializeOpCreateOpenIDConnectP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateOpenIDConnectProvider(response, &metadata)
 	}
@@ -1942,11 +1907,9 @@ func (m *awsAwsquery_deserializeOpCreatePolicy) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreatePolicy(response, &metadata)
 	}
@@ -2077,11 +2040,9 @@ func (m *awsAwsquery_deserializeOpCreatePolicyVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreatePolicyVersion(response, &metadata)
 	}
@@ -2209,11 +2170,9 @@ func (m *awsAwsquery_deserializeOpCreateRole) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateRole(response, &metadata)
 	}
@@ -2344,11 +2303,9 @@ func (m *awsAwsquery_deserializeOpCreateSAMLProvider) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateSAMLProvider(response, &metadata)
 	}
@@ -2476,11 +2433,9 @@ func (m *awsAwsquery_deserializeOpCreateServiceLinkedR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateServiceLinkedRole(response, &metadata)
 	}
@@ -2605,11 +2560,9 @@ func (m *awsAwsquery_deserializeOpCreateServiceSpecifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateServiceSpecificCredential(response, &metadata)
 	}
@@ -2731,11 +2684,9 @@ func (m *awsAwsquery_deserializeOpCreateUser) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateUser(response, &metadata)
 	}
@@ -2866,11 +2817,9 @@ func (m *awsAwsquery_deserializeOpCreateVirtualMFADevi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateVirtualMFADevice(response, &metadata)
 	}
@@ -2998,11 +2947,9 @@ func (m *awsAwsquery_deserializeOpDeactivateMFADevice)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeactivateMFADevice(response, &metadata)
 	}
@@ -3095,11 +3042,9 @@ func (m *awsAwsquery_deserializeOpDeleteAccessKey) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteAccessKey(response, &metadata)
 	}
@@ -3186,11 +3131,9 @@ func (m *awsAwsquery_deserializeOpDeleteAccountAlias) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteAccountAlias(response, &metadata)
 	}
@@ -3280,11 +3223,9 @@ func (m *awsAwsquery_deserializeOpDeleteAccountPasswor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteAccountPasswordPolicy(response, &metadata)
 	}
@@ -3371,11 +3312,9 @@ func (m *awsAwsquery_deserializeOpDeleteGroup) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteGroup(response, &metadata)
 	}
@@ -3465,11 +3404,9 @@ func (m *awsAwsquery_deserializeOpDeleteGroupPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteGroupPolicy(response, &metadata)
 	}
@@ -3556,11 +3493,9 @@ func (m *awsAwsquery_deserializeOpDeleteInstanceProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteInstanceProfile(response, &metadata)
 	}
@@ -3650,11 +3585,9 @@ func (m *awsAwsquery_deserializeOpDeleteLoginProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteLoginProfile(response, &metadata)
 	}
@@ -3744,11 +3677,9 @@ func (m *awsAwsquery_deserializeOpDeleteOpenIDConnectP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteOpenIDConnectProvider(response, &metadata)
 	}
@@ -3835,11 +3766,9 @@ func (m *awsAwsquery_deserializeOpDeletePolicy) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeletePolicy(response, &metadata)
 	}
@@ -3932,11 +3861,9 @@ func (m *awsAwsquery_deserializeOpDeletePolicyVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeletePolicyVersion(response, &metadata)
 	}
@@ -4029,11 +3956,9 @@ func (m *awsAwsquery_deserializeOpDeleteRole) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteRole(response, &metadata)
 	}
@@ -4129,11 +4054,9 @@ func (m *awsAwsquery_deserializeOpDeleteRolePermission
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteRolePermissionsBoundary(response, &metadata)
 	}
@@ -4220,11 +4143,9 @@ func (m *awsAwsquery_deserializeOpDeleteRolePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteRolePolicy(response, &metadata)
 	}
@@ -4314,11 +4235,9 @@ func (m *awsAwsquery_deserializeOpDeleteSAMLProvider) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteSAMLProvider(response, &metadata)
 	}
@@ -4408,11 +4327,9 @@ func (m *awsAwsquery_deserializeOpDeleteServerCertific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteServerCertificate(response, &metadata)
 	}
@@ -4502,11 +4419,9 @@ func (m *awsAwsquery_deserializeOpDeleteServiceLinkedR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteServiceLinkedRole(response, &metadata)
 	}
@@ -4628,11 +4543,9 @@ func (m *awsAwsquery_deserializeOpDeleteServiceSpecifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteServiceSpecificCredential(response, &metadata)
 	}
@@ -4713,11 +4626,9 @@ func (m *awsAwsquery_deserializeOpDeleteSigningCertifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteSigningCertificate(response, &metadata)
 	}
@@ -4807,11 +4718,9 @@ func (m *awsAwsquery_deserializeOpDeleteSSHPublicKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteSSHPublicKey(response, &metadata)
 	}
@@ -4892,11 +4801,9 @@ func (m *awsAwsquery_deserializeOpDeleteUser) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteUser(response, &metadata)
 	}
@@ -4989,11 +4896,9 @@ func (m *awsAwsquery_deserializeOpDeleteUserPermission
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteUserPermissionsBoundary(response, &metadata)
 	}
@@ -5077,11 +4982,9 @@ func (m *awsAwsquery_deserializeOpDeleteUserPolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteUserPolicy(response, &metadata)
 	}
@@ -5168,11 +5071,9 @@ func (m *awsAwsquery_deserializeOpDeleteVirtualMFADevi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteVirtualMFADevice(response, &metadata)
 	}
@@ -5265,11 +5166,9 @@ func (m *awsAwsquery_deserializeOpDetachGroupPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDetachGroupPolicy(response, &metadata)
 	}
@@ -5359,11 +5258,9 @@ func (m *awsAwsquery_deserializeOpDetachRolePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDetachRolePolicy(response, &metadata)
 	}
@@ -5456,11 +5353,9 @@ func (m *awsAwsquery_deserializeOpDetachUserPolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDetachUserPolicy(response, &metadata)
 	}
@@ -5550,11 +5445,9 @@ func (m *awsAwsquery_deserializeOpDisableOrganizations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDisableOrganizationsRootCredentialsManagement(response, &metadata)
 	}
@@ -5679,11 +5572,9 @@ func (m *awsAwsquery_deserializeOpDisableOrganizations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDisableOrganizationsRootSessions(response, &metadata)
 	}
@@ -5808,11 +5699,9 @@ func (m *awsAwsquery_deserializeOpDisableOutboundWebId
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDisableOutboundWebIdentityFederation(response, &metadata)
 	}
@@ -5893,11 +5782,9 @@ func (m *awsAwsquery_deserializeOpEnableMFADevice) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorEnableMFADevice(response, &metadata)
 	}
@@ -5996,11 +5883,9 @@ func (m *awsAwsquery_deserializeOpEnableOrganizationsR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorEnableOrganizationsRootCredentialsManagement(response, &metadata)
 	}
@@ -6128,11 +6013,9 @@ func (m *awsAwsquery_deserializeOpEnableOrganizationsR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorEnableOrganizationsRootSessions(response, &metadata)
 	}
@@ -6260,11 +6143,9 @@ func (m *awsAwsquery_deserializeOpEnableOutboundWebIde
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorEnableOutboundWebIdentityFederation(response, &metadata)
 	}
@@ -6380,11 +6261,9 @@ func (m *awsAwsquery_deserializeOpGenerateCredentialRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGenerateCredentialReport(response, &metadata)
 	}
@@ -6503,11 +6382,9 @@ func (m *awsAwsquery_deserializeOpGenerateOrganization
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGenerateOrganizationsAccessReport(response, &metadata)
 	}
@@ -6623,11 +6500,9 @@ func (m *awsAwsquery_deserializeOpGenerateServiceLastA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGenerateServiceLastAccessedDetails(response, &metadata)
 	}
@@ -6746,11 +6621,9 @@ func (m *awsAwsquery_deserializeOpGetAccessKeyLastUsed
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccessKeyLastUsed(response, &metadata)
 	}
@@ -6863,11 +6736,9 @@ func (m *awsAwsquery_deserializeOpGetAccountAuthorizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccountAuthorizationDetails(response, &metadata)
 	}
@@ -6983,11 +6854,9 @@ func (m *awsAwsquery_deserializeOpGetAccountPasswordPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccountPasswordPolicy(response, &metadata)
 	}
@@ -7106,11 +6975,9 @@ func (m *awsAwsquery_deserializeOpGetAccountProperties
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccountProperties(response, &metadata)
 	}
@@ -7229,11 +7096,9 @@ func (m *awsAwsquery_deserializeOpGetAccountSummary) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccountSummary(response, &metadata)
 	}
@@ -7349,11 +7214,9 @@ func (m *awsAwsquery_deserializeOpGetContextKeysForCus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetContextKeysForCustomPolicy(response, &metadata)
 	}
@@ -7469,11 +7332,9 @@ func (m *awsAwsquery_deserializeOpGetContextKeysForPri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetContextKeysForPrincipalPolicy(response, &metadata)
 	}
@@ -7592,11 +7453,9 @@ func (m *awsAwsquery_deserializeOpGetCredentialReport)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetCredentialReport(response, &metadata)
 	}
@@ -7721,11 +7580,9 @@ func (m *awsAwsquery_deserializeOpGetDelegationRequest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetDelegationRequest(response, &metadata)
 	}
@@ -7844,11 +7701,9 @@ func (m *awsAwsquery_deserializeOpGetGroup) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetGroup(response, &metadata)
 	}
@@ -7967,11 +7822,9 @@ func (m *awsAwsquery_deserializeOpGetGroupPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetGroupPolicy(response, &metadata)
 	}
@@ -8090,11 +7943,9 @@ func (m *awsAwsquery_deserializeOpGetHumanReadableSumm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetHumanReadableSummary(response, &metadata)
 	}
@@ -8216,11 +8067,9 @@ func (m *awsAwsquery_deserializeOpGetInstanceProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetInstanceProfile(response, &metadata)
 	}
@@ -8339,11 +8188,9 @@ func (m *awsAwsquery_deserializeOpGetLoginProfile) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetLoginProfile(response, &metadata)
 	}
@@ -8462,11 +8309,9 @@ func (m *awsAwsquery_deserializeOpGetMFADevice) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetMFADevice(response, &metadata)
 	}
@@ -8585,11 +8430,9 @@ func (m *awsAwsquery_deserializeOpGetOpenIDConnectProv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetOpenIDConnectProvider(response, &metadata)
 	}
@@ -8711,11 +8554,9 @@ func (m *awsAwsquery_deserializeOpGetOrganizationsAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetOrganizationsAccessReport(response, &metadata)
 	}
@@ -8831,11 +8672,9 @@ func (m *awsAwsquery_deserializeOpGetOutboundWebIdenti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetOutboundWebIdentityFederationInfo(response, &metadata)
 	}
@@ -8951,11 +8790,9 @@ func (m *awsAwsquery_deserializeOpGetPolicy) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetPolicy(response, &metadata)
 	}
@@ -9077,11 +8914,9 @@ func (m *awsAwsquery_deserializeOpGetPolicyVersion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetPolicyVersion(response, &metadata)
 	}
@@ -9203,11 +9038,9 @@ func (m *awsAwsquery_deserializeOpGetRole) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetRole(response, &metadata)
 	}
@@ -9326,11 +9159,9 @@ func (m *awsAwsquery_deserializeOpGetRolePolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetRolePolicy(response, &metadata)
 	}
@@ -9449,11 +9280,9 @@ func (m *awsAwsquery_deserializeOpGetRoleTemplateVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetRoleTemplateVersion(response, &metadata)
 	}
@@ -9575,11 +9404,9 @@ func (m *awsAwsquery_deserializeOpGetSAMLProvider) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSAMLProvider(response, &metadata)
 	}
@@ -9701,11 +9528,9 @@ func (m *awsAwsquery_deserializeOpGetServerCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetServerCertificate(response, &metadata)
 	}
@@ -9824,11 +9649,9 @@ func (m *awsAwsquery_deserializeOpGetServiceLastAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetServiceLastAccessedDetails(response, &metadata)
 	}
@@ -9947,11 +9770,9 @@ func (m *awsAwsquery_deserializeOpGetServiceLastAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetServiceLastAccessedDetailsWithEntities(response, &metadata)
 	}
@@ -10070,11 +9891,9 @@ func (m *awsAwsquery_deserializeOpGetServiceLinkedRole
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetServiceLinkedRoleDeletionStatus(response, &metadata)
 	}
@@ -10196,11 +10015,9 @@ func (m *awsAwsquery_deserializeOpGetSSHPublicKey) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSSHPublicKey(response, &metadata)
 	}
@@ -10319,11 +10136,9 @@ func (m *awsAwsquery_deserializeOpGetUser) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetUser(response, &metadata)
 	}
@@ -10442,11 +10257,9 @@ func (m *awsAwsquery_deserializeOpGetUserPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetUserPolicy(response, &metadata)
 	}
@@ -10565,11 +10378,9 @@ func (m *awsAwsquery_deserializeOpListAccessKeys) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListAccessKeys(response, &metadata)
 	}
@@ -10688,11 +10499,9 @@ func (m *awsAwsquery_deserializeOpListAccountAliases) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListAccountAliases(response, &metadata)
 	}
@@ -10808,11 +10617,9 @@ func (m *awsAwsquery_deserializeOpListAttachedGroupPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListAttachedGroupPolicies(response, &metadata)
 	}
@@ -10934,11 +10741,9 @@ func (m *awsAwsquery_deserializeOpListAttachedRolePoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListAttachedRolePolicies(response, &metadata)
 	}
@@ -11060,11 +10865,9 @@ func (m *awsAwsquery_deserializeOpListAttachedUserPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListAttachedUserPolicies(response, &metadata)
 	}
@@ -11186,11 +10989,9 @@ func (m *awsAwsquery_deserializeOpListDelegationReques
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListDelegationRequests(response, &metadata)
 	}
@@ -11312,11 +11113,9 @@ func (m *awsAwsquery_deserializeOpListEntitiesForPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListEntitiesForPolicy(response, &metadata)
 	}
@@ -11438,11 +11237,9 @@ func (m *awsAwsquery_deserializeOpListGroupPolicies) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListGroupPolicies(response, &metadata)
 	}
@@ -11561,11 +11358,9 @@ func (m *awsAwsquery_deserializeOpListGroups) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListGroups(response, &metadata)
 	}
@@ -11681,11 +11476,9 @@ func (m *awsAwsquery_deserializeOpListGroupsForUser) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListGroupsForUser(response, &metadata)
 	}
@@ -11804,11 +11597,9 @@ func (m *awsAwsquery_deserializeOpListInstanceProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListInstanceProfiles(response, &metadata)
 	}
@@ -11924,11 +11715,9 @@ func (m *awsAwsquery_deserializeOpListInstanceProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListInstanceProfilesForRole(response, &metadata)
 	}
@@ -12047,11 +11836,9 @@ func (m *awsAwsquery_deserializeOpListInstanceProfileT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListInstanceProfileTags(response, &metadata)
 	}
@@ -12170,11 +11957,9 @@ func (m *awsAwsquery_deserializeOpListMFADevices) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListMFADevices(response, &metadata)
 	}
@@ -12293,11 +12078,9 @@ func (m *awsAwsquery_deserializeOpListMFADeviceTags) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListMFADeviceTags(response, &metadata)
 	}
@@ -12419,11 +12202,9 @@ func (m *awsAwsquery_deserializeOpListOpenIDConnectPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListOpenIDConnectProviders(response, &metadata)
 	}
@@ -12539,11 +12320,9 @@ func (m *awsAwsquery_deserializeOpListOpenIDConnectPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListOpenIDConnectProviderTags(response, &metadata)
 	}
@@ -12665,11 +12444,9 @@ func (m *awsAwsquery_deserializeOpListOrganizationsFea
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListOrganizationsFeatures(response, &metadata)
 	}
@@ -12794,11 +12571,9 @@ func (m *awsAwsquery_deserializeOpListPolicies) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListPolicies(response, &metadata)
 	}
@@ -12914,11 +12689,9 @@ func (m *awsAwsquery_deserializeOpListPoliciesGranting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListPoliciesGrantingServiceAccess(response, &metadata)
 	}
@@ -13037,11 +12810,9 @@ func (m *awsAwsquery_deserializeOpListPolicyTags) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListPolicyTags(response, &metadata)
 	}
@@ -13163,11 +12934,9 @@ func (m *awsAwsquery_deserializeOpListPolicyVersions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListPolicyVersions(response, &metadata)
 	}
@@ -13289,11 +13058,9 @@ func (m *awsAwsquery_deserializeOpListRolePolicies) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListRolePolicies(response, &metadata)
 	}
@@ -13412,11 +13179,9 @@ func (m *awsAwsquery_deserializeOpListRoles) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListRoles(response, &metadata)
 	}
@@ -13532,11 +13297,9 @@ func (m *awsAwsquery_deserializeOpListRoleTags) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListRoleTags(response, &metadata)
 	}
@@ -13655,11 +13418,9 @@ func (m *awsAwsquery_deserializeOpListSAMLProviders) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListSAMLProviders(response, &metadata)
 	}
@@ -13775,11 +13536,9 @@ func (m *awsAwsquery_deserializeOpListSAMLProviderTags
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListSAMLProviderTags(response, &metadata)
 	}
@@ -13901,11 +13660,9 @@ func (m *awsAwsquery_deserializeOpListServerCertificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListServerCertificates(response, &metadata)
 	}
@@ -14021,11 +13778,9 @@ func (m *awsAwsquery_deserializeOpListServerCertificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListServerCertificateTags(response, &metadata)
 	}
@@ -14144,11 +13899,9 @@ func (m *awsAwsquery_deserializeOpListServiceSpecificC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListServiceSpecificCredentials(response, &metadata)
 	}
@@ -14267,11 +14020,9 @@ func (m *awsAwsquery_deserializeOpListSigningCertifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListSigningCertificates(response, &metadata)
 	}
@@ -14390,11 +14141,9 @@ func (m *awsAwsquery_deserializeOpListSSHPublicKeys) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListSSHPublicKeys(response, &metadata)
 	}
@@ -14510,11 +14259,9 @@ func (m *awsAwsquery_deserializeOpListUserPolicies) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListUserPolicies(response, &metadata)
 	}
@@ -14633,11 +14380,9 @@ func (m *awsAwsquery_deserializeOpListUsers) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListUsers(response, &metadata)
 	}
@@ -14753,11 +14498,9 @@ func (m *awsAwsquery_deserializeOpListUserTags) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListUserTags(response, &metadata)
 	}
@@ -14876,11 +14619,9 @@ func (m *awsAwsquery_deserializeOpListVirtualMFADevice
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListVirtualMFADevices(response, &metadata)
 	}
@@ -14993,11 +14734,9 @@ func (m *awsAwsquery_deserializeOpPutAccountProperties
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPutAccountProperties(response, &metadata)
 	}
@@ -15119,11 +14858,9 @@ func (m *awsAwsquery_deserializeOpPutGroupPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPutGroupPolicy(response, &metadata)
 	}
@@ -15213,11 +14950,9 @@ func (m *awsAwsquery_deserializeOpPutRolePermissionsBo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPutRolePermissionsBoundary(response, &metadata)
 	}
@@ -15310,11 +15045,9 @@ func (m *awsAwsquery_deserializeOpPutRolePolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPutRolePolicy(response, &metadata)
 	}
@@ -15407,11 +15140,9 @@ func (m *awsAwsquery_deserializeOpPutUserPermissionsBo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPutUserPermissionsBoundary(response, &metadata)
 	}
@@ -15501,11 +15232,9 @@ func (m *awsAwsquery_deserializeOpPutUserPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPutUserPolicy(response, &metadata)
 	}
@@ -15595,11 +15324,9 @@ func (m *awsAwsquery_deserializeOpRejectDelegationRequ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorRejectDelegationRequest(response, &metadata)
 	}
@@ -15689,11 +15416,9 @@ func (m *awsAwsquery_deserializeOpRemoveClientIDFromOp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorRemoveClientIDFromOpenIDConnectProvider(response, &metadata)
 	}
@@ -15783,11 +15508,9 @@ func (m *awsAwsquery_deserializeOpRemoveRoleFromInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorRemoveRoleFromInstanceProfile(response, &metadata)
 	}
@@ -15877,11 +15600,9 @@ func (m *awsAwsquery_deserializeOpRemoveUserFromGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorRemoveUserFromGroup(response, &metadata)
 	}
@@ -15968,11 +15689,9 @@ func (m *awsAwsquery_deserializeOpResetServiceSpecific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorResetServiceSpecificCredential(response, &metadata)
 	}
@@ -16088,11 +15807,9 @@ func (m *awsAwsquery_deserializeOpResyncMFADevice) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorResyncMFADevice(response, &metadata)
 	}
@@ -16185,11 +15902,9 @@ func (m *awsAwsquery_deserializeOpSendDelegationToken)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSendDelegationToken(response, &metadata)
 	}
@@ -16279,11 +15994,9 @@ func (m *awsAwsquery_deserializeOpSetDefaultPolicyVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSetDefaultPolicyVersion(response, &metadata)
 	}
@@ -16373,11 +16086,9 @@ func (m *awsAwsquery_deserializeOpSetSecurityTokenServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSetSecurityTokenServicePreferences(response, &metadata)
 	}
@@ -16458,11 +16169,9 @@ func (m *awsAwsquery_deserializeOpSimulateCustomPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSimulateCustomPolicy(response, &metadata)
 	}
@@ -16581,11 +16290,9 @@ func (m *awsAwsquery_deserializeOpSimulatePrincipalPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSimulatePrincipalPolicy(response, &metadata)
 	}
@@ -16707,11 +16414,9 @@ func (m *awsAwsquery_deserializeOpTagInstanceProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagInstanceProfile(response, &metadata)
 	}
@@ -16804,11 +16509,9 @@ func (m *awsAwsquery_deserializeOpTagMFADevice) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagMFADevice(response, &metadata)
 	}
@@ -16901,11 +16604,9 @@ func (m *awsAwsquery_deserializeOpTagOpenIDConnectProv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagOpenIDConnectProvider(response, &metadata)
 	}
@@ -16998,11 +16699,9 @@ func (m *awsAwsquery_deserializeOpTagPolicy) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagPolicy(response, &metadata)
 	}
@@ -17095,11 +16794,9 @@ func (m *awsAwsquery_deserializeOpTagRole) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagRole(response, &metadata)
 	}
@@ -17192,11 +16889,9 @@ func (m *awsAwsquery_deserializeOpTagSAMLProvider) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagSAMLProvider(response, &metadata)
 	}
@@ -17289,11 +16984,9 @@ func (m *awsAwsquery_deserializeOpTagServerCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagServerCertificate(response, &metadata)
 	}
@@ -17386,11 +17079,9 @@ func (m *awsAwsquery_deserializeOpTagUser) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagUser(response, &metadata)
 	}
@@ -17483,11 +17174,9 @@ func (m *awsAwsquery_deserializeOpUntagInstanceProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagInstanceProfile(response, &metadata)
 	}
@@ -17577,11 +17266,9 @@ func (m *awsAwsquery_deserializeOpUntagMFADevice) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagMFADevice(response, &metadata)
 	}
@@ -17671,11 +17358,9 @@ func (m *awsAwsquery_deserializeOpUntagOpenIDConnectPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagOpenIDConnectProvider(response, &metadata)
 	}
@@ -17765,11 +17450,9 @@ func (m *awsAwsquery_deserializeOpUntagPolicy) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagPolicy(response, &metadata)
 	}
@@ -17859,11 +17542,9 @@ func (m *awsAwsquery_deserializeOpUntagRole) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagRole(response, &metadata)
 	}
@@ -17950,11 +17631,9 @@ func (m *awsAwsquery_deserializeOpUntagSAMLProvider) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagSAMLProvider(response, &metadata)
 	}
@@ -18044,11 +17723,9 @@ func (m *awsAwsquery_deserializeOpUntagServerCertifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagServerCertificate(response, &metadata)
 	}
@@ -18138,11 +17815,9 @@ func (m *awsAwsquery_deserializeOpUntagUser) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagUser(response, &metadata)
 	}
@@ -18229,11 +17904,9 @@ func (m *awsAwsquery_deserializeOpUpdateAccessKey) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateAccessKey(response, &metadata)
 	}
@@ -18323,11 +17996,9 @@ func (m *awsAwsquery_deserializeOpUpdateAccountPasswor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateAccountPasswordPolicy(response, &metadata)
 	}
@@ -18417,11 +18088,9 @@ func (m *awsAwsquery_deserializeOpUpdateAssumeRolePoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateAssumeRolePolicy(response, &metadata)
 	}
@@ -18514,11 +18183,9 @@ func (m *awsAwsquery_deserializeOpUpdateDelegationRequ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateDelegationRequest(response, &metadata)
 	}
@@ -18608,11 +18275,9 @@ func (m *awsAwsquery_deserializeOpUpdateGroup) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateGroup(response, &metadata)
 	}
@@ -18702,11 +18367,9 @@ func (m *awsAwsquery_deserializeOpUpdateLoginProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateLoginProfile(response, &metadata)
 	}
@@ -18799,11 +18462,9 @@ func (m *awsAwsquery_deserializeOpUpdateOpenIDConnectP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateOpenIDConnectProviderThumbprint(response, &metadata)
 	}
@@ -18893,11 +18554,9 @@ func (m *awsAwsquery_deserializeOpUpdateRole) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateRole(response, &metadata)
 	}
@@ -19019,11 +18678,9 @@ func (m *awsAwsquery_deserializeOpUpdateRoleDescriptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateRoleDescription(response, &metadata)
 	}
@@ -19145,11 +18802,9 @@ func (m *awsAwsquery_deserializeOpUpdateSAMLProvider) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateSAMLProvider(response, &metadata)
 	}
@@ -19277,11 +18932,9 @@ func (m *awsAwsquery_deserializeOpUpdateServerCertific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateServerCertificate(response, &metadata)
 	}
@@ -19371,11 +19024,9 @@ func (m *awsAwsquery_deserializeOpUpdateServiceSpecifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateServiceSpecificCredential(response, &metadata)
 	}
@@ -19456,11 +19107,9 @@ func (m *awsAwsquery_deserializeOpUpdateSigningCertifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateSigningCertificate(response, &metadata)
 	}
@@ -19550,11 +19199,9 @@ func (m *awsAwsquery_deserializeOpUpdateSSHPublicKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateSSHPublicKey(response, &metadata)
 	}
@@ -19638,11 +19285,9 @@ func (m *awsAwsquery_deserializeOpUpdateUser) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUpdateUser(response, &metadata)
 	}
@@ -19738,11 +19383,9 @@ func (m *awsAwsquery_deserializeOpUploadServerCertific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUploadServerCertificate(response, &metadata)
 	}
@@ -19876,11 +19519,9 @@ func (m *awsAwsquery_deserializeOpUploadSigningCertifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUploadSigningCertificate(response, &metadata)
 	}
@@ -20017,11 +19658,9 @@ func (m *awsAwsquery_deserializeOpUploadSSHPublicKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUploadSSHPublicKey(response, &metadata)
 	}
