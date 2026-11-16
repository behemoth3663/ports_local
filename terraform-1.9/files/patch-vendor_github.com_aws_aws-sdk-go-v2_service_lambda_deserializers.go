--- vendor/github.com/aws/aws-sdk-go-v2/service/lambda/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lambda/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpAddLayerVersionPerm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -91,7 +88,6 @@ func (m *awsRestjson1_deserializeOpAddLayerVersionPerm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -231,10 +227,8 @@ func (m *awsRestjson1_deserializeOpAddPermission) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -274,7 +268,6 @@ func (m *awsRestjson1_deserializeOpAddPermission) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -405,10 +398,8 @@ func (m *awsRestjson1_deserializeOpCreateAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -448,7 +439,6 @@ func (m *awsRestjson1_deserializeOpCreateAlias) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -614,10 +604,8 @@ func (m *awsRestjson1_deserializeOpCreateCodeSigningCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -657,7 +645,6 @@ func (m *awsRestjson1_deserializeOpCreateCodeSigningCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -769,10 +756,8 @@ func (m *awsRestjson1_deserializeOpCreateEventSourceMa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -812,7 +797,6 @@ func (m *awsRestjson1_deserializeOpCreateEventSourceMa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1198,10 +1182,8 @@ func (m *awsRestjson1_deserializeOpCreateFunction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1241,7 +1223,6 @@ func (m *awsRestjson1_deserializeOpCreateFunction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1657,10 +1638,8 @@ func (m *awsRestjson1_deserializeOpCreateFunctionUrlCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1700,7 +1679,6 @@ func (m *awsRestjson1_deserializeOpCreateFunctionUrlCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1866,10 +1844,8 @@ func (m *awsRestjson1_deserializeOpDeleteAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1887,7 +1863,6 @@ func (m *awsRestjson1_deserializeOpDeleteAlias) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1969,10 +1944,8 @@ func (m *awsRestjson1_deserializeOpDeleteCodeSigningCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1984,7 +1957,6 @@ func (m *awsRestjson1_deserializeOpDeleteCodeSigningCo
 	output := &DeleteCodeSigningConfigOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2066,10 +2038,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventSourceMa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2109,7 +2079,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventSourceMa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2498,10 +2467,8 @@ func (m *awsRestjson1_deserializeOpDeleteFunction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2519,7 +2486,6 @@ func (m *awsRestjson1_deserializeOpDeleteFunction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2604,10 +2570,8 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionCodeS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2625,7 +2589,6 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionCodeS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2713,10 +2676,8 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionConcu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2734,7 +2695,6 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionConcu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2819,10 +2779,8 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionEvent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2840,7 +2798,6 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionEvent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2925,10 +2882,8 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionUrlCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2946,7 +2901,6 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionUrlCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3028,10 +2982,8 @@ func (m *awsRestjson1_deserializeOpDeleteLayerVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3049,7 +3001,6 @@ func (m *awsRestjson1_deserializeOpDeleteLayerVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3125,10 +3076,8 @@ func (m *awsRestjson1_deserializeOpDeleteProvisionedCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3146,7 +3095,6 @@ func (m *awsRestjson1_deserializeOpDeleteProvisionedCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3231,10 +3179,8 @@ func (m *awsRestjson1_deserializeOpGetAccountSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3274,7 +3220,6 @@ func (m *awsRestjson1_deserializeOpGetAccountSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3391,10 +3336,8 @@ func (m *awsRestjson1_deserializeOpGetAlias) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3434,7 +3377,6 @@ func (m *awsRestjson1_deserializeOpGetAlias) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3597,10 +3539,8 @@ func (m *awsRestjson1_deserializeOpGetCodeSigningConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3640,7 +3580,6 @@ func (m *awsRestjson1_deserializeOpGetCodeSigningConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3755,10 +3694,8 @@ func (m *awsRestjson1_deserializeOpGetEventSourceMappi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3798,7 +3735,6 @@ func (m *awsRestjson1_deserializeOpGetEventSourceMappi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4181,10 +4117,8 @@ func (m *awsRestjson1_deserializeOpGetFunction) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4224,7 +4158,6 @@ func (m *awsRestjson1_deserializeOpGetFunction) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4362,10 +4295,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionCodeSign
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4405,7 +4336,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionCodeSign
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4536,10 +4466,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionConcurre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4579,7 +4507,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionConcurre
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4705,10 +4632,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4748,7 +4673,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5149,10 +5073,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionEventInv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5192,7 +5114,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionEventInv
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5361,10 +5282,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionRecursio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5404,7 +5323,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionRecursio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5526,10 +5444,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionUrlConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5569,7 +5485,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionUrlConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5741,10 +5656,8 @@ func (m *awsRestjson1_deserializeOpGetLayerVersion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5784,7 +5697,6 @@ func (m *awsRestjson1_deserializeOpGetLayerVersion) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5970,10 +5882,8 @@ func (m *awsRestjson1_deserializeOpGetLayerVersionByAr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6013,7 +5923,6 @@ func (m *awsRestjson1_deserializeOpGetLayerVersionByAr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6199,10 +6108,8 @@ func (m *awsRestjson1_deserializeOpGetLayerVersionPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6242,7 +6149,6 @@ func (m *awsRestjson1_deserializeOpGetLayerVersionPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6373,10 +6279,8 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6416,7 +6320,6 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6547,10 +6450,8 @@ func (m *awsRestjson1_deserializeOpGetProvisionedConcu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6590,7 +6491,6 @@ func (m *awsRestjson1_deserializeOpGetProvisionedConcu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6772,10 +6672,8 @@ func (m *awsRestjson1_deserializeOpGetRuntimeManagemen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6815,7 +6713,6 @@ func (m *awsRestjson1_deserializeOpGetRuntimeManagemen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6955,10 +6852,8 @@ func (m *awsRestjson1_deserializeOpInvoke) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6980,7 +6875,6 @@ func (m *awsRestjson1_deserializeOpInvoke) HandleDeser
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7186,10 +7080,8 @@ func (m *awsRestjson1_deserializeOpInvokeAsync) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7206,7 +7098,6 @@ func (m *awsRestjson1_deserializeOpInvokeAsync) Handle
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7301,10 +7192,8 @@ func (m *awsRestjson1_deserializeOpInvokeWithResponseS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7321,7 +7210,6 @@ func (m *awsRestjson1_deserializeOpInvokeWithResponseS
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7501,10 +7389,8 @@ func (m *awsRestjson1_deserializeOpListAliases) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7544,7 +7430,6 @@ func (m *awsRestjson1_deserializeOpListAliases) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7671,10 +7556,8 @@ func (m *awsRestjson1_deserializeOpListCodeSigningConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7714,7 +7597,6 @@ func (m *awsRestjson1_deserializeOpListCodeSigningConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7835,10 +7717,8 @@ func (m *awsRestjson1_deserializeOpListEventSourceMapp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7878,7 +7758,6 @@ func (m *awsRestjson1_deserializeOpListEventSourceMapp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8005,10 +7884,8 @@ func (m *awsRestjson1_deserializeOpListFunctionEventIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8048,7 +7925,6 @@ func (m *awsRestjson1_deserializeOpListFunctionEventIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8175,10 +8051,8 @@ func (m *awsRestjson1_deserializeOpListFunctions) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8218,7 +8092,6 @@ func (m *awsRestjson1_deserializeOpListFunctions) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8342,10 +8215,8 @@ func (m *awsRestjson1_deserializeOpListFunctionsByCode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8385,7 +8256,6 @@ func (m *awsRestjson1_deserializeOpListFunctionsByCode
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8509,10 +8379,8 @@ func (m *awsRestjson1_deserializeOpListFunctionUrlConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8552,7 +8420,6 @@ func (m *awsRestjson1_deserializeOpListFunctionUrlConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8679,10 +8546,8 @@ func (m *awsRestjson1_deserializeOpListLayers) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8722,7 +8587,6 @@ func (m *awsRestjson1_deserializeOpListLayers) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8846,10 +8710,8 @@ func (m *awsRestjson1_deserializeOpListLayerVersions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8889,7 +8751,6 @@ func (m *awsRestjson1_deserializeOpListLayerVersions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9016,10 +8877,8 @@ func (m *awsRestjson1_deserializeOpListProvisionedConc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9059,7 +8918,6 @@ func (m *awsRestjson1_deserializeOpListProvisionedConc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9186,10 +9044,8 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9229,7 +9085,6 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9347,10 +9202,8 @@ func (m *awsRestjson1_deserializeOpListVersionsByFunct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9390,7 +9243,6 @@ func (m *awsRestjson1_deserializeOpListVersionsByFunct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9517,10 +9369,8 @@ func (m *awsRestjson1_deserializeOpPublishLayerVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9560,7 +9410,6 @@ func (m *awsRestjson1_deserializeOpPublishLayerVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9749,10 +9598,8 @@ func (m *awsRestjson1_deserializeOpPublishVersion) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9792,7 +9639,6 @@ func (m *awsRestjson1_deserializeOpPublishVersion) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10202,10 +10048,8 @@ func (m *awsRestjson1_deserializeOpPutFunctionCodeSign
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10245,7 +10089,6 @@ func (m *awsRestjson1_deserializeOpPutFunctionCodeSign
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10382,10 +10225,8 @@ func (m *awsRestjson1_deserializeOpPutFunctionConcurre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10425,7 +10266,6 @@ func (m *awsRestjson1_deserializeOpPutFunctionConcurre
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10554,10 +10394,8 @@ func (m *awsRestjson1_deserializeOpPutFunctionEventInv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10597,7 +10435,6 @@ func (m *awsRestjson1_deserializeOpPutFunctionEventInv
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10769,10 +10606,8 @@ func (m *awsRestjson1_deserializeOpPutFunctionRecursio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10812,7 +10647,6 @@ func (m *awsRestjson1_deserializeOpPutFunctionRecursio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10937,10 +10771,8 @@ func (m *awsRestjson1_deserializeOpPutProvisionedConcu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10980,7 +10812,6 @@ func (m *awsRestjson1_deserializeOpPutProvisionedConcu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11162,10 +10993,8 @@ func (m *awsRestjson1_deserializeOpPutRuntimeManagemen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11205,7 +11034,6 @@ func (m *awsRestjson1_deserializeOpPutRuntimeManagemen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11348,10 +11176,8 @@ func (m *awsRestjson1_deserializeOpRemoveLayerVersionP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11369,7 +11195,6 @@ func (m *awsRestjson1_deserializeOpRemoveLayerVersionP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11454,10 +11279,8 @@ func (m *awsRestjson1_deserializeOpRemovePermission) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11475,7 +11298,6 @@ func (m *awsRestjson1_deserializeOpRemovePermission) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11560,10 +11382,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11581,7 +11401,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11666,10 +11485,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11687,7 +11504,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11772,10 +11588,8 @@ func (m *awsRestjson1_deserializeOpUpdateAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11815,7 +11629,6 @@ func (m *awsRestjson1_deserializeOpUpdateAlias) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11984,10 +11797,8 @@ func (m *awsRestjson1_deserializeOpUpdateCodeSigningCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12027,7 +11838,6 @@ func (m *awsRestjson1_deserializeOpUpdateCodeSigningCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12142,10 +11952,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventSourceMa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12185,7 +11993,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventSourceMa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12574,10 +12381,8 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionCode)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12617,7 +12422,6 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionCode)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13036,10 +12840,8 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13079,7 +12881,6 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13495,10 +13296,8 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionEvent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13538,7 +13337,6 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionEvent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13710,10 +13508,8 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionUrlCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13753,7 +13549,6 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionUrlCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
