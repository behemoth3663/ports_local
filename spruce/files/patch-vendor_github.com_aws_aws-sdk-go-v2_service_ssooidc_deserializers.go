--- vendor/github.com/aws/aws-sdk-go-v2/service/ssooidc/deserializers.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ssooidc/deserializers.go
@@ -13,7 +13,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -45,11 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateToken) Handle
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateToken(response, &metadata)
@@ -84,8 +80,6 @@ func (m *awsRestjson1_deserializeOpCreateToken) Handle
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -279,11 +273,8 @@ func (m *awsRestjson1_deserializeOpCreateTokenWithIAM)
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateTokenWithIAM(response, &metadata)
@@ -318,8 +309,6 @@ func (m *awsRestjson1_deserializeOpCreateTokenWithIAM)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -535,11 +524,8 @@ func (m *awsRestjson1_deserializeOpRegisterClient) Han
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorRegisterClient(response, &metadata)
@@ -574,8 +560,6 @@ func (m *awsRestjson1_deserializeOpRegisterClient) Han
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -770,11 +754,8 @@ func (m *awsRestjson1_deserializeOpStartDeviceAuthoriz
 	if err != nil {
 		return out, metadata, err
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorStartDeviceAuthorization(response, &metadata)
@@ -809,8 +790,6 @@ func (m *awsRestjson1_deserializeOpStartDeviceAuthoriz
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
