--- vendor/github.com/aws/aws-sdk-go-v2/service/sts/deserializers.go.orig	2026-10-28 11:20:31 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sts/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsAwsquery_deserializeOpAssumeRole) HandleDe
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRole(response, &metadata)
@@ -169,10 +166,8 @@ func (m *awsAwsquery_deserializeOpAssumeRoleWithSAML) 
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoleWithSAML(response, &metadata)
@@ -299,10 +294,8 @@ func (m *awsAwsquery_deserializeOpAssumeRoleWithWebIde
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoleWithWebIdentity(response, &metadata)
@@ -432,10 +425,8 @@ func (m *awsAwsquery_deserializeOpAssumeRoot) HandleDe
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoot(response, &metadata)
@@ -550,10 +541,8 @@ func (m *awsAwsquery_deserializeOpDecodeAuthorizationM
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDecodeAuthorizationMessage(response, &metadata)
@@ -665,10 +654,8 @@ func (m *awsAwsquery_deserializeOpGetAccessKeyInfo) Ha
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccessKeyInfo(response, &metadata)
@@ -777,10 +764,8 @@ func (m *awsAwsquery_deserializeOpGetCallerIdentity) H
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetCallerIdentity(response, &metadata)
@@ -889,10 +874,8 @@ func (m *awsAwsquery_deserializeOpGetDelegatedAccessTo
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetDelegatedAccessToken(response, &metadata)
@@ -1010,10 +993,8 @@ func (m *awsAwsquery_deserializeOpGetFederationToken) 
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetFederationToken(response, &metadata)
@@ -1131,10 +1112,8 @@ func (m *awsAwsquery_deserializeOpGetSessionToken) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSessionToken(response, &metadata)
@@ -1246,10 +1225,8 @@ func (m *awsAwsquery_deserializeOpGetWebIdentityToken)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetWebIdentityToken(response, &metadata)
