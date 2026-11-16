--- vendor/github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAddNotificationChan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAddNotificationChan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -212,10 +208,8 @@ func (m *awsRestjson1_deserializeOpBatchGetFrameMetric
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -255,7 +249,6 @@ func (m *awsRestjson1_deserializeOpBatchGetFrameMetric
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -418,10 +411,8 @@ func (m *awsRestjson1_deserializeOpConfigureAgent) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -461,7 +452,6 @@ func (m *awsRestjson1_deserializeOpConfigureAgent) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -579,10 +569,8 @@ func (m *awsRestjson1_deserializeOpCreateProfilingGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -622,7 +610,6 @@ func (m *awsRestjson1_deserializeOpCreateProfilingGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -743,10 +730,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfilingGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -758,7 +743,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfilingGrou
 	output := &DeleteProfilingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -843,10 +827,8 @@ func (m *awsRestjson1_deserializeOpDescribeProfilingGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -886,7 +868,6 @@ func (m *awsRestjson1_deserializeOpDescribeProfilingGr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1004,10 +985,8 @@ func (m *awsRestjson1_deserializeOpGetFindingsReportAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1047,7 +1026,6 @@ func (m *awsRestjson1_deserializeOpGetFindingsReportAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1171,10 +1149,8 @@ func (m *awsRestjson1_deserializeOpGetNotificationConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1214,7 +1190,6 @@ func (m *awsRestjson1_deserializeOpGetNotificationConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1332,10 +1307,8 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1375,7 +1348,6 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1503,10 +1475,8 @@ func (m *awsRestjson1_deserializeOpGetProfile) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1528,7 +1498,6 @@ func (m *awsRestjson1_deserializeOpGetProfile) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1649,10 +1618,8 @@ func (m *awsRestjson1_deserializeOpGetRecommendations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1692,7 +1659,6 @@ func (m *awsRestjson1_deserializeOpGetRecommendations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1850,10 +1816,8 @@ func (m *awsRestjson1_deserializeOpListFindingsReports
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1893,7 +1857,6 @@ func (m *awsRestjson1_deserializeOpListFindingsReports
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2020,10 +1983,8 @@ func (m *awsRestjson1_deserializeOpListProfileTimes) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2063,7 +2024,6 @@ func (m *awsRestjson1_deserializeOpListProfileTimes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2190,10 +2150,8 @@ func (m *awsRestjson1_deserializeOpListProfilingGroups
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2233,7 +2191,6 @@ func (m *awsRestjson1_deserializeOpListProfilingGroups
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2359,10 +2316,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2402,7 +2357,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2517,10 +2471,8 @@ func (m *awsRestjson1_deserializeOpPostAgentProfile) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2532,7 +2484,6 @@ func (m *awsRestjson1_deserializeOpPostAgentProfile) H
 	output := &PostAgentProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2614,10 +2565,8 @@ func (m *awsRestjson1_deserializeOpPutPermission) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2657,7 +2606,6 @@ func (m *awsRestjson1_deserializeOpPutPermission) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2791,10 +2739,8 @@ func (m *awsRestjson1_deserializeOpRemoveNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2834,7 +2780,6 @@ func (m *awsRestjson1_deserializeOpRemoveNotificationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2952,10 +2897,8 @@ func (m *awsRestjson1_deserializeOpRemovePermission) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2995,7 +2938,6 @@ func (m *awsRestjson1_deserializeOpRemovePermission) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3129,10 +3071,8 @@ func (m *awsRestjson1_deserializeOpSubmitFeedback) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3144,7 +3084,6 @@ func (m *awsRestjson1_deserializeOpSubmitFeedback) Han
 	output := &SubmitFeedbackOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3226,10 +3165,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3241,7 +3178,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3320,10 +3256,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3335,7 +3269,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3414,10 +3347,8 @@ func (m *awsRestjson1_deserializeOpUpdateProfilingGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3457,7 +3388,6 @@ func (m *awsRestjson1_deserializeOpUpdateProfilingGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
