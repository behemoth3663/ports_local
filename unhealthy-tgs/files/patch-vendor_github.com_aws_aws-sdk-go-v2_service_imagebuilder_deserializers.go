--- vendor/github.com/aws/aws-sdk-go-v2/service/imagebuilder/deserializers.go.orig	2026-11-21 14:46:20 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/imagebuilder/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -43,10 +42,8 @@ func (m *awsRestjson1_deserializeOpCancelImageCreation
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCancelImageCreation(response, &metadata)
@@ -82,7 +79,6 @@ func (m *awsRestjson1_deserializeOpCancelImageCreation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -241,10 +237,8 @@ func (m *awsRestjson1_deserializeOpCancelLifecycleExec
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCancelLifecycleExecution(response, &metadata)
@@ -280,7 +274,6 @@ func (m *awsRestjson1_deserializeOpCancelLifecycleExec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -421,10 +414,8 @@ func (m *awsRestjson1_deserializeOpCreateComponent) Ha
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateComponent(response, &metadata)
@@ -460,7 +451,6 @@ func (m *awsRestjson1_deserializeOpCreateComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -636,10 +626,8 @@ func (m *awsRestjson1_deserializeOpCreateContainerReci
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateContainerRecipe(response, &metadata)
@@ -675,7 +663,6 @@ func (m *awsRestjson1_deserializeOpCreateContainerReci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -851,10 +838,8 @@ func (m *awsRestjson1_deserializeOpCreateDistributionC
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateDistributionConfiguration(response, &metadata)
@@ -890,7 +875,6 @@ func (m *awsRestjson1_deserializeOpCreateDistributionC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1061,10 +1045,8 @@ func (m *awsRestjson1_deserializeOpCreateImage) Handle
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateImage(response, &metadata)
@@ -1100,7 +1082,6 @@ func (m *awsRestjson1_deserializeOpCreateImage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1267,10 +1248,8 @@ func (m *awsRestjson1_deserializeOpCreateImagePipeline
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateImagePipeline(response, &metadata)
@@ -1306,7 +1285,6 @@ func (m *awsRestjson1_deserializeOpCreateImagePipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1474,10 +1452,8 @@ func (m *awsRestjson1_deserializeOpCreateImageRecipe) 
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateImageRecipe(response, &metadata)
@@ -1513,7 +1489,6 @@ func (m *awsRestjson1_deserializeOpCreateImageRecipe) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1689,10 +1664,8 @@ func (m *awsRestjson1_deserializeOpCreateInfrastructur
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateInfrastructureConfiguration(response, &metadata)
@@ -1728,7 +1701,6 @@ func (m *awsRestjson1_deserializeOpCreateInfrastructur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1896,10 +1868,8 @@ func (m *awsRestjson1_deserializeOpCreateLifecyclePoli
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateLifecyclePolicy(response, &metadata)
@@ -1935,7 +1905,6 @@ func (m *awsRestjson1_deserializeOpCreateLifecyclePoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2094,10 +2063,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorCreateWorkflow(response, &metadata)
@@ -2133,7 +2100,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2300,10 +2266,8 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteComponent(response, &metadata)
@@ -2339,7 +2303,6 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2486,10 +2449,8 @@ func (m *awsRestjson1_deserializeOpDeleteContainerReci
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteContainerRecipe(response, &metadata)
@@ -2525,7 +2486,6 @@ func (m *awsRestjson1_deserializeOpDeleteContainerReci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2672,10 +2632,8 @@ func (m *awsRestjson1_deserializeOpDeleteDistributionC
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteDistributionConfiguration(response, &metadata)
@@ -2711,7 +2669,6 @@ func (m *awsRestjson1_deserializeOpDeleteDistributionC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2858,10 +2815,8 @@ func (m *awsRestjson1_deserializeOpDeleteImage) Handle
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteImage(response, &metadata)
@@ -2897,7 +2852,6 @@ func (m *awsRestjson1_deserializeOpDeleteImage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3044,10 +2998,8 @@ func (m *awsRestjson1_deserializeOpDeleteImagePipeline
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteImagePipeline(response, &metadata)
@@ -3083,7 +3035,6 @@ func (m *awsRestjson1_deserializeOpDeleteImagePipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3230,10 +3181,8 @@ func (m *awsRestjson1_deserializeOpDeleteImageRecipe) 
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteImageRecipe(response, &metadata)
@@ -3269,7 +3218,6 @@ func (m *awsRestjson1_deserializeOpDeleteImageRecipe) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3416,10 +3364,8 @@ func (m *awsRestjson1_deserializeOpDeleteInfrastructur
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteInfrastructureConfiguration(response, &metadata)
@@ -3455,7 +3401,6 @@ func (m *awsRestjson1_deserializeOpDeleteInfrastructur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3602,10 +3547,8 @@ func (m *awsRestjson1_deserializeOpDeleteLifecyclePoli
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteLifecyclePolicy(response, &metadata)
@@ -3641,7 +3584,6 @@ func (m *awsRestjson1_deserializeOpDeleteLifecyclePoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3779,10 +3721,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDeleteWorkflow(response, &metadata)
@@ -3818,7 +3758,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3956,10 +3895,8 @@ func (m *awsRestjson1_deserializeOpDistributeImage) Ha
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorDistributeImage(response, &metadata)
@@ -3995,7 +3932,6 @@ func (m *awsRestjson1_deserializeOpDistributeImage) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4157,10 +4093,8 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetComponent(response, &metadata)
@@ -4196,7 +4130,6 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4341,10 +4274,8 @@ func (m *awsRestjson1_deserializeOpGetComponentPolicy)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetComponentPolicy(response, &metadata)
@@ -4380,7 +4311,6 @@ func (m *awsRestjson1_deserializeOpGetComponentPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4524,10 +4454,8 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipe)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetContainerRecipe(response, &metadata)
@@ -4563,7 +4491,6 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipe)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4708,10 +4635,8 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipeP
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetContainerRecipePolicy(response, &metadata)
@@ -4747,7 +4672,6 @@ func (m *awsRestjson1_deserializeOpGetContainerRecipeP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4891,10 +4815,8 @@ func (m *awsRestjson1_deserializeOpGetDistributionConf
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetDistributionConfiguration(response, &metadata)
@@ -4930,7 +4852,6 @@ func (m *awsRestjson1_deserializeOpGetDistributionConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5070,10 +4991,8 @@ func (m *awsRestjson1_deserializeOpGetImage) HandleDes
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetImage(response, &metadata)
@@ -5109,7 +5028,6 @@ func (m *awsRestjson1_deserializeOpGetImage) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5254,10 +5172,8 @@ func (m *awsRestjson1_deserializeOpGetImagePipeline) H
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetImagePipeline(response, &metadata)
@@ -5293,7 +5209,6 @@ func (m *awsRestjson1_deserializeOpGetImagePipeline) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5433,10 +5348,8 @@ func (m *awsRestjson1_deserializeOpGetImagePolicy) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetImagePolicy(response, &metadata)
@@ -5472,7 +5385,6 @@ func (m *awsRestjson1_deserializeOpGetImagePolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5616,10 +5528,8 @@ func (m *awsRestjson1_deserializeOpGetImageRecipe) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetImageRecipe(response, &metadata)
@@ -5655,7 +5565,6 @@ func (m *awsRestjson1_deserializeOpGetImageRecipe) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5800,10 +5709,8 @@ func (m *awsRestjson1_deserializeOpGetImageRecipePolic
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetImageRecipePolicy(response, &metadata)
@@ -5839,7 +5746,6 @@ func (m *awsRestjson1_deserializeOpGetImageRecipePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5983,10 +5889,8 @@ func (m *awsRestjson1_deserializeOpGetInfrastructureCo
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetInfrastructureConfiguration(response, &metadata)
@@ -6022,7 +5926,6 @@ func (m *awsRestjson1_deserializeOpGetInfrastructureCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6162,10 +6065,8 @@ func (m *awsRestjson1_deserializeOpGetLifecycleExecuti
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetLifecycleExecution(response, &metadata)
@@ -6201,7 +6102,6 @@ func (m *awsRestjson1_deserializeOpGetLifecycleExecuti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6332,10 +6232,8 @@ func (m *awsRestjson1_deserializeOpGetLifecyclePolicy)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetLifecyclePolicy(response, &metadata)
@@ -6371,7 +6269,6 @@ func (m *awsRestjson1_deserializeOpGetLifecyclePolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6502,10 +6399,8 @@ func (m *awsRestjson1_deserializeOpGetMarketplaceResou
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetMarketplaceResource(response, &metadata)
@@ -6541,7 +6436,6 @@ func (m *awsRestjson1_deserializeOpGetMarketplaceResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6694,10 +6588,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetWorkflow(response, &metadata)
@@ -6733,7 +6625,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6869,10 +6760,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowExecutio
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetWorkflowExecution(response, &metadata)
@@ -6908,7 +6797,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowExecutio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7176,10 +7064,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStepExec
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorGetWorkflowStepExecution(response, &metadata)
@@ -7215,7 +7101,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowStepExec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7524,10 +7409,8 @@ func (m *awsRestjson1_deserializeOpImportComponent) Ha
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorImportComponent(response, &metadata)
@@ -7563,7 +7446,6 @@ func (m *awsRestjson1_deserializeOpImportComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7728,10 +7610,8 @@ func (m *awsRestjson1_deserializeOpImportDiskImage) Ha
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorImportDiskImage(response, &metadata)
@@ -7767,7 +7647,6 @@ func (m *awsRestjson1_deserializeOpImportDiskImage) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7908,10 +7787,8 @@ func (m *awsRestjson1_deserializeOpImportVmImage) Hand
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorImportVmImage(response, &metadata)
@@ -7947,7 +7824,6 @@ func (m *awsRestjson1_deserializeOpImportVmImage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8091,10 +7967,8 @@ func (m *awsRestjson1_deserializeOpListComponentBuildV
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListComponentBuildVersions(response, &metadata)
@@ -8130,7 +8004,6 @@ func (m *awsRestjson1_deserializeOpListComponentBuildV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8282,10 +8155,8 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListComponents(response, &metadata)
@@ -8321,7 +8192,6 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8473,10 +8343,8 @@ func (m *awsRestjson1_deserializeOpListContainerRecipe
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListContainerRecipes(response, &metadata)
@@ -8512,7 +8380,6 @@ func (m *awsRestjson1_deserializeOpListContainerRecipe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8664,10 +8531,8 @@ func (m *awsRestjson1_deserializeOpListDistributionCon
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListDistributionConfigurations(response, &metadata)
@@ -8703,7 +8568,6 @@ func (m *awsRestjson1_deserializeOpListDistributionCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8855,10 +8719,8 @@ func (m *awsRestjson1_deserializeOpListImageBuildVersi
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImageBuildVersions(response, &metadata)
@@ -8894,7 +8756,6 @@ func (m *awsRestjson1_deserializeOpListImageBuildVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9046,10 +8907,8 @@ func (m *awsRestjson1_deserializeOpListImagePackages) 
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImagePackages(response, &metadata)
@@ -9085,7 +8944,6 @@ func (m *awsRestjson1_deserializeOpListImagePackages) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9240,10 +9098,8 @@ func (m *awsRestjson1_deserializeOpListImagePipelineIm
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImagePipelineImages(response, &metadata)
@@ -9279,7 +9135,6 @@ func (m *awsRestjson1_deserializeOpListImagePipelineIm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9434,10 +9289,8 @@ func (m *awsRestjson1_deserializeOpListImagePipelines)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImagePipelines(response, &metadata)
@@ -9473,7 +9326,6 @@ func (m *awsRestjson1_deserializeOpListImagePipelines)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9625,10 +9477,8 @@ func (m *awsRestjson1_deserializeOpListImageRecipes) H
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImageRecipes(response, &metadata)
@@ -9664,7 +9514,6 @@ func (m *awsRestjson1_deserializeOpListImageRecipes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9816,10 +9665,8 @@ func (m *awsRestjson1_deserializeOpListImages) HandleD
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImages(response, &metadata)
@@ -9855,7 +9702,6 @@ func (m *awsRestjson1_deserializeOpListImages) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10007,10 +9853,8 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImageScanFindingAggregations(response, &metadata)
@@ -10046,7 +9890,6 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10207,10 +10050,8 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListImageScanFindings(response, &metadata)
@@ -10246,7 +10087,6 @@ func (m *awsRestjson1_deserializeOpListImageScanFindin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10398,10 +10238,8 @@ func (m *awsRestjson1_deserializeOpListInfrastructureC
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListInfrastructureConfigurations(response, &metadata)
@@ -10437,7 +10275,6 @@ func (m *awsRestjson1_deserializeOpListInfrastructureC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10589,10 +10426,8 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListLifecycleExecutionResources(response, &metadata)
@@ -10628,7 +10463,6 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10785,10 +10619,8 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListLifecycleExecutions(response, &metadata)
@@ -10824,7 +10656,6 @@ func (m *awsRestjson1_deserializeOpListLifecycleExecut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10967,10 +10798,8 @@ func (m *awsRestjson1_deserializeOpListLifecyclePolici
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListLifecyclePolicies(response, &metadata)
@@ -11006,7 +10835,6 @@ func (m *awsRestjson1_deserializeOpListLifecyclePolici
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11149,10 +10977,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListTagsForResource(response, &metadata)
@@ -11188,7 +11014,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11310,10 +11135,8 @@ func (m *awsRestjson1_deserializeOpListWaitingWorkflow
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListWaitingWorkflowSteps(response, &metadata)
@@ -11349,7 +11172,6 @@ func (m *awsRestjson1_deserializeOpListWaitingWorkflow
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11492,10 +11314,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowBuildVe
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListWorkflowBuildVersions(response, &metadata)
@@ -11531,7 +11351,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowBuildVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11674,10 +11493,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowExecuti
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListWorkflowExecutions(response, &metadata)
@@ -11713,7 +11530,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowExecuti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11883,10 +11699,8 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListWorkflows(response, &metadata)
@@ -11922,7 +11736,6 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12065,10 +11878,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowStepExe
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorListWorkflowStepExecutions(response, &metadata)
@@ -12104,7 +11915,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowStepExe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12292,10 +12102,8 @@ func (m *awsRestjson1_deserializeOpPutComponentPolicy)
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutComponentPolicy(response, &metadata)
@@ -12331,7 +12139,6 @@ func (m *awsRestjson1_deserializeOpPutComponentPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12481,10 +12288,8 @@ func (m *awsRestjson1_deserializeOpPutContainerRecipeP
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutContainerRecipePolicy(response, &metadata)
@@ -12520,7 +12325,6 @@ func (m *awsRestjson1_deserializeOpPutContainerRecipeP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12670,10 +12474,8 @@ func (m *awsRestjson1_deserializeOpPutImagePolicy) Han
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutImagePolicy(response, &metadata)
@@ -12709,7 +12511,6 @@ func (m *awsRestjson1_deserializeOpPutImagePolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12859,10 +12660,8 @@ func (m *awsRestjson1_deserializeOpPutImageRecipePolic
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorPutImageRecipePolicy(response, &metadata)
@@ -12898,7 +12697,6 @@ func (m *awsRestjson1_deserializeOpPutImageRecipePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13048,10 +12846,8 @@ func (m *awsRestjson1_deserializeOpRetryImage) HandleD
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorRetryImage(response, &metadata)
@@ -13087,7 +12883,6 @@ func (m *awsRestjson1_deserializeOpRetryImage) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13237,10 +13032,8 @@ func (m *awsRestjson1_deserializeOpSendWorkflowStepAct
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorSendWorkflowStepAction(response, &metadata)
@@ -13276,7 +13069,6 @@ func (m *awsRestjson1_deserializeOpSendWorkflowStepAct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13441,10 +13233,8 @@ func (m *awsRestjson1_deserializeOpStartImagePipelineE
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorStartImagePipelineExecution(response, &metadata)
@@ -13480,7 +13270,6 @@ func (m *awsRestjson1_deserializeOpStartImagePipelineE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13642,10 +13431,8 @@ func (m *awsRestjson1_deserializeOpStartResourceStateU
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorStartResourceStateUpdate(response, &metadata)
@@ -13681,7 +13468,6 @@ func (m *awsRestjson1_deserializeOpStartResourceStateU
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13834,10 +13620,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorTagResource(response, &metadata)
@@ -13845,7 +13629,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13931,10 +13714,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorUntagResource(response, &metadata)
@@ -13942,7 +13723,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14028,10 +13808,8 @@ func (m *awsRestjson1_deserializeOpUpdateDistributionC
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorUpdateDistributionConfiguration(response, &metadata)
@@ -14067,7 +13845,6 @@ func (m *awsRestjson1_deserializeOpUpdateDistributionC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14229,10 +14006,8 @@ func (m *awsRestjson1_deserializeOpUpdateImagePipeline
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorUpdateImagePipeline(response, &metadata)
@@ -14268,7 +14043,6 @@ func (m *awsRestjson1_deserializeOpUpdateImagePipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14427,10 +14201,8 @@ func (m *awsRestjson1_deserializeOpUpdateInfrastructur
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorUpdateInfrastructureConfiguration(response, &metadata)
@@ -14466,7 +14238,6 @@ func (m *awsRestjson1_deserializeOpUpdateInfrastructur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14625,10 +14396,8 @@ func (m *awsRestjson1_deserializeOpUpdateLifecyclePoli
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestjson1_deserializeOpErrorUpdateLifecyclePolicy(response, &metadata)
@@ -14664,7 +14433,6 @@ func (m *awsRestjson1_deserializeOpUpdateLifecyclePoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
