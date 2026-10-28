--- vendor/github.com/aws/aws-sdk-go-v2/service/ssooidc/deserializers.go.orig	2026-10-28 11:19:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ssooidc/deserializers.go
@@ -13,7 +13,6 @@ import (
 	smithyio "github.com/aws/smithy-go/io"
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -41,10 +40,8 @@ func (m *awsRestjson1_deserializeOpCreateToken) Handle
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateToken(response, &metadata)
@@ -79,8 +76,6 @@ func (m *awsRestjson1_deserializeOpCreateToken) Handle
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -270,10 +265,8 @@ func (m *awsRestjson1_deserializeOpCreateTokenWithIAM)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateTokenWithIAM(response, &metadata)
@@ -308,8 +301,6 @@ func (m *awsRestjson1_deserializeOpCreateTokenWithIAM)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -521,10 +512,8 @@ func (m *awsRestjson1_deserializeOpRegisterClient) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorRegisterClient(response, &metadata)
@@ -559,8 +548,6 @@ func (m *awsRestjson1_deserializeOpRegisterClient) Han
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -751,10 +738,8 @@ func (m *awsRestjson1_deserializeOpStartDeviceAuthoriz
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorStartDeviceAuthorization(response, &metadata)
@@ -789,8 +774,6 @@ func (m *awsRestjson1_deserializeOpStartDeviceAuthoriz
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
