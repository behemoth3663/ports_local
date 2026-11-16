--- vendor/github.com/aws/aws-sdk-go-v2/service/evidently/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/evidently/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchEvaluateFeatur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchEvaluateFeatur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -206,10 +202,8 @@ func (m *awsRestjson1_deserializeOpCreateExperiment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -249,7 +243,6 @@ func (m *awsRestjson1_deserializeOpCreateExperiment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -370,10 +363,8 @@ func (m *awsRestjson1_deserializeOpCreateFeature) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -413,7 +404,6 @@ func (m *awsRestjson1_deserializeOpCreateFeature) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -534,10 +524,8 @@ func (m *awsRestjson1_deserializeOpCreateLaunch) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -577,7 +565,6 @@ func (m *awsRestjson1_deserializeOpCreateLaunch) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -698,10 +685,8 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -741,7 +726,6 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -859,10 +843,8 @@ func (m *awsRestjson1_deserializeOpCreateSegment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -902,7 +884,6 @@ func (m *awsRestjson1_deserializeOpCreateSegment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1020,10 +1001,8 @@ func (m *awsRestjson1_deserializeOpDeleteExperiment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1035,7 +1014,6 @@ func (m *awsRestjson1_deserializeOpDeleteExperiment) H
 	output := &DeleteExperimentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1123,10 +1101,8 @@ func (m *awsRestjson1_deserializeOpDeleteFeature) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1138,7 +1114,6 @@ func (m *awsRestjson1_deserializeOpDeleteFeature) Hand
 	output := &DeleteFeatureOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1223,10 +1198,8 @@ func (m *awsRestjson1_deserializeOpDeleteLaunch) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1238,7 +1211,6 @@ func (m *awsRestjson1_deserializeOpDeleteLaunch) Handl
 	output := &DeleteLaunchOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1323,10 +1295,8 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1338,7 +1308,6 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 	output := &DeleteProjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1423,10 +1392,8 @@ func (m *awsRestjson1_deserializeOpDeleteSegment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1438,7 +1405,6 @@ func (m *awsRestjson1_deserializeOpDeleteSegment) Hand
 	output := &DeleteSegmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1523,10 +1489,8 @@ func (m *awsRestjson1_deserializeOpEvaluateFeature) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1566,7 +1530,6 @@ func (m *awsRestjson1_deserializeOpEvaluateFeature) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1711,10 +1674,8 @@ func (m *awsRestjson1_deserializeOpGetExperiment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1754,7 +1715,6 @@ func (m *awsRestjson1_deserializeOpGetExperiment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1872,10 +1832,8 @@ func (m *awsRestjson1_deserializeOpGetExperimentResult
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1915,7 +1873,6 @@ func (m *awsRestjson1_deserializeOpGetExperimentResult
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2055,10 +2012,8 @@ func (m *awsRestjson1_deserializeOpGetFeature) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2098,7 +2053,6 @@ func (m *awsRestjson1_deserializeOpGetFeature) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2216,10 +2170,8 @@ func (m *awsRestjson1_deserializeOpGetLaunch) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2259,7 +2211,6 @@ func (m *awsRestjson1_deserializeOpGetLaunch) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2377,10 +2328,8 @@ func (m *awsRestjson1_deserializeOpGetProject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2420,7 +2369,6 @@ func (m *awsRestjson1_deserializeOpGetProject) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2538,10 +2486,8 @@ func (m *awsRestjson1_deserializeOpGetSegment) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2581,7 +2527,6 @@ func (m *awsRestjson1_deserializeOpGetSegment) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2699,10 +2644,8 @@ func (m *awsRestjson1_deserializeOpListExperiments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2742,7 +2685,6 @@ func (m *awsRestjson1_deserializeOpListExperiments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2866,10 +2808,8 @@ func (m *awsRestjson1_deserializeOpListFeatures) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2909,7 +2849,6 @@ func (m *awsRestjson1_deserializeOpListFeatures) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3036,10 +2975,8 @@ func (m *awsRestjson1_deserializeOpListLaunches) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3079,7 +3016,6 @@ func (m *awsRestjson1_deserializeOpListLaunches) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3203,10 +3139,8 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3246,7 +3180,6 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3370,10 +3303,8 @@ func (m *awsRestjson1_deserializeOpListSegmentReferenc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3413,7 +3344,6 @@ func (m *awsRestjson1_deserializeOpListSegmentReferenc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3540,10 +3470,8 @@ func (m *awsRestjson1_deserializeOpListSegments) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3583,7 +3511,6 @@ func (m *awsRestjson1_deserializeOpListSegments) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3707,10 +3634,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3750,7 +3675,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3865,10 +3789,8 @@ func (m *awsRestjson1_deserializeOpPutProjectEvents) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3908,7 +3830,6 @@ func (m *awsRestjson1_deserializeOpPutProjectEvents) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4039,10 +3960,8 @@ func (m *awsRestjson1_deserializeOpStartExperiment) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4082,7 +4001,6 @@ func (m *awsRestjson1_deserializeOpStartExperiment) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4217,10 +4135,8 @@ func (m *awsRestjson1_deserializeOpStartLaunch) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4260,7 +4176,6 @@ func (m *awsRestjson1_deserializeOpStartLaunch) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4384,10 +4299,8 @@ func (m *awsRestjson1_deserializeOpStopExperiment) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4427,7 +4340,6 @@ func (m *awsRestjson1_deserializeOpStopExperiment) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4562,10 +4474,8 @@ func (m *awsRestjson1_deserializeOpStopLaunch) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4605,7 +4515,6 @@ func (m *awsRestjson1_deserializeOpStopLaunch) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4734,10 +4643,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4749,7 +4656,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4828,10 +4734,8 @@ func (m *awsRestjson1_deserializeOpTestSegmentPattern)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4871,7 +4775,6 @@ func (m *awsRestjson1_deserializeOpTestSegmentPattern)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4990,10 +4893,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5005,7 +4906,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5084,10 +4984,8 @@ func (m *awsRestjson1_deserializeOpUpdateExperiment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5127,7 +5025,6 @@ func (m *awsRestjson1_deserializeOpUpdateExperiment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5245,10 +5142,8 @@ func (m *awsRestjson1_deserializeOpUpdateFeature) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5288,7 +5183,6 @@ func (m *awsRestjson1_deserializeOpUpdateFeature) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5409,10 +5303,8 @@ func (m *awsRestjson1_deserializeOpUpdateLaunch) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5452,7 +5344,6 @@ func (m *awsRestjson1_deserializeOpUpdateLaunch) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5570,10 +5461,8 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5613,7 +5502,6 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5734,10 +5622,8 @@ func (m *awsRestjson1_deserializeOpUpdateProjectDataDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5777,7 +5663,6 @@ func (m *awsRestjson1_deserializeOpUpdateProjectDataDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
