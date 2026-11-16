--- vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateMediaCaptureP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateMediaCaptureP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -215,10 +211,8 @@ func (m *awsRestjson1_deserializeOpCreateMediaConcaten
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -258,7 +252,6 @@ func (m *awsRestjson1_deserializeOpCreateMediaConcaten
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -385,10 +378,8 @@ func (m *awsRestjson1_deserializeOpCreateMediaInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -428,7 +419,6 @@ func (m *awsRestjson1_deserializeOpCreateMediaInsights
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -558,10 +548,8 @@ func (m *awsRestjson1_deserializeOpCreateMediaInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -601,7 +589,6 @@ func (m *awsRestjson1_deserializeOpCreateMediaInsights
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -731,10 +718,8 @@ func (m *awsRestjson1_deserializeOpCreateMediaLiveConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -774,7 +759,6 @@ func (m *awsRestjson1_deserializeOpCreateMediaLiveConn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -901,10 +885,8 @@ func (m *awsRestjson1_deserializeOpCreateMediaPipeline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -944,7 +926,6 @@ func (m *awsRestjson1_deserializeOpCreateMediaPipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1074,10 +1055,8 @@ func (m *awsRestjson1_deserializeOpCreateMediaStreamPi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1117,7 +1096,6 @@ func (m *awsRestjson1_deserializeOpCreateMediaStreamPi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1247,10 +1225,8 @@ func (m *awsRestjson1_deserializeOpDeleteMediaCaptureP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1268,7 +1244,6 @@ func (m *awsRestjson1_deserializeOpDeleteMediaCaptureP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1359,10 +1334,8 @@ func (m *awsRestjson1_deserializeOpDeleteMediaInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1380,7 +1353,6 @@ func (m *awsRestjson1_deserializeOpDeleteMediaInsights
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1474,10 +1446,8 @@ func (m *awsRestjson1_deserializeOpDeleteMediaPipeline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1495,7 +1465,6 @@ func (m *awsRestjson1_deserializeOpDeleteMediaPipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1589,10 +1558,8 @@ func (m *awsRestjson1_deserializeOpDeleteMediaPipeline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1610,7 +1577,6 @@ func (m *awsRestjson1_deserializeOpDeleteMediaPipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1704,10 +1670,8 @@ func (m *awsRestjson1_deserializeOpGetMediaCapturePipe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1747,7 +1711,6 @@ func (m *awsRestjson1_deserializeOpGetMediaCapturePipe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1874,10 +1837,8 @@ func (m *awsRestjson1_deserializeOpGetMediaInsightsPip
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1917,7 +1878,6 @@ func (m *awsRestjson1_deserializeOpGetMediaInsightsPip
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2044,10 +2004,8 @@ func (m *awsRestjson1_deserializeOpGetMediaPipeline) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2087,7 +2045,6 @@ func (m *awsRestjson1_deserializeOpGetMediaPipeline) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2214,10 +2171,8 @@ func (m *awsRestjson1_deserializeOpGetMediaPipelineKin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2257,7 +2212,6 @@ func (m *awsRestjson1_deserializeOpGetMediaPipelineKin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2384,10 +2338,8 @@ func (m *awsRestjson1_deserializeOpGetSpeakerSearchTas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2427,7 +2379,6 @@ func (m *awsRestjson1_deserializeOpGetSpeakerSearchTas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2554,10 +2505,8 @@ func (m *awsRestjson1_deserializeOpGetVoiceToneAnalysi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2597,7 +2546,6 @@ func (m *awsRestjson1_deserializeOpGetVoiceToneAnalysi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2724,10 +2672,8 @@ func (m *awsRestjson1_deserializeOpListMediaCapturePip
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2767,7 +2713,6 @@ func (m *awsRestjson1_deserializeOpListMediaCapturePip
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2903,10 +2848,8 @@ func (m *awsRestjson1_deserializeOpListMediaInsightsPi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2946,7 +2889,6 @@ func (m *awsRestjson1_deserializeOpListMediaInsightsPi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3082,10 +3024,8 @@ func (m *awsRestjson1_deserializeOpListMediaPipelineKi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3125,7 +3065,6 @@ func (m *awsRestjson1_deserializeOpListMediaPipelineKi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3261,10 +3200,8 @@ func (m *awsRestjson1_deserializeOpListMediaPipelines)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3304,7 +3241,6 @@ func (m *awsRestjson1_deserializeOpListMediaPipelines)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3440,10 +3376,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3483,7 +3417,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3610,10 +3543,8 @@ func (m *awsRestjson1_deserializeOpStartSpeakerSearchT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3653,7 +3584,6 @@ func (m *awsRestjson1_deserializeOpStartSpeakerSearchT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3783,10 +3713,8 @@ func (m *awsRestjson1_deserializeOpStartVoiceToneAnaly
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3826,7 +3754,6 @@ func (m *awsRestjson1_deserializeOpStartVoiceToneAnaly
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3956,10 +3883,8 @@ func (m *awsRestjson1_deserializeOpStopSpeakerSearchTa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3977,7 +3902,6 @@ func (m *awsRestjson1_deserializeOpStopSpeakerSearchTa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4071,10 +3995,8 @@ func (m *awsRestjson1_deserializeOpStopVoiceToneAnalys
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4092,7 +4014,6 @@ func (m *awsRestjson1_deserializeOpStopVoiceToneAnalys
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4186,10 +4107,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4201,7 +4120,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4292,10 +4210,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4307,7 +4223,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4398,10 +4313,8 @@ func (m *awsRestjson1_deserializeOpUpdateMediaInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4441,7 +4354,6 @@ func (m *awsRestjson1_deserializeOpUpdateMediaInsights
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4571,10 +4483,8 @@ func (m *awsRestjson1_deserializeOpUpdateMediaInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4592,7 +4502,6 @@ func (m *awsRestjson1_deserializeOpUpdateMediaInsights
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4686,10 +4595,8 @@ func (m *awsRestjson1_deserializeOpUpdateMediaPipeline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4729,7 +4636,6 @@ func (m *awsRestjson1_deserializeOpUpdateMediaPipeline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
