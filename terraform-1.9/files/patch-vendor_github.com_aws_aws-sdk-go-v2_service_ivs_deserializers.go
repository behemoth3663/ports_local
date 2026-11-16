--- vendor/github.com/aws/aws-sdk-go-v2/service/ivs/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ivs/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchGetChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchGetChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -199,10 +195,8 @@ func (m *awsRestjson1_deserializeOpBatchGetStreamKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -242,7 +236,6 @@ func (m *awsRestjson1_deserializeOpBatchGetStreamKey) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -353,10 +346,8 @@ func (m *awsRestjson1_deserializeOpBatchStartViewerSes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -396,7 +387,6 @@ func (m *awsRestjson1_deserializeOpBatchStartViewerSes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -514,10 +504,8 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -557,7 +545,6 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -683,10 +670,8 @@ func (m *awsRestjson1_deserializeOpCreatePlaybackRestr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -726,7 +711,6 @@ func (m *awsRestjson1_deserializeOpCreatePlaybackRestr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -847,10 +831,8 @@ func (m *awsRestjson1_deserializeOpCreateRecordingConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -890,7 +872,6 @@ func (m *awsRestjson1_deserializeOpCreateRecordingConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1014,10 +995,8 @@ func (m *awsRestjson1_deserializeOpCreateStreamKey) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1057,7 +1036,6 @@ func (m *awsRestjson1_deserializeOpCreateStreamKey) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1178,10 +1156,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1199,7 +1175,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1284,10 +1259,8 @@ func (m *awsRestjson1_deserializeOpDeletePlaybackKeyPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1299,7 +1272,6 @@ func (m *awsRestjson1_deserializeOpDeletePlaybackKeyPa
 	output := &DeletePlaybackKeyPairOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1381,10 +1353,8 @@ func (m *awsRestjson1_deserializeOpDeletePlaybackRestr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1402,7 +1372,6 @@ func (m *awsRestjson1_deserializeOpDeletePlaybackRestr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1487,10 +1456,8 @@ func (m *awsRestjson1_deserializeOpDeleteRecordingConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1508,7 +1475,6 @@ func (m *awsRestjson1_deserializeOpDeleteRecordingConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1593,10 +1559,8 @@ func (m *awsRestjson1_deserializeOpDeleteStreamKey) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1614,7 +1578,6 @@ func (m *awsRestjson1_deserializeOpDeleteStreamKey) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1696,10 +1659,8 @@ func (m *awsRestjson1_deserializeOpGetChannel) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1739,7 +1700,6 @@ func (m *awsRestjson1_deserializeOpGetChannel) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1854,10 +1814,8 @@ func (m *awsRestjson1_deserializeOpGetPlaybackKeyPair)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1897,7 +1855,6 @@ func (m *awsRestjson1_deserializeOpGetPlaybackKeyPair)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2012,10 +1969,8 @@ func (m *awsRestjson1_deserializeOpGetPlaybackRestrict
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2055,7 +2010,6 @@ func (m *awsRestjson1_deserializeOpGetPlaybackRestrict
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2173,10 +2127,8 @@ func (m *awsRestjson1_deserializeOpGetRecordingConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2216,7 +2168,6 @@ func (m *awsRestjson1_deserializeOpGetRecordingConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2334,10 +2285,8 @@ func (m *awsRestjson1_deserializeOpGetStream) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2377,7 +2326,6 @@ func (m *awsRestjson1_deserializeOpGetStream) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2495,10 +2443,8 @@ func (m *awsRestjson1_deserializeOpGetStreamKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2538,7 +2484,6 @@ func (m *awsRestjson1_deserializeOpGetStreamKey) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2653,10 +2598,8 @@ func (m *awsRestjson1_deserializeOpGetStreamSession) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2696,7 +2639,6 @@ func (m *awsRestjson1_deserializeOpGetStreamSession) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2811,10 +2753,8 @@ func (m *awsRestjson1_deserializeOpImportPlaybackKeyPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2854,7 +2794,6 @@ func (m *awsRestjson1_deserializeOpImportPlaybackKeyPa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2975,10 +2914,8 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3018,7 +2955,6 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3142,10 +3078,8 @@ func (m *awsRestjson1_deserializeOpListPlaybackKeyPair
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3185,7 +3119,6 @@ func (m *awsRestjson1_deserializeOpListPlaybackKeyPair
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3306,10 +3239,8 @@ func (m *awsRestjson1_deserializeOpListPlaybackRestric
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3349,7 +3280,6 @@ func (m *awsRestjson1_deserializeOpListPlaybackRestric
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3476,10 +3406,8 @@ func (m *awsRestjson1_deserializeOpListRecordingConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3519,7 +3447,6 @@ func (m *awsRestjson1_deserializeOpListRecordingConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3643,10 +3570,8 @@ func (m *awsRestjson1_deserializeOpListStreamKeys) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3686,7 +3611,6 @@ func (m *awsRestjson1_deserializeOpListStreamKeys) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3810,10 +3734,8 @@ func (m *awsRestjson1_deserializeOpListStreams) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3853,7 +3775,6 @@ func (m *awsRestjson1_deserializeOpListStreams) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3974,10 +3895,8 @@ func (m *awsRestjson1_deserializeOpListStreamSessions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4017,7 +3936,6 @@ func (m *awsRestjson1_deserializeOpListStreamSessions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4141,10 +4059,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4184,7 +4100,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4299,10 +4214,8 @@ func (m *awsRestjson1_deserializeOpPutMetadata) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4320,7 +4233,6 @@ func (m *awsRestjson1_deserializeOpPutMetadata) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4405,10 +4317,8 @@ func (m *awsRestjson1_deserializeOpStartViewerSessionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4420,7 +4330,6 @@ func (m *awsRestjson1_deserializeOpStartViewerSessionR
 	output := &StartViewerSessionRevocationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4508,10 +4417,8 @@ func (m *awsRestjson1_deserializeOpStopStream) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4523,7 +4430,6 @@ func (m *awsRestjson1_deserializeOpStopStream) HandleD
 	output := &StopStreamOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4608,10 +4514,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4623,7 +4527,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4702,10 +4605,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4717,7 +4618,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4796,10 +4696,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4839,7 +4737,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4960,10 +4857,8 @@ func (m *awsRestjson1_deserializeOpUpdatePlaybackRestr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5003,7 +4898,6 @@ func (m *awsRestjson1_deserializeOpUpdatePlaybackRestr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
