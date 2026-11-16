--- vendor/github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateCell) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateCell) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -237,10 +233,8 @@ func (m *awsRestjson1_deserializeOpCreateCrossAccountA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -280,7 +274,6 @@ func (m *awsRestjson1_deserializeOpCreateCrossAccountA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -405,10 +398,8 @@ func (m *awsRestjson1_deserializeOpCreateReadinessChec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -448,7 +439,6 @@ func (m *awsRestjson1_deserializeOpCreateReadinessChec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -596,10 +586,8 @@ func (m *awsRestjson1_deserializeOpCreateRecoveryGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -639,7 +627,6 @@ func (m *awsRestjson1_deserializeOpCreateRecoveryGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -783,10 +770,8 @@ func (m *awsRestjson1_deserializeOpCreateResourceSet) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -826,7 +811,6 @@ func (m *awsRestjson1_deserializeOpCreateResourceSet) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -979,10 +963,8 @@ func (m *awsRestjson1_deserializeOpDeleteCell) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1000,7 +982,6 @@ func (m *awsRestjson1_deserializeOpDeleteCell) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1085,10 +1066,8 @@ func (m *awsRestjson1_deserializeOpDeleteCrossAccountA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1100,7 +1079,6 @@ func (m *awsRestjson1_deserializeOpDeleteCrossAccountA
 	output := &DeleteCrossAccountAuthorizationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1182,10 +1160,8 @@ func (m *awsRestjson1_deserializeOpDeleteReadinessChec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1203,7 +1179,6 @@ func (m *awsRestjson1_deserializeOpDeleteReadinessChec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1288,10 +1263,8 @@ func (m *awsRestjson1_deserializeOpDeleteRecoveryGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1309,7 +1282,6 @@ func (m *awsRestjson1_deserializeOpDeleteRecoveryGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1394,10 +1366,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourceSet) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1415,7 +1385,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourceSet) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1500,10 +1469,8 @@ func (m *awsRestjson1_deserializeOpGetArchitectureReco
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1543,7 +1510,6 @@ func (m *awsRestjson1_deserializeOpGetArchitectureReco
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1686,10 +1652,8 @@ func (m *awsRestjson1_deserializeOpGetCell) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1729,7 +1693,6 @@ func (m *awsRestjson1_deserializeOpGetCell) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1878,10 +1841,8 @@ func (m *awsRestjson1_deserializeOpGetCellReadinessSum
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1921,7 +1882,6 @@ func (m *awsRestjson1_deserializeOpGetCellReadinessSum
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2060,10 +2020,8 @@ func (m *awsRestjson1_deserializeOpGetReadinessCheck) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2103,7 +2061,6 @@ func (m *awsRestjson1_deserializeOpGetReadinessCheck) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2251,10 +2208,8 @@ func (m *awsRestjson1_deserializeOpGetReadinessCheckRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2294,7 +2249,6 @@ func (m *awsRestjson1_deserializeOpGetReadinessCheckRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2433,10 +2387,8 @@ func (m *awsRestjson1_deserializeOpGetReadinessCheckSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2476,7 +2428,6 @@ func (m *awsRestjson1_deserializeOpGetReadinessCheckSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2620,10 +2571,8 @@ func (m *awsRestjson1_deserializeOpGetRecoveryGroup) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2663,7 +2612,6 @@ func (m *awsRestjson1_deserializeOpGetRecoveryGroup) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2807,10 +2755,8 @@ func (m *awsRestjson1_deserializeOpGetRecoveryGroupRea
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2850,7 +2796,6 @@ func (m *awsRestjson1_deserializeOpGetRecoveryGroupRea
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2989,10 +2934,8 @@ func (m *awsRestjson1_deserializeOpGetResourceSet) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3032,7 +2975,6 @@ func (m *awsRestjson1_deserializeOpGetResourceSet) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3185,10 +3127,8 @@ func (m *awsRestjson1_deserializeOpListCells) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3228,7 +3168,6 @@ func (m *awsRestjson1_deserializeOpListCells) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3355,10 +3294,8 @@ func (m *awsRestjson1_deserializeOpListCrossAccountAut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3398,7 +3335,6 @@ func (m *awsRestjson1_deserializeOpListCrossAccountAut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3525,10 +3461,8 @@ func (m *awsRestjson1_deserializeOpListReadinessChecks
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3568,7 +3502,6 @@ func (m *awsRestjson1_deserializeOpListReadinessChecks
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3695,10 +3628,8 @@ func (m *awsRestjson1_deserializeOpListRecoveryGroups)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3738,7 +3669,6 @@ func (m *awsRestjson1_deserializeOpListRecoveryGroups)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3865,10 +3795,8 @@ func (m *awsRestjson1_deserializeOpListResourceSets) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3908,7 +3836,6 @@ func (m *awsRestjson1_deserializeOpListResourceSets) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4035,10 +3962,8 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4078,7 +4003,6 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4205,10 +4129,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4248,7 +4170,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4363,10 +4284,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4378,7 +4297,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4457,10 +4375,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4478,7 +4394,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4557,10 +4472,8 @@ func (m *awsRestjson1_deserializeOpUpdateCell) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4600,7 +4513,6 @@ func (m *awsRestjson1_deserializeOpUpdateCell) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4749,10 +4661,8 @@ func (m *awsRestjson1_deserializeOpUpdateReadinessChec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4792,7 +4702,6 @@ func (m *awsRestjson1_deserializeOpUpdateReadinessChec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4940,10 +4849,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecoveryGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4983,7 +4890,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecoveryGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5127,10 +5033,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceSet) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5170,7 +5074,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceSet) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
