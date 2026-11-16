--- vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -219,10 +215,8 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -262,7 +256,6 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -392,10 +385,8 @@ func (m *awsRestjson1_deserializeOpBatchDeletePhoneNum
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -435,7 +426,6 @@ func (m *awsRestjson1_deserializeOpBatchDeletePhoneNum
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -562,10 +552,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdatePhoneNum
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -605,7 +593,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdatePhoneNum
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -732,10 +719,8 @@ func (m *awsRestjson1_deserializeOpCreatePhoneNumberOr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -775,7 +760,6 @@ func (m *awsRestjson1_deserializeOpCreatePhoneNumberOr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -905,10 +889,8 @@ func (m *awsRestjson1_deserializeOpCreateProxySession)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -948,7 +930,6 @@ func (m *awsRestjson1_deserializeOpCreateProxySession)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1075,10 +1056,8 @@ func (m *awsRestjson1_deserializeOpCreateSipMediaAppli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1118,7 +1097,6 @@ func (m *awsRestjson1_deserializeOpCreateSipMediaAppli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1251,10 +1229,8 @@ func (m *awsRestjson1_deserializeOpCreateSipMediaAppli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1294,7 +1270,6 @@ func (m *awsRestjson1_deserializeOpCreateSipMediaAppli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1424,10 +1399,8 @@ func (m *awsRestjson1_deserializeOpCreateSipRule) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1467,7 +1440,6 @@ func (m *awsRestjson1_deserializeOpCreateSipRule) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1600,10 +1572,8 @@ func (m *awsRestjson1_deserializeOpCreateVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1643,7 +1613,6 @@ func (m *awsRestjson1_deserializeOpCreateVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1773,10 +1742,8 @@ func (m *awsRestjson1_deserializeOpCreateVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1816,7 +1783,6 @@ func (m *awsRestjson1_deserializeOpCreateVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1946,10 +1912,8 @@ func (m *awsRestjson1_deserializeOpCreateVoiceProfile)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1989,7 +1953,6 @@ func (m *awsRestjson1_deserializeOpCreateVoiceProfile)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2128,10 +2091,8 @@ func (m *awsRestjson1_deserializeOpCreateVoiceProfileD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2171,7 +2132,6 @@ func (m *awsRestjson1_deserializeOpCreateVoiceProfileD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2304,10 +2264,8 @@ func (m *awsRestjson1_deserializeOpDeletePhoneNumber) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2325,7 +2283,6 @@ func (m *awsRestjson1_deserializeOpDeletePhoneNumber) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2416,10 +2373,8 @@ func (m *awsRestjson1_deserializeOpDeleteProxySession)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2437,7 +2392,6 @@ func (m *awsRestjson1_deserializeOpDeleteProxySession)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2528,10 +2482,8 @@ func (m *awsRestjson1_deserializeOpDeleteSipMediaAppli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2549,7 +2501,6 @@ func (m *awsRestjson1_deserializeOpDeleteSipMediaAppli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2643,10 +2594,8 @@ func (m *awsRestjson1_deserializeOpDeleteSipRule) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2664,7 +2613,6 @@ func (m *awsRestjson1_deserializeOpDeleteSipRule) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2758,10 +2706,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2779,7 +2725,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2873,10 +2818,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2894,7 +2837,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2985,10 +2927,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3006,7 +2946,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3097,10 +3036,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3118,7 +3055,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3212,10 +3148,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3233,7 +3167,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3324,10 +3257,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3345,7 +3276,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3436,10 +3366,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3457,7 +3385,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3548,10 +3475,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3569,7 +3494,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3660,10 +3584,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3681,7 +3603,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3772,10 +3693,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceProfile)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3793,7 +3712,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceProfile)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3890,10 +3808,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceProfileD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3911,7 +3827,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceProfileD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4008,10 +3923,8 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4051,7 +3964,6 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4178,10 +4090,8 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4221,7 +4131,6 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4348,10 +4257,8 @@ func (m *awsRestjson1_deserializeOpGetGlobalSettings) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4391,7 +4298,6 @@ func (m *awsRestjson1_deserializeOpGetGlobalSettings) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4515,10 +4421,8 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumber) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4558,7 +4462,6 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumber) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4685,10 +4588,8 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberOrder
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4728,7 +4629,6 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberOrder
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4855,10 +4755,8 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4898,7 +4796,6 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5039,10 +4936,8 @@ func (m *awsRestjson1_deserializeOpGetProxySession) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5082,7 +4977,6 @@ func (m *awsRestjson1_deserializeOpGetProxySession) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5209,10 +5103,8 @@ func (m *awsRestjson1_deserializeOpGetSipMediaApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5252,7 +5144,6 @@ func (m *awsRestjson1_deserializeOpGetSipMediaApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5379,10 +5270,8 @@ func (m *awsRestjson1_deserializeOpGetSipMediaApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5422,7 +5311,6 @@ func (m *awsRestjson1_deserializeOpGetSipMediaApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5549,10 +5437,8 @@ func (m *awsRestjson1_deserializeOpGetSipMediaApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5592,7 +5478,6 @@ func (m *awsRestjson1_deserializeOpGetSipMediaApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5719,10 +5604,8 @@ func (m *awsRestjson1_deserializeOpGetSipRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5762,7 +5645,6 @@ func (m *awsRestjson1_deserializeOpGetSipRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5889,10 +5771,8 @@ func (m *awsRestjson1_deserializeOpGetSpeakerSearchTas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5932,7 +5812,6 @@ func (m *awsRestjson1_deserializeOpGetSpeakerSearchTas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6065,10 +5944,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnector) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6108,7 +5985,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnector) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6235,10 +6111,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorEm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6278,7 +6152,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorEm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6405,10 +6278,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorEx
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6448,7 +6319,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorEx
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6575,10 +6445,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6618,7 +6486,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorGr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6745,10 +6612,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorLo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6788,7 +6653,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorLo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6915,10 +6779,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorOr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6958,7 +6820,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorOr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7085,10 +6946,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7128,7 +6987,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7255,10 +7113,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7298,7 +7154,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7425,10 +7280,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7468,7 +7321,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7595,10 +7447,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7638,7 +7488,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceConnectorTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7765,10 +7614,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceProfile) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7808,7 +7655,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceProfile) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7938,10 +7784,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceProfileDoma
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7981,7 +7825,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceProfileDoma
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8111,10 +7954,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceToneAnalysi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8154,7 +7995,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceToneAnalysi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8287,10 +8127,8 @@ func (m *awsRestjson1_deserializeOpListAvailableVoiceC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8330,7 +8168,6 @@ func (m *awsRestjson1_deserializeOpListAvailableVoiceC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8454,10 +8291,8 @@ func (m *awsRestjson1_deserializeOpListPhoneNumberOrde
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8497,7 +8332,6 @@ func (m *awsRestjson1_deserializeOpListPhoneNumberOrde
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8630,10 +8464,8 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbers) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8673,7 +8505,6 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbers) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8809,10 +8640,8 @@ func (m *awsRestjson1_deserializeOpListProxySessions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8852,7 +8681,6 @@ func (m *awsRestjson1_deserializeOpListProxySessions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8988,10 +8816,8 @@ func (m *awsRestjson1_deserializeOpListSipMediaApplica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9031,7 +8857,6 @@ func (m *awsRestjson1_deserializeOpListSipMediaApplica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9164,10 +8989,8 @@ func (m *awsRestjson1_deserializeOpListSipRules) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9207,7 +9030,6 @@ func (m *awsRestjson1_deserializeOpListSipRules) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9340,10 +9162,8 @@ func (m *awsRestjson1_deserializeOpListSupportedPhoneN
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9383,7 +9203,6 @@ func (m *awsRestjson1_deserializeOpListSupportedPhoneN
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9510,10 +9329,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9553,7 +9370,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9677,10 +9493,8 @@ func (m *awsRestjson1_deserializeOpListVoiceConnectorG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9720,7 +9534,6 @@ func (m *awsRestjson1_deserializeOpListVoiceConnectorG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9853,10 +9666,8 @@ func (m *awsRestjson1_deserializeOpListVoiceConnectors
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9896,7 +9707,6 @@ func (m *awsRestjson1_deserializeOpListVoiceConnectors
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10029,10 +9839,8 @@ func (m *awsRestjson1_deserializeOpListVoiceConnectorT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10072,7 +9880,6 @@ func (m *awsRestjson1_deserializeOpListVoiceConnectorT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10199,10 +10006,8 @@ func (m *awsRestjson1_deserializeOpListVoiceProfileDom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10242,7 +10047,6 @@ func (m *awsRestjson1_deserializeOpListVoiceProfileDom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10378,10 +10182,8 @@ func (m *awsRestjson1_deserializeOpListVoiceProfiles) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10421,7 +10223,6 @@ func (m *awsRestjson1_deserializeOpListVoiceProfiles) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10557,10 +10358,8 @@ func (m *awsRestjson1_deserializeOpPutSipMediaApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10600,7 +10399,6 @@ func (m *awsRestjson1_deserializeOpPutSipMediaApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10727,10 +10525,8 @@ func (m *awsRestjson1_deserializeOpPutSipMediaApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10770,7 +10566,6 @@ func (m *awsRestjson1_deserializeOpPutSipMediaApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10897,10 +10692,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorEm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10940,7 +10733,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorEm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11067,10 +10859,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorEx
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11110,7 +10900,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorEx
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11240,10 +11029,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorLo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11283,7 +11070,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorLo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11410,10 +11196,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorOr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11453,7 +11237,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorOr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11580,10 +11363,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11623,7 +11404,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11753,10 +11533,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11796,7 +11574,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11923,10 +11700,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11966,7 +11741,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12096,10 +11870,8 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12117,7 +11889,6 @@ func (m *awsRestjson1_deserializeOpPutVoiceConnectorTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12208,10 +11979,8 @@ func (m *awsRestjson1_deserializeOpRestorePhoneNumber)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12251,7 +12020,6 @@ func (m *awsRestjson1_deserializeOpRestorePhoneNumber)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12381,10 +12149,8 @@ func (m *awsRestjson1_deserializeOpSearchAvailablePhon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12424,7 +12190,6 @@ func (m *awsRestjson1_deserializeOpSearchAvailablePhon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12560,10 +12325,8 @@ func (m *awsRestjson1_deserializeOpStartSpeakerSearchT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12603,7 +12366,6 @@ func (m *awsRestjson1_deserializeOpStartSpeakerSearchT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12745,10 +12507,8 @@ func (m *awsRestjson1_deserializeOpStartVoiceToneAnaly
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12788,7 +12548,6 @@ func (m *awsRestjson1_deserializeOpStartVoiceToneAnaly
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12930,10 +12689,8 @@ func (m *awsRestjson1_deserializeOpStopSpeakerSearchTa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12951,7 +12708,6 @@ func (m *awsRestjson1_deserializeOpStopSpeakerSearchTa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13051,10 +12807,8 @@ func (m *awsRestjson1_deserializeOpStopVoiceToneAnalys
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13072,7 +12826,6 @@ func (m *awsRestjson1_deserializeOpStopVoiceToneAnalys
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13172,10 +12925,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13193,7 +12944,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13284,10 +13034,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13305,7 +13053,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13393,10 +13140,8 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13414,7 +13159,6 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13502,10 +13246,8 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumber) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13545,7 +13287,6 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumber) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13675,10 +13416,8 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumberSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13696,7 +13435,6 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumberSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13784,10 +13522,8 @@ func (m *awsRestjson1_deserializeOpUpdateProxySession)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13827,7 +13563,6 @@ func (m *awsRestjson1_deserializeOpUpdateProxySession)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13954,10 +13689,8 @@ func (m *awsRestjson1_deserializeOpUpdateSipMediaAppli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13997,7 +13730,6 @@ func (m *awsRestjson1_deserializeOpUpdateSipMediaAppli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14127,10 +13859,8 @@ func (m *awsRestjson1_deserializeOpUpdateSipMediaAppli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14170,7 +13900,6 @@ func (m *awsRestjson1_deserializeOpUpdateSipMediaAppli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14300,10 +14029,8 @@ func (m *awsRestjson1_deserializeOpUpdateSipRule) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14343,7 +14070,6 @@ func (m *awsRestjson1_deserializeOpUpdateSipRule) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14476,10 +14202,8 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14519,7 +14243,6 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14646,10 +14369,8 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14689,7 +14410,6 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14819,10 +14539,8 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceProfile)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14862,7 +14580,6 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceProfile)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14998,10 +14715,8 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceProfileD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15041,7 +14756,6 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceProfileD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15171,10 +14885,8 @@ func (m *awsRestjson1_deserializeOpValidateE911Address
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15214,7 +14926,6 @@ func (m *awsRestjson1_deserializeOpValidateE911Address
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
