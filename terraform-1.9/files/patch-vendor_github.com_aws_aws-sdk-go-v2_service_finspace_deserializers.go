--- vendor/github.com/aws/aws-sdk-go-v2/service/finspace/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/finspace/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -234,10 +230,8 @@ func (m *awsRestjson1_deserializeOpCreateKxChangeset) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -277,7 +271,6 @@ func (m *awsRestjson1_deserializeOpCreateKxChangeset) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -477,10 +470,8 @@ func (m *awsRestjson1_deserializeOpCreateKxCluster) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -520,7 +511,6 @@ func (m *awsRestjson1_deserializeOpCreateKxCluster) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -828,10 +818,8 @@ func (m *awsRestjson1_deserializeOpCreateKxDatabase) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -871,7 +859,6 @@ func (m *awsRestjson1_deserializeOpCreateKxDatabase) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1064,10 +1051,8 @@ func (m *awsRestjson1_deserializeOpCreateKxDataview) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1107,7 +1092,6 @@ func (m *awsRestjson1_deserializeOpCreateKxDataview) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1359,10 +1343,8 @@ func (m *awsRestjson1_deserializeOpCreateKxEnvironment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1402,7 +1384,6 @@ func (m *awsRestjson1_deserializeOpCreateKxEnvironment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1594,10 +1575,8 @@ func (m *awsRestjson1_deserializeOpCreateKxScalingGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1637,7 +1616,6 @@ func (m *awsRestjson1_deserializeOpCreateKxScalingGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1836,10 +1814,8 @@ func (m *awsRestjson1_deserializeOpCreateKxUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1879,7 +1855,6 @@ func (m *awsRestjson1_deserializeOpCreateKxUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2040,10 +2015,8 @@ func (m *awsRestjson1_deserializeOpCreateKxVolume) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2083,7 +2056,6 @@ func (m *awsRestjson1_deserializeOpCreateKxVolume) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2306,10 +2278,8 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2321,7 +2291,6 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironment) 
 	output := &DeleteEnvironmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2406,10 +2375,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxCluster) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2421,7 +2388,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxCluster) Ha
 	output := &DeleteKxClusterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2512,10 +2478,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxClusterNode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2527,7 +2491,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxClusterNode
 	output := &DeleteKxClusterNodeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2612,10 +2575,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxDatabase) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2627,7 +2588,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxDatabase) H
 	output := &DeleteKxDatabaseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2715,10 +2675,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxDataview) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2730,7 +2688,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxDataview) H
 	output := &DeleteKxDataviewOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2818,10 +2775,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxEnvironment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2833,7 +2788,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxEnvironment
 	output := &DeleteKxEnvironmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2921,10 +2875,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxScalingGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2936,7 +2888,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxScalingGrou
 	output := &DeleteKxScalingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3027,10 +2978,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3042,7 +2991,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxUser) Handl
 	output := &DeleteKxUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3130,10 +3078,8 @@ func (m *awsRestjson1_deserializeOpDeleteKxVolume) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3145,7 +3091,6 @@ func (m *awsRestjson1_deserializeOpDeleteKxVolume) Han
 	output := &DeleteKxVolumeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3236,10 +3181,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironment) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3279,7 +3222,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironment) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3397,10 +3339,8 @@ func (m *awsRestjson1_deserializeOpGetKxChangeset) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3440,7 +3380,6 @@ func (m *awsRestjson1_deserializeOpGetKxChangeset) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3650,10 +3589,8 @@ func (m *awsRestjson1_deserializeOpGetKxCluster) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3693,7 +3630,6 @@ func (m *awsRestjson1_deserializeOpGetKxCluster) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3992,10 +3928,8 @@ func (m *awsRestjson1_deserializeOpGetKxConnectionStri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4035,7 +3969,6 @@ func (m *awsRestjson1_deserializeOpGetKxConnectionStri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4160,10 +4093,8 @@ func (m *awsRestjson1_deserializeOpGetKxDatabase) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4203,7 +4134,6 @@ func (m *awsRestjson1_deserializeOpGetKxDatabase) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4435,10 +4365,8 @@ func (m *awsRestjson1_deserializeOpGetKxDataview) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4478,7 +4406,6 @@ func (m *awsRestjson1_deserializeOpGetKxDataview) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4735,10 +4662,8 @@ func (m *awsRestjson1_deserializeOpGetKxEnvironment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4778,7 +4703,6 @@ func (m *awsRestjson1_deserializeOpGetKxEnvironment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5049,10 +4973,8 @@ func (m *awsRestjson1_deserializeOpGetKxScalingGroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5092,7 +5014,6 @@ func (m *awsRestjson1_deserializeOpGetKxScalingGroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5305,10 +5226,8 @@ func (m *awsRestjson1_deserializeOpGetKxUser) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5348,7 +5267,6 @@ func (m *awsRestjson1_deserializeOpGetKxUser) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5500,10 +5418,8 @@ func (m *awsRestjson1_deserializeOpGetKxVolume) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5543,7 +5459,6 @@ func (m *awsRestjson1_deserializeOpGetKxVolume) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5784,10 +5699,8 @@ func (m *awsRestjson1_deserializeOpListEnvironments) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5827,7 +5740,6 @@ func (m *awsRestjson1_deserializeOpListEnvironments) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5951,10 +5863,8 @@ func (m *awsRestjson1_deserializeOpListKxChangesets) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5994,7 +5904,6 @@ func (m *awsRestjson1_deserializeOpListKxChangesets) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6124,10 +6033,8 @@ func (m *awsRestjson1_deserializeOpListKxClusterNodes)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6167,7 +6074,6 @@ func (m *awsRestjson1_deserializeOpListKxClusterNodes)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6300,10 +6206,8 @@ func (m *awsRestjson1_deserializeOpListKxClusters) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6343,7 +6247,6 @@ func (m *awsRestjson1_deserializeOpListKxClusters) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6479,10 +6382,8 @@ func (m *awsRestjson1_deserializeOpListKxDatabases) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6522,7 +6423,6 @@ func (m *awsRestjson1_deserializeOpListKxDatabases) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6652,10 +6552,8 @@ func (m *awsRestjson1_deserializeOpListKxDataviews) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6695,7 +6593,6 @@ func (m *awsRestjson1_deserializeOpListKxDataviews) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6825,10 +6722,8 @@ func (m *awsRestjson1_deserializeOpListKxEnvironments)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6868,7 +6763,6 @@ func (m *awsRestjson1_deserializeOpListKxEnvironments)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6992,10 +6886,8 @@ func (m *awsRestjson1_deserializeOpListKxScalingGroups
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7035,7 +6927,6 @@ func (m *awsRestjson1_deserializeOpListKxScalingGroups
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7171,10 +7062,8 @@ func (m *awsRestjson1_deserializeOpListKxUsers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7214,7 +7103,6 @@ func (m *awsRestjson1_deserializeOpListKxUsers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7344,10 +7232,8 @@ func (m *awsRestjson1_deserializeOpListKxVolumes) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7387,7 +7273,6 @@ func (m *awsRestjson1_deserializeOpListKxVolumes) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7523,10 +7408,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7566,7 +7449,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7681,10 +7563,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7696,7 +7576,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7775,10 +7654,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7790,7 +7667,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7869,10 +7745,8 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7912,7 +7786,6 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8033,10 +7906,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxClusterCode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8048,7 +7919,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxClusterCode
 	output := &UpdateKxClusterCodeConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8139,10 +8009,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxClusterData
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8154,7 +8022,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxClusterData
 	output := &UpdateKxClusterDatabasesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8245,10 +8112,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxDatabase) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8288,7 +8153,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxDatabase) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8450,10 +8314,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxDataview) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8493,7 +8355,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxDataview) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8747,10 +8608,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxEnvironment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8790,7 +8649,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxEnvironment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9055,10 +8913,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxEnvironment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9098,7 +8954,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxEnvironment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9363,10 +9218,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9406,7 +9259,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9564,10 +9416,8 @@ func (m *awsRestjson1_deserializeOpUpdateKxVolume) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9607,7 +9457,6 @@ func (m *awsRestjson1_deserializeOpUpdateKxVolume) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
