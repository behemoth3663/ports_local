--- vendor/github.com/aws/aws-sdk-go-v2/service/signin/deserializers.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/signin/deserializers.go
@@ -13,7 +13,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -45,11 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2Token) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateOAuth2Token(response, &metadata)
@@ -84,8 +80,6 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2Token) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -214,11 +208,8 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2TokenWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateOAuth2TokenWithIAM(response, &metadata)
@@ -253,8 +244,6 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2TokenWi
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -409,11 +398,8 @@ func (m *awsRestjson1_deserializeOpDeleteConsoleAuthor
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteConsoleAuthorizationConfiguration(response, &metadata)
@@ -448,8 +434,6 @@ func (m *awsRestjson1_deserializeOpDeleteConsoleAuthor
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -603,19 +587,14 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePermi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteResourcePermissionStatement(response, &metadata)
 	}
 	output := &DeleteResourcePermissionStatementOutput{}
 	out.Result = output
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -711,11 +690,8 @@ func (m *awsRestjson1_deserializeOpGetConsoleAuthoriza
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetConsoleAuthorizationConfiguration(response, &metadata)
@@ -750,8 +726,6 @@ func (m *awsRestjson1_deserializeOpGetConsoleAuthoriza
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -905,11 +879,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetResourcePolicy(response, &metadata)
@@ -944,8 +915,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1074,11 +1043,8 @@ func (m *awsRestjson1_deserializeOpIntrospectOAuth2Tok
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorIntrospectOAuth2TokenWithIAM(response, &metadata)
@@ -1113,8 +1079,6 @@ func (m *awsRestjson1_deserializeOpIntrospectOAuth2Tok
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1376,11 +1340,8 @@ func (m *awsRestjson1_deserializeOpListResourcePermiss
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListResourcePermissionStatements(response, &metadata)
@@ -1415,8 +1376,6 @@ func (m *awsRestjson1_deserializeOpListResourcePermiss
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1557,11 +1516,8 @@ func (m *awsRestjson1_deserializeOpPutConsoleAuthoriza
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutConsoleAuthorizationConfiguration(response, &metadata)
@@ -1596,8 +1552,6 @@ func (m *awsRestjson1_deserializeOpPutConsoleAuthoriza
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1754,11 +1708,8 @@ func (m *awsRestjson1_deserializeOpPutResourcePermissi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutResourcePermissionStatement(response, &metadata)
@@ -1793,8 +1744,6 @@ func (m *awsRestjson1_deserializeOpPutResourcePermissi
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1933,19 +1882,14 @@ func (m *awsRestjson1_deserializeOpRevokeOAuth2TokenWi
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorRevokeOAuth2TokenWithIAM(response, &metadata)
 	}
 	output := &RevokeOAuth2TokenWithIAMOutput{}
 	out.Result = output
-
-	span.End()
 	return out, metadata, err
 }
 
