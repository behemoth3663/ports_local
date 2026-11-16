--- vendor/github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -226,10 +222,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -269,7 +263,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -476,10 +469,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkflowStep)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -519,7 +510,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkflowStep)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -668,10 +658,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkflowStepG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -711,7 +699,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkflowStepG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -891,10 +878,8 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -906,7 +891,6 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 	output := &DeleteTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -991,10 +975,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1034,7 +1016,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1177,10 +1158,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflowStep)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1192,7 +1171,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflowStep)
 	output := &DeleteWorkflowStepOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1277,10 +1255,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflowStepG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1292,7 +1268,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflowStepG
 	output := &DeleteWorkflowStepGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1377,10 +1352,8 @@ func (m *awsRestjson1_deserializeOpGetTemplate) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1420,7 +1393,6 @@ func (m *awsRestjson1_deserializeOpGetTemplate) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1636,10 +1608,8 @@ func (m *awsRestjson1_deserializeOpGetTemplateStep) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1679,7 +1649,6 @@ func (m *awsRestjson1_deserializeOpGetTemplateStep) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1878,10 +1847,8 @@ func (m *awsRestjson1_deserializeOpGetTemplateStepGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1921,7 +1888,6 @@ func (m *awsRestjson1_deserializeOpGetTemplateStepGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2129,10 +2095,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2172,7 +2136,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2499,10 +2462,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStep) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2542,7 +2503,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStep) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2857,10 +2817,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStepGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2900,7 +2858,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStepGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3133,10 +3090,8 @@ func (m *awsRestjson1_deserializeOpListPlugins) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3176,7 +3131,6 @@ func (m *awsRestjson1_deserializeOpListPlugins) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3300,10 +3254,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3343,7 +3295,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3455,10 +3406,8 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3498,7 +3447,6 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3622,10 +3570,8 @@ func (m *awsRestjson1_deserializeOpListTemplateStepGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3665,7 +3611,6 @@ func (m *awsRestjson1_deserializeOpListTemplateStepGro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3792,10 +3737,8 @@ func (m *awsRestjson1_deserializeOpListTemplateSteps) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3835,7 +3778,6 @@ func (m *awsRestjson1_deserializeOpListTemplateSteps) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3965,10 +3907,8 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4008,7 +3948,6 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4138,10 +4077,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowStepGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4181,7 +4118,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowStepGro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4311,10 +4247,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowSteps) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4354,7 +4288,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowSteps) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4481,10 +4414,8 @@ func (m *awsRestjson1_deserializeOpRetryWorkflowStep) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4524,7 +4455,6 @@ func (m *awsRestjson1_deserializeOpRetryWorkflowStep) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4673,10 +4603,8 @@ func (m *awsRestjson1_deserializeOpStartWorkflow) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4716,7 +4644,6 @@ func (m *awsRestjson1_deserializeOpStartWorkflow) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4884,10 +4811,8 @@ func (m *awsRestjson1_deserializeOpStopWorkflow) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4927,7 +4852,6 @@ func (m *awsRestjson1_deserializeOpStopWorkflow) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5095,10 +5019,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5110,7 +5032,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5186,10 +5107,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5201,7 +5120,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5277,10 +5195,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5320,7 +5236,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5459,10 +5374,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5502,7 +5415,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5728,10 +5640,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflowStep)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5771,7 +5681,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflowStep)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5920,10 +5829,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflowStepG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5963,7 +5870,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflowStepG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
