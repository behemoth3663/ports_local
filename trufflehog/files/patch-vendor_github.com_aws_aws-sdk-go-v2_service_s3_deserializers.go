--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/deserializers.go.orig	2026-10-28 15:34:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3/deserializers.go
@@ -20,7 +20,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -57,11 +56,8 @@ func (m *awsRestxml_deserializeOpAbortMultipartUpload)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorAbortMultipartUpload(response, &metadata)
@@ -73,8 +69,6 @@ func (m *awsRestxml_deserializeOpAbortMultipartUpload)
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -155,11 +149,8 @@ func (m *awsRestxml_deserializeOpCompleteMultipartUplo
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCompleteMultipartUpload(response, &metadata)
@@ -199,8 +190,6 @@ func (m *awsRestxml_deserializeOpCompleteMultipartUplo
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -537,11 +526,8 @@ func (m *awsRestxml_deserializeOpCopyObject) HandleDes
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCopyObject(response, &metadata)
@@ -581,8 +567,6 @@ func (m *awsRestxml_deserializeOpCopyObject) HandleDes
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -753,11 +737,8 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateBucket(response, &metadata)
@@ -769,8 +750,6 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -859,11 +838,8 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataC
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateBucketMetadataConfiguration(response, &metadata)
@@ -876,8 +852,6 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataC
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -942,11 +916,8 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataT
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateBucketMetadataTableConfiguration(response, &metadata)
@@ -959,8 +930,6 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataT
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1025,11 +994,8 @@ func (m *awsRestxml_deserializeOpCreateMultipartUpload
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateMultipartUpload(response, &metadata)
@@ -1069,8 +1035,6 @@ func (m *awsRestxml_deserializeOpCreateMultipartUpload
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1280,11 +1244,8 @@ func (m *awsRestxml_deserializeOpCreateSession) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateSession(response, &metadata)
@@ -1324,8 +1285,6 @@ func (m *awsRestxml_deserializeOpCreateSession) Handle
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1466,11 +1425,8 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucket(response, &metadata)
@@ -1483,8 +1439,6 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1549,11 +1503,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketAnalytics
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketAnalyticsConfiguration(response, &metadata)
@@ -1566,8 +1517,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketAnalytics
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1632,11 +1581,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketCors) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketCors(response, &metadata)
@@ -1649,8 +1595,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketCors) Han
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1715,11 +1659,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketEncryptio
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketEncryption(response, &metadata)
@@ -1732,8 +1673,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketEncryptio
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1798,11 +1737,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketIntellige
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketIntelligentTieringConfiguration(response, &metadata)
@@ -1815,8 +1751,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketIntellige
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1881,11 +1815,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketInventory
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketInventoryConfiguration(response, &metadata)
@@ -1898,8 +1829,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketInventory
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -1964,11 +1893,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketLifecycle(response, &metadata)
@@ -1981,8 +1907,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2047,11 +1971,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataC
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketMetadataConfiguration(response, &metadata)
@@ -2064,8 +1985,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataC
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2130,11 +2049,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataT
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketMetadataTableConfiguration(response, &metadata)
@@ -2147,8 +2063,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataT
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2213,11 +2127,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetricsCo
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketMetricsConfiguration(response, &metadata)
@@ -2230,8 +2141,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetricsCo
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2296,11 +2205,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketOwnership
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketOwnershipControls(response, &metadata)
@@ -2313,8 +2219,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketOwnership
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2379,11 +2283,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketPolicy(response, &metadata)
@@ -2396,8 +2297,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2462,11 +2361,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketReplication(response, &metadata)
@@ -2479,8 +2375,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2545,11 +2439,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketTagging(response, &metadata)
@@ -2562,8 +2453,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2628,11 +2517,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketWebsite) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketWebsite(response, &metadata)
@@ -2645,8 +2531,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketWebsite) 
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2711,11 +2595,8 @@ func (m *awsRestxml_deserializeOpDeleteObject) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObject(response, &metadata)
@@ -2727,8 +2608,6 @@ func (m *awsRestxml_deserializeOpDeleteObject) HandleD
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2820,11 +2699,8 @@ func (m *awsRestxml_deserializeOpDeleteObjectAnnotatio
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObjectAnnotation(response, &metadata)
@@ -2836,8 +2712,6 @@ func (m *awsRestxml_deserializeOpDeleteObjectAnnotatio
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -2926,11 +2800,8 @@ func (m *awsRestxml_deserializeOpDeleteObjects) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObjects(response, &metadata)
@@ -2970,8 +2841,6 @@ func (m *awsRestxml_deserializeOpDeleteObjects) Handle
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -3096,11 +2965,8 @@ func (m *awsRestxml_deserializeOpDeleteObjectTagging) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObjectTagging(response, &metadata)
@@ -3112,8 +2978,6 @@ func (m *awsRestxml_deserializeOpDeleteObjectTagging) 
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -3191,11 +3055,8 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeletePublicAccessBlock(response, &metadata)
@@ -3208,8 +3069,6 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -3274,11 +3133,8 @@ func (m *awsRestxml_deserializeOpGetBucketAbac) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAbac(response, &metadata)
@@ -3313,8 +3169,6 @@ func (m *awsRestxml_deserializeOpGetBucketAbac) Handle
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -3421,11 +3275,8 @@ func (m *awsRestxml_deserializeOpGetBucketAccelerateCo
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAccelerateConfiguration(response, &metadata)
@@ -3465,8 +3316,6 @@ func (m *awsRestxml_deserializeOpGetBucketAccelerateCo
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -3592,11 +3441,8 @@ func (m *awsRestxml_deserializeOpGetBucketAcl) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAcl(response, &metadata)
@@ -3631,8 +3477,6 @@ func (m *awsRestxml_deserializeOpGetBucketAcl) HandleD
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -3745,11 +3589,8 @@ func (m *awsRestxml_deserializeOpGetBucketAnalyticsCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAnalyticsConfiguration(response, &metadata)
@@ -3784,8 +3625,6 @@ func (m *awsRestxml_deserializeOpGetBucketAnalyticsCon
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -3892,11 +3731,8 @@ func (m *awsRestxml_deserializeOpGetBucketCors) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketCors(response, &metadata)
@@ -3931,8 +3767,6 @@ func (m *awsRestxml_deserializeOpGetBucketCors) Handle
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -4039,11 +3873,8 @@ func (m *awsRestxml_deserializeOpGetBucketEncryption) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketEncryption(response, &metadata)
@@ -4078,8 +3909,6 @@ func (m *awsRestxml_deserializeOpGetBucketEncryption) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -4186,11 +4015,8 @@ func (m *awsRestxml_deserializeOpGetBucketIntelligentT
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketIntelligentTieringConfiguration(response, &metadata)
@@ -4225,8 +4051,6 @@ func (m *awsRestxml_deserializeOpGetBucketIntelligentT
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -4333,11 +4157,8 @@ func (m *awsRestxml_deserializeOpGetBucketInventoryCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketInventoryConfiguration(response, &metadata)
@@ -4372,8 +4193,6 @@ func (m *awsRestxml_deserializeOpGetBucketInventoryCon
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -4480,11 +4299,8 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLifecycleConfiguration(response, &metadata)
@@ -4524,8 +4340,6 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -4644,11 +4458,8 @@ func (m *awsRestxml_deserializeOpGetBucketLocation) Ha
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLocation(response, &metadata)
@@ -4683,8 +4494,6 @@ func (m *awsRestxml_deserializeOpGetBucketLocation) Ha
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -4798,11 +4607,8 @@ func (m *awsRestxml_deserializeOpGetBucketLogging) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLogging(response, &metadata)
@@ -4837,8 +4643,6 @@ func (m *awsRestxml_deserializeOpGetBucketLogging) Han
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -4945,11 +4749,8 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataConf
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketMetadataConfiguration(response, &metadata)
@@ -4984,8 +4785,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataConf
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -5092,11 +4891,8 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataTabl
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketMetadataTableConfiguration(response, &metadata)
@@ -5131,8 +4927,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataTabl
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -5239,11 +5033,8 @@ func (m *awsRestxml_deserializeOpGetBucketMetricsConfi
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketMetricsConfiguration(response, &metadata)
@@ -5278,8 +5069,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetricsConfi
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -5386,11 +5175,8 @@ func (m *awsRestxml_deserializeOpGetBucketNotification
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketNotificationConfiguration(response, &metadata)
@@ -5425,8 +5211,6 @@ func (m *awsRestxml_deserializeOpGetBucketNotification
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -5551,11 +5335,8 @@ func (m *awsRestxml_deserializeOpGetBucketOwnershipCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketOwnershipControls(response, &metadata)
@@ -5590,8 +5371,6 @@ func (m *awsRestxml_deserializeOpGetBucketOwnershipCon
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -5698,11 +5477,8 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketPolicy(response, &metadata)
@@ -5714,8 +5490,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -5801,11 +5575,8 @@ func (m *awsRestxml_deserializeOpGetBucketPolicyStatus
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketPolicyStatus(response, &metadata)
@@ -5840,8 +5611,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicyStatus
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -5948,11 +5717,8 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketReplication(response, &metadata)
@@ -5987,8 +5753,6 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -6095,11 +5859,8 @@ func (m *awsRestxml_deserializeOpGetBucketRequestPayme
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketRequestPayment(response, &metadata)
@@ -6134,8 +5895,6 @@ func (m *awsRestxml_deserializeOpGetBucketRequestPayme
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -6249,11 +6008,8 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketTagging(response, &metadata)
@@ -6288,8 +6044,6 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -6396,11 +6150,8 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketVersioning(response, &metadata)
@@ -6435,8 +6186,6 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -6563,11 +6312,8 @@ func (m *awsRestxml_deserializeOpGetBucketWebsite) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketWebsite(response, &metadata)
@@ -6602,8 +6348,6 @@ func (m *awsRestxml_deserializeOpGetBucketWebsite) Han
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -6728,11 +6472,8 @@ func (m *awsRestxml_deserializeOpGetObject) HandleDese
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, true, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObject(response, &metadata)
@@ -6749,8 +6490,6 @@ func (m *awsRestxml_deserializeOpGetObject) HandleDese
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -7092,11 +6831,8 @@ func (m *awsRestxml_deserializeOpGetObjectAcl) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectAcl(response, &metadata)
@@ -7136,8 +6872,6 @@ func (m *awsRestxml_deserializeOpGetObjectAcl) HandleD
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -7265,11 +6999,8 @@ func (m *awsRestxml_deserializeOpGetObjectAnnotation) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, true, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectAnnotation(response, &metadata)
@@ -7286,8 +7017,6 @@ func (m *awsRestxml_deserializeOpGetObjectAnnotation) 
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -7474,11 +7203,8 @@ func (m *awsRestxml_deserializeOpGetObjectAttributes) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectAttributes(response, &metadata)
@@ -7518,8 +7244,6 @@ func (m *awsRestxml_deserializeOpGetObjectAttributes) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -7713,11 +7437,8 @@ func (m *awsRestxml_deserializeOpGetObjectLegalHold) H
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectLegalHold(response, &metadata)
@@ -7752,8 +7473,6 @@ func (m *awsRestxml_deserializeOpGetObjectLegalHold) H
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -7860,11 +7579,8 @@ func (m *awsRestxml_deserializeOpGetObjectLockConfigur
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectLockConfiguration(response, &metadata)
@@ -7899,8 +7615,6 @@ func (m *awsRestxml_deserializeOpGetObjectLockConfigur
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -8007,11 +7721,8 @@ func (m *awsRestxml_deserializeOpGetObjectRetention) H
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectRetention(response, &metadata)
@@ -8046,8 +7757,6 @@ func (m *awsRestxml_deserializeOpGetObjectRetention) H
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -8154,11 +7863,8 @@ func (m *awsRestxml_deserializeOpGetObjectTagging) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectTagging(response, &metadata)
@@ -8198,8 +7904,6 @@ func (m *awsRestxml_deserializeOpGetObjectTagging) Han
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -8318,11 +8022,8 @@ func (m *awsRestxml_deserializeOpGetObjectTorrent) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, true, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectTorrent(response, &metadata)
@@ -8339,8 +8040,6 @@ func (m *awsRestxml_deserializeOpGetObjectTorrent) Han
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -8425,11 +8124,8 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetPublicAccessBlock(response, &metadata)
@@ -8464,8 +8160,6 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -8572,11 +8266,8 @@ func (m *awsRestxml_deserializeOpHeadBucket) HandleDes
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorHeadBucket(response, &metadata)
@@ -8588,8 +8279,6 @@ func (m *awsRestxml_deserializeOpHeadBucket) HandleDes
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -8694,11 +8383,8 @@ func (m *awsRestxml_deserializeOpHeadObject) HandleDes
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorHeadObject(response, &metadata)
@@ -8710,8 +8396,6 @@ func (m *awsRestxml_deserializeOpHeadObject) HandleDes
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -9048,11 +8732,8 @@ func (m *awsRestxml_deserializeOpListBucketAnalyticsCo
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketAnalyticsConfigurations(response, &metadata)
@@ -9087,8 +8768,6 @@ func (m *awsRestxml_deserializeOpListBucketAnalyticsCo
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -9237,11 +8916,8 @@ func (m *awsRestxml_deserializeOpListBucketIntelligent
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketIntelligentTieringConfigurations(response, &metadata)
@@ -9276,8 +8952,6 @@ func (m *awsRestxml_deserializeOpListBucketIntelligent
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -9426,11 +9100,8 @@ func (m *awsRestxml_deserializeOpListBucketInventoryCo
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketInventoryConfigurations(response, &metadata)
@@ -9465,8 +9136,6 @@ func (m *awsRestxml_deserializeOpListBucketInventoryCo
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -9615,11 +9284,8 @@ func (m *awsRestxml_deserializeOpListBucketMetricsConf
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketMetricsConfigurations(response, &metadata)
@@ -9654,8 +9320,6 @@ func (m *awsRestxml_deserializeOpListBucketMetricsConf
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -9804,11 +9468,8 @@ func (m *awsRestxml_deserializeOpListBuckets) HandleDe
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBuckets(response, &metadata)
@@ -9843,8 +9504,6 @@ func (m *awsRestxml_deserializeOpListBuckets) HandleDe
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -9983,11 +9642,8 @@ func (m *awsRestxml_deserializeOpListDirectoryBuckets)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDirectoryBuckets(response, &metadata)
@@ -10022,8 +9678,6 @@ func (m *awsRestxml_deserializeOpListDirectoryBuckets)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -10143,11 +9797,8 @@ func (m *awsRestxml_deserializeOpListMultipartUploads)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListMultipartUploads(response, &metadata)
@@ -10187,8 +9838,6 @@ func (m *awsRestxml_deserializeOpListMultipartUploads)
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -10450,11 +10099,8 @@ func (m *awsRestxml_deserializeOpListObjectAnnotations
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjectAnnotations(response, &metadata)
@@ -10494,8 +10140,6 @@ func (m *awsRestxml_deserializeOpListObjectAnnotations
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -10727,11 +10371,8 @@ func (m *awsRestxml_deserializeOpListObjects) HandleDe
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjects(response, &metadata)
@@ -10771,8 +10412,6 @@ func (m *awsRestxml_deserializeOpListObjects) HandleDe
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -11011,11 +10650,8 @@ func (m *awsRestxml_deserializeOpListObjectsV2) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjectsV2(response, &metadata)
@@ -11055,8 +10691,6 @@ func (m *awsRestxml_deserializeOpListObjectsV2) Handle
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -11325,11 +10959,8 @@ func (m *awsRestxml_deserializeOpListObjectVersions) H
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjectVersions(response, &metadata)
@@ -11369,8 +11000,6 @@ func (m *awsRestxml_deserializeOpListObjectVersions) H
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -11638,11 +11267,8 @@ func (m *awsRestxml_deserializeOpListParts) HandleDese
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListParts(response, &metadata)
@@ -11682,8 +11308,6 @@ func (m *awsRestxml_deserializeOpListParts) HandleDese
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -11965,11 +11589,8 @@ func (m *awsRestxml_deserializeOpPutBucketAbac) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAbac(response, &metadata)
@@ -11982,8 +11603,6 @@ func (m *awsRestxml_deserializeOpPutBucketAbac) Handle
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12048,11 +11667,8 @@ func (m *awsRestxml_deserializeOpPutBucketAccelerateCo
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAccelerateConfiguration(response, &metadata)
@@ -12065,8 +11681,6 @@ func (m *awsRestxml_deserializeOpPutBucketAccelerateCo
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12131,11 +11745,8 @@ func (m *awsRestxml_deserializeOpPutBucketAcl) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAcl(response, &metadata)
@@ -12148,8 +11759,6 @@ func (m *awsRestxml_deserializeOpPutBucketAcl) HandleD
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12214,11 +11823,8 @@ func (m *awsRestxml_deserializeOpPutBucketAnalyticsCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAnalyticsConfiguration(response, &metadata)
@@ -12231,8 +11837,6 @@ func (m *awsRestxml_deserializeOpPutBucketAnalyticsCon
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12297,11 +11901,8 @@ func (m *awsRestxml_deserializeOpPutBucketCors) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketCors(response, &metadata)
@@ -12314,8 +11915,6 @@ func (m *awsRestxml_deserializeOpPutBucketCors) Handle
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12380,11 +11979,8 @@ func (m *awsRestxml_deserializeOpPutBucketEncryption) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketEncryption(response, &metadata)
@@ -12397,8 +11993,6 @@ func (m *awsRestxml_deserializeOpPutBucketEncryption) 
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12463,11 +12057,8 @@ func (m *awsRestxml_deserializeOpPutBucketIntelligentT
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketIntelligentTieringConfiguration(response, &metadata)
@@ -12480,8 +12071,6 @@ func (m *awsRestxml_deserializeOpPutBucketIntelligentT
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12546,11 +12135,8 @@ func (m *awsRestxml_deserializeOpPutBucketInventoryCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketInventoryConfiguration(response, &metadata)
@@ -12563,8 +12149,6 @@ func (m *awsRestxml_deserializeOpPutBucketInventoryCon
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12629,11 +12213,8 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketLifecycleConfiguration(response, &metadata)
@@ -12645,8 +12226,6 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12724,11 +12303,8 @@ func (m *awsRestxml_deserializeOpPutBucketLogging) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketLogging(response, &metadata)
@@ -12741,8 +12317,6 @@ func (m *awsRestxml_deserializeOpPutBucketLogging) Han
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12807,11 +12381,8 @@ func (m *awsRestxml_deserializeOpPutBucketMetricsConfi
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketMetricsConfiguration(response, &metadata)
@@ -12824,8 +12395,6 @@ func (m *awsRestxml_deserializeOpPutBucketMetricsConfi
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12890,11 +12459,8 @@ func (m *awsRestxml_deserializeOpPutBucketNotification
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketNotificationConfiguration(response, &metadata)
@@ -12907,8 +12473,6 @@ func (m *awsRestxml_deserializeOpPutBucketNotification
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -12973,11 +12537,8 @@ func (m *awsRestxml_deserializeOpPutBucketOwnershipCon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketOwnershipControls(response, &metadata)
@@ -12990,8 +12551,6 @@ func (m *awsRestxml_deserializeOpPutBucketOwnershipCon
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13056,11 +12615,8 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketPolicy(response, &metadata)
@@ -13073,8 +12629,6 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13139,11 +12693,8 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketReplication(response, &metadata)
@@ -13156,8 +12707,6 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13222,11 +12771,8 @@ func (m *awsRestxml_deserializeOpPutBucketRequestPayme
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketRequestPayment(response, &metadata)
@@ -13239,8 +12785,6 @@ func (m *awsRestxml_deserializeOpPutBucketRequestPayme
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13305,11 +12849,8 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketTagging(response, &metadata)
@@ -13322,8 +12863,6 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13388,11 +12927,8 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketVersioning(response, &metadata)
@@ -13405,8 +12941,6 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13471,11 +13005,8 @@ func (m *awsRestxml_deserializeOpPutBucketWebsite) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketWebsite(response, &metadata)
@@ -13488,8 +13019,6 @@ func (m *awsRestxml_deserializeOpPutBucketWebsite) Han
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13554,11 +13083,8 @@ func (m *awsRestxml_deserializeOpPutObject) HandleDese
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObject(response, &metadata)
@@ -13570,8 +13096,6 @@ func (m *awsRestxml_deserializeOpPutObject) HandleDese
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13774,11 +13298,8 @@ func (m *awsRestxml_deserializeOpPutObjectAcl) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectAcl(response, &metadata)
@@ -13790,8 +13311,6 @@ func (m *awsRestxml_deserializeOpPutObjectAcl) HandleD
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -13872,11 +13391,8 @@ func (m *awsRestxml_deserializeOpPutObjectAnnotation) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectAnnotation(response, &metadata)
@@ -13916,8 +13432,6 @@ func (m *awsRestxml_deserializeOpPutObjectAnnotation) 
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14147,11 +13661,8 @@ func (m *awsRestxml_deserializeOpPutObjectLegalHold) H
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectLegalHold(response, &metadata)
@@ -14163,8 +13674,6 @@ func (m *awsRestxml_deserializeOpPutObjectLegalHold) H
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14242,11 +13751,8 @@ func (m *awsRestxml_deserializeOpPutObjectLockConfigur
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectLockConfiguration(response, &metadata)
@@ -14258,8 +13764,6 @@ func (m *awsRestxml_deserializeOpPutObjectLockConfigur
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14337,11 +13841,8 @@ func (m *awsRestxml_deserializeOpPutObjectRetention) H
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectRetention(response, &metadata)
@@ -14353,8 +13854,6 @@ func (m *awsRestxml_deserializeOpPutObjectRetention) H
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14432,11 +13931,8 @@ func (m *awsRestxml_deserializeOpPutObjectTagging) Han
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectTagging(response, &metadata)
@@ -14448,8 +13944,6 @@ func (m *awsRestxml_deserializeOpPutObjectTagging) Han
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14527,11 +14021,8 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutPublicAccessBlock(response, &metadata)
@@ -14544,8 +14035,6 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14610,19 +14099,14 @@ func (m *awsRestxml_deserializeOpRenameObject) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorRenameObject(response, &metadata)
 	}
 	output := &RenameObjectOutput{}
 	out.Result = output
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14690,11 +14174,8 @@ func (m *awsRestxml_deserializeOpRestoreObject) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorRestoreObject(response, &metadata)
@@ -14706,8 +14187,6 @@ func (m *awsRestxml_deserializeOpRestoreObject) Handle
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14791,19 +14270,14 @@ func (m *awsRestxml_deserializeOpSelectObjectContent) 
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
 	}
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorSelectObjectContent(response, &metadata)
 	}
 	output := &SelectObjectContentOutput{}
 	out.Result = output
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14868,11 +14342,8 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataA
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateBucketMetadataAnnotationTableConfiguration(response, &metadata)
@@ -14885,8 +14356,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataA
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -14951,11 +14420,8 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataI
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateBucketMetadataInventoryTableConfiguration(response, &metadata)
@@ -14968,8 +14434,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataI
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -15034,11 +14498,8 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataJ
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateBucketMetadataJournalTableConfiguration(response, &metadata)
@@ -15051,8 +14512,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataJ
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -15117,11 +14576,8 @@ func (m *awsRestxml_deserializeOpUpdateObjectEncryptio
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateObjectEncryption(response, &metadata)
@@ -15133,8 +14589,6 @@ func (m *awsRestxml_deserializeOpUpdateObjectEncryptio
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -15221,11 +14675,8 @@ func (m *awsRestxml_deserializeOpUploadPart) HandleDes
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUploadPart(response, &metadata)
@@ -15237,8 +14688,6 @@ func (m *awsRestxml_deserializeOpUploadPart) HandleDes
 	if err != nil {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -15400,11 +14849,8 @@ func (m *awsRestxml_deserializeOpUploadPartCopy) Handl
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUploadPartCopy(response, &metadata)
@@ -15444,8 +14890,6 @@ func (m *awsRestxml_deserializeOpUploadPartCopy) Handl
 			Snapshot: snapshot.Bytes(),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
@@ -15598,11 +15042,8 @@ func (m *awsRestxml_deserializeOpWriteGetObjectRespons
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorWriteGetObjectResponse(response, &metadata)
@@ -15615,8 +15056,6 @@ func (m *awsRestxml_deserializeOpWriteGetObjectRespons
 			Err: fmt.Errorf("failed to discard response body, %w", err),
 		}
 	}
-
-	span.End()
 	return out, metadata, err
 }
 
