--- vendor/github.com/aws/aws-sdk-go-v2/service/batch/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/batch/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 	output := &CancelJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -136,10 +132,8 @@ func (m *awsRestjson1_deserializeOpCreateComputeEnviro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -179,7 +173,6 @@ func (m *awsRestjson1_deserializeOpCreateComputeEnviro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -304,10 +297,8 @@ func (m *awsRestjson1_deserializeOpCreateConsumableRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -347,7 +338,6 @@ func (m *awsRestjson1_deserializeOpCreateConsumableRes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -472,10 +462,8 @@ func (m *awsRestjson1_deserializeOpCreateJobQueue) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -515,7 +503,6 @@ func (m *awsRestjson1_deserializeOpCreateJobQueue) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -640,10 +627,8 @@ func (m *awsRestjson1_deserializeOpCreateSchedulingPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -683,7 +668,6 @@ func (m *awsRestjson1_deserializeOpCreateSchedulingPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -808,10 +792,8 @@ func (m *awsRestjson1_deserializeOpDeleteComputeEnviro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -823,7 +805,6 @@ func (m *awsRestjson1_deserializeOpDeleteComputeEnviro
 	output := &DeleteComputeEnvironmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -899,10 +880,8 @@ func (m *awsRestjson1_deserializeOpDeleteConsumableRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -914,7 +893,6 @@ func (m *awsRestjson1_deserializeOpDeleteConsumableRes
 	output := &DeleteConsumableResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -990,10 +968,8 @@ func (m *awsRestjson1_deserializeOpDeleteJobQueue) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1005,7 +981,6 @@ func (m *awsRestjson1_deserializeOpDeleteJobQueue) Han
 	output := &DeleteJobQueueOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1081,10 +1056,8 @@ func (m *awsRestjson1_deserializeOpDeleteSchedulingPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1096,7 +1069,6 @@ func (m *awsRestjson1_deserializeOpDeleteSchedulingPol
 	output := &DeleteSchedulingPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1172,10 +1144,8 @@ func (m *awsRestjson1_deserializeOpDeregisterJobDefini
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1187,7 +1157,6 @@ func (m *awsRestjson1_deserializeOpDeregisterJobDefini
 	output := &DeregisterJobDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1263,10 +1232,8 @@ func (m *awsRestjson1_deserializeOpDescribeComputeEnvi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1306,7 +1273,6 @@ func (m *awsRestjson1_deserializeOpDescribeComputeEnvi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1427,10 +1393,8 @@ func (m *awsRestjson1_deserializeOpDescribeConsumableR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1470,7 +1434,6 @@ func (m *awsRestjson1_deserializeOpDescribeConsumableR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1661,10 +1624,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobDefiniti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1704,7 +1665,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobDefiniti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1825,10 +1785,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobQueues) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1868,7 +1826,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobQueues) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1989,10 +1946,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobs) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2032,7 +1987,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobs) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2144,10 +2098,8 @@ func (m *awsRestjson1_deserializeOpDescribeSchedulingP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2187,7 +2139,6 @@ func (m *awsRestjson1_deserializeOpDescribeSchedulingP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2299,10 +2250,8 @@ func (m *awsRestjson1_deserializeOpGetJobQueueSnapshot
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2342,7 +2291,6 @@ func (m *awsRestjson1_deserializeOpGetJobQueueSnapshot
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2454,10 +2402,8 @@ func (m *awsRestjson1_deserializeOpListConsumableResou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2497,7 +2443,6 @@ func (m *awsRestjson1_deserializeOpListConsumableResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2618,10 +2563,8 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2661,7 +2604,6 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2782,10 +2724,8 @@ func (m *awsRestjson1_deserializeOpListJobsByConsumabl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2825,7 +2765,6 @@ func (m *awsRestjson1_deserializeOpListJobsByConsumabl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2946,10 +2885,8 @@ func (m *awsRestjson1_deserializeOpListSchedulingPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2989,7 +2926,6 @@ func (m *awsRestjson1_deserializeOpListSchedulingPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3110,10 +3046,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3153,7 +3087,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3265,10 +3198,8 @@ func (m *awsRestjson1_deserializeOpRegisterJobDefiniti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3308,7 +3239,6 @@ func (m *awsRestjson1_deserializeOpRegisterJobDefiniti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3446,10 +3376,8 @@ func (m *awsRestjson1_deserializeOpSubmitJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3489,7 +3417,6 @@ func (m *awsRestjson1_deserializeOpSubmitJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3623,10 +3550,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3638,7 +3563,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3714,10 +3638,8 @@ func (m *awsRestjson1_deserializeOpTerminateJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3729,7 +3651,6 @@ func (m *awsRestjson1_deserializeOpTerminateJob) Handl
 	output := &TerminateJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3805,10 +3726,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3820,7 +3739,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3896,10 +3814,8 @@ func (m *awsRestjson1_deserializeOpUpdateComputeEnviro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3939,7 +3855,6 @@ func (m *awsRestjson1_deserializeOpUpdateComputeEnviro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4064,10 +3979,8 @@ func (m *awsRestjson1_deserializeOpUpdateConsumableRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4107,7 +4020,6 @@ func (m *awsRestjson1_deserializeOpUpdateConsumableRes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4245,10 +4157,8 @@ func (m *awsRestjson1_deserializeOpUpdateJobQueue) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4288,7 +4198,6 @@ func (m *awsRestjson1_deserializeOpUpdateJobQueue) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4413,10 +4322,8 @@ func (m *awsRestjson1_deserializeOpUpdateSchedulingPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4428,7 +4335,6 @@ func (m *awsRestjson1_deserializeOpUpdateSchedulingPol
 	output := &UpdateSchedulingPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
