--- vendor/github.com/aws/aws-sdk-go-v2/service/imagebuilder/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/imagebuilder/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelImageCreation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCancelImageCreation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -240,10 +236,8 @@ func (m *awsRestjson1_deserializeOpCancelLifecycleExec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -283,7 +277,6 @@ func (m *awsRestjson1_deserializeOpCancelLifecycleExec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -417,10 +410,8 @@ func (m *awsRestjson1_deserializeOpCreateComponent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -460,7 +451,6 @@ func (m *awsRestjson1_deserializeOpCreateComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -621,10 +611,8 @@ func (m *awsRestjson1_deserializeOpCreateContainerReci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -664,7 +652,6 @@ func (m *awsRestjson1_deserializeOpCreateContainerReci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -825,10 +812,8 @@ func (m *awsRestjson1_deserializeOpCreateDistributionC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -868,7 +853,6 @@ func (m *awsRestjson1_deserializeOpCreateDistributionC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1029,10 +1013,8 @@ func (m *awsRestjson1_deserializeOpCreateImage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1072,7 +1054,6 @@ func (m *awsRestjson1_deserializeOpCreateImage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1227,10 +1208,8 @@ func (m *awsRestjson1_deserializeOpCreateImagePipeline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1270,7 +1249,6 @@ func (m *awsRestjson1_deserializeOpCreateImagePipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1428,10 +1406,8 @@ func (m *awsRestjson1_deserializeOpCreateImageRecipe) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1471,7 +1447,6 @@ func (m *awsRestjson1_deserializeOpCreateImageRecipe) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1632,10 +1607,8 @@ func (m *awsRestjson1_deserializeOpCreateInfrastructur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1675,7 +1648,6 @@ func (m *awsRestjson1_deserializeOpCreateInfrastructur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1833,10 +1805,8 @@ func (m *awsRestjson1_deserializeOpCreateLifecyclePoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1876,7 +1846,6 @@ func (m *awsRestjson1_deserializeOpCreateLifecyclePoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2025,10 +1994,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2068,7 +2035,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2220,10 +2186,8 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2263,7 +2227,6 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2403,10 +2366,8 @@ func (m *awsRestjson1_deserializeOpDeleteContainerReci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2446,7 +2407,6 @@ func (m *awsRestjson1_deserializeOpDeleteContainerReci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2586,10 +2546,8 @@ func (m *awsRestjson1_deserializeOpDeleteDistributionC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2629,7 +2587,6 @@ func (m *awsRestjson1_deserializeOpDeleteDistributionC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2769,10 +2726,8 @@ func (m *awsRestjson1_deserializeOpDeleteImage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2812,7 +2767,6 @@ func (m *awsRestjson1_deserializeOpDeleteImage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2952,10 +2906,8 @@ func (m *awsRestjson1_deserializeOpDeleteImagePipeline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2995,7 +2947,6 @@ func (m *awsRestjson1_deserializeOpDeleteImagePipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3135,10 +3086,8 @@ func (m *awsRestjson1_deserializeOpDeleteImageRecipe) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3178,7 +3127,6 @@ func (m *awsRestjson1_deserializeOpDeleteImageRecipe) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3318,10 +3266,8 @@ func (m *awsRestjson1_deserializeOpDeleteInfrastructur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3361,7 +3307,6 @@ func (m *awsRestjson1_deserializeOpDeleteInfrastructur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3501,10 +3446,8 @@ func (m *awsRestjson1_deserializeOpDeleteLifecyclePoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3544,7 +3487,6 @@ func (m *awsRestjson1_deserializeOpDeleteLifecyclePoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3675,10 +3617,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3718,7 +3658,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3849,10 +3788,8 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3892,7 +3829,6 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4025,10 +3961,8 @@ func (m *awsRestjson1_deserializeOpGetComponentPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4068,7 +4002,6 @@ func (m *awsRestjson1_deserializeOpGetComponentPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4205,10 +4138,8 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipe)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4248,7 +4179,6 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipe)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4381,10 +4311,8 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipeP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4424,7 +4352,6 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipeP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4561,10 +4488,8 @@ func (m *awsRestjson1_deserializeOpGetDistributionConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4604,7 +4529,6 @@ func (m *awsRestjson1_deserializeOpGetDistributionConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4737,10 +4661,8 @@ func (m *awsRestjson1_deserializeOpGetImage) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4780,7 +4702,6 @@ func (m *awsRestjson1_deserializeOpGetImage) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4913,10 +4834,8 @@ func (m *awsRestjson1_deserializeOpGetImagePipeline) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4956,7 +4875,6 @@ func (m *awsRestjson1_deserializeOpGetImagePipeline) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5089,10 +5007,8 @@ func (m *awsRestjson1_deserializeOpGetImagePolicy) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5132,7 +5048,6 @@ func (m *awsRestjson1_deserializeOpGetImagePolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5269,10 +5184,8 @@ func (m *awsRestjson1_deserializeOpGetImageRecipe) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5312,7 +5225,6 @@ func (m *awsRestjson1_deserializeOpGetImageRecipe) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5445,10 +5357,8 @@ func (m *awsRestjson1_deserializeOpGetImageRecipePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5488,7 +5398,6 @@ func (m *awsRestjson1_deserializeOpGetImageRecipePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5625,10 +5534,8 @@ func (m *awsRestjson1_deserializeOpGetInfrastructureCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5668,7 +5575,6 @@ func (m *awsRestjson1_deserializeOpGetInfrastructureCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5801,10 +5707,8 @@ func (m *awsRestjson1_deserializeOpGetLifecycleExecuti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5844,7 +5748,6 @@ func (m *awsRestjson1_deserializeOpGetLifecycleExecuti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5968,10 +5871,8 @@ func (m *awsRestjson1_deserializeOpGetLifecyclePolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6011,7 +5912,6 @@ func (m *awsRestjson1_deserializeOpGetLifecyclePolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6135,10 +6035,8 @@ func (m *awsRestjson1_deserializeOpGetMarketplaceResou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6178,7 +6076,6 @@ func (m *awsRestjson1_deserializeOpGetMarketplaceResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6324,10 +6221,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6367,7 +6262,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6491,10 +6385,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowExecutio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6534,7 +6426,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowExecutio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6795,10 +6686,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStepExec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6838,7 +6727,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStepExec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7114,10 +7002,8 @@ func (m *awsRestjson1_deserializeOpImportComponent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7157,7 +7043,6 @@ func (m *awsRestjson1_deserializeOpImportComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7315,10 +7200,8 @@ func (m *awsRestjson1_deserializeOpImportDiskImage) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7358,7 +7241,6 @@ func (m *awsRestjson1_deserializeOpImportDiskImage) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7486,10 +7368,8 @@ func (m *awsRestjson1_deserializeOpImportVmImage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7529,7 +7409,6 @@ func (m *awsRestjson1_deserializeOpImportVmImage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7666,10 +7545,8 @@ func (m *awsRestjson1_deserializeOpListComponentBuildV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7709,7 +7586,6 @@ func (m *awsRestjson1_deserializeOpListComponentBuildV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7854,10 +7730,8 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7897,7 +7771,6 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8042,10 +7915,8 @@ func (m *awsRestjson1_deserializeOpListContainerRecipe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8085,7 +7956,6 @@ func (m *awsRestjson1_deserializeOpListContainerRecipe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8230,10 +8100,8 @@ func (m *awsRestjson1_deserializeOpListDistributionCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8273,7 +8141,6 @@ func (m *awsRestjson1_deserializeOpListDistributionCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8418,10 +8285,8 @@ func (m *awsRestjson1_deserializeOpListImageBuildVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8461,7 +8326,6 @@ func (m *awsRestjson1_deserializeOpListImageBuildVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8606,10 +8470,8 @@ func (m *awsRestjson1_deserializeOpListImagePackages) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8649,7 +8511,6 @@ func (m *awsRestjson1_deserializeOpListImagePackages) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8797,10 +8658,8 @@ func (m *awsRestjson1_deserializeOpListImagePipelineIm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8840,7 +8699,6 @@ func (m *awsRestjson1_deserializeOpListImagePipelineIm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8988,10 +8846,8 @@ func (m *awsRestjson1_deserializeOpListImagePipelines)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9031,7 +8887,6 @@ func (m *awsRestjson1_deserializeOpListImagePipelines)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9176,10 +9031,8 @@ func (m *awsRestjson1_deserializeOpListImageRecipes) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9219,7 +9072,6 @@ func (m *awsRestjson1_deserializeOpListImageRecipes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9364,10 +9216,8 @@ func (m *awsRestjson1_deserializeOpListImages) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9407,7 +9257,6 @@ func (m *awsRestjson1_deserializeOpListImages) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9552,10 +9401,8 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9595,7 +9442,6 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9749,10 +9595,8 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9792,7 +9636,6 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9937,10 +9780,8 @@ func (m *awsRestjson1_deserializeOpListInfrastructureC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9980,7 +9821,6 @@ func (m *awsRestjson1_deserializeOpListInfrastructureC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10125,10 +9965,8 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10168,7 +10006,6 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10318,10 +10155,8 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10361,7 +10196,6 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10497,10 +10331,8 @@ func (m *awsRestjson1_deserializeOpListLifecyclePolici
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10540,7 +10372,6 @@ func (m *awsRestjson1_deserializeOpListLifecyclePolici
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10676,10 +10507,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10719,7 +10548,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10834,10 +10662,8 @@ func (m *awsRestjson1_deserializeOpListWaitingWorkflow
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10877,7 +10703,6 @@ func (m *awsRestjson1_deserializeOpListWaitingWorkflow
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11013,10 +10838,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowBuildVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11056,7 +10879,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowBuildVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11192,10 +11014,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowExecuti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11235,7 +11055,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowExecuti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11398,10 +11217,8 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11441,7 +11258,6 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11577,10 +11393,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowStepExe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11620,7 +11434,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowStepExe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11801,10 +11614,8 @@ func (m *awsRestjson1_deserializeOpPutComponentPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11844,7 +11655,6 @@ func (m *awsRestjson1_deserializeOpPutComponentPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11987,10 +11797,8 @@ func (m *awsRestjson1_deserializeOpPutContainerRecipeP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12030,7 +11838,6 @@ func (m *awsRestjson1_deserializeOpPutContainerRecipeP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12173,10 +11980,8 @@ func (m *awsRestjson1_deserializeOpPutImagePolicy) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12216,7 +12021,6 @@ func (m *awsRestjson1_deserializeOpPutImagePolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12359,10 +12163,8 @@ func (m *awsRestjson1_deserializeOpPutImageRecipePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12402,7 +12204,6 @@ func (m *awsRestjson1_deserializeOpPutImageRecipePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12545,10 +12346,8 @@ func (m *awsRestjson1_deserializeOpSendWorkflowStepAct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12588,7 +12387,6 @@ func (m *awsRestjson1_deserializeOpSendWorkflowStepAct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12746,10 +12544,8 @@ func (m *awsRestjson1_deserializeOpStartImagePipelineE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12789,7 +12585,6 @@ func (m *awsRestjson1_deserializeOpStartImagePipelineE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12944,10 +12739,8 @@ func (m *awsRestjson1_deserializeOpStartResourceStateU
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12987,7 +12780,6 @@ func (m *awsRestjson1_deserializeOpStartResourceStateU
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13133,10 +12925,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13148,7 +12938,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13227,10 +13016,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13242,7 +13029,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13321,10 +13107,8 @@ func (m *awsRestjson1_deserializeOpUpdateDistributionC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13364,7 +13148,6 @@ func (m *awsRestjson1_deserializeOpUpdateDistributionC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13519,10 +13302,8 @@ func (m *awsRestjson1_deserializeOpUpdateImagePipeline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13562,7 +13343,6 @@ func (m *awsRestjson1_deserializeOpUpdateImagePipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13714,10 +13494,8 @@ func (m *awsRestjson1_deserializeOpUpdateInfrastructur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13757,7 +13535,6 @@ func (m *awsRestjson1_deserializeOpUpdateInfrastructur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13909,10 +13686,8 @@ func (m *awsRestjson1_deserializeOpUpdateLifecyclePoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13952,7 +13727,6 @@ func (m *awsRestjson1_deserializeOpUpdateLifecyclePoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
