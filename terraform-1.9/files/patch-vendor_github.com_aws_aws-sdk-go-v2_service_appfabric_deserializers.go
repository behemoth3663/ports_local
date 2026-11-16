--- vendor/github.com/aws/aws-sdk-go-v2/service/appfabric/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/appfabric/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchGetUserAccessT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchGetUserAccessT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -209,10 +205,8 @@ func (m *awsRestjson1_deserializeOpConnectAppAuthoriza
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -252,7 +246,6 @@ func (m *awsRestjson1_deserializeOpConnectAppAuthoriza
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -373,10 +366,8 @@ func (m *awsRestjson1_deserializeOpCreateAppAuthorizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -416,7 +407,6 @@ func (m *awsRestjson1_deserializeOpCreateAppAuthorizat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -543,10 +533,8 @@ func (m *awsRestjson1_deserializeOpCreateAppBundle) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -586,7 +574,6 @@ func (m *awsRestjson1_deserializeOpCreateAppBundle) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -710,10 +697,8 @@ func (m *awsRestjson1_deserializeOpCreateIngestion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -753,7 +738,6 @@ func (m *awsRestjson1_deserializeOpCreateIngestion) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -877,10 +861,8 @@ func (m *awsRestjson1_deserializeOpCreateIngestionDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -920,7 +902,6 @@ func (m *awsRestjson1_deserializeOpCreateIngestionDest
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1044,10 +1025,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppAuthorizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1059,7 +1038,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppAuthorizat
 	output := &DeleteAppAuthorizationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1144,10 +1122,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppBundle) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1159,7 +1135,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppBundle) Ha
 	output := &DeleteAppBundleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1244,10 +1219,8 @@ func (m *awsRestjson1_deserializeOpDeleteIngestion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1259,7 +1232,6 @@ func (m *awsRestjson1_deserializeOpDeleteIngestion) Ha
 	output := &DeleteIngestionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1344,10 +1316,8 @@ func (m *awsRestjson1_deserializeOpDeleteIngestionDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1359,7 +1329,6 @@ func (m *awsRestjson1_deserializeOpDeleteIngestionDest
 	output := &DeleteIngestionDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1444,10 +1413,8 @@ func (m *awsRestjson1_deserializeOpGetAppAuthorization
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1487,7 +1454,6 @@ func (m *awsRestjson1_deserializeOpGetAppAuthorization
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1608,10 +1574,8 @@ func (m *awsRestjson1_deserializeOpGetAppBundle) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1651,7 +1615,6 @@ func (m *awsRestjson1_deserializeOpGetAppBundle) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1772,10 +1735,8 @@ func (m *awsRestjson1_deserializeOpGetIngestion) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1815,7 +1776,6 @@ func (m *awsRestjson1_deserializeOpGetIngestion) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1936,10 +1896,8 @@ func (m *awsRestjson1_deserializeOpGetIngestionDestina
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1979,7 +1937,6 @@ func (m *awsRestjson1_deserializeOpGetIngestionDestina
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2100,10 +2057,8 @@ func (m *awsRestjson1_deserializeOpListAppAuthorizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2143,7 +2098,6 @@ func (m *awsRestjson1_deserializeOpListAppAuthorizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2273,10 +2227,8 @@ func (m *awsRestjson1_deserializeOpListAppBundles) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2316,7 +2268,6 @@ func (m *awsRestjson1_deserializeOpListAppBundles) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2443,10 +2394,8 @@ func (m *awsRestjson1_deserializeOpListIngestionDestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2486,7 +2435,6 @@ func (m *awsRestjson1_deserializeOpListIngestionDestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2616,10 +2564,8 @@ func (m *awsRestjson1_deserializeOpListIngestions) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2659,7 +2605,6 @@ func (m *awsRestjson1_deserializeOpListIngestions) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2789,10 +2734,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2832,7 +2775,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2953,10 +2895,8 @@ func (m *awsRestjson1_deserializeOpStartIngestion) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2968,7 +2908,6 @@ func (m *awsRestjson1_deserializeOpStartIngestion) Han
 	output := &StartIngestionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3056,10 +2995,8 @@ func (m *awsRestjson1_deserializeOpStartUserAccessTask
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3099,7 +3036,6 @@ func (m *awsRestjson1_deserializeOpStartUserAccessTask
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3220,10 +3156,8 @@ func (m *awsRestjson1_deserializeOpStopIngestion) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3235,7 +3169,6 @@ func (m *awsRestjson1_deserializeOpStopIngestion) Hand
 	output := &StopIngestionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3323,10 +3256,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3338,7 +3269,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3423,10 +3353,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3438,7 +3366,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3523,10 +3450,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppAuthorizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3566,7 +3491,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppAuthorizat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3687,10 +3611,8 @@ func (m *awsRestjson1_deserializeOpUpdateIngestionDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3730,7 +3652,6 @@ func (m *awsRestjson1_deserializeOpUpdateIngestionDest
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
