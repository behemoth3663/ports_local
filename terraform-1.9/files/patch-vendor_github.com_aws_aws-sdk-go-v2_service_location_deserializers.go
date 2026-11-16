--- vendor/github.com/aws/aws-sdk-go-v2/service/location/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/location/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAssociateTrackerCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -61,7 +58,6 @@ func (m *awsRestjson1_deserializeOpAssociateTrackerCon
 	output := &AssociateTrackerConsumerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -152,10 +148,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDevicePo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -195,7 +189,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDevicePo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -316,10 +309,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteGeofence
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -359,7 +350,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteGeofence
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -480,10 +470,8 @@ func (m *awsRestjson1_deserializeOpBatchEvaluateGeofen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -523,7 +511,6 @@ func (m *awsRestjson1_deserializeOpBatchEvaluateGeofen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -644,10 +631,8 @@ func (m *awsRestjson1_deserializeOpBatchGetDevicePosit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -687,7 +672,6 @@ func (m *awsRestjson1_deserializeOpBatchGetDevicePosit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -813,10 +797,8 @@ func (m *awsRestjson1_deserializeOpBatchPutGeofence) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -856,7 +838,6 @@ func (m *awsRestjson1_deserializeOpBatchPutGeofence) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -982,10 +963,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateDevicePo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1025,7 +1004,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateDevicePo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1146,10 +1124,8 @@ func (m *awsRestjson1_deserializeOpCalculateRoute) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1189,7 +1165,6 @@ func (m *awsRestjson1_deserializeOpCalculateRoute) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1315,10 +1290,8 @@ func (m *awsRestjson1_deserializeOpCalculateRouteMatri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1358,7 +1331,6 @@ func (m *awsRestjson1_deserializeOpCalculateRouteMatri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1494,10 +1466,8 @@ func (m *awsRestjson1_deserializeOpCreateGeofenceColle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1537,7 +1507,6 @@ func (m *awsRestjson1_deserializeOpCreateGeofenceColle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1687,10 +1656,8 @@ func (m *awsRestjson1_deserializeOpCreateKey) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1730,7 +1697,6 @@ func (m *awsRestjson1_deserializeOpCreateKey) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1889,10 +1855,8 @@ func (m *awsRestjson1_deserializeOpCreateMap) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1932,7 +1896,6 @@ func (m *awsRestjson1_deserializeOpCreateMap) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2082,10 +2045,8 @@ func (m *awsRestjson1_deserializeOpCreatePlaceIndex) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2125,7 +2086,6 @@ func (m *awsRestjson1_deserializeOpCreatePlaceIndex) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2275,10 +2235,8 @@ func (m *awsRestjson1_deserializeOpCreateRouteCalculat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2318,7 +2276,6 @@ func (m *awsRestjson1_deserializeOpCreateRouteCalculat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2468,10 +2425,8 @@ func (m *awsRestjson1_deserializeOpCreateTracker) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2511,7 +2466,6 @@ func (m *awsRestjson1_deserializeOpCreateTracker) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2661,10 +2615,8 @@ func (m *awsRestjson1_deserializeOpDeleteGeofenceColle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2676,7 +2628,6 @@ func (m *awsRestjson1_deserializeOpDeleteGeofenceColle
 	output := &DeleteGeofenceCollectionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2761,10 +2712,8 @@ func (m *awsRestjson1_deserializeOpDeleteKey) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2776,7 +2725,6 @@ func (m *awsRestjson1_deserializeOpDeleteKey) HandleDe
 	output := &DeleteKeyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2861,10 +2809,8 @@ func (m *awsRestjson1_deserializeOpDeleteMap) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2876,7 +2822,6 @@ func (m *awsRestjson1_deserializeOpDeleteMap) HandleDe
 	output := &DeleteMapOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2961,10 +2906,8 @@ func (m *awsRestjson1_deserializeOpDeletePlaceIndex) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2976,7 +2919,6 @@ func (m *awsRestjson1_deserializeOpDeletePlaceIndex) H
 	output := &DeletePlaceIndexOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3061,10 +3003,8 @@ func (m *awsRestjson1_deserializeOpDeleteRouteCalculat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3076,7 +3016,6 @@ func (m *awsRestjson1_deserializeOpDeleteRouteCalculat
 	output := &DeleteRouteCalculatorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3161,10 +3100,8 @@ func (m *awsRestjson1_deserializeOpDeleteTracker) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3176,7 +3113,6 @@ func (m *awsRestjson1_deserializeOpDeleteTracker) Hand
 	output := &DeleteTrackerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3261,10 +3197,8 @@ func (m *awsRestjson1_deserializeOpDescribeGeofenceCol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3304,7 +3238,6 @@ func (m *awsRestjson1_deserializeOpDescribeGeofenceCol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3518,10 +3451,8 @@ func (m *awsRestjson1_deserializeOpDescribeKey) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3561,7 +3492,6 @@ func (m *awsRestjson1_deserializeOpDescribeKey) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3762,10 +3692,8 @@ func (m *awsRestjson1_deserializeOpDescribeMap) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3805,7 +3733,6 @@ func (m *awsRestjson1_deserializeOpDescribeMap) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4002,10 +3929,8 @@ func (m *awsRestjson1_deserializeOpDescribePlaceIndex)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4045,7 +3970,6 @@ func (m *awsRestjson1_deserializeOpDescribePlaceIndex)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4242,10 +4166,8 @@ func (m *awsRestjson1_deserializeOpDescribeRouteCalcul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4285,7 +4207,6 @@ func (m *awsRestjson1_deserializeOpDescribeRouteCalcul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4477,10 +4398,8 @@ func (m *awsRestjson1_deserializeOpDescribeTracker) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4520,7 +4439,6 @@ func (m *awsRestjson1_deserializeOpDescribeTracker) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4748,10 +4666,8 @@ func (m *awsRestjson1_deserializeOpDisassociateTracker
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4763,7 +4679,6 @@ func (m *awsRestjson1_deserializeOpDisassociateTracker
 	output := &DisassociateTrackerConsumerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4848,10 +4763,8 @@ func (m *awsRestjson1_deserializeOpForecastGeofenceEve
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4891,7 +4804,6 @@ func (m *awsRestjson1_deserializeOpForecastGeofenceEve
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5039,10 +4951,8 @@ func (m *awsRestjson1_deserializeOpGetDevicePosition) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5082,7 +4992,6 @@ func (m *awsRestjson1_deserializeOpGetDevicePosition) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5248,10 +5157,8 @@ func (m *awsRestjson1_deserializeOpGetDevicePositionHi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5291,7 +5198,6 @@ func (m *awsRestjson1_deserializeOpGetDevicePositionHi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5421,10 +5327,8 @@ func (m *awsRestjson1_deserializeOpGetGeofence) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5464,7 +5368,6 @@ func (m *awsRestjson1_deserializeOpGetGeofence) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5634,10 +5537,8 @@ func (m *awsRestjson1_deserializeOpGetMapGlyphs) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5659,7 +5560,6 @@ func (m *awsRestjson1_deserializeOpGetMapGlyphs) Handl
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5783,10 +5683,8 @@ func (m *awsRestjson1_deserializeOpGetMapSprites) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5808,7 +5706,6 @@ func (m *awsRestjson1_deserializeOpGetMapSprites) Hand
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5932,10 +5829,8 @@ func (m *awsRestjson1_deserializeOpGetMapStyleDescript
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5957,7 +5852,6 @@ func (m *awsRestjson1_deserializeOpGetMapStyleDescript
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6081,10 +5975,8 @@ func (m *awsRestjson1_deserializeOpGetMapTile) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6106,7 +5998,6 @@ func (m *awsRestjson1_deserializeOpGetMapTile) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6230,10 +6121,8 @@ func (m *awsRestjson1_deserializeOpGetPlace) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6273,7 +6162,6 @@ func (m *awsRestjson1_deserializeOpGetPlace) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6394,10 +6282,8 @@ func (m *awsRestjson1_deserializeOpListDevicePositions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6437,7 +6323,6 @@ func (m *awsRestjson1_deserializeOpListDevicePositions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6564,10 +6449,8 @@ func (m *awsRestjson1_deserializeOpListGeofenceCollect
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6607,7 +6490,6 @@ func (m *awsRestjson1_deserializeOpListGeofenceCollect
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6734,10 +6616,8 @@ func (m *awsRestjson1_deserializeOpListGeofences) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6777,7 +6657,6 @@ func (m *awsRestjson1_deserializeOpListGeofences) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6907,10 +6786,8 @@ func (m *awsRestjson1_deserializeOpListKeys) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6950,7 +6827,6 @@ func (m *awsRestjson1_deserializeOpListKeys) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7077,10 +6953,8 @@ func (m *awsRestjson1_deserializeOpListMaps) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7120,7 +6994,6 @@ func (m *awsRestjson1_deserializeOpListMaps) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7247,10 +7120,8 @@ func (m *awsRestjson1_deserializeOpListPlaceIndexes) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7290,7 +7161,6 @@ func (m *awsRestjson1_deserializeOpListPlaceIndexes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7417,10 +7287,8 @@ func (m *awsRestjson1_deserializeOpListRouteCalculator
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7460,7 +7328,6 @@ func (m *awsRestjson1_deserializeOpListRouteCalculator
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7587,10 +7454,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7630,7 +7495,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7751,10 +7615,8 @@ func (m *awsRestjson1_deserializeOpListTrackerConsumer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7794,7 +7656,6 @@ func (m *awsRestjson1_deserializeOpListTrackerConsumer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7924,10 +7785,8 @@ func (m *awsRestjson1_deserializeOpListTrackers) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7967,7 +7826,6 @@ func (m *awsRestjson1_deserializeOpListTrackers) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8094,10 +7952,8 @@ func (m *awsRestjson1_deserializeOpPutGeofence) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8137,7 +7993,6 @@ func (m *awsRestjson1_deserializeOpPutGeofence) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8291,10 +8146,8 @@ func (m *awsRestjson1_deserializeOpSearchPlaceIndexFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8334,7 +8187,6 @@ func (m *awsRestjson1_deserializeOpSearchPlaceIndexFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8460,10 +8312,8 @@ func (m *awsRestjson1_deserializeOpSearchPlaceIndexFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8503,7 +8353,6 @@ func (m *awsRestjson1_deserializeOpSearchPlaceIndexFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8629,10 +8478,8 @@ func (m *awsRestjson1_deserializeOpSearchPlaceIndexFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8672,7 +8519,6 @@ func (m *awsRestjson1_deserializeOpSearchPlaceIndexFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8798,10 +8644,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8813,7 +8657,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8898,10 +8741,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8913,7 +8754,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8998,10 +8838,8 @@ func (m *awsRestjson1_deserializeOpUpdateGeofenceColle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9041,7 +8879,6 @@ func (m *awsRestjson1_deserializeOpUpdateGeofenceColle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9188,10 +9025,8 @@ func (m *awsRestjson1_deserializeOpUpdateKey) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9231,7 +9066,6 @@ func (m *awsRestjson1_deserializeOpUpdateKey) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9378,10 +9212,8 @@ func (m *awsRestjson1_deserializeOpUpdateMap) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9421,7 +9253,6 @@ func (m *awsRestjson1_deserializeOpUpdateMap) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9568,10 +9399,8 @@ func (m *awsRestjson1_deserializeOpUpdatePlaceIndex) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9611,7 +9440,6 @@ func (m *awsRestjson1_deserializeOpUpdatePlaceIndex) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9758,10 +9586,8 @@ func (m *awsRestjson1_deserializeOpUpdateRouteCalculat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9801,7 +9627,6 @@ func (m *awsRestjson1_deserializeOpUpdateRouteCalculat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9948,10 +9773,8 @@ func (m *awsRestjson1_deserializeOpUpdateTracker) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9991,7 +9814,6 @@ func (m *awsRestjson1_deserializeOpUpdateTracker) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10138,10 +9960,8 @@ func (m *awsRestjson1_deserializeOpVerifyDevicePositio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10181,7 +10001,6 @@ func (m *awsRestjson1_deserializeOpVerifyDevicePositio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
