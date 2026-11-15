--- vendor/github.com/aws/aws-sdk-go-v2/service/sts/deserializers.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sts/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -50,11 +49,9 @@ func (m *awsAwsquery_deserializeOpAssumeRole) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRole(response, &metadata)
 	}
@@ -179,11 +176,9 @@ func (m *awsAwsquery_deserializeOpAssumeRoleWithSAML) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoleWithSAML(response, &metadata)
 	}
@@ -314,11 +309,9 @@ func (m *awsAwsquery_deserializeOpAssumeRoleWithWebIde
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoleWithWebIdentity(response, &metadata)
 	}
@@ -452,11 +445,9 @@ func (m *awsAwsquery_deserializeOpAssumeRoot) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAssumeRoot(response, &metadata)
 	}
@@ -575,11 +566,9 @@ func (m *awsAwsquery_deserializeOpDecodeAuthorizationM
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDecodeAuthorizationMessage(response, &metadata)
 	}
@@ -695,11 +684,9 @@ func (m *awsAwsquery_deserializeOpGetAccessKeyInfo) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetAccessKeyInfo(response, &metadata)
 	}
@@ -812,11 +799,9 @@ func (m *awsAwsquery_deserializeOpGetCallerIdentity) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetCallerIdentity(response, &metadata)
 	}
@@ -929,11 +914,9 @@ func (m *awsAwsquery_deserializeOpGetDelegatedAccessTo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetDelegatedAccessToken(response, &metadata)
 	}
@@ -1055,11 +1038,9 @@ func (m *awsAwsquery_deserializeOpGetFederationToken) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetFederationToken(response, &metadata)
 	}
@@ -1181,11 +1162,9 @@ func (m *awsAwsquery_deserializeOpGetSessionToken) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSessionToken(response, &metadata)
 	}
@@ -1301,11 +1280,9 @@ func (m *awsAwsquery_deserializeOpGetWebIdentityToken)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetWebIdentityToken(response, &metadata)
 	}
