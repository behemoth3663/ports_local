--- vendor/github.com/aws/aws-sdk-go-v2/service/vpclattice/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/vpclattice/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateRule) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateRule) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -217,10 +213,8 @@ func (m *awsRestjson1_deserializeOpCreateAccessLogSubs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -260,7 +254,6 @@ func (m *awsRestjson1_deserializeOpCreateAccessLogSubs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -433,10 +426,8 @@ func (m *awsRestjson1_deserializeOpCreateListener) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -476,7 +467,6 @@ func (m *awsRestjson1_deserializeOpCreateListener) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -670,10 +660,8 @@ func (m *awsRestjson1_deserializeOpCreateResourceConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -713,7 +701,6 @@ func (m *awsRestjson1_deserializeOpCreateResourceConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -948,10 +935,8 @@ func (m *awsRestjson1_deserializeOpCreateResourceGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -991,7 +976,6 @@ func (m *awsRestjson1_deserializeOpCreateResourceGatew
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1177,10 +1161,8 @@ func (m *awsRestjson1_deserializeOpCreateRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1220,7 +1202,6 @@ func (m *awsRestjson1_deserializeOpCreateRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1392,10 +1373,8 @@ func (m *awsRestjson1_deserializeOpCreateService) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1435,7 +1414,6 @@ func (m *awsRestjson1_deserializeOpCreateService) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1625,10 +1603,8 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1668,7 +1644,6 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1831,10 +1806,8 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1874,7 +1847,6 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2032,10 +2004,8 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2075,7 +2045,6 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2247,10 +2216,8 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2290,7 +2257,6 @@ func (m *awsRestjson1_deserializeOpCreateServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2453,10 +2419,8 @@ func (m *awsRestjson1_deserializeOpCreateTargetGroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2496,7 +2460,6 @@ func (m *awsRestjson1_deserializeOpCreateTargetGroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2668,10 +2631,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccessLogSubs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2683,7 +2644,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccessLogSubs
 	output := &DeleteAccessLogSubscriptionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2768,10 +2728,8 @@ func (m *awsRestjson1_deserializeOpDeleteAuthPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2783,7 +2741,6 @@ func (m *awsRestjson1_deserializeOpDeleteAuthPolicy) H
 	output := &DeleteAuthPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2868,10 +2825,8 @@ func (m *awsRestjson1_deserializeOpDeleteListener) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2883,7 +2838,6 @@ func (m *awsRestjson1_deserializeOpDeleteListener) Han
 	output := &DeleteListenerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2971,10 +2925,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourceConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2986,7 +2938,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourceConfi
 	output := &DeleteResourceConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3074,10 +3025,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourceEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3117,7 +3066,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourceEndpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3278,10 +3226,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourceGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3321,7 +3267,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourceGatew
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3476,10 +3421,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3491,7 +3434,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 	output := &DeleteResourcePolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3576,10 +3518,8 @@ func (m *awsRestjson1_deserializeOpDeleteRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3591,7 +3531,6 @@ func (m *awsRestjson1_deserializeOpDeleteRule) HandleD
 	output := &DeleteRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3679,10 +3618,8 @@ func (m *awsRestjson1_deserializeOpDeleteService) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3722,7 +3659,6 @@ func (m *awsRestjson1_deserializeOpDeleteService) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3877,10 +3813,8 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3892,7 +3826,6 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 	output := &DeleteServiceNetworkOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3980,10 +3913,8 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4023,7 +3954,6 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4169,10 +4099,8 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4212,7 +4140,6 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4358,10 +4285,8 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4401,7 +4326,6 @@ func (m *awsRestjson1_deserializeOpDeleteServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4547,10 +4471,8 @@ func (m *awsRestjson1_deserializeOpDeleteTargetGroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4590,7 +4512,6 @@ func (m *awsRestjson1_deserializeOpDeleteTargetGroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4733,10 +4654,8 @@ func (m *awsRestjson1_deserializeOpDeregisterTargets) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4776,7 +4695,6 @@ func (m *awsRestjson1_deserializeOpDeregisterTargets) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4905,10 +4823,8 @@ func (m *awsRestjson1_deserializeOpGetAccessLogSubscri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4948,7 +4864,6 @@ func (m *awsRestjson1_deserializeOpGetAccessLogSubscri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5144,10 +5059,8 @@ func (m *awsRestjson1_deserializeOpGetAuthPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5187,7 +5100,6 @@ func (m *awsRestjson1_deserializeOpGetAuthPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5347,10 +5259,8 @@ func (m *awsRestjson1_deserializeOpGetListener) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5390,7 +5300,6 @@ func (m *awsRestjson1_deserializeOpGetListener) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5604,10 +5513,8 @@ func (m *awsRestjson1_deserializeOpGetResourceConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5647,7 +5554,6 @@ func (m *awsRestjson1_deserializeOpGetResourceConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5907,10 +5813,8 @@ func (m *awsRestjson1_deserializeOpGetResourceGateway)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5950,7 +5854,6 @@ func (m *awsRestjson1_deserializeOpGetResourceGateway)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6156,10 +6059,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6199,7 +6100,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6324,10 +6224,8 @@ func (m *awsRestjson1_deserializeOpGetRule) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6367,7 +6265,6 @@ func (m *awsRestjson1_deserializeOpGetRule) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6568,10 +6465,8 @@ func (m *awsRestjson1_deserializeOpGetService) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6611,7 +6506,6 @@ func (m *awsRestjson1_deserializeOpGetService) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6839,10 +6733,8 @@ func (m *awsRestjson1_deserializeOpGetServiceNetwork) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6882,7 +6774,6 @@ func (m *awsRestjson1_deserializeOpGetServiceNetwork) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7091,10 +6982,8 @@ func (m *awsRestjson1_deserializeOpGetServiceNetworkRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7134,7 +7023,6 @@ func (m *awsRestjson1_deserializeOpGetServiceNetworkRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7403,10 +7291,8 @@ func (m *awsRestjson1_deserializeOpGetServiceNetworkSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7446,7 +7332,6 @@ func (m *awsRestjson1_deserializeOpGetServiceNetworkSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7697,10 +7582,8 @@ func (m *awsRestjson1_deserializeOpGetServiceNetworkVp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7740,7 +7623,6 @@ func (m *awsRestjson1_deserializeOpGetServiceNetworkVp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7977,10 +7859,8 @@ func (m *awsRestjson1_deserializeOpGetTargetGroup) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8020,7 +7900,6 @@ func (m *awsRestjson1_deserializeOpGetTargetGroup) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8235,10 +8114,8 @@ func (m *awsRestjson1_deserializeOpListAccessLogSubscr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8278,7 +8155,6 @@ func (m *awsRestjson1_deserializeOpListAccessLogSubscr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8405,10 +8281,8 @@ func (m *awsRestjson1_deserializeOpListListeners) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8448,7 +8322,6 @@ func (m *awsRestjson1_deserializeOpListListeners) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8578,10 +8451,8 @@ func (m *awsRestjson1_deserializeOpListResourceConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8621,7 +8492,6 @@ func (m *awsRestjson1_deserializeOpListResourceConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8748,10 +8618,8 @@ func (m *awsRestjson1_deserializeOpListResourceEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8791,7 +8659,6 @@ func (m *awsRestjson1_deserializeOpListResourceEndpoin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8918,10 +8785,8 @@ func (m *awsRestjson1_deserializeOpListResourceGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8961,7 +8826,6 @@ func (m *awsRestjson1_deserializeOpListResourceGateway
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9088,10 +8952,8 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9131,7 +8993,6 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9261,10 +9122,8 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9304,7 +9163,6 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9431,10 +9289,8 @@ func (m *awsRestjson1_deserializeOpListServiceNetworks
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9474,7 +9330,6 @@ func (m *awsRestjson1_deserializeOpListServiceNetworks
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9601,10 +9456,8 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9644,7 +9497,6 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9771,10 +9623,8 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9814,7 +9664,6 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9941,10 +9790,8 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9984,7 +9831,6 @@ func (m *awsRestjson1_deserializeOpListServiceNetworkV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10111,10 +9957,8 @@ func (m *awsRestjson1_deserializeOpListServices) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10154,7 +9998,6 @@ func (m *awsRestjson1_deserializeOpListServices) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10281,10 +10124,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10324,7 +10165,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10442,10 +10282,8 @@ func (m *awsRestjson1_deserializeOpListTargetGroups) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10485,7 +10323,6 @@ func (m *awsRestjson1_deserializeOpListTargetGroups) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10612,10 +10449,8 @@ func (m *awsRestjson1_deserializeOpListTargets) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10655,7 +10490,6 @@ func (m *awsRestjson1_deserializeOpListTargets) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10785,10 +10619,8 @@ func (m *awsRestjson1_deserializeOpPutAuthPolicy) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10828,7 +10660,6 @@ func (m *awsRestjson1_deserializeOpPutAuthPolicy) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10962,10 +10793,8 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10977,7 +10806,6 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 	output := &PutResourcePolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11062,10 +10890,8 @@ func (m *awsRestjson1_deserializeOpRegisterTargets) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11105,7 +10931,6 @@ func (m *awsRestjson1_deserializeOpRegisterTargets) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11237,10 +11062,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11252,7 +11075,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11337,10 +11159,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11352,7 +11172,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11434,10 +11253,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccessLogSubs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11477,7 +11294,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccessLogSubs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11641,10 +11457,8 @@ func (m *awsRestjson1_deserializeOpUpdateListener) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11684,7 +11498,6 @@ func (m *awsRestjson1_deserializeOpUpdateListener) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11878,10 +11691,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11921,7 +11732,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12131,10 +11941,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12174,7 +11982,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceGatew
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12354,10 +12161,8 @@ func (m *awsRestjson1_deserializeOpUpdateRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12397,7 +12202,6 @@ func (m *awsRestjson1_deserializeOpUpdateRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12578,10 +12382,8 @@ func (m *awsRestjson1_deserializeOpUpdateService) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12621,7 +12423,6 @@ func (m *awsRestjson1_deserializeOpUpdateService) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12797,10 +12598,8 @@ func (m *awsRestjson1_deserializeOpUpdateServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12840,7 +12639,6 @@ func (m *awsRestjson1_deserializeOpUpdateServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12995,10 +12793,8 @@ func (m *awsRestjson1_deserializeOpUpdateServiceNetwor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13038,7 +12834,6 @@ func (m *awsRestjson1_deserializeOpUpdateServiceNetwor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13198,10 +12993,8 @@ func (m *awsRestjson1_deserializeOpUpdateTargetGroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13241,7 +13034,6 @@ func (m *awsRestjson1_deserializeOpUpdateTargetGroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
