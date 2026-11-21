--- vendor/github.com/aws/aws-sdk-go-v2/service/s3control/deserializers.go.orig	2026-11-21 14:46:05 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3control/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -50,10 +49,8 @@ func (m *awsRestxml_deserializeOpAssociateAccessGrants
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorAssociateAccessGrantsIdentityCenter(response, &metadata)
@@ -67,7 +64,6 @@ func (m *awsRestxml_deserializeOpAssociateAccessGrants
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -138,10 +134,8 @@ func (m *awsRestxml_deserializeOpCreateAccessGrant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateAccessGrant(response, &metadata)
@@ -177,7 +171,6 @@ func (m *awsRestxml_deserializeOpCreateAccessGrant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -391,10 +384,8 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateAccessGrantsInstance(response, &metadata)
@@ -430,7 +421,6 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -619,10 +609,8 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateAccessGrantsLocation(response, &metadata)
@@ -658,7 +646,6 @@ func (m *awsRestxml_deserializeOpCreateAccessGrantsLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -834,10 +821,8 @@ func (m *awsRestxml_deserializeOpCreateAccessPoint) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateAccessPoint(response, &metadata)
@@ -873,7 +858,6 @@ func (m *awsRestxml_deserializeOpCreateAccessPoint) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1006,10 +990,8 @@ func (m *awsRestxml_deserializeOpCreateAccessPointForO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateAccessPointForObjectLambda(response, &metadata)
@@ -1045,7 +1027,6 @@ func (m *awsRestxml_deserializeOpCreateAccessPointForO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1171,10 +1152,8 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateBucket(response, &metadata)
@@ -1215,7 +1194,6 @@ func (m *awsRestxml_deserializeOpCreateBucket) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1353,10 +1331,8 @@ func (m *awsRestxml_deserializeOpCreateJob) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateJob(response, &metadata)
@@ -1392,7 +1368,6 @@ func (m *awsRestxml_deserializeOpCreateJob) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1524,10 +1499,8 @@ func (m *awsRestxml_deserializeOpCreateMultiRegionAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateMultiRegionAccessPoint(response, &metadata)
@@ -1563,7 +1536,6 @@ func (m *awsRestxml_deserializeOpCreateMultiRegionAcce
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1683,10 +1655,8 @@ func (m *awsRestxml_deserializeOpCreateStorageLensGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateStorageLensGroup(response, &metadata)
@@ -1700,7 +1670,6 @@ func (m *awsRestxml_deserializeOpCreateStorageLensGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1771,10 +1740,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessGrant(response, &metadata)
@@ -1788,7 +1755,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1859,10 +1825,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessGrantsInstance(response, &metadata)
@@ -1876,7 +1840,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1947,10 +1910,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessGrantsInstanceResourcePolicy(response, &metadata)
@@ -1964,7 +1925,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2035,10 +1995,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessGrantsLocation(response, &metadata)
@@ -2052,7 +2010,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessGrantsLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2123,10 +2080,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPoint) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessPoint(response, &metadata)
@@ -2140,7 +2095,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPoint) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2211,10 +2165,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointForO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessPointForObjectLambda(response, &metadata)
@@ -2228,7 +2180,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointForO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2299,10 +2250,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessPointPolicy(response, &metadata)
@@ -2316,7 +2265,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2387,10 +2335,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessPointPolicyForObjectLambda(response, &metadata)
@@ -2404,7 +2350,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2475,10 +2420,8 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointScop
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAccessPointScope(response, &metadata)
@@ -2492,7 +2435,6 @@ func (m *awsRestxml_deserializeOpDeleteAccessPointScop
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2563,10 +2505,8 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucket(response, &metadata)
@@ -2580,7 +2520,6 @@ func (m *awsRestxml_deserializeOpDeleteBucket) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2651,10 +2590,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketLifecycleConfiguration(response, &metadata)
@@ -2668,7 +2605,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketLifecycle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2739,10 +2675,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketPolicy(response, &metadata)
@@ -2756,7 +2690,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2827,10 +2760,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketReplication(response, &metadata)
@@ -2844,7 +2775,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketReplicati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2915,10 +2845,8 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteBucketTagging(response, &metadata)
@@ -2932,7 +2860,6 @@ func (m *awsRestxml_deserializeOpDeleteBucketTagging) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3003,10 +2930,8 @@ func (m *awsRestxml_deserializeOpDeleteJobTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteJobTagging(response, &metadata)
@@ -3014,7 +2939,6 @@ func (m *awsRestxml_deserializeOpDeleteJobTagging) Han
 	output := &DeleteJobTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3094,10 +3018,8 @@ func (m *awsRestxml_deserializeOpDeleteMultiRegionAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteMultiRegionAccessPoint(response, &metadata)
@@ -3133,7 +3055,6 @@ func (m *awsRestxml_deserializeOpDeleteMultiRegionAcce
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3253,10 +3174,8 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeletePublicAccessBlock(response, &metadata)
@@ -3270,7 +3189,6 @@ func (m *awsRestxml_deserializeOpDeletePublicAccessBlo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3341,10 +3259,8 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteStorageLensConfiguration(response, &metadata)
@@ -3358,7 +3274,6 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3429,10 +3344,8 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteStorageLensConfigurationTagging(response, &metadata)
@@ -3440,7 +3353,6 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensConf
 	output := &DeleteStorageLensConfigurationTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3511,10 +3423,8 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteStorageLensGroup(response, &metadata)
@@ -3528,7 +3438,6 @@ func (m *awsRestxml_deserializeOpDeleteStorageLensGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3599,10 +3508,8 @@ func (m *awsRestxml_deserializeOpDescribeJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDescribeJob(response, &metadata)
@@ -3638,7 +3545,6 @@ func (m *awsRestxml_deserializeOpDescribeJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3763,10 +3669,8 @@ func (m *awsRestxml_deserializeOpDescribeMultiRegionAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDescribeMultiRegionAccessPointOperation(response, &metadata)
@@ -3802,7 +3706,6 @@ func (m *awsRestxml_deserializeOpDescribeMultiRegionAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3915,10 +3818,8 @@ func (m *awsRestxml_deserializeOpDissociateAccessGrant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDissociateAccessGrantsIdentityCenter(response, &metadata)
@@ -3932,7 +3833,6 @@ func (m *awsRestxml_deserializeOpDissociateAccessGrant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4003,10 +3903,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrant) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessGrant(response, &metadata)
@@ -4042,7 +3940,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrant) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4256,10 +4153,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessGrantsInstance(response, &metadata)
@@ -4295,7 +4190,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4484,10 +4378,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessGrantsInstanceForPrefix(response, &metadata)
@@ -4523,7 +4415,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4656,10 +4547,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessGrantsInstanceResourcePolicy(response, &metadata)
@@ -4695,7 +4584,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4845,10 +4733,8 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsLocati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessGrantsLocation(response, &metadata)
@@ -4884,7 +4770,6 @@ func (m *awsRestxml_deserializeOpGetAccessGrantsLocati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5060,10 +4945,8 @@ func (m *awsRestxml_deserializeOpGetAccessPoint) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPoint(response, &metadata)
@@ -5099,7 +4982,6 @@ func (m *awsRestxml_deserializeOpGetAccessPoint) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5345,10 +5227,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPointConfigurationForObjectLambda(response, &metadata)
@@ -5384,7 +5264,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5497,10 +5376,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointForObje
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPointForObjectLambda(response, &metadata)
@@ -5536,7 +5413,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointForObje
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5685,10 +5561,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPointPolicy(response, &metadata)
@@ -5724,7 +5598,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5844,10 +5717,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPointPolicyForObjectLambda(response, &metadata)
@@ -5883,7 +5754,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6003,10 +5873,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPointPolicyStatus(response, &metadata)
@@ -6042,7 +5910,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6155,10 +6022,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPointPolicyStatusForObjectLambda(response, &metadata)
@@ -6194,7 +6059,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointPolicyS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6307,10 +6171,8 @@ func (m *awsRestxml_deserializeOpGetAccessPointScope) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccessPointScope(response, &metadata)
@@ -6346,7 +6208,6 @@ func (m *awsRestxml_deserializeOpGetAccessPointScope) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6459,10 +6320,8 @@ func (m *awsRestxml_deserializeOpGetBucket) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucket(response, &metadata)
@@ -6498,7 +6357,6 @@ func (m *awsRestxml_deserializeOpGetBucket) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6651,10 +6509,8 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLifecycleConfiguration(response, &metadata)
@@ -6690,7 +6546,6 @@ func (m *awsRestxml_deserializeOpGetBucketLifecycleCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6803,10 +6658,8 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketPolicy(response, &metadata)
@@ -6842,7 +6695,6 @@ func (m *awsRestxml_deserializeOpGetBucketPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6962,10 +6814,8 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketReplication(response, &metadata)
@@ -7001,7 +6851,6 @@ func (m *awsRestxml_deserializeOpGetBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7114,10 +6963,8 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketTagging(response, &metadata)
@@ -7153,7 +7000,6 @@ func (m *awsRestxml_deserializeOpGetBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7266,10 +7112,8 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetBucketVersioning(response, &metadata)
@@ -7305,7 +7149,6 @@ func (m *awsRestxml_deserializeOpGetBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7438,10 +7281,8 @@ func (m *awsRestxml_deserializeOpGetDataAccess) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetDataAccess(response, &metadata)
@@ -7477,7 +7318,6 @@ func (m *awsRestxml_deserializeOpGetDataAccess) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7622,10 +7462,8 @@ func (m *awsRestxml_deserializeOpGetJobTagging) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetJobTagging(response, &metadata)
@@ -7661,7 +7499,6 @@ func (m *awsRestxml_deserializeOpGetJobTagging) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7783,10 +7620,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetMultiRegionAccessPoint(response, &metadata)
@@ -7822,7 +7657,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7935,10 +7769,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetMultiRegionAccessPointPolicy(response, &metadata)
@@ -7974,7 +7806,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8087,10 +7918,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetMultiRegionAccessPointPolicyStatus(response, &metadata)
@@ -8126,7 +7955,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8239,10 +8067,8 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetMultiRegionAccessPointRoutes(response, &metadata)
@@ -8278,7 +8104,6 @@ func (m *awsRestxml_deserializeOpGetMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8404,10 +8229,8 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetPublicAccessBlock(response, &metadata)
@@ -8443,7 +8266,6 @@ func (m *awsRestxml_deserializeOpGetPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8559,10 +8381,8 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetStorageLensConfiguration(response, &metadata)
@@ -8598,7 +8418,6 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8711,10 +8530,8 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetStorageLensConfigurationTagging(response, &metadata)
@@ -8750,7 +8567,6 @@ func (m *awsRestxml_deserializeOpGetStorageLensConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8863,10 +8679,8 @@ func (m *awsRestxml_deserializeOpGetStorageLensGroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetStorageLensGroup(response, &metadata)
@@ -8902,7 +8716,6 @@ func (m *awsRestxml_deserializeOpGetStorageLensGroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9015,10 +8828,8 @@ func (m *awsRestxml_deserializeOpListAccessGrants) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListAccessGrants(response, &metadata)
@@ -9054,7 +8865,6 @@ func (m *awsRestxml_deserializeOpListAccessGrants) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9180,10 +8990,8 @@ func (m *awsRestxml_deserializeOpListAccessGrantsInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListAccessGrantsInstances(response, &metadata)
@@ -9219,7 +9027,6 @@ func (m *awsRestxml_deserializeOpListAccessGrantsInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9345,10 +9152,8 @@ func (m *awsRestxml_deserializeOpListAccessGrantsLocat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListAccessGrantsLocations(response, &metadata)
@@ -9384,7 +9189,6 @@ func (m *awsRestxml_deserializeOpListAccessGrantsLocat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9510,10 +9314,8 @@ func (m *awsRestxml_deserializeOpListAccessPoints) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListAccessPoints(response, &metadata)
@@ -9549,7 +9351,6 @@ func (m *awsRestxml_deserializeOpListAccessPoints) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9675,10 +9476,8 @@ func (m *awsRestxml_deserializeOpListAccessPointsForDi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListAccessPointsForDirectoryBuckets(response, &metadata)
@@ -9714,7 +9513,6 @@ func (m *awsRestxml_deserializeOpListAccessPointsForDi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9840,10 +9638,8 @@ func (m *awsRestxml_deserializeOpListAccessPointsForOb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListAccessPointsForObjectLambda(response, &metadata)
@@ -9879,7 +9675,6 @@ func (m *awsRestxml_deserializeOpListAccessPointsForOb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10005,10 +9800,8 @@ func (m *awsRestxml_deserializeOpListCallerAccessGrant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListCallerAccessGrants(response, &metadata)
@@ -10044,7 +9837,6 @@ func (m *awsRestxml_deserializeOpListCallerAccessGrant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10170,10 +9962,8 @@ func (m *awsRestxml_deserializeOpListJobs) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListJobs(response, &metadata)
@@ -10209,7 +9999,6 @@ func (m *awsRestxml_deserializeOpListJobs) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10344,10 +10133,8 @@ func (m *awsRestxml_deserializeOpListMultiRegionAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListMultiRegionAccessPoints(response, &metadata)
@@ -10383,7 +10170,6 @@ func (m *awsRestxml_deserializeOpListMultiRegionAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10509,10 +10295,8 @@ func (m *awsRestxml_deserializeOpListRegionalBuckets) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListRegionalBuckets(response, &metadata)
@@ -10548,7 +10332,6 @@ func (m *awsRestxml_deserializeOpListRegionalBuckets) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10674,10 +10457,8 @@ func (m *awsRestxml_deserializeOpListStorageLensConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListStorageLensConfigurations(response, &metadata)
@@ -10713,7 +10494,6 @@ func (m *awsRestxml_deserializeOpListStorageLensConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10839,10 +10619,8 @@ func (m *awsRestxml_deserializeOpListStorageLensGroups
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListStorageLensGroups(response, &metadata)
@@ -10878,7 +10656,6 @@ func (m *awsRestxml_deserializeOpListStorageLensGroups
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11004,10 +10781,8 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTagsForResource(response, &metadata)
@@ -11043,7 +10818,6 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11156,10 +10930,8 @@ func (m *awsRestxml_deserializeOpPutAccessGrantsInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutAccessGrantsInstanceResourcePolicy(response, &metadata)
@@ -11195,7 +10967,6 @@ func (m *awsRestxml_deserializeOpPutAccessGrantsInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11345,10 +11116,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutAccessPointConfigurationForObjectLambda(response, &metadata)
@@ -11362,7 +11131,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11433,10 +11201,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutAccessPointPolicy(response, &metadata)
@@ -11450,7 +11216,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11521,10 +11286,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicyF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutAccessPointPolicyForObjectLambda(response, &metadata)
@@ -11538,7 +11301,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointPolicyF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11609,10 +11371,8 @@ func (m *awsRestxml_deserializeOpPutAccessPointScope) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutAccessPointScope(response, &metadata)
@@ -11626,7 +11386,6 @@ func (m *awsRestxml_deserializeOpPutAccessPointScope) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11697,10 +11456,8 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketLifecycleConfiguration(response, &metadata)
@@ -11714,7 +11471,6 @@ func (m *awsRestxml_deserializeOpPutBucketLifecycleCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11785,10 +11541,8 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketPolicy(response, &metadata)
@@ -11802,7 +11556,6 @@ func (m *awsRestxml_deserializeOpPutBucketPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11873,10 +11626,8 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketReplication(response, &metadata)
@@ -11890,7 +11641,6 @@ func (m *awsRestxml_deserializeOpPutBucketReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11961,10 +11711,8 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketTagging(response, &metadata)
@@ -11978,7 +11726,6 @@ func (m *awsRestxml_deserializeOpPutBucketTagging) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12049,10 +11796,8 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutBucketVersioning(response, &metadata)
@@ -12066,7 +11811,6 @@ func (m *awsRestxml_deserializeOpPutBucketVersioning) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12137,10 +11881,8 @@ func (m *awsRestxml_deserializeOpPutJobTagging) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutJobTagging(response, &metadata)
@@ -12148,7 +11890,6 @@ func (m *awsRestxml_deserializeOpPutJobTagging) Handle
 	output := &PutJobTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12231,10 +11972,8 @@ func (m *awsRestxml_deserializeOpPutMultiRegionAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutMultiRegionAccessPointPolicy(response, &metadata)
@@ -12270,7 +12009,6 @@ func (m *awsRestxml_deserializeOpPutMultiRegionAccessP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12390,10 +12128,8 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutPublicAccessBlock(response, &metadata)
@@ -12407,7 +12143,6 @@ func (m *awsRestxml_deserializeOpPutPublicAccessBlock)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12478,10 +12213,8 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutStorageLensConfiguration(response, &metadata)
@@ -12495,7 +12228,6 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12566,10 +12298,8 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutStorageLensConfigurationTagging(response, &metadata)
@@ -12577,7 +12307,6 @@ func (m *awsRestxml_deserializeOpPutStorageLensConfigu
 	output := &PutStorageLensConfigurationTaggingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12648,10 +12377,8 @@ func (m *awsRestxml_deserializeOpSubmitMultiRegionAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorSubmitMultiRegionAccessPointRoutes(response, &metadata)
@@ -12659,7 +12386,6 @@ func (m *awsRestxml_deserializeOpSubmitMultiRegionAcce
 	output := &SubmitMultiRegionAccessPointRoutesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12730,10 +12456,8 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorTagResource(response, &metadata)
@@ -12741,7 +12465,6 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12812,10 +12535,8 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUntagResource(response, &metadata)
@@ -12823,7 +12544,6 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12894,10 +12614,8 @@ func (m *awsRestxml_deserializeOpUpdateAccessGrantsLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateAccessGrantsLocation(response, &metadata)
@@ -12933,7 +12651,6 @@ func (m *awsRestxml_deserializeOpUpdateAccessGrantsLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13109,10 +12826,8 @@ func (m *awsRestxml_deserializeOpUpdateJobPriority) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateJobPriority(response, &metadata)
@@ -13148,7 +12863,6 @@ func (m *awsRestxml_deserializeOpUpdateJobPriority) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13297,10 +13011,8 @@ func (m *awsRestxml_deserializeOpUpdateJobStatus) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateJobStatus(response, &metadata)
@@ -13336,7 +13048,6 @@ func (m *awsRestxml_deserializeOpUpdateJobStatus) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13497,10 +13208,8 @@ func (m *awsRestxml_deserializeOpUpdateStorageLensGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateStorageLensGroup(response, &metadata)
@@ -13514,7 +13223,6 @@ func (m *awsRestxml_deserializeOpUpdateStorageLensGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
