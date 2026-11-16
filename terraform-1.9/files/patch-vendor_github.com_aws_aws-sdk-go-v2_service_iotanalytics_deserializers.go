--- vendor/github.com/aws/aws-sdk-go-v2/service/iotanalytics/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotanalytics/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpBatchPutMessage) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpBatchPutMessage) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -211,10 +207,8 @@ func (m *awsRestjson1_deserializeOpCancelPipelineRepro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -226,7 +220,6 @@ func (m *awsRestjson1_deserializeOpCancelPipelineRepro
 	output := &CancelPipelineReprocessingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -311,10 +304,8 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -354,7 +345,6 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -496,10 +486,8 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -539,7 +527,6 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -681,10 +668,8 @@ func (m *awsRestjson1_deserializeOpCreateDatasetConten
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -724,7 +709,6 @@ func (m *awsRestjson1_deserializeOpCreateDatasetConten
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -849,10 +833,8 @@ func (m *awsRestjson1_deserializeOpCreateDatastore) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -892,7 +874,6 @@ func (m *awsRestjson1_deserializeOpCreateDatastore) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1034,10 +1015,8 @@ func (m *awsRestjson1_deserializeOpCreatePipeline) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1077,7 +1056,6 @@ func (m *awsRestjson1_deserializeOpCreatePipeline) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1214,10 +1192,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1235,7 +1211,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1320,10 +1295,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1341,7 +1314,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1426,10 +1398,8 @@ func (m *awsRestjson1_deserializeOpDeleteDatasetConten
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1447,7 +1417,6 @@ func (m *awsRestjson1_deserializeOpDeleteDatasetConten
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1532,10 +1501,8 @@ func (m *awsRestjson1_deserializeOpDeleteDatastore) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1553,7 +1520,6 @@ func (m *awsRestjson1_deserializeOpDeleteDatastore) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1638,10 +1604,8 @@ func (m *awsRestjson1_deserializeOpDeletePipeline) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1659,7 +1623,6 @@ func (m *awsRestjson1_deserializeOpDeletePipeline) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1744,10 +1707,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1787,7 +1748,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1913,10 +1873,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1956,7 +1914,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2077,10 +2034,8 @@ func (m *awsRestjson1_deserializeOpDescribeDatastore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2120,7 +2075,6 @@ func (m *awsRestjson1_deserializeOpDescribeDatastore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2246,10 +2200,8 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingOpti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2289,7 +2241,6 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingOpti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2410,10 +2361,8 @@ func (m *awsRestjson1_deserializeOpDescribePipeline) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2453,7 +2402,6 @@ func (m *awsRestjson1_deserializeOpDescribePipeline) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2574,10 +2522,8 @@ func (m *awsRestjson1_deserializeOpGetDatasetContent) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2617,7 +2563,6 @@ func (m *awsRestjson1_deserializeOpGetDatasetContent) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2759,10 +2704,8 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2802,7 +2745,6 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2929,10 +2871,8 @@ func (m *awsRestjson1_deserializeOpListDatasetContents
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2972,7 +2912,6 @@ func (m *awsRestjson1_deserializeOpListDatasetContents
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3102,10 +3041,8 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3145,7 +3082,6 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3272,10 +3208,8 @@ func (m *awsRestjson1_deserializeOpListDatastores) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3315,7 +3249,6 @@ func (m *awsRestjson1_deserializeOpListDatastores) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3442,10 +3375,8 @@ func (m *awsRestjson1_deserializeOpListPipelines) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3485,7 +3416,6 @@ func (m *awsRestjson1_deserializeOpListPipelines) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3612,10 +3542,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3655,7 +3583,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3779,10 +3706,8 @@ func (m *awsRestjson1_deserializeOpPutLoggingOptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3800,7 +3725,6 @@ func (m *awsRestjson1_deserializeOpPutLoggingOptions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3882,10 +3806,8 @@ func (m *awsRestjson1_deserializeOpRunPipelineActivity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3925,7 +3847,6 @@ func (m *awsRestjson1_deserializeOpRunPipelineActivity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4052,10 +3973,8 @@ func (m *awsRestjson1_deserializeOpSampleChannelData) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4095,7 +4014,6 @@ func (m *awsRestjson1_deserializeOpSampleChannelData) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4216,10 +4134,8 @@ func (m *awsRestjson1_deserializeOpStartPipelineReproc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4259,7 +4175,6 @@ func (m *awsRestjson1_deserializeOpStartPipelineReproc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4387,10 +4302,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4402,7 +4315,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4490,10 +4402,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4505,7 +4415,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4593,10 +4502,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4614,7 +4521,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4699,10 +4605,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4720,7 +4624,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4805,10 +4708,8 @@ func (m *awsRestjson1_deserializeOpUpdateDatastore) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4826,7 +4727,6 @@ func (m *awsRestjson1_deserializeOpUpdateDatastore) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4911,10 +4811,8 @@ func (m *awsRestjson1_deserializeOpUpdatePipeline) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4932,7 +4830,6 @@ func (m *awsRestjson1_deserializeOpUpdatePipeline) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
