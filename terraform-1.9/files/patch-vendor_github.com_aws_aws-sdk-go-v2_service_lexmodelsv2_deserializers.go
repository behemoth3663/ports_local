--- vendor/github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpBatchCreateCustomVo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpBatchCreateCustomVo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -243,10 +239,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteCustomVo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -286,7 +280,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteCustomVo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -439,10 +432,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateCustomVo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -482,7 +473,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateCustomVo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -635,10 +625,8 @@ func (m *awsRestjson1_deserializeOpBuildBotLocale) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -678,7 +666,6 @@ func (m *awsRestjson1_deserializeOpBuildBotLocale) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -849,10 +836,8 @@ func (m *awsRestjson1_deserializeOpCreateBot) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -892,7 +877,6 @@ func (m *awsRestjson1_deserializeOpCreateBot) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1119,10 +1103,8 @@ func (m *awsRestjson1_deserializeOpCreateBotAlias) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1162,7 +1144,6 @@ func (m *awsRestjson1_deserializeOpCreateBotAlias) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1371,10 +1352,8 @@ func (m *awsRestjson1_deserializeOpCreateBotLocale) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1414,7 +1393,6 @@ func (m *awsRestjson1_deserializeOpCreateBotLocale) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1647,10 +1625,8 @@ func (m *awsRestjson1_deserializeOpCreateBotReplica) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1690,7 +1666,6 @@ func (m *awsRestjson1_deserializeOpCreateBotReplica) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1861,10 +1836,8 @@ func (m *awsRestjson1_deserializeOpCreateBotVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1904,7 +1877,6 @@ func (m *awsRestjson1_deserializeOpCreateBotVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2080,10 +2052,8 @@ func (m *awsRestjson1_deserializeOpCreateExport) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2123,7 +2093,6 @@ func (m *awsRestjson1_deserializeOpCreateExport) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2290,10 +2259,8 @@ func (m *awsRestjson1_deserializeOpCreateIntent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2333,7 +2300,6 @@ func (m *awsRestjson1_deserializeOpCreateIntent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2586,10 +2552,8 @@ func (m *awsRestjson1_deserializeOpCreateResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2629,7 +2593,6 @@ func (m *awsRestjson1_deserializeOpCreateResourcePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2766,10 +2729,8 @@ func (m *awsRestjson1_deserializeOpCreateResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2809,7 +2770,6 @@ func (m *awsRestjson1_deserializeOpCreateResourcePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2949,10 +2909,8 @@ func (m *awsRestjson1_deserializeOpCreateSlot) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2992,7 +2950,6 @@ func (m *awsRestjson1_deserializeOpCreateSlot) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3219,10 +3176,8 @@ func (m *awsRestjson1_deserializeOpCreateSlotType) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3262,7 +3217,6 @@ func (m *awsRestjson1_deserializeOpCreateSlotType) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3480,10 +3434,8 @@ func (m *awsRestjson1_deserializeOpCreateTestSetDiscre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3523,7 +3475,6 @@ func (m *awsRestjson1_deserializeOpCreateTestSetDiscre
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3681,10 +3632,8 @@ func (m *awsRestjson1_deserializeOpCreateUploadUrl) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3724,7 +3673,6 @@ func (m *awsRestjson1_deserializeOpCreateUploadUrl) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3858,10 +3806,8 @@ func (m *awsRestjson1_deserializeOpDeleteBot) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3901,7 +3847,6 @@ func (m *awsRestjson1_deserializeOpDeleteBot) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4038,10 +3983,8 @@ func (m *awsRestjson1_deserializeOpDeleteBotAlias) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4081,7 +4024,6 @@ func (m *awsRestjson1_deserializeOpDeleteBotAlias) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4227,10 +4169,8 @@ func (m *awsRestjson1_deserializeOpDeleteBotLocale) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4270,7 +4210,6 @@ func (m *awsRestjson1_deserializeOpDeleteBotLocale) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4425,10 +4364,8 @@ func (m *awsRestjson1_deserializeOpDeleteBotReplica) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4468,7 +4405,6 @@ func (m *awsRestjson1_deserializeOpDeleteBotReplica) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4614,10 +4550,8 @@ func (m *awsRestjson1_deserializeOpDeleteBotVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4657,7 +4591,6 @@ func (m *awsRestjson1_deserializeOpDeleteBotVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4803,10 +4736,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomVocabul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4846,7 +4777,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomVocabul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5001,10 +4931,8 @@ func (m *awsRestjson1_deserializeOpDeleteExport) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5044,7 +4972,6 @@ func (m *awsRestjson1_deserializeOpDeleteExport) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5178,10 +5105,8 @@ func (m *awsRestjson1_deserializeOpDeleteImport) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5221,7 +5146,6 @@ func (m *awsRestjson1_deserializeOpDeleteImport) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5355,10 +5279,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5376,7 +5298,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5464,10 +5385,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5507,7 +5426,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5638,10 +5556,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5681,7 +5597,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5812,10 +5727,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlot) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5833,7 +5746,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlot) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5921,10 +5833,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlotType) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5942,7 +5852,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlotType) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6030,10 +5939,8 @@ func (m *awsRestjson1_deserializeOpDeleteTestSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6051,7 +5958,6 @@ func (m *awsRestjson1_deserializeOpDeleteTestSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6139,10 +6045,8 @@ func (m *awsRestjson1_deserializeOpDeleteUtterances) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6154,7 +6058,6 @@ func (m *awsRestjson1_deserializeOpDeleteUtterances) H
 	output := &DeleteUtterancesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6233,10 +6136,8 @@ func (m *awsRestjson1_deserializeOpDescribeBot) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6276,7 +6177,6 @@ func (m *awsRestjson1_deserializeOpDescribeBot) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6511,10 +6411,8 @@ func (m *awsRestjson1_deserializeOpDescribeBotAlias) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6554,7 +6452,6 @@ func (m *awsRestjson1_deserializeOpDescribeBotAlias) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6781,10 +6678,8 @@ func (m *awsRestjson1_deserializeOpDescribeBotLocale) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6824,7 +6719,6 @@ func (m *awsRestjson1_deserializeOpDescribeBotLocale) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7127,10 +7021,8 @@ func (m *awsRestjson1_deserializeOpDescribeBotRecommen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7170,7 +7062,6 @@ func (m *awsRestjson1_deserializeOpDescribeBotRecommen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7380,10 +7271,8 @@ func (m *awsRestjson1_deserializeOpDescribeBotReplica)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7423,7 +7312,6 @@ func (m *awsRestjson1_deserializeOpDescribeBotReplica)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7596,10 +7484,8 @@ func (m *awsRestjson1_deserializeOpDescribeBotResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7639,7 +7525,6 @@ func (m *awsRestjson1_deserializeOpDescribeBotResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7861,10 +7746,8 @@ func (m *awsRestjson1_deserializeOpDescribeBotVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7904,7 +7787,6 @@ func (m *awsRestjson1_deserializeOpDescribeBotVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8132,10 +8014,8 @@ func (m *awsRestjson1_deserializeOpDescribeCustomVocab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8175,7 +8055,6 @@ func (m *awsRestjson1_deserializeOpDescribeCustomVocab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8359,10 +8238,8 @@ func (m *awsRestjson1_deserializeOpDescribeExport) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8402,7 +8279,6 @@ func (m *awsRestjson1_deserializeOpDescribeExport) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8593,10 +8469,8 @@ func (m *awsRestjson1_deserializeOpDescribeImport) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8636,7 +8510,6 @@ func (m *awsRestjson1_deserializeOpDescribeImport) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8836,10 +8709,8 @@ func (m *awsRestjson1_deserializeOpDescribeIntent) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8879,7 +8750,6 @@ func (m *awsRestjson1_deserializeOpDescribeIntent) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9150,10 +9020,8 @@ func (m *awsRestjson1_deserializeOpDescribeResourcePol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9193,7 +9061,6 @@ func (m *awsRestjson1_deserializeOpDescribeResourcePol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9330,10 +9197,8 @@ func (m *awsRestjson1_deserializeOpDescribeSlot) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9373,7 +9238,6 @@ func (m *awsRestjson1_deserializeOpDescribeSlot) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9613,10 +9477,8 @@ func (m *awsRestjson1_deserializeOpDescribeSlotType) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9656,7 +9518,6 @@ func (m *awsRestjson1_deserializeOpDescribeSlotType) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9887,10 +9748,8 @@ func (m *awsRestjson1_deserializeOpDescribeTestExecuti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9930,7 +9789,6 @@ func (m *awsRestjson1_deserializeOpDescribeTestExecuti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10142,10 +10000,8 @@ func (m *awsRestjson1_deserializeOpDescribeTestSet) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10185,7 +10041,6 @@ func (m *awsRestjson1_deserializeOpDescribeTestSet) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10405,10 +10260,8 @@ func (m *awsRestjson1_deserializeOpDescribeTestSetDisc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10448,7 +10301,6 @@ func (m *awsRestjson1_deserializeOpDescribeTestSetDisc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10647,10 +10499,8 @@ func (m *awsRestjson1_deserializeOpDescribeTestSetGene
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10690,7 +10540,6 @@ func (m *awsRestjson1_deserializeOpDescribeTestSetGene
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10907,10 +10756,8 @@ func (m *awsRestjson1_deserializeOpGenerateBotElement)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10950,7 +10797,6 @@ func (m *awsRestjson1_deserializeOpGenerateBotElement)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11113,10 +10959,8 @@ func (m *awsRestjson1_deserializeOpGetTestExecutionArt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11156,7 +11000,6 @@ func (m *awsRestjson1_deserializeOpGetTestExecutionArt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11290,10 +11133,8 @@ func (m *awsRestjson1_deserializeOpListAggregatedUtter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11333,7 +11174,6 @@ func (m *awsRestjson1_deserializeOpListAggregatedUtter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11549,10 +11389,8 @@ func (m *awsRestjson1_deserializeOpListBotAliases) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11592,7 +11430,6 @@ func (m *awsRestjson1_deserializeOpListBotAliases) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11728,10 +11565,8 @@ func (m *awsRestjson1_deserializeOpListBotAliasReplica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11771,7 +11606,6 @@ func (m *awsRestjson1_deserializeOpListBotAliasReplica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11925,10 +11759,8 @@ func (m *awsRestjson1_deserializeOpListBotLocales) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11968,7 +11800,6 @@ func (m *awsRestjson1_deserializeOpListBotLocales) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12113,10 +11944,8 @@ func (m *awsRestjson1_deserializeOpListBotRecommendati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12156,7 +11985,6 @@ func (m *awsRestjson1_deserializeOpListBotRecommendati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12310,10 +12138,8 @@ func (m *awsRestjson1_deserializeOpListBotReplicas) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12353,7 +12179,6 @@ func (m *awsRestjson1_deserializeOpListBotReplicas) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12489,10 +12314,8 @@ func (m *awsRestjson1_deserializeOpListBotResourceGene
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12532,7 +12355,6 @@ func (m *awsRestjson1_deserializeOpListBotResourceGene
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12686,10 +12508,8 @@ func (m *awsRestjson1_deserializeOpListBots) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12729,7 +12549,6 @@ func (m *awsRestjson1_deserializeOpListBots) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12856,10 +12675,8 @@ func (m *awsRestjson1_deserializeOpListBotVersionRepli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12899,7 +12716,6 @@ func (m *awsRestjson1_deserializeOpListBotVersionRepli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13053,10 +12869,8 @@ func (m *awsRestjson1_deserializeOpListBotVersions) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13096,7 +12910,6 @@ func (m *awsRestjson1_deserializeOpListBotVersions) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13232,10 +13045,8 @@ func (m *awsRestjson1_deserializeOpListBuiltInIntents)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13275,7 +13086,6 @@ func (m *awsRestjson1_deserializeOpListBuiltInIntents)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13411,10 +13221,8 @@ func (m *awsRestjson1_deserializeOpListBuiltInSlotType
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13454,7 +13262,6 @@ func (m *awsRestjson1_deserializeOpListBuiltInSlotType
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13590,10 +13397,8 @@ func (m *awsRestjson1_deserializeOpListCustomVocabular
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13633,7 +13438,6 @@ func (m *awsRestjson1_deserializeOpListCustomVocabular
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13790,10 +13594,8 @@ func (m *awsRestjson1_deserializeOpListExports) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13833,7 +13635,6 @@ func (m *awsRestjson1_deserializeOpListExports) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13984,10 +13785,8 @@ func (m *awsRestjson1_deserializeOpListImports) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14027,7 +13826,6 @@ func (m *awsRestjson1_deserializeOpListImports) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14178,10 +13976,8 @@ func (m *awsRestjson1_deserializeOpListIntentMetrics) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14221,7 +14017,6 @@ func (m *awsRestjson1_deserializeOpListIntentMetrics) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14360,10 +14155,8 @@ func (m *awsRestjson1_deserializeOpListIntentPaths) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14403,7 +14196,6 @@ func (m *awsRestjson1_deserializeOpListIntentPaths) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14524,10 +14316,8 @@ func (m *awsRestjson1_deserializeOpListIntents) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14567,7 +14357,6 @@ func (m *awsRestjson1_deserializeOpListIntents) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14721,10 +14510,8 @@ func (m *awsRestjson1_deserializeOpListIntentStageMetr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14764,7 +14551,6 @@ func (m *awsRestjson1_deserializeOpListIntentStageMetr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14903,10 +14689,8 @@ func (m *awsRestjson1_deserializeOpListRecommendedInte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14946,7 +14730,6 @@ func (m *awsRestjson1_deserializeOpListRecommendedInte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15112,10 +14895,8 @@ func (m *awsRestjson1_deserializeOpListSessionAnalytic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15155,7 +14936,6 @@ func (m *awsRestjson1_deserializeOpListSessionAnalytic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15294,10 +15074,8 @@ func (m *awsRestjson1_deserializeOpListSessionMetrics)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15337,7 +15115,6 @@ func (m *awsRestjson1_deserializeOpListSessionMetrics)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15476,10 +15253,8 @@ func (m *awsRestjson1_deserializeOpListSlots) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15519,7 +15294,6 @@ func (m *awsRestjson1_deserializeOpListSlots) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15682,10 +15456,8 @@ func (m *awsRestjson1_deserializeOpListSlotTypes) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15725,7 +15497,6 @@ func (m *awsRestjson1_deserializeOpListSlotTypes) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15879,10 +15650,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15922,7 +15691,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16040,10 +15808,8 @@ func (m *awsRestjson1_deserializeOpListTestExecutionRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16083,7 +15849,6 @@ func (m *awsRestjson1_deserializeOpListTestExecutionRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16213,10 +15978,8 @@ func (m *awsRestjson1_deserializeOpListTestExecutions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16256,7 +16019,6 @@ func (m *awsRestjson1_deserializeOpListTestExecutions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16383,10 +16145,8 @@ func (m *awsRestjson1_deserializeOpListTestSetRecords)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16426,7 +16186,6 @@ func (m *awsRestjson1_deserializeOpListTestSetRecords)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16556,10 +16315,8 @@ func (m *awsRestjson1_deserializeOpListTestSets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16599,7 +16356,6 @@ func (m *awsRestjson1_deserializeOpListTestSets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16726,10 +16482,8 @@ func (m *awsRestjson1_deserializeOpListUtteranceAnalyt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16769,7 +16523,6 @@ func (m *awsRestjson1_deserializeOpListUtteranceAnalyt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16908,10 +16661,8 @@ func (m *awsRestjson1_deserializeOpListUtteranceMetric
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16951,7 +16702,6 @@ func (m *awsRestjson1_deserializeOpListUtteranceMetric
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17090,10 +16840,8 @@ func (m *awsRestjson1_deserializeOpSearchAssociatedTra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17133,7 +16881,6 @@ func (m *awsRestjson1_deserializeOpSearchAssociatedTra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17316,10 +17063,8 @@ func (m *awsRestjson1_deserializeOpStartBotRecommendat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17359,7 +17104,6 @@ func (m *awsRestjson1_deserializeOpStartBotRecommendat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17552,10 +17296,8 @@ func (m *awsRestjson1_deserializeOpStartBotResourceGen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17595,7 +17337,6 @@ func (m *awsRestjson1_deserializeOpStartBotResourceGen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17784,10 +17525,8 @@ func (m *awsRestjson1_deserializeOpStartImport) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17827,7 +17566,6 @@ func (m *awsRestjson1_deserializeOpStartImport) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17994,10 +17732,8 @@ func (m *awsRestjson1_deserializeOpStartTestExecution)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18037,7 +17773,6 @@ func (m *awsRestjson1_deserializeOpStartTestExecution)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18213,10 +17948,8 @@ func (m *awsRestjson1_deserializeOpStartTestSetGenerat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18256,7 +17989,6 @@ func (m *awsRestjson1_deserializeOpStartTestSetGenerat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18451,10 +18183,8 @@ func (m *awsRestjson1_deserializeOpStopBotRecommendati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18494,7 +18224,6 @@ func (m *awsRestjson1_deserializeOpStopBotRecommendati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18661,10 +18390,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18676,7 +18403,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18758,10 +18484,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18773,7 +18497,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18855,10 +18578,8 @@ func (m *awsRestjson1_deserializeOpUpdateBot) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18898,7 +18619,6 @@ func (m *awsRestjson1_deserializeOpUpdateBot) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19131,10 +18851,8 @@ func (m *awsRestjson1_deserializeOpUpdateBotAlias) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19174,7 +18892,6 @@ func (m *awsRestjson1_deserializeOpUpdateBotAlias) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19394,10 +19111,8 @@ func (m *awsRestjson1_deserializeOpUpdateBotLocale) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19437,7 +19152,6 @@ func (m *awsRestjson1_deserializeOpUpdateBotLocale) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19696,10 +19410,8 @@ func (m *awsRestjson1_deserializeOpUpdateBotRecommenda
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19739,7 +19451,6 @@ func (m *awsRestjson1_deserializeOpUpdateBotRecommenda
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19948,10 +19659,8 @@ func (m *awsRestjson1_deserializeOpUpdateExport) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19991,7 +19700,6 @@ func (m *awsRestjson1_deserializeOpUpdateExport) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20174,10 +19882,8 @@ func (m *awsRestjson1_deserializeOpUpdateIntent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20217,7 +19923,6 @@ func (m *awsRestjson1_deserializeOpUpdateIntent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20491,10 +20196,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20534,7 +20237,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourcePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20671,10 +20373,8 @@ func (m *awsRestjson1_deserializeOpUpdateSlot) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20714,7 +20414,6 @@ func (m *awsRestjson1_deserializeOpUpdateSlot) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20957,10 +20656,8 @@ func (m *awsRestjson1_deserializeOpUpdateSlotType) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21000,7 +20697,6 @@ func (m *awsRestjson1_deserializeOpUpdateSlotType) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21234,10 +20930,8 @@ func (m *awsRestjson1_deserializeOpUpdateTestSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21277,7 +20971,6 @@ func (m *awsRestjson1_deserializeOpUpdateTestSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
