--- vendor/github.com/aws/aws-sdk-go-v2/service/sso/deserializers.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sso/deserializers.go
@@ -13,7 +13,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -45,11 +44,8 @@ func (m *awsRestjson1_deserializeOpGetRoleCredentials)
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetRoleCredentials(response, &metadata)
@@ -84,8 +80,6 @@ func (m *awsRestjson1_deserializeOpGetRoleCredentials)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -214,11 +208,8 @@ func (m *awsRestjson1_deserializeOpListAccountRoles) H
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListAccountRoles(response, &metadata)
@@ -253,8 +244,6 @@ func (m *awsRestjson1_deserializeOpListAccountRoles) H
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -392,11 +381,8 @@ func (m *awsRestjson1_deserializeOpListAccounts) Handl
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListAccounts(response, &metadata)
@@ -431,8 +417,6 @@ func (m *awsRestjson1_deserializeOpListAccounts) Handl
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -570,11 +554,8 @@ func (m *awsRestjson1_deserializeOpLogout) HandleDeser
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorLogout(response, &metadata)
@@ -587,8 +568,6 @@ func (m *awsRestjson1_deserializeOpLogout) HandleDeser
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
