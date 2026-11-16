--- vendor/github.com/aws/aws-sdk-go-v2/service/xray/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/xray/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchGetTraces) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchGetTraces) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -214,10 +210,8 @@ func (m *awsRestjson1_deserializeOpCancelTraceRetrieva
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -229,7 +223,6 @@ func (m *awsRestjson1_deserializeOpCancelTraceRetrieva
 	output := &CancelTraceRetrievalOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -308,10 +301,8 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -351,7 +342,6 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -463,10 +453,8 @@ func (m *awsRestjson1_deserializeOpCreateSamplingRule)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -506,7 +494,6 @@ func (m *awsRestjson1_deserializeOpCreateSamplingRule)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -621,10 +608,8 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -636,7 +621,6 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 	output := &DeleteGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -712,10 +696,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -727,7 +709,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 	output := &DeleteResourcePolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -806,10 +787,8 @@ func (m *awsRestjson1_deserializeOpDeleteSamplingRule)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -849,7 +828,6 @@ func (m *awsRestjson1_deserializeOpDeleteSamplingRule)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -961,10 +939,8 @@ func (m *awsRestjson1_deserializeOpGetEncryptionConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1004,7 +980,6 @@ func (m *awsRestjson1_deserializeOpGetEncryptionConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1116,10 +1091,8 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1159,7 +1132,6 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1271,10 +1243,8 @@ func (m *awsRestjson1_deserializeOpGetGroups) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1314,7 +1284,6 @@ func (m *awsRestjson1_deserializeOpGetGroups) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1435,10 +1404,8 @@ func (m *awsRestjson1_deserializeOpGetIndexingRules) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1478,7 +1445,6 @@ func (m *awsRestjson1_deserializeOpGetIndexingRules) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1599,10 +1565,8 @@ func (m *awsRestjson1_deserializeOpGetInsight) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1642,7 +1606,6 @@ func (m *awsRestjson1_deserializeOpGetInsight) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1754,10 +1717,8 @@ func (m *awsRestjson1_deserializeOpGetInsightEvents) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1797,7 +1758,6 @@ func (m *awsRestjson1_deserializeOpGetInsightEvents) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1918,10 +1878,8 @@ func (m *awsRestjson1_deserializeOpGetInsightImpactGra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1961,7 +1919,6 @@ func (m *awsRestjson1_deserializeOpGetInsightImpactGra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2155,10 +2112,8 @@ func (m *awsRestjson1_deserializeOpGetInsightSummaries
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2198,7 +2153,6 @@ func (m *awsRestjson1_deserializeOpGetInsightSummaries
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2319,10 +2273,8 @@ func (m *awsRestjson1_deserializeOpGetRetrievedTracesG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2362,7 +2314,6 @@ func (m *awsRestjson1_deserializeOpGetRetrievedTracesG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2495,10 +2446,8 @@ func (m *awsRestjson1_deserializeOpGetSamplingRules) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2538,7 +2487,6 @@ func (m *awsRestjson1_deserializeOpGetSamplingRules) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2659,10 +2607,8 @@ func (m *awsRestjson1_deserializeOpGetSamplingStatisti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2702,7 +2648,6 @@ func (m *awsRestjson1_deserializeOpGetSamplingStatisti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2823,10 +2768,8 @@ func (m *awsRestjson1_deserializeOpGetSamplingTargets)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2866,7 +2809,6 @@ func (m *awsRestjson1_deserializeOpGetSamplingTargets)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2999,10 +2941,8 @@ func (m *awsRestjson1_deserializeOpGetServiceGraph) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3042,7 +2982,6 @@ func (m *awsRestjson1_deserializeOpGetServiceGraph) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3204,10 +3143,8 @@ func (m *awsRestjson1_deserializeOpGetTimeSeriesServic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3247,7 +3184,6 @@ func (m *awsRestjson1_deserializeOpGetTimeSeriesServic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3377,10 +3313,8 @@ func (m *awsRestjson1_deserializeOpGetTraceGraph) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3420,7 +3354,6 @@ func (m *awsRestjson1_deserializeOpGetTraceGraph) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3541,10 +3474,8 @@ func (m *awsRestjson1_deserializeOpGetTraceSegmentDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3584,7 +3515,6 @@ func (m *awsRestjson1_deserializeOpGetTraceSegmentDest
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3709,10 +3639,8 @@ func (m *awsRestjson1_deserializeOpGetTraceSummaries) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3752,7 +3680,6 @@ func (m *awsRestjson1_deserializeOpGetTraceSummaries) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3902,10 +3829,8 @@ func (m *awsRestjson1_deserializeOpListResourcePolicie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3945,7 +3870,6 @@ func (m *awsRestjson1_deserializeOpListResourcePolicie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4066,10 +3990,8 @@ func (m *awsRestjson1_deserializeOpListRetrievedTraces
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4109,7 +4031,6 @@ func (m *awsRestjson1_deserializeOpListRetrievedTraces
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4251,10 +4172,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4294,7 +4213,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4418,10 +4336,8 @@ func (m *awsRestjson1_deserializeOpPutEncryptionConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4461,7 +4377,6 @@ func (m *awsRestjson1_deserializeOpPutEncryptionConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4573,10 +4488,8 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4616,7 +4529,6 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4740,10 +4652,8 @@ func (m *awsRestjson1_deserializeOpPutTelemetryRecords
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4755,7 +4665,6 @@ func (m *awsRestjson1_deserializeOpPutTelemetryRecords
 	output := &PutTelemetryRecordsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4831,10 +4740,8 @@ func (m *awsRestjson1_deserializeOpPutTraceSegments) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4874,7 +4781,6 @@ func (m *awsRestjson1_deserializeOpPutTraceSegments) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4986,10 +4892,8 @@ func (m *awsRestjson1_deserializeOpStartTraceRetrieval
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5029,7 +4933,6 @@ func (m *awsRestjson1_deserializeOpStartTraceRetrieval
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5148,10 +5051,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5163,7 +5064,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5245,10 +5145,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5260,7 +5158,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5339,10 +5236,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5382,7 +5277,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5494,10 +5388,8 @@ func (m *awsRestjson1_deserializeOpUpdateIndexingRule)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5537,7 +5429,6 @@ func (m *awsRestjson1_deserializeOpUpdateIndexingRule)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5652,10 +5543,8 @@ func (m *awsRestjson1_deserializeOpUpdateSamplingRule)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5695,7 +5584,6 @@ func (m *awsRestjson1_deserializeOpUpdateSamplingRule)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5807,10 +5695,8 @@ func (m *awsRestjson1_deserializeOpUpdateTraceSegmentD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5850,7 +5736,6 @@ func (m *awsRestjson1_deserializeOpUpdateTraceSegmentD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
