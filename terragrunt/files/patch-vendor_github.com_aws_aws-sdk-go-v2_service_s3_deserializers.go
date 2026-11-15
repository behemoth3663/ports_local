--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/deserializers.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3/deserializers.go
@@ -20,7 +20,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -63,11 +62,9 @@ func (m *awsRestxml_deserializeOpAbortMultipartUpload)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorAbortMultipartUpload(response, &metadata)
 	}
@@ -79,7 +76,6 @@ func (m *awsRestxml_deserializeOpAbortMultipartUpload)
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -166,11 +162,9 @@ func (m *awsRestxml_deserializeOpCompleteMultipartUplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCompleteMultipartUpload(response, &metadata)
 	}
@@ -210,7 +204,6 @@ func (m *awsRestxml_deserializeOpCompleteMultipartUplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -553,11 +546,9 @@ func (m *awsRestxml_deserializeOpCopyObject) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCopyObject(response, &metadata)
 	}
@@ -597,7 +588,6 @@ func (m *awsRestxml_deserializeOpCopyObject) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -774,11 +764,9 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateBucket(response, &metadata)
 	}
@@ -790,7 +778,6 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -885,11 +872,9 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateBucketMetadataConfiguration(response, &metadata)
 	}
@@ -902,7 +887,6 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -973,11 +957,9 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateBucketMetadataTableConfiguration(response, &metadata)
 	}
@@ -990,7 +972,6 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1061,11 +1042,9 @@ func (m *awsRestxml_deserializeOpCreateMultipartUpload
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateMultipartUpload(response, &metadata)
 	}
@@ -1105,7 +1084,6 @@ func (m *awsRestxml_deserializeOpCreateMultipartUpload
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1321,11 +1299,9 @@ func (m *awsRestxml_deserializeOpCreateSession) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateSession(response, &metadata)
 	}
@@ -1365,7 +1341,6 @@ func (m *awsRestxml_deserializeOpCreateSession) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1512,11 +1487,9 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucket(response, &metadata)
 	}
@@ -1529,7 +1502,6 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1600,11 +1572,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketAnalytics
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketAnalyticsConfiguration(response, &metadata)
 	}
@@ -1617,7 +1587,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketAnalytics
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1688,11 +1657,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketCors) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketCors(response, &metadata)
 	}
@@ -1705,7 +1672,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketCors) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1776,11 +1742,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketEncryptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketEncryption(response, &metadata)
 	}
@@ -1793,7 +1757,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketEncryptio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1864,11 +1827,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketIntellige
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketIntelligentTieringConfiguration(response, &metadata)
 	}
@@ -1881,7 +1842,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketIntellige
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1952,11 +1912,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketInventory
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketInventoryConfiguration(response, &metadata)
 	}
@@ -1969,7 +1927,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketInventory
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2040,11 +1997,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketLifecycle(response, &metadata)
 	}
@@ -2057,7 +2012,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2128,11 +2082,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketMetadataConfiguration(response, &metadata)
 	}
@@ -2145,7 +2097,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2216,11 +2167,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketMetadataTableConfiguration(response, &metadata)
 	}
@@ -2233,7 +2182,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2304,11 +2252,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetricsCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketMetricsConfiguration(response, &metadata)
 	}
@@ -2321,7 +2267,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetricsCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2392,11 +2337,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketOwnership
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketOwnershipControls(response, &metadata)
 	}
@@ -2409,7 +2352,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketOwnership
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2480,11 +2422,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketPolicy(response, &metadata)
 	}
@@ -2497,7 +2437,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2568,11 +2507,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketReplication(response, &metadata)
 	}
@@ -2585,7 +2522,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2656,11 +2592,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketTagging(response, &metadata)
 	}
@@ -2673,7 +2607,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2744,11 +2677,9 @@ func (m *awsRestxml_deserializeOpDeleteBucketWebsite) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketWebsite(response, &metadata)
 	}
@@ -2761,7 +2692,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketWebsite) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2832,11 +2762,9 @@ func (m *awsRestxml_deserializeOpDeleteObject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObject(response, &metadata)
 	}
@@ -2848,7 +2776,6 @@ func (m *awsRestxml_deserializeOpDeleteObject) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2946,11 +2873,9 @@ func (m *awsRestxml_deserializeOpDeleteObjectAnnotatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObjectAnnotation(response, &metadata)
 	}
@@ -2962,7 +2887,6 @@ func (m *awsRestxml_deserializeOpDeleteObjectAnnotatio
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3057,11 +2981,9 @@ func (m *awsRestxml_deserializeOpDeleteObjects) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObjects(response, &metadata)
 	}
@@ -3101,7 +3023,6 @@ func (m *awsRestxml_deserializeOpDeleteObjects) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3232,11 +3153,9 @@ func (m *awsRestxml_deserializeOpDeleteObjectTagging) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteObjectTagging(response, &metadata)
 	}
@@ -3248,7 +3167,6 @@ func (m *awsRestxml_deserializeOpDeleteObjectTagging) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3332,11 +3250,9 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeletePublicAccessBlock(response, &metadata)
 	}
@@ -3349,7 +3265,6 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3420,11 +3335,9 @@ func (m *awsRestxml_deserializeOpGetBucketAbac) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAbac(response, &metadata)
 	}
@@ -3459,7 +3372,6 @@ func (m *awsRestxml_deserializeOpGetBucketAbac) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3572,11 +3484,9 @@ func (m *awsRestxml_deserializeOpGetBucketAccelerateCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAccelerateConfiguration(response, &metadata)
 	}
@@ -3616,7 +3526,6 @@ func (m *awsRestxml_deserializeOpGetBucketAccelerateCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3748,11 +3657,9 @@ func (m *awsRestxml_deserializeOpGetBucketAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAcl(response, &metadata)
 	}
@@ -3787,7 +3694,6 @@ func (m *awsRestxml_deserializeOpGetBucketAcl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3906,11 +3812,9 @@ func (m *awsRestxml_deserializeOpGetBucketAnalyticsCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketAnalyticsConfiguration(response, &metadata)
 	}
@@ -3945,7 +3849,6 @@ func (m *awsRestxml_deserializeOpGetBucketAnalyticsCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4058,11 +3961,9 @@ func (m *awsRestxml_deserializeOpGetBucketCors) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketCors(response, &metadata)
 	}
@@ -4097,7 +3998,6 @@ func (m *awsRestxml_deserializeOpGetBucketCors) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4210,11 +4110,9 @@ func (m *awsRestxml_deserializeOpGetBucketEncryption) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketEncryption(response, &metadata)
 	}
@@ -4249,7 +4147,6 @@ func (m *awsRestxml_deserializeOpGetBucketEncryption) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4362,11 +4259,9 @@ func (m *awsRestxml_deserializeOpGetBucketIntelligentT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketIntelligentTieringConfiguration(response, &metadata)
 	}
@@ -4401,7 +4296,6 @@ func (m *awsRestxml_deserializeOpGetBucketIntelligentT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4514,11 +4408,9 @@ func (m *awsRestxml_deserializeOpGetBucketInventoryCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketInventoryConfiguration(response, &metadata)
 	}
@@ -4553,7 +4445,6 @@ func (m *awsRestxml_deserializeOpGetBucketInventoryCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4666,11 +4557,9 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLifecycleConfiguration(response, &metadata)
 	}
@@ -4710,7 +4599,6 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4835,11 +4723,9 @@ func (m *awsRestxml_deserializeOpGetBucketLocation) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLocation(response, &metadata)
 	}
@@ -4874,7 +4760,6 @@ func (m *awsRestxml_deserializeOpGetBucketLocation) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4994,11 +4879,9 @@ func (m *awsRestxml_deserializeOpGetBucketLogging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLogging(response, &metadata)
 	}
@@ -5033,7 +4916,6 @@ func (m *awsRestxml_deserializeOpGetBucketLogging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5146,11 +5028,9 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketMetadataConfiguration(response, &metadata)
 	}
@@ -5185,7 +5065,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5298,11 +5177,9 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataTabl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketMetadataTableConfiguration(response, &metadata)
 	}
@@ -5337,7 +5214,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataTabl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5450,11 +5326,9 @@ func (m *awsRestxml_deserializeOpGetBucketMetricsConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketMetricsConfiguration(response, &metadata)
 	}
@@ -5489,7 +5363,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetricsConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5602,11 +5475,9 @@ func (m *awsRestxml_deserializeOpGetBucketNotification
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketNotificationConfiguration(response, &metadata)
 	}
@@ -5641,7 +5512,6 @@ func (m *awsRestxml_deserializeOpGetBucketNotification
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5772,11 +5642,9 @@ func (m *awsRestxml_deserializeOpGetBucketOwnershipCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketOwnershipControls(response, &metadata)
 	}
@@ -5811,7 +5679,6 @@ func (m *awsRestxml_deserializeOpGetBucketOwnershipCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5924,11 +5791,9 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketPolicy(response, &metadata)
 	}
@@ -5940,7 +5805,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6032,11 +5896,9 @@ func (m *awsRestxml_deserializeOpGetBucketPolicyStatus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketPolicyStatus(response, &metadata)
 	}
@@ -6071,7 +5933,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicyStatus
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6184,11 +6045,9 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketReplication(response, &metadata)
 	}
@@ -6223,7 +6082,6 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6336,11 +6194,9 @@ func (m *awsRestxml_deserializeOpGetBucketRequestPayme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketRequestPayment(response, &metadata)
 	}
@@ -6375,7 +6231,6 @@ func (m *awsRestxml_deserializeOpGetBucketRequestPayme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6495,11 +6350,9 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketTagging(response, &metadata)
 	}
@@ -6534,7 +6387,6 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6647,11 +6499,9 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketVersioning(response, &metadata)
 	}
@@ -6686,7 +6536,6 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6819,11 +6668,9 @@ func (m *awsRestxml_deserializeOpGetBucketWebsite) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketWebsite(response, &metadata)
 	}
@@ -6858,7 +6705,6 @@ func (m *awsRestxml_deserializeOpGetBucketWebsite) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6989,11 +6835,9 @@ func (m *awsRestxml_deserializeOpGetObject) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObject(response, &metadata)
 	}
@@ -7010,7 +6854,6 @@ func (m *awsRestxml_deserializeOpGetObject) HandleDese
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7381,11 +7224,9 @@ func (m *awsRestxml_deserializeOpGetObjectAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectAcl(response, &metadata)
 	}
@@ -7425,7 +7266,6 @@ func (m *awsRestxml_deserializeOpGetObjectAcl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7559,11 +7399,9 @@ func (m *awsRestxml_deserializeOpGetObjectAnnotation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectAnnotation(response, &metadata)
 	}
@@ -7580,7 +7418,6 @@ func (m *awsRestxml_deserializeOpGetObjectAnnotation) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7773,11 +7610,9 @@ func (m *awsRestxml_deserializeOpGetObjectAttributes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectAttributes(response, &metadata)
 	}
@@ -7817,7 +7652,6 @@ func (m *awsRestxml_deserializeOpGetObjectAttributes) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8017,11 +7851,9 @@ func (m *awsRestxml_deserializeOpGetObjectLegalHold) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectLegalHold(response, &metadata)
 	}
@@ -8056,7 +7888,6 @@ func (m *awsRestxml_deserializeOpGetObjectLegalHold) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8169,11 +8000,9 @@ func (m *awsRestxml_deserializeOpGetObjectLockConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectLockConfiguration(response, &metadata)
 	}
@@ -8208,7 +8037,6 @@ func (m *awsRestxml_deserializeOpGetObjectLockConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8321,11 +8149,9 @@ func (m *awsRestxml_deserializeOpGetObjectRetention) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectRetention(response, &metadata)
 	}
@@ -8360,7 +8186,6 @@ func (m *awsRestxml_deserializeOpGetObjectRetention) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8473,11 +8298,9 @@ func (m *awsRestxml_deserializeOpGetObjectTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectTagging(response, &metadata)
 	}
@@ -8517,7 +8340,6 @@ func (m *awsRestxml_deserializeOpGetObjectTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8642,11 +8464,9 @@ func (m *awsRestxml_deserializeOpGetObjectTorrent) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetObjectTorrent(response, &metadata)
 	}
@@ -8663,7 +8483,6 @@ func (m *awsRestxml_deserializeOpGetObjectTorrent) Han
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8754,11 +8573,9 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetPublicAccessBlock(response, &metadata)
 	}
@@ -8793,7 +8610,6 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8906,11 +8722,9 @@ func (m *awsRestxml_deserializeOpHeadBucket) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorHeadBucket(response, &metadata)
 	}
@@ -8922,7 +8736,6 @@ func (m *awsRestxml_deserializeOpHeadBucket) HandleDes
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9033,11 +8846,9 @@ func (m *awsRestxml_deserializeOpHeadObject) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorHeadObject(response, &metadata)
 	}
@@ -9049,7 +8860,6 @@ func (m *awsRestxml_deserializeOpHeadObject) HandleDes
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9415,11 +9225,9 @@ func (m *awsRestxml_deserializeOpListBucketAnalyticsCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketAnalyticsConfigurations(response, &metadata)
 	}
@@ -9454,7 +9262,6 @@ func (m *awsRestxml_deserializeOpListBucketAnalyticsCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9609,11 +9416,9 @@ func (m *awsRestxml_deserializeOpListBucketIntelligent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketIntelligentTieringConfigurations(response, &metadata)
 	}
@@ -9648,7 +9453,6 @@ func (m *awsRestxml_deserializeOpListBucketIntelligent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9803,11 +9607,9 @@ func (m *awsRestxml_deserializeOpListBucketInventoryCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketInventoryConfigurations(response, &metadata)
 	}
@@ -9842,7 +9644,6 @@ func (m *awsRestxml_deserializeOpListBucketInventoryCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9997,11 +9798,9 @@ func (m *awsRestxml_deserializeOpListBucketMetricsConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBucketMetricsConfigurations(response, &metadata)
 	}
@@ -10036,7 +9835,6 @@ func (m *awsRestxml_deserializeOpListBucketMetricsConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10191,11 +9989,9 @@ func (m *awsRestxml_deserializeOpListBuckets) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListBuckets(response, &metadata)
 	}
@@ -10230,7 +10026,6 @@ func (m *awsRestxml_deserializeOpListBuckets) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10375,11 +10170,9 @@ func (m *awsRestxml_deserializeOpListDirectoryBuckets)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDirectoryBuckets(response, &metadata)
 	}
@@ -10414,7 +10207,6 @@ func (m *awsRestxml_deserializeOpListDirectoryBuckets)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10540,11 +10332,9 @@ func (m *awsRestxml_deserializeOpListMultipartUploads)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListMultipartUploads(response, &metadata)
 	}
@@ -10584,7 +10374,6 @@ func (m *awsRestxml_deserializeOpListMultipartUploads)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10852,11 +10641,9 @@ func (m *awsRestxml_deserializeOpListObjectAnnotations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjectAnnotations(response, &metadata)
 	}
@@ -10896,7 +10683,6 @@ func (m *awsRestxml_deserializeOpListObjectAnnotations
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11134,11 +10920,9 @@ func (m *awsRestxml_deserializeOpListObjects) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjects(response, &metadata)
 	}
@@ -11178,7 +10962,6 @@ func (m *awsRestxml_deserializeOpListObjects) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11423,11 +11206,9 @@ func (m *awsRestxml_deserializeOpListObjectsV2) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjectsV2(response, &metadata)
 	}
@@ -11467,7 +11248,6 @@ func (m *awsRestxml_deserializeOpListObjectsV2) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11742,11 +11522,9 @@ func (m *awsRestxml_deserializeOpListObjectVersions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListObjectVersions(response, &metadata)
 	}
@@ -11786,7 +11564,6 @@ func (m *awsRestxml_deserializeOpListObjectVersions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12060,11 +11837,9 @@ func (m *awsRestxml_deserializeOpListParts) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListParts(response, &metadata)
 	}
@@ -12104,7 +11879,6 @@ func (m *awsRestxml_deserializeOpListParts) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12392,11 +12166,9 @@ func (m *awsRestxml_deserializeOpPutBucketAbac) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAbac(response, &metadata)
 	}
@@ -12409,7 +12181,6 @@ func (m *awsRestxml_deserializeOpPutBucketAbac) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12480,11 +12251,9 @@ func (m *awsRestxml_deserializeOpPutBucketAccelerateCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAccelerateConfiguration(response, &metadata)
 	}
@@ -12497,7 +12266,6 @@ func (m *awsRestxml_deserializeOpPutBucketAccelerateCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12568,11 +12336,9 @@ func (m *awsRestxml_deserializeOpPutBucketAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAcl(response, &metadata)
 	}
@@ -12585,7 +12351,6 @@ func (m *awsRestxml_deserializeOpPutBucketAcl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12656,11 +12421,9 @@ func (m *awsRestxml_deserializeOpPutBucketAnalyticsCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketAnalyticsConfiguration(response, &metadata)
 	}
@@ -12673,7 +12436,6 @@ func (m *awsRestxml_deserializeOpPutBucketAnalyticsCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12744,11 +12506,9 @@ func (m *awsRestxml_deserializeOpPutBucketCors) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketCors(response, &metadata)
 	}
@@ -12761,7 +12521,6 @@ func (m *awsRestxml_deserializeOpPutBucketCors) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12832,11 +12591,9 @@ func (m *awsRestxml_deserializeOpPutBucketEncryption) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketEncryption(response, &metadata)
 	}
@@ -12849,7 +12606,6 @@ func (m *awsRestxml_deserializeOpPutBucketEncryption) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12920,11 +12676,9 @@ func (m *awsRestxml_deserializeOpPutBucketIntelligentT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketIntelligentTieringConfiguration(response, &metadata)
 	}
@@ -12937,7 +12691,6 @@ func (m *awsRestxml_deserializeOpPutBucketIntelligentT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13008,11 +12761,9 @@ func (m *awsRestxml_deserializeOpPutBucketInventoryCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketInventoryConfiguration(response, &metadata)
 	}
@@ -13025,7 +12776,6 @@ func (m *awsRestxml_deserializeOpPutBucketInventoryCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13096,11 +12846,9 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketLifecycleConfiguration(response, &metadata)
 	}
@@ -13112,7 +12860,6 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13196,11 +12943,9 @@ func (m *awsRestxml_deserializeOpPutBucketLogging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketLogging(response, &metadata)
 	}
@@ -13213,7 +12958,6 @@ func (m *awsRestxml_deserializeOpPutBucketLogging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13284,11 +13028,9 @@ func (m *awsRestxml_deserializeOpPutBucketMetricsConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketMetricsConfiguration(response, &metadata)
 	}
@@ -13301,7 +13043,6 @@ func (m *awsRestxml_deserializeOpPutBucketMetricsConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13372,11 +13113,9 @@ func (m *awsRestxml_deserializeOpPutBucketNotification
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketNotificationConfiguration(response, &metadata)
 	}
@@ -13389,7 +13128,6 @@ func (m *awsRestxml_deserializeOpPutBucketNotification
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13460,11 +13198,9 @@ func (m *awsRestxml_deserializeOpPutBucketOwnershipCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketOwnershipControls(response, &metadata)
 	}
@@ -13477,7 +13213,6 @@ func (m *awsRestxml_deserializeOpPutBucketOwnershipCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13548,11 +13283,9 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketPolicy(response, &metadata)
 	}
@@ -13565,7 +13298,6 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13636,11 +13368,9 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketReplication(response, &metadata)
 	}
@@ -13653,7 +13383,6 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13724,11 +13453,9 @@ func (m *awsRestxml_deserializeOpPutBucketRequestPayme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketRequestPayment(response, &metadata)
 	}
@@ -13741,7 +13468,6 @@ func (m *awsRestxml_deserializeOpPutBucketRequestPayme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13812,11 +13538,9 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketTagging(response, &metadata)
 	}
@@ -13829,7 +13553,6 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13900,11 +13623,9 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketVersioning(response, &metadata)
 	}
@@ -13917,7 +13638,6 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13988,11 +13708,9 @@ func (m *awsRestxml_deserializeOpPutBucketWebsite) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketWebsite(response, &metadata)
 	}
@@ -14005,7 +13723,6 @@ func (m *awsRestxml_deserializeOpPutBucketWebsite) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14076,11 +13793,9 @@ func (m *awsRestxml_deserializeOpPutObject) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObject(response, &metadata)
 	}
@@ -14092,7 +13807,6 @@ func (m *awsRestxml_deserializeOpPutObject) HandleDese
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14301,11 +14015,9 @@ func (m *awsRestxml_deserializeOpPutObjectAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectAcl(response, &metadata)
 	}
@@ -14317,7 +14029,6 @@ func (m *awsRestxml_deserializeOpPutObjectAcl) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14404,11 +14115,9 @@ func (m *awsRestxml_deserializeOpPutObjectAnnotation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectAnnotation(response, &metadata)
 	}
@@ -14448,7 +14157,6 @@ func (m *awsRestxml_deserializeOpPutObjectAnnotation) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14684,11 +14392,9 @@ func (m *awsRestxml_deserializeOpPutObjectLegalHold) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectLegalHold(response, &metadata)
 	}
@@ -14700,7 +14406,6 @@ func (m *awsRestxml_deserializeOpPutObjectLegalHold) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14784,11 +14489,9 @@ func (m *awsRestxml_deserializeOpPutObjectLockConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectLockConfiguration(response, &metadata)
 	}
@@ -14800,7 +14503,6 @@ func (m *awsRestxml_deserializeOpPutObjectLockConfigur
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14884,11 +14586,9 @@ func (m *awsRestxml_deserializeOpPutObjectRetention) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectRetention(response, &metadata)
 	}
@@ -14900,7 +14600,6 @@ func (m *awsRestxml_deserializeOpPutObjectRetention) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14984,11 +14683,9 @@ func (m *awsRestxml_deserializeOpPutObjectTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutObjectTagging(response, &metadata)
 	}
@@ -15000,7 +14697,6 @@ func (m *awsRestxml_deserializeOpPutObjectTagging) Han
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15084,11 +14780,9 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutPublicAccessBlock(response, &metadata)
 	}
@@ -15101,7 +14795,6 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15172,18 +14865,15 @@ func (m *awsRestxml_deserializeOpRenameObject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorRenameObject(response, &metadata)
 	}
 	output := &RenameObjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15257,11 +14947,9 @@ func (m *awsRestxml_deserializeOpRestoreObject) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorRestoreObject(response, &metadata)
 	}
@@ -15273,7 +14961,6 @@ func (m *awsRestxml_deserializeOpRestoreObject) Handle
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15363,18 +15050,15 @@ func (m *awsRestxml_deserializeOpSelectObjectContent) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorSelectObjectContent(response, &metadata)
 	}
 	output := &SelectObjectContentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15445,11 +15129,9 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateBucketMetadataAnnotationTableConfiguration(response, &metadata)
 	}
@@ -15462,7 +15144,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15533,11 +15214,9 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateBucketMetadataInventoryTableConfiguration(response, &metadata)
 	}
@@ -15550,7 +15229,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15621,11 +15299,9 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateBucketMetadataJournalTableConfiguration(response, &metadata)
 	}
@@ -15638,7 +15314,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15709,11 +15384,9 @@ func (m *awsRestxml_deserializeOpUpdateObjectEncryptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateObjectEncryption(response, &metadata)
 	}
@@ -15725,7 +15398,6 @@ func (m *awsRestxml_deserializeOpUpdateObjectEncryptio
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15818,11 +15490,9 @@ func (m *awsRestxml_deserializeOpUploadPart) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUploadPart(response, &metadata)
 	}
@@ -15834,7 +15504,6 @@ func (m *awsRestxml_deserializeOpUploadPart) HandleDes
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16002,11 +15671,9 @@ func (m *awsRestxml_deserializeOpUploadPartCopy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUploadPartCopy(response, &metadata)
 	}
@@ -16046,7 +15713,6 @@ func (m *awsRestxml_deserializeOpUploadPartCopy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16205,11 +15871,9 @@ func (m *awsRestxml_deserializeOpWriteGetObjectRespons
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
-
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorWriteGetObjectResponse(response, &metadata)
 	}
@@ -16222,7 +15886,6 @@ func (m *awsRestxml_deserializeOpWriteGetObjectRespons
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
