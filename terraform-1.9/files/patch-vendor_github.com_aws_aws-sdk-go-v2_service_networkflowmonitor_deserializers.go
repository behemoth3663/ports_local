--- vendor/github.com/aws/aws-sdk-go-v2/service/networkflowmonitor/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/networkflowmonitor/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateMonitor) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateMonitor) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -281,10 +277,8 @@ func (m *awsRestjson1_deserializeOpCreateScope) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -324,7 +318,6 @@ func (m *awsRestjson1_deserializeOpCreateScope) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -475,10 +468,8 @@ func (m *awsRestjson1_deserializeOpDeleteMonitor) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -490,7 +481,6 @@ func (m *awsRestjson1_deserializeOpDeleteMonitor) Hand
 	output := &DeleteMonitorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -575,10 +565,8 @@ func (m *awsRestjson1_deserializeOpDeleteScope) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -590,7 +578,6 @@ func (m *awsRestjson1_deserializeOpDeleteScope) Handle
 	output := &DeleteScopeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -681,10 +668,8 @@ func (m *awsRestjson1_deserializeOpGetMonitor) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -724,7 +709,6 @@ func (m *awsRestjson1_deserializeOpGetMonitor) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -914,10 +898,8 @@ func (m *awsRestjson1_deserializeOpGetQueryResultsMoni
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -957,7 +939,6 @@ func (m *awsRestjson1_deserializeOpGetQueryResultsMoni
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1099,10 +1080,8 @@ func (m *awsRestjson1_deserializeOpGetQueryResultsWork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1142,7 +1121,6 @@ func (m *awsRestjson1_deserializeOpGetQueryResultsWork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1275,10 +1253,8 @@ func (m *awsRestjson1_deserializeOpGetQueryResultsWork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1318,7 +1294,6 @@ func (m *awsRestjson1_deserializeOpGetQueryResultsWork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1460,10 +1435,8 @@ func (m *awsRestjson1_deserializeOpGetQueryStatusMonit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1503,7 +1476,6 @@ func (m *awsRestjson1_deserializeOpGetQueryStatusMonit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1628,10 +1600,8 @@ func (m *awsRestjson1_deserializeOpGetQueryStatusWorkl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1671,7 +1641,6 @@ func (m *awsRestjson1_deserializeOpGetQueryStatusWorkl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1796,10 +1765,8 @@ func (m *awsRestjson1_deserializeOpGetQueryStatusWorkl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1839,7 +1806,6 @@ func (m *awsRestjson1_deserializeOpGetQueryStatusWorkl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1964,10 +1930,8 @@ func (m *awsRestjson1_deserializeOpGetScope) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2007,7 +1971,6 @@ func (m *awsRestjson1_deserializeOpGetScope) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2163,10 +2126,8 @@ func (m *awsRestjson1_deserializeOpListMonitors) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2206,7 +2167,6 @@ func (m *awsRestjson1_deserializeOpListMonitors) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2333,10 +2293,8 @@ func (m *awsRestjson1_deserializeOpListScopes) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2376,7 +2334,6 @@ func (m *awsRestjson1_deserializeOpListScopes) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2506,10 +2463,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2549,7 +2504,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2673,10 +2627,8 @@ func (m *awsRestjson1_deserializeOpStartQueryMonitorTo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2716,7 +2668,6 @@ func (m *awsRestjson1_deserializeOpStartQueryMonitorTo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2841,10 +2792,8 @@ func (m *awsRestjson1_deserializeOpStartQueryWorkloadI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2884,7 +2833,6 @@ func (m *awsRestjson1_deserializeOpStartQueryWorkloadI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3009,10 +2957,8 @@ func (m *awsRestjson1_deserializeOpStartQueryWorkloadI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3052,7 +2998,6 @@ func (m *awsRestjson1_deserializeOpStartQueryWorkloadI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3177,10 +3122,8 @@ func (m *awsRestjson1_deserializeOpStopQueryMonitorTop
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3192,7 +3135,6 @@ func (m *awsRestjson1_deserializeOpStopQueryMonitorTop
 	output := &StopQueryMonitorTopContributorsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3277,10 +3219,8 @@ func (m *awsRestjson1_deserializeOpStopQueryWorkloadIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3292,7 +3232,6 @@ func (m *awsRestjson1_deserializeOpStopQueryWorkloadIn
 	output := &StopQueryWorkloadInsightsTopContributorsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3377,10 +3316,8 @@ func (m *awsRestjson1_deserializeOpStopQueryWorkloadIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3392,7 +3329,6 @@ func (m *awsRestjson1_deserializeOpStopQueryWorkloadIn
 	output := &StopQueryWorkloadInsightsTopContributorsDataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3477,10 +3413,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3492,7 +3426,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3580,10 +3513,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3595,7 +3526,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3683,10 +3613,8 @@ func (m *awsRestjson1_deserializeOpUpdateMonitor) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3726,7 +3654,6 @@ func (m *awsRestjson1_deserializeOpUpdateMonitor) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3916,10 +3843,8 @@ func (m *awsRestjson1_deserializeOpUpdateScope) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3959,7 +3884,6 @@ func (m *awsRestjson1_deserializeOpUpdateScope) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
