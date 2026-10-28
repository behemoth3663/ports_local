--- vendor/github.com/aws/aws-sdk-go-v2/service/signin/deserializers.go.orig	2026-10-28 15:34:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/signin/deserializers.go
@@ -13,7 +13,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -41,10 +40,8 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2Token) 
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateOAuth2Token(response, &metadata)
@@ -79,8 +76,6 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2Token) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -205,10 +200,8 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2TokenWi
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateOAuth2TokenWithIAM(response, &metadata)
@@ -243,8 +236,6 @@ func (m *awsRestjson1_deserializeOpCreateOAuth2TokenWi
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -395,10 +386,8 @@ func (m *awsRestjson1_deserializeOpDeleteConsoleAuthor
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteConsoleAuthorizationConfiguration(response, &metadata)
@@ -433,8 +422,6 @@ func (m *awsRestjson1_deserializeOpDeleteConsoleAuthor
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -584,10 +571,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePermi
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteResourcePermissionStatement(response, &metadata)
@@ -595,7 +580,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePermi
 	output := &DeleteResourcePermissionStatementOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -687,10 +671,8 @@ func (m *awsRestjson1_deserializeOpGetConsoleAuthoriza
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetConsoleAuthorizationConfiguration(response, &metadata)
@@ -725,8 +707,6 @@ func (m *awsRestjson1_deserializeOpGetConsoleAuthoriza
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -876,10 +856,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetResourcePolicy(response, &metadata)
@@ -914,8 +892,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1040,10 +1016,8 @@ func (m *awsRestjson1_deserializeOpIntrospectOAuth2Tok
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorIntrospectOAuth2TokenWithIAM(response, &metadata)
@@ -1078,8 +1052,6 @@ func (m *awsRestjson1_deserializeOpIntrospectOAuth2Tok
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1337,10 +1309,8 @@ func (m *awsRestjson1_deserializeOpListResourcePermiss
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListResourcePermissionStatements(response, &metadata)
@@ -1375,8 +1345,6 @@ func (m *awsRestjson1_deserializeOpListResourcePermiss
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1513,10 +1481,8 @@ func (m *awsRestjson1_deserializeOpPutConsoleAuthoriza
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutConsoleAuthorizationConfiguration(response, &metadata)
@@ -1551,8 +1517,6 @@ func (m *awsRestjson1_deserializeOpPutConsoleAuthoriza
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1705,10 +1669,8 @@ func (m *awsRestjson1_deserializeOpPutResourcePermissi
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutResourcePermissionStatement(response, &metadata)
@@ -1743,8 +1705,6 @@ func (m *awsRestjson1_deserializeOpPutResourcePermissi
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1879,10 +1839,8 @@ func (m *awsRestjson1_deserializeOpRevokeOAuth2TokenWi
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorRevokeOAuth2TokenWithIAM(response, &metadata)
@@ -1890,7 +1848,6 @@ func (m *awsRestjson1_deserializeOpRevokeOAuth2TokenWi
 	output := &RevokeOAuth2TokenWithIAMOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
