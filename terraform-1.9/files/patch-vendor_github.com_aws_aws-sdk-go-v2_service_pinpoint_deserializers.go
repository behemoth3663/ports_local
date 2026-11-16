--- vendor/github.com/aws/aws-sdk-go-v2/service/pinpoint/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/pinpoint/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -216,10 +212,8 @@ func (m *awsRestjson1_deserializeOpCreateCampaign) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -259,7 +253,6 @@ func (m *awsRestjson1_deserializeOpCreateCampaign) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -386,10 +379,8 @@ func (m *awsRestjson1_deserializeOpCreateEmailTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -429,7 +420,6 @@ func (m *awsRestjson1_deserializeOpCreateEmailTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -550,10 +540,8 @@ func (m *awsRestjson1_deserializeOpCreateExportJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -593,7 +581,6 @@ func (m *awsRestjson1_deserializeOpCreateExportJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -720,10 +707,8 @@ func (m *awsRestjson1_deserializeOpCreateImportJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -763,7 +748,6 @@ func (m *awsRestjson1_deserializeOpCreateImportJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -890,10 +874,8 @@ func (m *awsRestjson1_deserializeOpCreateInAppTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -933,7 +915,6 @@ func (m *awsRestjson1_deserializeOpCreateInAppTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1054,10 +1035,8 @@ func (m *awsRestjson1_deserializeOpCreateJourney) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1097,7 +1076,6 @@ func (m *awsRestjson1_deserializeOpCreateJourney) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1224,10 +1202,8 @@ func (m *awsRestjson1_deserializeOpCreatePushTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1267,7 +1243,6 @@ func (m *awsRestjson1_deserializeOpCreatePushTemplate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1388,10 +1363,8 @@ func (m *awsRestjson1_deserializeOpCreateRecommenderCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1431,7 +1404,6 @@ func (m *awsRestjson1_deserializeOpCreateRecommenderCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1558,10 +1530,8 @@ func (m *awsRestjson1_deserializeOpCreateSegment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1601,7 +1571,6 @@ func (m *awsRestjson1_deserializeOpCreateSegment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1728,10 +1697,8 @@ func (m *awsRestjson1_deserializeOpCreateSmsTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1771,7 +1738,6 @@ func (m *awsRestjson1_deserializeOpCreateSmsTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1892,10 +1858,8 @@ func (m *awsRestjson1_deserializeOpCreateVoiceTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1935,7 +1899,6 @@ func (m *awsRestjson1_deserializeOpCreateVoiceTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2056,10 +2019,8 @@ func (m *awsRestjson1_deserializeOpDeleteAdmChannel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2099,7 +2060,6 @@ func (m *awsRestjson1_deserializeOpDeleteAdmChannel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2226,10 +2186,8 @@ func (m *awsRestjson1_deserializeOpDeleteApnsChannel) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2269,7 +2227,6 @@ func (m *awsRestjson1_deserializeOpDeleteApnsChannel) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2396,10 +2353,8 @@ func (m *awsRestjson1_deserializeOpDeleteApnsSandboxCh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2439,7 +2394,6 @@ func (m *awsRestjson1_deserializeOpDeleteApnsSandboxCh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2566,10 +2520,8 @@ func (m *awsRestjson1_deserializeOpDeleteApnsVoipChann
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2609,7 +2561,6 @@ func (m *awsRestjson1_deserializeOpDeleteApnsVoipChann
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2736,10 +2687,8 @@ func (m *awsRestjson1_deserializeOpDeleteApnsVoipSandb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2779,7 +2728,6 @@ func (m *awsRestjson1_deserializeOpDeleteApnsVoipSandb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2906,10 +2854,8 @@ func (m *awsRestjson1_deserializeOpDeleteApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2949,7 +2895,6 @@ func (m *awsRestjson1_deserializeOpDeleteApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3076,10 +3021,8 @@ func (m *awsRestjson1_deserializeOpDeleteBaiduChannel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3119,7 +3062,6 @@ func (m *awsRestjson1_deserializeOpDeleteBaiduChannel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3246,10 +3188,8 @@ func (m *awsRestjson1_deserializeOpDeleteCampaign) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3289,7 +3229,6 @@ func (m *awsRestjson1_deserializeOpDeleteCampaign) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3416,10 +3355,8 @@ func (m *awsRestjson1_deserializeOpDeleteEmailChannel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3459,7 +3396,6 @@ func (m *awsRestjson1_deserializeOpDeleteEmailChannel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3586,10 +3522,8 @@ func (m *awsRestjson1_deserializeOpDeleteEmailTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3629,7 +3563,6 @@ func (m *awsRestjson1_deserializeOpDeleteEmailTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3756,10 +3689,8 @@ func (m *awsRestjson1_deserializeOpDeleteEndpoint) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3799,7 +3730,6 @@ func (m *awsRestjson1_deserializeOpDeleteEndpoint) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3926,10 +3856,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventStream) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3969,7 +3897,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventStream) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4096,10 +4023,8 @@ func (m *awsRestjson1_deserializeOpDeleteGcmChannel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4139,7 +4064,6 @@ func (m *awsRestjson1_deserializeOpDeleteGcmChannel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4266,10 +4190,8 @@ func (m *awsRestjson1_deserializeOpDeleteInAppTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4309,7 +4231,6 @@ func (m *awsRestjson1_deserializeOpDeleteInAppTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4436,10 +4357,8 @@ func (m *awsRestjson1_deserializeOpDeleteJourney) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4479,7 +4398,6 @@ func (m *awsRestjson1_deserializeOpDeleteJourney) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4606,10 +4524,8 @@ func (m *awsRestjson1_deserializeOpDeletePushTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4649,7 +4565,6 @@ func (m *awsRestjson1_deserializeOpDeletePushTemplate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4776,10 +4691,8 @@ func (m *awsRestjson1_deserializeOpDeleteRecommenderCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4819,7 +4732,6 @@ func (m *awsRestjson1_deserializeOpDeleteRecommenderCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4946,10 +4858,8 @@ func (m *awsRestjson1_deserializeOpDeleteSegment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4989,7 +4899,6 @@ func (m *awsRestjson1_deserializeOpDeleteSegment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5116,10 +5025,8 @@ func (m *awsRestjson1_deserializeOpDeleteSmsChannel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5159,7 +5066,6 @@ func (m *awsRestjson1_deserializeOpDeleteSmsChannel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5286,10 +5192,8 @@ func (m *awsRestjson1_deserializeOpDeleteSmsTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5329,7 +5233,6 @@ func (m *awsRestjson1_deserializeOpDeleteSmsTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5456,10 +5359,8 @@ func (m *awsRestjson1_deserializeOpDeleteUserEndpoints
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5499,7 +5400,6 @@ func (m *awsRestjson1_deserializeOpDeleteUserEndpoints
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5626,10 +5526,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceChannel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5669,7 +5567,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceChannel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5796,10 +5693,8 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5839,7 +5734,6 @@ func (m *awsRestjson1_deserializeOpDeleteVoiceTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5966,10 +5860,8 @@ func (m *awsRestjson1_deserializeOpGetAdmChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6009,7 +5901,6 @@ func (m *awsRestjson1_deserializeOpGetAdmChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6136,10 +6027,8 @@ func (m *awsRestjson1_deserializeOpGetApnsChannel) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6179,7 +6068,6 @@ func (m *awsRestjson1_deserializeOpGetApnsChannel) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6306,10 +6194,8 @@ func (m *awsRestjson1_deserializeOpGetApnsSandboxChann
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6349,7 +6235,6 @@ func (m *awsRestjson1_deserializeOpGetApnsSandboxChann
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6476,10 +6361,8 @@ func (m *awsRestjson1_deserializeOpGetApnsVoipChannel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6519,7 +6402,6 @@ func (m *awsRestjson1_deserializeOpGetApnsVoipChannel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6646,10 +6528,8 @@ func (m *awsRestjson1_deserializeOpGetApnsVoipSandboxC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6689,7 +6569,6 @@ func (m *awsRestjson1_deserializeOpGetApnsVoipSandboxC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6816,10 +6695,8 @@ func (m *awsRestjson1_deserializeOpGetApp) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6859,7 +6736,6 @@ func (m *awsRestjson1_deserializeOpGetApp) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6986,10 +6862,8 @@ func (m *awsRestjson1_deserializeOpGetApplicationDateR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7029,7 +6903,6 @@ func (m *awsRestjson1_deserializeOpGetApplicationDateR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7156,10 +7029,8 @@ func (m *awsRestjson1_deserializeOpGetApplicationSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7199,7 +7070,6 @@ func (m *awsRestjson1_deserializeOpGetApplicationSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7326,10 +7196,8 @@ func (m *awsRestjson1_deserializeOpGetApps) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7369,7 +7237,6 @@ func (m *awsRestjson1_deserializeOpGetApps) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7496,10 +7363,8 @@ func (m *awsRestjson1_deserializeOpGetBaiduChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7539,7 +7404,6 @@ func (m *awsRestjson1_deserializeOpGetBaiduChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7666,10 +7530,8 @@ func (m *awsRestjson1_deserializeOpGetCampaign) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7709,7 +7571,6 @@ func (m *awsRestjson1_deserializeOpGetCampaign) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7836,10 +7697,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignActiviti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7879,7 +7738,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignActiviti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8006,10 +7864,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignDateRang
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8049,7 +7905,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignDateRang
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8176,10 +8031,8 @@ func (m *awsRestjson1_deserializeOpGetCampaigns) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8219,7 +8072,6 @@ func (m *awsRestjson1_deserializeOpGetCampaigns) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8346,10 +8198,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8389,7 +8239,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8516,10 +8365,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignVersions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8559,7 +8406,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignVersions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8686,10 +8532,8 @@ func (m *awsRestjson1_deserializeOpGetChannels) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8729,7 +8573,6 @@ func (m *awsRestjson1_deserializeOpGetChannels) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8856,10 +8699,8 @@ func (m *awsRestjson1_deserializeOpGetEmailChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8899,7 +8740,6 @@ func (m *awsRestjson1_deserializeOpGetEmailChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9026,10 +8866,8 @@ func (m *awsRestjson1_deserializeOpGetEmailTemplate) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9069,7 +8907,6 @@ func (m *awsRestjson1_deserializeOpGetEmailTemplate) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9196,10 +9033,8 @@ func (m *awsRestjson1_deserializeOpGetEndpoint) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9239,7 +9074,6 @@ func (m *awsRestjson1_deserializeOpGetEndpoint) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9366,10 +9200,8 @@ func (m *awsRestjson1_deserializeOpGetEventStream) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9409,7 +9241,6 @@ func (m *awsRestjson1_deserializeOpGetEventStream) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9536,10 +9367,8 @@ func (m *awsRestjson1_deserializeOpGetExportJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9579,7 +9408,6 @@ func (m *awsRestjson1_deserializeOpGetExportJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9706,10 +9534,8 @@ func (m *awsRestjson1_deserializeOpGetExportJobs) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9749,7 +9575,6 @@ func (m *awsRestjson1_deserializeOpGetExportJobs) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9876,10 +9701,8 @@ func (m *awsRestjson1_deserializeOpGetGcmChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9919,7 +9742,6 @@ func (m *awsRestjson1_deserializeOpGetGcmChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10046,10 +9868,8 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10089,7 +9909,6 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10216,10 +10035,8 @@ func (m *awsRestjson1_deserializeOpGetImportJobs) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10259,7 +10076,6 @@ func (m *awsRestjson1_deserializeOpGetImportJobs) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10386,10 +10202,8 @@ func (m *awsRestjson1_deserializeOpGetInAppMessages) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10429,7 +10243,6 @@ func (m *awsRestjson1_deserializeOpGetInAppMessages) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10556,10 +10369,8 @@ func (m *awsRestjson1_deserializeOpGetInAppTemplate) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10599,7 +10410,6 @@ func (m *awsRestjson1_deserializeOpGetInAppTemplate) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10726,10 +10536,8 @@ func (m *awsRestjson1_deserializeOpGetJourney) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10769,7 +10577,6 @@ func (m *awsRestjson1_deserializeOpGetJourney) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10896,10 +10703,8 @@ func (m *awsRestjson1_deserializeOpGetJourneyDateRange
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10939,7 +10744,6 @@ func (m *awsRestjson1_deserializeOpGetJourneyDateRange
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11066,10 +10870,8 @@ func (m *awsRestjson1_deserializeOpGetJourneyExecution
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11109,7 +10911,6 @@ func (m *awsRestjson1_deserializeOpGetJourneyExecution
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11236,10 +11037,8 @@ func (m *awsRestjson1_deserializeOpGetJourneyExecution
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11279,7 +11078,6 @@ func (m *awsRestjson1_deserializeOpGetJourneyExecution
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11406,10 +11204,8 @@ func (m *awsRestjson1_deserializeOpGetJourneyRunExecut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11449,7 +11245,6 @@ func (m *awsRestjson1_deserializeOpGetJourneyRunExecut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11576,10 +11371,8 @@ func (m *awsRestjson1_deserializeOpGetJourneyRunExecut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11619,7 +11412,6 @@ func (m *awsRestjson1_deserializeOpGetJourneyRunExecut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11746,10 +11538,8 @@ func (m *awsRestjson1_deserializeOpGetJourneyRuns) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11789,7 +11579,6 @@ func (m *awsRestjson1_deserializeOpGetJourneyRuns) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11916,10 +11705,8 @@ func (m *awsRestjson1_deserializeOpGetPushTemplate) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11959,7 +11746,6 @@ func (m *awsRestjson1_deserializeOpGetPushTemplate) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12086,10 +11872,8 @@ func (m *awsRestjson1_deserializeOpGetRecommenderConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12129,7 +11913,6 @@ func (m *awsRestjson1_deserializeOpGetRecommenderConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12256,10 +12039,8 @@ func (m *awsRestjson1_deserializeOpGetRecommenderConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12299,7 +12080,6 @@ func (m *awsRestjson1_deserializeOpGetRecommenderConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12426,10 +12206,8 @@ func (m *awsRestjson1_deserializeOpGetSegment) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12469,7 +12247,6 @@ func (m *awsRestjson1_deserializeOpGetSegment) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12596,10 +12373,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentExportJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12639,7 +12414,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentExportJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12766,10 +12540,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentImportJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12809,7 +12581,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentImportJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12936,10 +12707,8 @@ func (m *awsRestjson1_deserializeOpGetSegments) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12979,7 +12748,6 @@ func (m *awsRestjson1_deserializeOpGetSegments) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13106,10 +12874,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13149,7 +12915,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentVersion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13276,10 +13041,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentVersions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13319,7 +13082,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentVersions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13446,10 +13208,8 @@ func (m *awsRestjson1_deserializeOpGetSmsChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13489,7 +13249,6 @@ func (m *awsRestjson1_deserializeOpGetSmsChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13616,10 +13375,8 @@ func (m *awsRestjson1_deserializeOpGetSmsTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13659,7 +13416,6 @@ func (m *awsRestjson1_deserializeOpGetSmsTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13786,10 +13542,8 @@ func (m *awsRestjson1_deserializeOpGetUserEndpoints) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13829,7 +13583,6 @@ func (m *awsRestjson1_deserializeOpGetUserEndpoints) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13956,10 +13709,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13999,7 +13750,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14126,10 +13876,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceTemplate) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14169,7 +13917,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceTemplate) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14296,10 +14043,8 @@ func (m *awsRestjson1_deserializeOpListJourneys) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14339,7 +14084,6 @@ func (m *awsRestjson1_deserializeOpListJourneys) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14466,10 +14210,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14509,7 +14251,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14615,10 +14356,8 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14658,7 +14397,6 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14779,10 +14517,8 @@ func (m *awsRestjson1_deserializeOpListTemplateVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14822,7 +14558,6 @@ func (m *awsRestjson1_deserializeOpListTemplateVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14949,10 +14684,8 @@ func (m *awsRestjson1_deserializeOpPhoneNumberValidate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14992,7 +14725,6 @@ func (m *awsRestjson1_deserializeOpPhoneNumberValidate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15119,10 +14851,8 @@ func (m *awsRestjson1_deserializeOpPutEvents) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15162,7 +14892,6 @@ func (m *awsRestjson1_deserializeOpPutEvents) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15289,10 +15018,8 @@ func (m *awsRestjson1_deserializeOpPutEventStream) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15332,7 +15059,6 @@ func (m *awsRestjson1_deserializeOpPutEventStream) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15459,10 +15185,8 @@ func (m *awsRestjson1_deserializeOpRemoveAttributes) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15502,7 +15226,6 @@ func (m *awsRestjson1_deserializeOpRemoveAttributes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15629,10 +15352,8 @@ func (m *awsRestjson1_deserializeOpSendMessages) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15672,7 +15393,6 @@ func (m *awsRestjson1_deserializeOpSendMessages) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15799,10 +15519,8 @@ func (m *awsRestjson1_deserializeOpSendOTPMessage) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15842,7 +15560,6 @@ func (m *awsRestjson1_deserializeOpSendOTPMessage) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15969,10 +15686,8 @@ func (m *awsRestjson1_deserializeOpSendUsersMessages) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16012,7 +15727,6 @@ func (m *awsRestjson1_deserializeOpSendUsersMessages) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16139,10 +15853,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16160,7 +15872,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16230,10 +15941,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16251,7 +15960,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16321,10 +16029,8 @@ func (m *awsRestjson1_deserializeOpUpdateAdmChannel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16364,7 +16070,6 @@ func (m *awsRestjson1_deserializeOpUpdateAdmChannel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16491,10 +16196,8 @@ func (m *awsRestjson1_deserializeOpUpdateApnsChannel) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16534,7 +16237,6 @@ func (m *awsRestjson1_deserializeOpUpdateApnsChannel) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16661,10 +16363,8 @@ func (m *awsRestjson1_deserializeOpUpdateApnsSandboxCh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16704,7 +16404,6 @@ func (m *awsRestjson1_deserializeOpUpdateApnsSandboxCh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16831,10 +16530,8 @@ func (m *awsRestjson1_deserializeOpUpdateApnsVoipChann
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16874,7 +16571,6 @@ func (m *awsRestjson1_deserializeOpUpdateApnsVoipChann
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17001,10 +16697,8 @@ func (m *awsRestjson1_deserializeOpUpdateApnsVoipSandb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17044,7 +16738,6 @@ func (m *awsRestjson1_deserializeOpUpdateApnsVoipSandb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17171,10 +16864,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplicationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17214,7 +16905,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplicationSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17341,10 +17031,8 @@ func (m *awsRestjson1_deserializeOpUpdateBaiduChannel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17384,7 +17072,6 @@ func (m *awsRestjson1_deserializeOpUpdateBaiduChannel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17511,10 +17198,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaign) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17554,7 +17239,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaign) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17681,10 +17365,8 @@ func (m *awsRestjson1_deserializeOpUpdateEmailChannel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17724,7 +17406,6 @@ func (m *awsRestjson1_deserializeOpUpdateEmailChannel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17851,10 +17532,8 @@ func (m *awsRestjson1_deserializeOpUpdateEmailTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17894,7 +17573,6 @@ func (m *awsRestjson1_deserializeOpUpdateEmailTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18021,10 +17699,8 @@ func (m *awsRestjson1_deserializeOpUpdateEndpoint) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18064,7 +17740,6 @@ func (m *awsRestjson1_deserializeOpUpdateEndpoint) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18191,10 +17866,8 @@ func (m *awsRestjson1_deserializeOpUpdateEndpointsBatc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18234,7 +17907,6 @@ func (m *awsRestjson1_deserializeOpUpdateEndpointsBatc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18361,10 +18033,8 @@ func (m *awsRestjson1_deserializeOpUpdateGcmChannel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18404,7 +18074,6 @@ func (m *awsRestjson1_deserializeOpUpdateGcmChannel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18531,10 +18200,8 @@ func (m *awsRestjson1_deserializeOpUpdateInAppTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18574,7 +18241,6 @@ func (m *awsRestjson1_deserializeOpUpdateInAppTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18701,10 +18367,8 @@ func (m *awsRestjson1_deserializeOpUpdateJourney) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18744,7 +18408,6 @@ func (m *awsRestjson1_deserializeOpUpdateJourney) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18874,10 +18537,8 @@ func (m *awsRestjson1_deserializeOpUpdateJourneyState)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18917,7 +18578,6 @@ func (m *awsRestjson1_deserializeOpUpdateJourneyState)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19044,10 +18704,8 @@ func (m *awsRestjson1_deserializeOpUpdatePushTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19087,7 +18745,6 @@ func (m *awsRestjson1_deserializeOpUpdatePushTemplate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19214,10 +18871,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecommenderCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19257,7 +18912,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecommenderCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19384,10 +19038,8 @@ func (m *awsRestjson1_deserializeOpUpdateSegment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19427,7 +19079,6 @@ func (m *awsRestjson1_deserializeOpUpdateSegment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19554,10 +19205,8 @@ func (m *awsRestjson1_deserializeOpUpdateSmsChannel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19597,7 +19246,6 @@ func (m *awsRestjson1_deserializeOpUpdateSmsChannel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19724,10 +19372,8 @@ func (m *awsRestjson1_deserializeOpUpdateSmsTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19767,7 +19413,6 @@ func (m *awsRestjson1_deserializeOpUpdateSmsTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19894,10 +19539,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplateActiv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19937,7 +19580,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplateActiv
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20064,10 +19706,8 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceChannel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20107,7 +19747,6 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceChannel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20234,10 +19873,8 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20277,7 +19914,6 @@ func (m *awsRestjson1_deserializeOpUpdateVoiceTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20404,10 +20040,8 @@ func (m *awsRestjson1_deserializeOpVerifyOTPMessage) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20447,7 +20081,6 @@ func (m *awsRestjson1_deserializeOpVerifyOTPMessage) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
