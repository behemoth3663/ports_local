--- vendor/github.com/aws/aws-sdk-go-v2/service/sts/deserializers.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sts/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -49,11 +48,8 @@ func (m *awsAwsquery_deserializeOpAssumeRole) HandleDe
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRole(response, &metadata)
@@ -178,11 +174,8 @@ func (m *awsAwsquery_deserializeOpAssumeRoleWithSAML) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoleWithSAML(response, &metadata)
@@ -313,11 +306,8 @@ func (m *awsAwsquery_deserializeOpAssumeRoleWithWebIde
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoleWithWebIdentity(response, &metadata)
@@ -451,11 +441,8 @@ func (m *awsAwsquery_deserializeOpAssumeRoot) HandleDe
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoot(response, &metadata)
@@ -574,11 +561,8 @@ func (m *awsAwsquery_deserializeOpDecodeAuthorizationM
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDecodeAuthorizationMessage(response, &metadata)
@@ -694,11 +678,8 @@ func (m *awsAwsquery_deserializeOpGetAccessKeyInfo) Ha
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccessKeyInfo(response, &metadata)
@@ -811,11 +792,8 @@ func (m *awsAwsquery_deserializeOpGetCallerIdentity) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetCallerIdentity(response, &metadata)
@@ -928,11 +906,8 @@ func (m *awsAwsquery_deserializeOpGetDelegatedAccessTo
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetDelegatedAccessToken(response, &metadata)
@@ -1054,11 +1029,8 @@ func (m *awsAwsquery_deserializeOpGetFederationToken) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetFederationToken(response, &metadata)
@@ -1180,11 +1152,8 @@ func (m *awsAwsquery_deserializeOpGetSessionToken) Han
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSessionToken(response, &metadata)
@@ -1300,11 +1269,8 @@ func (m *awsAwsquery_deserializeOpGetWebIdentityToken)
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetWebIdentityToken(response, &metadata)
