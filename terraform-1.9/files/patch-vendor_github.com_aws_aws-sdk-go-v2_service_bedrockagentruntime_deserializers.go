--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/deserializers.go
@@ -19,7 +19,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -50,10 +49,8 @@ func (m *awsRestjson1_deserializeOpCreateInvocation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -93,7 +90,6 @@ func (m *awsRestjson1_deserializeOpCreateInvocation) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -246,10 +242,8 @@ func (m *awsRestjson1_deserializeOpCreateSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -289,7 +283,6 @@ func (m *awsRestjson1_deserializeOpCreateSession) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -448,10 +441,8 @@ func (m *awsRestjson1_deserializeOpDeleteAgentMemory) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -463,7 +454,6 @@ func (m *awsRestjson1_deserializeOpDeleteAgentMemory) 
 	output := &DeleteAgentMemoryOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -560,10 +550,8 @@ func (m *awsRestjson1_deserializeOpDeleteSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -575,7 +563,6 @@ func (m *awsRestjson1_deserializeOpDeleteSession) Hand
 	output := &DeleteSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -663,10 +650,8 @@ func (m *awsRestjson1_deserializeOpEndSession) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -706,7 +691,6 @@ func (m *awsRestjson1_deserializeOpEndSession) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -852,10 +836,8 @@ func (m *awsRestjson1_deserializeOpGenerateQuery) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -895,7 +877,6 @@ func (m *awsRestjson1_deserializeOpGenerateQuery) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1028,10 +1009,8 @@ func (m *awsRestjson1_deserializeOpGetAgentMemory) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1071,7 +1050,6 @@ func (m *awsRestjson1_deserializeOpGetAgentMemory) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1213,10 +1191,8 @@ func (m *awsRestjson1_deserializeOpGetExecutionFlowSna
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1256,7 +1232,6 @@ func (m *awsRestjson1_deserializeOpGetExecutionFlowSna
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1426,10 +1401,8 @@ func (m *awsRestjson1_deserializeOpGetFlowExecution) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1469,7 +1442,6 @@ func (m *awsRestjson1_deserializeOpGetFlowExecution) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1661,10 +1633,8 @@ func (m *awsRestjson1_deserializeOpGetInvocationStep) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1704,7 +1674,6 @@ func (m *awsRestjson1_deserializeOpGetInvocationStep) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1825,10 +1794,8 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1868,7 +1835,6 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2051,10 +2017,8 @@ func (m *awsRestjson1_deserializeOpInvokeAgent) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2071,7 +2035,6 @@ func (m *awsRestjson1_deserializeOpInvokeAgent) Handle
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2194,10 +2157,8 @@ func (m *awsRestjson1_deserializeOpInvokeFlow) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2214,7 +2175,6 @@ func (m *awsRestjson1_deserializeOpInvokeFlow) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2324,10 +2284,8 @@ func (m *awsRestjson1_deserializeOpInvokeInlineAgent) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2344,7 +2302,6 @@ func (m *awsRestjson1_deserializeOpInvokeInlineAgent) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2459,10 +2416,8 @@ func (m *awsRestjson1_deserializeOpListFlowExecutionEv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2502,7 +2457,6 @@ func (m *awsRestjson1_deserializeOpListFlowExecutionEv
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2632,10 +2586,8 @@ func (m *awsRestjson1_deserializeOpListFlowExecutions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2675,7 +2627,6 @@ func (m *awsRestjson1_deserializeOpListFlowExecutions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2805,10 +2756,8 @@ func (m *awsRestjson1_deserializeOpListInvocations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2848,7 +2797,6 @@ func (m *awsRestjson1_deserializeOpListInvocations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2978,10 +2926,8 @@ func (m *awsRestjson1_deserializeOpListInvocationSteps
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3021,7 +2967,6 @@ func (m *awsRestjson1_deserializeOpListInvocationSteps
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3151,10 +3096,8 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3194,7 +3137,6 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3321,10 +3263,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3364,7 +3304,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3485,10 +3424,8 @@ func (m *awsRestjson1_deserializeOpOptimizePrompt) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3500,7 +3437,6 @@ func (m *awsRestjson1_deserializeOpOptimizePrompt) Han
 	output := &OptimizePromptOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3588,10 +3524,8 @@ func (m *awsRestjson1_deserializeOpPutInvocationStep) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3631,7 +3565,6 @@ func (m *awsRestjson1_deserializeOpPutInvocationStep) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3762,10 +3695,8 @@ func (m *awsRestjson1_deserializeOpRerank) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3805,7 +3736,6 @@ func (m *awsRestjson1_deserializeOpRerank) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3947,10 +3877,8 @@ func (m *awsRestjson1_deserializeOpRetrieve) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3990,7 +3918,6 @@ func (m *awsRestjson1_deserializeOpRetrieve) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4141,10 +4068,8 @@ func (m *awsRestjson1_deserializeOpRetrieveAndGenerate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4184,7 +4109,6 @@ func (m *awsRestjson1_deserializeOpRetrieveAndGenerate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4340,10 +4264,8 @@ func (m *awsRestjson1_deserializeOpRetrieveAndGenerate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4360,7 +4282,6 @@ func (m *awsRestjson1_deserializeOpRetrieveAndGenerate
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4470,10 +4391,8 @@ func (m *awsRestjson1_deserializeOpStartFlowExecution)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4513,7 +4432,6 @@ func (m *awsRestjson1_deserializeOpStartFlowExecution)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4650,10 +4568,8 @@ func (m *awsRestjson1_deserializeOpStopFlowExecution) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4693,7 +4609,6 @@ func (m *awsRestjson1_deserializeOpStopFlowExecution) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4836,10 +4751,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4851,7 +4764,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4939,10 +4851,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4954,7 +4864,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5039,10 +4948,8 @@ func (m *awsRestjson1_deserializeOpUpdateSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5082,7 +4989,6 @@ func (m *awsRestjson1_deserializeOpUpdateSession) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
