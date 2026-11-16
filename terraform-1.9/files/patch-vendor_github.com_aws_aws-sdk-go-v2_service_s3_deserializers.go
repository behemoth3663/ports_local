--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3/deserializers.go
@@ -20,7 +20,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -52,10 +51,8 @@ func (m *awsRestxml_deserializeOpAbortMultipartUpload)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -72,7 +69,6 @@ func (m *awsRestxml_deserializeOpAbortMultipartUpload)
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -147,10 +143,8 @@ func (m *awsRestxml_deserializeOpCompleteMultipartUplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -195,7 +189,6 @@ func (m *awsRestxml_deserializeOpCompleteMultipartUplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -461,10 +454,8 @@ func (m *awsRestxml_deserializeOpCopyObject) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -509,7 +500,6 @@ func (m *awsRestxml_deserializeOpCopyObject) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -674,10 +664,8 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -694,7 +682,6 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -777,10 +764,8 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -798,7 +783,6 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -857,10 +841,8 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -878,7 +860,6 @@ func (m *awsRestxml_deserializeOpCreateBucketMetadataT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -937,10 +918,8 @@ func (m *awsRestxml_deserializeOpCreateMultipartUpload
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -985,7 +964,6 @@ func (m *awsRestxml_deserializeOpCreateMultipartUpload
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1189,10 +1167,8 @@ func (m *awsRestxml_deserializeOpCreateSession) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1237,7 +1213,6 @@ func (m *awsRestxml_deserializeOpCreateSession) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1372,10 +1347,8 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1393,7 +1366,6 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1452,10 +1424,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketAnalytics
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1473,7 +1443,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketAnalytics
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1532,10 +1501,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketCors) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1553,7 +1520,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketCors) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1612,10 +1578,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketEncryptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1633,7 +1597,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketEncryptio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1692,10 +1655,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketIntellige
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1713,7 +1674,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketIntellige
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1772,10 +1732,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketInventory
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1793,7 +1751,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketInventory
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1852,10 +1809,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1873,7 +1828,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1932,10 +1886,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1953,7 +1905,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2012,10 +1963,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2033,7 +1982,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetadataT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2092,10 +2040,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetricsCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2113,7 +2059,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketMetricsCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2172,10 +2117,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketOwnership
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2193,7 +2136,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketOwnership
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2252,10 +2194,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2273,7 +2213,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2332,10 +2271,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2353,7 +2290,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2412,10 +2348,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2433,7 +2367,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2492,10 +2425,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketWebsite) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2513,7 +2444,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketWebsite) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2572,10 +2502,8 @@ func (m *awsRestxml_deserializeOpDeleteObject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2592,7 +2520,6 @@ func (m *awsRestxml_deserializeOpDeleteObject) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2678,10 +2605,8 @@ func (m *awsRestxml_deserializeOpDeleteObjects) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2726,7 +2651,6 @@ func (m *awsRestxml_deserializeOpDeleteObjects) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2845,10 +2769,8 @@ func (m *awsRestxml_deserializeOpDeleteObjectTagging) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2865,7 +2787,6 @@ func (m *awsRestxml_deserializeOpDeleteObjectTagging) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2937,10 +2858,8 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2958,7 +2877,6 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3017,10 +2935,8 @@ func (m *awsRestxml_deserializeOpGetBucketAccelerateCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3065,7 +2981,6 @@ func (m *awsRestxml_deserializeOpGetBucketAccelerateCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3185,10 +3100,8 @@ func (m *awsRestxml_deserializeOpGetBucketAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3228,7 +3141,6 @@ func (m *awsRestxml_deserializeOpGetBucketAcl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3335,10 +3247,8 @@ func (m *awsRestxml_deserializeOpGetBucketAnalyticsCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3378,7 +3288,6 @@ func (m *awsRestxml_deserializeOpGetBucketAnalyticsCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3479,10 +3388,8 @@ func (m *awsRestxml_deserializeOpGetBucketCors) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3522,7 +3429,6 @@ func (m *awsRestxml_deserializeOpGetBucketCors) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3623,10 +3529,8 @@ func (m *awsRestxml_deserializeOpGetBucketEncryption) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3666,7 +3570,6 @@ func (m *awsRestxml_deserializeOpGetBucketEncryption) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3767,10 +3670,8 @@ func (m *awsRestxml_deserializeOpGetBucketIntelligentT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3810,7 +3711,6 @@ func (m *awsRestxml_deserializeOpGetBucketIntelligentT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3911,10 +3811,8 @@ func (m *awsRestxml_deserializeOpGetBucketInventoryCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3954,7 +3852,6 @@ func (m *awsRestxml_deserializeOpGetBucketInventoryCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4055,10 +3952,8 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4103,7 +3998,6 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4216,10 +4110,8 @@ func (m *awsRestxml_deserializeOpGetBucketLocation) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4259,7 +4151,6 @@ func (m *awsRestxml_deserializeOpGetBucketLocation) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4367,10 +4258,8 @@ func (m *awsRestxml_deserializeOpGetBucketLogging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4410,7 +4299,6 @@ func (m *awsRestxml_deserializeOpGetBucketLogging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4511,10 +4399,8 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4554,7 +4440,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4655,10 +4540,8 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataTabl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4698,7 +4581,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetadataTabl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4799,10 +4681,8 @@ func (m *awsRestxml_deserializeOpGetBucketMetricsConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4842,7 +4722,6 @@ func (m *awsRestxml_deserializeOpGetBucketMetricsConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4943,10 +4822,8 @@ func (m *awsRestxml_deserializeOpGetBucketNotification
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4986,7 +4863,6 @@ func (m *awsRestxml_deserializeOpGetBucketNotification
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5105,10 +4981,8 @@ func (m *awsRestxml_deserializeOpGetBucketOwnershipCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5148,7 +5022,6 @@ func (m *awsRestxml_deserializeOpGetBucketOwnershipCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5249,10 +5122,8 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5269,7 +5140,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5349,10 +5219,8 @@ func (m *awsRestxml_deserializeOpGetBucketPolicyStatus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5392,7 +5260,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicyStatus
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5493,10 +5360,8 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5536,7 +5401,6 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5637,10 +5501,8 @@ func (m *awsRestxml_deserializeOpGetBucketRequestPayme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5680,7 +5542,6 @@ func (m *awsRestxml_deserializeOpGetBucketRequestPayme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5788,10 +5649,8 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5831,7 +5690,6 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5932,10 +5790,8 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5975,7 +5831,6 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6096,10 +5951,8 @@ func (m *awsRestxml_deserializeOpGetBucketWebsite) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6139,7 +5992,6 @@ func (m *awsRestxml_deserializeOpGetBucketWebsite) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6258,10 +6110,8 @@ func (m *awsRestxml_deserializeOpGetObject) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6283,7 +6133,6 @@ func (m *awsRestxml_deserializeOpGetObject) HandleDese
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6594,10 +6443,8 @@ func (m *awsRestxml_deserializeOpGetObjectAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6642,7 +6489,6 @@ func (m *awsRestxml_deserializeOpGetObjectAcl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6764,10 +6610,8 @@ func (m *awsRestxml_deserializeOpGetObjectAttributes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6812,7 +6656,6 @@ func (m *awsRestxml_deserializeOpGetObjectAttributes) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7000,10 +6843,8 @@ func (m *awsRestxml_deserializeOpGetObjectLegalHold) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7043,7 +6884,6 @@ func (m *awsRestxml_deserializeOpGetObjectLegalHold) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7144,10 +6984,8 @@ func (m *awsRestxml_deserializeOpGetObjectLockConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7187,7 +7025,6 @@ func (m *awsRestxml_deserializeOpGetObjectLockConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7288,10 +7125,8 @@ func (m *awsRestxml_deserializeOpGetObjectRetention) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7331,7 +7166,6 @@ func (m *awsRestxml_deserializeOpGetObjectRetention) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7432,10 +7266,8 @@ func (m *awsRestxml_deserializeOpGetObjectTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7480,7 +7312,6 @@ func (m *awsRestxml_deserializeOpGetObjectTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7593,10 +7424,8 @@ func (m *awsRestxml_deserializeOpGetObjectTorrent) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7618,7 +7447,6 @@ func (m *awsRestxml_deserializeOpGetObjectTorrent) Han
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7697,10 +7525,8 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7740,7 +7566,6 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7841,10 +7666,8 @@ func (m *awsRestxml_deserializeOpHeadBucket) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7861,7 +7684,6 @@ func (m *awsRestxml_deserializeOpHeadBucket) HandleDes
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7960,10 +7782,8 @@ func (m *awsRestxml_deserializeOpHeadObject) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7980,7 +7800,6 @@ func (m *awsRestxml_deserializeOpHeadObject) HandleDes
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8286,10 +8105,8 @@ func (m *awsRestxml_deserializeOpListBucketAnalyticsCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8329,7 +8146,6 @@ func (m *awsRestxml_deserializeOpListBucketAnalyticsCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8472,10 +8288,8 @@ func (m *awsRestxml_deserializeOpListBucketIntelligent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8515,7 +8329,6 @@ func (m *awsRestxml_deserializeOpListBucketIntelligent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8658,10 +8471,8 @@ func (m *awsRestxml_deserializeOpListBucketInventoryCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8701,7 +8512,6 @@ func (m *awsRestxml_deserializeOpListBucketInventoryCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8844,10 +8654,8 @@ func (m *awsRestxml_deserializeOpListBucketMetricsConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8887,7 +8695,6 @@ func (m *awsRestxml_deserializeOpListBucketMetricsConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9030,10 +8837,8 @@ func (m *awsRestxml_deserializeOpListBuckets) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9073,7 +8878,6 @@ func (m *awsRestxml_deserializeOpListBuckets) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9206,10 +9010,8 @@ func (m *awsRestxml_deserializeOpListDirectoryBuckets)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9249,7 +9051,6 @@ func (m *awsRestxml_deserializeOpListDirectoryBuckets)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9363,10 +9164,8 @@ func (m *awsRestxml_deserializeOpListMultipartUploads)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9411,7 +9210,6 @@ func (m *awsRestxml_deserializeOpListMultipartUploads)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9667,10 +9465,8 @@ func (m *awsRestxml_deserializeOpListObjects) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9715,7 +9511,6 @@ func (m *awsRestxml_deserializeOpListObjects) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9948,10 +9743,8 @@ func (m *awsRestxml_deserializeOpListObjectsV2) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9996,7 +9789,6 @@ func (m *awsRestxml_deserializeOpListObjectsV2) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10259,10 +10051,8 @@ func (m *awsRestxml_deserializeOpListObjectVersions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10307,7 +10097,6 @@ func (m *awsRestxml_deserializeOpListObjectVersions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10569,10 +10358,8 @@ func (m *awsRestxml_deserializeOpListParts) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10617,7 +10404,6 @@ func (m *awsRestxml_deserializeOpListParts) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10893,10 +10679,8 @@ func (m *awsRestxml_deserializeOpPutBucketAccelerateCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10914,7 +10698,6 @@ func (m *awsRestxml_deserializeOpPutBucketAccelerateCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10973,10 +10756,8 @@ func (m *awsRestxml_deserializeOpPutBucketAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10994,7 +10775,6 @@ func (m *awsRestxml_deserializeOpPutBucketAcl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11053,10 +10833,8 @@ func (m *awsRestxml_deserializeOpPutBucketAnalyticsCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11074,7 +10852,6 @@ func (m *awsRestxml_deserializeOpPutBucketAnalyticsCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11133,10 +10910,8 @@ func (m *awsRestxml_deserializeOpPutBucketCors) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11154,7 +10929,6 @@ func (m *awsRestxml_deserializeOpPutBucketCors) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11213,10 +10987,8 @@ func (m *awsRestxml_deserializeOpPutBucketEncryption) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11234,7 +11006,6 @@ func (m *awsRestxml_deserializeOpPutBucketEncryption) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11293,10 +11064,8 @@ func (m *awsRestxml_deserializeOpPutBucketIntelligentT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11314,7 +11083,6 @@ func (m *awsRestxml_deserializeOpPutBucketIntelligentT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11373,10 +11141,8 @@ func (m *awsRestxml_deserializeOpPutBucketInventoryCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11394,7 +11160,6 @@ func (m *awsRestxml_deserializeOpPutBucketInventoryCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11453,10 +11218,8 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11473,7 +11236,6 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11545,10 +11307,8 @@ func (m *awsRestxml_deserializeOpPutBucketLogging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11566,7 +11326,6 @@ func (m *awsRestxml_deserializeOpPutBucketLogging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11625,10 +11384,8 @@ func (m *awsRestxml_deserializeOpPutBucketMetricsConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11646,7 +11403,6 @@ func (m *awsRestxml_deserializeOpPutBucketMetricsConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11705,10 +11461,8 @@ func (m *awsRestxml_deserializeOpPutBucketNotification
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11726,7 +11480,6 @@ func (m *awsRestxml_deserializeOpPutBucketNotification
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11785,10 +11538,8 @@ func (m *awsRestxml_deserializeOpPutBucketOwnershipCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11806,7 +11557,6 @@ func (m *awsRestxml_deserializeOpPutBucketOwnershipCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11865,10 +11615,8 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11886,7 +11634,6 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11945,10 +11692,8 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11966,7 +11711,6 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12025,10 +11769,8 @@ func (m *awsRestxml_deserializeOpPutBucketRequestPayme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12046,7 +11788,6 @@ func (m *awsRestxml_deserializeOpPutBucketRequestPayme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12105,10 +11846,8 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12126,7 +11865,6 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12185,10 +11923,8 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12206,7 +11942,6 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12265,10 +12000,8 @@ func (m *awsRestxml_deserializeOpPutBucketWebsite) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12286,7 +12019,6 @@ func (m *awsRestxml_deserializeOpPutBucketWebsite) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12345,10 +12077,8 @@ func (m *awsRestxml_deserializeOpPutObject) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12365,7 +12095,6 @@ func (m *awsRestxml_deserializeOpPutObject) HandleDese
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12537,10 +12266,8 @@ func (m *awsRestxml_deserializeOpPutObjectAcl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12557,7 +12284,6 @@ func (m *awsRestxml_deserializeOpPutObjectAcl) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12632,10 +12358,8 @@ func (m *awsRestxml_deserializeOpPutObjectLegalHold) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12652,7 +12376,6 @@ func (m *awsRestxml_deserializeOpPutObjectLegalHold) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12724,10 +12447,8 @@ func (m *awsRestxml_deserializeOpPutObjectLockConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12744,7 +12465,6 @@ func (m *awsRestxml_deserializeOpPutObjectLockConfigur
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12816,10 +12536,8 @@ func (m *awsRestxml_deserializeOpPutObjectRetention) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12836,7 +12554,6 @@ func (m *awsRestxml_deserializeOpPutObjectRetention) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12908,10 +12625,8 @@ func (m *awsRestxml_deserializeOpPutObjectTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12928,7 +12643,6 @@ func (m *awsRestxml_deserializeOpPutObjectTagging) Han
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13000,10 +12714,8 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13021,7 +12733,6 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13080,10 +12791,8 @@ func (m *awsRestxml_deserializeOpRenameObject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13095,7 +12804,6 @@ func (m *awsRestxml_deserializeOpRenameObject) HandleD
 	output := &RenameObjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13157,10 +12865,8 @@ func (m *awsRestxml_deserializeOpRestoreObject) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13177,7 +12883,6 @@ func (m *awsRestxml_deserializeOpRestoreObject) Handle
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13257,10 +12962,8 @@ func (m *awsRestxml_deserializeOpSelectObjectContent) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13272,7 +12975,6 @@ func (m *awsRestxml_deserializeOpSelectObjectContent) 
 	output := &SelectObjectContentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13331,10 +13033,8 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13352,7 +13052,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13411,10 +13110,8 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13432,7 +13129,6 @@ func (m *awsRestxml_deserializeOpUpdateBucketMetadataJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13491,10 +13187,8 @@ func (m *awsRestxml_deserializeOpUploadPart) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13511,7 +13205,6 @@ func (m *awsRestxml_deserializeOpUploadPart) HandleDes
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13642,10 +13335,8 @@ func (m *awsRestxml_deserializeOpUploadPartCopy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13690,7 +13381,6 @@ func (m *awsRestxml_deserializeOpUploadPartCopy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13837,10 +13527,8 @@ func (m *awsRestxml_deserializeOpWriteGetObjectRespons
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13858,7 +13546,6 @@ func (m *awsRestxml_deserializeOpWriteGetObjectRespons
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
