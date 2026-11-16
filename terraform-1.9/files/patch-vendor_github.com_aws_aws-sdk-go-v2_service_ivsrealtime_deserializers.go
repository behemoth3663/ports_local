--- vendor/github.com/aws/aws-sdk-go-v2/service/ivsrealtime/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ivsrealtime/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateEncoderConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateEncoderConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -215,10 +211,8 @@ func (m *awsRestjson1_deserializeOpCreateIngestConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -258,7 +252,6 @@ func (m *awsRestjson1_deserializeOpCreateIngestConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -376,10 +369,8 @@ func (m *awsRestjson1_deserializeOpCreateParticipantTo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -419,7 +410,6 @@ func (m *awsRestjson1_deserializeOpCreateParticipantTo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -540,10 +530,8 @@ func (m *awsRestjson1_deserializeOpCreateStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -583,7 +571,6 @@ func (m *awsRestjson1_deserializeOpCreateStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -706,10 +693,8 @@ func (m *awsRestjson1_deserializeOpCreateStorageConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -749,7 +734,6 @@ func (m *awsRestjson1_deserializeOpCreateStorageConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -876,10 +860,8 @@ func (m *awsRestjson1_deserializeOpDeleteEncoderConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -891,7 +873,6 @@ func (m *awsRestjson1_deserializeOpDeleteEncoderConfig
 	output := &DeleteEncoderConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -979,10 +960,8 @@ func (m *awsRestjson1_deserializeOpDeleteIngestConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -994,7 +973,6 @@ func (m *awsRestjson1_deserializeOpDeleteIngestConfigu
 	output := &DeleteIngestConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1079,10 +1057,8 @@ func (m *awsRestjson1_deserializeOpDeletePublicKey) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1094,7 +1070,6 @@ func (m *awsRestjson1_deserializeOpDeletePublicKey) Ha
 	output := &DeletePublicKeyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1179,10 +1154,8 @@ func (m *awsRestjson1_deserializeOpDeleteStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1194,7 +1167,6 @@ func (m *awsRestjson1_deserializeOpDeleteStage) Handle
 	output := &DeleteStageOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1279,10 +1251,8 @@ func (m *awsRestjson1_deserializeOpDeleteStorageConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1294,7 +1264,6 @@ func (m *awsRestjson1_deserializeOpDeleteStorageConfig
 	output := &DeleteStorageConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1382,10 +1351,8 @@ func (m *awsRestjson1_deserializeOpDisconnectParticipa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1397,7 +1364,6 @@ func (m *awsRestjson1_deserializeOpDisconnectParticipa
 	output := &DisconnectParticipantOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1479,10 +1445,8 @@ func (m *awsRestjson1_deserializeOpGetComposition) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1522,7 +1486,6 @@ func (m *awsRestjson1_deserializeOpGetComposition) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1646,10 +1609,8 @@ func (m *awsRestjson1_deserializeOpGetEncoderConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1689,7 +1650,6 @@ func (m *awsRestjson1_deserializeOpGetEncoderConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1813,10 +1773,8 @@ func (m *awsRestjson1_deserializeOpGetIngestConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1856,7 +1814,6 @@ func (m *awsRestjson1_deserializeOpGetIngestConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1971,10 +1928,8 @@ func (m *awsRestjson1_deserializeOpGetParticipant) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2014,7 +1969,6 @@ func (m *awsRestjson1_deserializeOpGetParticipant) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2129,10 +2083,8 @@ func (m *awsRestjson1_deserializeOpGetPublicKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2172,7 +2124,6 @@ func (m *awsRestjson1_deserializeOpGetPublicKey) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2287,10 +2238,8 @@ func (m *awsRestjson1_deserializeOpGetStage) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2330,7 +2279,6 @@ func (m *awsRestjson1_deserializeOpGetStage) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2445,10 +2393,8 @@ func (m *awsRestjson1_deserializeOpGetStageSession) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2488,7 +2434,6 @@ func (m *awsRestjson1_deserializeOpGetStageSession) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2603,10 +2548,8 @@ func (m *awsRestjson1_deserializeOpGetStorageConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2646,7 +2589,6 @@ func (m *awsRestjson1_deserializeOpGetStorageConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2770,10 +2712,8 @@ func (m *awsRestjson1_deserializeOpImportPublicKey) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2813,7 +2753,6 @@ func (m *awsRestjson1_deserializeOpImportPublicKey) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2934,10 +2873,8 @@ func (m *awsRestjson1_deserializeOpListCompositions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2977,7 +2914,6 @@ func (m *awsRestjson1_deserializeOpListCompositions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3107,10 +3043,8 @@ func (m *awsRestjson1_deserializeOpListEncoderConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3150,7 +3084,6 @@ func (m *awsRestjson1_deserializeOpListEncoderConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3280,10 +3213,8 @@ func (m *awsRestjson1_deserializeOpListIngestConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3323,7 +3254,6 @@ func (m *awsRestjson1_deserializeOpListIngestConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3444,10 +3374,8 @@ func (m *awsRestjson1_deserializeOpListParticipantEven
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3487,7 +3415,6 @@ func (m *awsRestjson1_deserializeOpListParticipantEven
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3608,10 +3535,8 @@ func (m *awsRestjson1_deserializeOpListParticipantRepl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3651,7 +3576,6 @@ func (m *awsRestjson1_deserializeOpListParticipantRepl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3772,10 +3696,8 @@ func (m *awsRestjson1_deserializeOpListParticipants) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3815,7 +3737,6 @@ func (m *awsRestjson1_deserializeOpListParticipants) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3936,10 +3857,8 @@ func (m *awsRestjson1_deserializeOpListPublicKeys) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3979,7 +3898,6 @@ func (m *awsRestjson1_deserializeOpListPublicKeys) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4100,10 +4018,8 @@ func (m *awsRestjson1_deserializeOpListStages) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4143,7 +4059,6 @@ func (m *awsRestjson1_deserializeOpListStages) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4267,10 +4182,8 @@ func (m *awsRestjson1_deserializeOpListStageSessions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4310,7 +4223,6 @@ func (m *awsRestjson1_deserializeOpListStageSessions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4431,10 +4343,8 @@ func (m *awsRestjson1_deserializeOpListStorageConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4474,7 +4384,6 @@ func (m *awsRestjson1_deserializeOpListStorageConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4604,10 +4513,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4647,7 +4554,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4762,10 +4668,8 @@ func (m *awsRestjson1_deserializeOpStartComposition) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4805,7 +4709,6 @@ func (m *awsRestjson1_deserializeOpStartComposition) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4932,10 +4835,8 @@ func (m *awsRestjson1_deserializeOpStartParticipantRep
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4952,7 +4853,6 @@ func (m *awsRestjson1_deserializeOpStartParticipantRep
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5086,10 +4986,8 @@ func (m *awsRestjson1_deserializeOpStopComposition) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5101,7 +4999,6 @@ func (m *awsRestjson1_deserializeOpStopComposition) Ha
 	output := &StopCompositionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5189,10 +5086,8 @@ func (m *awsRestjson1_deserializeOpStopParticipantRepl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5209,7 +5104,6 @@ func (m *awsRestjson1_deserializeOpStopParticipantRepl
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5334,10 +5228,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5349,7 +5241,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5428,10 +5319,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5443,7 +5332,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5522,10 +5410,8 @@ func (m *awsRestjson1_deserializeOpUpdateIngestConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5565,7 +5451,6 @@ func (m *awsRestjson1_deserializeOpUpdateIngestConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5686,10 +5571,8 @@ func (m *awsRestjson1_deserializeOpUpdateStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5729,7 +5612,6 @@ func (m *awsRestjson1_deserializeOpUpdateStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
