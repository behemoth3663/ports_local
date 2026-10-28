--- vendor/github.com/aws/aws-sdk-go-v2/service/sso/deserializers.go.orig	2026-10-28 15:34:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sso/deserializers.go
@@ -13,7 +13,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -41,10 +40,8 @@ func (m *awsRestjson1_deserializeOpGetRoleCredentials)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetRoleCredentials(response, &metadata)
@@ -79,8 +76,6 @@ func (m *awsRestjson1_deserializeOpGetRoleCredentials)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -205,10 +200,8 @@ func (m *awsRestjson1_deserializeOpListAccountRoles) H
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListAccountRoles(response, &metadata)
@@ -243,8 +236,6 @@ func (m *awsRestjson1_deserializeOpListAccountRoles) H
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -378,10 +369,8 @@ func (m *awsRestjson1_deserializeOpListAccounts) Handl
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListAccounts(response, &metadata)
@@ -416,8 +405,6 @@ func (m *awsRestjson1_deserializeOpListAccounts) Handl
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -551,10 +538,8 @@ func (m *awsRestjson1_deserializeOpLogout) HandleDeser
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorLogout(response, &metadata)
@@ -567,8 +552,6 @@ func (m *awsRestjson1_deserializeOpLogout) HandleDeser
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
