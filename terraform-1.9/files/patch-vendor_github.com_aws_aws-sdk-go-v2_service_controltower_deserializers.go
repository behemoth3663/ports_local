--- vendor/github.com/aws/aws-sdk-go-v2/service/controltower/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/controltower/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpCreateLandingZone) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpCreateLandingZone) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -224,10 +220,8 @@ func (m *awsRestjson1_deserializeOpDeleteLandingZone) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -267,7 +261,6 @@ func (m *awsRestjson1_deserializeOpDeleteLandingZone) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -395,10 +388,8 @@ func (m *awsRestjson1_deserializeOpDisableBaseline) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -438,7 +429,6 @@ func (m *awsRestjson1_deserializeOpDisableBaseline) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -569,10 +559,8 @@ func (m *awsRestjson1_deserializeOpDisableControl) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -612,7 +600,6 @@ func (m *awsRestjson1_deserializeOpDisableControl) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -743,10 +730,8 @@ func (m *awsRestjson1_deserializeOpEnableBaseline) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -786,7 +771,6 @@ func (m *awsRestjson1_deserializeOpEnableBaseline) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -926,10 +910,8 @@ func (m *awsRestjson1_deserializeOpEnableControl) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -969,7 +951,6 @@ func (m *awsRestjson1_deserializeOpEnableControl) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1109,10 +1090,8 @@ func (m *awsRestjson1_deserializeOpGetBaseline) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1152,7 +1131,6 @@ func (m *awsRestjson1_deserializeOpGetBaseline) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1295,10 +1273,8 @@ func (m *awsRestjson1_deserializeOpGetBaselineOperatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1338,7 +1314,6 @@ func (m *awsRestjson1_deserializeOpGetBaselineOperatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1459,10 +1434,8 @@ func (m *awsRestjson1_deserializeOpGetControlOperation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1502,7 +1475,6 @@ func (m *awsRestjson1_deserializeOpGetControlOperation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1623,10 +1595,8 @@ func (m *awsRestjson1_deserializeOpGetEnabledBaseline)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1666,7 +1636,6 @@ func (m *awsRestjson1_deserializeOpGetEnabledBaseline)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1787,10 +1756,8 @@ func (m *awsRestjson1_deserializeOpGetEnabledControl) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1830,7 +1797,6 @@ func (m *awsRestjson1_deserializeOpGetEnabledControl) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1951,10 +1917,8 @@ func (m *awsRestjson1_deserializeOpGetLandingZone) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1994,7 +1958,6 @@ func (m *awsRestjson1_deserializeOpGetLandingZone) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2115,10 +2078,8 @@ func (m *awsRestjson1_deserializeOpGetLandingZoneOpera
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2158,7 +2119,6 @@ func (m *awsRestjson1_deserializeOpGetLandingZoneOpera
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2279,10 +2239,8 @@ func (m *awsRestjson1_deserializeOpListBaselines) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2322,7 +2280,6 @@ func (m *awsRestjson1_deserializeOpListBaselines) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2449,10 +2406,8 @@ func (m *awsRestjson1_deserializeOpListControlOperatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2492,7 +2447,6 @@ func (m *awsRestjson1_deserializeOpListControlOperatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2619,10 +2573,8 @@ func (m *awsRestjson1_deserializeOpListEnabledBaseline
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2662,7 +2614,6 @@ func (m *awsRestjson1_deserializeOpListEnabledBaseline
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2789,10 +2740,8 @@ func (m *awsRestjson1_deserializeOpListEnabledControls
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2832,7 +2781,6 @@ func (m *awsRestjson1_deserializeOpListEnabledControls
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2962,10 +2910,8 @@ func (m *awsRestjson1_deserializeOpListLandingZoneOper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3005,7 +2951,6 @@ func (m *awsRestjson1_deserializeOpListLandingZoneOper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3132,10 +3077,8 @@ func (m *awsRestjson1_deserializeOpListLandingZones) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3175,7 +3118,6 @@ func (m *awsRestjson1_deserializeOpListLandingZones) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3302,10 +3244,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3345,7 +3285,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3460,10 +3399,8 @@ func (m *awsRestjson1_deserializeOpResetEnabledBaselin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3503,7 +3440,6 @@ func (m *awsRestjson1_deserializeOpResetEnabledBaselin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3634,10 +3570,8 @@ func (m *awsRestjson1_deserializeOpResetEnabledControl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3677,7 +3611,6 @@ func (m *awsRestjson1_deserializeOpResetEnabledControl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3808,10 +3741,8 @@ func (m *awsRestjson1_deserializeOpResetLandingZone) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3851,7 +3782,6 @@ func (m *awsRestjson1_deserializeOpResetLandingZone) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3979,10 +3909,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3994,7 +3922,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4073,10 +4000,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4088,7 +4013,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4167,10 +4091,8 @@ func (m *awsRestjson1_deserializeOpUpdateEnabledBaseli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4210,7 +4132,6 @@ func (m *awsRestjson1_deserializeOpUpdateEnabledBaseli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4341,10 +4262,8 @@ func (m *awsRestjson1_deserializeOpUpdateEnabledContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4384,7 +4303,6 @@ func (m *awsRestjson1_deserializeOpUpdateEnabledContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4515,10 +4433,8 @@ func (m *awsRestjson1_deserializeOpUpdateLandingZone) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4558,7 +4474,6 @@ func (m *awsRestjson1_deserializeOpUpdateLandingZone) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
