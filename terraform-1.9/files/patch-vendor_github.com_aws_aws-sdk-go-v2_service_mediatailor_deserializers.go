--- vendor/github.com/aws/aws-sdk-go-v2/service/mediatailor/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/mediatailor/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpConfigureLogsForCha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpConfigureLogsForCha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -203,10 +199,8 @@ func (m *awsRestjson1_deserializeOpConfigureLogsForPla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -246,7 +240,6 @@ func (m *awsRestjson1_deserializeOpConfigureLogsForPla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -384,10 +377,8 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -427,7 +418,6 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -630,10 +620,8 @@ func (m *awsRestjson1_deserializeOpCreateLiveSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -673,7 +661,6 @@ func (m *awsRestjson1_deserializeOpCreateLiveSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -843,10 +830,8 @@ func (m *awsRestjson1_deserializeOpCreatePrefetchSched
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -886,7 +871,6 @@ func (m *awsRestjson1_deserializeOpCreatePrefetchSched
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1047,10 +1031,8 @@ func (m *awsRestjson1_deserializeOpCreateProgram) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1090,7 +1072,6 @@ func (m *awsRestjson1_deserializeOpCreateProgram) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1305,10 +1286,8 @@ func (m *awsRestjson1_deserializeOpCreateSourceLocatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1348,7 +1327,6 @@ func (m *awsRestjson1_deserializeOpCreateSourceLocatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1524,10 +1502,8 @@ func (m *awsRestjson1_deserializeOpCreateVodSource) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1567,7 +1543,6 @@ func (m *awsRestjson1_deserializeOpCreateVodSource) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1737,10 +1712,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1752,7 +1725,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 	output := &DeleteChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1822,10 +1794,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1837,7 +1807,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelPolicy
 	output := &DeleteChannelPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1907,10 +1876,8 @@ func (m *awsRestjson1_deserializeOpDeleteLiveSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1922,7 +1889,6 @@ func (m *awsRestjson1_deserializeOpDeleteLiveSource) H
 	output := &DeleteLiveSourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1992,10 +1958,8 @@ func (m *awsRestjson1_deserializeOpDeletePlaybackConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2007,7 +1971,6 @@ func (m *awsRestjson1_deserializeOpDeletePlaybackConfi
 	output := &DeletePlaybackConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2077,10 +2040,8 @@ func (m *awsRestjson1_deserializeOpDeletePrefetchSched
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2092,7 +2053,6 @@ func (m *awsRestjson1_deserializeOpDeletePrefetchSched
 	output := &DeletePrefetchScheduleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2162,10 +2122,8 @@ func (m *awsRestjson1_deserializeOpDeleteProgram) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2177,7 +2135,6 @@ func (m *awsRestjson1_deserializeOpDeleteProgram) Hand
 	output := &DeleteProgramOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2247,10 +2204,8 @@ func (m *awsRestjson1_deserializeOpDeleteSourceLocatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2262,7 +2217,6 @@ func (m *awsRestjson1_deserializeOpDeleteSourceLocatio
 	output := &DeleteSourceLocationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2332,10 +2286,8 @@ func (m *awsRestjson1_deserializeOpDeleteVodSource) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2347,7 +2299,6 @@ func (m *awsRestjson1_deserializeOpDeleteVodSource) Ha
 	output := &DeleteVodSourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2417,10 +2368,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2460,7 +2409,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2668,10 +2616,8 @@ func (m *awsRestjson1_deserializeOpDescribeLiveSource)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2711,7 +2657,6 @@ func (m *awsRestjson1_deserializeOpDescribeLiveSource)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2881,10 +2826,8 @@ func (m *awsRestjson1_deserializeOpDescribeProgram) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2924,7 +2867,6 @@ func (m *awsRestjson1_deserializeOpDescribeProgram) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3139,10 +3081,8 @@ func (m *awsRestjson1_deserializeOpDescribeSourceLocat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3182,7 +3122,6 @@ func (m *awsRestjson1_deserializeOpDescribeSourceLocat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3358,10 +3297,8 @@ func (m *awsRestjson1_deserializeOpDescribeVodSource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3401,7 +3338,6 @@ func (m *awsRestjson1_deserializeOpDescribeVodSource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3576,10 +3512,8 @@ func (m *awsRestjson1_deserializeOpGetChannelPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3619,7 +3553,6 @@ func (m *awsRestjson1_deserializeOpGetChannelPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3729,10 +3662,8 @@ func (m *awsRestjson1_deserializeOpGetChannelSchedule)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3772,7 +3703,6 @@ func (m *awsRestjson1_deserializeOpGetChannelSchedule)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3887,10 +3817,8 @@ func (m *awsRestjson1_deserializeOpGetPlaybackConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3930,7 +3858,6 @@ func (m *awsRestjson1_deserializeOpGetPlaybackConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4180,10 +4107,8 @@ func (m *awsRestjson1_deserializeOpGetPrefetchSchedule
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4223,7 +4148,6 @@ func (m *awsRestjson1_deserializeOpGetPrefetchSchedule
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4384,10 +4308,8 @@ func (m *awsRestjson1_deserializeOpListAlerts) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4427,7 +4349,6 @@ func (m *awsRestjson1_deserializeOpListAlerts) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4542,10 +4463,8 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4585,7 +4504,6 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4700,10 +4618,8 @@ func (m *awsRestjson1_deserializeOpListLiveSources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4743,7 +4659,6 @@ func (m *awsRestjson1_deserializeOpListLiveSources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4858,10 +4773,8 @@ func (m *awsRestjson1_deserializeOpListPlaybackConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4901,7 +4814,6 @@ func (m *awsRestjson1_deserializeOpListPlaybackConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5016,10 +4928,8 @@ func (m *awsRestjson1_deserializeOpListPrefetchSchedul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5059,7 +4969,6 @@ func (m *awsRestjson1_deserializeOpListPrefetchSchedul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5174,10 +5083,8 @@ func (m *awsRestjson1_deserializeOpListSourceLocations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5217,7 +5124,6 @@ func (m *awsRestjson1_deserializeOpListSourceLocations
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5332,10 +5238,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5375,7 +5279,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5484,10 +5387,8 @@ func (m *awsRestjson1_deserializeOpListVodSources) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5527,7 +5428,6 @@ func (m *awsRestjson1_deserializeOpListVodSources) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5642,10 +5542,8 @@ func (m *awsRestjson1_deserializeOpPutChannelPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5657,7 +5555,6 @@ func (m *awsRestjson1_deserializeOpPutChannelPolicy) H
 	output := &PutChannelPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5727,10 +5624,8 @@ func (m *awsRestjson1_deserializeOpPutPlaybackConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5770,7 +5665,6 @@ func (m *awsRestjson1_deserializeOpPutPlaybackConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6020,10 +5914,8 @@ func (m *awsRestjson1_deserializeOpStartChannel) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6035,7 +5927,6 @@ func (m *awsRestjson1_deserializeOpStartChannel) Handl
 	output := &StartChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6105,10 +5996,8 @@ func (m *awsRestjson1_deserializeOpStopChannel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6120,7 +6009,6 @@ func (m *awsRestjson1_deserializeOpStopChannel) Handle
 	output := &StopChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6190,10 +6078,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6211,7 +6097,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6284,10 +6169,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6305,7 +6188,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6378,10 +6260,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6421,7 +6301,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6624,10 +6503,8 @@ func (m *awsRestjson1_deserializeOpUpdateLiveSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6667,7 +6544,6 @@ func (m *awsRestjson1_deserializeOpUpdateLiveSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6837,10 +6713,8 @@ func (m *awsRestjson1_deserializeOpUpdateProgram) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6880,7 +6754,6 @@ func (m *awsRestjson1_deserializeOpUpdateProgram) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7095,10 +6968,8 @@ func (m *awsRestjson1_deserializeOpUpdateSourceLocatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7138,7 +7009,6 @@ func (m *awsRestjson1_deserializeOpUpdateSourceLocatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7314,10 +7184,8 @@ func (m *awsRestjson1_deserializeOpUpdateVodSource) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7357,7 +7225,6 @@ func (m *awsRestjson1_deserializeOpUpdateVodSource) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
