--- vendor/github.com/aws/aws-sdk-go-v2/service/s3control/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3control/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -48,10 +47,8 @@ func (m *awsRestxml_deserializeOpAssociateAccessGrants
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -69,7 +66,6 @@ func (m *awsRestxml_deserializeOpAssociateAccessGrants
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -128,10 +124,8 @@ func (m *awsRestxml_deserializeOpCreateAccessGrant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -171,7 +165,6 @@ func (m *awsRestxml_deserializeOpCreateAccessGrant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -373,10 +366,8 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -416,7 +407,6 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -593,10 +583,8 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -636,7 +624,6 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -800,10 +787,8 @@ func (m *awsRestxml_deserializeOpCreateAccessPoint) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -843,7 +828,6 @@ func (m *awsRestxml_deserializeOpCreateAccessPoint) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -964,10 +948,8 @@ func (m *awsRestxml_deserializeOpCreateAccessPointForO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1007,7 +989,6 @@ func (m *awsRestxml_deserializeOpCreateAccessPointForO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1121,10 +1102,8 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1169,7 +1148,6 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1295,10 +1273,8 @@ func (m *awsRestxml_deserializeOpCreateJob) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1338,7 +1314,6 @@ func (m *awsRestxml_deserializeOpCreateJob) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1458,10 +1433,8 @@ func (m *awsRestxml_deserializeOpCreateMultiRegionAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1501,7 +1474,6 @@ func (m *awsRestxml_deserializeOpCreateMultiRegionAcce
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1609,10 +1581,8 @@ func (m *awsRestxml_deserializeOpCreateStorageLensGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1630,7 +1600,6 @@ func (m *awsRestxml_deserializeOpCreateStorageLensGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1689,10 +1658,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1710,7 +1677,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1769,10 +1735,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1790,7 +1754,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1849,10 +1812,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1870,7 +1831,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1929,10 +1889,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1950,7 +1908,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2009,10 +1966,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPoint) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2030,7 +1985,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPoint) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2089,10 +2043,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointForO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2110,7 +2062,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointForO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2169,10 +2120,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2190,7 +2139,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2249,10 +2197,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2270,7 +2216,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2329,10 +2274,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointScop
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2350,7 +2293,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointScop
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2409,10 +2351,8 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2430,7 +2370,6 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2489,10 +2428,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2510,7 +2447,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2569,10 +2505,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2590,7 +2524,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2649,10 +2582,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2670,7 +2601,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2729,10 +2659,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2750,7 +2678,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2809,10 +2736,8 @@ func (m *awsRestxml_deserializeOpDeleteJobTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2824,7 +2749,6 @@ func (m *awsRestxml_deserializeOpDeleteJobTagging) Han
 	output := &DeleteJobTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2892,10 +2816,8 @@ func (m *awsRestxml_deserializeOpDeleteMultiRegionAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2935,7 +2857,6 @@ func (m *awsRestxml_deserializeOpDeleteMultiRegionAcce
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3043,10 +2964,8 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3064,7 +2983,6 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3123,10 +3041,8 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3144,7 +3060,6 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3203,10 +3118,8 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3218,7 +3131,6 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 	output := &DeleteStorageLensConfigurationTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3277,10 +3189,8 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3298,7 +3208,6 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3357,10 +3266,8 @@ func (m *awsRestxml_deserializeOpDescribeJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3400,7 +3307,6 @@ func (m *awsRestxml_deserializeOpDescribeJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3513,10 +3419,8 @@ func (m *awsRestxml_deserializeOpDescribeMultiRegionAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3556,7 +3460,6 @@ func (m *awsRestxml_deserializeOpDescribeMultiRegionAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3657,10 +3560,8 @@ func (m *awsRestxml_deserializeOpDissociateAccessGrant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3678,7 +3579,6 @@ func (m *awsRestxml_deserializeOpDissociateAccessGrant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3737,10 +3637,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrant) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3780,7 +3678,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrant) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3982,10 +3879,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4025,7 +3920,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4202,10 +4096,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4245,7 +4137,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4366,10 +4257,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4409,7 +4298,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4547,10 +4435,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsLocati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4590,7 +4476,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsLocati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4754,10 +4639,8 @@ func (m *awsRestxml_deserializeOpGetAccessPoint) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4797,7 +4680,6 @@ func (m *awsRestxml_deserializeOpGetAccessPoint) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5031,10 +4913,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5074,7 +4954,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5175,10 +5054,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointForObje
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5218,7 +5095,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointForObje
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5355,10 +5231,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5398,7 +5272,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5506,10 +5379,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5549,7 +5420,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5657,10 +5527,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5700,7 +5568,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5801,10 +5668,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5844,7 +5709,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5945,10 +5809,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointScope) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5988,7 +5850,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointScope) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6089,10 +5950,8 @@ func (m *awsRestxml_deserializeOpGetBucket) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6132,7 +5991,6 @@ func (m *awsRestxml_deserializeOpGetBucket) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6273,10 +6131,8 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6316,7 +6172,6 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6417,10 +6272,8 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6460,7 +6313,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6568,10 +6420,8 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6611,7 +6461,6 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6712,10 +6561,8 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6755,7 +6602,6 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6856,10 +6702,8 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6899,7 +6743,6 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7020,10 +6863,8 @@ func (m *awsRestxml_deserializeOpGetDataAccess) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7063,7 +6904,6 @@ func (m *awsRestxml_deserializeOpGetDataAccess) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7183,10 +7023,8 @@ func (m *awsRestxml_deserializeOpGetJobTagging) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7226,7 +7064,6 @@ func (m *awsRestxml_deserializeOpGetJobTagging) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7336,10 +7173,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7379,7 +7214,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7480,10 +7314,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7523,7 +7355,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7624,10 +7455,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7667,7 +7496,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7768,10 +7596,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7811,7 +7637,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7925,10 +7750,8 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7968,7 +7791,6 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8072,10 +7894,8 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8115,7 +7935,6 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8216,10 +8035,8 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8259,7 +8076,6 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8360,10 +8176,8 @@ func (m *awsRestxml_deserializeOpGetStorageLensGroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8403,7 +8217,6 @@ func (m *awsRestxml_deserializeOpGetStorageLensGroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8504,10 +8317,8 @@ func (m *awsRestxml_deserializeOpListAccessGrants) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8547,7 +8358,6 @@ func (m *awsRestxml_deserializeOpListAccessGrants) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8661,10 +8471,8 @@ func (m *awsRestxml_deserializeOpListAccessGrantsInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8704,7 +8512,6 @@ func (m *awsRestxml_deserializeOpListAccessGrantsInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8818,10 +8625,8 @@ func (m *awsRestxml_deserializeOpListAccessGrantsLocat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8861,7 +8666,6 @@ func (m *awsRestxml_deserializeOpListAccessGrantsLocat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8975,10 +8779,8 @@ func (m *awsRestxml_deserializeOpListAccessPoints) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9018,7 +8820,6 @@ func (m *awsRestxml_deserializeOpListAccessPoints) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9132,10 +8933,8 @@ func (m *awsRestxml_deserializeOpListAccessPointsForDi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9175,7 +8974,6 @@ func (m *awsRestxml_deserializeOpListAccessPointsForDi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9289,10 +9087,8 @@ func (m *awsRestxml_deserializeOpListAccessPointsForOb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9332,7 +9128,6 @@ func (m *awsRestxml_deserializeOpListAccessPointsForOb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9446,10 +9241,8 @@ func (m *awsRestxml_deserializeOpListCallerAccessGrant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9489,7 +9282,6 @@ func (m *awsRestxml_deserializeOpListCallerAccessGrant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9603,10 +9395,8 @@ func (m *awsRestxml_deserializeOpListJobs) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9646,7 +9436,6 @@ func (m *awsRestxml_deserializeOpListJobs) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9769,10 +9558,8 @@ func (m *awsRestxml_deserializeOpListMultiRegionAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9812,7 +9599,6 @@ func (m *awsRestxml_deserializeOpListMultiRegionAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9926,10 +9712,8 @@ func (m *awsRestxml_deserializeOpListRegionalBuckets) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9969,7 +9753,6 @@ func (m *awsRestxml_deserializeOpListRegionalBuckets) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10083,10 +9866,8 @@ func (m *awsRestxml_deserializeOpListStorageLensConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10126,7 +9907,6 @@ func (m *awsRestxml_deserializeOpListStorageLensConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10240,10 +10020,8 @@ func (m *awsRestxml_deserializeOpListStorageLensGroups
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10283,7 +10061,6 @@ func (m *awsRestxml_deserializeOpListStorageLensGroups
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10397,10 +10174,8 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10440,7 +10215,6 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10541,10 +10315,8 @@ func (m *awsRestxml_deserializeOpPutAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10584,7 +10356,6 @@ func (m *awsRestxml_deserializeOpPutAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10722,10 +10493,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10743,7 +10512,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10802,10 +10570,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10823,7 +10589,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10882,10 +10647,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicyF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10903,7 +10666,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicyF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10962,10 +10724,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointScope) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10983,7 +10743,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointScope) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11042,10 +10801,8 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11063,7 +10820,6 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11122,10 +10878,8 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11143,7 +10897,6 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11202,10 +10955,8 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11223,7 +10974,6 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11282,10 +11032,8 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11303,7 +11051,6 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11362,10 +11109,8 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11383,7 +11128,6 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11442,10 +11186,8 @@ func (m *awsRestxml_deserializeOpPutJobTagging) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11457,7 +11199,6 @@ func (m *awsRestxml_deserializeOpPutJobTagging) Handle
 	output := &PutJobTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11528,10 +11269,8 @@ func (m *awsRestxml_deserializeOpPutMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11571,7 +11310,6 @@ func (m *awsRestxml_deserializeOpPutMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11679,10 +11417,8 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11700,7 +11436,6 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11759,10 +11494,8 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11780,7 +11513,6 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11839,10 +11571,8 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11854,7 +11584,6 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 	output := &PutStorageLensConfigurationTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11913,10 +11642,8 @@ func (m *awsRestxml_deserializeOpSubmitMultiRegionAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11928,7 +11655,6 @@ func (m *awsRestxml_deserializeOpSubmitMultiRegionAcce
 	output := &SubmitMultiRegionAccessPointRoutesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11987,10 +11713,8 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12002,7 +11726,6 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12061,10 +11784,8 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12076,7 +11797,6 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12135,10 +11855,8 @@ func (m *awsRestxml_deserializeOpUpdateAccessGrantsLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12178,7 +11896,6 @@ func (m *awsRestxml_deserializeOpUpdateAccessGrantsLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12342,10 +12059,8 @@ func (m *awsRestxml_deserializeOpUpdateJobPriority) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12385,7 +12100,6 @@ func (m *awsRestxml_deserializeOpUpdateJobPriority) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12522,10 +12236,8 @@ func (m *awsRestxml_deserializeOpUpdateJobStatus) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12565,7 +12277,6 @@ func (m *awsRestxml_deserializeOpUpdateJobStatus) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12714,10 +12425,8 @@ func (m *awsRestxml_deserializeOpUpdateStorageLensGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12735,7 +12444,6 @@ func (m *awsRestxml_deserializeOpUpdateStorageLensGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
