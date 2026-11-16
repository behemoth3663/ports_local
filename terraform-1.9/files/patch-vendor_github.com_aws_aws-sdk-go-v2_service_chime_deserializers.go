--- vendor/github.com/aws/aws-sdk-go-v2/service/chime/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chime/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 	output := &AssociatePhoneNumberWithUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -154,10 +150,8 @@ func (m *awsRestjson1_deserializeOpAssociateSigninDele
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -169,7 +163,6 @@ func (m *awsRestjson1_deserializeOpAssociateSigninDele
 	output := &AssociateSigninDelegateGroupsWithAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -260,10 +253,8 @@ func (m *awsRestjson1_deserializeOpBatchCreateRoomMemb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -303,7 +294,6 @@ func (m *awsRestjson1_deserializeOpBatchCreateRoomMemb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -430,10 +420,8 @@ func (m *awsRestjson1_deserializeOpBatchDeletePhoneNum
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -473,7 +461,6 @@ func (m *awsRestjson1_deserializeOpBatchDeletePhoneNum
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -600,10 +587,8 @@ func (m *awsRestjson1_deserializeOpBatchSuspendUser) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -643,7 +628,6 @@ func (m *awsRestjson1_deserializeOpBatchSuspendUser) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -770,10 +754,8 @@ func (m *awsRestjson1_deserializeOpBatchUnsuspendUser)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -813,7 +795,6 @@ func (m *awsRestjson1_deserializeOpBatchUnsuspendUser)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -940,10 +921,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdatePhoneNum
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -983,7 +962,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdatePhoneNum
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1110,10 +1088,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateUser) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1153,7 +1129,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateUser) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1280,10 +1255,8 @@ func (m *awsRestjson1_deserializeOpCreateAccount) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1323,7 +1296,6 @@ func (m *awsRestjson1_deserializeOpCreateAccount) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1450,10 +1422,8 @@ func (m *awsRestjson1_deserializeOpCreateBot) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1493,7 +1463,6 @@ func (m *awsRestjson1_deserializeOpCreateBot) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1623,10 +1592,8 @@ func (m *awsRestjson1_deserializeOpCreateMeetingDialOu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1666,7 +1633,6 @@ func (m *awsRestjson1_deserializeOpCreateMeetingDialOu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1800,10 +1766,8 @@ func (m *awsRestjson1_deserializeOpCreatePhoneNumberOr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1843,7 +1807,6 @@ func (m *awsRestjson1_deserializeOpCreatePhoneNumberOr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1973,10 +1936,8 @@ func (m *awsRestjson1_deserializeOpCreateRoom) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2016,7 +1977,6 @@ func (m *awsRestjson1_deserializeOpCreateRoom) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2146,10 +2106,8 @@ func (m *awsRestjson1_deserializeOpCreateRoomMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2189,7 +2147,6 @@ func (m *awsRestjson1_deserializeOpCreateRoomMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2322,10 +2279,8 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2365,7 +2320,6 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2495,10 +2449,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccount) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2510,7 +2462,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccount) Hand
 	output := &DeleteAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2604,10 +2555,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventsConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2625,7 +2574,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventsConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2713,10 +2661,8 @@ func (m *awsRestjson1_deserializeOpDeletePhoneNumber) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2734,7 +2680,6 @@ func (m *awsRestjson1_deserializeOpDeletePhoneNumber) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2825,10 +2770,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoom) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2846,7 +2789,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoom) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2937,10 +2879,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoomMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2958,7 +2898,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoomMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3049,10 +2988,8 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3064,7 +3001,6 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 	output := &DisassociatePhoneNumberFromUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3155,10 +3091,8 @@ func (m *awsRestjson1_deserializeOpDisassociateSigninD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3170,7 +3104,6 @@ func (m *awsRestjson1_deserializeOpDisassociateSigninD
 	output := &DisassociateSigninDelegateGroupsFromAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3261,10 +3194,8 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3304,7 +3235,6 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3431,10 +3361,8 @@ func (m *awsRestjson1_deserializeOpGetAccountSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3474,7 +3402,6 @@ func (m *awsRestjson1_deserializeOpGetAccountSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3601,10 +3528,8 @@ func (m *awsRestjson1_deserializeOpGetBot) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3644,7 +3569,6 @@ func (m *awsRestjson1_deserializeOpGetBot) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3771,10 +3695,8 @@ func (m *awsRestjson1_deserializeOpGetEventsConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3814,7 +3736,6 @@ func (m *awsRestjson1_deserializeOpGetEventsConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3941,10 +3862,8 @@ func (m *awsRestjson1_deserializeOpGetGlobalSettings) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3984,7 +3903,6 @@ func (m *awsRestjson1_deserializeOpGetGlobalSettings) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4113,10 +4031,8 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumber) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4156,7 +4072,6 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumber) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4283,10 +4198,8 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberOrder
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4326,7 +4239,6 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberOrder
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4453,10 +4365,8 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4496,7 +4406,6 @@ func (m *awsRestjson1_deserializeOpGetPhoneNumberSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4637,10 +4546,8 @@ func (m *awsRestjson1_deserializeOpGetRetentionSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4680,7 +4587,6 @@ func (m *awsRestjson1_deserializeOpGetRetentionSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4820,10 +4726,8 @@ func (m *awsRestjson1_deserializeOpGetRoom) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4863,7 +4767,6 @@ func (m *awsRestjson1_deserializeOpGetRoom) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4990,10 +4893,8 @@ func (m *awsRestjson1_deserializeOpGetUser) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5033,7 +4934,6 @@ func (m *awsRestjson1_deserializeOpGetUser) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5160,10 +5060,8 @@ func (m *awsRestjson1_deserializeOpGetUserSettings) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5203,7 +5101,6 @@ func (m *awsRestjson1_deserializeOpGetUserSettings) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5330,10 +5227,8 @@ func (m *awsRestjson1_deserializeOpInviteUsers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5373,7 +5268,6 @@ func (m *awsRestjson1_deserializeOpInviteUsers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5500,10 +5394,8 @@ func (m *awsRestjson1_deserializeOpListAccounts) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5543,7 +5435,6 @@ func (m *awsRestjson1_deserializeOpListAccounts) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5679,10 +5570,8 @@ func (m *awsRestjson1_deserializeOpListBots) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5722,7 +5611,6 @@ func (m *awsRestjson1_deserializeOpListBots) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5858,10 +5746,8 @@ func (m *awsRestjson1_deserializeOpListPhoneNumberOrde
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5901,7 +5787,6 @@ func (m *awsRestjson1_deserializeOpListPhoneNumberOrde
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6034,10 +5919,8 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbers) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6077,7 +5960,6 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbers) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6213,10 +6095,8 @@ func (m *awsRestjson1_deserializeOpListRoomMemberships
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6256,7 +6136,6 @@ func (m *awsRestjson1_deserializeOpListRoomMemberships
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6392,10 +6271,8 @@ func (m *awsRestjson1_deserializeOpListRooms) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6435,7 +6312,6 @@ func (m *awsRestjson1_deserializeOpListRooms) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6571,10 +6447,8 @@ func (m *awsRestjson1_deserializeOpListSupportedPhoneN
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6614,7 +6488,6 @@ func (m *awsRestjson1_deserializeOpListSupportedPhoneN
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6741,10 +6614,8 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6784,7 +6655,6 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6920,10 +6790,8 @@ func (m *awsRestjson1_deserializeOpLogoutUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6935,7 +6803,6 @@ func (m *awsRestjson1_deserializeOpLogoutUser) HandleD
 	output := &LogoutUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7026,10 +6893,8 @@ func (m *awsRestjson1_deserializeOpPutEventsConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7069,7 +6934,6 @@ func (m *awsRestjson1_deserializeOpPutEventsConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7196,10 +7060,8 @@ func (m *awsRestjson1_deserializeOpPutRetentionSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7239,7 +7101,6 @@ func (m *awsRestjson1_deserializeOpPutRetentionSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7382,10 +7243,8 @@ func (m *awsRestjson1_deserializeOpRedactConversationM
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7397,7 +7256,6 @@ func (m *awsRestjson1_deserializeOpRedactConversationM
 	output := &RedactConversationMessageOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7488,10 +7346,8 @@ func (m *awsRestjson1_deserializeOpRedactRoomMessage) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7503,7 +7359,6 @@ func (m *awsRestjson1_deserializeOpRedactRoomMessage) 
 	output := &RedactRoomMessageOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7594,10 +7449,8 @@ func (m *awsRestjson1_deserializeOpRegenerateSecurityT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7637,7 +7490,6 @@ func (m *awsRestjson1_deserializeOpRegenerateSecurityT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7764,10 +7616,8 @@ func (m *awsRestjson1_deserializeOpResetPersonalPIN) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7807,7 +7657,6 @@ func (m *awsRestjson1_deserializeOpResetPersonalPIN) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7934,10 +7783,8 @@ func (m *awsRestjson1_deserializeOpRestorePhoneNumber)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7977,7 +7824,6 @@ func (m *awsRestjson1_deserializeOpRestorePhoneNumber)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8107,10 +7953,8 @@ func (m *awsRestjson1_deserializeOpSearchAvailablePhon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8150,7 +7994,6 @@ func (m *awsRestjson1_deserializeOpSearchAvailablePhon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8286,10 +8129,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccount) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8329,7 +8170,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccount) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8456,10 +8296,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8471,7 +8309,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountSettin
 	output := &UpdateAccountSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8565,10 +8402,8 @@ func (m *awsRestjson1_deserializeOpUpdateBot) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8608,7 +8443,6 @@ func (m *awsRestjson1_deserializeOpUpdateBot) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8735,10 +8569,8 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8756,7 +8588,6 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8844,10 +8675,8 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumber) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8887,7 +8716,6 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumber) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9017,10 +8845,8 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumberSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9038,7 +8864,6 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumberSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9126,10 +8951,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoom) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9169,7 +8992,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoom) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9296,10 +9118,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoomMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9339,7 +9159,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoomMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9466,10 +9285,8 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9509,7 +9326,6 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9636,10 +9452,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9657,7 +9471,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
