--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrock/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrock/deserializers.go
@@ -17,7 +17,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteEvaluati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -91,7 +88,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteEvaluati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -220,10 +216,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomModel) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -263,7 +257,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomModel) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -397,10 +390,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomModelDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -440,7 +431,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomModelDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -571,10 +561,8 @@ func (m *awsRestjson1_deserializeOpCreateEvaluationJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -614,7 +602,6 @@ func (m *awsRestjson1_deserializeOpCreateEvaluationJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -745,10 +732,8 @@ func (m *awsRestjson1_deserializeOpCreateFoundationMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -788,7 +773,6 @@ func (m *awsRestjson1_deserializeOpCreateFoundationMod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -916,10 +900,8 @@ func (m *awsRestjson1_deserializeOpCreateGuardrail) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -959,7 +941,6 @@ func (m *awsRestjson1_deserializeOpCreateGuardrail) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1124,10 +1105,8 @@ func (m *awsRestjson1_deserializeOpCreateGuardrailVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1167,7 +1146,6 @@ func (m *awsRestjson1_deserializeOpCreateGuardrailVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1307,10 +1285,8 @@ func (m *awsRestjson1_deserializeOpCreateInferenceProf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1350,7 +1326,6 @@ func (m *awsRestjson1_deserializeOpCreateInferenceProf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1493,10 +1468,8 @@ func (m *awsRestjson1_deserializeOpCreateMarketplaceMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1536,7 +1509,6 @@ func (m *awsRestjson1_deserializeOpCreateMarketplaceMo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1663,10 +1635,8 @@ func (m *awsRestjson1_deserializeOpCreateModelCopyJob)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1706,7 +1676,6 @@ func (m *awsRestjson1_deserializeOpCreateModelCopyJob)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1828,10 +1797,8 @@ func (m *awsRestjson1_deserializeOpCreateModelCustomiz
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1871,7 +1838,6 @@ func (m *awsRestjson1_deserializeOpCreateModelCustomiz
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2005,10 +1971,8 @@ func (m *awsRestjson1_deserializeOpCreateModelImportJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2048,7 +2012,6 @@ func (m *awsRestjson1_deserializeOpCreateModelImportJo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2182,10 +2145,8 @@ func (m *awsRestjson1_deserializeOpCreateModelInvocati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2225,7 +2186,6 @@ func (m *awsRestjson1_deserializeOpCreateModelInvocati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2356,10 +2316,8 @@ func (m *awsRestjson1_deserializeOpCreatePromptRouter)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2399,7 +2357,6 @@ func (m *awsRestjson1_deserializeOpCreatePromptRouter)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2533,10 +2490,8 @@ func (m *awsRestjson1_deserializeOpCreateProvisionedMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2576,7 +2531,6 @@ func (m *awsRestjson1_deserializeOpCreateProvisionedMo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2707,10 +2661,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomModel) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2722,7 +2674,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomModel) 
 	output := &DeleteCustomModelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2810,10 +2761,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomModelDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2825,7 +2774,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomModelDe
 	output := &DeleteCustomModelDeploymentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2913,10 +2861,8 @@ func (m *awsRestjson1_deserializeOpDeleteFoundationMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2928,7 +2874,6 @@ func (m *awsRestjson1_deserializeOpDeleteFoundationMod
 	output := &DeleteFoundationModelAgreementOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3016,10 +2961,8 @@ func (m *awsRestjson1_deserializeOpDeleteGuardrail) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3031,7 +2974,6 @@ func (m *awsRestjson1_deserializeOpDeleteGuardrail) Ha
 	output := &DeleteGuardrailOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3119,10 +3061,8 @@ func (m *awsRestjson1_deserializeOpDeleteImportedModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3134,7 +3074,6 @@ func (m *awsRestjson1_deserializeOpDeleteImportedModel
 	output := &DeleteImportedModelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3222,10 +3161,8 @@ func (m *awsRestjson1_deserializeOpDeleteInferenceProf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3237,7 +3174,6 @@ func (m *awsRestjson1_deserializeOpDeleteInferenceProf
 	output := &DeleteInferenceProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3325,10 +3261,8 @@ func (m *awsRestjson1_deserializeOpDeleteMarketplaceMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3340,7 +3274,6 @@ func (m *awsRestjson1_deserializeOpDeleteMarketplaceMo
 	output := &DeleteMarketplaceModelEndpointOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3425,10 +3358,8 @@ func (m *awsRestjson1_deserializeOpDeleteModelInvocati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3440,7 +3371,6 @@ func (m *awsRestjson1_deserializeOpDeleteModelInvocati
 	output := &DeleteModelInvocationLoggingConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3519,10 +3449,8 @@ func (m *awsRestjson1_deserializeOpDeletePromptRouter)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3534,7 +3462,6 @@ func (m *awsRestjson1_deserializeOpDeletePromptRouter)
 	output := &DeletePromptRouterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3619,10 +3546,8 @@ func (m *awsRestjson1_deserializeOpDeleteProvisionedMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3634,7 +3559,6 @@ func (m *awsRestjson1_deserializeOpDeleteProvisionedMo
 	output := &DeleteProvisionedModelThroughputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3722,10 +3646,8 @@ func (m *awsRestjson1_deserializeOpDeregisterMarketpla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3737,7 +3659,6 @@ func (m *awsRestjson1_deserializeOpDeregisterMarketpla
 	output := &DeregisterMarketplaceModelEndpointOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3825,10 +3746,8 @@ func (m *awsRestjson1_deserializeOpGetCustomModel) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3868,7 +3787,6 @@ func (m *awsRestjson1_deserializeOpGetCustomModel) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4113,10 +4031,8 @@ func (m *awsRestjson1_deserializeOpGetCustomModelDeplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4156,7 +4072,6 @@ func (m *awsRestjson1_deserializeOpGetCustomModelDeplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4352,10 +4267,8 @@ func (m *awsRestjson1_deserializeOpGetEvaluationJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4395,7 +4308,6 @@ func (m *awsRestjson1_deserializeOpGetEvaluationJob) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4629,10 +4541,8 @@ func (m *awsRestjson1_deserializeOpGetFoundationModel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4672,7 +4582,6 @@ func (m *awsRestjson1_deserializeOpGetFoundationModel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4793,10 +4702,8 @@ func (m *awsRestjson1_deserializeOpGetFoundationModelA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4836,7 +4743,6 @@ func (m *awsRestjson1_deserializeOpGetFoundationModelA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4993,10 +4899,8 @@ func (m *awsRestjson1_deserializeOpGetGuardrail) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5036,7 +4940,6 @@ func (m *awsRestjson1_deserializeOpGetGuardrail) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5299,10 +5202,8 @@ func (m *awsRestjson1_deserializeOpGetImportedModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5342,7 +5243,6 @@ func (m *awsRestjson1_deserializeOpGetImportedModel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5544,10 +5444,8 @@ func (m *awsRestjson1_deserializeOpGetInferenceProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5587,7 +5485,6 @@ func (m *awsRestjson1_deserializeOpGetInferenceProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5788,10 +5685,8 @@ func (m *awsRestjson1_deserializeOpGetMarketplaceModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5831,7 +5726,6 @@ func (m *awsRestjson1_deserializeOpGetMarketplaceModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5952,10 +5846,8 @@ func (m *awsRestjson1_deserializeOpGetModelCopyJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5995,7 +5887,6 @@ func (m *awsRestjson1_deserializeOpGetModelCopyJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6210,10 +6101,8 @@ func (m *awsRestjson1_deserializeOpGetModelCustomizati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6253,7 +6142,6 @@ func (m *awsRestjson1_deserializeOpGetModelCustomizati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6552,10 +6440,8 @@ func (m *awsRestjson1_deserializeOpGetModelImportJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6595,7 +6481,6 @@ func (m *awsRestjson1_deserializeOpGetModelImportJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6832,10 +6717,8 @@ func (m *awsRestjson1_deserializeOpGetModelInvocationJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6875,7 +6758,6 @@ func (m *awsRestjson1_deserializeOpGetModelInvocationJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7134,10 +7016,8 @@ func (m *awsRestjson1_deserializeOpGetModelInvocationL
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7177,7 +7057,6 @@ func (m *awsRestjson1_deserializeOpGetModelInvocationL
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7292,10 +7171,8 @@ func (m *awsRestjson1_deserializeOpGetPromptRouter) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7335,7 +7212,6 @@ func (m *awsRestjson1_deserializeOpGetPromptRouter) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7537,10 +7413,8 @@ func (m *awsRestjson1_deserializeOpGetProvisionedModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7580,7 +7454,6 @@ func (m *awsRestjson1_deserializeOpGetProvisionedModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7833,10 +7706,8 @@ func (m *awsRestjson1_deserializeOpGetUseCaseForModelA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7876,7 +7747,6 @@ func (m *awsRestjson1_deserializeOpGetUseCaseForModelA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8002,10 +7872,8 @@ func (m *awsRestjson1_deserializeOpListCustomModelDepl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8045,7 +7913,6 @@ func (m *awsRestjson1_deserializeOpListCustomModelDepl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8172,10 +8039,8 @@ func (m *awsRestjson1_deserializeOpListCustomModels) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8215,7 +8080,6 @@ func (m *awsRestjson1_deserializeOpListCustomModels) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8342,10 +8206,8 @@ func (m *awsRestjson1_deserializeOpListEvaluationJobs)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8385,7 +8247,6 @@ func (m *awsRestjson1_deserializeOpListEvaluationJobs)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8512,10 +8373,8 @@ func (m *awsRestjson1_deserializeOpListFoundationModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8555,7 +8414,6 @@ func (m *awsRestjson1_deserializeOpListFoundationModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8685,10 +8543,8 @@ func (m *awsRestjson1_deserializeOpListFoundationModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8728,7 +8584,6 @@ func (m *awsRestjson1_deserializeOpListFoundationModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8846,10 +8701,8 @@ func (m *awsRestjson1_deserializeOpListGuardrails) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8889,7 +8742,6 @@ func (m *awsRestjson1_deserializeOpListGuardrails) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9019,10 +8871,8 @@ func (m *awsRestjson1_deserializeOpListImportedModels)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9062,7 +8912,6 @@ func (m *awsRestjson1_deserializeOpListImportedModels)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9189,10 +9038,8 @@ func (m *awsRestjson1_deserializeOpListInferenceProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9232,7 +9079,6 @@ func (m *awsRestjson1_deserializeOpListInferenceProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9359,10 +9205,8 @@ func (m *awsRestjson1_deserializeOpListMarketplaceMode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9402,7 +9246,6 @@ func (m *awsRestjson1_deserializeOpListMarketplaceMode
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9532,10 +9375,8 @@ func (m *awsRestjson1_deserializeOpListModelCopyJobs) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9575,7 +9416,6 @@ func (m *awsRestjson1_deserializeOpListModelCopyJobs) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9705,10 +9545,8 @@ func (m *awsRestjson1_deserializeOpListModelCustomizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9748,7 +9586,6 @@ func (m *awsRestjson1_deserializeOpListModelCustomizat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9875,10 +9712,8 @@ func (m *awsRestjson1_deserializeOpListModelImportJobs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9918,7 +9753,6 @@ func (m *awsRestjson1_deserializeOpListModelImportJobs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10045,10 +9879,8 @@ func (m *awsRestjson1_deserializeOpListModelInvocation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10088,7 +9920,6 @@ func (m *awsRestjson1_deserializeOpListModelInvocation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10215,10 +10046,8 @@ func (m *awsRestjson1_deserializeOpListPromptRouters) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10258,7 +10087,6 @@ func (m *awsRestjson1_deserializeOpListPromptRouters) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10385,10 +10213,8 @@ func (m *awsRestjson1_deserializeOpListProvisionedMode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10428,7 +10254,6 @@ func (m *awsRestjson1_deserializeOpListProvisionedMode
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10555,10 +10380,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10598,7 +10421,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10719,10 +10541,8 @@ func (m *awsRestjson1_deserializeOpPutModelInvocationL
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10734,7 +10554,6 @@ func (m *awsRestjson1_deserializeOpPutModelInvocationL
 	output := &PutModelInvocationLoggingConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10816,10 +10635,8 @@ func (m *awsRestjson1_deserializeOpPutUseCaseForModelA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10831,7 +10648,6 @@ func (m *awsRestjson1_deserializeOpPutUseCaseForModelA
 	output := &PutUseCaseForModelAccessOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10913,10 +10729,8 @@ func (m *awsRestjson1_deserializeOpRegisterMarketplace
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10956,7 +10770,6 @@ func (m *awsRestjson1_deserializeOpRegisterMarketplace
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11080,10 +10893,8 @@ func (m *awsRestjson1_deserializeOpStopEvaluationJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11095,7 +10906,6 @@ func (m *awsRestjson1_deserializeOpStopEvaluationJob) 
 	output := &StopEvaluationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11183,10 +10993,8 @@ func (m *awsRestjson1_deserializeOpStopModelCustomizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11198,7 +11006,6 @@ func (m *awsRestjson1_deserializeOpStopModelCustomizat
 	output := &StopModelCustomizationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11286,10 +11093,8 @@ func (m *awsRestjson1_deserializeOpStopModelInvocation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11301,7 +11106,6 @@ func (m *awsRestjson1_deserializeOpStopModelInvocation
 	output := &StopModelInvocationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11389,10 +11193,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11404,7 +11206,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11492,10 +11293,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11507,7 +11306,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11592,10 +11390,8 @@ func (m *awsRestjson1_deserializeOpUpdateGuardrail) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11635,7 +11431,6 @@ func (m *awsRestjson1_deserializeOpUpdateGuardrail) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11797,10 +11592,8 @@ func (m *awsRestjson1_deserializeOpUpdateMarketplaceMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11840,7 +11633,6 @@ func (m *awsRestjson1_deserializeOpUpdateMarketplaceMo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11967,10 +11759,8 @@ func (m *awsRestjson1_deserializeOpUpdateProvisionedMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11982,7 +11772,6 @@ func (m *awsRestjson1_deserializeOpUpdateProvisionedMo
 	output := &UpdateProvisionedModelThroughputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
