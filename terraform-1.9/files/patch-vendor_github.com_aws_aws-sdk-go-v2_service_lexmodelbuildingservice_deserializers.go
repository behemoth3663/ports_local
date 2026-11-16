--- vendor/github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateBotVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateBotVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -367,10 +363,8 @@ func (m *awsRestjson1_deserializeOpCreateIntentVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -410,7 +404,6 @@ func (m *awsRestjson1_deserializeOpCreateIntentVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -661,10 +654,8 @@ func (m *awsRestjson1_deserializeOpCreateSlotTypeVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -704,7 +695,6 @@ func (m *awsRestjson1_deserializeOpCreateSlotTypeVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -919,10 +909,8 @@ func (m *awsRestjson1_deserializeOpDeleteBot) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -940,7 +928,6 @@ func (m *awsRestjson1_deserializeOpDeleteBot) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1028,10 +1015,8 @@ func (m *awsRestjson1_deserializeOpDeleteBotAlias) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1049,7 +1034,6 @@ func (m *awsRestjson1_deserializeOpDeleteBotAlias) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1137,10 +1121,8 @@ func (m *awsRestjson1_deserializeOpDeleteBotChannelAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1158,7 +1140,6 @@ func (m *awsRestjson1_deserializeOpDeleteBotChannelAss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1243,10 +1224,8 @@ func (m *awsRestjson1_deserializeOpDeleteBotVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1264,7 +1243,6 @@ func (m *awsRestjson1_deserializeOpDeleteBotVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1352,10 +1330,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1373,7 +1349,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1461,10 +1436,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntentVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1482,7 +1455,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntentVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1570,10 +1542,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlotType) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1591,7 +1561,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlotType) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1679,10 +1648,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlotTypeVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1700,7 +1667,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlotTypeVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1788,10 +1754,8 @@ func (m *awsRestjson1_deserializeOpDeleteUtterances) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1809,7 +1773,6 @@ func (m *awsRestjson1_deserializeOpDeleteUtterances) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1891,10 +1854,8 @@ func (m *awsRestjson1_deserializeOpGetBot) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1934,7 +1895,6 @@ func (m *awsRestjson1_deserializeOpGetBot) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2240,10 +2200,8 @@ func (m *awsRestjson1_deserializeOpGetBotAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2283,7 +2241,6 @@ func (m *awsRestjson1_deserializeOpGetBotAlias) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2478,10 +2435,8 @@ func (m *awsRestjson1_deserializeOpGetBotAliases) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2521,7 +2476,6 @@ func (m *awsRestjson1_deserializeOpGetBotAliases) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2645,10 +2599,8 @@ func (m *awsRestjson1_deserializeOpGetBotChannelAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2688,7 +2640,6 @@ func (m *awsRestjson1_deserializeOpGetBotChannelAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2885,10 +2836,8 @@ func (m *awsRestjson1_deserializeOpGetBotChannelAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2928,7 +2877,6 @@ func (m *awsRestjson1_deserializeOpGetBotChannelAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3052,10 +3000,8 @@ func (m *awsRestjson1_deserializeOpGetBots) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3095,7 +3041,6 @@ func (m *awsRestjson1_deserializeOpGetBots) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3222,10 +3167,8 @@ func (m *awsRestjson1_deserializeOpGetBotVersions) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3265,7 +3208,6 @@ func (m *awsRestjson1_deserializeOpGetBotVersions) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3392,10 +3334,8 @@ func (m *awsRestjson1_deserializeOpGetBuiltinIntent) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3435,7 +3375,6 @@ func (m *awsRestjson1_deserializeOpGetBuiltinIntent) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3567,10 +3506,8 @@ func (m *awsRestjson1_deserializeOpGetBuiltinIntents) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3610,7 +3547,6 @@ func (m *awsRestjson1_deserializeOpGetBuiltinIntents) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3734,10 +3670,8 @@ func (m *awsRestjson1_deserializeOpGetBuiltinSlotTypes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3777,7 +3711,6 @@ func (m *awsRestjson1_deserializeOpGetBuiltinSlotTypes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3901,10 +3834,8 @@ func (m *awsRestjson1_deserializeOpGetExport) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3944,7 +3875,6 @@ func (m *awsRestjson1_deserializeOpGetExport) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4120,10 +4050,8 @@ func (m *awsRestjson1_deserializeOpGetImport) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4163,7 +4091,6 @@ func (m *awsRestjson1_deserializeOpGetImport) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4342,10 +4269,8 @@ func (m *awsRestjson1_deserializeOpGetIntent) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4385,7 +4310,6 @@ func (m *awsRestjson1_deserializeOpGetIntent) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4630,10 +4554,8 @@ func (m *awsRestjson1_deserializeOpGetIntents) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4673,7 +4595,6 @@ func (m *awsRestjson1_deserializeOpGetIntents) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4800,10 +4721,8 @@ func (m *awsRestjson1_deserializeOpGetIntentVersions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4843,7 +4762,6 @@ func (m *awsRestjson1_deserializeOpGetIntentVersions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4970,10 +4888,8 @@ func (m *awsRestjson1_deserializeOpGetMigration) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5013,7 +4929,6 @@ func (m *awsRestjson1_deserializeOpGetMigration) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5219,10 +5134,8 @@ func (m *awsRestjson1_deserializeOpGetMigrations) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5262,7 +5175,6 @@ func (m *awsRestjson1_deserializeOpGetMigrations) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5386,10 +5298,8 @@ func (m *awsRestjson1_deserializeOpGetSlotType) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5429,7 +5339,6 @@ func (m *awsRestjson1_deserializeOpGetSlotType) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5638,10 +5547,8 @@ func (m *awsRestjson1_deserializeOpGetSlotTypes) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5681,7 +5588,6 @@ func (m *awsRestjson1_deserializeOpGetSlotTypes) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5808,10 +5714,8 @@ func (m *awsRestjson1_deserializeOpGetSlotTypeVersions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5851,7 +5755,6 @@ func (m *awsRestjson1_deserializeOpGetSlotTypeVersions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5978,10 +5881,8 @@ func (m *awsRestjson1_deserializeOpGetUtterancesView) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6021,7 +5922,6 @@ func (m *awsRestjson1_deserializeOpGetUtterancesView) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6145,10 +6045,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6188,7 +6086,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6306,10 +6203,8 @@ func (m *awsRestjson1_deserializeOpPutBot) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6349,7 +6244,6 @@ func (m *awsRestjson1_deserializeOpPutBot) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6672,10 +6566,8 @@ func (m *awsRestjson1_deserializeOpPutBotAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6715,7 +6607,6 @@ func (m *awsRestjson1_deserializeOpPutBotAlias) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6918,10 +6809,8 @@ func (m *awsRestjson1_deserializeOpPutIntent) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6961,7 +6850,6 @@ func (m *awsRestjson1_deserializeOpPutIntent) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7218,10 +7106,8 @@ func (m *awsRestjson1_deserializeOpPutSlotType) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7261,7 +7147,6 @@ func (m *awsRestjson1_deserializeOpPutSlotType) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7482,10 +7367,8 @@ func (m *awsRestjson1_deserializeOpStartImport) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7525,7 +7408,6 @@ func (m *awsRestjson1_deserializeOpStartImport) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7701,10 +7583,8 @@ func (m *awsRestjson1_deserializeOpStartMigration) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7744,7 +7624,6 @@ func (m *awsRestjson1_deserializeOpStartMigration) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7939,10 +7818,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7954,7 +7831,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8039,10 +7915,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8054,7 +7928,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
